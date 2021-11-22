// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service Virtual Network Association (this is for the [Regional VNet Integration](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration)).
//
// > **Note:** This resource can be used for both `appservice.AppService` and `appservice.FunctionApp`.
//
// > **Note:** There is a hard limit of [one VNet integration per App Service Plan](https://docs.microsoft.com/en-us/azure/app-service/web-sites-integrate-with-vnet#regional-vnet-integration).
// Multiple apps in the same App Service plan can use the same VNet.
//
// ## Example Usage
// ### With App Service)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("example-delegation"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Web/serverFarms"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewVirtualNetworkSwiftConnection(ctx, "exampleVirtualNetworkSwiftConnection", &appservice.VirtualNetworkSwiftConnectionArgs{
// 			AppServiceId: exampleAppService.ID(),
// 			SubnetId:     exampleSubnet.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Function App)
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 			Delegations: network.SubnetDelegationArray{
// 				&network.SubnetDelegationArgs{
// 					Name: pulumi.String("example-delegation"),
// 					ServiceDelegation: &network.SubnetDelegationServiceDelegationArgs{
// 						Name: pulumi.String("Microsoft.Web/serverFarms"),
// 						Actions: pulumi.StringArray{
// 							pulumi.String("Microsoft.Network/virtualNetworks/subnets/action"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFunctionApp, err := appservice.NewFunctionApp(ctx, "exampleFunctionApp", &appservice.FunctionAppArgs{
// 			Location:                exampleResourceGroup.Location,
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			AppServicePlanId:        examplePlan.ID(),
// 			StorageAccountName:      exampleAccount.Name,
// 			StorageAccountAccessKey: exampleAccount.PrimaryAccessKey,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewVirtualNetworkSwiftConnection(ctx, "exampleVirtualNetworkSwiftConnection", &appservice.VirtualNetworkSwiftConnectionArgs{
// 			AppServiceId: exampleFunctionApp.ID(),
// 			SubnetId:     exampleSubnet.ID(),
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
// App Service Virtual Network Associations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection myassociation /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/config/virtualNetwork
// ```
type VirtualNetworkSwiftConnection struct {
	pulumi.CustomResourceState

	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringOutput `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewVirtualNetworkSwiftConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkSwiftConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceId == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource VirtualNetworkSwiftConnection
	err := ctx.RegisterResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkSwiftConnection gets an existing VirtualNetworkSwiftConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkSwiftConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	var resource VirtualNetworkSwiftConnection
	err := ctx.ReadResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkSwiftConnection resources.
type virtualNetworkSwiftConnectionState struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId *string `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkSwiftConnectionState struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringPtrInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkSwiftConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionState)(nil)).Elem()
}

type virtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId string `pulumi:"appServiceId"`
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkSwiftConnection resource.
type VirtualNetworkSwiftConnectionArgs struct {
	// The ID of the App Service or Function App to associate to the VNet. Changing this forces a new resource to be created.
	AppServiceId pulumi.StringInput
	// The ID of the subnet the app service will be associated to (the subnet must have a `serviceDelegation` configured for `Microsoft.Web/serverFarms`).
	SubnetId pulumi.StringInput
}

func (VirtualNetworkSwiftConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionArgs)(nil)).Elem()
}

type VirtualNetworkSwiftConnectionInput interface {
	pulumi.Input

	ToVirtualNetworkSwiftConnectionOutput() VirtualNetworkSwiftConnectionOutput
	ToVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionOutput
}

func (*VirtualNetworkSwiftConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkSwiftConnection)(nil))
}

func (i *VirtualNetworkSwiftConnection) ToVirtualNetworkSwiftConnectionOutput() VirtualNetworkSwiftConnectionOutput {
	return i.ToVirtualNetworkSwiftConnectionOutputWithContext(context.Background())
}

func (i *VirtualNetworkSwiftConnection) ToVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkSwiftConnectionOutput)
}

