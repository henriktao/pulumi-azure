// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ExpressRoute circuit.
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
// 		_, err = network.NewExpressRouteCircuit(ctx, "exampleExpressRouteCircuit", &network.ExpressRouteCircuitArgs{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Location:            exampleResourceGroup.Location,
// 			ServiceProviderName: pulumi.String("Equinix"),
// 			PeeringLocation:     pulumi.String("Silicon Valley"),
// 			BandwidthInMbps:     pulumi.Int(50),
// 			Sku: &network.ExpressRouteCircuitSkuArgs{
// 				Tier:   pulumi.String("Standard"),
// 				Family: pulumi.String("MeteredData"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// ExpressRoute circuits can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/expressRouteCircuit:ExpressRouteCircuit myExpressRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/expressRouteCircuits/myExpressRoute
// ```
type ExpressRouteCircuit struct {
	pulumi.CustomResourceState

	// Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
	AllowClassicOperations pulumi.BoolPtrOutput `pulumi:"allowClassicOperations"`
	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps pulumi.Float64PtrOutput `pulumi:"bandwidthInGbps"`
	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps pulumi.IntPtrOutput `pulumi:"bandwidthInMbps"`
	// The ID of the Express Route Port this Express Route Circuit is based on.
	ExpressRoutePortId pulumi.StringPtrOutput `pulumi:"expressRoutePortId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation pulumi.StringPtrOutput `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey pulumi.StringOutput `pulumi:"serviceKey"`
	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName pulumi.StringPtrOutput `pulumi:"serviceProviderName"`
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState pulumi.StringOutput `pulumi:"serviceProviderProvisioningState"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewExpressRouteCircuit registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	var resource ExpressRouteCircuit
	err := ctx.RegisterResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuit gets an existing ExpressRouteCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	var resource ExpressRouteCircuit
	err := ctx.ReadResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuit resources.
type expressRouteCircuitState struct {
	// Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps *float64 `pulumi:"bandwidthInGbps"`
	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps *int `pulumi:"bandwidthInMbps"`
	// The ID of the Express Route Port this Express Route Circuit is based on.
	ExpressRoutePortId *string `pulumi:"expressRoutePortId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey *string `pulumi:"serviceKey"`
	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName *string `pulumi:"serviceProviderName"`
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState *string `pulumi:"serviceProviderProvisioningState"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku *ExpressRouteCircuitSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ExpressRouteCircuitState struct {
	// Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
	AllowClassicOperations pulumi.BoolPtrInput
	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps pulumi.Float64PtrInput
	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps pulumi.IntPtrInput
	// The ID of the Express Route Port this Express Route Circuit is based on.
	ExpressRoutePortId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation pulumi.StringPtrInput
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey pulumi.StringPtrInput
	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName pulumi.StringPtrInput
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState pulumi.StringPtrInput
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitState)(nil)).Elem()
}

type expressRouteCircuitArgs struct {
	// Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps *float64 `pulumi:"bandwidthInGbps"`
	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps *int `pulumi:"bandwidthInMbps"`
	// The ID of the Express Route Port this Express Route Circuit is based on.
	ExpressRoutePortId *string `pulumi:"expressRoutePortId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName *string `pulumi:"serviceProviderName"`
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ExpressRouteCircuit resource.
type ExpressRouteCircuitArgs struct {
	// Allow the circuit to interact with classic (RDFE) resources. Defaults to `false`.
	AllowClassicOperations pulumi.BoolPtrInput
	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	BandwidthInGbps pulumi.Float64PtrInput
	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	BandwidthInMbps pulumi.IntPtrInput
	// The ID of the Express Route Port this Express Route Circuit is based on.
	ExpressRoutePortId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the peering location and **not** the Azure resource location. Changing this forces a new resource to be created.
	PeeringLocation pulumi.StringPtrInput
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	ServiceProviderName pulumi.StringPtrInput
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku ExpressRouteCircuitSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ExpressRouteCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitArgs)(nil)).Elem()
}

type ExpressRouteCircuitInput interface {
	pulumi.Input

	ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput
	ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput
}

func (*ExpressRouteCircuit) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuit)(nil))
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return i.ToExpressRouteCircuitOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitOutput)
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitPtrOutput() ExpressRouteCircuitPtrOutput {
	return i.ToExpressRouteCircuitPtrOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPtrOutput)
}

type ExpressRouteCircuitPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPtrOutput() ExpressRouteCircuitPtrOutput
	ToExpressRouteCircuitPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPtrOutput
}

type expressRouteCircuitPtrType ExpressRouteCircuitArgs

func (*expressRouteCircuitPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil))
}

