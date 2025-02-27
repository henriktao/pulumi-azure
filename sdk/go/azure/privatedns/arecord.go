// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS A Records within Azure Private DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/privatedns"
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
// 		exampleZone, err := privatedns.NewZone(ctx, "exampleZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewARecord(ctx, "exampleARecord", &privatedns.ARecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: pulumi.StringArray{
// 				pulumi.String("10.0.180.17"),
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
// Private DNS A Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/aRecord:ARecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/zone1/A/myrecord1
// ```
type ARecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS A Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of IPv4 Addresses.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewARecord registers a new resource with the given unique name, arguments, and options.
func NewARecord(ctx *pulumi.Context,
	name string, args *ARecordArgs, opts ...pulumi.ResourceOption) (*ARecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Records == nil {
		return nil, errors.New("invalid value for required argument 'Records'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	var resource ARecord
	err := ctx.RegisterResource("azure:privatedns/aRecord:ARecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetARecord gets an existing ARecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetARecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ARecordState, opts ...pulumi.ResourceOption) (*ARecord, error) {
	var resource ARecord
	err := ctx.ReadResource("azure:privatedns/aRecord:ARecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ARecord resources.
type arecordState struct {
	// The FQDN of the DNS A Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type ARecordState struct {
	// The FQDN of the DNS A Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (ARecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*arecordState)(nil)).Elem()
}

type arecordArgs struct {
	// The name of the DNS A Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a ARecord resource.
type ARecordArgs struct {
	// The name of the DNS A Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the Private DNS Zone exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (ARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arecordArgs)(nil)).Elem()
}

type ARecordInput interface {
	pulumi.Input

	ToARecordOutput() ARecordOutput
	ToARecordOutputWithContext(ctx context.Context) ARecordOutput
}

func (*ARecord) ElementType() reflect.Type {
	return reflect.TypeOf((**ARecord)(nil)).Elem()
}

func (i *ARecord) ToARecordOutput() ARecordOutput {
	return i.ToARecordOutputWithContext(context.Background())
}

func (i *ARecord) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordOutput)
}

// ARecordArrayInput is an input type that accepts ARecordArray and ARecordArrayOutput values.
// You can construct a concrete instance of `ARecordArrayInput` via:
//
//          ARecordArray{ ARecordArgs{...} }
type ARecordArrayInput interface {
	pulumi.Input

	ToARecordArrayOutput() ARecordArrayOutput
	ToARecordArrayOutputWithContext(context.Context) ARecordArrayOutput
}

type ARecordArray []ARecordInput

func (ARecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ARecord)(nil)).Elem()
}

func (i ARecordArray) ToARecordArrayOutput() ARecordArrayOutput {
	return i.ToARecordArrayOutputWithContext(context.Background())
}

func (i ARecordArray) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordArrayOutput)
}

// ARecordMapInput is an input type that accepts ARecordMap and ARecordMapOutput values.
// You can construct a concrete instance of `ARecordMapInput` via:
//
//          ARecordMap{ "key": ARecordArgs{...} }
type ARecordMapInput interface {
	pulumi.Input

	ToARecordMapOutput() ARecordMapOutput
	ToARecordMapOutputWithContext(context.Context) ARecordMapOutput
}

type ARecordMap map[string]ARecordInput

func (ARecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ARecord)(nil)).Elem()
}

func (i ARecordMap) ToARecordMapOutput() ARecordMapOutput {
	return i.ToARecordMapOutputWithContext(context.Background())
}

func (i ARecordMap) ToARecordMapOutputWithContext(ctx context.Context) ARecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordMapOutput)
}

type ARecordOutput struct{ *pulumi.OutputState }

func (ARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ARecord)(nil)).Elem()
}

func (o ARecordOutput) ToARecordOutput() ARecordOutput {
	return o
}

func (o ARecordOutput) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return o
}

type ARecordArrayOutput struct{ *pulumi.OutputState }

func (ARecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ARecord)(nil)).Elem()
}

func (o ARecordArrayOutput) ToARecordArrayOutput() ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) Index(i pulumi.IntInput) ARecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ARecord {
		return vs[0].([]*ARecord)[vs[1].(int)]
	}).(ARecordOutput)
}

type ARecordMapOutput struct{ *pulumi.OutputState }

func (ARecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ARecord)(nil)).Elem()
}

func (o ARecordMapOutput) ToARecordMapOutput() ARecordMapOutput {
	return o
}

func (o ARecordMapOutput) ToARecordMapOutputWithContext(ctx context.Context) ARecordMapOutput {
	return o
}

func (o ARecordMapOutput) MapIndex(k pulumi.StringInput) ARecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ARecord {
		return vs[0].(map[string]*ARecord)[vs[1].(string)]
	}).(ARecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordInput)(nil)).Elem(), &ARecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordArrayInput)(nil)).Elem(), ARecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordMapInput)(nil)).Elem(), ARecordMap{})
	pulumi.RegisterOutputType(ARecordOutput{})
	pulumi.RegisterOutputType(ARecordArrayOutput{})
	pulumi.RegisterOutputType(ARecordMapOutput{})
}
