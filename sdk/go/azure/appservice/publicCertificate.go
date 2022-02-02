// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Service Public Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/base64"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/appservice"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func filebase64OrPanic(path string) pulumi.StringPtrInput {
// 	if fileData, err := ioutil.ReadFile(path); err == nil {
// 		return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
// 	} else {
// 		panic(err.Error())
// 	}
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePlan, err := appservice.NewPlan(ctx, "examplePlan", &appservice.PlanArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku: &appservice.PlanSkuArgs{
// 				Tier: pulumi.String("Standard"),
// 				Size: pulumi.String("S1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAppService, err := appservice.NewAppService(ctx, "exampleAppService", &appservice.AppServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AppServicePlanId:  examplePlan.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appservice.NewPublicCertificate(ctx, "examplePublicCertificate", &appservice.PublicCertificateArgs{
// 			ResourceGroupName:   exampleResourceGroup.Name,
// 			AppServiceName:      exampleAppService.Name,
// 			CertificateName:     pulumi.String("example-public-certificate"),
// 			CertificateLocation: pulumi.String("Unknown"),
// 			Blob:                filebase64OrPanic("app_service_public_certificate.cer"),
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
// App Service Public Certificates can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:appservice/publicCertificate:PublicCertificate example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/site1/publicCertificates/publicCertificate1
// ```
type PublicCertificate struct {
	pulumi.CustomResourceState

	// The name of the App Service. Changing this forces a new App Service Public Certificate to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The base64-encoded contents of the certificate. Changing this forces a new App Service Public Certificate to be created.
	Blob pulumi.StringOutput `pulumi:"blob"`
	// The location of the certificate. Possible values are `CurrentUserMy`, `LocalMachineMy` and `Unknown`.
	CertificateLocation pulumi.StringOutput `pulumi:"certificateLocation"`
	// The name of the public certificate. Changing this forces a new App Service Public Certificate to be created.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// The name of the Resource Group where the App Service Public Certificate should exist. Changing this forces a new App Service Public Certificate to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The thumbprint of the public certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
}

// NewPublicCertificate registers a new resource with the given unique name, arguments, and options.
func NewPublicCertificate(ctx *pulumi.Context,
	name string, args *PublicCertificateArgs, opts ...pulumi.ResourceOption) (*PublicCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppServiceName == nil {
		return nil, errors.New("invalid value for required argument 'AppServiceName'")
	}
	if args.Blob == nil {
		return nil, errors.New("invalid value for required argument 'Blob'")
	}
	if args.CertificateLocation == nil {
		return nil, errors.New("invalid value for required argument 'CertificateLocation'")
	}
	if args.CertificateName == nil {
		return nil, errors.New("invalid value for required argument 'CertificateName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PublicCertificate
	err := ctx.RegisterResource("azure:appservice/publicCertificate:PublicCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicCertificate gets an existing PublicCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicCertificateState, opts ...pulumi.ResourceOption) (*PublicCertificate, error) {
	var resource PublicCertificate
	err := ctx.ReadResource("azure:appservice/publicCertificate:PublicCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicCertificate resources.
type publicCertificateState struct {
	// The name of the App Service. Changing this forces a new App Service Public Certificate to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The base64-encoded contents of the certificate. Changing this forces a new App Service Public Certificate to be created.
	Blob *string `pulumi:"blob"`
	// The location of the certificate. Possible values are `CurrentUserMy`, `LocalMachineMy` and `Unknown`.
	CertificateLocation *string `pulumi:"certificateLocation"`
	// The name of the public certificate. Changing this forces a new App Service Public Certificate to be created.
	CertificateName *string `pulumi:"certificateName"`
	// The name of the Resource Group where the App Service Public Certificate should exist. Changing this forces a new App Service Public Certificate to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The thumbprint of the public certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

type PublicCertificateState struct {
	// The name of the App Service. Changing this forces a new App Service Public Certificate to be created.
	AppServiceName pulumi.StringPtrInput
	// The base64-encoded contents of the certificate. Changing this forces a new App Service Public Certificate to be created.
	Blob pulumi.StringPtrInput
	// The location of the certificate. Possible values are `CurrentUserMy`, `LocalMachineMy` and `Unknown`.
	CertificateLocation pulumi.StringPtrInput
	// The name of the public certificate. Changing this forces a new App Service Public Certificate to be created.
	CertificateName pulumi.StringPtrInput
	// The name of the Resource Group where the App Service Public Certificate should exist. Changing this forces a new App Service Public Certificate to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The thumbprint of the public certificate.
	Thumbprint pulumi.StringPtrInput
}

func (PublicCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicCertificateState)(nil)).Elem()
}

type publicCertificateArgs struct {
	// The name of the App Service. Changing this forces a new App Service Public Certificate to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The base64-encoded contents of the certificate. Changing this forces a new App Service Public Certificate to be created.
	Blob string `pulumi:"blob"`
	// The location of the certificate. Possible values are `CurrentUserMy`, `LocalMachineMy` and `Unknown`.
	CertificateLocation string `pulumi:"certificateLocation"`
	// The name of the public certificate. Changing this forces a new App Service Public Certificate to be created.
	CertificateName string `pulumi:"certificateName"`
	// The name of the Resource Group where the App Service Public Certificate should exist. Changing this forces a new App Service Public Certificate to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PublicCertificate resource.
type PublicCertificateArgs struct {
	// The name of the App Service. Changing this forces a new App Service Public Certificate to be created.
	AppServiceName pulumi.StringInput
	// The base64-encoded contents of the certificate. Changing this forces a new App Service Public Certificate to be created.
	Blob pulumi.StringInput
	// The location of the certificate. Possible values are `CurrentUserMy`, `LocalMachineMy` and `Unknown`.
	CertificateLocation pulumi.StringInput
	// The name of the public certificate. Changing this forces a new App Service Public Certificate to be created.
	CertificateName pulumi.StringInput
	// The name of the Resource Group where the App Service Public Certificate should exist. Changing this forces a new App Service Public Certificate to be created.
	ResourceGroupName pulumi.StringInput
}

func (PublicCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicCertificateArgs)(nil)).Elem()
}

type PublicCertificateInput interface {
	pulumi.Input

	ToPublicCertificateOutput() PublicCertificateOutput
	ToPublicCertificateOutputWithContext(ctx context.Context) PublicCertificateOutput
}

func (*PublicCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicCertificate)(nil)).Elem()
}

func (i *PublicCertificate) ToPublicCertificateOutput() PublicCertificateOutput {
	return i.ToPublicCertificateOutputWithContext(context.Background())
}

func (i *PublicCertificate) ToPublicCertificateOutputWithContext(ctx context.Context) PublicCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicCertificateOutput)
}

// PublicCertificateArrayInput is an input type that accepts PublicCertificateArray and PublicCertificateArrayOutput values.
// You can construct a concrete instance of `PublicCertificateArrayInput` via:
//
//          PublicCertificateArray{ PublicCertificateArgs{...} }
type PublicCertificateArrayInput interface {
	pulumi.Input

	ToPublicCertificateArrayOutput() PublicCertificateArrayOutput
	ToPublicCertificateArrayOutputWithContext(context.Context) PublicCertificateArrayOutput
}

type PublicCertificateArray []PublicCertificateInput

func (PublicCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicCertificate)(nil)).Elem()
}

