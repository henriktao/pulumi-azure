// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a HTTP Request Trigger within a Logic App Workflow
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
// 		_, err = logicapps.NewTriggerHttpRequest(ctx, "exampleTriggerHttpRequest", &logicapps.TriggerHttpRequestArgs{
// 			LogicAppId: exampleWorkflow.ID(),
// 			Schema:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "{\n", "    \"type\": \"object\",\n", "    \"properties\": {\n", "        \"hello\": {\n", "            \"type\": \"string\"\n", "        }\n", "    }\n", "}\n")),
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
// Logic App HTTP Request Triggers can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/triggerHttpRequest:TriggerHttpRequest request1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/request1
// ```
type TriggerHttpRequest struct {
	pulumi.CustomResourceState

	// The URL for the workflow trigger
	CallbackUrl pulumi.StringOutput `pulumi:"callbackUrl"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
	Method pulumi.StringPtrOutput `pulumi:"method"`
	// Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Relative Path used for this Request.
	RelativePath pulumi.StringPtrOutput `pulumi:"relativePath"`
	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema pulumi.StringOutput `pulumi:"schema"`
}

// NewTriggerHttpRequest registers a new resource with the given unique name, arguments, and options.
func NewTriggerHttpRequest(ctx *pulumi.Context,
	name string, args *TriggerHttpRequestArgs, opts ...pulumi.ResourceOption) (*TriggerHttpRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogicAppId == nil {
		return nil, errors.New("invalid value for required argument 'LogicAppId'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	var resource TriggerHttpRequest
	err := ctx.RegisterResource("azure:logicapps/triggerHttpRequest:TriggerHttpRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerHttpRequest gets an existing TriggerHttpRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerHttpRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerHttpRequestState, opts ...pulumi.ResourceOption) (*TriggerHttpRequest, error) {
	var resource TriggerHttpRequest
	err := ctx.ReadResource("azure:logicapps/triggerHttpRequest:TriggerHttpRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerHttpRequest resources.
type triggerHttpRequestState struct {
	// The URL for the workflow trigger
	CallbackUrl *string `pulumi:"callbackUrl"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
	Method *string `pulumi:"method"`
	// Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Relative Path used for this Request.
	RelativePath *string `pulumi:"relativePath"`
	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema *string `pulumi:"schema"`
}

type TriggerHttpRequestState struct {
	// The URL for the workflow trigger
	CallbackUrl pulumi.StringPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
	Method pulumi.StringPtrInput
	// Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Relative Path used for this Request.
	RelativePath pulumi.StringPtrInput
	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema pulumi.StringPtrInput
}

func (TriggerHttpRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerHttpRequestState)(nil)).Elem()
}

type triggerHttpRequestArgs struct {
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
	Method *string `pulumi:"method"`
	// Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Relative Path used for this Request.
	RelativePath *string `pulumi:"relativePath"`
	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a TriggerHttpRequest resource.
type TriggerHttpRequestArgs struct {
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the HTTP Method which the request be using. Possible values include `DELETE`, `GET`, `PATCH`, `POST` or `PUT`.
	Method pulumi.StringPtrInput
	// Specifies the name of the HTTP Request Trigger to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Relative Path used for this Request.
	RelativePath pulumi.StringPtrInput
	// A JSON Blob defining the Schema of the incoming request. This needs to be valid JSON.
	Schema pulumi.StringInput
}

func (TriggerHttpRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerHttpRequestArgs)(nil)).Elem()
}

type TriggerHttpRequestInput interface {
	pulumi.Input

	ToTriggerHttpRequestOutput() TriggerHttpRequestOutput
	ToTriggerHttpRequestOutputWithContext(ctx context.Context) TriggerHttpRequestOutput
}

func (*TriggerHttpRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerHttpRequest)(nil))
}

