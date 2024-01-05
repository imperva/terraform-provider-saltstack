package saltstack

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Host          string `validate:"required"`
	Port          int
	Username      string `validate:"required_if=UseToken false"`
	Password      string `validate:"required_if=UseToken false"`
	Debug         bool
	SSLSkipVerify bool
	Eauth         string
	Scheme        string
	UseToken      bool
	Token         string `validate:"required_if=UseToken true"`
}

type Client struct {
	Config       Config
	Client       *http.Client
	sessionToken string
}

type LoginReadResult struct {
	Return []struct {
		Token  string   `json:"token"`
		Expire string   `json:"expire"`
		Start  string   `json:"start"`
		User   string   `json:"user"`
		Eauth  string   `json:"eauth"`
		Perms  []string `json:"perms"`
	} `json:"return"`
}

func NewClient(config Config) (*Client, error) {
	supportedAuthTypesKeys := [...]string{"pam", "file", "sharedsecret"}

	supportedAuthTypes := map[string]bool{}
	for _, t := range supportedAuthTypesKeys {
		supportedAuthTypes[t] = true
	}

	if !supportedAuthTypes[config.Eauth] {
		return nil, fmt.Errorf("the Eauth type %s is not supported. The valid types are: %v", config.Eauth, strings.Join(supportedAuthTypesKeys[:], ", "))
	}

	if config.Port == 0 {
		config.Port = 8000
	}

	if config.Scheme == "" {
		config.Scheme = "https"
	}

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return nil, err.(validator.ValidationErrors)
	}

	c := Client{Config: config}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.SSLSkipVerify,
		},
	}
	c.Client = &http.Client{Timeout: 60 * time.Second, Transport: tr}

	return &c, nil
}

func (c *Client) Login() error {
	if c.sessionToken != "" {
		// We are already logged in and already have a session token
		return nil
	}

	if c.Config.UseToken {
		return fmt.Errorf("Unable to login when configured to use a token. Should configure username/password when trying to login")
	}

	reqData := map[string]interface{}{
		"username": c.Config.Username,
		"password": c.Config.Password,
		"eauth":    c.Config.Eauth,
	}

	reqBody, err := convertToJSONString(reqData)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s://%s:%s/login", c.Config.Scheme, c.Config.Host, strconv.Itoa(c.Config.Port))
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}

	var rd LoginReadResult
	err = parseResponseBody(resp, &rd)
	if err != nil {
		return err
	}

	if len(rd.Return) == 0 {
		return fmt.Errorf("Empty return from API while trying to login")
	}
	c.sessionToken = rd.Return[0].Token
	return nil
}

func (c *Client) getSessionToken() (string, error) {
	if c.sessionToken == "" {
		// We need to login
		if err := c.Login(); err != nil {
			return "", err
		}
	}

	return c.sessionToken, nil
}

func (c *Client) Post(uri string, data map[string]interface{}) (*http.Response, error) {
	reqData := make(map[string]interface{})

	url := fmt.Sprintf("%s://%s:%s%s", c.Config.Scheme, c.Config.Host, strconv.Itoa(c.Config.Port), uri)

	if c.Config.UseToken {
		if uri == "/run" {
			reqData["token"] = c.Config.Token
		}
	} else {
		reqData["username"] = c.Config.Username
		reqData["password"] = c.Config.Password
		reqData["eauth"] = c.Config.Eauth
	}

	for k, v := range data {
		reqData[k] = v
	}

	reqBody, err := convertToJSONString(reqData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.Config.UseToken && uri != "/run" {
		var sessionToken string
		if sessionToken, err = c.getSessionToken(); err != nil {
			return nil, err
		}
		req.Header.Set("X-Auth-Token", sessionToken)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}

func convertToJSONString(data map[string]interface{}) (string, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
