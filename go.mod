module github.com/CiscoDevNet/terraform-provider-intersight

go 1.14

require (
	github.com/bflad/tfproviderlint v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

require (
	github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk