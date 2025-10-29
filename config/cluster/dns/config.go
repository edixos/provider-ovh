package dns

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortName = "dns"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_domain_zone", func(r *config.Resource) {
		r.ShortGroup = shortName
	})
	p.AddResourceConfigurator("ovh_domain_zone_dnssec", func(r *config.Resource) {
		r.ShortGroup = shortName
	})
	p.AddResourceConfigurator("ovh_domain_zone_record", func(r *config.Resource) {
		r.ShortGroup = shortName
	})
	p.AddResourceConfigurator("ovh_domain_zone_redirection", func(r *config.Resource) {
		r.ShortGroup = shortName
	})

}
