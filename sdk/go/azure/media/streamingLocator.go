// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Media Streaming Locator.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/media"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServiceAccount, err := media.NewServiceAccount(ctx, "exampleServiceAccount", &media.ServiceAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StorageAccounts: media.ServiceAccountStorageAccountArray{
// 				&media.ServiceAccountStorageAccountArgs{
// 					Id:        exampleAccount.ID(),
// 					IsPrimary: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAsset, err := media.NewAsset(ctx, "exampleAsset", &media.AssetArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			Description:              pulumi.String("Asset description"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = media.NewStreamingLocator(ctx, "exampleStreamingLocator", &media.StreamingLocatorArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			AssetName:                exampleAsset.Name,
// 			StreamingPolicyName:      pulumi.String("Predefined_ClearStreamingOnly"),
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
// Streaming Locators can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/streamingLocator:StreamingLocator example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/streaminglocators/locator1
// ```
type StreamingLocator struct {
	pulumi.CustomResourceState

	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaId pulumi.StringPtrOutput `pulumi:"alternativeMediaId"`
	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName pulumi.StringOutput `pulumi:"assetName"`
	// One or more `contentKey` blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKeys StreamingLocatorContentKeyArrayOutput `pulumi:"contentKeys"`
	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrOutput `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName pulumi.StringOutput `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorId pulumi.StringOutput `pulumi:"streamingLocatorId"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: `Predefined_DownloadOnly`, `Predefined_ClearStreamingOnly`, `Predefined_DownloadAndClearStreaming`, `Predefined_ClearKey`, `Predefined_MultiDrmCencStreaming` and `Predefined_MultiDrmStreaming`. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName pulumi.StringOutput `pulumi:"streamingPolicyName"`
}

// NewStreamingLocator registers a new resource with the given unique name, arguments, and options.
func NewStreamingLocator(ctx *pulumi.Context,
	name string, args *StreamingLocatorArgs, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.MediaServicesAccountName == nil {
		return nil, errors.New("invalid value for required argument 'MediaServicesAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StreamingPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'StreamingPolicyName'")
	}
	var resource StreamingLocator
	err := ctx.RegisterResource("azure:media/streamingLocator:StreamingLocator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingLocator gets an existing StreamingLocator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingLocator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingLocatorState, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	var resource StreamingLocator
	err := ctx.ReadResource("azure:media/streamingLocator:StreamingLocator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingLocator resources.
type streamingLocatorState struct {
	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaId *string `pulumi:"alternativeMediaId"`
	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName *string `pulumi:"assetName"`
	// One or more `contentKey` blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKeys []StreamingLocatorContentKey `pulumi:"contentKeys"`
	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime *string `pulumi:"endTime"`
	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName *string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime *string `pulumi:"startTime"`
	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorId *string `pulumi:"streamingLocatorId"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: `Predefined_DownloadOnly`, `Predefined_ClearStreamingOnly`, `Predefined_DownloadAndClearStreaming`, `Predefined_ClearKey`, `Predefined_MultiDrmCencStreaming` and `Predefined_MultiDrmStreaming`. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName *string `pulumi:"streamingPolicyName"`
}

type StreamingLocatorState struct {
	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaId pulumi.StringPtrInput
	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName pulumi.StringPtrInput
	// One or more `contentKey` blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKeys StreamingLocatorContentKeyArrayInput
	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName pulumi.StringPtrInput
	// The name which should be used for this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime pulumi.StringPtrInput
	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorId pulumi.StringPtrInput
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: `Predefined_DownloadOnly`, `Predefined_ClearStreamingOnly`, `Predefined_DownloadAndClearStreaming`, `Predefined_ClearKey`, `Predefined_MultiDrmCencStreaming` and `Predefined_MultiDrmStreaming`. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName pulumi.StringPtrInput
}

func (StreamingLocatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorState)(nil)).Elem()
}

