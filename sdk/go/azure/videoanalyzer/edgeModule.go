// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Video Analyzer Edge Module.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/authorization"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/videoanalyzer"
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
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleUserAssignedIdentity, err := authorization.NewUserAssignedIdentity(ctx, "exampleUserAssignedIdentity", &authorization.UserAssignedIdentityArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		contributor, err := authorization.NewAssignment(ctx, "contributor", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Storage Blob Data Contributor"),
// 			PrincipalId:        exampleUserAssignedIdentity.PrincipalId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		reader, err := authorization.NewAssignment(ctx, "reader", &authorization.AssignmentArgs{
// 			Scope:              exampleAccount.ID(),
// 			RoleDefinitionName: pulumi.String("Reader"),
// 			PrincipalId:        exampleUserAssignedIdentity.PrincipalId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAnalyzer, err := videoanalyzer.NewAnalyzer(ctx, "exampleAnalyzer", &videoanalyzer.AnalyzerArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StorageAccount: &videoanalyzer.AnalyzerStorageAccountArgs{
// 				Id:                     exampleAccount.ID(),
// 				UserAssignedIdentityId: exampleUserAssignedIdentity.ID(),
// 			},
// 			Identity: &videoanalyzer.AnalyzerIdentityArgs{
// 				Type: pulumi.String("UserAssigned"),
// 				IdentityIds: pulumi.StringArray{
// 					exampleUserAssignedIdentity.ID(),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("staging"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleUserAssignedIdentity,
// 			contributor,
// 			reader,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = videoanalyzer.NewEdgeModule(ctx, "exampleEdgeModule", &videoanalyzer.EdgeModuleArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			VideoAnalyzerName: exampleAnalyzer.Name,
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
// Video Analyzer Edge Module can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:videoanalyzer/edgeModule:EdgeModule edge /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1/edgeModules/edge1
// ```
type EdgeModule struct {
	pulumi.CustomResourceState

	// Specifies the name of the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The name of the Video Analyzer in which to create the Edge Module. Changing this forces a new resource to be created.
	VideoAnalyzerName pulumi.StringOutput `pulumi:"videoAnalyzerName"`
}

// NewEdgeModule registers a new resource with the given unique name, arguments, and options.
func NewEdgeModule(ctx *pulumi.Context,
	name string, args *EdgeModuleArgs, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VideoAnalyzerName == nil {
		return nil, errors.New("invalid value for required argument 'VideoAnalyzerName'")
	}
	var resource EdgeModule
	err := ctx.RegisterResource("azure:videoanalyzer/edgeModule:EdgeModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeModule gets an existing EdgeModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeModuleState, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
	var resource EdgeModule
	err := ctx.ReadResource("azure:videoanalyzer/edgeModule:EdgeModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeModule resources.
type edgeModuleState struct {
	// Specifies the name of the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The name of the Video Analyzer in which to create the Edge Module. Changing this forces a new resource to be created.
	VideoAnalyzerName *string `pulumi:"videoAnalyzerName"`
}

type EdgeModuleState struct {
	// Specifies the name of the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The name of the Video Analyzer in which to create the Edge Module. Changing this forces a new resource to be created.
	VideoAnalyzerName pulumi.StringPtrInput
}

func (EdgeModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleState)(nil)).Elem()
}

type edgeModuleArgs struct {
	// Specifies the name of the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Video Analyzer in which to create the Edge Module. Changing this forces a new resource to be created.
	VideoAnalyzerName string `pulumi:"videoAnalyzerName"`
}

// The set of arguments for constructing a EdgeModule resource.
type EdgeModuleArgs struct {
	// Specifies the name of the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Video Analyzer Edge Module. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The name of the Video Analyzer in which to create the Edge Module. Changing this forces a new resource to be created.
	VideoAnalyzerName pulumi.StringInput
}

func (EdgeModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleArgs)(nil)).Elem()
}

type EdgeModuleInput interface {
	pulumi.Input

	ToEdgeModuleOutput() EdgeModuleOutput
	ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput
}

func (*EdgeModule) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeModule)(nil))
}

func (i *EdgeModule) ToEdgeModuleOutput() EdgeModuleOutput {
	return i.ToEdgeModuleOutputWithContext(context.Background())
}

func (i *EdgeModule) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModuleOutput)
}

func (i *EdgeModule) ToEdgeModulePtrOutput() EdgeModulePtrOutput {
	return i.ToEdgeModulePtrOutputWithContext(context.Background())
}

func (i *EdgeModule) ToEdgeModulePtrOutputWithContext(ctx context.Context) EdgeModulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModulePtrOutput)
}

type EdgeModulePtrInput interface {
	pulumi.Input

	ToEdgeModulePtrOutput() EdgeModulePtrOutput
	ToEdgeModulePtrOutputWithContext(ctx context.Context) EdgeModulePtrOutput
}

type edgeModulePtrType EdgeModuleArgs

func (*edgeModulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeModule)(nil))
}

func (i *edgeModulePtrType) ToEdgeModulePtrOutput() EdgeModulePtrOutput {
	return i.ToEdgeModulePtrOutputWithContext(context.Background())
}

func (i *edgeModulePtrType) ToEdgeModulePtrOutputWithContext(ctx context.Context) EdgeModulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModulePtrOutput)
}

