// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Hostname Binding within an App Service Slot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleSlot, err := appservice.NewSlot(ctx, "exampleSlot", &appservice.SlotArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServiceName:    exampleAppService.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewSlotCustomHostnameBinding(ctx, "exampleSlotCustomHostnameBinding", &appservice.SlotCustomHostnameBindingArgs{
// 			AppServiceSlotId: exampleSlot.ID(),
// 			Hostname:         pulumi.String("www.mywebsite.com"),
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
// App Service Custom Hostname Bindings can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/slotCustomHostnameBinding:SlotCustomHostnameBinding mywebsite /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/hostNameBindings/mywebsite.com
// ```
type SlotCustomHostnameBinding struct {
	pulumi.CustomResourceState

	// The ID of the App Service Slot. Changing this forces a new resource to be created.
	AppServiceSlotId pulumi.StringOutput `pulumi:"appServiceSlotId"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringOutput `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp pulumi.StringOutput `pulumi:"virtualIp"`
}

// NewSlotCustomHostnameBinding registers a new resource with the given unique name, arguments, and options.
func NewSlotCustomHostnameBinding(ctx *pulumi.Context,
	name string, args *SlotCustomHostnameBindingArgs, opts ...pulumi.ResourceOption) (*SlotCustomHostnameBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceSlotId == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceSlotId'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	var resource SlotCustomHostnameBinding
	err := ctx.RegisterResource("azure:appservice/slotCustomHostnameBinding:SlotCustomHostnameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlotCustomHostnameBinding gets an existing SlotCustomHostnameBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlotCustomHostnameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlotCustomHostnameBindingState, opts ...pulumi.ResourceOption) (*SlotCustomHostnameBinding, error) {
	var resource SlotCustomHostnameBinding
	err := ctx.ReadResource("azure:appservice/slotCustomHostnameBinding:SlotCustomHostnameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SlotCustomHostnameBinding resources.
type slotCustomHostnameBindingState struct {
	// The ID of the App Service Slot. Changing this forces a new resource to be created.
	AppServiceSlotId *string `pulumi:"appServiceSlotId"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname *string `pulumi:"hostname"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState *string `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint *string `pulumi:"thumbprint"`
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp *string `pulumi:"virtualIp"`
}

type SlotCustomHostnameBindingState struct {
	// The ID of the App Service Slot. Changing this forces a new resource to be created.
	AppServiceSlotId pulumi.StringPtrInput
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringPtrInput
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringPtrInput
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringPtrInput
	// The virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIp pulumi.StringPtrInput
}

func (SlotCustomHostnameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*slotCustomHostnameBindingState)(nil)).Elem()
}

type slotCustomHostnameBindingArgs struct {
	// The ID of the App Service Slot. Changing this forces a new resource to be created.
	AppServiceSlotId string `pulumi:"appServiceSlotId"`
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname string `pulumi:"hostname"`
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState *string `pulumi:"sslState"`
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a SlotCustomHostnameBinding resource.
type SlotCustomHostnameBindingArgs struct {
	// The ID of the App Service Slot. Changing this forces a new resource to be created.
	AppServiceSlotId pulumi.StringInput
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname pulumi.StringInput
	// The SSL type. Possible values are `IpBasedEnabled` and `SniEnabled`. Changing this forces a new resource to be created.
	SslState pulumi.StringPtrInput
	// The SSL certificate thumbprint. Changing this forces a new resource to be created.
	Thumbprint pulumi.StringPtrInput
}

func (SlotCustomHostnameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slotCustomHostnameBindingArgs)(nil)).Elem()
}

type SlotCustomHostnameBindingInput interface {
	pulumi.Input

	ToSlotCustomHostnameBindingOutput() SlotCustomHostnameBindingOutput
	ToSlotCustomHostnameBindingOutputWithContext(ctx context.Context) SlotCustomHostnameBindingOutput
}

func (*SlotCustomHostnameBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotCustomHostnameBinding)(nil))
}

func (i *SlotCustomHostnameBinding) ToSlotCustomHostnameBindingOutput() SlotCustomHostnameBindingOutput {
	return i.ToSlotCustomHostnameBindingOutputWithContext(context.Background())
}

func (i *SlotCustomHostnameBinding) ToSlotCustomHostnameBindingOutputWithContext(ctx context.Context) SlotCustomHostnameBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotCustomHostnameBindingOutput)
}

func (i *SlotCustomHostnameBinding) ToSlotCustomHostnameBindingPtrOutput() SlotCustomHostnameBindingPtrOutput {
	return i.ToSlotCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (i *SlotCustomHostnameBinding) ToSlotCustomHostnameBindingPtrOutputWithContext(ctx context.Context) SlotCustomHostnameBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotCustomHostnameBindingPtrOutput)
}

type SlotCustomHostnameBindingPtrInput interface {
	pulumi.Input

	ToSlotCustomHostnameBindingPtrOutput() SlotCustomHostnameBindingPtrOutput
	ToSlotCustomHostnameBindingPtrOutputWithContext(ctx context.Context) SlotCustomHostnameBindingPtrOutput
}

type slotCustomHostnameBindingPtrType SlotCustomHostnameBindingArgs

func (*slotCustomHostnameBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotCustomHostnameBinding)(nil))
}

func (i *slotCustomHostnameBindingPtrType) ToSlotCustomHostnameBindingPtrOutput() SlotCustomHostnameBindingPtrOutput {
	return i.ToSlotCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (i *slotCustomHostnameBindingPtrType) ToSlotCustomHostnameBindingPtrOutputWithContext(ctx context.Context) SlotCustomHostnameBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotCustomHostnameBindingPtrOutput)
}

// SlotCustomHostnameBindingArrayInput is an input type that accepts SlotCustomHostnameBindingArray and SlotCustomHostnameBindingArrayOutput values.
// You can construct a concrete instance of `SlotCustomHostnameBindingArrayInput` via:
//
//          SlotCustomHostnameBindingArray{ SlotCustomHostnameBindingArgs{...} }
type SlotCustomHostnameBindingArrayInput interface {
	pulumi.Input

	ToSlotCustomHostnameBindingArrayOutput() SlotCustomHostnameBindingArrayOutput
	ToSlotCustomHostnameBindingArrayOutputWithContext(context.Context) SlotCustomHostnameBindingArrayOutput
}

type SlotCustomHostnameBindingArray []SlotCustomHostnameBindingInput

func (SlotCustomHostnameBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SlotCustomHostnameBinding)(nil)).Elem()
}

func (i SlotCustomHostnameBindingArray) ToSlotCustomHostnameBindingArrayOutput() SlotCustomHostnameBindingArrayOutput {
	return i.ToSlotCustomHostnameBindingArrayOutputWithContext(context.Background())
}

func (i SlotCustomHostnameBindingArray) ToSlotCustomHostnameBindingArrayOutputWithContext(ctx context.Context) SlotCustomHostnameBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotCustomHostnameBindingArrayOutput)
}

// SlotCustomHostnameBindingMapInput is an input type that accepts SlotCustomHostnameBindingMap and SlotCustomHostnameBindingMapOutput values.
// You can construct a concrete instance of `SlotCustomHostnameBindingMapInput` via:
//
//          SlotCustomHostnameBindingMap{ "key": SlotCustomHostnameBindingArgs{...} }
type SlotCustomHostnameBindingMapInput interface {
	pulumi.Input

	ToSlotCustomHostnameBindingMapOutput() SlotCustomHostnameBindingMapOutput
	ToSlotCustomHostnameBindingMapOutputWithContext(context.Context) SlotCustomHostnameBindingMapOutput
}

type SlotCustomHostnameBindingMap map[string]SlotCustomHostnameBindingInput

func (SlotCustomHostnameBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SlotCustomHostnameBinding)(nil)).Elem()
}

func (i SlotCustomHostnameBindingMap) ToSlotCustomHostnameBindingMapOutput() SlotCustomHostnameBindingMapOutput {
	return i.ToSlotCustomHostnameBindingMapOutputWithContext(context.Background())
}

func (i SlotCustomHostnameBindingMap) ToSlotCustomHostnameBindingMapOutputWithContext(ctx context.Context) SlotCustomHostnameBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotCustomHostnameBindingMapOutput)
}

type SlotCustomHostnameBindingOutput struct{ *pulumi.OutputState }

func (SlotCustomHostnameBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotCustomHostnameBinding)(nil))
}

func (o SlotCustomHostnameBindingOutput) ToSlotCustomHostnameBindingOutput() SlotCustomHostnameBindingOutput {
	return o
}

func (o SlotCustomHostnameBindingOutput) ToSlotCustomHostnameBindingOutputWithContext(ctx context.Context) SlotCustomHostnameBindingOutput {
	return o
}

func (o SlotCustomHostnameBindingOutput) ToSlotCustomHostnameBindingPtrOutput() SlotCustomHostnameBindingPtrOutput {
	return o.ToSlotCustomHostnameBindingPtrOutputWithContext(context.Background())
}

func (o SlotCustomHostnameBindingOutput) ToSlotCustomHostnameBindingPtrOutputWithContext(ctx context.Context) SlotCustomHostnameBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlotCustomHostnameBinding) *SlotCustomHostnameBinding {
		return &v
	}).(SlotCustomHostnameBindingPtrOutput)
}

type SlotCustomHostnameBindingPtrOutput struct{ *pulumi.OutputState }

func (SlotCustomHostnameBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotCustomHostnameBinding)(nil))
}

func (o SlotCustomHostnameBindingPtrOutput) ToSlotCustomHostnameBindingPtrOutput() SlotCustomHostnameBindingPtrOutput {
	return o
}

func (o SlotCustomHostnameBindingPtrOutput) ToSlotCustomHostnameBindingPtrOutputWithContext(ctx context.Context) SlotCustomHostnameBindingPtrOutput {
	return o
}

func (o SlotCustomHostnameBindingPtrOutput) Elem() SlotCustomHostnameBindingOutput {
	return o.ApplyT(func(v *SlotCustomHostnameBinding) SlotCustomHostnameBinding {
		if v != nil {
			return *v
		}
		var ret SlotCustomHostnameBinding
		return ret
	}).(SlotCustomHostnameBindingOutput)
}

type SlotCustomHostnameBindingArrayOutput struct{ *pulumi.OutputState }

func (SlotCustomHostnameBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlotCustomHostnameBinding)(nil))
}

func (o SlotCustomHostnameBindingArrayOutput) ToSlotCustomHostnameBindingArrayOutput() SlotCustomHostnameBindingArrayOutput {
	return o
}

func (o SlotCustomHostnameBindingArrayOutput) ToSlotCustomHostnameBindingArrayOutputWithContext(ctx context.Context) SlotCustomHostnameBindingArrayOutput {
	return o
}

func (o SlotCustomHostnameBindingArrayOutput) Index(i pulumi.IntInput) SlotCustomHostnameBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlotCustomHostnameBinding {
		return vs[0].([]SlotCustomHostnameBinding)[vs[1].(int)]
	}).(SlotCustomHostnameBindingOutput)
}

type SlotCustomHostnameBindingMapOutput struct{ *pulumi.OutputState }

func (SlotCustomHostnameBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SlotCustomHostnameBinding)(nil))
}

func (o SlotCustomHostnameBindingMapOutput) ToSlotCustomHostnameBindingMapOutput() SlotCustomHostnameBindingMapOutput {
	return o
}

func (o SlotCustomHostnameBindingMapOutput) ToSlotCustomHostnameBindingMapOutputWithContext(ctx context.Context) SlotCustomHostnameBindingMapOutput {
	return o
}

func (o SlotCustomHostnameBindingMapOutput) MapIndex(k pulumi.StringInput) SlotCustomHostnameBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SlotCustomHostnameBinding {
		return vs[0].(map[string]SlotCustomHostnameBinding)[vs[1].(string)]
	}).(SlotCustomHostnameBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SlotCustomHostnameBindingInput)(nil)).Elem(), &SlotCustomHostnameBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotCustomHostnameBindingPtrInput)(nil)).Elem(), &SlotCustomHostnameBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotCustomHostnameBindingArrayInput)(nil)).Elem(), SlotCustomHostnameBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlotCustomHostnameBindingMapInput)(nil)).Elem(), SlotCustomHostnameBindingMap{})
	pulumi.RegisterOutputType(SlotCustomHostnameBindingOutput{})
	pulumi.RegisterOutputType(SlotCustomHostnameBindingPtrOutput{})
	pulumi.RegisterOutputType(SlotCustomHostnameBindingArrayOutput{})
	pulumi.RegisterOutputType(SlotCustomHostnameBindingMapOutput{})
}
