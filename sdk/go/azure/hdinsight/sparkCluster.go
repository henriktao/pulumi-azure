// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a HDInsight Spark Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/hdinsight"
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
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hdinsight.NewSparkCluster(ctx, "exampleSparkCluster", &hdinsight.SparkClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("3.6"),
// 			Tier:              pulumi.String("Standard"),
// 			ComponentVersion: &hdinsight.SparkClusterComponentVersionArgs{
// 				Spark: pulumi.String("2.3"),
// 			},
// 			Gateway: &hdinsight.SparkClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("Password123!"),
// 			},
// 			StorageAccounts: hdinsight.SparkClusterStorageAccountArray{
// 				&hdinsight.SparkClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.SparkClusterRolesArgs{
// 				HeadNode: &hdinsight.SparkClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_A3"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.SparkClusterRolesWorkerNodeArgs{
// 					VmSize:              pulumi.String("Standard_A3"),
// 					Username:            pulumi.String("acctestusrvm"),
// 					Password:            pulumi.String("AccTestvdSC4daf986!"),
// 					TargetInstanceCount: pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.SparkClusterRolesZookeeperNodeArgs{
// 					VmSize:   pulumi.String("Medium"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
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
// HDInsight Spark Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:hdinsight/sparkCluster:SparkCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
// ```
type SparkCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion SparkClusterComponentVersionOutput `pulumi:"componentVersion"`
	// Whether encryption in transit is enabled for this Cluster. Changing this forces a new resource to be created.
	EncryptionInTransitEnabled pulumi.BoolOutput `pulumi:"encryptionInTransitEnabled"`
	// A `gateway` block as defined below.
	Gateway SparkClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores SparkClusterMetastoresPtrOutput `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor SparkClusterMonitorPtrOutput `pulumi:"monitor"`
	// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `network` block as defined below.
	Network SparkClusterNetworkPtrOutput `pulumi:"network"`
	// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles SparkClusterRolesOutput `pulumi:"roles"`
	// A `securityProfile` block as defined below.
	SecurityProfile SparkClusterSecurityProfilePtrOutput `pulumi:"securityProfile"`
	// The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 SparkClusterStorageAccountGen2PtrOutput `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts SparkClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Spark Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewSparkCluster registers a new resource with the given unique name, arguments, and options.
func NewSparkCluster(ctx *pulumi.Context,
	name string, args *SparkClusterArgs, opts ...pulumi.ResourceOption) (*SparkCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'ClusterVersion'")
	}
	if args.ComponentVersion == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersion'")
	}
	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	var resource SparkCluster
	err := ctx.RegisterResource("azure:hdinsight/sparkCluster:SparkCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSparkCluster gets an existing SparkCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSparkCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SparkClusterState, opts ...pulumi.ResourceOption) (*SparkCluster, error) {
	var resource SparkCluster
	err := ctx.ReadResource("azure:hdinsight/sparkCluster:SparkCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SparkCluster resources.
type sparkClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion *SparkClusterComponentVersion `pulumi:"componentVersion"`
	// Whether encryption in transit is enabled for this Cluster. Changing this forces a new resource to be created.
	EncryptionInTransitEnabled *bool `pulumi:"encryptionInTransitEnabled"`
	// A `gateway` block as defined below.
	Gateway *SparkClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *SparkClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *SparkClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `network` block as defined below.
	Network *SparkClusterNetwork `pulumi:"network"`
	// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *SparkClusterRoles `pulumi:"roles"`
	// A `securityProfile` block as defined below.
	SecurityProfile *SparkClusterSecurityProfile `pulumi:"securityProfile"`
	// The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *SparkClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []SparkClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Spark Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type SparkClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// A `componentVersion` block as defined below.
	ComponentVersion SparkClusterComponentVersionPtrInput
	// Whether encryption in transit is enabled for this Cluster. Changing this forces a new resource to be created.
	EncryptionInTransitEnabled pulumi.BoolPtrInput
	// A `gateway` block as defined below.
	Gateway SparkClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight Spark Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores SparkClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor SparkClusterMonitorPtrInput
	// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `network` block as defined below.
	Network SparkClusterNetworkPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles SparkClusterRolesPtrInput
	// A `securityProfile` block as defined below.
	SecurityProfile SparkClusterSecurityProfilePtrInput
	// The SSH Connectivity Endpoint for this HDInsight Spark Cluster.
	SshEndpoint pulumi.StringPtrInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 SparkClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts SparkClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Spark Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (SparkClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*sparkClusterState)(nil)).Elem()
}

type sparkClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion SparkClusterComponentVersion `pulumi:"componentVersion"`
	// Whether encryption in transit is enabled for this Cluster. Changing this forces a new resource to be created.
	EncryptionInTransitEnabled *bool `pulumi:"encryptionInTransitEnabled"`
	// A `gateway` block as defined below.
	Gateway SparkClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *SparkClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *SparkClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `network` block as defined below.
	Network *SparkClusterNetwork `pulumi:"network"`
	// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles SparkClusterRoles `pulumi:"roles"`
	// A `securityProfile` block as defined below.
	SecurityProfile *SparkClusterSecurityProfile `pulumi:"securityProfile"`
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 *SparkClusterStorageAccountGen2 `pulumi:"storageAccountGen2"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []SparkClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Spark Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a SparkCluster resource.
type SparkClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `componentVersion` block as defined below.
	ComponentVersion SparkClusterComponentVersionInput
	// Whether encryption in transit is enabled for this Cluster. Changing this forces a new resource to be created.
	EncryptionInTransitEnabled pulumi.BoolPtrInput
	// A `gateway` block as defined below.
	Gateway SparkClusterGatewayInput
	// Specifies the Azure Region which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores SparkClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor SparkClusterMonitorPtrInput
	// Specifies the name for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `network` block as defined below.
	Network SparkClusterNetworkPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Spark Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles SparkClusterRolesInput
	// A `securityProfile` block as defined below.
	SecurityProfile SparkClusterSecurityProfilePtrInput
	// A `storageAccountGen2` block as defined below.
	StorageAccountGen2 SparkClusterStorageAccountGen2PtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts SparkClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Spark Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Spark Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (SparkClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sparkClusterArgs)(nil)).Elem()
}

type SparkClusterInput interface {
	pulumi.Input

	ToSparkClusterOutput() SparkClusterOutput
	ToSparkClusterOutputWithContext(ctx context.Context) SparkClusterOutput
}

func (*SparkCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkCluster)(nil))
}

func (i *SparkCluster) ToSparkClusterOutput() SparkClusterOutput {
	return i.ToSparkClusterOutputWithContext(context.Background())
}

func (i *SparkCluster) ToSparkClusterOutputWithContext(ctx context.Context) SparkClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkClusterOutput)
}

func (i *SparkCluster) ToSparkClusterPtrOutput() SparkClusterPtrOutput {
	return i.ToSparkClusterPtrOutputWithContext(context.Background())
}

func (i *SparkCluster) ToSparkClusterPtrOutputWithContext(ctx context.Context) SparkClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkClusterPtrOutput)
}

type SparkClusterPtrInput interface {
	pulumi.Input

	ToSparkClusterPtrOutput() SparkClusterPtrOutput
	ToSparkClusterPtrOutputWithContext(ctx context.Context) SparkClusterPtrOutput
}

type sparkClusterPtrType SparkClusterArgs

func (*sparkClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkCluster)(nil))
}

func (i *sparkClusterPtrType) ToSparkClusterPtrOutput() SparkClusterPtrOutput {
	return i.ToSparkClusterPtrOutputWithContext(context.Background())
}

func (i *sparkClusterPtrType) ToSparkClusterPtrOutputWithContext(ctx context.Context) SparkClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkClusterPtrOutput)
}

// SparkClusterArrayInput is an input type that accepts SparkClusterArray and SparkClusterArrayOutput values.
// You can construct a concrete instance of `SparkClusterArrayInput` via:
//
//          SparkClusterArray{ SparkClusterArgs{...} }
type SparkClusterArrayInput interface {
	pulumi.Input

	ToSparkClusterArrayOutput() SparkClusterArrayOutput
	ToSparkClusterArrayOutputWithContext(context.Context) SparkClusterArrayOutput
}

type SparkClusterArray []SparkClusterInput

func (SparkClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SparkCluster)(nil))
}

func (i SparkClusterArray) ToSparkClusterArrayOutput() SparkClusterArrayOutput {
	return i.ToSparkClusterArrayOutputWithContext(context.Background())
}

func (i SparkClusterArray) ToSparkClusterArrayOutputWithContext(ctx context.Context) SparkClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkClusterArrayOutput)
}

// SparkClusterMapInput is an input type that accepts SparkClusterMap and SparkClusterMapOutput values.
// You can construct a concrete instance of `SparkClusterMapInput` via:
//
//          SparkClusterMap{ "key": SparkClusterArgs{...} }
type SparkClusterMapInput interface {
	pulumi.Input

	ToSparkClusterMapOutput() SparkClusterMapOutput
	ToSparkClusterMapOutputWithContext(context.Context) SparkClusterMapOutput
}

type SparkClusterMap map[string]SparkClusterInput

func (SparkClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SparkCluster)(nil))
}

func (i SparkClusterMap) ToSparkClusterMapOutput() SparkClusterMapOutput {
	return i.ToSparkClusterMapOutputWithContext(context.Background())
}

func (i SparkClusterMap) ToSparkClusterMapOutputWithContext(ctx context.Context) SparkClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkClusterMapOutput)
}

type SparkClusterOutput struct {
	*pulumi.OutputState
}

func (SparkClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkCluster)(nil))
}

func (o SparkClusterOutput) ToSparkClusterOutput() SparkClusterOutput {
	return o
}

func (o SparkClusterOutput) ToSparkClusterOutputWithContext(ctx context.Context) SparkClusterOutput {
	return o
}

func (o SparkClusterOutput) ToSparkClusterPtrOutput() SparkClusterPtrOutput {
	return o.ToSparkClusterPtrOutputWithContext(context.Background())
}

func (o SparkClusterOutput) ToSparkClusterPtrOutputWithContext(ctx context.Context) SparkClusterPtrOutput {
	return o.ApplyT(func(v SparkCluster) *SparkCluster {
		return &v
	}).(SparkClusterPtrOutput)
}

type SparkClusterPtrOutput struct {
	*pulumi.OutputState
}

func (SparkClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkCluster)(nil))
}

func (o SparkClusterPtrOutput) ToSparkClusterPtrOutput() SparkClusterPtrOutput {
	return o
}

func (o SparkClusterPtrOutput) ToSparkClusterPtrOutputWithContext(ctx context.Context) SparkClusterPtrOutput {
	return o
}

type SparkClusterArrayOutput struct{ *pulumi.OutputState }

func (SparkClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkCluster)(nil))
}

func (o SparkClusterArrayOutput) ToSparkClusterArrayOutput() SparkClusterArrayOutput {
	return o
}

func (o SparkClusterArrayOutput) ToSparkClusterArrayOutputWithContext(ctx context.Context) SparkClusterArrayOutput {
	return o
}

func (o SparkClusterArrayOutput) Index(i pulumi.IntInput) SparkClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SparkCluster {
		return vs[0].([]SparkCluster)[vs[1].(int)]
	}).(SparkClusterOutput)
}

type SparkClusterMapOutput struct{ *pulumi.OutputState }

func (SparkClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SparkCluster)(nil))
}

func (o SparkClusterMapOutput) ToSparkClusterMapOutput() SparkClusterMapOutput {
	return o
}

func (o SparkClusterMapOutput) ToSparkClusterMapOutputWithContext(ctx context.Context) SparkClusterMapOutput {
	return o
}

func (o SparkClusterMapOutput) MapIndex(k pulumi.StringInput) SparkClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SparkCluster {
		return vs[0].(map[string]SparkCluster)[vs[1].(string)]
	}).(SparkClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(SparkClusterOutput{})
	pulumi.RegisterOutputType(SparkClusterPtrOutput{})
	pulumi.RegisterOutputType(SparkClusterArrayOutput{})
	pulumi.RegisterOutputType(SparkClusterMapOutput{})
}
