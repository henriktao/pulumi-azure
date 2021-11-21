// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Cost Management Export for a Subscription.
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
// 		_, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
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
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName: pulumi.Any(azurerm_storage_account.Test.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewSubscriptionCostManagementExport(ctx, "exampleSubscriptionCostManagementExport", &core.SubscriptionCostManagementExportArgs{
// 			SubscriptionId:            pulumi.Any(azurerm_subscription.Example.Id),
// 			RecurrenceType:            pulumi.String("Monthly"),
// 			RecurrencePeriodStartDate: pulumi.String("2020-08-18T00:00:00Z"),
// 			RecurrencePeriodEndDate:   pulumi.String("2020-09-18T00:00:00Z"),
// 			ExportDataStorageLocation: &core.SubscriptionCostManagementExportExportDataStorageLocationArgs{
// 				ContainerId:    exampleContainer.ResourceManagerId,
// 				RootFolderPath: pulumi.String("/root/updated"),
// 			},
// 			ExportDataOptions: &core.SubscriptionCostManagementExportExportDataOptionsArgs{
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
// Subscription Cost Management Exports can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:core/subscriptionCostManagementExport:SubscriptionCostManagementExport example /subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/exports/export1
// ```
type SubscriptionCostManagementExport struct {
	pulumi.CustomResourceState

	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions SubscriptionCostManagementExportExportDataOptionsOutput `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation SubscriptionCostManagementExportExportDataStorageLocationOutput `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name                    pulumi.StringOutput `pulumi:"name"`
	RecurrencePeriodEndDate pulumi.StringOutput `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringOutput `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringOutput `pulumi:"recurrenceType"`
	// The id of the subscription on which to create an export.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
}

// NewSubscriptionCostManagementExport registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionCostManagementExport(ctx *pulumi.Context,
	name string, args *SubscriptionCostManagementExportArgs, opts ...pulumi.ResourceOption) (*SubscriptionCostManagementExport, error) {
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
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	var resource SubscriptionCostManagementExport
	err := ctx.RegisterResource("azure:core/subscriptionCostManagementExport:SubscriptionCostManagementExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionCostManagementExport gets an existing SubscriptionCostManagementExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionCostManagementExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionCostManagementExportState, opts ...pulumi.ResourceOption) (*SubscriptionCostManagementExport, error) {
	var resource SubscriptionCostManagementExport
	err := ctx.ReadResource("azure:core/subscriptionCostManagementExport:SubscriptionCostManagementExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionCostManagementExport resources.
type subscriptionCostManagementExportState struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions *SubscriptionCostManagementExportExportDataOptions `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation *SubscriptionCostManagementExportExportDataStorageLocation `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name                    *string `pulumi:"name"`
	RecurrencePeriodEndDate *string `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate *string `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType *string `pulumi:"recurrenceType"`
	// The id of the subscription on which to create an export.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

type SubscriptionCostManagementExportState struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `exportDataOptions` block as defined below.
	ExportDataOptions SubscriptionCostManagementExportExportDataOptionsPtrInput
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation SubscriptionCostManagementExportExportDataStorageLocationPtrInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name                    pulumi.StringPtrInput
	RecurrencePeriodEndDate pulumi.StringPtrInput
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringPtrInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringPtrInput
	// The id of the subscription on which to create an export.
	SubscriptionId pulumi.StringPtrInput
}

func (SubscriptionCostManagementExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionCostManagementExportState)(nil)).Elem()
}

type subscriptionCostManagementExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active *bool `pulumi:"active"`
	// A `exportDataOptions` block as defined below.
	ExportDataOptions SubscriptionCostManagementExportExportDataOptions `pulumi:"exportDataOptions"`
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation SubscriptionCostManagementExportExportDataStorageLocation `pulumi:"exportDataStorageLocation"`
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name                    *string `pulumi:"name"`
	RecurrencePeriodEndDate string  `pulumi:"recurrencePeriodEndDate"`
	// The date the export will start capturing information.
	RecurrencePeriodStartDate string `pulumi:"recurrencePeriodStartDate"`
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType string `pulumi:"recurrenceType"`
	// The id of the subscription on which to create an export.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a SubscriptionCostManagementExport resource.
type SubscriptionCostManagementExportArgs struct {
	// Is the cost management export active? Default is `true`.
	Active pulumi.BoolPtrInput
	// A `exportDataOptions` block as defined below.
	ExportDataOptions SubscriptionCostManagementExportExportDataOptionsInput
	// A `exportDataStorageLocation` block as defined below.
	ExportDataStorageLocation SubscriptionCostManagementExportExportDataStorageLocationInput
	// Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
	Name                    pulumi.StringPtrInput
	RecurrencePeriodEndDate pulumi.StringInput
	// The date the export will start capturing information.
	RecurrencePeriodStartDate pulumi.StringInput
	// How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
	RecurrenceType pulumi.StringInput
	// The id of the subscription on which to create an export.
	SubscriptionId pulumi.StringInput
}

func (SubscriptionCostManagementExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionCostManagementExportArgs)(nil)).Elem()
}

