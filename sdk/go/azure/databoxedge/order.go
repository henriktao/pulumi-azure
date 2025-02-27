// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Databox Edge Order.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/databoxedge"
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
// 		exampleDevice, err := databoxedge.NewDevice(ctx, "exampleDevice", &databoxedge.DeviceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			SkuName:           pulumi.String("EdgeP_Base-Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = databoxedge.NewOrder(ctx, "exampleOrder", &databoxedge.OrderArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DeviceName:        exampleDevice.Name,
// 			Contact: &databoxedge.OrderContactArgs{
// 				CompanyName: pulumi.String("Contoso Corporation"),
// 				Name:        pulumi.String("Bart"),
// 				EmailLists: []string{
// 					"bart@example.com",
// 				},
// 				Phone: "(800) 867-5309",
// 			},
// 			ShipmentAddress: &databoxedge.OrderShipmentAddressArgs{
// 				Addresses: pulumi.StringArray{
// 					pulumi.String("740 Evergreen Terrace"),
// 				},
// 				City:       pulumi.String("Springfield"),
// 				Country:    pulumi.String("United States"),
// 				PostalCode: pulumi.String("97403"),
// 				State:      pulumi.String("OR"),
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
// Databox Edge Orders can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:databoxedge/order:Order example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1/orders/default
// ```
type Order struct {
	pulumi.CustomResourceState

	// A `contact` block as defined below.
	Contact OrderContactOutput `pulumi:"contact"`
	// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// The contact person name. Changing this forces a new Databox Edge Order to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Tracking information for the package returned from the customer whether it has an original or a replacement device. A `returnTracking` block as defined below.
	ReturnTrackings OrderReturnTrackingArrayOutput `pulumi:"returnTrackings"`
	// Serial number of the device being tracked.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// A `shipmentAddress block as defined below.
	ShipmentAddress OrderShipmentAddressOutput `pulumi:"shipmentAddress"`
	// List of status changes in the order. A `shipmentHistory` block as defined below.
	ShipmentHistories OrderShipmentHistoryArrayOutput `pulumi:"shipmentHistories"`
	// Tracking information for the package delivered to the customer whether it has an original or a replacement device. A `shipmentTracking` block as defined below.
	ShipmentTrackings OrderShipmentTrackingArrayOutput `pulumi:"shipmentTrackings"`
	// The current status of the order. A `status` block as defined below.
	Statuses OrderStatusArrayOutput `pulumi:"statuses"`
}

// NewOrder registers a new resource with the given unique name, arguments, and options.
func NewOrder(ctx *pulumi.Context,
	name string, args *OrderArgs, opts ...pulumi.ResourceOption) (*Order, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Contact == nil {
		return nil, errors.New("invalid value for required argument 'Contact'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShipmentAddress == nil {
		return nil, errors.New("invalid value for required argument 'ShipmentAddress'")
	}
	var resource Order
	err := ctx.RegisterResource("azure:databoxedge/order:Order", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrder gets an existing Order resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderState, opts ...pulumi.ResourceOption) (*Order, error) {
	var resource Order
	err := ctx.ReadResource("azure:databoxedge/order:Order", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Order resources.
type orderState struct {
	// A `contact` block as defined below.
	Contact *OrderContact `pulumi:"contact"`
	// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
	DeviceName *string `pulumi:"deviceName"`
	// The contact person name. Changing this forces a new Databox Edge Order to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Tracking information for the package returned from the customer whether it has an original or a replacement device. A `returnTracking` block as defined below.
	ReturnTrackings []OrderReturnTracking `pulumi:"returnTrackings"`
	// Serial number of the device being tracked.
	SerialNumber *string `pulumi:"serialNumber"`
	// A `shipmentAddress block as defined below.
	ShipmentAddress *OrderShipmentAddress `pulumi:"shipmentAddress"`
	// List of status changes in the order. A `shipmentHistory` block as defined below.
	ShipmentHistories []OrderShipmentHistory `pulumi:"shipmentHistories"`
	// Tracking information for the package delivered to the customer whether it has an original or a replacement device. A `shipmentTracking` block as defined below.
	ShipmentTrackings []OrderShipmentTracking `pulumi:"shipmentTrackings"`
	// The current status of the order. A `status` block as defined below.
	Statuses []OrderStatus `pulumi:"statuses"`
}

type OrderState struct {
	// A `contact` block as defined below.
	Contact OrderContactPtrInput
	// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
	DeviceName pulumi.StringPtrInput
	// The contact person name. Changing this forces a new Databox Edge Order to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Tracking information for the package returned from the customer whether it has an original or a replacement device. A `returnTracking` block as defined below.
	ReturnTrackings OrderReturnTrackingArrayInput
	// Serial number of the device being tracked.
	SerialNumber pulumi.StringPtrInput
	// A `shipmentAddress block as defined below.
	ShipmentAddress OrderShipmentAddressPtrInput
	// List of status changes in the order. A `shipmentHistory` block as defined below.
	ShipmentHistories OrderShipmentHistoryArrayInput
	// Tracking information for the package delivered to the customer whether it has an original or a replacement device. A `shipmentTracking` block as defined below.
	ShipmentTrackings OrderShipmentTrackingArrayInput
	// The current status of the order. A `status` block as defined below.
	Statuses OrderStatusArrayInput
}

func (OrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderState)(nil)).Elem()
}

type orderArgs struct {
	// A `contact` block as defined below.
	Contact OrderContact `pulumi:"contact"`
	// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
	DeviceName string `pulumi:"deviceName"`
	// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `shipmentAddress block as defined below.
	ShipmentAddress OrderShipmentAddress `pulumi:"shipmentAddress"`
}

// The set of arguments for constructing a Order resource.
type OrderArgs struct {
	// A `contact` block as defined below.
	Contact OrderContactInput
	// The name of the Databox Edge Device this order is for. Changing this forces a new Databox Edge Order to be created.
	DeviceName pulumi.StringInput
	// The name of the Resource Group where the Databox Edge Order should exist. Changing this forces a new Databox Edge Order to be created.
	ResourceGroupName pulumi.StringInput
	// A `shipmentAddress block as defined below.
	ShipmentAddress OrderShipmentAddressInput
}

func (OrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderArgs)(nil)).Elem()
}

