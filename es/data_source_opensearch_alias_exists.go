package es

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/olivere/elastic/uritemplates"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	elastic7 "github.com/olivere/elastic/v7"
	elastic6 "gopkg.in/olivere/elastic.v6"
)

func dataSourceOpenSearchAliasExists() *schema.Resource {
	return &schema.Resource{
		Description: "`opensearch_alias_exists` can be used to checks if an alias exists.",
		Read:        dataSourceOpenSearchAliasExistsRead,

		Schema: map[string]*schema.Schema{
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "the alias name",
			},
			"exists": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "is alias exists",
			},
		},
	}
}

func dataSourceOpenSearchAliasExistsRead(d *schema.ResourceData, m interface{}) error {

	// The upstream elastic client does not export the property for the urls
	// it's using. Presumably the URLS would be available where the client is
	// intantiated, but in terraform, that's not always practicable.
	var err error
	response := new(GetPolicyResponse)
	path, err := uritemplates.Expand("/_alias/{alias_name}", map[string]string{
		"alias_name": d.Id(),
	})

	if err != nil {
		return fmt.Errorf("error building URL path for policy: %+v", err)
	}

	var body *json.RawMessage

	esClient, err := getClient(m.(*ProviderConf))
	if err != nil {
		return err
	}

	switch client := esClient.(type) {
	case *elastic7.Client:
		var res *elastic7.Response
		res, err = client.PerformRequest(context.TODO(), elastic7.PerformRequestOptions{
			Method: "HEAD",
			Path:   path,
		})

		if err != nil {
			return fmt.Errorf("error getting alias: %+v : %+v", path, err)
		}
		body = &res.Body
	case *elastic6.Client:
		var res *elastic6.Response
		res, err = client.PerformRequest(context.TODO(), elastic6.PerformRequestOptions{
			Method: "HEAD",
			Path:   path,
		})

		if err != nil {
			return fmt.Errorf("error getting alias: %+v : %+v", path, err)
		}
		body = &res.Body
	default:
		return errors.New("this version of Elasticsearch is not supported")
	}
	d.SetId("sss")
	err = d.Set("url", "sdd")
	if err := json.Unmarshal(*body, &response); err != nil {
		return fmt.Errorf("error unmarshalling policy body: %+v: %+v", err, body)
	}
	return err
}
