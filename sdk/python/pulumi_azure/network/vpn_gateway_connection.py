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

__all__ = ['VpnGatewayConnectionArgs', 'VpnGatewayConnection']

@pulumi.input_type
class VpnGatewayConnectionArgs:
    def __init__(__self__, *,
                 remote_vpn_site_id: pulumi.Input[str],
                 vpn_gateway_id: pulumi.Input[str],
                 vpn_links: pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]],
                 internet_security_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routings: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]] = None):
        """
        The set of arguments for constructing a VpnGatewayConnection resource.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]] vpn_links: One or more `vpn_link` blocks as defined below.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        """
        pulumi.set(__self__, "remote_vpn_site_id", remote_vpn_site_id)
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        pulumi.set(__self__, "vpn_links", vpn_links)
        if internet_security_enabled is not None:
            pulumi.set(__self__, "internet_security_enabled", internet_security_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if routings is not None:
            pulumi.set(__self__, "routings", routings)

    @property
    @pulumi.getter(name="remoteVpnSiteId")
    def remote_vpn_site_id(self) -> pulumi.Input[str]:
        """
        The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "remote_vpn_site_id")

    @remote_vpn_site_id.setter
    def remote_vpn_site_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_vpn_site_id", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_gateway_id", value)

    @property
    @pulumi.getter(name="vpnLinks")
    def vpn_links(self) -> pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]]:
        """
        One or more `vpn_link` blocks as defined below.
        """
        return pulumi.get(self, "vpn_links")

    @vpn_links.setter
    def vpn_links(self, value: pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]]):
        pulumi.set(self, "vpn_links", value)

    @property
    @pulumi.getter(name="internetSecurityEnabled")
    def internet_security_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        """
        return pulumi.get(self, "internet_security_enabled")

    @internet_security_enabled.setter
    def internet_security_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internet_security_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def routings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]]:
        """
        A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        """
        return pulumi.get(self, "routings")

    @routings.setter
    def routings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]]):
        pulumi.set(self, "routings", value)


