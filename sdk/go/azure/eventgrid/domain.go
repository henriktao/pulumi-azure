// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EventGrid Domain
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventgrid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West US 2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eventgrid.NewDomain(ctx, "exampleDomain", &eventgrid.DomainArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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
// EventGrid Domains can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:eventgrid/domain:Domain domain1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The Endpoint associated with the EventGrid Domain.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules DomainInboundIpRuleArrayOutput `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues DomainInputMappingDefaultValuesPtrOutput `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields DomainInputMappingFieldsPtrOutput `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Domain.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrOutput `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Domain.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure:eventhub/domain:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure:eventgrid/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure:eventgrid/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The Endpoint associated with the EventGrid Domain.
	Endpoint *string `pulumi:"endpoint"`
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules []DomainInboundIpRule `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *DomainInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *DomainInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Primary Shared Access Key associated with the EventGrid Domain.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Shared Access Key associated with the EventGrid Domain.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type DomainState struct {
	// The Endpoint associated with the EventGrid Domain.
	Endpoint pulumi.StringPtrInput
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules DomainInboundIpRuleArrayInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues DomainInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields DomainInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Primary Shared Access Key associated with the EventGrid Domain.
	PrimaryAccessKey pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Shared Access Key associated with the EventGrid Domain.
	SecondaryAccessKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules []DomainInboundIpRule `pulumi:"inboundIpRules"`
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues *DomainInputMappingDefaultValues `pulumi:"inputMappingDefaultValues"`
	// A `inputMappingFields` block as defined below.
	InputMappingFields *DomainInputMappingFields `pulumi:"inputMappingFields"`
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema *string `pulumi:"inputSchema"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled *bool `pulumi:"publicNetworkAccessEnabled"`
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// One or more `inboundIpRule` blocks as defined below.
	InboundIpRules DomainInboundIpRuleArrayInput
	// A `inputMappingDefaultValues` block as defined below.
	InputMappingDefaultValues DomainInputMappingDefaultValuesPtrInput
	// A `inputMappingFields` block as defined below.
	InputMappingFields DomainInputMappingFieldsPtrInput
	// Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
	InputSchema pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Whether or not public network access is allowed for this server. Defaults to `true`.
	PublicNetworkAccessEnabled pulumi.BoolPtrInput
	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (i Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOutput)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
