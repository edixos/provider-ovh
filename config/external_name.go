/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// isOVHNotFoundError checks if the diagnostic is an OVH 404 error
func isOVHNotFoundError(diags []*tfprotov6.Diagnostic) bool {
	fmt.Printf("isOVHNotFoundError called with %d diagnostics\n", len(diags))
	for i, d := range diags {
		if d == nil {
			fmt.Printf("Diagnostic %d is nil\n", i)
			continue
		}
		fmt.Printf("Diagnostic %d - Severity: %v\n", i, d.Severity)
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			summary := d.Summary
			detail := d.Detail

			fmt.Printf("SUMMARY: %s\n", summary)
			fmt.Printf("DETAIL: %s\n", detail)
			// Check if it's an OVH 404 error
			if strings.Contains(summary, "Error calling Get") &&
				(strings.Contains(detail, "status code 404") ||
					strings.Contains(detail, "Client::NotFound") ||
					strings.Contains(detail, "not found: NotFound")) {
				fmt.Printf("DETECTED 404 ERROR - returning true\n")
				return true
			}
		}
	}
	fmt.Printf("No 404 error detected - returning false\n")
	return false
}

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

func clusterID(parameters map[string]any) (string, error) {
	clusterID, ok := parameters["cluster_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "cluster_id")
	}
	clusterIDStr, ok := clusterID.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "cluster_id")
	}
	return clusterIDStr, nil
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

var privateNetworkIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		// If external-name is empty, the resource hasn't been created yet,
		// so we should return empty string instead of constructing an incomplete ID
		if externalName == "" {
			return "", nil
		}
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}
		return fmt.Sprintf("%s/%s", serviceName, externalName), nil
	},
	DisableNameInitializer: true,
}

var subnetIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		// If external-name is empty, the resource hasn't been created yet,
		// so we should return empty string instead of constructing an incomplete ID
		if externalName == "" {
			return "", nil
		}
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		networkID, ok := parameters["network_id"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "network_id")
		}
		networkIDStr, ok := networkID.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "network_id")
		}

		return fmt.Sprintf("%s/%s/%s", serviceName, networkIDStr, externalName), nil
	},
	DisableNameInitializer: true,
}

var userIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		id, err := config.IDAsExternalName(tfstate)
		if err != nil {
			return "", err
		}
		return userID(id)
	},
	GetIDFn: func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		// If external-name is empty, the resource hasn't been created yet,
		// so we should return empty string instead of constructing an incomplete ID
		if externalName == "" {
			return "", nil
		}
		return userID(externalName)
	},
	DisableNameInitializer: true,
}

func userID(externalName string) (string, error) {
	id := externalName
	if parts := strings.SplitN(id, "/", 2); len(parts) == 2 {
		id = parts[1]
	}
	if id == "" || strings.Contains(id, "/") {
		return "", errors.Errorf("user external-name %q is not <user_id> or <service_name>/<user_id>", externalName)
	}
	return id, nil
}

var postgresqlUserIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		id, err := config.IDAsExternalName(tfstate)
		if err != nil {
			return "", err
		}
		return postgresqlUserID(id)
	},
	// The upstream Read puts the Terraform ID straight into the request path,
	// so it must be the bare user id. Orphans left by a create that failed
	// with 409/500 are recovered by name in the upstream create since 2.23.0.
	GetIDFn: func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if externalName == "" {
			return "", nil
		}
		return postgresqlUserID(externalName)
	},
	DisableNameInitializer: true,
}

// postgresqlUserID accepts the bare user id or the legacy
// <service_name>/<cluster_id>/<user_id> external-name and returns the user id.
func postgresqlUserID(externalName string) (string, error) {
	id := externalName
	if parts := strings.SplitN(id, "/", 3); len(parts) == 3 {
		id = parts[2]
	}
	if id == "" || strings.Contains(id, "/") {
		return "", errors.Errorf("postgresql user external-name %q is not <user_id> or <service_name>/<cluster_id>/<user_id>", externalName)
	}
	return id, nil
}

