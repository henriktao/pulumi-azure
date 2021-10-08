// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Linked Service (connection) between a SFTP Server and Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/datafactory"
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
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceSftp(ctx, "exampleLinkedServiceSftp", &datafactory.LinkedServiceSftpArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Basic"),
// 			Host:               pulumi.String("http://www.bing.com"),
// 			Port:               pulumi.Int(22),
// 			Username:           pulumi.String("foo"),
// 			Password:           pulumi.String("bar"),
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
// Data Factory Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceSftp:LinkedServiceSftp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceSftp struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The SFTP server hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint pulumi.StringPtrOutput `pulumi:"hostKeyFingerprint"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringOutput `pulumi:"password"`
	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port pulumi.IntOutput `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Whether to validate host key fingerprint while connecting. If set to `false`, `hostKeyFingerprint` must also be set.
	SkipHostKeyValidation pulumi.BoolPtrOutput `pulumi:"skipHostKeyValidation"`
	// The username used to log on to the SFTP server.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewLinkedServiceSftp registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSftp(ctx *pulumi.Context,
	name string, args *LinkedServiceSftpArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSftp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationType'")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource LinkedServiceSftp
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSftp:LinkedServiceSftp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSftp gets an existing LinkedServiceSftp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSftp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSftpState, opts ...pulumi.ResourceOption) (*LinkedServiceSftp, error) {
	var resource LinkedServiceSftp
	err := ctx.ReadResource("azure:datafactory/linkedServiceSftp:LinkedServiceSftp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSftp resources.
type linkedServiceSftpState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType *string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The SFTP server hostname.
	Host *string `pulumi:"host"`
	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint *string `pulumi:"hostKeyFingerprint"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password *string `pulumi:"password"`
	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port *int `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Whether to validate host key fingerprint while connecting. If set to `false`, `hostKeyFingerprint` must also be set.
	SkipHostKeyValidation *bool `pulumi:"skipHostKeyValidation"`
	// The username used to log on to the SFTP server.
	Username *string `pulumi:"username"`
}

type LinkedServiceSftpState struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The SFTP server hostname.
	Host pulumi.StringPtrInput
	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringPtrInput
	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port pulumi.IntPtrInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// Whether to validate host key fingerprint while connecting. If set to `false`, `hostKeyFingerprint` must also be set.
	SkipHostKeyValidation pulumi.BoolPtrInput
	// The username used to log on to the SFTP server.
	Username pulumi.StringPtrInput
}

func (LinkedServiceSftpState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSftpState)(nil)).Elem()
}

type linkedServiceSftpArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []string `pulumi:"annotations"`
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType string `pulumi:"authenticationType"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service.
	Description *string `pulumi:"description"`
	// The SFTP server hostname.
	Host string `pulumi:"host"`
	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint *string `pulumi:"hostKeyFingerprint"`
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]string `pulumi:"parameters"`
	// Password to logon to the SFTP Server for Basic Authentication.
	Password string `pulumi:"password"`
	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port int `pulumi:"port"`
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Whether to validate host key fingerprint while connecting. If set to `false`, `hostKeyFingerprint` must also be set.
	SkipHostKeyValidation *bool `pulumi:"skipHostKeyValidation"`
	// The username used to log on to the SFTP server.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a LinkedServiceSftp resource.
type LinkedServiceSftpArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations pulumi.StringArrayInput
	// The type of authentication used to connect to the web table source. Valid options are `Anonymous`, `Basic` and `ClientCertificate`.
	AuthenticationType pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service.
	Description pulumi.StringPtrInput
	// The SFTP server hostname.
	Host pulumi.StringInput
	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters pulumi.StringMapInput
	// Password to logon to the SFTP Server for Basic Authentication.
	Password pulumi.StringInput
	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port pulumi.IntInput
	// The name of the resource group in which to create the Data Factory Linked Service. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// Whether to validate host key fingerprint while connecting. If set to `false`, `hostKeyFingerprint` must also be set.
	SkipHostKeyValidation pulumi.BoolPtrInput
	// The username used to log on to the SFTP server.
	Username pulumi.StringInput
}

func (LinkedServiceSftpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSftpArgs)(nil)).Elem()
}

type LinkedServiceSftpInput interface {
	pulumi.Input

	ToLinkedServiceSftpOutput() LinkedServiceSftpOutput
	ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput
}

func (*LinkedServiceSftp) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSftp)(nil))
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpOutput() LinkedServiceSftpOutput {
	return i.ToLinkedServiceSftpOutputWithContext(context.Background())
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpOutput)
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpPtrOutput() LinkedServiceSftpPtrOutput {
	return i.ToLinkedServiceSftpPtrOutputWithContext(context.Background())
}

func (i *LinkedServiceSftp) ToLinkedServiceSftpPtrOutputWithContext(ctx context.Context) LinkedServiceSftpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpPtrOutput)
}

type LinkedServiceSftpPtrInput interface {
	pulumi.Input

	ToLinkedServiceSftpPtrOutput() LinkedServiceSftpPtrOutput
	ToLinkedServiceSftpPtrOutputWithContext(ctx context.Context) LinkedServiceSftpPtrOutput
}

type linkedServiceSftpPtrType LinkedServiceSftpArgs

func (*linkedServiceSftpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSftp)(nil))
}

func (i *linkedServiceSftpPtrType) ToLinkedServiceSftpPtrOutput() LinkedServiceSftpPtrOutput {
	return i.ToLinkedServiceSftpPtrOutputWithContext(context.Background())
}

