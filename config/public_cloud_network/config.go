package public_cloud_network

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_network_private", func(r *config.Resource) {
		r.ShortGroup = "publiccloudnetwork"
	})
	p.AddResourceConfigurator("ovh_cloud_project_network_private_subnet", func(r *config.Resource) {
		r.ShortGroup = "publiccloudnetwork"
	})
}