// ipFirewallID builds the deterministic <ip>/<ip_on_firewall> identity of an
// ovh_ip_firewall, or <ip>/<ip_on_firewall>/<sequence> of an
// ovh_ip_firewall_rule when withSequence is set.
func ipFirewallID(attrs map[string]any, withSequence bool) (string, error) {
	ip, ok := attrs["ip"].(string)
	if !ok || ip == "" {
		return "", errors.Errorf(ErrFmtNoAttribute, "ip")
	}
	ipOnFirewall, ok := attrs["ip_on_firewall"].(string)
	if !ok || ipOnFirewall == "" {
		return "", errors.Errorf(ErrFmtNoAttribute, "ip_on_firewall")
	}
	if !withSequence {
		return fmt.Sprintf("%s/%s", ip, ipOnFirewall), nil
	}
	// spec.forProvider decodes numbers as float64; the Framework state as int64.
	var sequence int64
	switch v := attrs["sequence"].(type) {
	case float64:
		sequence = int64(v)
	case int64:
		sequence = v
	default:
		return "", errors.Errorf(ErrFmtNoAttribute, "sequence")
	}
	return fmt.Sprintf("%s/%s/%d", ip, ipOnFirewall, sequence), nil
}

// ipFirewallIdentifier is shared by ovh_ip_firewall and ovh_ip_firewall_rule.
// Their upstream Read returns an error on a 404 and does not set the id when
// adopting an existing object, so the identity is rebuilt from the arguments
// and a 404 is reported as "resource does not exist" to let Create proceed.
func ipFirewallIdentifier(withSequence bool) config.ExternalName {
	return config.ExternalName{
		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
		GetExternalNameFn: func(tfstate map[string]any) (string, error) {
			return ipFirewallID(tfstate, withSequence)
		},
		GetIDFn: func(_ context.Context, _ string, parameters map[string]any, _ map[string]any) (string, error) {
			return ipFirewallID(parameters, withSequence)
		},
		DisableNameInitializer: true,
		IsNotFoundDiagnosticFn: isOVHNotFoundError,
	}
}

// serviceNameIdentifierFromProvider builds the "<service_name>/<id>" Terraform
// identifier used by resources scoped to a single service.
var serviceNameIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		if externalName == "" {
			return "", nil
		}
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}
		return fmt.Sprintf("%s/%s", serviceName, externalName), nil
	},
	DisableNameInitializer: true,
}

// databaseClusterIdentifierFromProvider builds the
// "<service_name>/<cluster_id>/<id>" Terraform identifier used by resources
// that hang off a managed database cluster.
var databaseClusterIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		if externalName == "" {
			return "", nil
		}
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}
		clusterID, err := clusterID(parameters)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s/%s/%s", serviceName, clusterID, externalName), nil
	},
	DisableNameInitializer: true,
}

// databaseLogSubscriptionIdentifierFromProvider builds the
// "<service_name>/<engine>/<cluster_id>/<id>" Terraform identifier used by
// ovh_cloud_project_database_log_subscription.
var databaseLogSubscriptionIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		if externalName == "" {
			return "", nil
		}
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}
		clusterID, err := clusterID(parameters)
		if err != nil {
			return "", err
		}
		engine, ok := parameters["engine"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "engine")
		}
		engineStr, ok := engine.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "engine")
		}
		return fmt.Sprintf("%s/%s/%s/%s", serviceName, engineStr, clusterID, externalName), nil
	},
	DisableNameInitializer: true,
}

