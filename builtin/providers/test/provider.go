package test

import (
	"github.com/eugenetaranov/terraform/helper/schema"
	"github.com/eugenetaranov/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"test_resource":         testResource(),
			"test_resource_gh12183": testResourceGH12183(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"test_data_source": testDataSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
