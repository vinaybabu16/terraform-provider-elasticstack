## [Unreleased]

## [0.3.2] - 2022-02-28
### Fixed
- Properly apply `number_of_replicas` setting in `allocate` action in ILM ([#80](https://github.com/elastic/terraform-provider-elasticstack/issues/80))

## [0.3.1] - 2022-02-24
### Fixed
- Add new field `allow_custom_routing` in `data_stream` section of [`index_template`](https://www.elastic.co/guide/en/elasticsearch/reference/8.0/indices-put-template.html#put-index-template-api-request-body), which appears only in Elasticsearch version **8.0.0**. Make sure `index_template` resource can work with both **7.x** and **8.x** versions ([#72](https://github.com/elastic/terraform-provider-elasticstack/pull/72))
- Panic using `elasticstack_elasticsearch_security_user_data_source` when the user absent or there is not enough permissions to fetch users from ES API ([#73](https://github.com/elastic/terraform-provider-elasticstack/issues/73))
- Fix typo in the index alias model ([#78](https://github.com/elastic/terraform-provider-elasticstack/issues/78))

## [0.3.0] - 2022-02-17
### Added
- New resource `elasticstack_elasticsearch_data_stream` to manage Elasticsearch [data streams](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html) ([#45](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/45))
- New resource `elasticstack_elasticsearch_ingest_pipeline` to manage Elasticsearch [ingest pipelines](https://www.elastic.co/guide/en/elasticsearch/reference/7.16/ingest.html) ([#56](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerissues/56))
- New resource `elasticstack_elasticsearch_component_template` to manage Elasticsearch [component templates](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html) ([#39](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/39))
- New helper data sources to create [processorts](https://www.elastic.co/guide/en/elasticsearch/reference/current/processors.html) for ingest pipelines ([#67](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/67))

### Fixed
- Update only changed index settings ([#52](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerissues/52))
- Enable import of index settings ([#53](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerissues/53))
- Handle `allocate` action in ILM policy ([#59](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerissues/59))
- Send only initialized values to Elasticsearch API when managing the users ([#66](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerissues/66))

## [0.2.0] - 2022-01-27
### Added
- New resource to manage Elasticsearch indices ([#38](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/38))
- The `insecure` option to the Elasticsearch provider configuration to ignore certificate verification ([#36](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/36))
- The `ca_file` option to the Elasticsearch provider configuration to provide path to the custom root certificate file ([#35](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/35))
- Documentation templates for all the exisiting resources ([#32](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/32))

### Fixed
- Identify missing or deleted resources in the Elasticsearch cluster and make sure Terraform can re-create them ([#40](https://github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/providerpull/40))

### Changed
- **[Breaking]** Rename `aliases` configuration option in
`elasticstack_elasticsearch_index_template` resource to singular `alias`

## [0.1.0] - 2021-12-20
### Added
- Initial set of the resources to gather feedback from the community
- Initial set of docs
- CI integration

[Unreleased]: https://github.com/elastic/terraform-provider-elasticstack/compare/v0.3.2...HEAD
[0.3.1]: https://github.com/elastic/terraform-provider-elasticstack/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/elastic/terraform-provider-elasticstack/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/elastic/terraform-provider-elasticstack/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/elastic/terraform-provider-elasticstack/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/elastic/terraform-provider-elasticstack/releases/tag/v0.1.0
