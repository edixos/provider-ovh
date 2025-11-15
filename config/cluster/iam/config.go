package iam

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "iam"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_iam_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IAMPolicy"
	})
	p.AddResourceConfigurator("ovh_iam_resource_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IAMResourceGroup"
	})
	p.AddResourceConfigurator("ovh_iam_permissions_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IAMPermissionsGroup"
	})
}
