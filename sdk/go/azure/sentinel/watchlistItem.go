// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Sentinel Watchlist Item.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sentinel"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAnalyticsSolution, err := operationalinsights.NewAnalyticsSolution(ctx, "exampleAnalyticsSolution", &operationalinsights.AnalyticsSolutionArgs{
// 			SolutionName:        pulumi.String("SecurityInsights"),
// 			Location:            exampleResourceGroup.Location,
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			WorkspaceResourceId: exampleAnalyticsWorkspace.ID(),
// 			WorkspaceName:       exampleAnalyticsWorkspace.Name,
// 			Plan: &operationalinsights.AnalyticsSolutionPlanArgs{
// 				Publisher: pulumi.String("Microsoft"),
// 				Product:   pulumi.String("OMSGallery/SecurityInsights"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWatchlist, err := sentinel.NewWatchlist(ctx, "exampleWatchlist", &sentinel.WatchlistArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsSolution.WorkspaceResourceId,
// 			DisplayName:             pulumi.String("example-wl"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sentinel.NewWatchlistItem(ctx, "exampleWatchlistItem", &sentinel.WatchlistItemArgs{
// 			WatchlistId: exampleWatchlist.ID(),
// 			Properties: pulumi.StringMap{
// 				"k1": pulumi.String("v1"),
// 				"k2": pulumi.String("v2"),
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
// Sentinel Watchlist Items can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/watchlistItem:WatchlistItem example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1/watchlistItems/item1
// ```
type WatchlistItem struct {
	pulumi.CustomResourceState

	// The name in UUID format which should be used for this Sentinel Watchlist Item. Changing this forces a new Sentinel Watchlist Item to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The key value pairs of the Sentinel Watchlist Item.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// The ID of the Sentinel Watchlist that this Item resides in. Changing this forces a new Sentinel Watchlist Item to be created.
	WatchlistId pulumi.StringOutput `pulumi:"watchlistId"`
}

// NewWatchlistItem registers a new resource with the given unique name, arguments, and options.
func NewWatchlistItem(ctx *pulumi.Context,
	name string, args *WatchlistItemArgs, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.WatchlistId == nil {
		return nil, errors.New("invalid value for required argument 'WatchlistId'")
	}
	var resource WatchlistItem
	err := ctx.RegisterResource("azure:sentinel/watchlistItem:WatchlistItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatchlistItem gets an existing WatchlistItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatchlistItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistItemState, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	var resource WatchlistItem
	err := ctx.ReadResource("azure:sentinel/watchlistItem:WatchlistItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WatchlistItem resources.
type watchlistItemState struct {
	// The name in UUID format which should be used for this Sentinel Watchlist Item. Changing this forces a new Sentinel Watchlist Item to be created.
	Name *string `pulumi:"name"`
	// The key value pairs of the Sentinel Watchlist Item.
	Properties map[string]string `pulumi:"properties"`
	// The ID of the Sentinel Watchlist that this Item resides in. Changing this forces a new Sentinel Watchlist Item to be created.
	WatchlistId *string `pulumi:"watchlistId"`
}

type WatchlistItemState struct {
	// The name in UUID format which should be used for this Sentinel Watchlist Item. Changing this forces a new Sentinel Watchlist Item to be created.
	Name pulumi.StringPtrInput
	// The key value pairs of the Sentinel Watchlist Item.
	Properties pulumi.StringMapInput
	// The ID of the Sentinel Watchlist that this Item resides in. Changing this forces a new Sentinel Watchlist Item to be created.
	WatchlistId pulumi.StringPtrInput
}

func (WatchlistItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemState)(nil)).Elem()
}

type watchlistItemArgs struct {
	// The name in UUID format which should be used for this Sentinel Watchlist Item. Changing this forces a new Sentinel Watchlist Item to be created.
	Name *string `pulumi:"name"`
	// The key value pairs of the Sentinel Watchlist Item.
	Properties map[string]string `pulumi:"properties"`
	// The ID of the Sentinel Watchlist that this Item resides in. Changing this forces a new Sentinel Watchlist Item to be created.
	WatchlistId string `pulumi:"watchlistId"`
}

// The set of arguments for constructing a WatchlistItem resource.
type WatchlistItemArgs struct {
	// The name in UUID format which should be used for this Sentinel Watchlist Item. Changing this forces a new Sentinel Watchlist Item to be created.
	Name pulumi.StringPtrInput
	// The key value pairs of the Sentinel Watchlist Item.
	Properties pulumi.StringMapInput
	// The ID of the Sentinel Watchlist that this Item resides in. Changing this forces a new Sentinel Watchlist Item to be created.
	WatchlistId pulumi.StringInput
}

func (WatchlistItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemArgs)(nil)).Elem()
}

