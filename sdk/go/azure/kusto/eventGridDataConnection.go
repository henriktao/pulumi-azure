// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Event Grid Data Connection
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/eventhub"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/kusto"
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
// 		exampleCluster, err := kusto.NewCluster(ctx, "exampleCluster", &kusto.ClusterArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDatabase, err := kusto.NewDatabase(ctx, "exampleDatabase", &kusto.DatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterName:       exampleCluster.Name,
// 			HotCachePeriod:    pulumi.String("P7D"),
// 			SoftDeletePeriod:  pulumi.String("P31D"),
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
// 		_, err = eventhub.NewEventHubNamespace(ctx, "testEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventhub.NewEventHub(ctx, "testEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     pulumi.Any(azurerm_eventhub_namespace.Example.Name),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(1),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConsumerGroup, err := eventhub.NewConsumerGroup(ctx, "exampleConsumerGroup", &eventhub.ConsumerGroupArgs{
// 			NamespaceName:     pulumi.Any(azurerm_eventhub_namespace.Example.Name),
// 			EventhubName:      pulumi.Any(azurerm_eventhub.Example.Name),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventSubscription, err := eventgrid.NewEventSubscription(ctx, "exampleEventSubscription", &eventgrid.EventSubscriptionArgs{
// 			Scope:               exampleAccount.ID(),
// 			EventhubEndpointId:  pulumi.Any(azurerm_eventhub.Example.Id),
// 			EventDeliverySchema: pulumi.String("EventGridSchema"),
// 			IncludedEventTypes: pulumi.StringArray{
// 				pulumi.String("Microsoft.Storage.BlobCreated"),
// 				pulumi.String("Microsoft.Storage.BlobRenamed"),
// 			},
// 			RetryPolicy: &eventgrid.EventSubscriptionRetryPolicyArgs{
// 				EventTimeToLive:     pulumi.Int(144),
// 				MaxDeliveryAttempts: pulumi.Int(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewEventGridDataConnection(ctx, "exampleEventGridDataConnection", &kusto.EventGridDataConnectionArgs{
// 			ResourceGroupName:         exampleResourceGroup.Name,
// 			Location:                  exampleResourceGroup.Location,
// 			ClusterName:               exampleCluster.Name,
// 			DatabaseName:              exampleDatabase.Name,
// 			StorageAccountId:          exampleAccount.ID(),
// 			EventhubId:                pulumi.Any(azurerm_eventhub.Example.Id),
// 			EventhubConsumerGroupName: exampleConsumerGroup.Name,
// 			TableName:                 pulumi.String("my-table"),
// 			MappingRuleName:           pulumi.String("my-table-mapping"),
// 			DataFormat:                pulumi.String("JSON"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleEventSubscription,
// 		}))
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
// Kusto Event Grid Data Connections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/eventGridDataConnection:EventGridDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
// ```
type EventGridDataConnection struct {
	pulumi.CustomResourceState

	// Specifies the blob storage event type that needs to be processed. Possible
	// Values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobRenamed`. Defaults
	// to `Microsoft.Storage.BlobCreated`.
	BlobStorageEventType pulumi.StringPtrOutput `pulumi:"blobStorageEventType"`
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Specifies the Event Hub consumer group this data connection will use for
	// ingestion. Changing this forces a new resource to be created.
	EventhubConsumerGroupName pulumi.StringOutput `pulumi:"eventhubConsumerGroupName"`
	// Specifies the resource id of the Event Hub this data connection will use for ingestion.
	// Changing this forces a new resource to be created.
	EventhubId pulumi.StringOutput `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the Kusto Event Grid Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// is the first record of every file ignored? Defaults to `false`.
	SkipFirstRecord pulumi.BoolPtrOutput `pulumi:"skipFirstRecord"`
	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewEventGridDataConnection registers a new resource with the given unique name, arguments, and options.
func NewEventGridDataConnection(ctx *pulumi.Context,
	name string, args *EventGridDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventhubConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'EventhubConsumerGroupName'")
	}
	if args.EventhubId == nil {
		return nil, errors.New("invalid value for required argument 'EventhubId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	var resource EventGridDataConnection
	err := ctx.RegisterResource("azure:kusto/eventGridDataConnection:EventGridDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventGridDataConnection gets an existing EventGridDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventGridDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridDataConnectionState, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	var resource EventGridDataConnection
	err := ctx.ReadResource("azure:kusto/eventGridDataConnection:EventGridDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventGridDataConnection resources.
type eventGridDataConnectionState struct {
	// Specifies the blob storage event type that needs to be processed. Possible
	// Values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobRenamed`. Defaults
	// to `Microsoft.Storage.BlobCreated`.
	BlobStorageEventType *string `pulumi:"blobStorageEventType"`
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies the Event Hub consumer group this data connection will use for
	// ingestion. Changing this forces a new resource to be created.
	EventhubConsumerGroupName *string `pulumi:"eventhubConsumerGroupName"`
	// Specifies the resource id of the Event Hub this data connection will use for ingestion.
	// Changing this forces a new resource to be created.
	EventhubId *string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto Event Grid Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// is the first record of every file ignored? Defaults to `false`.
	SkipFirstRecord *bool `pulumi:"skipFirstRecord"`
	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

type EventGridDataConnectionState struct {
	// Specifies the blob storage event type that needs to be processed. Possible
	// Values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobRenamed`. Defaults
	// to `Microsoft.Storage.BlobCreated`.
	BlobStorageEventType pulumi.StringPtrInput
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// Specifies the Event Hub consumer group this data connection will use for
	// ingestion. Changing this forces a new resource to be created.
	EventhubConsumerGroupName pulumi.StringPtrInput
	// Specifies the resource id of the Event Hub this data connection will use for ingestion.
	// Changing this forces a new resource to be created.
	EventhubId pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto Event Grid Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// is the first record of every file ignored? Defaults to `false`.
	SkipFirstRecord pulumi.BoolPtrInput
	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventGridDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionState)(nil)).Elem()
}

