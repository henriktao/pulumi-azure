// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a PostgreSQL Flexible Server Firewall Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/postgresql"
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
// 		exampleFlexibleServer, err := postgresql.NewFlexibleServer(ctx, "exampleFlexibleServer", &postgresql.FlexibleServerArgs{
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			Location:              exampleResourceGroup.Location,
// 			Version:               pulumi.String("12"),
// 			AdministratorLogin:    pulumi.String("psqladmin"),
// 			AdministratorPassword: pulumi.String("H@Sh1CoR3!"),
// 			StorageMb:             pulumi.Int(32768),
// 			SkuName:               pulumi.String("GP_Standard_D4s_v3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = postgresql.NewFlexibleServerFirewallRule(ctx, "exampleFlexibleServerFirewallRule", &postgresql.FlexibleServerFirewallRuleArgs{
// 			ServerId:       exampleFlexibleServer.ID(),
// 			StartIpAddress: pulumi.String("122.122.0.0"),
// 			EndIpAddress:   pulumi.String("122.122.0.0"),
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
// PostgreSQL Flexible Server Firewall Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:postgresql/flexibleServerFirewallRule:FlexibleServerFirewallRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
// ```
type FlexibleServerFirewallRule struct {
	pulumi.CustomResourceState

	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The name which should be used for this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewFlexibleServerFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFlexibleServerFirewallRule(ctx *pulumi.Context,
	name string, args *FlexibleServerFirewallRuleArgs, opts ...pulumi.ResourceOption) (*FlexibleServerFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource FlexibleServerFirewallRule
	err := ctx.RegisterResource("azure:postgresql/flexibleServerFirewallRule:FlexibleServerFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleServerFirewallRule gets an existing FlexibleServerFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleServerFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleServerFirewallRuleState, opts ...pulumi.ResourceOption) (*FlexibleServerFirewallRule, error) {
	var resource FlexibleServerFirewallRule
	err := ctx.ReadResource("azure:postgresql/flexibleServerFirewallRule:FlexibleServerFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleServerFirewallRule resources.
type flexibleServerFirewallRuleState struct {
	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// The name which should be used for this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	Name *string `pulumi:"name"`
	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerId *string `pulumi:"serverId"`
	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FlexibleServerFirewallRuleState struct {
	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIpAddress pulumi.StringPtrInput
	// The name which should be used for this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	Name pulumi.StringPtrInput
	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerId pulumi.StringPtrInput
	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIpAddress pulumi.StringPtrInput
}

func (FlexibleServerFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerFirewallRuleState)(nil)).Elem()
}

type flexibleServerFirewallRuleArgs struct {
	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The name which should be used for this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	Name *string `pulumi:"name"`
	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerId string `pulumi:"serverId"`
	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FlexibleServerFirewallRule resource.
type FlexibleServerFirewallRuleArgs struct {
	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIpAddress pulumi.StringInput
	// The name which should be used for this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	Name pulumi.StringPtrInput
	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerId pulumi.StringInput
	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIpAddress pulumi.StringInput
}

func (FlexibleServerFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerFirewallRuleArgs)(nil)).Elem()
}

type FlexibleServerFirewallRuleInput interface {
	pulumi.Input

	ToFlexibleServerFirewallRuleOutput() FlexibleServerFirewallRuleOutput
	ToFlexibleServerFirewallRuleOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleOutput
}

func (*FlexibleServerFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerFirewallRule)(nil)).Elem()
}

func (i *FlexibleServerFirewallRule) ToFlexibleServerFirewallRuleOutput() FlexibleServerFirewallRuleOutput {
	return i.ToFlexibleServerFirewallRuleOutputWithContext(context.Background())
}

func (i *FlexibleServerFirewallRule) ToFlexibleServerFirewallRuleOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerFirewallRuleOutput)
}

// FlexibleServerFirewallRuleArrayInput is an input type that accepts FlexibleServerFirewallRuleArray and FlexibleServerFirewallRuleArrayOutput values.
// You can construct a concrete instance of `FlexibleServerFirewallRuleArrayInput` via:
//
//          FlexibleServerFirewallRuleArray{ FlexibleServerFirewallRuleArgs{...} }
type FlexibleServerFirewallRuleArrayInput interface {
	pulumi.Input

	ToFlexibleServerFirewallRuleArrayOutput() FlexibleServerFirewallRuleArrayOutput
	ToFlexibleServerFirewallRuleArrayOutputWithContext(context.Context) FlexibleServerFirewallRuleArrayOutput
}

