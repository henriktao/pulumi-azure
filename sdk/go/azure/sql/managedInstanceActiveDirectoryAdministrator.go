// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to set a user or group as the AD administrator for an Azure SQL Managed Instance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleManagedInstance, err := sql.NewManagedInstance(ctx, "exampleManagedInstance", &sql.ManagedInstanceArgs{
// 			ResourceGroupName:          pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:                   pulumi.Any(azurerm_resource_group.Example.Location),
// 			AdministratorLogin:         pulumi.String("mradministrator"),
// 			AdministratorLoginPassword: pulumi.String("thisIsDog11"),
// 			LicenseType:                pulumi.String("BasePrice"),
// 			SubnetId:                   pulumi.Any(azurerm_subnet.Example.Id),
// 			SkuName:                    pulumi.String("GP_Gen5"),
// 			Vcores:                     pulumi.Int(4),
// 			StorageSizeInGb:            pulumi.Int(32),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			azurerm_subnet_network_security_group_association.Example,
// 			azurerm_subnet_route_table_association.Example,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewManagedInstanceActiveDirectoryAdministrator(ctx, "exampleManagedInstanceActiveDirectoryAdministrator", &sql.ManagedInstanceActiveDirectoryAdministratorArgs{
// 			ManagedInstanceName: exampleManagedInstance.Name,
// 			ResourceGroupName:   pulumi.Any(azurerm_resource_group.Example.Name),
// 			Login:               pulumi.String("sqladmin"),
// 			TenantId:            pulumi.String(current.TenantId),
// 			ObjectId:            pulumi.String(current.ObjectId),
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
// A SQL Active Directory Administrator can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/managedInstanceActiveDirectoryAdministrator:ManagedInstanceActiveDirectoryAdministrator administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/mymanagedinstance/administrators/activeDirectory
// ```
type ManagedInstanceActiveDirectoryAdministrator struct {
	pulumi.CustomResourceState

	// Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`). Defaults to `false`.
	AzureadAuthenticationOnly pulumi.BoolPtrOutput `pulumi:"azureadAuthenticationOnly"`
	// The login name of the principal to set as the Managed Instance administrator
	Login pulumi.StringOutput `pulumi:"login"`
	// The name of the SQL Managed Instance on which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceName pulumi.StringOutput `pulumi:"managedInstanceName"`
	// The ID of the principal to set as the Managed Instance administrator
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The name of the resource group for the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Azure Tenant ID
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewManagedInstanceActiveDirectoryAdministrator registers a new resource with the given unique name, arguments, and options.
func NewManagedInstanceActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, args *ManagedInstanceActiveDirectoryAdministratorArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceActiveDirectoryAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource ManagedInstanceActiveDirectoryAdministrator
	err := ctx.RegisterResource("azure:sql/managedInstanceActiveDirectoryAdministrator:ManagedInstanceActiveDirectoryAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstanceActiveDirectoryAdministrator gets an existing ManagedInstanceActiveDirectoryAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstanceActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceActiveDirectoryAdministratorState, opts ...pulumi.ResourceOption) (*ManagedInstanceActiveDirectoryAdministrator, error) {
	var resource ManagedInstanceActiveDirectoryAdministrator
	err := ctx.ReadResource("azure:sql/managedInstanceActiveDirectoryAdministrator:ManagedInstanceActiveDirectoryAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstanceActiveDirectoryAdministrator resources.
type managedInstanceActiveDirectoryAdministratorState struct {
	// Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`). Defaults to `false`.
	AzureadAuthenticationOnly *bool `pulumi:"azureadAuthenticationOnly"`
	// The login name of the principal to set as the Managed Instance administrator
	Login *string `pulumi:"login"`
	// The name of the SQL Managed Instance on which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceName *string `pulumi:"managedInstanceName"`
	// The ID of the principal to set as the Managed Instance administrator
	ObjectId *string `pulumi:"objectId"`
	// The name of the resource group for the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Azure Tenant ID
	TenantId *string `pulumi:"tenantId"`
}

type ManagedInstanceActiveDirectoryAdministratorState struct {
	// Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`). Defaults to `false`.
	AzureadAuthenticationOnly pulumi.BoolPtrInput
	// The login name of the principal to set as the Managed Instance administrator
	Login pulumi.StringPtrInput
	// The name of the SQL Managed Instance on which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceName pulumi.StringPtrInput
	// The ID of the principal to set as the Managed Instance administrator
	ObjectId pulumi.StringPtrInput
	// The name of the resource group for the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Azure Tenant ID
	TenantId pulumi.StringPtrInput
}

func (ManagedInstanceActiveDirectoryAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceActiveDirectoryAdministratorState)(nil)).Elem()
}

type managedInstanceActiveDirectoryAdministratorArgs struct {
	// Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`). Defaults to `false`.
	AzureadAuthenticationOnly *bool `pulumi:"azureadAuthenticationOnly"`
	// The login name of the principal to set as the Managed Instance administrator
	Login string `pulumi:"login"`
	// The name of the SQL Managed Instance on which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The ID of the principal to set as the Managed Instance administrator
	ObjectId string `pulumi:"objectId"`
	// The name of the resource group for the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Azure Tenant ID
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ManagedInstanceActiveDirectoryAdministrator resource.
type ManagedInstanceActiveDirectoryAdministratorArgs struct {
	// Specifies whether only AD Users and administrators can be used to login (`true`) or also local database users (`false`). Defaults to `false`.
	AzureadAuthenticationOnly pulumi.BoolPtrInput
	// The login name of the principal to set as the Managed Instance administrator
	Login pulumi.StringInput
	// The name of the SQL Managed Instance on which to set the administrator. Changing this forces a new resource to be created.
	ManagedInstanceName pulumi.StringInput
	// The ID of the principal to set as the Managed Instance administrator
	ObjectId pulumi.StringInput
	// The name of the resource group for the SQL Managed Instance. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Azure Tenant ID
	TenantId pulumi.StringInput
}

func (ManagedInstanceActiveDirectoryAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceActiveDirectoryAdministratorArgs)(nil)).Elem()
}

