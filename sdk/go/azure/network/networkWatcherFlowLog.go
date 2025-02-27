// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Network Watcher Flow Log.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testResourceGroup, err := core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "testNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          testResourceGroup.Location,
// 			ResourceGroupName: testResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testNetworkWatcher, err := network.NewNetworkWatcher(ctx, "testNetworkWatcher", &network.NetworkWatcherArgs{
// 			Location:          testResourceGroup.Location,
// 			ResourceGroupName: testResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testAccount, err := storage.NewAccount(ctx, "testAccount", &storage.AccountArgs{
// 			ResourceGroupName:      testResourceGroup.Name,
// 			Location:               testResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountKind:            pulumi.String("StorageV2"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			EnableHttpsTrafficOnly: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "testAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          testResourceGroup.Location,
// 			ResourceGroupName: testResourceGroup.Name,
// 			Sku:               pulumi.String("PerGB2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkWatcherFlowLog(ctx, "testNetworkWatcherFlowLog", &network.NetworkWatcherFlowLogArgs{
// 			NetworkWatcherName:     testNetworkWatcher.Name,
// 			ResourceGroupName:      testResourceGroup.Name,
// 			NetworkSecurityGroupId: testNetworkSecurityGroup.ID(),
// 			StorageAccountId:       testAccount.ID(),
// 			Enabled:                pulumi.Bool(true),
// 			RetentionPolicy: &network.NetworkWatcherFlowLogRetentionPolicyArgs{
// 				Enabled: pulumi.Bool(true),
// 				Days:    pulumi.Int(7),
// 			},
// 			TrafficAnalytics: &network.NetworkWatcherFlowLogTrafficAnalyticsArgs{
// 				Enabled:             pulumi.Bool(true),
// 				WorkspaceId:         testAnalyticsWorkspace.WorkspaceId,
// 				WorkspaceRegion:     testAnalyticsWorkspace.Location,
// 				WorkspaceResourceId: testAnalyticsWorkspace.ID(),
// 				IntervalInMinutes:   pulumi.Int(10),
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
// Network Watcher Flow Logs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog watcher1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/flowLogs/log1
// ```
type NetworkWatcherFlowLog struct {
	pulumi.CustomResourceState

	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the `location` of the Network Watcher.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Network Watcher Flow Log. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringOutput `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyOutput `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrOutput `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNetworkWatcherFlowLog registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, args *NetworkWatcherFlowLogArgs, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.NetworkSecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupId'")
	}
	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionPolicy == nil {
		return nil, errors.New("invalid value for required argument 'RetentionPolicy'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource NetworkWatcherFlowLog
	err := ctx.RegisterResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkWatcherFlowLog gets an existing NetworkWatcherFlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherFlowLogState, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	var resource NetworkWatcherFlowLog
	err := ctx.ReadResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkWatcherFlowLog resources.
type networkWatcherFlowLogState struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled *bool `pulumi:"enabled"`
	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the `location` of the Network Watcher.
	Location *string `pulumi:"location"`
	// The name of the Network Watcher Flow Log. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId *string `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy *NetworkWatcherFlowLogRetentionPolicy `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags map[string]string `pulumi:"tags"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics *NetworkWatcherFlowLogTrafficAnalytics `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version *int `pulumi:"version"`
}

type NetworkWatcherFlowLogState struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolPtrInput
	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the `location` of the Network Watcher.
	Location pulumi.StringPtrInput
	// The name of the Network Watcher Flow Log. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyPtrInput
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags pulumi.StringMapInput
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrInput
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntPtrInput
}

func (NetworkWatcherFlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogState)(nil)).Elem()
}

type networkWatcherFlowLogArgs struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled bool `pulumi:"enabled"`
	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the `location` of the Network Watcher.
	Location *string `pulumi:"location"`
	// The name of the Network Watcher Flow Log. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicy `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId string `pulumi:"storageAccountId"`
	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags map[string]string `pulumi:"tags"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics *NetworkWatcherFlowLogTrafficAnalytics `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a NetworkWatcherFlowLog resource.
type NetworkWatcherFlowLogArgs struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolInput
	// The location where the Network Watcher Flow Log resides. Changing this forces a new resource to be created. Defaults to the `location` of the Network Watcher.
	Location pulumi.StringPtrInput
	// The name of the Network Watcher Flow Log. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyInput
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringInput
	// A mapping of tags which should be assigned to the Network Watcher Flow Log.
	Tags pulumi.StringMapInput
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrInput
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntPtrInput
}

func (NetworkWatcherFlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogArgs)(nil)).Elem()
}

type NetworkWatcherFlowLogInput interface {
	pulumi.Input

	ToNetworkWatcherFlowLogOutput() NetworkWatcherFlowLogOutput
	ToNetworkWatcherFlowLogOutputWithContext(ctx context.Context) NetworkWatcherFlowLogOutput
}

func (*NetworkWatcherFlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcherFlowLog)(nil)).Elem()
}

func (i *NetworkWatcherFlowLog) ToNetworkWatcherFlowLogOutput() NetworkWatcherFlowLogOutput {
	return i.ToNetworkWatcherFlowLogOutputWithContext(context.Background())
}

func (i *NetworkWatcherFlowLog) ToNetworkWatcherFlowLogOutputWithContext(ctx context.Context) NetworkWatcherFlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherFlowLogOutput)
}

