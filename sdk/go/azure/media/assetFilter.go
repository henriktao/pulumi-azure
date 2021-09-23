// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Media Asset Filter.
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
// 		_, err = media.NewAssetFilter(ctx, "exampleAssetFilter", &media.AssetFilterArgs{
// 			AssetId:             exampleAsset.ID(),
// 			FirstQualityBitrate: pulumi.Int(128000),
// 			PresentationTimeRange: &media.AssetFilterPresentationTimeRangeArgs{
// 				StartInUnits:               pulumi.Int(0),
// 				EndInUnits:                 pulumi.Int(15),
// 				PresentationWindowInUnits:  pulumi.Int(90),
// 				LiveBackoffInUnits:         pulumi.Int(0),
// 				UnitTimescaleInMiliseconds: pulumi.Int(1000),
// 				ForceEnd:                   pulumi.Bool(false),
// 			},
// 			TrackSelections: media.AssetFilterTrackSelectionArray{
// 				&media.AssetFilterTrackSelectionArgs{
// 					Conditions: media.AssetFilterTrackSelectionConditionArray{
// 						&media.AssetFilterTrackSelectionConditionArgs{
// 							Property:  pulumi.String("Type"),
// 							Operation: pulumi.String("Equal"),
// 							Value:     pulumi.String("Audio"),
// 						},
// 						&media.AssetFilterTrackSelectionConditionArgs{
// 							Property:  pulumi.String("Language"),
// 							Operation: pulumi.String("NotEqual"),
// 							Value:     pulumi.String("en"),
// 						},
// 						&media.AssetFilterTrackSelectionConditionArgs{
// 							Property:  pulumi.String("FourCC"),
// 							Operation: pulumi.String("NotEqual"),
// 							Value:     pulumi.String("EC-3"),
// 						},
// 					},
// 				},
// 				&media.AssetFilterTrackSelectionArgs{
// 					Conditions: media.AssetFilterTrackSelectionConditionArray{
// 						&media.AssetFilterTrackSelectionConditionArgs{
// 							Property:  pulumi.String("Type"),
// 							Operation: pulumi.String("Equal"),
// 							Value:     pulumi.String("Video"),
// 						},
// 						&media.AssetFilterTrackSelectionConditionArgs{
// 							Property:  pulumi.String("Bitrate"),
// 							Operation: pulumi.String("Equal"),
// 							Value:     pulumi.String("3000000-5000000"),
// 						},
// 					},
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
// Asset Filters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/assetFilter:AssetFilter example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/assets/asset1/assetFilters/filter1
// ```
type AssetFilter struct {
	pulumi.CustomResourceState

	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetId pulumi.StringOutput `pulumi:"assetId"`
	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate pulumi.IntPtrOutput `pulumi:"firstQualityBitrate"`
	// The name which should be used for this Asset Filter. Changing this forces a new Asset Filter to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `presentationTimeRange` block as defined below.
	PresentationTimeRange AssetFilterPresentationTimeRangePtrOutput `pulumi:"presentationTimeRange"`
	// One or more `trackSelection` blocks as defined below.
	TrackSelections AssetFilterTrackSelectionArrayOutput `pulumi:"trackSelections"`
}

