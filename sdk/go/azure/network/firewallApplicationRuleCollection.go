// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Application Rule Collection within an Azure Firewall.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
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
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFirewall, err := network.NewFirewall(ctx, "exampleFirewall", &network.FirewallArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.FirewallIpConfigurationArray{
// 				&network.FirewallIpConfigurationArgs{
// 					Name:              pulumi.String("configuration"),
// 					SubnetId:          exampleSubnet.ID(),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewFirewallApplicationRuleCollection(ctx, "exampleFirewallApplicationRuleCollection", &network.FirewallApplicationRuleCollectionArgs{
// 			AzureFirewallName: exampleFirewall.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Priority:          pulumi.Int(100),
// 			Action:            pulumi.String("Allow"),
// 			Rules: network.FirewallApplicationRuleCollectionRuleArray{
// 				&network.FirewallApplicationRuleCollectionRuleArgs{
// 					Name: pulumi.String("testrule"),
// 					SourceAddresses: pulumi.StringArray{
// 						pulumi.String("10.0.0.0/16"),
// 					},
// 					TargetFqdns: pulumi.StringArray{
// 						pulumi.String("*.google.com"),
// 					},
// 					Protocols: network.FirewallApplicationRuleCollectionRuleProtocolArray{
// 						&network.FirewallApplicationRuleCollectionRuleProtocolArgs{
// 							Port: pulumi.Int(443),
// 							Type: pulumi.String("Https"),
// 						},
// 					},
// 				},
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
// Firewall Application Rule Collections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/applicationRuleCollections/mycollection
// ```
type FirewallApplicationRuleCollection struct {
	pulumi.CustomResourceState

	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringOutput `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayOutput `pulumi:"rules"`
}

// NewFirewallApplicationRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, args *FirewallApplicationRuleCollectionArgs, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.AzureFirewallName == nil {
		return nil, errors.New("invalid value for required argument 'AzureFirewallName'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource FirewallApplicationRuleCollection
	err := ctx.RegisterResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallApplicationRuleCollection gets an existing FirewallApplicationRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallApplicationRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallApplicationRuleCollectionState, opts ...pulumi.ResourceOption) (*FirewallApplicationRuleCollection, error) {
	var resource FirewallApplicationRuleCollection
	err := ctx.ReadResource("azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallApplicationRuleCollection resources.
type firewallApplicationRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action *string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName *string `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority *int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallApplicationRuleCollectionRule `pulumi:"rules"`
}

type FirewallApplicationRuleCollectionState struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringPtrInput
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringPtrInput
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntPtrInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayInput
}

func (FirewallApplicationRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallApplicationRuleCollectionState)(nil)).Elem()
}

type firewallApplicationRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action string `pulumi:"action"`
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName string `pulumi:"azureFirewallName"`
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority int `pulumi:"priority"`
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// One or more `rule` blocks as defined below.
	Rules []FirewallApplicationRuleCollectionRule `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallApplicationRuleCollection resource.
type FirewallApplicationRuleCollectionArgs struct {
	// Specifies the action the rule will apply to matching traffic. Possible values are `Allow` and `Deny`.
	Action pulumi.StringInput
	// Specifies the name of the Firewall in which the Application Rule Collection should be created. Changing this forces a new resource to be created.
	AzureFirewallName pulumi.StringInput
	// Specifies the name of the Application Rule Collection which must be unique within the Firewall. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the priority of the rule collection. Possible values are between `100` - `65000`.
	Priority pulumi.IntInput
	// Specifies the name of the Resource Group in which the Firewall exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// One or more `rule` blocks as defined below.
	Rules FirewallApplicationRuleCollectionRuleArrayInput
}

func (FirewallApplicationRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallApplicationRuleCollectionArgs)(nil)).Elem()
}

type FirewallApplicationRuleCollectionInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionOutput() FirewallApplicationRuleCollectionOutput
	ToFirewallApplicationRuleCollectionOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionOutput
}

func (*FirewallApplicationRuleCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallApplicationRuleCollection)(nil)).Elem()
}

func (i *FirewallApplicationRuleCollection) ToFirewallApplicationRuleCollectionOutput() FirewallApplicationRuleCollectionOutput {
	return i.ToFirewallApplicationRuleCollectionOutputWithContext(context.Background())
}

func (i *FirewallApplicationRuleCollection) ToFirewallApplicationRuleCollectionOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallApplicationRuleCollectionOutput)
}

// FirewallApplicationRuleCollectionArrayInput is an input type that accepts FirewallApplicationRuleCollectionArray and FirewallApplicationRuleCollectionArrayOutput values.
// You can construct a concrete instance of `FirewallApplicationRuleCollectionArrayInput` via:
//
//          FirewallApplicationRuleCollectionArray{ FirewallApplicationRuleCollectionArgs{...} }
type FirewallApplicationRuleCollectionArrayInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionArrayOutput() FirewallApplicationRuleCollectionArrayOutput
	ToFirewallApplicationRuleCollectionArrayOutputWithContext(context.Context) FirewallApplicationRuleCollectionArrayOutput
}

type FirewallApplicationRuleCollectionArray []FirewallApplicationRuleCollectionInput

func (FirewallApplicationRuleCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallApplicationRuleCollection)(nil)).Elem()
}

func (i FirewallApplicationRuleCollectionArray) ToFirewallApplicationRuleCollectionArrayOutput() FirewallApplicationRuleCollectionArrayOutput {
	return i.ToFirewallApplicationRuleCollectionArrayOutputWithContext(context.Background())
}

func (i FirewallApplicationRuleCollectionArray) ToFirewallApplicationRuleCollectionArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallApplicationRuleCollectionArrayOutput)
}

// FirewallApplicationRuleCollectionMapInput is an input type that accepts FirewallApplicationRuleCollectionMap and FirewallApplicationRuleCollectionMapOutput values.
// You can construct a concrete instance of `FirewallApplicationRuleCollectionMapInput` via:
//
//          FirewallApplicationRuleCollectionMap{ "key": FirewallApplicationRuleCollectionArgs{...} }
type FirewallApplicationRuleCollectionMapInput interface {
	pulumi.Input

	ToFirewallApplicationRuleCollectionMapOutput() FirewallApplicationRuleCollectionMapOutput
	ToFirewallApplicationRuleCollectionMapOutputWithContext(context.Context) FirewallApplicationRuleCollectionMapOutput
}

type FirewallApplicationRuleCollectionMap map[string]FirewallApplicationRuleCollectionInput

func (FirewallApplicationRuleCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallApplicationRuleCollection)(nil)).Elem()
}

func (i FirewallApplicationRuleCollectionMap) ToFirewallApplicationRuleCollectionMapOutput() FirewallApplicationRuleCollectionMapOutput {
	return i.ToFirewallApplicationRuleCollectionMapOutputWithContext(context.Background())
}

func (i FirewallApplicationRuleCollectionMap) ToFirewallApplicationRuleCollectionMapOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallApplicationRuleCollectionMapOutput)
}

type FirewallApplicationRuleCollectionOutput struct{ *pulumi.OutputState }

func (FirewallApplicationRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallApplicationRuleCollection)(nil)).Elem()
}

func (o FirewallApplicationRuleCollectionOutput) ToFirewallApplicationRuleCollectionOutput() FirewallApplicationRuleCollectionOutput {
	return o
}

func (o FirewallApplicationRuleCollectionOutput) ToFirewallApplicationRuleCollectionOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionOutput {
	return o
}

type FirewallApplicationRuleCollectionArrayOutput struct{ *pulumi.OutputState }

func (FirewallApplicationRuleCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallApplicationRuleCollection)(nil)).Elem()
}

func (o FirewallApplicationRuleCollectionArrayOutput) ToFirewallApplicationRuleCollectionArrayOutput() FirewallApplicationRuleCollectionArrayOutput {
	return o
}

func (o FirewallApplicationRuleCollectionArrayOutput) ToFirewallApplicationRuleCollectionArrayOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionArrayOutput {
	return o
}

func (o FirewallApplicationRuleCollectionArrayOutput) Index(i pulumi.IntInput) FirewallApplicationRuleCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallApplicationRuleCollection {
		return vs[0].([]*FirewallApplicationRuleCollection)[vs[1].(int)]
	}).(FirewallApplicationRuleCollectionOutput)
}

type FirewallApplicationRuleCollectionMapOutput struct{ *pulumi.OutputState }

func (FirewallApplicationRuleCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallApplicationRuleCollection)(nil)).Elem()
}

func (o FirewallApplicationRuleCollectionMapOutput) ToFirewallApplicationRuleCollectionMapOutput() FirewallApplicationRuleCollectionMapOutput {
	return o
}

func (o FirewallApplicationRuleCollectionMapOutput) ToFirewallApplicationRuleCollectionMapOutputWithContext(ctx context.Context) FirewallApplicationRuleCollectionMapOutput {
	return o
}

func (o FirewallApplicationRuleCollectionMapOutput) MapIndex(k pulumi.StringInput) FirewallApplicationRuleCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallApplicationRuleCollection {
		return vs[0].(map[string]*FirewallApplicationRuleCollection)[vs[1].(string)]
	}).(FirewallApplicationRuleCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallApplicationRuleCollectionInput)(nil)).Elem(), &FirewallApplicationRuleCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallApplicationRuleCollectionArrayInput)(nil)).Elem(), FirewallApplicationRuleCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallApplicationRuleCollectionMapInput)(nil)).Elem(), FirewallApplicationRuleCollectionMap{})
	pulumi.RegisterOutputType(FirewallApplicationRuleCollectionOutput{})
	pulumi.RegisterOutputType(FirewallApplicationRuleCollectionArrayOutput{})
	pulumi.RegisterOutputType(FirewallApplicationRuleCollectionMapOutput{})
}
