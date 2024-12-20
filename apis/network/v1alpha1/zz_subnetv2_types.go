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

type AllocationPoolInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type AllocationPoolParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type HostRouteInitParameters struct {
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	Nexthop *string `json:"nexthop,omitempty" tf:"nexthop,omitempty"`
}

type HostRouteObservation struct {
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	Nexthop *string `json:"nexthop,omitempty" tf:"nexthop,omitempty"`
}

type HostRouteParameters struct {

	// +kubebuilder:validation:Optional
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	Nexthop *string `json:"nexthop" tf:"nexthop,omitempty"`
}

type SubnetV2InitParameters struct {

	// DHCP allocation pools of subnet
	AllocationPool []AllocationPoolInitParameters `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// CIDR of subnet
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Enable DHCP in subnet
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// List of DNS nameservers, default: 213.186.33.99
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Enable gateway IP in subnet
	EnableGatewayIP *bool `json:"enableGatewayIp,omitempty" tf:"enable_gateway_ip,omitempty"`

	// Gateway IP of subnet
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Static host routes of subnet
	HostRoute []HostRouteInitParameters `json:"hostRoute,omitempty" tf:"host_route,omitempty"`

	// Name of subnet
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network ID of subnet
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Region of network/subnet
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type SubnetV2Observation struct {

	// DHCP allocation pools of subnet
	AllocationPool []AllocationPoolObservation `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// CIDR of subnet
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Enable DHCP in subnet
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// List of DNS nameservers, default: 213.186.33.99
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Enable gateway IP in subnet
	EnableGatewayIP *bool `json:"enableGatewayIp,omitempty" tf:"enable_gateway_ip,omitempty"`

	// Gateway IP of subnet
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Static host routes of subnet
	HostRoute []HostRouteObservation `json:"hostRoute,omitempty" tf:"host_route,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of subnet
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network ID of subnet
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Region of network/subnet
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type SubnetV2Parameters struct {

	// DHCP allocation pools of subnet
	// +kubebuilder:validation:Optional
	AllocationPool []AllocationPoolParameters `json:"allocationPool,omitempty" tf:"allocation_pool,omitempty"`

	// CIDR of subnet
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Enable DHCP in subnet
	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// List of DNS nameservers, default: 213.186.33.99
	// +kubebuilder:validation:Optional
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// Enable gateway IP in subnet
	// +kubebuilder:validation:Optional
	EnableGatewayIP *bool `json:"enableGatewayIp,omitempty" tf:"enable_gateway_ip,omitempty"`

	// Gateway IP of subnet
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Static host routes of subnet
	// +kubebuilder:validation:Optional
	HostRoute []HostRouteParameters `json:"hostRoute,omitempty" tf:"host_route,omitempty"`

	// Name of subnet
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network ID of subnet
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Region of network/subnet
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Service name of the resource representing the id of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// SubnetV2Spec defines the desired state of SubnetV2
type SubnetV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetV2Parameters `json:"forProvider"`
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
	InitProvider SubnetV2InitParameters `json:"initProvider,omitempty"`
}

// SubnetV2Status defines the observed state of SubnetV2.
type SubnetV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubnetV2 is the Schema for the SubnetV2s API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type SubnetV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   SubnetV2Spec   `json:"spec"`
	Status SubnetV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetV2List contains a list of SubnetV2s
type SubnetV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetV2 `json:"items"`
}

// Repository type metadata.
var (
	SubnetV2_Kind             = "SubnetV2"
	SubnetV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetV2_Kind}.String()
	SubnetV2_KindAPIVersion   = SubnetV2_Kind + "." + CRDGroupVersion.String()
	SubnetV2_GroupVersionKind = CRDGroupVersion.WithKind(SubnetV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetV2{}, &SubnetV2List{})
}
