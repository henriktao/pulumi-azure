// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CaaRecordRecord struct {
	// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
	Flags int `pulumi:"flags"`
	// A property tag, options are issue, issuewild and iodef.
	Tag string `pulumi:"tag"`
	// A property value such as a registrar domain.
	Value string `pulumi:"value"`
}

// CaaRecordRecordInput is an input type that accepts CaaRecordRecordArgs and CaaRecordRecordOutput values.
// You can construct a concrete instance of `CaaRecordRecordInput` via:
//
//          CaaRecordRecordArgs{...}
type CaaRecordRecordInput interface {
	pulumi.Input

	ToCaaRecordRecordOutput() CaaRecordRecordOutput
	ToCaaRecordRecordOutputWithContext(context.Context) CaaRecordRecordOutput
}

type CaaRecordRecordArgs struct {
	// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
	Flags pulumi.IntInput `pulumi:"flags"`
	// A property tag, options are issue, issuewild and iodef.
	Tag pulumi.StringInput `pulumi:"tag"`
	// A property value such as a registrar domain.
	Value pulumi.StringInput `pulumi:"value"`
}

func (CaaRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CaaRecordRecord)(nil)).Elem()
}

func (i CaaRecordRecordArgs) ToCaaRecordRecordOutput() CaaRecordRecordOutput {
	return i.ToCaaRecordRecordOutputWithContext(context.Background())
}

func (i CaaRecordRecordArgs) ToCaaRecordRecordOutputWithContext(ctx context.Context) CaaRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaaRecordRecordOutput)
}

// CaaRecordRecordArrayInput is an input type that accepts CaaRecordRecordArray and CaaRecordRecordArrayOutput values.
// You can construct a concrete instance of `CaaRecordRecordArrayInput` via:
//
//          CaaRecordRecordArray{ CaaRecordRecordArgs{...} }
type CaaRecordRecordArrayInput interface {
	pulumi.Input

	ToCaaRecordRecordArrayOutput() CaaRecordRecordArrayOutput
	ToCaaRecordRecordArrayOutputWithContext(context.Context) CaaRecordRecordArrayOutput
}

type CaaRecordRecordArray []CaaRecordRecordInput

func (CaaRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CaaRecordRecord)(nil)).Elem()
}

func (i CaaRecordRecordArray) ToCaaRecordRecordArrayOutput() CaaRecordRecordArrayOutput {
	return i.ToCaaRecordRecordArrayOutputWithContext(context.Background())
}

func (i CaaRecordRecordArray) ToCaaRecordRecordArrayOutputWithContext(ctx context.Context) CaaRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaaRecordRecordArrayOutput)
}

type CaaRecordRecordOutput struct{ *pulumi.OutputState }

func (CaaRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaaRecordRecord)(nil)).Elem()
}

func (o CaaRecordRecordOutput) ToCaaRecordRecordOutput() CaaRecordRecordOutput {
	return o
}

func (o CaaRecordRecordOutput) ToCaaRecordRecordOutputWithContext(ctx context.Context) CaaRecordRecordOutput {
	return o
}

// Extensible CAA flags, currently only 1 is implemented to set the issuer critical flag.
func (o CaaRecordRecordOutput) Flags() pulumi.IntOutput {
	return o.ApplyT(func(v CaaRecordRecord) int { return v.Flags }).(pulumi.IntOutput)
}

// A property tag, options are issue, issuewild and iodef.
func (o CaaRecordRecordOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v CaaRecordRecord) string { return v.Tag }).(pulumi.StringOutput)
}

// A property value such as a registrar domain.
func (o CaaRecordRecordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CaaRecordRecord) string { return v.Value }).(pulumi.StringOutput)
}

type CaaRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (CaaRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CaaRecordRecord)(nil)).Elem()
}

func (o CaaRecordRecordArrayOutput) ToCaaRecordRecordArrayOutput() CaaRecordRecordArrayOutput {
	return o
}