type eventGridDataConnectionArgs struct {
	// Specifies the blob storage event type that needs to be processed. Possible
	// Values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobRenamed`. Defaults
	// to `Microsoft.Storage.BlobCreated`.
	BlobStorageEventType *string `pulumi:"blobStorageEventType"`
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies the Event Hub consumer group this data connection will use for
	// ingestion. Changing this forces a new resource to be created.
	EventhubConsumerGroupName string `pulumi:"eventhubConsumerGroupName"`
	// Specifies the resource id of the Event Hub this data connection will use for ingestion.
	// Changing this forces a new resource to be created.
	EventhubId string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto Event Grid Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// is the first record of every file ignored? Defaults to `false`.
	SkipFirstRecord *bool `pulumi:"skipFirstRecord"`
	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a EventGridDataConnection resource.
type EventGridDataConnectionArgs struct {
	// Specifies the blob storage event type that needs to be processed. Possible
	// Values are `Microsoft.Storage.BlobCreated` and `Microsoft.Storage.BlobRenamed`. Defaults
	// to `Microsoft.Storage.BlobCreated`.
	BlobStorageEventType pulumi.StringPtrInput
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// Specifies the Event Hub consumer group this data connection will use for
	// ingestion. Changing this forces a new resource to be created.
	EventhubConsumerGroupName pulumi.StringInput
	// Specifies the resource id of the Event Hub this data connection will use for ingestion.
	// Changing this forces a new resource to be created.
	EventhubId pulumi.StringInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto Event Grid Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// is the first record of every file ignored? Defaults to `false`.
	SkipFirstRecord pulumi.BoolPtrInput
	// Specifies the resource id of the Storage Account this data connection will use for ingestion. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventGridDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionArgs)(nil)).Elem()
}

type EventGridDataConnectionInput interface {
	pulumi.Input

	ToEventGridDataConnectionOutput() EventGridDataConnectionOutput
	ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput
}

func (*EventGridDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataConnection)(nil))
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return i.ToEventGridDataConnectionOutputWithContext(context.Background())
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionOutput)
}

func (i *EventGridDataConnection) ToEventGridDataConnectionPtrOutput() EventGridDataConnectionPtrOutput {
	return i.ToEventGridDataConnectionPtrOutputWithContext(context.Background())
}

func (i *EventGridDataConnection) ToEventGridDataConnectionPtrOutputWithContext(ctx context.Context) EventGridDataConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionPtrOutput)
}

type EventGridDataConnectionPtrInput interface {
	pulumi.Input

	ToEventGridDataConnectionPtrOutput() EventGridDataConnectionPtrOutput
	ToEventGridDataConnectionPtrOutputWithContext(ctx context.Context) EventGridDataConnectionPtrOutput
}

type eventGridDataConnectionPtrType EventGridDataConnectionArgs

func (*eventGridDataConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil))
}

func (i *eventGridDataConnectionPtrType) ToEventGridDataConnectionPtrOutput() EventGridDataConnectionPtrOutput {
	return i.ToEventGridDataConnectionPtrOutputWithContext(context.Background())
}

func (i *eventGridDataConnectionPtrType) ToEventGridDataConnectionPtrOutputWithContext(ctx context.Context) EventGridDataConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionPtrOutput)
}

