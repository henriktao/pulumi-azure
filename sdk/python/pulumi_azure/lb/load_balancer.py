# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LoadBalancer(pulumi.CustomResource):
    frontend_ip_configurations: pulumi.Output[list]
    """
    One or multiple `frontend_ip_configuration` blocks as documented below.

      * `id` (`str`) - The id of the Frontend IP Configuration.
      * `inbound_nat_rules` (`list`) - The list of IDs of inbound rules that use this frontend IP.
      * `load_balancer_rules` (`list`) - The list of IDs of load balancing rules that use this frontend IP.
      * `name` (`str`) - Specifies the name of the frontend ip configuration.
      * `outboundRules` (`list`) - The list of IDs outbound rules that use this frontend IP.
      * `private_ip_address` (`str`) - Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
      * `privateIpAddressAllocation` (`str`) - The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
      * `privateIpAddressVersion` (`str`) - The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
      * `public_ip_address_id` (`str`) - The ID of a Public IP Address which should be associated with the Load Balancer.
      * `public_ip_prefix_id` (`str`) - The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
      * `subnet_id` (`str`) - The ID of the Subnet which should be associated with the IP Configuration.
      * `zones` (`str`) - A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure Region where the Load Balancer should be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Load Balancer.
    """
    private_ip_address: pulumi.Output[str]
    """
    Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
    """
    private_ip_addresses: pulumi.Output[list]
    """
    The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which to create the Load Balancer.
    """
    sku: pulumi.Output[str]
    """
    The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, frontend_ip_configurations=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Load Balancer Resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_public_ip = azure.network.PublicIp("examplePublicIp",
            location="West US",
            resource_group_name=example_resource_group.name,
            allocation_method="Static")
        example_load_balancer = azure.lb.LoadBalancer("exampleLoadBalancer",
            location="West US",
            resource_group_name=example_resource_group.name,
            frontend_ip_configuration=[{
                "name": "PublicIPAddress",
                "public_ip_address_id": example_public_ip.id,
            }])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] frontend_ip_configurations: One or multiple `frontend_ip_configuration` blocks as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure Region where the Load Balancer should be created.
        :param pulumi.Input[str] name: Specifies the name of the Load Balancer.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the Load Balancer.
        :param pulumi.Input[str] sku: The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **frontend_ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - The id of the Frontend IP Configuration.
          * `inbound_nat_rules` (`pulumi.Input[list]`) - The list of IDs of inbound rules that use this frontend IP.
          * `load_balancer_rules` (`pulumi.Input[list]`) - The list of IDs of load balancing rules that use this frontend IP.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the frontend ip configuration.
          * `outboundRules` (`pulumi.Input[list]`) - The list of IDs outbound rules that use this frontend IP.
          * `private_ip_address` (`pulumi.Input[str]`) - Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
          * `privateIpAddressAllocation` (`pulumi.Input[str]`) - The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
          * `privateIpAddressVersion` (`pulumi.Input[str]`) - The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
          * `public_ip_address_id` (`pulumi.Input[str]`) - The ID of a Public IP Address which should be associated with the Load Balancer.
          * `public_ip_prefix_id` (`pulumi.Input[str]`) - The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
          * `subnet_id` (`pulumi.Input[str]`) - The ID of the Subnet which should be associated with the IP Configuration.
          * `zones` (`pulumi.Input[str]`) - A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['frontend_ip_configurations'] = frontend_ip_configurations
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['private_ip_address'] = None
            __props__['private_ip_addresses'] = None
        super(LoadBalancer, __self__).__init__(
            'azure:lb/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, frontend_ip_configurations=None, location=None, name=None, private_ip_address=None, private_ip_addresses=None, resource_group_name=None, sku=None, tags=None):
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] frontend_ip_configurations: One or multiple `frontend_ip_configuration` blocks as documented below.
        :param pulumi.Input[str] location: Specifies the supported Azure Region where the Load Balancer should be created.
        :param pulumi.Input[str] name: Specifies the name of the Load Balancer.
        :param pulumi.Input[str] private_ip_address: Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
        :param pulumi.Input[list] private_ip_addresses: The list of private IP address assigned to the load balancer in `frontend_ip_configuration` blocks, if any.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the Load Balancer.
        :param pulumi.Input[str] sku: The SKU of the Azure Load Balancer. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **frontend_ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - The id of the Frontend IP Configuration.
          * `inbound_nat_rules` (`pulumi.Input[list]`) - The list of IDs of inbound rules that use this frontend IP.
          * `load_balancer_rules` (`pulumi.Input[list]`) - The list of IDs of load balancing rules that use this frontend IP.
          * `name` (`pulumi.Input[str]`) - Specifies the name of the frontend ip configuration.
          * `outboundRules` (`pulumi.Input[list]`) - The list of IDs outbound rules that use this frontend IP.
          * `private_ip_address` (`pulumi.Input[str]`) - Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.
          * `privateIpAddressAllocation` (`pulumi.Input[str]`) - The allocation method for the Private IP Address used by this Load Balancer. Possible values as `Dynamic` and `Static`.
          * `privateIpAddressVersion` (`pulumi.Input[str]`) - The version of IP that the Private IP Address is. Possible values are `IPv4` or `IPv6`.
          * `public_ip_address_id` (`pulumi.Input[str]`) - The ID of a Public IP Address which should be associated with the Load Balancer.
          * `public_ip_prefix_id` (`pulumi.Input[str]`) - The ID of a Public IP Prefix which should be associated with the Load Balancer. Public IP Prefix can only be used with outbound rules.
          * `subnet_id` (`pulumi.Input[str]`) - The ID of the Subnet which should be associated with the IP Configuration.
          * `zones` (`pulumi.Input[str]`) - A list of Availability Zones which the Load Balancer's IP Addresses should be created in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["frontend_ip_configurations"] = frontend_ip_configurations
        __props__["location"] = location
        __props__["name"] = name
        __props__["private_ip_address"] = private_ip_address
        __props__["private_ip_addresses"] = private_ip_addresses
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

