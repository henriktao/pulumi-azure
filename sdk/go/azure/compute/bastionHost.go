// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Bastion Host.
//
// ## Example Usage
//
// This example deploys an Azure Bastion Host Instance to a target virtual network.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
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
// 				pulumi.String("192.168.1.0/24"),
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
// 				pulumi.String("192.168.1.224/27"),
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
// 		_, err = compute.NewBastionHost(ctx, "exampleBastionHost", &compute.BastionHostArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfiguration: &compute.BastionHostIpConfigurationArgs{
// 				Name:              pulumi.String("configuration"),
// 				SubnetId:          exampleSubnet.ID(),
// 				PublicIpAddressId: examplePublicIp.ID(),
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
// Bastion Hosts can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/bastionHost:BastionHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
// ```
type BastionHost struct {
	pulumi.CustomResourceState

	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
	CopyPasteEnabled pulumi.BoolPtrOutput `pulumi:"copyPasteEnabled"`
	// The FQDN for the Bastion Host.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
	FileCopyEnabled pulumi.BoolPtrOutput `pulumi:"fileCopyEnabled"`
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrOutput `pulumi:"ipConfiguration"`
	// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
	IpConnectEnabled pulumi.BoolPtrOutput `pulumi:"ipConnectEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
	ScaleUnits pulumi.IntPtrOutput `pulumi:"scaleUnits"`
	// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
	ShareableLinkEnabled pulumi.BoolPtrOutput `pulumi:"shareableLinkEnabled"`
	// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
	TunnelingEnabled pulumi.BoolPtrOutput `pulumi:"tunnelingEnabled"`
}

// NewBastionHost registers a new resource with the given unique name, arguments, and options.
func NewBastionHost(ctx *pulumi.Context,
	name string, args *BastionHostArgs, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource BastionHost
	err := ctx.RegisterResource("azure:compute/bastionHost:BastionHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBastionHost gets an existing BastionHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBastionHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostState, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	var resource BastionHost
	err := ctx.ReadResource("azure:compute/bastionHost:BastionHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BastionHost resources.
type bastionHostState struct {
	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
	CopyPasteEnabled *bool `pulumi:"copyPasteEnabled"`
	// The FQDN for the Bastion Host.
	DnsName *string `pulumi:"dnsName"`
	// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
	FileCopyEnabled *bool `pulumi:"fileCopyEnabled"`
	// A `ipConfiguration` block as defined below.
	IpConfiguration *BastionHostIpConfiguration `pulumi:"ipConfiguration"`
	// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
	IpConnectEnabled *bool `pulumi:"ipConnectEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
	Location *string `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
	ShareableLinkEnabled *bool `pulumi:"shareableLinkEnabled"`
	// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
	TunnelingEnabled *bool `pulumi:"tunnelingEnabled"`
}

type BastionHostState struct {
	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
	CopyPasteEnabled pulumi.BoolPtrInput
	// The FQDN for the Bastion Host.
	DnsName pulumi.StringPtrInput
	// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
	FileCopyEnabled pulumi.BoolPtrInput
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrInput
	// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
	IpConnectEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
	Location pulumi.StringPtrInput
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringPtrInput
	// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
	ScaleUnits pulumi.IntPtrInput
	// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
	ShareableLinkEnabled pulumi.BoolPtrInput
	// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
	TunnelingEnabled pulumi.BoolPtrInput
}

func (BastionHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostState)(nil)).Elem()
}

type bastionHostArgs struct {
	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
	CopyPasteEnabled *bool `pulumi:"copyPasteEnabled"`
	// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
	FileCopyEnabled *bool `pulumi:"fileCopyEnabled"`
	// A `ipConfiguration` block as defined below.
	IpConfiguration *BastionHostIpConfiguration `pulumi:"ipConfiguration"`
	// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
	IpConnectEnabled *bool `pulumi:"ipConnectEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
	Location *string `pulumi:"location"`
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
	ShareableLinkEnabled *bool `pulumi:"shareableLinkEnabled"`
	// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku *string `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
	TunnelingEnabled *bool `pulumi:"tunnelingEnabled"`
}

