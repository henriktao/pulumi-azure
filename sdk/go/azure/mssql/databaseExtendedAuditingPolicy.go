// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Ms Sql Database Extended Auditing Policy.
//
// > **NOTE:** The Database Extended Auditing Policy Can be set inline here as well as with the mssqlDatabaseExtendedAuditingPolicy resource resource. You can only use one or the other and using both will cause a conflict.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mssql"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 		exampleServer, err := mssql.NewServer(ctx, "exampleServer", &mssql.ServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("missadministrator"),
// 			AdministratorLoginPassword: pulumi.String("AdminPassword123!"),
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewDatabaseExtendedAuditingPolicy(ctx, "exampleDatabaseExtendedAuditingPolicy", &mssql.DatabaseExtendedAuditingPolicyArgs{
// 			DatabaseId:                         exampleDatabase.ID(),
// 			StorageEndpoint:                    exampleAccount.PrimaryBlobEndpoint,
// 			StorageAccountAccessKey:            exampleAccount.PrimaryAccessKey,
// 			StorageAccountAccessKeyIsSecondary: pulumi.Bool(false),
// 			RetentionInDays:                    pulumi.Int(6),
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
// Ms Sql Database Extended Auditing Policys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
// ```
type DatabaseExtendedAuditingPolicy struct {
	pulumi.CustomResourceState

	// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
	DatabaseId           pulumi.StringOutput  `pulumi:"databaseId"`
	LogMonitoringEnabled pulumi.BoolPtrOutput `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrOutput `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
}

// NewDatabaseExtendedAuditingPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, args *DatabaseExtendedAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseExtendedAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	var resource DatabaseExtendedAuditingPolicy
	err := ctx.RegisterResource("azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseExtendedAuditingPolicy gets an existing DatabaseExtendedAuditingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseExtendedAuditingPolicyState, opts ...pulumi.ResourceOption) (*DatabaseExtendedAuditingPolicy, error) {
	var resource DatabaseExtendedAuditingPolicy
	err := ctx.ReadResource("azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseExtendedAuditingPolicy resources.
type databaseExtendedAuditingPolicyState struct {
	// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
	DatabaseId           *string `pulumi:"databaseId"`
	LogMonitoringEnabled *bool   `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

type DatabaseExtendedAuditingPolicyState struct {
	// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
	DatabaseId           pulumi.StringPtrInput
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (DatabaseExtendedAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseExtendedAuditingPolicyState)(nil)).Elem()
}

type databaseExtendedAuditingPolicyArgs struct {
	// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
	DatabaseId           string `pulumi:"databaseId"`
	LogMonitoringEnabled *bool  `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a DatabaseExtendedAuditingPolicy resource.
type DatabaseExtendedAuditingPolicyArgs struct {
	// The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
	DatabaseId           pulumi.StringInput
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (DatabaseExtendedAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseExtendedAuditingPolicyArgs)(nil)).Elem()
}

type DatabaseExtendedAuditingPolicyInput interface {
	pulumi.Input

	ToDatabaseExtendedAuditingPolicyOutput() DatabaseExtendedAuditingPolicyOutput
	ToDatabaseExtendedAuditingPolicyOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyOutput
}

func (*DatabaseExtendedAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseExtendedAuditingPolicy)(nil))
}

func (i *DatabaseExtendedAuditingPolicy) ToDatabaseExtendedAuditingPolicyOutput() DatabaseExtendedAuditingPolicyOutput {
	return i.ToDatabaseExtendedAuditingPolicyOutputWithContext(context.Background())
}

func (i *DatabaseExtendedAuditingPolicy) ToDatabaseExtendedAuditingPolicyOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExtendedAuditingPolicyOutput)
}

func (i *DatabaseExtendedAuditingPolicy) ToDatabaseExtendedAuditingPolicyPtrOutput() DatabaseExtendedAuditingPolicyPtrOutput {
	return i.ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *DatabaseExtendedAuditingPolicy) ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExtendedAuditingPolicyPtrOutput)
}

type DatabaseExtendedAuditingPolicyPtrInput interface {
	pulumi.Input

	ToDatabaseExtendedAuditingPolicyPtrOutput() DatabaseExtendedAuditingPolicyPtrOutput
	ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyPtrOutput
}

type databaseExtendedAuditingPolicyPtrType DatabaseExtendedAuditingPolicyArgs

func (*databaseExtendedAuditingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseExtendedAuditingPolicy)(nil))
}

func (i *databaseExtendedAuditingPolicyPtrType) ToDatabaseExtendedAuditingPolicyPtrOutput() DatabaseExtendedAuditingPolicyPtrOutput {
	return i.ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *databaseExtendedAuditingPolicyPtrType) ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExtendedAuditingPolicyPtrOutput)
}

// DatabaseExtendedAuditingPolicyArrayInput is an input type that accepts DatabaseExtendedAuditingPolicyArray and DatabaseExtendedAuditingPolicyArrayOutput values.
// You can construct a concrete instance of `DatabaseExtendedAuditingPolicyArrayInput` via:
//
//          DatabaseExtendedAuditingPolicyArray{ DatabaseExtendedAuditingPolicyArgs{...} }
type DatabaseExtendedAuditingPolicyArrayInput interface {
	pulumi.Input

	ToDatabaseExtendedAuditingPolicyArrayOutput() DatabaseExtendedAuditingPolicyArrayOutput
	ToDatabaseExtendedAuditingPolicyArrayOutputWithContext(context.Context) DatabaseExtendedAuditingPolicyArrayOutput
}

