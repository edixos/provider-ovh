/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ContainerRegistryOIDCInitParameters struct {
	DeleteUsers *bool `json:"deleteUsers,omitempty" tf:"delete_users,omitempty"`

	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/registry/v1alpha1.ContainerRegistry
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a ContainerRegistry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a ContainerRegistry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ContainerRegistryOIDCObservation struct {
	DeleteUsers *bool `json:"deleteUsers,omitempty" tf:"delete_users,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ContainerRegistryOIDCParameters struct {

	// +kubebuilder:validation:Optional
	DeleteUsers *bool `json:"deleteUsers,omitempty" tf:"delete_users,omitempty"`

	// +kubebuilder:validation:Optional
	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	// +kubebuilder:validation:Optional
	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	// +kubebuilder:validation:Optional
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// +kubebuilder:validation:Optional
	OidcClientSecretSecretRef v1.SecretKeySelector `json:"oidcClientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// +kubebuilder:validation:Optional
	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	// +kubebuilder:validation:Optional
	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	// +kubebuilder:validation:Optional
	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	// +kubebuilder:validation:Optional
	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/registry/v1alpha1.ContainerRegistry
	// +kubebuilder:validation:Optional
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a ContainerRegistry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a ContainerRegistry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ContainerRegistryOIDCSpec defines the desired state of ContainerRegistryOIDC
type ContainerRegistryOIDCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerRegistryOIDCParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ContainerRegistryOIDCInitParameters `json:"initProvider,omitempty"`
}

// ContainerRegistryOIDCStatus defines the observed state of ContainerRegistryOIDC.
type ContainerRegistryOIDCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerRegistryOIDCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ContainerRegistryOIDC is the Schema for the ContainerRegistryOIDCs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ContainerRegistryOIDC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oidcClientId) || (has(self.initProvider) && has(self.initProvider.oidcClientId))",message="spec.forProvider.oidcClientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oidcClientSecretSecretRef)",message="spec.forProvider.oidcClientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oidcEndpoint) || (has(self.initProvider) && has(self.initProvider.oidcEndpoint))",message="spec.forProvider.oidcEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oidcName) || (has(self.initProvider) && has(self.initProvider.oidcName))",message="spec.forProvider.oidcName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oidcScope) || (has(self.initProvider) && has(self.initProvider.oidcScope))",message="spec.forProvider.oidcScope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ContainerRegistryOIDCSpec   `json:"spec"`
	Status ContainerRegistryOIDCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRegistryOIDCList contains a list of ContainerRegistryOIDCs
type ContainerRegistryOIDCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerRegistryOIDC `json:"items"`
}

// Repository type metadata.
var (
	ContainerRegistryOIDC_Kind             = "ContainerRegistryOIDC"
	ContainerRegistryOIDC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerRegistryOIDC_Kind}.String()
	ContainerRegistryOIDC_KindAPIVersion   = ContainerRegistryOIDC_Kind + "." + CRDGroupVersion.String()
	ContainerRegistryOIDC_GroupVersionKind = CRDGroupVersion.WithKind(ContainerRegistryOIDC_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerRegistryOIDC{}, &ContainerRegistryOIDCList{})
}