type SubscriptionCostManagementExportInput interface {
	pulumi.Input

	ToSubscriptionCostManagementExportOutput() SubscriptionCostManagementExportOutput
	ToSubscriptionCostManagementExportOutputWithContext(ctx context.Context) SubscriptionCostManagementExportOutput
}

func (*SubscriptionCostManagementExport) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionCostManagementExport)(nil))
}

func (i *SubscriptionCostManagementExport) ToSubscriptionCostManagementExportOutput() SubscriptionCostManagementExportOutput {
	return i.ToSubscriptionCostManagementExportOutputWithContext(context.Background())
}

func (i *SubscriptionCostManagementExport) ToSubscriptionCostManagementExportOutputWithContext(ctx context.Context) SubscriptionCostManagementExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionCostManagementExportOutput)
}

func (i *SubscriptionCostManagementExport) ToSubscriptionCostManagementExportPtrOutput() SubscriptionCostManagementExportPtrOutput {
	return i.ToSubscriptionCostManagementExportPtrOutputWithContext(context.Background())
}

func (i *SubscriptionCostManagementExport) ToSubscriptionCostManagementExportPtrOutputWithContext(ctx context.Context) SubscriptionCostManagementExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionCostManagementExportPtrOutput)
}

type SubscriptionCostManagementExportPtrInput interface {
	pulumi.Input

	ToSubscriptionCostManagementExportPtrOutput() SubscriptionCostManagementExportPtrOutput
	ToSubscriptionCostManagementExportPtrOutputWithContext(ctx context.Context) SubscriptionCostManagementExportPtrOutput
}

type subscriptionCostManagementExportPtrType SubscriptionCostManagementExportArgs

func (*subscriptionCostManagementExportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionCostManagementExport)(nil))
}

func (i *subscriptionCostManagementExportPtrType) ToSubscriptionCostManagementExportPtrOutput() SubscriptionCostManagementExportPtrOutput {
	return i.ToSubscriptionCostManagementExportPtrOutputWithContext(context.Background())
}

func (i *subscriptionCostManagementExportPtrType) ToSubscriptionCostManagementExportPtrOutputWithContext(ctx context.Context) SubscriptionCostManagementExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionCostManagementExportPtrOutput)
}

// SubscriptionCostManagementExportArrayInput is an input type that accepts SubscriptionCostManagementExportArray and SubscriptionCostManagementExportArrayOutput values.
// You can construct a concrete instance of `SubscriptionCostManagementExportArrayInput` via:
//
//          SubscriptionCostManagementExportArray{ SubscriptionCostManagementExportArgs{...} }
type SubscriptionCostManagementExportArrayInput interface {
	pulumi.Input

	ToSubscriptionCostManagementExportArrayOutput() SubscriptionCostManagementExportArrayOutput
	ToSubscriptionCostManagementExportArrayOutputWithContext(context.Context) SubscriptionCostManagementExportArrayOutput
}

type SubscriptionCostManagementExportArray []SubscriptionCostManagementExportInput

func (SubscriptionCostManagementExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionCostManagementExport)(nil)).Elem()
}

func (i SubscriptionCostManagementExportArray) ToSubscriptionCostManagementExportArrayOutput() SubscriptionCostManagementExportArrayOutput {
	return i.ToSubscriptionCostManagementExportArrayOutputWithContext(context.Background())
}

func (i SubscriptionCostManagementExportArray) ToSubscriptionCostManagementExportArrayOutputWithContext(ctx context.Context) SubscriptionCostManagementExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionCostManagementExportArrayOutput)
}

// SubscriptionCostManagementExportMapInput is an input type that accepts SubscriptionCostManagementExportMap and SubscriptionCostManagementExportMapOutput values.
// You can construct a concrete instance of `SubscriptionCostManagementExportMapInput` via:
//
//          SubscriptionCostManagementExportMap{ "key": SubscriptionCostManagementExportArgs{...} }
type SubscriptionCostManagementExportMapInput interface {
	pulumi.Input

	ToSubscriptionCostManagementExportMapOutput() SubscriptionCostManagementExportMapOutput
	ToSubscriptionCostManagementExportMapOutputWithContext(context.Context) SubscriptionCostManagementExportMapOutput
}

