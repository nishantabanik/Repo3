package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/kriyaanshtechnology/terraform-provider-sonarcloud/sonarcloud"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: sonarcloud.Provider,
		},
	)
}
