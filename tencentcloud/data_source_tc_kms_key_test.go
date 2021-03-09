package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTencentCloudKmsKeyDataSource(t *testing.T) {
	dataSourceName := "data.tencentcloud_kms_key.test"
	rName := fmt.Sprintf("tf-testacc-kms-key-%s", acctest.RandString(13))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKmsKeyConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceKmsKeyCheck(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.key_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.create_time"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.description"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.key_state"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.key_usage"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.creator_uin"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.key_rotation_enabled"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.owner"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.next_rotate_time"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.origin"),
					resource.TestCheckResourceAttrSet(dataSourceName, "key_list.0.valid_to"),
				),
			},
		},
	})
}

func testAccDataSourceKmsKeyCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		return nil
	}
}

func testAccDataSourceKmsKeyConfig(rName string) string {
	return fmt.Sprintf(`
resource "tencentcloud_kms_key" "test" {
  	alias = %[1]q
	description = %[1]q
  	key_state = "Disabled"
	key_rotation_enabled = true
}
data "tencentcloud_kms_key" "test" {
  search_key_alias = tencentcloud_kms_key.test.alias
}
`, rName)
}
