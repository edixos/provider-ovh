package registry

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "registry"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_containerregistry_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
