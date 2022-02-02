// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Connection for a Virtual Hub.
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
// 				pulumi.String("172.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualWan, err := network.NewVirtualWan(ctx, "exampleVirtualWan", &network.VirtualWanArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualHub, err := network.NewVirtualHub(ctx, "exampleVirtualHub", &network.VirtualHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			VirtualWanId:      exampleVirtualWan.ID(),
// 			AddressPrefix:     pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewVirtualHubConnection(ctx, "exampleVirtualHubConnection", &network.VirtualHubConnectionArgs{
// 			VirtualHubId:           exampleVirtualHub.ID(),
// 			RemoteVirtualNetworkId: exampleVirtualNetwork.ID(),
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
// Virtual Hub Connection's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/virtualHubConnection:VirtualHubConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/connection1
// ```
type VirtualHubConnection struct {
	pulumi.CustomResourceState

	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrOutput `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrOutput `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringOutput `pulumi:"remoteVirtualNetworkId"`
	// A `routing` block as defined below.
	Routing VirtualHubConnectionRoutingOutput `pulumi:"routing"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringOutput `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrOutput `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

// NewVirtualHubConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubConnection(ctx *pulumi.Context,
	name string, args *VirtualHubConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualHubConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RemoteVirtualNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'RemoteVirtualNetworkId'")
	}
	if args.VirtualHubId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubId'")
	}
	var resource VirtualHubConnection
	err := ctx.RegisterResource("azure:network/virtualHubConnection:VirtualHubConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubConnection gets an existing VirtualHubConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubConnectionState, opts ...pulumi.ResourceOption) (*VirtualHubConnection, error) {
	var resource VirtualHubConnection
	err := ctx.ReadResource("azure:network/virtualHubConnection:VirtualHubConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubConnection resources.
type virtualHubConnectionState struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed *bool `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled *bool `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId *string `pulumi:"remoteVirtualNetworkId"`
	// A `routing` block as defined below.
	Routing *VirtualHubConnectionRouting `pulumi:"routing"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId *string `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed *bool `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

type VirtualHubConnectionState struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrInput
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrInput
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringPtrInput
	// A `routing` block as defined below.
	Routing VirtualHubConnectionRoutingPtrInput
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringPtrInput
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrInput
}

func (VirtualHubConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubConnectionState)(nil)).Elem()
}

type virtualHubConnectionArgs struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed *bool `pulumi:"hubToVitualNetworkTrafficAllowed"`
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled *bool `pulumi:"internetSecurityEnabled"`
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId string `pulumi:"remoteVirtualNetworkId"`
	// A `routing` block as defined below.
	Routing *VirtualHubConnectionRouting `pulumi:"routing"`
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId string `pulumi:"virtualHubId"`
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed *bool `pulumi:"vitualNetworkToHubGatewaysTrafficAllowed"`
}

// The set of arguments for constructing a VirtualHubConnection resource.
type VirtualHubConnectionArgs struct {
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	HubToVitualNetworkTrafficAllowed pulumi.BoolPtrInput
	// Should Internet Security be enabled to secure internet traffic? Changing this forces a new resource to be created. Defaults to `false`.
	InternetSecurityEnabled pulumi.BoolPtrInput
	// The Name which should be used for this Connection, which must be unique within the Virtual Hub. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	RemoteVirtualNetworkId pulumi.StringInput
	// A `routing` block as defined below.
	Routing VirtualHubConnectionRoutingPtrInput
	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	VirtualHubId pulumi.StringInput
	// Deprecated: Due to a breaking behavioural change in the Azure API this property is no longer functional and will be removed in version 3.0 of the provider
	VitualNetworkToHubGatewaysTrafficAllowed pulumi.BoolPtrInput
}

func (VirtualHubConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubConnectionArgs)(nil)).Elem()
}

type VirtualHubConnectionInput interface {
	pulumi.Input

	ToVirtualHubConnectionOutput() VirtualHubConnectionOutput
	ToVirtualHubConnectionOutputWithContext(ctx context.Context) VirtualHubConnectionOutput
}

func (*VirtualHubConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubConnection)(nil)).Elem()
}

func (i *VirtualHubConnection) ToVirtualHubConnectionOutput() VirtualHubConnectionOutput {
	return i.ToVirtualHubConnectionOutputWithContext(context.Background())
}

func (i *VirtualHubConnection) ToVirtualHubConnectionOutputWithContext(ctx context.Context) VirtualHubConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubConnectionOutput)
}

