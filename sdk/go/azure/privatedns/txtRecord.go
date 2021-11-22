// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS TXT Records within Azure Private DNS.
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
// 		_, err := core.NewResourceGroup(ctx, "example", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testZone, err := privatedns.NewZone(ctx, "testZone", &privatedns.ZoneArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = privatedns.NewTxtRecord(ctx, "testTxtRecord", &privatedns.TxtRecordArgs{
// 			ResourceGroupName: pulumi.Any(azurerm_resource_group.Test.Name),
// 			ZoneName:          testZone.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: privatedns.TxtRecordRecordArray{
// 				&privatedns.TxtRecordRecordArgs{
// 					Value: pulumi.String("v=spf1 mx ~all"),
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
// Private DNS TXT Records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:privatedns/txtRecord:TxtRecord test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/privateDnsZones/contoso.com/TXT/test
// ```
type TxtRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Ttl  pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewTxtRecord registers a new resource with the given unique name, arguments, and options.
func NewTxtRecord(ctx *pulumi.Context,
	name string, args *TxtRecordArgs, opts ...pulumi.ResourceOption) (*TxtRecord, error) {
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
	var resource TxtRecord
	err := ctx.RegisterResource("azure:privatedns/txtRecord:TxtRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTxtRecord gets an existing TxtRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTxtRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TxtRecordState, opts ...pulumi.ResourceOption) (*TxtRecord, error) {
	var resource TxtRecord
	err := ctx.ReadResource("azure:privatedns/txtRecord:TxtRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TxtRecord resources.
type txtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  *int              `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type TxtRecordState struct {
	// The FQDN of the DNS TXT Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntPtrInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (TxtRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordState)(nil)).Elem()
}

type txtRecordArgs struct {
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// One or more `record` blocks as defined below.
	Records []TxtRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	Ttl  int               `pulumi:"ttl"`
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a TxtRecord resource.
type TxtRecordArgs struct {
	// The name of the DNS TXT Record. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// One or more `record` blocks as defined below.
	Records TxtRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	Ttl  pulumi.IntInput
	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*txtRecordArgs)(nil)).Elem()
}

type TxtRecordInput interface {
	pulumi.Input

	ToTxtRecordOutput() TxtRecordOutput
	ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput
}

func (*TxtRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil))
}

func (i *TxtRecord) ToTxtRecordOutput() TxtRecordOutput {
	return i.ToTxtRecordOutputWithContext(context.Background())
}

func (i *TxtRecord) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordOutput)
}

func (i *TxtRecord) ToTxtRecordPtrOutput() TxtRecordPtrOutput {
	return i.ToTxtRecordPtrOutputWithContext(context.Background())
}

func (i *TxtRecord) ToTxtRecordPtrOutputWithContext(ctx context.Context) TxtRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordPtrOutput)
}

type TxtRecordPtrInput interface {
	pulumi.Input

	ToTxtRecordPtrOutput() TxtRecordPtrOutput
	ToTxtRecordPtrOutputWithContext(ctx context.Context) TxtRecordPtrOutput
}

type txtRecordPtrType TxtRecordArgs

func (*txtRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TxtRecord)(nil))
}

func (i *txtRecordPtrType) ToTxtRecordPtrOutput() TxtRecordPtrOutput {
	return i.ToTxtRecordPtrOutputWithContext(context.Background())
}

func (i *txtRecordPtrType) ToTxtRecordPtrOutputWithContext(ctx context.Context) TxtRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordPtrOutput)
}

// TxtRecordArrayInput is an input type that accepts TxtRecordArray and TxtRecordArrayOutput values.
// You can construct a concrete instance of `TxtRecordArrayInput` via:
//
//          TxtRecordArray{ TxtRecordArgs{...} }
type TxtRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordArrayOutput() TxtRecordArrayOutput
	ToTxtRecordArrayOutputWithContext(context.Context) TxtRecordArrayOutput
}

type TxtRecordArray []TxtRecordInput

func (TxtRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TxtRecord)(nil)).Elem()
}

func (i TxtRecordArray) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return i.ToTxtRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordArray) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordArrayOutput)
}

// TxtRecordMapInput is an input type that accepts TxtRecordMap and TxtRecordMapOutput values.
// You can construct a concrete instance of `TxtRecordMapInput` via:
//
//          TxtRecordMap{ "key": TxtRecordArgs{...} }
type TxtRecordMapInput interface {
	pulumi.Input

	ToTxtRecordMapOutput() TxtRecordMapOutput
	ToTxtRecordMapOutputWithContext(context.Context) TxtRecordMapOutput
}

type TxtRecordMap map[string]TxtRecordInput

func (TxtRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TxtRecord)(nil)).Elem()
}

func (i TxtRecordMap) ToTxtRecordMapOutput() TxtRecordMapOutput {
	return i.ToTxtRecordMapOutputWithContext(context.Background())
}

func (i TxtRecordMap) ToTxtRecordMapOutputWithContext(ctx context.Context) TxtRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordMapOutput)
}

type TxtRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil))
}

func (o TxtRecordOutput) ToTxtRecordOutput() TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordPtrOutput() TxtRecordPtrOutput {
	return o.ToTxtRecordPtrOutputWithContext(context.Background())
}

func (o TxtRecordOutput) ToTxtRecordPtrOutputWithContext(ctx context.Context) TxtRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TxtRecord) *TxtRecord {
		return &v
	}).(TxtRecordPtrOutput)
}

type TxtRecordPtrOutput struct{ *pulumi.OutputState }

func (TxtRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TxtRecord)(nil))
}

func (o TxtRecordPtrOutput) ToTxtRecordPtrOutput() TxtRecordPtrOutput {
	return o
}

func (o TxtRecordPtrOutput) ToTxtRecordPtrOutputWithContext(ctx context.Context) TxtRecordPtrOutput {
	return o
}

func (o TxtRecordPtrOutput) Elem() TxtRecordOutput {
	return o.ApplyT(func(v *TxtRecord) TxtRecord {
		if v != nil {
			return *v
		}
		var ret TxtRecord
		return ret
	}).(TxtRecordOutput)
}

type TxtRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil))
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].([]TxtRecord)[vs[1].(int)]
	}).(TxtRecordOutput)
}

type TxtRecordMapOutput struct{ *pulumi.OutputState }

func (TxtRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TxtRecord)(nil))
}

func (o TxtRecordMapOutput) ToTxtRecordMapOutput() TxtRecordMapOutput {
	return o
}

func (o TxtRecordMapOutput) ToTxtRecordMapOutputWithContext(ctx context.Context) TxtRecordMapOutput {
	return o
}

func (o TxtRecordMapOutput) MapIndex(k pulumi.StringInput) TxtRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].(map[string]TxtRecord)[vs[1].(string)]
	}).(TxtRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordInput)(nil)).Elem(), &TxtRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordPtrInput)(nil)).Elem(), &TxtRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordArrayInput)(nil)).Elem(), TxtRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordMapInput)(nil)).Elem(), TxtRecordMap{})
	pulumi.RegisterOutputType(TxtRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordPtrOutput{})
	pulumi.RegisterOutputType(TxtRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordMapOutput{})
}