type FlexibleServerFirewallRuleArray []FlexibleServerFirewallRuleInput

func (FlexibleServerFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleServerFirewallRule)(nil)).Elem()
}

func (i FlexibleServerFirewallRuleArray) ToFlexibleServerFirewallRuleArrayOutput() FlexibleServerFirewallRuleArrayOutput {
	return i.ToFlexibleServerFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FlexibleServerFirewallRuleArray) ToFlexibleServerFirewallRuleArrayOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerFirewallRuleArrayOutput)
}

// FlexibleServerFirewallRuleMapInput is an input type that accepts FlexibleServerFirewallRuleMap and FlexibleServerFirewallRuleMapOutput values.
// You can construct a concrete instance of `FlexibleServerFirewallRuleMapInput` via:
//
//          FlexibleServerFirewallRuleMap{ "key": FlexibleServerFirewallRuleArgs{...} }
type FlexibleServerFirewallRuleMapInput interface {
	pulumi.Input

	ToFlexibleServerFirewallRuleMapOutput() FlexibleServerFirewallRuleMapOutput
	ToFlexibleServerFirewallRuleMapOutputWithContext(context.Context) FlexibleServerFirewallRuleMapOutput
}

type FlexibleServerFirewallRuleMap map[string]FlexibleServerFirewallRuleInput

func (FlexibleServerFirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleServerFirewallRule)(nil)).Elem()
}

func (i FlexibleServerFirewallRuleMap) ToFlexibleServerFirewallRuleMapOutput() FlexibleServerFirewallRuleMapOutput {
	return i.ToFlexibleServerFirewallRuleMapOutputWithContext(context.Background())
}

func (i FlexibleServerFirewallRuleMap) ToFlexibleServerFirewallRuleMapOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerFirewallRuleMapOutput)
}

type FlexibleServerFirewallRuleOutput struct{ *pulumi.OutputState }

func (FlexibleServerFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerFirewallRule)(nil)).Elem()
}

func (o FlexibleServerFirewallRuleOutput) ToFlexibleServerFirewallRuleOutput() FlexibleServerFirewallRuleOutput {
	return o
}

func (o FlexibleServerFirewallRuleOutput) ToFlexibleServerFirewallRuleOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleOutput {
	return o
}

type FlexibleServerFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FlexibleServerFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleServerFirewallRule)(nil)).Elem()
}

func (o FlexibleServerFirewallRuleArrayOutput) ToFlexibleServerFirewallRuleArrayOutput() FlexibleServerFirewallRuleArrayOutput {
	return o
}

func (o FlexibleServerFirewallRuleArrayOutput) ToFlexibleServerFirewallRuleArrayOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleArrayOutput {
	return o
}

func (o FlexibleServerFirewallRuleArrayOutput) Index(i pulumi.IntInput) FlexibleServerFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlexibleServerFirewallRule {
		return vs[0].([]*FlexibleServerFirewallRule)[vs[1].(int)]
	}).(FlexibleServerFirewallRuleOutput)
}

type FlexibleServerFirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FlexibleServerFirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleServerFirewallRule)(nil)).Elem()
}

func (o FlexibleServerFirewallRuleMapOutput) ToFlexibleServerFirewallRuleMapOutput() FlexibleServerFirewallRuleMapOutput {
	return o
}

func (o FlexibleServerFirewallRuleMapOutput) ToFlexibleServerFirewallRuleMapOutputWithContext(ctx context.Context) FlexibleServerFirewallRuleMapOutput {
	return o
}

func (o FlexibleServerFirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FlexibleServerFirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlexibleServerFirewallRule {
		return vs[0].(map[string]*FlexibleServerFirewallRule)[vs[1].(string)]
	}).(FlexibleServerFirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleServerFirewallRuleInput)(nil)).Elem(), &FlexibleServerFirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleServerFirewallRuleArrayInput)(nil)).Elem(), FlexibleServerFirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleServerFirewallRuleMapInput)(nil)).Elem(), FlexibleServerFirewallRuleMap{})
	pulumi.RegisterOutputType(FlexibleServerFirewallRuleOutput{})
	pulumi.RegisterOutputType(FlexibleServerFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FlexibleServerFirewallRuleMapOutput{})
}
