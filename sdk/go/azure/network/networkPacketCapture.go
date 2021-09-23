// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures Network Packet Capturing against a Virtual Machine using a Network Watcher.
//
// ## Import
//
// Packet Captures can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkPacketCapture:NetworkPacketCapture capture1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
// ```
type NetworkPacketCapture struct {
	pulumi.CustomResourceState

	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayOutput `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrOutput `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrOutput `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrOutput `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationOutput `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewNetworkPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewNetworkPacketCapture(ctx *pulumi.Context,
	name string, args *NetworkPacketCaptureArgs, opts ...pulumi.ResourceOption) (*NetworkPacketCapture, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageLocation == nil {
		return nil, errors.New("invalid value for required argument 'StorageLocation'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource NetworkPacketCapture
	err := ctx.RegisterResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPacketCapture gets an existing NetworkPacketCapture resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPacketCaptureState, opts ...pulumi.ResourceOption) (*NetworkPacketCapture, error) {
	var resource NetworkPacketCapture
	err := ctx.ReadResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPacketCapture resources.
type networkPacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []NetworkPacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation *NetworkPacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type NetworkPacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationPtrInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (NetworkPacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPacketCaptureState)(nil)).Elem()
}

type networkPacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []NetworkPacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a NetworkPacketCapture resource.
type NetworkPacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (NetworkPacketCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPacketCaptureArgs)(nil)).Elem()
}

type NetworkPacketCaptureInput interface {
	pulumi.Input

	ToNetworkPacketCaptureOutput() NetworkPacketCaptureOutput
	ToNetworkPacketCaptureOutputWithContext(ctx context.Context) NetworkPacketCaptureOutput
}

func (*NetworkPacketCapture) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPacketCapture)(nil))
}

func (i *NetworkPacketCapture) ToNetworkPacketCaptureOutput() NetworkPacketCaptureOutput {
	return i.ToNetworkPacketCaptureOutputWithContext(context.Background())
}

func (i *NetworkPacketCapture) ToNetworkPacketCaptureOutputWithContext(ctx context.Context) NetworkPacketCaptureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPacketCaptureOutput)
}

func (i *NetworkPacketCapture) ToNetworkPacketCapturePtrOutput() NetworkPacketCapturePtrOutput {
	return i.ToNetworkPacketCapturePtrOutputWithContext(context.Background())
}

func (i *NetworkPacketCapture) ToNetworkPacketCapturePtrOutputWithContext(ctx context.Context) NetworkPacketCapturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPacketCapturePtrOutput)
}

type NetworkPacketCapturePtrInput interface {
	pulumi.Input

	ToNetworkPacketCapturePtrOutput() NetworkPacketCapturePtrOutput
	ToNetworkPacketCapturePtrOutputWithContext(ctx context.Context) NetworkPacketCapturePtrOutput
}

type networkPacketCapturePtrType NetworkPacketCaptureArgs

func (*networkPacketCapturePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPacketCapture)(nil))
}

func (i *networkPacketCapturePtrType) ToNetworkPacketCapturePtrOutput() NetworkPacketCapturePtrOutput {
	return i.ToNetworkPacketCapturePtrOutputWithContext(context.Background())
}

func (i *networkPacketCapturePtrType) ToNetworkPacketCapturePtrOutputWithContext(ctx context.Context) NetworkPacketCapturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPacketCapturePtrOutput)
}

// NetworkPacketCaptureArrayInput is an input type that accepts NetworkPacketCaptureArray and NetworkPacketCaptureArrayOutput values.
// You can construct a concrete instance of `NetworkPacketCaptureArrayInput` via:
//
//          NetworkPacketCaptureArray{ NetworkPacketCaptureArgs{...} }
type NetworkPacketCaptureArrayInput interface {
	pulumi.Input

	ToNetworkPacketCaptureArrayOutput() NetworkPacketCaptureArrayOutput
	ToNetworkPacketCaptureArrayOutputWithContext(context.Context) NetworkPacketCaptureArrayOutput
}

type NetworkPacketCaptureArray []NetworkPacketCaptureInput

func (NetworkPacketCaptureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPacketCapture)(nil)).Elem()
}

func (i NetworkPacketCaptureArray) ToNetworkPacketCaptureArrayOutput() NetworkPacketCaptureArrayOutput {
	return i.ToNetworkPacketCaptureArrayOutputWithContext(context.Background())
}

func (i NetworkPacketCaptureArray) ToNetworkPacketCaptureArrayOutputWithContext(ctx context.Context) NetworkPacketCaptureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPacketCaptureArrayOutput)
}

// NetworkPacketCaptureMapInput is an input type that accepts NetworkPacketCaptureMap and NetworkPacketCaptureMapOutput values.
// You can construct a concrete instance of `NetworkPacketCaptureMapInput` via:
//
//          NetworkPacketCaptureMap{ "key": NetworkPacketCaptureArgs{...} }
type NetworkPacketCaptureMapInput interface {
	pulumi.Input

	ToNetworkPacketCaptureMapOutput() NetworkPacketCaptureMapOutput
	ToNetworkPacketCaptureMapOutputWithContext(context.Context) NetworkPacketCaptureMapOutput
}

type NetworkPacketCaptureMap map[string]NetworkPacketCaptureInput

func (NetworkPacketCaptureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPacketCapture)(nil)).Elem()
}

func (i NetworkPacketCaptureMap) ToNetworkPacketCaptureMapOutput() NetworkPacketCaptureMapOutput {
	return i.ToNetworkPacketCaptureMapOutputWithContext(context.Background())
}

func (i NetworkPacketCaptureMap) ToNetworkPacketCaptureMapOutputWithContext(ctx context.Context) NetworkPacketCaptureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPacketCaptureMapOutput)
}

type NetworkPacketCaptureOutput struct{ *pulumi.OutputState }

func (NetworkPacketCaptureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPacketCapture)(nil))
}

func (o NetworkPacketCaptureOutput) ToNetworkPacketCaptureOutput() NetworkPacketCaptureOutput {
	return o
}

func (o NetworkPacketCaptureOutput) ToNetworkPacketCaptureOutputWithContext(ctx context.Context) NetworkPacketCaptureOutput {
	return o
}

func (o NetworkPacketCaptureOutput) ToNetworkPacketCapturePtrOutput() NetworkPacketCapturePtrOutput {
	return o.ToNetworkPacketCapturePtrOutputWithContext(context.Background())
}

func (o NetworkPacketCaptureOutput) ToNetworkPacketCapturePtrOutputWithContext(ctx context.Context) NetworkPacketCapturePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPacketCapture) *NetworkPacketCapture {
		return &v
	}).(NetworkPacketCapturePtrOutput)
}

type NetworkPacketCapturePtrOutput struct{ *pulumi.OutputState }

func (NetworkPacketCapturePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPacketCapture)(nil))
}

func (o NetworkPacketCapturePtrOutput) ToNetworkPacketCapturePtrOutput() NetworkPacketCapturePtrOutput {
	return o
}

func (o NetworkPacketCapturePtrOutput) ToNetworkPacketCapturePtrOutputWithContext(ctx context.Context) NetworkPacketCapturePtrOutput {
	return o
}

func (o NetworkPacketCapturePtrOutput) Elem() NetworkPacketCaptureOutput {
	return o.ApplyT(func(v *NetworkPacketCapture) NetworkPacketCapture {
		if v != nil {
			return *v
		}
		var ret NetworkPacketCapture
		return ret
	}).(NetworkPacketCaptureOutput)
}

type NetworkPacketCaptureArrayOutput struct{ *pulumi.OutputState }

func (NetworkPacketCaptureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkPacketCapture)(nil))
}

func (o NetworkPacketCaptureArrayOutput) ToNetworkPacketCaptureArrayOutput() NetworkPacketCaptureArrayOutput {
	return o
}

func (o NetworkPacketCaptureArrayOutput) ToNetworkPacketCaptureArrayOutputWithContext(ctx context.Context) NetworkPacketCaptureArrayOutput {
	return o
}

func (o NetworkPacketCaptureArrayOutput) Index(i pulumi.IntInput) NetworkPacketCaptureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkPacketCapture {
		return vs[0].([]NetworkPacketCapture)[vs[1].(int)]
	}).(NetworkPacketCaptureOutput)
}

type NetworkPacketCaptureMapOutput struct{ *pulumi.OutputState }

func (NetworkPacketCaptureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkPacketCapture)(nil))
}

func (o NetworkPacketCaptureMapOutput) ToNetworkPacketCaptureMapOutput() NetworkPacketCaptureMapOutput {
	return o
}

func (o NetworkPacketCaptureMapOutput) ToNetworkPacketCaptureMapOutputWithContext(ctx context.Context) NetworkPacketCaptureMapOutput {
	return o
}

func (o NetworkPacketCaptureMapOutput) MapIndex(k pulumi.StringInput) NetworkPacketCaptureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkPacketCapture {
		return vs[0].(map[string]NetworkPacketCapture)[vs[1].(string)]
	}).(NetworkPacketCaptureOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkPacketCaptureOutput{})
	pulumi.RegisterOutputType(NetworkPacketCapturePtrOutput{})
	pulumi.RegisterOutputType(NetworkPacketCaptureArrayOutput{})
	pulumi.RegisterOutputType(NetworkPacketCaptureMapOutput{})
}
