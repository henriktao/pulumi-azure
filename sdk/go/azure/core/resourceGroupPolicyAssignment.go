// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Resource Group Policy Assignment.
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
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/policy"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
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
// 		_, err = core.NewResourceGroupPolicyAssignment(ctx, "exampleResourceGroupPolicyAssignment", &core.ResourceGroupPolicyAssignmentArgs{
// 			ResourceGroupId:    exampleResourceGroup.ID(),
// 			PolicyDefinitionId: exampleDefinition.ID(),
// 			Parameters:         pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "      \"tagName\": {\n", "        \"value\": \"Business Unit\"\n", "      },\n", "      \"tagValue\": {\n", "        \"value\": \"BU\"\n", "      }\n")),
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
// Resource Group Policy Assignments can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/resourceGroupPolicyAssignment:ResourceGroupPolicyAssignment example /subscriptions/00000000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
// ```
type ResourceGroupPolicyAssignment struct {
	pulumi.CustomResourceState

	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// An `identity` block as defined below.
	Identity ResourceGroupPolicyAssignmentIdentityPtrOutput `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Policy Assignment to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `nonComplianceMessage` blocks as defined below.
	NonComplianceMessages ResourceGroupPolicyAssignmentNonComplianceMessageArrayOutput `pulumi:"nonComplianceMessages"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayOutput `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
}

// NewResourceGroupPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewResourceGroupPolicyAssignment(ctx *pulumi.Context,
	name string, args *ResourceGroupPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*ResourceGroupPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitionId'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	var resource ResourceGroupPolicyAssignment
	err := ctx.RegisterResource("azure:core/resourceGroupPolicyAssignment:ResourceGroupPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroupPolicyAssignment gets an existing ResourceGroupPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroupPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupPolicyAssignmentState, opts ...pulumi.ResourceOption) (*ResourceGroupPolicyAssignment, error) {
	var resource ResourceGroupPolicyAssignment
	err := ctx.ReadResource("azure:core/resourceGroupPolicyAssignment:ResourceGroupPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroupPolicyAssignment resources.
type resourceGroupPolicyAssignmentState struct {
	// A description which should be used for this Policy Assignment.
	Description *string `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName *string `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce *bool `pulumi:"enforce"`
	// An `identity` block as defined below.
	Identity *ResourceGroupPolicyAssignmentIdentity `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Policy Assignment to be created.
	Name *string `pulumi:"name"`
	// One or more `nonComplianceMessage` blocks as defined below.
	NonComplianceMessages []ResourceGroupPolicyAssignmentNonComplianceMessage `pulumi:"nonComplianceMessages"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []string `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type ResourceGroupPolicyAssignmentState struct {
	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrInput
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrInput
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity ResourceGroupPolicyAssignmentIdentityPtrInput
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringPtrInput
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Policy Assignment. Changing this forces a new Policy Assignment to be created.
	Name pulumi.StringPtrInput
	// One or more `nonComplianceMessage` blocks as defined below.
	NonComplianceMessages ResourceGroupPolicyAssignmentNonComplianceMessageArrayInput
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayInput
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringPtrInput
	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupId pulumi.StringPtrInput
}

func (ResourceGroupPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupPolicyAssignmentState)(nil)).Elem()
}

type resourceGroupPolicyAssignmentArgs struct {
	// A description which should be used for this Policy Assignment.
	Description *string `pulumi:"description"`
	// The Display Name for this Policy Assignment.
	DisplayName *string `pulumi:"displayName"`
	// Specifies if this Policy should be enforced or not?
	Enforce *bool `pulumi:"enforce"`
	// An `identity` block as defined below.
	Identity *ResourceGroupPolicyAssignmentIdentity `pulumi:"identity"`
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `pulumi:"location"`
	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Policy Assignment. Changing this forces a new Policy Assignment to be created.
	Name *string `pulumi:"name"`
	// One or more `nonComplianceMessage` blocks as defined below.
	NonComplianceMessages []ResourceGroupPolicyAssignmentNonComplianceMessage `pulumi:"nonComplianceMessages"`
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []string `pulumi:"notScopes"`
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters *string `pulumi:"parameters"`
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId string `pulumi:"policyDefinitionId"`
	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupId string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a ResourceGroupPolicyAssignment resource.
type ResourceGroupPolicyAssignmentArgs struct {
	// A description which should be used for this Policy Assignment.
	Description pulumi.StringPtrInput
	// The Display Name for this Policy Assignment.
	DisplayName pulumi.StringPtrInput
	// Specifies if this Policy should be enforced or not?
	Enforce pulumi.BoolPtrInput
	// An `identity` block as defined below.
	Identity ResourceGroupPolicyAssignmentIdentityPtrInput
	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location pulumi.StringPtrInput
	// A JSON mapping of any Metadata for this Policy.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Policy Assignment. Changing this forces a new Policy Assignment to be created.
	Name pulumi.StringPtrInput
	// One or more `nonComplianceMessage` blocks as defined below.
	NonComplianceMessages ResourceGroupPolicyAssignmentNonComplianceMessageArrayInput
	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes pulumi.StringArrayInput
	// A JSON mapping of any Parameters for this Policy. Changing this forces a new Management Group Policy Assignment to be created.
	Parameters pulumi.StringPtrInput
	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionId pulumi.StringInput
	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	ResourceGroupId pulumi.StringInput
}

func (ResourceGroupPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupPolicyAssignmentArgs)(nil)).Elem()
}

type ResourceGroupPolicyAssignmentInput interface {
	pulumi.Input

	ToResourceGroupPolicyAssignmentOutput() ResourceGroupPolicyAssignmentOutput
	ToResourceGroupPolicyAssignmentOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentOutput
}

