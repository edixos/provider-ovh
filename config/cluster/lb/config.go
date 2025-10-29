package lb

import "github.com/crossplane/upjet/v2/pkg/config"

const (
	shortGroup = "lb"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_iploadbalancing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_http_farm", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_http_farm_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_http_frontend", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_http_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_http_route_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_refresh", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_tcp_farm", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_tcp_farm_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_tcp_frontend", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_tcp_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_tcp_route_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_vrack_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_udp_frontend", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_cloud_project_region_loadbalancer_log_subscription", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_ssl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_udp_farm", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("ovh_iploadbalancing_udp_farm_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
