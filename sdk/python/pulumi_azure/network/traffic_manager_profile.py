# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['TrafficManagerProfile']


class TrafficManagerProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_config: Optional[pulumi.Input[pulumi.InputType['TrafficManagerProfileDnsConfigArgs']]] = None,
                 max_return: Optional[pulumi.Input[int]] = None,
                 monitor_config: Optional[pulumi.Input[pulumi.InputType['TrafficManagerProfileMonitorConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile_status: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_routing_method: Optional[pulumi.Input[str]] = None,
                 traffic_view_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Traffic Manager Profile to which multiple endpoints can be attached.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        server = random.RandomId("server",
            keepers={
                "azi_id": 1,
            },
            byte_length=8)
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_traffic_manager_profile = azure.network.TrafficManagerProfile("exampleTrafficManagerProfile",
            resource_group_name=example_resource_group.name,
            traffic_routing_method="Weighted",
            dns_config=azure.network.TrafficManagerProfileDnsConfigArgs(
                relative_name=server.hex,
                ttl=100,
            ),
            monitor_config=azure.network.TrafficManagerProfileMonitorConfigArgs(
                protocol="http",
                port=80,
                path="/",
                interval_in_seconds=30,
                timeout_in_seconds=9,
                tolerated_number_of_failures=3,
            ),
            tags={
                "environment": "Production",
            })
        ```

        ## Import

        Traffic Manager Profiles can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/trafficManagerProfile:TrafficManagerProfile exampleProfile /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TrafficManagerProfileDnsConfigArgs']] dns_config: This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        :param pulumi.Input[int] max_return: The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
        :param pulumi.Input[pulumi.InputType['TrafficManagerProfileMonitorConfigArgs']] monitor_config: This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        :param pulumi.Input[str] name: The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        :param pulumi.Input[str] profile_status: The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Traffic Manager profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] traffic_routing_method: Specifies the algorithm used to route traffic, possible values are:
        :param pulumi.Input[bool] traffic_view_enabled: Indicates whether Traffic View is enabled for the Traffic Manager profile.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if dns_config is None and not opts.urn:
                raise TypeError("Missing required property 'dns_config'")
            __props__['dns_config'] = dns_config
            __props__['max_return'] = max_return
            if monitor_config is None and not opts.urn:
                raise TypeError("Missing required property 'monitor_config'")
            __props__['monitor_config'] = monitor_config
            __props__['name'] = name
            __props__['profile_status'] = profile_status
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if traffic_routing_method is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_routing_method'")
            __props__['traffic_routing_method'] = traffic_routing_method
            __props__['traffic_view_enabled'] = traffic_view_enabled
            __props__['fqdn'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:trafficmanager/profile:Profile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(TrafficManagerProfile, __self__).__init__(
            'azure:network/trafficManagerProfile:TrafficManagerProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_config: Optional[pulumi.Input[pulumi.InputType['TrafficManagerProfileDnsConfigArgs']]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            max_return: Optional[pulumi.Input[int]] = None,
            monitor_config: Optional[pulumi.Input[pulumi.InputType['TrafficManagerProfileMonitorConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            profile_status: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            traffic_routing_method: Optional[pulumi.Input[str]] = None,
            traffic_view_enabled: Optional[pulumi.Input[bool]] = None) -> 'TrafficManagerProfile':
        """
        Get an existing TrafficManagerProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TrafficManagerProfileDnsConfigArgs']] dns_config: This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        :param pulumi.Input[str] fqdn: The FQDN of the created Profile.
        :param pulumi.Input[int] max_return: The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
        :param pulumi.Input[pulumi.InputType['TrafficManagerProfileMonitorConfigArgs']] monitor_config: This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        :param pulumi.Input[str] name: The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        :param pulumi.Input[str] profile_status: The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Traffic Manager profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] traffic_routing_method: Specifies the algorithm used to route traffic, possible values are:
        :param pulumi.Input[bool] traffic_view_enabled: Indicates whether Traffic View is enabled for the Traffic Manager profile.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dns_config"] = dns_config
        __props__["fqdn"] = fqdn
        __props__["max_return"] = max_return
        __props__["monitor_config"] = monitor_config
        __props__["name"] = name
        __props__["profile_status"] = profile_status
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["traffic_routing_method"] = traffic_routing_method
        __props__["traffic_view_enabled"] = traffic_view_enabled
        return TrafficManagerProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> pulumi.Output['outputs.TrafficManagerProfileDnsConfig']:
        """
        This block specifies the DNS configuration of the Profile, it supports the fields documented below.
        """
        return pulumi.get(self, "dns_config")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN of the created Profile.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="maxReturn")
    def max_return(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of endpoints to return for DNS queries to this Profile. Possible values range from `1` to `8`.
        """
        return pulumi.get(self, "max_return")

    @property
    @pulumi.getter(name="monitorConfig")
    def monitor_config(self) -> pulumi.Output['outputs.TrafficManagerProfileMonitorConfig']:
        """
        This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
        """
        return pulumi.get(self, "monitor_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Traffic Manager profile. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileStatus")
    def profile_status(self) -> pulumi.Output[str]:
        """
        The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
        """
        return pulumi.get(self, "profile_status")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to create the Traffic Manager profile.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficRoutingMethod")
    def traffic_routing_method(self) -> pulumi.Output[str]:
        """
        Specifies the algorithm used to route traffic, possible values are:
        """
        return pulumi.get(self, "traffic_routing_method")

    @property
    @pulumi.getter(name="trafficViewEnabled")
    def traffic_view_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether Traffic View is enabled for the Traffic Manager profile.
        """
        return pulumi.get(self, "traffic_view_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

