// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Factory Azure Integration Runtime.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewIntegrationRuntimeRule(ctx, "exampleIntegrationRuntimeRule", &datafactory.IntegrationRuntimeRuleArgs{
// 			DataFactoryName:   exampleFactory.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
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
// Data Factory Azure Integration Runtimes can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
// ```
type IntegrationRuntimeRule struct {
	pulumi.CustomResourceState

	// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
	ComputeType pulumi.StringPtrOutput `pulumi:"computeType"`
	// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
	CoreCount pulumi.IntPtrOutput `pulumi:"coreCount"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
	TimeToLiveMin pulumi.IntPtrOutput `pulumi:"timeToLiveMin"`
	// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
	VirtualNetworkEnabled pulumi.BoolPtrOutput `pulumi:"virtualNetworkEnabled"`
}

// NewIntegrationRuntimeRule registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntimeRule(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeRuleArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntimeRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationRuntimeRule
	err := ctx.RegisterResource("azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntimeRule gets an existing IntegrationRuntimeRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntimeRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeRuleState, opts ...pulumi.ResourceOption) (*IntegrationRuntimeRule, error) {
	var resource IntegrationRuntimeRule
	err := ctx.ReadResource("azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntimeRule resources.
type integrationRuntimeRuleState struct {
	// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
	ComputeType *string `pulumi:"computeType"`
	// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
	CoreCount *int `pulumi:"coreCount"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
	TimeToLiveMin *int `pulumi:"timeToLiveMin"`
	// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
	VirtualNetworkEnabled *bool `pulumi:"virtualNetworkEnabled"`
}

type IntegrationRuntimeRuleState struct {
	// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
	ComputeType pulumi.StringPtrInput
	// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
	CoreCount pulumi.IntPtrInput
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringPtrInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
	TimeToLiveMin pulumi.IntPtrInput
	// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
	VirtualNetworkEnabled pulumi.BoolPtrInput
}

func (IntegrationRuntimeRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeRuleState)(nil)).Elem()
}

type integrationRuntimeRuleArgs struct {
	// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
	ComputeType *string `pulumi:"computeType"`
	// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
	CoreCount *int `pulumi:"coreCount"`
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// Integration runtime description.
	Description *string `pulumi:"description"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
	TimeToLiveMin *int `pulumi:"timeToLiveMin"`
	// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
	VirtualNetworkEnabled *bool `pulumi:"virtualNetworkEnabled"`
}

// The set of arguments for constructing a IntegrationRuntimeRule resource.
type IntegrationRuntimeRuleArgs struct {
	// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
	ComputeType pulumi.StringPtrInput
	// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
	CoreCount pulumi.IntPtrInput
	// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringInput
	// Integration runtime description.
	Description pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
	TimeToLiveMin pulumi.IntPtrInput
	// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
	VirtualNetworkEnabled pulumi.BoolPtrInput
}

func (IntegrationRuntimeRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeRuleArgs)(nil)).Elem()
}

type IntegrationRuntimeRuleInput interface {
	pulumi.Input

	ToIntegrationRuntimeRuleOutput() IntegrationRuntimeRuleOutput
	ToIntegrationRuntimeRuleOutputWithContext(ctx context.Context) IntegrationRuntimeRuleOutput
}

func (*IntegrationRuntimeRule) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeRule)(nil))
}

func (i *IntegrationRuntimeRule) ToIntegrationRuntimeRuleOutput() IntegrationRuntimeRuleOutput {
	return i.ToIntegrationRuntimeRuleOutputWithContext(context.Background())
}

func (i *IntegrationRuntimeRule) ToIntegrationRuntimeRuleOutputWithContext(ctx context.Context) IntegrationRuntimeRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeRuleOutput)
}

func (i *IntegrationRuntimeRule) ToIntegrationRuntimeRulePtrOutput() IntegrationRuntimeRulePtrOutput {
	return i.ToIntegrationRuntimeRulePtrOutputWithContext(context.Background())
}

func (i *IntegrationRuntimeRule) ToIntegrationRuntimeRulePtrOutputWithContext(ctx context.Context) IntegrationRuntimeRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeRulePtrOutput)
}

type IntegrationRuntimeRulePtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeRulePtrOutput() IntegrationRuntimeRulePtrOutput
	ToIntegrationRuntimeRulePtrOutputWithContext(ctx context.Context) IntegrationRuntimeRulePtrOutput
}

type integrationRuntimeRulePtrType IntegrationRuntimeRuleArgs

func (*integrationRuntimeRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeRule)(nil))
}

func (i *integrationRuntimeRulePtrType) ToIntegrationRuntimeRulePtrOutput() IntegrationRuntimeRulePtrOutput {
	return i.ToIntegrationRuntimeRulePtrOutputWithContext(context.Background())
}

func (i *integrationRuntimeRulePtrType) ToIntegrationRuntimeRulePtrOutputWithContext(ctx context.Context) IntegrationRuntimeRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeRulePtrOutput)
}

// IntegrationRuntimeRuleArrayInput is an input type that accepts IntegrationRuntimeRuleArray and IntegrationRuntimeRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationRuntimeRuleArrayInput` via:
//
//          IntegrationRuntimeRuleArray{ IntegrationRuntimeRuleArgs{...} }
type IntegrationRuntimeRuleArrayInput interface {
	pulumi.Input

	ToIntegrationRuntimeRuleArrayOutput() IntegrationRuntimeRuleArrayOutput
	ToIntegrationRuntimeRuleArrayOutputWithContext(context.Context) IntegrationRuntimeRuleArrayOutput
}

type IntegrationRuntimeRuleArray []IntegrationRuntimeRuleInput

func (IntegrationRuntimeRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IntegrationRuntimeRule)(nil))
}

