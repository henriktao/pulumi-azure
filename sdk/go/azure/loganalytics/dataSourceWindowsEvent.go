// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Log Analytics Windows Event DataSource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/loganalytics"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loganalytics.NewDataSourceWindowsEvent(ctx, "exampleDataSourceWindowsEvent", &loganalytics.DataSourceWindowsEventArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			WorkspaceName:     exampleAnalyticsWorkspace.Name,
// 			EventLogName:      pulumi.String("Application"),
// 			EventTypes: pulumi.StringArray{
// 				pulumi.String("error"),
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
// Log Analytics Windows Event DataSources can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:loganalytics/dataSourceWindowsEvent:DataSourceWindowsEvent example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/datasources/datasource1
// ```
type DataSourceWindowsEvent struct {
	pulumi.CustomResourceState

	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName pulumi.StringOutput `pulumi:"eventLogName"`
	// Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName pulumi.StringOutput `pulumi:"workspaceName"`
}

// NewDataSourceWindowsEvent registers a new resource with the given unique name, arguments, and options.
func NewDataSourceWindowsEvent(ctx *pulumi.Context,
	name string, args *DataSourceWindowsEventArgs, opts ...pulumi.ResourceOption) (*DataSourceWindowsEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventLogName == nil {
		return nil, errors.New("invalid value for required argument 'EventLogName'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	var resource DataSourceWindowsEvent
	err := ctx.RegisterResource("azure:loganalytics/dataSourceWindowsEvent:DataSourceWindowsEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSourceWindowsEvent gets an existing DataSourceWindowsEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSourceWindowsEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceWindowsEventState, opts ...pulumi.ResourceOption) (*DataSourceWindowsEvent, error) {
	var resource DataSourceWindowsEvent
	err := ctx.ReadResource("azure:loganalytics/dataSourceWindowsEvent:DataSourceWindowsEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSourceWindowsEvent resources.
type dataSourceWindowsEventState struct {
	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName *string `pulumi:"eventLogName"`
	// Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
	EventTypes []string `pulumi:"eventTypes"`
	// The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName *string `pulumi:"workspaceName"`
}

type DataSourceWindowsEventState struct {
	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName pulumi.StringPtrInput
	// Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
	EventTypes pulumi.StringArrayInput
	// The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName pulumi.StringPtrInput
}

func (DataSourceWindowsEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceWindowsEventState)(nil)).Elem()
}

type dataSourceWindowsEventArgs struct {
	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName string `pulumi:"eventLogName"`
	// Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
	EventTypes []string `pulumi:"eventTypes"`
	// The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataSourceWindowsEvent resource.
type DataSourceWindowsEventArgs struct {
	// Specifies the name of the Windows Event Log to collect events from.
	EventLogName pulumi.StringInput
	// Specifies an array of event types applied to the specified event log. Possible values include `error`, `warning` and `information`.
	EventTypes pulumi.StringArrayInput
	// The name which should be used for this Log Analytics Windows Event DataSource. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Log Analytics Workspace where the Log Analytics Windows Event DataSource should exist. Changing this forces a new Log Analytics Windows Event DataSource to be created.
	WorkspaceName pulumi.StringInput
}

func (DataSourceWindowsEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceWindowsEventArgs)(nil)).Elem()
}

type DataSourceWindowsEventInput interface {
	pulumi.Input

	ToDataSourceWindowsEventOutput() DataSourceWindowsEventOutput
	ToDataSourceWindowsEventOutputWithContext(ctx context.Context) DataSourceWindowsEventOutput
}

func (*DataSourceWindowsEvent) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceWindowsEvent)(nil)).Elem()
}

func (i *DataSourceWindowsEvent) ToDataSourceWindowsEventOutput() DataSourceWindowsEventOutput {
	return i.ToDataSourceWindowsEventOutputWithContext(context.Background())
}

func (i *DataSourceWindowsEvent) ToDataSourceWindowsEventOutputWithContext(ctx context.Context) DataSourceWindowsEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceWindowsEventOutput)
}

// DataSourceWindowsEventArrayInput is an input type that accepts DataSourceWindowsEventArray and DataSourceWindowsEventArrayOutput values.
// You can construct a concrete instance of `DataSourceWindowsEventArrayInput` via:
//
//          DataSourceWindowsEventArray{ DataSourceWindowsEventArgs{...} }
type DataSourceWindowsEventArrayInput interface {
	pulumi.Input

	ToDataSourceWindowsEventArrayOutput() DataSourceWindowsEventArrayOutput
	ToDataSourceWindowsEventArrayOutputWithContext(context.Context) DataSourceWindowsEventArrayOutput
}

