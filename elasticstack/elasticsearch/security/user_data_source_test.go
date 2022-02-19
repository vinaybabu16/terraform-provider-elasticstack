package security_test

import (
	"github.com/elastic/terraform-provider-elasticstack/elasticstack/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSecurityUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecurityUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_security_user.test", "username", "elastic"),
					resource.TestCheckTypeSetElemAttr("data.elasticstack_elasticsearch_security_user.test", "roles.*", "superuser"),
				),
			},
		},
	})
}

const testAccDataSourceSecurityUser = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_security_user" "test" {
  username = "elastic"
}
`
