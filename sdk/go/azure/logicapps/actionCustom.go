// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Custom Action within a Logic App Workflow
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/logicapps"
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
// 		exampleWorkflow, err := logicapps.NewWorkflow(ctx, "exampleWorkflow", &logicapps.WorkflowArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logicapps.NewActionCustom(ctx, "exampleActionCustom", &logicapps.ActionCustomArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Body:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"description\": \"A variable to configure the auto expiration age in days. Configured in negative number. Default is -30 (30 days old).\",\n", "    \"inputs\": {\n", "        \"variables\": [\n", "            {\n", "                \"name\": \"ExpirationAgeInDays\",\n", "                \"type\": \"Integer\",\n", "                \"value\": -30\n", "            }\n", "        ]\n", "    },\n", "    \"runAfter\": {},\n", "    \"type\": \"InitializeVariable\"\n", "}\n")),
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
// Logic App Custom Actions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/actionCustom:ActionCustom custom1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/custom1
// ```
type ActionCustom struct {
	pulumi.CustomResourceState

	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringOutput `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewActionCustom registers a new resource with the given unique name, arguments, and options.
func NewActionCustom(ctx *pulumi.Context,
	name string, args *ActionCustomArgs, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	var resource ActionCustom
	err := ctx.RegisterResource("azure:logicapps/actionCustom:ActionCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionCustom gets an existing ActionCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionCustomState, opts ...pulumi.ResourceOption) (*ActionCustom, error) {
	var resource ActionCustom
	err := ctx.ReadResource("azure:logicapps/actionCustom:ActionCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionCustom resources.
type actionCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body *string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

type ActionCustomState struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ActionCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionCustomState)(nil)).Elem()
}

type actionCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body string `pulumi:"body"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ActionCustom resource.
type ActionCustomArgs struct {
	// Specifies the JSON Blob defining the Body of this Custom Action.
	Body pulumi.StringInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (ActionCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionCustomArgs)(nil)).Elem()
}

type ActionCustomInput interface {
	pulumi.Input

	ToActionCustomOutput() ActionCustomOutput
	ToActionCustomOutputWithContext(ctx context.Context) ActionCustomOutput
}

func (*ActionCustom) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionCustom)(nil))
}

func (i *ActionCustom) ToActionCustomOutput() ActionCustomOutput {
	return i.ToActionCustomOutputWithContext(context.Background())
}

func (i *ActionCustom) ToActionCustomOutputWithContext(ctx context.Context) ActionCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionCustomOutput)
}

func (i *ActionCustom) ToActionCustomPtrOutput() ActionCustomPtrOutput {
	return i.ToActionCustomPtrOutputWithContext(context.Background())
}

func (i *ActionCustom) ToActionCustomPtrOutputWithContext(ctx context.Context) ActionCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionCustomPtrOutput)
}

type ActionCustomPtrInput interface {
	pulumi.Input

	ToActionCustomPtrOutput() ActionCustomPtrOutput
	ToActionCustomPtrOutputWithContext(ctx context.Context) ActionCustomPtrOutput
}

type actionCustomPtrType ActionCustomArgs

func (*actionCustomPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionCustom)(nil))
}

func (i *actionCustomPtrType) ToActionCustomPtrOutput() ActionCustomPtrOutput {
	return i.ToActionCustomPtrOutputWithContext(context.Background())
}

func (i *actionCustomPtrType) ToActionCustomPtrOutputWithContext(ctx context.Context) ActionCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionCustomPtrOutput)
}

// ActionCustomArrayInput is an input type that accepts ActionCustomArray and ActionCustomArrayOutput values.
// You can construct a concrete instance of `ActionCustomArrayInput` via:
//
//          ActionCustomArray{ ActionCustomArgs{...} }
type ActionCustomArrayInput interface {
	pulumi.Input

	ToActionCustomArrayOutput() ActionCustomArrayOutput
	ToActionCustomArrayOutputWithContext(context.Context) ActionCustomArrayOutput
}

type ActionCustomArray []ActionCustomInput

func (ActionCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionCustom)(nil)).Elem()
}

