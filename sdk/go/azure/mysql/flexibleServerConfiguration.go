// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets a MySQL Flexible Server Configuration value on a MySQL Flexible Server.
//
// ## Disclaimers
//
// > **Note:** Since this resource is provisioned by default, the Azure Provider will not check for the presence of an existing resource prior to attempting to create it.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mysql"
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
// 		_, err = mysql.NewFlexibleServer(ctx, "exampleFlexibleServer", &mysql.FlexibleServerArgs{
// 			ResourceGroupName:     pulumi.Any(azurerm_resource_group.Test.Name),
// 			Location:              pulumi.Any(azurerm_resource_group.Test.Location),
// 			AdministratorLogin:    pulumi.String("adminTerraform"),
// 			AdministratorPassword: pulumi.String("H@Sh1CoR3!"),
// 			SkuName:               pulumi.String("GP_Standard_D2ds_v4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewFlexibleServerConfiguration(ctx, "exampleFlexibleServerConfiguration", &mysql.FlexibleServerConfigurationArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        pulumi.Any(azurerm_mysql_server.Example.Name),
// 			Value:             pulumi.String("600"),
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
// MySQL Flexible Server Configurations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mysql/flexibleServerConfiguration:FlexibleServerConfiguration interactive_timeout /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/flexibleServers/flexibleServer1/configurations/interactive_timeout
// ```
type FlexibleServerConfiguration struct {
	pulumi.CustomResourceState

	// Specifies the name of the MySQL Flexible Server Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewFlexibleServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewFlexibleServerConfiguration(ctx *pulumi.Context,
	name string, args *FlexibleServerConfigurationArgs, opts ...pulumi.ResourceOption) (*FlexibleServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource FlexibleServerConfiguration
	err := ctx.RegisterResource("azure:mysql/flexibleServerConfiguration:FlexibleServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleServerConfiguration gets an existing FlexibleServerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleServerConfigurationState, opts ...pulumi.ResourceOption) (*FlexibleServerConfiguration, error) {
	var resource FlexibleServerConfiguration
	err := ctx.ReadResource("azure:mysql/flexibleServerConfiguration:FlexibleServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleServerConfiguration resources.
type flexibleServerConfigurationState struct {
	// Specifies the name of the MySQL Flexible Server Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value *string `pulumi:"value"`
}

type FlexibleServerConfigurationState struct {
	// Specifies the name of the MySQL Flexible Server Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringPtrInput
}

func (FlexibleServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerConfigurationState)(nil)).Elem()
}

type flexibleServerConfigurationArgs struct {
	// Specifies the name of the MySQL Flexible Server Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a FlexibleServerConfiguration resource.
type FlexibleServerConfigurationArgs struct {
	// Specifies the name of the MySQL Flexible Server Configuration, which needs [to be a valid MySQL configuration name](https://dev.mysql.com/doc/refman/5.7/en/server-configuration.html). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MySQL Flexible Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MySQL Flexible Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the value of the MySQL Flexible Server Configuration. See the MySQL documentation for valid values.
	Value pulumi.StringInput
}

func (FlexibleServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerConfigurationArgs)(nil)).Elem()
}

type FlexibleServerConfigurationInput interface {
	pulumi.Input

	ToFlexibleServerConfigurationOutput() FlexibleServerConfigurationOutput
	ToFlexibleServerConfigurationOutputWithContext(ctx context.Context) FlexibleServerConfigurationOutput
}

func (*FlexibleServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServerConfiguration)(nil))
}

func (i *FlexibleServerConfiguration) ToFlexibleServerConfigurationOutput() FlexibleServerConfigurationOutput {
	return i.ToFlexibleServerConfigurationOutputWithContext(context.Background())
}

func (i *FlexibleServerConfiguration) ToFlexibleServerConfigurationOutputWithContext(ctx context.Context) FlexibleServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerConfigurationOutput)
}

func (i *FlexibleServerConfiguration) ToFlexibleServerConfigurationPtrOutput() FlexibleServerConfigurationPtrOutput {
	return i.ToFlexibleServerConfigurationPtrOutputWithContext(context.Background())
}

func (i *FlexibleServerConfiguration) ToFlexibleServerConfigurationPtrOutputWithContext(ctx context.Context) FlexibleServerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerConfigurationPtrOutput)
}

type FlexibleServerConfigurationPtrInput interface {
	pulumi.Input

	ToFlexibleServerConfigurationPtrOutput() FlexibleServerConfigurationPtrOutput
	ToFlexibleServerConfigurationPtrOutputWithContext(ctx context.Context) FlexibleServerConfigurationPtrOutput
}

type flexibleServerConfigurationPtrType FlexibleServerConfigurationArgs

func (*flexibleServerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerConfiguration)(nil))
}

func (i *flexibleServerConfigurationPtrType) ToFlexibleServerConfigurationPtrOutput() FlexibleServerConfigurationPtrOutput {
	return i.ToFlexibleServerConfigurationPtrOutputWithContext(context.Background())
}

func (i *flexibleServerConfigurationPtrType) ToFlexibleServerConfigurationPtrOutputWithContext(ctx context.Context) FlexibleServerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerConfigurationPtrOutput)
}

// FlexibleServerConfigurationArrayInput is an input type that accepts FlexibleServerConfigurationArray and FlexibleServerConfigurationArrayOutput values.
// You can construct a concrete instance of `FlexibleServerConfigurationArrayInput` via:
//
//          FlexibleServerConfigurationArray{ FlexibleServerConfigurationArgs{...} }
type FlexibleServerConfigurationArrayInput interface {
	pulumi.Input

	ToFlexibleServerConfigurationArrayOutput() FlexibleServerConfigurationArrayOutput
	ToFlexibleServerConfigurationArrayOutputWithContext(context.Context) FlexibleServerConfigurationArrayOutput
}