func (o CaaRecordRecordArrayOutput) ToCaaRecordRecordArrayOutputWithContext(ctx context.Context) CaaRecordRecordArrayOutput {
	return o
}

func (o CaaRecordRecordArrayOutput) Index(i pulumi.IntInput) CaaRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CaaRecordRecord {
		return vs[0].([]CaaRecordRecord)[vs[1].(int)]
	}).(CaaRecordRecordOutput)
}

type MxRecordRecord struct {
	// The mail server responsible for the domain covered by the MX record.
	Exchange string `pulumi:"exchange"`
	// String representing the "preference” value of the MX records. Records with lower preference value take priority.
	Preference string `pulumi:"preference"`
}

// MxRecordRecordInput is an input type that accepts MxRecordRecordArgs and MxRecordRecordOutput values.
// You can construct a concrete instance of `MxRecordRecordInput` via:
//
//          MxRecordRecordArgs{...}
type MxRecordRecordInput interface {
	pulumi.Input

	ToMxRecordRecordOutput() MxRecordRecordOutput
	ToMxRecordRecordOutputWithContext(context.Context) MxRecordRecordOutput
}

type MxRecordRecordArgs struct {
	// The mail server responsible for the domain covered by the MX record.
	Exchange pulumi.StringInput `pulumi:"exchange"`
	// String representing the "preference” value of the MX records. Records with lower preference value take priority.
	Preference pulumi.StringInput `pulumi:"preference"`
}

func (MxRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return i.ToMxRecordRecordOutputWithContext(context.Background())
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordOutput)
}

// MxRecordRecordArrayInput is an input type that accepts MxRecordRecordArray and MxRecordRecordArrayOutput values.
// You can construct a concrete instance of `MxRecordRecordArrayInput` via:
//
//          MxRecordRecordArray{ MxRecordRecordArgs{...} }
type MxRecordRecordArrayInput interface {
	pulumi.Input

	ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput
	ToMxRecordRecordArrayOutputWithContext(context.Context) MxRecordRecordArrayOutput
}

type MxRecordRecordArray []MxRecordRecordInput

func (MxRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return i.ToMxRecordRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordArrayOutput)
}

type MxRecordRecordOutput struct{ *pulumi.OutputState }

func (MxRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return o
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return o
}

// The mail server responsible for the domain covered by the MX record.
func (o MxRecordRecordOutput) Exchange() pulumi.StringOutput {
	return o.ApplyT(func(v MxRecordRecord) string { return v.Exchange }).(pulumi.StringOutput)
}

// String representing the "preference” value of the MX records. Records with lower preference value take priority.
func (o MxRecordRecordOutput) Preference() pulumi.StringOutput {
	return o.ApplyT(func(v MxRecordRecord) string { return v.Preference }).(pulumi.StringOutput)
}

type MxRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) Index(i pulumi.IntInput) MxRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecordRecord {
		return vs[0].([]MxRecordRecord)[vs[1].(int)]
	}).(MxRecordRecordOutput)
}

type SrvRecordRecord struct {
	// Port the service is listening on.
	Port int `pulumi:"port"`
	// Priority of the SRV record.
	Priority int `pulumi:"priority"`
	// FQDN of the service.
	Target string `pulumi:"target"`
	// Weight of the SRV record.
	Weight int `pulumi:"weight"`
}

// SrvRecordRecordInput is an input type that accepts SrvRecordRecordArgs and SrvRecordRecordOutput values.
// You can construct a concrete instance of `SrvRecordRecordInput` via:
//
//          SrvRecordRecordArgs{...}
type SrvRecordRecordInput interface {
	pulumi.Input

	ToSrvRecordRecordOutput() SrvRecordRecordOutput
	ToSrvRecordRecordOutputWithContext(context.Context) SrvRecordRecordOutput
}

