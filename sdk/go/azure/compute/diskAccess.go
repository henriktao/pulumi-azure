// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Disk Access.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewDiskAccess(ctx, "example", &compute.DiskAccessArgs{
// 			Location:          pulumi.String("West Europe"),
// 			ResourceGroupName: pulumi.String("example"),
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
// Disk Access resource can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/diskAccess:DiskAccess example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/diskAccesses/diskAccess1
// ```
type DiskAccess struct {
	pulumi.CustomResourceState

	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Disk Access. Changing this forces a new Disk Access to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Disk Access.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDiskAccess registers a new resource with the given unique name, arguments, and options.
func NewDiskAccess(ctx *pulumi.Context,
	name string, args *DiskAccessArgs, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DiskAccess
	err := ctx.RegisterResource("azure:compute/diskAccess:DiskAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAccess gets an existing DiskAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessState, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	var resource DiskAccess
	err := ctx.ReadResource("azure:compute/diskAccess:DiskAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAccess resources.
type diskAccessState struct {
	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Disk Access. Changing this forces a new Disk Access to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Disk Access.
	Tags map[string]string `pulumi:"tags"`
}

type DiskAccessState struct {
	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Disk Access. Changing this forces a new Disk Access to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Disk Access.
	Tags pulumi.StringMapInput
}

func (DiskAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessState)(nil)).Elem()
}

type diskAccessArgs struct {
	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Disk Access. Changing this forces a new Disk Access to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Disk Access.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskAccess resource.
type DiskAccessArgs struct {
	// The Azure Region where the Disk Access should exist. Changing this forces a new Disk to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Disk Access. Changing this forces a new Disk Access to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Disk Access should exist. Changing this forces a new Disk Access to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Disk Access.
	Tags pulumi.StringMapInput
}

func (DiskAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessArgs)(nil)).Elem()
}

type DiskAccessInput interface {
	pulumi.Input

	ToDiskAccessOutput() DiskAccessOutput
	ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput
}

func (*DiskAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (i *DiskAccess) ToDiskAccessOutput() DiskAccessOutput {
	return i.ToDiskAccessOutputWithContext(context.Background())
}

func (i *DiskAccess) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessOutput)
}

// DiskAccessArrayInput is an input type that accepts DiskAccessArray and DiskAccessArrayOutput values.
// You can construct a concrete instance of `DiskAccessArrayInput` via:
//
//          DiskAccessArray{ DiskAccessArgs{...} }
type DiskAccessArrayInput interface {
	pulumi.Input

	ToDiskAccessArrayOutput() DiskAccessArrayOutput
	ToDiskAccessArrayOutputWithContext(context.Context) DiskAccessArrayOutput
}

type DiskAccessArray []DiskAccessInput

func (DiskAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskAccess)(nil)).Elem()
}

func (i DiskAccessArray) ToDiskAccessArrayOutput() DiskAccessArrayOutput {
	return i.ToDiskAccessArrayOutputWithContext(context.Background())
}

func (i DiskAccessArray) ToDiskAccessArrayOutputWithContext(ctx context.Context) DiskAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessArrayOutput)
}

// DiskAccessMapInput is an input type that accepts DiskAccessMap and DiskAccessMapOutput values.
// You can construct a concrete instance of `DiskAccessMapInput` via:
//
//          DiskAccessMap{ "key": DiskAccessArgs{...} }
type DiskAccessMapInput interface {
	pulumi.Input

	ToDiskAccessMapOutput() DiskAccessMapOutput
	ToDiskAccessMapOutputWithContext(context.Context) DiskAccessMapOutput
}

type DiskAccessMap map[string]DiskAccessInput

func (DiskAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskAccess)(nil)).Elem()
}

func (i DiskAccessMap) ToDiskAccessMapOutput() DiskAccessMapOutput {
	return i.ToDiskAccessMapOutputWithContext(context.Background())
}

func (i DiskAccessMap) ToDiskAccessMapOutputWithContext(ctx context.Context) DiskAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessMapOutput)
}

type DiskAccessOutput struct{ *pulumi.OutputState }

func (DiskAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (o DiskAccessOutput) ToDiskAccessOutput() DiskAccessOutput {
	return o
}

func (o DiskAccessOutput) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return o
}

type DiskAccessArrayOutput struct{ *pulumi.OutputState }

func (DiskAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskAccess)(nil)).Elem()
}

func (o DiskAccessArrayOutput) ToDiskAccessArrayOutput() DiskAccessArrayOutput {
	return o
}

func (o DiskAccessArrayOutput) ToDiskAccessArrayOutputWithContext(ctx context.Context) DiskAccessArrayOutput {
	return o
}

func (o DiskAccessArrayOutput) Index(i pulumi.IntInput) DiskAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DiskAccess {
		return vs[0].([]*DiskAccess)[vs[1].(int)]
	}).(DiskAccessOutput)
}

type DiskAccessMapOutput struct{ *pulumi.OutputState }

func (DiskAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskAccess)(nil)).Elem()
}

func (o DiskAccessMapOutput) ToDiskAccessMapOutput() DiskAccessMapOutput {
	return o
}

func (o DiskAccessMapOutput) ToDiskAccessMapOutputWithContext(ctx context.Context) DiskAccessMapOutput {
	return o
}

func (o DiskAccessMapOutput) MapIndex(k pulumi.StringInput) DiskAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DiskAccess {
		return vs[0].(map[string]*DiskAccess)[vs[1].(string)]
	}).(DiskAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskAccessInput)(nil)).Elem(), &DiskAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskAccessArrayInput)(nil)).Elem(), DiskAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskAccessMapInput)(nil)).Elem(), DiskAccessMap{})
	pulumi.RegisterOutputType(DiskAccessOutput{})
	pulumi.RegisterOutputType(DiskAccessArrayOutput{})
	pulumi.RegisterOutputType(DiskAccessMapOutput{})
}
