// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a logz Monitor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/monitoring"
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
// 		_, err = monitoring.NewLogzMonitor(ctx, "exampleLogzMonitor", &monitoring.LogzMonitorArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Plan: &monitoring.LogzMonitorPlanArgs{
// 				BillingCycle:  pulumi.String("Monthly"),
// 				EffectiveDate: pulumi.String("2022-06-06T00:00:00Z"),
// 				PlanId:        pulumi.String("100gb14days"),
// 				UsageType:     pulumi.String("Committed"),
// 			},
// 			User: &monitoring.LogzMonitorUserArgs{
// 				Email:       pulumi.String("user@example.com"),
// 				FirstName:   pulumi.String("Example"),
// 				LastName:    pulumi.String("User"),
// 				PhoneNumber: pulumi.String("+12313803556"),
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
// logz Monitors can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/logzMonitor:LogzMonitor example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1
// ```
type LogzMonitor struct {
	pulumi.CustomResourceState

	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName pulumi.StringPtrOutput `pulumi:"companyName"`
	// Whether the resource monitoring is enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppId pulumi.StringPtrOutput `pulumi:"enterpriseAppId"`
	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID associated with the logz organization of this logz Monitor.
	LogzOrganizationId pulumi.StringOutput `pulumi:"logzOrganizationId"`
	// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `plan` block as defined below.
	Plan LogzMonitorPlanOutput `pulumi:"plan"`
	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The single sign on url associated with the logz organization of this logz Monitor.
	SingleSignOnUrl pulumi.StringOutput `pulumi:"singleSignOnUrl"`
	// A mapping of tags which should be assigned to the logz Monitor.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `user` block as defined below.
	User LogzMonitorUserOutput `pulumi:"user"`
}

// NewLogzMonitor registers a new resource with the given unique name, arguments, and options.
func NewLogzMonitor(ctx *pulumi.Context,
	name string, args *LogzMonitorArgs, opts ...pulumi.ResourceOption) (*LogzMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource LogzMonitor
	err := ctx.RegisterResource("azure:monitoring/logzMonitor:LogzMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogzMonitor gets an existing LogzMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogzMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogzMonitorState, opts ...pulumi.ResourceOption) (*LogzMonitor, error) {
	var resource LogzMonitor
	err := ctx.ReadResource("azure:monitoring/logzMonitor:LogzMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogzMonitor resources.
type logzMonitorState struct {
	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName *string `pulumi:"companyName"`
	// Whether the resource monitoring is enabled?
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location *string `pulumi:"location"`
	// The ID associated with the logz organization of this logz Monitor.
	LogzOrganizationId *string `pulumi:"logzOrganizationId"`
	// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
	Name *string `pulumi:"name"`
	// A `plan` block as defined below.
	Plan *LogzMonitorPlan `pulumi:"plan"`
	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The single sign on url associated with the logz organization of this logz Monitor.
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
	// A mapping of tags which should be assigned to the logz Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `user` block as defined below.
	User *LogzMonitorUser `pulumi:"user"`
}

type LogzMonitorState struct {
	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName pulumi.StringPtrInput
	// Whether the resource monitoring is enabled?
	Enabled pulumi.BoolPtrInput
	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppId pulumi.StringPtrInput
	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location pulumi.StringPtrInput
	// The ID associated with the logz organization of this logz Monitor.
	LogzOrganizationId pulumi.StringPtrInput
	// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
	Name pulumi.StringPtrInput
	// A `plan` block as defined below.
	Plan LogzMonitorPlanPtrInput
	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The single sign on url associated with the logz organization of this logz Monitor.
	SingleSignOnUrl pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the logz Monitor.
	Tags pulumi.StringMapInput
	// A `user` block as defined below.
	User LogzMonitorUserPtrInput
}

func (LogzMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*logzMonitorState)(nil)).Elem()
}

type logzMonitorArgs struct {
	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName *string `pulumi:"companyName"`
	// Whether the resource monitoring is enabled?
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
	Name *string `pulumi:"name"`
	// A `plan` block as defined below.
	Plan LogzMonitorPlan `pulumi:"plan"`
	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the logz Monitor.
	Tags map[string]string `pulumi:"tags"`
	// A `user` block as defined below.
	User LogzMonitorUser `pulumi:"user"`
}

// The set of arguments for constructing a LogzMonitor resource.
type LogzMonitorArgs struct {
	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName pulumi.StringPtrInput
	// Whether the resource monitoring is enabled?
	Enabled pulumi.BoolPtrInput
	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppId pulumi.StringPtrInput
	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
	Name pulumi.StringPtrInput
	// A `plan` block as defined below.
	Plan LogzMonitorPlanInput
	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the logz Monitor.
	Tags pulumi.StringMapInput
	// A `user` block as defined below.
	User LogzMonitorUserInput
}

func (LogzMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logzMonitorArgs)(nil)).Elem()
}