type OrderInput interface {
	pulumi.Input

	ToOrderOutput() OrderOutput
	ToOrderOutputWithContext(ctx context.Context) OrderOutput
}

func (*Order) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (i *Order) ToOrderOutput() OrderOutput {
	return i.ToOrderOutputWithContext(context.Background())
}

func (i *Order) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderOutput)
}

// OrderArrayInput is an input type that accepts OrderArray and OrderArrayOutput values.
// You can construct a concrete instance of `OrderArrayInput` via:
//
//          OrderArray{ OrderArgs{...} }
type OrderArrayInput interface {
	pulumi.Input

	ToOrderArrayOutput() OrderArrayOutput
	ToOrderArrayOutputWithContext(context.Context) OrderArrayOutput
}

type OrderArray []OrderInput

func (OrderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Order)(nil)).Elem()
}

func (i OrderArray) ToOrderArrayOutput() OrderArrayOutput {
	return i.ToOrderArrayOutputWithContext(context.Background())
}

func (i OrderArray) ToOrderArrayOutputWithContext(ctx context.Context) OrderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderArrayOutput)
}

// OrderMapInput is an input type that accepts OrderMap and OrderMapOutput values.
// You can construct a concrete instance of `OrderMapInput` via:
//
//          OrderMap{ "key": OrderArgs{...} }
type OrderMapInput interface {
	pulumi.Input

	ToOrderMapOutput() OrderMapOutput
	ToOrderMapOutputWithContext(context.Context) OrderMapOutput
}

type OrderMap map[string]OrderInput

func (OrderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Order)(nil)).Elem()
}

func (i OrderMap) ToOrderMapOutput() OrderMapOutput {
	return i.ToOrderMapOutputWithContext(context.Background())
}

func (i OrderMap) ToOrderMapOutputWithContext(ctx context.Context) OrderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderMapOutput)
}

type OrderOutput struct{ *pulumi.OutputState }

func (OrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (o OrderOutput) ToOrderOutput() OrderOutput {
	return o
}

func (o OrderOutput) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return o
}

type OrderArrayOutput struct{ *pulumi.OutputState }

func (OrderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Order)(nil)).Elem()
}

func (o OrderArrayOutput) ToOrderArrayOutput() OrderArrayOutput {
	return o
}

func (o OrderArrayOutput) ToOrderArrayOutputWithContext(ctx context.Context) OrderArrayOutput {
	return o
}

func (o OrderArrayOutput) Index(i pulumi.IntInput) OrderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Order {
		return vs[0].([]*Order)[vs[1].(int)]
	}).(OrderOutput)
}

type OrderMapOutput struct{ *pulumi.OutputState }

func (OrderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Order)(nil)).Elem()
}

func (o OrderMapOutput) ToOrderMapOutput() OrderMapOutput {
	return o
}

func (o OrderMapOutput) ToOrderMapOutputWithContext(ctx context.Context) OrderMapOutput {
	return o
}

func (o OrderMapOutput) MapIndex(k pulumi.StringInput) OrderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Order {
		return vs[0].(map[string]*Order)[vs[1].(string)]
	}).(OrderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrderInput)(nil)).Elem(), &Order{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderArrayInput)(nil)).Elem(), OrderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderMapInput)(nil)).Elem(), OrderMap{})
	pulumi.RegisterOutputType(OrderOutput{})
	pulumi.RegisterOutputType(OrderArrayOutput{})
	pulumi.RegisterOutputType(OrderMapOutput{})
}
