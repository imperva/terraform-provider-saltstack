package saltstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_HOST", nil),
				Description: "Salt Master hostname.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_PORT", nil),
				Description: "Salt Master API port.",
			},
			"scheme": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_SCHEME", "https"),
				Description: "Connection scheme. Can be http or https. Defaults to `https`.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_USERNAME", nil),
				Description: "Salt Master API username.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_PASSWORD", nil),
				Description: "Salt Master API password.",
			},
			"eauth": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_EAUTH", "pam"),
				Description: "Salt Master API External Authentication system. Currently supports: `pam`, `sharedsecret`. Reference: https://docs.saltproject.io/en/latest/topics/eauth/index.html. Defaults to `pam`",
			},
			"use_token": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_USE_TOKEN", false),
				Description: "Whether or not to use token authentication. Reference: https://docs.saltproject.io/en/latest/topics/eauth/index.html#tokens. Defaults to `false`",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_TOKEN", nil),
				Description: "Authentication token if `use_token` is true.",
			},
			"ssl_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_SSL_SKIP_VERIFY", false),
				Description: "Skip SSL verification. Defaults to `false`",
			},
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SALTSTACK_DEBUG", false),
				Description: "Run provider in DEBUG mode. Defaults to `false`",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"saltstack_minion_key_pair": resourceMinionAcceptedKeyPair(),
		},
		ConfigureContextFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	config := Config{
		Host:          d.Get("host").(string),
		Port:          d.Get("port").(int),
		Scheme:        d.Get("scheme").(string),
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		Token:         d.Get("token").(string),
		Eauth:         d.Get("eauth").(string),
		UseToken:      d.Get("use_token").(bool),
		Debug:         d.Get("debug").(bool),
		SSLSkipVerify: d.Get("ssl_skip_verify").(bool),
	}

	var diags diag.Diagnostics

	c, err := NewClient(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