func (i *TriggerHttpRequest) ToTriggerHttpRequestOutput() TriggerHttpRequestOutput {
	return i.ToTriggerHttpRequestOutputWithContext(context.Background())
}

func (i *TriggerHttpRequest) ToTriggerHttpRequestOutputWithContext(ctx context.Context) TriggerHttpRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerHttpRequestOutput)
}

func (i *TriggerHttpRequest) ToTriggerHttpRequestPtrOutput() TriggerHttpRequestPtrOutput {
	return i.ToTriggerHttpRequestPtrOutputWithContext(context.Background())
}

func (i *TriggerHttpRequest) ToTriggerHttpRequestPtrOutputWithContext(ctx context.Context) TriggerHttpRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerHttpRequestPtrOutput)
}

type TriggerHttpRequestPtrInput interface {
	pulumi.Input

	ToTriggerHttpRequestPtrOutput() TriggerHttpRequestPtrOutput
	ToTriggerHttpRequestPtrOutputWithContext(ctx context.Context) TriggerHttpRequestPtrOutput
}

type triggerHttpRequestPtrType TriggerHttpRequestArgs

func (*triggerHttpRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerHttpRequest)(nil))
}

func (i *triggerHttpRequestPtrType) ToTriggerHttpRequestPtrOutput() TriggerHttpRequestPtrOutput {
	return i.ToTriggerHttpRequestPtrOutputWithContext(context.Background())
}

func (i *triggerHttpRequestPtrType) ToTriggerHttpRequestPtrOutputWithContext(ctx context.Context) TriggerHttpRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerHttpRequestPtrOutput)
}

// TriggerHttpRequestArrayInput is an input type that accepts TriggerHttpRequestArray and TriggerHttpRequestArrayOutput values.
// You can construct a concrete instance of `TriggerHttpRequestArrayInput` via:
//
//          TriggerHttpRequestArray{ TriggerHttpRequestArgs{...} }
type TriggerHttpRequestArrayInput interface {
	pulumi.Input

	ToTriggerHttpRequestArrayOutput() TriggerHttpRequestArrayOutput
	ToTriggerHttpRequestArrayOutputWithContext(context.Context) TriggerHttpRequestArrayOutput
}

type TriggerHttpRequestArray []TriggerHttpRequestInput

func (TriggerHttpRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TriggerHttpRequest)(nil)).Elem()
}

func (i TriggerHttpRequestArray) ToTriggerHttpRequestArrayOutput() TriggerHttpRequestArrayOutput {
	return i.ToTriggerHttpRequestArrayOutputWithContext(context.Background())
}

func (i TriggerHttpRequestArray) ToTriggerHttpRequestArrayOutputWithContext(ctx context.Context) TriggerHttpRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerHttpRequestArrayOutput)
}

// TriggerHttpRequestMapInput is an input type that accepts TriggerHttpRequestMap and TriggerHttpRequestMapOutput values.
// You can construct a concrete instance of `TriggerHttpRequestMapInput` via:
//
//          TriggerHttpRequestMap{ "key": TriggerHttpRequestArgs{...} }
type TriggerHttpRequestMapInput interface {
	pulumi.Input

	ToTriggerHttpRequestMapOutput() TriggerHttpRequestMapOutput
	ToTriggerHttpRequestMapOutputWithContext(context.Context) TriggerHttpRequestMapOutput
}

type TriggerHttpRequestMap map[string]TriggerHttpRequestInput

func (TriggerHttpRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TriggerHttpRequest)(nil)).Elem()
}

func (i TriggerHttpRequestMap) ToTriggerHttpRequestMapOutput() TriggerHttpRequestMapOutput {
	return i.ToTriggerHttpRequestMapOutputWithContext(context.Background())
}

func (i TriggerHttpRequestMap) ToTriggerHttpRequestMapOutputWithContext(ctx context.Context) TriggerHttpRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerHttpRequestMapOutput)
}

type TriggerHttpRequestOutput struct{ *pulumi.OutputState }

