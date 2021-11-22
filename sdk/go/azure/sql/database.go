// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// SQL Databases can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/database:Database database1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/databases/database1
// ```
type Database struct {
	pulumi.CustomResourceState

	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// The creation date of the SQL Database.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringOutput `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringOutput `pulumi:"elasticPoolName"`
	Encryption      pulumi.StringOutput `pulumi:"encryption"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyOutput `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrOutput `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringOutput `pulumi:"maxSizeBytes"`
	MaxSizeGb    pulumi.StringOutput `pulumi:"maxSizeGb"`
	// The name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrOutput `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringOutput `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: `shell az sql db list-editions -l westus -o table`. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringOutput `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringOutput `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringOutput `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringOutput `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyOutput `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource Database
	err := ctx.RegisterResource("azure:sql/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure:sql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode *string `pulumi:"createMode"`
	// The creation date of the SQL Database.
	CreationDate *string `pulumi:"creationDate"`
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation *string `pulumi:"defaultSecondaryLocation"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition *string `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	Encryption      *string `pulumi:"encryption"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy *DatabaseExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import *DatabaseImport `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	MaxSizeGb    *string `pulumi:"maxSizeGb"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale *bool `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: `shell az sql db list-editions -l westus -o table`. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName *string `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *DatabaseThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type DatabaseState struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrInput
	// The creation date of the SQL Database.
	CreationDate pulumi.StringPtrInput
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation pulumi.StringPtrInput
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringPtrInput
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringPtrInput
	Encryption      pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyPtrInput
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringPtrInput
	MaxSizeGb    pulumi.StringPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrInput
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringPtrInput
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: `shell az sql db list-editions -l westus -o table`. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringPtrInput
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringPtrInput
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringPtrInput
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode *string `pulumi:"createMode"`
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition *string `pulumi:"edition"`
	// The name of the elastic database pool.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy *DatabaseExtendedAuditingPolicy `pulumi:"extendedAuditingPolicy"`
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import *DatabaseImport `pulumi:"import"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	MaxSizeGb    *string `pulumi:"maxSizeGb"`
	// The name of the database.
	Name *string `pulumi:"name"`
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale *bool `pulumi:"readScale"`
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: `shell az sql db list-editions -l westus -o table`. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the SQL Server on which to create the database.
	ServerName string `pulumi:"serverName"`
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy *DatabaseThreatDetectionPolicy `pulumi:"threatDetectionPolicy"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
	CreateMode pulumi.StringPtrInput
	// The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	Edition pulumi.StringPtrInput
	// The name of the elastic database pool.
	ElasticPoolName pulumi.StringPtrInput
	// A `extendedAuditingPolicy` block as defined below.
	//
	// Deprecated: the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
	ExtendedAuditingPolicy DatabaseExtendedAuditingPolicyPtrInput
	// A Database Import block as documented below. `createMode` must be set to `Default`.
	Import DatabaseImportPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
	MaxSizeBytes pulumi.StringPtrInput
	MaxSizeGb    pulumi.StringPtrInput
	// The name of the database.
	Name pulumi.StringPtrInput
	// Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
	ReadScale pulumi.BoolPtrInput
	// A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
	// .
	RequestedServiceObjectiveId pulumi.StringPtrInput
	// The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: `shell az sql db list-editions -l westus -o table`. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
	ResourceGroupName pulumi.StringInput
	// The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
	RestorePointInTime pulumi.StringPtrInput
	// The name of the SQL Server on which to create the database.
	ServerName pulumi.StringInput
	// The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The URI of the source database if `createMode` value is not `Default`.
	SourceDatabaseId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
	ThreatDetectionPolicy DatabaseThreatDetectionPolicyPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

func (i *Database) ToDatabasePtrOutput() DatabasePtrOutput {
	return i.ToDatabasePtrOutputWithContext(context.Background())
}

func (i *Database) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePtrOutput)
}

type DatabasePtrInput interface {
	pulumi.Input

	ToDatabasePtrOutput() DatabasePtrOutput
	ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput
}

type databasePtrType DatabaseArgs

func (*databasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil))
}

func (i *databasePtrType) ToDatabasePtrOutput() DatabasePtrOutput {
	return i.ToDatabasePtrOutputWithContext(context.Background())
}

func (i *databasePtrType) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePtrOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//          DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//          DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabasePtrOutput() DatabasePtrOutput {
	return o.ToDatabasePtrOutputWithContext(context.Background())
}

func (o DatabaseOutput) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Database) *Database {
		return &v
	}).(DatabasePtrOutput)
}

type DatabasePtrOutput struct{ *pulumi.OutputState }

func (DatabasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil))
}

func (o DatabasePtrOutput) ToDatabasePtrOutput() DatabasePtrOutput {
	return o
}

func (o DatabasePtrOutput) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return o
}

func (o DatabasePtrOutput) Elem() DatabaseOutput {
	return o.ApplyT(func(v *Database) Database {
		if v != nil {
			return *v
		}
		var ret Database
		return ret
	}).(DatabaseOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Database)(nil))
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Database {
		return vs[0].([]Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Database)(nil))
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Database {
		return vs[0].(map[string]Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePtrInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabasePtrOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