type SrvRecordRecordArgs struct {
	// Port the service is listening on.
	Port pulumi.IntInput `pulumi:"port"`
	// Priority of the SRV record.
	Priority pulumi.IntInput `pulumi:"priority"`
	// FQDN of the service.
	Target pulumi.StringInput `pulumi:"target"`
	// Weight of the SRV record.
	Weight pulumi.IntInput `pulumi:"weight"`
}

func (SrvRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordRecord)(nil)).Elem()
}

func (i SrvRecordRecordArgs) ToSrvRecordRecordOutput() SrvRecordRecordOutput {
	return i.ToSrvRecordRecordOutputWithContext(context.Background())
}

func (i SrvRecordRecordArgs) ToSrvRecordRecordOutputWithContext(ctx context.Context) SrvRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordRecordOutput)
}

// SrvRecordRecordArrayInput is an input type that accepts SrvRecordRecordArray and SrvRecordRecordArrayOutput values.
// You can construct a concrete instance of `SrvRecordRecordArrayInput` via:
//
//          SrvRecordRecordArray{ SrvRecordRecordArgs{...} }
type SrvRecordRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordRecordArrayOutput() SrvRecordRecordArrayOutput
	ToSrvRecordRecordArrayOutputWithContext(context.Context) SrvRecordRecordArrayOutput
}

type SrvRecordRecordArray []SrvRecordRecordInput

func (SrvRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordRecord)(nil)).Elem()
}

func (i SrvRecordRecordArray) ToSrvRecordRecordArrayOutput() SrvRecordRecordArrayOutput {
	return i.ToSrvRecordRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordRecordArray) ToSrvRecordRecordArrayOutputWithContext(ctx context.Context) SrvRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordRecordArrayOutput)
}

type SrvRecordRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordRecord)(nil)).Elem()
}

func (o SrvRecordRecordOutput) ToSrvRecordRecordOutput() SrvRecordRecordOutput {
	return o
}

func (o SrvRecordRecordOutput) ToSrvRecordRecordOutputWithContext(ctx context.Context) SrvRecordRecordOutput {
	return o
}

// Port the service is listening on.
func (o SrvRecordRecordOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v SrvRecordRecord) int { return v.Port }).(pulumi.IntOutput)
}

// Priority of the SRV record.
func (o SrvRecordRecordOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v SrvRecordRecord) int { return v.Priority }).(pulumi.IntOutput)
}

// FQDN of the service.
func (o SrvRecordRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v SrvRecordRecord) string { return v.Target }).(pulumi.StringOutput)
}

// Weight of the SRV record.
func (o SrvRecordRecordOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v SrvRecordRecord) int { return v.Weight }).(pulumi.IntOutput)
}

type SrvRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordRecord)(nil)).Elem()
}

func (o SrvRecordRecordArrayOutput) ToSrvRecordRecordArrayOutput() SrvRecordRecordArrayOutput {
	return o
}

func (o SrvRecordRecordArrayOutput) ToSrvRecordRecordArrayOutputWithContext(ctx context.Context) SrvRecordRecordArrayOutput {
	return o
}

func (o SrvRecordRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecordRecord {
		return vs[0].([]SrvRecordRecord)[vs[1].(int)]
	}).(SrvRecordRecordOutput)
}

type TxtRecordRecord struct {
	// The value of the record. Max length: 1024 characters
	Value string `pulumi:"value"`
}

// TxtRecordRecordInput is an input type that accepts TxtRecordRecordArgs and TxtRecordRecordOutput values.
// You can construct a concrete instance of `TxtRecordRecordInput` via:
//
//          TxtRecordRecordArgs{...}
type TxtRecordRecordInput interface {
	pulumi.Input

	ToTxtRecordRecordOutput() TxtRecordRecordOutput
	ToTxtRecordRecordOutputWithContext(context.Context) TxtRecordRecordOutput
}

type TxtRecordRecordArgs struct {
	// The value of the record. Max length: 1024 characters
	Value pulumi.StringInput `pulumi:"value"`
}

func (TxtRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordRecord)(nil)).Elem()
}