// NewAssetFilter registers a new resource with the given unique name, arguments, and options.
func NewAssetFilter(ctx *pulumi.Context,
	name string, args *AssetFilterArgs, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssetId == nil {
		return nil, errors.New("invalid value for required argument 'AssetId'")
	}
	var resource AssetFilter
	err := ctx.RegisterResource("azure:media/assetFilter:AssetFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssetFilter gets an existing AssetFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetFilterState, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	var resource AssetFilter
	err := ctx.ReadResource("azure:media/assetFilter:AssetFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssetFilter resources.
type assetFilterState struct {
	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetId *string `pulumi:"assetId"`
	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *int `pulumi:"firstQualityBitrate"`
	// The name which should be used for this Asset Filter. Changing this forces a new Asset Filter to be created.
	Name *string `pulumi:"name"`
	// A `presentationTimeRange` block as defined below.
	PresentationTimeRange *AssetFilterPresentationTimeRange `pulumi:"presentationTimeRange"`
	// One or more `trackSelection` blocks as defined below.
	TrackSelections []AssetFilterTrackSelection `pulumi:"trackSelections"`
}

type AssetFilterState struct {
	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetId pulumi.StringPtrInput
	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate pulumi.IntPtrInput
	// The name which should be used for this Asset Filter. Changing this forces a new Asset Filter to be created.
	Name pulumi.StringPtrInput
	// A `presentationTimeRange` block as defined below.
	PresentationTimeRange AssetFilterPresentationTimeRangePtrInput
	// One or more `trackSelection` blocks as defined below.
	TrackSelections AssetFilterTrackSelectionArrayInput
}

func (AssetFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterState)(nil)).Elem()
}

type assetFilterArgs struct {
	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetId string `pulumi:"assetId"`
	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *int `pulumi:"firstQualityBitrate"`
	// The name which should be used for this Asset Filter. Changing this forces a new Asset Filter to be created.
	Name *string `pulumi:"name"`
	// A `presentationTimeRange` block as defined below.
	PresentationTimeRange *AssetFilterPresentationTimeRange `pulumi:"presentationTimeRange"`
	// One or more `trackSelection` blocks as defined below.
	TrackSelections []AssetFilterTrackSelection `pulumi:"trackSelections"`
}

// The set of arguments for constructing a AssetFilter resource.
type AssetFilterArgs struct {
	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetId pulumi.StringInput
	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate pulumi.IntPtrInput
	// The name which should be used for this Asset Filter. Changing this forces a new Asset Filter to be created.
	Name pulumi.StringPtrInput
	// A `presentationTimeRange` block as defined below.
	PresentationTimeRange AssetFilterPresentationTimeRangePtrInput
	// One or more `trackSelection` blocks as defined below.
	TrackSelections AssetFilterTrackSelectionArrayInput
}

func (AssetFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterArgs)(nil)).Elem()
}

type AssetFilterInput interface {
	pulumi.Input

	ToAssetFilterOutput() AssetFilterOutput
	ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput
}

func (*AssetFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetFilter)(nil))
}

func (i *AssetFilter) ToAssetFilterOutput() AssetFilterOutput {
	return i.ToAssetFilterOutputWithContext(context.Background())
}

func (i *AssetFilter) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterOutput)
}

func (i *AssetFilter) ToAssetFilterPtrOutput() AssetFilterPtrOutput {
	return i.ToAssetFilterPtrOutputWithContext(context.Background())
}

func (i *AssetFilter) ToAssetFilterPtrOutputWithContext(ctx context.Context) AssetFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterPtrOutput)
}

type AssetFilterPtrInput interface {
	pulumi.Input

	ToAssetFilterPtrOutput() AssetFilterPtrOutput
	ToAssetFilterPtrOutputWithContext(ctx context.Context) AssetFilterPtrOutput
}

type assetFilterPtrType AssetFilterArgs

func (*assetFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil))
}

func (i *assetFilterPtrType) ToAssetFilterPtrOutput() AssetFilterPtrOutput {
	return i.ToAssetFilterPtrOutputWithContext(context.Background())
}

func (i *assetFilterPtrType) ToAssetFilterPtrOutputWithContext(ctx context.Context) AssetFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterPtrOutput)
}

// AssetFilterArrayInput is an input type that accepts AssetFilterArray and AssetFilterArrayOutput values.
// You can construct a concrete instance of `AssetFilterArrayInput` via:
//
//          AssetFilterArray{ AssetFilterArgs{...} }
type AssetFilterArrayInput interface {
	pulumi.Input

	ToAssetFilterArrayOutput() AssetFilterArrayOutput
	ToAssetFilterArrayOutputWithContext(context.Context) AssetFilterArrayOutput
}

type AssetFilterArray []AssetFilterInput

func (AssetFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssetFilter)(nil)).Elem()
}

func (i AssetFilterArray) ToAssetFilterArrayOutput() AssetFilterArrayOutput {
	return i.ToAssetFilterArrayOutputWithContext(context.Background())
}

