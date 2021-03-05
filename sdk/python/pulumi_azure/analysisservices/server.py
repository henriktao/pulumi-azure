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

__all__ = ['Server']


class Server(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_blob_container_uri: Optional[pulumi.Input[str]] = None,
                 enable_power_bi_service: Optional[pulumi.Input[bool]] = None,
                 ipv4_firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerIpv4FirewallRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 querypool_connection_mode: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Analysis Services Server.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        server = azure.analysisservices.Server("server",
            location="northeurope",
            resource_group_name=rg.name,
            sku="S0",
            admin_users=["myuser@domain.tld"],
            enable_power_bi_service=True,
            ipv4_firewall_rules=[azure.analysisservices.ServerIpv4FirewallRuleArgs(
                name="myRule1",
                range_start="210.117.252.0",
                range_end="210.117.252.255",
            )],
            tags={
                "abc": "123",
            })
        ```

        > **NOTE:** The server resource will automatically be started and stopped during an update if it is in `paused` state.

        ## Import

        Analysis Services Server can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:analysisservices/server:Server server /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AnalysisServices/servers/server1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_users: List of email addresses of admin users.
        :param pulumi.Input[str] backup_blob_container_uri: URI and SAS token for a blob container to store backups.
        :param pulumi.Input[bool] enable_power_bi_service: Indicates if the Power BI service is allowed to access or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerIpv4FirewallRuleArgs']]]] ipv4_firewall_rules: One or more `ipv4_firewall_rule` block(s) as defined below.
        :param pulumi.Input[str] location: The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the firewall rule.
        :param pulumi.Input[str] querypool_connection_mode: Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku: SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
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

            __props__['admin_users'] = admin_users
            __props__['backup_blob_container_uri'] = backup_blob_container_uri
            __props__['enable_power_bi_service'] = enable_power_bi_service
            __props__['ipv4_firewall_rules'] = ipv4_firewall_rules
            __props__['location'] = location
            __props__['name'] = name
            __props__['querypool_connection_mode'] = querypool_connection_mode
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['server_full_name'] = None
        super(Server, __self__).__init__(
            'azure:analysisservices/server:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backup_blob_container_uri: Optional[pulumi.Input[str]] = None,
            enable_power_bi_service: Optional[pulumi.Input[bool]] = None,
            ipv4_firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerIpv4FirewallRuleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            querypool_connection_mode: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            server_full_name: Optional[pulumi.Input[str]] = None,
            sku: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_users: List of email addresses of admin users.
        :param pulumi.Input[str] backup_blob_container_uri: URI and SAS token for a blob container to store backups.
        :param pulumi.Input[bool] enable_power_bi_service: Indicates if the Power BI service is allowed to access or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerIpv4FirewallRuleArgs']]]] ipv4_firewall_rules: One or more `ipv4_firewall_rule` block(s) as defined below.
        :param pulumi.Input[str] location: The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the firewall rule.
        :param pulumi.Input[str] querypool_connection_mode: Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] server_full_name: The full name of the Analysis Services Server.
        :param pulumi.Input[str] sku: SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_users"] = admin_users
        __props__["backup_blob_container_uri"] = backup_blob_container_uri
        __props__["enable_power_bi_service"] = enable_power_bi_service
        __props__["ipv4_firewall_rules"] = ipv4_firewall_rules
        __props__["location"] = location
        __props__["name"] = name
        __props__["querypool_connection_mode"] = querypool_connection_mode
        __props__["resource_group_name"] = resource_group_name
        __props__["server_full_name"] = server_full_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        return Server(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminUsers")
    def admin_users(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of email addresses of admin users.
        """
        return pulumi.get(self, "admin_users")

    @property
    @pulumi.getter(name="backupBlobContainerUri")
    def backup_blob_container_uri(self) -> pulumi.Output[Optional[str]]:
        """
        URI and SAS token for a blob container to store backups.
        """
        return pulumi.get(self, "backup_blob_container_uri")

    @property
    @pulumi.getter(name="enablePowerBiService")
    def enable_power_bi_service(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the Power BI service is allowed to access or not.
        """
        return pulumi.get(self, "enable_power_bi_service")

    @property
    @pulumi.getter(name="ipv4FirewallRules")
    def ipv4_firewall_rules(self) -> pulumi.Output[Optional[Sequence['outputs.ServerIpv4FirewallRule']]]:
        """
        One or more `ipv4_firewall_rule` block(s) as defined below.
        """
        return pulumi.get(self, "ipv4_firewall_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where the Analysis Services Server exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the firewall rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="querypoolConnectionMode")
    def querypool_connection_mode(self) -> pulumi.Output[str]:
        """
        Controls how the read-write server is used in the query pool. If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
        """
        return pulumi.get(self, "querypool_connection_mode")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which the Analysis Services Server should be exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serverFullName")
    def server_full_name(self) -> pulumi.Output[str]:
        """
        The full name of the Analysis Services Server.
        """
        return pulumi.get(self, "server_full_name")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[str]:
        """
        SKU for the Analysis Services Server. Possible values are: `D1`, `B1`, `B2`, `S0`, `S1`, `S2`, `S4`, `S8`, `S9`, `S8v2` and `S9v2`.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

