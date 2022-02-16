// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Logic App Standard instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := logicapps.LookupStandard(ctx, &logicapps.LookupStandardArgs{
// 			Name:              "logicappstd",
// 			ResourceGroupName: "example-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", example.Id)
// 		return nil
// 	})
// }
// ```
func LookupStandard(ctx *pulumi.Context, args *LookupStandardArgs, opts ...pulumi.InvokeOption) (*LookupStandardResult, error) {
	var rv LookupStandardResult
	err := ctx.Invoke("azure:logicapps/getStandard:getStandard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStandard.
type LookupStandardArgs struct {
	// The name of this Logic App.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Logic App exists.
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	SiteConfig        *GetStandardSiteConfig `pulumi:"siteConfig"`
	Tags              map[string]string      `pulumi:"tags"`
}

// A collection of values returned by getStandard.
type LookupStandardResult struct {
	AppServicePlanId           string                        `pulumi:"appServicePlanId"`
	AppSettings                map[string]string             `pulumi:"appSettings"`
	BundleVersion              string                        `pulumi:"bundleVersion"`
	ClientAffinityEnabled      bool                          `pulumi:"clientAffinityEnabled"`
	ClientCertificateMode      string                        `pulumi:"clientCertificateMode"`
	ConnectionStrings          []GetStandardConnectionString `pulumi:"connectionStrings"`
	CustomDomainVerificationId string                        `pulumi:"customDomainVerificationId"`
	DefaultHostname            string                        `pulumi:"defaultHostname"`
	Enabled                    bool                          `pulumi:"enabled"`
	HttpsOnly                  bool                          `pulumi:"httpsOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An `identity` block as defined below.
	Identities []GetStandardIdentity `pulumi:"identities"`
	Kind       string                `pulumi:"kind"`
	// The Azure location where the Logic App Standard exists.
	Location                    string                      `pulumi:"location"`
	Name                        string                      `pulumi:"name"`
	OutboundIpAddresses         string                      `pulumi:"outboundIpAddresses"`
	PossibleOutboundIpAddresses string                      `pulumi:"possibleOutboundIpAddresses"`
	ResourceGroupName           string                      `pulumi:"resourceGroupName"`
	SiteConfig                  GetStandardSiteConfig       `pulumi:"siteConfig"`
	SiteCredentials             []GetStandardSiteCredential `pulumi:"siteCredentials"`
	StorageAccountAccessKey     string                      `pulumi:"storageAccountAccessKey"`
	StorageAccountName          string                      `pulumi:"storageAccountName"`
	StorageAccountShareName     string                      `pulumi:"storageAccountShareName"`
	Tags                        map[string]string           `pulumi:"tags"`
	UseExtensionBundle          bool                        `pulumi:"useExtensionBundle"`
	Version                     string                      `pulumi:"version"`
}

func LookupStandardOutput(ctx *pulumi.Context, args LookupStandardOutputArgs, opts ...pulumi.InvokeOption) LookupStandardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStandardResult, error) {
			args := v.(LookupStandardArgs)
			r, err := LookupStandard(ctx, &args, opts...)
			return *r, err
		}).(LookupStandardResultOutput)
}

// A collection of arguments for invoking getStandard.
type LookupStandardOutputArgs struct {
	// The name of this Logic App.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Logic App exists.
	ResourceGroupName pulumi.StringInput            `pulumi:"resourceGroupName"`
	SiteConfig        GetStandardSiteConfigPtrInput `pulumi:"siteConfig"`
	Tags              pulumi.StringMapInput         `pulumi:"tags"`
}

func (LookupStandardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardArgs)(nil)).Elem()
}

// A collection of values returned by getStandard.
type LookupStandardResultOutput struct{ *pulumi.OutputState }

func (LookupStandardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardResult)(nil)).Elem()
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutput() LookupStandardResultOutput {
	return o
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutputWithContext(ctx context.Context) LookupStandardResultOutput {
	return o
}

func (o LookupStandardResultOutput) AppServicePlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.AppServicePlanId }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) AppSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStandardResult) map[string]string { return v.AppSettings }).(pulumi.StringMapOutput)
}

func (o LookupStandardResultOutput) BundleVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.BundleVersion }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) ClientAffinityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStandardResult) bool { return v.ClientAffinityEnabled }).(pulumi.BoolOutput)
}

func (o LookupStandardResultOutput) ClientCertificateMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.ClientCertificateMode }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) ConnectionStrings() GetStandardConnectionStringArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []GetStandardConnectionString { return v.ConnectionStrings }).(GetStandardConnectionStringArrayOutput)
}

func (o LookupStandardResultOutput) CustomDomainVerificationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.CustomDomainVerificationId }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.DefaultHostname }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStandardResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupStandardResultOutput) HttpsOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStandardResult) bool { return v.HttpsOnly }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStandardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Id }).(pulumi.StringOutput)
}

// An `identity` block as defined below.
func (o LookupStandardResultOutput) Identities() GetStandardIdentityArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []GetStandardIdentity { return v.Identities }).(GetStandardIdentityArrayOutput)
}

func (o LookupStandardResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The Azure location where the Logic App Standard exists.
func (o LookupStandardResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) SiteConfig() GetStandardSiteConfigOutput {
	return o.ApplyT(func(v LookupStandardResult) GetStandardSiteConfig { return v.SiteConfig }).(GetStandardSiteConfigOutput)
}

func (o LookupStandardResultOutput) SiteCredentials() GetStandardSiteCredentialArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []GetStandardSiteCredential { return v.SiteCredentials }).(GetStandardSiteCredentialArrayOutput)
}

func (o LookupStandardResultOutput) StorageAccountAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.StorageAccountAccessKey }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) StorageAccountShareName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.StorageAccountShareName }).(pulumi.StringOutput)
}

func (o LookupStandardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStandardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStandardResultOutput) UseExtensionBundle() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStandardResult) bool { return v.UseExtensionBundle }).(pulumi.BoolOutput)
}

func (o LookupStandardResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStandardResultOutput{})
}
