//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaObservation) DeepCopyInto(out *ReplicaObservation) {
	*out = *in
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaObservation.
func (in *ReplicaObservation) DeepCopy() *ReplicaObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaParameters) DeepCopyInto(out *ReplicaParameters) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaParameters.
func (in *ReplicaParameters) DeepCopy() *ReplicaParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationRulesObservation) DeepCopyInto(out *RotationRulesObservation) {
	*out = *in
	if in.AutomaticallyAfterDays != nil {
		in, out := &in.AutomaticallyAfterDays, &out.AutomaticallyAfterDays
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationRulesObservation.
func (in *RotationRulesObservation) DeepCopy() *RotationRulesObservation {
	if in == nil {
		return nil
	}
	out := new(RotationRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationRulesParameters) DeepCopyInto(out *RotationRulesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationRulesParameters.
func (in *RotationRulesParameters) DeepCopy() *RotationRulesParameters {
	if in == nil {
		return nil
	}
	out := new(RotationRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Secret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretList) DeepCopyInto(out *SecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretList.
func (in *SecretList) DeepCopy() *SecretList {
	if in == nil {
		return nil
	}
	out := new(SecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObservation) DeepCopyInto(out *SecretObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]ReplicaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RotationEnabled != nil {
		in, out := &in.RotationEnabled, &out.RotationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RotationLambdaArn != nil {
		in, out := &in.RotationLambdaArn, &out.RotationLambdaArn
		*out = new(string)
		**out = **in
	}
	if in.RotationRules != nil {
		in, out := &in.RotationRules, &out.RotationRules
		*out = make([]RotationRulesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObservation.
func (in *SecretObservation) DeepCopy() *SecretObservation {
	if in == nil {
		return nil
	}
	out := new(SecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretParameters) DeepCopyInto(out *SecretParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceOverwriteReplicaSecret != nil {
		in, out := &in.ForceOverwriteReplicaSecret, &out.ForceOverwriteReplicaSecret
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyIDRef != nil {
		in, out := &in.KMSKeyIDRef, &out.KMSKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyIDSelector != nil {
		in, out := &in.KMSKeyIDSelector, &out.KMSKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryWindowInDays != nil {
		in, out := &in.RecoveryWindowInDays, &out.RecoveryWindowInDays
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = make([]ReplicaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretParameters.
func (in *SecretParameters) DeepCopy() *SecretParameters {
	if in == nil {
		return nil
	}
	out := new(SecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicy) DeepCopyInto(out *SecretPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicy.
func (in *SecretPolicy) DeepCopy() *SecretPolicy {
	if in == nil {
		return nil
	}
	out := new(SecretPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicyList) DeepCopyInto(out *SecretPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicyList.
func (in *SecretPolicyList) DeepCopy() *SecretPolicyList {
	if in == nil {
		return nil
	}
	out := new(SecretPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicyObservation) DeepCopyInto(out *SecretPolicyObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicyObservation.
func (in *SecretPolicyObservation) DeepCopy() *SecretPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SecretPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicyParameters) DeepCopyInto(out *SecretPolicyParameters) {
	*out = *in
	if in.BlockPublicPolicy != nil {
		in, out := &in.BlockPublicPolicy, &out.BlockPublicPolicy
		*out = new(bool)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecretArn != nil {
		in, out := &in.SecretArn, &out.SecretArn
		*out = new(string)
		**out = **in
	}
	if in.SecretArnRef != nil {
		in, out := &in.SecretArnRef, &out.SecretArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretArnSelector != nil {
		in, out := &in.SecretArnSelector, &out.SecretArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicyParameters.
func (in *SecretPolicyParameters) DeepCopy() *SecretPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SecretPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicySpec) DeepCopyInto(out *SecretPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicySpec.
func (in *SecretPolicySpec) DeepCopy() *SecretPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SecretPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPolicyStatus) DeepCopyInto(out *SecretPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPolicyStatus.
func (in *SecretPolicyStatus) DeepCopy() *SecretPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SecretPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotation) DeepCopyInto(out *SecretRotation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotation.
func (in *SecretRotation) DeepCopy() *SecretRotation {
	if in == nil {
		return nil
	}
	out := new(SecretRotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretRotation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationList) DeepCopyInto(out *SecretRotationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretRotation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationList.
func (in *SecretRotationList) DeepCopy() *SecretRotationList {
	if in == nil {
		return nil
	}
	out := new(SecretRotationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretRotationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationObservation) DeepCopyInto(out *SecretRotationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RotationEnabled != nil {
		in, out := &in.RotationEnabled, &out.RotationEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationObservation.
func (in *SecretRotationObservation) DeepCopy() *SecretRotationObservation {
	if in == nil {
		return nil
	}
	out := new(SecretRotationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationParameters) DeepCopyInto(out *SecretRotationParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RotationLambdaArn != nil {
		in, out := &in.RotationLambdaArn, &out.RotationLambdaArn
		*out = new(string)
		**out = **in
	}
	if in.RotationLambdaArnRef != nil {
		in, out := &in.RotationLambdaArnRef, &out.RotationLambdaArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RotationLambdaArnSelector != nil {
		in, out := &in.RotationLambdaArnSelector, &out.RotationLambdaArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RotationRules != nil {
		in, out := &in.RotationRules, &out.RotationRules
		*out = make([]SecretRotationRotationRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	if in.SecretIDRef != nil {
		in, out := &in.SecretIDRef, &out.SecretIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretIDSelector != nil {
		in, out := &in.SecretIDSelector, &out.SecretIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationParameters.
func (in *SecretRotationParameters) DeepCopy() *SecretRotationParameters {
	if in == nil {
		return nil
	}
	out := new(SecretRotationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationRotationRulesObservation) DeepCopyInto(out *SecretRotationRotationRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationRotationRulesObservation.
func (in *SecretRotationRotationRulesObservation) DeepCopy() *SecretRotationRotationRulesObservation {
	if in == nil {
		return nil
	}
	out := new(SecretRotationRotationRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationRotationRulesParameters) DeepCopyInto(out *SecretRotationRotationRulesParameters) {
	*out = *in
	if in.AutomaticallyAfterDays != nil {
		in, out := &in.AutomaticallyAfterDays, &out.AutomaticallyAfterDays
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationRotationRulesParameters.
func (in *SecretRotationRotationRulesParameters) DeepCopy() *SecretRotationRotationRulesParameters {
	if in == nil {
		return nil
	}
	out := new(SecretRotationRotationRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationSpec) DeepCopyInto(out *SecretRotationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationSpec.
func (in *SecretRotationSpec) DeepCopy() *SecretRotationSpec {
	if in == nil {
		return nil
	}
	out := new(SecretRotationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRotationStatus) DeepCopyInto(out *SecretRotationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRotationStatus.
func (in *SecretRotationStatus) DeepCopy() *SecretRotationStatus {
	if in == nil {
		return nil
	}
	out := new(SecretRotationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersion) DeepCopyInto(out *SecretVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersion.
func (in *SecretVersion) DeepCopy() *SecretVersion {
	if in == nil {
		return nil
	}
	out := new(SecretVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionList) DeepCopyInto(out *SecretVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionList.
func (in *SecretVersionList) DeepCopy() *SecretVersionList {
	if in == nil {
		return nil
	}
	out := new(SecretVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionObservation) DeepCopyInto(out *SecretVersionObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionObservation.
func (in *SecretVersionObservation) DeepCopy() *SecretVersionObservation {
	if in == nil {
		return nil
	}
	out := new(SecretVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionParameters) DeepCopyInto(out *SecretVersionParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecretBinarySecretRef != nil {
		in, out := &in.SecretBinarySecretRef, &out.SecretBinarySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
	if in.SecretIDRef != nil {
		in, out := &in.SecretIDRef, &out.SecretIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretIDSelector != nil {
		in, out := &in.SecretIDSelector, &out.SecretIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretStringSecretRef != nil {
		in, out := &in.SecretStringSecretRef, &out.SecretStringSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.VersionStages != nil {
		in, out := &in.VersionStages, &out.VersionStages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionParameters.
func (in *SecretVersionParameters) DeepCopy() *SecretVersionParameters {
	if in == nil {
		return nil
	}
	out := new(SecretVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionSpec) DeepCopyInto(out *SecretVersionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionSpec.
func (in *SecretVersionSpec) DeepCopy() *SecretVersionSpec {
	if in == nil {
		return nil
	}
	out := new(SecretVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionStatus) DeepCopyInto(out *SecretVersionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionStatus.
func (in *SecretVersionStatus) DeepCopy() *SecretVersionStatus {
	if in == nil {
		return nil
	}
	out := new(SecretVersionStatus)
	in.DeepCopyInto(out)
	return out
}
