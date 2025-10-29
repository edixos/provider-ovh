package vrack

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "vrack"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_vrack", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_vrack_cloudproject", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_vrack_dedicated_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_vrack_dedicated_server_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_vrack_ip", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_vrack_iploadbalancing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