type ManagedInstanceActiveDirectoryAdministratorInput interface {
	pulumi.Input

	ToManagedInstanceActiveDirectoryAdministratorOutput() ManagedInstanceActiveDirectoryAdministratorOutput
	ToManagedInstanceActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorOutput
}

func (*ManagedInstanceActiveDirectoryAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (i *ManagedInstanceActiveDirectoryAdministrator) ToManagedInstanceActiveDirectoryAdministratorOutput() ManagedInstanceActiveDirectoryAdministratorOutput {
	return i.ToManagedInstanceActiveDirectoryAdministratorOutputWithContext(context.Background())
}

func (i *ManagedInstanceActiveDirectoryAdministrator) ToManagedInstanceActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceActiveDirectoryAdministratorOutput)
}

// ManagedInstanceActiveDirectoryAdministratorArrayInput is an input type that accepts ManagedInstanceActiveDirectoryAdministratorArray and ManagedInstanceActiveDirectoryAdministratorArrayOutput values.
// You can construct a concrete instance of `ManagedInstanceActiveDirectoryAdministratorArrayInput` via:
//
//          ManagedInstanceActiveDirectoryAdministratorArray{ ManagedInstanceActiveDirectoryAdministratorArgs{...} }
type ManagedInstanceActiveDirectoryAdministratorArrayInput interface {
	pulumi.Input

	ToManagedInstanceActiveDirectoryAdministratorArrayOutput() ManagedInstanceActiveDirectoryAdministratorArrayOutput
	ToManagedInstanceActiveDirectoryAdministratorArrayOutputWithContext(context.Context) ManagedInstanceActiveDirectoryAdministratorArrayOutput
}

type ManagedInstanceActiveDirectoryAdministratorArray []ManagedInstanceActiveDirectoryAdministratorInput

func (ManagedInstanceActiveDirectoryAdministratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (i ManagedInstanceActiveDirectoryAdministratorArray) ToManagedInstanceActiveDirectoryAdministratorArrayOutput() ManagedInstanceActiveDirectoryAdministratorArrayOutput {
	return i.ToManagedInstanceActiveDirectoryAdministratorArrayOutputWithContext(context.Background())
}

func (i ManagedInstanceActiveDirectoryAdministratorArray) ToManagedInstanceActiveDirectoryAdministratorArrayOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceActiveDirectoryAdministratorArrayOutput)
}

