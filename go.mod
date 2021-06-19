module github.com/phillbaker/terraform-provider-elasticsearch

go 1.12

require (
	github.com/aws/aws-sdk-go v1.38.17
	github.com/deoxxa/aws_signing_client v0.0.0-20161109131055-c20ee106809e
	github.com/hashicorp/go-version v1.2.1
	github.com/hashicorp/terraform-plugin-sdk v1.12.0
	github.com/olivere/elastic v6.2.26+incompatible
	github.com/olivere/elastic/v7 v7.0.25
	gopkg.in/olivere/elastic.v5 v5.0.85
	gopkg.in/olivere/elastic.v6 v6.2.35
)

replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
