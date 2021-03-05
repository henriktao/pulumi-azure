// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto Cluster Principal Assignment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/kusto"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		rg, err := core.NewResourceGroup(ctx, "rg", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCluster, err := kusto.NewCluster(ctx, "exampleCluster", &kusto.ClusterArgs{
// 			Location:          rg.Location,
// 			ResourceGroupName: rg.Name,
// 			Sku: &kusto.ClusterSkuArgs{
// 				Name:     pulumi.String("Standard_D13_v2"),
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kusto.NewClusterPrincipalAssignment(ctx, "exampleClusterPrincipalAssignment", &kusto.ClusterPrincipalAssignmentArgs{
// 			ResourceGroupName: rg.Name,
// 			ClusterName:       exampleCluster.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			PrincipalId:       pulumi.String(current.ClientId),
// 			PrincipalType:     pulumi.String("App"),
// 			Role:              pulumi.String("AllDatabasesAdmin"),
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
// Data Explorer Cluster Principal Assignments can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/PrincipalAssignments/assignment1
// ```
type ClusterPrincipalAssignment struct {
	pulumi.CustomResourceState

	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	Name        pulumi.StringOutput `pulumi:"name"`
	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The name of the principal.
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
	Role pulumi.StringOutput `pulumi:"role"`
	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The name of the tenant.
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
}

// NewClusterPrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, args *ClusterPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource ClusterPrincipalAssignment
	err := ctx.RegisterResource("azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPrincipalAssignment gets an existing ClusterPrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*ClusterPrincipalAssignment, error) {
	var resource ClusterPrincipalAssignment
	err := ctx.ReadResource("azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPrincipalAssignment resources.
type clusterPrincipalAssignmentState struct {
	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	Name        *string `pulumi:"name"`
	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalId *string `pulumi:"principalId"`
	// The name of the principal.
	PrincipalName *string `pulumi:"principalName"`
	// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	PrincipalType *string `pulumi:"principalType"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
	Role *string `pulumi:"role"`
	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantId *string `pulumi:"tenantId"`
	// The name of the tenant.
	TenantName *string `pulumi:"tenantName"`
}

type ClusterPrincipalAssignmentState struct {
	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringPtrInput
	// The name of the principal.
	PrincipalName pulumi.StringPtrInput
	// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	PrincipalType pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
	Role pulumi.StringPtrInput
	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantId pulumi.StringPtrInput
	// The name of the tenant.
	TenantName pulumi.StringPtrInput
}

func (ClusterPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentState)(nil)).Elem()
}

type clusterPrincipalAssignmentArgs struct {
	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName string  `pulumi:"clusterName"`
	Name        *string `pulumi:"name"`
	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalId string `pulumi:"principalId"`
	// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	PrincipalType string `pulumi:"principalType"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
	Role string `pulumi:"role"`
	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ClusterPrincipalAssignment resource.
type ClusterPrincipalAssignmentArgs struct {
	// The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	Name        pulumi.StringPtrInput
	// The object id of the principal. Changing this forces a new resource to be created.
	PrincipalId pulumi.StringInput
	// The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
	PrincipalType pulumi.StringInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
	Role pulumi.StringInput
	// The tenant id in which the principal resides. Changing this forces a new resource to be created.
	TenantId pulumi.StringInput
}

func (ClusterPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPrincipalAssignmentArgs)(nil)).Elem()
}

type ClusterPrincipalAssignmentInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput
	ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput
}

func (*ClusterPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalAssignment)(nil))
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return i.ToClusterPrincipalAssignmentOutputWithContext(context.Background())
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentOutput)
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentPtrOutput() ClusterPrincipalAssignmentPtrOutput {
	return i.ToClusterPrincipalAssignmentPtrOutputWithContext(context.Background())
}

func (i *ClusterPrincipalAssignment) ToClusterPrincipalAssignmentPtrOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentPtrOutput)
}

type ClusterPrincipalAssignmentPtrInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentPtrOutput() ClusterPrincipalAssignmentPtrOutput
	ToClusterPrincipalAssignmentPtrOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentPtrOutput
}

type clusterPrincipalAssignmentPtrType ClusterPrincipalAssignmentArgs

func (*clusterPrincipalAssignmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil))
}

func (i *clusterPrincipalAssignmentPtrType) ToClusterPrincipalAssignmentPtrOutput() ClusterPrincipalAssignmentPtrOutput {
	return i.ToClusterPrincipalAssignmentPtrOutputWithContext(context.Background())
}

func (i *clusterPrincipalAssignmentPtrType) ToClusterPrincipalAssignmentPtrOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentPtrOutput)
}

