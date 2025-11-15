/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

func serviceName(parameters map[string]any) (string, error) {
	serviceName, ok := parameters["service_name"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "service_name")
	}
	serviceNameStr, ok := serviceName.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "service_name")
	}
	return serviceNameStr, nil
}

var kubePoolIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		kubeID, ok := parameters["kube_id"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "kube_id")
		}
		kubeIDStr, ok := kubeID.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "kube_id")
		}

		return fmt.Sprintf("%s/%s/%s", serviceName, kubeIDStr, externalName), nil
	},
	DisableNameInitializer: true,
}

var kubeIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		return fmt.Sprintf("%s/%s", serviceName, externalName), nil
	},
	DisableNameInitializer: true,
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"ovh_cloud_project_network_private":           config.IdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet":    config.IdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet_v2": config.IdentifierFromProvider,

	"ovh_cloud_project_workflow_backup": config.NameAsIdentifier,
	"ovh_cloud_project":                 config.IdentifierFromProvider,
	// The ovh_cloud_project_alerting resource uses a nested type which is not supported yet in upjet.
	// there is an open issue in upjet regarding this issue: https://github.com/crossplane/upjet/v2/issues/372
	// "ovh_cloud_project_alerting":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_user":                        config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_credential":          config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_policy":              config.IdentifierFromProvider,
	"ovh_iam_policy":                                config.IdentifierFromProvider,
	"ovh_iam_resource_group":                        config.IdentifierFromProvider,
	"ovh_iam_permissions_group":                     config.IdentifierFromProvider,
	"ovh_me_identity_group":                         config.IdentifierFromProvider,
	"ovh_me_identity_user":                          config.IdentifierFromProvider,
	"ovh_me_ssh_key":                                config.IdentifierFromProvider,
	"ovh_me_api_oauth2_client":                      config.IdentifierFromProvider,
	"ovh_cloud_project_failover_ip_attach":          config.IdentifierFromProvider,
	"ovh_ip_reverse":                                config.IdentifierFromProvider,
	"ovh_ip_move":                                   config.IdentifierFromProvider,
	"ovh_ip_service":                                config.IdentifierFromProvider,
	"ovh_ip_firewall":                               config.IdentifierFromProvider,
	"ovh_ip_firewall_rule":                          config.IdentifierFromProvider,
	"ovh_ip_mitigation":                             config.IdentifierFromProvider,
	"ovh_dedicated_ceph_acl":                        config.IdentifierFromProvider,
	"ovh_dedicated_server_install_task":             config.IdentifierFromProvider,
	"ovh_dedicated_server_reboot_task":              config.IdentifierFromProvider,
	"ovh_dedicated_server_update":                   config.IdentifierFromProvider,
	"ovh_me_installation_template":                  config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme": config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme_hardware_raid": config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme_partition":     config.IdentifierFromProvider,
	"ovh_me_ipxe_script":                                             config.IdentifierFromProvider,
	"ovh_dedicated_server_networking":                                config.IdentifierFromProvider,
	"ovh_domain_zone":                                                config.IdentifierFromProvider,
	"ovh_domain_zone_dnssec":                                         config.IdentifierFromProvider,
	"ovh_domain_zone_record":                                         config.IdentifierFromProvider,
	"ovh_domain_zone_redirection":                                    config.IdentifierFromProvider,
	"ovh_iploadbalancing":                                            config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_frontend":                               config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_farm":                                  config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_farm_server":                           config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_frontend":                              config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_route":                                 config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_route_rule":                            config.IdentifierFromProvider,
	"ovh_iploadbalancing_refresh":                                    config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_farm":                                   config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_farm_server":                            config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_frontend":                               config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_route":                                  config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_route_rule":                             config.IdentifierFromProvider,
	"ovh_iploadbalancing_vrack_network":                              config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm_server":                            config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm":                                   config.IdentifierFromProvider,
	"ovh_iploadbalancing_ssl":                                        config.IdentifierFromProvider,
	"ovh_dbaas_logs_cluster":                                         config.IdentifierFromProvider,
	"ovh_dbaas_logs_graylog_output_stream":                           config.IdentifierFromProvider,
	"ovh_dbaas_logs_input":                                           config.IdentifierFromProvider,
	"ovh_dbaas_logs_token":                                           config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_alias":                         config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_index":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_database":                            config.IdentifierFromProvider,
	"ovh_cloud_project_database_integration":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_ip_restriction":                      config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_acl":                           config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_schemaregistryacl":             config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_topic":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_m3db_namespace":                      config.IdentifierFromProvider,
	"ovh_cloud_project_database_m3db_user":                           config.IdentifierFromProvider,
	"ovh_cloud_project_database_mongodb_user":                        config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_pattern":                  config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_user":                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_user":                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_connection_pool":          config.IdentifierFromProvider,
	"ovh_cloud_project_database_redis_user":                          config.IdentifierFromProvider,
	"ovh_cloud_project_database_user":                                config.IdentifierFromProvider,
	"ovh_cloud_project_kube":                                         kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_iprestrictions":                          kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_nodepool":                                kubePoolIdentifierFromProvider,
	"ovh_cloud_project_kube_oidc":                                    config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry":                            config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_oidc":                       config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_user":                       config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_ip_restrictions_management": config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_ip_restrictions_registry":   config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition":                                  config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition_access":                           config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition_snapshot":                         config.IdentifierFromProvider,
	"ovh_cloud_project_region_storage_presign":                       config.IdentifierFromProvider,
	"ovh_cloud_project_region_loadbalancer_log_subscription":         config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase":                                    config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_database":                           config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_user":                               config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_user_grant":                         config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_whitelist":                          config.IdentifierFromProvider,
	"ovh_vrack":                                                      config.IdentifierFromProvider,
	"ovh_vrack_cloudproject":                                         config.IdentifierFromProvider,
	"ovh_vrack_dedicated_server":                                     config.IdentifierFromProvider,
	"ovh_vrack_dedicated_server_interface":                           config.IdentifierFromProvider,
	"ovh_vrack_ip":                                                   config.IdentifierFromProvider,
	"ovh_vrack_iploadbalancing":                                      config.IdentifierFromProvider,
	// "ovh_vps":                                                        config.IdentifierFromProvider,
	"ovh_cloud_project_gateway": config.IdentifierFromProvider,
	// ovh_okms, ovh_okms_credential, ovh_okms_service_key are not supported yet in upjet. because they uses nested types.
	// this is an open issue in upjet regarding this issue:
	// https://github.com/crossplane/upjet/v2/issues/372
	// "ovh_okms":                  config.IdentifierFromProvider,
	// "ovh_okms_credential":       config.IdentifierFromProvider,
	// "ovh_okms_service_key":      config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
