// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Pool within a NetApp Account.
//
// ## NetApp Pool Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/netapp"
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
// 		exampleAccount, err := netapp.NewAccount(ctx, "exampleAccount", &netapp.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = netapp.NewPool(ctx, "examplePool", &netapp.PoolArgs{
// 			AccountName:       exampleAccount.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceLevel:      pulumi.String("Premium"),
// 			SizeInTb:          pulumi.Int(4),
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
// NetApp Pool can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:netapp/pool:Pool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1
// ```
type Pool struct {
	pulumi.CustomResourceState

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the NetApp Pool. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// QoS Type of the pool. Valid values include `Auto` or `Manual`.
	QosType pulumi.StringOutput `pulumi:"qosType"`
	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
	ServiceLevel pulumi.StringOutput `pulumi:"serviceLevel"`
	// Provisioned size of the pool in TB. Value must be between `4` and `500`.
	SizeInTb pulumi.IntOutput `pulumi:"sizeInTb"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceLevel == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLevel'")
	}
	if args.SizeInTb == nil {
		return nil, errors.New("invalid value for required argument 'SizeInTb'")
	}
	var resource Pool
	err := ctx.RegisterResource("azure:netapp/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure:netapp/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the NetApp Pool. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// QoS Type of the pool. Valid values include `Auto` or `Manual`.
	QosType *string `pulumi:"qosType"`
	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
	ServiceLevel *string `pulumi:"serviceLevel"`
	// Provisioned size of the pool in TB. Value must be between `4` and `500`.
	SizeInTb *int `pulumi:"sizeInTb"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type PoolState struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the NetApp Pool. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// QoS Type of the pool. Valid values include `Auto` or `Manual`.
	QosType pulumi.StringPtrInput
	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
	ServiceLevel pulumi.StringPtrInput
	// Provisioned size of the pool in TB. Value must be between `4` and `500`.
	SizeInTb pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the NetApp Pool. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// QoS Type of the pool. Valid values include `Auto` or `Manual`.
	QosType *string `pulumi:"qosType"`
	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool in TB. Value must be between `4` and `500`.
	SizeInTb int `pulumi:"sizeInTb"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the NetApp Pool. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// QoS Type of the pool. Valid values include `Auto` or `Manual`.
	QosType pulumi.StringPtrInput
	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The service level of the file system. Valid values include `Premium`, `Standard`, or `Ultra`. Changing this forces a new resource to be created.
	ServiceLevel pulumi.StringInput
	// Provisioned size of the pool in TB. Value must be between `4` and `500`.
	SizeInTb pulumi.IntInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//          PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//          PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].([]*Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].(map[string]*Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PoolInput)(nil)).Elem(), &Pool{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolArrayInput)(nil)).Elem(), PoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolMapInput)(nil)).Elem(), PoolMap{})
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
