package nas

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "nas"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_dedicated_nasha_partition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_dedicated_nasha_partition_access", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_dedicated_nasha_partition_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