type WatchlistItemInput interface {
	pulumi.Input

	ToWatchlistItemOutput() WatchlistItemOutput
	ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput
}

func (*WatchlistItem) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (i *WatchlistItem) ToWatchlistItemOutput() WatchlistItemOutput {
	return i.ToWatchlistItemOutputWithContext(context.Background())
}

func (i *WatchlistItem) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistItemOutput)
}

// WatchlistItemArrayInput is an input type that accepts WatchlistItemArray and WatchlistItemArrayOutput values.
// You can construct a concrete instance of `WatchlistItemArrayInput` via:
//
//          WatchlistItemArray{ WatchlistItemArgs{...} }
type WatchlistItemArrayInput interface {
	pulumi.Input

	ToWatchlistItemArrayOutput() WatchlistItemArrayOutput
	ToWatchlistItemArrayOutputWithContext(context.Context) WatchlistItemArrayOutput
}

type WatchlistItemArray []WatchlistItemInput

func (WatchlistItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WatchlistItem)(nil)).Elem()
}

func (i WatchlistItemArray) ToWatchlistItemArrayOutput() WatchlistItemArrayOutput {
	return i.ToWatchlistItemArrayOutputWithContext(context.Background())
}

func (i WatchlistItemArray) ToWatchlistItemArrayOutputWithContext(ctx context.Context) WatchlistItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistItemArrayOutput)
}

// WatchlistItemMapInput is an input type that accepts WatchlistItemMap and WatchlistItemMapOutput values.
// You can construct a concrete instance of `WatchlistItemMapInput` via:
//
//          WatchlistItemMap{ "key": WatchlistItemArgs{...} }
type WatchlistItemMapInput interface {
	pulumi.Input

	ToWatchlistItemMapOutput() WatchlistItemMapOutput
	ToWatchlistItemMapOutputWithContext(context.Context) WatchlistItemMapOutput
}

type WatchlistItemMap map[string]WatchlistItemInput

func (WatchlistItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WatchlistItem)(nil)).Elem()
}

func (i WatchlistItemMap) ToWatchlistItemMapOutput() WatchlistItemMapOutput {
	return i.ToWatchlistItemMapOutputWithContext(context.Background())
}

func (i WatchlistItemMap) ToWatchlistItemMapOutputWithContext(ctx context.Context) WatchlistItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistItemMapOutput)
}

type WatchlistItemOutput struct{ *pulumi.OutputState }

func (WatchlistItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (o WatchlistItemOutput) ToWatchlistItemOutput() WatchlistItemOutput {
	return o
}

func (o WatchlistItemOutput) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return o
}

type WatchlistItemArrayOutput struct{ *pulumi.OutputState }

func (WatchlistItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WatchlistItem)(nil)).Elem()
}

func (o WatchlistItemArrayOutput) ToWatchlistItemArrayOutput() WatchlistItemArrayOutput {
	return o
}

func (o WatchlistItemArrayOutput) ToWatchlistItemArrayOutputWithContext(ctx context.Context) WatchlistItemArrayOutput {
	return o
}

func (o WatchlistItemArrayOutput) Index(i pulumi.IntInput) WatchlistItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WatchlistItem {
		return vs[0].([]*WatchlistItem)[vs[1].(int)]
	}).(WatchlistItemOutput)
}

type WatchlistItemMapOutput struct{ *pulumi.OutputState }

func (WatchlistItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WatchlistItem)(nil)).Elem()
}

func (o WatchlistItemMapOutput) ToWatchlistItemMapOutput() WatchlistItemMapOutput {
	return o
}

func (o WatchlistItemMapOutput) ToWatchlistItemMapOutputWithContext(ctx context.Context) WatchlistItemMapOutput {
	return o
}

func (o WatchlistItemMapOutput) MapIndex(k pulumi.StringInput) WatchlistItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WatchlistItem {
		return vs[0].(map[string]*WatchlistItem)[vs[1].(string)]
	}).(WatchlistItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistItemInput)(nil)).Elem(), &WatchlistItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistItemArrayInput)(nil)).Elem(), WatchlistItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistItemMapInput)(nil)).Elem(), WatchlistItemMap{})
	pulumi.RegisterOutputType(WatchlistItemOutput{})
	pulumi.RegisterOutputType(WatchlistItemArrayOutput{})
	pulumi.RegisterOutputType(WatchlistItemMapOutput{})
}
