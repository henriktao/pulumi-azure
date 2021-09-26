// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Microsoft Azure SQL Failover Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
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
// 		primary, err := mssql.NewServer(ctx, "primary", &mssql.ServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("missadministrator"),
// 			AdministratorLoginPassword: pulumi.String("thisIsKat11"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := mssql.NewServer(ctx, "secondary", &mssql.ServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("missadministrator"),
// 			AdministratorLoginPassword: pulumi.String("thisIsKat12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDatabase, err := mssql.NewDatabase(ctx, "exampleDatabase", &mssql.DatabaseArgs{
// 			ServerId:  primary.ID(),
// 			SkuName:   pulumi.String("S1"),
// 			Collation: pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
// 			MaxSizeGb: pulumi.Int(200),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewFailoverGroup(ctx, "exampleFailoverGroup", &mssql.FailoverGroupArgs{
// 			ServerId: primary.ID(),
// 			Databases: pulumi.StringArray{
// 				exampleDatabase.ID(),
// 			},
// 			PartnerServers: mssql.FailoverGroupPartnerServerArray{
// 				&mssql.FailoverGroupPartnerServerArgs{
// 					Id: secondary.ID(),
// 				},
// 			},
// 			ReadWriteEndpointFailoverPolicy: &mssql.FailoverGroupReadWriteEndpointFailoverPolicyArgs{
// 				Mode:         pulumi.String("Automatic"),
// 				GraceMinutes: pulumi.Int(80),
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("prod"),
// 				"database":    pulumi.String("example"),
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
// Failover Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/failoverGroup:FailoverGroup example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/servers/server1/failoverGroups/failoverGroup1
// ```
type FailoverGroup struct {
	pulumi.CustomResourceState

	// A set of database names to include in the failover group.
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	// The name of the Failover Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `partnerServer` block as defined below.
	PartnerServers FailoverGroupPartnerServerArrayOutput `pulumi:"partnerServers"`
	// A `readWriteEndpointFailoverPolicy` block as defined below.
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyOutput `pulumi:"readWriteEndpointFailoverPolicy"`
	// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
	ReadonlyEndpointFailoverPolicyEnabled pulumi.BoolOutput `pulumi:"readonlyEndpointFailoverPolicyEnabled"`
	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServers == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServers'")
	}
	if args.ReadWriteEndpointFailoverPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpointFailoverPolicy'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	var resource FailoverGroup
	err := ctx.RegisterResource("azure:mssql/failoverGroup:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure:mssql/failoverGroup:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
	// A set of database names to include in the failover group.
	Databases []string `pulumi:"databases"`
	// The name of the Failover Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `partnerServer` block as defined below.
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A `readWriteEndpointFailoverPolicy` block as defined below.
	ReadWriteEndpointFailoverPolicy *FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
	ReadonlyEndpointFailoverPolicyEnabled *bool `pulumi:"readonlyEndpointFailoverPolicyEnabled"`
	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerId *string `pulumi:"serverId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type FailoverGroupState struct {
	// A set of database names to include in the failover group.
	Databases pulumi.StringArrayInput
	// The name of the Failover Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `partnerServer` block as defined below.
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A `readWriteEndpointFailoverPolicy` block as defined below.
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyPtrInput
	// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
	ReadonlyEndpointFailoverPolicyEnabled pulumi.BoolPtrInput
	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// A set of database names to include in the failover group.
	Databases []string `pulumi:"databases"`
	// The name of the Failover Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `partnerServer` block as defined below.
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A `readWriteEndpointFailoverPolicy` block as defined below.
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
	ReadonlyEndpointFailoverPolicyEnabled *bool `pulumi:"readonlyEndpointFailoverPolicyEnabled"`
	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerId string `pulumi:"serverId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// A set of database names to include in the failover group.
	Databases pulumi.StringArrayInput
	// The name of the Failover Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `partnerServer` block as defined below.
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A `readWriteEndpointFailoverPolicy` block as defined below.
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput
	// Whether failover is enabled for the readonly endpoint. Defaults to `false`.
	ReadonlyEndpointFailoverPolicyEnabled pulumi.BoolPtrInput
	// The ID of the primary SQL Server on which to create the failover group. Changing this forces a new resource to be created.
	ServerId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (*FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroup)(nil))
}

func (i *FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

func (i *FailoverGroup) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return i.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupPtrOutput)
}