// TerraformPluginSDKExternalNameConfigs contains all external name
// configurations for Terraform Plugin SDK resources
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"ovh_cloud_project_network_private":           privateNetworkIdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet":    subnetIdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet_v2": subnetIdentifierFromProvider,

	"ovh_cloud_project_workflow_backup": config.NameAsIdentifier,
	"ovh_cloud_project":                 config.IdentifierFromProvider,
	// The ovh_cloud_project_alerting resource uses a nested type which is not supported yet in upjet.
	// there is an open issue in upjet regarding this issue: https://github.com/crossplane/upjet/v2/issues/372
	// "ovh_cloud_project_alerting":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_user":                                         userIdentifierFromProvider,
	"ovh_cloud_project_user_s3_credential":                           config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_policy":                               config.IdentifierFromProvider,
	"ovh_iam_policy":                                                 config.IdentifierFromProvider,
	"ovh_iam_resource_group":                                         config.IdentifierFromProvider,
	"ovh_iam_permissions_group":                                      config.IdentifierFromProvider,
	"ovh_iam_resource_tags":                                          config.IdentifierFromProvider,
	"ovh_me_identity_group":                                          config.IdentifierFromProvider,
	"ovh_me_identity_user":                                           config.IdentifierFromProvider,
	"ovh_me_ssh_key":                                                 config.IdentifierFromProvider,
	"ovh_me_api_oauth2_client":                                       config.IdentifierFromProvider,
	"ovh_cloud_project_failover_ip_attach":                           config.IdentifierFromProvider,
	"ovh_ip_reverse":                                                 config.IdentifierFromProvider,
	"ovh_ip_move":                                                    config.IdentifierFromProvider,
	"ovh_ip_service":                                                 config.IdentifierFromProvider,
	"ovh_dedicated_ceph_acl":                                         config.IdentifierFromProvider,
	"ovh_dedicated_server_install_task":                              config.IdentifierFromProvider,
	"ovh_dedicated_server_reboot_task":                               config.IdentifierFromProvider,
	"ovh_dedicated_server_update":                                    config.IdentifierFromProvider,
	"ovh_me_installation_template":                                   config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme":                  config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme_hardware_raid":    config.IdentifierFromProvider,
	"ovh_me_installation_template_partition_scheme_partition":        config.IdentifierFromProvider,
	"ovh_me_ipxe_script":                                             config.IdentifierFromProvider,
	"ovh_dedicated_server_networking":                                config.IdentifierFromProvider,
	"ovh_dedicated_server_virtual_mac":                               config.IdentifierFromProvider,
	"ovh_domain_zone":                                                config.IdentifierFromProvider,
	"ovh_domain_zone_record":                                         config.IdentifierFromProvider,
	"ovh_domain_zone_redirection":                                    config.IdentifierFromProvider,
	"ovh_iploadbalancing":                                            config.IdentifierFromProvider,
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
	"ovh_dbaas_logs_cluster":                                         config.IdentifierFromProvider,
	"ovh_dbaas_logs_graylog_output_stream":                           config.IdentifierFromProvider,
	"ovh_dbaas_logs_input":                                           config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_alias":                         config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_index":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_clickhouse_user":                     databaseClusterIdentifierFromProvider,
	"ovh_cloud_project_database_log_subscription":                    databaseLogSubscriptionIdentifierFromProvider,
	"ovh_cloud_project_database_database":                            config.IdentifierFromProvider,
	"ovh_cloud_project_database_integration":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_acl":                           config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_schemaregistryacl":             config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_topic":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_mongodb_user":                        config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_pattern":                  config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_user":                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_user":                     postgresqlUserIdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_connection_pool":          config.IdentifierFromProvider,
	"ovh_cloud_project_database_user":                                config.IdentifierFromProvider,
	"ovh_cloud_project_kube":                                         kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_iprestrictions":                          kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_nodepool":                                kubePoolIdentifierFromProvider,
	"ovh_cloud_project_kube_log_subscription":                        kubePoolIdentifierFromProvider,
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
	"ovh_cloud_project_gateway":                                      config.IdentifierFromProvider,
	// Added resources from provider versions 1.2.0 through 2.9.0
	"ovh_savings_plan":                              config.IdentifierFromProvider,
	"ovh_cloud_project_instance":                    config.IdentifierFromProvider,
	"ovh_cloud_project_database_prometheus":         config.IdentifierFromProvider,
	"ovh_cloud_project_database_mongodb_prometheus": config.IdentifierFromProvider,
	"ovh_domain_name_servers":                       config.IdentifierFromProvider,
	"ovh_domain_ds_records":                         config.IdentifierFromProvider,
	// Framework resources moved to TerraformPluginFrameworkExternalNameConfigs
	"ovh_dbaas_logs_role":                     config.IdentifierFromProvider,
	"ovh_dbaas_logs_role_permission_stream":   config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_graylog_stream":    serviceNameIdentifierFromProvider,
	"ovh_vrack_ipv6":                          config.IdentifierFromProvider,
	"ovh_vrack_vrackservices":                 config.IdentifierFromProvider,
	"ovh_dedicated_server_reinstall_task":     config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_iam": config.IdentifierFromProvider,
	"ovh_cloud_project_database_valkey_user":  config.IdentifierFromProvider,
}

// sdkResourceConfigurator applies all external name configs from the
// TerraformPluginSDKExternalNameConfigs table.
func sdkResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// TerraformPluginSDKResourceList returns the list of all resources that have
// external name configured in TerraformPluginSDKExternalNameConfigs
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// frameworkResourceConfigurator applies all external name configs from the
// TerraformPluginFrameworkExternalNameConfigs table.
func frameworkResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		if e, configured := TerraformPluginFrameworkExternalNameConfigs[r.Name]; configured {
			r.ExternalName = e
		}
	}
}

