// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Monitor Action Rule which type is action group.
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
// 		exampleActionGroup, err := monitoring.NewActionGroup(ctx, "exampleActionGroup", &monitoring.ActionGroupArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ShortName:         pulumi.String("exampleactiongroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewActionRuleActionGroup(ctx, "exampleActionRuleActionGroup", &monitoring.ActionRuleActionGroupArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ActionGroupId:     exampleActionGroup.ID(),
// 			Scope: &monitoring.ActionRuleActionGroupScopeArgs{
// 				Type: pulumi.String("ResourceGroup"),
// 				ResourceIds: pulumi.StringArray{
// 					exampleResourceGroup.ID(),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
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
// Monitor Action Rule can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:monitoring/actionRuleActionGroup:ActionRuleActionGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
// ```
type ActionRuleActionGroup struct {
	pulumi.CustomResourceState

	// Specifies the resource id of monitor action group.
	ActionGroupId pulumi.StringOutput `pulumi:"actionGroupId"`
	// A `condition` block as defined below.
	Condition ActionRuleActionGroupConditionPtrOutput `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope ActionRuleActionGroupScopePtrOutput `pulumi:"scope"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActionRuleActionGroup registers a new resource with the given unique name, arguments, and options.
func NewActionRuleActionGroup(ctx *pulumi.Context,
	name string, args *ActionRuleActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionRuleActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ActionGroupId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ActionRuleActionGroup
	err := ctx.RegisterResource("azure:monitoring/actionRuleActionGroup:ActionRuleActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionRuleActionGroup gets an existing ActionRuleActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionRuleActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionRuleActionGroupState, opts ...pulumi.ResourceOption) (*ActionRuleActionGroup, error) {
	var resource ActionRuleActionGroup
	err := ctx.ReadResource("azure:monitoring/actionRuleActionGroup:ActionRuleActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionRuleActionGroup resources.
type actionRuleActionGroupState struct {
	// Specifies the resource id of monitor action group.
	ActionGroupId *string `pulumi:"actionGroupId"`
	// A `condition` block as defined below.
	Condition *ActionRuleActionGroupCondition `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description *string `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope *ActionRuleActionGroupScope `pulumi:"scope"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ActionRuleActionGroupState struct {
	// Specifies the resource id of monitor action group.
	ActionGroupId pulumi.StringPtrInput
	// A `condition` block as defined below.
	Condition ActionRuleActionGroupConditionPtrInput
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrInput
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `scope` block as defined below.
	Scope ActionRuleActionGroupScopePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActionRuleActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleActionGroupState)(nil)).Elem()
}

type actionRuleActionGroupArgs struct {
	// Specifies the resource id of monitor action group.
	ActionGroupId string `pulumi:"actionGroupId"`
	// A `condition` block as defined below.
	Condition *ActionRuleActionGroupCondition `pulumi:"condition"`
	// Specifies a description for the Action Rule.
	Description *string `pulumi:"description"`
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `scope` block as defined below.
	Scope *ActionRuleActionGroupScope `pulumi:"scope"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActionRuleActionGroup resource.
type ActionRuleActionGroupArgs struct {
	// Specifies the resource id of monitor action group.
	ActionGroupId pulumi.StringInput
	// A `condition` block as defined below.
	Condition ActionRuleActionGroupConditionPtrInput
	// Specifies a description for the Action Rule.
	Description pulumi.StringPtrInput
	// Is the Action Rule enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of the Monitor Action Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `scope` block as defined below.
	Scope ActionRuleActionGroupScopePtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActionRuleActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionRuleActionGroupArgs)(nil)).Elem()
}

type ActionRuleActionGroupInput interface {
	pulumi.Input

	ToActionRuleActionGroupOutput() ActionRuleActionGroupOutput
	ToActionRuleActionGroupOutputWithContext(ctx context.Context) ActionRuleActionGroupOutput
}

func (*ActionRuleActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleActionGroup)(nil)).Elem()
}

func (i *ActionRuleActionGroup) ToActionRuleActionGroupOutput() ActionRuleActionGroupOutput {
	return i.ToActionRuleActionGroupOutputWithContext(context.Background())
}

func (i *ActionRuleActionGroup) ToActionRuleActionGroupOutputWithContext(ctx context.Context) ActionRuleActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleActionGroupOutput)
}

