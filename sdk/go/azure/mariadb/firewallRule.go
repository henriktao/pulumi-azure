// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Firewall Rule for a MariaDB Server
//
// ## Example Usage
// ### Single IP Address)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mariadb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mariadb.NewFirewallRule(ctx, "example", &mariadb.FirewallRuleArgs{
// 			EndIpAddress:      pulumi.String("40.112.8.12"),
// 			ResourceGroupName: pulumi.String("test-rg"),
// 			ServerName:        pulumi.String("test-server"),
// 			StartIpAddress:    pulumi.String("40.112.8.12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### IP Range)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mariadb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mariadb.NewFirewallRule(ctx, "example", &mariadb.FirewallRuleArgs{
// 			EndIpAddress:      pulumi.String("40.112.255.255"),
// 			ResourceGroupName: pulumi.String("test-rg"),
// 			ServerName:        pulumi.String("test-server"),
// 			StartIpAddress:    pulumi.String("40.112.0.0"),
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
// MariaDB Firewall rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mariadb/firewallRule:FirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1/firewallRules/rule1
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	var resource FirewallRule
	err := ctx.RegisterResource("azure:mariadb/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure:mariadb/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress *string `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FirewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringPtrInput
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress pulumi.StringInput
	// Specifies the name of the MariaDB Firewall Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//          FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//          FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
