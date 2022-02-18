// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages the Registration Info for a Virtual Desktop Host Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/desktopvirtualization"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("westeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleHostPool, err := desktopvirtualization.NewHostPool(ctx, "exampleHostPool", &desktopvirtualization.HostPoolArgs{
// 			Location:            exampleResourceGroup.Location,
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			Type:                pulumi.String("Pooled"),
// 			ValidateEnvironment: pulumi.Bool(true),
// 			LoadBalancerType:    pulumi.String("BreadthFirst"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = desktopvirtualization.NewgetHostPoolRegistrationInfo(ctx, "examplegetHostPoolRegistrationInfo", &desktopvirtualization.getHostPoolRegistrationInfoArgs{
// 			HostpoolId:     exampleHostPool.ID(),
// 			ExpirationDate: pulumi.String("2022-01-01T23:40:52Z"),
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
// AVD Registration Infos can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:desktopvirtualization/getHostPoolRegistrationInfo:getHostPoolRegistrationInfo example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/hostPools/pool1/registrationInfo/default
// ```
type GetHostPoolRegistrationInfo struct {
	pulumi.CustomResourceState

	// A valid `RFC3339Time` for the expiration of the token..
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtualDesktopHostPoolRegistrationInfo resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolId pulumi.StringOutput `pulumi:"hostpoolId"`
	// The registration token generated by the Virtual Desktop Host Pool for registration of session hosts.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewGetHostPoolRegistrationInfo registers a new resource with the given unique name, arguments, and options.
func NewGetHostPoolRegistrationInfo(ctx *pulumi.Context,
	name string, args *GetHostPoolRegistrationInfoArgs, opts ...pulumi.ResourceOption) (*GetHostPoolRegistrationInfo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpirationDate == nil {
		return nil, errors.New("invalid value for required argument 'ExpirationDate'")
	}
	if args.HostpoolId == nil {
		return nil, errors.New("invalid value for required argument 'HostpoolId'")
	}
	var resource GetHostPoolRegistrationInfo
	err := ctx.RegisterResource("azure:desktopvirtualization/getHostPoolRegistrationInfo:getHostPoolRegistrationInfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGetHostPoolRegistrationInfo gets an existing GetHostPoolRegistrationInfo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGetHostPoolRegistrationInfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GetHostPoolRegistrationInfoState, opts ...pulumi.ResourceOption) (*GetHostPoolRegistrationInfo, error) {
	var resource GetHostPoolRegistrationInfo
	err := ctx.ReadResource("azure:desktopvirtualization/getHostPoolRegistrationInfo:getHostPoolRegistrationInfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GetHostPoolRegistrationInfo resources.
type getHostPoolRegistrationInfoState struct {
	// A valid `RFC3339Time` for the expiration of the token..
	ExpirationDate *string `pulumi:"expirationDate"`
	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtualDesktopHostPoolRegistrationInfo resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolId *string `pulumi:"hostpoolId"`
	// The registration token generated by the Virtual Desktop Host Pool for registration of session hosts.
	Token *string `pulumi:"token"`
}

type GetHostPoolRegistrationInfoState struct {
	// A valid `RFC3339Time` for the expiration of the token..
	ExpirationDate pulumi.StringPtrInput
	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtualDesktopHostPoolRegistrationInfo resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolId pulumi.StringPtrInput
	// The registration token generated by the Virtual Desktop Host Pool for registration of session hosts.
	Token pulumi.StringPtrInput
}

func (GetHostPoolRegistrationInfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*getHostPoolRegistrationInfoState)(nil)).Elem()
}

type getHostPoolRegistrationInfoArgs struct {
	// A valid `RFC3339Time` for the expiration of the token..
	ExpirationDate string `pulumi:"expirationDate"`
	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtualDesktopHostPoolRegistrationInfo resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolId string `pulumi:"hostpoolId"`
}

// The set of arguments for constructing a GetHostPoolRegistrationInfo resource.
type GetHostPoolRegistrationInfoArgs struct {
	// A valid `RFC3339Time` for the expiration of the token..
	ExpirationDate pulumi.StringInput
	// The ID of the Virtual Desktop Host Pool to link the Registration Info to. Changing this forces a new Registration Info resource to be created. Only a single virtualDesktopHostPoolRegistrationInfo resource should be associated with a given hostpool. Assigning multiple resources will produce inconsistent results.
	HostpoolId pulumi.StringInput
}

func (GetHostPoolRegistrationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*getHostPoolRegistrationInfoArgs)(nil)).Elem()
}

type GetHostPoolRegistrationInfoInput interface {
	pulumi.Input

	ToGetHostPoolRegistrationInfoOutput() GetHostPoolRegistrationInfoOutput
	ToGetHostPoolRegistrationInfoOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoOutput
}

