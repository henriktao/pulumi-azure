// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Stream Analytics Reference Input from MS SQL. Reference data (also known as a lookup table) is a finite data set that is static or slowly changing in nature, used to perform a lookup or to correlate with your data stream. Learn more [here](https://docs.microsoft.com/en-us/azure/stream-analytics/stream-analytics-use-reference-data#azure-sql-database).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleServer, err := mssql.NewServer(ctx, "exampleServer", &mssql.ServerArgs{
// 			ResourceGroupName:          pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:                   pulumi.Any(azurerm_resource_group.Example.Location),
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("admin"),
// 			AdministratorLoginPassword: pulumi.String("password"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDatabase, err := mssql.NewDatabase(ctx, "exampleDatabase", &mssql.DatabaseArgs{
// 			ServerId: exampleServer.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewReferenceInputMssql(ctx, "exampleReferenceInputMssql", &streamanalytics.ReferenceInputMssqlArgs{
// 			ResourceGroupName:       pulumi.Any(azurerm_stream_analytics_job.Example.Resource_group_name),
// 			StreamAnalyticsJobName:  pulumi.Any(azurerm_stream_analytics_job.Example.Name),
// 			Server:                  exampleServer.FullyQualifiedDomainName,
// 			Database:                exampleDatabase.Name,
// 			Username:                pulumi.String("exampleuser"),
// 			Password:                pulumi.String("examplepassword"),
// 			RefreshType:             pulumi.String("RefreshPeriodicallyWithFull"),
// 			RefreshIntervalDuration: pulumi.String("00:20:00"),
// 			FullSnapshotQuery:       pulumi.String(fmt.Sprintf("%v%v%v", "    SELECT *\n", "    INTO [YourOutputAlias]\n", "    FROM [YourInputAlias]\n")),
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
// Stream Analytics can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/referenceInputMssql:ReferenceInputMssql example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
// ```
type ReferenceInputMssql struct {
	pulumi.CustomResourceState

	// The MS SQL database name where the reference data exists.
	Database pulumi.StringOutput `pulumi:"database"`
	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when `refreshType` is `Static`.
	DeltaSnapshotQuery pulumi.StringPtrOutput `pulumi:"deltaSnapshotQuery"`
	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery pulumi.StringOutput `pulumi:"fullSnapshotQuery"`
	// The name of the Reference Input MS SQL data. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The username to connect to the MS SQL database.
	Password pulumi.StringOutput `pulumi:"password"`
	// The frequency in `hh:mm:ss` with which the reference data should be retrieved from the MS SQL database e.g. `00:20:00` for every 20 minutes. Must be set when `refreshType` is `RefreshPeriodicallyWithFull` or `RefreshPeriodicallyWithDelta`.
	RefreshIntervalDuration pulumi.StringPtrOutput `pulumi:"refreshIntervalDuration"`
	// Defines whether and how the reference data should be refreshed. Accepted values are `Static`, `RefreshPeriodicallyWithFull` and `RefreshPeriodicallyWithDelta`.
	RefreshType pulumi.StringOutput `pulumi:"refreshType"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The fully qualified domain name of the MS SQL server.
	Server pulumi.StringOutput `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// The username to connect to the MS SQL database.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewReferenceInputMssql registers a new resource with the given unique name, arguments, and options.
func NewReferenceInputMssql(ctx *pulumi.Context,
	name string, args *ReferenceInputMssqlArgs, opts ...pulumi.ResourceOption) (*ReferenceInputMssql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.FullSnapshotQuery == nil {
		return nil, errors.New("invalid value for required argument 'FullSnapshotQuery'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.RefreshType == nil {
		return nil, errors.New("invalid value for required argument 'RefreshType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource ReferenceInputMssql
	err := ctx.RegisterResource("azure:streamanalytics/referenceInputMssql:ReferenceInputMssql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceInputMssql gets an existing ReferenceInputMssql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceInputMssql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceInputMssqlState, opts ...pulumi.ResourceOption) (*ReferenceInputMssql, error) {
	var resource ReferenceInputMssql
	err := ctx.ReadResource("azure:streamanalytics/referenceInputMssql:ReferenceInputMssql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceInputMssql resources.
type referenceInputMssqlState struct {
	// The MS SQL database name where the reference data exists.
	Database *string `pulumi:"database"`
	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when `refreshType` is `Static`.
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery *string `pulumi:"fullSnapshotQuery"`
	// The name of the Reference Input MS SQL data. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The username to connect to the MS SQL database.
	Password *string `pulumi:"password"`
	// The frequency in `hh:mm:ss` with which the reference data should be retrieved from the MS SQL database e.g. `00:20:00` for every 20 minutes. Must be set when `refreshType` is `RefreshPeriodicallyWithFull` or `RefreshPeriodicallyWithDelta`.
	RefreshIntervalDuration *string `pulumi:"refreshIntervalDuration"`
	// Defines whether and how the reference data should be refreshed. Accepted values are `Static`, `RefreshPeriodicallyWithFull` and `RefreshPeriodicallyWithDelta`.
	RefreshType *string `pulumi:"refreshType"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The fully qualified domain name of the MS SQL server.
	Server *string `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// The username to connect to the MS SQL database.
	Username *string `pulumi:"username"`
}

type ReferenceInputMssqlState struct {
	// The MS SQL database name where the reference data exists.
	Database pulumi.StringPtrInput
	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when `refreshType` is `Static`.
	DeltaSnapshotQuery pulumi.StringPtrInput
	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery pulumi.StringPtrInput
	// The name of the Reference Input MS SQL data. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The username to connect to the MS SQL database.
	Password pulumi.StringPtrInput
	// The frequency in `hh:mm:ss` with which the reference data should be retrieved from the MS SQL database e.g. `00:20:00` for every 20 minutes. Must be set when `refreshType` is `RefreshPeriodicallyWithFull` or `RefreshPeriodicallyWithDelta`.
	RefreshIntervalDuration pulumi.StringPtrInput
	// Defines whether and how the reference data should be refreshed. Accepted values are `Static`, `RefreshPeriodicallyWithFull` and `RefreshPeriodicallyWithDelta`.
	RefreshType pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The fully qualified domain name of the MS SQL server.
	Server pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// The username to connect to the MS SQL database.
	Username pulumi.StringPtrInput
}

func (ReferenceInputMssqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceInputMssqlState)(nil)).Elem()
}

type referenceInputMssqlArgs struct {
	// The MS SQL database name where the reference data exists.
	Database string `pulumi:"database"`
	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when `refreshType` is `Static`.
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery string `pulumi:"fullSnapshotQuery"`
	// The name of the Reference Input MS SQL data. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The username to connect to the MS SQL database.
	Password string `pulumi:"password"`
	// The frequency in `hh:mm:ss` with which the reference data should be retrieved from the MS SQL database e.g. `00:20:00` for every 20 minutes. Must be set when `refreshType` is `RefreshPeriodicallyWithFull` or `RefreshPeriodicallyWithDelta`.
	RefreshIntervalDuration *string `pulumi:"refreshIntervalDuration"`
	// Defines whether and how the reference data should be refreshed. Accepted values are `Static`, `RefreshPeriodicallyWithFull` and `RefreshPeriodicallyWithDelta`.
	RefreshType string `pulumi:"refreshType"`
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The fully qualified domain name of the MS SQL server.
	Server string `pulumi:"server"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// The username to connect to the MS SQL database.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ReferenceInputMssql resource.
type ReferenceInputMssqlArgs struct {
	// The MS SQL database name where the reference data exists.
	Database pulumi.StringInput
	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when `refreshType` is `Static`.
	DeltaSnapshotQuery pulumi.StringPtrInput
	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery pulumi.StringInput
	// The name of the Reference Input MS SQL data. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The username to connect to the MS SQL database.
	Password pulumi.StringInput
	// The frequency in `hh:mm:ss` with which the reference data should be retrieved from the MS SQL database e.g. `00:20:00` for every 20 minutes. Must be set when `refreshType` is `RefreshPeriodicallyWithFull` or `RefreshPeriodicallyWithDelta`.
	RefreshIntervalDuration pulumi.StringPtrInput
	// Defines whether and how the reference data should be refreshed. Accepted values are `Static`, `RefreshPeriodicallyWithFull` and `RefreshPeriodicallyWithDelta`.
	RefreshType pulumi.StringInput
	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The fully qualified domain name of the MS SQL server.
	Server pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// The username to connect to the MS SQL database.
	Username pulumi.StringInput
}

func (ReferenceInputMssqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceInputMssqlArgs)(nil)).Elem()
}

type ReferenceInputMssqlInput interface {
	pulumi.Input

	ToReferenceInputMssqlOutput() ReferenceInputMssqlOutput
	ToReferenceInputMssqlOutputWithContext(ctx context.Context) ReferenceInputMssqlOutput
}

func (*ReferenceInputMssql) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputMssql)(nil))
}

func (i *ReferenceInputMssql) ToReferenceInputMssqlOutput() ReferenceInputMssqlOutput {
	return i.ToReferenceInputMssqlOutputWithContext(context.Background())
}

func (i *ReferenceInputMssql) ToReferenceInputMssqlOutputWithContext(ctx context.Context) ReferenceInputMssqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputMssqlOutput)
}

func (i *ReferenceInputMssql) ToReferenceInputMssqlPtrOutput() ReferenceInputMssqlPtrOutput {
	return i.ToReferenceInputMssqlPtrOutputWithContext(context.Background())
}

func (i *ReferenceInputMssql) ToReferenceInputMssqlPtrOutputWithContext(ctx context.Context) ReferenceInputMssqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputMssqlPtrOutput)
}

type ReferenceInputMssqlPtrInput interface {
	pulumi.Input

	ToReferenceInputMssqlPtrOutput() ReferenceInputMssqlPtrOutput
	ToReferenceInputMssqlPtrOutputWithContext(ctx context.Context) ReferenceInputMssqlPtrOutput
}

type referenceInputMssqlPtrType ReferenceInputMssqlArgs

func (*referenceInputMssqlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceInputMssql)(nil))
}

func (i *referenceInputMssqlPtrType) ToReferenceInputMssqlPtrOutput() ReferenceInputMssqlPtrOutput {
	return i.ToReferenceInputMssqlPtrOutputWithContext(context.Background())
}

func (i *referenceInputMssqlPtrType) ToReferenceInputMssqlPtrOutputWithContext(ctx context.Context) ReferenceInputMssqlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputMssqlPtrOutput)
}

// ReferenceInputMssqlArrayInput is an input type that accepts ReferenceInputMssqlArray and ReferenceInputMssqlArrayOutput values.
// You can construct a concrete instance of `ReferenceInputMssqlArrayInput` via:
//
//          ReferenceInputMssqlArray{ ReferenceInputMssqlArgs{...} }
type ReferenceInputMssqlArrayInput interface {
	pulumi.Input

	ToReferenceInputMssqlArrayOutput() ReferenceInputMssqlArrayOutput
	ToReferenceInputMssqlArrayOutputWithContext(context.Context) ReferenceInputMssqlArrayOutput
}

type ReferenceInputMssqlArray []ReferenceInputMssqlInput

func (ReferenceInputMssqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReferenceInputMssql)(nil)).Elem()
}

func (i ReferenceInputMssqlArray) ToReferenceInputMssqlArrayOutput() ReferenceInputMssqlArrayOutput {
	return i.ToReferenceInputMssqlArrayOutputWithContext(context.Background())
}

func (i ReferenceInputMssqlArray) ToReferenceInputMssqlArrayOutputWithContext(ctx context.Context) ReferenceInputMssqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputMssqlArrayOutput)
}

// ReferenceInputMssqlMapInput is an input type that accepts ReferenceInputMssqlMap and ReferenceInputMssqlMapOutput values.
// You can construct a concrete instance of `ReferenceInputMssqlMapInput` via:
//
//          ReferenceInputMssqlMap{ "key": ReferenceInputMssqlArgs{...} }
type ReferenceInputMssqlMapInput interface {
	pulumi.Input

	ToReferenceInputMssqlMapOutput() ReferenceInputMssqlMapOutput
	ToReferenceInputMssqlMapOutputWithContext(context.Context) ReferenceInputMssqlMapOutput
}

type ReferenceInputMssqlMap map[string]ReferenceInputMssqlInput

func (ReferenceInputMssqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReferenceInputMssql)(nil)).Elem()
}

func (i ReferenceInputMssqlMap) ToReferenceInputMssqlMapOutput() ReferenceInputMssqlMapOutput {
	return i.ToReferenceInputMssqlMapOutputWithContext(context.Background())
}

func (i ReferenceInputMssqlMap) ToReferenceInputMssqlMapOutputWithContext(ctx context.Context) ReferenceInputMssqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputMssqlMapOutput)
}

type ReferenceInputMssqlOutput struct{ *pulumi.OutputState }

func (ReferenceInputMssqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputMssql)(nil))
}

func (o ReferenceInputMssqlOutput) ToReferenceInputMssqlOutput() ReferenceInputMssqlOutput {
	return o
}

func (o ReferenceInputMssqlOutput) ToReferenceInputMssqlOutputWithContext(ctx context.Context) ReferenceInputMssqlOutput {
	return o
}

func (o ReferenceInputMssqlOutput) ToReferenceInputMssqlPtrOutput() ReferenceInputMssqlPtrOutput {
	return o.ToReferenceInputMssqlPtrOutputWithContext(context.Background())
}

func (o ReferenceInputMssqlOutput) ToReferenceInputMssqlPtrOutputWithContext(ctx context.Context) ReferenceInputMssqlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReferenceInputMssql) *ReferenceInputMssql {
		return &v
	}).(ReferenceInputMssqlPtrOutput)
}

type ReferenceInputMssqlPtrOutput struct{ *pulumi.OutputState }

func (ReferenceInputMssqlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceInputMssql)(nil))
}

func (o ReferenceInputMssqlPtrOutput) ToReferenceInputMssqlPtrOutput() ReferenceInputMssqlPtrOutput {
	return o
}

func (o ReferenceInputMssqlPtrOutput) ToReferenceInputMssqlPtrOutputWithContext(ctx context.Context) ReferenceInputMssqlPtrOutput {
	return o
}

func (o ReferenceInputMssqlPtrOutput) Elem() ReferenceInputMssqlOutput {
	return o.ApplyT(func(v *ReferenceInputMssql) ReferenceInputMssql {
		if v != nil {
			return *v
		}
		var ret ReferenceInputMssql
		return ret
	}).(ReferenceInputMssqlOutput)
}

type ReferenceInputMssqlArrayOutput struct{ *pulumi.OutputState }

func (ReferenceInputMssqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceInputMssql)(nil))
}

func (o ReferenceInputMssqlArrayOutput) ToReferenceInputMssqlArrayOutput() ReferenceInputMssqlArrayOutput {
	return o
}

func (o ReferenceInputMssqlArrayOutput) ToReferenceInputMssqlArrayOutputWithContext(ctx context.Context) ReferenceInputMssqlArrayOutput {
	return o
}

func (o ReferenceInputMssqlArrayOutput) Index(i pulumi.IntInput) ReferenceInputMssqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceInputMssql {
		return vs[0].([]ReferenceInputMssql)[vs[1].(int)]
	}).(ReferenceInputMssqlOutput)
}

type ReferenceInputMssqlMapOutput struct{ *pulumi.OutputState }

func (ReferenceInputMssqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReferenceInputMssql)(nil))
}

func (o ReferenceInputMssqlMapOutput) ToReferenceInputMssqlMapOutput() ReferenceInputMssqlMapOutput {
	return o
}

func (o ReferenceInputMssqlMapOutput) ToReferenceInputMssqlMapOutputWithContext(ctx context.Context) ReferenceInputMssqlMapOutput {
	return o
}

func (o ReferenceInputMssqlMapOutput) MapIndex(k pulumi.StringInput) ReferenceInputMssqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReferenceInputMssql {
		return vs[0].(map[string]ReferenceInputMssql)[vs[1].(string)]
	}).(ReferenceInputMssqlOutput)
}

func init() {
	pulumi.RegisterOutputType(ReferenceInputMssqlOutput{})
	pulumi.RegisterOutputType(ReferenceInputMssqlPtrOutput{})
	pulumi.RegisterOutputType(ReferenceInputMssqlArrayOutput{})
	pulumi.RegisterOutputType(ReferenceInputMssqlMapOutput{})
}
