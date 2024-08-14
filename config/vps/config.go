package vps

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.SkipList = append(p.SkipList, "ovh_vps")
}
