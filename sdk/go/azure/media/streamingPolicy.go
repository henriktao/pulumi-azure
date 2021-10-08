// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Streaming Policy.
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
// 		_, err = media.NewStreamingPolicy(ctx, "exampleStreamingPolicy", &media.StreamingPolicyArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			CommonEncryptionCenc: &media.StreamingPolicyCommonEncryptionCencArgs{
// 				EnabledProtocols: &media.StreamingPolicyCommonEncryptionCencEnabledProtocolsArgs{
// 					Download:        pulumi.Bool(false),
// 					Dash:            pulumi.Bool(true),
// 					Hls:             pulumi.Bool(false),
// 					SmoothStreaming: pulumi.Bool(false),
// 				},
// 				DrmPlayready: &media.StreamingPolicyCommonEncryptionCencDrmPlayreadyArgs{
// 					CustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
// 					CustomAttributes:                    pulumi.String("PlayReady CustomAttributes"),
// 				},
// 				DrmWidevineCustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId}"),
// 			},
// 			CommonEncryptionCbcs: &media.StreamingPolicyCommonEncryptionCbcsArgs{
// 				EnabledProtocols: &media.StreamingPolicyCommonEncryptionCbcsEnabledProtocolsArgs{
// 					Download:        pulumi.Bool(false),
// 					Dash:            pulumi.Bool(true),
// 					Hls:             pulumi.Bool(false),
// 					SmoothStreaming: pulumi.Bool(false),
// 				},
// 				DrmFairplay: &media.StreamingPolicyCommonEncryptionCbcsDrmFairplayArgs{
// 					CustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
// 					AllowPersistentLicense:              pulumi.Bool(true),
// 				},
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
// Streaming Policies can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/streamingPolicy:StreamingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/streamingpolicies/policy1
// ```
type StreamingPolicy struct {
	pulumi.CustomResourceState

	// A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCbcs StreamingPolicyCommonEncryptionCbcsPtrOutput `pulumi:"commonEncryptionCbcs"`
	// A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCenc StreamingPolicyCommonEncryptionCencPtrOutput `pulumi:"commonEncryptionCenc"`
	// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrOutput `pulumi:"defaultContentKeyPolicyName"`
	// The Media Services account name. Changing this forces a new Streaming Policy to be created.
	MediaServicesAccountName pulumi.StringOutput `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
	NoEncryptionEnabledProtocols StreamingPolicyNoEncryptionEnabledProtocolsPtrOutput `pulumi:"noEncryptionEnabledProtocols"`
	// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewStreamingPolicy registers a new resource with the given unique name, arguments, and options.
func NewStreamingPolicy(ctx *pulumi.Context,
	name string, args *StreamingPolicyArgs, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MediaServicesAccountName == nil {
		return nil, errors.New("invalid value for required argument 'MediaServicesAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource StreamingPolicy
	err := ctx.RegisterResource("azure:media/streamingPolicy:StreamingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingPolicy gets an existing StreamingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingPolicyState, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	var resource StreamingPolicy
	err := ctx.ReadResource("azure:media/streamingPolicy:StreamingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingPolicy resources.
type streamingPolicyState struct {
	// A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCbcs *StreamingPolicyCommonEncryptionCbcs `pulumi:"commonEncryptionCbcs"`
	// A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCenc *StreamingPolicyCommonEncryptionCenc `pulumi:"commonEncryptionCenc"`
	// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The Media Services account name. Changing this forces a new Streaming Policy to be created.
	MediaServicesAccountName *string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
	Name *string `pulumi:"name"`
	// A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
	NoEncryptionEnabledProtocols *StreamingPolicyNoEncryptionEnabledProtocols `pulumi:"noEncryptionEnabledProtocols"`
	// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type StreamingPolicyState struct {
	// A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCbcs StreamingPolicyCommonEncryptionCbcsPtrInput
	// A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCenc StreamingPolicyCommonEncryptionCencPtrInput
	// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Streaming Policy to be created.
	MediaServicesAccountName pulumi.StringPtrInput
	// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
	Name pulumi.StringPtrInput
	// A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
	NoEncryptionEnabledProtocols StreamingPolicyNoEncryptionEnabledProtocolsPtrInput
	// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (StreamingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyState)(nil)).Elem()
}

type streamingPolicyArgs struct {
	// A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCbcs *StreamingPolicyCommonEncryptionCbcs `pulumi:"commonEncryptionCbcs"`
	// A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCenc *StreamingPolicyCommonEncryptionCenc `pulumi:"commonEncryptionCenc"`
	// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The Media Services account name. Changing this forces a new Streaming Policy to be created.
	MediaServicesAccountName string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
	Name *string `pulumi:"name"`
	// A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
	NoEncryptionEnabledProtocols *StreamingPolicyNoEncryptionEnabledProtocols `pulumi:"noEncryptionEnabledProtocols"`
	// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StreamingPolicy resource.
type StreamingPolicyArgs struct {
	// A `commonEncryptionCbcs` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCbcs StreamingPolicyCommonEncryptionCbcsPtrInput
	// A `commonEncryptionCenc` block as defined below. Changing this forces a new Streaming Policy to be created.
	CommonEncryptionCenc StreamingPolicyCommonEncryptionCencPtrInput
	// Default Content Key used by current Streaming Policy. Changing this forces a new Streaming Policy to be created.
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Streaming Policy to be created.
	MediaServicesAccountName pulumi.StringInput
	// The name which should be used for this Streaming Policy. Changing this forces a new Streaming Policy to be created.
	Name pulumi.StringPtrInput
	// A `noEncryptionEnabledProtocols` block as defined below. Changing this forces a new Streaming Policy to be created.
	NoEncryptionEnabledProtocols StreamingPolicyNoEncryptionEnabledProtocolsPtrInput
	// The name of the Resource Group where the Streaming Policy should exist. Changing this forces a new Streaming Policy to be created.
	ResourceGroupName pulumi.StringInput
}

func (StreamingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyArgs)(nil)).Elem()
}

type StreamingPolicyInput interface {
	pulumi.Input

	ToStreamingPolicyOutput() StreamingPolicyOutput
	ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput
}

func (*StreamingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicy)(nil))
}

func (i *StreamingPolicy) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return i.ToStreamingPolicyOutputWithContext(context.Background())
}

func (i *StreamingPolicy) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyOutput)
}

func (i *StreamingPolicy) ToStreamingPolicyPtrOutput() StreamingPolicyPtrOutput {
	return i.ToStreamingPolicyPtrOutputWithContext(context.Background())
}

func (i *StreamingPolicy) ToStreamingPolicyPtrOutputWithContext(ctx context.Context) StreamingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPtrOutput)
}

type StreamingPolicyPtrInput interface {
	pulumi.Input

	ToStreamingPolicyPtrOutput() StreamingPolicyPtrOutput
	ToStreamingPolicyPtrOutputWithContext(ctx context.Context) StreamingPolicyPtrOutput
}

type streamingPolicyPtrType StreamingPolicyArgs

func (*streamingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil))
}

func (i *streamingPolicyPtrType) ToStreamingPolicyPtrOutput() StreamingPolicyPtrOutput {
	return i.ToStreamingPolicyPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyPtrType) ToStreamingPolicyPtrOutputWithContext(ctx context.Context) StreamingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPtrOutput)
}

// StreamingPolicyArrayInput is an input type that accepts StreamingPolicyArray and StreamingPolicyArrayOutput values.
// You can construct a concrete instance of `StreamingPolicyArrayInput` via:
//
//          StreamingPolicyArray{ StreamingPolicyArgs{...} }
type StreamingPolicyArrayInput interface {
	pulumi.Input

	ToStreamingPolicyArrayOutput() StreamingPolicyArrayOutput
	ToStreamingPolicyArrayOutputWithContext(context.Context) StreamingPolicyArrayOutput
}

type StreamingPolicyArray []StreamingPolicyInput

func (StreamingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamingPolicy)(nil)).Elem()
}

func (i StreamingPolicyArray) ToStreamingPolicyArrayOutput() StreamingPolicyArrayOutput {
	return i.ToStreamingPolicyArrayOutputWithContext(context.Background())
}

func (i StreamingPolicyArray) ToStreamingPolicyArrayOutputWithContext(ctx context.Context) StreamingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyArrayOutput)
}

// StreamingPolicyMapInput is an input type that accepts StreamingPolicyMap and StreamingPolicyMapOutput values.
// You can construct a concrete instance of `StreamingPolicyMapInput` via:
//
//          StreamingPolicyMap{ "key": StreamingPolicyArgs{...} }
type StreamingPolicyMapInput interface {
	pulumi.Input

	ToStreamingPolicyMapOutput() StreamingPolicyMapOutput
	ToStreamingPolicyMapOutputWithContext(context.Context) StreamingPolicyMapOutput
}

type StreamingPolicyMap map[string]StreamingPolicyInput

func (StreamingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamingPolicy)(nil)).Elem()
}

func (i StreamingPolicyMap) ToStreamingPolicyMapOutput() StreamingPolicyMapOutput {
	return i.ToStreamingPolicyMapOutputWithContext(context.Background())
}

func (i StreamingPolicyMap) ToStreamingPolicyMapOutputWithContext(ctx context.Context) StreamingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyMapOutput)
}

type StreamingPolicyOutput struct{ *pulumi.OutputState }

func (StreamingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicy)(nil))
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return o
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return o
}

func (o StreamingPolicyOutput) ToStreamingPolicyPtrOutput() StreamingPolicyPtrOutput {
	return o.ToStreamingPolicyPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyOutput) ToStreamingPolicyPtrOutputWithContext(ctx context.Context) StreamingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicy) *StreamingPolicy {
		return &v
	}).(StreamingPolicyPtrOutput)
}

type StreamingPolicyPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil))
}

func (o StreamingPolicyPtrOutput) ToStreamingPolicyPtrOutput() StreamingPolicyPtrOutput {
	return o
}

func (o StreamingPolicyPtrOutput) ToStreamingPolicyPtrOutputWithContext(ctx context.Context) StreamingPolicyPtrOutput {
	return o
}

func (o StreamingPolicyPtrOutput) Elem() StreamingPolicyOutput {
	return o.ApplyT(func(v *StreamingPolicy) StreamingPolicy {
		if v != nil {
			return *v
		}
		var ret StreamingPolicy
		return ret
	}).(StreamingPolicyOutput)
}

type StreamingPolicyArrayOutput struct{ *pulumi.OutputState }

func (StreamingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicy)(nil))
}

func (o StreamingPolicyArrayOutput) ToStreamingPolicyArrayOutput() StreamingPolicyArrayOutput {
	return o
}

func (o StreamingPolicyArrayOutput) ToStreamingPolicyArrayOutputWithContext(ctx context.Context) StreamingPolicyArrayOutput {
	return o
}

func (o StreamingPolicyArrayOutput) Index(i pulumi.IntInput) StreamingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPolicy {
		return vs[0].([]StreamingPolicy)[vs[1].(int)]
	}).(StreamingPolicyOutput)
}

type StreamingPolicyMapOutput struct{ *pulumi.OutputState }

func (StreamingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamingPolicy)(nil))
}

func (o StreamingPolicyMapOutput) ToStreamingPolicyMapOutput() StreamingPolicyMapOutput {
	return o
}

func (o StreamingPolicyMapOutput) ToStreamingPolicyMapOutputWithContext(ctx context.Context) StreamingPolicyMapOutput {
	return o
}

func (o StreamingPolicyMapOutput) MapIndex(k pulumi.StringInput) StreamingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StreamingPolicy {
		return vs[0].(map[string]StreamingPolicy)[vs[1].(string)]
	}).(StreamingPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingPolicyOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyMapOutput{})
}