// NetworkWatcherFlowLogArrayInput is an input type that accepts NetworkWatcherFlowLogArray and NetworkWatcherFlowLogArrayOutput values.
// You can construct a concrete instance of `NetworkWatcherFlowLogArrayInput` via:
//
//          NetworkWatcherFlowLogArray{ NetworkWatcherFlowLogArgs{...} }
type NetworkWatcherFlowLogArrayInput interface {
	pulumi.Input

	ToNetworkWatcherFlowLogArrayOutput() NetworkWatcherFlowLogArrayOutput
	ToNetworkWatcherFlowLogArrayOutputWithContext(context.Context) NetworkWatcherFlowLogArrayOutput
}

type NetworkWatcherFlowLogArray []NetworkWatcherFlowLogInput

func (NetworkWatcherFlowLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkWatcherFlowLog)(nil)).Elem()
}

func (i NetworkWatcherFlowLogArray) ToNetworkWatcherFlowLogArrayOutput() NetworkWatcherFlowLogArrayOutput {
	return i.ToNetworkWatcherFlowLogArrayOutputWithContext(context.Background())
}

func (i NetworkWatcherFlowLogArray) ToNetworkWatcherFlowLogArrayOutputWithContext(ctx context.Context) NetworkWatcherFlowLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherFlowLogArrayOutput)
}

// NetworkWatcherFlowLogMapInput is an input type that accepts NetworkWatcherFlowLogMap and NetworkWatcherFlowLogMapOutput values.
// You can construct a concrete instance of `NetworkWatcherFlowLogMapInput` via:
//
//          NetworkWatcherFlowLogMap{ "key": NetworkWatcherFlowLogArgs{...} }
type NetworkWatcherFlowLogMapInput interface {
	pulumi.Input

	ToNetworkWatcherFlowLogMapOutput() NetworkWatcherFlowLogMapOutput
	ToNetworkWatcherFlowLogMapOutputWithContext(context.Context) NetworkWatcherFlowLogMapOutput
}

type NetworkWatcherFlowLogMap map[string]NetworkWatcherFlowLogInput

func (NetworkWatcherFlowLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkWatcherFlowLog)(nil)).Elem()
}

func (i NetworkWatcherFlowLogMap) ToNetworkWatcherFlowLogMapOutput() NetworkWatcherFlowLogMapOutput {
	return i.ToNetworkWatcherFlowLogMapOutputWithContext(context.Background())
}

func (i NetworkWatcherFlowLogMap) ToNetworkWatcherFlowLogMapOutputWithContext(ctx context.Context) NetworkWatcherFlowLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherFlowLogMapOutput)
}

type NetworkWatcherFlowLogOutput struct{ *pulumi.OutputState }

func (NetworkWatcherFlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcherFlowLog)(nil)).Elem()
}

func (o NetworkWatcherFlowLogOutput) ToNetworkWatcherFlowLogOutput() NetworkWatcherFlowLogOutput {
	return o
}

func (o NetworkWatcherFlowLogOutput) ToNetworkWatcherFlowLogOutputWithContext(ctx context.Context) NetworkWatcherFlowLogOutput {
	return o
}

type NetworkWatcherFlowLogArrayOutput struct{ *pulumi.OutputState }

func (NetworkWatcherFlowLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkWatcherFlowLog)(nil)).Elem()
}

func (o NetworkWatcherFlowLogArrayOutput) ToNetworkWatcherFlowLogArrayOutput() NetworkWatcherFlowLogArrayOutput {
	return o
}

func (o NetworkWatcherFlowLogArrayOutput) ToNetworkWatcherFlowLogArrayOutputWithContext(ctx context.Context) NetworkWatcherFlowLogArrayOutput {
	return o
}

func (o NetworkWatcherFlowLogArrayOutput) Index(i pulumi.IntInput) NetworkWatcherFlowLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkWatcherFlowLog {
		return vs[0].([]*NetworkWatcherFlowLog)[vs[1].(int)]
	}).(NetworkWatcherFlowLogOutput)
}

type NetworkWatcherFlowLogMapOutput struct{ *pulumi.OutputState }

func (NetworkWatcherFlowLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkWatcherFlowLog)(nil)).Elem()
}

func (o NetworkWatcherFlowLogMapOutput) ToNetworkWatcherFlowLogMapOutput() NetworkWatcherFlowLogMapOutput {
	return o
}

func (o NetworkWatcherFlowLogMapOutput) ToNetworkWatcherFlowLogMapOutputWithContext(ctx context.Context) NetworkWatcherFlowLogMapOutput {
	return o
}

func (o NetworkWatcherFlowLogMapOutput) MapIndex(k pulumi.StringInput) NetworkWatcherFlowLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkWatcherFlowLog {
		return vs[0].(map[string]*NetworkWatcherFlowLog)[vs[1].(string)]
	}).(NetworkWatcherFlowLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkWatcherFlowLogInput)(nil)).Elem(), &NetworkWatcherFlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkWatcherFlowLogArrayInput)(nil)).Elem(), NetworkWatcherFlowLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkWatcherFlowLogMapInput)(nil)).Elem(), NetworkWatcherFlowLogMap{})
	pulumi.RegisterOutputType(NetworkWatcherFlowLogOutput{})
	pulumi.RegisterOutputType(NetworkWatcherFlowLogArrayOutput{})
	pulumi.RegisterOutputType(NetworkWatcherFlowLogMapOutput{})
}
