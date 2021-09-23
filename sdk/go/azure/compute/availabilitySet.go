// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Availability Set for Virtual Machines.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		_, err = compute.NewAvailabilitySet(ctx, "exampleAvailabilitySet", &compute.AvailabilitySetArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
// 			},
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
// Availability Sets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/availabilitySet:AvailabilitySet group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
// ```
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrOutput `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount pulumi.IntPtrOutput `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure:compute/availabilitySet:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure:compute/availabilitySet:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type availabilitySetState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed *bool `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AvailabilitySetState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed *bool `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used. Defaults to `3`. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies the number of update domains that are used. Defaults to `5`. Changing this forces a new resource to be created.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created.
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}

type AvailabilitySetInput interface {
	pulumi.Input

	ToAvailabilitySetOutput() AvailabilitySetOutput
	ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput
}

func (*AvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySet)(nil))
}

func (i *AvailabilitySet) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return i.ToAvailabilitySetOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetOutput)
}

func (i *AvailabilitySet) ToAvailabilitySetPtrOutput() AvailabilitySetPtrOutput {
	return i.ToAvailabilitySetPtrOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetPtrOutputWithContext(ctx context.Context) AvailabilitySetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetPtrOutput)
}

type AvailabilitySetPtrInput interface {
	pulumi.Input

	ToAvailabilitySetPtrOutput() AvailabilitySetPtrOutput
	ToAvailabilitySetPtrOutputWithContext(ctx context.Context) AvailabilitySetPtrOutput
}

type availabilitySetPtrType AvailabilitySetArgs

func (*availabilitySetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil))
}

func (i *availabilitySetPtrType) ToAvailabilitySetPtrOutput() AvailabilitySetPtrOutput {
	return i.ToAvailabilitySetPtrOutputWithContext(context.Background())
}

func (i *availabilitySetPtrType) ToAvailabilitySetPtrOutputWithContext(ctx context.Context) AvailabilitySetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetPtrOutput)
}

// AvailabilitySetArrayInput is an input type that accepts AvailabilitySetArray and AvailabilitySetArrayOutput values.
// You can construct a concrete instance of `AvailabilitySetArrayInput` via:
//
//          AvailabilitySetArray{ AvailabilitySetArgs{...} }
type AvailabilitySetArrayInput interface {
	pulumi.Input

	ToAvailabilitySetArrayOutput() AvailabilitySetArrayOutput
	ToAvailabilitySetArrayOutputWithContext(context.Context) AvailabilitySetArrayOutput
}

type AvailabilitySetArray []AvailabilitySetInput

func (AvailabilitySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AvailabilitySet)(nil)).Elem()
}

func (i AvailabilitySetArray) ToAvailabilitySetArrayOutput() AvailabilitySetArrayOutput {
	return i.ToAvailabilitySetArrayOutputWithContext(context.Background())
}

func (i AvailabilitySetArray) ToAvailabilitySetArrayOutputWithContext(ctx context.Context) AvailabilitySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetArrayOutput)
}

// AvailabilitySetMapInput is an input type that accepts AvailabilitySetMap and AvailabilitySetMapOutput values.
// You can construct a concrete instance of `AvailabilitySetMapInput` via:
//
//          AvailabilitySetMap{ "key": AvailabilitySetArgs{...} }
type AvailabilitySetMapInput interface {
	pulumi.Input

	ToAvailabilitySetMapOutput() AvailabilitySetMapOutput
	ToAvailabilitySetMapOutputWithContext(context.Context) AvailabilitySetMapOutput
}

type AvailabilitySetMap map[string]AvailabilitySetInput

func (AvailabilitySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AvailabilitySet)(nil)).Elem()
}

func (i AvailabilitySetMap) ToAvailabilitySetMapOutput() AvailabilitySetMapOutput {
	return i.ToAvailabilitySetMapOutputWithContext(context.Background())
}

func (i AvailabilitySetMap) ToAvailabilitySetMapOutputWithContext(ctx context.Context) AvailabilitySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetMapOutput)
}

type AvailabilitySetOutput struct{ *pulumi.OutputState }

func (AvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySet)(nil))
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetPtrOutput() AvailabilitySetPtrOutput {
	return o.ToAvailabilitySetPtrOutputWithContext(context.Background())
}

func (o AvailabilitySetOutput) ToAvailabilitySetPtrOutputWithContext(ctx context.Context) AvailabilitySetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AvailabilitySet) *AvailabilitySet {
		return &v
	}).(AvailabilitySetPtrOutput)
}

type AvailabilitySetPtrOutput struct{ *pulumi.OutputState }

func (AvailabilitySetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil))
}

func (o AvailabilitySetPtrOutput) ToAvailabilitySetPtrOutput() AvailabilitySetPtrOutput {
	return o
}

func (o AvailabilitySetPtrOutput) ToAvailabilitySetPtrOutputWithContext(ctx context.Context) AvailabilitySetPtrOutput {
	return o
}

func (o AvailabilitySetPtrOutput) Elem() AvailabilitySetOutput {
	return o.ApplyT(func(v *AvailabilitySet) AvailabilitySet {
		if v != nil {
			return *v
		}
		var ret AvailabilitySet
		return ret
	}).(AvailabilitySetOutput)
}

type AvailabilitySetArrayOutput struct{ *pulumi.OutputState }

func (AvailabilitySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailabilitySet)(nil))
}

func (o AvailabilitySetArrayOutput) ToAvailabilitySetArrayOutput() AvailabilitySetArrayOutput {
	return o
}

func (o AvailabilitySetArrayOutput) ToAvailabilitySetArrayOutputWithContext(ctx context.Context) AvailabilitySetArrayOutput {
	return o
}

func (o AvailabilitySetArrayOutput) Index(i pulumi.IntInput) AvailabilitySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AvailabilitySet {
		return vs[0].([]AvailabilitySet)[vs[1].(int)]
	}).(AvailabilitySetOutput)
}

type AvailabilitySetMapOutput struct{ *pulumi.OutputState }

func (AvailabilitySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AvailabilitySet)(nil))
}

func (o AvailabilitySetMapOutput) ToAvailabilitySetMapOutput() AvailabilitySetMapOutput {
	return o
}

func (o AvailabilitySetMapOutput) ToAvailabilitySetMapOutputWithContext(ctx context.Context) AvailabilitySetMapOutput {
	return o
}

func (o AvailabilitySetMapOutput) MapIndex(k pulumi.StringInput) AvailabilitySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AvailabilitySet {
		return vs[0].(map[string]AvailabilitySet)[vs[1].(string)]
	}).(AvailabilitySetOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
	pulumi.RegisterOutputType(AvailabilitySetPtrOutput{})
	pulumi.RegisterOutputType(AvailabilitySetArrayOutput{})
	pulumi.RegisterOutputType(AvailabilitySetMapOutput{})
}