func (i PublicCertificateArray) ToPublicCertificateArrayOutput() PublicCertificateArrayOutput {
	return i.ToPublicCertificateArrayOutputWithContext(context.Background())
}

func (i PublicCertificateArray) ToPublicCertificateArrayOutputWithContext(ctx context.Context) PublicCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicCertificateArrayOutput)
}

// PublicCertificateMapInput is an input type that accepts PublicCertificateMap and PublicCertificateMapOutput values.
// You can construct a concrete instance of `PublicCertificateMapInput` via:
//
//          PublicCertificateMap{ "key": PublicCertificateArgs{...} }
type PublicCertificateMapInput interface {
	pulumi.Input

	ToPublicCertificateMapOutput() PublicCertificateMapOutput
	ToPublicCertificateMapOutputWithContext(context.Context) PublicCertificateMapOutput
}

type PublicCertificateMap map[string]PublicCertificateInput

func (PublicCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicCertificate)(nil)).Elem()
}

func (i PublicCertificateMap) ToPublicCertificateMapOutput() PublicCertificateMapOutput {
	return i.ToPublicCertificateMapOutputWithContext(context.Background())
}

func (i PublicCertificateMap) ToPublicCertificateMapOutputWithContext(ctx context.Context) PublicCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicCertificateMapOutput)
}

type PublicCertificateOutput struct{ *pulumi.OutputState }

func (PublicCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicCertificate)(nil)).Elem()
}

func (o PublicCertificateOutput) ToPublicCertificateOutput() PublicCertificateOutput {
	return o
}

func (o PublicCertificateOutput) ToPublicCertificateOutputWithContext(ctx context.Context) PublicCertificateOutput {
	return o
}

type PublicCertificateArrayOutput struct{ *pulumi.OutputState }

func (PublicCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicCertificate)(nil)).Elem()
}

func (o PublicCertificateArrayOutput) ToPublicCertificateArrayOutput() PublicCertificateArrayOutput {
	return o
}

func (o PublicCertificateArrayOutput) ToPublicCertificateArrayOutputWithContext(ctx context.Context) PublicCertificateArrayOutput {
	return o
}

func (o PublicCertificateArrayOutput) Index(i pulumi.IntInput) PublicCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicCertificate {
		return vs[0].([]*PublicCertificate)[vs[1].(int)]
	}).(PublicCertificateOutput)
}

type PublicCertificateMapOutput struct{ *pulumi.OutputState }

func (PublicCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicCertificate)(nil)).Elem()
}

func (o PublicCertificateMapOutput) ToPublicCertificateMapOutput() PublicCertificateMapOutput {
	return o
}

func (o PublicCertificateMapOutput) ToPublicCertificateMapOutputWithContext(ctx context.Context) PublicCertificateMapOutput {
	return o
}

func (o PublicCertificateMapOutput) MapIndex(k pulumi.StringInput) PublicCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicCertificate {
		return vs[0].(map[string]*PublicCertificate)[vs[1].(string)]
	}).(PublicCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicCertificateInput)(nil)).Elem(), &PublicCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicCertificateArrayInput)(nil)).Elem(), PublicCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicCertificateMapInput)(nil)).Elem(), PublicCertificateMap{})
	pulumi.RegisterOutputType(PublicCertificateOutput{})
	pulumi.RegisterOutputType(PublicCertificateArrayOutput{})
	pulumi.RegisterOutputType(PublicCertificateMapOutput{})
}
