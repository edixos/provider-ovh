package object_storage

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "storage"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_region_storage_presign", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