type FlexibleServerConfigurationArray []FlexibleServerConfigurationInput

func (FlexibleServerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleServerConfiguration)(nil)).Elem()
}

func (i FlexibleServerConfigurationArray) ToFlexibleServerConfigurationArrayOutput() FlexibleServerConfigurationArrayOutput {
	return i.ToFlexibleServerConfigurationArrayOutputWithContext(context.Background())
}

func (i FlexibleServerConfigurationArray) ToFlexibleServerConfigurationArrayOutputWithContext(ctx context.Context) FlexibleServerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerConfigurationArrayOutput)
}

// FlexibleServerConfigurationMapInput is an input type that accepts FlexibleServerConfigurationMap and FlexibleServerConfigurationMapOutput values.
// You can construct a concrete instance of `FlexibleServerConfigurationMapInput` via:
//
//          FlexibleServerConfigurationMap{ "key": FlexibleServerConfigurationArgs{...} }
type FlexibleServerConfigurationMapInput interface {
	pulumi.Input

	ToFlexibleServerConfigurationMapOutput() FlexibleServerConfigurationMapOutput
	ToFlexibleServerConfigurationMapOutputWithContext(context.Context) FlexibleServerConfigurationMapOutput
}

type FlexibleServerConfigurationMap map[string]FlexibleServerConfigurationInput

func (FlexibleServerConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleServerConfiguration)(nil)).Elem()
}

func (i FlexibleServerConfigurationMap) ToFlexibleServerConfigurationMapOutput() FlexibleServerConfigurationMapOutput {
	return i.ToFlexibleServerConfigurationMapOutputWithContext(context.Background())
}

func (i FlexibleServerConfigurationMap) ToFlexibleServerConfigurationMapOutputWithContext(ctx context.Context) FlexibleServerConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerConfigurationMapOutput)
}

type FlexibleServerConfigurationOutput struct{ *pulumi.OutputState }

func (FlexibleServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServerConfiguration)(nil))
}

func (o FlexibleServerConfigurationOutput) ToFlexibleServerConfigurationOutput() FlexibleServerConfigurationOutput {
	return o
}

func (o FlexibleServerConfigurationOutput) ToFlexibleServerConfigurationOutputWithContext(ctx context.Context) FlexibleServerConfigurationOutput {
	return o
}

func (o FlexibleServerConfigurationOutput) ToFlexibleServerConfigurationPtrOutput() FlexibleServerConfigurationPtrOutput {
	return o.ToFlexibleServerConfigurationPtrOutputWithContext(context.Background())
}

func (o FlexibleServerConfigurationOutput) ToFlexibleServerConfigurationPtrOutputWithContext(ctx context.Context) FlexibleServerConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlexibleServerConfiguration) *FlexibleServerConfiguration {
		return &v
	}).(FlexibleServerConfigurationPtrOutput)
}

type FlexibleServerConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FlexibleServerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerConfiguration)(nil))
}

func (o FlexibleServerConfigurationPtrOutput) ToFlexibleServerConfigurationPtrOutput() FlexibleServerConfigurationPtrOutput {
	return o
}

func (o FlexibleServerConfigurationPtrOutput) ToFlexibleServerConfigurationPtrOutputWithContext(ctx context.Context) FlexibleServerConfigurationPtrOutput {
	return o
}

func (o FlexibleServerConfigurationPtrOutput) Elem() FlexibleServerConfigurationOutput {
	return o.ApplyT(func(v *FlexibleServerConfiguration) FlexibleServerConfiguration {
		if v != nil {
			return *v
		}
		var ret FlexibleServerConfiguration
		return ret
	}).(FlexibleServerConfigurationOutput)
}

type FlexibleServerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FlexibleServerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FlexibleServerConfiguration)(nil))
}

func (o FlexibleServerConfigurationArrayOutput) ToFlexibleServerConfigurationArrayOutput() FlexibleServerConfigurationArrayOutput {
	return o
}

func (o FlexibleServerConfigurationArrayOutput) ToFlexibleServerConfigurationArrayOutputWithContext(ctx context.Context) FlexibleServerConfigurationArrayOutput {
	return o
}

func (o FlexibleServerConfigurationArrayOutput) Index(i pulumi.IntInput) FlexibleServerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FlexibleServerConfiguration {
		return vs[0].([]FlexibleServerConfiguration)[vs[1].(int)]
	}).(FlexibleServerConfigurationOutput)
}

type FlexibleServerConfigurationMapOutput struct{ *pulumi.OutputState }

func (FlexibleServerConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlexibleServerConfiguration)(nil))
}

func (o FlexibleServerConfigurationMapOutput) ToFlexibleServerConfigurationMapOutput() FlexibleServerConfigurationMapOutput {
	return o
}

func (o FlexibleServerConfigurationMapOutput) ToFlexibleServerConfigurationMapOutputWithContext(ctx context.Context) FlexibleServerConfigurationMapOutput {
	return o
}

func (o FlexibleServerConfigurationMapOutput) MapIndex(k pulumi.StringInput) FlexibleServerConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlexibleServerConfiguration {
		return vs[0].(map[string]FlexibleServerConfiguration)[vs[1].(string)]
	}).(FlexibleServerConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(FlexibleServerConfigurationOutput{})
	pulumi.RegisterOutputType(FlexibleServerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FlexibleServerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FlexibleServerConfigurationMapOutput{})
}
