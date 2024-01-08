package cloud

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "cloud"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_credential", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Credentials"
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Policy"
	})
}
