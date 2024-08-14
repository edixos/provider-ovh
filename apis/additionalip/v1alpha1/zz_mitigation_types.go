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

type MitigationInitParameters struct {

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	IPOnMitigation *string `json:"ipOnMitigation,omitempty" tf:"ip_on_mitigation,omitempty"`

	// Set on true if your ip is on permanent mitigation
	Permanent *bool `json:"permanent,omitempty" tf:"permanent,omitempty"`
}

type MitigationObservation struct {

	// Set on true if your ip is on auto-mitigation
	Auto *bool `json:"auto,omitempty" tf:"auto,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	IPOnMitigation *string `json:"ipOnMitigation,omitempty" tf:"ip_on_mitigation,omitempty"`

	// Set on true if your ip is on permanent mitigation
	Permanent *bool `json:"permanent,omitempty" tf:"permanent,omitempty"`

	// Current state of your ip on mitigation
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type MitigationParameters struct {

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	// +kubebuilder:validation:Optional
	IPOnMitigation *string `json:"ipOnMitigation,omitempty" tf:"ip_on_mitigation,omitempty"`

	// Set on true if your ip is on permanent mitigation
	// +kubebuilder:validation:Optional
	Permanent *bool `json:"permanent,omitempty" tf:"permanent,omitempty"`
}

// MitigationSpec defines the desired state of Mitigation
type MitigationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MitigationParameters `json:"forProvider"`
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
	InitProvider MitigationInitParameters `json:"initProvider,omitempty"`
}

// MitigationStatus defines the observed state of Mitigation.
type MitigationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MitigationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Mitigation is the Schema for the Mitigations API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type Mitigation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ip) || (has(self.initProvider) && has(self.initProvider.ip))",message="spec.forProvider.ip is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipOnMitigation) || (has(self.initProvider) && has(self.initProvider.ipOnMitigation))",message="spec.forProvider.ipOnMitigation is a required parameter"
	Spec   MitigationSpec   `json:"spec"`
	Status MitigationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MitigationList contains a list of Mitigations
type MitigationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mitigation `json:"items"`
}

// Repository type metadata.
var (
	Mitigation_Kind             = "Mitigation"
	Mitigation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Mitigation_Kind}.String()
	Mitigation_KindAPIVersion   = Mitigation_Kind + "." + CRDGroupVersion.String()
	Mitigation_GroupVersionKind = CRDGroupVersion.WithKind(Mitigation_Kind)
)

func init() {
	SchemeBuilder.Register(&Mitigation{}, &MitigationList{})
}
