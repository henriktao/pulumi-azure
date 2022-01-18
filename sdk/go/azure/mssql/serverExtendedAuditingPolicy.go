// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Ms Sql Server Extended Auditing Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewServerExtendedAuditingPolicy(ctx, "exampleServerExtendedAuditingPolicy", &mssql.ServerExtendedAuditingPolicyArgs{
// 			ServerId:                           exampleServer.ID(),
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
// ### With Storage Account Behind VNet And Firewall
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/mssql"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/sql"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := core.LookupSubscription(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSubnet, err := network.NewSubnet(ctx, "exampleSubnet", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 			ServiceEndpoints: pulumi.StringArray{
// 				pulumi.String("Microsoft.Sql"),
// 				pulumi.String("Microsoft.Storage"),
// 			},
// 			EnforcePrivateLinkEndpointNetworkPolicies: pulumi.Bool(true),
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
// 			MinimumTlsVersion:          pulumi.String("1.2"),
// 			Identity: &mssql.ServerIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAssignment, err := authorization.NewAssignment(ctx, "exampleAssignment", &authorization.AssignmentArgs{
// 			Scope:              pulumi.String(primary.Id),
// 			RoleDefinitionName: pulumi.String("Storage Blob Data Contributor"),
// 			PrincipalId: exampleServer.Identity.ApplyT(func(identity mssql.ServerIdentity) (string, error) {
// 				return identity.PrincipalId, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewVirtualNetworkRule(ctx, "sqlvnetrule", &sql.VirtualNetworkRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			SubnetId:          exampleSubnet.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewFirewallRule(ctx, "exampleFirewallRule", &sql.FirewallRuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			StartIpAddress:    pulumi.String("0.0.0.0"),
// 			EndIpAddress:      pulumi.String("0.0.0.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("StorageV2"),
// 			AllowBlobPublicAccess:  pulumi.Bool(false),
// 			NetworkRules: &storage.AccountNetworkRulesArgs{
// 				DefaultAction: pulumi.String("Deny"),
// 				IpRules: pulumi.StringArray{
// 					pulumi.String("127.0.0.1"),
// 				},
// 				VirtualNetworkSubnetIds: pulumi.StringArray{
// 					exampleSubnet.ID(),
// 				},
// 				Bypasses: pulumi.StringArray{
// 					pulumi.String("AzureServices"),
// 				},
// 			},
// 			Identity: &storage.AccountIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mssql.NewServerExtendedAuditingPolicy(ctx, "exampleServerExtendedAuditingPolicy", &mssql.ServerExtendedAuditingPolicyArgs{
// 			StorageEndpoint:              exampleAccount.PrimaryBlobEndpoint,
// 			ServerId:                     exampleServer.ID(),
// 			RetentionInDays:              pulumi.Int(6),
// 			LogMonitoringEnabled:         pulumi.Bool(false),
// 			StorageAccountSubscriptionId: pulumi.Any(azurerm_subscription.Primary.Subscription_id),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAssignment,
// 			exampleAccount,
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
// Ms Sql Server Extended Auditing Policys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/extendedAuditingSettings/default
// ```
type ServerExtendedAuditingPolicy struct {
	pulumi.CustomResourceState

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrOutput `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrOutput `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The ID of the Subscription containing the Storage Account.
	StorageAccountSubscriptionId pulumi.StringPtrOutput `pulumi:"storageAccountSubscriptionId"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
}

// NewServerExtendedAuditingPolicy registers a new resource with the given unique name, arguments, and options.
func NewServerExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, args *ServerExtendedAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*ServerExtendedAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	var resource ServerExtendedAuditingPolicy
	err := ctx.RegisterResource("azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerExtendedAuditingPolicy gets an existing ServerExtendedAuditingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerExtendedAuditingPolicyState, opts ...pulumi.ResourceOption) (*ServerExtendedAuditingPolicy, error) {
	var resource ServerExtendedAuditingPolicy
	err := ctx.ReadResource("azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerExtendedAuditingPolicy resources.
type serverExtendedAuditingPolicyState struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
	LogMonitoringEnabled *bool `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerId *string `pulumi:"serverId"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The ID of the Subscription containing the Storage Account.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

type ServerExtendedAuditingPolicyState struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerId pulumi.StringPtrInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The ID of the Subscription containing the Storage Account.
	StorageAccountSubscriptionId pulumi.StringPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerExtendedAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverExtendedAuditingPolicyState)(nil)).Elem()
}

type serverExtendedAuditingPolicyArgs struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
	LogMonitoringEnabled *bool `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerId string `pulumi:"serverId"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The ID of the Subscription containing the Storage Account.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a ServerExtendedAuditingPolicy resource.
type ServerExtendedAuditingPolicyArgs struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
	ServerId pulumi.StringInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The ID of the Subscription containing the Storage Account.
	StorageAccountSubscriptionId pulumi.StringPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
}

func (ServerExtendedAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverExtendedAuditingPolicyArgs)(nil)).Elem()
}

type ServerExtendedAuditingPolicyInput interface {
	pulumi.Input

	ToServerExtendedAuditingPolicyOutput() ServerExtendedAuditingPolicyOutput
	ToServerExtendedAuditingPolicyOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyOutput
}

func (*ServerExtendedAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExtendedAuditingPolicy)(nil))
}

func (i *ServerExtendedAuditingPolicy) ToServerExtendedAuditingPolicyOutput() ServerExtendedAuditingPolicyOutput {
	return i.ToServerExtendedAuditingPolicyOutputWithContext(context.Background())
}

func (i *ServerExtendedAuditingPolicy) ToServerExtendedAuditingPolicyOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExtendedAuditingPolicyOutput)
}

func (i *ServerExtendedAuditingPolicy) ToServerExtendedAuditingPolicyPtrOutput() ServerExtendedAuditingPolicyPtrOutput {
	return i.ToServerExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *ServerExtendedAuditingPolicy) ToServerExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExtendedAuditingPolicyPtrOutput)
}

type ServerExtendedAuditingPolicyPtrInput interface {
	pulumi.Input

	ToServerExtendedAuditingPolicyPtrOutput() ServerExtendedAuditingPolicyPtrOutput
	ToServerExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyPtrOutput
}

type serverExtendedAuditingPolicyPtrType ServerExtendedAuditingPolicyArgs

func (*serverExtendedAuditingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExtendedAuditingPolicy)(nil))
}

func (i *serverExtendedAuditingPolicyPtrType) ToServerExtendedAuditingPolicyPtrOutput() ServerExtendedAuditingPolicyPtrOutput {
	return i.ToServerExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *serverExtendedAuditingPolicyPtrType) ToServerExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExtendedAuditingPolicyPtrOutput)
}

// ServerExtendedAuditingPolicyArrayInput is an input type that accepts ServerExtendedAuditingPolicyArray and ServerExtendedAuditingPolicyArrayOutput values.
// You can construct a concrete instance of `ServerExtendedAuditingPolicyArrayInput` via:
//
//          ServerExtendedAuditingPolicyArray{ ServerExtendedAuditingPolicyArgs{...} }
type ServerExtendedAuditingPolicyArrayInput interface {
	pulumi.Input

	ToServerExtendedAuditingPolicyArrayOutput() ServerExtendedAuditingPolicyArrayOutput
	ToServerExtendedAuditingPolicyArrayOutputWithContext(context.Context) ServerExtendedAuditingPolicyArrayOutput
}

type ServerExtendedAuditingPolicyArray []ServerExtendedAuditingPolicyInput

func (ServerExtendedAuditingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerExtendedAuditingPolicy)(nil)).Elem()
}

func (i ServerExtendedAuditingPolicyArray) ToServerExtendedAuditingPolicyArrayOutput() ServerExtendedAuditingPolicyArrayOutput {
	return i.ToServerExtendedAuditingPolicyArrayOutputWithContext(context.Background())
}

func (i ServerExtendedAuditingPolicyArray) ToServerExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExtendedAuditingPolicyArrayOutput)
}

// ServerExtendedAuditingPolicyMapInput is an input type that accepts ServerExtendedAuditingPolicyMap and ServerExtendedAuditingPolicyMapOutput values.
// You can construct a concrete instance of `ServerExtendedAuditingPolicyMapInput` via:
//
//          ServerExtendedAuditingPolicyMap{ "key": ServerExtendedAuditingPolicyArgs{...} }
type ServerExtendedAuditingPolicyMapInput interface {
	pulumi.Input

	ToServerExtendedAuditingPolicyMapOutput() ServerExtendedAuditingPolicyMapOutput
	ToServerExtendedAuditingPolicyMapOutputWithContext(context.Context) ServerExtendedAuditingPolicyMapOutput
}

type ServerExtendedAuditingPolicyMap map[string]ServerExtendedAuditingPolicyInput

func (ServerExtendedAuditingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerExtendedAuditingPolicy)(nil)).Elem()
}

func (i ServerExtendedAuditingPolicyMap) ToServerExtendedAuditingPolicyMapOutput() ServerExtendedAuditingPolicyMapOutput {
	return i.ToServerExtendedAuditingPolicyMapOutputWithContext(context.Background())
}

func (i ServerExtendedAuditingPolicyMap) ToServerExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExtendedAuditingPolicyMapOutput)
}

type ServerExtendedAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ServerExtendedAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExtendedAuditingPolicy)(nil))
}

func (o ServerExtendedAuditingPolicyOutput) ToServerExtendedAuditingPolicyOutput() ServerExtendedAuditingPolicyOutput {
	return o
}

func (o ServerExtendedAuditingPolicyOutput) ToServerExtendedAuditingPolicyOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyOutput {
	return o
}

func (o ServerExtendedAuditingPolicyOutput) ToServerExtendedAuditingPolicyPtrOutput() ServerExtendedAuditingPolicyPtrOutput {
	return o.ToServerExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (o ServerExtendedAuditingPolicyOutput) ToServerExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerExtendedAuditingPolicy) *ServerExtendedAuditingPolicy {
		return &v
	}).(ServerExtendedAuditingPolicyPtrOutput)
}

type ServerExtendedAuditingPolicyPtrOutput struct{ *pulumi.OutputState }

func (ServerExtendedAuditingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExtendedAuditingPolicy)(nil))
}

func (o ServerExtendedAuditingPolicyPtrOutput) ToServerExtendedAuditingPolicyPtrOutput() ServerExtendedAuditingPolicyPtrOutput {
	return o
}

func (o ServerExtendedAuditingPolicyPtrOutput) ToServerExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyPtrOutput {
	return o
}

func (o ServerExtendedAuditingPolicyPtrOutput) Elem() ServerExtendedAuditingPolicyOutput {
	return o.ApplyT(func(v *ServerExtendedAuditingPolicy) ServerExtendedAuditingPolicy {
		if v != nil {
			return *v
		}
		var ret ServerExtendedAuditingPolicy
		return ret
	}).(ServerExtendedAuditingPolicyOutput)
}

type ServerExtendedAuditingPolicyArrayOutput struct{ *pulumi.OutputState }

func (ServerExtendedAuditingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerExtendedAuditingPolicy)(nil))
}

func (o ServerExtendedAuditingPolicyArrayOutput) ToServerExtendedAuditingPolicyArrayOutput() ServerExtendedAuditingPolicyArrayOutput {
	return o
}

func (o ServerExtendedAuditingPolicyArrayOutput) ToServerExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyArrayOutput {
	return o
}

func (o ServerExtendedAuditingPolicyArrayOutput) Index(i pulumi.IntInput) ServerExtendedAuditingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerExtendedAuditingPolicy {
		return vs[0].([]ServerExtendedAuditingPolicy)[vs[1].(int)]
	}).(ServerExtendedAuditingPolicyOutput)
}

type ServerExtendedAuditingPolicyMapOutput struct{ *pulumi.OutputState }

func (ServerExtendedAuditingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServerExtendedAuditingPolicy)(nil))
}

func (o ServerExtendedAuditingPolicyMapOutput) ToServerExtendedAuditingPolicyMapOutput() ServerExtendedAuditingPolicyMapOutput {
	return o
}

func (o ServerExtendedAuditingPolicyMapOutput) ToServerExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) ServerExtendedAuditingPolicyMapOutput {
	return o
}

func (o ServerExtendedAuditingPolicyMapOutput) MapIndex(k pulumi.StringInput) ServerExtendedAuditingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServerExtendedAuditingPolicy {
		return vs[0].(map[string]ServerExtendedAuditingPolicy)[vs[1].(string)]
	}).(ServerExtendedAuditingPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExtendedAuditingPolicyInput)(nil)).Elem(), &ServerExtendedAuditingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExtendedAuditingPolicyPtrInput)(nil)).Elem(), &ServerExtendedAuditingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExtendedAuditingPolicyArrayInput)(nil)).Elem(), ServerExtendedAuditingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExtendedAuditingPolicyMapInput)(nil)).Elem(), ServerExtendedAuditingPolicyMap{})
	pulumi.RegisterOutputType(ServerExtendedAuditingPolicyOutput{})
	pulumi.RegisterOutputType(ServerExtendedAuditingPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServerExtendedAuditingPolicyArrayOutput{})
	pulumi.RegisterOutputType(ServerExtendedAuditingPolicyMapOutput{})
}
