// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Digital Twins instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/digitaltwins"
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
// 		_, err = digitaltwins.NewInstance(ctx, "exampleInstance", &digitaltwins.InstanceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// Digital Twins instances can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:digitaltwins/instance:Instance example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The Api endpoint to work with this Digital Twins instance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The Azure Region where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Digital Twins instance. Changing this forces a new Digital Twins instance to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Digital Twins instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Instance
	err := ctx.RegisterResource("azure:digitaltwins/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("azure:digitaltwins/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The Api endpoint to work with this Digital Twins instance.
	HostName *string `pulumi:"hostName"`
	// The Azure Region where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Digital Twins instance. Changing this forces a new Digital Twins instance to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Digital Twins instance.
	Tags map[string]string `pulumi:"tags"`
}

type InstanceState struct {
	// The Api endpoint to work with this Digital Twins instance.
	HostName pulumi.StringPtrInput
	// The Azure Region where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Digital Twins instance. Changing this forces a new Digital Twins instance to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Digital Twins instance.
	Tags pulumi.StringMapInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The Azure Region where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Digital Twins instance. Changing this forces a new Digital Twins instance to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Digital Twins instance.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The Azure Region where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Digital Twins instance. Changing this forces a new Digital Twins instance to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Digital Twins instance should exist. Changing this forces a new Digital Twins instance to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Digital Twins instance.
	Tags pulumi.StringMapInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
