// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Machine Learning Synapse Spark.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/keyvault"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/machinelearning"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/synapse"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("west europe"),
// 			Tags: pulumi.StringMap{
// 				"stage": pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			TenantId:               pulumi.String(current.TenantId),
// 			SkuName:                pulumi.String("standard"),
// 			PurgeProtectionEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			Location:               exampleResourceGroup.Location,
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWorkspace, err := machinelearning.NewWorkspace(ctx, "exampleWorkspace", &machinelearning.WorkspaceArgs{
// 			Location:              exampleResourceGroup.Location,
// 			ResourceGroupName:     exampleResourceGroup.Name,
// 			ApplicationInsightsId: exampleInsights.ID(),
// 			KeyVaultId:            exampleKeyVault.ID(),
// 			StorageAccountId:      exampleAccount.ID(),
// 			Identity: &machinelearning.WorkspaceIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
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
// 		_, err = synapse.NewWorkspace(ctx, "exampleSynapse_workspaceWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSparkPool, err := synapse.NewSparkPool(ctx, "exampleSparkPool", &synapse.SparkPoolArgs{
// 			SynapseWorkspaceId: exampleSynapse / workspaceWorkspace.Id,
// 			NodeSizeFamily:     pulumi.String("MemoryOptimized"),
// 			NodeSize:           pulumi.String("Small"),
// 			NodeCount:          pulumi.Int(3),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = machinelearning.NewSynapseSpark(ctx, "exampleSynapseSpark", &machinelearning.SynapseSparkArgs{
// 			MachineLearningWorkspaceId: exampleWorkspace.ID(),
// 			Location:                   exampleResourceGroup.Location,
// 			SynapseSparkPoolId:         exampleSparkPool.ID(),
// 			Identity: &machinelearning.SynapseSparkIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
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
// Machine Learning Synapse Sparks can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:machinelearning/synapseSpark:SynapseSpark example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
// ```
type SynapseSpark struct {
	pulumi.CustomResourceState

	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity SynapseSparkIdentityPtrOutput `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled pulumi.BoolPtrOutput `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceId pulumi.StringOutput `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolId pulumi.StringOutput `pulumi:"synapseSparkPoolId"`
	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSynapseSpark registers a new resource with the given unique name, arguments, and options.
func NewSynapseSpark(ctx *pulumi.Context,
	name string, args *SynapseSparkArgs, opts ...pulumi.ResourceOption) (*SynapseSpark, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineLearningWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'MachineLearningWorkspaceId'")
	}
	if args.SynapseSparkPoolId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseSparkPoolId'")
	}
	var resource SynapseSpark
	err := ctx.RegisterResource("azure:machinelearning/synapseSpark:SynapseSpark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynapseSpark gets an existing SynapseSpark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynapseSpark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynapseSparkState, opts ...pulumi.ResourceOption) (*SynapseSpark, error) {
	var resource SynapseSpark
	err := ctx.ReadResource("azure:machinelearning/synapseSpark:SynapseSpark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynapseSpark resources.
type synapseSparkState struct {
	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description *string `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity *SynapseSparkIdentity `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceId *string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Name *string `pulumi:"name"`
	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolId *string `pulumi:"synapseSparkPoolId"`
	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Tags map[string]string `pulumi:"tags"`
}

type SynapseSparkState struct {
	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity SynapseSparkIdentityPtrInput
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled pulumi.BoolPtrInput
	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Name pulumi.StringPtrInput
	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Tags pulumi.StringMapInput
}

func (SynapseSparkState) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseSparkState)(nil)).Elem()
}

type synapseSparkArgs struct {
	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description *string `pulumi:"description"`
	// A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity *SynapseSparkIdentity `pulumi:"identity"`
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled *bool `pulumi:"localAuthEnabled"`
	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location *string `pulumi:"location"`
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceId string `pulumi:"machineLearningWorkspaceId"`
	// The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Name *string `pulumi:"name"`
	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolId string `pulumi:"synapseSparkPoolId"`
	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SynapseSpark resource.
type SynapseSparkArgs struct {
	// The description of the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Description pulumi.StringPtrInput
	// A `identity` block as defined below. Changing this forces a new Machine Learning Synapse Spark to be created.
	Identity SynapseSparkIdentityPtrInput
	// Whether local authentication methods is enabled. Defaults to `true`. Changing this forces a new Machine Learning Synapse Spark to be created.
	LocalAuthEnabled pulumi.BoolPtrInput
	// The Azure Region where the Machine Learning Synapse Spark should exist. Changing this forces a new Machine Learning Synapse Spark to be created.
	Location pulumi.StringPtrInput
	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Synapse Spark to be created.
	MachineLearningWorkspaceId pulumi.StringInput
	// The name which should be used for this Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Name pulumi.StringPtrInput
	// The ID of the linked Synapse Spark Pool. Changing this forces a new Machine Learning Synapse Spark to be created.
	SynapseSparkPoolId pulumi.StringInput
	// A mapping of tags which should be assigned to the Machine Learning Synapse Spark. Changing this forces a new Machine Learning Synapse Spark to be created.
	Tags pulumi.StringMapInput
}

func (SynapseSparkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseSparkArgs)(nil)).Elem()
}

type SynapseSparkInput interface {
	pulumi.Input

	ToSynapseSparkOutput() SynapseSparkOutput
	ToSynapseSparkOutputWithContext(ctx context.Context) SynapseSparkOutput
}

func (*SynapseSpark) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSpark)(nil))
}

func (i *SynapseSpark) ToSynapseSparkOutput() SynapseSparkOutput {
	return i.ToSynapseSparkOutputWithContext(context.Background())
}

func (i *SynapseSpark) ToSynapseSparkOutputWithContext(ctx context.Context) SynapseSparkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkOutput)
}

func (i *SynapseSpark) ToSynapseSparkPtrOutput() SynapseSparkPtrOutput {
	return i.ToSynapseSparkPtrOutputWithContext(context.Background())
}

func (i *SynapseSpark) ToSynapseSparkPtrOutputWithContext(ctx context.Context) SynapseSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkPtrOutput)
}

type SynapseSparkPtrInput interface {
	pulumi.Input

	ToSynapseSparkPtrOutput() SynapseSparkPtrOutput
	ToSynapseSparkPtrOutputWithContext(ctx context.Context) SynapseSparkPtrOutput
}

type synapseSparkPtrType SynapseSparkArgs

func (*synapseSparkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSpark)(nil))
}

func (i *synapseSparkPtrType) ToSynapseSparkPtrOutput() SynapseSparkPtrOutput {
	return i.ToSynapseSparkPtrOutputWithContext(context.Background())
}

func (i *synapseSparkPtrType) ToSynapseSparkPtrOutputWithContext(ctx context.Context) SynapseSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkPtrOutput)
}

// SynapseSparkArrayInput is an input type that accepts SynapseSparkArray and SynapseSparkArrayOutput values.
// You can construct a concrete instance of `SynapseSparkArrayInput` via:
//
//          SynapseSparkArray{ SynapseSparkArgs{...} }
type SynapseSparkArrayInput interface {
	pulumi.Input

	ToSynapseSparkArrayOutput() SynapseSparkArrayOutput
	ToSynapseSparkArrayOutputWithContext(context.Context) SynapseSparkArrayOutput
}

type SynapseSparkArray []SynapseSparkInput

func (SynapseSparkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynapseSpark)(nil)).Elem()
}

func (i SynapseSparkArray) ToSynapseSparkArrayOutput() SynapseSparkArrayOutput {
	return i.ToSynapseSparkArrayOutputWithContext(context.Background())
}

func (i SynapseSparkArray) ToSynapseSparkArrayOutputWithContext(ctx context.Context) SynapseSparkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkArrayOutput)
}

// SynapseSparkMapInput is an input type that accepts SynapseSparkMap and SynapseSparkMapOutput values.
// You can construct a concrete instance of `SynapseSparkMapInput` via:
//
//          SynapseSparkMap{ "key": SynapseSparkArgs{...} }
type SynapseSparkMapInput interface {
	pulumi.Input

	ToSynapseSparkMapOutput() SynapseSparkMapOutput
	ToSynapseSparkMapOutputWithContext(context.Context) SynapseSparkMapOutput
}

type SynapseSparkMap map[string]SynapseSparkInput

func (SynapseSparkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynapseSpark)(nil)).Elem()
}

func (i SynapseSparkMap) ToSynapseSparkMapOutput() SynapseSparkMapOutput {
	return i.ToSynapseSparkMapOutputWithContext(context.Background())
}

func (i SynapseSparkMap) ToSynapseSparkMapOutputWithContext(ctx context.Context) SynapseSparkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkMapOutput)
}

type SynapseSparkOutput struct{ *pulumi.OutputState }

func (SynapseSparkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSpark)(nil))
}

func (o SynapseSparkOutput) ToSynapseSparkOutput() SynapseSparkOutput {
	return o
}

func (o SynapseSparkOutput) ToSynapseSparkOutputWithContext(ctx context.Context) SynapseSparkOutput {
	return o
}

func (o SynapseSparkOutput) ToSynapseSparkPtrOutput() SynapseSparkPtrOutput {
	return o.ToSynapseSparkPtrOutputWithContext(context.Background())
}

func (o SynapseSparkOutput) ToSynapseSparkPtrOutputWithContext(ctx context.Context) SynapseSparkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SynapseSpark) *SynapseSpark {
		return &v
	}).(SynapseSparkPtrOutput)
}

type SynapseSparkPtrOutput struct{ *pulumi.OutputState }

func (SynapseSparkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSpark)(nil))
}

func (o SynapseSparkPtrOutput) ToSynapseSparkPtrOutput() SynapseSparkPtrOutput {
	return o
}

func (o SynapseSparkPtrOutput) ToSynapseSparkPtrOutputWithContext(ctx context.Context) SynapseSparkPtrOutput {
	return o
}

func (o SynapseSparkPtrOutput) Elem() SynapseSparkOutput {
	return o.ApplyT(func(v *SynapseSpark) SynapseSpark {
		if v != nil {
			return *v
		}
		var ret SynapseSpark
		return ret
	}).(SynapseSparkOutput)
}

type SynapseSparkArrayOutput struct{ *pulumi.OutputState }

func (SynapseSparkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SynapseSpark)(nil))
}

func (o SynapseSparkArrayOutput) ToSynapseSparkArrayOutput() SynapseSparkArrayOutput {
	return o
}

func (o SynapseSparkArrayOutput) ToSynapseSparkArrayOutputWithContext(ctx context.Context) SynapseSparkArrayOutput {
	return o
}

func (o SynapseSparkArrayOutput) Index(i pulumi.IntInput) SynapseSparkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SynapseSpark {
		return vs[0].([]SynapseSpark)[vs[1].(int)]
	}).(SynapseSparkOutput)
}

type SynapseSparkMapOutput struct{ *pulumi.OutputState }

func (SynapseSparkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SynapseSpark)(nil))
}

func (o SynapseSparkMapOutput) ToSynapseSparkMapOutput() SynapseSparkMapOutput {
	return o
}

func (o SynapseSparkMapOutput) ToSynapseSparkMapOutputWithContext(ctx context.Context) SynapseSparkMapOutput {
	return o
}

func (o SynapseSparkMapOutput) MapIndex(k pulumi.StringInput) SynapseSparkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SynapseSpark {
		return vs[0].(map[string]SynapseSpark)[vs[1].(string)]
	}).(SynapseSparkOutput)
}

func init() {
	pulumi.RegisterOutputType(SynapseSparkOutput{})
	pulumi.RegisterOutputType(SynapseSparkPtrOutput{})
	pulumi.RegisterOutputType(SynapseSparkArrayOutput{})
	pulumi.RegisterOutputType(SynapseSparkMapOutput{})
}
