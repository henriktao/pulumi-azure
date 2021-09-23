// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Azure Spring Cloud Deployment with a Java runtime.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appplatform"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
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
// 		exampleSpringCloudService, err := appplatform.NewSpringCloudService(ctx, "exampleSpringCloudService", &appplatform.SpringCloudServiceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSpringCloudApp, err := appplatform.NewSpringCloudApp(ctx, "exampleSpringCloudApp", &appplatform.SpringCloudAppArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServiceName:       exampleSpringCloudService.Name,
// 			Identity: &appplatform.SpringCloudAppIdentityArgs{
// 				Type: pulumi.String("SystemAssigned"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appplatform.NewSpringCloudJavaDeployment(ctx, "exampleSpringCloudJavaDeployment", &appplatform.SpringCloudJavaDeploymentArgs{
// 			SpringCloudAppId: exampleSpringCloudApp.ID(),
// 			Cpu:              pulumi.Int(2),
// 			InstanceCount:    pulumi.Int(2),
// 			JvmOptions:       pulumi.String("-XX:+PrintGC"),
// 			MemoryInGb:       pulumi.Int(4),
// 			RuntimeVersion:   pulumi.String("Java_11"),
// 			EnvironmentVariables: pulumi.StringMap{
// 				"Foo": pulumi.String("Bar"),
// 				"Env": pulumi.String("Staging"),
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
// Spring Cloud Deployment can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/Spring/service1/apps/app1/deployments/deploy1
// ```
type SpringCloudJavaDeployment struct {
	pulumi.CustomResourceState

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
	Cpu pulumi.IntPtrOutput `pulumi:"cpu"`
	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions pulumi.StringPtrOutput `pulumi:"jvmOptions"`
	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
	MemoryInGb pulumi.IntPtrOutput `pulumi:"memoryInGb"`
	// Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
	RuntimeVersion pulumi.StringPtrOutput `pulumi:"runtimeVersion"`
	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringOutput `pulumi:"springCloudAppId"`
}

// NewSpringCloudJavaDeployment registers a new resource with the given unique name, arguments, and options.
func NewSpringCloudJavaDeployment(ctx *pulumi.Context,
	name string, args *SpringCloudJavaDeploymentArgs, opts ...pulumi.ResourceOption) (*SpringCloudJavaDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SpringCloudAppId == nil {
		return nil, errors.New("invalid value for required argument 'SpringCloudAppId'")
	}
	var resource SpringCloudJavaDeployment
	err := ctx.RegisterResource("azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringCloudJavaDeployment gets an existing SpringCloudJavaDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringCloudJavaDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringCloudJavaDeploymentState, opts ...pulumi.ResourceOption) (*SpringCloudJavaDeployment, error) {
	var resource SpringCloudJavaDeployment
	err := ctx.ReadResource("azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpringCloudJavaDeployment resources.
type springCloudJavaDeploymentState struct {
	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
	Cpu *int `pulumi:"cpu"`
	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
	InstanceCount *int `pulumi:"instanceCount"`
	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions *string `pulumi:"jvmOptions"`
	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
	MemoryInGb *int `pulumi:"memoryInGb"`
	// Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppId *string `pulumi:"springCloudAppId"`
}

type SpringCloudJavaDeploymentState struct {
	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
	Cpu pulumi.IntPtrInput
	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables pulumi.StringMapInput
	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
	InstanceCount pulumi.IntPtrInput
	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions pulumi.StringPtrInput
	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
	MemoryInGb pulumi.IntPtrInput
	// Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
	RuntimeVersion pulumi.StringPtrInput
	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringPtrInput
}

func (SpringCloudJavaDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudJavaDeploymentState)(nil)).Elem()
}

type springCloudJavaDeploymentArgs struct {
	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
	Cpu *int `pulumi:"cpu"`
	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
	InstanceCount *int `pulumi:"instanceCount"`
	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions *string `pulumi:"jvmOptions"`
	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
	MemoryInGb *int `pulumi:"memoryInGb"`
	// Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppId string `pulumi:"springCloudAppId"`
}

// The set of arguments for constructing a SpringCloudJavaDeployment resource.
type SpringCloudJavaDeploymentArgs struct {
	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are between `1` and `4`. Defaults to `1` if not specified.
	Cpu pulumi.IntPtrInput
	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	EnvironmentVariables pulumi.StringMapInput
	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between `1` and `500`. Defaults to `1` if not specified.
	InstanceCount pulumi.IntPtrInput
	// Specifies the jvm option of the Spring Cloud Deployment.
	JvmOptions pulumi.StringPtrInput
	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are between `1` and `8`. Defaults to `1` if not specified.
	MemoryInGb pulumi.IntPtrInput
	// Specifies the name of the Spring Cloud Deployment. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the runtime version of the Spring Cloud Deployment. Possible Values are `Java_8` and `Java_11`. Defaults to `Java_8`.
	RuntimeVersion pulumi.StringPtrInput
	// Specifies the id of the Spring Cloud Application in which to create the Deployment. Changing this forces a new resource to be created.
	SpringCloudAppId pulumi.StringInput
}

func (SpringCloudJavaDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springCloudJavaDeploymentArgs)(nil)).Elem()
}