// ManagedInstanceActiveDirectoryAdministratorMapInput is an input type that accepts ManagedInstanceActiveDirectoryAdministratorMap and ManagedInstanceActiveDirectoryAdministratorMapOutput values.
// You can construct a concrete instance of `ManagedInstanceActiveDirectoryAdministratorMapInput` via:
//
//          ManagedInstanceActiveDirectoryAdministratorMap{ "key": ManagedInstanceActiveDirectoryAdministratorArgs{...} }
type ManagedInstanceActiveDirectoryAdministratorMapInput interface {
	pulumi.Input

	ToManagedInstanceActiveDirectoryAdministratorMapOutput() ManagedInstanceActiveDirectoryAdministratorMapOutput
	ToManagedInstanceActiveDirectoryAdministratorMapOutputWithContext(context.Context) ManagedInstanceActiveDirectoryAdministratorMapOutput
}

type ManagedInstanceActiveDirectoryAdministratorMap map[string]ManagedInstanceActiveDirectoryAdministratorInput

func (ManagedInstanceActiveDirectoryAdministratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (i ManagedInstanceActiveDirectoryAdministratorMap) ToManagedInstanceActiveDirectoryAdministratorMapOutput() ManagedInstanceActiveDirectoryAdministratorMapOutput {
	return i.ToManagedInstanceActiveDirectoryAdministratorMapOutputWithContext(context.Background())
}

func (i ManagedInstanceActiveDirectoryAdministratorMap) ToManagedInstanceActiveDirectoryAdministratorMapOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceActiveDirectoryAdministratorMapOutput)
}

type ManagedInstanceActiveDirectoryAdministratorOutput struct{ *pulumi.OutputState }

func (ManagedInstanceActiveDirectoryAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (o ManagedInstanceActiveDirectoryAdministratorOutput) ToManagedInstanceActiveDirectoryAdministratorOutput() ManagedInstanceActiveDirectoryAdministratorOutput {
	return o
}

func (o ManagedInstanceActiveDirectoryAdministratorOutput) ToManagedInstanceActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorOutput {
	return o
}

type ManagedInstanceActiveDirectoryAdministratorArrayOutput struct{ *pulumi.OutputState }

func (ManagedInstanceActiveDirectoryAdministratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (o ManagedInstanceActiveDirectoryAdministratorArrayOutput) ToManagedInstanceActiveDirectoryAdministratorArrayOutput() ManagedInstanceActiveDirectoryAdministratorArrayOutput {
	return o
}

func (o ManagedInstanceActiveDirectoryAdministratorArrayOutput) ToManagedInstanceActiveDirectoryAdministratorArrayOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorArrayOutput {
	return o
}

func (o ManagedInstanceActiveDirectoryAdministratorArrayOutput) Index(i pulumi.IntInput) ManagedInstanceActiveDirectoryAdministratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedInstanceActiveDirectoryAdministrator {
		return vs[0].([]*ManagedInstanceActiveDirectoryAdministrator)[vs[1].(int)]
	}).(ManagedInstanceActiveDirectoryAdministratorOutput)
}

type ManagedInstanceActiveDirectoryAdministratorMapOutput struct{ *pulumi.OutputState }

func (ManagedInstanceActiveDirectoryAdministratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedInstanceActiveDirectoryAdministrator)(nil)).Elem()
}

func (o ManagedInstanceActiveDirectoryAdministratorMapOutput) ToManagedInstanceActiveDirectoryAdministratorMapOutput() ManagedInstanceActiveDirectoryAdministratorMapOutput {
	return o
}

func (o ManagedInstanceActiveDirectoryAdministratorMapOutput) ToManagedInstanceActiveDirectoryAdministratorMapOutputWithContext(ctx context.Context) ManagedInstanceActiveDirectoryAdministratorMapOutput {
	return o
}

func (o ManagedInstanceActiveDirectoryAdministratorMapOutput) MapIndex(k pulumi.StringInput) ManagedInstanceActiveDirectoryAdministratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedInstanceActiveDirectoryAdministrator {
		return vs[0].(map[string]*ManagedInstanceActiveDirectoryAdministrator)[vs[1].(string)]
	}).(ManagedInstanceActiveDirectoryAdministratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceActiveDirectoryAdministratorInput)(nil)).Elem(), &ManagedInstanceActiveDirectoryAdministrator{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceActiveDirectoryAdministratorArrayInput)(nil)).Elem(), ManagedInstanceActiveDirectoryAdministratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceActiveDirectoryAdministratorMapInput)(nil)).Elem(), ManagedInstanceActiveDirectoryAdministratorMap{})
	pulumi.RegisterOutputType(ManagedInstanceActiveDirectoryAdministratorOutput{})
	pulumi.RegisterOutputType(ManagedInstanceActiveDirectoryAdministratorArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstanceActiveDirectoryAdministratorMapOutput{})
}
