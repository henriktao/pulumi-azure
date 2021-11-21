// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Cassandra Datacenter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
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
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              exampleVirtualNetwork.ID(),
// 			RoleDefinitionName: pulumi.String("Network Contributor"),
// 			PrincipalId:        pulumi.String("e5007d2c-4b13-4a74-9b6a-605d99f03501"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCassandraCluster, err := cosmosdb.NewCassandraCluster(ctx, "exampleCassandraCluster", &cosmosdb.CassandraClusterArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			Location:                    exampleResourceGroup.Location,
// 			DelegatedManagementSubnetId: exampleSubnet.ID(),
// 			DefaultAdminPassword:        pulumi.String("Password1234"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewCassandraDatacenter(ctx, "exampleCassandraDatacenter", &cosmosdb.CassandraDatacenterArgs{
// 			Location:                    exampleCassandraCluster.Location,
// 			CassandraClusterId:          exampleCassandraCluster.ID(),
// 			DelegatedManagementSubnetId: exampleSubnet.ID(),
// 			NodeCount:                   pulumi.Int(3),
// 			DiskCount:                   pulumi.Int(4),
// 			SkuName:                     pulumi.String("Standard_DS14_v2"),
// 			AvailabilityZonesEnabled:    pulumi.Bool(false),
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
// Cassandra Datacenters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/cassandraDatacenter:CassandraDatacenter example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/cassandraClusters/cluster1/dataCenters/dc1
// ```
type CassandraDatacenter struct {
	pulumi.CustomResourceState

	// Determines whether availability zones are enabled. Defaults to `true`.
	AvailabilityZonesEnabled pulumi.BoolPtrOutput `pulumi:"availabilityZonesEnabled"`
	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterId pulumi.StringOutput `pulumi:"cassandraClusterId"`
	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetId pulumi.StringOutput `pulumi:"delegatedManagementSubnetId"`
	// Determines the number of p30 disks that are attached to each node. Defaults to `4`.
	DiskCount pulumi.IntPtrOutput `pulumi:"diskCount"`
	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than `3`. Defaults to `3`.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// Determines the selected sku. Defaults to Standard_DS14_v2.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
}

// NewCassandraDatacenter registers a new resource with the given unique name, arguments, and options.
func NewCassandraDatacenter(ctx *pulumi.Context,
	name string, args *CassandraDatacenterArgs, opts ...pulumi.ResourceOption) (*CassandraDatacenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CassandraClusterId == nil {
		return nil, errors.New("invalid value for required argument 'CassandraClusterId'")
	}
	if args.DelegatedManagementSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'DelegatedManagementSubnetId'")
	}
	var resource CassandraDatacenter
	err := ctx.RegisterResource("azure:cosmosdb/cassandraDatacenter:CassandraDatacenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraDatacenter gets an existing CassandraDatacenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraDatacenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraDatacenterState, opts ...pulumi.ResourceOption) (*CassandraDatacenter, error) {
	var resource CassandraDatacenter
	err := ctx.ReadResource("azure:cosmosdb/cassandraDatacenter:CassandraDatacenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraDatacenter resources.
type cassandraDatacenterState struct {
	// Determines whether availability zones are enabled. Defaults to `true`.
	AvailabilityZonesEnabled *bool `pulumi:"availabilityZonesEnabled"`
	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterId *string `pulumi:"cassandraClusterId"`
	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetId *string `pulumi:"delegatedManagementSubnetId"`
	// Determines the number of p30 disks that are attached to each node. Defaults to `4`.
	DiskCount *int `pulumi:"diskCount"`
	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	Name *string `pulumi:"name"`
	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than `3`. Defaults to `3`.
	NodeCount *int `pulumi:"nodeCount"`
	// Determines the selected sku. Defaults to Standard_DS14_v2.
	SkuName *string `pulumi:"skuName"`
}

type CassandraDatacenterState struct {
	// Determines whether availability zones are enabled. Defaults to `true`.
	AvailabilityZonesEnabled pulumi.BoolPtrInput
	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterId pulumi.StringPtrInput
	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetId pulumi.StringPtrInput
	// Determines the number of p30 disks that are attached to each node. Defaults to `4`.
	DiskCount pulumi.IntPtrInput
	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	Name pulumi.StringPtrInput
	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than `3`. Defaults to `3`.
	NodeCount pulumi.IntPtrInput
	// Determines the selected sku. Defaults to Standard_DS14_v2.
	SkuName pulumi.StringPtrInput
}

func (CassandraDatacenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDatacenterState)(nil)).Elem()
}

