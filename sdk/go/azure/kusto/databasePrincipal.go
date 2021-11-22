// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Database Principal
//
// > **NOTE:** This resource is being **deprecated** due to API updates and should no longer be used.  Please use kusto.DatabasePrincipalAssignment instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/kusto"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		cluster, err := kusto.NewCluster(ctx, "cluster", &kusto.ClusterArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewDatabase(ctx, "database", &kusto.DatabaseArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 			ClusterName:       cluster.Name,
// 			HotCachePeriod:    pulumi.String("P7D"),
// 			SoftDeletePeriod:  pulumi.String("P31D"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewDatabasePrincipal(ctx, "principal", &kusto.DatabasePrincipalArgs{
// 			ResourceGroupName: rg.Name,
// 			ClusterName:       cluster.Name,
// 			DatabaseName:      pulumi.Any(azurerm_kusto_database.Test.Name),
// 			Role:              pulumi.String("Viewer"),
// 			Type:              pulumi.String("User"),
// 			ClientId:          pulumi.String(current.TenantId),
// 			ObjectId:          pulumi.String(current.ClientId),
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
// Kusto Database Principals can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/databasePrincipal:DatabasePrincipal example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/Role/role1/FQN/some-guid
// ```
type DatabasePrincipal struct {
	pulumi.CustomResourceState

	// The app id, if not empty, of the principal.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The email, if not empty, of the principal.
	Email pulumi.StringOutput `pulumi:"email"`
	// The fully qualified name of the principal.
	FullyQualifiedName pulumi.StringOutput `pulumi:"fullyQualifiedName"`
	// The name of the Kusto Database Principal.
	Name pulumi.StringOutput `pulumi:"name"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringOutput `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabasePrincipal registers a new resource with the given unique name, arguments, and options.
func NewDatabasePrincipal(ctx *pulumi.Context,
	name string, args *DatabasePrincipalArgs, opts ...pulumi.ResourceOption) (*DatabasePrincipal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DatabasePrincipal
	err := ctx.RegisterResource("azure:kusto/databasePrincipal:DatabasePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasePrincipal gets an existing DatabasePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrincipalState, opts ...pulumi.ResourceOption) (*DatabasePrincipal, error) {
	var resource DatabasePrincipal
	err := ctx.ReadResource("azure:kusto/databasePrincipal:DatabasePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasePrincipal resources.
type databasePrincipalState struct {
	// The app id, if not empty, of the principal.
	AppId *string `pulumi:"appId"`
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
	ClientId *string `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The email, if not empty, of the principal.
	Email *string `pulumi:"email"`
	// The fully qualified name of the principal.
	FullyQualifiedName *string `pulumi:"fullyQualifiedName"`
	// The name of the Kusto Database Principal.
	Name *string `pulumi:"name"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
	ObjectId *string `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role *string `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
}

type DatabasePrincipalState struct {
	// The app id, if not empty, of the principal.
	AppId pulumi.StringPtrInput
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
	ClientId pulumi.StringPtrInput
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The email, if not empty, of the principal.
	Email pulumi.StringPtrInput
	// The fully qualified name of the principal.
	FullyQualifiedName pulumi.StringPtrInput
	// The name of the Kusto Database Principal.
	Name pulumi.StringPtrInput
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
	ObjectId pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringPtrInput
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
}

func (DatabasePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalState)(nil)).Elem()
}

type databasePrincipalArgs struct {
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
	ClientId string `pulumi:"clientId"`
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
	ObjectId string `pulumi:"objectId"`
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role string `pulumi:"role"`
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DatabasePrincipal resource.
type DatabasePrincipalArgs struct {
	// The Client ID that owns the specified `objectId`. Changing this forces a new resource to be created.
	ClientId pulumi.StringInput
	// Specifies the name of the Kusto Cluster this database principal will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specified the name of the Kusto Database this principal will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// An Object ID of a User, Group, or App. Changing this forces a new resource to be created.
	ObjectId pulumi.StringInput
	// Specifies the Resource Group where the Kusto Database Principal should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the permissions the Principal will have. Valid values include `Admin`, `Ingestor`, `Monitor`, `UnrestrictedViewers`, `User`, `Viewer`. Changing this forces a new resource to be created.
	Role pulumi.StringInput
	// Specifies the type of object the principal is. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	Type pulumi.StringInput
}

func (DatabasePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalArgs)(nil)).Elem()
}

type DatabasePrincipalInput interface {
	pulumi.Input

	ToDatabasePrincipalOutput() DatabasePrincipalOutput
	ToDatabasePrincipalOutputWithContext(ctx context.Context) DatabasePrincipalOutput
}

func (*DatabasePrincipal) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipal)(nil))
}

func (i *DatabasePrincipal) ToDatabasePrincipalOutput() DatabasePrincipalOutput {
	return i.ToDatabasePrincipalOutputWithContext(context.Background())
}

func (i *DatabasePrincipal) ToDatabasePrincipalOutputWithContext(ctx context.Context) DatabasePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalOutput)
}

func (i *DatabasePrincipal) ToDatabasePrincipalPtrOutput() DatabasePrincipalPtrOutput {
	return i.ToDatabasePrincipalPtrOutputWithContext(context.Background())
}

func (i *DatabasePrincipal) ToDatabasePrincipalPtrOutputWithContext(ctx context.Context) DatabasePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalPtrOutput)
}

type DatabasePrincipalPtrInput interface {
	pulumi.Input

	ToDatabasePrincipalPtrOutput() DatabasePrincipalPtrOutput
	ToDatabasePrincipalPtrOutputWithContext(ctx context.Context) DatabasePrincipalPtrOutput
}

type databasePrincipalPtrType DatabasePrincipalArgs

func (*databasePrincipalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipal)(nil))
}