func (i *VirtualNetworkSwiftConnection) ToVirtualNetworkSwiftConnectionPtrOutput() VirtualNetworkSwiftConnectionPtrOutput {
	return i.ToVirtualNetworkSwiftConnectionPtrOutputWithContext(context.Background())
}

func (i *VirtualNetworkSwiftConnection) ToVirtualNetworkSwiftConnectionPtrOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkSwiftConnectionPtrOutput)
}

type VirtualNetworkSwiftConnectionPtrInput interface {
	pulumi.Input

	ToVirtualNetworkSwiftConnectionPtrOutput() VirtualNetworkSwiftConnectionPtrOutput
	ToVirtualNetworkSwiftConnectionPtrOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionPtrOutput
}

type virtualNetworkSwiftConnectionPtrType VirtualNetworkSwiftConnectionArgs

func (*virtualNetworkSwiftConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkSwiftConnection)(nil))
}

func (i *virtualNetworkSwiftConnectionPtrType) ToVirtualNetworkSwiftConnectionPtrOutput() VirtualNetworkSwiftConnectionPtrOutput {
	return i.ToVirtualNetworkSwiftConnectionPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkSwiftConnectionPtrType) ToVirtualNetworkSwiftConnectionPtrOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkSwiftConnectionPtrOutput)
}

// VirtualNetworkSwiftConnectionArrayInput is an input type that accepts VirtualNetworkSwiftConnectionArray and VirtualNetworkSwiftConnectionArrayOutput values.
// You can construct a concrete instance of `VirtualNetworkSwiftConnectionArrayInput` via:
//
//          VirtualNetworkSwiftConnectionArray{ VirtualNetworkSwiftConnectionArgs{...} }
type VirtualNetworkSwiftConnectionArrayInput interface {
	pulumi.Input

	ToVirtualNetworkSwiftConnectionArrayOutput() VirtualNetworkSwiftConnectionArrayOutput
	ToVirtualNetworkSwiftConnectionArrayOutputWithContext(context.Context) VirtualNetworkSwiftConnectionArrayOutput
}

type VirtualNetworkSwiftConnectionArray []VirtualNetworkSwiftConnectionInput

func (VirtualNetworkSwiftConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNetworkSwiftConnection)(nil)).Elem()
}

func (i VirtualNetworkSwiftConnectionArray) ToVirtualNetworkSwiftConnectionArrayOutput() VirtualNetworkSwiftConnectionArrayOutput {
	return i.ToVirtualNetworkSwiftConnectionArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkSwiftConnectionArray) ToVirtualNetworkSwiftConnectionArrayOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkSwiftConnectionArrayOutput)
}

// VirtualNetworkSwiftConnectionMapInput is an input type that accepts VirtualNetworkSwiftConnectionMap and VirtualNetworkSwiftConnectionMapOutput values.
// You can construct a concrete instance of `VirtualNetworkSwiftConnectionMapInput` via:
//
//          VirtualNetworkSwiftConnectionMap{ "key": VirtualNetworkSwiftConnectionArgs{...} }
type VirtualNetworkSwiftConnectionMapInput interface {
	pulumi.Input

	ToVirtualNetworkSwiftConnectionMapOutput() VirtualNetworkSwiftConnectionMapOutput
	ToVirtualNetworkSwiftConnectionMapOutputWithContext(context.Context) VirtualNetworkSwiftConnectionMapOutput
}

type VirtualNetworkSwiftConnectionMap map[string]VirtualNetworkSwiftConnectionInput

func (VirtualNetworkSwiftConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNetworkSwiftConnection)(nil)).Elem()
}

func (i VirtualNetworkSwiftConnectionMap) ToVirtualNetworkSwiftConnectionMapOutput() VirtualNetworkSwiftConnectionMapOutput {
	return i.ToVirtualNetworkSwiftConnectionMapOutputWithContext(context.Background())
}