func (i *linkedServiceSftpPtrType) ToLinkedServiceSftpPtrOutputWithContext(ctx context.Context) LinkedServiceSftpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpPtrOutput)
}

// LinkedServiceSftpArrayInput is an input type that accepts LinkedServiceSftpArray and LinkedServiceSftpArrayOutput values.
// You can construct a concrete instance of `LinkedServiceSftpArrayInput` via:
//
//          LinkedServiceSftpArray{ LinkedServiceSftpArgs{...} }
type LinkedServiceSftpArrayInput interface {
	pulumi.Input

	ToLinkedServiceSftpArrayOutput() LinkedServiceSftpArrayOutput
	ToLinkedServiceSftpArrayOutputWithContext(context.Context) LinkedServiceSftpArrayOutput
}

type LinkedServiceSftpArray []LinkedServiceSftpInput

func (LinkedServiceSftpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LinkedServiceSftp)(nil)).Elem()
}

func (i LinkedServiceSftpArray) ToLinkedServiceSftpArrayOutput() LinkedServiceSftpArrayOutput {
	return i.ToLinkedServiceSftpArrayOutputWithContext(context.Background())
}

func (i LinkedServiceSftpArray) ToLinkedServiceSftpArrayOutputWithContext(ctx context.Context) LinkedServiceSftpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpArrayOutput)
}

// LinkedServiceSftpMapInput is an input type that accepts LinkedServiceSftpMap and LinkedServiceSftpMapOutput values.
// You can construct a concrete instance of `LinkedServiceSftpMapInput` via:
//
//          LinkedServiceSftpMap{ "key": LinkedServiceSftpArgs{...} }
type LinkedServiceSftpMapInput interface {
	pulumi.Input

	ToLinkedServiceSftpMapOutput() LinkedServiceSftpMapOutput
	ToLinkedServiceSftpMapOutputWithContext(context.Context) LinkedServiceSftpMapOutput
}

type LinkedServiceSftpMap map[string]LinkedServiceSftpInput

func (LinkedServiceSftpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LinkedServiceSftp)(nil)).Elem()
}

func (i LinkedServiceSftpMap) ToLinkedServiceSftpMapOutput() LinkedServiceSftpMapOutput {
	return i.ToLinkedServiceSftpMapOutputWithContext(context.Background())
}

func (i LinkedServiceSftpMap) ToLinkedServiceSftpMapOutputWithContext(ctx context.Context) LinkedServiceSftpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceSftpMapOutput)
}

type LinkedServiceSftpOutput struct{ *pulumi.OutputState }

func (LinkedServiceSftpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceSftp)(nil))
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpOutput() LinkedServiceSftpOutput {
	return o
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpOutputWithContext(ctx context.Context) LinkedServiceSftpOutput {
	return o
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpPtrOutput() LinkedServiceSftpPtrOutput {
	return o.ToLinkedServiceSftpPtrOutputWithContext(context.Background())
}

func (o LinkedServiceSftpOutput) ToLinkedServiceSftpPtrOutputWithContext(ctx context.Context) LinkedServiceSftpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceSftp) *LinkedServiceSftp {
		return &v
	}).(LinkedServiceSftpPtrOutput)
}

type LinkedServiceSftpPtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceSftpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceSftp)(nil))
}

func (o LinkedServiceSftpPtrOutput) ToLinkedServiceSftpPtrOutput() LinkedServiceSftpPtrOutput {
	return o
}

func (o LinkedServiceSftpPtrOutput) ToLinkedServiceSftpPtrOutputWithContext(ctx context.Context) LinkedServiceSftpPtrOutput {
	return o
}

func (o LinkedServiceSftpPtrOutput) Elem() LinkedServiceSftpOutput {
	return o.ApplyT(func(v *LinkedServiceSftp) LinkedServiceSftp {
		if v != nil {
			return *v
		}
		var ret LinkedServiceSftp
		return ret
	}).(LinkedServiceSftpOutput)
}

type LinkedServiceSftpArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceSftpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceSftp)(nil))
}

func (o LinkedServiceSftpArrayOutput) ToLinkedServiceSftpArrayOutput() LinkedServiceSftpArrayOutput {
	return o
}

func (o LinkedServiceSftpArrayOutput) ToLinkedServiceSftpArrayOutputWithContext(ctx context.Context) LinkedServiceSftpArrayOutput {
	return o
}

func (o LinkedServiceSftpArrayOutput) Index(i pulumi.IntInput) LinkedServiceSftpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceSftp {
		return vs[0].([]LinkedServiceSftp)[vs[1].(int)]
	}).(LinkedServiceSftpOutput)
}

type LinkedServiceSftpMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceSftpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceSftp)(nil))
}

func (o LinkedServiceSftpMapOutput) ToLinkedServiceSftpMapOutput() LinkedServiceSftpMapOutput {
	return o
}

func (o LinkedServiceSftpMapOutput) ToLinkedServiceSftpMapOutputWithContext(ctx context.Context) LinkedServiceSftpMapOutput {
	return o
}

func (o LinkedServiceSftpMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceSftpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceSftp {
		return vs[0].(map[string]LinkedServiceSftp)[vs[1].(string)]
	}).(LinkedServiceSftpOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceSftpOutput{})
	pulumi.RegisterOutputType(LinkedServiceSftpPtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceSftpArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceSftpMapOutput{})
}
