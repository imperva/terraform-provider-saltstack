package saltstack

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type KeyPairCreateResult struct {
	Return []struct {
		Data struct {
			Return map[string]string `json:"return"`
		} `json:"data"`
	} `json:"return"`
}

type KeyPairReadResult struct {
	Return []struct {
		Data struct {
			Return struct {
				Minions map[string]string `json:"minions"`
			} `json:"return"`
		} `json:"data"`
	} `json:"return"`
}

func resourceMinionAcceptedKeyPair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMinionAcceptedKeyPairCreate,
		ReadContext:   resourceMinionAcceptedKeyPairRead,
		DeleteContext: resourceMinionAcceptedKeyPairDelete,
		Schema: map[string]*schema.Schema{
			"minion_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of SaltStack minion.",
				ForceNew:    true,
				ValidateDiagFunc: func(v any, p cty.Path) diag.Diagnostics {
					value := v.(string)
					var diags diag.Diagnostics
					if matched, _ := regexp.MatchString("^[a-zA-Z0-9.-]+$", value); !matched {
						diag := diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Wrong value",
							Detail:   fmt.Sprintf("minion_id must be a valid RFC1123 hostname, the value %s is wrong.", value),
						}
						diags = append(diags, diag)
					}
					return diags
				},
			},
			"key_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The size of the key pair to generate. The size must be 2048, which is the default, or greater. If set to a value less than 2048, the key size will be rounded up to 2048.",
				Default:     2048,
				ForceNew:    true,
			},
			"private_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Minion's private key.",
				Sensitive:   true,
			},
			"public_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Minion's public key.",
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMinionAcceptedKeyPairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	api := m.(*Client)

	minionId := d.Get("minion_id").(string)
	keySize := d.Get("key_size").(int)

	reqData := map[string]interface{}{
		"client":  "wheel",
		"fun":     "key.gen_accept",
		"id_":     minionId,
		"keysize": keySize,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating key pair for minion %s", minionId), nil)
	resp, err := api.Post("/run", reqData)
	if err != nil {
		return diag.FromErr(err)
	}

	var rd KeyPairCreateResult

	err = parseResponseBody(resp, &rd)
	if err != nil {
		return diag.FromErr(err)
	}

	if key_pair := rd.Return[0].Data.Return; len(key_pair) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("The minion %s is already in use.", minionId),
		})
		return diags
	}
	tflog.Debug(ctx, fmt.Sprintf("Created key pair for minion %s", minionId), nil)
	d.Set("public_key", rd.Return[0].Data.Return["pub"])
	d.Set("private_key", rd.Return[0].Data.Return["priv"])
	d.SetId(minionId)

	return resourceMinionAcceptedKeyPairRead(ctx, d, m)
}

func resourceMinionAcceptedKeyPairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.IsNewResource() {
		return diags
	}

	api := m.(*Client)

	minionId := d.Get("minion_id").(string)

	data := map[string]interface{}{
		"client": "wheel",
		"fun":    "key.print",
		"match":  minionId,
	}

	resp, err := api.Post("/run", data)
	if err != nil {
		return diag.FromErr(err)
	}

	var rd KeyPairReadResult

	err = parseResponseBody(resp, &rd)
	if err != nil {
		return diag.FromErr(err)
	}

	if pub_key, ok := rd.Return[0].Data.Return.Minions[minionId]; ok {
		d.Set("public_key", pub_key)
	} else {
		d.SetId("")
	}

	return diags
}

func resourceMinionAcceptedKeyPairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	api := m.(*Client)

	minionId := d.Get("minion_id").(string)

	data := map[string]interface{}{
		"client": "wheel",
		"fun":    "key.delete",
		"match":  minionId,
	}
	tflog.Debug(ctx, fmt.Sprintf("Deleting key pair for minion %s", minionId), nil)
	_, err := api.Post("/run", data)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("Deleted key pair for minion %s", minionId), nil)
	return diags
}

func parseResponseBody(resp *http.Response, extractedData interface{}) error {
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, extractedData)
	if err != nil {
		return err
	}

	return nil
}
