package me

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "me"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_me_identity_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Group"
	})
	p.AddResourceConfigurator("ovh_me_identity_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.References["group"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/namespaced/me/v1alpha1.Group",
		}
	})
	p.AddResourceConfigurator("ovh_me_api_oauth2_client", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Oauth2Client"
	})
}
