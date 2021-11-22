// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Customer Managed Key for a EventHub Namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
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
// 		exampleCluster, err := eventhub.NewCluster(ctx, "exampleCluster", &eventhub.ClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("Dedicated_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:           exampleResourceGroup.Location,
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			Sku:                pulumi.String("Standard"),
// 			DedicatedClusterId: exampleCluster.ID(),
// 			Identity: &eventhub.EventHubNamespaceIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			TenantId:               pulumi.String(current.TenantId),
// 			SkuName:                pulumi.String("standard"),
// 			SoftDeleteEnabled:      pulumi.Bool(true),
// 			PurgeProtectionEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccessPolicy, err := keyvault.NewAccessPolicy(ctx, "exampleAccessPolicy", &keyvault.AccessPolicyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			TenantId: exampleEventHubNamespace.Identity.ApplyT(func(identity eventhub.EventHubNamespaceIdentity) (string, error) {
// 				return identity.TenantId, nil
// 			}).(pulumi.StringOutput),
// 			ObjectId: exampleEventHubNamespace.Identity.ApplyT(func(identity eventhub.EventHubNamespaceIdentity) (string, error) {
// 				return identity.PrincipalId, nil
// 			}).(pulumi.StringOutput),
// 			KeyPermissions: pulumi.StringArray{
// 				pulumi.String("get"),
// 				pulumi.String("unwrapkey"),
// 				pulumi.String("wrapkey"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example2, err := keyvault.NewAccessPolicy(ctx, "example2", &keyvault.AccessPolicyArgs{
// 			KeyVaultId: exampleKeyVault.ID(),
// 			TenantId:   pulumi.String(current.TenantId),
// 			ObjectId:   pulumi.String(current.ObjectId),
// 			KeyPermissions: pulumi.StringArray{
// 				pulumi.String("create"),
// 				pulumi.String("delete"),
// 				pulumi.String("get"),
// 				pulumi.String("list"),
// 				pulumi.String("purge"),
// 				pulumi.String("recover"),
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
// 			exampleAccessPolicy,
// 			example2,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewNamespaceCustomerManagedKey(ctx, "exampleNamespaceCustomerManagedKey", &eventhub.NamespaceCustomerManagedKeyArgs{
// 			EventhubNamespaceId: exampleEventHubNamespace.ID(),
// 			KeyVaultKeyIds: pulumi.StringArray{
// 				exampleKey.ID(),
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
// Customer Managed Keys for a EventHub Namespace can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventhub/namespaceCustomerManagedKey:NamespaceCustomerManagedKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1
// ```
type NamespaceCustomerManagedKey struct {
	pulumi.CustomResourceState

	// The ID of the EventHub Namespace. Changing this forces a new resource to be created.
	EventhubNamespaceId pulumi.StringOutput `pulumi:"eventhubNamespaceId"`
	// The list of keys of Key Vault.
	KeyVaultKeyIds pulumi.StringArrayOutput `pulumi:"keyVaultKeyIds"`
}

// NewNamespaceCustomerManagedKey registers a new resource with the given unique name, arguments, and options.
func NewNamespaceCustomerManagedKey(ctx *pulumi.Context,
	name string, args *NamespaceCustomerManagedKeyArgs, opts ...pulumi.ResourceOption) (*NamespaceCustomerManagedKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventhubNamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'EventhubNamespaceId'")
	}
	if args.KeyVaultKeyIds == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultKeyIds'")
	}
	var resource NamespaceCustomerManagedKey
	err := ctx.RegisterResource("azure:eventhub/namespaceCustomerManagedKey:NamespaceCustomerManagedKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceCustomerManagedKey gets an existing NamespaceCustomerManagedKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceCustomerManagedKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceCustomerManagedKeyState, opts ...pulumi.ResourceOption) (*NamespaceCustomerManagedKey, error) {
	var resource NamespaceCustomerManagedKey
	err := ctx.ReadResource("azure:eventhub/namespaceCustomerManagedKey:NamespaceCustomerManagedKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceCustomerManagedKey resources.
type namespaceCustomerManagedKeyState struct {
	// The ID of the EventHub Namespace. Changing this forces a new resource to be created.
	EventhubNamespaceId *string `pulumi:"eventhubNamespaceId"`
	// The list of keys of Key Vault.
	KeyVaultKeyIds []string `pulumi:"keyVaultKeyIds"`
}

type NamespaceCustomerManagedKeyState struct {
	// The ID of the EventHub Namespace. Changing this forces a new resource to be created.
	EventhubNamespaceId pulumi.StringPtrInput
	// The list of keys of Key Vault.
	KeyVaultKeyIds pulumi.StringArrayInput
}

func (NamespaceCustomerManagedKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceCustomerManagedKeyState)(nil)).Elem()
}

