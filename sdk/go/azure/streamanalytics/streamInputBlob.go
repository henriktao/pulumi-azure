// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Stream Analytics Stream Input Blob.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/streamanalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "example-resources",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleJob, err := streamanalytics.LookupJob(ctx, &streamanalytics.LookupJobArgs{
// 			Name:              "example-job",
// 			ResourceGroupName: azurerm_resource_group.Example.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      pulumi.Any(azurerm_resource_group.Example.Name),
// 			Location:               pulumi.Any(azurerm_resource_group.Example.Location),
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = streamanalytics.NewStreamInputBlob(ctx, "exampleStreamInputBlob", &streamanalytics.StreamInputBlobArgs{
// 			StreamAnalyticsJobName: pulumi.String(exampleJob.Name),
// 			ResourceGroupName:      pulumi.String(exampleJob.ResourceGroupName),
// 			StorageAccountName:     exampleAccount.Name,
// 			StorageAccountKey:      exampleAccount.PrimaryAccessKey,
// 			StorageContainerName:   exampleContainer.Name,
// 			PathPattern:            pulumi.String("some-random-pattern"),
// 			DateFormat:             pulumi.String("yyyy/MM/dd"),
// 			TimeFormat:             pulumi.String("HH"),
// 			Serialization: &streamanalytics.StreamInputBlobSerializationArgs{
// 				Type:     pulumi.String("Json"),
// 				Encoding: pulumi.String("UTF8"),
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
// Stream Analytics Stream Input Blob's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:streamanalytics/streamInputBlob:StreamInputBlob example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
// ```
type StreamInputBlob struct {
	pulumi.CustomResourceState

	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringOutput `pulumi:"dateFormat"`
	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringOutput `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputBlobSerializationOutput `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringOutput `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringOutput `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringOutput `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringOutput `pulumi:"timeFormat"`
}

// NewStreamInputBlob registers a new resource with the given unique name, arguments, and options.
func NewStreamInputBlob(ctx *pulumi.Context,
	name string, args *StreamInputBlobArgs, opts ...pulumi.ResourceOption) (*StreamInputBlob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DateFormat == nil {
		return nil, errors.New("invalid value for required argument 'DateFormat'")
	}
	if args.PathPattern == nil {
		return nil, errors.New("invalid value for required argument 'PathPattern'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Serialization == nil {
		return nil, errors.New("invalid value for required argument 'Serialization'")
	}
	if args.StorageAccountKey == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountKey'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	if args.StreamAnalyticsJobName == nil {
		return nil, errors.New("invalid value for required argument 'StreamAnalyticsJobName'")
	}
	if args.TimeFormat == nil {
		return nil, errors.New("invalid value for required argument 'TimeFormat'")
	}
	var resource StreamInputBlob
	err := ctx.RegisterResource("azure:streamanalytics/streamInputBlob:StreamInputBlob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamInputBlob gets an existing StreamInputBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamInputBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamInputBlobState, opts ...pulumi.ResourceOption) (*StreamInputBlob, error) {
	var resource StreamInputBlob
	err := ctx.ReadResource("azure:streamanalytics/streamInputBlob:StreamInputBlob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamInputBlob resources.
type streamInputBlobState struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat *string `pulumi:"dateFormat"`
	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern *string `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization *StreamInputBlobSerialization `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey *string `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName *string `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat *string `pulumi:"timeFormat"`
}

type StreamInputBlobState struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringPtrInput
	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringPtrInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `serialization` block as defined below.
	Serialization StreamInputBlobSerializationPtrInput
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringPtrInput
	// The name of the Storage Account.
	StorageAccountName pulumi.StringPtrInput
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringPtrInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringPtrInput
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringPtrInput
}

func (StreamInputBlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputBlobState)(nil)).Elem()
}

type streamInputBlobArgs struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat string `pulumi:"dateFormat"`
	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern string `pulumi:"pathPattern"`
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `serialization` block as defined below.
	Serialization StreamInputBlobSerialization `pulumi:"serialization"`
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey string `pulumi:"storageAccountKey"`
	// The name of the Storage Account.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the Container within the Storage Account.
	StorageContainerName string `pulumi:"storageContainerName"`
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName string `pulumi:"streamAnalyticsJobName"`
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat string `pulumi:"timeFormat"`
}

// The set of arguments for constructing a StreamInputBlob resource.
type StreamInputBlobArgs struct {
	// The date format. Wherever `{date}` appears in `pathPattern`, the value of this property is used as the date format instead.
	DateFormat pulumi.StringInput
	// The name of the Stream Input Blob. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
	PathPattern pulumi.StringInput
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `serialization` block as defined below.
	Serialization StreamInputBlobSerializationInput
	// The Access Key which should be used to connect to this Storage Account.
	StorageAccountKey pulumi.StringInput
	// The name of the Storage Account.
	StorageAccountName pulumi.StringInput
	// The name of the Container within the Storage Account.
	StorageContainerName pulumi.StringInput
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName pulumi.StringInput
	// The time format. Wherever `{time}` appears in `pathPattern`, the value of this property is used as the time format instead.
	TimeFormat pulumi.StringInput
}

func (StreamInputBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamInputBlobArgs)(nil)).Elem()
}

type StreamInputBlobInput interface {
	pulumi.Input

	ToStreamInputBlobOutput() StreamInputBlobOutput
	ToStreamInputBlobOutputWithContext(ctx context.Context) StreamInputBlobOutput
}

func (*StreamInputBlob) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputBlob)(nil))
}

func (i *StreamInputBlob) ToStreamInputBlobOutput() StreamInputBlobOutput {
	return i.ToStreamInputBlobOutputWithContext(context.Background())
}

func (i *StreamInputBlob) ToStreamInputBlobOutputWithContext(ctx context.Context) StreamInputBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputBlobOutput)
}

func (i *StreamInputBlob) ToStreamInputBlobPtrOutput() StreamInputBlobPtrOutput {
	return i.ToStreamInputBlobPtrOutputWithContext(context.Background())
}

func (i *StreamInputBlob) ToStreamInputBlobPtrOutputWithContext(ctx context.Context) StreamInputBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputBlobPtrOutput)
}

type StreamInputBlobPtrInput interface {
	pulumi.Input

	ToStreamInputBlobPtrOutput() StreamInputBlobPtrOutput
	ToStreamInputBlobPtrOutputWithContext(ctx context.Context) StreamInputBlobPtrOutput
}

type streamInputBlobPtrType StreamInputBlobArgs

func (*streamInputBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamInputBlob)(nil))
}

func (i *streamInputBlobPtrType) ToStreamInputBlobPtrOutput() StreamInputBlobPtrOutput {
	return i.ToStreamInputBlobPtrOutputWithContext(context.Background())
}

func (i *streamInputBlobPtrType) ToStreamInputBlobPtrOutputWithContext(ctx context.Context) StreamInputBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputBlobPtrOutput)
}

// StreamInputBlobArrayInput is an input type that accepts StreamInputBlobArray and StreamInputBlobArrayOutput values.
// You can construct a concrete instance of `StreamInputBlobArrayInput` via:
//
//          StreamInputBlobArray{ StreamInputBlobArgs{...} }
type StreamInputBlobArrayInput interface {
	pulumi.Input

	ToStreamInputBlobArrayOutput() StreamInputBlobArrayOutput
	ToStreamInputBlobArrayOutputWithContext(context.Context) StreamInputBlobArrayOutput
}

type StreamInputBlobArray []StreamInputBlobInput

func (StreamInputBlobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamInputBlob)(nil)).Elem()
}

func (i StreamInputBlobArray) ToStreamInputBlobArrayOutput() StreamInputBlobArrayOutput {
	return i.ToStreamInputBlobArrayOutputWithContext(context.Background())
}

func (i StreamInputBlobArray) ToStreamInputBlobArrayOutputWithContext(ctx context.Context) StreamInputBlobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputBlobArrayOutput)
}

// StreamInputBlobMapInput is an input type that accepts StreamInputBlobMap and StreamInputBlobMapOutput values.
// You can construct a concrete instance of `StreamInputBlobMapInput` via:
//
//          StreamInputBlobMap{ "key": StreamInputBlobArgs{...} }
type StreamInputBlobMapInput interface {
	pulumi.Input

	ToStreamInputBlobMapOutput() StreamInputBlobMapOutput
	ToStreamInputBlobMapOutputWithContext(context.Context) StreamInputBlobMapOutput
}

type StreamInputBlobMap map[string]StreamInputBlobInput

func (StreamInputBlobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamInputBlob)(nil)).Elem()
}

func (i StreamInputBlobMap) ToStreamInputBlobMapOutput() StreamInputBlobMapOutput {
	return i.ToStreamInputBlobMapOutputWithContext(context.Background())
}

func (i StreamInputBlobMap) ToStreamInputBlobMapOutputWithContext(ctx context.Context) StreamInputBlobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputBlobMapOutput)
}

type StreamInputBlobOutput struct{ *pulumi.OutputState }

func (StreamInputBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputBlob)(nil))
}

func (o StreamInputBlobOutput) ToStreamInputBlobOutput() StreamInputBlobOutput {
	return o
}

func (o StreamInputBlobOutput) ToStreamInputBlobOutputWithContext(ctx context.Context) StreamInputBlobOutput {
	return o
}

func (o StreamInputBlobOutput) ToStreamInputBlobPtrOutput() StreamInputBlobPtrOutput {
	return o.ToStreamInputBlobPtrOutputWithContext(context.Background())
}

func (o StreamInputBlobOutput) ToStreamInputBlobPtrOutputWithContext(ctx context.Context) StreamInputBlobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamInputBlob) *StreamInputBlob {
		return &v
	}).(StreamInputBlobPtrOutput)
}

type StreamInputBlobPtrOutput struct{ *pulumi.OutputState }

func (StreamInputBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamInputBlob)(nil))
}

func (o StreamInputBlobPtrOutput) ToStreamInputBlobPtrOutput() StreamInputBlobPtrOutput {
	return o
}

func (o StreamInputBlobPtrOutput) ToStreamInputBlobPtrOutputWithContext(ctx context.Context) StreamInputBlobPtrOutput {
	return o
}

func (o StreamInputBlobPtrOutput) Elem() StreamInputBlobOutput {
	return o.ApplyT(func(v *StreamInputBlob) StreamInputBlob {
		if v != nil {
			return *v
		}
		var ret StreamInputBlob
		return ret
	}).(StreamInputBlobOutput)
}

type StreamInputBlobArrayOutput struct{ *pulumi.OutputState }

func (StreamInputBlobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamInputBlob)(nil))
}

func (o StreamInputBlobArrayOutput) ToStreamInputBlobArrayOutput() StreamInputBlobArrayOutput {
	return o
}

func (o StreamInputBlobArrayOutput) ToStreamInputBlobArrayOutputWithContext(ctx context.Context) StreamInputBlobArrayOutput {
	return o
}

func (o StreamInputBlobArrayOutput) Index(i pulumi.IntInput) StreamInputBlobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamInputBlob {
		return vs[0].([]StreamInputBlob)[vs[1].(int)]
	}).(StreamInputBlobOutput)
}

type StreamInputBlobMapOutput struct{ *pulumi.OutputState }

func (StreamInputBlobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamInputBlob)(nil))
}

func (o StreamInputBlobMapOutput) ToStreamInputBlobMapOutput() StreamInputBlobMapOutput {
	return o
}

func (o StreamInputBlobMapOutput) ToStreamInputBlobMapOutputWithContext(ctx context.Context) StreamInputBlobMapOutput {
	return o
}

func (o StreamInputBlobMapOutput) MapIndex(k pulumi.StringInput) StreamInputBlobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StreamInputBlob {
		return vs[0].(map[string]StreamInputBlob)[vs[1].(string)]
	}).(StreamInputBlobOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamInputBlobOutput{})
	pulumi.RegisterOutputType(StreamInputBlobPtrOutput{})
	pulumi.RegisterOutputType(StreamInputBlobArrayOutput{})
	pulumi.RegisterOutputType(StreamInputBlobMapOutput{})
}
