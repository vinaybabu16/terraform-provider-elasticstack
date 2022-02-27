package ingest_test

import (
	"testing"

	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIngestProcessorTrim(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorTrim,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_trim.test", "field", "foo"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_trim.test", "json", expectedJsonTrim),
				),
			},
		},
	})
}

const expectedJsonTrim = `{
	"trim": {
		"field": "foo",
		"ignore_failure": false,
		"ignore_missing": false
	}
}`

const testAccDataSourceIngestProcessorTrim = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_trim" "test" {
  field = "foo"
}
`
