// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Video Analyzer.
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
// 		_, err = videoanalyzer.NewAnalyzer(ctx, "exampleAnalyzer", &videoanalyzer.AnalyzerArgs{
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
// 			contributor,
// 			reader,
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
// Video Analyzer can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:videoanalyzer/analyzer:Analyzer analyzer /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1
// ```
type Analyzer struct {
	pulumi.CustomResourceState

	// An `identity` block as defined below.
	Identity AnalyzerIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Video Analyzer. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `storageAccount` block as defined below.
	StorageAccount AnalyzerStorageAccountOutput `pulumi:"storageAccount"`
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAnalyzer registers a new resource with the given unique name, arguments, and options.
func NewAnalyzer(ctx *pulumi.Context,
	name string, args *AnalyzerArgs, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	var resource Analyzer
	err := ctx.RegisterResource("azure:videoanalyzer/analyzer:Analyzer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyzer gets an existing Analyzer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyzer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyzerState, opts ...pulumi.ResourceOption) (*Analyzer, error) {
	var resource Analyzer
	err := ctx.ReadResource("azure:videoanalyzer/analyzer:Analyzer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analyzer resources.
type analyzerState struct {
	// An `identity` block as defined below.
	Identity *AnalyzerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Video Analyzer. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `storageAccount` block as defined below.
	StorageAccount *AnalyzerStorageAccount `pulumi:"storageAccount"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AnalyzerState struct {
	// An `identity` block as defined below.
	Identity AnalyzerIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Video Analyzer. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Video Analyzer. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `storageAccount` block as defined below.
	StorageAccount AnalyzerStorageAccountPtrInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
}

func (AnalyzerState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerState)(nil)).Elem()
}

type analyzerArgs struct {
	// An `identity` block as defined below.
	Identity AnalyzerIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Video Analyzer. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Video Analyzer. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `storageAccount` block as defined below.
	StorageAccount AnalyzerStorageAccount `pulumi:"storageAccount"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Analyzer resource.
type AnalyzerArgs struct {
	// An `identity` block as defined below.
	Identity AnalyzerIdentityInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Video Analyzer. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Video Analyzer. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `storageAccount` block as defined below.
	StorageAccount AnalyzerStorageAccountInput
	// A mapping of tags assigned to the resource.
	Tags pulumi.StringMapInput
}

func (AnalyzerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyzerArgs)(nil)).Elem()
}

type AnalyzerInput interface {
	pulumi.Input

	ToAnalyzerOutput() AnalyzerOutput
	ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput
}

func (*Analyzer) ElementType() reflect.Type {
	return reflect.TypeOf((**Analyzer)(nil)).Elem()
}

func (i *Analyzer) ToAnalyzerOutput() AnalyzerOutput {
	return i.ToAnalyzerOutputWithContext(context.Background())
}

func (i *Analyzer) ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyzerOutput)
}

// AnalyzerArrayInput is an input type that accepts AnalyzerArray and AnalyzerArrayOutput values.
// You can construct a concrete instance of `AnalyzerArrayInput` via:
//
//          AnalyzerArray{ AnalyzerArgs{...} }
type AnalyzerArrayInput interface {
	pulumi.Input

	ToAnalyzerArrayOutput() AnalyzerArrayOutput
	ToAnalyzerArrayOutputWithContext(context.Context) AnalyzerArrayOutput
}

type AnalyzerArray []AnalyzerInput

func (AnalyzerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Analyzer)(nil)).Elem()
}

func (i AnalyzerArray) ToAnalyzerArrayOutput() AnalyzerArrayOutput {
	return i.ToAnalyzerArrayOutputWithContext(context.Background())
}

func (i AnalyzerArray) ToAnalyzerArrayOutputWithContext(ctx context.Context) AnalyzerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyzerArrayOutput)
}

// AnalyzerMapInput is an input type that accepts AnalyzerMap and AnalyzerMapOutput values.
// You can construct a concrete instance of `AnalyzerMapInput` via:
//
//          AnalyzerMap{ "key": AnalyzerArgs{...} }
type AnalyzerMapInput interface {
	pulumi.Input

	ToAnalyzerMapOutput() AnalyzerMapOutput
	ToAnalyzerMapOutputWithContext(context.Context) AnalyzerMapOutput
}

type AnalyzerMap map[string]AnalyzerInput

func (AnalyzerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Analyzer)(nil)).Elem()
}

func (i AnalyzerMap) ToAnalyzerMapOutput() AnalyzerMapOutput {
	return i.ToAnalyzerMapOutputWithContext(context.Background())
}

func (i AnalyzerMap) ToAnalyzerMapOutputWithContext(ctx context.Context) AnalyzerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyzerMapOutput)
}

type AnalyzerOutput struct{ *pulumi.OutputState }

func (AnalyzerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Analyzer)(nil)).Elem()
}

func (o AnalyzerOutput) ToAnalyzerOutput() AnalyzerOutput {
	return o
}

func (o AnalyzerOutput) ToAnalyzerOutputWithContext(ctx context.Context) AnalyzerOutput {
	return o
}

type AnalyzerArrayOutput struct{ *pulumi.OutputState }

func (AnalyzerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Analyzer)(nil)).Elem()
}

func (o AnalyzerArrayOutput) ToAnalyzerArrayOutput() AnalyzerArrayOutput {
	return o
}

func (o AnalyzerArrayOutput) ToAnalyzerArrayOutputWithContext(ctx context.Context) AnalyzerArrayOutput {
	return o
}

func (o AnalyzerArrayOutput) Index(i pulumi.IntInput) AnalyzerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Analyzer {
		return vs[0].([]*Analyzer)[vs[1].(int)]
	}).(AnalyzerOutput)
}

type AnalyzerMapOutput struct{ *pulumi.OutputState }

func (AnalyzerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Analyzer)(nil)).Elem()
}

func (o AnalyzerMapOutput) ToAnalyzerMapOutput() AnalyzerMapOutput {
	return o
}

func (o AnalyzerMapOutput) ToAnalyzerMapOutputWithContext(ctx context.Context) AnalyzerMapOutput {
	return o
}

func (o AnalyzerMapOutput) MapIndex(k pulumi.StringInput) AnalyzerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Analyzer {
		return vs[0].(map[string]*Analyzer)[vs[1].(string)]
	}).(AnalyzerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyzerInput)(nil)).Elem(), &Analyzer{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyzerArrayInput)(nil)).Elem(), AnalyzerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyzerMapInput)(nil)).Elem(), AnalyzerMap{})
	pulumi.RegisterOutputType(AnalyzerOutput{})
	pulumi.RegisterOutputType(AnalyzerArrayOutput{})
	pulumi.RegisterOutputType(AnalyzerMapOutput{})
}