type namespaceCustomerManagedKeyArgs struct {
	// The ID of the EventHub Namespace. Changing this forces a new resource to be created.
	EventhubNamespaceId string `pulumi:"eventhubNamespaceId"`
	// The list of keys of Key Vault.
	KeyVaultKeyIds []string `pulumi:"keyVaultKeyIds"`
}

// The set of arguments for constructing a NamespaceCustomerManagedKey resource.
type NamespaceCustomerManagedKeyArgs struct {
	// The ID of the EventHub Namespace. Changing this forces a new resource to be created.
	EventhubNamespaceId pulumi.StringInput
	// The list of keys of Key Vault.
	KeyVaultKeyIds pulumi.StringArrayInput
}

func (NamespaceCustomerManagedKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceCustomerManagedKeyArgs)(nil)).Elem()
}

type NamespaceCustomerManagedKeyInput interface {
	pulumi.Input

	ToNamespaceCustomerManagedKeyOutput() NamespaceCustomerManagedKeyOutput
	ToNamespaceCustomerManagedKeyOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyOutput
}

func (*NamespaceCustomerManagedKey) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceCustomerManagedKey)(nil))
}

func (i *NamespaceCustomerManagedKey) ToNamespaceCustomerManagedKeyOutput() NamespaceCustomerManagedKeyOutput {
	return i.ToNamespaceCustomerManagedKeyOutputWithContext(context.Background())
}

func (i *NamespaceCustomerManagedKey) ToNamespaceCustomerManagedKeyOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceCustomerManagedKeyOutput)
}

func (i *NamespaceCustomerManagedKey) ToNamespaceCustomerManagedKeyPtrOutput() NamespaceCustomerManagedKeyPtrOutput {
	return i.ToNamespaceCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (i *NamespaceCustomerManagedKey) ToNamespaceCustomerManagedKeyPtrOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceCustomerManagedKeyPtrOutput)
}

type NamespaceCustomerManagedKeyPtrInput interface {
	pulumi.Input

	ToNamespaceCustomerManagedKeyPtrOutput() NamespaceCustomerManagedKeyPtrOutput
	ToNamespaceCustomerManagedKeyPtrOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyPtrOutput
}

type namespaceCustomerManagedKeyPtrType NamespaceCustomerManagedKeyArgs

func (*namespaceCustomerManagedKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceCustomerManagedKey)(nil))
}

func (i *namespaceCustomerManagedKeyPtrType) ToNamespaceCustomerManagedKeyPtrOutput() NamespaceCustomerManagedKeyPtrOutput {
	return i.ToNamespaceCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (i *namespaceCustomerManagedKeyPtrType) ToNamespaceCustomerManagedKeyPtrOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceCustomerManagedKeyPtrOutput)
}

// NamespaceCustomerManagedKeyArrayInput is an input type that accepts NamespaceCustomerManagedKeyArray and NamespaceCustomerManagedKeyArrayOutput values.
// You can construct a concrete instance of `NamespaceCustomerManagedKeyArrayInput` via:
//
//          NamespaceCustomerManagedKeyArray{ NamespaceCustomerManagedKeyArgs{...} }
type NamespaceCustomerManagedKeyArrayInput interface {
	pulumi.Input

	ToNamespaceCustomerManagedKeyArrayOutput() NamespaceCustomerManagedKeyArrayOutput
	ToNamespaceCustomerManagedKeyArrayOutputWithContext(context.Context) NamespaceCustomerManagedKeyArrayOutput
}

type NamespaceCustomerManagedKeyArray []NamespaceCustomerManagedKeyInput

func (NamespaceCustomerManagedKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceCustomerManagedKey)(nil)).Elem()
}

func (i NamespaceCustomerManagedKeyArray) ToNamespaceCustomerManagedKeyArrayOutput() NamespaceCustomerManagedKeyArrayOutput {
	return i.ToNamespaceCustomerManagedKeyArrayOutputWithContext(context.Background())
}

func (i NamespaceCustomerManagedKeyArray) ToNamespaceCustomerManagedKeyArrayOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceCustomerManagedKeyArrayOutput)
}

// NamespaceCustomerManagedKeyMapInput is an input type that accepts NamespaceCustomerManagedKeyMap and NamespaceCustomerManagedKeyMapOutput values.
// You can construct a concrete instance of `NamespaceCustomerManagedKeyMapInput` via:
//
//          NamespaceCustomerManagedKeyMap{ "key": NamespaceCustomerManagedKeyArgs{...} }
type NamespaceCustomerManagedKeyMapInput interface {
	pulumi.Input

	ToNamespaceCustomerManagedKeyMapOutput() NamespaceCustomerManagedKeyMapOutput
	ToNamespaceCustomerManagedKeyMapOutputWithContext(context.Context) NamespaceCustomerManagedKeyMapOutput
}

type NamespaceCustomerManagedKeyMap map[string]NamespaceCustomerManagedKeyInput

func (NamespaceCustomerManagedKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceCustomerManagedKey)(nil)).Elem()
}