// VirtualHubConnectionArrayInput is an input type that accepts VirtualHubConnectionArray and VirtualHubConnectionArrayOutput values.
// You can construct a concrete instance of `VirtualHubConnectionArrayInput` via:
//
//          VirtualHubConnectionArray{ VirtualHubConnectionArgs{...} }
type VirtualHubConnectionArrayInput interface {
	pulumi.Input

	ToVirtualHubConnectionArrayOutput() VirtualHubConnectionArrayOutput
	ToVirtualHubConnectionArrayOutputWithContext(context.Context) VirtualHubConnectionArrayOutput
}

type VirtualHubConnectionArray []VirtualHubConnectionInput

func (VirtualHubConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHubConnection)(nil)).Elem()
}

func (i VirtualHubConnectionArray) ToVirtualHubConnectionArrayOutput() VirtualHubConnectionArrayOutput {
	return i.ToVirtualHubConnectionArrayOutputWithContext(context.Background())
}

func (i VirtualHubConnectionArray) ToVirtualHubConnectionArrayOutputWithContext(ctx context.Context) VirtualHubConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubConnectionArrayOutput)
}

// VirtualHubConnectionMapInput is an input type that accepts VirtualHubConnectionMap and VirtualHubConnectionMapOutput values.
// You can construct a concrete instance of `VirtualHubConnectionMapInput` via:
//
//          VirtualHubConnectionMap{ "key": VirtualHubConnectionArgs{...} }
type VirtualHubConnectionMapInput interface {
	pulumi.Input

	ToVirtualHubConnectionMapOutput() VirtualHubConnectionMapOutput
	ToVirtualHubConnectionMapOutputWithContext(context.Context) VirtualHubConnectionMapOutput
}

type VirtualHubConnectionMap map[string]VirtualHubConnectionInput

func (VirtualHubConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHubConnection)(nil)).Elem()
}

func (i VirtualHubConnectionMap) ToVirtualHubConnectionMapOutput() VirtualHubConnectionMapOutput {
	return i.ToVirtualHubConnectionMapOutputWithContext(context.Background())
}

func (i VirtualHubConnectionMap) ToVirtualHubConnectionMapOutputWithContext(ctx context.Context) VirtualHubConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubConnectionMapOutput)
}

type VirtualHubConnectionOutput struct{ *pulumi.OutputState }

func (VirtualHubConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubConnection)(nil)).Elem()
}

func (o VirtualHubConnectionOutput) ToVirtualHubConnectionOutput() VirtualHubConnectionOutput {
	return o
}

func (o VirtualHubConnectionOutput) ToVirtualHubConnectionOutputWithContext(ctx context.Context) VirtualHubConnectionOutput {
	return o
}

type VirtualHubConnectionArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHubConnection)(nil)).Elem()
}

func (o VirtualHubConnectionArrayOutput) ToVirtualHubConnectionArrayOutput() VirtualHubConnectionArrayOutput {
	return o
}

func (o VirtualHubConnectionArrayOutput) ToVirtualHubConnectionArrayOutputWithContext(ctx context.Context) VirtualHubConnectionArrayOutput {
	return o
}

func (o VirtualHubConnectionArrayOutput) Index(i pulumi.IntInput) VirtualHubConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualHubConnection {
		return vs[0].([]*VirtualHubConnection)[vs[1].(int)]
	}).(VirtualHubConnectionOutput)
}

type VirtualHubConnectionMapOutput struct{ *pulumi.OutputState }

func (VirtualHubConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHubConnection)(nil)).Elem()
}

func (o VirtualHubConnectionMapOutput) ToVirtualHubConnectionMapOutput() VirtualHubConnectionMapOutput {
	return o
}

func (o VirtualHubConnectionMapOutput) ToVirtualHubConnectionMapOutputWithContext(ctx context.Context) VirtualHubConnectionMapOutput {
	return o
}

func (o VirtualHubConnectionMapOutput) MapIndex(k pulumi.StringInput) VirtualHubConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualHubConnection {
		return vs[0].(map[string]*VirtualHubConnection)[vs[1].(string)]
	}).(VirtualHubConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHubConnectionInput)(nil)).Elem(), &VirtualHubConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHubConnectionArrayInput)(nil)).Elem(), VirtualHubConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHubConnectionMapInput)(nil)).Elem(), VirtualHubConnectionMap{})
	pulumi.RegisterOutputType(VirtualHubConnectionOutput{})
	pulumi.RegisterOutputType(VirtualHubConnectionArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubConnectionMapOutput{})
}
