package elasticstack

import (
	"github.com/elastic/terraform-provider-elasticstack/elasticstack/clients"
	cluster2 "github.com/elastic/terraform-provider-elasticstack/elasticstack/elasticsearch/cluster"
	index2 "github.com/elastic/terraform-provider-elasticstack/elasticstack/elasticsearch/index"
	security2 "github.com/elastic/terraform-provider-elasticstack/elasticstack/elasticsearch/security"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{

			Schema: map[string]*schema.Schema{
				"elasticsearch": {
					Description: "Default Elasticsearch connection configuration block.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"username": {
								Description: "Username to use for API authentication to Elasticsearch.",
								Type:        schema.TypeString,
								Optional:    true,
								DefaultFunc: schema.EnvDefaultFunc("ELASTICSEARCH_USERNAME", nil),
							},
							"password": {
								Description: "Password to use for API authentication to Elasticsearch.",
								Type:        schema.TypeString,
								Optional:    true,
								Sensitive:   true,
								DefaultFunc: schema.EnvDefaultFunc("ELASTICSEARCH_PASSWORD", nil),
							},
							"endpoints": {
								Description: "A comma-separated list of endpoints where the terraform provider will point to, this must include the http(s) schema and port number.",
								Type:        schema.TypeList,
								Optional:    true,
								Sensitive:   true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
							"insecure": {
								Description: "Disable TLS certificate validation",
								Type:        schema.TypeBool,
								Optional:    true,
								Default:     false,
							},
							"ca_file": {
								Description: "Path to a custom Certificate Authority certificate",
								Type:        schema.TypeString,
								Optional:    true,
							},
						},
					},
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"elasticstack_elasticsearch_security_user":       security2.DataSourceUser(),
				"elasticstack_elasticsearch_snapshot_repository": cluster2.DataSourceSnapshotRespository(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"elasticstack_elasticsearch_cluster_settings":    cluster2.ResourceSettings(),
				"elasticstack_elasticsearch_index":               index2.ResourceIndex(),
				"elasticstack_elasticsearch_index_lifecycle":     index2.ResourceIlm(),
				"elasticstack_elasticsearch_index_template":      index2.ResourceTemplate(),
				"elasticstack_elasticsearch_security_role":       security2.ResourceRole(),
				"elasticstack_elasticsearch_security_user":       security2.ResourceUser(),
				"elasticstack_elasticsearch_snapshot_lifecycle":  cluster2.ResourceSlm(),
				"elasticstack_elasticsearch_snapshot_repository": cluster2.ResourceSnapshotRepository(),
			},
		}

		p.ConfigureContextFunc = clients.NewApiClientFunc(version, p)

		return p
	}
}
