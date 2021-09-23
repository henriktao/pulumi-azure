// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Logic App Integration Account Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/logicapps"
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
// 		exampleIntegrationAccount, err := logicapps.NewIntegrationAccount(ctx, "exampleIntegrationAccount", &logicapps.IntegrationAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logicapps.NewIntegrationAccountCertificate(ctx, "exampleIntegrationAccountCertificate", &logicapps.IntegrationAccountCertificateArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			IntegrationAccountName: exampleIntegrationAccount.Name,
// 			PublicCertificate:      pulumi.String("MIIDbzCCAlegAwIBAgIJAIzjRD36sIbbMA0GCSqGSIb3DQEBCwUAME0xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApTb21lLVN0YXRlMRIwEAYDVQQKDAl0ZXJyYWZvcm0xFTATBgNVBAMMDHRlcnJhZm9ybS5pbzAgFw0xNzA0MjEyMDA1MjdaGA8yMTE3MDMyODIwMDUyN1owTTELMAkGA1UEBhMCVVMxEzARBgNVBAgMClNvbWUtU3RhdGUxEjAQBgNVBAoMCXRlcnJhZm9ybTEVMBMGA1UEAwwMdGVycmFmb3JtLmlvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3L9L5szT4+FLykTFNyyPjy/k3BQTYAfRQzP2dhnsuUKm3cdPC0NyZ+wEXIUGhoDO2YG6EYChOl8fsDqDOjloSUGKqYw++nlpHIuUgJx8IxxG2XkALCjFU7EmF+w7kn76d0ezpEIYxnLP+KG2DVornoEt1aLhv1MLmpgEZZPhDbMSLhSYWeTVRMayXLwqtfgnDumQSB+8d/1JuJqrSI4pD12JozVThzb6hsjfb6RMX4epPmrGn0PbTPEEA6awmsxBCXB0s13nNQt/O0hLM2agwvAyozilQV+s616Ckgk6DJoUkqZhDy7vPYMIRSr98fBws6zkrV6tTLjmD8xAvobePQIDAQABo1AwTjAdBgNVHQ4EFgQUXIqO421zMMmbcRRX9wctZFCQuPIwHwYDVR0jBBgwFoAUXIqO421zMMmbcRRX9wctZFCQuPIwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAr82NeT3BYJOKLlUL6Om5LjUF66ewcJjG9ltdvyQwVneMcq7t5UAPxgChzqNRVk4da8PzkXpjBJyWezHupdJNX3XqeUk2kSxqQ6/gmhqvfI3y7djrwoO6jvMEY26WqtkTNORWDP3THJJVimC3zV+KMU5UBVrEzhOVhHSU709lBP75o0BBn3xGsPqSq9k8IotIFfyAc6a+XP3+ZMpvh7wqAUml7vWa5wlcXExCx39h1balfDSLGNC4swWPCp9AMnQR0p+vMay9hNP1Eh+9QYUai14d5KS3cFV+KxE1cJR5HD/iLltnnOEbpMsB0eVOZWkFvE7Y5lW0oVSAfin5TwTJMQ=="),
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
// Logic App Integration Account Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/integrationAccountCertificate:IntegrationAccountCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/certificate1
// ```
type IntegrationAccountCertificate struct {
	pulumi.CustomResourceState

	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Certificate to be created.
	IntegrationAccountName pulumi.StringOutput `pulumi:"integrationAccountName"`
	// A `keyVaultKey` block as documented below.
	KeyVaultKey IntegrationAccountCertificateKeyVaultKeyPtrOutput `pulumi:"keyVaultKey"`
	// A JSON mapping of any Metadata for this Logic App Integration Account Certificate.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Certificate. Changing this forces a new Logic App Integration Account Certificate to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The public certificate for the Logic App Integration Account Certificate.
	PublicCertificate pulumi.StringPtrOutput `pulumi:"publicCertificate"`
	// The name of the Resource Group where the Logic App Integration Account Certificate should exist. Changing this forces a new Logic App Integration Account Certificate to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIntegrationAccountCertificate registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, args *IntegrationAccountCertificateArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationAccountCertificate
	err := ctx.RegisterResource("azure:logicapps/integrationAccountCertificate:IntegrationAccountCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountCertificate gets an existing IntegrationAccountCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountCertificateState, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	var resource IntegrationAccountCertificate
	err := ctx.ReadResource("azure:logicapps/integrationAccountCertificate:IntegrationAccountCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountCertificate resources.
type integrationAccountCertificateState struct {
	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Certificate to be created.
	IntegrationAccountName *string `pulumi:"integrationAccountName"`
	// A `keyVaultKey` block as documented below.
	KeyVaultKey *IntegrationAccountCertificateKeyVaultKey `pulumi:"keyVaultKey"`
	// A JSON mapping of any Metadata for this Logic App Integration Account Certificate.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Certificate. Changing this forces a new Logic App Integration Account Certificate to be created.
	Name *string `pulumi:"name"`
	// The public certificate for the Logic App Integration Account Certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The name of the Resource Group where the Logic App Integration Account Certificate should exist. Changing this forces a new Logic App Integration Account Certificate to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IntegrationAccountCertificateState struct {
	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Certificate to be created.
	IntegrationAccountName pulumi.StringPtrInput
	// A `keyVaultKey` block as documented below.
	KeyVaultKey IntegrationAccountCertificateKeyVaultKeyPtrInput
	// A JSON mapping of any Metadata for this Logic App Integration Account Certificate.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Logic App Integration Account Certificate. Changing this forces a new Logic App Integration Account Certificate to be created.
	Name pulumi.StringPtrInput
	// The public certificate for the Logic App Integration Account Certificate.
	PublicCertificate pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Certificate should exist. Changing this forces a new Logic App Integration Account Certificate to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IntegrationAccountCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateState)(nil)).Elem()
}

type integrationAccountCertificateArgs struct {
	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Certificate to be created.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// A `keyVaultKey` block as documented below.
	KeyVaultKey *IntegrationAccountCertificateKeyVaultKey `pulumi:"keyVaultKey"`
	// A JSON mapping of any Metadata for this Logic App Integration Account Certificate.
	Metadata *string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Certificate. Changing this forces a new Logic App Integration Account Certificate to be created.
	Name *string `pulumi:"name"`
	// The public certificate for the Logic App Integration Account Certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The name of the Resource Group where the Logic App Integration Account Certificate should exist. Changing this forces a new Logic App Integration Account Certificate to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IntegrationAccountCertificate resource.
type IntegrationAccountCertificateArgs struct {
	// The name of the Logic App Integration Account. Changing this forces a new Logic App Integration Account Certificate to be created.
	IntegrationAccountName pulumi.StringInput
	// A `keyVaultKey` block as documented below.
	KeyVaultKey IntegrationAccountCertificateKeyVaultKeyPtrInput
	// A JSON mapping of any Metadata for this Logic App Integration Account Certificate.
	Metadata pulumi.StringPtrInput
	// The name which should be used for this Logic App Integration Account Certificate. Changing this forces a new Logic App Integration Account Certificate to be created.
	Name pulumi.StringPtrInput
	// The public certificate for the Logic App Integration Account Certificate.
	PublicCertificate pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Certificate should exist. Changing this forces a new Logic App Integration Account Certificate to be created.
	ResourceGroupName pulumi.StringInput
}

func (IntegrationAccountCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateArgs)(nil)).Elem()
}

type IntegrationAccountCertificateInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput
	ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput
}

func (*IntegrationAccountCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return i.ToIntegrationAccountCertificateOutputWithContext(context.Background())
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateOutput)
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificatePtrOutput() IntegrationAccountCertificatePtrOutput {
	return i.ToIntegrationAccountCertificatePtrOutputWithContext(context.Background())
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificatePtrOutputWithContext(ctx context.Context) IntegrationAccountCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificatePtrOutput)
}

type IntegrationAccountCertificatePtrInput interface {
	pulumi.Input

	ToIntegrationAccountCertificatePtrOutput() IntegrationAccountCertificatePtrOutput
	ToIntegrationAccountCertificatePtrOutputWithContext(ctx context.Context) IntegrationAccountCertificatePtrOutput
}

type integrationAccountCertificatePtrType IntegrationAccountCertificateArgs

func (*integrationAccountCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountCertificate)(nil))
}

func (i *integrationAccountCertificatePtrType) ToIntegrationAccountCertificatePtrOutput() IntegrationAccountCertificatePtrOutput {
	return i.ToIntegrationAccountCertificatePtrOutputWithContext(context.Background())
}

func (i *integrationAccountCertificatePtrType) ToIntegrationAccountCertificatePtrOutputWithContext(ctx context.Context) IntegrationAccountCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificatePtrOutput)
}

// IntegrationAccountCertificateArrayInput is an input type that accepts IntegrationAccountCertificateArray and IntegrationAccountCertificateArrayOutput values.
// You can construct a concrete instance of `IntegrationAccountCertificateArrayInput` via:
//
//          IntegrationAccountCertificateArray{ IntegrationAccountCertificateArgs{...} }
type IntegrationAccountCertificateArrayInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateArrayOutput() IntegrationAccountCertificateArrayOutput
	ToIntegrationAccountCertificateArrayOutputWithContext(context.Context) IntegrationAccountCertificateArrayOutput
}

