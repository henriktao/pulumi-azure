// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Streaming Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
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
// 		_, err = media.NewStreamingEndpoint(ctx, "exampleStreamingEndpoint", &media.StreamingEndpointArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			Location:                 exampleResourceGroup.Location,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			ScaleUnits:               pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Access Control
//
// ```go
// package main
//
// import (
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
// 		_, err = media.NewStreamingEndpoint(ctx, "exampleStreamingEndpoint", &media.StreamingEndpointArgs{
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			Location:                 exampleResourceGroup.Location,
// 			MediaServicesAccountName: exampleServiceAccount.Name,
// 			ScaleUnits:               pulumi.Int(2),
// 			AccessControl: &media.StreamingEndpointAccessControlArgs{
// 				IpAllows: media.StreamingEndpointAccessControlIpAllowArray{
// 					&media.StreamingEndpointAccessControlIpAllowArgs{
// 						Name:    pulumi.String("AllowedIP"),
// 						Address: pulumi.String("192.168.1.1"),
// 					},
// 					&media.StreamingEndpointAccessControlIpAllowArgs{
// 						Name:    pulumi.String("AnotherIp"),
// 						Address: pulumi.String("192.168.1.2"),
// 					},
// 				},
// 				AkamaiSignatureHeaderAuthenticationKeys: media.StreamingEndpointAccessControlAkamaiSignatureHeaderAuthenticationKeyArray{
// 					&media.StreamingEndpointAccessControlAkamaiSignatureHeaderAuthenticationKeyArgs{
// 						Identifier: pulumi.String("id1"),
// 						Expiration: pulumi.String("2030-12-31T16:00:00Z"),
// 						Base64Key:  pulumi.String("dGVzdGlkMQ=="),
// 					},
// 					&media.StreamingEndpointAccessControlAkamaiSignatureHeaderAuthenticationKeyArgs{
// 						Identifier: pulumi.String("id2"),
// 						Expiration: pulumi.String("2032-01-28T16:00:00Z"),
// 						Base64Key:  pulumi.String("dGVzdGlkMQ=="),
// 					},
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
// Streaming Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:media/streamingEndpoint:StreamingEndpoint example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/service1/streamingendpoints/endpoint1
// ```
type StreamingEndpoint struct {
	pulumi.CustomResourceState

	// A `accessControl` block as defined below.
	AccessControl StreamingEndpointAccessControlPtrOutput `pulumi:"accessControl"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStartEnabled pulumi.BoolOutput `pulumi:"autoStartEnabled"`
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrOutput `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile pulumi.StringOutput `pulumi:"cdnProfile"`
	// The CDN provider name. Supported value are `StandardVerizon`,`PremiumVerizon` and `StandardAkamai`
	CdnProvider pulumi.StringOutput `pulumi:"cdnProvider"`
	// A `crossSiteAccessPolicy` block as defined below.
	CrossSiteAccessPolicy StreamingEndpointCrossSiteAccessPolicyPtrOutput `pulumi:"crossSiteAccessPolicy"`
	// The custom host names of the streaming endpoint.
	CustomHostNames pulumi.StringArrayOutput `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The host name of the Streaming Endpoint.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Max cache age in seconds.
	MaxCacheAgeSeconds pulumi.IntPtrOutput `pulumi:"maxCacheAgeSeconds"`
	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	MediaServicesAccountName pulumi.StringOutput `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The number of scale units.
	ScaleUnits pulumi.IntOutput `pulumi:"scaleUnits"`
	// A mapping of tags which should be assigned to the Streaming Endpoint.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewStreamingEndpoint registers a new resource with the given unique name, arguments, and options.
func NewStreamingEndpoint(ctx *pulumi.Context,
	name string, args *StreamingEndpointArgs, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MediaServicesAccountName == nil {
		return nil, errors.New("invalid value for required argument 'MediaServicesAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScaleUnits == nil {
		return nil, errors.New("invalid value for required argument 'ScaleUnits'")
	}
	var resource StreamingEndpoint
	err := ctx.RegisterResource("azure:media/streamingEndpoint:StreamingEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingEndpoint gets an existing StreamingEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingEndpointState, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	var resource StreamingEndpoint
	err := ctx.ReadResource("azure:media/streamingEndpoint:StreamingEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingEndpoint resources.
type streamingEndpointState struct {
	// A `accessControl` block as defined below.
	AccessControl *StreamingEndpointAccessControl `pulumi:"accessControl"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStartEnabled *bool `pulumi:"autoStartEnabled"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name. Supported value are `StandardVerizon`,`PremiumVerizon` and `StandardAkamai`
	CdnProvider *string `pulumi:"cdnProvider"`
	// A `crossSiteAccessPolicy` block as defined below.
	CrossSiteAccessPolicy *StreamingEndpointCrossSiteAccessPolicy `pulumi:"crossSiteAccessPolicy"`
	// The custom host names of the streaming endpoint.
	CustomHostNames []string `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description *string `pulumi:"description"`
	// The host name of the Streaming Endpoint.
	HostName *string `pulumi:"hostName"`
	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	Location *string `pulumi:"location"`
	// Max cache age in seconds.
	MaxCacheAgeSeconds *int `pulumi:"maxCacheAgeSeconds"`
	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	MediaServicesAccountName *string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The number of scale units.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// A mapping of tags which should be assigned to the Streaming Endpoint.
	Tags map[string]string `pulumi:"tags"`
}

type StreamingEndpointState struct {
	// A `accessControl` block as defined below.
	AccessControl StreamingEndpointAccessControlPtrInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStartEnabled pulumi.BoolPtrInput
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrInput
	// The CDN profile name.
	CdnProfile pulumi.StringPtrInput
	// The CDN provider name. Supported value are `StandardVerizon`,`PremiumVerizon` and `StandardAkamai`
	CdnProvider pulumi.StringPtrInput
	// A `crossSiteAccessPolicy` block as defined below.
	CrossSiteAccessPolicy StreamingEndpointCrossSiteAccessPolicyPtrInput
	// The custom host names of the streaming endpoint.
	CustomHostNames pulumi.StringArrayInput
	// The streaming endpoint description.
	Description pulumi.StringPtrInput
	// The host name of the Streaming Endpoint.
	HostName pulumi.StringPtrInput
	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	Location pulumi.StringPtrInput
	// Max cache age in seconds.
	MaxCacheAgeSeconds pulumi.IntPtrInput
	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	MediaServicesAccountName pulumi.StringPtrInput
	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The number of scale units.
	ScaleUnits pulumi.IntPtrInput
	// A mapping of tags which should be assigned to the Streaming Endpoint.
	Tags pulumi.StringMapInput
}

func (StreamingEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointState)(nil)).Elem()
}

type streamingEndpointArgs struct {
	// A `accessControl` block as defined below.
	AccessControl *StreamingEndpointAccessControl `pulumi:"accessControl"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStartEnabled *bool `pulumi:"autoStartEnabled"`
	// The CDN enabled flag.
	CdnEnabled *bool `pulumi:"cdnEnabled"`
	// The CDN profile name.
	CdnProfile *string `pulumi:"cdnProfile"`
	// The CDN provider name. Supported value are `StandardVerizon`,`PremiumVerizon` and `StandardAkamai`
	CdnProvider *string `pulumi:"cdnProvider"`
	// A `crossSiteAccessPolicy` block as defined below.
	CrossSiteAccessPolicy *StreamingEndpointCrossSiteAccessPolicy `pulumi:"crossSiteAccessPolicy"`
	// The custom host names of the streaming endpoint.
	CustomHostNames []string `pulumi:"customHostNames"`
	// The streaming endpoint description.
	Description *string `pulumi:"description"`
	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	Location *string `pulumi:"location"`
	// Max cache age in seconds.
	MaxCacheAgeSeconds *int `pulumi:"maxCacheAgeSeconds"`
	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	MediaServicesAccountName string `pulumi:"mediaServicesAccountName"`
	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of scale units.
	ScaleUnits int `pulumi:"scaleUnits"`
	// A mapping of tags which should be assigned to the Streaming Endpoint.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingEndpoint resource.
type StreamingEndpointArgs struct {
	// A `accessControl` block as defined below.
	AccessControl StreamingEndpointAccessControlPtrInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStartEnabled pulumi.BoolPtrInput
	// The CDN enabled flag.
	CdnEnabled pulumi.BoolPtrInput
	// The CDN profile name.
	CdnProfile pulumi.StringPtrInput
	// The CDN provider name. Supported value are `StandardVerizon`,`PremiumVerizon` and `StandardAkamai`
	CdnProvider pulumi.StringPtrInput
	// A `crossSiteAccessPolicy` block as defined below.
	CrossSiteAccessPolicy StreamingEndpointCrossSiteAccessPolicyPtrInput
	// The custom host names of the streaming endpoint.
	CustomHostNames pulumi.StringArrayInput
	// The streaming endpoint description.
	Description pulumi.StringPtrInput
	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	Location pulumi.StringPtrInput
	// Max cache age in seconds.
	MaxCacheAgeSeconds pulumi.IntPtrInput
	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	MediaServicesAccountName pulumi.StringInput
	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	ResourceGroupName pulumi.StringInput
	// The number of scale units.
	ScaleUnits pulumi.IntInput
	// A mapping of tags which should be assigned to the Streaming Endpoint.
	Tags pulumi.StringMapInput
}

func (StreamingEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointArgs)(nil)).Elem()
}

type StreamingEndpointInput interface {
	pulumi.Input

	ToStreamingEndpointOutput() StreamingEndpointOutput
	ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput
}

func (*StreamingEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (i *StreamingEndpoint) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return i.ToStreamingEndpointOutputWithContext(context.Background())
}

func (i *StreamingEndpoint) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointOutput)
}

// StreamingEndpointArrayInput is an input type that accepts StreamingEndpointArray and StreamingEndpointArrayOutput values.
// You can construct a concrete instance of `StreamingEndpointArrayInput` via:
//
//          StreamingEndpointArray{ StreamingEndpointArgs{...} }
type StreamingEndpointArrayInput interface {
	pulumi.Input

	ToStreamingEndpointArrayOutput() StreamingEndpointArrayOutput
	ToStreamingEndpointArrayOutputWithContext(context.Context) StreamingEndpointArrayOutput
}

type StreamingEndpointArray []StreamingEndpointInput

func (StreamingEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamingEndpoint)(nil)).Elem()
}

