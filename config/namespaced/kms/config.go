package kms

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "kms"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_okms", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_okms_service_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_okms_credential", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

}
