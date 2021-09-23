// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a list of secret names from an existing Key Vault Secret.
func GetSecrets(ctx *pulumi.Context, args *GetSecretsArgs, opts ...pulumi.InvokeOption) (*GetSecretsResult, error) {
	var rv GetSecretsResult
	err := ctx.Invoke("azure:keyvault/getSecrets:getSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecrets.
type GetSecretsArgs struct {
	// Specifies the ID of the Key Vault instance to fetch secret names from, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId string `pulumi:"keyVaultId"`
}

// A collection of values returned by getSecrets.
type GetSecretsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Key Vault ID.
	KeyVaultId string `pulumi:"keyVaultId"`
	// List containing names of secrets that exist in this Key Vault.
	Names []string `pulumi:"names"`
}

func GetSecretsOutput(ctx *pulumi.Context, args GetSecretsOutputArgs, opts ...pulumi.InvokeOption) GetSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretsResult, error) {
			args := v.(GetSecretsArgs)
			r, err := GetSecrets(ctx, &args, opts...)
			return *r, err
		}).(GetSecretsResultOutput)
}

// A collection of arguments for invoking getSecrets.
type GetSecretsOutputArgs struct {
	// Specifies the ID of the Key Vault instance to fetch secret names from, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId pulumi.StringInput `pulumi:"keyVaultId"`
}

func (GetSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsArgs)(nil)).Elem()
}

// A collection of values returned by getSecrets.
type GetSecretsResultOutput struct{ *pulumi.OutputState }

func (GetSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsResult)(nil)).Elem()
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutput() GetSecretsResultOutput {
	return o
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutputWithContext(ctx context.Context) GetSecretsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Key Vault ID.
func (o GetSecretsResultOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsResult) string { return v.KeyVaultId }).(pulumi.StringOutput)
}

// List containing names of secrets that exist in this Key Vault.
func (o GetSecretsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretsResultOutput{})
}
