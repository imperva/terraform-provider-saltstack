package saltstack

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"saltstack": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("SALTSTACK_HOST"); err == "" {
		t.Fatal("SALTSTACK_HOST must be set for acceptance tests")
	}
	if err := os.Getenv("SALTSTACK_PORT"); err == "" {
		t.Fatal("SALTSTACK_PORT must be set for acceptance tests")
	}
	if err := os.Getenv("SALTSTACK_USERNAME"); err == "" {
		t.Fatal("SALTSTACK_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("SALTSTACK_PASSWORD"); err == "" {
		t.Fatal("SALTSTACK_PASSWORD must be set for acceptance tests")
	}
}
