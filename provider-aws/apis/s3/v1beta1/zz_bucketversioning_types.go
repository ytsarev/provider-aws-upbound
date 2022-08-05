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

type BucketVersioningObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketVersioningParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// +kubebuilder:validation:Optional
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	VersioningConfiguration []VersioningConfigurationParameters `json:"versioningConfiguration" tf:"versioning_configuration,omitempty"`
}

type VersioningConfigurationObservation struct {
}

type VersioningConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	MfaDelete *string `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`

	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

// BucketVersioningSpec defines the desired state of BucketVersioning
type BucketVersioningSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketVersioningParameters `json:"forProvider"`
}

// BucketVersioningStatus defines the observed state of BucketVersioning.
type BucketVersioningStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketVersioningObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketVersioning is the Schema for the BucketVersionings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketVersioning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketVersioningSpec   `json:"spec"`
	Status            BucketVersioningStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketVersioningList contains a list of BucketVersionings
type BucketVersioningList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketVersioning `json:"items"`
}

// Repository type metadata.
var (
	BucketVersioning_Kind             = "BucketVersioning"
	BucketVersioning_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketVersioning_Kind}.String()
	BucketVersioning_KindAPIVersion   = BucketVersioning_Kind + "." + CRDGroupVersion.String()
	BucketVersioning_GroupVersionKind = CRDGroupVersion.WithKind(BucketVersioning_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketVersioning{}, &BucketVersioningList{})
}
