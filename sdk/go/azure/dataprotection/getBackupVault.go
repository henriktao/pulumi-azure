// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Backup Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/dataprotection"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := dataprotection.LookupBackupVault(ctx, &dataprotection.LookupBackupVaultArgs{
// 			Name:              "existing-backup-vault",
// 			ResourceGroupName: "existing-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("azurermDataProtectionBackupVaultId", data.Azurerm_vpn_gateway.Example.Id)
// 		ctx.Export("azurermDataProtectionBackupVaultPrincipalId", example.Identities[0].PrincipalId)
// 		return nil
// 	})
// }
// ```
func LookupBackupVault(ctx *pulumi.Context, args *LookupBackupVaultArgs, opts ...pulumi.InvokeOption) (*LookupBackupVaultResult, error) {
	var rv LookupBackupVaultResult
	err := ctx.Invoke("azure:dataprotection/getBackupVault:getBackupVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupVault.
type LookupBackupVaultArgs struct {
	// Specifies the name of the Backup Vault.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Backup Vault exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getBackupVault.
type LookupBackupVaultResult struct {
	// Specifies the type of the data store.
	DatastoreType string `pulumi:"datastoreType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A `identity` block as defined below.
	Identities []GetBackupVaultIdentity `pulumi:"identities"`
	// The Azure Region where the Backup Vault exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// Specifies the backup storage redundancy.
	Redundancy        string `pulumi:"redundancy"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which are assigned to the Backup Vault.
	Tags map[string]string `pulumi:"tags"`
}

func LookupBackupVaultOutput(ctx *pulumi.Context, args LookupBackupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupBackupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupVaultResult, error) {
			args := v.(LookupBackupVaultArgs)
			r, err := LookupBackupVault(ctx, &args, opts...)
			return *r, err
		}).(LookupBackupVaultResultOutput)
}

// A collection of arguments for invoking getBackupVault.
type LookupBackupVaultOutputArgs struct {
	// Specifies the name of the Backup Vault.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Resource Group where the Backup Vault exists.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultArgs)(nil)).Elem()
}

// A collection of values returned by getBackupVault.
type LookupBackupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupBackupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultResult)(nil)).Elem()
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutput() LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutputWithContext(ctx context.Context) LookupBackupVaultResultOutput {
	return o
}

// Specifies the type of the data store.
func (o LookupBackupVaultResultOutput) DatastoreType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.DatastoreType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBackupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// A `identity` block as defined below.
func (o LookupBackupVaultResultOutput) Identities() GetBackupVaultIdentityArrayOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) []GetBackupVaultIdentity { return v.Identities }).(GetBackupVaultIdentityArrayOutput)
}

// The Azure Region where the Backup Vault exists.
func (o LookupBackupVaultResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBackupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the backup storage redundancy.
func (o LookupBackupVaultResultOutput) Redundancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Redundancy }).(pulumi.StringOutput)
}

func (o LookupBackupVaultResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A mapping of tags which are assigned to the Backup Vault.
func (o LookupBackupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupVaultResultOutput{})
}
