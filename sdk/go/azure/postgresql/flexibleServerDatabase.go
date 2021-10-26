// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a PostgreSQL Flexible Server Database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/postgresql"
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
// 		exampleFlexibleServer, err := postgresql.NewFlexibleServer(ctx, "exampleFlexibleServer", &postgresql.FlexibleServerArgs{
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			Location:              exampleResourceGroup.Location,
// 			Version:               pulumi.String("12"),
// 			AdministratorLogin:    pulumi.String("psqladmin"),
// 			AdministratorPassword: pulumi.String("H@Sh1CoR3!"),
// 			StorageMb:             pulumi.Int(32768),
// 			SkuName:               pulumi.String("GP_Standard_D4s_v3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = postgresql.NewFlexibleServerDatabase(ctx, "exampleFlexibleServerDatabase", &postgresql.FlexibleServerDatabaseArgs{
// 			ServerId:  exampleFlexibleServer.ID(),
// 			Collation: pulumi.String("en_US.utf8"),
// 			Charset:   pulumi.String("utf8"),
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
// Azure PostgreSQL Flexible Server Database can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:postgresql/flexibleServerDatabase:FlexibleServerDatabase example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/databases/database1
// ```
type FlexibleServerDatabase struct {
	pulumi.CustomResourceState

	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Defaults to `UTF8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset pulumi.StringPtrOutput `pulumi:"charset"`
	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Defaults to `en_US.utf8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The name which should be used for this Azure PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
}

// NewFlexibleServerDatabase registers a new resource with the given unique name, arguments, and options.
func NewFlexibleServerDatabase(ctx *pulumi.Context,
	name string, args *FlexibleServerDatabaseArgs, opts ...pulumi.ResourceOption) (*FlexibleServerDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	var resource FlexibleServerDatabase
	err := ctx.RegisterResource("azure:postgresql/flexibleServerDatabase:FlexibleServerDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleServerDatabase gets an existing FlexibleServerDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleServerDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleServerDatabaseState, opts ...pulumi.ResourceOption) (*FlexibleServerDatabase, error) {
	var resource FlexibleServerDatabase
	err := ctx.ReadResource("azure:postgresql/flexibleServerDatabase:FlexibleServerDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleServerDatabase resources.
type flexibleServerDatabaseState struct {
	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Defaults to `UTF8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset *string `pulumi:"charset"`
	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Defaults to `en_US.utf8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation *string `pulumi:"collation"`
	// The name which should be used for this Azure PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Name *string `pulumi:"name"`
	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerId *string `pulumi:"serverId"`
}

type FlexibleServerDatabaseState struct {
	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Defaults to `UTF8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset pulumi.StringPtrInput
	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Defaults to `en_US.utf8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation pulumi.StringPtrInput
	// The name which should be used for this Azure PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Name pulumi.StringPtrInput
	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerId pulumi.StringPtrInput
}

func (FlexibleServerDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerDatabaseState)(nil)).Elem()
}

type flexibleServerDatabaseArgs struct {
	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Defaults to `UTF8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset *string `pulumi:"charset"`
	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Defaults to `en_US.utf8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation *string `pulumi:"collation"`
	// The name which should be used for this Azure PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Name *string `pulumi:"name"`
	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerId string `pulumi:"serverId"`
}

// The set of arguments for constructing a FlexibleServerDatabase resource.
type FlexibleServerDatabaseArgs struct {
	// Specifies the Charset for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Defaults to `UTF8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Charset pulumi.StringPtrInput
	// Specifies the Collation for the Azure PostgreSQL Flexible Server Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Defaults to `en_US.utf8`. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Collation pulumi.StringPtrInput
	// The name which should be used for this Azure PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	Name pulumi.StringPtrInput
	// The ID of the Azure PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Database. Changing this forces a new Azure PostgreSQL Flexible Server Database to be created.
	ServerId pulumi.StringInput
}

func (FlexibleServerDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleServerDatabaseArgs)(nil)).Elem()
}

type FlexibleServerDatabaseInput interface {
	pulumi.Input

	ToFlexibleServerDatabaseOutput() FlexibleServerDatabaseOutput
	ToFlexibleServerDatabaseOutputWithContext(ctx context.Context) FlexibleServerDatabaseOutput
}