func (i *databasePrincipalPtrType) ToDatabasePrincipalPtrOutput() DatabasePrincipalPtrOutput {
	return i.ToDatabasePrincipalPtrOutputWithContext(context.Background())
}

func (i *databasePrincipalPtrType) ToDatabasePrincipalPtrOutputWithContext(ctx context.Context) DatabasePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalPtrOutput)
}

// DatabasePrincipalArrayInput is an input type that accepts DatabasePrincipalArray and DatabasePrincipalArrayOutput values.
// You can construct a concrete instance of `DatabasePrincipalArrayInput` via:
//
//          DatabasePrincipalArray{ DatabasePrincipalArgs{...} }
type DatabasePrincipalArrayInput interface {
	pulumi.Input

	ToDatabasePrincipalArrayOutput() DatabasePrincipalArrayOutput
	ToDatabasePrincipalArrayOutputWithContext(context.Context) DatabasePrincipalArrayOutput
}

type DatabasePrincipalArray []DatabasePrincipalInput

func (DatabasePrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabasePrincipal)(nil)).Elem()
}

func (i DatabasePrincipalArray) ToDatabasePrincipalArrayOutput() DatabasePrincipalArrayOutput {
	return i.ToDatabasePrincipalArrayOutputWithContext(context.Background())
}

func (i DatabasePrincipalArray) ToDatabasePrincipalArrayOutputWithContext(ctx context.Context) DatabasePrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalArrayOutput)
}

// DatabasePrincipalMapInput is an input type that accepts DatabasePrincipalMap and DatabasePrincipalMapOutput values.
// You can construct a concrete instance of `DatabasePrincipalMapInput` via:
//
//          DatabasePrincipalMap{ "key": DatabasePrincipalArgs{...} }
type DatabasePrincipalMapInput interface {
	pulumi.Input

	ToDatabasePrincipalMapOutput() DatabasePrincipalMapOutput
	ToDatabasePrincipalMapOutputWithContext(context.Context) DatabasePrincipalMapOutput
}

type DatabasePrincipalMap map[string]DatabasePrincipalInput

func (DatabasePrincipalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabasePrincipal)(nil)).Elem()
}

func (i DatabasePrincipalMap) ToDatabasePrincipalMapOutput() DatabasePrincipalMapOutput {
	return i.ToDatabasePrincipalMapOutputWithContext(context.Background())
}

func (i DatabasePrincipalMap) ToDatabasePrincipalMapOutputWithContext(ctx context.Context) DatabasePrincipalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalMapOutput)
}

type DatabasePrincipalOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipal)(nil))
}

func (o DatabasePrincipalOutput) ToDatabasePrincipalOutput() DatabasePrincipalOutput {
	return o
}

func (o DatabasePrincipalOutput) ToDatabasePrincipalOutputWithContext(ctx context.Context) DatabasePrincipalOutput {
	return o
}

func (o DatabasePrincipalOutput) ToDatabasePrincipalPtrOutput() DatabasePrincipalPtrOutput {
	return o.ToDatabasePrincipalPtrOutputWithContext(context.Background())
}

func (o DatabasePrincipalOutput) ToDatabasePrincipalPtrOutputWithContext(ctx context.Context) DatabasePrincipalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabasePrincipal) *DatabasePrincipal {
		return &v
	}).(DatabasePrincipalPtrOutput)
}

type DatabasePrincipalPtrOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipal)(nil))
}

func (o DatabasePrincipalPtrOutput) ToDatabasePrincipalPtrOutput() DatabasePrincipalPtrOutput {
	return o
}

func (o DatabasePrincipalPtrOutput) ToDatabasePrincipalPtrOutputWithContext(ctx context.Context) DatabasePrincipalPtrOutput {
	return o
}

func (o DatabasePrincipalPtrOutput) Elem() DatabasePrincipalOutput {
	return o.ApplyT(func(v *DatabasePrincipal) DatabasePrincipal {
		if v != nil {
			return *v
		}
		var ret DatabasePrincipal
		return ret
	}).(DatabasePrincipalOutput)
}

type DatabasePrincipalArrayOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipal)(nil))
}

func (o DatabasePrincipalArrayOutput) ToDatabasePrincipalArrayOutput() DatabasePrincipalArrayOutput {
	return o
}

func (o DatabasePrincipalArrayOutput) ToDatabasePrincipalArrayOutputWithContext(ctx context.Context) DatabasePrincipalArrayOutput {
	return o
}

func (o DatabasePrincipalArrayOutput) Index(i pulumi.IntInput) DatabasePrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabasePrincipal {
		return vs[0].([]DatabasePrincipal)[vs[1].(int)]
	}).(DatabasePrincipalOutput)
}

type DatabasePrincipalMapOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabasePrincipal)(nil))
}

func (o DatabasePrincipalMapOutput) ToDatabasePrincipalMapOutput() DatabasePrincipalMapOutput {
	return o
}

func (o DatabasePrincipalMapOutput) ToDatabasePrincipalMapOutputWithContext(ctx context.Context) DatabasePrincipalMapOutput {
	return o
}

func (o DatabasePrincipalMapOutput) MapIndex(k pulumi.StringInput) DatabasePrincipalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabasePrincipal {
		return vs[0].(map[string]DatabasePrincipal)[vs[1].(string)]
	}).(DatabasePrincipalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalInput)(nil)).Elem(), &DatabasePrincipal{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalPtrInput)(nil)).Elem(), &DatabasePrincipal{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalArrayInput)(nil)).Elem(), DatabasePrincipalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePrincipalMapInput)(nil)).Elem(), DatabasePrincipalMap{})
	pulumi.RegisterOutputType(DatabasePrincipalOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalPtrOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalArrayOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalMapOutput{})
}
