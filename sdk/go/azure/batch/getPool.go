// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Batch pool
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/batch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.LookupPool(ctx, &batch.LookupPoolArgs{
// 			AccountName:       "testbatchaccount",
// 			Name:              "testbatchpool",
// 			ResourceGroupName: "test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure:batch/getPool:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPool.
type LookupPoolArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the endpoint.
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getPool.
type LookupPoolResult struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// A `autoScale` block that describes the scale settings when using auto scale.
	AutoScales []GetPoolAutoScale `pulumi:"autoScales"`
	// One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
	Certificates []GetPoolCertificate `pulumi:"certificates"`
	// The container configuration used in the pool's VMs.
	ContainerConfigurations []GetPoolContainerConfiguration `pulumi:"containerConfigurations"`
	DisplayName             string                          `pulumi:"displayName"`
	// A `fixedScale` block that describes the scale settings when using fixed scale.
	FixedScales []GetPoolFixedScale `pulumi:"fixedScales"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The maximum number of tasks that can run concurrently on a single compute node in the pool.
	MaxTasksPerNode int               `pulumi:"maxTasksPerNode"`
	Metadata        map[string]string `pulumi:"metadata"`
	// The name of the endpoint.
	Name                 string                      `pulumi:"name"`
	NetworkConfiguration GetPoolNetworkConfiguration `pulumi:"networkConfiguration"`
	// The Sku of the node agents in the Batch pool.
	NodeAgentSkuId    string `pulumi:"nodeAgentSkuId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `startTask` block that describes the start task settings for the Batch pool.
	StartTasks []GetPoolStartTask `pulumi:"startTasks"`
	// The reference of the storage image used by the nodes in the Batch pool.
	StorageImageReferences []GetPoolStorageImageReference `pulumi:"storageImageReferences"`
	// The size of the VM created in the Batch pool.
	VmSize string `pulumi:"vmSize"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			return *r, err
		}).(LookupPoolResultOutput)
}

// A collection of arguments for invoking getPool.
type LookupPoolOutputArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the endpoint.
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// A collection of values returned by getPool.
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// The name of the Batch account.
func (o LookupPoolResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// A `autoScale` block that describes the scale settings when using auto scale.
func (o LookupPoolResultOutput) AutoScales() GetPoolAutoScaleArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolAutoScale { return v.AutoScales }).(GetPoolAutoScaleArrayOutput)
}

// One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
func (o LookupPoolResultOutput) Certificates() GetPoolCertificateArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolCertificate { return v.Certificates }).(GetPoolCertificateArrayOutput)
}

// The container configuration used in the pool's VMs.
func (o LookupPoolResultOutput) ContainerConfigurations() GetPoolContainerConfigurationArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolContainerConfiguration { return v.ContainerConfigurations }).(GetPoolContainerConfigurationArrayOutput)
}

func (o LookupPoolResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A `fixedScale` block that describes the scale settings when using fixed scale.
func (o LookupPoolResultOutput) FixedScales() GetPoolFixedScaleArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolFixedScale { return v.FixedScales }).(GetPoolFixedScaleArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The maximum number of tasks that can run concurrently on a single compute node in the pool.
func (o LookupPoolResultOutput) MaxTasksPerNode() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPoolResult) int { return v.MaxTasksPerNode }).(pulumi.IntOutput)
}

func (o LookupPoolResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the endpoint.
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) NetworkConfiguration() GetPoolNetworkConfigurationOutput {
	return o.ApplyT(func(v LookupPoolResult) GetPoolNetworkConfiguration { return v.NetworkConfiguration }).(GetPoolNetworkConfigurationOutput)
}

// The Sku of the node agents in the Batch pool.
func (o LookupPoolResultOutput) NodeAgentSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.NodeAgentSkuId }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

// A `startTask` block that describes the start task settings for the Batch pool.
func (o LookupPoolResultOutput) StartTasks() GetPoolStartTaskArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolStartTask { return v.StartTasks }).(GetPoolStartTaskArrayOutput)
}

// The reference of the storage image used by the nodes in the Batch pool.
func (o LookupPoolResultOutput) StorageImageReferences() GetPoolStorageImageReferenceArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolStorageImageReference { return v.StorageImageReferences }).(GetPoolStorageImageReferenceArrayOutput)
}

// The size of the VM created in the Batch pool.
func (o LookupPoolResultOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.VmSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