type SpringCloudJavaDeploymentInput interface {
	pulumi.Input

	ToSpringCloudJavaDeploymentOutput() SpringCloudJavaDeploymentOutput
	ToSpringCloudJavaDeploymentOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentOutput
}

func (*SpringCloudJavaDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudJavaDeployment)(nil))
}

func (i *SpringCloudJavaDeployment) ToSpringCloudJavaDeploymentOutput() SpringCloudJavaDeploymentOutput {
	return i.ToSpringCloudJavaDeploymentOutputWithContext(context.Background())
}

func (i *SpringCloudJavaDeployment) ToSpringCloudJavaDeploymentOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudJavaDeploymentOutput)
}

func (i *SpringCloudJavaDeployment) ToSpringCloudJavaDeploymentPtrOutput() SpringCloudJavaDeploymentPtrOutput {
	return i.ToSpringCloudJavaDeploymentPtrOutputWithContext(context.Background())
}

func (i *SpringCloudJavaDeployment) ToSpringCloudJavaDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudJavaDeploymentPtrOutput)
}

type SpringCloudJavaDeploymentPtrInput interface {
	pulumi.Input

	ToSpringCloudJavaDeploymentPtrOutput() SpringCloudJavaDeploymentPtrOutput
	ToSpringCloudJavaDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentPtrOutput
}

type springCloudJavaDeploymentPtrType SpringCloudJavaDeploymentArgs

func (*springCloudJavaDeploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudJavaDeployment)(nil))
}

func (i *springCloudJavaDeploymentPtrType) ToSpringCloudJavaDeploymentPtrOutput() SpringCloudJavaDeploymentPtrOutput {
	return i.ToSpringCloudJavaDeploymentPtrOutputWithContext(context.Background())
}

func (i *springCloudJavaDeploymentPtrType) ToSpringCloudJavaDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudJavaDeploymentPtrOutput)
}

// SpringCloudJavaDeploymentArrayInput is an input type that accepts SpringCloudJavaDeploymentArray and SpringCloudJavaDeploymentArrayOutput values.
// You can construct a concrete instance of `SpringCloudJavaDeploymentArrayInput` via:
//
//          SpringCloudJavaDeploymentArray{ SpringCloudJavaDeploymentArgs{...} }
type SpringCloudJavaDeploymentArrayInput interface {
	pulumi.Input

	ToSpringCloudJavaDeploymentArrayOutput() SpringCloudJavaDeploymentArrayOutput
	ToSpringCloudJavaDeploymentArrayOutputWithContext(context.Context) SpringCloudJavaDeploymentArrayOutput
}

type SpringCloudJavaDeploymentArray []SpringCloudJavaDeploymentInput

func (SpringCloudJavaDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpringCloudJavaDeployment)(nil)).Elem()
}

func (i SpringCloudJavaDeploymentArray) ToSpringCloudJavaDeploymentArrayOutput() SpringCloudJavaDeploymentArrayOutput {
	return i.ToSpringCloudJavaDeploymentArrayOutputWithContext(context.Background())
}

func (i SpringCloudJavaDeploymentArray) ToSpringCloudJavaDeploymentArrayOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudJavaDeploymentArrayOutput)
}

// SpringCloudJavaDeploymentMapInput is an input type that accepts SpringCloudJavaDeploymentMap and SpringCloudJavaDeploymentMapOutput values.
// You can construct a concrete instance of `SpringCloudJavaDeploymentMapInput` via:
//
//          SpringCloudJavaDeploymentMap{ "key": SpringCloudJavaDeploymentArgs{...} }
type SpringCloudJavaDeploymentMapInput interface {
	pulumi.Input

	ToSpringCloudJavaDeploymentMapOutput() SpringCloudJavaDeploymentMapOutput
	ToSpringCloudJavaDeploymentMapOutputWithContext(context.Context) SpringCloudJavaDeploymentMapOutput
}