func (i StreamingEndpointArray) ToStreamingEndpointArrayOutput() StreamingEndpointArrayOutput {
	return i.ToStreamingEndpointArrayOutputWithContext(context.Background())
}

func (i StreamingEndpointArray) ToStreamingEndpointArrayOutputWithContext(ctx context.Context) StreamingEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointArrayOutput)
}

// StreamingEndpointMapInput is an input type that accepts StreamingEndpointMap and StreamingEndpointMapOutput values.
// You can construct a concrete instance of `StreamingEndpointMapInput` via:
//
//          StreamingEndpointMap{ "key": StreamingEndpointArgs{...} }
type StreamingEndpointMapInput interface {
	pulumi.Input

	ToStreamingEndpointMapOutput() StreamingEndpointMapOutput
	ToStreamingEndpointMapOutputWithContext(context.Context) StreamingEndpointMapOutput
}

type StreamingEndpointMap map[string]StreamingEndpointInput

func (StreamingEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamingEndpoint)(nil)).Elem()
}

func (i StreamingEndpointMap) ToStreamingEndpointMapOutput() StreamingEndpointMapOutput {
	return i.ToStreamingEndpointMapOutputWithContext(context.Background())
}

func (i StreamingEndpointMap) ToStreamingEndpointMapOutputWithContext(ctx context.Context) StreamingEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointMapOutput)
}

