package saltstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidClientWithAllRequiredConfig_pam(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		Username: "username",
		Password: "password",
		Eauth:    "pam",
	}

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestValidClientWithAllRequiredConfig_file(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		Username: "username",
		Password: "password",
		Eauth:    "file",
	}

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestValidClientWithAllRequiredConfig_sharedsecret(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		Username: "username",
		Password: "password",
		Eauth:    "sharedsecret",
	}

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestValidClientWithAllRequiredConfig_token(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		UseToken: true,
		Eauth:    "pam",
		Token:    "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	client, err := NewClient(config)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestNotValidClientWithNoHostProvided(t *testing.T) {
	config := Config{
		Port:     8000,
		Username: "username",
		Password: "password",
		Eauth:    "pam",
	}

	_, err := NewClient(config)
	assert.Error(t, err)
}

func TestNotValidClientWithNoUsernameProvided(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		Password: "password",
		Eauth:    "pam",
	}

	_, err := NewClient(config)
	assert.Error(t, err)
}

func TestNotValidClientWithNoPasswordProvided(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		Username: "username",
		Eauth:    "pam",
	}

	_, err := NewClient(config)
	assert.Error(t, err)
}

func TestNotValidClientWithNoTokenProvided(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8000,
		UseToken: true,
		Eauth:    "pam",
	}

	_, err := NewClient(config)
	assert.Error(t, err)
}
