// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Data Factory Managed Private Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:                     exampleResourceGroup.Location,
// 			ResourceGroupName:            exampleResourceGroup.Name,
// 			ManagedVirtualNetworkEnabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountKind:            pulumi.String("BlobStorage"),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewManagedPrivateEndpoint(ctx, "exampleManagedPrivateEndpoint", &datafactory.ManagedPrivateEndpointArgs{
// 			DataFactoryId:    exampleFactory.ID(),
// 			TargetResourceId: exampleAccount.ID(),
// 			SubresourceName:  pulumi.String("blob"),
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
// Data Factory Managed Private Endpoint can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
// ```
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryId pulumi.StringOutput `pulumi:"dataFactoryId"`
	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns pulumi.StringArrayOutput `pulumi:"fqdns"`
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringPtrOutput `pulumi:"subresourceName"`
	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryId == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryId'")
	}
	if args.TargetResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceId'")
	}
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure:datafactory/managedPrivateEndpoint:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryId *string `pulumi:"dataFactoryId"`
	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns []string `pulumi:"fqdns"`
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `pulumi:"subresourceName"`
	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type ManagedPrivateEndpointState struct {
	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryId pulumi.StringPtrInput
	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns pulumi.StringArrayInput
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringPtrInput
	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryId string `pulumi:"dataFactoryId"`
	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns []string `pulumi:"fqdns"`
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `pulumi:"subresourceName"`
	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryId pulumi.StringInput
	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns pulumi.StringArrayInput
	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName pulumi.StringPtrInput
	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointPtrOutput() ManagedPrivateEndpointPtrOutput {
	return i.ToManagedPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointPtrOutputWithContext(ctx context.Context) ManagedPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointPtrOutput)
}

type ManagedPrivateEndpointPtrInput interface {
	pulumi.Input

	ToManagedPrivateEndpointPtrOutput() ManagedPrivateEndpointPtrOutput
	ToManagedPrivateEndpointPtrOutputWithContext(ctx context.Context) ManagedPrivateEndpointPtrOutput
}

type managedPrivateEndpointPtrType ManagedPrivateEndpointArgs

func (*managedPrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil))
}

func (i *managedPrivateEndpointPtrType) ToManagedPrivateEndpointPtrOutput() ManagedPrivateEndpointPtrOutput {
	return i.ToManagedPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *managedPrivateEndpointPtrType) ToManagedPrivateEndpointPtrOutputWithContext(ctx context.Context) ManagedPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointPtrOutput)
}

// ManagedPrivateEndpointArrayInput is an input type that accepts ManagedPrivateEndpointArray and ManagedPrivateEndpointArrayOutput values.
// You can construct a concrete instance of `ManagedPrivateEndpointArrayInput` via:
//
//          ManagedPrivateEndpointArray{ ManagedPrivateEndpointArgs{...} }
type ManagedPrivateEndpointArrayInput interface {
	pulumi.Input

	ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput
	ToManagedPrivateEndpointArrayOutputWithContext(context.Context) ManagedPrivateEndpointArrayOutput
}

type ManagedPrivateEndpointArray []ManagedPrivateEndpointInput

func (ManagedPrivateEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (i ManagedPrivateEndpointArray) ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput {
	return i.ToManagedPrivateEndpointArrayOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointArray) ToManagedPrivateEndpointArrayOutputWithContext(ctx context.Context) ManagedPrivateEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointArrayOutput)
}

// ManagedPrivateEndpointMapInput is an input type that accepts ManagedPrivateEndpointMap and ManagedPrivateEndpointMapOutput values.
// You can construct a concrete instance of `ManagedPrivateEndpointMapInput` via:
//
//          ManagedPrivateEndpointMap{ "key": ManagedPrivateEndpointArgs{...} }
type ManagedPrivateEndpointMapInput interface {
	pulumi.Input

	ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput
	ToManagedPrivateEndpointMapOutputWithContext(context.Context) ManagedPrivateEndpointMapOutput
}

type ManagedPrivateEndpointMap map[string]ManagedPrivateEndpointInput

func (ManagedPrivateEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedPrivateEndpoint)(nil)).Elem()
}

func (i ManagedPrivateEndpointMap) ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput {
	return i.ToManagedPrivateEndpointMapOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointMap) ToManagedPrivateEndpointMapOutputWithContext(ctx context.Context) ManagedPrivateEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointMapOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointPtrOutput() ManagedPrivateEndpointPtrOutput {
	return o.ToManagedPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointPtrOutputWithContext(ctx context.Context) ManagedPrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedPrivateEndpoint) *ManagedPrivateEndpoint {
		return &v
	}).(ManagedPrivateEndpointPtrOutput)
}

type ManagedPrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointPtrOutput) ToManagedPrivateEndpointPtrOutput() ManagedPrivateEndpointPtrOutput {
	return o
}

func (o ManagedPrivateEndpointPtrOutput) ToManagedPrivateEndpointPtrOutputWithContext(ctx context.Context) ManagedPrivateEndpointPtrOutput {
	return o
}

func (o ManagedPrivateEndpointPtrOutput) Elem() ManagedPrivateEndpointOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) ManagedPrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret ManagedPrivateEndpoint
		return ret
	}).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointArrayOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointArrayOutput) ToManagedPrivateEndpointArrayOutput() ManagedPrivateEndpointArrayOutput {
	return o
}

func (o ManagedPrivateEndpointArrayOutput) ToManagedPrivateEndpointArrayOutputWithContext(ctx context.Context) ManagedPrivateEndpointArrayOutput {
	return o
}

func (o ManagedPrivateEndpointArrayOutput) Index(i pulumi.IntInput) ManagedPrivateEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedPrivateEndpoint {
		return vs[0].([]ManagedPrivateEndpoint)[vs[1].(int)]
	}).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointMapOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointMapOutput) ToManagedPrivateEndpointMapOutput() ManagedPrivateEndpointMapOutput {
	return o
}

func (o ManagedPrivateEndpointMapOutput) ToManagedPrivateEndpointMapOutputWithContext(ctx context.Context) ManagedPrivateEndpointMapOutput {
	return o
}

func (o ManagedPrivateEndpointMapOutput) MapIndex(k pulumi.StringInput) ManagedPrivateEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedPrivateEndpoint {
		return vs[0].(map[string]ManagedPrivateEndpoint)[vs[1].(string)]
	}).(ManagedPrivateEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointInput)(nil)).Elem(), &ManagedPrivateEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointPtrInput)(nil)).Elem(), &ManagedPrivateEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointArrayInput)(nil)).Elem(), ManagedPrivateEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrivateEndpointMapInput)(nil)).Elem(), ManagedPrivateEndpointMap{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointArrayOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointMapOutput{})
}