// The set of arguments for constructing a BastionHost resource.
type BastionHostArgs struct {
	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
	CopyPasteEnabled pulumi.BoolPtrInput
	// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
	FileCopyEnabled pulumi.BoolPtrInput
	// A `ipConfiguration` block as defined below.
	IpConfiguration BastionHostIpConfigurationPtrInput
	// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
	IpConnectEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
	Location pulumi.StringPtrInput
	// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Bastion Host.
	ResourceGroupName pulumi.StringInput
	// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
	ScaleUnits pulumi.IntPtrInput
	// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
	ShareableLinkEnabled pulumi.BoolPtrInput
	// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
	TunnelingEnabled pulumi.BoolPtrInput
}

func (BastionHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostArgs)(nil)).Elem()
}

type BastionHostInput interface {
	pulumi.Input

	ToBastionHostOutput() BastionHostOutput
	ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput
}

func (*BastionHost) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (i *BastionHost) ToBastionHostOutput() BastionHostOutput {
	return i.ToBastionHostOutputWithContext(context.Background())
}

func (i *BastionHost) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostOutput)
}

// BastionHostArrayInput is an input type that accepts BastionHostArray and BastionHostArrayOutput values.
// You can construct a concrete instance of `BastionHostArrayInput` via:
//
//          BastionHostArray{ BastionHostArgs{...} }
type BastionHostArrayInput interface {
	pulumi.Input

	ToBastionHostArrayOutput() BastionHostArrayOutput
	ToBastionHostArrayOutputWithContext(context.Context) BastionHostArrayOutput
}

type BastionHostArray []BastionHostInput

func (BastionHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BastionHost)(nil)).Elem()
}

func (i BastionHostArray) ToBastionHostArrayOutput() BastionHostArrayOutput {
	return i.ToBastionHostArrayOutputWithContext(context.Background())
}

func (i BastionHostArray) ToBastionHostArrayOutputWithContext(ctx context.Context) BastionHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostArrayOutput)
}

// BastionHostMapInput is an input type that accepts BastionHostMap and BastionHostMapOutput values.
// You can construct a concrete instance of `BastionHostMapInput` via:
//
//          BastionHostMap{ "key": BastionHostArgs{...} }
type BastionHostMapInput interface {
	pulumi.Input

	ToBastionHostMapOutput() BastionHostMapOutput
	ToBastionHostMapOutputWithContext(context.Context) BastionHostMapOutput
}

type BastionHostMap map[string]BastionHostInput

func (BastionHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BastionHost)(nil)).Elem()
}

func (i BastionHostMap) ToBastionHostMapOutput() BastionHostMapOutput {
	return i.ToBastionHostMapOutputWithContext(context.Background())
}

func (i BastionHostMap) ToBastionHostMapOutputWithContext(ctx context.Context) BastionHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostMapOutput)
}

type BastionHostOutput struct{ *pulumi.OutputState }

func (BastionHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (o BastionHostOutput) ToBastionHostOutput() BastionHostOutput {
	return o
}

func (o BastionHostOutput) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return o
}

type BastionHostArrayOutput struct{ *pulumi.OutputState }

func (BastionHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BastionHost)(nil)).Elem()
}

func (o BastionHostArrayOutput) ToBastionHostArrayOutput() BastionHostArrayOutput {
	return o
}

func (o BastionHostArrayOutput) ToBastionHostArrayOutputWithContext(ctx context.Context) BastionHostArrayOutput {
	return o
}

func (o BastionHostArrayOutput) Index(i pulumi.IntInput) BastionHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BastionHost {
		return vs[0].([]*BastionHost)[vs[1].(int)]
	}).(BastionHostOutput)
}

type BastionHostMapOutput struct{ *pulumi.OutputState }

func (BastionHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BastionHost)(nil)).Elem()
}

func (o BastionHostMapOutput) ToBastionHostMapOutput() BastionHostMapOutput {
	return o
}

func (o BastionHostMapOutput) ToBastionHostMapOutputWithContext(ctx context.Context) BastionHostMapOutput {
	return o
}

func (o BastionHostMapOutput) MapIndex(k pulumi.StringInput) BastionHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BastionHost {
		return vs[0].(map[string]*BastionHost)[vs[1].(string)]
	}).(BastionHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostInput)(nil)).Elem(), &BastionHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostArrayInput)(nil)).Elem(), BastionHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostMapInput)(nil)).Elem(), BastionHostMap{})
	pulumi.RegisterOutputType(BastionHostOutput{})
	pulumi.RegisterOutputType(BastionHostArrayOutput{})
	pulumi.RegisterOutputType(BastionHostMapOutput{})
}
