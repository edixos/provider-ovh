/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestUserIdentifierFromProvider(t *testing.T) {
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "",
		},
		"HappyPath": {
			externalName: "user-123",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "user-123",
		},
		"IgnoresMissingServiceName": {
			externalName: "user-123",
			params:       map[string]any{},
			want:         "user-123",
		},
		"IgnoresWrongTypeServiceName": {
			externalName: "user-123",
			params:       map[string]any{"service_name": 42},
			want:         "user-123",
		},
		"LegacyCompositeExternalName": {
			externalName: "svc-1/user-123",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "user-123",
		},
		"EmptyUserID": {
			externalName: "svc-1/",
			params:       map[string]any{"service_name": "svc-1"},
			wantErr:      "is not <user_id> or <service_name>/<user_id>",
		},
		"TrailingSlash": {
			externalName: "svc-1/user-123/",
			params:       map[string]any{"service_name": "svc-1"},
			wantErr:      "is not <user_id> or <service_name>/<user_id>",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := userIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestUserExternalNameFromState(t *testing.T) {
	cases := map[string]struct {
		tfstate map[string]any
		want    string
		wantErr bool
	}{
		"BareID": {
			tfstate: map[string]any{"id": "user-123"},
			want:    "user-123",
		},
		"LegacyComposite": {
			tfstate: map[string]any{"id": "svc-1/user-123"},
			want:    "user-123",
		},
		"MissingID": {
			tfstate: map[string]any{},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := userIdentifierFromProvider.GetExternalNameFn(tc.tfstate)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetExternalNameFn() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.want)
			}
		})
	}
}

// Terraform's Read for kube clusters uses d.Id() as the cluster ID, so the ID must be
// bare. A <service_name>/<kube_id> ID makes Read call /kube/<service_name>/<kube_id>,
// get a 404 and report the cluster as gone.
func TestKubeIdentifierFromProvider(t *testing.T) {
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "",
		},
		"HappyPath": {
			externalName: "kube-1",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "kube-1",
		},
		"IgnoresMissingServiceName": {
			externalName: "kube-1",
			params:       map[string]any{},
			want:         "kube-1",
		},
		"LegacyCompositeExternalName": {
			externalName: "svc-1/kube-1",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "kube-1",
		},
		"EmptyKubeID": {
			externalName: "svc-1/",
			params:       map[string]any{"service_name": "svc-1"},
			wantErr:      "is not <id> or its <service_name>/.../<id> form",
		},
		"TooManySegments": {
			externalName: "svc-1/kube-1/extra",
			params:       map[string]any{"service_name": "svc-1"},
			wantErr:      "is not <id> or its <service_name>/.../<id> form",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := kubeIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestKubeExternalNameFromState(t *testing.T) {
	cases := map[string]struct {
		tfstate map[string]any
		want    string
		wantErr bool
	}{
		"BareID": {
			tfstate: map[string]any{"id": "kube-1"},
			want:    "kube-1",
		},
		"LegacyComposite": {
			tfstate: map[string]any{"id": "svc-1/kube-1"},
			want:    "kube-1",
		},
		"MissingID": {
			tfstate: map[string]any{},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := kubeIdentifierFromProvider.GetExternalNameFn(tc.tfstate)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetExternalNameFn() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.want)
			}
		})
	}
}

