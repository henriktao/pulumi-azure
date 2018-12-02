// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Service.
type API struct {
	s *pulumi.ResourceState
}

// NewAPI registers a new resource with the given unique name, arguments, and options.
func NewAPI(ctx *pulumi.Context,
	name string, args *APIArgs, opts ...pulumi.ResourceOpt) (*API, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PublisherEmail == nil {
		return nil, errors.New("missing required argument 'PublisherEmail'")
	}
	if args == nil || args.PublisherName == nil {
		return nil, errors.New("missing required argument 'PublisherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalLocation"] = nil
		inputs["certificates"] = nil
		inputs["hostnameConfiguration"] = nil
		inputs["identity"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["notificationSenderEmail"] = nil
		inputs["publisherEmail"] = nil
		inputs["publisherName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["security"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["additionalLocation"] = args.AdditionalLocation
		inputs["certificates"] = args.Certificates
		inputs["hostnameConfiguration"] = args.HostnameConfiguration
		inputs["identity"] = args.Identity
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["notificationSenderEmail"] = args.NotificationSenderEmail
		inputs["publisherEmail"] = args.PublisherEmail
		inputs["publisherName"] = args.PublisherName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["security"] = args.Security
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["gatewayRegionalUrl"] = nil
	inputs["gatewayUrl"] = nil
	inputs["managementApiUrl"] = nil
	inputs["portalUrl"] = nil
	inputs["publicIpAddresses"] = nil
	inputs["scmUrl"] = nil
	s, err := ctx.RegisterResource("azure:apimanagement/aPI:API", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &API{s: s}, nil
}

// GetAPI gets an existing API resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPI(ctx *pulumi.Context,
	name string, id pulumi.ID, state *APIState, opts ...pulumi.ResourceOpt) (*API, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalLocation"] = state.AdditionalLocation
		inputs["certificates"] = state.Certificates
		inputs["gatewayRegionalUrl"] = state.GatewayRegionalUrl
		inputs["gatewayUrl"] = state.GatewayUrl
		inputs["hostnameConfiguration"] = state.HostnameConfiguration
		inputs["identity"] = state.Identity
		inputs["location"] = state.Location
		inputs["managementApiUrl"] = state.ManagementApiUrl
		inputs["name"] = state.Name
		inputs["notificationSenderEmail"] = state.NotificationSenderEmail
		inputs["portalUrl"] = state.PortalUrl
		inputs["publicIpAddresses"] = state.PublicIpAddresses
		inputs["publisherEmail"] = state.PublisherEmail
		inputs["publisherName"] = state.PublisherName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["scmUrl"] = state.ScmUrl
		inputs["security"] = state.Security
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:apimanagement/aPI:API", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &API{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *API) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *API) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// One or more `additional_location` blocks as defined below.
func (r *API) AdditionalLocation() *pulumi.Output {
	return r.s.State["additionalLocation"]
}

// One or more (up to 10) `certificate` blocks as defined below.
func (r *API) Certificates() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["certificates"])
}

// The URL of the Regional Gateway for the API Management Service in the specified region.
func (r *API) GatewayRegionalUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayRegionalUrl"])
}

// The URL of the Gateway for the API Management Service.
func (r *API) GatewayUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayUrl"])
}

// A `hostname_configuration` block as defined below.
func (r *API) HostnameConfiguration() *pulumi.Output {
	return r.s.State["hostnameConfiguration"]
}

// An `identity` block is documented below.
func (r *API) Identity() *pulumi.Output {
	return r.s.State["identity"]
}

// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
func (r *API) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The URL for the Management API associated with this API Management service.
func (r *API) ManagementApiUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managementApiUrl"])
}

// The name of the API Management Service. Changing this forces a new resource to be created.
func (r *API) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Email address from which the notification will be sent.
func (r *API) NotificationSenderEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["notificationSenderEmail"])
}

// The URL for the Publisher Portal associated with this API Management service.
func (r *API) PortalUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portalUrl"])
}

// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
func (r *API) PublicIpAddresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["publicIpAddresses"])
}

// The email of publisher/company.
func (r *API) PublisherEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publisherEmail"])
}

// The name of publisher/company.
func (r *API) PublisherName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publisherName"])
}

// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
func (r *API) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
func (r *API) ScmUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["scmUrl"])
}

// A `security` block as defined below.
func (r *API) Security() *pulumi.Output {
	return r.s.State["security"]
}

// A `sku` block as documented below.
func (r *API) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// A mapping of tags assigned to the resource.
func (r *API) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering API resources.
type APIState struct {
	// One or more `additional_location` blocks as defined below.
	AdditionalLocation interface{}
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates interface{}
	// The URL of the Regional Gateway for the API Management Service in the specified region.
	GatewayRegionalUrl interface{}
	// The URL of the Gateway for the API Management Service.
	GatewayUrl interface{}
	// A `hostname_configuration` block as defined below.
	HostnameConfiguration interface{}
	// An `identity` block is documented below.
	Identity interface{}
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location interface{}
	// The URL for the Management API associated with this API Management service.
	ManagementApiUrl interface{}
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// Email address from which the notification will be sent.
	NotificationSenderEmail interface{}
	// The URL for the Publisher Portal associated with this API Management service.
	PortalUrl interface{}
	// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
	PublicIpAddresses interface{}
	// The email of publisher/company.
	PublisherEmail interface{}
	// The name of publisher/company.
	PublisherName interface{}
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
	ScmUrl interface{}
	// A `security` block as defined below.
	Security interface{}
	// A `sku` block as documented below.
	Sku interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
}

// The set of arguments for constructing a API resource.
type APIArgs struct {
	// One or more `additional_location` blocks as defined below.
	AdditionalLocation interface{}
	// One or more (up to 10) `certificate` blocks as defined below.
	Certificates interface{}
	// A `hostname_configuration` block as defined below.
	HostnameConfiguration interface{}
	// An `identity` block is documented below.
	Identity interface{}
	// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the API Management Service. Changing this forces a new resource to be created.
	Name interface{}
	// Email address from which the notification will be sent.
	NotificationSenderEmail interface{}
	// The email of publisher/company.
	PublisherEmail interface{}
	// The name of publisher/company.
	PublisherName interface{}
	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `security` block as defined below.
	Security interface{}
	// A `sku` block as documented below.
	Sku interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
}
