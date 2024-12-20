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

type TCPFrontendInitParameters struct {

	// +listType=set
	AllowedSource []*string `json:"allowedSource,omitempty" tf:"allowed_source,omitempty"`

	// +listType=set
	DedicatedIpfo []*string `json:"dedicatedIpfo,omitempty" tf:"dedicated_ipfo,omitempty"`

	DefaultFarmID *float64 `json:"defaultFarmId,omitempty" tf:"default_farm_id,omitempty"`

	DefaultSSLID *float64 `json:"defaultSslId,omitempty" tf:"default_ssl_id,omitempty"`

	// +listType=set
	DeniedSource []*string `json:"deniedSource,omitempty" tf:"denied_source,omitempty"`

	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TCPFrontendObservation struct {

	// +listType=set
	AllowedSource []*string `json:"allowedSource,omitempty" tf:"allowed_source,omitempty"`

	// +listType=set
	DedicatedIpfo []*string `json:"dedicatedIpfo,omitempty" tf:"dedicated_ipfo,omitempty"`

	DefaultFarmID *float64 `json:"defaultFarmId,omitempty" tf:"default_farm_id,omitempty"`

	DefaultSSLID *float64 `json:"defaultSslId,omitempty" tf:"default_ssl_id,omitempty"`

	// +listType=set
	DeniedSource []*string `json:"deniedSource,omitempty" tf:"denied_source,omitempty"`

	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TCPFrontendParameters struct {

	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedSource []*string `json:"allowedSource,omitempty" tf:"allowed_source,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	DedicatedIpfo []*string `json:"dedicatedIpfo,omitempty" tf:"dedicated_ipfo,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultFarmID *float64 `json:"defaultFarmId,omitempty" tf:"default_farm_id,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultSSLID *float64 `json:"defaultSslId,omitempty" tf:"default_ssl_id,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	DeniedSource []*string `json:"deniedSource,omitempty" tf:"denied_source,omitempty"`

	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// TCPFrontendSpec defines the desired state of TCPFrontend
type TCPFrontendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TCPFrontendParameters `json:"forProvider"`
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
	InitProvider TCPFrontendInitParameters `json:"initProvider,omitempty"`
}

// TCPFrontendStatus defines the observed state of TCPFrontend.
type TCPFrontendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TCPFrontendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TCPFrontend is the Schema for the TCPFrontends API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type TCPFrontend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   TCPFrontendSpec   `json:"spec"`
	Status TCPFrontendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TCPFrontendList contains a list of TCPFrontends
type TCPFrontendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPFrontend `json:"items"`
}

// Repository type metadata.
var (
	TCPFrontend_Kind             = "TCPFrontend"
	TCPFrontend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TCPFrontend_Kind}.String()
	TCPFrontend_KindAPIVersion   = TCPFrontend_Kind + "." + CRDGroupVersion.String()
	TCPFrontend_GroupVersionKind = CRDGroupVersion.WithKind(TCPFrontend_Kind)
)

func init() {
	SchemeBuilder.Register(&TCPFrontend{}, &TCPFrontendList{})
}