// EventGridDataConnectionArrayInput is an input type that accepts EventGridDataConnectionArray and EventGridDataConnectionArrayOutput values.
// You can construct a concrete instance of `EventGridDataConnectionArrayInput` via:
//
//          EventGridDataConnectionArray{ EventGridDataConnectionArgs{...} }
type EventGridDataConnectionArrayInput interface {
	pulumi.Input

	ToEventGridDataConnectionArrayOutput() EventGridDataConnectionArrayOutput
	ToEventGridDataConnectionArrayOutputWithContext(context.Context) EventGridDataConnectionArrayOutput
}

type EventGridDataConnectionArray []EventGridDataConnectionInput

func (EventGridDataConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventGridDataConnection)(nil)).Elem()
}

func (i EventGridDataConnectionArray) ToEventGridDataConnectionArrayOutput() EventGridDataConnectionArrayOutput {
	return i.ToEventGridDataConnectionArrayOutputWithContext(context.Background())
}

func (i EventGridDataConnectionArray) ToEventGridDataConnectionArrayOutputWithContext(ctx context.Context) EventGridDataConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionArrayOutput)
}

// EventGridDataConnectionMapInput is an input type that accepts EventGridDataConnectionMap and EventGridDataConnectionMapOutput values.
// You can construct a concrete instance of `EventGridDataConnectionMapInput` via:
//
//          EventGridDataConnectionMap{ "key": EventGridDataConnectionArgs{...} }
type EventGridDataConnectionMapInput interface {
	pulumi.Input

	ToEventGridDataConnectionMapOutput() EventGridDataConnectionMapOutput
	ToEventGridDataConnectionMapOutputWithContext(context.Context) EventGridDataConnectionMapOutput
}

type EventGridDataConnectionMap map[string]EventGridDataConnectionInput

func (EventGridDataConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventGridDataConnection)(nil)).Elem()
}

func (i EventGridDataConnectionMap) ToEventGridDataConnectionMapOutput() EventGridDataConnectionMapOutput {
	return i.ToEventGridDataConnectionMapOutputWithContext(context.Background())
}

func (i EventGridDataConnectionMap) ToEventGridDataConnectionMapOutputWithContext(ctx context.Context) EventGridDataConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionMapOutput)
}

type EventGridDataConnectionOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataConnection)(nil))
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionPtrOutput() EventGridDataConnectionPtrOutput {
	return o.ToEventGridDataConnectionPtrOutputWithContext(context.Background())
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionPtrOutputWithContext(ctx context.Context) EventGridDataConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventGridDataConnection) *EventGridDataConnection {
		return &v
	}).(EventGridDataConnectionPtrOutput)
}

type EventGridDataConnectionPtrOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil))
}

func (o EventGridDataConnectionPtrOutput) ToEventGridDataConnectionPtrOutput() EventGridDataConnectionPtrOutput {
	return o
}

func (o EventGridDataConnectionPtrOutput) ToEventGridDataConnectionPtrOutputWithContext(ctx context.Context) EventGridDataConnectionPtrOutput {
	return o
}

func (o EventGridDataConnectionPtrOutput) Elem() EventGridDataConnectionOutput {
	return o.ApplyT(func(v *EventGridDataConnection) EventGridDataConnection {
		if v != nil {
			return *v
		}
		var ret EventGridDataConnection
		return ret
	}).(EventGridDataConnectionOutput)
}

type EventGridDataConnectionArrayOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventGridDataConnection)(nil))
}

func (o EventGridDataConnectionArrayOutput) ToEventGridDataConnectionArrayOutput() EventGridDataConnectionArrayOutput {
	return o
}

func (o EventGridDataConnectionArrayOutput) ToEventGridDataConnectionArrayOutputWithContext(ctx context.Context) EventGridDataConnectionArrayOutput {
	return o
}

func (o EventGridDataConnectionArrayOutput) Index(i pulumi.IntInput) EventGridDataConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventGridDataConnection {
		return vs[0].([]EventGridDataConnection)[vs[1].(int)]
	}).(EventGridDataConnectionOutput)
}

type EventGridDataConnectionMapOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventGridDataConnection)(nil))
}

func (o EventGridDataConnectionMapOutput) ToEventGridDataConnectionMapOutput() EventGridDataConnectionMapOutput {
	return o
}

func (o EventGridDataConnectionMapOutput) ToEventGridDataConnectionMapOutputWithContext(ctx context.Context) EventGridDataConnectionMapOutput {
	return o
}

func (o EventGridDataConnectionMapOutput) MapIndex(k pulumi.StringInput) EventGridDataConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventGridDataConnection {
		return vs[0].(map[string]EventGridDataConnection)[vs[1].(string)]
	}).(EventGridDataConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(EventGridDataConnectionOutput{})
	pulumi.RegisterOutputType(EventGridDataConnectionPtrOutput{})
	pulumi.RegisterOutputType(EventGridDataConnectionArrayOutput{})
	pulumi.RegisterOutputType(EventGridDataConnectionMapOutput{})
}
