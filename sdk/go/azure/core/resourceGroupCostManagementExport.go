// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Cost Management Export for a Resource Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		_, err = storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName: pulumi.Any(azurerm_storage_account.Test.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewResourceGroupCostManagementExport(ctx, "exampleResourceGroupCostManagementExport", &core.ResourceGroupCostManagementExportArgs{
// 			ResourceGroupId:           exampleResourceGroup.ID(),
// 			RecurrenceType:            pulumi.String("Monthly"),
// 			RecurrencePeriodStartDate: pulumi.String("2020-08-18T00:00:00Z"),
// 			RecurrencePeriodEndDate:   pulumi.String("2020-09-18T00:00:00Z"),
// 			ExportDataStorageLocation: &core.ResourceGroupCostManagementExportExportDataStorageLocationArgs{
// 				ContainerId:    pulumi.Any(azurerm_storage_container.Test.Resource_manager_id),
// 				RootFolderPath: pulumi.String("/root/updated"),
// 			},
// 			ExportDataOptions: &core.ResourceGroupCostManagementExportExportDataOptionsArgs{
// 				Type:      pulumi.String("Usage"),
// 				TimeFrame: pulumi.String("WeekToDate"),
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
// Cost Management Export for a Resource Group can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/resourceGroupCostManagementExport:ResourceGroupCostManagementExport example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.CostManagement/exports/export1
// ```
type ResourceGroupCostManagementExport struct {
	pulumi.CustomResourceState

	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions ResourceGroupCostManagementExportExportDataOptionsOutput `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation ResourceGroupCostManagementExportExportDataStorageLocationOutput `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The date the export will stop capturing information.
	RecurrencePeriodEndDate pulumi.StringOutput `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringOutput `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringOutput `pulumi:"recurrenceType"`
	// The id of the resource group on which to create an export.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
}

// NewResourceGroupCostManagementExport registers a new resource with the given unique name, arguments, and options.
func NewResourceGroupCostManagementExport(ctx *pulumi.Context,
	name string, args *ResourceGroupCostManagementExportArgs, opts ...pulumi.ResourceOption) (*ResourceGroupCostManagementExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExportDataOptions == nil {
		return nil, errors.New("invalid value for required argument 'ExportDataOptions'")
	}
	if args.ExportDataStorageLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExportDataStorageLocation'")
	}
	if args.RecurrencePeriodEndDate == nil {
		return nil, errors.New("invalid value for required argument 'RecurrencePeriodEndDate'")
	}
	if args.RecurrencePeriodStartDate == nil {
		return nil, errors.New("invalid value for required argument 'RecurrencePeriodStartDate'")
	}
	if args.RecurrenceType == nil {
		return nil, errors.New("invalid value for required argument 'RecurrenceType'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	var resource ResourceGroupCostManagementExport
	err := ctx.RegisterResource("azure:core/resourceGroupCostManagementExport:ResourceGroupCostManagementExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroupCostManagementExport gets an existing ResourceGroupCostManagementExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroupCostManagementExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupCostManagementExportState, opts ...pulumi.ResourceOption) (*ResourceGroupCostManagementExport, error) {
	var resource ResourceGroupCostManagementExport
	err := ctx.ReadResource("azure:core/resourceGroupCostManagementExport:ResourceGroupCostManagementExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroupCostManagementExport resources.
type resourceGroupCostManagementExportState struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions *ResourceGroupCostManagementExportExportDataOptions `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation *ResourceGroupCostManagementExportExportDataStorageLocation `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The date the export will stop capturing information.
	RecurrencePeriodEndDate *string `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate *string `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType *string `pulumi:"recurrenceType"`
	// The id of the resource group on which to create an export.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type ResourceGroupCostManagementExportState struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `exportDataOptions` block as defined below.
	ExportDataOptions ResourceGroupCostManagementExportExportDataOptionsPtrInput
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation ResourceGroupCostManagementExportExportDataStorageLocationPtrInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The date the export will stop capturing information.
	RecurrencePeriodEndDate pulumi.StringPtrInput
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringPtrInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringPtrInput
	// The id of the resource group on which to create an export.
	ResourceGroupId pulumi.StringPtrInput
}

func (ResourceGroupCostManagementExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupCostManagementExportState)(nil)).Elem()
}

type resourceGroupCostManagementExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions ResourceGroupCostManagementExportExportDataOptions `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation ResourceGroupCostManagementExportExportDataStorageLocation `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The date the export will stop capturing information.
	RecurrencePeriodEndDate string `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate string `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType string `pulumi:"recurrenceType"`
	// The id of the resource group on which to create an export.
	ResourceGroupId string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a ResourceGroupCostManagementExport resource.
type ResourceGroupCostManagementExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `exportDataOptions` block as defined below.
	ExportDataOptions ResourceGroupCostManagementExportExportDataOptionsInput
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation ResourceGroupCostManagementExportExportDataStorageLocationInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The date the export will stop capturing information.
	RecurrencePeriodEndDate pulumi.StringInput
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringInput
	// The id of the resource group on which to create an export.
	ResourceGroupId pulumi.StringInput
}

func (ResourceGroupCostManagementExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupCostManagementExportArgs)(nil)).Elem()
}

type ResourceGroupCostManagementExportInput interface {
	pulumi.Input

	ToResourceGroupCostManagementExportOutput() ResourceGroupCostManagementExportOutput
	ToResourceGroupCostManagementExportOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportOutput
}

func (*ResourceGroupCostManagementExport) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupCostManagementExport)(nil))
}

func (i *ResourceGroupCostManagementExport) ToResourceGroupCostManagementExportOutput() ResourceGroupCostManagementExportOutput {
	return i.ToResourceGroupCostManagementExportOutputWithContext(context.Background())
}

func (i *ResourceGroupCostManagementExport) ToResourceGroupCostManagementExportOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupCostManagementExportOutput)
}

func (i *ResourceGroupCostManagementExport) ToResourceGroupCostManagementExportPtrOutput() ResourceGroupCostManagementExportPtrOutput {
	return i.ToResourceGroupCostManagementExportPtrOutputWithContext(context.Background())
}

func (i *ResourceGroupCostManagementExport) ToResourceGroupCostManagementExportPtrOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupCostManagementExportPtrOutput)
}

type ResourceGroupCostManagementExportPtrInput interface {
	pulumi.Input

	ToResourceGroupCostManagementExportPtrOutput() ResourceGroupCostManagementExportPtrOutput
	ToResourceGroupCostManagementExportPtrOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportPtrOutput
}

type resourceGroupCostManagementExportPtrType ResourceGroupCostManagementExportArgs

func (*resourceGroupCostManagementExportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupCostManagementExport)(nil))
}

func (i *resourceGroupCostManagementExportPtrType) ToResourceGroupCostManagementExportPtrOutput() ResourceGroupCostManagementExportPtrOutput {
	return i.ToResourceGroupCostManagementExportPtrOutputWithContext(context.Background())
}

func (i *resourceGroupCostManagementExportPtrType) ToResourceGroupCostManagementExportPtrOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupCostManagementExportPtrOutput)
}

// ResourceGroupCostManagementExportArrayInput is an input type that accepts ResourceGroupCostManagementExportArray and ResourceGroupCostManagementExportArrayOutput values.
// You can construct a concrete instance of `ResourceGroupCostManagementExportArrayInput` via:
//
//          ResourceGroupCostManagementExportArray{ ResourceGroupCostManagementExportArgs{...} }
type ResourceGroupCostManagementExportArrayInput interface {
	pulumi.Input

	ToResourceGroupCostManagementExportArrayOutput() ResourceGroupCostManagementExportArrayOutput
	ToResourceGroupCostManagementExportArrayOutputWithContext(context.Context) ResourceGroupCostManagementExportArrayOutput
}

type ResourceGroupCostManagementExportArray []ResourceGroupCostManagementExportInput

func (ResourceGroupCostManagementExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceGroupCostManagementExport)(nil)).Elem()
}

func (i ResourceGroupCostManagementExportArray) ToResourceGroupCostManagementExportArrayOutput() ResourceGroupCostManagementExportArrayOutput {
	return i.ToResourceGroupCostManagementExportArrayOutputWithContext(context.Background())
}

func (i ResourceGroupCostManagementExportArray) ToResourceGroupCostManagementExportArrayOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupCostManagementExportArrayOutput)
}

// ResourceGroupCostManagementExportMapInput is an input type that accepts ResourceGroupCostManagementExportMap and ResourceGroupCostManagementExportMapOutput values.
// You can construct a concrete instance of `ResourceGroupCostManagementExportMapInput` via:
//
//          ResourceGroupCostManagementExportMap{ "key": ResourceGroupCostManagementExportArgs{...} }
type ResourceGroupCostManagementExportMapInput interface {
	pulumi.Input

	ToResourceGroupCostManagementExportMapOutput() ResourceGroupCostManagementExportMapOutput
	ToResourceGroupCostManagementExportMapOutputWithContext(context.Context) ResourceGroupCostManagementExportMapOutput
}

type ResourceGroupCostManagementExportMap map[string]ResourceGroupCostManagementExportInput

func (ResourceGroupCostManagementExportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceGroupCostManagementExport)(nil)).Elem()
}

func (i ResourceGroupCostManagementExportMap) ToResourceGroupCostManagementExportMapOutput() ResourceGroupCostManagementExportMapOutput {
	return i.ToResourceGroupCostManagementExportMapOutputWithContext(context.Background())
}

func (i ResourceGroupCostManagementExportMap) ToResourceGroupCostManagementExportMapOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupCostManagementExportMapOutput)
}

type ResourceGroupCostManagementExportOutput struct{ *pulumi.OutputState }

func (ResourceGroupCostManagementExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupCostManagementExport)(nil))
}

func (o ResourceGroupCostManagementExportOutput) ToResourceGroupCostManagementExportOutput() ResourceGroupCostManagementExportOutput {
	return o
}

func (o ResourceGroupCostManagementExportOutput) ToResourceGroupCostManagementExportOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportOutput {
	return o
}

func (o ResourceGroupCostManagementExportOutput) ToResourceGroupCostManagementExportPtrOutput() ResourceGroupCostManagementExportPtrOutput {
	return o.ToResourceGroupCostManagementExportPtrOutputWithContext(context.Background())
}

func (o ResourceGroupCostManagementExportOutput) ToResourceGroupCostManagementExportPtrOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceGroupCostManagementExport) *ResourceGroupCostManagementExport {
		return &v
	}).(ResourceGroupCostManagementExportPtrOutput)
}

type ResourceGroupCostManagementExportPtrOutput struct{ *pulumi.OutputState }

func (ResourceGroupCostManagementExportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupCostManagementExport)(nil))
}

func (o ResourceGroupCostManagementExportPtrOutput) ToResourceGroupCostManagementExportPtrOutput() ResourceGroupCostManagementExportPtrOutput {
	return o
}

func (o ResourceGroupCostManagementExportPtrOutput) ToResourceGroupCostManagementExportPtrOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportPtrOutput {
	return o
}

func (o ResourceGroupCostManagementExportPtrOutput) Elem() ResourceGroupCostManagementExportOutput {
	return o.ApplyT(func(v *ResourceGroupCostManagementExport) ResourceGroupCostManagementExport {
		if v != nil {
			return *v
		}
		var ret ResourceGroupCostManagementExport
		return ret
	}).(ResourceGroupCostManagementExportOutput)
}

type ResourceGroupCostManagementExportArrayOutput struct{ *pulumi.OutputState }

func (ResourceGroupCostManagementExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGroupCostManagementExport)(nil))
}

func (o ResourceGroupCostManagementExportArrayOutput) ToResourceGroupCostManagementExportArrayOutput() ResourceGroupCostManagementExportArrayOutput {
	return o
}

func (o ResourceGroupCostManagementExportArrayOutput) ToResourceGroupCostManagementExportArrayOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportArrayOutput {
	return o
}

func (o ResourceGroupCostManagementExportArrayOutput) Index(i pulumi.IntInput) ResourceGroupCostManagementExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceGroupCostManagementExport {
		return vs[0].([]ResourceGroupCostManagementExport)[vs[1].(int)]
	}).(ResourceGroupCostManagementExportOutput)
}

type ResourceGroupCostManagementExportMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupCostManagementExportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupCostManagementExport)(nil))
}

func (o ResourceGroupCostManagementExportMapOutput) ToResourceGroupCostManagementExportMapOutput() ResourceGroupCostManagementExportMapOutput {
	return o
}

func (o ResourceGroupCostManagementExportMapOutput) ToResourceGroupCostManagementExportMapOutputWithContext(ctx context.Context) ResourceGroupCostManagementExportMapOutput {
	return o
}

func (o ResourceGroupCostManagementExportMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupCostManagementExportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupCostManagementExport {
		return vs[0].(map[string]ResourceGroupCostManagementExport)[vs[1].(string)]
	}).(ResourceGroupCostManagementExportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupCostManagementExportInput)(nil)).Elem(), &ResourceGroupCostManagementExport{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupCostManagementExportPtrInput)(nil)).Elem(), &ResourceGroupCostManagementExport{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupCostManagementExportArrayInput)(nil)).Elem(), ResourceGroupCostManagementExportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupCostManagementExportMapInput)(nil)).Elem(), ResourceGroupCostManagementExportMap{})
	pulumi.RegisterOutputType(ResourceGroupCostManagementExportOutput{})
	pulumi.RegisterOutputType(ResourceGroupCostManagementExportPtrOutput{})
	pulumi.RegisterOutputType(ResourceGroupCostManagementExportArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupCostManagementExportMapOutput{})
}
