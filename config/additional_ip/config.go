package additional_ip

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "additionalip"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_failover_ip_attach", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_ip_reverse", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_ip_move", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_ip_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
