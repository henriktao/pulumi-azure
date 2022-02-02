// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Logic App Integration Account Agreement.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/logicapps"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testResourceGroup, err := core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testIntegrationAccount, err := logicapps.NewIntegrationAccount(ctx, "testIntegrationAccount", &logicapps.IntegrationAccountArgs{
// 			Location:          testResourceGroup.Location,
// 			ResourceGroupName: testResourceGroup.Name,
// 			SkuName:           pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		host, err := logicapps.NewIntegrationAccountPartner(ctx, "host", &logicapps.IntegrationAccountPartnerArgs{
// 			ResourceGroupName:      testResourceGroup.Name,
// 			IntegrationAccountName: testIntegrationAccount.Name,
// 			BusinessIdentities: logicapps.IntegrationAccountPartnerBusinessIdentityArray{
// 				&logicapps.IntegrationAccountPartnerBusinessIdentityArgs{
// 					Qualifier: pulumi.String("AS2Identity"),
// 					Value:     pulumi.String("FabrikamNY"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		guest, err := logicapps.NewIntegrationAccountPartner(ctx, "guest", &logicapps.IntegrationAccountPartnerArgs{
// 			ResourceGroupName:      testResourceGroup.Name,
// 			IntegrationAccountName: testIntegrationAccount.Name,
// 			BusinessIdentities: logicapps.IntegrationAccountPartnerBusinessIdentityArray{
// 				&logicapps.IntegrationAccountPartnerBusinessIdentityArgs{
// 					Qualifier: pulumi.String("AS2Identity"),
// 					Value:     pulumi.String("FabrikamDC"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logicapps.NewIntegrationAccountAgreement(ctx, "testIntegrationAccountAgreement", &logicapps.IntegrationAccountAgreementArgs{
// 			ResourceGroupName:      testResourceGroup.Name,
// 			IntegrationAccountName: testIntegrationAccount.Name,
// 			AgreementType:          pulumi.String("AS2"),
// 			HostPartnerName:        host.Name,
// 			GuestPartnerName:       guest.Name,
// 			Content:                readFileOrPanic("testdata/integration_account_agreement_content_as2.json"),
// 			HostIdentity: &logicapps.IntegrationAccountAgreementHostIdentityArgs{
// 				Qualifier: pulumi.String("AS2Identity"),
// 				Value:     pulumi.String("FabrikamNY"),
// 			},
// 			GuestIdentity: &logicapps.IntegrationAccountAgreementGuestIdentityArgs{
// 				Qualifier: pulumi.String("AS2Identity"),
// 				Value:     pulumi.String("FabrikamDC"),
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
// Logic App Integration Account Agreements can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:logicapps/integrationAccountAgreement:IntegrationAccountAgreement example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/agreements/agreement1
// ```
type IntegrationAccountAgreement struct {
	pulumi.CustomResourceState

	// The type of the Logic App Integration Account Agreement. Possible values are `AS2`, `X12` and `Edifact`.
	AgreementType pulumi.StringOutput `pulumi:"agreementType"`
	// The content of the Logic App Integration Account Agreement.
	Content pulumi.StringOutput `pulumi:"content"`
	// A `guestIdentity` block as documented below.
	GuestIdentity IntegrationAccountAgreementGuestIdentityOutput `pulumi:"guestIdentity"`
	// The name of the guest Logic App Integration Account Partner.
	GuestPartnerName pulumi.StringOutput `pulumi:"guestPartnerName"`
	// A `hostIdentity` block as documented below.
	HostIdentity IntegrationAccountAgreementHostIdentityOutput `pulumi:"hostIdentity"`
	// The name of the host Logic App Integration Account Partner.
	HostPartnerName pulumi.StringOutput `pulumi:"hostPartnerName"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringOutput `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Agreement.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Agreement. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Agreement should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIntegrationAccountAgreement registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, args *IntegrationAccountAgreementArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgreementType == nil {
		return nil, errors.New("invalid value for required argument 'AgreementType'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.GuestIdentity == nil {
		return nil, errors.New("invalid value for required argument 'GuestIdentity'")
	}
	if args.GuestPartnerName == nil {
		return nil, errors.New("invalid value for required argument 'GuestPartnerName'")
	}
	if args.HostIdentity == nil {
		return nil, errors.New("invalid value for required argument 'HostIdentity'")
	}
	if args.HostPartnerName == nil {
		return nil, errors.New("invalid value for required argument 'HostPartnerName'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IntegrationAccountAgreement
	err := ctx.RegisterResource("azure:logicapps/integrationAccountAgreement:IntegrationAccountAgreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountAgreement gets an existing IntegrationAccountAgreement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountAgreementState, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	var resource IntegrationAccountAgreement
	err := ctx.ReadResource("azure:logicapps/integrationAccountAgreement:IntegrationAccountAgreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountAgreement resources.
type integrationAccountAgreementState struct {
	// The type of the Logic App Integration Account Agreement. Possible values are `AS2`, `X12` and `Edifact`.
	AgreementType *string `pulumi:"agreementType"`
	// The content of the Logic App Integration Account Agreement.
	Content *string `pulumi:"content"`
	// A `guestIdentity` block as documented below.
	GuestIdentity *IntegrationAccountAgreementGuestIdentity `pulumi:"guestIdentity"`
	// The name of the guest Logic App Integration Account Partner.
	GuestPartnerName *string `pulumi:"guestPartnerName"`
	// A `hostIdentity` block as documented below.
	HostIdentity *IntegrationAccountAgreementHostIdentity `pulumi:"hostIdentity"`
	// The name of the host Logic App Integration Account Partner.
	HostPartnerName *string `pulumi:"hostPartnerName"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName *string `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Agreement.
	Metadata map[string]string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Agreement. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Agreement should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IntegrationAccountAgreementState struct {
	// The type of the Logic App Integration Account Agreement. Possible values are `AS2`, `X12` and `Edifact`.
	AgreementType pulumi.StringPtrInput
	// The content of the Logic App Integration Account Agreement.
	Content pulumi.StringPtrInput
	// A `guestIdentity` block as documented below.
	GuestIdentity IntegrationAccountAgreementGuestIdentityPtrInput
	// The name of the guest Logic App Integration Account Partner.
	GuestPartnerName pulumi.StringPtrInput
	// A `hostIdentity` block as documented below.
	HostIdentity IntegrationAccountAgreementHostIdentityPtrInput
	// The name of the host Logic App Integration Account Partner.
	HostPartnerName pulumi.StringPtrInput
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringPtrInput
	// The metadata of the Logic App Integration Account Agreement.
	Metadata pulumi.StringMapInput
	// The name which should be used for this Logic App Integration Account Agreement. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Agreement should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IntegrationAccountAgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementState)(nil)).Elem()
}

type integrationAccountAgreementArgs struct {
	// The type of the Logic App Integration Account Agreement. Possible values are `AS2`, `X12` and `Edifact`.
	AgreementType string `pulumi:"agreementType"`
	// The content of the Logic App Integration Account Agreement.
	Content string `pulumi:"content"`
	// A `guestIdentity` block as documented below.
	GuestIdentity IntegrationAccountAgreementGuestIdentity `pulumi:"guestIdentity"`
	// The name of the guest Logic App Integration Account Partner.
	GuestPartnerName string `pulumi:"guestPartnerName"`
	// A `hostIdentity` block as documented below.
	HostIdentity IntegrationAccountAgreementHostIdentity `pulumi:"hostIdentity"`
	// The name of the host Logic App Integration Account Partner.
	HostPartnerName string `pulumi:"hostPartnerName"`
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The metadata of the Logic App Integration Account Agreement.
	Metadata map[string]string `pulumi:"metadata"`
	// The name which should be used for this Logic App Integration Account Agreement. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Logic App Integration Account Agreement should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IntegrationAccountAgreement resource.
type IntegrationAccountAgreementArgs struct {
	// The type of the Logic App Integration Account Agreement. Possible values are `AS2`, `X12` and `Edifact`.
	AgreementType pulumi.StringInput
	// The content of the Logic App Integration Account Agreement.
	Content pulumi.StringInput
	// A `guestIdentity` block as documented below.
	GuestIdentity IntegrationAccountAgreementGuestIdentityInput
	// The name of the guest Logic App Integration Account Partner.
	GuestPartnerName pulumi.StringInput
	// A `hostIdentity` block as documented below.
	HostIdentity IntegrationAccountAgreementHostIdentityInput
	// The name of the host Logic App Integration Account Partner.
	HostPartnerName pulumi.StringInput
	// The name of the Logic App Integration Account. Changing this forces a new resource to be created.
	IntegrationAccountName pulumi.StringInput
	// The metadata of the Logic App Integration Account Agreement.
	Metadata pulumi.StringMapInput
	// The name which should be used for this Logic App Integration Account Agreement. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Logic App Integration Account Agreement should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IntegrationAccountAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementArgs)(nil)).Elem()
}

type IntegrationAccountAgreementInput interface {
	pulumi.Input

	ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput
	ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput
}

func (*IntegrationAccountAgreement) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAgreement)(nil)).Elem()
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return i.ToIntegrationAccountAgreementOutputWithContext(context.Background())
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAgreementOutput)
}

// IntegrationAccountAgreementArrayInput is an input type that accepts IntegrationAccountAgreementArray and IntegrationAccountAgreementArrayOutput values.
// You can construct a concrete instance of `IntegrationAccountAgreementArrayInput` via:
//
//          IntegrationAccountAgreementArray{ IntegrationAccountAgreementArgs{...} }
type IntegrationAccountAgreementArrayInput interface {
	pulumi.Input

	ToIntegrationAccountAgreementArrayOutput() IntegrationAccountAgreementArrayOutput
	ToIntegrationAccountAgreementArrayOutputWithContext(context.Context) IntegrationAccountAgreementArrayOutput
}

type IntegrationAccountAgreementArray []IntegrationAccountAgreementInput

func (IntegrationAccountAgreementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAccountAgreement)(nil)).Elem()
}

func (i IntegrationAccountAgreementArray) ToIntegrationAccountAgreementArrayOutput() IntegrationAccountAgreementArrayOutput {
	return i.ToIntegrationAccountAgreementArrayOutputWithContext(context.Background())
}

func (i IntegrationAccountAgreementArray) ToIntegrationAccountAgreementArrayOutputWithContext(ctx context.Context) IntegrationAccountAgreementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAgreementArrayOutput)
}

