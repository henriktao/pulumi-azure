# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VirtualNetworkPeering']


class VirtualNetworkPeering(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_forwarded_traffic: Optional[pulumi.Input[bool]] = None,
                 allow_gateway_transit: Optional[pulumi.Input[bool]] = None,
                 allow_virtual_network_access: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_virtual_network_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 use_remote_gateways: Optional[pulumi.Input[bool]] = None,
                 virtual_network_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a virtual network peering which allows resources to access other
        resources in the linked virtual network.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West Europe")
        example_1_virtual_network = azure.network.VirtualNetwork("example-1VirtualNetwork",
            resource_group_name=example.name,
            address_spaces=["10.0.1.0/24"],
            location="West US")
        example_2_virtual_network = azure.network.VirtualNetwork("example-2VirtualNetwork",
            resource_group_name=example.name,
            address_spaces=["10.0.2.0/24"],
            location="West US")
        example_1_virtual_network_peering = azure.network.VirtualNetworkPeering("example-1VirtualNetworkPeering",
            resource_group_name=example.name,
            virtual_network_name=example_1_virtual_network.name,
            remote_virtual_network_id=example_2_virtual_network.id)
        example_2_virtual_network_peering = azure.network.VirtualNetworkPeering("example-2VirtualNetworkPeering",
            resource_group_name=example.name,
            virtual_network_name=example_2_virtual_network.name,
            remote_virtual_network_id=example_1_virtual_network.id)
        ```
        ## Note

        Virtual Network peerings cannot be created, updated or deleted concurrently.

        ## Import

        Virtual Network Peerings can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/virtualNetworkPeering:VirtualNetworkPeering examplePeering /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/virtualNetworkPeerings/myvnet1peering
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_forwarded_traffic: Controls if forwarded traffic from  VMs
               in the remote virtual network is allowed. Defaults to false.
        :param pulumi.Input[bool] allow_gateway_transit: Controls gatewayLinks can be used in the
               remote virtual network’s link to the local virtual network.
        :param pulumi.Input[bool] allow_virtual_network_access: Controls if the VMs in the remote
               virtual network can access VMs in the local virtual network. Defaults to
               true.
        :param pulumi.Input[str] name: The name of the virtual network peering. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] remote_virtual_network_id: The full Azure resource ID of the
               remote virtual network.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the virtual network peering. Changing this forces a new resource to be
               created.
        :param pulumi.Input[bool] use_remote_gateways: Controls if remote gateways can be used on
               the local virtual network. If the flag is set to `true`, and
               `allow_gateway_transit` on the remote peering is also `true`, virtual network will
               use gateways of remote virtual network for transit. Only one peering can
               have this flag set to `true`. This flag cannot be set if virtual network
               already has a gateway. Defaults to `false`.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network. Changing
               this forces a new resource to be created.
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

            __props__['allow_forwarded_traffic'] = allow_forwarded_traffic
            __props__['allow_gateway_transit'] = allow_gateway_transit
            __props__['allow_virtual_network_access'] = allow_virtual_network_access
            __props__['name'] = name
            if remote_virtual_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'remote_virtual_network_id'")
            __props__['remote_virtual_network_id'] = remote_virtual_network_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['use_remote_gateways'] = use_remote_gateways
            if virtual_network_name is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
        super(VirtualNetworkPeering, __self__).__init__(
            'azure:network/virtualNetworkPeering:VirtualNetworkPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_forwarded_traffic: Optional[pulumi.Input[bool]] = None,
            allow_gateway_transit: Optional[pulumi.Input[bool]] = None,
            allow_virtual_network_access: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_virtual_network_id: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            use_remote_gateways: Optional[pulumi.Input[bool]] = None,
            virtual_network_name: Optional[pulumi.Input[str]] = None) -> 'VirtualNetworkPeering':
        """
        Get an existing VirtualNetworkPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_forwarded_traffic: Controls if forwarded traffic from  VMs
               in the remote virtual network is allowed. Defaults to false.
        :param pulumi.Input[bool] allow_gateway_transit: Controls gatewayLinks can be used in the
               remote virtual network’s link to the local virtual network.
        :param pulumi.Input[bool] allow_virtual_network_access: Controls if the VMs in the remote
               virtual network can access VMs in the local virtual network. Defaults to
               true.
        :param pulumi.Input[str] name: The name of the virtual network peering. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] remote_virtual_network_id: The full Azure resource ID of the
               remote virtual network.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the virtual network peering. Changing this forces a new resource to be
               created.
        :param pulumi.Input[bool] use_remote_gateways: Controls if remote gateways can be used on
               the local virtual network. If the flag is set to `true`, and
               `allow_gateway_transit` on the remote peering is also `true`, virtual network will
               use gateways of remote virtual network for transit. Only one peering can
               have this flag set to `true`. This flag cannot be set if virtual network
               already has a gateway. Defaults to `false`.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network. Changing
               this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_forwarded_traffic"] = allow_forwarded_traffic
        __props__["allow_gateway_transit"] = allow_gateway_transit
        __props__["allow_virtual_network_access"] = allow_virtual_network_access
        __props__["name"] = name
        __props__["remote_virtual_network_id"] = remote_virtual_network_id
        __props__["resource_group_name"] = resource_group_name
        __props__["use_remote_gateways"] = use_remote_gateways
        __props__["virtual_network_name"] = virtual_network_name
        return VirtualNetworkPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowForwardedTraffic")
    def allow_forwarded_traffic(self) -> pulumi.Output[bool]:
        """
        Controls if forwarded traffic from  VMs
        in the remote virtual network is allowed. Defaults to false.
        """
        return pulumi.get(self, "allow_forwarded_traffic")

    @property
    @pulumi.getter(name="allowGatewayTransit")
    def allow_gateway_transit(self) -> pulumi.Output[bool]:
        """
        Controls gatewayLinks can be used in the
        remote virtual network’s link to the local virtual network.
        """
        return pulumi.get(self, "allow_gateway_transit")

    @property
    @pulumi.getter(name="allowVirtualNetworkAccess")
    def allow_virtual_network_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Controls if the VMs in the remote
        virtual network can access VMs in the local virtual network. Defaults to
        true.
        """
        return pulumi.get(self, "allow_virtual_network_access")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the virtual network peering. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteVirtualNetworkId")
    def remote_virtual_network_id(self) -> pulumi.Output[str]:
        """
        The full Azure resource ID of the
        remote virtual network.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "remote_virtual_network_id")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which to
        create the virtual network peering. Changing this forces a new resource to be
        created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="useRemoteGateways")
    def use_remote_gateways(self) -> pulumi.Output[bool]:
        """
        Controls if remote gateways can be used on
        the local virtual network. If the flag is set to `true`, and
        `allow_gateway_transit` on the remote peering is also `true`, virtual network will
        use gateways of remote virtual network for transit. Only one peering can
        have this flag set to `true`. This flag cannot be set if virtual network
        already has a gateway. Defaults to `false`.
        """
        return pulumi.get(self, "use_remote_gateways")

    @property
    @pulumi.getter(name="virtualNetworkName")
    def virtual_network_name(self) -> pulumi.Output[str]:
        """
        The name of the virtual network. Changing
        this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

