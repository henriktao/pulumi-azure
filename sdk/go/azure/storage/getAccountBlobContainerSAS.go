// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.
//
// Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		storage, err := storage.NewAccount(ctx, "storage", &storage.AccountArgs{
// 			ResourceGroupName:      rg.Name,
// 			Location:               rg.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		container, err := storage.NewContainer(ctx, "container", &storage.ContainerArgs{
// 			StorageAccountName:  storage.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example := storage.GetAccountBlobContainerSASOutput(ctx, storage.GetAccountBlobContainerSASOutputArgs{
// 			ConnectionString: storage.PrimaryConnectionString,
// 			ContainerName:    container.Name,
// 			HttpsOnly:        pulumi.Bool(true),
// 			IpAddress:        pulumi.String("168.1.5.65"),
// 			Start:            pulumi.String("2018-03-21"),
// 			Expiry:           pulumi.String("2018-03-21"),
// 			Permissions: &storage.GetAccountBlobContainerSASPermissionsArgs{
// 				Read:   pulumi.Bool(true),
// 				Add:    pulumi.Bool(true),
// 				Create: pulumi.Bool(false),
// 				Write:  pulumi.Bool(false),
// 				Delete: pulumi.Bool(true),
// 				List:   pulumi.Bool(true),
// 			},
// 			CacheControl:       pulumi.String("max-age=5"),
// 			ContentDisposition: pulumi.String("inline"),
// 			ContentEncoding:    pulumi.String("deflate"),
// 			ContentLanguage:    pulumi.String("en-US"),
// 			ContentType:        pulumi.String("application/json"),
// 		}, nil)
// 		ctx.Export("sasUrlQueryString", example.ApplyT(func(example storage.GetAccountBlobContainerSASResult) (string, error) {
// 			return example.Sas, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetAccountBlobContainerSAS(ctx *pulumi.Context, args *GetAccountBlobContainerSASArgs, opts ...pulumi.InvokeOption) (*GetAccountBlobContainerSASResult, error) {
	var rv GetAccountBlobContainerSASResult
	err := ctx.Invoke("azure:storage/getAccountBlobContainerSAS:getAccountBlobContainerSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountBlobContainerSAS.
type GetAccountBlobContainerSASArgs struct {
	// The `Cache-Control` response header that is sent when this SAS token is used.
	CacheControl *string `pulumi:"cacheControl"`
	// The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of an `storage.Account` resource.
	ConnectionString string `pulumi:"connectionString"`
	// Name of the container.
	ContainerName string `pulumi:"containerName"`
	// The `Content-Disposition` response header that is sent when this SAS token is used.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// The `Content-Encoding` response header that is sent when this SAS token is used.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The `Content-Language` response header that is sent when this SAS token is used.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// The `Content-Type` response header that is sent when this SAS token is used.
	ContentType *string `pulumi:"contentType"`
	// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
	Expiry string `pulumi:"expiry"`
	// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Single ipv4 address or range (connected with a dash) of ipv4 addresses.
	IpAddress *string `pulumi:"ipAddress"`
	// A `permissions` block as defined below.
	Permissions GetAccountBlobContainerSASPermissions `pulumi:"permissions"`
	// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
	Start string `pulumi:"start"`
}

// A collection of values returned by getAccountBlobContainerSAS.
type GetAccountBlobContainerSASResult struct {
	CacheControl       *string `pulumi:"cacheControl"`
	ConnectionString   string  `pulumi:"connectionString"`
	ContainerName      string  `pulumi:"containerName"`
	ContentDisposition *string `pulumi:"contentDisposition"`
	ContentEncoding    *string `pulumi:"contentEncoding"`
	ContentLanguage    *string `pulumi:"contentLanguage"`
	ContentType        *string `pulumi:"contentType"`
	Expiry             string  `pulumi:"expiry"`
	HttpsOnly          *bool   `pulumi:"httpsOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                                `pulumi:"id"`
	IpAddress   *string                               `pulumi:"ipAddress"`
	Permissions GetAccountBlobContainerSASPermissions `pulumi:"permissions"`
	// The computed Blob Container Shared Access Signature (SAS).
	Sas   string `pulumi:"sas"`
	Start string `pulumi:"start"`
}

func GetAccountBlobContainerSASOutput(ctx *pulumi.Context, args GetAccountBlobContainerSASOutputArgs, opts ...pulumi.InvokeOption) GetAccountBlobContainerSASResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountBlobContainerSASResult, error) {
			args := v.(GetAccountBlobContainerSASArgs)
			r, err := GetAccountBlobContainerSAS(ctx, &args, opts...)
			return *r, err
		}).(GetAccountBlobContainerSASResultOutput)
}

// A collection of arguments for invoking getAccountBlobContainerSAS.
type GetAccountBlobContainerSASOutputArgs struct {
	// The `Cache-Control` response header that is sent when this SAS token is used.
	CacheControl pulumi.StringPtrInput `pulumi:"cacheControl"`
	// The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of an `storage.Account` resource.
	ConnectionString pulumi.StringInput `pulumi:"connectionString"`
	// Name of the container.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The `Content-Disposition` response header that is sent when this SAS token is used.
	ContentDisposition pulumi.StringPtrInput `pulumi:"contentDisposition"`
	// The `Content-Encoding` response header that is sent when this SAS token is used.
	ContentEncoding pulumi.StringPtrInput `pulumi:"contentEncoding"`
	// The `Content-Language` response header that is sent when this SAS token is used.
	ContentLanguage pulumi.StringPtrInput `pulumi:"contentLanguage"`
	// The `Content-Type` response header that is sent when this SAS token is used.
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	// The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
	Expiry pulumi.StringInput `pulumi:"expiry"`
	// Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
	HttpsOnly pulumi.BoolPtrInput `pulumi:"httpsOnly"`
	// Single ipv4 address or range (connected with a dash) of ipv4 addresses.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// A `permissions` block as defined below.
	Permissions GetAccountBlobContainerSASPermissionsInput `pulumi:"permissions"`
	// The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
	Start pulumi.StringInput `pulumi:"start"`
}

func (GetAccountBlobContainerSASOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountBlobContainerSASArgs)(nil)).Elem()
}

// A collection of values returned by getAccountBlobContainerSAS.
type GetAccountBlobContainerSASResultOutput struct{ *pulumi.OutputState }

func (GetAccountBlobContainerSASResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountBlobContainerSASResult)(nil)).Elem()
}

func (o GetAccountBlobContainerSASResultOutput) ToGetAccountBlobContainerSASResultOutput() GetAccountBlobContainerSASResultOutput {
	return o
}

func (o GetAccountBlobContainerSASResultOutput) ToGetAccountBlobContainerSASResultOutputWithContext(ctx context.Context) GetAccountBlobContainerSASResultOutput {
	return o
}

func (o GetAccountBlobContainerSASResultOutput) CacheControl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.CacheControl }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ContentDisposition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.ContentDisposition }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ContentLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.ContentLanguage }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o GetAccountBlobContainerSASResultOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountBlobContainerSASResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountBlobContainerSASResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o GetAccountBlobContainerSASResultOutput) Permissions() GetAccountBlobContainerSASPermissionsOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) GetAccountBlobContainerSASPermissions { return v.Permissions }).(GetAccountBlobContainerSASPermissionsOutput)
}

// The computed Blob Container Shared Access Signature (SAS).
func (o GetAccountBlobContainerSASResultOutput) Sas() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.Sas }).(pulumi.StringOutput)
}

func (o GetAccountBlobContainerSASResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountBlobContainerSASResult) string { return v.Start }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountBlobContainerSASResultOutput{})
}
