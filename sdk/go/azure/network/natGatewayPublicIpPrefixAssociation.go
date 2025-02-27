// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages the association between a Nat Gateway and a Public IP Prefix.
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
// 		examplePublicIpPrefix, err := network.NewPublicIpPrefix(ctx, "examplePublicIpPrefix", &network.PublicIpPrefixArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PrefixLength:      pulumi.Int(30),
// 			Zones: pulumi.String{
// 				"1",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNatGateway, err := network.NewNatGateway(ctx, "exampleNatGateway", &network.NatGatewayArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNatGatewayPublicIpPrefixAssociation(ctx, "exampleNatGatewayPublicIpPrefixAssociation", &network.NatGatewayPublicIpPrefixAssociationArgs{
// 			NatGatewayId:     exampleNatGateway.ID(),
// 			PublicIpPrefixId: examplePublicIpPrefix.ID(),
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
// Associations between Nat Gateway and Public IP Prefixes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/natGatewayPublicIpPrefixAssociation:NatGatewayPublicIpPrefixAssociation example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPPrefixes/myPublicIpPrefix1"
// ```
type NatGatewayPublicIpPrefixAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// The ID of the Public IP Prefix which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpPrefixId pulumi.StringOutput `pulumi:"publicIpPrefixId"`
}

// NewNatGatewayPublicIpPrefixAssociation registers a new resource with the given unique name, arguments, and options.
func NewNatGatewayPublicIpPrefixAssociation(ctx *pulumi.Context,
	name string, args *NatGatewayPublicIpPrefixAssociationArgs, opts ...pulumi.ResourceOption) (*NatGatewayPublicIpPrefixAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	if args.PublicIpPrefixId == nil {
		return nil, errors.New("invalid value for required argument 'PublicIpPrefixId'")
	}
	var resource NatGatewayPublicIpPrefixAssociation
	err := ctx.RegisterResource("azure:network/natGatewayPublicIpPrefixAssociation:NatGatewayPublicIpPrefixAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGatewayPublicIpPrefixAssociation gets an existing NatGatewayPublicIpPrefixAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGatewayPublicIpPrefixAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayPublicIpPrefixAssociationState, opts ...pulumi.ResourceOption) (*NatGatewayPublicIpPrefixAssociation, error) {
	var resource NatGatewayPublicIpPrefixAssociation
	err := ctx.ReadResource("azure:network/natGatewayPublicIpPrefixAssociation:NatGatewayPublicIpPrefixAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGatewayPublicIpPrefixAssociation resources.
type natGatewayPublicIpPrefixAssociationState struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// The ID of the Public IP Prefix which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpPrefixId *string `pulumi:"publicIpPrefixId"`
}

type NatGatewayPublicIpPrefixAssociationState struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringPtrInput
	// The ID of the Public IP Prefix which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpPrefixId pulumi.StringPtrInput
}

func (NatGatewayPublicIpPrefixAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayPublicIpPrefixAssociationState)(nil)).Elem()
}

type natGatewayPublicIpPrefixAssociationArgs struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId string `pulumi:"natGatewayId"`
	// The ID of the Public IP Prefix which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpPrefixId string `pulumi:"publicIpPrefixId"`
}

// The set of arguments for constructing a NatGatewayPublicIpPrefixAssociation resource.
type NatGatewayPublicIpPrefixAssociationArgs struct {
	// The ID of the Nat Gateway. Changing this forces a new resource to be created.
	NatGatewayId pulumi.StringInput
	// The ID of the Public IP Prefix which this Nat Gateway which should be connected to. Changing this forces a new resource to be created.
	PublicIpPrefixId pulumi.StringInput
}

func (NatGatewayPublicIpPrefixAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayPublicIpPrefixAssociationArgs)(nil)).Elem()
}

type NatGatewayPublicIpPrefixAssociationInput interface {
	pulumi.Input

	ToNatGatewayPublicIpPrefixAssociationOutput() NatGatewayPublicIpPrefixAssociationOutput
	ToNatGatewayPublicIpPrefixAssociationOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationOutput
}

func (*NatGatewayPublicIpPrefixAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (i *NatGatewayPublicIpPrefixAssociation) ToNatGatewayPublicIpPrefixAssociationOutput() NatGatewayPublicIpPrefixAssociationOutput {
	return i.ToNatGatewayPublicIpPrefixAssociationOutputWithContext(context.Background())
}

func (i *NatGatewayPublicIpPrefixAssociation) ToNatGatewayPublicIpPrefixAssociationOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPublicIpPrefixAssociationOutput)
}

// NatGatewayPublicIpPrefixAssociationArrayInput is an input type that accepts NatGatewayPublicIpPrefixAssociationArray and NatGatewayPublicIpPrefixAssociationArrayOutput values.
// You can construct a concrete instance of `NatGatewayPublicIpPrefixAssociationArrayInput` via:
//
//          NatGatewayPublicIpPrefixAssociationArray{ NatGatewayPublicIpPrefixAssociationArgs{...} }
type NatGatewayPublicIpPrefixAssociationArrayInput interface {
	pulumi.Input

	ToNatGatewayPublicIpPrefixAssociationArrayOutput() NatGatewayPublicIpPrefixAssociationArrayOutput
	ToNatGatewayPublicIpPrefixAssociationArrayOutputWithContext(context.Context) NatGatewayPublicIpPrefixAssociationArrayOutput
}