type LogzMonitorInput interface {
	pulumi.Input

	ToLogzMonitorOutput() LogzMonitorOutput
	ToLogzMonitorOutputWithContext(ctx context.Context) LogzMonitorOutput
}

func (*LogzMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**LogzMonitor)(nil)).Elem()
}

func (i *LogzMonitor) ToLogzMonitorOutput() LogzMonitorOutput {
	return i.ToLogzMonitorOutputWithContext(context.Background())
}

func (i *LogzMonitor) ToLogzMonitorOutputWithContext(ctx context.Context) LogzMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzMonitorOutput)
}

// LogzMonitorArrayInput is an input type that accepts LogzMonitorArray and LogzMonitorArrayOutput values.
// You can construct a concrete instance of `LogzMonitorArrayInput` via:
//
//          LogzMonitorArray{ LogzMonitorArgs{...} }
type LogzMonitorArrayInput interface {
	pulumi.Input

	ToLogzMonitorArrayOutput() LogzMonitorArrayOutput
	ToLogzMonitorArrayOutputWithContext(context.Context) LogzMonitorArrayOutput
}

type LogzMonitorArray []LogzMonitorInput

func (LogzMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogzMonitor)(nil)).Elem()
}

func (i LogzMonitorArray) ToLogzMonitorArrayOutput() LogzMonitorArrayOutput {
	return i.ToLogzMonitorArrayOutputWithContext(context.Background())
}

func (i LogzMonitorArray) ToLogzMonitorArrayOutputWithContext(ctx context.Context) LogzMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzMonitorArrayOutput)
}

// LogzMonitorMapInput is an input type that accepts LogzMonitorMap and LogzMonitorMapOutput values.
// You can construct a concrete instance of `LogzMonitorMapInput` via:
//
//          LogzMonitorMap{ "key": LogzMonitorArgs{...} }
type LogzMonitorMapInput interface {
	pulumi.Input

	ToLogzMonitorMapOutput() LogzMonitorMapOutput
	ToLogzMonitorMapOutputWithContext(context.Context) LogzMonitorMapOutput
}

type LogzMonitorMap map[string]LogzMonitorInput

func (LogzMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogzMonitor)(nil)).Elem()
}

func (i LogzMonitorMap) ToLogzMonitorMapOutput() LogzMonitorMapOutput {
	return i.ToLogzMonitorMapOutputWithContext(context.Background())
}

func (i LogzMonitorMap) ToLogzMonitorMapOutputWithContext(ctx context.Context) LogzMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzMonitorMapOutput)
}

type LogzMonitorOutput struct{ *pulumi.OutputState }

func (LogzMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogzMonitor)(nil)).Elem()
}

func (o LogzMonitorOutput) ToLogzMonitorOutput() LogzMonitorOutput {
	return o
}

func (o LogzMonitorOutput) ToLogzMonitorOutputWithContext(ctx context.Context) LogzMonitorOutput {
	return o
}

type LogzMonitorArrayOutput struct{ *pulumi.OutputState }

func (LogzMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogzMonitor)(nil)).Elem()
}

func (o LogzMonitorArrayOutput) ToLogzMonitorArrayOutput() LogzMonitorArrayOutput {
	return o
}

func (o LogzMonitorArrayOutput) ToLogzMonitorArrayOutputWithContext(ctx context.Context) LogzMonitorArrayOutput {
	return o
}

func (o LogzMonitorArrayOutput) Index(i pulumi.IntInput) LogzMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogzMonitor {
		return vs[0].([]*LogzMonitor)[vs[1].(int)]
	}).(LogzMonitorOutput)
}

type LogzMonitorMapOutput struct{ *pulumi.OutputState }

func (LogzMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogzMonitor)(nil)).Elem()
}

func (o LogzMonitorMapOutput) ToLogzMonitorMapOutput() LogzMonitorMapOutput {
	return o
}

func (o LogzMonitorMapOutput) ToLogzMonitorMapOutputWithContext(ctx context.Context) LogzMonitorMapOutput {
	return o
}

func (o LogzMonitorMapOutput) MapIndex(k pulumi.StringInput) LogzMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogzMonitor {
		return vs[0].(map[string]*LogzMonitor)[vs[1].(string)]
	}).(LogzMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogzMonitorInput)(nil)).Elem(), &LogzMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogzMonitorArrayInput)(nil)).Elem(), LogzMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogzMonitorMapInput)(nil)).Elem(), LogzMonitorMap{})
	pulumi.RegisterOutputType(LogzMonitorOutput{})
	pulumi.RegisterOutputType(LogzMonitorArrayOutput{})
	pulumi.RegisterOutputType(LogzMonitorMapOutput{})
}