func (*ResourceGroupPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (i *ResourceGroupPolicyAssignment) ToResourceGroupPolicyAssignmentOutput() ResourceGroupPolicyAssignmentOutput {
	return i.ToResourceGroupPolicyAssignmentOutputWithContext(context.Background())
}

func (i *ResourceGroupPolicyAssignment) ToResourceGroupPolicyAssignmentOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPolicyAssignmentOutput)
}

// ResourceGroupPolicyAssignmentArrayInput is an input type that accepts ResourceGroupPolicyAssignmentArray and ResourceGroupPolicyAssignmentArrayOutput values.
// You can construct a concrete instance of `ResourceGroupPolicyAssignmentArrayInput` via:
//
//          ResourceGroupPolicyAssignmentArray{ ResourceGroupPolicyAssignmentArgs{...} }
type ResourceGroupPolicyAssignmentArrayInput interface {
	pulumi.Input

	ToResourceGroupPolicyAssignmentArrayOutput() ResourceGroupPolicyAssignmentArrayOutput
	ToResourceGroupPolicyAssignmentArrayOutputWithContext(context.Context) ResourceGroupPolicyAssignmentArrayOutput
}

type ResourceGroupPolicyAssignmentArray []ResourceGroupPolicyAssignmentInput

func (ResourceGroupPolicyAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (i ResourceGroupPolicyAssignmentArray) ToResourceGroupPolicyAssignmentArrayOutput() ResourceGroupPolicyAssignmentArrayOutput {
	return i.ToResourceGroupPolicyAssignmentArrayOutputWithContext(context.Background())
}

func (i ResourceGroupPolicyAssignmentArray) ToResourceGroupPolicyAssignmentArrayOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPolicyAssignmentArrayOutput)
}

// ResourceGroupPolicyAssignmentMapInput is an input type that accepts ResourceGroupPolicyAssignmentMap and ResourceGroupPolicyAssignmentMapOutput values.
// You can construct a concrete instance of `ResourceGroupPolicyAssignmentMapInput` via:
//
//          ResourceGroupPolicyAssignmentMap{ "key": ResourceGroupPolicyAssignmentArgs{...} }
type ResourceGroupPolicyAssignmentMapInput interface {
	pulumi.Input

	ToResourceGroupPolicyAssignmentMapOutput() ResourceGroupPolicyAssignmentMapOutput
	ToResourceGroupPolicyAssignmentMapOutputWithContext(context.Context) ResourceGroupPolicyAssignmentMapOutput
}

type ResourceGroupPolicyAssignmentMap map[string]ResourceGroupPolicyAssignmentInput

func (ResourceGroupPolicyAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (i ResourceGroupPolicyAssignmentMap) ToResourceGroupPolicyAssignmentMapOutput() ResourceGroupPolicyAssignmentMapOutput {
	return i.ToResourceGroupPolicyAssignmentMapOutputWithContext(context.Background())
}

func (i ResourceGroupPolicyAssignmentMap) ToResourceGroupPolicyAssignmentMapOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPolicyAssignmentMapOutput)
}

type ResourceGroupPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (ResourceGroupPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (o ResourceGroupPolicyAssignmentOutput) ToResourceGroupPolicyAssignmentOutput() ResourceGroupPolicyAssignmentOutput {
	return o
}

func (o ResourceGroupPolicyAssignmentOutput) ToResourceGroupPolicyAssignmentOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentOutput {
	return o
}

type ResourceGroupPolicyAssignmentArrayOutput struct{ *pulumi.OutputState }

func (ResourceGroupPolicyAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (o ResourceGroupPolicyAssignmentArrayOutput) ToResourceGroupPolicyAssignmentArrayOutput() ResourceGroupPolicyAssignmentArrayOutput {
	return o
}

func (o ResourceGroupPolicyAssignmentArrayOutput) ToResourceGroupPolicyAssignmentArrayOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentArrayOutput {
	return o
}

func (o ResourceGroupPolicyAssignmentArrayOutput) Index(i pulumi.IntInput) ResourceGroupPolicyAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceGroupPolicyAssignment {
		return vs[0].([]*ResourceGroupPolicyAssignment)[vs[1].(int)]
	}).(ResourceGroupPolicyAssignmentOutput)
}

type ResourceGroupPolicyAssignmentMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupPolicyAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceGroupPolicyAssignment)(nil)).Elem()
}

func (o ResourceGroupPolicyAssignmentMapOutput) ToResourceGroupPolicyAssignmentMapOutput() ResourceGroupPolicyAssignmentMapOutput {
	return o
}

func (o ResourceGroupPolicyAssignmentMapOutput) ToResourceGroupPolicyAssignmentMapOutputWithContext(ctx context.Context) ResourceGroupPolicyAssignmentMapOutput {
	return o
}

func (o ResourceGroupPolicyAssignmentMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupPolicyAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceGroupPolicyAssignment {
		return vs[0].(map[string]*ResourceGroupPolicyAssignment)[vs[1].(string)]
	}).(ResourceGroupPolicyAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPolicyAssignmentInput)(nil)).Elem(), &ResourceGroupPolicyAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPolicyAssignmentArrayInput)(nil)).Elem(), ResourceGroupPolicyAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPolicyAssignmentMapInput)(nil)).Elem(), ResourceGroupPolicyAssignmentMap{})
	pulumi.RegisterOutputType(ResourceGroupPolicyAssignmentOutput{})
	pulumi.RegisterOutputType(ResourceGroupPolicyAssignmentArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupPolicyAssignmentMapOutput{})
}