func (i NamespaceCustomerManagedKeyMap) ToNamespaceCustomerManagedKeyMapOutput() NamespaceCustomerManagedKeyMapOutput {
	return i.ToNamespaceCustomerManagedKeyMapOutputWithContext(context.Background())
}

func (i NamespaceCustomerManagedKeyMap) ToNamespaceCustomerManagedKeyMapOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceCustomerManagedKeyMapOutput)
}

type NamespaceCustomerManagedKeyOutput struct{ *pulumi.OutputState }

func (NamespaceCustomerManagedKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceCustomerManagedKey)(nil))
}

func (o NamespaceCustomerManagedKeyOutput) ToNamespaceCustomerManagedKeyOutput() NamespaceCustomerManagedKeyOutput {
	return o
}

func (o NamespaceCustomerManagedKeyOutput) ToNamespaceCustomerManagedKeyOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyOutput {
	return o
}

func (o NamespaceCustomerManagedKeyOutput) ToNamespaceCustomerManagedKeyPtrOutput() NamespaceCustomerManagedKeyPtrOutput {
	return o.ToNamespaceCustomerManagedKeyPtrOutputWithContext(context.Background())
}

func (o NamespaceCustomerManagedKeyOutput) ToNamespaceCustomerManagedKeyPtrOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamespaceCustomerManagedKey) *NamespaceCustomerManagedKey {
		return &v
	}).(NamespaceCustomerManagedKeyPtrOutput)
}

type NamespaceCustomerManagedKeyPtrOutput struct{ *pulumi.OutputState }

func (NamespaceCustomerManagedKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceCustomerManagedKey)(nil))
}

func (o NamespaceCustomerManagedKeyPtrOutput) ToNamespaceCustomerManagedKeyPtrOutput() NamespaceCustomerManagedKeyPtrOutput {
	return o
}

func (o NamespaceCustomerManagedKeyPtrOutput) ToNamespaceCustomerManagedKeyPtrOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyPtrOutput {
	return o
}

func (o NamespaceCustomerManagedKeyPtrOutput) Elem() NamespaceCustomerManagedKeyOutput {
	return o.ApplyT(func(v *NamespaceCustomerManagedKey) NamespaceCustomerManagedKey {
		if v != nil {
			return *v
		}
		var ret NamespaceCustomerManagedKey
		return ret
	}).(NamespaceCustomerManagedKeyOutput)
}

type NamespaceCustomerManagedKeyArrayOutput struct{ *pulumi.OutputState }

func (NamespaceCustomerManagedKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceCustomerManagedKey)(nil))
}

func (o NamespaceCustomerManagedKeyArrayOutput) ToNamespaceCustomerManagedKeyArrayOutput() NamespaceCustomerManagedKeyArrayOutput {
	return o
}

func (o NamespaceCustomerManagedKeyArrayOutput) ToNamespaceCustomerManagedKeyArrayOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyArrayOutput {
	return o
}

func (o NamespaceCustomerManagedKeyArrayOutput) Index(i pulumi.IntInput) NamespaceCustomerManagedKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceCustomerManagedKey {
		return vs[0].([]NamespaceCustomerManagedKey)[vs[1].(int)]
	}).(NamespaceCustomerManagedKeyOutput)
}

type NamespaceCustomerManagedKeyMapOutput struct{ *pulumi.OutputState }

func (NamespaceCustomerManagedKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NamespaceCustomerManagedKey)(nil))
}

func (o NamespaceCustomerManagedKeyMapOutput) ToNamespaceCustomerManagedKeyMapOutput() NamespaceCustomerManagedKeyMapOutput {
	return o
}

func (o NamespaceCustomerManagedKeyMapOutput) ToNamespaceCustomerManagedKeyMapOutputWithContext(ctx context.Context) NamespaceCustomerManagedKeyMapOutput {
	return o
}

func (o NamespaceCustomerManagedKeyMapOutput) MapIndex(k pulumi.StringInput) NamespaceCustomerManagedKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NamespaceCustomerManagedKey {
		return vs[0].(map[string]NamespaceCustomerManagedKey)[vs[1].(string)]
	}).(NamespaceCustomerManagedKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceCustomerManagedKeyInput)(nil)).Elem(), &NamespaceCustomerManagedKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceCustomerManagedKeyPtrInput)(nil)).Elem(), &NamespaceCustomerManagedKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceCustomerManagedKeyArrayInput)(nil)).Elem(), NamespaceCustomerManagedKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceCustomerManagedKeyMapInput)(nil)).Elem(), NamespaceCustomerManagedKeyMap{})
	pulumi.RegisterOutputType(NamespaceCustomerManagedKeyOutput{})
	pulumi.RegisterOutputType(NamespaceCustomerManagedKeyPtrOutput{})
	pulumi.RegisterOutputType(NamespaceCustomerManagedKeyArrayOutput{})
	pulumi.RegisterOutputType(NamespaceCustomerManagedKeyMapOutput{})
}
