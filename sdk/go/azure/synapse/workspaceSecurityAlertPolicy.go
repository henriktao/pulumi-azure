// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Security Alert Policy for a Synapse Workspace.
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
// 			AccountKind:            pulumi.String("StorageV2"),
// 			IsHnsEnabled:           pulumi.Bool(true),
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
// 			AadAdmin: &synapse.WorkspaceAadAdminArgs{
// 				Login:    pulumi.String("AzureAD Admin"),
// 				ObjectId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 				TenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Env": pulumi.String("production"),
// 			},
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
// 		_, err = synapse.NewWorkspaceSecurityAlertPolicy(ctx, "exampleWorkspaceSecurityAlertPolicy", &synapse.WorkspaceSecurityAlertPolicyArgs{
// 			SynapseWorkspaceId:      exampleWorkspace.ID(),
// 			PolicyState:             pulumi.String("Enabled"),
// 			StorageEndpoint:         auditLogs.PrimaryBlobEndpoint,
// 			StorageAccountAccessKey: auditLogs.PrimaryAccessKey,
// 			DisabledAlerts: pulumi.StringArray{
// 				pulumi.String("Sql_Injection"),
// 				pulumi.String("Data_Exfiltration"),
// 			},
// 			RetentionDays: pulumi.Int(20),
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
// Synapse Workspace Security Alert Policies can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/workspaceSecurityAlertPolicy:WorkspaceSecurityAlertPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/securityAlertPolicies/Default
// ```
type WorkspaceSecurityAlertPolicy struct {
	pulumi.CustomResourceState

	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayOutput `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdminsEnabled pulumi.BoolPtrOutput `pulumi:"emailAccountAdminsEnabled"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Allowed values are: `Disabled`, `Enabled`.
	PolicyState pulumi.StringOutput `pulumi:"policyState"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrOutput `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringOutput `pulumi:"synapseWorkspaceId"`
}

// NewWorkspaceSecurityAlertPolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *WorkspaceSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*WorkspaceSecurityAlertPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyState == nil {
		return nil, errors.New("invalid value for required argument 'PolicyState'")
	}
	if args.SynapseWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceId'")
	}
	var resource WorkspaceSecurityAlertPolicy
	err := ctx.RegisterResource("azure:synapse/workspaceSecurityAlertPolicy:WorkspaceSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceSecurityAlertPolicy gets an existing WorkspaceSecurityAlertPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*WorkspaceSecurityAlertPolicy, error) {
	var resource WorkspaceSecurityAlertPolicy
	err := ctx.ReadResource("azure:synapse/workspaceSecurityAlertPolicy:WorkspaceSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceSecurityAlertPolicy resources.
type workspaceSecurityAlertPolicyState struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdminsEnabled *bool `pulumi:"emailAccountAdminsEnabled"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Allowed values are: `Disabled`, `Enabled`.
	PolicyState *string `pulumi:"policyState"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
}

type WorkspaceSecurityAlertPolicyState struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayInput
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdminsEnabled pulumi.BoolPtrInput
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Allowed values are: `Disabled`, `Enabled`.
	PolicyState pulumi.StringPtrInput
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
}

func (WorkspaceSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSecurityAlertPolicyState)(nil)).Elem()
}

type workspaceSecurityAlertPolicyArgs struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts []string `pulumi:"disabledAlerts"`
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdminsEnabled *bool `pulumi:"emailAccountAdminsEnabled"`
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Allowed values are: `Disabled`, `Enabled`.
	PolicyState string `pulumi:"policyState"`
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceId string `pulumi:"synapseWorkspaceId"`
}

// The set of arguments for constructing a WorkspaceSecurityAlertPolicy resource.
type WorkspaceSecurityAlertPolicyArgs struct {
	// Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
	DisabledAlerts pulumi.StringArrayInput
	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
	EmailAccountAdminsEnabled pulumi.BoolPtrInput
	// Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses pulumi.StringArrayInput
	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific workspace. Allowed values are: `Disabled`, `Enabled`.
	PolicyState pulumi.StringInput
	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
	RetentionDays pulumi.IntPtrInput
	// Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	StorageEndpoint pulumi.StringPtrInput
	// Specifies the ID of the Synapse Workspace. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringInput
}

