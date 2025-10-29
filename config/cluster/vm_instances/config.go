package vm_instances

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_workflow_backup", func(r *config.Resource) {
		r.ShortGroup = "vminstances"
	})
}
