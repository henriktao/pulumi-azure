// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Content Key Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/media"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("GRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServiceAccount, err := media.NewServiceAccount(ctx, "exampleServiceAccount", &media.ServiceAccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			StorageAccounts: media.ServiceAccountStorageAccountArray{
// 				&media.ServiceAccountStorageAccountArgs{
// 					Id:        exampleAccount.ID(),
// 					IsPrimary: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"allowed_track_types": "SD_HD",
// 			"content_key_specs": []map[string]interface{}{
// 				map[string]interface{}{
// 					"track_type":     "SD",
// 					"security_level": 1,
// 					"required_output_protection": map[string]interface{}{
// 						"hdcp": "HDCP_V2",
// 					},
// 				},
// 			},
// 			"policy_overrides": map[string]interface{}{
// 				"can_play":    true,
// 				"can_persist": true,
// 				"can_renew":   false,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err = media.NewContentKeyPolicy(ctx, "exampleContentKeyPolicy", &media.ContentKeyPolicyArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			PolicyOptions: media.ContentKeyPolicyPolicyOptionArray{
// 				&media.ContentKeyPolicyPolicyOptionArgs{
// 					Name: pulumi.String("fairPlay"),
// 					FairplayConfiguration: &media.ContentKeyPolicyPolicyOptionFairplayConfigurationArgs{
// 						Ask:                   pulumi.String("bb566284cc124a21c435a92cd3c108c4"),
// 						Pfx:                   pulumi.String("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"),
// 						PfxPassword:           pulumi.String("password"),
// 						RentalDurationSeconds: pulumi.Int(2249),
// 						RentalAndLeaseKeyType: pulumi.String("PersistentUnlimited"),
// 					},
// 					OpenRestrictionEnabled: pulumi.Bool(true),
// 				},
// 				&media.ContentKeyPolicyPolicyOptionArgs{
// 					Name: pulumi.String("playReady"),
// 					PlayreadyConfigurationLicenses: media.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArray{
// 						&media.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArgs{
// 							AllowTestDevices: pulumi.Bool(true),
// 							BeginDate:        pulumi.String("2017-10-16T18:22:53Z"),
// 							PlayRight: &media.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightArgs{
// 								ScmsRestriction:                                    pulumi.Int(2),
// 								DigitalVideoOnlyContentRestriction:                 pulumi.Bool(false),
// 								ImageConstraintForAnalogComponentVideoRestriction:  pulumi.Bool(false),
// 								ImageConstraintForAnalogComputerMonitorRestriction: pulumi.Bool(false),
// 								AllowPassingVideoContentToUnknownOutput:            pulumi.String("NotAllowed"),
// 								UncompressedDigitalVideoOpl:                        pulumi.Int(100),
// 								UncompressedDigitalAudioOpl:                        pulumi.Int(100),
// 								AnalogVideoOpl:                                     pulumi.Int(150),
// 								CompressedDigitalAudioOpl:                          pulumi.Int(150),
// 							},
// 							LicenseType:                         pulumi.String("Persistent"),
// 							ContentType:                         pulumi.String("UltraVioletDownload"),
// 							ContentKeyLocationFromHeaderEnabled: pulumi.Bool(true),
// 						},
// 					},
// 					OpenRestrictionEnabled: pulumi.Bool(true),
// 				},
// 				&media.ContentKeyPolicyPolicyOptionArgs{
// 					Name:                         pulumi.String("clearKey"),
// 					ClearKeyConfigurationEnabled: pulumi.Bool(true),
// 					TokenRestriction: &media.ContentKeyPolicyPolicyOptionTokenRestrictionArgs{
// 						Issuer:                   pulumi.String("urn:issuer"),
// 						Audience:                 pulumi.String("urn:audience"),
// 						TokenType:                pulumi.String("Swt"),
// 						PrimarySymmetricTokenKey: pulumi.String("AAAAAAAAAAAAAAAAAAAAAA=="),
// 					},
// 				},
// 				&media.ContentKeyPolicyPolicyOptionArgs{
// 					Name:                          pulumi.String("widevine"),
// 					WidevineConfigurationTemplate: pulumi.String(json0),
// 					OpenRestrictionEnabled:        pulumi.Bool(true),
// 				},
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
// Resource Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/contentKeyPolicy:ContentKeyPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/contentkeypolicies/policy1
// ```
type ContentKeyPolicy struct {
	pulumi.CustomResourceState

	// A description for the Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName pulumi.StringOutput `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more `policyOption` blocks as defined below.
	PolicyOptions ContentKeyPolicyPolicyOptionArrayOutput `pulumi:"policyOptions"`
	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewContentKeyPolicy registers a new resource with the given unique name, arguments, and options.
func NewContentKeyPolicy(ctx *pulumi.Context,
	name string, args *ContentKeyPolicyArgs, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MediaServicesAccountName == nil {
		return nil, errors.New("invalid value for required argument 'MediaServicesAccountName'")
	}
	if args.PolicyOptions == nil {
		return nil, errors.New("invalid value for required argument 'PolicyOptions'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ContentKeyPolicy
	err := ctx.RegisterResource("azure:media/contentKeyPolicy:ContentKeyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentKeyPolicy gets an existing ContentKeyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentKeyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentKeyPolicyState, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	var resource ContentKeyPolicy
	err := ctx.ReadResource("azure:media/contentKeyPolicy:ContentKeyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentKeyPolicy resources.
type contentKeyPolicyState struct {
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName *string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
	Name *string `pulumi:"name"`
	// One or more `policyOption` blocks as defined below.
	PolicyOptions []ContentKeyPolicyPolicyOption `pulumi:"policyOptions"`
	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ContentKeyPolicyState struct {
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName pulumi.StringPtrInput
	// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
	Name pulumi.StringPtrInput
	// One or more `policyOption` blocks as defined below.
	PolicyOptions ContentKeyPolicyPolicyOptionArrayInput
	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ContentKeyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyState)(nil)).Elem()
}

type contentKeyPolicyArgs struct {
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
	Name *string `pulumi:"name"`
	// One or more `policyOption` blocks as defined below.
	PolicyOptions []ContentKeyPolicyPolicyOption `pulumi:"policyOptions"`
	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ContentKeyPolicy resource.
type ContentKeyPolicyArgs struct {
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName pulumi.StringInput
	// The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
	Name pulumi.StringPtrInput
	// One or more `policyOption` blocks as defined below.
	PolicyOptions ContentKeyPolicyPolicyOptionArrayInput
	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName pulumi.StringInput
}

func (ContentKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyArgs)(nil)).Elem()
}

type ContentKeyPolicyInput interface {
	pulumi.Input

	ToContentKeyPolicyOutput() ContentKeyPolicyOutput
	ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput
}

func (*ContentKeyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicy)(nil)).Elem()
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return i.ToContentKeyPolicyOutputWithContext(context.Background())
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOutput)
}

// ContentKeyPolicyArrayInput is an input type that accepts ContentKeyPolicyArray and ContentKeyPolicyArrayOutput values.
// You can construct a concrete instance of `ContentKeyPolicyArrayInput` via:
//
//          ContentKeyPolicyArray{ ContentKeyPolicyArgs{...} }
type ContentKeyPolicyArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyArrayOutput() ContentKeyPolicyArrayOutput
	ToContentKeyPolicyArrayOutputWithContext(context.Context) ContentKeyPolicyArrayOutput
}

