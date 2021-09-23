// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Custom Domain for a CDN Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cdn"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/dns"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("west europe"),
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
// 		exampleProfile, err := cdn.NewProfile(ctx, "exampleProfile", &cdn.ProfileArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard_Verizon"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEndpoint, err := cdn.NewEndpoint(ctx, "exampleEndpoint", &cdn.EndpointArgs{
// 			ProfileName:       exampleProfile.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Origins: cdn.EndpointOriginArray{
// 				&cdn.EndpointOriginArgs{
// 					Name:     pulumi.String("example"),
// 					HostName: exampleAccount.PrimaryBlobHost,
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "domain-rg"
// 		exampleZone, err := dns.LookupZone(ctx, &dns.LookupZoneArgs{
// 			Name:              "example-domain.com",
// 			ResourceGroupName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleCNameRecord, err := dns.NewCNameRecord(ctx, "exampleCNameRecord", &dns.CNameRecordArgs{
// 			ZoneName:          pulumi.String(exampleZone.Name),
// 			ResourceGroupName: pulumi.String(exampleZone.ResourceGroupName),
// 			Ttl:               pulumi.Int(3600),
// 			TargetResourceId:  exampleEndpoint.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cdn.NewEndpointCustomDomain(ctx, "exampleEndpointCustomDomain", &cdn.EndpointCustomDomainArgs{
// 			CdnEndpointId: exampleEndpoint.ID(),
// 			HostName: exampleCNameRecord.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v%v", name, ".", exampleZone.Name), nil
// 			}).(pulumi.StringOutput),
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
// CDN Endpoint Custom Domains can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cdn/endpointCustomDomain:EndpointCustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
// ```
type EndpointCustomDomain struct {
	pulumi.CustomResourceState

	// The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
	CdnEndpointId pulumi.StringOutput `pulumi:"cdnEndpointId"`
	// The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEndpointCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewEndpointCustomDomain(ctx *pulumi.Context,
	name string, args *EndpointCustomDomainArgs, opts ...pulumi.ResourceOption) (*EndpointCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CdnEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'CdnEndpointId'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	var resource EndpointCustomDomain
	err := ctx.RegisterResource("azure:cdn/endpointCustomDomain:EndpointCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointCustomDomain gets an existing EndpointCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointCustomDomainState, opts ...pulumi.ResourceOption) (*EndpointCustomDomain, error) {
	var resource EndpointCustomDomain
	err := ctx.ReadResource("azure:cdn/endpointCustomDomain:EndpointCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointCustomDomain resources.
type endpointCustomDomainState struct {
	// The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
	CdnEndpointId *string `pulumi:"cdnEndpointId"`
	// The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	HostName *string `pulumi:"hostName"`
	// The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	Name *string `pulumi:"name"`
}

type EndpointCustomDomainState struct {
	// The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
	CdnEndpointId pulumi.StringPtrInput
	// The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	HostName pulumi.StringPtrInput
	// The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	Name pulumi.StringPtrInput
}

func (EndpointCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointCustomDomainState)(nil)).Elem()
}

type endpointCustomDomainArgs struct {
	// The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
	CdnEndpointId string `pulumi:"cdnEndpointId"`
	// The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	HostName string `pulumi:"hostName"`
	// The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a EndpointCustomDomain resource.
type EndpointCustomDomainArgs struct {
	// The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
	CdnEndpointId pulumi.StringInput
	// The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	HostName pulumi.StringInput
	// The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
	Name pulumi.StringPtrInput
}

func (EndpointCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointCustomDomainArgs)(nil)).Elem()
}

type EndpointCustomDomainInput interface {
	pulumi.Input

	ToEndpointCustomDomainOutput() EndpointCustomDomainOutput
	ToEndpointCustomDomainOutputWithContext(ctx context.Context) EndpointCustomDomainOutput
}

func (*EndpointCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointCustomDomain)(nil))
}

func (i *EndpointCustomDomain) ToEndpointCustomDomainOutput() EndpointCustomDomainOutput {
	return i.ToEndpointCustomDomainOutputWithContext(context.Background())
}

func (i *EndpointCustomDomain) ToEndpointCustomDomainOutputWithContext(ctx context.Context) EndpointCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointCustomDomainOutput)
}

func (i *EndpointCustomDomain) ToEndpointCustomDomainPtrOutput() EndpointCustomDomainPtrOutput {
	return i.ToEndpointCustomDomainPtrOutputWithContext(context.Background())
}

func (i *EndpointCustomDomain) ToEndpointCustomDomainPtrOutputWithContext(ctx context.Context) EndpointCustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointCustomDomainPtrOutput)
}

type EndpointCustomDomainPtrInput interface {
	pulumi.Input

	ToEndpointCustomDomainPtrOutput() EndpointCustomDomainPtrOutput
	ToEndpointCustomDomainPtrOutputWithContext(ctx context.Context) EndpointCustomDomainPtrOutput
}

type endpointCustomDomainPtrType EndpointCustomDomainArgs

func (*endpointCustomDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointCustomDomain)(nil))
}

func (i *endpointCustomDomainPtrType) ToEndpointCustomDomainPtrOutput() EndpointCustomDomainPtrOutput {
	return i.ToEndpointCustomDomainPtrOutputWithContext(context.Background())
}

func (i *endpointCustomDomainPtrType) ToEndpointCustomDomainPtrOutputWithContext(ctx context.Context) EndpointCustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointCustomDomainPtrOutput)
}

