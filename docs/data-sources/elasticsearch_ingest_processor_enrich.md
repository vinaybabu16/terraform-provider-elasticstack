---
subcategory: "Ingest"
layout: ""
page_title: "Elasticstack: elasticstack_elasticsearch_ingest_processor_enrich Data Source"
description: |-
  Helper data source to create a processor which enriches documents with data from another index.
---

# Data Source: elasticstack_elasticsearch_ingest_processor_enrich

The enrich processor can enrich documents with data from another index. See enrich data section for more information about how to set this up.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-enriching-data.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-processor.html

## Example Usage

```terraform
provider "elasticstack" {
  elasticsearch {}
}

// the policy must exist before using this processor
// See example at: https://www.elastic.co/guide/en/elasticsearch/reference/current/match-enrich-policy-type.html
data "elasticstack_elasticsearch_ingest_processor_enrich" "enrich" {
  policy_name  = "users-policy"
  field        = "email"
  target_field = "user"
  max_matches  = 1
}

resource "elasticstack_elasticsearch_ingest_pipeline" "my_ingest_pipeline" {
  name = "enrich-ingest"

  processors = [
    data.elasticstack_elasticsearch_ingest_processor_enrich.enrich.json
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **field** (String) The field in the input document that matches the policies match_field used to retrieve the enrichment data.
- **policy_name** (String) The name of the enrich policy to use.
- **target_field** (String) Field added to incoming documents to contain enrich data.

### Optional

- **description** (String) Description of the processor.
- **if** (String) Conditionally execute the processor
- **ignore_failure** (Boolean) Ignore failures for the processor.
- **ignore_missing** (Boolean) If `true` and `field` does not exist or is `null`, the processor quietly exits without modifying the document.
- **max_matches** (Number) The maximum number of matched documents to include under the configured target field.
- **on_failure** (List of String) Handle failures for the processor.
- **override** (Boolean) If processor will update fields with pre-existing non-null-valued field.
- **shape_relation** (String) A spatial relation operator used to match the geoshape of incoming documents to documents in the enrich index.
- **tag** (String) Identifier for the processor.

### Read-Only

- **id** (String) Internal identifier of the resource
- **json** (String) JSON representation of this data source.
