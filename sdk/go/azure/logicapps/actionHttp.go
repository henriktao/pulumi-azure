// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an HTTP Action within a Logic App Workflow
//
// ## Example Usage
//
// ```go
// package main
//
// import (
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
// 		_, err = logicapps.NewActionHttp(ctx, "exampleActionHttp", &logicapps.ActionHttpArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Method:     pulumi.String("GET"),
// 			Uri:        pulumi.String("http://example.com/some-webhook"),
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
// Logic App HTTP Actions can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/actionHttp:ActionHttp webhook1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
// ```
type ActionHttp struct {
	pulumi.CustomResourceState

	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapOutput `pulumi:"headers"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringOutput `pulumi:"method"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `runAfter` block is as defined below.
	RunAfters ActionHttpRunAfterArrayOutput `pulumi:"runAfters"`
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewActionHttp registers a new resource with the given unique name, arguments, and options.
func NewActionHttp(ctx *pulumi.Context,
	name string, args *ActionHttpArgs, opts ...pulumi.ResourceOption) (*ActionHttp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	if args.Uri == nil {
		return nil, errors.New("invalid value for required argument 'Uri'")
	}
	var resource ActionHttp
	err := ctx.RegisterResource("azure:logicapps/actionHttp:ActionHttp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionHttp gets an existing ActionHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionHttpState, opts ...pulumi.ResourceOption) (*ActionHttp, error) {
	var resource ActionHttp
	err := ctx.ReadResource("azure:logicapps/actionHttp:ActionHttp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionHttp resources.
type actionHttpState struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body *string `pulumi:"body"`
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers map[string]string `pulumi:"headers"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method *string `pulumi:"method"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `runAfter` block is as defined below.
	RunAfters []ActionHttpRunAfter `pulumi:"runAfters"`
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri *string `pulumi:"uri"`
}

type ActionHttpState struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringPtrInput
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringPtrInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `runAfter` block is as defined below.
	RunAfters ActionHttpRunAfterArrayInput
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringPtrInput
}

func (ActionHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionHttpState)(nil)).Elem()
}

