package alicloud

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAlicloudAcmBasic(t *testing.T) {

	// resourceID := "alicloud_acm_configuration.default"
	// name := fmt.Sprintf("tf-testacc-object-%d", rand)
	// rand := acctest.RandInt()
	// testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceOssBucketObjectConfigDependence)

	// resource.Test(t, resource.TestCase{
	// 	PreCheck: func() {
	// 		testAccPreCheck(t)
	// 	},

	// 	// module name
	// 	IDRefreshName: resourceID,
	// 	Providers:     testAccProviders,
	// 	CheckDestroy:  testAccCheckAcmConfigDestroy,
	// 	Steps: []resource.TestStep{
	// 		{
	// 			Config: testAccCheckAcmConfigBasic(rand),
	// 			Check: resource.ComposeTestCheckFunc(
	// 				testAccCheck(map[string]string{
	// 					"data_id": fmt.Sprintf("dataid-%d", rand),
	// 					"group":   "tf-testAccCenConfigDescription",
	// 				}),
	// 			),
	// 		},
	// 	},
	// })
	rand := acctest.RandInt()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAcmConfigBasic(rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("alicloud_acm_configuration.default", "data_id", fmt.Sprintf("dataid-%d", rand)),
					resource.TestCheckResourceAttr("alicloud_acm_configuration.default", "group", fmt.Sprintf("group-%d", rand)),
					resource.TestCheckResourceAttr("alicloud_acm_configuration.default", "content", "content"),
				),
			},
		},
	})

}

func TestAcmPublishService(t *testing.T) {
	status, err := acmPublishConfigurations("19f267f0-d8b2-4139-a842-b20fe7a3676a", "dataid-tf", "group-tf", "content", "LTAIgL9i0txm1XFs", "hZd8GJJMjsvHrqIYctxJjIj6iY9vhg", "")
	log.Println(status)
	if err != nil {
		log.Printf(err.Error())
	}
	// tenant string, dataID string, group string, content string, ak string, sk string, token string
}

func testAccCheckAcmConfigBasic(rand int) string {
	return fmt.Sprintf(
		`
resource "alicloud_acm_configuration" "default" {
	data_id = "dataid-%d"
	group = "group-%d"
	namespace = "19f267f0-d8b2-4139-a842-b20fe7a3676a"
	content = "content"
}
`, rand, rand)
}