type DataSourceWindowsEventArray []DataSourceWindowsEventInput

func (DataSourceWindowsEventArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourceWindowsEvent)(nil)).Elem()
}

func (i DataSourceWindowsEventArray) ToDataSourceWindowsEventArrayOutput() DataSourceWindowsEventArrayOutput {
	return i.ToDataSourceWindowsEventArrayOutputWithContext(context.Background())
}

func (i DataSourceWindowsEventArray) ToDataSourceWindowsEventArrayOutputWithContext(ctx context.Context) DataSourceWindowsEventArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceWindowsEventArrayOutput)
}

// DataSourceWindowsEventMapInput is an input type that accepts DataSourceWindowsEventMap and DataSourceWindowsEventMapOutput values.
// You can construct a concrete instance of `DataSourceWindowsEventMapInput` via:
//
//          DataSourceWindowsEventMap{ "key": DataSourceWindowsEventArgs{...} }
type DataSourceWindowsEventMapInput interface {
	pulumi.Input

	ToDataSourceWindowsEventMapOutput() DataSourceWindowsEventMapOutput
	ToDataSourceWindowsEventMapOutputWithContext(context.Context) DataSourceWindowsEventMapOutput
}

type DataSourceWindowsEventMap map[string]DataSourceWindowsEventInput

func (DataSourceWindowsEventMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourceWindowsEvent)(nil)).Elem()
}

func (i DataSourceWindowsEventMap) ToDataSourceWindowsEventMapOutput() DataSourceWindowsEventMapOutput {
	return i.ToDataSourceWindowsEventMapOutputWithContext(context.Background())
}

func (i DataSourceWindowsEventMap) ToDataSourceWindowsEventMapOutputWithContext(ctx context.Context) DataSourceWindowsEventMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceWindowsEventMapOutput)
}

type DataSourceWindowsEventOutput struct{ *pulumi.OutputState }

func (DataSourceWindowsEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceWindowsEvent)(nil)).Elem()
}

func (o DataSourceWindowsEventOutput) ToDataSourceWindowsEventOutput() DataSourceWindowsEventOutput {
	return o
}

func (o DataSourceWindowsEventOutput) ToDataSourceWindowsEventOutputWithContext(ctx context.Context) DataSourceWindowsEventOutput {
	return o
}

type DataSourceWindowsEventArrayOutput struct{ *pulumi.OutputState }

func (DataSourceWindowsEventArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourceWindowsEvent)(nil)).Elem()
}

func (o DataSourceWindowsEventArrayOutput) ToDataSourceWindowsEventArrayOutput() DataSourceWindowsEventArrayOutput {
	return o
}

func (o DataSourceWindowsEventArrayOutput) ToDataSourceWindowsEventArrayOutputWithContext(ctx context.Context) DataSourceWindowsEventArrayOutput {
	return o
}

func (o DataSourceWindowsEventArrayOutput) Index(i pulumi.IntInput) DataSourceWindowsEventOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSourceWindowsEvent {
		return vs[0].([]*DataSourceWindowsEvent)[vs[1].(int)]
	}).(DataSourceWindowsEventOutput)
}

type DataSourceWindowsEventMapOutput struct{ *pulumi.OutputState }

func (DataSourceWindowsEventMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourceWindowsEvent)(nil)).Elem()
}

func (o DataSourceWindowsEventMapOutput) ToDataSourceWindowsEventMapOutput() DataSourceWindowsEventMapOutput {
	return o
}

func (o DataSourceWindowsEventMapOutput) ToDataSourceWindowsEventMapOutputWithContext(ctx context.Context) DataSourceWindowsEventMapOutput {
	return o
}

func (o DataSourceWindowsEventMapOutput) MapIndex(k pulumi.StringInput) DataSourceWindowsEventOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSourceWindowsEvent {
		return vs[0].(map[string]*DataSourceWindowsEvent)[vs[1].(string)]
	}).(DataSourceWindowsEventOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceWindowsEventInput)(nil)).Elem(), &DataSourceWindowsEvent{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceWindowsEventArrayInput)(nil)).Elem(), DataSourceWindowsEventArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceWindowsEventMapInput)(nil)).Elem(), DataSourceWindowsEventMap{})
	pulumi.RegisterOutputType(DataSourceWindowsEventOutput{})
	pulumi.RegisterOutputType(DataSourceWindowsEventArrayOutput{})
	pulumi.RegisterOutputType(DataSourceWindowsEventMapOutput{})
}
