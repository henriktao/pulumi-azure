// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Subscription within a API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_subscription.html.markdown.
type Subscription struct {
	s *pulumi.ResourceState
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOpt) (*Subscription, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["displayName"] = nil
		inputs["primaryKey"] = nil
		inputs["productId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["secondaryKey"] = nil
		inputs["state"] = nil
		inputs["subscriptionId"] = nil
		inputs["userId"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["displayName"] = args.DisplayName
		inputs["primaryKey"] = args.PrimaryKey
		inputs["productId"] = args.ProductId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["secondaryKey"] = args.SecondaryKey
		inputs["state"] = args.State
		inputs["subscriptionId"] = args.SubscriptionId
		inputs["userId"] = args.UserId
	}
	s, err := ctx.RegisterResource("azure:apimanagement/subscription:Subscription", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subscription{s: s}, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubscriptionState, opts ...pulumi.ResourceOpt) (*Subscription, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["displayName"] = state.DisplayName
		inputs["primaryKey"] = state.PrimaryKey
		inputs["productId"] = state.ProductId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["secondaryKey"] = state.SecondaryKey
		inputs["state"] = state.State
		inputs["subscriptionId"] = state.SubscriptionId
		inputs["userId"] = state.UserId
	}
	s, err := ctx.ReadResource("azure:apimanagement/subscription:Subscription", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subscription{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Subscription) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Subscription) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
func (r *Subscription) ApiManagementName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The display name of this Subscription.
func (r *Subscription) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

func (r *Subscription) PrimaryKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["primaryKey"])
}

// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
func (r *Subscription) ProductId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["productId"])
}

// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
func (r *Subscription) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

func (r *Subscription) SecondaryKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secondaryKey"])
}

// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
func (r *Subscription) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
func (r *Subscription) SubscriptionId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subscriptionId"])
}

// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
func (r *Subscription) UserId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userId"])
}

// Input properties used for looking up and filtering Subscription resources.
type SubscriptionState struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The display name of this Subscription.
	DisplayName interface{}
	PrimaryKey interface{}
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	SecondaryKey interface{}
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State interface{}
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId interface{}
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId interface{}
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The name of the API Management Service where this Subscription should be created. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The display name of this Subscription.
	DisplayName interface{}
	PrimaryKey interface{}
	// The ID of the Product which should be assigned to this Subscription. Changing this forces a new resource to be created.
	ProductId interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	SecondaryKey interface{}
	// The state of this Subscription. Possible values are `active`, `cancelled`, `expired`, `rejected`, `submitted` and `suspended`. Defaults to `submitted`.
	State interface{}
	// An Identifier which should used as the ID of this Subscription. If not specified a new Subscription ID will be generated. Changing this forces a new resource to be created.
	SubscriptionId interface{}
	// The ID of the User which should be assigned to this Subscription. Changing this forces a new resource to be created.
	UserId interface{}
}
