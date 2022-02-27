package ingest_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestAccDataSourceIngestProcessorFail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIngestProcessorFail,
				Check: resource.ComposeTestCheckFunc(
					CheckResourceJson("data.elasticstack_elasticsearch_ingest_processor_fail.test", "json", expectedJsonFail),
				),
			},
		},
	})
}

const expectedJsonFail = `{
  "fail": {
		"message": "The production tag is not present, found tags: {{{tags}}}",
		"ignore_failure": false,
		"if" : "ctx.tags.contains('production') != true"
	}
}
`

const testAccDataSourceIngestProcessorFail = `
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_fail" "test" {
  if      = "ctx.tags.contains('production') != true"
  message = "The production tag is not present, found tags: {{{tags}}}"
}
`
