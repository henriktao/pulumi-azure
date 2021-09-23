// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proximity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a proximity placement group for virtual machines, virtual machine scale sets and availability sets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/proximity"
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
// 		_, err = proximity.NewPlacementGroup(ctx, "examplePlacementGroup", &proximity.PlacementGroupArgs{
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
// Proximity Placement Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:proximity/placementGroup:PlacementGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/proximityPlacementGroups/example-ppg
// ```
type PlacementGroup struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PlacementGroup
	err := ctx.RegisterResource("azure:proximity/placementGroup:PlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementGroupState, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	var resource PlacementGroup
	err := ctx.ReadResource("azure:proximity/placementGroup:PlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementGroup resources.
type placementGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type PlacementGroupState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

type PlacementGroupInput interface {
	pulumi.Input

	ToPlacementGroupOutput() PlacementGroupOutput
	ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput
}

func (*PlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (i *PlacementGroup) ToPlacementGroupOutput() PlacementGroupOutput {
	return i.ToPlacementGroupOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupOutput)
}

func (i *PlacementGroup) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return i.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupPtrOutput)
}

type PlacementGroupPtrInput interface {
	pulumi.Input

	ToPlacementGroupPtrOutput() PlacementGroupPtrOutput
	ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput
}

type placementGroupPtrType PlacementGroupArgs

func (*placementGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil))
}

func (i *placementGroupPtrType) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return i.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (i *placementGroupPtrType) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupPtrOutput)
}

// PlacementGroupArrayInput is an input type that accepts PlacementGroupArray and PlacementGroupArrayOutput values.
// You can construct a concrete instance of `PlacementGroupArrayInput` via:
//
//          PlacementGroupArray{ PlacementGroupArgs{...} }
type PlacementGroupArrayInput interface {
	pulumi.Input

	ToPlacementGroupArrayOutput() PlacementGroupArrayOutput
	ToPlacementGroupArrayOutputWithContext(context.Context) PlacementGroupArrayOutput
}

type PlacementGroupArray []PlacementGroupInput

func (PlacementGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlacementGroup)(nil)).Elem()
}

func (i PlacementGroupArray) ToPlacementGroupArrayOutput() PlacementGroupArrayOutput {
	return i.ToPlacementGroupArrayOutputWithContext(context.Background())
}

func (i PlacementGroupArray) ToPlacementGroupArrayOutputWithContext(ctx context.Context) PlacementGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupArrayOutput)
}

// PlacementGroupMapInput is an input type that accepts PlacementGroupMap and PlacementGroupMapOutput values.
// You can construct a concrete instance of `PlacementGroupMapInput` via:
//
//          PlacementGroupMap{ "key": PlacementGroupArgs{...} }
type PlacementGroupMapInput interface {
	pulumi.Input

	ToPlacementGroupMapOutput() PlacementGroupMapOutput
	ToPlacementGroupMapOutputWithContext(context.Context) PlacementGroupMapOutput
}

type PlacementGroupMap map[string]PlacementGroupInput

func (PlacementGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlacementGroup)(nil)).Elem()
}

func (i PlacementGroupMap) ToPlacementGroupMapOutput() PlacementGroupMapOutput {
	return i.ToPlacementGroupMapOutputWithContext(context.Background())
}

func (i PlacementGroupMap) ToPlacementGroupMapOutputWithContext(ctx context.Context) PlacementGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupMapOutput)
}

type PlacementGroupOutput struct{ *pulumi.OutputState }

func (PlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (o PlacementGroupOutput) ToPlacementGroupOutput() PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return o.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (o PlacementGroupOutput) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlacementGroup) *PlacementGroup {
		return &v
	}).(PlacementGroupPtrOutput)
}

type PlacementGroupPtrOutput struct{ *pulumi.OutputState }

func (PlacementGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil))
}

func (o PlacementGroupPtrOutput) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return o
}

func (o PlacementGroupPtrOutput) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return o
}

func (o PlacementGroupPtrOutput) Elem() PlacementGroupOutput {
	return o.ApplyT(func(v *PlacementGroup) PlacementGroup {
		if v != nil {
			return *v
		}
		var ret PlacementGroup
		return ret
	}).(PlacementGroupOutput)
}

type PlacementGroupArrayOutput struct{ *pulumi.OutputState }

func (PlacementGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementGroup)(nil))
}

func (o PlacementGroupArrayOutput) ToPlacementGroupArrayOutput() PlacementGroupArrayOutput {
	return o
}

func (o PlacementGroupArrayOutput) ToPlacementGroupArrayOutputWithContext(ctx context.Context) PlacementGroupArrayOutput {
	return o
}

func (o PlacementGroupArrayOutput) Index(i pulumi.IntInput) PlacementGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlacementGroup {
		return vs[0].([]PlacementGroup)[vs[1].(int)]
	}).(PlacementGroupOutput)
}

type PlacementGroupMapOutput struct{ *pulumi.OutputState }

func (PlacementGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PlacementGroup)(nil))
}

func (o PlacementGroupMapOutput) ToPlacementGroupMapOutput() PlacementGroupMapOutput {
	return o
}

func (o PlacementGroupMapOutput) ToPlacementGroupMapOutputWithContext(ctx context.Context) PlacementGroupMapOutput {
	return o
}

func (o PlacementGroupMapOutput) MapIndex(k pulumi.StringInput) PlacementGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PlacementGroup {
		return vs[0].(map[string]PlacementGroup)[vs[1].(string)]
	}).(PlacementGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(PlacementGroupOutput{})
	pulumi.RegisterOutputType(PlacementGroupPtrOutput{})
	pulumi.RegisterOutputType(PlacementGroupArrayOutput{})
	pulumi.RegisterOutputType(PlacementGroupMapOutput{})
}
