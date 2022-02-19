package acctest

import (
	"github.com/elastic/terraform-provider-elasticstack/elasticstack"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Providers map[string]func() (*schema.Provider, error)
var Provider *schema.Provider

func init() {
	Provider = elasticstack.New("dev")()
	Providers = map[string]func() (*schema.Provider, error){
		"elasticstack": func() (*schema.Provider, error) {
			return Provider, nil
		},
	}
}

func PreCheck(t *testing.T) {
	_, endpointsOk := os.LookupEnv("ELASTICSEARCH_ENDPOINTS")
	_, userOk := os.LookupEnv("ELASTICSEARCH_USERNAME")
	_, passOk := os.LookupEnv("ELASTICSEARCH_PASSWORD")

	if !endpointsOk || !userOk || !passOk {
		t.Fatal("ELASTICSEARCH_ENDPOINTS, ELASTICSEARCH_USERNAME, ELASTICSEARCH_PASSWORD must be set for acceptance tests to run")
	}
}