type streamingLocatorArgs struct {
	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaId *string `pulumi:"alternativeMediaId"`
	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName string `pulumi:"assetName"`
	// One or more `contentKey` blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKeys []StreamingLocatorContentKey `pulumi:"contentKeys"`
	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime *string `pulumi:"endTime"`
	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime *string `pulumi:"startTime"`
	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorId *string `pulumi:"streamingLocatorId"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: `Predefined_DownloadOnly`, `Predefined_ClearStreamingOnly`, `Predefined_DownloadAndClearStreaming`, `Predefined_ClearKey`, `Predefined_MultiDrmCencStreaming` and `Predefined_MultiDrmStreaming`. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
}

// The set of arguments for constructing a StreamingLocator resource.
type StreamingLocatorArgs struct {
	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaId pulumi.StringPtrInput
	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName pulumi.StringInput
	// One or more `contentKey` blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKeys StreamingLocatorContentKeyArrayInput
	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName pulumi.StringInput
	// The name which should be used for this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName pulumi.StringInput
	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime pulumi.StringPtrInput
	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorId pulumi.StringPtrInput
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: `Predefined_DownloadOnly`, `Predefined_ClearStreamingOnly`, `Predefined_DownloadAndClearStreaming`, `Predefined_ClearKey`, `Predefined_MultiDrmCencStreaming` and `Predefined_MultiDrmStreaming`. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName pulumi.StringInput
}

func (StreamingLocatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorArgs)(nil)).Elem()
}

type StreamingLocatorInput interface {
	pulumi.Input

	ToStreamingLocatorOutput() StreamingLocatorOutput
	ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput
}

func (*StreamingLocator) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (i *StreamingLocator) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return i.ToStreamingLocatorOutputWithContext(context.Background())
}

func (i *StreamingLocator) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorOutput)
}

// StreamingLocatorArrayInput is an input type that accepts StreamingLocatorArray and StreamingLocatorArrayOutput values.
// You can construct a concrete instance of `StreamingLocatorArrayInput` via:
//
//          StreamingLocatorArray{ StreamingLocatorArgs{...} }
type StreamingLocatorArrayInput interface {
	pulumi.Input

	ToStreamingLocatorArrayOutput() StreamingLocatorArrayOutput
	ToStreamingLocatorArrayOutputWithContext(context.Context) StreamingLocatorArrayOutput
}

type StreamingLocatorArray []StreamingLocatorInput

func (StreamingLocatorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamingLocator)(nil)).Elem()
}

func (i StreamingLocatorArray) ToStreamingLocatorArrayOutput() StreamingLocatorArrayOutput {
	return i.ToStreamingLocatorArrayOutputWithContext(context.Background())
}

func (i StreamingLocatorArray) ToStreamingLocatorArrayOutputWithContext(ctx context.Context) StreamingLocatorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorArrayOutput)
}

// StreamingLocatorMapInput is an input type that accepts StreamingLocatorMap and StreamingLocatorMapOutput values.
// You can construct a concrete instance of `StreamingLocatorMapInput` via:
//
//          StreamingLocatorMap{ "key": StreamingLocatorArgs{...} }
type StreamingLocatorMapInput interface {
	pulumi.Input

	ToStreamingLocatorMapOutput() StreamingLocatorMapOutput
	ToStreamingLocatorMapOutputWithContext(context.Context) StreamingLocatorMapOutput
}

type StreamingLocatorMap map[string]StreamingLocatorInput

func (StreamingLocatorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamingLocator)(nil)).Elem()
}

func (i StreamingLocatorMap) ToStreamingLocatorMapOutput() StreamingLocatorMapOutput {
	return i.ToStreamingLocatorMapOutputWithContext(context.Background())
}

func (i StreamingLocatorMap) ToStreamingLocatorMapOutputWithContext(ctx context.Context) StreamingLocatorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorMapOutput)
}

type StreamingLocatorOutput struct{ *pulumi.OutputState }

func (StreamingLocatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return o
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return o
}

type StreamingLocatorArrayOutput struct{ *pulumi.OutputState }

func (StreamingLocatorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamingLocator)(nil)).Elem()
}

func (o StreamingLocatorArrayOutput) ToStreamingLocatorArrayOutput() StreamingLocatorArrayOutput {
	return o
}

func (o StreamingLocatorArrayOutput) ToStreamingLocatorArrayOutputWithContext(ctx context.Context) StreamingLocatorArrayOutput {
	return o
}

func (o StreamingLocatorArrayOutput) Index(i pulumi.IntInput) StreamingLocatorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StreamingLocator {
		return vs[0].([]*StreamingLocator)[vs[1].(int)]
	}).(StreamingLocatorOutput)
}

type StreamingLocatorMapOutput struct{ *pulumi.OutputState }

func (StreamingLocatorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamingLocator)(nil)).Elem()
}

func (o StreamingLocatorMapOutput) ToStreamingLocatorMapOutput() StreamingLocatorMapOutput {
	return o
}

func (o StreamingLocatorMapOutput) ToStreamingLocatorMapOutputWithContext(ctx context.Context) StreamingLocatorMapOutput {
	return o
}

func (o StreamingLocatorMapOutput) MapIndex(k pulumi.StringInput) StreamingLocatorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StreamingLocator {
		return vs[0].(map[string]*StreamingLocator)[vs[1].(string)]
	}).(StreamingLocatorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorInput)(nil)).Elem(), &StreamingLocator{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorArrayInput)(nil)).Elem(), StreamingLocatorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorMapInput)(nil)).Elem(), StreamingLocatorMap{})
	pulumi.RegisterOutputType(StreamingLocatorOutput{})
	pulumi.RegisterOutputType(StreamingLocatorArrayOutput{})
	pulumi.RegisterOutputType(StreamingLocatorMapOutput{})
}
