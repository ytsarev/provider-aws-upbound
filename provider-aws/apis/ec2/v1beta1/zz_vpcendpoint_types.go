/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DNSEntryObservation struct {
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`
}

type DNSEntryParameters struct {
}

type VPCEndpointObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	DNSEntry []DNSEntryObservation `json:"dnsEntry,omitempty" tf:"dns_entry,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	RequesterManaged *bool `json:"requesterManaged,omitempty" tf:"requester_managed,omitempty"`

	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCEndpointParameters struct {

	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("service_name",true)
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCEndpointSpec defines the desired state of VPCEndpoint
type VPCEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointParameters `json:"forProvider"`
}

// VPCEndpointStatus defines the observed state of VPCEndpoint.
type VPCEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpoint is the Schema for the VPCEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointSpec   `json:"spec"`
	Status            VPCEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointList contains a list of VPCEndpoints
type VPCEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpoint `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpoint_Kind             = "VPCEndpoint"
	VPCEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpoint_Kind}.String()
	VPCEndpoint_KindAPIVersion   = VPCEndpoint_Kind + "." + CRDGroupVersion.String()
	VPCEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpoint{}, &VPCEndpointList{})
}