type actionHttpArgs struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body *string `pulumi:"body"`
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers map[string]string `pulumi:"headers"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method string `pulumi:"method"`
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `runAfter` block is as defined below.
	RunAfters []ActionHttpRunAfter `pulumi:"runAfters"`
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri string `pulumi:"uri"`
}

// The set of arguments for constructing a ActionHttp resource.
type ActionHttpArgs struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body pulumi.StringPtrInput
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers pulumi.StringMapInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method pulumi.StringInput
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the place of the HTTP Action in the Logic App Workflow. If not specified, the HTTP Action is right after the Trigger. A `runAfter` block is as defined below.
	RunAfters ActionHttpRunAfterArrayInput
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri pulumi.StringInput
}

func (ActionHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionHttpArgs)(nil)).Elem()
}

type ActionHttpInput interface {
	pulumi.Input

	ToActionHttpOutput() ActionHttpOutput
	ToActionHttpOutputWithContext(ctx context.Context) ActionHttpOutput
}

func (*ActionHttp) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionHttp)(nil))
}

func (i *ActionHttp) ToActionHttpOutput() ActionHttpOutput {
	return i.ToActionHttpOutputWithContext(context.Background())
}

func (i *ActionHttp) ToActionHttpOutputWithContext(ctx context.Context) ActionHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpOutput)
}

func (i *ActionHttp) ToActionHttpPtrOutput() ActionHttpPtrOutput {
	return i.ToActionHttpPtrOutputWithContext(context.Background())
}

func (i *ActionHttp) ToActionHttpPtrOutputWithContext(ctx context.Context) ActionHttpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpPtrOutput)
}

type ActionHttpPtrInput interface {
	pulumi.Input

	ToActionHttpPtrOutput() ActionHttpPtrOutput
	ToActionHttpPtrOutputWithContext(ctx context.Context) ActionHttpPtrOutput
}

type actionHttpPtrType ActionHttpArgs

func (*actionHttpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionHttp)(nil))
}

func (i *actionHttpPtrType) ToActionHttpPtrOutput() ActionHttpPtrOutput {
	return i.ToActionHttpPtrOutputWithContext(context.Background())
}

func (i *actionHttpPtrType) ToActionHttpPtrOutputWithContext(ctx context.Context) ActionHttpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpPtrOutput)
}

// ActionHttpArrayInput is an input type that accepts ActionHttpArray and ActionHttpArrayOutput values.
// You can construct a concrete instance of `ActionHttpArrayInput` via:
//
//          ActionHttpArray{ ActionHttpArgs{...} }
type ActionHttpArrayInput interface {
	pulumi.Input

	ToActionHttpArrayOutput() ActionHttpArrayOutput
	ToActionHttpArrayOutputWithContext(context.Context) ActionHttpArrayOutput
}

type ActionHttpArray []ActionHttpInput

func (ActionHttpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionHttp)(nil)).Elem()
}

func (i ActionHttpArray) ToActionHttpArrayOutput() ActionHttpArrayOutput {
	return i.ToActionHttpArrayOutputWithContext(context.Background())
}

func (i ActionHttpArray) ToActionHttpArrayOutputWithContext(ctx context.Context) ActionHttpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpArrayOutput)
}

// ActionHttpMapInput is an input type that accepts ActionHttpMap and ActionHttpMapOutput values.
// You can construct a concrete instance of `ActionHttpMapInput` via:
//
//          ActionHttpMap{ "key": ActionHttpArgs{...} }
type ActionHttpMapInput interface {
	pulumi.Input

	ToActionHttpMapOutput() ActionHttpMapOutput
	ToActionHttpMapOutputWithContext(context.Context) ActionHttpMapOutput
}

type ActionHttpMap map[string]ActionHttpInput

func (ActionHttpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionHttp)(nil)).Elem()
}

func (i ActionHttpMap) ToActionHttpMapOutput() ActionHttpMapOutput {
	return i.ToActionHttpMapOutputWithContext(context.Background())
}

func (i ActionHttpMap) ToActionHttpMapOutputWithContext(ctx context.Context) ActionHttpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionHttpMapOutput)
}

type ActionHttpOutput struct{ *pulumi.OutputState }

func (ActionHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionHttp)(nil))
}

func (o ActionHttpOutput) ToActionHttpOutput() ActionHttpOutput {
	return o
}

func (o ActionHttpOutput) ToActionHttpOutputWithContext(ctx context.Context) ActionHttpOutput {
	return o
}

func (o ActionHttpOutput) ToActionHttpPtrOutput() ActionHttpPtrOutput {
	return o.ToActionHttpPtrOutputWithContext(context.Background())
}

func (o ActionHttpOutput) ToActionHttpPtrOutputWithContext(ctx context.Context) ActionHttpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionHttp) *ActionHttp {
		return &v
	}).(ActionHttpPtrOutput)
}

type ActionHttpPtrOutput struct{ *pulumi.OutputState }

func (ActionHttpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionHttp)(nil))
}

func (o ActionHttpPtrOutput) ToActionHttpPtrOutput() ActionHttpPtrOutput {
	return o
}

func (o ActionHttpPtrOutput) ToActionHttpPtrOutputWithContext(ctx context.Context) ActionHttpPtrOutput {
	return o
}

func (o ActionHttpPtrOutput) Elem() ActionHttpOutput {
	return o.ApplyT(func(v *ActionHttp) ActionHttp {
		if v != nil {
			return *v
		}
		var ret ActionHttp
		return ret
	}).(ActionHttpOutput)
}

type ActionHttpArrayOutput struct{ *pulumi.OutputState }

func (ActionHttpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionHttp)(nil))
}

func (o ActionHttpArrayOutput) ToActionHttpArrayOutput() ActionHttpArrayOutput {
	return o
}

func (o ActionHttpArrayOutput) ToActionHttpArrayOutputWithContext(ctx context.Context) ActionHttpArrayOutput {
	return o
}

func (o ActionHttpArrayOutput) Index(i pulumi.IntInput) ActionHttpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionHttp {
		return vs[0].([]ActionHttp)[vs[1].(int)]
	}).(ActionHttpOutput)
}

type ActionHttpMapOutput struct{ *pulumi.OutputState }

func (ActionHttpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActionHttp)(nil))
}

func (o ActionHttpMapOutput) ToActionHttpMapOutput() ActionHttpMapOutput {
	return o
}

func (o ActionHttpMapOutput) ToActionHttpMapOutputWithContext(ctx context.Context) ActionHttpMapOutput {
	return o
}

func (o ActionHttpMapOutput) MapIndex(k pulumi.StringInput) ActionHttpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActionHttp {
		return vs[0].(map[string]ActionHttp)[vs[1].(string)]
	}).(ActionHttpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionHttpInput)(nil)).Elem(), &ActionHttp{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionHttpPtrInput)(nil)).Elem(), &ActionHttp{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionHttpArrayInput)(nil)).Elem(), ActionHttpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionHttpMapInput)(nil)).Elem(), ActionHttpMap{})
	pulumi.RegisterOutputType(ActionHttpOutput{})
	pulumi.RegisterOutputType(ActionHttpPtrOutput{})
	pulumi.RegisterOutputType(ActionHttpArrayOutput{})
	pulumi.RegisterOutputType(ActionHttpMapOutput{})
}