type NatGatewayPublicIpPrefixAssociationArray []NatGatewayPublicIpPrefixAssociationInput

func (NatGatewayPublicIpPrefixAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (i NatGatewayPublicIpPrefixAssociationArray) ToNatGatewayPublicIpPrefixAssociationArrayOutput() NatGatewayPublicIpPrefixAssociationArrayOutput {
	return i.ToNatGatewayPublicIpPrefixAssociationArrayOutputWithContext(context.Background())
}

func (i NatGatewayPublicIpPrefixAssociationArray) ToNatGatewayPublicIpPrefixAssociationArrayOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPublicIpPrefixAssociationArrayOutput)
}

// NatGatewayPublicIpPrefixAssociationMapInput is an input type that accepts NatGatewayPublicIpPrefixAssociationMap and NatGatewayPublicIpPrefixAssociationMapOutput values.
// You can construct a concrete instance of `NatGatewayPublicIpPrefixAssociationMapInput` via:
//
//          NatGatewayPublicIpPrefixAssociationMap{ "key": NatGatewayPublicIpPrefixAssociationArgs{...} }
type NatGatewayPublicIpPrefixAssociationMapInput interface {
	pulumi.Input

	ToNatGatewayPublicIpPrefixAssociationMapOutput() NatGatewayPublicIpPrefixAssociationMapOutput
	ToNatGatewayPublicIpPrefixAssociationMapOutputWithContext(context.Context) NatGatewayPublicIpPrefixAssociationMapOutput
}

type NatGatewayPublicIpPrefixAssociationMap map[string]NatGatewayPublicIpPrefixAssociationInput

func (NatGatewayPublicIpPrefixAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (i NatGatewayPublicIpPrefixAssociationMap) ToNatGatewayPublicIpPrefixAssociationMapOutput() NatGatewayPublicIpPrefixAssociationMapOutput {
	return i.ToNatGatewayPublicIpPrefixAssociationMapOutputWithContext(context.Background())
}

func (i NatGatewayPublicIpPrefixAssociationMap) ToNatGatewayPublicIpPrefixAssociationMapOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPublicIpPrefixAssociationMapOutput)
}

type NatGatewayPublicIpPrefixAssociationOutput struct{ *pulumi.OutputState }

func (NatGatewayPublicIpPrefixAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (o NatGatewayPublicIpPrefixAssociationOutput) ToNatGatewayPublicIpPrefixAssociationOutput() NatGatewayPublicIpPrefixAssociationOutput {
	return o
}

func (o NatGatewayPublicIpPrefixAssociationOutput) ToNatGatewayPublicIpPrefixAssociationOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationOutput {
	return o
}

type NatGatewayPublicIpPrefixAssociationArrayOutput struct{ *pulumi.OutputState }

func (NatGatewayPublicIpPrefixAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (o NatGatewayPublicIpPrefixAssociationArrayOutput) ToNatGatewayPublicIpPrefixAssociationArrayOutput() NatGatewayPublicIpPrefixAssociationArrayOutput {
	return o
}

func (o NatGatewayPublicIpPrefixAssociationArrayOutput) ToNatGatewayPublicIpPrefixAssociationArrayOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationArrayOutput {
	return o
}

func (o NatGatewayPublicIpPrefixAssociationArrayOutput) Index(i pulumi.IntInput) NatGatewayPublicIpPrefixAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NatGatewayPublicIpPrefixAssociation {
		return vs[0].([]*NatGatewayPublicIpPrefixAssociation)[vs[1].(int)]
	}).(NatGatewayPublicIpPrefixAssociationOutput)
}

type NatGatewayPublicIpPrefixAssociationMapOutput struct{ *pulumi.OutputState }

func (NatGatewayPublicIpPrefixAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGatewayPublicIpPrefixAssociation)(nil)).Elem()
}

func (o NatGatewayPublicIpPrefixAssociationMapOutput) ToNatGatewayPublicIpPrefixAssociationMapOutput() NatGatewayPublicIpPrefixAssociationMapOutput {
	return o
}

func (o NatGatewayPublicIpPrefixAssociationMapOutput) ToNatGatewayPublicIpPrefixAssociationMapOutputWithContext(ctx context.Context) NatGatewayPublicIpPrefixAssociationMapOutput {
	return o
}

func (o NatGatewayPublicIpPrefixAssociationMapOutput) MapIndex(k pulumi.StringInput) NatGatewayPublicIpPrefixAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NatGatewayPublicIpPrefixAssociation {
		return vs[0].(map[string]*NatGatewayPublicIpPrefixAssociation)[vs[1].(string)]
	}).(NatGatewayPublicIpPrefixAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayPublicIpPrefixAssociationInput)(nil)).Elem(), &NatGatewayPublicIpPrefixAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayPublicIpPrefixAssociationArrayInput)(nil)).Elem(), NatGatewayPublicIpPrefixAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayPublicIpPrefixAssociationMapInput)(nil)).Elem(), NatGatewayPublicIpPrefixAssociationMap{})
	pulumi.RegisterOutputType(NatGatewayPublicIpPrefixAssociationOutput{})
	pulumi.RegisterOutputType(NatGatewayPublicIpPrefixAssociationArrayOutput{})
	pulumi.RegisterOutputType(NatGatewayPublicIpPrefixAssociationMapOutput{})
}
