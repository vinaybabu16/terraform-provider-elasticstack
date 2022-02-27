package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorCSV(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorCSV,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_csv.test", "field", "my_field"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_csv.test", "json", expectedJsonCSV),
				),
			},
		},
	})
}

const expectedJsonCSV = `{
	"csv": {
		"field": "my_field",
		"target_fields": ["field1", "field2"],
		"separator": ",",
		"trim": false,
		"quote": "\"",
		"ignore_failure": false,
		"ignore_missing": false
	}
}`

const testAccDataSourceIngestProcessorCSV = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_csv" "test" {
  field         = "my_field"
  target_fields = ["field1", "field2"]
}
`
