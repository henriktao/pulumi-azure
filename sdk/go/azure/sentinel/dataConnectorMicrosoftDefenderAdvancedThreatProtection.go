// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Microsoft Defender Advanced Threat Protection Data Connector.
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
// 		_, err = sentinel.NewDataConnectorMicrosoftDefenderAdvancedThreatProtection(ctx, "exampleDataConnectorMicrosoftDefenderAdvancedThreatProtection", &sentinel.DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs{
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
// Microsoft Defender Advanced Threat Protection Data Connectors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/dataConnectorMicrosoftDefenderAdvancedThreatProtection:DataConnectorMicrosoftDefenderAdvancedThreatProtection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
// ```
type DataConnectorMicrosoftDefenderAdvancedThreatProtection struct {
	pulumi.CustomResourceState

	// The ID of the Log Analytics Workspace that this Microsoft Defender Advanced Threat Protection Data Connector resides in. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Microsoft Defender Advanced Threat Protection Data Connector. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the tenant that this Microsoft Defender Advanced Threat Protection Data Connector connects to. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewDataConnectorMicrosoftDefenderAdvancedThreatProtection registers a new resource with the given unique name, arguments, and options.
func NewDataConnectorMicrosoftDefenderAdvancedThreatProtection(ctx *pulumi.Context,
	name string, args *DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs, opts ...pulumi.ResourceOption) (*DataConnectorMicrosoftDefenderAdvancedThreatProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	var resource DataConnectorMicrosoftDefenderAdvancedThreatProtection
	err := ctx.RegisterResource("azure:sentinel/dataConnectorMicrosoftDefenderAdvancedThreatProtection:DataConnectorMicrosoftDefenderAdvancedThreatProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataConnectorMicrosoftDefenderAdvancedThreatProtection gets an existing DataConnectorMicrosoftDefenderAdvancedThreatProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataConnectorMicrosoftDefenderAdvancedThreatProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorMicrosoftDefenderAdvancedThreatProtectionState, opts ...pulumi.ResourceOption) (*DataConnectorMicrosoftDefenderAdvancedThreatProtection, error) {
	var resource DataConnectorMicrosoftDefenderAdvancedThreatProtection
	err := ctx.ReadResource("azure:sentinel/dataConnectorMicrosoftDefenderAdvancedThreatProtection:DataConnectorMicrosoftDefenderAdvancedThreatProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataConnectorMicrosoftDefenderAdvancedThreatProtection resources.
type dataConnectorMicrosoftDefenderAdvancedThreatProtectionState struct {
	// The ID of the Log Analytics Workspace that this Microsoft Defender Advanced Threat Protection Data Connector resides in. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Microsoft Defender Advanced Threat Protection Data Connector. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the tenant that this Microsoft Defender Advanced Threat Protection Data Connector connects to. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	TenantId *string `pulumi:"tenantId"`
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionState struct {
	// The ID of the Log Analytics Workspace that this Microsoft Defender Advanced Threat Protection Data Connector resides in. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Microsoft Defender Advanced Threat Protection Data Connector. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the tenant that this Microsoft Defender Advanced Threat Protection Data Connector connects to. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	TenantId pulumi.StringPtrInput
}

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorMicrosoftDefenderAdvancedThreatProtectionState)(nil)).Elem()
}

type dataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs struct {
	// The ID of the Log Analytics Workspace that this Microsoft Defender Advanced Threat Protection Data Connector resides in. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Microsoft Defender Advanced Threat Protection Data Connector. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	Name *string `pulumi:"name"`
	// The ID of the tenant that this Microsoft Defender Advanced Threat Protection Data Connector connects to. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a DataConnectorMicrosoftDefenderAdvancedThreatProtection resource.
type DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs struct {
	// The ID of the Log Analytics Workspace that this Microsoft Defender Advanced Threat Protection Data Connector resides in. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Microsoft Defender Advanced Threat Protection Data Connector. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	Name pulumi.StringPtrInput
	// The ID of the tenant that this Microsoft Defender Advanced Threat Protection Data Connector connects to. Changing this forces a new Microsoft Defender Advanced Threat Protection Data Connector to be created.
	TenantId pulumi.StringPtrInput
}

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs)(nil)).Elem()
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionInput interface {
	pulumi.Input

	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput
	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput
}

func (*DataConnectorMicrosoftDefenderAdvancedThreatProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (i *DataConnectorMicrosoftDefenderAdvancedThreatProtection) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return i.ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutputWithContext(context.Background())
}

func (i *DataConnectorMicrosoftDefenderAdvancedThreatProtection) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput)
}

// DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayInput is an input type that accepts DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray and DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput values.
// You can construct a concrete instance of `DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayInput` via:
//
//          DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray{ DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs{...} }
type DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayInput interface {
	pulumi.Input

	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput
	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutputWithContext(context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray []DataConnectorMicrosoftDefenderAdvancedThreatProtectionInput

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (i DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput {
	return i.ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutputWithContext(context.Background())
}

func (i DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput)
}

// DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapInput is an input type that accepts DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap and DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput values.
// You can construct a concrete instance of `DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapInput` via:
//
//          DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap{ "key": DataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs{...} }
type DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapInput interface {
	pulumi.Input

	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput
	ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutputWithContext(context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap map[string]DataConnectorMicrosoftDefenderAdvancedThreatProtectionInput

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (i DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput {
	return i.ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutputWithContext(context.Background())
}

func (i DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput)
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput struct{ *pulumi.OutputState }

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return o
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return o
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput struct{ *pulumi.OutputState }

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput {
	return o
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput {
	return o
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput) Index(i pulumi.IntInput) DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataConnectorMicrosoftDefenderAdvancedThreatProtection {
		return vs[0].([]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)[vs[1].(int)]
	}).(DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput)
}

type DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput struct{ *pulumi.OutputState }

func (DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)).Elem()
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput() DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput {
	return o
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput) ToDataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutputWithContext(ctx context.Context) DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput {
	return o
}

func (o DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput) MapIndex(k pulumi.StringInput) DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataConnectorMicrosoftDefenderAdvancedThreatProtection {
		return vs[0].(map[string]*DataConnectorMicrosoftDefenderAdvancedThreatProtection)[vs[1].(string)]
	}).(DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataConnectorMicrosoftDefenderAdvancedThreatProtectionInput)(nil)).Elem(), &DataConnectorMicrosoftDefenderAdvancedThreatProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayInput)(nil)).Elem(), DataConnectorMicrosoftDefenderAdvancedThreatProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapInput)(nil)).Elem(), DataConnectorMicrosoftDefenderAdvancedThreatProtectionMap{})
	pulumi.RegisterOutputType(DataConnectorMicrosoftDefenderAdvancedThreatProtectionOutput{})
	pulumi.RegisterOutputType(DataConnectorMicrosoftDefenderAdvancedThreatProtectionArrayOutput{})
	pulumi.RegisterOutputType(DataConnectorMicrosoftDefenderAdvancedThreatProtectionMapOutput{})
}
