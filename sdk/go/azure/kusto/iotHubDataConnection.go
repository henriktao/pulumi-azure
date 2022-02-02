// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) IotHub Data Connection
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/iot"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/kusto"
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
// 		exampleIoTHub, err := iot.NewIoTHub(ctx, "exampleIoTHub", &iot.IoTHubArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &iot.IoTHubSkuArgs{
// 				Name:     pulumi.String("B1"),
// 				Capacity: pulumi.Int(1),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSharedAccessPolicy, err := iot.NewSharedAccessPolicy(ctx, "exampleSharedAccessPolicy", &iot.SharedAccessPolicyArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IothubName:        exampleIoTHub.Name,
// 			RegistryRead:      pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleConsumerGroup, err := iot.NewConsumerGroup(ctx, "exampleConsumerGroup", &iot.ConsumerGroupArgs{
// 			ResourceGroupName:    exampleResourceGroup.Name,
// 			IothubName:           exampleIoTHub.Name,
// 			EventhubEndpointName: pulumi.String("events"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewIotHubDataConnection(ctx, "exampleIotHubDataConnection", &kusto.IotHubDataConnectionArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			ClusterName:            exampleCluster.Name,
// 			DatabaseName:           exampleDatabase.Name,
// 			IothubId:               exampleIoTHub.ID(),
// 			ConsumerGroup:          exampleConsumerGroup.Name,
// 			SharedAccessPolicyName: exampleSharedAccessPolicy.Name,
// 			EventSystemProperties: pulumi.StringArray{
// 				pulumi.String("message-id"),
// 				pulumi.String("sequence-number"),
// 				pulumi.String("to"),
// 			},
// 			TableName:       pulumi.String("my-table"),
// 			MappingRuleName: pulumi.String("my-table-mapping"),
// 			DataFormat:      pulumi.String("JSON"),
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
// Kusto IotHub Data Connections can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/iotHubDataConnection:IotHubDataConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
// ```
type IotHubDataConnection struct {
	pulumi.CustomResourceState

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// Specifies the data format of the IoTHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `ORC`, `PARQUET`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV`, `TSVE`, `TXT` and `W3CLOGFILE`.
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
	EventSystemProperties pulumi.StringArrayOutput `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	IothubId pulumi.StringOutput `pulumi:"iothubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	SharedAccessPolicyName pulumi.StringOutput `pulumi:"sharedAccessPolicyName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewIotHubDataConnection registers a new resource with the given unique name, arguments, and options.
func NewIotHubDataConnection(ctx *pulumi.Context,
	name string, args *IotHubDataConnectionArgs, opts ...pulumi.ResourceOption) (*IotHubDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.IothubId == nil {
		return nil, errors.New("invalid value for required argument 'IothubId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SharedAccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyName'")
	}
	var resource IotHubDataConnection
	err := ctx.RegisterResource("azure:kusto/iotHubDataConnection:IotHubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubDataConnection gets an existing IotHubDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubDataConnectionState, opts ...pulumi.ResourceOption) (*IotHubDataConnection, error) {
	var resource IotHubDataConnection
	err := ctx.ReadResource("azure:kusto/iotHubDataConnection:IotHubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubDataConnection resources.
type iotHubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup *string `pulumi:"consumerGroup"`
	// Specifies the data format of the IoTHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `ORC`, `PARQUET`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV`, `TSVE`, `TXT` and `W3CLOGFILE`.
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	IothubId *string `pulumi:"iothubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

type IotHubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringPtrInput
	// Specifies the data format of the IoTHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `ORC`, `PARQUET`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV`, `TSVE`, `TXT` and `W3CLOGFILE`.
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
	EventSystemProperties pulumi.StringArrayInput
	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	IothubId pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	SharedAccessPolicyName pulumi.StringPtrInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (IotHubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDataConnectionState)(nil)).Elem()
}

type iotHubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// Specifies the data format of the IoTHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `ORC`, `PARQUET`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV`, `TSVE`, `TXT` and `W3CLOGFILE`.
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	IothubId string `pulumi:"iothubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	SharedAccessPolicyName string `pulumi:"sharedAccessPolicyName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a IotHubDataConnection resource.
type IotHubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringInput
	// Specifies the data format of the IoTHub messages. Allowed values: `APACHEAVRO`, `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `ORC`, `PARQUET`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV`, `TSVE`, `TXT` and `W3CLOGFILE`.
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created.
	EventSystemProperties pulumi.StringArrayInput
	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	IothubId pulumi.StringInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	SharedAccessPolicyName pulumi.StringInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (IotHubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDataConnectionArgs)(nil)).Elem()
}

type IotHubDataConnectionInput interface {
	pulumi.Input

	ToIotHubDataConnectionOutput() IotHubDataConnectionOutput
	ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput
}

func (*IotHubDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubDataConnection)(nil)).Elem()
}

