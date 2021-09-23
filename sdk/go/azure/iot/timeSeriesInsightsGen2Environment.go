// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure IoT Time Series Insights Gen2 Environment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
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
// 		storage, err := storage.NewAccount(ctx, "storage", &storage.AccountArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iot.NewTimeSeriesInsightsGen2Environment(ctx, "exampleTimeSeriesInsightsGen2Environment", &iot.TimeSeriesInsightsGen2EnvironmentArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			SkuName:                    pulumi.String("L1"),
// 			WarmStoreDataRetentionTime: pulumi.String("P30D"),
// 			IdProperties: pulumi.StringArray{
// 				pulumi.String("id"),
// 			},
// 			Storage: &iot.TimeSeriesInsightsGen2EnvironmentStorageArgs{
// 				Name: storage.Name,
// 				Key:  storage.PrimaryAccessKey,
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
// Azure IoT Time Series Insights Gen2 Environment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
// ```
type TimeSeriesInsightsGen2Environment struct {
	pulumi.CustomResourceState

	// The FQDN used to access the environment data.
	DataAccessFqdn pulumi.StringOutput `pulumi:"dataAccessFqdn"`
	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
	IdProperties pulumi.StringArrayOutput `pulumi:"idProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A `storage` block as defined below.
	Storage TimeSeriesInsightsGen2EnvironmentStorageOutput `pulumi:"storage"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	WarmStoreDataRetentionTime pulumi.StringPtrOutput `pulumi:"warmStoreDataRetentionTime"`
}

// NewTimeSeriesInsightsGen2Environment registers a new resource with the given unique name, arguments, and options.
func NewTimeSeriesInsightsGen2Environment(ctx *pulumi.Context,
	name string, args *TimeSeriesInsightsGen2EnvironmentArgs, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsGen2Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdProperties == nil {
		return nil, errors.New("invalid value for required argument 'IdProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	var resource TimeSeriesInsightsGen2Environment
	err := ctx.RegisterResource("azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeSeriesInsightsGen2Environment gets an existing TimeSeriesInsightsGen2Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeSeriesInsightsGen2Environment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesInsightsGen2EnvironmentState, opts ...pulumi.ResourceOption) (*TimeSeriesInsightsGen2Environment, error) {
	var resource TimeSeriesInsightsGen2Environment
	err := ctx.ReadResource("azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeSeriesInsightsGen2Environment resources.
type timeSeriesInsightsGen2EnvironmentState struct {
	// The FQDN used to access the environment data.
	DataAccessFqdn *string `pulumi:"dataAccessFqdn"`
	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
	IdProperties []string `pulumi:"idProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
	SkuName *string `pulumi:"skuName"`
	// A `storage` block as defined below.
	Storage *TimeSeriesInsightsGen2EnvironmentStorage `pulumi:"storage"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	WarmStoreDataRetentionTime *string `pulumi:"warmStoreDataRetentionTime"`
}

type TimeSeriesInsightsGen2EnvironmentState struct {
	// The FQDN used to access the environment data.
	DataAccessFqdn pulumi.StringPtrInput
	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
	IdProperties pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
	SkuName pulumi.StringPtrInput
	// A `storage` block as defined below.
	Storage TimeSeriesInsightsGen2EnvironmentStoragePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	WarmStoreDataRetentionTime pulumi.StringPtrInput
}

func (TimeSeriesInsightsGen2EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsGen2EnvironmentState)(nil)).Elem()
}

type timeSeriesInsightsGen2EnvironmentArgs struct {
	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
	IdProperties []string `pulumi:"idProperties"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
	SkuName string `pulumi:"skuName"`
	// A `storage` block as defined below.
	Storage TimeSeriesInsightsGen2EnvironmentStorage `pulumi:"storage"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	WarmStoreDataRetentionTime *string `pulumi:"warmStoreDataRetentionTime"`
}

// The set of arguments for constructing a TimeSeriesInsightsGen2Environment resource.
type TimeSeriesInsightsGen2EnvironmentArgs struct {
	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment
	IdProperties pulumi.StringArrayInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created. Must be globally unique.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment.
	ResourceGroupName pulumi.StringInput
	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only `L1`. For gen2, capacity cannot be specified.
	SkuName pulumi.StringInput
	// A `storage` block as defined below.
	Storage TimeSeriesInsightsGen2EnvironmentStorageInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	WarmStoreDataRetentionTime pulumi.StringPtrInput
}

func (TimeSeriesInsightsGen2EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesInsightsGen2EnvironmentArgs)(nil)).Elem()
}

type TimeSeriesInsightsGen2EnvironmentInput interface {
	pulumi.Input

	ToTimeSeriesInsightsGen2EnvironmentOutput() TimeSeriesInsightsGen2EnvironmentOutput
	ToTimeSeriesInsightsGen2EnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentOutput
}

func (*TimeSeriesInsightsGen2Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsGen2Environment)(nil))
}

