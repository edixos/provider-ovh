package logs

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "logs"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_dbaas_logs_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_dbaas_logs_graylog_output_stream", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_dbaas_logs_input", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_dbaas_logs_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})

}