func (i *expressRouteCircuitPtrType) ToExpressRouteCircuitPtrOutput() ExpressRouteCircuitPtrOutput {
	return i.ToExpressRouteCircuitPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitPtrType) ToExpressRouteCircuitPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPtrOutput)
}

// ExpressRouteCircuitArrayInput is an input type that accepts ExpressRouteCircuitArray and ExpressRouteCircuitArrayOutput values.
// You can construct a concrete instance of `ExpressRouteCircuitArrayInput` via:
//
//          ExpressRouteCircuitArray{ ExpressRouteCircuitArgs{...} }
type ExpressRouteCircuitArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitArrayOutput() ExpressRouteCircuitArrayOutput
	ToExpressRouteCircuitArrayOutputWithContext(context.Context) ExpressRouteCircuitArrayOutput
}

type ExpressRouteCircuitArray []ExpressRouteCircuitInput

func (ExpressRouteCircuitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExpressRouteCircuit)(nil)).Elem()
}

func (i ExpressRouteCircuitArray) ToExpressRouteCircuitArrayOutput() ExpressRouteCircuitArrayOutput {
	return i.ToExpressRouteCircuitArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitArray) ToExpressRouteCircuitArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitArrayOutput)
}

// ExpressRouteCircuitMapInput is an input type that accepts ExpressRouteCircuitMap and ExpressRouteCircuitMapOutput values.
// You can construct a concrete instance of `ExpressRouteCircuitMapInput` via:
//
//          ExpressRouteCircuitMap{ "key": ExpressRouteCircuitArgs{...} }
type ExpressRouteCircuitMapInput interface {
	pulumi.Input

	ToExpressRouteCircuitMapOutput() ExpressRouteCircuitMapOutput
	ToExpressRouteCircuitMapOutputWithContext(context.Context) ExpressRouteCircuitMapOutput
}

type ExpressRouteCircuitMap map[string]ExpressRouteCircuitInput

func (ExpressRouteCircuitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExpressRouteCircuit)(nil)).Elem()
}

func (i ExpressRouteCircuitMap) ToExpressRouteCircuitMapOutput() ExpressRouteCircuitMapOutput {
	return i.ToExpressRouteCircuitMapOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitMap) ToExpressRouteCircuitMapOutputWithContext(ctx context.Context) ExpressRouteCircuitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitMapOutput)
}

type ExpressRouteCircuitOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuit)(nil))
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return o
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return o
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitPtrOutput() ExpressRouteCircuitPtrOutput {
	return o.ToExpressRouteCircuitPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuit) *ExpressRouteCircuit {
		return &v
	}).(ExpressRouteCircuitPtrOutput)
}

type ExpressRouteCircuitPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil))
}

func (o ExpressRouteCircuitPtrOutput) ToExpressRouteCircuitPtrOutput() ExpressRouteCircuitPtrOutput {
	return o
}

func (o ExpressRouteCircuitPtrOutput) ToExpressRouteCircuitPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPtrOutput {
	return o
}

func (o ExpressRouteCircuitPtrOutput) Elem() ExpressRouteCircuitOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuit {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuit
		return ret
	}).(ExpressRouteCircuitOutput)
}

type ExpressRouteCircuitArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuit)(nil))
}

func (o ExpressRouteCircuitArrayOutput) ToExpressRouteCircuitArrayOutput() ExpressRouteCircuitArrayOutput {
	return o
}

func (o ExpressRouteCircuitArrayOutput) ToExpressRouteCircuitArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitArrayOutput {
	return o
}

func (o ExpressRouteCircuitArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuit {
		return vs[0].([]ExpressRouteCircuit)[vs[1].(int)]
	}).(ExpressRouteCircuitOutput)
}

type ExpressRouteCircuitMapOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExpressRouteCircuit)(nil))
}

func (o ExpressRouteCircuitMapOutput) ToExpressRouteCircuitMapOutput() ExpressRouteCircuitMapOutput {
	return o
}

func (o ExpressRouteCircuitMapOutput) ToExpressRouteCircuitMapOutputWithContext(ctx context.Context) ExpressRouteCircuitMapOutput {
	return o
}

func (o ExpressRouteCircuitMapOutput) MapIndex(k pulumi.StringInput) ExpressRouteCircuitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExpressRouteCircuit {
		return vs[0].(map[string]ExpressRouteCircuit)[vs[1].(string)]
	}).(ExpressRouteCircuitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitInput)(nil)).Elem(), &ExpressRouteCircuit{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitPtrInput)(nil)).Elem(), &ExpressRouteCircuit{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitArrayInput)(nil)).Elem(), ExpressRouteCircuitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressRouteCircuitMapInput)(nil)).Elem(), ExpressRouteCircuitMap{})
	pulumi.RegisterOutputType(ExpressRouteCircuitOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitMapOutput{})
}