func (*FlexibleServerDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServerDatabase)(nil))
}

func (i *FlexibleServerDatabase) ToFlexibleServerDatabaseOutput() FlexibleServerDatabaseOutput {
	return i.ToFlexibleServerDatabaseOutputWithContext(context.Background())
}

func (i *FlexibleServerDatabase) ToFlexibleServerDatabaseOutputWithContext(ctx context.Context) FlexibleServerDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerDatabaseOutput)
}

func (i *FlexibleServerDatabase) ToFlexibleServerDatabasePtrOutput() FlexibleServerDatabasePtrOutput {
	return i.ToFlexibleServerDatabasePtrOutputWithContext(context.Background())
}

func (i *FlexibleServerDatabase) ToFlexibleServerDatabasePtrOutputWithContext(ctx context.Context) FlexibleServerDatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerDatabasePtrOutput)
}

type FlexibleServerDatabasePtrInput interface {
	pulumi.Input

	ToFlexibleServerDatabasePtrOutput() FlexibleServerDatabasePtrOutput
	ToFlexibleServerDatabasePtrOutputWithContext(ctx context.Context) FlexibleServerDatabasePtrOutput
}

type flexibleServerDatabasePtrType FlexibleServerDatabaseArgs

func (*flexibleServerDatabasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerDatabase)(nil))
}

func (i *flexibleServerDatabasePtrType) ToFlexibleServerDatabasePtrOutput() FlexibleServerDatabasePtrOutput {
	return i.ToFlexibleServerDatabasePtrOutputWithContext(context.Background())
}

func (i *flexibleServerDatabasePtrType) ToFlexibleServerDatabasePtrOutputWithContext(ctx context.Context) FlexibleServerDatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerDatabasePtrOutput)
}

// FlexibleServerDatabaseArrayInput is an input type that accepts FlexibleServerDatabaseArray and FlexibleServerDatabaseArrayOutput values.
// You can construct a concrete instance of `FlexibleServerDatabaseArrayInput` via:
//
//          FlexibleServerDatabaseArray{ FlexibleServerDatabaseArgs{...} }
type FlexibleServerDatabaseArrayInput interface {
	pulumi.Input

	ToFlexibleServerDatabaseArrayOutput() FlexibleServerDatabaseArrayOutput
	ToFlexibleServerDatabaseArrayOutputWithContext(context.Context) FlexibleServerDatabaseArrayOutput
}

type FlexibleServerDatabaseArray []FlexibleServerDatabaseInput

func (FlexibleServerDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleServerDatabase)(nil)).Elem()
}

func (i FlexibleServerDatabaseArray) ToFlexibleServerDatabaseArrayOutput() FlexibleServerDatabaseArrayOutput {
	return i.ToFlexibleServerDatabaseArrayOutputWithContext(context.Background())
}

func (i FlexibleServerDatabaseArray) ToFlexibleServerDatabaseArrayOutputWithContext(ctx context.Context) FlexibleServerDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerDatabaseArrayOutput)
}

// FlexibleServerDatabaseMapInput is an input type that accepts FlexibleServerDatabaseMap and FlexibleServerDatabaseMapOutput values.
// You can construct a concrete instance of `FlexibleServerDatabaseMapInput` via:
//
//          FlexibleServerDatabaseMap{ "key": FlexibleServerDatabaseArgs{...} }
type FlexibleServerDatabaseMapInput interface {
	pulumi.Input

	ToFlexibleServerDatabaseMapOutput() FlexibleServerDatabaseMapOutput
	ToFlexibleServerDatabaseMapOutputWithContext(context.Context) FlexibleServerDatabaseMapOutput
}

type FlexibleServerDatabaseMap map[string]FlexibleServerDatabaseInput

func (FlexibleServerDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleServerDatabase)(nil)).Elem()
}

func (i FlexibleServerDatabaseMap) ToFlexibleServerDatabaseMapOutput() FlexibleServerDatabaseMapOutput {
	return i.ToFlexibleServerDatabaseMapOutputWithContext(context.Background())
}

