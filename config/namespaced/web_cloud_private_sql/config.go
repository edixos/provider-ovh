package web_cloud_private_sql

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "privatesql"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_hosting_privatedatabase", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_hosting_privatedatabase_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_hosting_privatedatabase_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_hosting_privatedatabase_user_grant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_hosting_privatedatabase_whitelist", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
