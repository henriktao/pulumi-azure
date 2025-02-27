// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Azure Site Recovery protection container. Protection containers serve as containers for replicated VMs and belong to a single region / recovery fabric. Protection containers can contain more than one replicated VM. To replicate a VM, a container must exist in both the source and target Azure regions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/recoveryservices"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/siterecovery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.NewResourceGroup(ctx, "primary", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := core.NewResourceGroup(ctx, "secondary", &core.ResourceGroupArgs{
// 			Location: pulumi.String("East US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vault, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
// 			Location:          secondary.Location,
// 			ResourceGroupName: secondary.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fabric, err := siterecovery.NewFabric(ctx, "fabric", &siterecovery.FabricArgs{
// 			ResourceGroupName: secondary.Name,
// 			RecoveryVaultName: vault.Name,
// 			Location:          primary.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = siterecovery.NewProtectionContainer(ctx, "protection-container", &siterecovery.ProtectionContainerArgs{
// 			ResourceGroupName:  secondary.Name,
// 			RecoveryVaultName:  vault.Name,
// 			RecoveryFabricName: fabric.Name,
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
// Site Recovery Protection Containers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:siterecovery/protectionContainer:ProtectionContainer mycontainer /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name
// ```
type ProtectionContainer struct {
	pulumi.CustomResourceState

	// The name of the protection container.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of fabric that should contain this protection container.
	RecoveryFabricName pulumi.StringOutput `pulumi:"recoveryFabricName"`
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewProtectionContainer registers a new resource with the given unique name, arguments, and options.
func NewProtectionContainer(ctx *pulumi.Context,
	name string, args *ProtectionContainerArgs, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecoveryFabricName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryFabricName'")
	}
	if args.RecoveryVaultName == nil {
		return nil, errors.New("invalid value for required argument 'RecoveryVaultName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ProtectionContainer
	err := ctx.RegisterResource("azure:siterecovery/protectionContainer:ProtectionContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectionContainer gets an existing ProtectionContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectionContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionContainerState, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	var resource ProtectionContainer
	err := ctx.ReadResource("azure:siterecovery/protectionContainer:ProtectionContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectionContainer resources.
type protectionContainerState struct {
	// The name of the protection container.
	Name *string `pulumi:"name"`
	// Name of fabric that should contain this protection container.
	RecoveryFabricName *string `pulumi:"recoveryFabricName"`
	// The name of the vault that should be updated.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ProtectionContainerState struct {
	// The name of the protection container.
	Name pulumi.StringPtrInput
	// Name of fabric that should contain this protection container.
	RecoveryFabricName pulumi.StringPtrInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringPtrInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringPtrInput
}

func (ProtectionContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerState)(nil)).Elem()
}

type protectionContainerArgs struct {
	// The name of the protection container.
	Name *string `pulumi:"name"`
	// Name of fabric that should contain this protection container.
	RecoveryFabricName string `pulumi:"recoveryFabricName"`
	// The name of the vault that should be updated.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProtectionContainer resource.
type ProtectionContainerArgs struct {
	// The name of the protection container.
	Name pulumi.StringPtrInput
	// Name of fabric that should contain this protection container.
	RecoveryFabricName pulumi.StringInput
	// The name of the vault that should be updated.
	RecoveryVaultName pulumi.StringInput
	// Name of the resource group where the vault that should be updated is located.
	ResourceGroupName pulumi.StringInput
}

func (ProtectionContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerArgs)(nil)).Elem()
}

type ProtectionContainerInput interface {
	pulumi.Input

	ToProtectionContainerOutput() ProtectionContainerOutput
	ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput
}

func (*ProtectionContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (i *ProtectionContainer) ToProtectionContainerOutput() ProtectionContainerOutput {
	return i.ToProtectionContainerOutputWithContext(context.Background())
}

func (i *ProtectionContainer) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerOutput)
}

// ProtectionContainerArrayInput is an input type that accepts ProtectionContainerArray and ProtectionContainerArrayOutput values.
// You can construct a concrete instance of `ProtectionContainerArrayInput` via:
//
//          ProtectionContainerArray{ ProtectionContainerArgs{...} }
type ProtectionContainerArrayInput interface {
	pulumi.Input

	ToProtectionContainerArrayOutput() ProtectionContainerArrayOutput
	ToProtectionContainerArrayOutputWithContext(context.Context) ProtectionContainerArrayOutput
}

type ProtectionContainerArray []ProtectionContainerInput

func (ProtectionContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProtectionContainer)(nil)).Elem()
}