// EndpointCustomDomainArrayInput is an input type that accepts EndpointCustomDomainArray and EndpointCustomDomainArrayOutput values.
// You can construct a concrete instance of `EndpointCustomDomainArrayInput` via:
//
//          EndpointCustomDomainArray{ EndpointCustomDomainArgs{...} }
type EndpointCustomDomainArrayInput interface {
	pulumi.Input

	ToEndpointCustomDomainArrayOutput() EndpointCustomDomainArrayOutput
	ToEndpointCustomDomainArrayOutputWithContext(context.Context) EndpointCustomDomainArrayOutput
}

type EndpointCustomDomainArray []EndpointCustomDomainInput

func (EndpointCustomDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointCustomDomain)(nil)).Elem()
}

func (i EndpointCustomDomainArray) ToEndpointCustomDomainArrayOutput() EndpointCustomDomainArrayOutput {
	return i.ToEndpointCustomDomainArrayOutputWithContext(context.Background())
}

func (i EndpointCustomDomainArray) ToEndpointCustomDomainArrayOutputWithContext(ctx context.Context) EndpointCustomDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointCustomDomainArrayOutput)
}

// EndpointCustomDomainMapInput is an input type that accepts EndpointCustomDomainMap and EndpointCustomDomainMapOutput values.
// You can construct a concrete instance of `EndpointCustomDomainMapInput` via:
//
//          EndpointCustomDomainMap{ "key": EndpointCustomDomainArgs{...} }
type EndpointCustomDomainMapInput interface {
	pulumi.Input

	ToEndpointCustomDomainMapOutput() EndpointCustomDomainMapOutput
	ToEndpointCustomDomainMapOutputWithContext(context.Context) EndpointCustomDomainMapOutput
}

type EndpointCustomDomainMap map[string]EndpointCustomDomainInput

func (EndpointCustomDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointCustomDomain)(nil)).Elem()
}

func (i EndpointCustomDomainMap) ToEndpointCustomDomainMapOutput() EndpointCustomDomainMapOutput {
	return i.ToEndpointCustomDomainMapOutputWithContext(context.Background())
}

func (i EndpointCustomDomainMap) ToEndpointCustomDomainMapOutputWithContext(ctx context.Context) EndpointCustomDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointCustomDomainMapOutput)
}

type EndpointCustomDomainOutput struct{ *pulumi.OutputState }

func (EndpointCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointCustomDomain)(nil))
}

func (o EndpointCustomDomainOutput) ToEndpointCustomDomainOutput() EndpointCustomDomainOutput {
	return o
}

func (o EndpointCustomDomainOutput) ToEndpointCustomDomainOutputWithContext(ctx context.Context) EndpointCustomDomainOutput {
	return o
}

func (o EndpointCustomDomainOutput) ToEndpointCustomDomainPtrOutput() EndpointCustomDomainPtrOutput {
	return o.ToEndpointCustomDomainPtrOutputWithContext(context.Background())
}

func (o EndpointCustomDomainOutput) ToEndpointCustomDomainPtrOutputWithContext(ctx context.Context) EndpointCustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointCustomDomain) *EndpointCustomDomain {
		return &v
	}).(EndpointCustomDomainPtrOutput)
}

type EndpointCustomDomainPtrOutput struct{ *pulumi.OutputState }

func (EndpointCustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointCustomDomain)(nil))
}

func (o EndpointCustomDomainPtrOutput) ToEndpointCustomDomainPtrOutput() EndpointCustomDomainPtrOutput {
	return o
}

func (o EndpointCustomDomainPtrOutput) ToEndpointCustomDomainPtrOutputWithContext(ctx context.Context) EndpointCustomDomainPtrOutput {
	return o
}

func (o EndpointCustomDomainPtrOutput) Elem() EndpointCustomDomainOutput {
	return o.ApplyT(func(v *EndpointCustomDomain) EndpointCustomDomain {
		if v != nil {
			return *v
		}
		var ret EndpointCustomDomain
		return ret
	}).(EndpointCustomDomainOutput)
}

type EndpointCustomDomainArrayOutput struct{ *pulumi.OutputState }

func (EndpointCustomDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointCustomDomain)(nil))
}

func (o EndpointCustomDomainArrayOutput) ToEndpointCustomDomainArrayOutput() EndpointCustomDomainArrayOutput {
	return o
}

func (o EndpointCustomDomainArrayOutput) ToEndpointCustomDomainArrayOutputWithContext(ctx context.Context) EndpointCustomDomainArrayOutput {
	return o
}

func (o EndpointCustomDomainArrayOutput) Index(i pulumi.IntInput) EndpointCustomDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointCustomDomain {
		return vs[0].([]EndpointCustomDomain)[vs[1].(int)]
	}).(EndpointCustomDomainOutput)
}

type EndpointCustomDomainMapOutput struct{ *pulumi.OutputState }

func (EndpointCustomDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointCustomDomain)(nil))
}

func (o EndpointCustomDomainMapOutput) ToEndpointCustomDomainMapOutput() EndpointCustomDomainMapOutput {
	return o
}

func (o EndpointCustomDomainMapOutput) ToEndpointCustomDomainMapOutputWithContext(ctx context.Context) EndpointCustomDomainMapOutput {
	return o
}

func (o EndpointCustomDomainMapOutput) MapIndex(k pulumi.StringInput) EndpointCustomDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointCustomDomain {
		return vs[0].(map[string]EndpointCustomDomain)[vs[1].(string)]
	}).(EndpointCustomDomainOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointCustomDomainOutput{})
	pulumi.RegisterOutputType(EndpointCustomDomainPtrOutput{})
	pulumi.RegisterOutputType(EndpointCustomDomainArrayOutput{})
	pulumi.RegisterOutputType(EndpointCustomDomainMapOutput{})
}