func (i FlexibleServerDatabaseMap) ToFlexibleServerDatabaseMapOutputWithContext(ctx context.Context) FlexibleServerDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleServerDatabaseMapOutput)
}

type FlexibleServerDatabaseOutput struct{ *pulumi.OutputState }

func (FlexibleServerDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexibleServerDatabase)(nil))
}

func (o FlexibleServerDatabaseOutput) ToFlexibleServerDatabaseOutput() FlexibleServerDatabaseOutput {
	return o
}

func (o FlexibleServerDatabaseOutput) ToFlexibleServerDatabaseOutputWithContext(ctx context.Context) FlexibleServerDatabaseOutput {
	return o
}

func (o FlexibleServerDatabaseOutput) ToFlexibleServerDatabasePtrOutput() FlexibleServerDatabasePtrOutput {
	return o.ToFlexibleServerDatabasePtrOutputWithContext(context.Background())
}

func (o FlexibleServerDatabaseOutput) ToFlexibleServerDatabasePtrOutputWithContext(ctx context.Context) FlexibleServerDatabasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlexibleServerDatabase) *FlexibleServerDatabase {
		return &v
	}).(FlexibleServerDatabasePtrOutput)
}

type FlexibleServerDatabasePtrOutput struct{ *pulumi.OutputState }

func (FlexibleServerDatabasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleServerDatabase)(nil))
}

func (o FlexibleServerDatabasePtrOutput) ToFlexibleServerDatabasePtrOutput() FlexibleServerDatabasePtrOutput {
	return o
}

func (o FlexibleServerDatabasePtrOutput) ToFlexibleServerDatabasePtrOutputWithContext(ctx context.Context) FlexibleServerDatabasePtrOutput {
	return o
}

func (o FlexibleServerDatabasePtrOutput) Elem() FlexibleServerDatabaseOutput {
	return o.ApplyT(func(v *FlexibleServerDatabase) FlexibleServerDatabase {
		if v != nil {
			return *v
		}
		var ret FlexibleServerDatabase
		return ret
	}).(FlexibleServerDatabaseOutput)
}

type FlexibleServerDatabaseArrayOutput struct{ *pulumi.OutputState }

func (FlexibleServerDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FlexibleServerDatabase)(nil))
}

func (o FlexibleServerDatabaseArrayOutput) ToFlexibleServerDatabaseArrayOutput() FlexibleServerDatabaseArrayOutput {
	return o
}

func (o FlexibleServerDatabaseArrayOutput) ToFlexibleServerDatabaseArrayOutputWithContext(ctx context.Context) FlexibleServerDatabaseArrayOutput {
	return o
}

func (o FlexibleServerDatabaseArrayOutput) Index(i pulumi.IntInput) FlexibleServerDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FlexibleServerDatabase {
		return vs[0].([]FlexibleServerDatabase)[vs[1].(int)]
	}).(FlexibleServerDatabaseOutput)
}

type FlexibleServerDatabaseMapOutput struct{ *pulumi.OutputState }

func (FlexibleServerDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlexibleServerDatabase)(nil))
}

func (o FlexibleServerDatabaseMapOutput) ToFlexibleServerDatabaseMapOutput() FlexibleServerDatabaseMapOutput {
	return o
}

func (o FlexibleServerDatabaseMapOutput) ToFlexibleServerDatabaseMapOutputWithContext(ctx context.Context) FlexibleServerDatabaseMapOutput {
	return o
}

func (o FlexibleServerDatabaseMapOutput) MapIndex(k pulumi.StringInput) FlexibleServerDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlexibleServerDatabase {
		return vs[0].(map[string]FlexibleServerDatabase)[vs[1].(string)]
	}).(FlexibleServerDatabaseOutput)
}

func init() {
	pulumi.RegisterOutputType(FlexibleServerDatabaseOutput{})
	pulumi.RegisterOutputType(FlexibleServerDatabasePtrOutput{})
	pulumi.RegisterOutputType(FlexibleServerDatabaseArrayOutput{})
	pulumi.RegisterOutputType(FlexibleServerDatabaseMapOutput{})
}
