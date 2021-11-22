// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfiguration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure App Configuration Feature.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appconfiguration"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appconf, err := appconfiguration.NewConfigurationStore(ctx, "appconf", &appconfiguration.ConfigurationStoreArgs{
// 			ResourceGroupName: rg.Name,
// 			Location:          rg.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appconfiguration.NewConfigurationFeature(ctx, "test", &appconfiguration.ConfigurationFeatureArgs{
// 			ConfigurationStoreId: appconf.ID(),
// 			Description:          pulumi.String("test description"),
// 			Label:                pulumi.String(fmt.Sprintf("%v%v%v", "acctest-ackeylabel-", "%", "d")),
// 			Enabled:              pulumi.Bool(true),
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
// App Configuration Features can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appconfiguration/configurationFeature:ConfigurationFeature test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/label1
// ```
//
//  If you wish to import a key with an empty label then sustitute the label's name with `%00`, like this
//
// ```sh
//  $ pulumi import azure:appconfiguration/configurationFeature:ConfigurationFeature test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/%00
// ```
type ConfigurationFeature struct {
	pulumi.CustomResourceState

	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringOutput `pulumi:"configurationStoreId"`
	// The description of the App Configuration Feature.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The status of the App Configuration Feature. By default, this is set to false.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	Etag    pulumi.StringOutput  `pulumi:"etag"`
	// The label of the App Configuration Feature.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// Should this App Configuration Feature be Locked to prevent changes?
	Locked pulumi.BoolPtrOutput `pulumi:"locked"`
	// The name of the App Configuration Feature. Changing this foces a new resource to be crearted.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of one or more numbers representing the value of the percentage required to enable this feature.
	PercentageFilterValue pulumi.IntPtrOutput `pulumi:"percentageFilterValue"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `targetingFilter` block as defined below.
	TargetingFilters ConfigurationFeatureTargetingFilterArrayOutput `pulumi:"targetingFilters"`
	// A `targetingFilter` block `timewindowFilter` as defined below.
	TimewindowFilters ConfigurationFeatureTimewindowFilterArrayOutput `pulumi:"timewindowFilters"`
}

// NewConfigurationFeature registers a new resource with the given unique name, arguments, and options.
func NewConfigurationFeature(ctx *pulumi.Context,
	name string, args *ConfigurationFeatureArgs, opts ...pulumi.ResourceOption) (*ConfigurationFeature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationStoreId'")
	}
	var resource ConfigurationFeature
	err := ctx.RegisterResource("azure:appconfiguration/configurationFeature:ConfigurationFeature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationFeature gets an existing ConfigurationFeature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationFeatureState, opts ...pulumi.ResourceOption) (*ConfigurationFeature, error) {
	var resource ConfigurationFeature
	err := ctx.ReadResource("azure:appconfiguration/configurationFeature:ConfigurationFeature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationFeature resources.
type configurationFeatureState struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId *string `pulumi:"configurationStoreId"`
	// The description of the App Configuration Feature.
	Description *string `pulumi:"description"`
	// The status of the App Configuration Feature. By default, this is set to false.
	Enabled *bool   `pulumi:"enabled"`
	Etag    *string `pulumi:"etag"`
	// The label of the App Configuration Feature.  Changing this forces a new resource to be created.
	Label *string `pulumi:"label"`
	// Should this App Configuration Feature be Locked to prevent changes?
	Locked *bool `pulumi:"locked"`
	// The name of the App Configuration Feature. Changing this foces a new resource to be crearted.
	Name *string `pulumi:"name"`
	// A list of one or more numbers representing the value of the percentage required to enable this feature.
	PercentageFilterValue *int `pulumi:"percentageFilterValue"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `targetingFilter` block as defined below.
	TargetingFilters []ConfigurationFeatureTargetingFilter `pulumi:"targetingFilters"`
	// A `targetingFilter` block `timewindowFilter` as defined below.
	TimewindowFilters []ConfigurationFeatureTimewindowFilter `pulumi:"timewindowFilters"`
}

type ConfigurationFeatureState struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringPtrInput
	// The description of the App Configuration Feature.
	Description pulumi.StringPtrInput
	// The status of the App Configuration Feature. By default, this is set to false.
	Enabled pulumi.BoolPtrInput
	Etag    pulumi.StringPtrInput
	// The label of the App Configuration Feature.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrInput
	// Should this App Configuration Feature be Locked to prevent changes?
	Locked pulumi.BoolPtrInput
	// The name of the App Configuration Feature. Changing this foces a new resource to be crearted.
	Name pulumi.StringPtrInput
	// A list of one or more numbers representing the value of the percentage required to enable this feature.
	PercentageFilterValue pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `targetingFilter` block as defined below.
	TargetingFilters ConfigurationFeatureTargetingFilterArrayInput
	// A `targetingFilter` block `timewindowFilter` as defined below.
	TimewindowFilters ConfigurationFeatureTimewindowFilterArrayInput
}

func (ConfigurationFeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationFeatureState)(nil)).Elem()
}

