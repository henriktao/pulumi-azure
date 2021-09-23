// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Promotes an App Service Slot to Production within an App Service.
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomId(ctx, "server", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewPlan(ctx, "examplePlan", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleSlot, err := appservice.NewSlot(ctx, "exampleSlot", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewActiveSlot(ctx, "exampleActiveSlot", &appservice.ActiveSlotArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			AppServiceName:     exampleAppService.Name,
// 			AppServiceSlotName: exampleSlot.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ActiveSlot struct {
	pulumi.CustomResourceState

	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringOutput `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewActiveSlot registers a new resource with the given unique name, arguments, and options.
func NewActiveSlot(ctx *pulumi.Context,
	name string, args *ActiveSlotArgs, opts ...pulumi.ResourceOption) (*ActiveSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceName'")
	}
	if args.AppServiceSlotName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceSlotName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ActiveSlot
	err := ctx.RegisterResource("azure:appservice/activeSlot:ActiveSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveSlot gets an existing ActiveSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveSlotState, opts ...pulumi.ResourceOption) (*ActiveSlot, error) {
	var resource ActiveSlot
	err := ctx.ReadResource("azure:appservice/activeSlot:ActiveSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveSlot resources.
type activeSlotState struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName *string `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ActiveSlotState struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringPtrInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ActiveSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeSlotState)(nil)).Elem()
}

type activeSlotArgs struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName string `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ActiveSlot resource.
type ActiveSlotArgs struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ActiveSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeSlotArgs)(nil)).Elem()
}

type ActiveSlotInput interface {
	pulumi.Input

	ToActiveSlotOutput() ActiveSlotOutput
	ToActiveSlotOutputWithContext(ctx context.Context) ActiveSlotOutput
}

func (*ActiveSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSlot)(nil))
}

func (i *ActiveSlot) ToActiveSlotOutput() ActiveSlotOutput {
	return i.ToActiveSlotOutputWithContext(context.Background())
}

func (i *ActiveSlot) ToActiveSlotOutputWithContext(ctx context.Context) ActiveSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSlotOutput)
}

func (i *ActiveSlot) ToActiveSlotPtrOutput() ActiveSlotPtrOutput {
	return i.ToActiveSlotPtrOutputWithContext(context.Background())
}

func (i *ActiveSlot) ToActiveSlotPtrOutputWithContext(ctx context.Context) ActiveSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSlotPtrOutput)
}

type ActiveSlotPtrInput interface {
	pulumi.Input

	ToActiveSlotPtrOutput() ActiveSlotPtrOutput
	ToActiveSlotPtrOutputWithContext(ctx context.Context) ActiveSlotPtrOutput
}

type activeSlotPtrType ActiveSlotArgs

func (*activeSlotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveSlot)(nil))
}

func (i *activeSlotPtrType) ToActiveSlotPtrOutput() ActiveSlotPtrOutput {
	return i.ToActiveSlotPtrOutputWithContext(context.Background())
}

func (i *activeSlotPtrType) ToActiveSlotPtrOutputWithContext(ctx context.Context) ActiveSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSlotPtrOutput)
}

// ActiveSlotArrayInput is an input type that accepts ActiveSlotArray and ActiveSlotArrayOutput values.
// You can construct a concrete instance of `ActiveSlotArrayInput` via:
//
//          ActiveSlotArray{ ActiveSlotArgs{...} }
type ActiveSlotArrayInput interface {
	pulumi.Input

	ToActiveSlotArrayOutput() ActiveSlotArrayOutput
	ToActiveSlotArrayOutputWithContext(context.Context) ActiveSlotArrayOutput
}

type ActiveSlotArray []ActiveSlotInput

func (ActiveSlotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActiveSlot)(nil)).Elem()
}

func (i ActiveSlotArray) ToActiveSlotArrayOutput() ActiveSlotArrayOutput {
	return i.ToActiveSlotArrayOutputWithContext(context.Background())
}

func (i ActiveSlotArray) ToActiveSlotArrayOutputWithContext(ctx context.Context) ActiveSlotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSlotArrayOutput)
}

// ActiveSlotMapInput is an input type that accepts ActiveSlotMap and ActiveSlotMapOutput values.
// You can construct a concrete instance of `ActiveSlotMapInput` via:
//
//          ActiveSlotMap{ "key": ActiveSlotArgs{...} }
type ActiveSlotMapInput interface {
	pulumi.Input

	ToActiveSlotMapOutput() ActiveSlotMapOutput
	ToActiveSlotMapOutputWithContext(context.Context) ActiveSlotMapOutput
}

type ActiveSlotMap map[string]ActiveSlotInput

func (ActiveSlotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActiveSlot)(nil)).Elem()
}

func (i ActiveSlotMap) ToActiveSlotMapOutput() ActiveSlotMapOutput {
	return i.ToActiveSlotMapOutputWithContext(context.Background())
}

func (i ActiveSlotMap) ToActiveSlotMapOutputWithContext(ctx context.Context) ActiveSlotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveSlotMapOutput)
}

type ActiveSlotOutput struct{ *pulumi.OutputState }

func (ActiveSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveSlot)(nil))
}

func (o ActiveSlotOutput) ToActiveSlotOutput() ActiveSlotOutput {
	return o
}

func (o ActiveSlotOutput) ToActiveSlotOutputWithContext(ctx context.Context) ActiveSlotOutput {
	return o
}

func (o ActiveSlotOutput) ToActiveSlotPtrOutput() ActiveSlotPtrOutput {
	return o.ToActiveSlotPtrOutputWithContext(context.Background())
}

func (o ActiveSlotOutput) ToActiveSlotPtrOutputWithContext(ctx context.Context) ActiveSlotPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveSlot) *ActiveSlot {
		return &v
	}).(ActiveSlotPtrOutput)
}

type ActiveSlotPtrOutput struct{ *pulumi.OutputState }

func (ActiveSlotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveSlot)(nil))
}

func (o ActiveSlotPtrOutput) ToActiveSlotPtrOutput() ActiveSlotPtrOutput {
	return o
}

func (o ActiveSlotPtrOutput) ToActiveSlotPtrOutputWithContext(ctx context.Context) ActiveSlotPtrOutput {
	return o
}

func (o ActiveSlotPtrOutput) Elem() ActiveSlotOutput {
	return o.ApplyT(func(v *ActiveSlot) ActiveSlot {
		if v != nil {
			return *v
		}
		var ret ActiveSlot
		return ret
	}).(ActiveSlotOutput)
}

type ActiveSlotArrayOutput struct{ *pulumi.OutputState }

func (ActiveSlotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveSlot)(nil))
}

func (o ActiveSlotArrayOutput) ToActiveSlotArrayOutput() ActiveSlotArrayOutput {
	return o
}

func (o ActiveSlotArrayOutput) ToActiveSlotArrayOutputWithContext(ctx context.Context) ActiveSlotArrayOutput {
	return o
}

func (o ActiveSlotArrayOutput) Index(i pulumi.IntInput) ActiveSlotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveSlot {
		return vs[0].([]ActiveSlot)[vs[1].(int)]
	}).(ActiveSlotOutput)
}

type ActiveSlotMapOutput struct{ *pulumi.OutputState }

func (ActiveSlotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActiveSlot)(nil))
}

func (o ActiveSlotMapOutput) ToActiveSlotMapOutput() ActiveSlotMapOutput {
	return o
}

func (o ActiveSlotMapOutput) ToActiveSlotMapOutputWithContext(ctx context.Context) ActiveSlotMapOutput {
	return o
}

func (o ActiveSlotMapOutput) MapIndex(k pulumi.StringInput) ActiveSlotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActiveSlot {
		return vs[0].(map[string]ActiveSlot)[vs[1].(string)]
	}).(ActiveSlotOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveSlotOutput{})
	pulumi.RegisterOutputType(ActiveSlotPtrOutput{})
	pulumi.RegisterOutputType(ActiveSlotArrayOutput{})
	pulumi.RegisterOutputType(ActiveSlotMapOutput{})
}