type cassandraDatacenterArgs struct {
	// Determines whether availability zones are enabled. Defaults to `true`.
	AvailabilityZonesEnabled *bool `pulumi:"availabilityZonesEnabled"`
	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterId string `pulumi:"cassandraClusterId"`
	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetId string `pulumi:"delegatedManagementSubnetId"`
	// Determines the number of p30 disks that are attached to each node. Defaults to `4`.
	DiskCount *int `pulumi:"diskCount"`
	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	Name *string `pulumi:"name"`
	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than `3`. Defaults to `3`.
	NodeCount *int `pulumi:"nodeCount"`
	// Determines the selected sku. Defaults to Standard_DS14_v2.
	SkuName *string `pulumi:"skuName"`
}

// The set of arguments for constructing a CassandraDatacenter resource.
type CassandraDatacenterArgs struct {
	// Determines whether availability zones are enabled. Defaults to `true`.
	AvailabilityZonesEnabled pulumi.BoolPtrInput
	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterId pulumi.StringInput
	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetId pulumi.StringInput
	// Determines the number of p30 disks that are attached to each node. Defaults to `4`.
	DiskCount pulumi.IntPtrInput
	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	Name pulumi.StringPtrInput
	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than `3`. Defaults to `3`.
	NodeCount pulumi.IntPtrInput
	// Determines the selected sku. Defaults to Standard_DS14_v2.
	SkuName pulumi.StringPtrInput
}

func (CassandraDatacenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDatacenterArgs)(nil)).Elem()
}

type CassandraDatacenterInput interface {
	pulumi.Input

	ToCassandraDatacenterOutput() CassandraDatacenterOutput
	ToCassandraDatacenterOutputWithContext(ctx context.Context) CassandraDatacenterOutput
}

func (*CassandraDatacenter) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraDatacenter)(nil))
}

func (i *CassandraDatacenter) ToCassandraDatacenterOutput() CassandraDatacenterOutput {
	return i.ToCassandraDatacenterOutputWithContext(context.Background())
}

func (i *CassandraDatacenter) ToCassandraDatacenterOutputWithContext(ctx context.Context) CassandraDatacenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDatacenterOutput)
}

func (i *CassandraDatacenter) ToCassandraDatacenterPtrOutput() CassandraDatacenterPtrOutput {
	return i.ToCassandraDatacenterPtrOutputWithContext(context.Background())
}

func (i *CassandraDatacenter) ToCassandraDatacenterPtrOutputWithContext(ctx context.Context) CassandraDatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDatacenterPtrOutput)
}

type CassandraDatacenterPtrInput interface {
	pulumi.Input

	ToCassandraDatacenterPtrOutput() CassandraDatacenterPtrOutput
	ToCassandraDatacenterPtrOutputWithContext(ctx context.Context) CassandraDatacenterPtrOutput
}

type cassandraDatacenterPtrType CassandraDatacenterArgs

func (*cassandraDatacenterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDatacenter)(nil))
}

func (i *cassandraDatacenterPtrType) ToCassandraDatacenterPtrOutput() CassandraDatacenterPtrOutput {
	return i.ToCassandraDatacenterPtrOutputWithContext(context.Background())
}

func (i *cassandraDatacenterPtrType) ToCassandraDatacenterPtrOutputWithContext(ctx context.Context) CassandraDatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDatacenterPtrOutput)
}

// CassandraDatacenterArrayInput is an input type that accepts CassandraDatacenterArray and CassandraDatacenterArrayOutput values.
// You can construct a concrete instance of `CassandraDatacenterArrayInput` via:
//
//          CassandraDatacenterArray{ CassandraDatacenterArgs{...} }
type CassandraDatacenterArrayInput interface {
	pulumi.Input

	ToCassandraDatacenterArrayOutput() CassandraDatacenterArrayOutput
	ToCassandraDatacenterArrayOutputWithContext(context.Context) CassandraDatacenterArrayOutput
}

type CassandraDatacenterArray []CassandraDatacenterInput

func (CassandraDatacenterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CassandraDatacenter)(nil)).Elem()
}

func (i CassandraDatacenterArray) ToCassandraDatacenterArrayOutput() CassandraDatacenterArrayOutput {
	return i.ToCassandraDatacenterArrayOutputWithContext(context.Background())
}

func (i CassandraDatacenterArray) ToCassandraDatacenterArrayOutputWithContext(ctx context.Context) CassandraDatacenterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDatacenterArrayOutput)
}