func (WorkspaceSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSecurityAlertPolicyArgs)(nil)).Elem()
}

type WorkspaceSecurityAlertPolicyInput interface {
	pulumi.Input

	ToWorkspaceSecurityAlertPolicyOutput() WorkspaceSecurityAlertPolicyOutput
	ToWorkspaceSecurityAlertPolicyOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyOutput
}

func (*WorkspaceSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSecurityAlertPolicy)(nil))
}

func (i *WorkspaceSecurityAlertPolicy) ToWorkspaceSecurityAlertPolicyOutput() WorkspaceSecurityAlertPolicyOutput {
	return i.ToWorkspaceSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *WorkspaceSecurityAlertPolicy) ToWorkspaceSecurityAlertPolicyOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSecurityAlertPolicyOutput)
}

func (i *WorkspaceSecurityAlertPolicy) ToWorkspaceSecurityAlertPolicyPtrOutput() WorkspaceSecurityAlertPolicyPtrOutput {
	return i.ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (i *WorkspaceSecurityAlertPolicy) ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSecurityAlertPolicyPtrOutput)
}

type WorkspaceSecurityAlertPolicyPtrInput interface {
	pulumi.Input

	ToWorkspaceSecurityAlertPolicyPtrOutput() WorkspaceSecurityAlertPolicyPtrOutput
	ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyPtrOutput
}

type workspaceSecurityAlertPolicyPtrType WorkspaceSecurityAlertPolicyArgs

func (*workspaceSecurityAlertPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSecurityAlertPolicy)(nil))
}

func (i *workspaceSecurityAlertPolicyPtrType) ToWorkspaceSecurityAlertPolicyPtrOutput() WorkspaceSecurityAlertPolicyPtrOutput {
	return i.ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (i *workspaceSecurityAlertPolicyPtrType) ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSecurityAlertPolicyPtrOutput)
}

// WorkspaceSecurityAlertPolicyArrayInput is an input type that accepts WorkspaceSecurityAlertPolicyArray and WorkspaceSecurityAlertPolicyArrayOutput values.
// You can construct a concrete instance of `WorkspaceSecurityAlertPolicyArrayInput` via:
//
//          WorkspaceSecurityAlertPolicyArray{ WorkspaceSecurityAlertPolicyArgs{...} }
type WorkspaceSecurityAlertPolicyArrayInput interface {
	pulumi.Input

	ToWorkspaceSecurityAlertPolicyArrayOutput() WorkspaceSecurityAlertPolicyArrayOutput
	ToWorkspaceSecurityAlertPolicyArrayOutputWithContext(context.Context) WorkspaceSecurityAlertPolicyArrayOutput
}

type WorkspaceSecurityAlertPolicyArray []WorkspaceSecurityAlertPolicyInput

func (WorkspaceSecurityAlertPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkspaceSecurityAlertPolicy)(nil)).Elem()
}

func (i WorkspaceSecurityAlertPolicyArray) ToWorkspaceSecurityAlertPolicyArrayOutput() WorkspaceSecurityAlertPolicyArrayOutput {
	return i.ToWorkspaceSecurityAlertPolicyArrayOutputWithContext(context.Background())
}

func (i WorkspaceSecurityAlertPolicyArray) ToWorkspaceSecurityAlertPolicyArrayOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSecurityAlertPolicyArrayOutput)
}

// WorkspaceSecurityAlertPolicyMapInput is an input type that accepts WorkspaceSecurityAlertPolicyMap and WorkspaceSecurityAlertPolicyMapOutput values.
// You can construct a concrete instance of `WorkspaceSecurityAlertPolicyMapInput` via:
//
//          WorkspaceSecurityAlertPolicyMap{ "key": WorkspaceSecurityAlertPolicyArgs{...} }
type WorkspaceSecurityAlertPolicyMapInput interface {
	pulumi.Input

	ToWorkspaceSecurityAlertPolicyMapOutput() WorkspaceSecurityAlertPolicyMapOutput
	ToWorkspaceSecurityAlertPolicyMapOutputWithContext(context.Context) WorkspaceSecurityAlertPolicyMapOutput
}

