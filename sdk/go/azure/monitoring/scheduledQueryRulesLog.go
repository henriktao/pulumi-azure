// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a LogToMetricAction Scheduled Query Rules resource within Azure Monitor.
//
// ## Import
//
// Scheduled Query Rule Log can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
// ```
type ScheduledQueryRulesLog struct {
	pulumi.CustomResourceState

	AuthorizedResourceIds pulumi.StringArrayOutput `pulumi:"authorizedResourceIds"`
	// A `criteria` block as defined below.
	Criteria ScheduledQueryRulesLogCriteriaOutput `pulumi:"criteria"`
	// The resource uri over which log search query is to be run.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled  pulumi.BoolPtrOutput `pulumi:"enabled"`
	Location pulumi.StringOutput  `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringOutput    `pulumi:"resourceGroupName"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
}

// NewScheduledQueryRulesLog registers a new resource with the given unique name, arguments, and options.
func NewScheduledQueryRulesLog(ctx *pulumi.Context,
	name string, args *ScheduledQueryRulesLogArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRulesLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.DataSourceId == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ScheduledQueryRulesLog
	err := ctx.RegisterResource("azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledQueryRulesLog gets an existing ScheduledQueryRulesLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledQueryRulesLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRulesLogState, opts ...pulumi.ResourceOption) (*ScheduledQueryRulesLog, error) {
	var resource ScheduledQueryRulesLog
	err := ctx.ReadResource("azure:monitoring/scheduledQueryRulesLog:ScheduledQueryRulesLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledQueryRulesLog resources.
type scheduledQueryRulesLogState struct {
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// A `criteria` block as defined below.
	Criteria *ScheduledQueryRulesLogCriteria `pulumi:"criteria"`
	// The resource uri over which log search query is to be run.
	DataSourceId *string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled  *bool   `pulumi:"enabled"`
	Location *string `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName *string           `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

type ScheduledQueryRulesLogState struct {
	AuthorizedResourceIds pulumi.StringArrayInput
	// A `criteria` block as defined below.
	Criteria ScheduledQueryRulesLogCriteriaPtrInput
	// The resource uri over which log search query is to be run.
	DataSourceId pulumi.StringPtrInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled  pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (ScheduledQueryRulesLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRulesLogState)(nil)).Elem()
}

type scheduledQueryRulesLogArgs struct {
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// A `criteria` block as defined below.
	Criteria ScheduledQueryRulesLogCriteria `pulumi:"criteria"`
	// The resource uri over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled  *bool   `pulumi:"enabled"`
	Location *string `pulumi:"location"`
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ScheduledQueryRulesLog resource.
type ScheduledQueryRulesLogArgs struct {
	AuthorizedResourceIds pulumi.StringArrayInput
	// A `criteria` block as defined below.
	Criteria ScheduledQueryRulesLogCriteriaInput
	// The resource uri over which log search query is to be run.
	DataSourceId pulumi.StringInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// Whether this scheduled query rule is enabled.  Default is `true`.
	Enabled  pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the scheduled query rule instance.
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ScheduledQueryRulesLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRulesLogArgs)(nil)).Elem()
}

type ScheduledQueryRulesLogInput interface {
	pulumi.Input

	ToScheduledQueryRulesLogOutput() ScheduledQueryRulesLogOutput
	ToScheduledQueryRulesLogOutputWithContext(ctx context.Context) ScheduledQueryRulesLogOutput
}

func (*ScheduledQueryRulesLog) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRulesLog)(nil))
}

func (i *ScheduledQueryRulesLog) ToScheduledQueryRulesLogOutput() ScheduledQueryRulesLogOutput {
	return i.ToScheduledQueryRulesLogOutputWithContext(context.Background())
}

func (i *ScheduledQueryRulesLog) ToScheduledQueryRulesLogOutputWithContext(ctx context.Context) ScheduledQueryRulesLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesLogOutput)
}

func (i *ScheduledQueryRulesLog) ToScheduledQueryRulesLogPtrOutput() ScheduledQueryRulesLogPtrOutput {
	return i.ToScheduledQueryRulesLogPtrOutputWithContext(context.Background())
}

func (i *ScheduledQueryRulesLog) ToScheduledQueryRulesLogPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesLogPtrOutput)
}

type ScheduledQueryRulesLogPtrInput interface {
	pulumi.Input

	ToScheduledQueryRulesLogPtrOutput() ScheduledQueryRulesLogPtrOutput
	ToScheduledQueryRulesLogPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesLogPtrOutput
}

type scheduledQueryRulesLogPtrType ScheduledQueryRulesLogArgs

func (*scheduledQueryRulesLogPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRulesLog)(nil))
}

func (i *scheduledQueryRulesLogPtrType) ToScheduledQueryRulesLogPtrOutput() ScheduledQueryRulesLogPtrOutput {
	return i.ToScheduledQueryRulesLogPtrOutputWithContext(context.Background())
}

func (i *scheduledQueryRulesLogPtrType) ToScheduledQueryRulesLogPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesLogPtrOutput)
}

// ScheduledQueryRulesLogArrayInput is an input type that accepts ScheduledQueryRulesLogArray and ScheduledQueryRulesLogArrayOutput values.
// You can construct a concrete instance of `ScheduledQueryRulesLogArrayInput` via:
//
//          ScheduledQueryRulesLogArray{ ScheduledQueryRulesLogArgs{...} }
type ScheduledQueryRulesLogArrayInput interface {
	pulumi.Input

	ToScheduledQueryRulesLogArrayOutput() ScheduledQueryRulesLogArrayOutput
	ToScheduledQueryRulesLogArrayOutputWithContext(context.Context) ScheduledQueryRulesLogArrayOutput
}

