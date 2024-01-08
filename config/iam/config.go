package iam

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "iam"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_iam_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IamPolicy"
	})
}
