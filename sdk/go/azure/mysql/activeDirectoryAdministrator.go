// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows you to set a user or group as the AD administrator for an MySQL server in Azure
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mysql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServer, err := mysql.NewServer(ctx, "exampleServer", &mysql.ServerArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			AdministratorLogin:         pulumi.String("mysqladminun"),
// 			AdministratorLoginPassword: pulumi.String("H@Sh1CoR3!"),
// 			SkuName:                    pulumi.String("B_Gen5_2"),
// 			StorageMb:                  pulumi.Int(5120),
// 			Version:                    pulumi.String("5.7"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewActiveDirectoryAdministrator(ctx, "exampleActiveDirectoryAdministrator", &mysql.ActiveDirectoryAdministratorArgs{
// 			ServerName:        exampleServer.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Login:             pulumi.String("sqladmin"),
// 			TenantId:          pulumi.String(current.TenantId),
// 			ObjectId:          pulumi.String(current.ObjectId),
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
// A MySQL Active Directory Administrator can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator administrator /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
// ```
type ActiveDirectoryAdministrator struct {
	pulumi.CustomResourceState

	// The login name of the principal to set as the server administrator
	Login pulumi.StringOutput `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewActiveDirectoryAdministrator registers a new resource with the given unique name, arguments, and options.
func NewActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, args *ActiveDirectoryAdministratorArgs, opts ...pulumi.ResourceOption) (*ActiveDirectoryAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource ActiveDirectoryAdministrator
	err := ctx.RegisterResource("azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveDirectoryAdministrator gets an existing ActiveDirectoryAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveDirectoryAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveDirectoryAdministratorState, opts ...pulumi.ResourceOption) (*ActiveDirectoryAdministrator, error) {
	var resource ActiveDirectoryAdministrator
	err := ctx.ReadResource("azure:mysql/activeDirectoryAdministrator:ActiveDirectoryAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveDirectoryAdministrator resources.
type activeDirectoryAdministratorState struct {
	// The login name of the principal to set as the server administrator
	Login *string `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId *string `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId *string `pulumi:"tenantId"`
}

type ActiveDirectoryAdministratorState struct {
	// The login name of the principal to set as the server administrator
	Login pulumi.StringPtrInput
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringPtrInput
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// The Azure Tenant ID
	TenantId pulumi.StringPtrInput
}

func (ActiveDirectoryAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryAdministratorState)(nil)).Elem()
}

type activeDirectoryAdministratorArgs struct {
	// The login name of the principal to set as the server administrator
	Login string `pulumi:"login"`
	// The ID of the principal to set as the server administrator
	ObjectId string `pulumi:"objectId"`
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// The Azure Tenant ID
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ActiveDirectoryAdministrator resource.
type ActiveDirectoryAdministratorArgs struct {
	// The login name of the principal to set as the server administrator
	Login pulumi.StringInput
	// The ID of the principal to set as the server administrator
	ObjectId pulumi.StringInput
	// The name of the resource group for the MySQL server. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the MySQL Server on which to set the administrator. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// The Azure Tenant ID
	TenantId pulumi.StringInput
}

func (ActiveDirectoryAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryAdministratorArgs)(nil)).Elem()
}

type ActiveDirectoryAdministratorInput interface {
	pulumi.Input

	ToActiveDirectoryAdministratorOutput() ActiveDirectoryAdministratorOutput
	ToActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorOutput
}

func (*ActiveDirectoryAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryAdministrator)(nil))
}

func (i *ActiveDirectoryAdministrator) ToActiveDirectoryAdministratorOutput() ActiveDirectoryAdministratorOutput {
	return i.ToActiveDirectoryAdministratorOutputWithContext(context.Background())
}

func (i *ActiveDirectoryAdministrator) ToActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryAdministratorOutput)
}

func (i *ActiveDirectoryAdministrator) ToActiveDirectoryAdministratorPtrOutput() ActiveDirectoryAdministratorPtrOutput {
	return i.ToActiveDirectoryAdministratorPtrOutputWithContext(context.Background())
}

func (i *ActiveDirectoryAdministrator) ToActiveDirectoryAdministratorPtrOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryAdministratorPtrOutput)
}

type ActiveDirectoryAdministratorPtrInput interface {
	pulumi.Input

	ToActiveDirectoryAdministratorPtrOutput() ActiveDirectoryAdministratorPtrOutput
	ToActiveDirectoryAdministratorPtrOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorPtrOutput
}

type activeDirectoryAdministratorPtrType ActiveDirectoryAdministratorArgs

func (*activeDirectoryAdministratorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryAdministrator)(nil))
}

func (i *activeDirectoryAdministratorPtrType) ToActiveDirectoryAdministratorPtrOutput() ActiveDirectoryAdministratorPtrOutput {
	return i.ToActiveDirectoryAdministratorPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryAdministratorPtrType) ToActiveDirectoryAdministratorPtrOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryAdministratorPtrOutput)
}

// ActiveDirectoryAdministratorArrayInput is an input type that accepts ActiveDirectoryAdministratorArray and ActiveDirectoryAdministratorArrayOutput values.
// You can construct a concrete instance of `ActiveDirectoryAdministratorArrayInput` via:
//
//          ActiveDirectoryAdministratorArray{ ActiveDirectoryAdministratorArgs{...} }
type ActiveDirectoryAdministratorArrayInput interface {
	pulumi.Input

	ToActiveDirectoryAdministratorArrayOutput() ActiveDirectoryAdministratorArrayOutput
	ToActiveDirectoryAdministratorArrayOutputWithContext(context.Context) ActiveDirectoryAdministratorArrayOutput
}

type ActiveDirectoryAdministratorArray []ActiveDirectoryAdministratorInput

func (ActiveDirectoryAdministratorArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ActiveDirectoryAdministrator)(nil))
}

func (i ActiveDirectoryAdministratorArray) ToActiveDirectoryAdministratorArrayOutput() ActiveDirectoryAdministratorArrayOutput {
	return i.ToActiveDirectoryAdministratorArrayOutputWithContext(context.Background())
}

func (i ActiveDirectoryAdministratorArray) ToActiveDirectoryAdministratorArrayOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryAdministratorArrayOutput)
}

// ActiveDirectoryAdministratorMapInput is an input type that accepts ActiveDirectoryAdministratorMap and ActiveDirectoryAdministratorMapOutput values.
// You can construct a concrete instance of `ActiveDirectoryAdministratorMapInput` via:
//
//          ActiveDirectoryAdministratorMap{ "key": ActiveDirectoryAdministratorArgs{...} }
type ActiveDirectoryAdministratorMapInput interface {
	pulumi.Input

	ToActiveDirectoryAdministratorMapOutput() ActiveDirectoryAdministratorMapOutput
	ToActiveDirectoryAdministratorMapOutputWithContext(context.Context) ActiveDirectoryAdministratorMapOutput
}

type ActiveDirectoryAdministratorMap map[string]ActiveDirectoryAdministratorInput

func (ActiveDirectoryAdministratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ActiveDirectoryAdministrator)(nil))
}

func (i ActiveDirectoryAdministratorMap) ToActiveDirectoryAdministratorMapOutput() ActiveDirectoryAdministratorMapOutput {
	return i.ToActiveDirectoryAdministratorMapOutputWithContext(context.Background())
}

func (i ActiveDirectoryAdministratorMap) ToActiveDirectoryAdministratorMapOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryAdministratorMapOutput)
}

type ActiveDirectoryAdministratorOutput struct {
	*pulumi.OutputState
}

func (ActiveDirectoryAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryAdministrator)(nil))
}

func (o ActiveDirectoryAdministratorOutput) ToActiveDirectoryAdministratorOutput() ActiveDirectoryAdministratorOutput {
	return o
}

func (o ActiveDirectoryAdministratorOutput) ToActiveDirectoryAdministratorOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorOutput {
	return o
}

func (o ActiveDirectoryAdministratorOutput) ToActiveDirectoryAdministratorPtrOutput() ActiveDirectoryAdministratorPtrOutput {
	return o.ToActiveDirectoryAdministratorPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryAdministratorOutput) ToActiveDirectoryAdministratorPtrOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorPtrOutput {
	return o.ApplyT(func(v ActiveDirectoryAdministrator) *ActiveDirectoryAdministrator {
		return &v
	}).(ActiveDirectoryAdministratorPtrOutput)
}

type ActiveDirectoryAdministratorPtrOutput struct {
	*pulumi.OutputState
}

func (ActiveDirectoryAdministratorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryAdministrator)(nil))
}

func (o ActiveDirectoryAdministratorPtrOutput) ToActiveDirectoryAdministratorPtrOutput() ActiveDirectoryAdministratorPtrOutput {
	return o
}

func (o ActiveDirectoryAdministratorPtrOutput) ToActiveDirectoryAdministratorPtrOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorPtrOutput {
	return o
}

type ActiveDirectoryAdministratorArrayOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryAdministratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveDirectoryAdministrator)(nil))
}

func (o ActiveDirectoryAdministratorArrayOutput) ToActiveDirectoryAdministratorArrayOutput() ActiveDirectoryAdministratorArrayOutput {
	return o
}

func (o ActiveDirectoryAdministratorArrayOutput) ToActiveDirectoryAdministratorArrayOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorArrayOutput {
	return o
}

func (o ActiveDirectoryAdministratorArrayOutput) Index(i pulumi.IntInput) ActiveDirectoryAdministratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveDirectoryAdministrator {
		return vs[0].([]ActiveDirectoryAdministrator)[vs[1].(int)]
	}).(ActiveDirectoryAdministratorOutput)
}

type ActiveDirectoryAdministratorMapOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryAdministratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActiveDirectoryAdministrator)(nil))
}

func (o ActiveDirectoryAdministratorMapOutput) ToActiveDirectoryAdministratorMapOutput() ActiveDirectoryAdministratorMapOutput {
	return o
}

func (o ActiveDirectoryAdministratorMapOutput) ToActiveDirectoryAdministratorMapOutputWithContext(ctx context.Context) ActiveDirectoryAdministratorMapOutput {
	return o
}

func (o ActiveDirectoryAdministratorMapOutput) MapIndex(k pulumi.StringInput) ActiveDirectoryAdministratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActiveDirectoryAdministrator {
		return vs[0].(map[string]ActiveDirectoryAdministrator)[vs[1].(string)]
	}).(ActiveDirectoryAdministratorOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryAdministratorOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryAdministratorPtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryAdministratorArrayOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryAdministratorMapOutput{})
}
