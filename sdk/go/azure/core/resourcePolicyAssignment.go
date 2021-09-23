// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Policy Assignment to a Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVirtualNetwork, err := network.LookupVirtualNetwork(ctx, &network.LookupVirtualNetworkArgs{
// 			Name:              "production",
// 			ResourceGroupName: "networking",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleDefinition, err := policy.NewDefinition(ctx, "exampleDefinition", &policy.DefinitionArgs{
// 			PolicyType: pulumi.String("Custom"),
// 			Mode:       pulumi.String("All"),
// 			PolicyRule: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "	{\n", "    \"if\": {\n", "      \"not\": {\n", "        \"field\": \"location\",\n", "        \"equals\": \"westeurope\"\n", "      }\n", "    },\n", "    \"then\": {\n", "      \"effect\": \"Deny\"\n", "    }\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewResourcePolicyAssignment(ctx, "exampleResourcePolicyAssignment", &core.ResourcePolicyAssignmentArgs{
// 			ResourceId:         pulumi.String(exampleVirtualNetwork.Id),
// 			PolicyDefinitionId: exampleDefinition.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Resource Policy Assignments can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/resourcePolicyAssignment:ResourcePolicyAssignment example "{resource}/providers/Microsoft.Authorization/policyAssignments/assignment1"
// ```
//
//  where `{resource}` is a Resource ID in the form `/subscriptions/00000000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/network1`.
type ResourcePolicyAssignment struct {
	pulumi.CustomResourceState

	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// A `identity` block as defined below.
	Identity ResourcePolicyAssignmentIdentityPtrOutput `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Resource Policy Assignment to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayOutput `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// The ID of the Resource (or Resource Scope) where this should be applied. Changing this forces a new Resource Policy Assignment to be created.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewResourcePolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicyAssignment(ctx *pulumi.Context,
	name string, args *ResourcePolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*ResourcePolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitionId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource ResourcePolicyAssignment
	err := ctx.RegisterResource("azure:core/resourcePolicyAssignment:ResourcePolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicyAssignment gets an existing ResourcePolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyAssignmentState, opts ...pulumi.ResourceOption) (*ResourcePolicyAssignment, error) {
	var resource ResourcePolicyAssignment
	err := ctx.ReadResource("azure:core/resourcePolicyAssignment:ResourcePolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicyAssignment resources.
type resourcePolicyAssignmentState struct {
	// A description which should be used for this Policy Assignment.
	Description *string `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName *string `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce *bool `pulumi:"enforce"`
	// A `identity` block as defined below.
	Identity *ResourcePolicyAssignmentIdentity `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Resource Policy Assignment to be created.
	Name *string `pulumi:"name"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []string `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The ID of the Resource (or Resource Scope) where this should be applied. Changing this forces a new Resource Policy Assignment to be created.
	ResourceId *string `pulumi:"resourceId"`
}

type ResourcePolicyAssignmentState struct {
	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrInput
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrInput
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity ResourcePolicyAssignmentIdentityPtrInput
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringPtrInput
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Policy Assignment. Changing this forces a new Resource Policy Assignment to be created.
	Name pulumi.StringPtrInput
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayInput
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringPtrInput
	// The ID of the Resource (or Resource Scope) where this should be applied. Changing this forces a new Resource Policy Assignment to be created.
	ResourceId pulumi.StringPtrInput
}

func (ResourcePolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyAssignmentState)(nil)).Elem()
}

type resourcePolicyAssignmentArgs struct {
	// A description which should be used for this Policy Assignment.
	Description *string `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName *string `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce *bool `pulumi:"enforce"`
	// A `identity` block as defined below.
	Identity *ResourcePolicyAssignmentIdentity `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Resource Policy Assignment to be created.
	Name *string `pulumi:"name"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []string `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// The ID of the Resource (or Resource Scope) where this should be applied. Changing this forces a new Resource Policy Assignment to be created.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResourcePolicyAssignment resource.
type ResourcePolicyAssignmentArgs struct {
	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrInput
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrInput
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity ResourcePolicyAssignmentIdentityPtrInput
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringPtrInput
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Policy Assignment. Changing this forces a new Resource Policy Assignment to be created.
	Name pulumi.StringPtrInput
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayInput
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringInput
	// The ID of the Resource (or Resource Scope) where this should be applied. Changing this forces a new Resource Policy Assignment to be created.
	ResourceId pulumi.StringInput
}

func (ResourcePolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyAssignmentArgs)(nil)).Elem()
}

type ResourcePolicyAssignmentInput interface {
	pulumi.Input

	ToResourcePolicyAssignmentOutput() ResourcePolicyAssignmentOutput
	ToResourcePolicyAssignmentOutputWithContext(ctx context.Context) ResourcePolicyAssignmentOutput
}

func (*ResourcePolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicyAssignment)(nil))
}