type ContentKeyPolicyArray []ContentKeyPolicyInput

func (ContentKeyPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContentKeyPolicy)(nil)).Elem()
}

func (i ContentKeyPolicyArray) ToContentKeyPolicyArrayOutput() ContentKeyPolicyArrayOutput {
	return i.ToContentKeyPolicyArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyArray) ToContentKeyPolicyArrayOutputWithContext(ctx context.Context) ContentKeyPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyArrayOutput)
}

// ContentKeyPolicyMapInput is an input type that accepts ContentKeyPolicyMap and ContentKeyPolicyMapOutput values.
// You can construct a concrete instance of `ContentKeyPolicyMapInput` via:
//
//          ContentKeyPolicyMap{ "key": ContentKeyPolicyArgs{...} }
type ContentKeyPolicyMapInput interface {
	pulumi.Input

	ToContentKeyPolicyMapOutput() ContentKeyPolicyMapOutput
	ToContentKeyPolicyMapOutputWithContext(context.Context) ContentKeyPolicyMapOutput
}

type ContentKeyPolicyMap map[string]ContentKeyPolicyInput

func (ContentKeyPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContentKeyPolicy)(nil)).Elem()
}

func (i ContentKeyPolicyMap) ToContentKeyPolicyMapOutput() ContentKeyPolicyMapOutput {
	return i.ToContentKeyPolicyMapOutputWithContext(context.Background())
}

func (i ContentKeyPolicyMap) ToContentKeyPolicyMapOutputWithContext(ctx context.Context) ContentKeyPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyMapOutput)
}

type ContentKeyPolicyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicy)(nil)).Elem()
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return o
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return o
}

type ContentKeyPolicyArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContentKeyPolicy)(nil)).Elem()
}

func (o ContentKeyPolicyArrayOutput) ToContentKeyPolicyArrayOutput() ContentKeyPolicyArrayOutput {
	return o
}

func (o ContentKeyPolicyArrayOutput) ToContentKeyPolicyArrayOutputWithContext(ctx context.Context) ContentKeyPolicyArrayOutput {
	return o
}

func (o ContentKeyPolicyArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContentKeyPolicy {
		return vs[0].([]*ContentKeyPolicy)[vs[1].(int)]
	}).(ContentKeyPolicyOutput)
}

type ContentKeyPolicyMapOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContentKeyPolicy)(nil)).Elem()
}

func (o ContentKeyPolicyMapOutput) ToContentKeyPolicyMapOutput() ContentKeyPolicyMapOutput {
	return o
}

func (o ContentKeyPolicyMapOutput) ToContentKeyPolicyMapOutputWithContext(ctx context.Context) ContentKeyPolicyMapOutput {
	return o
}

func (o ContentKeyPolicyMapOutput) MapIndex(k pulumi.StringInput) ContentKeyPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContentKeyPolicy {
		return vs[0].(map[string]*ContentKeyPolicy)[vs[1].(string)]
	}).(ContentKeyPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyInput)(nil)).Elem(), &ContentKeyPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyArrayInput)(nil)).Elem(), ContentKeyPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyMapInput)(nil)).Elem(), ContentKeyPolicyMap{})
	pulumi.RegisterOutputType(ContentKeyPolicyOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyMapOutput{})
}