func (i ProtectionContainerArray) ToProtectionContainerArrayOutput() ProtectionContainerArrayOutput {
	return i.ToProtectionContainerArrayOutputWithContext(context.Background())
}

func (i ProtectionContainerArray) ToProtectionContainerArrayOutputWithContext(ctx context.Context) ProtectionContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerArrayOutput)
}

// ProtectionContainerMapInput is an input type that accepts ProtectionContainerMap and ProtectionContainerMapOutput values.
// You can construct a concrete instance of `ProtectionContainerMapInput` via:
//
//          ProtectionContainerMap{ "key": ProtectionContainerArgs{...} }
type ProtectionContainerMapInput interface {
	pulumi.Input

	ToProtectionContainerMapOutput() ProtectionContainerMapOutput
	ToProtectionContainerMapOutputWithContext(context.Context) ProtectionContainerMapOutput
}

type ProtectionContainerMap map[string]ProtectionContainerInput

func (ProtectionContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProtectionContainer)(nil)).Elem()
}

func (i ProtectionContainerMap) ToProtectionContainerMapOutput() ProtectionContainerMapOutput {
	return i.ToProtectionContainerMapOutputWithContext(context.Background())
}

func (i ProtectionContainerMap) ToProtectionContainerMapOutputWithContext(ctx context.Context) ProtectionContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerMapOutput)
}

type ProtectionContainerOutput struct{ *pulumi.OutputState }

func (ProtectionContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (o ProtectionContainerOutput) ToProtectionContainerOutput() ProtectionContainerOutput {
	return o
}

func (o ProtectionContainerOutput) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return o
}

type ProtectionContainerArrayOutput struct{ *pulumi.OutputState }

func (ProtectionContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProtectionContainer)(nil)).Elem()
}

func (o ProtectionContainerArrayOutput) ToProtectionContainerArrayOutput() ProtectionContainerArrayOutput {
	return o
}

func (o ProtectionContainerArrayOutput) ToProtectionContainerArrayOutputWithContext(ctx context.Context) ProtectionContainerArrayOutput {
	return o
}

func (o ProtectionContainerArrayOutput) Index(i pulumi.IntInput) ProtectionContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProtectionContainer {
		return vs[0].([]*ProtectionContainer)[vs[1].(int)]
	}).(ProtectionContainerOutput)
}

type ProtectionContainerMapOutput struct{ *pulumi.OutputState }

func (ProtectionContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProtectionContainer)(nil)).Elem()
}

func (o ProtectionContainerMapOutput) ToProtectionContainerMapOutput() ProtectionContainerMapOutput {
	return o
}

func (o ProtectionContainerMapOutput) ToProtectionContainerMapOutputWithContext(ctx context.Context) ProtectionContainerMapOutput {
	return o
}

func (o ProtectionContainerMapOutput) MapIndex(k pulumi.StringInput) ProtectionContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProtectionContainer {
		return vs[0].(map[string]*ProtectionContainer)[vs[1].(string)]
	}).(ProtectionContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionContainerInput)(nil)).Elem(), &ProtectionContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionContainerArrayInput)(nil)).Elem(), ProtectionContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionContainerMapInput)(nil)).Elem(), ProtectionContainerMap{})
	pulumi.RegisterOutputType(ProtectionContainerOutput{})
	pulumi.RegisterOutputType(ProtectionContainerArrayOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMapOutput{})
}