// IntegrationAccountAgreementMapInput is an input type that accepts IntegrationAccountAgreementMap and IntegrationAccountAgreementMapOutput values.
// You can construct a concrete instance of `IntegrationAccountAgreementMapInput` via:
//
//          IntegrationAccountAgreementMap{ "key": IntegrationAccountAgreementArgs{...} }
type IntegrationAccountAgreementMapInput interface {
	pulumi.Input

	ToIntegrationAccountAgreementMapOutput() IntegrationAccountAgreementMapOutput
	ToIntegrationAccountAgreementMapOutputWithContext(context.Context) IntegrationAccountAgreementMapOutput
}

type IntegrationAccountAgreementMap map[string]IntegrationAccountAgreementInput

func (IntegrationAccountAgreementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAccountAgreement)(nil)).Elem()
}

func (i IntegrationAccountAgreementMap) ToIntegrationAccountAgreementMapOutput() IntegrationAccountAgreementMapOutput {
	return i.ToIntegrationAccountAgreementMapOutputWithContext(context.Background())
}

func (i IntegrationAccountAgreementMap) ToIntegrationAccountAgreementMapOutputWithContext(ctx context.Context) IntegrationAccountAgreementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAgreementMapOutput)
}

type IntegrationAccountAgreementOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAgreement)(nil)).Elem()
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return o
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return o
}

type IntegrationAccountAgreementArrayOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAgreementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationAccountAgreement)(nil)).Elem()
}

func (o IntegrationAccountAgreementArrayOutput) ToIntegrationAccountAgreementArrayOutput() IntegrationAccountAgreementArrayOutput {
	return o
}

func (o IntegrationAccountAgreementArrayOutput) ToIntegrationAccountAgreementArrayOutputWithContext(ctx context.Context) IntegrationAccountAgreementArrayOutput {
	return o
}

func (o IntegrationAccountAgreementArrayOutput) Index(i pulumi.IntInput) IntegrationAccountAgreementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationAccountAgreement {
		return vs[0].([]*IntegrationAccountAgreement)[vs[1].(int)]
	}).(IntegrationAccountAgreementOutput)
}

type IntegrationAccountAgreementMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAgreementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationAccountAgreement)(nil)).Elem()
}

func (o IntegrationAccountAgreementMapOutput) ToIntegrationAccountAgreementMapOutput() IntegrationAccountAgreementMapOutput {
	return o
}

func (o IntegrationAccountAgreementMapOutput) ToIntegrationAccountAgreementMapOutputWithContext(ctx context.Context) IntegrationAccountAgreementMapOutput {
	return o
}

func (o IntegrationAccountAgreementMapOutput) MapIndex(k pulumi.StringInput) IntegrationAccountAgreementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationAccountAgreement {
		return vs[0].(map[string]*IntegrationAccountAgreement)[vs[1].(string)]
	}).(IntegrationAccountAgreementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountAgreementInput)(nil)).Elem(), &IntegrationAccountAgreement{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountAgreementArrayInput)(nil)).Elem(), IntegrationAccountAgreementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountAgreementMapInput)(nil)).Elem(), IntegrationAccountAgreementMap{})
	pulumi.RegisterOutputType(IntegrationAccountAgreementOutput{})
	pulumi.RegisterOutputType(IntegrationAccountAgreementArrayOutput{})
	pulumi.RegisterOutputType(IntegrationAccountAgreementMapOutput{})
}
