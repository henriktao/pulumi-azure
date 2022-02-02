// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a MariaDB Database within a MariaDB Server
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mariadb"
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
// 		exampleServer, err := mariadb.NewServer(ctx, "exampleServer", &mariadb.ServerArgs{
// 			Location:                   exampleResourceGroup.Location,
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			SkuName:                    pulumi.String("B_Gen5_2"),
// 			StorageMb:                  pulumi.Int(51200),
// 			BackupRetentionDays:        pulumi.Int(7),
// 			GeoRedundantBackupEnabled:  pulumi.Bool(false),
// 			AdministratorLogin:         pulumi.String("acctestun"),
// 			AdministratorLoginPassword: pulumi.String("H@Sh1CoR3!"),
// 			Version:                    pulumi.String("10.2"),
// 			SslEnforcementEnabled:      pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mariadb.NewDatabase(ctx, "exampleDatabase", &mariadb.DatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			Charset:           pulumi.String("utf8"),
// 			Collation:         pulumi.String("utf8_general_ci"),
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
// MariaDB Database's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mariadb/database:Database database1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1/databases/database1
// ```
type Database struct {
	pulumi.CustomResourceState

	// Specifies the Charset for the MariaDB Database, which needs [to be a valid MariaDB Charset](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Charset pulumi.StringOutput `pulumi:"charset"`
	// Specifies the Collation for the MariaDB Database, which needs [to be a valid MariaDB Collation](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// Specifies the name of the MariaDB Database, which needs [to be a valid MariaDB identifier](https://mariadb.com/kb/en/library/identifier-names/). Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Charset == nil {
		return nil, errors.New("invalid value for required argument 'Charset'")
	}
	if args.Collation == nil {
		return nil, errors.New("invalid value for required argument 'Collation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource Database
	err := ctx.RegisterResource("azure:mariadb/database:Database", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:mariadb/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Specifies the Charset for the MariaDB Database, which needs [to be a valid MariaDB Charset](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Charset *string `pulumi:"charset"`
	// Specifies the Collation for the MariaDB Database, which needs [to be a valid MariaDB Collation](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies the name of the MariaDB Database, which needs [to be a valid MariaDB identifier](https://mariadb.com/kb/en/library/identifier-names/). Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
}

type DatabaseState struct {
	// Specifies the Charset for the MariaDB Database, which needs [to be a valid MariaDB Charset](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Charset pulumi.StringPtrInput
	// Specifies the Collation for the MariaDB Database, which needs [to be a valid MariaDB Collation](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies the name of the MariaDB Database, which needs [to be a valid MariaDB identifier](https://mariadb.com/kb/en/library/identifier-names/). Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Specifies the Charset for the MariaDB Database, which needs [to be a valid MariaDB Charset](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Charset string `pulumi:"charset"`
	// Specifies the Collation for the MariaDB Database, which needs [to be a valid MariaDB Collation](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Collation string `pulumi:"collation"`
	// Specifies the name of the MariaDB Database, which needs [to be a valid MariaDB identifier](https://mariadb.com/kb/en/library/identifier-names/). Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Specifies the Charset for the MariaDB Database, which needs [to be a valid MariaDB Charset](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Charset pulumi.StringInput
	// Specifies the Collation for the MariaDB Database, which needs [to be a valid MariaDB Collation](https://mariadb.com/kb/en/library/setting-character-sets-and-collations). Changing this forces a new resource to be created.
	Collation pulumi.StringInput
	// Specifies the name of the MariaDB Database, which needs [to be a valid MariaDB identifier](https://mariadb.com/kb/en/library/identifier-names/). Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
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
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
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
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
