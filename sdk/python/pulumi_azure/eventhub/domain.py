# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 identity: Optional[pulumi.Input['DomainIdentityArgs']] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']] = None,
                 input_mapping_fields: Optional[pulumi.Input['DomainInputMappingFieldsArgs']] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input['DomainIdentityArgs'] identity: An `identity` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input['DomainInputMappingDefaultValuesArgs'] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input['DomainInputMappingFieldsArgs'] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if inbound_ip_rules is not None:
            pulumi.set(__self__, "inbound_ip_rules", inbound_ip_rules)
        if input_mapping_default_values is not None:
            pulumi.set(__self__, "input_mapping_default_values", input_mapping_default_values)
        if input_mapping_fields is not None:
            pulumi.set(__self__, "input_mapping_fields", input_mapping_fields)
        if input_schema is not None:
            pulumi.set(__self__, "input_schema", input_schema)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_network_access_enabled is not None:
            pulumi.set(__self__, "public_network_access_enabled", public_network_access_enabled)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['DomainIdentityArgs']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['DomainIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]]:
        """
        One or more `inbound_ip_rule` blocks as defined below.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @inbound_ip_rules.setter
    def inbound_ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]]):
        pulumi.set(self, "inbound_ip_rules", value)

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @input_mapping_default_values.setter
    def input_mapping_default_values(self, value: Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']]):
        pulumi.set(self, "input_mapping_default_values", value)

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> Optional[pulumi.Input['DomainInputMappingFieldsArgs']]:
        """
        A `input_mapping_fields` block as defined below.
        """
        return pulumi.get(self, "input_mapping_fields")

    @input_mapping_fields.setter
    def input_mapping_fields(self, value: Optional[pulumi.Input['DomainInputMappingFieldsArgs']]):
        pulumi.set(self, "input_mapping_fields", value)

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "input_schema")

    @input_schema.setter
    def input_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_schema", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not public network access is allowed for this server. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @public_network_access_enabled.setter
    def public_network_access_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public_network_access_enabled", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DomainState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input['DomainIdentityArgs']] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']] = None,
                 input_mapping_fields: Optional[pulumi.Input['DomainInputMappingFieldsArgs']] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 primary_access_key: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 secondary_access_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Domain resources.
        :param pulumi.Input[str] endpoint: The Endpoint associated with the EventGrid Domain.
        :param pulumi.Input['DomainIdentityArgs'] identity: An `identity` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input['DomainInputMappingDefaultValuesArgs'] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input['DomainInputMappingFieldsArgs'] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The Primary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Secondary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if inbound_ip_rules is not None:
            pulumi.set(__self__, "inbound_ip_rules", inbound_ip_rules)
        if input_mapping_default_values is not None:
            pulumi.set(__self__, "input_mapping_default_values", input_mapping_default_values)
        if input_mapping_fields is not None:
            pulumi.set(__self__, "input_mapping_fields", input_mapping_fields)
        if input_schema is not None:
            pulumi.set(__self__, "input_schema", input_schema)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if primary_access_key is not None:
            pulumi.set(__self__, "primary_access_key", primary_access_key)
        if public_network_access_enabled is not None:
            pulumi.set(__self__, "public_network_access_enabled", public_network_access_enabled)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_access_key is not None:
            pulumi.set(__self__, "secondary_access_key", secondary_access_key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Endpoint associated with the EventGrid Domain.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['DomainIdentityArgs']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['DomainIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]]:
        """
        One or more `inbound_ip_rule` blocks as defined below.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @inbound_ip_rules.setter
    def inbound_ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DomainInboundIpRuleArgs']]]]):
        pulumi.set(self, "inbound_ip_rules", value)

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @input_mapping_default_values.setter
    def input_mapping_default_values(self, value: Optional[pulumi.Input['DomainInputMappingDefaultValuesArgs']]):
        pulumi.set(self, "input_mapping_default_values", value)

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> Optional[pulumi.Input['DomainInputMappingFieldsArgs']]:
        """
        A `input_mapping_fields` block as defined below.
        """
        return pulumi.get(self, "input_mapping_fields")

    @input_mapping_fields.setter
    def input_mapping_fields(self, value: Optional[pulumi.Input['DomainInputMappingFieldsArgs']]):
        pulumi.set(self, "input_mapping_fields", value)

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "input_schema")

    @input_schema.setter
    def input_schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_schema", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Primary Shared Access Key associated with the EventGrid Domain.
        """
        return pulumi.get(self, "primary_access_key")

    @primary_access_key.setter
    def primary_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_access_key", value)

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not public network access is allowed for this server. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @public_network_access_enabled.setter
    def public_network_access_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public_network_access_enabled", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Secondary Shared Access Key associated with the EventGrid Domain.
        """
        return pulumi.get(self, "secondary_access_key")

    @secondary_access_key.setter
    def secondary_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_access_key", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""azure.eventhub.Domain has been deprecated in favor of azure.eventgrid.Domain""", DeprecationWarning)


