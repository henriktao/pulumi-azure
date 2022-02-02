// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package stack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Stack HCI Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/stack"
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example-app"
// 		exampleApplication, err := azuread.LookupApplication(ctx, &GetApplicationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = stack.NewHciCluster(ctx, "exampleHciCluster", &stack.HciClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClientId:          pulumi.String(exampleApplication.ApplicationId),
// 			TenantId:          pulumi.String(current.TenantId),
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
// Azure Stack HCI Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:stack/hciCluster:HciCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1
// ```
type HciCluster struct {
	pulumi.CustomResourceState

	// The Client ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The Azure Region where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Azure Stack HCI Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Tenant ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewHciCluster registers a new resource with the given unique name, arguments, and options.
func NewHciCluster(ctx *pulumi.Context,
	name string, args *HciClusterArgs, opts ...pulumi.ResourceOption) (*HciCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource HciCluster
	err := ctx.RegisterResource("azure:stack/hciCluster:HciCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHciCluster gets an existing HciCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHciCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HciClusterState, opts ...pulumi.ResourceOption) (*HciCluster, error) {
	var resource HciCluster
	err := ctx.ReadResource("azure:stack/hciCluster:HciCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HciCluster resources.
type hciClusterState struct {
	// The Client ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	ClientId *string `pulumi:"clientId"`
	// The Azure Region where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Azure Stack HCI Cluster.
	Tags map[string]string `pulumi:"tags"`
	// The Tenant ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	TenantId *string `pulumi:"tenantId"`
}

type HciClusterState struct {
	// The Client ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	ClientId pulumi.StringPtrInput
	// The Azure Region where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Azure Stack HCI Cluster.
	Tags pulumi.StringMapInput
	// The Tenant ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	TenantId pulumi.StringPtrInput
}

func (HciClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hciClusterState)(nil)).Elem()
}

type hciClusterArgs struct {
	// The Client ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	ClientId string `pulumi:"clientId"`
	// The Azure Region where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Azure Stack HCI Cluster.
	Tags map[string]string `pulumi:"tags"`
	// The Tenant ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a HciCluster resource.
type HciClusterArgs struct {
	// The Client ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	ClientId pulumi.StringInput
	// The Azure Region where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Azure Stack HCI Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Azure Stack HCI Cluster.
	Tags pulumi.StringMapInput
	// The Tenant ID of the Azure Active Directory which is used by the Azure Stack HCI Cluster. Changing this forces a new resource to be created.
	TenantId pulumi.StringPtrInput
}

func (HciClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hciClusterArgs)(nil)).Elem()
}

type HciClusterInput interface {
	pulumi.Input

	ToHciClusterOutput() HciClusterOutput
	ToHciClusterOutputWithContext(ctx context.Context) HciClusterOutput
}

func (*HciCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**HciCluster)(nil)).Elem()
}

func (i *HciCluster) ToHciClusterOutput() HciClusterOutput {
	return i.ToHciClusterOutputWithContext(context.Background())
}

func (i *HciCluster) ToHciClusterOutputWithContext(ctx context.Context) HciClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HciClusterOutput)
}

// HciClusterArrayInput is an input type that accepts HciClusterArray and HciClusterArrayOutput values.
// You can construct a concrete instance of `HciClusterArrayInput` via:
//
//          HciClusterArray{ HciClusterArgs{...} }
type HciClusterArrayInput interface {
	pulumi.Input

	ToHciClusterArrayOutput() HciClusterArrayOutput
	ToHciClusterArrayOutputWithContext(context.Context) HciClusterArrayOutput
}

type HciClusterArray []HciClusterInput

func (HciClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HciCluster)(nil)).Elem()
}

func (i HciClusterArray) ToHciClusterArrayOutput() HciClusterArrayOutput {
	return i.ToHciClusterArrayOutputWithContext(context.Background())
}

func (i HciClusterArray) ToHciClusterArrayOutputWithContext(ctx context.Context) HciClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HciClusterArrayOutput)
}

// HciClusterMapInput is an input type that accepts HciClusterMap and HciClusterMapOutput values.
// You can construct a concrete instance of `HciClusterMapInput` via:
//
//          HciClusterMap{ "key": HciClusterArgs{...} }
type HciClusterMapInput interface {
	pulumi.Input

	ToHciClusterMapOutput() HciClusterMapOutput
	ToHciClusterMapOutputWithContext(context.Context) HciClusterMapOutput
}

type HciClusterMap map[string]HciClusterInput

func (HciClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HciCluster)(nil)).Elem()
}

func (i HciClusterMap) ToHciClusterMapOutput() HciClusterMapOutput {
	return i.ToHciClusterMapOutputWithContext(context.Background())
}

func (i HciClusterMap) ToHciClusterMapOutputWithContext(ctx context.Context) HciClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HciClusterMapOutput)
}

type HciClusterOutput struct{ *pulumi.OutputState }

func (HciClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HciCluster)(nil)).Elem()
}

func (o HciClusterOutput) ToHciClusterOutput() HciClusterOutput {
	return o
}

func (o HciClusterOutput) ToHciClusterOutputWithContext(ctx context.Context) HciClusterOutput {
	return o
}

type HciClusterArrayOutput struct{ *pulumi.OutputState }

func (HciClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HciCluster)(nil)).Elem()
}

func (o HciClusterArrayOutput) ToHciClusterArrayOutput() HciClusterArrayOutput {
	return o
}

func (o HciClusterArrayOutput) ToHciClusterArrayOutputWithContext(ctx context.Context) HciClusterArrayOutput {
	return o
}

func (o HciClusterArrayOutput) Index(i pulumi.IntInput) HciClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HciCluster {
		return vs[0].([]*HciCluster)[vs[1].(int)]
	}).(HciClusterOutput)
}

type HciClusterMapOutput struct{ *pulumi.OutputState }

func (HciClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HciCluster)(nil)).Elem()
}

func (o HciClusterMapOutput) ToHciClusterMapOutput() HciClusterMapOutput {
	return o
}

func (o HciClusterMapOutput) ToHciClusterMapOutputWithContext(ctx context.Context) HciClusterMapOutput {
	return o
}

func (o HciClusterMapOutput) MapIndex(k pulumi.StringInput) HciClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HciCluster {
		return vs[0].(map[string]*HciCluster)[vs[1].(string)]
	}).(HciClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HciClusterInput)(nil)).Elem(), &HciCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*HciClusterArrayInput)(nil)).Elem(), HciClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HciClusterMapInput)(nil)).Elem(), HciClusterMap{})
	pulumi.RegisterOutputType(HciClusterOutput{})
	pulumi.RegisterOutputType(HciClusterArrayOutput{})
	pulumi.RegisterOutputType(HciClusterMapOutput{})
}
