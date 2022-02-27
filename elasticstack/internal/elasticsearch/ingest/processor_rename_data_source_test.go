package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorRename(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorRename,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_rename.test", "field", "provider"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_rename.test", "json", expectedJsonRename),
				),
			},
		},
	})
}

const expectedJsonRename = `{
	"rename": {
		"field": "provider",
		"target_field": "cloud.provider",
		"ignore_failure": false,
		"ignore_missing": false
	}
}`

const testAccDataSourceIngestProcessorRename = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_rename" "test" {
  field        = "provider"
  target_field = "cloud.provider"
}
`
