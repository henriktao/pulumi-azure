// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Load Balancer NAT pool.
//
// > **NOTE:** This resource cannot be used with with virtual machines, instead use the `lb.NatRule` resource.
//
// > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/lb"
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
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Static"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			FrontendIpConfigurations: lb.LoadBalancerFrontendIpConfigurationArray{
// 				&lb.LoadBalancerFrontendIpConfigurationArgs{
// 					Name:              pulumi.String("PublicIPAddress"),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lb.NewNatPool(ctx, "exampleNatPool", &lb.NatPoolArgs{
// 			ResourceGroupName:           exampleResourceGroup.Name,
// 			LoadbalancerId:              exampleLoadBalancer.ID(),
// 			Protocol:                    pulumi.String("Tcp"),
// 			FrontendPortStart:           pulumi.Int(80),
// 			FrontendPortEnd:             pulumi.Int(81),
// 			BackendPort:                 pulumi.Int(8080),
// 			FrontendIpConfigurationName: pulumi.String("PublicIPAddress"),
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
// Load Balancer NAT Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:lb/natPool:NatPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1
// ```
type NatPool struct {
	pulumi.CustomResourceState

	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntOutput `pulumi:"backendPort"`
	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	FloatingIpEnabled         pulumi.BoolPtrOutput `pulumi:"floatingIpEnabled"`
	FrontendIpConfigurationId pulumi.StringOutput  `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringOutput `pulumi:"frontendIpConfigurationName"`
	// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd pulumi.IntOutput `pulumi:"frontendPortEnd"`
	// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart pulumi.IntOutput `pulumi:"frontendPortStart"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30`. Defaults to `4`.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT pool.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	TcpResetEnabled pulumi.BoolPtrOutput `pulumi:"tcpResetEnabled"`
}

// NewNatPool registers a new resource with the given unique name, arguments, and options.
func NewNatPool(ctx *pulumi.Context,
	name string, args *NatPoolArgs, opts ...pulumi.ResourceOption) (*NatPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendPort == nil {
		return nil, errors.New("invalid value for required argument 'BackendPort'")
	}
	if args.FrontendIpConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'FrontendIpConfigurationName'")
	}
	if args.FrontendPortEnd == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPortEnd'")
	}
	if args.FrontendPortStart == nil {
		return nil, errors.New("invalid value for required argument 'FrontendPortStart'")
	}
	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NatPool
	err := ctx.RegisterResource("azure:lb/natPool:NatPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatPool gets an existing NatPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatPoolState, opts ...pulumi.ResourceOption) (*NatPool, error) {
	var resource NatPool
	err := ctx.ReadResource("azure:lb/natPool:NatPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatPool resources.
type natPoolState struct {
	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort *int `pulumi:"backendPort"`
	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	FloatingIpEnabled         *bool   `pulumi:"floatingIpEnabled"`
	FrontendIpConfigurationId *string `pulumi:"frontendIpConfigurationId"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName *string `pulumi:"frontendIpConfigurationName"`
	// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd *int `pulumi:"frontendPortEnd"`
	// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart *int `pulumi:"frontendPortStart"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30`. Defaults to `4`.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT pool.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT pool.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	TcpResetEnabled *bool `pulumi:"tcpResetEnabled"`
}

type NatPoolState struct {
	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntPtrInput
	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	FloatingIpEnabled         pulumi.BoolPtrInput
	FrontendIpConfigurationId pulumi.StringPtrInput
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringPtrInput
	// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd pulumi.IntPtrInput
	// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart pulumi.IntPtrInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30`. Defaults to `4`.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the NAT pool.
	LoadbalancerId pulumi.StringPtrInput
	// Specifies the name of the NAT pool.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
	Protocol pulumi.StringPtrInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringPtrInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	TcpResetEnabled pulumi.BoolPtrInput
}

func (NatPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*natPoolState)(nil)).Elem()
}

