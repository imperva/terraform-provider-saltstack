package saltstack

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/imperva/terraform-provider-saltstack/helper"
)

func TestAccSaltstackMinionKeyPair_basic(t *testing.T) {
	minionId := "test-1.domain.com"
	keySize := 2048
	resourceName := "saltstack_minion_key_pair.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSaltstackMinionKeyPairDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSaltstackMinionKeyPairConfigBasic(minionId, keySize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSaltstackMinionKeyPairExists(resourceName),
					testAccCheckSaltstackMinionPrivateKey(resourceName),
					testAccCheckSaltstackMinionPublicKey(resourceName),
				),
			},
		},
	})
}

func TestAccSaltstackMinionKeyPair_overwrite(t *testing.T) {
	minionId := "test-1.domain.com"
	keySize := 2048
	resourceName := "saltstack_minion_key_pair.test"
	resourceNameForced := "saltstack_minion_key_pair.forced"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSaltstackMinionKeyPairDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSaltstackMinionKeyPairConfigBasic(minionId, keySize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSaltstackMinionKeyPairExists(resourceName),
					testAccCheckSaltstackMinionPrivateKey(resourceName),
					testAccCheckSaltstackMinionPublicKey(resourceName),
				),
			},
			{
				Config: strings.Join([]string{
					testAccCheckSaltstackMinionKeyPairConfigBasic(minionId, keySize),
					testAccCheckSaltstackMinionKeyPairConfigForced(minionId, keySize, true),
				}, "\n"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSaltstackMinionKeyPairExists(resourceNameForced),
					testAccCheckSaltstackMinionPrivateKey(resourceNameForced),
					testAccCheckSaltstackMinionPublicKey(resourceNameForced),
				),
			},
		},
	})
}

func testAccCheckSaltstackMinionKeyPairConfigBasic(minion_id string, key_size int) string {
	return fmt.Sprintf(`
	resource saltstack_minion_key_pair test {
		minion_id = "%s"
		key_size = %d
	}
	`, minion_id, key_size)
}

func testAccCheckSaltstackMinionKeyPairConfigForced(minion_id string, key_size int, force bool) string {
	return fmt.Sprintf(`
    resource saltstack_minion_key_pair forced {
		minion_id = "%s"
		key_size = %d
		force = %t
	}
	`, minion_id, key_size, force)
}

func testAccCheckSaltstackMinionKeyPairExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No minion_id set")
		}

		return nil
	}
}

func testAccCheckSaltstackMinionPrivateKey(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		privateKey, ok := rs.Primary.Attributes["private_key"]
		if !ok {
			return fmt.Errorf("No private_key found in resource %s", resourceName)
		}

		err := helper.ValidateRsaPrivateKey(privateKey)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckSaltstackMinionPublicKey(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		publicKey, ok := rs.Primary.Attributes["public_key"]
		if !ok {
			return fmt.Errorf("No public_key found in resource %s", resourceName)
		}

		err := helper.ValidateRsaPublicKey(publicKey)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckSaltstackMinionKeyPairDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "saltstack_minion_key_pair" {
			continue
		}

		minionId := rs.Primary.ID

		data := map[string]interface{}{
			"client": "wheel",
			"fun":    "key.delete",
			"match":  minionId,
		}

		_, err := c.Post("/run", data)
		if err != nil {
			return err
		}
	}

	return nil
}
