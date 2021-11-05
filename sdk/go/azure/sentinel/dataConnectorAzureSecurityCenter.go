// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Azure Security Center Data Connector.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sentinel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("west europe"),
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
// 		exampleAnalyticsSolution, err := operationalinsights.NewAnalyticsSolution(ctx, "exampleAnalyticsSolution", &operationalinsights.AnalyticsSolutionArgs{
// 			SolutionName:        pulumi.String("SecurityInsights"),
// 			Location:            exampleResourceGroup.Location,
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			WorkspaceResourceId: exampleAnalyticsWorkspace.ID(),
// 			WorkspaceName:       exampleAnalyticsWorkspace.Name,
// 			Plan: &operationalinsights.AnalyticsSolutionPlanArgs{
// 				Publisher: pulumi.String("Microsoft"),
// 				Product:   pulumi.String("OMSGallery/SecurityInsights"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sentinel.NewDataConnectorAzureSecurityCenter(ctx, "exampleDataConnectorAzureSecurityCenter", &sentinel.DataConnectorAzureSecurityCenterArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsSolution.WorkspaceResourceId,
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
// Azure Security Center Data Connectors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/dataConnectorAzureSecurityCenter:DataConnectorAzureSecurityCenter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
// ```
type DataConnectorAzureSecurityCenter struct {
	pulumi.CustomResourceState

	// The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
}

// NewDataConnectorAzureSecurityCenter registers a new resource with the given unique name, arguments, and options.
func NewDataConnectorAzureSecurityCenter(ctx *pulumi.Context,
	name string, args *DataConnectorAzureSecurityCenterArgs, opts ...pulumi.ResourceOption) (*DataConnectorAzureSecurityCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	var resource DataConnectorAzureSecurityCenter
	err := ctx.RegisterResource("azure:sentinel/dataConnectorAzureSecurityCenter:DataConnectorAzureSecurityCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataConnectorAzureSecurityCenter gets an existing DataConnectorAzureSecurityCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataConnectorAzureSecurityCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorAzureSecurityCenterState, opts ...pulumi.ResourceOption) (*DataConnectorAzureSecurityCenter, error) {
	var resource DataConnectorAzureSecurityCenter
	err := ctx.ReadResource("azure:sentinel/dataConnectorAzureSecurityCenter:DataConnectorAzureSecurityCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataConnectorAzureSecurityCenter resources.
type dataConnectorAzureSecurityCenterState struct {
	// The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

type DataConnectorAzureSecurityCenterState struct {
	// The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
	SubscriptionId pulumi.StringPtrInput
}

func (DataConnectorAzureSecurityCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorAzureSecurityCenterState)(nil)).Elem()
}

type dataConnectorAzureSecurityCenterArgs struct {
	// The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a DataConnectorAzureSecurityCenter resource.
type DataConnectorAzureSecurityCenterArgs struct {
	// The ID of the Log Analytics Workspace that this Azure Security Center Data Connector resides in. Changing this forces a new Azure Security Center Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Azure Security Center Data Connector. Changing this forces a new Azure Security Center Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the subscription that this Azure Security Center Data Connector connects to. Changing this forces a new Azure Security Center Data Connector to be created.
	SubscriptionId pulumi.StringPtrInput
}

func (DataConnectorAzureSecurityCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorAzureSecurityCenterArgs)(nil)).Elem()
}

type DataConnectorAzureSecurityCenterInput interface {
	pulumi.Input

	ToDataConnectorAzureSecurityCenterOutput() DataConnectorAzureSecurityCenterOutput
	ToDataConnectorAzureSecurityCenterOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterOutput
}

func (*DataConnectorAzureSecurityCenter) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorAzureSecurityCenter)(nil))
}

func (i *DataConnectorAzureSecurityCenter) ToDataConnectorAzureSecurityCenterOutput() DataConnectorAzureSecurityCenterOutput {
	return i.ToDataConnectorAzureSecurityCenterOutputWithContext(context.Background())
}