type SubscriptionCostManagementExportMap map[string]SubscriptionCostManagementExportInput

func (SubscriptionCostManagementExportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionCostManagementExport)(nil)).Elem()
}

func (i SubscriptionCostManagementExportMap) ToSubscriptionCostManagementExportMapOutput() SubscriptionCostManagementExportMapOutput {
	return i.ToSubscriptionCostManagementExportMapOutputWithContext(context.Background())
}

func (i SubscriptionCostManagementExportMap) ToSubscriptionCostManagementExportMapOutputWithContext(ctx context.Context) SubscriptionCostManagementExportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionCostManagementExportMapOutput)
}

type SubscriptionCostManagementExportOutput struct{ *pulumi.OutputState }

func (SubscriptionCostManagementExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionCostManagementExport)(nil))
}

func (o SubscriptionCostManagementExportOutput) ToSubscriptionCostManagementExportOutput() SubscriptionCostManagementExportOutput {
	return o
}

func (o SubscriptionCostManagementExportOutput) ToSubscriptionCostManagementExportOutputWithContext(ctx context.Context) SubscriptionCostManagementExportOutput {
	return o
}

func (o SubscriptionCostManagementExportOutput) ToSubscriptionCostManagementExportPtrOutput() SubscriptionCostManagementExportPtrOutput {
	return o.ToSubscriptionCostManagementExportPtrOutputWithContext(context.Background())
}

func (o SubscriptionCostManagementExportOutput) ToSubscriptionCostManagementExportPtrOutputWithContext(ctx context.Context) SubscriptionCostManagementExportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionCostManagementExport) *SubscriptionCostManagementExport {
		return &v
	}).(SubscriptionCostManagementExportPtrOutput)
}

type SubscriptionCostManagementExportPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionCostManagementExportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionCostManagementExport)(nil))
}

func (o SubscriptionCostManagementExportPtrOutput) ToSubscriptionCostManagementExportPtrOutput() SubscriptionCostManagementExportPtrOutput {
	return o
}

func (o SubscriptionCostManagementExportPtrOutput) ToSubscriptionCostManagementExportPtrOutputWithContext(ctx context.Context) SubscriptionCostManagementExportPtrOutput {
	return o
}

func (o SubscriptionCostManagementExportPtrOutput) Elem() SubscriptionCostManagementExportOutput {
	return o.ApplyT(func(v *SubscriptionCostManagementExport) SubscriptionCostManagementExport {
		if v != nil {
			return *v
		}
		var ret SubscriptionCostManagementExport
		return ret
	}).(SubscriptionCostManagementExportOutput)
}

type SubscriptionCostManagementExportArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionCostManagementExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionCostManagementExport)(nil))
}

func (o SubscriptionCostManagementExportArrayOutput) ToSubscriptionCostManagementExportArrayOutput() SubscriptionCostManagementExportArrayOutput {
	return o
}

func (o SubscriptionCostManagementExportArrayOutput) ToSubscriptionCostManagementExportArrayOutputWithContext(ctx context.Context) SubscriptionCostManagementExportArrayOutput {
	return o
}

func (o SubscriptionCostManagementExportArrayOutput) Index(i pulumi.IntInput) SubscriptionCostManagementExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionCostManagementExport {
		return vs[0].([]SubscriptionCostManagementExport)[vs[1].(int)]
	}).(SubscriptionCostManagementExportOutput)
}

type SubscriptionCostManagementExportMapOutput struct{ *pulumi.OutputState }

func (SubscriptionCostManagementExportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SubscriptionCostManagementExport)(nil))
}

func (o SubscriptionCostManagementExportMapOutput) ToSubscriptionCostManagementExportMapOutput() SubscriptionCostManagementExportMapOutput {
	return o
}

func (o SubscriptionCostManagementExportMapOutput) ToSubscriptionCostManagementExportMapOutputWithContext(ctx context.Context) SubscriptionCostManagementExportMapOutput {
	return o
}

func (o SubscriptionCostManagementExportMapOutput) MapIndex(k pulumi.StringInput) SubscriptionCostManagementExportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SubscriptionCostManagementExport {
		return vs[0].(map[string]SubscriptionCostManagementExport)[vs[1].(string)]
	}).(SubscriptionCostManagementExportOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionCostManagementExportOutput{})
	pulumi.RegisterOutputType(SubscriptionCostManagementExportPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionCostManagementExportArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionCostManagementExportMapOutput{})
}