func (i *IotHubDataConnection) ToIotHubDataConnectionOutput() IotHubDataConnectionOutput {
	return i.ToIotHubDataConnectionOutputWithContext(context.Background())
}

func (i *IotHubDataConnection) ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDataConnectionOutput)
}

// IotHubDataConnectionArrayInput is an input type that accepts IotHubDataConnectionArray and IotHubDataConnectionArrayOutput values.
// You can construct a concrete instance of `IotHubDataConnectionArrayInput` via:
//
//          IotHubDataConnectionArray{ IotHubDataConnectionArgs{...} }
type IotHubDataConnectionArrayInput interface {
	pulumi.Input

	ToIotHubDataConnectionArrayOutput() IotHubDataConnectionArrayOutput
	ToIotHubDataConnectionArrayOutputWithContext(context.Context) IotHubDataConnectionArrayOutput
}

type IotHubDataConnectionArray []IotHubDataConnectionInput

func (IotHubDataConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHubDataConnection)(nil)).Elem()
}

func (i IotHubDataConnectionArray) ToIotHubDataConnectionArrayOutput() IotHubDataConnectionArrayOutput {
	return i.ToIotHubDataConnectionArrayOutputWithContext(context.Background())
}

func (i IotHubDataConnectionArray) ToIotHubDataConnectionArrayOutputWithContext(ctx context.Context) IotHubDataConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDataConnectionArrayOutput)
}

// IotHubDataConnectionMapInput is an input type that accepts IotHubDataConnectionMap and IotHubDataConnectionMapOutput values.
// You can construct a concrete instance of `IotHubDataConnectionMapInput` via:
//
//          IotHubDataConnectionMap{ "key": IotHubDataConnectionArgs{...} }
type IotHubDataConnectionMapInput interface {
	pulumi.Input

	ToIotHubDataConnectionMapOutput() IotHubDataConnectionMapOutput
	ToIotHubDataConnectionMapOutputWithContext(context.Context) IotHubDataConnectionMapOutput
}

type IotHubDataConnectionMap map[string]IotHubDataConnectionInput

func (IotHubDataConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHubDataConnection)(nil)).Elem()
}

func (i IotHubDataConnectionMap) ToIotHubDataConnectionMapOutput() IotHubDataConnectionMapOutput {
	return i.ToIotHubDataConnectionMapOutputWithContext(context.Background())
}

func (i IotHubDataConnectionMap) ToIotHubDataConnectionMapOutputWithContext(ctx context.Context) IotHubDataConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDataConnectionMapOutput)
}

type IotHubDataConnectionOutput struct{ *pulumi.OutputState }

func (IotHubDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubDataConnection)(nil)).Elem()
}

func (o IotHubDataConnectionOutput) ToIotHubDataConnectionOutput() IotHubDataConnectionOutput {
	return o
}

func (o IotHubDataConnectionOutput) ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput {
	return o
}

type IotHubDataConnectionArrayOutput struct{ *pulumi.OutputState }

func (IotHubDataConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotHubDataConnection)(nil)).Elem()
}

func (o IotHubDataConnectionArrayOutput) ToIotHubDataConnectionArrayOutput() IotHubDataConnectionArrayOutput {
	return o
}

func (o IotHubDataConnectionArrayOutput) ToIotHubDataConnectionArrayOutputWithContext(ctx context.Context) IotHubDataConnectionArrayOutput {
	return o
}

func (o IotHubDataConnectionArrayOutput) Index(i pulumi.IntInput) IotHubDataConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IotHubDataConnection {
		return vs[0].([]*IotHubDataConnection)[vs[1].(int)]
	}).(IotHubDataConnectionOutput)
}

type IotHubDataConnectionMapOutput struct{ *pulumi.OutputState }

func (IotHubDataConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotHubDataConnection)(nil)).Elem()
}

func (o IotHubDataConnectionMapOutput) ToIotHubDataConnectionMapOutput() IotHubDataConnectionMapOutput {
	return o
}

func (o IotHubDataConnectionMapOutput) ToIotHubDataConnectionMapOutputWithContext(ctx context.Context) IotHubDataConnectionMapOutput {
	return o
}

func (o IotHubDataConnectionMapOutput) MapIndex(k pulumi.StringInput) IotHubDataConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IotHubDataConnection {
		return vs[0].(map[string]*IotHubDataConnection)[vs[1].(string)]
	}).(IotHubDataConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubDataConnectionInput)(nil)).Elem(), &IotHubDataConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubDataConnectionArrayInput)(nil)).Elem(), IotHubDataConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubDataConnectionMapInput)(nil)).Elem(), IotHubDataConnectionMap{})
	pulumi.RegisterOutputType(IotHubDataConnectionOutput{})
	pulumi.RegisterOutputType(IotHubDataConnectionArrayOutput{})
	pulumi.RegisterOutputType(IotHubDataConnectionMapOutput{})
}
