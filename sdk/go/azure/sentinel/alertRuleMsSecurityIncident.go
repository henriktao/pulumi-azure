// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Sentinel MS Security Incident Alert Rule.
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
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("pergb2018"),
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
// 		_, err = sentinel.NewAlertRuleMsSecurityIncident(ctx, "exampleAlertRuleMsSecurityIncident", &sentinel.AlertRuleMsSecurityIncidentArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsSolution.WorkspaceResourceId,
// 			ProductFilter:           pulumi.String("Microsoft Cloud App Security"),
// 			DisplayName:             pulumi.String("example rule"),
// 			SeverityFilters: pulumi.StringArray{
// 				pulumi.String("High"),
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
// Sentinel MS Security Incident Alert Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
// ```
type AlertRuleMsSecurityIncident struct {
	pulumi.CustomResourceState

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrOutput `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilters pulumi.StringArrayOutput `pulumi:"displayNameExcludeFilters"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayOutput `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT`, `Microsoft Cloud App Security`, `Microsoft Defender Advanced Threat Protection` and `Office 365 Advanced Threat Protection`.
	ProductFilter pulumi.StringOutput `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayOutput `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayOutput `pulumi:"textWhitelists"`
}

// NewAlertRuleMsSecurityIncident registers a new resource with the given unique name, arguments, and options.
func NewAlertRuleMsSecurityIncident(ctx *pulumi.Context,
	name string, args *AlertRuleMsSecurityIncidentArgs, opts ...pulumi.ResourceOption) (*AlertRuleMsSecurityIncident, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	if args.ProductFilter == nil {
		return nil, errors.New("invalid value for required argument 'ProductFilter'")
	}
	if args.SeverityFilters == nil {
		return nil, errors.New("invalid value for required argument 'SeverityFilters'")
	}
	var resource AlertRuleMsSecurityIncident
	err := ctx.RegisterResource("azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRuleMsSecurityIncident gets an existing AlertRuleMsSecurityIncident resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRuleMsSecurityIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleMsSecurityIncidentState, opts ...pulumi.ResourceOption) (*AlertRuleMsSecurityIncident, error) {
	var resource AlertRuleMsSecurityIncident
	err := ctx.ReadResource("azure:sentinel/alertRuleMsSecurityIncident:AlertRuleMsSecurityIncident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRuleMsSecurityIncident resources.
type alertRuleMsSecurityIncidentState struct {
	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGuid *string `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName *string `pulumi:"displayName"`
	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilters []string `pulumi:"displayNameExcludeFilters"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters []string `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT`, `Microsoft Cloud App Security`, `Microsoft Defender Advanced Threat Protection` and `Office 365 Advanced Threat Protection`.
	ProductFilter *string `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters []string `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists []string `pulumi:"textWhitelists"`
}

type AlertRuleMsSecurityIncidentState struct {
	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrInput
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringPtrInput
	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilters pulumi.StringArrayInput
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayInput
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT`, `Microsoft Cloud App Security`, `Microsoft Defender Advanced Threat Protection` and `Office 365 Advanced Threat Protection`.
	ProductFilter pulumi.StringPtrInput
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayInput
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayInput
}

func (AlertRuleMsSecurityIncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleMsSecurityIncidentState)(nil)).Elem()
}

type alertRuleMsSecurityIncidentArgs struct {
	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGuid *string `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName string `pulumi:"displayName"`
	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilters []string `pulumi:"displayNameExcludeFilters"`
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters []string `pulumi:"displayNameFilters"`
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT`, `Microsoft Cloud App Security`, `Microsoft Defender Advanced Threat Protection` and `Office 365 Advanced Threat Protection`.
	ProductFilter string `pulumi:"productFilter"`
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters []string `pulumi:"severityFilters"`
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists []string `pulumi:"textWhitelists"`
}

// The set of arguments for constructing a AlertRuleMsSecurityIncident resource.
type AlertRuleMsSecurityIncidentArgs struct {
	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrInput
	// The description of this Sentinel MS Security Incident Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName pulumi.StringInput
	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilters pulumi.StringArrayInput
	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilters pulumi.StringArrayInput
	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Sentinel MS Security Incident Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The Microsoft Security Service from where the alert will be generated. Possible values are `Azure Active Directory Identity Protection`, `Azure Advanced Threat Protection`, `Azure Security Center`, `Azure Security Center for IoT`, `Microsoft Cloud App Security`, `Microsoft Defender Advanced Threat Protection` and `Office 365 Advanced Threat Protection`.
	ProductFilter pulumi.StringInput
	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are `High`, `Medium`, `Low` and `Informational`.
	SeverityFilters pulumi.StringArrayInput
	// Deprecated: this property has been renamed to display_name_filter to better match the SDK & API
	TextWhitelists pulumi.StringArrayInput
}

func (AlertRuleMsSecurityIncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleMsSecurityIncidentArgs)(nil)).Elem()
}

type AlertRuleMsSecurityIncidentInput interface {
	pulumi.Input

	ToAlertRuleMsSecurityIncidentOutput() AlertRuleMsSecurityIncidentOutput
	ToAlertRuleMsSecurityIncidentOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentOutput
}

func (*AlertRuleMsSecurityIncident) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (i *AlertRuleMsSecurityIncident) ToAlertRuleMsSecurityIncidentOutput() AlertRuleMsSecurityIncidentOutput {
	return i.ToAlertRuleMsSecurityIncidentOutputWithContext(context.Background())
}

func (i *AlertRuleMsSecurityIncident) ToAlertRuleMsSecurityIncidentOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleMsSecurityIncidentOutput)
}

// AlertRuleMsSecurityIncidentArrayInput is an input type that accepts AlertRuleMsSecurityIncidentArray and AlertRuleMsSecurityIncidentArrayOutput values.
// You can construct a concrete instance of `AlertRuleMsSecurityIncidentArrayInput` via:
//
//          AlertRuleMsSecurityIncidentArray{ AlertRuleMsSecurityIncidentArgs{...} }
type AlertRuleMsSecurityIncidentArrayInput interface {
	pulumi.Input

	ToAlertRuleMsSecurityIncidentArrayOutput() AlertRuleMsSecurityIncidentArrayOutput
	ToAlertRuleMsSecurityIncidentArrayOutputWithContext(context.Context) AlertRuleMsSecurityIncidentArrayOutput
}

type AlertRuleMsSecurityIncidentArray []AlertRuleMsSecurityIncidentInput

func (AlertRuleMsSecurityIncidentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (i AlertRuleMsSecurityIncidentArray) ToAlertRuleMsSecurityIncidentArrayOutput() AlertRuleMsSecurityIncidentArrayOutput {
	return i.ToAlertRuleMsSecurityIncidentArrayOutputWithContext(context.Background())
}

func (i AlertRuleMsSecurityIncidentArray) ToAlertRuleMsSecurityIncidentArrayOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleMsSecurityIncidentArrayOutput)
}

// AlertRuleMsSecurityIncidentMapInput is an input type that accepts AlertRuleMsSecurityIncidentMap and AlertRuleMsSecurityIncidentMapOutput values.
// You can construct a concrete instance of `AlertRuleMsSecurityIncidentMapInput` via:
//
//          AlertRuleMsSecurityIncidentMap{ "key": AlertRuleMsSecurityIncidentArgs{...} }
type AlertRuleMsSecurityIncidentMapInput interface {
	pulumi.Input

	ToAlertRuleMsSecurityIncidentMapOutput() AlertRuleMsSecurityIncidentMapOutput
	ToAlertRuleMsSecurityIncidentMapOutputWithContext(context.Context) AlertRuleMsSecurityIncidentMapOutput
}

type AlertRuleMsSecurityIncidentMap map[string]AlertRuleMsSecurityIncidentInput

func (AlertRuleMsSecurityIncidentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (i AlertRuleMsSecurityIncidentMap) ToAlertRuleMsSecurityIncidentMapOutput() AlertRuleMsSecurityIncidentMapOutput {
	return i.ToAlertRuleMsSecurityIncidentMapOutputWithContext(context.Background())
}

func (i AlertRuleMsSecurityIncidentMap) ToAlertRuleMsSecurityIncidentMapOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleMsSecurityIncidentMapOutput)
}

type AlertRuleMsSecurityIncidentOutput struct{ *pulumi.OutputState }

func (AlertRuleMsSecurityIncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (o AlertRuleMsSecurityIncidentOutput) ToAlertRuleMsSecurityIncidentOutput() AlertRuleMsSecurityIncidentOutput {
	return o
}

func (o AlertRuleMsSecurityIncidentOutput) ToAlertRuleMsSecurityIncidentOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentOutput {
	return o
}

type AlertRuleMsSecurityIncidentArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleMsSecurityIncidentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (o AlertRuleMsSecurityIncidentArrayOutput) ToAlertRuleMsSecurityIncidentArrayOutput() AlertRuleMsSecurityIncidentArrayOutput {
	return o
}

func (o AlertRuleMsSecurityIncidentArrayOutput) ToAlertRuleMsSecurityIncidentArrayOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentArrayOutput {
	return o
}

func (o AlertRuleMsSecurityIncidentArrayOutput) Index(i pulumi.IntInput) AlertRuleMsSecurityIncidentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlertRuleMsSecurityIncident {
		return vs[0].([]*AlertRuleMsSecurityIncident)[vs[1].(int)]
	}).(AlertRuleMsSecurityIncidentOutput)
}

type AlertRuleMsSecurityIncidentMapOutput struct{ *pulumi.OutputState }

func (AlertRuleMsSecurityIncidentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertRuleMsSecurityIncident)(nil)).Elem()
}

func (o AlertRuleMsSecurityIncidentMapOutput) ToAlertRuleMsSecurityIncidentMapOutput() AlertRuleMsSecurityIncidentMapOutput {
	return o
}

func (o AlertRuleMsSecurityIncidentMapOutput) ToAlertRuleMsSecurityIncidentMapOutputWithContext(ctx context.Context) AlertRuleMsSecurityIncidentMapOutput {
	return o
}

func (o AlertRuleMsSecurityIncidentMapOutput) MapIndex(k pulumi.StringInput) AlertRuleMsSecurityIncidentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlertRuleMsSecurityIncident {
		return vs[0].(map[string]*AlertRuleMsSecurityIncident)[vs[1].(string)]
	}).(AlertRuleMsSecurityIncidentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertRuleMsSecurityIncidentInput)(nil)).Elem(), &AlertRuleMsSecurityIncident{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertRuleMsSecurityIncidentArrayInput)(nil)).Elem(), AlertRuleMsSecurityIncidentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertRuleMsSecurityIncidentMapInput)(nil)).Elem(), AlertRuleMsSecurityIncidentMap{})
	pulumi.RegisterOutputType(AlertRuleMsSecurityIncidentOutput{})
	pulumi.RegisterOutputType(AlertRuleMsSecurityIncidentArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleMsSecurityIncidentMapOutput{})
}