// TerraformPluginFrameworkExternalNameConfigs contains external name
// configurations for resources using Terraform Plugin Framework.
// These resources use the native Plugin Framework provider and support
// features like IsNotFoundDiagnosticFn.
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	// Cloud Project Framework resources
	"ovh_cloud_project_gateway_interface": config.IdentifierFromProvider,
	"ovh_cloud_project_instance_snapshot": config.IdentifierFromProvider,
	"ovh_cloud_project_loadbalancer":      config.IdentifierFromProvider,
	"ovh_cloud_project_rancher":           config.IdentifierFromProvider,
	"ovh_cloud_project_region":            config.IdentifierFromProvider,
	"ovh_cloud_project_region_network":    config.IdentifierFromProvider,
	"ovh_cloud_project_ssh_key":           config.IdentifierFromProvider,
	// ovh_cloud_project_storage needs IsNotFoundDiagnosticFn for 404 handling
	"ovh_cloud_project_storage": {
		// Use name as the external identifier
		SetIdentifierArgumentFn: func(base map[string]any, externalName string) {
			base["name"] = externalName
		},
		GetExternalNameFn: func(tfstate map[string]any) (string, error) {
			// Extract the name from tfstate
			if name, ok := tfstate["name"].(string); ok && name != "" {
				return name, nil
			}
			return "", errors.New("name field not found or empty in tfstate")
		},
		GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
			// Construct composite ID: <service_name>/<region_name>/<storage_name>
			serviceName, ok := parameters["service_name"].(string)
			if !ok {
				return "", errors.Errorf(ErrFmtNoAttribute, "service_name")
			}
			regionName, ok := parameters["region_name"].(string)
			if !ok {
				return "", errors.Errorf(ErrFmtNoAttribute, "region_name")
			}
			return fmt.Sprintf("%s/%s/%s", serviceName, regionName, externalName), nil
		},
		OmittedFields: []string{
			"limit",
			"marker",
			"prefix",
		},
		IsNotFoundDiagnosticFn: isOVHNotFoundError,
	},
	"ovh_cloud_project_volume":        config.IdentifierFromProvider,
	"ovh_cloud_project_volume_backup": config.IdentifierFromProvider,

	// DBaaS Logs Framework resources
	"ovh_dbaas_logs_token": config.IdentifierFromProvider,

	// Dedicated Server Framework resources
	"ovh_dedicated_server": config.IdentifierFromProvider,

	// Domain Framework resources
	"ovh_domain_name":                config.IdentifierFromProvider,
	"ovh_domain_zone_dnssec":         config.IdentifierFromProvider,
	"ovh_domain_zone_dynhost_login":  config.IdentifierFromProvider,
	"ovh_domain_zone_dynhost_record": config.IdentifierFromProvider,
	"ovh_domain_zone_import":         config.IdentifierFromProvider,

	// IP Framework resources
	"ovh_ip_firewall":      ipFirewallIdentifier(false),
	"ovh_ip_firewall_rule": ipFirewallIdentifier(true),
	"ovh_ip_mitigation":    config.IdentifierFromProvider,

	// IP Load Balancing Framework resources
	"ovh_iploadbalancing_ssl":             config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm":        config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm_server": config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_frontend":    config.IdentifierFromProvider,

	// OKMS Framework resources
	"ovh_okms":                 config.IdentifierFromProvider,
	"ovh_okms_credential":      config.IdentifierFromProvider,
	"ovh_okms_secret":          config.IdentifierFromProvider,
	"ovh_okms_service_key":     config.IdentifierFromProvider,
	"ovh_okms_service_key_jwk": config.IdentifierFromProvider,

	// OVH Cloud Connect Framework resources
	"ovh_ovhcloud_connect_pop_config":                  config.IdentifierFromProvider,
	"ovh_ovhcloud_connect_pop_datacenter_config":       config.IdentifierFromProvider,
	"ovh_ovhcloud_connect_pop_datacenter_extra_config": config.IdentifierFromProvider,

	// Storage EFS Framework resources
	"ovh_storage_efs_share":          config.IdentifierFromProvider,
	"ovh_storage_efs_share_acl":      config.IdentifierFromProvider,
	"ovh_storage_efs_share_snapshot": config.IdentifierFromProvider,

	// VPS Framework resource
	"ovh_vps": config.IdentifierFromProvider,

	// vRack Framework resources
	"ovh_vrack_dedicated_cloud":            config.IdentifierFromProvider,
	"ovh_vrack_dedicated_cloud_datacenter": config.IdentifierFromProvider,
	"ovh_vrack_ipv6_routed_subrange":       config.IdentifierFromProvider,
	"ovh_vrack_ovhcloudconnect":            config.IdentifierFromProvider,
	"ovh_vrack_public_routing_priority":    config.IdentifierFromProvider,
	"ovh_vrackservices":                    config.IdentifierFromProvider,

	// Public Cloud Framework resources
	"ovh_cloud_floating_ip":              config.IdentifierFromProvider,
	"ovh_cloud_gateway":                  config.IdentifierFromProvider,
	"ovh_cloud_instance":                 config.IdentifierFromProvider,
	"ovh_cloud_instance_group":           config.IdentifierFromProvider,
	"ovh_cloud_instance_snapshot":        config.IdentifierFromProvider,
	"ovh_cloud_loadbalancer":             config.IdentifierFromProvider,
	"ovh_cloud_loadbalancer_l7policy":    config.IdentifierFromProvider,
	"ovh_cloud_loadbalancer_listener":    config.IdentifierFromProvider,
	"ovh_cloud_loadbalancer_pool":        config.IdentifierFromProvider,
	"ovh_cloud_loadbalancer_pool_member": config.IdentifierFromProvider,
	"ovh_cloud_quota":                    config.IdentifierFromProvider,
	"ovh_cloud_security_group":           config.IdentifierFromProvider,

	// Public Cloud vRack network Framework resources
	"ovh_cloud_network_private_vrack":        config.IdentifierFromProvider,
	"ovh_cloud_network_private_vrack_subnet": config.IdentifierFromProvider,

	// Key Management Service (Key Manager) Framework resources
	"ovh_cloud_key_manager_container":          config.IdentifierFromProvider,
	"ovh_cloud_key_manager_container_consumer": config.IdentifierFromProvider,
	"ovh_cloud_key_manager_secret":             config.IdentifierFromProvider,
	"ovh_cloud_key_manager_secret_consumer":    config.IdentifierFromProvider,

	// Block and file storage Framework resources
	"ovh_cloud_storage_block_volume":                                  config.IdentifierFromProvider,
	"ovh_cloud_storage_block_volume_backup":                           config.IdentifierFromProvider,
	"ovh_cloud_storage_block_volume_snapshot":                         config.IdentifierFromProvider,
	"ovh_cloud_storage_file_share":                                    config.IdentifierFromProvider,
	"ovh_cloud_storage_file_share_acl":                                config.IdentifierFromProvider,
	"ovh_cloud_storage_file_share_network":                            config.IdentifierFromProvider,
	"ovh_cloud_storage_file_share_snapshot":                           config.IdentifierFromProvider,
	"ovh_cloud_storage_object_bucket":                                 config.IdentifierFromProvider,
	"ovh_cloud_project_file_storage_share":                            config.IdentifierFromProvider,
	"ovh_cloud_project_file_storage_share_network":                    config.IdentifierFromProvider,
	"ovh_cloud_project_storage_object_bucket_lifecycle_configuration": config.IdentifierFromProvider,
	"ovh_cloud_project_storage_replication_job":                       config.IdentifierFromProvider,
	"ovh_storage_efs":                                                 config.IdentifierFromProvider,

	// DBaaS Logs Framework resources
	"ovh_dbaas_logs_encryption_key": config.IdentifierFromProvider,

	// Email domain Framework resources
	"ovh_email_domain_account":     config.IdentifierFromProvider,
	"ovh_email_domain_dkim":        config.IdentifierFromProvider,
	"ovh_email_domain_redirection": config.IdentifierFromProvider,

	// Cloud Disk Array (dedicated Ceph) Framework resources
	"ovh_dedicated_ceph_pool": config.IdentifierFromProvider,
	"ovh_dedicated_ceph_user": config.IdentifierFromProvider,

	// Web Cloud private database Framework resources
	"ovh_hosting_privatedatabase_webhosting_network": config.IdentifierFromProvider,

	// Account log subscription Framework resources
	"ovh_iam_log_subscription":           config.IdentifierFromProvider,
	"ovh_me_api_log_subscription":        config.IdentifierFromProvider,
	"ovh_me_logs_audit_log_subscription": config.IdentifierFromProvider,

	// Framework resources whose schema exposes no "id" attribute; their
	// Terraform identity is the user-supplied name.
	"ovh_cloud_ssh_key":          config.NameAsIdentifier,
	"ovh_me_identity_user_token": config.NameAsIdentifier,
}

// CLIReconciledExternalNameConfigs contains external name configurations
// for resources to be reconciled using the Terraform CLI.
var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// TerraformPluginFrameworkResourceList returns the list of resources that have
// external name configured and to be reconciled under the Terraform Plugin
// Framework architecture.
func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// CLIReconciledResourceList returns the list of resources to be reconciled
// using Terraform CLI.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