func (i ActionCustomArray) ToActionCustomArrayOutput() ActionCustomArrayOutput {
	return i.ToActionCustomArrayOutputWithContext(context.Background())
}

func (i ActionCustomArray) ToActionCustomArrayOutputWithContext(ctx context.Context) ActionCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionCustomArrayOutput)
}

// ActionCustomMapInput is an input type that accepts ActionCustomMap and ActionCustomMapOutput values.
// You can construct a concrete instance of `ActionCustomMapInput` via:
//
//          ActionCustomMap{ "key": ActionCustomArgs{...} }
type ActionCustomMapInput interface {
	pulumi.Input

	ToActionCustomMapOutput() ActionCustomMapOutput
	ToActionCustomMapOutputWithContext(context.Context) ActionCustomMapOutput
}

type ActionCustomMap map[string]ActionCustomInput

func (ActionCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionCustom)(nil)).Elem()
}

func (i ActionCustomMap) ToActionCustomMapOutput() ActionCustomMapOutput {
	return i.ToActionCustomMapOutputWithContext(context.Background())
}

func (i ActionCustomMap) ToActionCustomMapOutputWithContext(ctx context.Context) ActionCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionCustomMapOutput)
}

type ActionCustomOutput struct{ *pulumi.OutputState }

func (ActionCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionCustom)(nil))
}

func (o ActionCustomOutput) ToActionCustomOutput() ActionCustomOutput {
	return o
}

func (o ActionCustomOutput) ToActionCustomOutputWithContext(ctx context.Context) ActionCustomOutput {
	return o
}

func (o ActionCustomOutput) ToActionCustomPtrOutput() ActionCustomPtrOutput {
	return o.ToActionCustomPtrOutputWithContext(context.Background())
}

func (o ActionCustomOutput) ToActionCustomPtrOutputWithContext(ctx context.Context) ActionCustomPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionCustom) *ActionCustom {
		return &v
	}).(ActionCustomPtrOutput)
}

type ActionCustomPtrOutput struct{ *pulumi.OutputState }

func (ActionCustomPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionCustom)(nil))
}

func (o ActionCustomPtrOutput) ToActionCustomPtrOutput() ActionCustomPtrOutput {
	return o
}

func (o ActionCustomPtrOutput) ToActionCustomPtrOutputWithContext(ctx context.Context) ActionCustomPtrOutput {
	return o
}

func (o ActionCustomPtrOutput) Elem() ActionCustomOutput {
	return o.ApplyT(func(v *ActionCustom) ActionCustom {
		if v != nil {
			return *v
		}
		var ret ActionCustom
		return ret
	}).(ActionCustomOutput)
}

type ActionCustomArrayOutput struct{ *pulumi.OutputState }

func (ActionCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionCustom)(nil))
}

func (o ActionCustomArrayOutput) ToActionCustomArrayOutput() ActionCustomArrayOutput {
	return o
}

func (o ActionCustomArrayOutput) ToActionCustomArrayOutputWithContext(ctx context.Context) ActionCustomArrayOutput {
	return o
}

func (o ActionCustomArrayOutput) Index(i pulumi.IntInput) ActionCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionCustom {
		return vs[0].([]ActionCustom)[vs[1].(int)]
	}).(ActionCustomOutput)
}

type ActionCustomMapOutput struct{ *pulumi.OutputState }

func (ActionCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActionCustom)(nil))
}

func (o ActionCustomMapOutput) ToActionCustomMapOutput() ActionCustomMapOutput {
	return o
}

func (o ActionCustomMapOutput) ToActionCustomMapOutputWithContext(ctx context.Context) ActionCustomMapOutput {
	return o
}

func (o ActionCustomMapOutput) MapIndex(k pulumi.StringInput) ActionCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActionCustom {
		return vs[0].(map[string]ActionCustom)[vs[1].(string)]
	}).(ActionCustomOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionCustomOutput{})
	pulumi.RegisterOutputType(ActionCustomPtrOutput{})
	pulumi.RegisterOutputType(ActionCustomArrayOutput{})
	pulumi.RegisterOutputType(ActionCustomMapOutput{})
}
