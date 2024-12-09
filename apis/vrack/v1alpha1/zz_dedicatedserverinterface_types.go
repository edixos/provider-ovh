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

type DedicatedServerInterfaceInitParameters struct {
	InterfaceID *string `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Service name of the vrack resource.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type DedicatedServerInterfaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InterfaceID *string `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Service name of the vrack resource.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type DedicatedServerInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	InterfaceID *string `json:"interfaceId,omitempty" tf:"interface_id,omitempty"`

	// Service name of the vrack resource.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// DedicatedServerInterfaceSpec defines the desired state of DedicatedServerInterface
type DedicatedServerInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedServerInterfaceParameters `json:"forProvider"`
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
	InitProvider DedicatedServerInterfaceInitParameters `json:"initProvider,omitempty"`
}

// DedicatedServerInterfaceStatus defines the observed state of DedicatedServerInterface.
type DedicatedServerInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedServerInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DedicatedServerInterface is the Schema for the DedicatedServerInterfaces API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type DedicatedServerInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interfaceId) || (has(self.initProvider) && has(self.initProvider.interfaceId))",message="spec.forProvider.interfaceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   DedicatedServerInterfaceSpec   `json:"spec"`
	Status DedicatedServerInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedServerInterfaceList contains a list of DedicatedServerInterfaces
type DedicatedServerInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedServerInterface `json:"items"`
}

// Repository type metadata.
var (
	DedicatedServerInterface_Kind             = "DedicatedServerInterface"
	DedicatedServerInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedServerInterface_Kind}.String()
	DedicatedServerInterface_KindAPIVersion   = DedicatedServerInterface_Kind + "." + CRDGroupVersion.String()
	DedicatedServerInterface_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedServerInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedServerInterface{}, &DedicatedServerInterfaceList{})
}