// ActionRuleActionGroupArrayInput is an input type that accepts ActionRuleActionGroupArray and ActionRuleActionGroupArrayOutput values.
// You can construct a concrete instance of `ActionRuleActionGroupArrayInput` via:
//
//          ActionRuleActionGroupArray{ ActionRuleActionGroupArgs{...} }
type ActionRuleActionGroupArrayInput interface {
	pulumi.Input

	ToActionRuleActionGroupArrayOutput() ActionRuleActionGroupArrayOutput
	ToActionRuleActionGroupArrayOutputWithContext(context.Context) ActionRuleActionGroupArrayOutput
}

type ActionRuleActionGroupArray []ActionRuleActionGroupInput

func (ActionRuleActionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionRuleActionGroup)(nil)).Elem()
}

func (i ActionRuleActionGroupArray) ToActionRuleActionGroupArrayOutput() ActionRuleActionGroupArrayOutput {
	return i.ToActionRuleActionGroupArrayOutputWithContext(context.Background())
}

func (i ActionRuleActionGroupArray) ToActionRuleActionGroupArrayOutputWithContext(ctx context.Context) ActionRuleActionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleActionGroupArrayOutput)
}

// ActionRuleActionGroupMapInput is an input type that accepts ActionRuleActionGroupMap and ActionRuleActionGroupMapOutput values.
// You can construct a concrete instance of `ActionRuleActionGroupMapInput` via:
//
//          ActionRuleActionGroupMap{ "key": ActionRuleActionGroupArgs{...} }
type ActionRuleActionGroupMapInput interface {
	pulumi.Input

	ToActionRuleActionGroupMapOutput() ActionRuleActionGroupMapOutput
	ToActionRuleActionGroupMapOutputWithContext(context.Context) ActionRuleActionGroupMapOutput
}

type ActionRuleActionGroupMap map[string]ActionRuleActionGroupInput

func (ActionRuleActionGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionRuleActionGroup)(nil)).Elem()
}

func (i ActionRuleActionGroupMap) ToActionRuleActionGroupMapOutput() ActionRuleActionGroupMapOutput {
	return i.ToActionRuleActionGroupMapOutputWithContext(context.Background())
}

func (i ActionRuleActionGroupMap) ToActionRuleActionGroupMapOutputWithContext(ctx context.Context) ActionRuleActionGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRuleActionGroupMapOutput)
}

type ActionRuleActionGroupOutput struct{ *pulumi.OutputState }

func (ActionRuleActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleActionGroup)(nil)).Elem()
}

func (o ActionRuleActionGroupOutput) ToActionRuleActionGroupOutput() ActionRuleActionGroupOutput {
	return o
}

func (o ActionRuleActionGroupOutput) ToActionRuleActionGroupOutputWithContext(ctx context.Context) ActionRuleActionGroupOutput {
	return o
}

type ActionRuleActionGroupArrayOutput struct{ *pulumi.OutputState }

func (ActionRuleActionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionRuleActionGroup)(nil)).Elem()
}

func (o ActionRuleActionGroupArrayOutput) ToActionRuleActionGroupArrayOutput() ActionRuleActionGroupArrayOutput {
	return o
}

func (o ActionRuleActionGroupArrayOutput) ToActionRuleActionGroupArrayOutputWithContext(ctx context.Context) ActionRuleActionGroupArrayOutput {
	return o
}

func (o ActionRuleActionGroupArrayOutput) Index(i pulumi.IntInput) ActionRuleActionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionRuleActionGroup {
		return vs[0].([]*ActionRuleActionGroup)[vs[1].(int)]
	}).(ActionRuleActionGroupOutput)
}

type ActionRuleActionGroupMapOutput struct{ *pulumi.OutputState }

func (ActionRuleActionGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionRuleActionGroup)(nil)).Elem()
}

func (o ActionRuleActionGroupMapOutput) ToActionRuleActionGroupMapOutput() ActionRuleActionGroupMapOutput {
	return o
}

func (o ActionRuleActionGroupMapOutput) ToActionRuleActionGroupMapOutputWithContext(ctx context.Context) ActionRuleActionGroupMapOutput {
	return o
}

func (o ActionRuleActionGroupMapOutput) MapIndex(k pulumi.StringInput) ActionRuleActionGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionRuleActionGroup {
		return vs[0].(map[string]*ActionRuleActionGroup)[vs[1].(string)]
	}).(ActionRuleActionGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionRuleActionGroupInput)(nil)).Elem(), &ActionRuleActionGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionRuleActionGroupArrayInput)(nil)).Elem(), ActionRuleActionGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionRuleActionGroupMapInput)(nil)).Elem(), ActionRuleActionGroupMap{})
	pulumi.RegisterOutputType(ActionRuleActionGroupOutput{})
	pulumi.RegisterOutputType(ActionRuleActionGroupArrayOutput{})
	pulumi.RegisterOutputType(ActionRuleActionGroupMapOutput{})
}