// EdgeModuleArrayInput is an input type that accepts EdgeModuleArray and EdgeModuleArrayOutput values.
// You can construct a concrete instance of `EdgeModuleArrayInput` via:
//
//          EdgeModuleArray{ EdgeModuleArgs{...} }
type EdgeModuleArrayInput interface {
	pulumi.Input

	ToEdgeModuleArrayOutput() EdgeModuleArrayOutput
	ToEdgeModuleArrayOutputWithContext(context.Context) EdgeModuleArrayOutput
}

type EdgeModuleArray []EdgeModuleInput

func (EdgeModuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeModule)(nil)).Elem()
}

func (i EdgeModuleArray) ToEdgeModuleArrayOutput() EdgeModuleArrayOutput {
	return i.ToEdgeModuleArrayOutputWithContext(context.Background())
}

func (i EdgeModuleArray) ToEdgeModuleArrayOutputWithContext(ctx context.Context) EdgeModuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModuleArrayOutput)
}

// EdgeModuleMapInput is an input type that accepts EdgeModuleMap and EdgeModuleMapOutput values.
// You can construct a concrete instance of `EdgeModuleMapInput` via:
//
//          EdgeModuleMap{ "key": EdgeModuleArgs{...} }
type EdgeModuleMapInput interface {
	pulumi.Input

	ToEdgeModuleMapOutput() EdgeModuleMapOutput
	ToEdgeModuleMapOutputWithContext(context.Context) EdgeModuleMapOutput
}

type EdgeModuleMap map[string]EdgeModuleInput

func (EdgeModuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeModule)(nil)).Elem()
}

func (i EdgeModuleMap) ToEdgeModuleMapOutput() EdgeModuleMapOutput {
	return i.ToEdgeModuleMapOutputWithContext(context.Background())
}

func (i EdgeModuleMap) ToEdgeModuleMapOutputWithContext(ctx context.Context) EdgeModuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModuleMapOutput)
}

type EdgeModuleOutput struct{ *pulumi.OutputState }

func (EdgeModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeModule)(nil))
}

func (o EdgeModuleOutput) ToEdgeModuleOutput() EdgeModuleOutput {
	return o
}

func (o EdgeModuleOutput) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return o
}

func (o EdgeModuleOutput) ToEdgeModulePtrOutput() EdgeModulePtrOutput {
	return o.ToEdgeModulePtrOutputWithContext(context.Background())
}

func (o EdgeModuleOutput) ToEdgeModulePtrOutputWithContext(ctx context.Context) EdgeModulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeModule) *EdgeModule {
		return &v
	}).(EdgeModulePtrOutput)
}

type EdgeModulePtrOutput struct{ *pulumi.OutputState }

func (EdgeModulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeModule)(nil))
}

func (o EdgeModulePtrOutput) ToEdgeModulePtrOutput() EdgeModulePtrOutput {
	return o
}

func (o EdgeModulePtrOutput) ToEdgeModulePtrOutputWithContext(ctx context.Context) EdgeModulePtrOutput {
	return o
}

func (o EdgeModulePtrOutput) Elem() EdgeModuleOutput {
	return o.ApplyT(func(v *EdgeModule) EdgeModule {
		if v != nil {
			return *v
		}
		var ret EdgeModule
		return ret
	}).(EdgeModuleOutput)
}

type EdgeModuleArrayOutput struct{ *pulumi.OutputState }

func (EdgeModuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdgeModule)(nil))
}

func (o EdgeModuleArrayOutput) ToEdgeModuleArrayOutput() EdgeModuleArrayOutput {
	return o
}

func (o EdgeModuleArrayOutput) ToEdgeModuleArrayOutputWithContext(ctx context.Context) EdgeModuleArrayOutput {
	return o
}

func (o EdgeModuleArrayOutput) Index(i pulumi.IntInput) EdgeModuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdgeModule {
		return vs[0].([]EdgeModule)[vs[1].(int)]
	}).(EdgeModuleOutput)
}

type EdgeModuleMapOutput struct{ *pulumi.OutputState }

func (EdgeModuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EdgeModule)(nil))
}

func (o EdgeModuleMapOutput) ToEdgeModuleMapOutput() EdgeModuleMapOutput {
	return o
}

func (o EdgeModuleMapOutput) ToEdgeModuleMapOutputWithContext(ctx context.Context) EdgeModuleMapOutput {
	return o
}

func (o EdgeModuleMapOutput) MapIndex(k pulumi.StringInput) EdgeModuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EdgeModule {
		return vs[0].(map[string]EdgeModule)[vs[1].(string)]
	}).(EdgeModuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeModuleInput)(nil)).Elem(), &EdgeModule{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeModulePtrInput)(nil)).Elem(), &EdgeModule{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeModuleArrayInput)(nil)).Elem(), EdgeModuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeModuleMapInput)(nil)).Elem(), EdgeModuleMap{})
	pulumi.RegisterOutputType(EdgeModuleOutput{})
	pulumi.RegisterOutputType(EdgeModulePtrOutput{})
	pulumi.RegisterOutputType(EdgeModuleArrayOutput{})
	pulumi.RegisterOutputType(EdgeModuleMapOutput{})
}