func (i TxtRecordRecordArgs) ToTxtRecordRecordOutput() TxtRecordRecordOutput {
	return i.ToTxtRecordRecordOutputWithContext(context.Background())
}

func (i TxtRecordRecordArgs) ToTxtRecordRecordOutputWithContext(ctx context.Context) TxtRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordRecordOutput)
}

// TxtRecordRecordArrayInput is an input type that accepts TxtRecordRecordArray and TxtRecordRecordArrayOutput values.
// You can construct a concrete instance of `TxtRecordRecordArrayInput` via:
//
//          TxtRecordRecordArray{ TxtRecordRecordArgs{...} }
type TxtRecordRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput
	ToTxtRecordRecordArrayOutputWithContext(context.Context) TxtRecordRecordArrayOutput
}

type TxtRecordRecordArray []TxtRecordRecordInput

func (TxtRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordRecord)(nil)).Elem()
}

func (i TxtRecordRecordArray) ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput {
	return i.ToTxtRecordRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordRecordArray) ToTxtRecordRecordArrayOutputWithContext(ctx context.Context) TxtRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordRecordArrayOutput)
}

type TxtRecordRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordRecord)(nil)).Elem()
}

func (o TxtRecordRecordOutput) ToTxtRecordRecordOutput() TxtRecordRecordOutput {
	return o
}

func (o TxtRecordRecordOutput) ToTxtRecordRecordOutputWithContext(ctx context.Context) TxtRecordRecordOutput {
	return o
}

// The value of the record. Max length: 1024 characters
func (o TxtRecordRecordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TxtRecordRecord) string { return v.Value }).(pulumi.StringOutput)
}

type TxtRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordRecord)(nil)).Elem()
}

func (o TxtRecordRecordArrayOutput) ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput {
	return o
}

func (o TxtRecordRecordArrayOutput) ToTxtRecordRecordArrayOutputWithContext(ctx context.Context) TxtRecordRecordArrayOutput {
	return o
}

func (o TxtRecordRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordRecord {
		return vs[0].([]TxtRecordRecord)[vs[1].(int)]
	}).(TxtRecordRecordOutput)
}

type ZoneSoaRecord struct {
	// The email contact for the SOA record.
	Email string `pulumi:"email"`
	// The expire time for the SOA record. Defaults to `2419200`.
	ExpireTime *int    `pulumi:"expireTime"`
	Fqdn       *string `pulumi:"fqdn"`
	// The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
	HostName string `pulumi:"hostName"`
	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
	MinimumTtl *int `pulumi:"minimumTtl"`
	// The refresh time for the SOA record. Defaults to `3600`.
	RefreshTime *int `pulumi:"refreshTime"`
	// The retry time for the SOA record. Defaults to `300`.
	RetryTime *int `pulumi:"retryTime"`
	// The serial number for the SOA record. Defaults to `1`.
	SerialNumber *int `pulumi:"serialNumber"`
	// A mapping of tags to assign to the Record Set.
	Tags map[string]string `pulumi:"tags"`
	// The Time To Live of the SOA Record in seconds. Defaults to `3600`.
	Ttl *int `pulumi:"ttl"`
}

// ZoneSoaRecordInput is an input type that accepts ZoneSoaRecordArgs and ZoneSoaRecordOutput values.
// You can construct a concrete instance of `ZoneSoaRecordInput` via:
//
//          ZoneSoaRecordArgs{...}
type ZoneSoaRecordInput interface {
	pulumi.Input

	ToZoneSoaRecordOutput() ZoneSoaRecordOutput
	ToZoneSoaRecordOutputWithContext(context.Context) ZoneSoaRecordOutput
}