type SpringCloudJavaDeploymentMap map[string]SpringCloudJavaDeploymentInput

func (SpringCloudJavaDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpringCloudJavaDeployment)(nil)).Elem()
}

func (i SpringCloudJavaDeploymentMap) ToSpringCloudJavaDeploymentMapOutput() SpringCloudJavaDeploymentMapOutput {
	return i.ToSpringCloudJavaDeploymentMapOutputWithContext(context.Background())
}

func (i SpringCloudJavaDeploymentMap) ToSpringCloudJavaDeploymentMapOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringCloudJavaDeploymentMapOutput)
}

type SpringCloudJavaDeploymentOutput struct{ *pulumi.OutputState }

func (SpringCloudJavaDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpringCloudJavaDeployment)(nil))
}

func (o SpringCloudJavaDeploymentOutput) ToSpringCloudJavaDeploymentOutput() SpringCloudJavaDeploymentOutput {
	return o
}

func (o SpringCloudJavaDeploymentOutput) ToSpringCloudJavaDeploymentOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentOutput {
	return o
}

func (o SpringCloudJavaDeploymentOutput) ToSpringCloudJavaDeploymentPtrOutput() SpringCloudJavaDeploymentPtrOutput {
	return o.ToSpringCloudJavaDeploymentPtrOutputWithContext(context.Background())
}

func (o SpringCloudJavaDeploymentOutput) ToSpringCloudJavaDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SpringCloudJavaDeployment) *SpringCloudJavaDeployment {
		return &v
	}).(SpringCloudJavaDeploymentPtrOutput)
}

type SpringCloudJavaDeploymentPtrOutput struct{ *pulumi.OutputState }

func (SpringCloudJavaDeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpringCloudJavaDeployment)(nil))
}

func (o SpringCloudJavaDeploymentPtrOutput) ToSpringCloudJavaDeploymentPtrOutput() SpringCloudJavaDeploymentPtrOutput {
	return o
}

func (o SpringCloudJavaDeploymentPtrOutput) ToSpringCloudJavaDeploymentPtrOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentPtrOutput {
	return o
}

func (o SpringCloudJavaDeploymentPtrOutput) Elem() SpringCloudJavaDeploymentOutput {
	return o.ApplyT(func(v *SpringCloudJavaDeployment) SpringCloudJavaDeployment {
		if v != nil {
			return *v
		}
		var ret SpringCloudJavaDeployment
		return ret
	}).(SpringCloudJavaDeploymentOutput)
}

type SpringCloudJavaDeploymentArrayOutput struct{ *pulumi.OutputState }

func (SpringCloudJavaDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpringCloudJavaDeployment)(nil))
}

func (o SpringCloudJavaDeploymentArrayOutput) ToSpringCloudJavaDeploymentArrayOutput() SpringCloudJavaDeploymentArrayOutput {
	return o
}

func (o SpringCloudJavaDeploymentArrayOutput) ToSpringCloudJavaDeploymentArrayOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentArrayOutput {
	return o
}

func (o SpringCloudJavaDeploymentArrayOutput) Index(i pulumi.IntInput) SpringCloudJavaDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpringCloudJavaDeployment {
		return vs[0].([]SpringCloudJavaDeployment)[vs[1].(int)]
	}).(SpringCloudJavaDeploymentOutput)
}

type SpringCloudJavaDeploymentMapOutput struct{ *pulumi.OutputState }

func (SpringCloudJavaDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpringCloudJavaDeployment)(nil))
}

func (o SpringCloudJavaDeploymentMapOutput) ToSpringCloudJavaDeploymentMapOutput() SpringCloudJavaDeploymentMapOutput {
	return o
}

func (o SpringCloudJavaDeploymentMapOutput) ToSpringCloudJavaDeploymentMapOutputWithContext(ctx context.Context) SpringCloudJavaDeploymentMapOutput {
	return o
}

func (o SpringCloudJavaDeploymentMapOutput) MapIndex(k pulumi.StringInput) SpringCloudJavaDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpringCloudJavaDeployment {
		return vs[0].(map[string]SpringCloudJavaDeployment)[vs[1].(string)]
	}).(SpringCloudJavaDeploymentOutput)
}

func init() {
	pulumi.RegisterOutputType(SpringCloudJavaDeploymentOutput{})
	pulumi.RegisterOutputType(SpringCloudJavaDeploymentPtrOutput{})
	pulumi.RegisterOutputType(SpringCloudJavaDeploymentArrayOutput{})
	pulumi.RegisterOutputType(SpringCloudJavaDeploymentMapOutput{})
}
