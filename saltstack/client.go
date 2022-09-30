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
	Config Config
	Client *http.Client
}

func NewClient(config Config) (*Client, error) {
	supportedAuthTypesKeys := [...]string{"pam", "sharedsecret"}

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

func (c *Client) Post(uri string, data map[string]interface{}) (*http.Response, error) {
	reqData := make(map[string]interface{})

	url := fmt.Sprintf("%s://%s:%s%s", c.Config.Scheme, c.Config.Host, strconv.Itoa(c.Config.Port), uri)

	if !c.Config.UseToken {
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
	if c.Config.UseToken {
		req.Header.Set("X-Auth-Token", c.Config.Token)
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
