// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Healthcare Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewService(ctx, "example", &healthcare.ServiceArgs{
// 			AccessPolicyObjectIds: pulumi.StringArray{
// 				pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
// 			},
// 			AuthenticationConfiguration: &healthcare.ServiceAuthenticationConfigurationArgs{
// 				Audience:          pulumi.String("https://azurehealthcareapis.com/"),
// 				Authority:         pulumi.String(fmt.Sprintf("%v%v%v%v%v%v", "https://login.microsoftonline.com/", "$", "%", "7Bdata.azurerm_client_config.current.tenant_id", "%", "7D")),
// 				SmartProxyEnabled: pulumi.Bool(true),
// 			},
// 			CorsConfiguration: &healthcare.ServiceCorsConfigurationArgs{
// 				AllowCredentials: pulumi.Bool(true),
// 				AllowedHeaders: pulumi.StringArray{
// 					pulumi.String("x-tempo-*"),
// 					pulumi.String("x-tempo2-*"),
// 				},
// 				AllowedMethods: pulumi.StringArray{
// 					pulumi.String("GET"),
// 					pulumi.String("PUT"),
// 				},
// 				AllowedOrigins: pulumi.StringArray{
// 					pulumi.String("http://www.example.com"),
// 					pulumi.String("http://www.example2.com"),
// 				},
// 				MaxAgeInSeconds: pulumi.Int(500),
// 			},
// 			CosmosdbThroughput: pulumi.Int(2000),
// 			Kind:               pulumi.String("fhir-R4"),
// 			Location:           pulumi.String("westus2"),
// 			ResourceGroupName:  pulumi.String("sample-resource-group"),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("testenv"),
// 				"purpose":     pulumi.String("AcceptanceTests"),
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
// Healthcare Service can be imported using the resource`id`, e.g.
//
// ```sh
//  $ pulumi import azure:healthcare/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource_group/providers/Microsoft.HealthcareApis/services/service_name
// ```
type Service struct {
	pulumi.CustomResourceState

	AccessPolicyObjectIds pulumi.StringArrayOutput `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationOutput `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationOutput `pulumi:"corsConfiguration"`
	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosdbKeyVaultKeyVersionlessId pulumi.StringPtrOutput `pulumi:"cosmosdbKeyVaultKeyVersionlessId"`
	// The provisioned throughput for the backing database. Range of `400`-`10000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrOutput `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether public network access is enabled or disabled for this service instance.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Service
	err := ctx.RegisterResource("azure:healthcare/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure:healthcare/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	AccessPolicyObjectIds []string `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration *ServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration *ServiceCorsConfiguration `pulumi:"corsConfiguration"`
	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosdbKeyVaultKeyVersionlessId *string `pulumi:"cosmosdbKeyVaultKeyVersionlessId"`
	// The provisioned throughput for the backing database. Range of `400`-`10000`. Defaults to `400`.
	CosmosdbThroughput *int `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location *string `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name *string `pulumi:"name"`
	// Whether public network access is enabled or disabled for this service instance.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ServiceState struct {
	AccessPolicyObjectIds pulumi.StringArrayInput
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationPtrInput
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationPtrInput
	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosdbKeyVaultKeyVersionlessId pulumi.StringPtrInput
	// The provisioned throughput for the backing database. Range of `400`-`10000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrInput
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringPtrInput
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringPtrInput
	// Whether public network access is enabled or disabled for this service instance.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	AccessPolicyObjectIds []string `pulumi:"accessPolicyObjectIds"`
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration *ServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// A `corsConfiguration` block as defined below.
	CorsConfiguration *ServiceCorsConfiguration `pulumi:"corsConfiguration"`
	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosdbKeyVaultKeyVersionlessId *string `pulumi:"cosmosdbKeyVaultKeyVersionlessId"`
	// The provisioned throughput for the backing database. Range of `400`-`10000`. Defaults to `400`.
	CosmosdbThroughput *int `pulumi:"cosmosdbThroughput"`
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind *string `pulumi:"kind"`
	// Specifies the supported Azure Region where the Service should be created.
	Location *string `pulumi:"location"`
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name *string `pulumi:"name"`
	// Whether public network access is enabled or disabled for this service instance.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	AccessPolicyObjectIds pulumi.StringArrayInput
	// An `authenticationConfiguration` block as defined below.
	AuthenticationConfiguration ServiceAuthenticationConfigurationPtrInput
	// A `corsConfiguration` block as defined below.
	CorsConfiguration ServiceCorsConfigurationPtrInput
	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosdbKeyVaultKeyVersionlessId pulumi.StringPtrInput
	// The provisioned throughput for the backing database. Range of `400`-`10000`. Defaults to `400`.
	CosmosdbThroughput pulumi.IntPtrInput
	// The type of the service. Values at time of publication are: `fhir`, `fhir-Stu3` and `fhir-R4`. Default value is `fhir`.
	Kind pulumi.StringPtrInput
	// Specifies the supported Azure Region where the Service should be created.
	Location pulumi.StringPtrInput
	// The name of the service instance. Used for service endpoint, must be unique within the audience.
	Name pulumi.StringPtrInput
	// Whether public network access is enabled or disabled for this service instance.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the Resource Group in which to create the Service.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
