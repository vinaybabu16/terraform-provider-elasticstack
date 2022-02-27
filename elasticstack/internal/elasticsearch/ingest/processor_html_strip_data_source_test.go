package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorHtmlStrip(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorHtmlStrip,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.elasticstack_elasticsearch_ingest_processor_html_strip.test", "field", "foo"),
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_html_strip.test", "json", expectedJsonHtmlStrip),
				),
			},
		},
	})
}

const expectedJsonHtmlStrip = `{
	"html_strip": {
		"field": "foo",
		"ignore_failure": false,
		"ignore_missing": false
	}
}`

const testAccDataSourceIngestProcessorHtmlStrip = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_html_strip" "test" {
  field = "foo"
}
`