func (i VirtualNetworkSwiftConnectionMap) ToVirtualNetworkSwiftConnectionMapOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkSwiftConnectionMapOutput)
}

type VirtualNetworkSwiftConnectionOutput struct{ *pulumi.OutputState }

func (VirtualNetworkSwiftConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkSwiftConnection)(nil))
}

func (o VirtualNetworkSwiftConnectionOutput) ToVirtualNetworkSwiftConnectionOutput() VirtualNetworkSwiftConnectionOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionOutput) ToVirtualNetworkSwiftConnectionOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionOutput) ToVirtualNetworkSwiftConnectionPtrOutput() VirtualNetworkSwiftConnectionPtrOutput {
	return o.ToVirtualNetworkSwiftConnectionPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkSwiftConnectionOutput) ToVirtualNetworkSwiftConnectionPtrOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkSwiftConnection) *VirtualNetworkSwiftConnection {
		return &v
	}).(VirtualNetworkSwiftConnectionPtrOutput)
}

type VirtualNetworkSwiftConnectionPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkSwiftConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkSwiftConnection)(nil))
}

func (o VirtualNetworkSwiftConnectionPtrOutput) ToVirtualNetworkSwiftConnectionPtrOutput() VirtualNetworkSwiftConnectionPtrOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionPtrOutput) ToVirtualNetworkSwiftConnectionPtrOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionPtrOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionPtrOutput) Elem() VirtualNetworkSwiftConnectionOutput {
	return o.ApplyT(func(v *VirtualNetworkSwiftConnection) VirtualNetworkSwiftConnection {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkSwiftConnection
		return ret
	}).(VirtualNetworkSwiftConnectionOutput)
}

type VirtualNetworkSwiftConnectionArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkSwiftConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkSwiftConnection)(nil))
}

func (o VirtualNetworkSwiftConnectionArrayOutput) ToVirtualNetworkSwiftConnectionArrayOutput() VirtualNetworkSwiftConnectionArrayOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionArrayOutput) ToVirtualNetworkSwiftConnectionArrayOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionArrayOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionArrayOutput) Index(i pulumi.IntInput) VirtualNetworkSwiftConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkSwiftConnection {
		return vs[0].([]VirtualNetworkSwiftConnection)[vs[1].(int)]
	}).(VirtualNetworkSwiftConnectionOutput)
}

type VirtualNetworkSwiftConnectionMapOutput struct{ *pulumi.OutputState }

func (VirtualNetworkSwiftConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualNetworkSwiftConnection)(nil))
}

func (o VirtualNetworkSwiftConnectionMapOutput) ToVirtualNetworkSwiftConnectionMapOutput() VirtualNetworkSwiftConnectionMapOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionMapOutput) ToVirtualNetworkSwiftConnectionMapOutputWithContext(ctx context.Context) VirtualNetworkSwiftConnectionMapOutput {
	return o
}

func (o VirtualNetworkSwiftConnectionMapOutput) MapIndex(k pulumi.StringInput) VirtualNetworkSwiftConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualNetworkSwiftConnection {
		return vs[0].(map[string]VirtualNetworkSwiftConnection)[vs[1].(string)]
	}).(VirtualNetworkSwiftConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkSwiftConnectionInput)(nil)).Elem(), &VirtualNetworkSwiftConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkSwiftConnectionPtrInput)(nil)).Elem(), &VirtualNetworkSwiftConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkSwiftConnectionArrayInput)(nil)).Elem(), VirtualNetworkSwiftConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkSwiftConnectionMapInput)(nil)).Elem(), VirtualNetworkSwiftConnectionMap{})
	pulumi.RegisterOutputType(VirtualNetworkSwiftConnectionOutput{})
	pulumi.RegisterOutputType(VirtualNetworkSwiftConnectionPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkSwiftConnectionArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkSwiftConnectionMapOutput{})
}
