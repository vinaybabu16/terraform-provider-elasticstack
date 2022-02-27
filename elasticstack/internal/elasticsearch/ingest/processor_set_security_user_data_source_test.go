package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorSetSecurityUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorSetSecurityUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_set_security_user.test", "field", "user"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_set_security_user.test", "json", expectedJsonSetSecurityUser),
				),
			},
		},
	})
}

const expectedJsonSetSecurityUser = `{
	"set_security_user": {
		"field": "user",
		"ignore_failure": false
	}
}`

const testAccDataSourceIngestProcessorSetSecurityUser = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_set_security_user" "test" {
  field = "user"
}
`
