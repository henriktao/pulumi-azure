// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Disk Pool.
//
// !> **Note:** This resource has been deprecated in favour of `compute.DiskPool` and will be removed in version 3.0 of the Azure Provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleVirtualNetwork.ResourceGroupName,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("diskspool"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/read"),
// 						},
// 						Name: pulumi.String("Microsoft.StoragePool/diskPools"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewDisksPool(ctx, "exampleDisksPool", &storage.DisksPoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SubnetId:          exampleSubnet.ID(),
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("1"),
// 			},
// 			SkuName: pulumi.String("Basic_B1"),
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
// Disk Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/disksPool:DisksPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/disksPool1
// ```
type DisksPool struct {
	pulumi.CustomResourceState

	// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A mapping of tags which should be assigned to the Disks Pool.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDisksPool registers a new resource with the given unique name, arguments, and options.
func NewDisksPool(ctx *pulumi.Context,
	name string, args *DisksPoolArgs, opts ...pulumi.ResourceOption) (*DisksPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZones'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource DisksPool
	err := ctx.RegisterResource("azure:storage/disksPool:DisksPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisksPool gets an existing DisksPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisksPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisksPoolState, opts ...pulumi.ResourceOption) (*DisksPool, error) {
	var resource DisksPool
	err := ctx.ReadResource("azure:storage/disksPool:DisksPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DisksPool resources.
type disksPoolState struct {
	// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
	Location *string `pulumi:"location"`
	// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
	SkuName *string `pulumi:"skuName"`
	// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags which should be assigned to the Disks Pool.
	Tags map[string]string `pulumi:"tags"`
}

type DisksPoolState struct {
	// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
	AvailabilityZones pulumi.StringArrayInput
	// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
	Location pulumi.StringPtrInput
	// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
	SkuName pulumi.StringPtrInput
	// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Disks Pool.
	Tags pulumi.StringMapInput
}

func (DisksPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*disksPoolState)(nil)).Elem()
}

type disksPoolArgs struct {
	// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
	Location *string `pulumi:"location"`
	// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
	SkuName string `pulumi:"skuName"`
	// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags which should be assigned to the Disks Pool.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DisksPool resource.
type DisksPoolArgs struct {
	// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
	AvailabilityZones pulumi.StringArrayInput
	// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
	Location pulumi.StringPtrInput
	// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
	ResourceGroupName pulumi.StringInput
	// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
	SkuName pulumi.StringInput
	// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
	SubnetId pulumi.StringInput
	// A mapping of tags which should be assigned to the Disks Pool.
	Tags pulumi.StringMapInput
}

func (DisksPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disksPoolArgs)(nil)).Elem()
}

type DisksPoolInput interface {
	pulumi.Input

	ToDisksPoolOutput() DisksPoolOutput
	ToDisksPoolOutputWithContext(ctx context.Context) DisksPoolOutput
}

func (*DisksPool) ElementType() reflect.Type {
	return reflect.TypeOf((**DisksPool)(nil)).Elem()
}

func (i *DisksPool) ToDisksPoolOutput() DisksPoolOutput {
	return i.ToDisksPoolOutputWithContext(context.Background())
}

func (i *DisksPool) ToDisksPoolOutputWithContext(ctx context.Context) DisksPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisksPoolOutput)
}

// DisksPoolArrayInput is an input type that accepts DisksPoolArray and DisksPoolArrayOutput values.
// You can construct a concrete instance of `DisksPoolArrayInput` via:
//
//          DisksPoolArray{ DisksPoolArgs{...} }
type DisksPoolArrayInput interface {
	pulumi.Input

	ToDisksPoolArrayOutput() DisksPoolArrayOutput
	ToDisksPoolArrayOutputWithContext(context.Context) DisksPoolArrayOutput
}

type DisksPoolArray []DisksPoolInput

func (DisksPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DisksPool)(nil)).Elem()
}

func (i DisksPoolArray) ToDisksPoolArrayOutput() DisksPoolArrayOutput {
	return i.ToDisksPoolArrayOutputWithContext(context.Background())
}

func (i DisksPoolArray) ToDisksPoolArrayOutputWithContext(ctx context.Context) DisksPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisksPoolArrayOutput)
}

// DisksPoolMapInput is an input type that accepts DisksPoolMap and DisksPoolMapOutput values.
// You can construct a concrete instance of `DisksPoolMapInput` via:
//
//          DisksPoolMap{ "key": DisksPoolArgs{...} }
type DisksPoolMapInput interface {
	pulumi.Input

	ToDisksPoolMapOutput() DisksPoolMapOutput
	ToDisksPoolMapOutputWithContext(context.Context) DisksPoolMapOutput
}

type DisksPoolMap map[string]DisksPoolInput

func (DisksPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DisksPool)(nil)).Elem()
}

func (i DisksPoolMap) ToDisksPoolMapOutput() DisksPoolMapOutput {
	return i.ToDisksPoolMapOutputWithContext(context.Background())
}

func (i DisksPoolMap) ToDisksPoolMapOutputWithContext(ctx context.Context) DisksPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisksPoolMapOutput)
}

type DisksPoolOutput struct{ *pulumi.OutputState }

func (DisksPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisksPool)(nil)).Elem()
}

func (o DisksPoolOutput) ToDisksPoolOutput() DisksPoolOutput {
	return o
}

func (o DisksPoolOutput) ToDisksPoolOutputWithContext(ctx context.Context) DisksPoolOutput {
	return o
}

type DisksPoolArrayOutput struct{ *pulumi.OutputState }

func (DisksPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DisksPool)(nil)).Elem()
}

func (o DisksPoolArrayOutput) ToDisksPoolArrayOutput() DisksPoolArrayOutput {
	return o
}

func (o DisksPoolArrayOutput) ToDisksPoolArrayOutputWithContext(ctx context.Context) DisksPoolArrayOutput {
	return o
}

func (o DisksPoolArrayOutput) Index(i pulumi.IntInput) DisksPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DisksPool {
		return vs[0].([]*DisksPool)[vs[1].(int)]
	}).(DisksPoolOutput)
}

type DisksPoolMapOutput struct{ *pulumi.OutputState }

func (DisksPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DisksPool)(nil)).Elem()
}

func (o DisksPoolMapOutput) ToDisksPoolMapOutput() DisksPoolMapOutput {
	return o
}

func (o DisksPoolMapOutput) ToDisksPoolMapOutputWithContext(ctx context.Context) DisksPoolMapOutput {
	return o
}

func (o DisksPoolMapOutput) MapIndex(k pulumi.StringInput) DisksPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DisksPool {
		return vs[0].(map[string]*DisksPool)[vs[1].(string)]
	}).(DisksPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DisksPoolInput)(nil)).Elem(), &DisksPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisksPoolArrayInput)(nil)).Elem(), DisksPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisksPoolMapInput)(nil)).Elem(), DisksPoolMap{})
	pulumi.RegisterOutputType(DisksPoolOutput{})
	pulumi.RegisterOutputType(DisksPoolArrayOutput{})
	pulumi.RegisterOutputType(DisksPoolMapOutput{})
}