type IntegrationAccountCertificateArray []IntegrationAccountCertificateInput

func (IntegrationAccountCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAccountCertificate)(nil)).Elem()
}

func (i IntegrationAccountCertificateArray) ToIntegrationAccountCertificateArrayOutput() IntegrationAccountCertificateArrayOutput {
	return i.ToIntegrationAccountCertificateArrayOutputWithContext(context.Background())
}

func (i IntegrationAccountCertificateArray) ToIntegrationAccountCertificateArrayOutputWithContext(ctx context.Context) IntegrationAccountCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateArrayOutput)
}

// IntegrationAccountCertificateMapInput is an input type that accepts IntegrationAccountCertificateMap and IntegrationAccountCertificateMapOutput values.
// You can construct a concrete instance of `IntegrationAccountCertificateMapInput` via:
//
//          IntegrationAccountCertificateMap{ "key": IntegrationAccountCertificateArgs{...} }
type IntegrationAccountCertificateMapInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateMapOutput() IntegrationAccountCertificateMapOutput
	ToIntegrationAccountCertificateMapOutputWithContext(context.Context) IntegrationAccountCertificateMapOutput
}

type IntegrationAccountCertificateMap map[string]IntegrationAccountCertificateInput

func (IntegrationAccountCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAccountCertificate)(nil)).Elem()
}

func (i IntegrationAccountCertificateMap) ToIntegrationAccountCertificateMapOutput() IntegrationAccountCertificateMapOutput {
	return i.ToIntegrationAccountCertificateMapOutputWithContext(context.Background())
}

func (i IntegrationAccountCertificateMap) ToIntegrationAccountCertificateMapOutputWithContext(ctx context.Context) IntegrationAccountCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateMapOutput)
}

type IntegrationAccountCertificateOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificatePtrOutput() IntegrationAccountCertificatePtrOutput {
	return o.ToIntegrationAccountCertificatePtrOutputWithContext(context.Background())
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificatePtrOutputWithContext(ctx context.Context) IntegrationAccountCertificatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationAccountCertificate) *IntegrationAccountCertificate {
		return &v
	}).(IntegrationAccountCertificatePtrOutput)
}

type IntegrationAccountCertificatePtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificatePtrOutput) ToIntegrationAccountCertificatePtrOutput() IntegrationAccountCertificatePtrOutput {
	return o
}

func (o IntegrationAccountCertificatePtrOutput) ToIntegrationAccountCertificatePtrOutputWithContext(ctx context.Context) IntegrationAccountCertificatePtrOutput {
	return o
}

func (o IntegrationAccountCertificatePtrOutput) Elem() IntegrationAccountCertificateOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) IntegrationAccountCertificate {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountCertificate
		return ret
	}).(IntegrationAccountCertificateOutput)
}

type IntegrationAccountCertificateArrayOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificateArrayOutput) ToIntegrationAccountCertificateArrayOutput() IntegrationAccountCertificateArrayOutput {
	return o
}

func (o IntegrationAccountCertificateArrayOutput) ToIntegrationAccountCertificateArrayOutputWithContext(ctx context.Context) IntegrationAccountCertificateArrayOutput {
	return o
}

func (o IntegrationAccountCertificateArrayOutput) Index(i pulumi.IntInput) IntegrationAccountCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationAccountCertificate {
		return vs[0].([]IntegrationAccountCertificate)[vs[1].(int)]
	}).(IntegrationAccountCertificateOutput)
}

type IntegrationAccountCertificateMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificateMapOutput) ToIntegrationAccountCertificateMapOutput() IntegrationAccountCertificateMapOutput {
	return o
}

func (o IntegrationAccountCertificateMapOutput) ToIntegrationAccountCertificateMapOutputWithContext(ctx context.Context) IntegrationAccountCertificateMapOutput {
	return o
}

func (o IntegrationAccountCertificateMapOutput) MapIndex(k pulumi.StringInput) IntegrationAccountCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IntegrationAccountCertificate {
		return vs[0].(map[string]IntegrationAccountCertificate)[vs[1].(string)]
	}).(IntegrationAccountCertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountCertificateOutput{})
	pulumi.RegisterOutputType(IntegrationAccountCertificatePtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountCertificateArrayOutput{})
	pulumi.RegisterOutputType(IntegrationAccountCertificateMapOutput{})
}
