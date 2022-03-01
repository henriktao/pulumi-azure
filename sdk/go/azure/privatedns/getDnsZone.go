// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Private DNS Zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/privatedns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := privatedns.GetDnsZone(ctx, &privatedns.GetDnsZoneArgs{
// 			Name:              "contoso.internal",
// 			ResourceGroupName: pulumi.StringRef("contoso-dns"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("privateDnsZoneId", example.Id)
// 		return nil
// 	})
// }
// ```
func GetDnsZone(ctx *pulumi.Context, args *GetDnsZoneArgs, opts ...pulumi.InvokeOption) (*GetDnsZoneResult, error) {
	var rv GetDnsZoneResult
	err := ctx.Invoke("azure:privatedns/getDnsZone:getDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsZone.
type GetDnsZoneArgs struct {
	// The name of the Private DNS Zone.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the Private DNS Zone exists.
	// If the Name of the Resource Group is not provided, the first Private DNS Zone from the list of Private
	// DNS Zones in your subscription that matches `name` will be returned.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getDnsZone.
type GetDnsZoneResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of recordsets that can be created in this Private Zone.
	MaxNumberOfRecordSets int `pulumi:"maxNumberOfRecordSets"`
	// Maximum number of Virtual Networks that can be linked to this Private Zone.
	MaxNumberOfVirtualNetworkLinks int `pulumi:"maxNumberOfVirtualNetworkLinks"`
	// Maximum number of Virtual Networks that can be linked to this Private Zone with registration enabled.
	MaxNumberOfVirtualNetworkLinksWithRegistration int    `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	Name                                           string `pulumi:"name"`
	// The number of recordsets currently in the zone.
	NumberOfRecordSets int    `pulumi:"numberOfRecordSets"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	// A mapping of tags for the zone.
	Tags map[string]string `pulumi:"tags"`
}

func GetDnsZoneOutput(ctx *pulumi.Context, args GetDnsZoneOutputArgs, opts ...pulumi.InvokeOption) GetDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsZoneResult, error) {
			args := v.(GetDnsZoneArgs)
			r, err := GetDnsZone(ctx, &args, opts...)
			return *r, err
		}).(GetDnsZoneResultOutput)
}

// A collection of arguments for invoking getDnsZone.
type GetDnsZoneOutputArgs struct {
	// The name of the Private DNS Zone.
	Name pulumi.StringInput `pulumi:"name"`
	// The Name of the Resource Group where the Private DNS Zone exists.
	// If the Name of the Resource Group is not provided, the first Private DNS Zone from the list of Private
	// DNS Zones in your subscription that matches `name` will be returned.
	ResourceGroupName pulumi.StringPtrInput `pulumi:"resourceGroupName"`
}

func (GetDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDnsZone.
type GetDnsZoneResultOutput struct{ *pulumi.OutputState }

func (GetDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsZoneResult)(nil)).Elem()
}

func (o GetDnsZoneResultOutput) ToGetDnsZoneResultOutput() GetDnsZoneResultOutput {
	return o
}

func (o GetDnsZoneResultOutput) ToGetDnsZoneResultOutputWithContext(ctx context.Context) GetDnsZoneResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of recordsets that can be created in this Private Zone.
func (o GetDnsZoneResultOutput) MaxNumberOfRecordSets() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.MaxNumberOfRecordSets }).(pulumi.IntOutput)
}

// Maximum number of Virtual Networks that can be linked to this Private Zone.
func (o GetDnsZoneResultOutput) MaxNumberOfVirtualNetworkLinks() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.MaxNumberOfVirtualNetworkLinks }).(pulumi.IntOutput)
}

// Maximum number of Virtual Networks that can be linked to this Private Zone with registration enabled.
func (o GetDnsZoneResultOutput) MaxNumberOfVirtualNetworkLinksWithRegistration() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.MaxNumberOfVirtualNetworkLinksWithRegistration }).(pulumi.IntOutput)
}

func (o GetDnsZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of recordsets currently in the zone.
func (o GetDnsZoneResultOutput) NumberOfRecordSets() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.NumberOfRecordSets }).(pulumi.IntOutput)
}

func (o GetDnsZoneResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A mapping of tags for the zone.
func (o GetDnsZoneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDnsZoneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsZoneResultOutput{})
}
