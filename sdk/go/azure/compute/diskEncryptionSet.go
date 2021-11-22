// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Disk Encryption Set.
//
// > **NOTE:** At this time the Key Vault used to store the Active Key for this Disk Encryption Set must have both Soft Delete & Purge Protection enabled - which are not yet supported by this provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:                 exampleResourceGroup.Location,
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			TenantId:                 pulumi.String(current.TenantId),
// 			SkuName:                  pulumi.String("premium"),
// 			EnabledForDiskEncryption: pulumi.Bool(true),
// 			SoftDeleteEnabled:        pulumi.Bool(true),
// 			PurgeProtectionEnabled:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewAccessPolicy(ctx, "example_user", &keyvault.AccessPolicyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			TenantId:   pulumi.String(current.TenantId),
// 			ObjectId:   pulumi.String(current.ObjectId),
// 			KeyPermissions: pulumi.StringArray{
// 				pulumi.String("get"),
// 				pulumi.String("create"),
// 				pulumi.String("delete"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKey, err := keyvault.NewKey(ctx, "exampleKey", &keyvault.KeyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			KeyType:    pulumi.String("RSA"),
// 			KeySize:    pulumi.Int(2048),
// 			KeyOpts: pulumi.StringArray{
// 				pulumi.String("decrypt"),
// 				pulumi.String("encrypt"),
// 				pulumi.String("sign"),
// 				pulumi.String("unwrapKey"),
// 				pulumi.String("verify"),
// 				pulumi.String("wrapKey"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			example_user,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		exampleDiskEncryptionSet, err := compute.NewDiskEncryptionSet(ctx, "exampleDiskEncryptionSet", &compute.DiskEncryptionSetArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			KeyVaultKeyId:     exampleKey.ID(),
// 			Identity: &compute.DiskEncryptionSetIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.NewAccessPolicy(ctx, "example_disk", &keyvault.AccessPolicyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			TenantId: exampleDiskEncryptionSet.Identity.ApplyT(func(identity compute.DiskEncryptionSetIdentity) (string, error) {
// 				return identity.TenantId, nil
// 			}).(pulumi.StringOutput),
// 			ObjectId: exampleDiskEncryptionSet.Identity.ApplyT(func(identity compute.DiskEncryptionSetIdentity) (string, error) {
// 				return identity.PrincipalId, nil
// 			}).(pulumi.StringOutput),
// 			KeyPermissions: pulumi.StringArray{
// 				pulumi.String("Get"),
// 				pulumi.String("WrapKey"),
// 				pulumi.String("UnwrapKey"),
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
// Disk Encryption Sets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:compute/diskEncryptionSet:DiskEncryptionSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
// ```
type DiskEncryptionSet struct {
	pulumi.CustomResourceState

	// Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
	AutoKeyRotationEnabled pulumi.BoolPtrOutput `pulumi:"autoKeyRotationEnabled"`
	// The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// An `identity` block as defined below.
	Identity DiskEncryptionSetIdentityOutput `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringOutput `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDiskEncryptionSet registers a new resource with the given unique name, arguments, and options.
func NewDiskEncryptionSet(ctx *pulumi.Context,
	name string, args *DiskEncryptionSetArgs, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.KeyVaultKeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultKeyId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource DiskEncryptionSet
	err := ctx.RegisterResource("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskEncryptionSet gets an existing DiskEncryptionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskEncryptionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskEncryptionSetState, opts ...pulumi.ResourceOption) (*DiskEncryptionSet, error) {
	var resource DiskEncryptionSet
	err := ctx.ReadResource("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskEncryptionSet resources.
type diskEncryptionSetState struct {
	// Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
	AutoKeyRotationEnabled *bool `pulumi:"autoKeyRotationEnabled"`
	// The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
	EncryptionType *string `pulumi:"encryptionType"`
	// An `identity` block as defined below.
	Identity *DiskEncryptionSetIdentity `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId *string `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags map[string]string `pulumi:"tags"`
}

type DiskEncryptionSetState struct {
	// Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
	AutoKeyRotationEnabled pulumi.BoolPtrInput
	// The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
	EncryptionType pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity DiskEncryptionSetIdentityPtrInput
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringPtrInput
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapInput
}

func (DiskEncryptionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetState)(nil)).Elem()
}

type diskEncryptionSetArgs struct {
	// Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
	AutoKeyRotationEnabled *bool `pulumi:"autoKeyRotationEnabled"`
	// The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
	EncryptionType *string `pulumi:"encryptionType"`
	// An `identity` block as defined below.
	Identity DiskEncryptionSetIdentity `pulumi:"identity"`
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId string `pulumi:"keyVaultKeyId"`
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskEncryptionSet resource.
type DiskEncryptionSetArgs struct {
	// Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
	AutoKeyRotationEnabled pulumi.BoolPtrInput
	// The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
	EncryptionType pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity DiskEncryptionSetIdentityInput
	// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
	KeyVaultKeyId pulumi.StringInput
	// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Disk Encryption Set.
	Tags pulumi.StringMapInput
}

func (DiskEncryptionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskEncryptionSetArgs)(nil)).Elem()
}

type DiskEncryptionSetInput interface {
	pulumi.Input

	ToDiskEncryptionSetOutput() DiskEncryptionSetOutput
	ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput
}

func (*DiskEncryptionSet) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSet)(nil))
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return i.ToDiskEncryptionSetOutputWithContext(context.Background())
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetOutput)
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return i.ToDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (i *DiskEncryptionSet) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetPtrOutput)
}

type DiskEncryptionSetPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput
	ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput
}

type diskEncryptionSetPtrType DiskEncryptionSetArgs

func (*diskEncryptionSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil))
}

func (i *diskEncryptionSetPtrType) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return i.ToDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionSetPtrType) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetPtrOutput)
}

// DiskEncryptionSetArrayInput is an input type that accepts DiskEncryptionSetArray and DiskEncryptionSetArrayOutput values.
// You can construct a concrete instance of `DiskEncryptionSetArrayInput` via:
//
//          DiskEncryptionSetArray{ DiskEncryptionSetArgs{...} }
type DiskEncryptionSetArrayInput interface {
	pulumi.Input

	ToDiskEncryptionSetArrayOutput() DiskEncryptionSetArrayOutput
	ToDiskEncryptionSetArrayOutputWithContext(context.Context) DiskEncryptionSetArrayOutput
}

type DiskEncryptionSetArray []DiskEncryptionSetInput

func (DiskEncryptionSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskEncryptionSet)(nil)).Elem()
}

func (i DiskEncryptionSetArray) ToDiskEncryptionSetArrayOutput() DiskEncryptionSetArrayOutput {
	return i.ToDiskEncryptionSetArrayOutputWithContext(context.Background())
}

func (i DiskEncryptionSetArray) ToDiskEncryptionSetArrayOutputWithContext(ctx context.Context) DiskEncryptionSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetArrayOutput)
}

// DiskEncryptionSetMapInput is an input type that accepts DiskEncryptionSetMap and DiskEncryptionSetMapOutput values.
// You can construct a concrete instance of `DiskEncryptionSetMapInput` via:
//
//          DiskEncryptionSetMap{ "key": DiskEncryptionSetArgs{...} }
type DiskEncryptionSetMapInput interface {
	pulumi.Input

	ToDiskEncryptionSetMapOutput() DiskEncryptionSetMapOutput
	ToDiskEncryptionSetMapOutputWithContext(context.Context) DiskEncryptionSetMapOutput
}

type DiskEncryptionSetMap map[string]DiskEncryptionSetInput

func (DiskEncryptionSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskEncryptionSet)(nil)).Elem()
}

func (i DiskEncryptionSetMap) ToDiskEncryptionSetMapOutput() DiskEncryptionSetMapOutput {
	return i.ToDiskEncryptionSetMapOutputWithContext(context.Background())
}

func (i DiskEncryptionSetMap) ToDiskEncryptionSetMapOutputWithContext(ctx context.Context) DiskEncryptionSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetMapOutput)
}

type DiskEncryptionSetOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSet)(nil))
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutput() DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetOutputWithContext(ctx context.Context) DiskEncryptionSetOutput {
	return o
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return o.ToDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetOutput) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSet) *DiskEncryptionSet {
		return &v
	}).(DiskEncryptionSetPtrOutput)
}

type DiskEncryptionSetPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSet)(nil))
}

func (o DiskEncryptionSetPtrOutput) ToDiskEncryptionSetPtrOutput() DiskEncryptionSetPtrOutput {
	return o
}

func (o DiskEncryptionSetPtrOutput) ToDiskEncryptionSetPtrOutputWithContext(ctx context.Context) DiskEncryptionSetPtrOutput {
	return o
}

func (o DiskEncryptionSetPtrOutput) Elem() DiskEncryptionSetOutput {
	return o.ApplyT(func(v *DiskEncryptionSet) DiskEncryptionSet {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSet
		return ret
	}).(DiskEncryptionSetOutput)
}

type DiskEncryptionSetArrayOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskEncryptionSet)(nil))
}

func (o DiskEncryptionSetArrayOutput) ToDiskEncryptionSetArrayOutput() DiskEncryptionSetArrayOutput {
	return o
}

func (o DiskEncryptionSetArrayOutput) ToDiskEncryptionSetArrayOutputWithContext(ctx context.Context) DiskEncryptionSetArrayOutput {
	return o
}

func (o DiskEncryptionSetArrayOutput) Index(i pulumi.IntInput) DiskEncryptionSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskEncryptionSet {
		return vs[0].([]DiskEncryptionSet)[vs[1].(int)]
	}).(DiskEncryptionSetOutput)
}

type DiskEncryptionSetMapOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DiskEncryptionSet)(nil))
}

func (o DiskEncryptionSetMapOutput) ToDiskEncryptionSetMapOutput() DiskEncryptionSetMapOutput {
	return o
}

func (o DiskEncryptionSetMapOutput) ToDiskEncryptionSetMapOutputWithContext(ctx context.Context) DiskEncryptionSetMapOutput {
	return o
}

func (o DiskEncryptionSetMapOutput) MapIndex(k pulumi.StringInput) DiskEncryptionSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DiskEncryptionSet {
		return vs[0].(map[string]DiskEncryptionSet)[vs[1].(string)]
	}).(DiskEncryptionSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetInput)(nil)).Elem(), &DiskEncryptionSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetPtrInput)(nil)).Elem(), &DiskEncryptionSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetArrayInput)(nil)).Elem(), DiskEncryptionSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetMapInput)(nil)).Elem(), DiskEncryptionSetMap{})
	pulumi.RegisterOutputType(DiskEncryptionSetOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetArrayOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetMapOutput{})
}