type ZoneSoaRecordArgs struct {
	// The email contact for the SOA record.
	Email pulumi.StringInput `pulumi:"email"`
	// The expire time for the SOA record. Defaults to `2419200`.
	ExpireTime pulumi.IntPtrInput    `pulumi:"expireTime"`
	Fqdn       pulumi.StringPtrInput `pulumi:"fqdn"`
	// The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
	HostName pulumi.StringInput `pulumi:"hostName"`
	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
	MinimumTtl pulumi.IntPtrInput `pulumi:"minimumTtl"`
	// The refresh time for the SOA record. Defaults to `3600`.
	RefreshTime pulumi.IntPtrInput `pulumi:"refreshTime"`
	// The retry time for the SOA record. Defaults to `300`.
	RetryTime pulumi.IntPtrInput `pulumi:"retryTime"`
	// The serial number for the SOA record. Defaults to `1`.
	SerialNumber pulumi.IntPtrInput `pulumi:"serialNumber"`
	// A mapping of tags to assign to the Record Set.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The Time To Live of the SOA Record in seconds. Defaults to `3600`.
	Ttl pulumi.IntPtrInput `pulumi:"ttl"`
}

func (ZoneSoaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneSoaRecord)(nil)).Elem()
}

func (i ZoneSoaRecordArgs) ToZoneSoaRecordOutput() ZoneSoaRecordOutput {
	return i.ToZoneSoaRecordOutputWithContext(context.Background())
}

func (i ZoneSoaRecordArgs) ToZoneSoaRecordOutputWithContext(ctx context.Context) ZoneSoaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSoaRecordOutput)
}

func (i ZoneSoaRecordArgs) ToZoneSoaRecordPtrOutput() ZoneSoaRecordPtrOutput {
	return i.ToZoneSoaRecordPtrOutputWithContext(context.Background())
}

func (i ZoneSoaRecordArgs) ToZoneSoaRecordPtrOutputWithContext(ctx context.Context) ZoneSoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSoaRecordOutput).ToZoneSoaRecordPtrOutputWithContext(ctx)
}

// ZoneSoaRecordPtrInput is an input type that accepts ZoneSoaRecordArgs, ZoneSoaRecordPtr and ZoneSoaRecordPtrOutput values.
// You can construct a concrete instance of `ZoneSoaRecordPtrInput` via:
//
//          ZoneSoaRecordArgs{...}
//
//  or:
//
//          nil
type ZoneSoaRecordPtrInput interface {
	pulumi.Input

	ToZoneSoaRecordPtrOutput() ZoneSoaRecordPtrOutput
	ToZoneSoaRecordPtrOutputWithContext(context.Context) ZoneSoaRecordPtrOutput
}

type zoneSoaRecordPtrType ZoneSoaRecordArgs

func ZoneSoaRecordPtr(v *ZoneSoaRecordArgs) ZoneSoaRecordPtrInput {
	return (*zoneSoaRecordPtrType)(v)
}

func (*zoneSoaRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneSoaRecord)(nil)).Elem()
}

func (i *zoneSoaRecordPtrType) ToZoneSoaRecordPtrOutput() ZoneSoaRecordPtrOutput {
	return i.ToZoneSoaRecordPtrOutputWithContext(context.Background())
}

func (i *zoneSoaRecordPtrType) ToZoneSoaRecordPtrOutputWithContext(ctx context.Context) ZoneSoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneSoaRecordPtrOutput)
}

type ZoneSoaRecordOutput struct{ *pulumi.OutputState }

func (ZoneSoaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneSoaRecord)(nil)).Elem()
}

func (o ZoneSoaRecordOutput) ToZoneSoaRecordOutput() ZoneSoaRecordOutput {
	return o
}

func (o ZoneSoaRecordOutput) ToZoneSoaRecordOutputWithContext(ctx context.Context) ZoneSoaRecordOutput {
	return o
}

func (o ZoneSoaRecordOutput) ToZoneSoaRecordPtrOutput() ZoneSoaRecordPtrOutput {
	return o.ToZoneSoaRecordPtrOutputWithContext(context.Background())
}

func (o ZoneSoaRecordOutput) ToZoneSoaRecordPtrOutputWithContext(ctx context.Context) ZoneSoaRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneSoaRecord) *ZoneSoaRecord {
		return &v
	}).(ZoneSoaRecordPtrOutput)
}

