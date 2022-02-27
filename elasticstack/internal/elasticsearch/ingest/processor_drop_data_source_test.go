package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorDrop(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorDrop,
				Check: resource.ComposeTestCheckFunc(
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_drop.test", "json", expectedJsonDrop),
				),
			},
		},
	})
}

const expectedJsonDrop = `{
  "drop": {
		"ignore_failure": false,
		"if" : "ctx.network_name == 'Guest'"
	}
}
`

const testAccDataSourceIngestProcessorDrop = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_drop" "test" {
  if = "ctx.network_name == 'Guest'"
}
`
