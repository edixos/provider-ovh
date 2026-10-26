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
		r.Kind = "ProjectStorage"
		// UseAsync is needed for resources that may take time to become available
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ovh_cloud_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Instance"
		r.References["group_id"] = config.Reference{
			TerraformName: "ovh_cloud_instance_group",
		}
		r.References["ssh_key_name"] = config.Reference{
			TerraformName: "ovh_cloud_ssh_key",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_instance_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InstanceGroup"
	})
	p.AddResourceConfigurator("ovh_cloud_instance_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InstanceSnapshot"
		r.References["instance_id"] = config.Reference{
			TerraformName: "ovh_cloud_instance",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_loadbalancer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Loadbalancer"
	})
	p.AddResourceConfigurator("ovh_cloud_loadbalancer_listener", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadbalancerListener"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer",
		}
		r.References["default_pool_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer_pool",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_loadbalancer_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadbalancerPool"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_loadbalancer_pool_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadbalancerPoolMember"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer",
		}
		r.References["pool_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer_pool",
		}
	})
	p.AddResourceConfigurator("ovh_cloud_loadbalancer_l7policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadbalancerL7Policy"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer",
		}
		r.References["listener_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer_listener",
		}
		r.References["redirect_pool_id"] = config.Reference{
			TerraformName: "ovh_cloud_loadbalancer_pool",
		}
	})
}
