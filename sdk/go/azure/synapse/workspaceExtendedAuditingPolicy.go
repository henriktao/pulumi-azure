// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Workspace Extended Auditing Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/synapse"
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
// 			AccountKind:            pulumi.String("BlobStorage"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkspace, err := synapse.NewWorkspace(ctx, "exampleWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		auditLogs, err := storage.NewAccount(ctx, "auditLogs", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewWorkspaceExtendedAuditingPolicy(ctx, "exampleWorkspaceExtendedAuditingPolicy", &synapse.WorkspaceExtendedAuditingPolicyArgs{
// 			SynapseWorkspaceId:                 exampleWorkspace.ID(),
// 			StorageEndpoint:                    auditLogs.PrimaryBlobEndpoint,
// 			StorageAccountAccessKey:            auditLogs.PrimaryAccessKey,
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
// Synapse Workspace Extended Auditing Policys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/workspaceExtendedAuditingPolicy:WorkspaceExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/extendedAuditingSettings/default
// ```
type WorkspaceExtendedAuditingPolicy struct {
	pulumi.CustomResourceState

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrOutput `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrOutput `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
}

// NewWorkspaceExtendedAuditingPolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, args *WorkspaceExtendedAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*WorkspaceExtendedAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	var resource WorkspaceExtendedAuditingPolicy
	err := ctx.RegisterResource("azure:synapse/workspaceExtendedAuditingPolicy:WorkspaceExtendedAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceExtendedAuditingPolicy gets an existing WorkspaceExtendedAuditingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceExtendedAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceExtendedAuditingPolicyState, opts ...pulumi.ResourceOption) (*WorkspaceExtendedAuditingPolicy, error) {
	var resource WorkspaceExtendedAuditingPolicy
	err := ctx.ReadResource("azure:synapse/workspaceExtendedAuditingPolicy:WorkspaceExtendedAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceExtendedAuditingPolicy resources.
type workspaceExtendedAuditingPolicyState struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor.
	LogMonitoringEnabled *bool `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
}

type WorkspaceExtendedAuditingPolicyState struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
}

func (WorkspaceExtendedAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceExtendedAuditingPolicyState)(nil)).Elem()
}

type workspaceExtendedAuditingPolicyArgs struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor.
	LogMonitoringEnabled *bool `pulumi:"logMonitoringEnabled"`
	// The number of days to retain logs for in the storage account.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `pulumi:"storageAccountAccessKeyIsSecondary"`
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
}

// The set of arguments for constructing a WorkspaceExtendedAuditingPolicy resource.
type WorkspaceExtendedAuditingPolicyArgs struct {
	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor.
	LogMonitoringEnabled pulumi.BoolPtrInput
	// The number of days to retain logs for in the storage account.
	RetentionInDays pulumi.IntPtrInput
	// The access key to use for the auditing storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Is `storageAccountAccessKey` value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary pulumi.BoolPtrInput
	// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint pulumi.StringPtrInput
	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringInput
}

func (WorkspaceExtendedAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceExtendedAuditingPolicyArgs)(nil)).Elem()
}

type WorkspaceExtendedAuditingPolicyInput interface {
	pulumi.Input

	ToWorkspaceExtendedAuditingPolicyOutput() WorkspaceExtendedAuditingPolicyOutput
	ToWorkspaceExtendedAuditingPolicyOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyOutput
}

func (*WorkspaceExtendedAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceExtendedAuditingPolicy)(nil))
}

func (i *WorkspaceExtendedAuditingPolicy) ToWorkspaceExtendedAuditingPolicyOutput() WorkspaceExtendedAuditingPolicyOutput {
	return i.ToWorkspaceExtendedAuditingPolicyOutputWithContext(context.Background())
}

func (i *WorkspaceExtendedAuditingPolicy) ToWorkspaceExtendedAuditingPolicyOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceExtendedAuditingPolicyOutput)
}

func (i *WorkspaceExtendedAuditingPolicy) ToWorkspaceExtendedAuditingPolicyPtrOutput() WorkspaceExtendedAuditingPolicyPtrOutput {
	return i.ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *WorkspaceExtendedAuditingPolicy) ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceExtendedAuditingPolicyPtrOutput)
}

type WorkspaceExtendedAuditingPolicyPtrInput interface {
	pulumi.Input

	ToWorkspaceExtendedAuditingPolicyPtrOutput() WorkspaceExtendedAuditingPolicyPtrOutput
	ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyPtrOutput
}

type workspaceExtendedAuditingPolicyPtrType WorkspaceExtendedAuditingPolicyArgs

func (*workspaceExtendedAuditingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceExtendedAuditingPolicy)(nil))
}

func (i *workspaceExtendedAuditingPolicyPtrType) ToWorkspaceExtendedAuditingPolicyPtrOutput() WorkspaceExtendedAuditingPolicyPtrOutput {
	return i.ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (i *workspaceExtendedAuditingPolicyPtrType) ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceExtendedAuditingPolicyPtrOutput)
}

// WorkspaceExtendedAuditingPolicyArrayInput is an input type that accepts WorkspaceExtendedAuditingPolicyArray and WorkspaceExtendedAuditingPolicyArrayOutput values.
// You can construct a concrete instance of `WorkspaceExtendedAuditingPolicyArrayInput` via:
//
//          WorkspaceExtendedAuditingPolicyArray{ WorkspaceExtendedAuditingPolicyArgs{...} }
type WorkspaceExtendedAuditingPolicyArrayInput interface {
	pulumi.Input

	ToWorkspaceExtendedAuditingPolicyArrayOutput() WorkspaceExtendedAuditingPolicyArrayOutput
	ToWorkspaceExtendedAuditingPolicyArrayOutputWithContext(context.Context) WorkspaceExtendedAuditingPolicyArrayOutput
}

type WorkspaceExtendedAuditingPolicyArray []WorkspaceExtendedAuditingPolicyInput

func (WorkspaceExtendedAuditingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkspaceExtendedAuditingPolicy)(nil)).Elem()
}

func (i WorkspaceExtendedAuditingPolicyArray) ToWorkspaceExtendedAuditingPolicyArrayOutput() WorkspaceExtendedAuditingPolicyArrayOutput {
	return i.ToWorkspaceExtendedAuditingPolicyArrayOutputWithContext(context.Background())
}

func (i WorkspaceExtendedAuditingPolicyArray) ToWorkspaceExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceExtendedAuditingPolicyArrayOutput)
}

// WorkspaceExtendedAuditingPolicyMapInput is an input type that accepts WorkspaceExtendedAuditingPolicyMap and WorkspaceExtendedAuditingPolicyMapOutput values.
// You can construct a concrete instance of `WorkspaceExtendedAuditingPolicyMapInput` via:
//
//          WorkspaceExtendedAuditingPolicyMap{ "key": WorkspaceExtendedAuditingPolicyArgs{...} }
type WorkspaceExtendedAuditingPolicyMapInput interface {
	pulumi.Input

	ToWorkspaceExtendedAuditingPolicyMapOutput() WorkspaceExtendedAuditingPolicyMapOutput
	ToWorkspaceExtendedAuditingPolicyMapOutputWithContext(context.Context) WorkspaceExtendedAuditingPolicyMapOutput
}

type WorkspaceExtendedAuditingPolicyMap map[string]WorkspaceExtendedAuditingPolicyInput

func (WorkspaceExtendedAuditingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkspaceExtendedAuditingPolicy)(nil)).Elem()
}

func (i WorkspaceExtendedAuditingPolicyMap) ToWorkspaceExtendedAuditingPolicyMapOutput() WorkspaceExtendedAuditingPolicyMapOutput {
	return i.ToWorkspaceExtendedAuditingPolicyMapOutputWithContext(context.Background())
}

func (i WorkspaceExtendedAuditingPolicyMap) ToWorkspaceExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceExtendedAuditingPolicyMapOutput)
}

type WorkspaceExtendedAuditingPolicyOutput struct{ *pulumi.OutputState }

func (WorkspaceExtendedAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceExtendedAuditingPolicy)(nil))
}

func (o WorkspaceExtendedAuditingPolicyOutput) ToWorkspaceExtendedAuditingPolicyOutput() WorkspaceExtendedAuditingPolicyOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyOutput) ToWorkspaceExtendedAuditingPolicyOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyOutput) ToWorkspaceExtendedAuditingPolicyPtrOutput() WorkspaceExtendedAuditingPolicyPtrOutput {
	return o.ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(context.Background())
}

func (o WorkspaceExtendedAuditingPolicyOutput) ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceExtendedAuditingPolicy) *WorkspaceExtendedAuditingPolicy {
		return &v
	}).(WorkspaceExtendedAuditingPolicyPtrOutput)
}

type WorkspaceExtendedAuditingPolicyPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceExtendedAuditingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceExtendedAuditingPolicy)(nil))
}

func (o WorkspaceExtendedAuditingPolicyPtrOutput) ToWorkspaceExtendedAuditingPolicyPtrOutput() WorkspaceExtendedAuditingPolicyPtrOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyPtrOutput) ToWorkspaceExtendedAuditingPolicyPtrOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyPtrOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyPtrOutput) Elem() WorkspaceExtendedAuditingPolicyOutput {
	return o.ApplyT(func(v *WorkspaceExtendedAuditingPolicy) WorkspaceExtendedAuditingPolicy {
		if v != nil {
			return *v
		}
		var ret WorkspaceExtendedAuditingPolicy
		return ret
	}).(WorkspaceExtendedAuditingPolicyOutput)
}

type WorkspaceExtendedAuditingPolicyArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceExtendedAuditingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceExtendedAuditingPolicy)(nil))
}

func (o WorkspaceExtendedAuditingPolicyArrayOutput) ToWorkspaceExtendedAuditingPolicyArrayOutput() WorkspaceExtendedAuditingPolicyArrayOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyArrayOutput) ToWorkspaceExtendedAuditingPolicyArrayOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyArrayOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyArrayOutput) Index(i pulumi.IntInput) WorkspaceExtendedAuditingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkspaceExtendedAuditingPolicy {
		return vs[0].([]WorkspaceExtendedAuditingPolicy)[vs[1].(int)]
	}).(WorkspaceExtendedAuditingPolicyOutput)
}

type WorkspaceExtendedAuditingPolicyMapOutput struct{ *pulumi.OutputState }

func (WorkspaceExtendedAuditingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkspaceExtendedAuditingPolicy)(nil))
}

func (o WorkspaceExtendedAuditingPolicyMapOutput) ToWorkspaceExtendedAuditingPolicyMapOutput() WorkspaceExtendedAuditingPolicyMapOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyMapOutput) ToWorkspaceExtendedAuditingPolicyMapOutputWithContext(ctx context.Context) WorkspaceExtendedAuditingPolicyMapOutput {
	return o
}

func (o WorkspaceExtendedAuditingPolicyMapOutput) MapIndex(k pulumi.StringInput) WorkspaceExtendedAuditingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkspaceExtendedAuditingPolicy {
		return vs[0].(map[string]WorkspaceExtendedAuditingPolicy)[vs[1].(string)]
	}).(WorkspaceExtendedAuditingPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceExtendedAuditingPolicyOutput{})
	pulumi.RegisterOutputType(WorkspaceExtendedAuditingPolicyPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceExtendedAuditingPolicyArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceExtendedAuditingPolicyMapOutput{})
}
