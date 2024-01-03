package public_cloud_network

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "pcn"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_network_private", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_network_private_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