type natPoolArgs struct {
	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort int `pulumi:"backendPort"`
	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	FloatingIpEnabled *bool `pulumi:"floatingIpEnabled"`
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName string `pulumi:"frontendIpConfigurationName"`
	// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd int `pulumi:"frontendPortEnd"`
	// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart int `pulumi:"frontendPortStart"`
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30`. Defaults to `4`.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The ID of the Load Balancer in which to create the NAT pool.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Specifies the name of the NAT pool.
	Name *string `pulumi:"name"`
	// The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group in which to create the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	TcpResetEnabled *bool `pulumi:"tcpResetEnabled"`
}

// The set of arguments for constructing a NatPool resource.
type NatPoolArgs struct {
	// The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive.
	BackendPort pulumi.IntInput
	// Are the floating IPs enabled for this Load Balancer Rule? A floating IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to `false`.
	FloatingIpEnabled pulumi.BoolPtrInput
	// The name of the frontend IP configuration exposing this rule.
	FrontendIpConfigurationName pulumi.StringInput
	// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortEnd pulumi.IntInput
	// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.
	FrontendPortStart pulumi.IntInput
	// Specifies the idle timeout in minutes for TCP connections. Valid values are between `4` and `30`. Defaults to `4`.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The ID of the Load Balancer in which to create the NAT pool.
	LoadbalancerId pulumi.StringInput
	// Specifies the name of the NAT pool.
	Name pulumi.StringPtrInput
	// The transport protocol for the external endpoint. Possible values are `Udp` or `Tcp`.
	Protocol pulumi.StringInput
	// The name of the resource group in which to create the resource.
	ResourceGroupName pulumi.StringInput
	// Is TCP Reset enabled for this Load Balancer Rule? Defaults to `false`.
	TcpResetEnabled pulumi.BoolPtrInput
}

func (NatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natPoolArgs)(nil)).Elem()
}

type NatPoolInput interface {
	pulumi.Input

	ToNatPoolOutput() NatPoolOutput
	ToNatPoolOutputWithContext(ctx context.Context) NatPoolOutput
}

func (*NatPool) ElementType() reflect.Type {
	return reflect.TypeOf((*NatPool)(nil))
}

func (i *NatPool) ToNatPoolOutput() NatPoolOutput {
	return i.ToNatPoolOutputWithContext(context.Background())
}

func (i *NatPool) ToNatPoolOutputWithContext(ctx context.Context) NatPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPoolOutput)
}

func (i *NatPool) ToNatPoolPtrOutput() NatPoolPtrOutput {
	return i.ToNatPoolPtrOutputWithContext(context.Background())
}

func (i *NatPool) ToNatPoolPtrOutputWithContext(ctx context.Context) NatPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPoolPtrOutput)
}

type NatPoolPtrInput interface {
	pulumi.Input

	ToNatPoolPtrOutput() NatPoolPtrOutput
	ToNatPoolPtrOutputWithContext(ctx context.Context) NatPoolPtrOutput
}

type natPoolPtrType NatPoolArgs

func (*natPoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NatPool)(nil))
}

func (i *natPoolPtrType) ToNatPoolPtrOutput() NatPoolPtrOutput {
	return i.ToNatPoolPtrOutputWithContext(context.Background())
}

func (i *natPoolPtrType) ToNatPoolPtrOutputWithContext(ctx context.Context) NatPoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPoolPtrOutput)
}

// NatPoolArrayInput is an input type that accepts NatPoolArray and NatPoolArrayOutput values.
// You can construct a concrete instance of `NatPoolArrayInput` via:
//
//          NatPoolArray{ NatPoolArgs{...} }
type NatPoolArrayInput interface {
	pulumi.Input

	ToNatPoolArrayOutput() NatPoolArrayOutput
	ToNatPoolArrayOutputWithContext(context.Context) NatPoolArrayOutput
}

type NatPoolArray []NatPoolInput

func (NatPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatPool)(nil)).Elem()
}

func (i NatPoolArray) ToNatPoolArrayOutput() NatPoolArrayOutput {
	return i.ToNatPoolArrayOutputWithContext(context.Background())
}

func (i NatPoolArray) ToNatPoolArrayOutputWithContext(ctx context.Context) NatPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPoolArrayOutput)
}

// NatPoolMapInput is an input type that accepts NatPoolMap and NatPoolMapOutput values.
// You can construct a concrete instance of `NatPoolMapInput` via:
//
//          NatPoolMap{ "key": NatPoolArgs{...} }
type NatPoolMapInput interface {
	pulumi.Input

	ToNatPoolMapOutput() NatPoolMapOutput
	ToNatPoolMapOutputWithContext(context.Context) NatPoolMapOutput
}

type NatPoolMap map[string]NatPoolInput

func (NatPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatPool)(nil)).Elem()
}

func (i NatPoolMap) ToNatPoolMapOutput() NatPoolMapOutput {
	return i.ToNatPoolMapOutputWithContext(context.Background())
}

func (i NatPoolMap) ToNatPoolMapOutputWithContext(ctx context.Context) NatPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPoolMapOutput)
}

type NatPoolOutput struct{ *pulumi.OutputState }

func (NatPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatPool)(nil))
}

func (o NatPoolOutput) ToNatPoolOutput() NatPoolOutput {
	return o
}

func (o NatPoolOutput) ToNatPoolOutputWithContext(ctx context.Context) NatPoolOutput {
	return o
}

func (o NatPoolOutput) ToNatPoolPtrOutput() NatPoolPtrOutput {
	return o.ToNatPoolPtrOutputWithContext(context.Background())
}

func (o NatPoolOutput) ToNatPoolPtrOutputWithContext(ctx context.Context) NatPoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NatPool) *NatPool {
		return &v
	}).(NatPoolPtrOutput)
}

type NatPoolPtrOutput struct{ *pulumi.OutputState }

func (NatPoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatPool)(nil))
}

func (o NatPoolPtrOutput) ToNatPoolPtrOutput() NatPoolPtrOutput {
	return o
}

func (o NatPoolPtrOutput) ToNatPoolPtrOutputWithContext(ctx context.Context) NatPoolPtrOutput {
	return o
}

func (o NatPoolPtrOutput) Elem() NatPoolOutput {
	return o.ApplyT(func(v *NatPool) NatPool {
		if v != nil {
			return *v
		}
		var ret NatPool
		return ret
	}).(NatPoolOutput)
}

type NatPoolArrayOutput struct{ *pulumi.OutputState }

func (NatPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NatPool)(nil))
}

func (o NatPoolArrayOutput) ToNatPoolArrayOutput() NatPoolArrayOutput {
	return o
}

func (o NatPoolArrayOutput) ToNatPoolArrayOutputWithContext(ctx context.Context) NatPoolArrayOutput {
	return o
}

func (o NatPoolArrayOutput) Index(i pulumi.IntInput) NatPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NatPool {
		return vs[0].([]NatPool)[vs[1].(int)]
	}).(NatPoolOutput)
}

type NatPoolMapOutput struct{ *pulumi.OutputState }

func (NatPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NatPool)(nil))
}

func (o NatPoolMapOutput) ToNatPoolMapOutput() NatPoolMapOutput {
	return o
}

func (o NatPoolMapOutput) ToNatPoolMapOutputWithContext(ctx context.Context) NatPoolMapOutput {
	return o
}

func (o NatPoolMapOutput) MapIndex(k pulumi.StringInput) NatPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NatPool {
		return vs[0].(map[string]NatPool)[vs[1].(string)]
	}).(NatPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatPoolInput)(nil)).Elem(), &NatPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatPoolPtrInput)(nil)).Elem(), &NatPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatPoolArrayInput)(nil)).Elem(), NatPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatPoolMapInput)(nil)).Elem(), NatPoolMap{})
	pulumi.RegisterOutputType(NatPoolOutput{})
	pulumi.RegisterOutputType(NatPoolPtrOutput{})
	pulumi.RegisterOutputType(NatPoolArrayOutput{})
	pulumi.RegisterOutputType(NatPoolMapOutput{})
}
