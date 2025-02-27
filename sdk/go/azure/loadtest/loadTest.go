// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadtest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Load Test.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/loadtest"
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
// 		_, err = loadtest.NewLoadTest(ctx, "exampleLoadTest", &loadtest.LoadTestArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          pulumi.String("West Europe"),
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
// Load tests can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:loadtest/loadTest:LoadTest example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.LoadTestService/loadtests/example
// ```
type LoadTest struct {
	pulumi.CustomResourceState

	// Public URI of the Data Plane.
	DataplaneUri pulumi.StringOutput `pulumi:"dataplaneUri"`
	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name which should be used for this Load Test. Changing this forces a new Load Test to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Load Test.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLoadTest registers a new resource with the given unique name, arguments, and options.
func NewLoadTest(ctx *pulumi.Context,
	name string, args *LoadTestArgs, opts ...pulumi.ResourceOption) (*LoadTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LoadTest
	err := ctx.RegisterResource("azure:loadtest/loadTest:LoadTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadTest gets an existing LoadTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadTestState, opts ...pulumi.ResourceOption) (*LoadTest, error) {
	var resource LoadTest
	err := ctx.ReadResource("azure:loadtest/loadTest:LoadTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadTest resources.
type loadTestState struct {
	// Public URI of the Data Plane.
	DataplaneUri *string `pulumi:"dataplaneUri"`
	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Load Test. Changing this forces a new Load Test to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Load Test.
	Tags map[string]string `pulumi:"tags"`
}

type LoadTestState struct {
	// Public URI of the Data Plane.
	DataplaneUri pulumi.StringPtrInput
	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Load Test. Changing this forces a new Load Test to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Load Test.
	Tags pulumi.StringMapInput
}

func (LoadTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadTestState)(nil)).Elem()
}

type loadTestArgs struct {
	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location *string `pulumi:"location"`
	// The name which should be used for this Load Test. Changing this forces a new Load Test to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags which should be assigned to the Load Test.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LoadTest resource.
type LoadTestArgs struct {
	// The Azure Region where the Load Test should exist. Changing this forces a new Load Test to be created.
	Location pulumi.StringPtrInput
	// The name which should be used for this Load Test. Changing this forces a new Load Test to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Load Test should exist. Changing this forces a new Load Test to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags which should be assigned to the Load Test.
	Tags pulumi.StringMapInput
}

func (LoadTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadTestArgs)(nil)).Elem()
}

type LoadTestInput interface {
	pulumi.Input

	ToLoadTestOutput() LoadTestOutput
	ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput
}

func (*LoadTest) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadTest)(nil)).Elem()
}

func (i *LoadTest) ToLoadTestOutput() LoadTestOutput {
	return i.ToLoadTestOutputWithContext(context.Background())
}

func (i *LoadTest) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadTestOutput)
}

// LoadTestArrayInput is an input type that accepts LoadTestArray and LoadTestArrayOutput values.
// You can construct a concrete instance of `LoadTestArrayInput` via:
//
//          LoadTestArray{ LoadTestArgs{...} }
type LoadTestArrayInput interface {
	pulumi.Input

	ToLoadTestArrayOutput() LoadTestArrayOutput
	ToLoadTestArrayOutputWithContext(context.Context) LoadTestArrayOutput
}

type LoadTestArray []LoadTestInput

func (LoadTestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadTest)(nil)).Elem()
}

func (i LoadTestArray) ToLoadTestArrayOutput() LoadTestArrayOutput {
	return i.ToLoadTestArrayOutputWithContext(context.Background())
}

func (i LoadTestArray) ToLoadTestArrayOutputWithContext(ctx context.Context) LoadTestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadTestArrayOutput)
}

// LoadTestMapInput is an input type that accepts LoadTestMap and LoadTestMapOutput values.
// You can construct a concrete instance of `LoadTestMapInput` via:
//
//          LoadTestMap{ "key": LoadTestArgs{...} }
type LoadTestMapInput interface {
	pulumi.Input

	ToLoadTestMapOutput() LoadTestMapOutput
	ToLoadTestMapOutputWithContext(context.Context) LoadTestMapOutput
}

type LoadTestMap map[string]LoadTestInput

func (LoadTestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadTest)(nil)).Elem()
}

func (i LoadTestMap) ToLoadTestMapOutput() LoadTestMapOutput {
	return i.ToLoadTestMapOutputWithContext(context.Background())
}

func (i LoadTestMap) ToLoadTestMapOutputWithContext(ctx context.Context) LoadTestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadTestMapOutput)
}

type LoadTestOutput struct{ *pulumi.OutputState }

func (LoadTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadTest)(nil)).Elem()
}

func (o LoadTestOutput) ToLoadTestOutput() LoadTestOutput {
	return o
}

func (o LoadTestOutput) ToLoadTestOutputWithContext(ctx context.Context) LoadTestOutput {
	return o
}

type LoadTestArrayOutput struct{ *pulumi.OutputState }

func (LoadTestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadTest)(nil)).Elem()
}

func (o LoadTestArrayOutput) ToLoadTestArrayOutput() LoadTestArrayOutput {
	return o
}

func (o LoadTestArrayOutput) ToLoadTestArrayOutputWithContext(ctx context.Context) LoadTestArrayOutput {
	return o
}

func (o LoadTestArrayOutput) Index(i pulumi.IntInput) LoadTestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadTest {
		return vs[0].([]*LoadTest)[vs[1].(int)]
	}).(LoadTestOutput)
}

type LoadTestMapOutput struct{ *pulumi.OutputState }

func (LoadTestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadTest)(nil)).Elem()
}

func (o LoadTestMapOutput) ToLoadTestMapOutput() LoadTestMapOutput {
	return o
}

func (o LoadTestMapOutput) ToLoadTestMapOutputWithContext(ctx context.Context) LoadTestMapOutput {
	return o
}

func (o LoadTestMapOutput) MapIndex(k pulumi.StringInput) LoadTestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadTest {
		return vs[0].(map[string]*LoadTest)[vs[1].(string)]
	}).(LoadTestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadTestInput)(nil)).Elem(), &LoadTest{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadTestArrayInput)(nil)).Elem(), LoadTestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadTestMapInput)(nil)).Elem(), LoadTestMap{})
	pulumi.RegisterOutputType(LoadTestOutput{})
	pulumi.RegisterOutputType(LoadTestArrayOutput{})
	pulumi.RegisterOutputType(LoadTestMapOutput{})
}