func (i *ResourcePolicyAssignment) ToResourcePolicyAssignmentOutput() ResourcePolicyAssignmentOutput {
	return i.ToResourcePolicyAssignmentOutputWithContext(context.Background())
}

func (i *ResourcePolicyAssignment) ToResourcePolicyAssignmentOutputWithContext(ctx context.Context) ResourcePolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyAssignmentOutput)
}

func (i *ResourcePolicyAssignment) ToResourcePolicyAssignmentPtrOutput() ResourcePolicyAssignmentPtrOutput {
	return i.ToResourcePolicyAssignmentPtrOutputWithContext(context.Background())
}

func (i *ResourcePolicyAssignment) ToResourcePolicyAssignmentPtrOutputWithContext(ctx context.Context) ResourcePolicyAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyAssignmentPtrOutput)
}

type ResourcePolicyAssignmentPtrInput interface {
	pulumi.Input

	ToResourcePolicyAssignmentPtrOutput() ResourcePolicyAssignmentPtrOutput
	ToResourcePolicyAssignmentPtrOutputWithContext(ctx context.Context) ResourcePolicyAssignmentPtrOutput
}

type resourcePolicyAssignmentPtrType ResourcePolicyAssignmentArgs

func (*resourcePolicyAssignmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicyAssignment)(nil))
}

func (i *resourcePolicyAssignmentPtrType) ToResourcePolicyAssignmentPtrOutput() ResourcePolicyAssignmentPtrOutput {
	return i.ToResourcePolicyAssignmentPtrOutputWithContext(context.Background())
}

func (i *resourcePolicyAssignmentPtrType) ToResourcePolicyAssignmentPtrOutputWithContext(ctx context.Context) ResourcePolicyAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyAssignmentPtrOutput)
}

// ResourcePolicyAssignmentArrayInput is an input type that accepts ResourcePolicyAssignmentArray and ResourcePolicyAssignmentArrayOutput values.
// You can construct a concrete instance of `ResourcePolicyAssignmentArrayInput` via:
//
//          ResourcePolicyAssignmentArray{ ResourcePolicyAssignmentArgs{...} }
type ResourcePolicyAssignmentArrayInput interface {
	pulumi.Input

	ToResourcePolicyAssignmentArrayOutput() ResourcePolicyAssignmentArrayOutput
	ToResourcePolicyAssignmentArrayOutputWithContext(context.Context) ResourcePolicyAssignmentArrayOutput
}

type ResourcePolicyAssignmentArray []ResourcePolicyAssignmentInput

func (ResourcePolicyAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePolicyAssignment)(nil)).Elem()
}

func (i ResourcePolicyAssignmentArray) ToResourcePolicyAssignmentArrayOutput() ResourcePolicyAssignmentArrayOutput {
	return i.ToResourcePolicyAssignmentArrayOutputWithContext(context.Background())
}

func (i ResourcePolicyAssignmentArray) ToResourcePolicyAssignmentArrayOutputWithContext(ctx context.Context) ResourcePolicyAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyAssignmentArrayOutput)
}

// ResourcePolicyAssignmentMapInput is an input type that accepts ResourcePolicyAssignmentMap and ResourcePolicyAssignmentMapOutput values.
// You can construct a concrete instance of `ResourcePolicyAssignmentMapInput` via:
//
//          ResourcePolicyAssignmentMap{ "key": ResourcePolicyAssignmentArgs{...} }
type ResourcePolicyAssignmentMapInput interface {
	pulumi.Input

	ToResourcePolicyAssignmentMapOutput() ResourcePolicyAssignmentMapOutput
	ToResourcePolicyAssignmentMapOutputWithContext(context.Context) ResourcePolicyAssignmentMapOutput
}

type ResourcePolicyAssignmentMap map[string]ResourcePolicyAssignmentInput

func (ResourcePolicyAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePolicyAssignment)(nil)).Elem()
}

