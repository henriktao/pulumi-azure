// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables you to manage DNS SRV Records within Azure DNS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/dns"
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
// 		exampleZone, err := dns.NewZone(ctx, "exampleZone", &dns.ZoneArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewSrvRecord(ctx, "exampleSrvRecord", &dns.SrvRecordArgs{
// 			ZoneName:          exampleZone.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Ttl:               pulumi.Int(300),
// 			Records: dns.SrvRecordRecordArray{
// 				&dns.SrvRecordRecordArgs{
// 					Priority: pulumi.Int(1),
// 					Weight:   pulumi.Int(5),
// 					Port:     pulumi.Int(8080),
// 					Target:   pulumi.String("target1.contoso.com"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
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
// SRV records can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:dns/srvRecord:SrvRecord example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnszones/zone1/SRV/myrecord1
// ```
type SrvRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS SRV Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS SRV Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records SrvRecordRecordArrayOutput `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewSrvRecord registers a new resource with the given unique name, arguments, and options.
func NewSrvRecord(ctx *pulumi.Context,
	name string, args *SrvRecordArgs, opts ...pulumi.ResourceOption) (*SrvRecord, error) {
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
	var resource SrvRecord
	err := ctx.RegisterResource("azure:dns/srvRecord:SrvRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSrvRecord gets an existing SrvRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSrvRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SrvRecordState, opts ...pulumi.ResourceOption) (*SrvRecord, error) {
	var resource SrvRecord
	err := ctx.ReadResource("azure:dns/srvRecord:SrvRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SrvRecord resources.
type srvRecordState struct {
	// The FQDN of the DNS SRV Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS SRV Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records []SrvRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl *int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type SrvRecordState struct {
	// The FQDN of the DNS SRV Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS SRV Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records SrvRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntPtrInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (SrvRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*srvRecordState)(nil)).Elem()
}

type srvRecordArgs struct {
	// The name of the DNS SRV Record.
	Name *string `pulumi:"name"`
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records []SrvRecordRecord `pulumi:"records"`
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl int `pulumi:"ttl"`
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a SrvRecord resource.
type SrvRecordArgs struct {
	// The name of the DNS SRV Record.
	Name pulumi.StringPtrInput
	// A list of values that make up the SRV record. Each `record` block supports fields documented below.
	Records SrvRecordRecordArrayInput
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Time To Live (TTL) of the DNS record in seconds.
	Ttl pulumi.IntInput
	// Specifies the DNS Zone where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (SrvRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*srvRecordArgs)(nil)).Elem()
}

type SrvRecordInput interface {
	pulumi.Input

	ToSrvRecordOutput() SrvRecordOutput
	ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput
}

func (*SrvRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil))
}

func (i *SrvRecord) ToSrvRecordOutput() SrvRecordOutput {
	return i.ToSrvRecordOutputWithContext(context.Background())
}

func (i *SrvRecord) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordOutput)
}

func (i *SrvRecord) ToSrvRecordPtrOutput() SrvRecordPtrOutput {
	return i.ToSrvRecordPtrOutputWithContext(context.Background())
}

func (i *SrvRecord) ToSrvRecordPtrOutputWithContext(ctx context.Context) SrvRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordPtrOutput)
}

type SrvRecordPtrInput interface {
	pulumi.Input

	ToSrvRecordPtrOutput() SrvRecordPtrOutput
	ToSrvRecordPtrOutputWithContext(ctx context.Context) SrvRecordPtrOutput
}

type srvRecordPtrType SrvRecordArgs

func (*srvRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SrvRecord)(nil))
}

func (i *srvRecordPtrType) ToSrvRecordPtrOutput() SrvRecordPtrOutput {
	return i.ToSrvRecordPtrOutputWithContext(context.Background())
}

func (i *srvRecordPtrType) ToSrvRecordPtrOutputWithContext(ctx context.Context) SrvRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordPtrOutput)
}

// SrvRecordArrayInput is an input type that accepts SrvRecordArray and SrvRecordArrayOutput values.
// You can construct a concrete instance of `SrvRecordArrayInput` via:
//
//          SrvRecordArray{ SrvRecordArgs{...} }
type SrvRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordArrayOutput() SrvRecordArrayOutput
	ToSrvRecordArrayOutputWithContext(context.Context) SrvRecordArrayOutput
}

type SrvRecordArray []SrvRecordInput

func (SrvRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SrvRecord)(nil)).Elem()
}

func (i SrvRecordArray) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return i.ToSrvRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordArray) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordArrayOutput)
}

// SrvRecordMapInput is an input type that accepts SrvRecordMap and SrvRecordMapOutput values.
// You can construct a concrete instance of `SrvRecordMapInput` via:
//
//          SrvRecordMap{ "key": SrvRecordArgs{...} }
type SrvRecordMapInput interface {
	pulumi.Input

	ToSrvRecordMapOutput() SrvRecordMapOutput
	ToSrvRecordMapOutputWithContext(context.Context) SrvRecordMapOutput
}

type SrvRecordMap map[string]SrvRecordInput

func (SrvRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SrvRecord)(nil)).Elem()
}

func (i SrvRecordMap) ToSrvRecordMapOutput() SrvRecordMapOutput {
	return i.ToSrvRecordMapOutputWithContext(context.Background())
}

func (i SrvRecordMap) ToSrvRecordMapOutputWithContext(ctx context.Context) SrvRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordMapOutput)
}

type SrvRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil))
}

func (o SrvRecordOutput) ToSrvRecordOutput() SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordPtrOutput() SrvRecordPtrOutput {
	return o.ToSrvRecordPtrOutputWithContext(context.Background())
}

func (o SrvRecordOutput) ToSrvRecordPtrOutputWithContext(ctx context.Context) SrvRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SrvRecord) *SrvRecord {
		return &v
	}).(SrvRecordPtrOutput)
}

type SrvRecordPtrOutput struct{ *pulumi.OutputState }

func (SrvRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SrvRecord)(nil))
}

func (o SrvRecordPtrOutput) ToSrvRecordPtrOutput() SrvRecordPtrOutput {
	return o
}

func (o SrvRecordPtrOutput) ToSrvRecordPtrOutputWithContext(ctx context.Context) SrvRecordPtrOutput {
	return o
}

func (o SrvRecordPtrOutput) Elem() SrvRecordOutput {
	return o.ApplyT(func(v *SrvRecord) SrvRecord {
		if v != nil {
			return *v
		}
		var ret SrvRecord
		return ret
	}).(SrvRecordOutput)
}

type SrvRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil))
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].([]SrvRecord)[vs[1].(int)]
	}).(SrvRecordOutput)
}

type SrvRecordMapOutput struct{ *pulumi.OutputState }

func (SrvRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SrvRecord)(nil))
}

func (o SrvRecordMapOutput) ToSrvRecordMapOutput() SrvRecordMapOutput {
	return o
}

func (o SrvRecordMapOutput) ToSrvRecordMapOutputWithContext(ctx context.Context) SrvRecordMapOutput {
	return o
}

func (o SrvRecordMapOutput) MapIndex(k pulumi.StringInput) SrvRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].(map[string]SrvRecord)[vs[1].(string)]
	}).(SrvRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordInput)(nil)).Elem(), &SrvRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordPtrInput)(nil)).Elem(), &SrvRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordArrayInput)(nil)).Elem(), SrvRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordMapInput)(nil)).Elem(), SrvRecordMap{})
	pulumi.RegisterOutputType(SrvRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordPtrOutput{})
	pulumi.RegisterOutputType(SrvRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordMapOutput{})
}
