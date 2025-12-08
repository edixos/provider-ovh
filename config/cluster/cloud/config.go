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

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			details := make(map[string][]byte)
			if accessKey, ok := attr["access_key_id"].(string); ok {
				details["access_key_id"] = []byte(accessKey)
			}

			return details, nil
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "S3Policy"
		r.References["user_id"] = config.Reference{
			TerraformName: "ovh_cloud_project_user",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_project_volume", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_savings_plan", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_instance_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_rancher", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_ssh_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_region", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_storage", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