type FailoverGroupPtrInput interface {
	pulumi.Input

	ToFailoverGroupPtrOutput() FailoverGroupPtrOutput
	ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput
}

type failoverGroupPtrType FailoverGroupArgs

func (*failoverGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil))
}

func (i *failoverGroupPtrType) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return i.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (i *failoverGroupPtrType) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupPtrOutput)
}

// FailoverGroupArrayInput is an input type that accepts FailoverGroupArray and FailoverGroupArrayOutput values.
// You can construct a concrete instance of `FailoverGroupArrayInput` via:
//
//          FailoverGroupArray{ FailoverGroupArgs{...} }
type FailoverGroupArrayInput interface {
	pulumi.Input

	ToFailoverGroupArrayOutput() FailoverGroupArrayOutput
	ToFailoverGroupArrayOutputWithContext(context.Context) FailoverGroupArrayOutput
}

type FailoverGroupArray []FailoverGroupInput

func (FailoverGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FailoverGroup)(nil)).Elem()
}

func (i FailoverGroupArray) ToFailoverGroupArrayOutput() FailoverGroupArrayOutput {
	return i.ToFailoverGroupArrayOutputWithContext(context.Background())
}

func (i FailoverGroupArray) ToFailoverGroupArrayOutputWithContext(ctx context.Context) FailoverGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupArrayOutput)
}

// FailoverGroupMapInput is an input type that accepts FailoverGroupMap and FailoverGroupMapOutput values.
// You can construct a concrete instance of `FailoverGroupMapInput` via:
//
//          FailoverGroupMap{ "key": FailoverGroupArgs{...} }
type FailoverGroupMapInput interface {
	pulumi.Input

	ToFailoverGroupMapOutput() FailoverGroupMapOutput
	ToFailoverGroupMapOutputWithContext(context.Context) FailoverGroupMapOutput
}

type FailoverGroupMap map[string]FailoverGroupInput

func (FailoverGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FailoverGroup)(nil)).Elem()
}

func (i FailoverGroupMap) ToFailoverGroupMapOutput() FailoverGroupMapOutput {
	return i.ToFailoverGroupMapOutputWithContext(context.Background())
}

func (i FailoverGroupMap) ToFailoverGroupMapOutputWithContext(ctx context.Context) FailoverGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupMapOutput)
}

type FailoverGroupOutput struct{ *pulumi.OutputState }

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroup)(nil))
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return o.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (o FailoverGroupOutput) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroup) *FailoverGroup {
		return &v
	}).(FailoverGroupPtrOutput)
}

type FailoverGroupPtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil))
}

func (o FailoverGroupPtrOutput) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return o
}

func (o FailoverGroupPtrOutput) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return o
}

func (o FailoverGroupPtrOutput) Elem() FailoverGroupOutput {
	return o.ApplyT(func(v *FailoverGroup) FailoverGroup {
		if v != nil {
			return *v
		}
		var ret FailoverGroup
		return ret
	}).(FailoverGroupOutput)
}

type FailoverGroupArrayOutput struct{ *pulumi.OutputState }

func (FailoverGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverGroup)(nil))
}

func (o FailoverGroupArrayOutput) ToFailoverGroupArrayOutput() FailoverGroupArrayOutput {
	return o
}

func (o FailoverGroupArrayOutput) ToFailoverGroupArrayOutputWithContext(ctx context.Context) FailoverGroupArrayOutput {
	return o
}

func (o FailoverGroupArrayOutput) Index(i pulumi.IntInput) FailoverGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverGroup {
		return vs[0].([]FailoverGroup)[vs[1].(int)]
	}).(FailoverGroupOutput)
}

type FailoverGroupMapOutput struct{ *pulumi.OutputState }

func (FailoverGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FailoverGroup)(nil))
}

func (o FailoverGroupMapOutput) ToFailoverGroupMapOutput() FailoverGroupMapOutput {
	return o
}

func (o FailoverGroupMapOutput) ToFailoverGroupMapOutputWithContext(ctx context.Context) FailoverGroupMapOutput {
	return o
}

func (o FailoverGroupMapOutput) MapIndex(k pulumi.StringInput) FailoverGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FailoverGroup {
		return vs[0].(map[string]FailoverGroup)[vs[1].(string)]
	}).(FailoverGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
	pulumi.RegisterOutputType(FailoverGroupPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupArrayOutput{})
	pulumi.RegisterOutputType(FailoverGroupMapOutput{})
}
