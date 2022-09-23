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

type GroupMembershipObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupMembershipParameters struct {

	// –  The IAM Group name to attach the list of users to
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Reference to a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// Selector for a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// The name to identify the Group Membership
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// References to User to populate users.
	// +kubebuilder:validation:Optional
	UserRefs []v1.Reference `json:"userRefs,omitempty" tf:"-"`

	// Selector for a list of User to populate users.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`

	// A list of IAM User names to associate with the Group
	// +crossplane:generate:reference:type=User
	// +crossplane:generate:reference:refFieldName=UserRefs
	// +crossplane:generate:reference:selectorFieldName=UserSelector
	// +kubebuilder:validation:Optional
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

// GroupMembershipSpec defines the desired state of GroupMembership
type GroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembershipParameters `json:"forProvider"`
}

// GroupMembershipStatus defines the observed state of GroupMembership.
type GroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembership is the Schema for the GroupMemberships API. Provides a top level resource to manage IAM Group membership for IAM Users.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMembershipSpec   `json:"spec"`
	Status            GroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembershipList contains a list of GroupMemberships
type GroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembership `json:"items"`
}

// Repository type metadata.
var (
	GroupMembership_Kind             = "GroupMembership"
	GroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMembership_Kind}.String()
	GroupMembership_KindAPIVersion   = GroupMembership_Kind + "." + CRDGroupVersion.String()
	GroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(GroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMembership{}, &GroupMembershipList{})
}