// The email contact for the SOA record.
func (o ZoneSoaRecordOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneSoaRecord) string { return v.Email }).(pulumi.StringOutput)
}

// The expire time for the SOA record. Defaults to `2419200`.
func (o ZoneSoaRecordOutput) ExpireTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.ExpireTime }).(pulumi.IntPtrOutput)
}

func (o ZoneSoaRecordOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
func (o ZoneSoaRecordOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneSoaRecord) string { return v.HostName }).(pulumi.StringOutput)
}

// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
func (o ZoneSoaRecordOutput) MinimumTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.MinimumTtl }).(pulumi.IntPtrOutput)
}

// The refresh time for the SOA record. Defaults to `3600`.
func (o ZoneSoaRecordOutput) RefreshTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.RefreshTime }).(pulumi.IntPtrOutput)
}

// The retry time for the SOA record. Defaults to `300`.
func (o ZoneSoaRecordOutput) RetryTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.RetryTime }).(pulumi.IntPtrOutput)
}

// The serial number for the SOA record. Defaults to `1`.
func (o ZoneSoaRecordOutput) SerialNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.SerialNumber }).(pulumi.IntPtrOutput)
}

// A mapping of tags to assign to the Record Set.
func (o ZoneSoaRecordOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ZoneSoaRecord) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The Time To Live of the SOA Record in seconds. Defaults to `3600`.
func (o ZoneSoaRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ZoneSoaRecord) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

type ZoneSoaRecordPtrOutput struct{ *pulumi.OutputState }

func (ZoneSoaRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneSoaRecord)(nil)).Elem()
}

func (o ZoneSoaRecordPtrOutput) ToZoneSoaRecordPtrOutput() ZoneSoaRecordPtrOutput {
	return o
}

func (o ZoneSoaRecordPtrOutput) ToZoneSoaRecordPtrOutputWithContext(ctx context.Context) ZoneSoaRecordPtrOutput {
	return o
}

func (o ZoneSoaRecordPtrOutput) Elem() ZoneSoaRecordOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) ZoneSoaRecord {
		if v != nil {
			return *v
		}
		var ret ZoneSoaRecord
		return ret
	}).(ZoneSoaRecordOutput)
}

// The email contact for the SOA record.
func (o ZoneSoaRecordPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

// The expire time for the SOA record. Defaults to `2419200`.
func (o ZoneSoaRecordPtrOutput) ExpireTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.IntPtrOutput)
}

func (o ZoneSoaRecordPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

// The domain name of the authoritative name server for the SOA record. Defaults to `ns1-03.azure-dns.com.`.
func (o ZoneSoaRecordPtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *string {
		if v == nil {
			return nil
		}
		return &v.HostName
	}).(pulumi.StringPtrOutput)
}

// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to `300`.
func (o ZoneSoaRecordPtrOutput) MinimumTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.IntPtrOutput)
}

// The refresh time for the SOA record. Defaults to `3600`.
func (o ZoneSoaRecordPtrOutput) RefreshTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.IntPtrOutput)
}

// The retry time for the SOA record. Defaults to `300`.
func (o ZoneSoaRecordPtrOutput) RetryTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.IntPtrOutput)
}

// The serial number for the SOA record. Defaults to `1`.
func (o ZoneSoaRecordPtrOutput) SerialNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.IntPtrOutput)
}

// A mapping of tags to assign to the Record Set.
func (o ZoneSoaRecordPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

// The Time To Live of the SOA Record in seconds. Defaults to `3600`.
func (o ZoneSoaRecordPtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneSoaRecord) *int {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CaaRecordRecordOutput{})
	pulumi.RegisterOutputType(CaaRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(MxRecordRecordOutput{})
	pulumi.RegisterOutputType(MxRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(ZoneSoaRecordOutput{})
	pulumi.RegisterOutputType(ZoneSoaRecordPtrOutput{})
}