class Domain(pulumi.CustomResource):
    warnings.warn("""azure.eventhub.Domain has been deprecated in favor of azure.eventgrid.Domain""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['DomainIdentityArgs']]] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainInboundIpRuleArgs']]]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingDefaultValuesArgs']]] = None,
                 input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingFieldsArgs']]] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an EventGrid Domain

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_domain = azure.eventgrid.Domain("exampleDomain",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        EventGrid Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/domain:Domain domain1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DomainIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainInboundIpRuleArgs']]]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['DomainInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['DomainInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EventGrid Domain

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_domain = azure.eventgrid.Domain("exampleDomain",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        EventGrid Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:eventhub/domain:Domain domain1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
        ```

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['DomainIdentityArgs']]] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainInboundIpRuleArgs']]]]] = None,
                 input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingDefaultValuesArgs']]] = None,
                 input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingFieldsArgs']]] = None,
                 input_schema: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        pulumi.log.warn("""Domain is deprecated: azure.eventhub.Domain has been deprecated in favor of azure.eventgrid.Domain""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["identity"] = identity
            __props__.__dict__["inbound_ip_rules"] = inbound_ip_rules
            __props__.__dict__["input_mapping_default_values"] = input_mapping_default_values
            __props__.__dict__["input_mapping_fields"] = input_mapping_fields
            __props__.__dict__["input_schema"] = input_schema
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["public_network_access_enabled"] = public_network_access_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["primary_access_key"] = None
            __props__.__dict__["secondary_access_key"] = None
        super(Domain, __self__).__init__(
            'azure:eventhub/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['DomainIdentityArgs']]] = None,
            inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainInboundIpRuleArgs']]]]] = None,
            input_mapping_default_values: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingDefaultValuesArgs']]] = None,
            input_mapping_fields: Optional[pulumi.Input[pulumi.InputType['DomainInputMappingFieldsArgs']]] = None,
            input_schema: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_access_key: Optional[pulumi.Input[str]] = None,
            public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_access_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The Endpoint associated with the EventGrid Domain.
        :param pulumi.Input[pulumi.InputType['DomainIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DomainInboundIpRuleArgs']]]] inbound_ip_rules: One or more `inbound_ip_rule` blocks as defined below.
        :param pulumi.Input[pulumi.InputType['DomainInputMappingDefaultValuesArgs']] input_mapping_default_values: A `input_mapping_default_values` block as defined below.
        :param pulumi.Input[pulumi.InputType['DomainInputMappingFieldsArgs']] input_mapping_fields: A `input_mapping_fields` block as defined below.
        :param pulumi.Input[str] input_schema: Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_access_key: The Primary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this server. Defaults to `true`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_access_key: The Secondary Shared Access Key associated with the EventGrid Domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainState.__new__(_DomainState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["identity"] = identity
        __props__.__dict__["inbound_ip_rules"] = inbound_ip_rules
        __props__.__dict__["input_mapping_default_values"] = input_mapping_default_values
        __props__.__dict__["input_mapping_fields"] = input_mapping_fields
        __props__.__dict__["input_schema"] = input_schema
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["primary_access_key"] = primary_access_key
        __props__.__dict__["public_network_access_enabled"] = public_network_access_enabled
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["secondary_access_key"] = secondary_access_key
        __props__.__dict__["tags"] = tags
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The Endpoint associated with the EventGrid Domain.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.DomainIdentity']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.DomainInboundIpRule']]]:
        """
        One or more `inbound_ip_rule` blocks as defined below.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @property
    @pulumi.getter(name="inputMappingDefaultValues")
    def input_mapping_default_values(self) -> pulumi.Output[Optional['outputs.DomainInputMappingDefaultValues']]:
        """
        A `input_mapping_default_values` block as defined below.
        """
        return pulumi.get(self, "input_mapping_default_values")

    @property
    @pulumi.getter(name="inputMappingFields")
    def input_mapping_fields(self) -> pulumi.Output[Optional['outputs.DomainInputMappingFields']]:
        """
        A `input_mapping_fields` block as defined below.
        """
        return pulumi.get(self, "input_mapping_fields")

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "input_schema")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> pulumi.Output[str]:
        """
        The Primary Shared Access Key associated with the EventGrid Domain.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not public network access is allowed for this server. Defaults to `true`.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> pulumi.Output[str]:
        """
        The Secondary Shared Access Key associated with the EventGrid Domain.
        """
        return pulumi.get(self, "secondary_access_key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

