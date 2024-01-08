package me

import "github.com/crossplane/upjet/pkg/config"

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
	})
	p.AddResourceConfigurator("ovh_me_ssh_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SSHKey"
	})
}