// ClusterPrincipalAssignmentArrayInput is an input type that accepts ClusterPrincipalAssignmentArray and ClusterPrincipalAssignmentArrayOutput values.
// You can construct a concrete instance of `ClusterPrincipalAssignmentArrayInput` via:
//
//          ClusterPrincipalAssignmentArray{ ClusterPrincipalAssignmentArgs{...} }
type ClusterPrincipalAssignmentArrayInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentArrayOutput() ClusterPrincipalAssignmentArrayOutput
	ToClusterPrincipalAssignmentArrayOutputWithContext(context.Context) ClusterPrincipalAssignmentArrayOutput
}

type ClusterPrincipalAssignmentArray []ClusterPrincipalAssignmentInput

func (ClusterPrincipalAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClusterPrincipalAssignment)(nil))
}

func (i ClusterPrincipalAssignmentArray) ToClusterPrincipalAssignmentArrayOutput() ClusterPrincipalAssignmentArrayOutput {
	return i.ToClusterPrincipalAssignmentArrayOutputWithContext(context.Background())
}

func (i ClusterPrincipalAssignmentArray) ToClusterPrincipalAssignmentArrayOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentArrayOutput)
}

// ClusterPrincipalAssignmentMapInput is an input type that accepts ClusterPrincipalAssignmentMap and ClusterPrincipalAssignmentMapOutput values.
// You can construct a concrete instance of `ClusterPrincipalAssignmentMapInput` via:
//
//          ClusterPrincipalAssignmentMap{ "key": ClusterPrincipalAssignmentArgs{...} }
type ClusterPrincipalAssignmentMapInput interface {
	pulumi.Input

	ToClusterPrincipalAssignmentMapOutput() ClusterPrincipalAssignmentMapOutput
	ToClusterPrincipalAssignmentMapOutputWithContext(context.Context) ClusterPrincipalAssignmentMapOutput
}

type ClusterPrincipalAssignmentMap map[string]ClusterPrincipalAssignmentInput

func (ClusterPrincipalAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClusterPrincipalAssignment)(nil))
}

func (i ClusterPrincipalAssignmentMap) ToClusterPrincipalAssignmentMapOutput() ClusterPrincipalAssignmentMapOutput {
	return i.ToClusterPrincipalAssignmentMapOutputWithContext(context.Background())
}

func (i ClusterPrincipalAssignmentMap) ToClusterPrincipalAssignmentMapOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPrincipalAssignmentMapOutput)
}

type ClusterPrincipalAssignmentOutput struct {
	*pulumi.OutputState
}

func (ClusterPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalAssignment)(nil))
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutput() ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentOutput {
	return o
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentPtrOutput() ClusterPrincipalAssignmentPtrOutput {
	return o.ToClusterPrincipalAssignmentPtrOutputWithContext(context.Background())
}

func (o ClusterPrincipalAssignmentOutput) ToClusterPrincipalAssignmentPtrOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentPtrOutput {
	return o.ApplyT(func(v ClusterPrincipalAssignment) *ClusterPrincipalAssignment {
		return &v
	}).(ClusterPrincipalAssignmentPtrOutput)
}

type ClusterPrincipalAssignmentPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterPrincipalAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalAssignment)(nil))
}

func (o ClusterPrincipalAssignmentPtrOutput) ToClusterPrincipalAssignmentPtrOutput() ClusterPrincipalAssignmentPtrOutput {
	return o
}

func (o ClusterPrincipalAssignmentPtrOutput) ToClusterPrincipalAssignmentPtrOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentPtrOutput {
	return o
}

type ClusterPrincipalAssignmentArrayOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterPrincipalAssignment)(nil))
}

func (o ClusterPrincipalAssignmentArrayOutput) ToClusterPrincipalAssignmentArrayOutput() ClusterPrincipalAssignmentArrayOutput {
	return o
}

func (o ClusterPrincipalAssignmentArrayOutput) ToClusterPrincipalAssignmentArrayOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentArrayOutput {
	return o
}

func (o ClusterPrincipalAssignmentArrayOutput) Index(i pulumi.IntInput) ClusterPrincipalAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterPrincipalAssignment {
		return vs[0].([]ClusterPrincipalAssignment)[vs[1].(int)]
	}).(ClusterPrincipalAssignmentOutput)
}

type ClusterPrincipalAssignmentMapOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterPrincipalAssignment)(nil))
}

func (o ClusterPrincipalAssignmentMapOutput) ToClusterPrincipalAssignmentMapOutput() ClusterPrincipalAssignmentMapOutput {
	return o
}

func (o ClusterPrincipalAssignmentMapOutput) ToClusterPrincipalAssignmentMapOutputWithContext(ctx context.Context) ClusterPrincipalAssignmentMapOutput {
	return o
}

func (o ClusterPrincipalAssignmentMapOutput) MapIndex(k pulumi.StringInput) ClusterPrincipalAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterPrincipalAssignment {
		return vs[0].(map[string]ClusterPrincipalAssignment)[vs[1].(string)]
	}).(ClusterPrincipalAssignmentOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentOutput{})
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentPtrOutput{})
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentArrayOutput{})
	pulumi.RegisterOutputType(ClusterPrincipalAssignmentMapOutput{})
}
