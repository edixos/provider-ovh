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

type FirewallRuleInitParameters struct {

	// Possible values for action
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Fragments option
	Fragments *bool `json:"fragments,omitempty" tf:"fragments,omitempty"`

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	IPOnFirewall *string `json:"ipOnFirewall,omitempty" tf:"ip_on_firewall,omitempty"`

	// Possible values for protocol
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Possible values for action
	Sequence *float64 `json:"sequence,omitempty" tf:"sequence,omitempty"`

	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// TCP option on your rule
	TCPOption *string `json:"tcpOption,omitempty" tf:"tcp_option,omitempty"`
}

type FirewallRuleObservation struct {

	// Possible values for action
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// Destination ip for your rule
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPortDesc *string `json:"destinationPortDesc,omitempty" tf:"destination_port_desc,omitempty"`

	// Fragments option
	Fragments *bool `json:"fragments,omitempty" tf:"fragments,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	IPOnFirewall *string `json:"ipOnFirewall,omitempty" tf:"ip_on_firewall,omitempty"`

	// Possible values for protocol
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`

	// Possible values for action
	Sequence *float64 `json:"sequence,omitempty" tf:"sequence,omitempty"`

	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// Source port for your rule. Only with TCP/UDP protocol
	SourcePortDesc *string `json:"sourcePortDesc,omitempty" tf:"source_port_desc,omitempty"`

	// Current state of your rule
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// TCP option on your rule
	TCPOption *string `json:"tcpOption,omitempty" tf:"tcp_option,omitempty"`
}

type FirewallRuleParameters struct {

	// Possible values for action
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Destination port for your rule. Only with TCP/UDP protocol
	// +kubebuilder:validation:Optional
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Fragments option
	// +kubebuilder:validation:Optional
	Fragments *bool `json:"fragments,omitempty" tf:"fragments,omitempty"`

	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IPv4 address (e.g., 192.0.2.0)
	// +kubebuilder:validation:Optional
	IPOnFirewall *string `json:"ipOnFirewall,omitempty" tf:"ip_on_firewall,omitempty"`

	// Possible values for protocol
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Possible values for action
	// +kubebuilder:validation:Optional
	Sequence *float64 `json:"sequence,omitempty" tf:"sequence,omitempty"`

	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Source port for your rule. Only with TCP/UDP protocol
	// +kubebuilder:validation:Optional
	SourcePort *float64 `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// TCP option on your rule
	// +kubebuilder:validation:Optional
	TCPOption *string `json:"tcpOption,omitempty" tf:"tcp_option,omitempty"`
}

// FirewallRuleSpec defines the desired state of FirewallRule
type FirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleParameters `json:"forProvider"`
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
	InitProvider FirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FirewallRule is the Schema for the FirewallRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ip) || (has(self.initProvider) && has(self.initProvider.ip))",message="spec.forProvider.ip is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipOnFirewall) || (has(self.initProvider) && has(self.initProvider.ipOnFirewall))",message="spec.forProvider.ipOnFirewall is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sequence) || (has(self.initProvider) && has(self.initProvider.sequence))",message="spec.forProvider.sequence is a required parameter"
	Spec   FirewallRuleSpec   `json:"spec"`
	Status FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallRule_Kind             = "FirewallRule"
	FirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRule_Kind}.String()
	FirewallRule_KindAPIVersion   = FirewallRule_Kind + "." + CRDGroupVersion.String()
	FirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}