func (i *DataConnectorAzureSecurityCenter) ToDataConnectorAzureSecurityCenterOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureSecurityCenterOutput)
}

func (i *DataConnectorAzureSecurityCenter) ToDataConnectorAzureSecurityCenterPtrOutput() DataConnectorAzureSecurityCenterPtrOutput {
	return i.ToDataConnectorAzureSecurityCenterPtrOutputWithContext(context.Background())
}

func (i *DataConnectorAzureSecurityCenter) ToDataConnectorAzureSecurityCenterPtrOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureSecurityCenterPtrOutput)
}

type DataConnectorAzureSecurityCenterPtrInput interface {
	pulumi.Input

	ToDataConnectorAzureSecurityCenterPtrOutput() DataConnectorAzureSecurityCenterPtrOutput
	ToDataConnectorAzureSecurityCenterPtrOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterPtrOutput
}

type dataConnectorAzureSecurityCenterPtrType DataConnectorAzureSecurityCenterArgs

func (*dataConnectorAzureSecurityCenterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorAzureSecurityCenter)(nil))
}

func (i *dataConnectorAzureSecurityCenterPtrType) ToDataConnectorAzureSecurityCenterPtrOutput() DataConnectorAzureSecurityCenterPtrOutput {
	return i.ToDataConnectorAzureSecurityCenterPtrOutputWithContext(context.Background())
}

func (i *dataConnectorAzureSecurityCenterPtrType) ToDataConnectorAzureSecurityCenterPtrOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureSecurityCenterPtrOutput)
}

// DataConnectorAzureSecurityCenterArrayInput is an input type that accepts DataConnectorAzureSecurityCenterArray and DataConnectorAzureSecurityCenterArrayOutput values.
// You can construct a concrete instance of `DataConnectorAzureSecurityCenterArrayInput` via:
//
//          DataConnectorAzureSecurityCenterArray{ DataConnectorAzureSecurityCenterArgs{...} }
type DataConnectorAzureSecurityCenterArrayInput interface {
	pulumi.Input

	ToDataConnectorAzureSecurityCenterArrayOutput() DataConnectorAzureSecurityCenterArrayOutput
	ToDataConnectorAzureSecurityCenterArrayOutputWithContext(context.Context) DataConnectorAzureSecurityCenterArrayOutput
}

type DataConnectorAzureSecurityCenterArray []DataConnectorAzureSecurityCenterInput

func (DataConnectorAzureSecurityCenterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataConnectorAzureSecurityCenter)(nil)).Elem()
}

func (i DataConnectorAzureSecurityCenterArray) ToDataConnectorAzureSecurityCenterArrayOutput() DataConnectorAzureSecurityCenterArrayOutput {
	return i.ToDataConnectorAzureSecurityCenterArrayOutputWithContext(context.Background())
}

func (i DataConnectorAzureSecurityCenterArray) ToDataConnectorAzureSecurityCenterArrayOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureSecurityCenterArrayOutput)
}

// DataConnectorAzureSecurityCenterMapInput is an input type that accepts DataConnectorAzureSecurityCenterMap and DataConnectorAzureSecurityCenterMapOutput values.
// You can construct a concrete instance of `DataConnectorAzureSecurityCenterMapInput` via:
//
//          DataConnectorAzureSecurityCenterMap{ "key": DataConnectorAzureSecurityCenterArgs{...} }
type DataConnectorAzureSecurityCenterMapInput interface {
	pulumi.Input

	ToDataConnectorAzureSecurityCenterMapOutput() DataConnectorAzureSecurityCenterMapOutput
	ToDataConnectorAzureSecurityCenterMapOutputWithContext(context.Context) DataConnectorAzureSecurityCenterMapOutput
}

type DataConnectorAzureSecurityCenterMap map[string]DataConnectorAzureSecurityCenterInput

func (DataConnectorAzureSecurityCenterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataConnectorAzureSecurityCenter)(nil)).Elem()
}