func (i *TimeSeriesInsightsGen2Environment) ToTimeSeriesInsightsGen2EnvironmentOutput() TimeSeriesInsightsGen2EnvironmentOutput {
	return i.ToTimeSeriesInsightsGen2EnvironmentOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsGen2Environment) ToTimeSeriesInsightsGen2EnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsGen2EnvironmentOutput)
}

func (i *TimeSeriesInsightsGen2Environment) ToTimeSeriesInsightsGen2EnvironmentPtrOutput() TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return i.ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(context.Background())
}

func (i *TimeSeriesInsightsGen2Environment) ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsGen2EnvironmentPtrOutput)
}

type TimeSeriesInsightsGen2EnvironmentPtrInput interface {
	pulumi.Input

	ToTimeSeriesInsightsGen2EnvironmentPtrOutput() TimeSeriesInsightsGen2EnvironmentPtrOutput
	ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentPtrOutput
}

type timeSeriesInsightsGen2EnvironmentPtrType TimeSeriesInsightsGen2EnvironmentArgs

func (*timeSeriesInsightsGen2EnvironmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsGen2Environment)(nil))
}

func (i *timeSeriesInsightsGen2EnvironmentPtrType) ToTimeSeriesInsightsGen2EnvironmentPtrOutput() TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return i.ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(context.Background())
}

func (i *timeSeriesInsightsGen2EnvironmentPtrType) ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsGen2EnvironmentPtrOutput)
}

// TimeSeriesInsightsGen2EnvironmentArrayInput is an input type that accepts TimeSeriesInsightsGen2EnvironmentArray and TimeSeriesInsightsGen2EnvironmentArrayOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsGen2EnvironmentArrayInput` via:
//
//          TimeSeriesInsightsGen2EnvironmentArray{ TimeSeriesInsightsGen2EnvironmentArgs{...} }
type TimeSeriesInsightsGen2EnvironmentArrayInput interface {
	pulumi.Input

	ToTimeSeriesInsightsGen2EnvironmentArrayOutput() TimeSeriesInsightsGen2EnvironmentArrayOutput
	ToTimeSeriesInsightsGen2EnvironmentArrayOutputWithContext(context.Context) TimeSeriesInsightsGen2EnvironmentArrayOutput
}

type TimeSeriesInsightsGen2EnvironmentArray []TimeSeriesInsightsGen2EnvironmentInput

func (TimeSeriesInsightsGen2EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TimeSeriesInsightsGen2Environment)(nil)).Elem()
}

func (i TimeSeriesInsightsGen2EnvironmentArray) ToTimeSeriesInsightsGen2EnvironmentArrayOutput() TimeSeriesInsightsGen2EnvironmentArrayOutput {
	return i.ToTimeSeriesInsightsGen2EnvironmentArrayOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsGen2EnvironmentArray) ToTimeSeriesInsightsGen2EnvironmentArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsGen2EnvironmentArrayOutput)
}

// TimeSeriesInsightsGen2EnvironmentMapInput is an input type that accepts TimeSeriesInsightsGen2EnvironmentMap and TimeSeriesInsightsGen2EnvironmentMapOutput values.
// You can construct a concrete instance of `TimeSeriesInsightsGen2EnvironmentMapInput` via:
//
//          TimeSeriesInsightsGen2EnvironmentMap{ "key": TimeSeriesInsightsGen2EnvironmentArgs{...} }
type TimeSeriesInsightsGen2EnvironmentMapInput interface {
	pulumi.Input

	ToTimeSeriesInsightsGen2EnvironmentMapOutput() TimeSeriesInsightsGen2EnvironmentMapOutput
	ToTimeSeriesInsightsGen2EnvironmentMapOutputWithContext(context.Context) TimeSeriesInsightsGen2EnvironmentMapOutput
}

