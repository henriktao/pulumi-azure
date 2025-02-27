// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Synapse Role Assignment.
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
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFirewallRule, err := synapse.NewFirewallRule(ctx, "exampleFirewallRule", &synapse.FirewallRuleArgs{
// 			SynapseWorkspaceId: pulumi.Any(azurerm_synapse_workspace.Test.Id),
// 			StartIpAddress:     pulumi.String("0.0.0.0"),
// 			EndIpAddress:       pulumi.String("255.255.255.255"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewRoleAssignment(ctx, "exampleRoleAssignment", &synapse.RoleAssignmentArgs{
// 			SynapseWorkspaceId: exampleWorkspace.ID(),
// 			RoleName:           pulumi.String("Synapse SQL Administrator"),
// 			PrincipalId:        pulumi.String(current.ObjectId),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleFirewallRule,
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
// Synapse Role Assignment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/roleAssignment:RoleAssignment example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1|000000000000"
// ```
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The ID of the Principal (User, Group or Service Principal) to assign the Synapse Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The Role Name of the Synapse Built-In Role. Changing this forces a new resource to be created.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The Synapse Spark Pool which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseSparkPoolId pulumi.StringPtrOutput `pulumi:"synapseSparkPoolId"`
	// The Synapse Workspace which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrOutput `pulumi:"synapseWorkspaceId"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	var resource RoleAssignment
	err := ctx.RegisterResource("azure:synapse/roleAssignment:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure:synapse/roleAssignment:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
	// The ID of the Principal (User, Group or Service Principal) to assign the Synapse Role Definition to. Changing this forces a new resource to be created.
	PrincipalId *string `pulumi:"principalId"`
	// The Role Name of the Synapse Built-In Role. Changing this forces a new resource to be created.
	RoleName *string `pulumi:"roleName"`
	// The Synapse Spark Pool which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseSparkPoolId *string `pulumi:"synapseSparkPoolId"`
	// The Synapse Workspace which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
}

type RoleAssignmentState struct {
	// The ID of the Principal (User, Group or Service Principal) to assign the Synapse Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringPtrInput
	// The Role Name of the Synapse Built-In Role. Changing this forces a new resource to be created.
	RoleName pulumi.StringPtrInput
	// The Synapse Spark Pool which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseSparkPoolId pulumi.StringPtrInput
	// The Synapse Workspace which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// The ID of the Principal (User, Group or Service Principal) to assign the Synapse Role Definition to. Changing this forces a new resource to be created.
	PrincipalId string `pulumi:"principalId"`
	// The Role Name of the Synapse Built-In Role. Changing this forces a new resource to be created.
	RoleName string `pulumi:"roleName"`
	// The Synapse Spark Pool which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseSparkPoolId *string `pulumi:"synapseSparkPoolId"`
	// The Synapse Workspace which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseWorkspaceId *string `pulumi:"synapseWorkspaceId"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The ID of the Principal (User, Group or Service Principal) to assign the Synapse Role Definition to. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringInput
	// The Role Name of the Synapse Built-In Role. Changing this forces a new resource to be created.
	RoleName pulumi.StringInput
	// The Synapse Spark Pool which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseSparkPoolId pulumi.StringPtrInput
	// The Synapse Workspace which the Synapse Role Assignment applies to. Changing this forces a new resource to be created.
	SynapseWorkspaceId pulumi.StringPtrInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

// RoleAssignmentArrayInput is an input type that accepts RoleAssignmentArray and RoleAssignmentArrayOutput values.
// You can construct a concrete instance of `RoleAssignmentArrayInput` via:
//
//          RoleAssignmentArray{ RoleAssignmentArgs{...} }
type RoleAssignmentArrayInput interface {
	pulumi.Input

	ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput
	ToRoleAssignmentArrayOutputWithContext(context.Context) RoleAssignmentArrayOutput
}

type RoleAssignmentArray []RoleAssignmentInput

func (RoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return i.ToRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentArrayOutput)
}

// RoleAssignmentMapInput is an input type that accepts RoleAssignmentMap and RoleAssignmentMapOutput values.
// You can construct a concrete instance of `RoleAssignmentMapInput` via:
//
//          RoleAssignmentMap{ "key": RoleAssignmentArgs{...} }
type RoleAssignmentMapInput interface {
	pulumi.Input

	ToRoleAssignmentMapOutput() RoleAssignmentMapOutput
	ToRoleAssignmentMapOutputWithContext(context.Context) RoleAssignmentMapOutput
}

type RoleAssignmentMap map[string]RoleAssignmentInput

func (RoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return i.ToRoleAssignmentMapOutputWithContext(context.Background())
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentMapOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

type RoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (RoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) Index(i pulumi.IntInput) RoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleAssignment {
		return vs[0].([]*RoleAssignment)[vs[1].(int)]
	}).(RoleAssignmentOutput)
}

type RoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (RoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) RoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleAssignment {
		return vs[0].(map[string]*RoleAssignment)[vs[1].(string)]
	}).(RoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentInput)(nil)).Elem(), &RoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentArrayInput)(nil)).Elem(), RoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentMapInput)(nil)).Elem(), RoleAssignmentMap{})
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
	pulumi.RegisterOutputType(RoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(RoleAssignmentMapOutput{})
}