type DatabaseExtendedAuditingPolicyArray []DatabaseExtendedAuditingPolicyInput

func (DatabaseExtendedAuditingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatabaseExtendedAuditingPolicy)(nil))
}

func (i DatabaseExtendedAuditingPolicyArray) ToDatabaseExtendedAuditingPolicyArrayOutput() DatabaseExtendedAuditingPolicyArrayOutput {
	return i.ToDatabaseExtendedAuditingPolicyArrayOutputWithContext(context.Background())
}

func (i DatabaseExtendedAuditingPolicyArray) ToDatabaseExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExtendedAuditingPolicyArrayOutput)
}

// DatabaseExtendedAuditingPolicyMapInput is an input type that accepts DatabaseExtendedAuditingPolicyMap and DatabaseExtendedAuditingPolicyMapOutput values.
// You can construct a concrete instance of `DatabaseExtendedAuditingPolicyMapInput` via:
//
//          DatabaseExtendedAuditingPolicyMap{ "key": DatabaseExtendedAuditingPolicyArgs{...} }
type DatabaseExtendedAuditingPolicyMapInput interface {
	pulumi.Input

	ToDatabaseExtendedAuditingPolicyMapOutput() DatabaseExtendedAuditingPolicyMapOutput
	ToDatabaseExtendedAuditingPolicyMapOutputWithContext(context.Context) DatabaseExtendedAuditingPolicyMapOutput
}

type DatabaseExtendedAuditingPolicyMap map[string]DatabaseExtendedAuditingPolicyInput

func (DatabaseExtendedAuditingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatabaseExtendedAuditingPolicy)(nil))
}

func (i DatabaseExtendedAuditingPolicyMap) ToDatabaseExtendedAuditingPolicyMapOutput() DatabaseExtendedAuditingPolicyMapOutput {
	return i.ToDatabaseExtendedAuditingPolicyMapOutputWithContext(context.Background())
}

func (i DatabaseExtendedAuditingPolicyMap) ToDatabaseExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExtendedAuditingPolicyMapOutput)
}

type DatabaseExtendedAuditingPolicyOutput struct {
	*pulumi.OutputState
}

func (DatabaseExtendedAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseExtendedAuditingPolicy)(nil))
}

func (o DatabaseExtendedAuditingPolicyOutput) ToDatabaseExtendedAuditingPolicyOutput() DatabaseExtendedAuditingPolicyOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyOutput) ToDatabaseExtendedAuditingPolicyOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyOutput) ToDatabaseExtendedAuditingPolicyPtrOutput() DatabaseExtendedAuditingPolicyPtrOutput {
	return o.ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (o DatabaseExtendedAuditingPolicyOutput) ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyPtrOutput {
	return o.ApplyT(func(v DatabaseExtendedAuditingPolicy) *DatabaseExtendedAuditingPolicy {
		return &v
	}).(DatabaseExtendedAuditingPolicyPtrOutput)
}

type DatabaseExtendedAuditingPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (DatabaseExtendedAuditingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseExtendedAuditingPolicy)(nil))
}

func (o DatabaseExtendedAuditingPolicyPtrOutput) ToDatabaseExtendedAuditingPolicyPtrOutput() DatabaseExtendedAuditingPolicyPtrOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyPtrOutput) ToDatabaseExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyPtrOutput {
	return o
}

type DatabaseExtendedAuditingPolicyArrayOutput struct{ *pulumi.OutputState }

func (DatabaseExtendedAuditingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseExtendedAuditingPolicy)(nil))
}

func (o DatabaseExtendedAuditingPolicyArrayOutput) ToDatabaseExtendedAuditingPolicyArrayOutput() DatabaseExtendedAuditingPolicyArrayOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyArrayOutput) ToDatabaseExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyArrayOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyArrayOutput) Index(i pulumi.IntInput) DatabaseExtendedAuditingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseExtendedAuditingPolicy {
		return vs[0].([]DatabaseExtendedAuditingPolicy)[vs[1].(int)]
	}).(DatabaseExtendedAuditingPolicyOutput)
}

type DatabaseExtendedAuditingPolicyMapOutput struct{ *pulumi.OutputState }

func (DatabaseExtendedAuditingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseExtendedAuditingPolicy)(nil))
}

func (o DatabaseExtendedAuditingPolicyMapOutput) ToDatabaseExtendedAuditingPolicyMapOutput() DatabaseExtendedAuditingPolicyMapOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyMapOutput) ToDatabaseExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) DatabaseExtendedAuditingPolicyMapOutput {
	return o
}

func (o DatabaseExtendedAuditingPolicyMapOutput) MapIndex(k pulumi.StringInput) DatabaseExtendedAuditingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseExtendedAuditingPolicy {
		return vs[0].(map[string]DatabaseExtendedAuditingPolicy)[vs[1].(string)]
	}).(DatabaseExtendedAuditingPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseExtendedAuditingPolicyOutput{})
	pulumi.RegisterOutputType(DatabaseExtendedAuditingPolicyPtrOutput{})
	pulumi.RegisterOutputType(DatabaseExtendedAuditingPolicyArrayOutput{})
	pulumi.RegisterOutputType(DatabaseExtendedAuditingPolicyMapOutput{})
}