type TimeSeriesInsightsGen2EnvironmentMap map[string]TimeSeriesInsightsGen2EnvironmentInput

func (TimeSeriesInsightsGen2EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TimeSeriesInsightsGen2Environment)(nil)).Elem()
}

func (i TimeSeriesInsightsGen2EnvironmentMap) ToTimeSeriesInsightsGen2EnvironmentMapOutput() TimeSeriesInsightsGen2EnvironmentMapOutput {
	return i.ToTimeSeriesInsightsGen2EnvironmentMapOutputWithContext(context.Background())
}

func (i TimeSeriesInsightsGen2EnvironmentMap) ToTimeSeriesInsightsGen2EnvironmentMapOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesInsightsGen2EnvironmentMapOutput)
}

type TimeSeriesInsightsGen2EnvironmentOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsGen2EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeSeriesInsightsGen2Environment)(nil))
}

func (o TimeSeriesInsightsGen2EnvironmentOutput) ToTimeSeriesInsightsGen2EnvironmentOutput() TimeSeriesInsightsGen2EnvironmentOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentOutput) ToTimeSeriesInsightsGen2EnvironmentOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentOutput) ToTimeSeriesInsightsGen2EnvironmentPtrOutput() TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return o.ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(context.Background())
}

func (o TimeSeriesInsightsGen2EnvironmentOutput) ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeSeriesInsightsGen2Environment) *TimeSeriesInsightsGen2Environment {
		return &v
	}).(TimeSeriesInsightsGen2EnvironmentPtrOutput)
}

type TimeSeriesInsightsGen2EnvironmentPtrOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsGen2EnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesInsightsGen2Environment)(nil))
}

func (o TimeSeriesInsightsGen2EnvironmentPtrOutput) ToTimeSeriesInsightsGen2EnvironmentPtrOutput() TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentPtrOutput) ToTimeSeriesInsightsGen2EnvironmentPtrOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentPtrOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentPtrOutput) Elem() TimeSeriesInsightsGen2EnvironmentOutput {
	return o.ApplyT(func(v *TimeSeriesInsightsGen2Environment) TimeSeriesInsightsGen2Environment {
		if v != nil {
			return *v
		}
		var ret TimeSeriesInsightsGen2Environment
		return ret
	}).(TimeSeriesInsightsGen2EnvironmentOutput)
}

type TimeSeriesInsightsGen2EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsGen2EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeSeriesInsightsGen2Environment)(nil))
}

func (o TimeSeriesInsightsGen2EnvironmentArrayOutput) ToTimeSeriesInsightsGen2EnvironmentArrayOutput() TimeSeriesInsightsGen2EnvironmentArrayOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentArrayOutput) ToTimeSeriesInsightsGen2EnvironmentArrayOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentArrayOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentArrayOutput) Index(i pulumi.IntInput) TimeSeriesInsightsGen2EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeSeriesInsightsGen2Environment {
		return vs[0].([]TimeSeriesInsightsGen2Environment)[vs[1].(int)]
	}).(TimeSeriesInsightsGen2EnvironmentOutput)
}

type TimeSeriesInsightsGen2EnvironmentMapOutput struct{ *pulumi.OutputState }

func (TimeSeriesInsightsGen2EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TimeSeriesInsightsGen2Environment)(nil))
}

func (o TimeSeriesInsightsGen2EnvironmentMapOutput) ToTimeSeriesInsightsGen2EnvironmentMapOutput() TimeSeriesInsightsGen2EnvironmentMapOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentMapOutput) ToTimeSeriesInsightsGen2EnvironmentMapOutputWithContext(ctx context.Context) TimeSeriesInsightsGen2EnvironmentMapOutput {
	return o
}

func (o TimeSeriesInsightsGen2EnvironmentMapOutput) MapIndex(k pulumi.StringInput) TimeSeriesInsightsGen2EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TimeSeriesInsightsGen2Environment {
		return vs[0].(map[string]TimeSeriesInsightsGen2Environment)[vs[1].(string)]
	}).(TimeSeriesInsightsGen2EnvironmentOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesInsightsGen2EnvironmentOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsGen2EnvironmentPtrOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsGen2EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(TimeSeriesInsightsGen2EnvironmentMapOutput{})
}