type configurationFeatureArgs struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId string `pulumi:"configurationStoreId"`
	// The description of the App Configuration Feature.
	Description *string `pulumi:"description"`
	// The status of the App Configuration Feature. By default, this is set to false.
	Enabled *bool   `pulumi:"enabled"`
	Etag    *string `pulumi:"etag"`
	// The label of the App Configuration Feature.  Changing this forces a new resource to be created.
	Label *string `pulumi:"label"`
	// Should this App Configuration Feature be Locked to prevent changes?
	Locked *bool `pulumi:"locked"`
	// The name of the App Configuration Feature. Changing this foces a new resource to be crearted.
	Name *string `pulumi:"name"`
	// A list of one or more numbers representing the value of the percentage required to enable this feature.
	PercentageFilterValue *int `pulumi:"percentageFilterValue"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A `targetingFilter` block as defined below.
	TargetingFilters []ConfigurationFeatureTargetingFilter `pulumi:"targetingFilters"`
	// A `targetingFilter` block `timewindowFilter` as defined below.
	TimewindowFilters []ConfigurationFeatureTimewindowFilter `pulumi:"timewindowFilters"`
}

// The set of arguments for constructing a ConfigurationFeature resource.
type ConfigurationFeatureArgs struct {
	// Specifies the id of the App Configuration. Changing this forces a new resource to be created.
	ConfigurationStoreId pulumi.StringInput
	// The description of the App Configuration Feature.
	Description pulumi.StringPtrInput
	// The status of the App Configuration Feature. By default, this is set to false.
	Enabled pulumi.BoolPtrInput
	Etag    pulumi.StringPtrInput
	// The label of the App Configuration Feature.  Changing this forces a new resource to be created.
	Label pulumi.StringPtrInput
	// Should this App Configuration Feature be Locked to prevent changes?
	Locked pulumi.BoolPtrInput
	// The name of the App Configuration Feature. Changing this foces a new resource to be crearted.
	Name pulumi.StringPtrInput
	// A list of one or more numbers representing the value of the percentage required to enable this feature.
	PercentageFilterValue pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A `targetingFilter` block as defined below.
	TargetingFilters ConfigurationFeatureTargetingFilterArrayInput
	// A `targetingFilter` block `timewindowFilter` as defined below.
	TimewindowFilters ConfigurationFeatureTimewindowFilterArrayInput
}

func (ConfigurationFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationFeatureArgs)(nil)).Elem()
}

type ConfigurationFeatureInput interface {
	pulumi.Input

	ToConfigurationFeatureOutput() ConfigurationFeatureOutput
	ToConfigurationFeatureOutputWithContext(ctx context.Context) ConfigurationFeatureOutput
}

func (*ConfigurationFeature) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFeature)(nil))
}

func (i *ConfigurationFeature) ToConfigurationFeatureOutput() ConfigurationFeatureOutput {
	return i.ToConfigurationFeatureOutputWithContext(context.Background())
}

func (i *ConfigurationFeature) ToConfigurationFeatureOutputWithContext(ctx context.Context) ConfigurationFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFeatureOutput)
}

func (i *ConfigurationFeature) ToConfigurationFeaturePtrOutput() ConfigurationFeaturePtrOutput {
	return i.ToConfigurationFeaturePtrOutputWithContext(context.Background())
}

func (i *ConfigurationFeature) ToConfigurationFeaturePtrOutputWithContext(ctx context.Context) ConfigurationFeaturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFeaturePtrOutput)
}

type ConfigurationFeaturePtrInput interface {
	pulumi.Input

	ToConfigurationFeaturePtrOutput() ConfigurationFeaturePtrOutput
	ToConfigurationFeaturePtrOutputWithContext(ctx context.Context) ConfigurationFeaturePtrOutput
}

type configurationFeaturePtrType ConfigurationFeatureArgs

func (*configurationFeaturePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationFeature)(nil))
}

func (i *configurationFeaturePtrType) ToConfigurationFeaturePtrOutput() ConfigurationFeaturePtrOutput {
	return i.ToConfigurationFeaturePtrOutputWithContext(context.Background())
}

func (i *configurationFeaturePtrType) ToConfigurationFeaturePtrOutputWithContext(ctx context.Context) ConfigurationFeaturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFeaturePtrOutput)
}

// ConfigurationFeatureArrayInput is an input type that accepts ConfigurationFeatureArray and ConfigurationFeatureArrayOutput values.
// You can construct a concrete instance of `ConfigurationFeatureArrayInput` via:
//
//          ConfigurationFeatureArray{ ConfigurationFeatureArgs{...} }
type ConfigurationFeatureArrayInput interface {
	pulumi.Input

	ToConfigurationFeatureArrayOutput() ConfigurationFeatureArrayOutput
	ToConfigurationFeatureArrayOutputWithContext(context.Context) ConfigurationFeatureArrayOutput
}

type ConfigurationFeatureArray []ConfigurationFeatureInput

func (ConfigurationFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigurationFeature)(nil)).Elem()
}

func (i ConfigurationFeatureArray) ToConfigurationFeatureArrayOutput() ConfigurationFeatureArrayOutput {
	return i.ToConfigurationFeatureArrayOutputWithContext(context.Background())
}

func (i ConfigurationFeatureArray) ToConfigurationFeatureArrayOutputWithContext(ctx context.Context) ConfigurationFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFeatureArrayOutput)
}

// ConfigurationFeatureMapInput is an input type that accepts ConfigurationFeatureMap and ConfigurationFeatureMapOutput values.
// You can construct a concrete instance of `ConfigurationFeatureMapInput` via:
//
//          ConfigurationFeatureMap{ "key": ConfigurationFeatureArgs{...} }
type ConfigurationFeatureMapInput interface {
	pulumi.Input

	ToConfigurationFeatureMapOutput() ConfigurationFeatureMapOutput
	ToConfigurationFeatureMapOutputWithContext(context.Context) ConfigurationFeatureMapOutput
}

type ConfigurationFeatureMap map[string]ConfigurationFeatureInput

func (ConfigurationFeatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigurationFeature)(nil)).Elem()
}

func (i ConfigurationFeatureMap) ToConfigurationFeatureMapOutput() ConfigurationFeatureMapOutput {
	return i.ToConfigurationFeatureMapOutputWithContext(context.Background())
}

func (i ConfigurationFeatureMap) ToConfigurationFeatureMapOutputWithContext(ctx context.Context) ConfigurationFeatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationFeatureMapOutput)
}

type ConfigurationFeatureOutput struct{ *pulumi.OutputState }

func (ConfigurationFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationFeature)(nil))
}

func (o ConfigurationFeatureOutput) ToConfigurationFeatureOutput() ConfigurationFeatureOutput {
	return o
}

func (o ConfigurationFeatureOutput) ToConfigurationFeatureOutputWithContext(ctx context.Context) ConfigurationFeatureOutput {
	return o
}

func (o ConfigurationFeatureOutput) ToConfigurationFeaturePtrOutput() ConfigurationFeaturePtrOutput {
	return o.ToConfigurationFeaturePtrOutputWithContext(context.Background())
}

func (o ConfigurationFeatureOutput) ToConfigurationFeaturePtrOutputWithContext(ctx context.Context) ConfigurationFeaturePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationFeature) *ConfigurationFeature {
		return &v
	}).(ConfigurationFeaturePtrOutput)
}

type ConfigurationFeaturePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationFeaturePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationFeature)(nil))
}

func (o ConfigurationFeaturePtrOutput) ToConfigurationFeaturePtrOutput() ConfigurationFeaturePtrOutput {
	return o
}

func (o ConfigurationFeaturePtrOutput) ToConfigurationFeaturePtrOutputWithContext(ctx context.Context) ConfigurationFeaturePtrOutput {
	return o
}

func (o ConfigurationFeaturePtrOutput) Elem() ConfigurationFeatureOutput {
	return o.ApplyT(func(v *ConfigurationFeature) ConfigurationFeature {
		if v != nil {
			return *v
		}
		var ret ConfigurationFeature
		return ret
	}).(ConfigurationFeatureOutput)
}

type ConfigurationFeatureArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationFeature)(nil))
}

func (o ConfigurationFeatureArrayOutput) ToConfigurationFeatureArrayOutput() ConfigurationFeatureArrayOutput {
	return o
}

func (o ConfigurationFeatureArrayOutput) ToConfigurationFeatureArrayOutputWithContext(ctx context.Context) ConfigurationFeatureArrayOutput {
	return o
}

func (o ConfigurationFeatureArrayOutput) Index(i pulumi.IntInput) ConfigurationFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationFeature {
		return vs[0].([]ConfigurationFeature)[vs[1].(int)]
	}).(ConfigurationFeatureOutput)
}

type ConfigurationFeatureMapOutput struct{ *pulumi.OutputState }

func (ConfigurationFeatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConfigurationFeature)(nil))
}

func (o ConfigurationFeatureMapOutput) ToConfigurationFeatureMapOutput() ConfigurationFeatureMapOutput {
	return o
}

func (o ConfigurationFeatureMapOutput) ToConfigurationFeatureMapOutputWithContext(ctx context.Context) ConfigurationFeatureMapOutput {
	return o
}

func (o ConfigurationFeatureMapOutput) MapIndex(k pulumi.StringInput) ConfigurationFeatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConfigurationFeature {
		return vs[0].(map[string]ConfigurationFeature)[vs[1].(string)]
	}).(ConfigurationFeatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFeatureInput)(nil)).Elem(), &ConfigurationFeature{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFeaturePtrInput)(nil)).Elem(), &ConfigurationFeature{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFeatureArrayInput)(nil)).Elem(), ConfigurationFeatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationFeatureMapInput)(nil)).Elem(), ConfigurationFeatureMap{})
	pulumi.RegisterOutputType(ConfigurationFeatureOutput{})
	pulumi.RegisterOutputType(ConfigurationFeaturePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationFeatureArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationFeatureMapOutput{})
}