type WorkspaceSecurityAlertPolicyMap map[string]WorkspaceSecurityAlertPolicyInput

func (WorkspaceSecurityAlertPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkspaceSecurityAlertPolicy)(nil)).Elem()
}

func (i WorkspaceSecurityAlertPolicyMap) ToWorkspaceSecurityAlertPolicyMapOutput() WorkspaceSecurityAlertPolicyMapOutput {
	return i.ToWorkspaceSecurityAlertPolicyMapOutputWithContext(context.Background())
}

func (i WorkspaceSecurityAlertPolicyMap) ToWorkspaceSecurityAlertPolicyMapOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSecurityAlertPolicyMapOutput)
}

type WorkspaceSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (WorkspaceSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSecurityAlertPolicy)(nil))
}

func (o WorkspaceSecurityAlertPolicyOutput) ToWorkspaceSecurityAlertPolicyOutput() WorkspaceSecurityAlertPolicyOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyOutput) ToWorkspaceSecurityAlertPolicyOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyOutput) ToWorkspaceSecurityAlertPolicyPtrOutput() WorkspaceSecurityAlertPolicyPtrOutput {
	return o.ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(context.Background())
}

func (o WorkspaceSecurityAlertPolicyOutput) ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSecurityAlertPolicy) *WorkspaceSecurityAlertPolicy {
		return &v
	}).(WorkspaceSecurityAlertPolicyPtrOutput)
}

type WorkspaceSecurityAlertPolicyPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSecurityAlertPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSecurityAlertPolicy)(nil))
}

func (o WorkspaceSecurityAlertPolicyPtrOutput) ToWorkspaceSecurityAlertPolicyPtrOutput() WorkspaceSecurityAlertPolicyPtrOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyPtrOutput) ToWorkspaceSecurityAlertPolicyPtrOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyPtrOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyPtrOutput) Elem() WorkspaceSecurityAlertPolicyOutput {
	return o.ApplyT(func(v *WorkspaceSecurityAlertPolicy) WorkspaceSecurityAlertPolicy {
		if v != nil {
			return *v
		}
		var ret WorkspaceSecurityAlertPolicy
		return ret
	}).(WorkspaceSecurityAlertPolicyOutput)
}

type WorkspaceSecurityAlertPolicyArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceSecurityAlertPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceSecurityAlertPolicy)(nil))
}

func (o WorkspaceSecurityAlertPolicyArrayOutput) ToWorkspaceSecurityAlertPolicyArrayOutput() WorkspaceSecurityAlertPolicyArrayOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyArrayOutput) ToWorkspaceSecurityAlertPolicyArrayOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyArrayOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyArrayOutput) Index(i pulumi.IntInput) WorkspaceSecurityAlertPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkspaceSecurityAlertPolicy {
		return vs[0].([]WorkspaceSecurityAlertPolicy)[vs[1].(int)]
	}).(WorkspaceSecurityAlertPolicyOutput)
}

type WorkspaceSecurityAlertPolicyMapOutput struct{ *pulumi.OutputState }

func (WorkspaceSecurityAlertPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkspaceSecurityAlertPolicy)(nil))
}

func (o WorkspaceSecurityAlertPolicyMapOutput) ToWorkspaceSecurityAlertPolicyMapOutput() WorkspaceSecurityAlertPolicyMapOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyMapOutput) ToWorkspaceSecurityAlertPolicyMapOutputWithContext(ctx context.Context) WorkspaceSecurityAlertPolicyMapOutput {
	return o
}

func (o WorkspaceSecurityAlertPolicyMapOutput) MapIndex(k pulumi.StringInput) WorkspaceSecurityAlertPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkspaceSecurityAlertPolicy {
		return vs[0].(map[string]WorkspaceSecurityAlertPolicy)[vs[1].(string)]
	}).(WorkspaceSecurityAlertPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceSecurityAlertPolicyOutput{})
	pulumi.RegisterOutputType(WorkspaceSecurityAlertPolicyPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSecurityAlertPolicyArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceSecurityAlertPolicyMapOutput{})
}
