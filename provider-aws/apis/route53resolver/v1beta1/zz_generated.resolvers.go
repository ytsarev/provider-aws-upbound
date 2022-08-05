/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DNSSECConfig.
func (mg *DNSSECConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Endpoint.
func (mg *Endpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.IPAddress); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPAddress[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IPAddress[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.IPAddress[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IPAddress[i3].SubnetID")
		}
		mg.Spec.ForProvider.IPAddress[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IPAddress[i3].SubnetIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this FirewallConfig.
func (mg *FirewallConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FirewallRule.
func (mg *FirewallRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirewallDomainListID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.FirewallDomainListIDRef,
		Selector:     mg.Spec.ForProvider.FirewallDomainListIDSelector,
		To: reference.To{
			List:    &FirewallDomainListList{},
			Managed: &FirewallDomainList{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirewallDomainListID")
	}
	mg.Spec.ForProvider.FirewallDomainListID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirewallDomainListIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirewallRuleGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.FirewallRuleGroupIDRef,
		Selector:     mg.Spec.ForProvider.FirewallRuleGroupIDSelector,
		To: reference.To{
			List:    &FirewallRuleGroupList{},
			Managed: &FirewallRuleGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirewallRuleGroupID")
	}
	mg.Spec.ForProvider.FirewallRuleGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirewallRuleGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FirewallRuleGroupAssociation.
func (mg *FirewallRuleGroupAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirewallRuleGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.FirewallRuleGroupIDRef,
		Selector:     mg.Spec.ForProvider.FirewallRuleGroupIDSelector,
		To: reference.To{
			List:    &FirewallRuleGroupList{},
			Managed: &FirewallRuleGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirewallRuleGroupID")
	}
	mg.Spec.ForProvider.FirewallRuleGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirewallRuleGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QueryLogConfigAssociation.
func (mg *QueryLogConfigAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResolverQueryLogConfigID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResolverQueryLogConfigIDRef,
		Selector:     mg.Spec.ForProvider.ResolverQueryLogConfigIDSelector,
		To: reference.To{
			List:    &QueryLogConfigList{},
			Managed: &QueryLogConfig{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResolverQueryLogConfigID")
	}
	mg.Spec.ForProvider.ResolverQueryLogConfigID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResolverQueryLogConfigIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResolverEndpointID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResolverEndpointIDRef,
		Selector:     mg.Spec.ForProvider.ResolverEndpointIDSelector,
		To: reference.To{
			List:    &EndpointList{},
			Managed: &Endpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResolverEndpointID")
	}
	mg.Spec.ForProvider.ResolverEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResolverEndpointIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleAssociation.
func (mg *RuleAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResolverRuleID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResolverRuleIDRef,
		Selector:     mg.Spec.ForProvider.ResolverRuleIDSelector,
		To: reference.To{
			List:    &RuleList{},
			Managed: &Rule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResolverRuleID")
	}
	mg.Spec.ForProvider.ResolverRuleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResolverRuleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}