// CassandraDatacenterMapInput is an input type that accepts CassandraDatacenterMap and CassandraDatacenterMapOutput values.
// You can construct a concrete instance of `CassandraDatacenterMapInput` via:
//
//          CassandraDatacenterMap{ "key": CassandraDatacenterArgs{...} }
type CassandraDatacenterMapInput interface {
	pulumi.Input

	ToCassandraDatacenterMapOutput() CassandraDatacenterMapOutput
	ToCassandraDatacenterMapOutputWithContext(context.Context) CassandraDatacenterMapOutput
}

type CassandraDatacenterMap map[string]CassandraDatacenterInput

func (CassandraDatacenterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CassandraDatacenter)(nil)).Elem()
}

func (i CassandraDatacenterMap) ToCassandraDatacenterMapOutput() CassandraDatacenterMapOutput {
	return i.ToCassandraDatacenterMapOutputWithContext(context.Background())
}

func (i CassandraDatacenterMap) ToCassandraDatacenterMapOutputWithContext(ctx context.Context) CassandraDatacenterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDatacenterMapOutput)
}

type CassandraDatacenterOutput struct{ *pulumi.OutputState }

func (CassandraDatacenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraDatacenter)(nil))
}

func (o CassandraDatacenterOutput) ToCassandraDatacenterOutput() CassandraDatacenterOutput {
	return o
}

func (o CassandraDatacenterOutput) ToCassandraDatacenterOutputWithContext(ctx context.Context) CassandraDatacenterOutput {
	return o
}

func (o CassandraDatacenterOutput) ToCassandraDatacenterPtrOutput() CassandraDatacenterPtrOutput {
	return o.ToCassandraDatacenterPtrOutputWithContext(context.Background())
}

func (o CassandraDatacenterOutput) ToCassandraDatacenterPtrOutputWithContext(ctx context.Context) CassandraDatacenterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraDatacenter) *CassandraDatacenter {
		return &v
	}).(CassandraDatacenterPtrOutput)
}

type CassandraDatacenterPtrOutput struct{ *pulumi.OutputState }

func (CassandraDatacenterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDatacenter)(nil))
}

func (o CassandraDatacenterPtrOutput) ToCassandraDatacenterPtrOutput() CassandraDatacenterPtrOutput {
	return o
}

func (o CassandraDatacenterPtrOutput) ToCassandraDatacenterPtrOutputWithContext(ctx context.Context) CassandraDatacenterPtrOutput {
	return o
}

func (o CassandraDatacenterPtrOutput) Elem() CassandraDatacenterOutput {
	return o.ApplyT(func(v *CassandraDatacenter) CassandraDatacenter {
		if v != nil {
			return *v
		}
		var ret CassandraDatacenter
		return ret
	}).(CassandraDatacenterOutput)
}

type CassandraDatacenterArrayOutput struct{ *pulumi.OutputState }

func (CassandraDatacenterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraDatacenter)(nil))
}

func (o CassandraDatacenterArrayOutput) ToCassandraDatacenterArrayOutput() CassandraDatacenterArrayOutput {
	return o
}

func (o CassandraDatacenterArrayOutput) ToCassandraDatacenterArrayOutputWithContext(ctx context.Context) CassandraDatacenterArrayOutput {
	return o
}

func (o CassandraDatacenterArrayOutput) Index(i pulumi.IntInput) CassandraDatacenterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraDatacenter {
		return vs[0].([]CassandraDatacenter)[vs[1].(int)]
	}).(CassandraDatacenterOutput)
}

type CassandraDatacenterMapOutput struct{ *pulumi.OutputState }

func (CassandraDatacenterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CassandraDatacenter)(nil))
}

func (o CassandraDatacenterMapOutput) ToCassandraDatacenterMapOutput() CassandraDatacenterMapOutput {
	return o
}

func (o CassandraDatacenterMapOutput) ToCassandraDatacenterMapOutputWithContext(ctx context.Context) CassandraDatacenterMapOutput {
	return o
}

func (o CassandraDatacenterMapOutput) MapIndex(k pulumi.StringInput) CassandraDatacenterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CassandraDatacenter {
		return vs[0].(map[string]CassandraDatacenter)[vs[1].(string)]
	}).(CassandraDatacenterOutput)
}

func init() {
	pulumi.RegisterOutputType(CassandraDatacenterOutput{})
	pulumi.RegisterOutputType(CassandraDatacenterPtrOutput{})
	pulumi.RegisterOutputType(CassandraDatacenterArrayOutput{})
	pulumi.RegisterOutputType(CassandraDatacenterMapOutput{})
}
