package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/kriyaanshtechnology/terraform-sonarcloud-provider/sonarcloud"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: sonarcloud.Provider,
		},
	)
}