func (i ResourcePolicyAssignmentMap) ToResourcePolicyAssignmentMapOutput() ResourcePolicyAssignmentMapOutput {
	return i.ToResourcePolicyAssignmentMapOutputWithContext(context.Background())
}

func (i ResourcePolicyAssignmentMap) ToResourcePolicyAssignmentMapOutputWithContext(ctx context.Context) ResourcePolicyAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyAssignmentMapOutput)
}

type ResourcePolicyAssignmentOutput struct{ *pulumi.OutputState }

func (ResourcePolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicyAssignment)(nil))
}

func (o ResourcePolicyAssignmentOutput) ToResourcePolicyAssignmentOutput() ResourcePolicyAssignmentOutput {
	return o
}

func (o ResourcePolicyAssignmentOutput) ToResourcePolicyAssignmentOutputWithContext(ctx context.Context) ResourcePolicyAssignmentOutput {
	return o
}

func (o ResourcePolicyAssignmentOutput) ToResourcePolicyAssignmentPtrOutput() ResourcePolicyAssignmentPtrOutput {
	return o.ToResourcePolicyAssignmentPtrOutputWithContext(context.Background())
}

func (o ResourcePolicyAssignmentOutput) ToResourcePolicyAssignmentPtrOutputWithContext(ctx context.Context) ResourcePolicyAssignmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourcePolicyAssignment) *ResourcePolicyAssignment {
		return &v
	}).(ResourcePolicyAssignmentPtrOutput)
}

type ResourcePolicyAssignmentPtrOutput struct{ *pulumi.OutputState }

func (ResourcePolicyAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicyAssignment)(nil))
}

func (o ResourcePolicyAssignmentPtrOutput) ToResourcePolicyAssignmentPtrOutput() ResourcePolicyAssignmentPtrOutput {
	return o
}

func (o ResourcePolicyAssignmentPtrOutput) ToResourcePolicyAssignmentPtrOutputWithContext(ctx context.Context) ResourcePolicyAssignmentPtrOutput {
	return o
}

func (o ResourcePolicyAssignmentPtrOutput) Elem() ResourcePolicyAssignmentOutput {
	return o.ApplyT(func(v *ResourcePolicyAssignment) ResourcePolicyAssignment {
		if v != nil {
			return *v
		}
		var ret ResourcePolicyAssignment
		return ret
	}).(ResourcePolicyAssignmentOutput)
}

type ResourcePolicyAssignmentArrayOutput struct{ *pulumi.OutputState }

func (ResourcePolicyAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcePolicyAssignment)(nil))
}

func (o ResourcePolicyAssignmentArrayOutput) ToResourcePolicyAssignmentArrayOutput() ResourcePolicyAssignmentArrayOutput {
	return o
}

func (o ResourcePolicyAssignmentArrayOutput) ToResourcePolicyAssignmentArrayOutputWithContext(ctx context.Context) ResourcePolicyAssignmentArrayOutput {
	return o
}

func (o ResourcePolicyAssignmentArrayOutput) Index(i pulumi.IntInput) ResourcePolicyAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourcePolicyAssignment {
		return vs[0].([]ResourcePolicyAssignment)[vs[1].(int)]
	}).(ResourcePolicyAssignmentOutput)
}

type ResourcePolicyAssignmentMapOutput struct{ *pulumi.OutputState }

func (ResourcePolicyAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourcePolicyAssignment)(nil))
}

func (o ResourcePolicyAssignmentMapOutput) ToResourcePolicyAssignmentMapOutput() ResourcePolicyAssignmentMapOutput {
	return o
}

func (o ResourcePolicyAssignmentMapOutput) ToResourcePolicyAssignmentMapOutputWithContext(ctx context.Context) ResourcePolicyAssignmentMapOutput {
	return o
}

func (o ResourcePolicyAssignmentMapOutput) MapIndex(k pulumi.StringInput) ResourcePolicyAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourcePolicyAssignment {
		return vs[0].(map[string]ResourcePolicyAssignment)[vs[1].(string)]
	}).(ResourcePolicyAssignmentOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourcePolicyAssignmentOutput{})
	pulumi.RegisterOutputType(ResourcePolicyAssignmentPtrOutput{})
	pulumi.RegisterOutputType(ResourcePolicyAssignmentArrayOutput{})
	pulumi.RegisterOutputType(ResourcePolicyAssignmentMapOutput{})
}
