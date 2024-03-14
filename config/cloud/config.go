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
	p.AddResourceConfigurator("ovh_cloud_project_alerting", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectAlerting"
	})
	p.AddResourceConfigurator("ovh_cloud_project_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_credential", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Credentials"
		r.References["user_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cloud/v1alpha1.User",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Policy"
		r.References["user_id"] = config.Reference{
			Type: "github.com/edixos/provider-ovh/apis/cloud/v1alpha1.User",
		}
	})
}
