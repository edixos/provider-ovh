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

type TCPRouteRuleInitParameters struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	RouteID *string `json:"routeId,omitempty" tf:"route_id,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	SubField *string `json:"subField,omitempty" tf:"sub_field,omitempty"`
}

type TCPRouteRuleObservation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	RouteID *string `json:"routeId,omitempty" tf:"route_id,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	SubField *string `json:"subField,omitempty" tf:"sub_field,omitempty"`
}

type TCPRouteRuleParameters struct {

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// +kubebuilder:validation:Optional
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// +kubebuilder:validation:Optional
	RouteID *string `json:"routeId,omitempty" tf:"route_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	SubField *string `json:"subField,omitempty" tf:"sub_field,omitempty"`
}

// TCPRouteRuleSpec defines the desired state of TCPRouteRule
type TCPRouteRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TCPRouteRuleParameters `json:"forProvider"`
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
	InitProvider TCPRouteRuleInitParameters `json:"initProvider,omitempty"`
}

// TCPRouteRuleStatus defines the observed state of TCPRouteRule.
type TCPRouteRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TCPRouteRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TCPRouteRule is the Schema for the TCPRouteRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,lb}
type TCPRouteRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.field) || (has(self.initProvider) && has(self.initProvider.field))",message="spec.forProvider.field is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.match) || (has(self.initProvider) && has(self.initProvider.match))",message="spec.forProvider.match is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeId) || (has(self.initProvider) && has(self.initProvider.routeId))",message="spec.forProvider.routeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   TCPRouteRuleSpec   `json:"spec"`
	Status TCPRouteRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TCPRouteRuleList contains a list of TCPRouteRules
type TCPRouteRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPRouteRule `json:"items"`
}

// Repository type metadata.
var (
	TCPRouteRule_Kind             = "TCPRouteRule"
	TCPRouteRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TCPRouteRule_Kind}.String()
	TCPRouteRule_KindAPIVersion   = TCPRouteRule_Kind + "." + CRDGroupVersion.String()
	TCPRouteRule_GroupVersionKind = CRDGroupVersion.WithKind(TCPRouteRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TCPRouteRule{}, &TCPRouteRuleList{})
}
