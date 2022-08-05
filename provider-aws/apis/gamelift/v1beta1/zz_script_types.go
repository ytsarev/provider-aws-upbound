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

type ScriptObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ScriptParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageLocation []ScriptStorageLocationParameters `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`
}

type ScriptStorageLocationObservation struct {
}

type ScriptStorageLocationParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Object
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("key",false)
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// ScriptSpec defines the desired state of Script
type ScriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScriptParameters `json:"forProvider"`
}

// ScriptStatus defines the observed state of Script.
type ScriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Script is the Schema for the Scripts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Script struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScriptSpec   `json:"spec"`
	Status            ScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScriptList contains a list of Scripts
type ScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Script `json:"items"`
}

// Repository type metadata.
var (
	Script_Kind             = "Script"
	Script_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Script_Kind}.String()
	Script_KindAPIVersion   = Script_Kind + "." + CRDGroupVersion.String()
	Script_GroupVersionKind = CRDGroupVersion.WithKind(Script_Kind)
)

func init() {
	SchemeBuilder.Register(&Script{}, &ScriptList{})
}