func (i IntegrationRuntimeRuleArray) ToIntegrationRuntimeRuleArrayOutput() IntegrationRuntimeRuleArrayOutput {
	return i.ToIntegrationRuntimeRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationRuntimeRuleArray) ToIntegrationRuntimeRuleArrayOutputWithContext(ctx context.Context) IntegrationRuntimeRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeRuleArrayOutput)
}

// IntegrationRuntimeRuleMapInput is an input type that accepts IntegrationRuntimeRuleMap and IntegrationRuntimeRuleMapOutput values.
// You can construct a concrete instance of `IntegrationRuntimeRuleMapInput` via:
//
//          IntegrationRuntimeRuleMap{ "key": IntegrationRuntimeRuleArgs{...} }
type IntegrationRuntimeRuleMapInput interface {
	pulumi.Input

	ToIntegrationRuntimeRuleMapOutput() IntegrationRuntimeRuleMapOutput
	ToIntegrationRuntimeRuleMapOutputWithContext(context.Context) IntegrationRuntimeRuleMapOutput
}

type IntegrationRuntimeRuleMap map[string]IntegrationRuntimeRuleInput

func (IntegrationRuntimeRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IntegrationRuntimeRule)(nil))
}

func (i IntegrationRuntimeRuleMap) ToIntegrationRuntimeRuleMapOutput() IntegrationRuntimeRuleMapOutput {
	return i.ToIntegrationRuntimeRuleMapOutputWithContext(context.Background())
}

func (i IntegrationRuntimeRuleMap) ToIntegrationRuntimeRuleMapOutputWithContext(ctx context.Context) IntegrationRuntimeRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeRuleMapOutput)
}

type IntegrationRuntimeRuleOutput struct {
	*pulumi.OutputState
}

func (IntegrationRuntimeRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeRule)(nil))
}

func (o IntegrationRuntimeRuleOutput) ToIntegrationRuntimeRuleOutput() IntegrationRuntimeRuleOutput {
	return o
}

func (o IntegrationRuntimeRuleOutput) ToIntegrationRuntimeRuleOutputWithContext(ctx context.Context) IntegrationRuntimeRuleOutput {
	return o
}

func (o IntegrationRuntimeRuleOutput) ToIntegrationRuntimeRulePtrOutput() IntegrationRuntimeRulePtrOutput {
	return o.ToIntegrationRuntimeRulePtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeRuleOutput) ToIntegrationRuntimeRulePtrOutputWithContext(ctx context.Context) IntegrationRuntimeRulePtrOutput {
	return o.ApplyT(func(v IntegrationRuntimeRule) *IntegrationRuntimeRule {
		return &v
	}).(IntegrationRuntimeRulePtrOutput)
}

type IntegrationRuntimeRulePtrOutput struct {
	*pulumi.OutputState
}

func (IntegrationRuntimeRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeRule)(nil))
}

func (o IntegrationRuntimeRulePtrOutput) ToIntegrationRuntimeRulePtrOutput() IntegrationRuntimeRulePtrOutput {
	return o
}

func (o IntegrationRuntimeRulePtrOutput) ToIntegrationRuntimeRulePtrOutputWithContext(ctx context.Context) IntegrationRuntimeRulePtrOutput {
	return o
}

type IntegrationRuntimeRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationRuntimeRule)(nil))
}

func (o IntegrationRuntimeRuleArrayOutput) ToIntegrationRuntimeRuleArrayOutput() IntegrationRuntimeRuleArrayOutput {
	return o
}

func (o IntegrationRuntimeRuleArrayOutput) ToIntegrationRuntimeRuleArrayOutputWithContext(ctx context.Context) IntegrationRuntimeRuleArrayOutput {
	return o
}

func (o IntegrationRuntimeRuleArrayOutput) Index(i pulumi.IntInput) IntegrationRuntimeRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationRuntimeRule {
		return vs[0].([]IntegrationRuntimeRule)[vs[1].(int)]
	}).(IntegrationRuntimeRuleOutput)
}

type IntegrationRuntimeRuleMapOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IntegrationRuntimeRule)(nil))
}

func (o IntegrationRuntimeRuleMapOutput) ToIntegrationRuntimeRuleMapOutput() IntegrationRuntimeRuleMapOutput {
	return o
}

func (o IntegrationRuntimeRuleMapOutput) ToIntegrationRuntimeRuleMapOutputWithContext(ctx context.Context) IntegrationRuntimeRuleMapOutput {
	return o
}

func (o IntegrationRuntimeRuleMapOutput) MapIndex(k pulumi.StringInput) IntegrationRuntimeRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IntegrationRuntimeRule {
		return vs[0].(map[string]IntegrationRuntimeRule)[vs[1].(string)]
	}).(IntegrationRuntimeRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationRuntimeRuleOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeRulePtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeRuleArrayOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeRuleMapOutput{})
}
