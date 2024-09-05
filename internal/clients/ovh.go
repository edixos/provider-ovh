/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/edixos/provider-ovh/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal ovh credentials as JSON"
	errNoValidCredentials   = "cannot find valid credentials"
)

type OVHCredentials struct {
	Endpoint          string `json:"endpoint"`
	ApplicationKey    string `json:"application_key,omitempty"`
	ApplicationSecret string `json:"application_secret,omitempty"`
	ConsumerKey       string `json:"consumer_key,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
	ClientSecret      string `json:"client_secret,omitempty"`
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		cfg, err := configureClient(ctx, pc, client)
		if err != nil {
			return ps, err
		}

		ps.Configuration = cfg

		return ps, nil
	}
}

func configureClient(ctx context.Context, pc *v1beta1.ProviderConfig, client client.Client) (map[string]any, error) {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}

	var creds OVHCredentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}

	// Set credentials in Terraform provider configuration.
	cfg := map[string]any{
		"endpoint": creds.Endpoint,
	}

	if creds.ApplicationKey != "" && creds.ApplicationSecret != "" && creds.ConsumerKey != "" {
		cfg["application_key"] = creds.ApplicationKey
		cfg["application_secret"] = creds.ApplicationSecret
		cfg["consumer_key"] = creds.ConsumerKey
		return cfg, nil
	}

	if creds.ClientID != "" && creds.ClientSecret != "" {
		cfg["client_id"] = creds.ClientID
		cfg["client_secret"] = creds.ClientSecret
		return cfg, nil
	}

	return cfg, errors.New(errNoValidCredentials)
}
