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

type RefreshInitParameters struct {
	Keepers []*string `json:"keepers,omitempty" tf:"keepers,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type RefreshObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Keepers []*string `json:"keepers,omitempty" tf:"keepers,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type RefreshParameters struct {

	// +kubebuilder:validation:Optional
	Keepers []*string `json:"keepers,omitempty" tf:"keepers,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// RefreshSpec defines the desired state of Refresh
type RefreshSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RefreshParameters `json:"forProvider"`
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
	InitProvider RefreshInitParameters `json:"initProvider,omitempty"`
}

// RefreshStatus defines the observed state of Refresh.
type RefreshStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RefreshObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Refresh is the Schema for the Refreshs API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type Refresh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keepers) || (has(self.initProvider) && has(self.initProvider.keepers))",message="spec.forProvider.keepers is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   RefreshSpec   `json:"spec"`
	Status RefreshStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RefreshList contains a list of Refreshs
type RefreshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Refresh `json:"items"`
}

// Repository type metadata.
var (
	Refresh_Kind             = "Refresh"
	Refresh_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Refresh_Kind}.String()
	Refresh_KindAPIVersion   = Refresh_Kind + "." + CRDGroupVersion.String()
	Refresh_GroupVersionKind = CRDGroupVersion.WithKind(Refresh_Kind)
)

func init() {
	SchemeBuilder.Register(&Refresh{}, &RefreshList{})
}
