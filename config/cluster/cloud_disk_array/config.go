package cloud_disk_array

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_dedicated_ceph_acl", func(r *config.Resource) {
		r.ShortGroup = "clouddiskarray"
	})
}