@pulumi.input_type
class _VpnGatewayConnectionState:
    def __init__(__self__, *,
                 internet_security_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
                 routings: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]]] = None):
        """
        Input properties used for looking up and filtering VpnGatewayConnection resources.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]] vpn_links: One or more `vpn_link` blocks as defined below.
        """
        if internet_security_enabled is not None:
            pulumi.set(__self__, "internet_security_enabled", internet_security_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote_vpn_site_id is not None:
            pulumi.set(__self__, "remote_vpn_site_id", remote_vpn_site_id)
        if routings is not None:
            pulumi.set(__self__, "routings", routings)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        if vpn_links is not None:
            pulumi.set(__self__, "vpn_links", vpn_links)

    @property
    @pulumi.getter(name="internetSecurityEnabled")
    def internet_security_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        """
        return pulumi.get(self, "internet_security_enabled")

    @internet_security_enabled.setter
    def internet_security_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internet_security_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="remoteVpnSiteId")
    def remote_vpn_site_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "remote_vpn_site_id")

    @remote_vpn_site_id.setter
    def remote_vpn_site_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_vpn_site_id", value)

    @property
    @pulumi.getter
    def routings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]]:
        """
        A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        """
        return pulumi.get(self, "routings")

    @routings.setter
    def routings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionRoutingArgs']]]]):
        pulumi.set(self, "routings", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)

    @property
    @pulumi.getter(name="vpnLinks")
    def vpn_links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]]]:
        """
        One or more `vpn_link` blocks as defined below.
        """
        return pulumi.get(self, "vpn_links")

    @vpn_links.setter
    def vpn_links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnGatewayConnectionVpnLinkArgs']]]]):
        pulumi.set(self, "vpn_links", value)


class VpnGatewayConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internet_security_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
                 routings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]]] = None,
                 __props__=None):
        """
        Manages a VPN Gateway Connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_wan = azure.network.VirtualWan("exampleVirtualWan",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_virtual_hub = azure.network.VirtualHub("exampleVirtualHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=example_virtual_wan.id,
            address_prefix="10.0.0.0/24")
        example_vpn_gateway = azure.network.VpnGateway("exampleVpnGateway",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            virtual_hub_id=example_virtual_hub.id)
        example_vpn_site = azure.network.VpnSite("exampleVpnSite",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            virtual_wan_id=example_virtual_wan.id,
            links=[
                azure.network.VpnSiteLinkArgs(
                    name="link1",
                    ip_address="10.1.0.0",
                ),
                azure.network.VpnSiteLinkArgs(
                    name="link2",
                    ip_address="10.2.0.0",
                ),
            ])
        example_vpn_gateway_connection = azure.network.VpnGatewayConnection("exampleVpnGatewayConnection",
            vpn_gateway_id=example_vpn_gateway.id,
            remote_vpn_site_id=example_vpn_site.id,
            vpn_links=[
                azure.network.VpnGatewayConnectionVpnLinkArgs(
                    name="link1",
                    vpn_site_link_id=example_vpn_site.links[0].id,
                ),
                azure.network.VpnGatewayConnectionVpnLinkArgs(
                    name="link2",
                    vpn_site_link_id=example_vpn_site.links[1].id,
                ),
            ])
        ```

        ## Import

        VPN Gateway Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/vpnGatewayConnection:VpnGatewayConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/conn1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]] vpn_links: One or more `vpn_link` blocks as defined below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnGatewayConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a VPN Gateway Connection.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_wan = azure.network.VirtualWan("exampleVirtualWan",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_virtual_hub = azure.network.VirtualHub("exampleVirtualHub",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            virtual_wan_id=example_virtual_wan.id,
            address_prefix="10.0.0.0/24")
        example_vpn_gateway = azure.network.VpnGateway("exampleVpnGateway",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            virtual_hub_id=example_virtual_hub.id)
        example_vpn_site = azure.network.VpnSite("exampleVpnSite",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            virtual_wan_id=example_virtual_wan.id,
            links=[
                azure.network.VpnSiteLinkArgs(
                    name="link1",
                    ip_address="10.1.0.0",
                ),
                azure.network.VpnSiteLinkArgs(
                    name="link2",
                    ip_address="10.2.0.0",
                ),
            ])
        example_vpn_gateway_connection = azure.network.VpnGatewayConnection("exampleVpnGatewayConnection",
            vpn_gateway_id=example_vpn_gateway.id,
            remote_vpn_site_id=example_vpn_site.id,
            vpn_links=[
                azure.network.VpnGatewayConnectionVpnLinkArgs(
                    name="link1",
                    vpn_site_link_id=example_vpn_site.links[0].id,
                ),
                azure.network.VpnGatewayConnectionVpnLinkArgs(
                    name="link2",
                    vpn_site_link_id=example_vpn_site.links[1].id,
                ),
            ])
        ```

        ## Import

        VPN Gateway Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/vpnGatewayConnection:VpnGatewayConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/conn1
        ```

        :param str resource_name: The name of the resource.
        :param VpnGatewayConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnGatewayConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internet_security_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
                 routings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]]] = None,
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
            __props__ = VpnGatewayConnectionArgs.__new__(VpnGatewayConnectionArgs)

            __props__.__dict__["internet_security_enabled"] = internet_security_enabled
            __props__.__dict__["name"] = name
            if remote_vpn_site_id is None and not opts.urn:
                raise TypeError("Missing required property 'remote_vpn_site_id'")
            __props__.__dict__["remote_vpn_site_id"] = remote_vpn_site_id
            __props__.__dict__["routings"] = routings
            if vpn_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
            if vpn_links is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_links'")
            __props__.__dict__["vpn_links"] = vpn_links
        super(VpnGatewayConnection, __self__).__init__(
            'azure:network/vpnGatewayConnection:VpnGatewayConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            internet_security_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_vpn_site_id: Optional[pulumi.Input[str]] = None,
            routings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None,
            vpn_links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]]] = None) -> 'VpnGatewayConnection':
        """
        Get an existing VpnGatewayConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internet_security_enabled: Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        :param pulumi.Input[str] name: The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[str] remote_vpn_site_id: The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionRoutingArgs']]]] routings: A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnGatewayConnectionVpnLinkArgs']]]] vpn_links: One or more `vpn_link` blocks as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnGatewayConnectionState.__new__(_VpnGatewayConnectionState)

        __props__.__dict__["internet_security_enabled"] = internet_security_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["remote_vpn_site_id"] = remote_vpn_site_id
        __props__.__dict__["routings"] = routings
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        __props__.__dict__["vpn_links"] = vpn_links
        return VpnGatewayConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="internetSecurityEnabled")
    def internet_security_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Internet Security is enabled for this VPN Connection. Defaults to `false`.
        """
        return pulumi.get(self, "internet_security_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this VPN Gateway Connection. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteVpnSiteId")
    def remote_vpn_site_id(self) -> pulumi.Output[str]:
        """
        The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "remote_vpn_site_id")

    @property
    @pulumi.getter
    def routings(self) -> pulumi.Output[Sequence['outputs.VpnGatewayConnectionRouting']]:
        """
        A `routing` block as defined below. If this is not specified, there will be a default route table created implicitly.
        """
        return pulumi.get(self, "routings")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @property
    @pulumi.getter(name="vpnLinks")
    def vpn_links(self) -> pulumi.Output[Sequence['outputs.VpnGatewayConnectionVpnLink']]:
        """
        One or more `vpn_link` blocks as defined below.
        """
        return pulumi.get(self, "vpn_links")