func (i AssetFilterArray) ToAssetFilterArrayOutputWithContext(ctx context.Context) AssetFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterArrayOutput)
}

// AssetFilterMapInput is an input type that accepts AssetFilterMap and AssetFilterMapOutput values.
// You can construct a concrete instance of `AssetFilterMapInput` via:
//
//          AssetFilterMap{ "key": AssetFilterArgs{...} }
type AssetFilterMapInput interface {
	pulumi.Input

	ToAssetFilterMapOutput() AssetFilterMapOutput
	ToAssetFilterMapOutputWithContext(context.Context) AssetFilterMapOutput
}

type AssetFilterMap map[string]AssetFilterInput

func (AssetFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssetFilter)(nil)).Elem()
}

func (i AssetFilterMap) ToAssetFilterMapOutput() AssetFilterMapOutput {
	return i.ToAssetFilterMapOutputWithContext(context.Background())
}

func (i AssetFilterMap) ToAssetFilterMapOutputWithContext(ctx context.Context) AssetFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFilterMapOutput)
}

type AssetFilterOutput struct{ *pulumi.OutputState }

func (AssetFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetFilter)(nil))
}

func (o AssetFilterOutput) ToAssetFilterOutput() AssetFilterOutput {
	return o
}

func (o AssetFilterOutput) ToAssetFilterOutputWithContext(ctx context.Context) AssetFilterOutput {
	return o
}

func (o AssetFilterOutput) ToAssetFilterPtrOutput() AssetFilterPtrOutput {
	return o.ToAssetFilterPtrOutputWithContext(context.Background())
}

func (o AssetFilterOutput) ToAssetFilterPtrOutputWithContext(ctx context.Context) AssetFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetFilter) *AssetFilter {
		return &v
	}).(AssetFilterPtrOutput)
}

type AssetFilterPtrOutput struct{ *pulumi.OutputState }

func (AssetFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetFilter)(nil))
}

func (o AssetFilterPtrOutput) ToAssetFilterPtrOutput() AssetFilterPtrOutput {
	return o
}

func (o AssetFilterPtrOutput) ToAssetFilterPtrOutputWithContext(ctx context.Context) AssetFilterPtrOutput {
	return o
}

func (o AssetFilterPtrOutput) Elem() AssetFilterOutput {
	return o.ApplyT(func(v *AssetFilter) AssetFilter {
		if v != nil {
			return *v
		}
		var ret AssetFilter
		return ret
	}).(AssetFilterOutput)
}

type AssetFilterArrayOutput struct{ *pulumi.OutputState }

func (AssetFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssetFilter)(nil))
}

func (o AssetFilterArrayOutput) ToAssetFilterArrayOutput() AssetFilterArrayOutput {
	return o
}

func (o AssetFilterArrayOutput) ToAssetFilterArrayOutputWithContext(ctx context.Context) AssetFilterArrayOutput {
	return o
}

func (o AssetFilterArrayOutput) Index(i pulumi.IntInput) AssetFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssetFilter {
		return vs[0].([]AssetFilter)[vs[1].(int)]
	}).(AssetFilterOutput)
}

type AssetFilterMapOutput struct{ *pulumi.OutputState }

func (AssetFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssetFilter)(nil))
}

func (o AssetFilterMapOutput) ToAssetFilterMapOutput() AssetFilterMapOutput {
	return o
}

func (o AssetFilterMapOutput) ToAssetFilterMapOutputWithContext(ctx context.Context) AssetFilterMapOutput {
	return o
}

func (o AssetFilterMapOutput) MapIndex(k pulumi.StringInput) AssetFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AssetFilter {
		return vs[0].(map[string]AssetFilter)[vs[1].(string)]
	}).(AssetFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetFilterOutput{})
	pulumi.RegisterOutputType(AssetFilterPtrOutput{})
	pulumi.RegisterOutputType(AssetFilterArrayOutput{})
	pulumi.RegisterOutputType(AssetFilterMapOutput{})
}