type StreamingEndpointOutput struct{ *pulumi.OutputState }

func (StreamingEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return o
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return o
}

type StreamingEndpointArrayOutput struct{ *pulumi.OutputState }

func (StreamingEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StreamingEndpoint)(nil)).Elem()
}

func (o StreamingEndpointArrayOutput) ToStreamingEndpointArrayOutput() StreamingEndpointArrayOutput {
	return o
}

func (o StreamingEndpointArrayOutput) ToStreamingEndpointArrayOutputWithContext(ctx context.Context) StreamingEndpointArrayOutput {
	return o
}

func (o StreamingEndpointArrayOutput) Index(i pulumi.IntInput) StreamingEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StreamingEndpoint {
		return vs[0].([]*StreamingEndpoint)[vs[1].(int)]
	}).(StreamingEndpointOutput)
}

type StreamingEndpointMapOutput struct{ *pulumi.OutputState }

func (StreamingEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StreamingEndpoint)(nil)).Elem()
}

func (o StreamingEndpointMapOutput) ToStreamingEndpointMapOutput() StreamingEndpointMapOutput {
	return o
}

func (o StreamingEndpointMapOutput) ToStreamingEndpointMapOutputWithContext(ctx context.Context) StreamingEndpointMapOutput {
	return o
}

func (o StreamingEndpointMapOutput) MapIndex(k pulumi.StringInput) StreamingEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StreamingEndpoint {
		return vs[0].(map[string]*StreamingEndpoint)[vs[1].(string)]
	}).(StreamingEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointInput)(nil)).Elem(), &StreamingEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointArrayInput)(nil)).Elem(), StreamingEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointMapInput)(nil)).Elem(), StreamingEndpointMap{})
	pulumi.RegisterOutputType(StreamingEndpointOutput{})
	pulumi.RegisterOutputType(StreamingEndpointArrayOutput{})
	pulumi.RegisterOutputType(StreamingEndpointMapOutput{})
}
