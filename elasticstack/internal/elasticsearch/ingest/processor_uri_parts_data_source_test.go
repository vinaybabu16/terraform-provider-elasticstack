package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorUriParts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorUriParts,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_uri_parts.test", "field", "input_field"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_uri_parts.test", "json", expectedJsonUriParts),
				),
			},
		},
	})
}

const expectedJsonUriParts = `{
	"uri_parts": {
		"field": "input_field",
		"ignore_failure": false,
		"keep_original": true,
		"remove_if_successful": false,
		"target_field": "url"
	}
}`

const testAccDataSourceIngestProcessorUriParts = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_uri_parts" "test" {
  field                = "input_field"
  target_field         = "url"
  keep_original        = true
  remove_if_successful = false
}
`
