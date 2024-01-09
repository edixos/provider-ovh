/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/upjet/pkg/config"
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
	"ovh_cloud_project_network_private":                           config.NameAsIdentifier,
	"ovh_cloud_project_network_private_subnet":                    config.NameAsIdentifier,
	"ovh_cloud_project_workflow_backup":                           config.NameAsIdentifier,
	"ovh_cloud_project":                                           config.IdentifierFromProvider,
	"ovh_cloud_project_user":                                      config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_credential":                        config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_policy":                            config.IdentifierFromProvider,
	"ovh_iam_policy":                                              config.NameAsIdentifier,
	"ovh_me_identity_group":                                       config.NameAsIdentifier,
	"ovh_me_identity_user":                                        config.NameAsIdentifier,
	"ovh_me_ssh_key":                                              config.NameAsIdentifier,
	"ovh_cloud_project_failover_ip_attach":                        config.NameAsIdentifier,
	"ovh_ip_reverse":                                              config.NameAsIdentifier,
	"ovh_ip_service":                                              config.NameAsIdentifier,
	"ovh_dedicated_ceph_acl":                                      config.NameAsIdentifier,
	"ovh_dedicated_server_install_task":                           config.NameAsIdentifier,
	"ovh_dedicated_server_reboot_task":                            config.NameAsIdentifier,
	"ovh_dedicated_server_update":                                 config.NameAsIdentifier,
	"ovh_me_installation_template":                                config.NameAsIdentifier,
	"ovh_me_installation_template_partition_scheme":               config.NameAsIdentifier,
	"ovh_me_installation_template_partition_scheme_hardware_raid": config.NameAsIdentifier,
	"ovh_me_installation_template_partition_scheme_partition":     config.NameAsIdentifier,
	"ovh_me_ipxe_script":                                          config.NameAsIdentifier,
	"ovh_dedicated_server_networking":                             config.NameAsIdentifier,
	"ovh_domain_zone":                                             config.NameAsIdentifier,
	"ovh_domain_zone_record":                                      config.NameAsIdentifier,
	"ovh_domain_zone_redirection":                                 config.NameAsIdentifier,
	"ovh_iploadbalancing":                                         config.NameAsIdentifier,
	"ovh_iploadbalancing_http_farm":                               config.NameAsIdentifier,
	"ovh_iploadbalancing_http_farm_server":                        config.NameAsIdentifier,
	"ovh_iploadbalancing_http_frontend":                           config.NameAsIdentifier,
	"ovh_iploadbalancing_http_route":                              config.NameAsIdentifier,
	"ovh_iploadbalancing_http_route_rule":                         config.NameAsIdentifier,
	"ovh_iploadbalancing_refresh":                                 config.NameAsIdentifier,
	"ovh_iploadbalancing_tcp_farm":                                config.NameAsIdentifier,
	"ovh_iploadbalancing_tcp_farm_server":                         config.NameAsIdentifier,
	"ovh_iploadbalancing_tcp_frontend":                            config.NameAsIdentifier,
	"ovh_iploadbalancing_tcp_route":                               config.NameAsIdentifier,
	"ovh_iploadbalancing_tcp_route_rule":                          config.NameAsIdentifier,
	"ovh_iploadbalancing_vrack_network":                           config.NameAsIdentifier,
	"ovh_dbaas_logs_cluster":                                      config.NameAsIdentifier,
	"ovh_dbaas_logs_graylog_output_stream":                        config.NameAsIdentifier,
	"ovh_dbaas_logs_input":                                        config.NameAsIdentifier,
	"ovh_cloud_project_database":                                  config.NameAsIdentifier,
	"ovh_cloud_project_database_database":                         config.NameAsIdentifier,
	"ovh_cloud_project_database_integration":                      config.NameAsIdentifier,
	"ovh_cloud_project_database_ip_restriction":                   config.NameAsIdentifier,
	"ovh_cloud_project_database_kafka_acl":                        config.NameAsIdentifier,
	"ovh_cloud_project_database_kafka_schemaregistryacl":          config.NameAsIdentifier,
	"ovh_cloud_project_database_kafka_topic":                      config.NameAsIdentifier,
	"ovh_cloud_project_database_m3db_namespace":                   config.NameAsIdentifier,
	"ovh_cloud_project_database_m3db_user":                        config.NameAsIdentifier,
	"ovh_cloud_project_database_mongodb_user":                     config.NameAsIdentifier,
	"ovh_cloud_project_database_opensearch_pattern":               config.NameAsIdentifier,
	"ovh_cloud_project_database_opensearch_user":                  config.NameAsIdentifier,
	"ovh_cloud_project_database_postgresql_user":                  config.NameAsIdentifier,
	"ovh_cloud_project_database_redis_user":                       config.NameAsIdentifier,
	"ovh_cloud_project_database_user":                             config.NameAsIdentifier,
	"ovh_cloud_project_kube":                                      kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_iprestrictions":                       kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_nodepool":                             kubePoolIdentifierFromProvider,
	"ovh_cloud_project_kube_oidc":                                 config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry":                         config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_oidc":                    config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_user":                    config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition":                               config.NameAsIdentifier,
	"ovh_dedicated_nasha_partition_access":                        config.NameAsIdentifier,
	"ovh_dedicated_nasha_partition_snapshot":                      config.NameAsIdentifier,
	"ovh_cloud_project_region_storage_presign":                    config.NameAsIdentifier,
	"ovh_hosting_privatedatabase":                                 config.NameAsIdentifier,
	"ovh_hosting_privatedatabase_database":                        config.NameAsIdentifier,
	"ovh_hosting_privatedatabase_user":                            config.NameAsIdentifier,
	"ovh_hosting_privatedatabase_user_grant":                      config.NameAsIdentifier,
	"ovh_hosting_privatedatabase_whitelist":                       config.NameAsIdentifier,
	"ovh_vrack":                                                   config.NameAsIdentifier,
	"ovh_vrack_cloudproject":                                      config.NameAsIdentifier,
	"ovh_vrack_dedicated_server":                                  config.NameAsIdentifier,
	"ovh_vrack_dedicated_server_interface":                        config.NameAsIdentifier,
	"ovh_vrack_ip":                                                config.NameAsIdentifier,
	"ovh_vrack_iploadbalancing":                                   config.NameAsIdentifier,
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
