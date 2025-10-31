package cloud

import "github.com/crossplane/upjet/v2/pkg/config"

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
	// The ovh_cloud_project_alerting resource uses a nested type which is not supported yet in upjet.
	// there is an open issue in upjet regarding this issue: https://github.com/crossplane/upjet/v2/issues/372

	// p.AddResourceConfigurator("ovh_cloud_project_alerting", func(r *config.Resource) {
	// 	r.ShortGroup = shortGroup
	// 	r.Kind = "ProjectAlerting"
	// })
	p.AddResourceConfigurator("ovh_cloud_project_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_credential", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Credentials"
		r.References["user_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_user",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Policy"
		r.References["user_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_user",
		}
	})
}