func (i DataConnectorAzureSecurityCenterMap) ToDataConnectorAzureSecurityCenterMapOutput() DataConnectorAzureSecurityCenterMapOutput {
	return i.ToDataConnectorAzureSecurityCenterMapOutputWithContext(context.Background())
}

func (i DataConnectorAzureSecurityCenterMap) ToDataConnectorAzureSecurityCenterMapOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorAzureSecurityCenterMapOutput)
}

type DataConnectorAzureSecurityCenterOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureSecurityCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorAzureSecurityCenter)(nil))
}

func (o DataConnectorAzureSecurityCenterOutput) ToDataConnectorAzureSecurityCenterOutput() DataConnectorAzureSecurityCenterOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterOutput) ToDataConnectorAzureSecurityCenterOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterOutput) ToDataConnectorAzureSecurityCenterPtrOutput() DataConnectorAzureSecurityCenterPtrOutput {
	return o.ToDataConnectorAzureSecurityCenterPtrOutputWithContext(context.Background())
}

func (o DataConnectorAzureSecurityCenterOutput) ToDataConnectorAzureSecurityCenterPtrOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataConnectorAzureSecurityCenter) *DataConnectorAzureSecurityCenter {
		return &v
	}).(DataConnectorAzureSecurityCenterPtrOutput)
}

type DataConnectorAzureSecurityCenterPtrOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureSecurityCenterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorAzureSecurityCenter)(nil))
}

func (o DataConnectorAzureSecurityCenterPtrOutput) ToDataConnectorAzureSecurityCenterPtrOutput() DataConnectorAzureSecurityCenterPtrOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterPtrOutput) ToDataConnectorAzureSecurityCenterPtrOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterPtrOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterPtrOutput) Elem() DataConnectorAzureSecurityCenterOutput {
	return o.ApplyT(func(v *DataConnectorAzureSecurityCenter) DataConnectorAzureSecurityCenter {
		if v != nil {
			return *v
		}
		var ret DataConnectorAzureSecurityCenter
		return ret
	}).(DataConnectorAzureSecurityCenterOutput)
}

type DataConnectorAzureSecurityCenterArrayOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureSecurityCenterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataConnectorAzureSecurityCenter)(nil))
}

func (o DataConnectorAzureSecurityCenterArrayOutput) ToDataConnectorAzureSecurityCenterArrayOutput() DataConnectorAzureSecurityCenterArrayOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterArrayOutput) ToDataConnectorAzureSecurityCenterArrayOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterArrayOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterArrayOutput) Index(i pulumi.IntInput) DataConnectorAzureSecurityCenterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataConnectorAzureSecurityCenter {
		return vs[0].([]DataConnectorAzureSecurityCenter)[vs[1].(int)]
	}).(DataConnectorAzureSecurityCenterOutput)
}

type DataConnectorAzureSecurityCenterMapOutput struct{ *pulumi.OutputState }

func (DataConnectorAzureSecurityCenterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataConnectorAzureSecurityCenter)(nil))
}

func (o DataConnectorAzureSecurityCenterMapOutput) ToDataConnectorAzureSecurityCenterMapOutput() DataConnectorAzureSecurityCenterMapOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterMapOutput) ToDataConnectorAzureSecurityCenterMapOutputWithContext(ctx context.Context) DataConnectorAzureSecurityCenterMapOutput {
	return o
}

func (o DataConnectorAzureSecurityCenterMapOutput) MapIndex(k pulumi.StringInput) DataConnectorAzureSecurityCenterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataConnectorAzureSecurityCenter {
		return vs[0].(map[string]DataConnectorAzureSecurityCenter)[vs[1].(string)]
	}).(DataConnectorAzureSecurityCenterOutput)
}

func init() {
	pulumi.RegisterOutputType(DataConnectorAzureSecurityCenterOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureSecurityCenterPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureSecurityCenterArrayOutput{})
	pulumi.RegisterOutputType(DataConnectorAzureSecurityCenterMapOutput{})
}
