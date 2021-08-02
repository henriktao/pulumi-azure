// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Direct Line Speech integration for a Bot Channel
//
// ## Import
//
// Direct Line Speech Channels can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
// ```
type ChannelDirectLineSpeech struct {
	pulumi.CustomResourceState

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringOutput `pulumi:"botName"`
	// The access key to access the Cognitive Service.
	CognitiveServiceAccessKey pulumi.StringOutput `pulumi:"cognitiveServiceAccessKey"`
	// Specifies the supported Azure location where the Cognitive Service resource exists.
	CognitiveServiceLocation pulumi.StringOutput `pulumi:"cognitiveServiceLocation"`
	// The custom speech model id for the Direct Line Speech Channel.
	CustomSpeechModelId pulumi.StringPtrOutput `pulumi:"customSpeechModelId"`
	// The custom voice deployment id for the Direct Line Speech Channel.
	CustomVoiceDeploymentId pulumi.StringPtrOutput `pulumi:"customVoiceDeploymentId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewChannelDirectLineSpeech registers a new resource with the given unique name, arguments, and options.
func NewChannelDirectLineSpeech(ctx *pulumi.Context,
	name string, args *ChannelDirectLineSpeechArgs, opts ...pulumi.ResourceOption) (*ChannelDirectLineSpeech, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BotName == nil {
		return nil, errors.New("invalid value for required argument 'BotName'")
	}
	if args.CognitiveServiceAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'CognitiveServiceAccessKey'")
	}
	if args.CognitiveServiceLocation == nil {
		return nil, errors.New("invalid value for required argument 'CognitiveServiceLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ChannelDirectLineSpeech
	err := ctx.RegisterResource("azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelDirectLineSpeech gets an existing ChannelDirectLineSpeech resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelDirectLineSpeech(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelDirectLineSpeechState, opts ...pulumi.ResourceOption) (*ChannelDirectLineSpeech, error) {
	var resource ChannelDirectLineSpeech
	err := ctx.ReadResource("azure:bot/channelDirectLineSpeech:ChannelDirectLineSpeech", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelDirectLineSpeech resources.
type channelDirectLineSpeechState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `pulumi:"botName"`
	// The access key to access the Cognitive Service.
	CognitiveServiceAccessKey *string `pulumi:"cognitiveServiceAccessKey"`
	// Specifies the supported Azure location where the Cognitive Service resource exists.
	CognitiveServiceLocation *string `pulumi:"cognitiveServiceLocation"`
	// The custom speech model id for the Direct Line Speech Channel.
	CustomSpeechModelId *string `pulumi:"customSpeechModelId"`
	// The custom voice deployment id for the Direct Line Speech Channel.
	CustomVoiceDeploymentId *string `pulumi:"customVoiceDeploymentId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ChannelDirectLineSpeechState struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringPtrInput
	// The access key to access the Cognitive Service.
	CognitiveServiceAccessKey pulumi.StringPtrInput
	// Specifies the supported Azure location where the Cognitive Service resource exists.
	CognitiveServiceLocation pulumi.StringPtrInput
	// The custom speech model id for the Direct Line Speech Channel.
	CustomSpeechModelId pulumi.StringPtrInput
	// The custom voice deployment id for the Direct Line Speech Channel.
	CustomVoiceDeploymentId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ChannelDirectLineSpeechState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelDirectLineSpeechState)(nil)).Elem()
}

type channelDirectLineSpeechArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName string `pulumi:"botName"`
	// The access key to access the Cognitive Service.
	CognitiveServiceAccessKey string `pulumi:"cognitiveServiceAccessKey"`
	// Specifies the supported Azure location where the Cognitive Service resource exists.
	CognitiveServiceLocation string `pulumi:"cognitiveServiceLocation"`
	// The custom speech model id for the Direct Line Speech Channel.
	CustomSpeechModelId *string `pulumi:"customSpeechModelId"`
	// The custom voice deployment id for the Direct Line Speech Channel.
	CustomVoiceDeploymentId *string `pulumi:"customVoiceDeploymentId"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ChannelDirectLineSpeech resource.
type ChannelDirectLineSpeechArgs struct {
	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName pulumi.StringInput
	// The access key to access the Cognitive Service.
	CognitiveServiceAccessKey pulumi.StringInput
	// Specifies the supported Azure location where the Cognitive Service resource exists.
	CognitiveServiceLocation pulumi.StringInput
	// The custom speech model id for the Direct Line Speech Channel.
	CustomSpeechModelId pulumi.StringPtrInput
	// The custom voice deployment id for the Direct Line Speech Channel.
	CustomVoiceDeploymentId pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the resource group where the Direct Line Speech Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ChannelDirectLineSpeechArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelDirectLineSpeechArgs)(nil)).Elem()
}

type ChannelDirectLineSpeechInput interface {
	pulumi.Input

	ToChannelDirectLineSpeechOutput() ChannelDirectLineSpeechOutput
	ToChannelDirectLineSpeechOutputWithContext(ctx context.Context) ChannelDirectLineSpeechOutput
}

func (*ChannelDirectLineSpeech) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDirectLineSpeech)(nil))
}

func (i *ChannelDirectLineSpeech) ToChannelDirectLineSpeechOutput() ChannelDirectLineSpeechOutput {
	return i.ToChannelDirectLineSpeechOutputWithContext(context.Background())
}

func (i *ChannelDirectLineSpeech) ToChannelDirectLineSpeechOutputWithContext(ctx context.Context) ChannelDirectLineSpeechOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSpeechOutput)
}

func (i *ChannelDirectLineSpeech) ToChannelDirectLineSpeechPtrOutput() ChannelDirectLineSpeechPtrOutput {
	return i.ToChannelDirectLineSpeechPtrOutputWithContext(context.Background())
}

func (i *ChannelDirectLineSpeech) ToChannelDirectLineSpeechPtrOutputWithContext(ctx context.Context) ChannelDirectLineSpeechPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSpeechPtrOutput)
}

type ChannelDirectLineSpeechPtrInput interface {
	pulumi.Input

	ToChannelDirectLineSpeechPtrOutput() ChannelDirectLineSpeechPtrOutput
	ToChannelDirectLineSpeechPtrOutputWithContext(ctx context.Context) ChannelDirectLineSpeechPtrOutput
}

type channelDirectLineSpeechPtrType ChannelDirectLineSpeechArgs

func (*channelDirectLineSpeechPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelDirectLineSpeech)(nil))
}

func (i *channelDirectLineSpeechPtrType) ToChannelDirectLineSpeechPtrOutput() ChannelDirectLineSpeechPtrOutput {
	return i.ToChannelDirectLineSpeechPtrOutputWithContext(context.Background())
}

func (i *channelDirectLineSpeechPtrType) ToChannelDirectLineSpeechPtrOutputWithContext(ctx context.Context) ChannelDirectLineSpeechPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSpeechPtrOutput)
}

// ChannelDirectLineSpeechArrayInput is an input type that accepts ChannelDirectLineSpeechArray and ChannelDirectLineSpeechArrayOutput values.
// You can construct a concrete instance of `ChannelDirectLineSpeechArrayInput` via:
//
//          ChannelDirectLineSpeechArray{ ChannelDirectLineSpeechArgs{...} }
type ChannelDirectLineSpeechArrayInput interface {
	pulumi.Input

	ToChannelDirectLineSpeechArrayOutput() ChannelDirectLineSpeechArrayOutput
	ToChannelDirectLineSpeechArrayOutputWithContext(context.Context) ChannelDirectLineSpeechArrayOutput
}

type ChannelDirectLineSpeechArray []ChannelDirectLineSpeechInput

func (ChannelDirectLineSpeechArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ChannelDirectLineSpeech)(nil))
}

func (i ChannelDirectLineSpeechArray) ToChannelDirectLineSpeechArrayOutput() ChannelDirectLineSpeechArrayOutput {
	return i.ToChannelDirectLineSpeechArrayOutputWithContext(context.Background())
}

func (i ChannelDirectLineSpeechArray) ToChannelDirectLineSpeechArrayOutputWithContext(ctx context.Context) ChannelDirectLineSpeechArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSpeechArrayOutput)
}

// ChannelDirectLineSpeechMapInput is an input type that accepts ChannelDirectLineSpeechMap and ChannelDirectLineSpeechMapOutput values.
// You can construct a concrete instance of `ChannelDirectLineSpeechMapInput` via:
//
//          ChannelDirectLineSpeechMap{ "key": ChannelDirectLineSpeechArgs{...} }
type ChannelDirectLineSpeechMapInput interface {
	pulumi.Input

	ToChannelDirectLineSpeechMapOutput() ChannelDirectLineSpeechMapOutput
	ToChannelDirectLineSpeechMapOutputWithContext(context.Context) ChannelDirectLineSpeechMapOutput
}

type ChannelDirectLineSpeechMap map[string]ChannelDirectLineSpeechInput

func (ChannelDirectLineSpeechMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ChannelDirectLineSpeech)(nil))
}

func (i ChannelDirectLineSpeechMap) ToChannelDirectLineSpeechMapOutput() ChannelDirectLineSpeechMapOutput {
	return i.ToChannelDirectLineSpeechMapOutputWithContext(context.Background())
}

func (i ChannelDirectLineSpeechMap) ToChannelDirectLineSpeechMapOutputWithContext(ctx context.Context) ChannelDirectLineSpeechMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelDirectLineSpeechMapOutput)
}

type ChannelDirectLineSpeechOutput struct {
	*pulumi.OutputState
}

func (ChannelDirectLineSpeechOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDirectLineSpeech)(nil))
}

func (o ChannelDirectLineSpeechOutput) ToChannelDirectLineSpeechOutput() ChannelDirectLineSpeechOutput {
	return o
}

func (o ChannelDirectLineSpeechOutput) ToChannelDirectLineSpeechOutputWithContext(ctx context.Context) ChannelDirectLineSpeechOutput {
	return o
}

func (o ChannelDirectLineSpeechOutput) ToChannelDirectLineSpeechPtrOutput() ChannelDirectLineSpeechPtrOutput {
	return o.ToChannelDirectLineSpeechPtrOutputWithContext(context.Background())
}

func (o ChannelDirectLineSpeechOutput) ToChannelDirectLineSpeechPtrOutputWithContext(ctx context.Context) ChannelDirectLineSpeechPtrOutput {
	return o.ApplyT(func(v ChannelDirectLineSpeech) *ChannelDirectLineSpeech {
		return &v
	}).(ChannelDirectLineSpeechPtrOutput)
}

type ChannelDirectLineSpeechPtrOutput struct {
	*pulumi.OutputState
}

func (ChannelDirectLineSpeechPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelDirectLineSpeech)(nil))
}

func (o ChannelDirectLineSpeechPtrOutput) ToChannelDirectLineSpeechPtrOutput() ChannelDirectLineSpeechPtrOutput {
	return o
}

func (o ChannelDirectLineSpeechPtrOutput) ToChannelDirectLineSpeechPtrOutputWithContext(ctx context.Context) ChannelDirectLineSpeechPtrOutput {
	return o
}

type ChannelDirectLineSpeechArrayOutput struct{ *pulumi.OutputState }

func (ChannelDirectLineSpeechArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelDirectLineSpeech)(nil))
}

func (o ChannelDirectLineSpeechArrayOutput) ToChannelDirectLineSpeechArrayOutput() ChannelDirectLineSpeechArrayOutput {
	return o
}

func (o ChannelDirectLineSpeechArrayOutput) ToChannelDirectLineSpeechArrayOutputWithContext(ctx context.Context) ChannelDirectLineSpeechArrayOutput {
	return o
}

func (o ChannelDirectLineSpeechArrayOutput) Index(i pulumi.IntInput) ChannelDirectLineSpeechOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelDirectLineSpeech {
		return vs[0].([]ChannelDirectLineSpeech)[vs[1].(int)]
	}).(ChannelDirectLineSpeechOutput)
}

type ChannelDirectLineSpeechMapOutput struct{ *pulumi.OutputState }

func (ChannelDirectLineSpeechMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ChannelDirectLineSpeech)(nil))
}

func (o ChannelDirectLineSpeechMapOutput) ToChannelDirectLineSpeechMapOutput() ChannelDirectLineSpeechMapOutput {
	return o
}

func (o ChannelDirectLineSpeechMapOutput) ToChannelDirectLineSpeechMapOutputWithContext(ctx context.Context) ChannelDirectLineSpeechMapOutput {
	return o
}

func (o ChannelDirectLineSpeechMapOutput) MapIndex(k pulumi.StringInput) ChannelDirectLineSpeechOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ChannelDirectLineSpeech {
		return vs[0].(map[string]ChannelDirectLineSpeech)[vs[1].(string)]
	}).(ChannelDirectLineSpeechOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelDirectLineSpeechOutput{})
	pulumi.RegisterOutputType(ChannelDirectLineSpeechPtrOutput{})
	pulumi.RegisterOutputType(ChannelDirectLineSpeechArrayOutput{})
	pulumi.RegisterOutputType(ChannelDirectLineSpeechMapOutput{})
}
