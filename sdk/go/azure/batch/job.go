// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Batch Job.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/batch"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("west europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := batch.NewAccount(ctx, "exampleAccount", &batch.AccountArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePool, err := batch.NewPool(ctx, "examplePool", &batch.PoolArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AccountName:       exampleAccount.Name,
// 			NodeAgentSkuId:    pulumi.String("batch.node.ubuntu 16.04"),
// 			VmSize:            pulumi.String("Standard_A1"),
// 			FixedScale: &batch.PoolFixedScaleArgs{
// 				TargetDedicatedNodes: pulumi.Int(1),
// 			},
// 			StorageImageReference: &batch.PoolStorageImageReferenceArgs{
// 				Publisher: pulumi.String("Canonical"),
// 				Offer:     pulumi.String("UbuntuServer"),
// 				Sku:       pulumi.String("16.04.0-LTS"),
// 				Version:   pulumi.String("latest"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = batch.NewJob(ctx, "exampleJob", &batch.JobArgs{
// 			BatchPoolId: examplePool.ID(),
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
// Batch Jobs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:batch/job:Job example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Batch/batchAccounts/account1/pools/pool1/jobs/job1
// ```
type Job struct {
	pulumi.CustomResourceState

	// The ID of the Batch Pool. Changing this forces a new Batch Job to be created.
	BatchPoolId pulumi.StringOutput `pulumi:"batchPoolId"`
	// Specifies a map of common environment settings applied to this Batch Job. Changing this forces a new Batch Job to be created.
	CommonEnvironmentProperties pulumi.StringMapOutput `pulumi:"commonEnvironmentProperties"`
	// The display name of this Batch Job. Changing this forces a new Batch Job to be created.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name which should be used for this Batch Job. Changing this forces a new Batch Job to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of this Batch Job, possible values can range from -1000 (lowest) to 1000 (highest). Defaults to `0`.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The number of retries to each Batch Task belongs to this Batch Job. If this is set to `0`, the Batch service does not retry Tasks. If this is set to `-1`, the Batch service retries Batch Tasks without limit. Default value is `0`.
	TaskRetryMaximum pulumi.IntPtrOutput `pulumi:"taskRetryMaximum"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BatchPoolId == nil {
		return nil, errors.New("invalid value for required argument 'BatchPoolId'")
	}
	var resource Job
	err := ctx.RegisterResource("azure:batch/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure:batch/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// The ID of the Batch Pool. Changing this forces a new Batch Job to be created.
	BatchPoolId *string `pulumi:"batchPoolId"`
	// Specifies a map of common environment settings applied to this Batch Job. Changing this forces a new Batch Job to be created.
	CommonEnvironmentProperties map[string]string `pulumi:"commonEnvironmentProperties"`
	// The display name of this Batch Job. Changing this forces a new Batch Job to be created.
	DisplayName *string `pulumi:"displayName"`
	// The name which should be used for this Batch Job. Changing this forces a new Batch Job to be created.
	Name *string `pulumi:"name"`
	// The priority of this Batch Job, possible values can range from -1000 (lowest) to 1000 (highest). Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// The number of retries to each Batch Task belongs to this Batch Job. If this is set to `0`, the Batch service does not retry Tasks. If this is set to `-1`, the Batch service retries Batch Tasks without limit. Default value is `0`.
	TaskRetryMaximum *int `pulumi:"taskRetryMaximum"`
}

type JobState struct {
	// The ID of the Batch Pool. Changing this forces a new Batch Job to be created.
	BatchPoolId pulumi.StringPtrInput
	// Specifies a map of common environment settings applied to this Batch Job. Changing this forces a new Batch Job to be created.
	CommonEnvironmentProperties pulumi.StringMapInput
	// The display name of this Batch Job. Changing this forces a new Batch Job to be created.
	DisplayName pulumi.StringPtrInput
	// The name which should be used for this Batch Job. Changing this forces a new Batch Job to be created.
	Name pulumi.StringPtrInput
	// The priority of this Batch Job, possible values can range from -1000 (lowest) to 1000 (highest). Defaults to `0`.
	Priority pulumi.IntPtrInput
	// The number of retries to each Batch Task belongs to this Batch Job. If this is set to `0`, the Batch service does not retry Tasks. If this is set to `-1`, the Batch service retries Batch Tasks without limit. Default value is `0`.
	TaskRetryMaximum pulumi.IntPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The ID of the Batch Pool. Changing this forces a new Batch Job to be created.
	BatchPoolId string `pulumi:"batchPoolId"`
	// Specifies a map of common environment settings applied to this Batch Job. Changing this forces a new Batch Job to be created.
	CommonEnvironmentProperties map[string]string `pulumi:"commonEnvironmentProperties"`
	// The display name of this Batch Job. Changing this forces a new Batch Job to be created.
	DisplayName *string `pulumi:"displayName"`
	// The name which should be used for this Batch Job. Changing this forces a new Batch Job to be created.
	Name *string `pulumi:"name"`
	// The priority of this Batch Job, possible values can range from -1000 (lowest) to 1000 (highest). Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// The number of retries to each Batch Task belongs to this Batch Job. If this is set to `0`, the Batch service does not retry Tasks. If this is set to `-1`, the Batch service retries Batch Tasks without limit. Default value is `0`.
	TaskRetryMaximum *int `pulumi:"taskRetryMaximum"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The ID of the Batch Pool. Changing this forces a new Batch Job to be created.
	BatchPoolId pulumi.StringInput
	// Specifies a map of common environment settings applied to this Batch Job. Changing this forces a new Batch Job to be created.
	CommonEnvironmentProperties pulumi.StringMapInput
	// The display name of this Batch Job. Changing this forces a new Batch Job to be created.
	DisplayName pulumi.StringPtrInput
	// The name which should be used for this Batch Job. Changing this forces a new Batch Job to be created.
	Name pulumi.StringPtrInput
	// The priority of this Batch Job, possible values can range from -1000 (lowest) to 1000 (highest). Defaults to `0`.
	Priority pulumi.IntPtrInput
	// The number of retries to each Batch Task belongs to this Batch Job. If this is set to `0`, the Batch service does not retry Tasks. If this is set to `-1`, the Batch service retries Batch Tasks without limit. Default value is `0`.
	TaskRetryMaximum pulumi.IntPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//          JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//          JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