func (TriggerHttpRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerHttpRequest)(nil))
}

func (o TriggerHttpRequestOutput) ToTriggerHttpRequestOutput() TriggerHttpRequestOutput {
	return o
}

func (o TriggerHttpRequestOutput) ToTriggerHttpRequestOutputWithContext(ctx context.Context) TriggerHttpRequestOutput {
	return o
}

func (o TriggerHttpRequestOutput) ToTriggerHttpRequestPtrOutput() TriggerHttpRequestPtrOutput {
	return o.ToTriggerHttpRequestPtrOutputWithContext(context.Background())
}

func (o TriggerHttpRequestOutput) ToTriggerHttpRequestPtrOutputWithContext(ctx context.Context) TriggerHttpRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerHttpRequest) *TriggerHttpRequest {
		return &v
	}).(TriggerHttpRequestPtrOutput)
}

type TriggerHttpRequestPtrOutput struct{ *pulumi.OutputState }

func (TriggerHttpRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerHttpRequest)(nil))
}

func (o TriggerHttpRequestPtrOutput) ToTriggerHttpRequestPtrOutput() TriggerHttpRequestPtrOutput {
	return o
}

func (o TriggerHttpRequestPtrOutput) ToTriggerHttpRequestPtrOutputWithContext(ctx context.Context) TriggerHttpRequestPtrOutput {
	return o
}

func (o TriggerHttpRequestPtrOutput) Elem() TriggerHttpRequestOutput {
	return o.ApplyT(func(v *TriggerHttpRequest) TriggerHttpRequest {
		if v != nil {
			return *v
		}
		var ret TriggerHttpRequest
		return ret
	}).(TriggerHttpRequestOutput)
}

type TriggerHttpRequestArrayOutput struct{ *pulumi.OutputState }

func (TriggerHttpRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerHttpRequest)(nil))
}

func (o TriggerHttpRequestArrayOutput) ToTriggerHttpRequestArrayOutput() TriggerHttpRequestArrayOutput {
	return o
}

func (o TriggerHttpRequestArrayOutput) ToTriggerHttpRequestArrayOutputWithContext(ctx context.Context) TriggerHttpRequestArrayOutput {
	return o
}

func (o TriggerHttpRequestArrayOutput) Index(i pulumi.IntInput) TriggerHttpRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerHttpRequest {
		return vs[0].([]TriggerHttpRequest)[vs[1].(int)]
	}).(TriggerHttpRequestOutput)
}

type TriggerHttpRequestMapOutput struct{ *pulumi.OutputState }

func (TriggerHttpRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TriggerHttpRequest)(nil))
}

func (o TriggerHttpRequestMapOutput) ToTriggerHttpRequestMapOutput() TriggerHttpRequestMapOutput {
	return o
}

func (o TriggerHttpRequestMapOutput) ToTriggerHttpRequestMapOutputWithContext(ctx context.Context) TriggerHttpRequestMapOutput {
	return o
}

func (o TriggerHttpRequestMapOutput) MapIndex(k pulumi.StringInput) TriggerHttpRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TriggerHttpRequest {
		return vs[0].(map[string]TriggerHttpRequest)[vs[1].(string)]
	}).(TriggerHttpRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerHttpRequestInput)(nil)).Elem(), &TriggerHttpRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerHttpRequestPtrInput)(nil)).Elem(), &TriggerHttpRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerHttpRequestArrayInput)(nil)).Elem(), TriggerHttpRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerHttpRequestMapInput)(nil)).Elem(), TriggerHttpRequestMap{})
	pulumi.RegisterOutputType(TriggerHttpRequestOutput{})
	pulumi.RegisterOutputType(TriggerHttpRequestPtrOutput{})
	pulumi.RegisterOutputType(TriggerHttpRequestArrayOutput{})
	pulumi.RegisterOutputType(TriggerHttpRequestMapOutput{})
}
