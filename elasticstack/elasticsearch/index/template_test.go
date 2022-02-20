package index_test

import (
	"fmt"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/acctest"
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/clients"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceIndexTemplate(t *testing.T) {
	// generate random template name
	templateName := sdkacctest.RandStringFromCharSet(10, sdkacctest.CharSetAlphaNum)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		CheckDestroy:      checkResourceIndexTemplateDestroy,
		ProviderFactories: acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIndexTemplateCreate(templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("elasticstack_elasticsearch_index_template.test", "name", templateName),
					resource.TestCheckTypeSetElemAttr("elasticstack_elasticsearch_index_template.test", "index_patterns.*", fmt.Sprintf("%s-logs-*", templateName)),
					resource.TestCheckResourceAttr("elasticstack_elasticsearch_index_template.test", "priority", "42"),
				),
			},
		},
	})
}

func testAccResourceIndexTemplateCreate(name string) string {
	return fmt.Sprintf(`
provider "elasticstack" {
  elasticsearch {}
}

resource "elasticstack_elasticsearch_index_template" "test" {
  name = "%s"

  priority       = 42
  index_patterns = ["%s-logs-*"]

  template {
    alias {
      name = "my_template_test"
    }

    settings = jsonencode({
      number_of_shards = "3"
    })
  }
}
	`, name, name)
}

func checkResourceIndexTemplateDestroy(s *terraform.State) error {
	client := acctest.Provider.Meta().(*clients.ApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "elasticstack_elasticsearch_index_template" {
			continue
		}
		compId, _ := clients.CompositeIdFromStr(rs.Primary.ID)

		req := client.GetESClient().Indices.GetIndexTemplate.WithName(compId.ResourceId)
		res, err := client.GetESClient().Indices.GetIndexTemplate(req)
		if err != nil {
			return err
		}

		if res.StatusCode != 404 {
			return fmt.Errorf("Index template (%s) still exists", compId.ResourceId)
		}
	}
	return nil
}
