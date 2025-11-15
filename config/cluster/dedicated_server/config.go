package dedicated_server

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "dedicatedserver"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_dedicated_server_install_task", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_dedicated_server_reboot_task", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_dedicated_server_update", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_me_installation_template", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_me_installation_template_partition_scheme", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_me_installation_template_partition_scheme_hardware_raid", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_me_installation_template_partition_scheme_partition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_me_ipxe_script", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("ovh_dedicated_server_networking", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
