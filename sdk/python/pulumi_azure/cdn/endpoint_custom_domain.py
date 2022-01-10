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

__all__ = ['EndpointCustomDomainArgs', 'EndpointCustomDomain']

@pulumi.input_type
class EndpointCustomDomainArgs:
    def __init__(__self__, *,
                 cdn_endpoint_id: pulumi.Input[str],
                 host_name: pulumi.Input[str],
                 cdn_managed_https: Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 user_managed_https: Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']] = None):
        """
        The set of arguments for constructing a EndpointCustomDomain resource.
        :param pulumi.Input[str] cdn_endpoint_id: The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[str] host_name: The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs'] cdn_managed_https: A `cdn_managed_https` block as defined below.
        :param pulumi.Input[str] name: The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs'] user_managed_https: A `user_managed_https` block as defined below.
        """
        pulumi.set(__self__, "cdn_endpoint_id", cdn_endpoint_id)
        pulumi.set(__self__, "host_name", host_name)
        if cdn_managed_https is not None:
            pulumi.set(__self__, "cdn_managed_https", cdn_managed_https)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if user_managed_https is not None:
            pulumi.set(__self__, "user_managed_https", user_managed_https)

    @property
    @pulumi.getter(name="cdnEndpointId")
    def cdn_endpoint_id(self) -> pulumi.Input[str]:
        """
        The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "cdn_endpoint_id")

    @cdn_endpoint_id.setter
    def cdn_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cdn_endpoint_id", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Input[str]:
        """
        The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="cdnManagedHttps")
    def cdn_managed_https(self) -> Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']]:
        """
        A `cdn_managed_https` block as defined below.
        """
        return pulumi.get(self, "cdn_managed_https")

    @cdn_managed_https.setter
    def cdn_managed_https(self, value: Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']]):
        pulumi.set(self, "cdn_managed_https", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="userManagedHttps")
    def user_managed_https(self) -> Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']]:
        """
        A `user_managed_https` block as defined below.
        """
        return pulumi.get(self, "user_managed_https")

    @user_managed_https.setter
    def user_managed_https(self, value: Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']]):
        pulumi.set(self, "user_managed_https", value)


@pulumi.input_type
class _EndpointCustomDomainState:
    def __init__(__self__, *,
                 cdn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 cdn_managed_https: Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 user_managed_https: Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']] = None):
        """
        Input properties used for looking up and filtering EndpointCustomDomain resources.
        :param pulumi.Input[str] cdn_endpoint_id: The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs'] cdn_managed_https: A `cdn_managed_https` block as defined below.
        :param pulumi.Input[str] host_name: The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[str] name: The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs'] user_managed_https: A `user_managed_https` block as defined below.
        """
        if cdn_endpoint_id is not None:
            pulumi.set(__self__, "cdn_endpoint_id", cdn_endpoint_id)
        if cdn_managed_https is not None:
            pulumi.set(__self__, "cdn_managed_https", cdn_managed_https)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if user_managed_https is not None:
            pulumi.set(__self__, "user_managed_https", user_managed_https)

    @property
    @pulumi.getter(name="cdnEndpointId")
    def cdn_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "cdn_endpoint_id")

    @cdn_endpoint_id.setter
    def cdn_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cdn_endpoint_id", value)

    @property
    @pulumi.getter(name="cdnManagedHttps")
    def cdn_managed_https(self) -> Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']]:
        """
        A `cdn_managed_https` block as defined below.
        """
        return pulumi.get(self, "cdn_managed_https")

    @cdn_managed_https.setter
    def cdn_managed_https(self, value: Optional[pulumi.Input['EndpointCustomDomainCdnManagedHttpsArgs']]):
        pulumi.set(self, "cdn_managed_https", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="userManagedHttps")
    def user_managed_https(self) -> Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']]:
        """
        A `user_managed_https` block as defined below.
        """
        return pulumi.get(self, "user_managed_https")

    @user_managed_https.setter
    def user_managed_https(self, value: Optional[pulumi.Input['EndpointCustomDomainUserManagedHttpsArgs']]):
        pulumi.set(self, "user_managed_https", value)


class EndpointCustomDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 cdn_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainCdnManagedHttpsArgs']]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 user_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainUserManagedHttpsArgs']]] = None,
                 __props__=None):
        """
        Manages a Custom Domain for a CDN Endpoint.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="west europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_profile = azure.cdn.Profile("exampleProfile",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard_Verizon")
        example_endpoint = azure.cdn.Endpoint("exampleEndpoint",
            profile_name=example_profile.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            origins=[azure.cdn.EndpointOriginArgs(
                name="example",
                host_name=example_account.primary_blob_host,
            )])
        example_zone = azure.dns.get_zone(name="example-domain.com",
            resource_group_name="domain-rg")
        example_c_name_record = azure.dns.CNameRecord("exampleCNameRecord",
            zone_name=example_zone.name,
            resource_group_name=example_zone.resource_group_name,
            ttl=3600,
            target_resource_id=example_endpoint.id)
        example_endpoint_custom_domain = azure.cdn.EndpointCustomDomain("exampleEndpointCustomDomain",
            cdn_endpoint_id=example_endpoint.id,
            host_name=example_c_name_record.name.apply(lambda name: f"{name}.{example_zone.name}"))
        ```

        ## Import

        CDN Endpoint Custom Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:cdn/endpointCustomDomain:EndpointCustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_endpoint_id: The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[pulumi.InputType['EndpointCustomDomainCdnManagedHttpsArgs']] cdn_managed_https: A `cdn_managed_https` block as defined below.
        :param pulumi.Input[str] host_name: The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[str] name: The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[pulumi.InputType['EndpointCustomDomainUserManagedHttpsArgs']] user_managed_https: A `user_managed_https` block as defined below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointCustomDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Custom Domain for a CDN Endpoint.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="west europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_profile = azure.cdn.Profile("exampleProfile",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard_Verizon")
        example_endpoint = azure.cdn.Endpoint("exampleEndpoint",
            profile_name=example_profile.name,
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            origins=[azure.cdn.EndpointOriginArgs(
                name="example",
                host_name=example_account.primary_blob_host,
            )])
        example_zone = azure.dns.get_zone(name="example-domain.com",
            resource_group_name="domain-rg")
        example_c_name_record = azure.dns.CNameRecord("exampleCNameRecord",
            zone_name=example_zone.name,
            resource_group_name=example_zone.resource_group_name,
            ttl=3600,
            target_resource_id=example_endpoint.id)
        example_endpoint_custom_domain = azure.cdn.EndpointCustomDomain("exampleEndpointCustomDomain",
            cdn_endpoint_id=example_endpoint.id,
            host_name=example_c_name_record.name.apply(lambda name: f"{name}.{example_zone.name}"))
        ```

        ## Import

        CDN Endpoint Custom Domains can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:cdn/endpointCustomDomain:EndpointCustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
        ```

        :param str resource_name: The name of the resource.
        :param EndpointCustomDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointCustomDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cdn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 cdn_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainCdnManagedHttpsArgs']]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 user_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainUserManagedHttpsArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointCustomDomainArgs.__new__(EndpointCustomDomainArgs)

            if cdn_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'cdn_endpoint_id'")
            __props__.__dict__["cdn_endpoint_id"] = cdn_endpoint_id
            __props__.__dict__["cdn_managed_https"] = cdn_managed_https
            if host_name is None and not opts.urn:
                raise TypeError("Missing required property 'host_name'")
            __props__.__dict__["host_name"] = host_name
            __props__.__dict__["name"] = name
            __props__.__dict__["user_managed_https"] = user_managed_https
        super(EndpointCustomDomain, __self__).__init__(
            'azure:cdn/endpointCustomDomain:EndpointCustomDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cdn_endpoint_id: Optional[pulumi.Input[str]] = None,
            cdn_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainCdnManagedHttpsArgs']]] = None,
            host_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            user_managed_https: Optional[pulumi.Input[pulumi.InputType['EndpointCustomDomainUserManagedHttpsArgs']]] = None) -> 'EndpointCustomDomain':
        """
        Get an existing EndpointCustomDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cdn_endpoint_id: The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[pulumi.InputType['EndpointCustomDomainCdnManagedHttpsArgs']] cdn_managed_https: A `cdn_managed_https` block as defined below.
        :param pulumi.Input[str] host_name: The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[str] name: The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        :param pulumi.Input[pulumi.InputType['EndpointCustomDomainUserManagedHttpsArgs']] user_managed_https: A `user_managed_https` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointCustomDomainState.__new__(_EndpointCustomDomainState)

        __props__.__dict__["cdn_endpoint_id"] = cdn_endpoint_id
        __props__.__dict__["cdn_managed_https"] = cdn_managed_https
        __props__.__dict__["host_name"] = host_name
        __props__.__dict__["name"] = name
        __props__.__dict__["user_managed_https"] = user_managed_https
        return EndpointCustomDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cdnEndpointId")
    def cdn_endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the CDN Endpoint. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "cdn_endpoint_id")

    @property
    @pulumi.getter(name="cdnManagedHttps")
    def cdn_managed_https(self) -> pulumi.Output[Optional['outputs.EndpointCustomDomainCdnManagedHttps']]:
        """
        A `cdn_managed_https` block as defined below.
        """
        return pulumi.get(self, "cdn_managed_https")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Output[str]:
        """
        The host name of the custom domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this CDN Endpoint Custom Domain. Changing this forces a new CDN Endpoint Custom Domain to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="userManagedHttps")
    def user_managed_https(self) -> pulumi.Output[Optional['outputs.EndpointCustomDomainUserManagedHttps']]:
        """
        A `user_managed_https` block as defined below.
        """
        return pulumi.get(self, "user_managed_https")

