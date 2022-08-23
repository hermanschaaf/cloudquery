package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/teams"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "heroku",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"teams": teams.Teams(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