type ScheduledQueryRulesLogArray []ScheduledQueryRulesLogInput

func (ScheduledQueryRulesLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledQueryRulesLog)(nil)).Elem()
}

func (i ScheduledQueryRulesLogArray) ToScheduledQueryRulesLogArrayOutput() ScheduledQueryRulesLogArrayOutput {
	return i.ToScheduledQueryRulesLogArrayOutputWithContext(context.Background())
}

func (i ScheduledQueryRulesLogArray) ToScheduledQueryRulesLogArrayOutputWithContext(ctx context.Context) ScheduledQueryRulesLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesLogArrayOutput)
}

// ScheduledQueryRulesLogMapInput is an input type that accepts ScheduledQueryRulesLogMap and ScheduledQueryRulesLogMapOutput values.
// You can construct a concrete instance of `ScheduledQueryRulesLogMapInput` via:
//
//          ScheduledQueryRulesLogMap{ "key": ScheduledQueryRulesLogArgs{...} }
type ScheduledQueryRulesLogMapInput interface {
	pulumi.Input

	ToScheduledQueryRulesLogMapOutput() ScheduledQueryRulesLogMapOutput
	ToScheduledQueryRulesLogMapOutputWithContext(context.Context) ScheduledQueryRulesLogMapOutput
}

type ScheduledQueryRulesLogMap map[string]ScheduledQueryRulesLogInput

func (ScheduledQueryRulesLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledQueryRulesLog)(nil)).Elem()
}

func (i ScheduledQueryRulesLogMap) ToScheduledQueryRulesLogMapOutput() ScheduledQueryRulesLogMapOutput {
	return i.ToScheduledQueryRulesLogMapOutputWithContext(context.Background())
}

func (i ScheduledQueryRulesLogMap) ToScheduledQueryRulesLogMapOutputWithContext(ctx context.Context) ScheduledQueryRulesLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRulesLogMapOutput)
}

type ScheduledQueryRulesLogOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRulesLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRulesLog)(nil))
}

func (o ScheduledQueryRulesLogOutput) ToScheduledQueryRulesLogOutput() ScheduledQueryRulesLogOutput {
	return o
}

func (o ScheduledQueryRulesLogOutput) ToScheduledQueryRulesLogOutputWithContext(ctx context.Context) ScheduledQueryRulesLogOutput {
	return o
}

func (o ScheduledQueryRulesLogOutput) ToScheduledQueryRulesLogPtrOutput() ScheduledQueryRulesLogPtrOutput {
	return o.ToScheduledQueryRulesLogPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryRulesLogOutput) ToScheduledQueryRulesLogPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesLogPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryRulesLog) *ScheduledQueryRulesLog {
		return &v
	}).(ScheduledQueryRulesLogPtrOutput)
}

type ScheduledQueryRulesLogPtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRulesLogPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRulesLog)(nil))
}

func (o ScheduledQueryRulesLogPtrOutput) ToScheduledQueryRulesLogPtrOutput() ScheduledQueryRulesLogPtrOutput {
	return o
}

func (o ScheduledQueryRulesLogPtrOutput) ToScheduledQueryRulesLogPtrOutputWithContext(ctx context.Context) ScheduledQueryRulesLogPtrOutput {
	return o
}

func (o ScheduledQueryRulesLogPtrOutput) Elem() ScheduledQueryRulesLogOutput {
	return o.ApplyT(func(v *ScheduledQueryRulesLog) ScheduledQueryRulesLog {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryRulesLog
		return ret
	}).(ScheduledQueryRulesLogOutput)
}

type ScheduledQueryRulesLogArrayOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRulesLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduledQueryRulesLog)(nil))
}

func (o ScheduledQueryRulesLogArrayOutput) ToScheduledQueryRulesLogArrayOutput() ScheduledQueryRulesLogArrayOutput {
	return o
}

func (o ScheduledQueryRulesLogArrayOutput) ToScheduledQueryRulesLogArrayOutputWithContext(ctx context.Context) ScheduledQueryRulesLogArrayOutput {
	return o
}

func (o ScheduledQueryRulesLogArrayOutput) Index(i pulumi.IntInput) ScheduledQueryRulesLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduledQueryRulesLog {
		return vs[0].([]ScheduledQueryRulesLog)[vs[1].(int)]
	}).(ScheduledQueryRulesLogOutput)
}

type ScheduledQueryRulesLogMapOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRulesLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ScheduledQueryRulesLog)(nil))
}

func (o ScheduledQueryRulesLogMapOutput) ToScheduledQueryRulesLogMapOutput() ScheduledQueryRulesLogMapOutput {
	return o
}

func (o ScheduledQueryRulesLogMapOutput) ToScheduledQueryRulesLogMapOutputWithContext(ctx context.Context) ScheduledQueryRulesLogMapOutput {
	return o
}

func (o ScheduledQueryRulesLogMapOutput) MapIndex(k pulumi.StringInput) ScheduledQueryRulesLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ScheduledQueryRulesLog {
		return vs[0].(map[string]ScheduledQueryRulesLog)[vs[1].(string)]
	}).(ScheduledQueryRulesLogOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledQueryRulesLogOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRulesLogPtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRulesLogArrayOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRulesLogMapOutput{})
}