// Node pools and log subscriptions: Terraform's Read takes service_name and kube_id
// from the parameters and uses d.Id() as the last path segment, so the ID must be bare.
func TestKubePoolIdentifierFromProvider(t *testing.T) {
	params := map[string]any{"service_name": "svc-1", "kube_id": "kube-1"}
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       params,
			want:         "",
		},
		"HappyPath": {
			externalName: "pool-1",
			params:       params,
			want:         "pool-1",
		},
		"IgnoresMissingKubeID": {
			externalName: "pool-1",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "pool-1",
		},
		"KubeIDPrefix": {
			externalName: "kube-1/pool-1",
			params:       params,
			want:         "pool-1",
		},
		"LegacyCompositeExternalName": {
			externalName: "svc-1/kube-1/pool-1",
			params:       params,
			want:         "pool-1",
		},
		"EmptyPoolID": {
			externalName: "svc-1/kube-1/",
			params:       params,
			wantErr:      "is not <id> or its <service_name>/.../<id> form",
		},
		"TooManySegments": {
			externalName: "a/svc-1/kube-1/pool-1",
			params:       params,
			wantErr:      "is not <id> or its <service_name>/.../<id> form",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := kubePoolIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestKubePoolExternalNameFromState(t *testing.T) {
	cases := map[string]struct {
		tfstate map[string]any
		want    string
		wantErr bool
	}{
		"BareID": {
			tfstate: map[string]any{"id": "pool-1"},
			want:    "pool-1",
		},
		"LegacyComposite": {
			tfstate: map[string]any{"id": "svc-1/kube-1/pool-1"},
			want:    "pool-1",
		},
		"MissingID": {
			tfstate: map[string]any{},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := kubePoolIdentifierFromProvider.GetExternalNameFn(tc.tfstate)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetExternalNameFn() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestServiceNameIdentifierFromProvider(t *testing.T) {
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "",
		},
		"HappyPath": {
			externalName: "stream-1",
			params:       map[string]any{"service_name": "svc-1"},
			want:         "svc-1/stream-1",
		},
		"MissingServiceName": {
			externalName: "stream-1",
			params:       map[string]any{},
			wantErr:      "service_name",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := serviceNameIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestDatabaseClusterIdentifierFromProvider(t *testing.T) {
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1"},
			want:         "",
		},
		"HappyPath": {
			externalName: "user-1",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1"},
			want:         "svc-1/cluster-1/user-1",
		},
		"MissingClusterID": {
			externalName: "user-1",
			params:       map[string]any{"service_name": "svc-1"},
			wantErr:      "cluster_id",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := databaseClusterIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestDatabaseLogSubscriptionIdentifierFromProvider(t *testing.T) {
	cases := map[string]struct {
		externalName string
		params       map[string]any
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1", "engine": "kafka"},
			want:         "",
		},
		"HappyPath": {
			externalName: "sub-1",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1", "engine": "kafka"},
			want:         "svc-1/kafka/cluster-1/sub-1",
		},
		"MissingEngine": {
			externalName: "sub-1",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1"},
			wantErr:      "engine",
		},
		"WrongTypeEngine": {
			externalName: "sub-1",
			params:       map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1", "engine": 42},
			wantErr:      "engine",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := databaseLogSubscriptionIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, tc.params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestSDKMapBindingsRestored(t *testing.T) {
	cases := map[string]string{
		"ovh_cloud_project_user":                     "u",
		"ovh_cloud_project_database_postgresql_user": "u",
		"ovh_cloud_project_database_clickhouse_user": "svc-1/cluster-1/u",
		// service_name/engine/cluster_id/id
		"ovh_cloud_project_database_log_subscription": "svc-1/postgresql/cluster-1/u",
		// bare subscription_id: Terraform's Read uses d.Id() as the last path segment
		"ovh_cloud_project_kube_log_subscription": "u",
		// service_name/stream_id
		"ovh_dbaas_logs_output_graylog_stream": "svc-1/u",
	}

	params := map[string]any{
		"service_name": "svc-1",
		"cluster_id":   "cluster-1",
		"user_id":      "uid",
		"region":       "GRA9",
		"engine":       "postgresql",
		"kube_id":      "kube-1",
	}

	for name, want := range cases {
		t.Run(name, func(t *testing.T) {
			cfg, ok := TerraformPluginSDKExternalNameConfigs[name]
			if !ok {
				t.Fatalf("%s missing from TerraformPluginSDKExternalNameConfigs", name)
			}
			got, err := cfg.GetIDFn(context.Background(), "u", params, nil)
			if err != nil {
				t.Fatalf("%s: unexpected error: %v", name, err)
			}
			if got != want {
				t.Errorf("%s: got %q, want %q (regression — entry likely set to IdentifierFromProvider)", name, got, want)
			}
		})
	}
}

// The kube entries return the bare ID, which IdentifierFromProvider would too for a
// bare external-name. A composite (legacy) external-name tells them apart:
// IdentifierFromProvider passes it through unchanged, which breaks Read.
func TestKubeMapBindingsStripCompositeExternalName(t *testing.T) {
	cases := map[string]string{
		"ovh_cloud_project_kube":                  "svc-1/kube-1",
		"ovh_cloud_project_kube_iprestrictions":   "svc-1/kube-1",
		"ovh_cloud_project_kube_nodepool":         "svc-1/kube-1/pool-1",
		"ovh_cloud_project_kube_log_subscription": "svc-1/kube-1/sub-1",
	}
	params := map[string]any{"service_name": "svc-1", "kube_id": "kube-1"}

	for name, externalName := range cases {
		t.Run(name, func(t *testing.T) {
			cfg, ok := TerraformPluginSDKExternalNameConfigs[name]
			if !ok {
				t.Fatalf("%s missing from TerraformPluginSDKExternalNameConfigs", name)
			}
			got, err := cfg.GetIDFn(context.Background(), externalName, params, nil)
			if err != nil {
				t.Fatalf("%s: unexpected error: %v", name, err)
			}
			want := externalName[strings.LastIndex(externalName, "/")+1:]
			if got != want {
				t.Errorf("%s: got %q, want bare %q", name, got, want)
			}
		})
	}
}

func TestPostgresqlUserIdentifierFromProvider(t *testing.T) {
	params := map[string]any{"service_name": "svc-1", "cluster_id": "cluster-1", "name": "johndoe"}
	cases := map[string]struct {
		externalName string
		want         string
		wantErr      string
	}{
		"EmptyExternalNameReturnsEmpty": {
			externalName: "",
			want:         "",
		},
		"BareID": {
			externalName: "user-123",
			want:         "user-123",
		},
		"LegacyCompositeExternalName": {
			externalName: "svc-1/cluster-1/user-123",
			want:         "user-123",
		},
		"EmptyUserID": {
			externalName: "svc-1/cluster-1/",
			wantErr:      "is not <user_id> or <service_name>/<cluster_id>/<user_id>",
		},
		"TrailingSlash": {
			externalName: "svc-1/cluster-1/user-123/",
			wantErr:      "is not <user_id> or <service_name>/<cluster_id>/<user_id>",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := postgresqlUserIdentifierFromProvider.GetIDFn(context.Background(), tc.externalName, params, nil)
			assertGetID(t, got, err, tc.want, tc.wantErr)
		})
	}
}

func TestPostgresqlUserExternalNameFromState(t *testing.T) {
	cases := map[string]struct {
		tfstate map[string]any
		want    string
		wantErr bool
	}{
		"BareID": {
			tfstate: map[string]any{"id": "user-123"},
			want:    "user-123",
		},
		"LegacyComposite": {
			tfstate: map[string]any{"id": "svc-1/cluster-1/user-123"},
			want:    "user-123",
		},
		"MissingID": {
			tfstate: map[string]any{},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := postgresqlUserIdentifierFromProvider.GetExternalNameFn(tc.tfstate)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetExternalNameFn() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestIPFirewallIdentifier(t *testing.T) {
	firewall := TerraformPluginFrameworkExternalNameConfigs["ovh_ip_firewall"]
	rule := TerraformPluginFrameworkExternalNameConfigs["ovh_ip_firewall_rule"]

	t.Run("FirewallIDFromParameters", func(t *testing.T) {
		got, err := firewall.GetIDFn(context.Background(), "", map[string]any{
			"ip":             "192.0.2.10/32",
			"ip_on_firewall": "192.0.2.10",
		}, nil)
		assertGetID(t, got, err, "192.0.2.10/32/192.0.2.10", "")
	})

	t.Run("FirewallMissingIPOnFirewall", func(t *testing.T) {
		got, err := firewall.GetIDFn(context.Background(), "", map[string]any{"ip": "192.0.2.10"}, nil)
		assertGetID(t, got, err, "", "ip_on_firewall")
	})

	t.Run("FirewallExternalNameFromStateWithoutID", func(t *testing.T) {
		got, err := firewall.GetExternalNameFn(map[string]any{
			"id":             "",
			"ip":             "192.0.2.10",
			"ip_on_firewall": "192.0.2.10",
		})
		assertGetID(t, got, err, "192.0.2.10/192.0.2.10", "")
	})

	t.Run("RuleIDFromParameters", func(t *testing.T) {
		// spec.forProvider numbers decode as float64
		got, err := rule.GetIDFn(context.Background(), "", map[string]any{
			"ip":             "192.0.2.10",
			"ip_on_firewall": "192.0.2.10",
			"sequence":       float64(0),
		}, nil)
		assertGetID(t, got, err, "192.0.2.10/192.0.2.10/0", "")
	})

	t.Run("RuleExternalNameFromState", func(t *testing.T) {
		// Framework state integers convert to int64
		got, err := rule.GetExternalNameFn(map[string]any{
			"ip":             "192.0.2.10",
			"ip_on_firewall": "192.0.2.10",
			"sequence":       int64(18),
		})
		assertGetID(t, got, err, "192.0.2.10/192.0.2.10/18", "")
	})

	t.Run("RuleMissingSequence", func(t *testing.T) {
		got, err := rule.GetIDFn(context.Background(), "", map[string]any{
			"ip":             "192.0.2.10",
			"ip_on_firewall": "192.0.2.10",
		}, nil)
		assertGetID(t, got, err, "", "sequence")
	})

	t.Run("NotFoundDiagnosticTreatedAsMissing", func(t *testing.T) {
		diags := []*tfprotov6.Diagnostic{{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error calling Get /ip/192.0.2.10/firewall/192.0.2.10/rule/0",
			Detail:   `OVHcloud API error (status code 404): Client::NotFound: "The requested object (sequence = 0) does not exist"`,
		}}
		if !firewall.IsNotFoundDiagnosticFn(diags) {
			t.Error("firewall: 404 diagnostic not treated as resource not found")
		}
		if !rule.IsNotFoundDiagnosticFn(diags) {
			t.Error("rule: 404 diagnostic not treated as resource not found")
		}
	})
}

func assertGetID(t *testing.T, got string, err error, want, wantErr string) {
	t.Helper()
	if wantErr != "" {
		if err == nil {
			t.Fatalf("expected error containing %q, got nil (id=%q)", wantErr, got)
		}
		if !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("error %q does not contain %q", err.Error(), wantErr)
		}
		return
	}
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