func (*GetHostPoolRegistrationInfo) ElementType() reflect.Type {
	return reflect.TypeOf((**GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (i *GetHostPoolRegistrationInfo) ToGetHostPoolRegistrationInfoOutput() GetHostPoolRegistrationInfoOutput {
	return i.ToGetHostPoolRegistrationInfoOutputWithContext(context.Background())
}

func (i *GetHostPoolRegistrationInfo) ToGetHostPoolRegistrationInfoOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetHostPoolRegistrationInfoOutput)
}

// GetHostPoolRegistrationInfoArrayInput is an input type that accepts GetHostPoolRegistrationInfoArray and GetHostPoolRegistrationInfoArrayOutput values.
// You can construct a concrete instance of `GetHostPoolRegistrationInfoArrayInput` via:
//
//          GetHostPoolRegistrationInfoArray{ GetHostPoolRegistrationInfoArgs{...} }
type GetHostPoolRegistrationInfoArrayInput interface {
	pulumi.Input

	ToGetHostPoolRegistrationInfoArrayOutput() GetHostPoolRegistrationInfoArrayOutput
	ToGetHostPoolRegistrationInfoArrayOutputWithContext(context.Context) GetHostPoolRegistrationInfoArrayOutput
}

type GetHostPoolRegistrationInfoArray []GetHostPoolRegistrationInfoInput

func (GetHostPoolRegistrationInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (i GetHostPoolRegistrationInfoArray) ToGetHostPoolRegistrationInfoArrayOutput() GetHostPoolRegistrationInfoArrayOutput {
	return i.ToGetHostPoolRegistrationInfoArrayOutputWithContext(context.Background())
}

func (i GetHostPoolRegistrationInfoArray) ToGetHostPoolRegistrationInfoArrayOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetHostPoolRegistrationInfoArrayOutput)
}

// GetHostPoolRegistrationInfoMapInput is an input type that accepts GetHostPoolRegistrationInfoMap and GetHostPoolRegistrationInfoMapOutput values.
// You can construct a concrete instance of `GetHostPoolRegistrationInfoMapInput` via:
//
//          GetHostPoolRegistrationInfoMap{ "key": GetHostPoolRegistrationInfoArgs{...} }
type GetHostPoolRegistrationInfoMapInput interface {
	pulumi.Input

	ToGetHostPoolRegistrationInfoMapOutput() GetHostPoolRegistrationInfoMapOutput
	ToGetHostPoolRegistrationInfoMapOutputWithContext(context.Context) GetHostPoolRegistrationInfoMapOutput
}

type GetHostPoolRegistrationInfoMap map[string]GetHostPoolRegistrationInfoInput

func (GetHostPoolRegistrationInfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (i GetHostPoolRegistrationInfoMap) ToGetHostPoolRegistrationInfoMapOutput() GetHostPoolRegistrationInfoMapOutput {
	return i.ToGetHostPoolRegistrationInfoMapOutputWithContext(context.Background())
}

func (i GetHostPoolRegistrationInfoMap) ToGetHostPoolRegistrationInfoMapOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetHostPoolRegistrationInfoMapOutput)
}

type GetHostPoolRegistrationInfoOutput struct{ *pulumi.OutputState }

func (GetHostPoolRegistrationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (o GetHostPoolRegistrationInfoOutput) ToGetHostPoolRegistrationInfoOutput() GetHostPoolRegistrationInfoOutput {
	return o
}

func (o GetHostPoolRegistrationInfoOutput) ToGetHostPoolRegistrationInfoOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoOutput {
	return o
}

type GetHostPoolRegistrationInfoArrayOutput struct{ *pulumi.OutputState }

func (GetHostPoolRegistrationInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (o GetHostPoolRegistrationInfoArrayOutput) ToGetHostPoolRegistrationInfoArrayOutput() GetHostPoolRegistrationInfoArrayOutput {
	return o
}

func (o GetHostPoolRegistrationInfoArrayOutput) ToGetHostPoolRegistrationInfoArrayOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoArrayOutput {
	return o
}

func (o GetHostPoolRegistrationInfoArrayOutput) Index(i pulumi.IntInput) GetHostPoolRegistrationInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GetHostPoolRegistrationInfo {
		return vs[0].([]*GetHostPoolRegistrationInfo)[vs[1].(int)]
	}).(GetHostPoolRegistrationInfoOutput)
}

type GetHostPoolRegistrationInfoMapOutput struct{ *pulumi.OutputState }

func (GetHostPoolRegistrationInfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GetHostPoolRegistrationInfo)(nil)).Elem()
}

func (o GetHostPoolRegistrationInfoMapOutput) ToGetHostPoolRegistrationInfoMapOutput() GetHostPoolRegistrationInfoMapOutput {
	return o
}

func (o GetHostPoolRegistrationInfoMapOutput) ToGetHostPoolRegistrationInfoMapOutputWithContext(ctx context.Context) GetHostPoolRegistrationInfoMapOutput {
	return o
}

func (o GetHostPoolRegistrationInfoMapOutput) MapIndex(k pulumi.StringInput) GetHostPoolRegistrationInfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GetHostPoolRegistrationInfo {
		return vs[0].(map[string]*GetHostPoolRegistrationInfo)[vs[1].(string)]
	}).(GetHostPoolRegistrationInfoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetHostPoolRegistrationInfoInput)(nil)).Elem(), &GetHostPoolRegistrationInfo{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetHostPoolRegistrationInfoArrayInput)(nil)).Elem(), GetHostPoolRegistrationInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetHostPoolRegistrationInfoMapInput)(nil)).Elem(), GetHostPoolRegistrationInfoMap{})
	pulumi.RegisterOutputType(GetHostPoolRegistrationInfoOutput{})
	pulumi.RegisterOutputType(GetHostPoolRegistrationInfoArrayOutput{})
	pulumi.RegisterOutputType(GetHostPoolRegistrationInfoMapOutput{})
}
