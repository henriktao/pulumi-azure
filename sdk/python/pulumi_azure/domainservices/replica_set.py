# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReplicaSetArgs', 'ReplicaSet']

@pulumi.input_type
class ReplicaSetArgs:
    def __init__(__self__, *,
                 domain_service_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReplicaSet resource.
        :param pulumi.Input[str] domain_service_id: The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet in which to place this Replica Set.
        :param pulumi.Input[str] location: The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "domain_service_id", domain_service_id)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter(name="domainServiceId")
    def domain_service_id(self) -> pulumi.Input[str]:
        """
        The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "domain_service_id")

    @domain_service_id.setter
    def domain_service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_service_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the subnet in which to place this Replica Set.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)


@pulumi.input_type
class _ReplicaSetState:
    def __init__(__self__, *,
                 domain_controller_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_service_id: Optional[pulumi.Input[str]] = None,
                 external_access_ip_address: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 service_status: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicaSet resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_controller_ip_addresses: A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        :param pulumi.Input[str] domain_service_id: The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] external_access_ip_address: The publicly routable IP address for the domain controllers in this Replica Set.
        :param pulumi.Input[str] location: The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_status: The current service status for the replica set.
        :param pulumi.Input[str] subnet_id: The ID of the subnet in which to place this Replica Set.
        """
        if domain_controller_ip_addresses is not None:
            pulumi.set(__self__, "domain_controller_ip_addresses", domain_controller_ip_addresses)
        if domain_service_id is not None:
            pulumi.set(__self__, "domain_service_id", domain_service_id)
        if external_access_ip_address is not None:
            pulumi.set(__self__, "external_access_ip_address", external_access_ip_address)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if service_status is not None:
            pulumi.set(__self__, "service_status", service_status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="domainControllerIpAddresses")
    def domain_controller_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        """
        return pulumi.get(self, "domain_controller_ip_addresses")

    @domain_controller_ip_addresses.setter
    def domain_controller_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domain_controller_ip_addresses", value)

    @property
    @pulumi.getter(name="domainServiceId")
    def domain_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "domain_service_id")

    @domain_service_id.setter
    def domain_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_service_id", value)

    @property
    @pulumi.getter(name="externalAccessIpAddress")
    def external_access_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The publicly routable IP address for the domain controllers in this Replica Set.
        """
        return pulumi.get(self, "external_access_ip_address")

    @external_access_ip_address.setter
    def external_access_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_access_ip_address", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> Optional[pulumi.Input[str]]:
        """
        The current service status for the replica set.
        """
        return pulumi.get(self, "service_status")

    @service_status.setter
    def service_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the subnet in which to place this Replica Set.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class ReplicaSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_service_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Replica Set for an Active Directory Domain Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        primary_resource_group = azure.core.ResourceGroup("primaryResourceGroup", location="West Europe")
        primary_virtual_network = azure.network.VirtualNetwork("primaryVirtualNetwork",
            location=primary_resource_group.location,
            resource_group_name=primary_resource_group.name,
            address_spaces=["10.0.1.0/16"])
        primary_subnet = azure.network.Subnet("primarySubnet",
            resource_group_name=primary_resource_group.name,
            virtual_network_name=primary_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        primary_network_security_group = azure.network.NetworkSecurityGroup("primaryNetworkSecurityGroup",
            location=primary_resource_group.location,
            resource_group_name=primary_resource_group.name,
            security_rules=[
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowSyncWithAzureAD",
                    priority=101,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="443",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowRD",
                    priority=201,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="3389",
                    source_address_prefix="CorpNetSaw",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowPSRemoting",
                    priority=301,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="5986",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowLDAPS",
                    priority=401,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="636",
                    source_address_prefix="*",
                    destination_address_prefix="*",
                ),
            ])
        primary_subnet_network_security_group_association = azure.network.SubnetNetworkSecurityGroupAssociation("primarySubnetNetworkSecurityGroupAssociation",
            subnet_id=primary_subnet.id,
            network_security_group_id=primary_network_security_group.id)
        dc_admins = azuread.Group("dcAdmins")
        admin_user = azuread.User("adminUser",
            user_principal_name="dc-admin@$hashicorp-example.net",
            display_name="DC Administrator",
            password="Pa55w0Rd!!1")
        admin_group_member = azuread.GroupMember("adminGroupMember",
            group_object_id=dc_admins.object_id,
            member_object_id=admin_user.object_id)
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id="2565bd9d-da50-47d4-8b85-4c97f669dc36")
        # published app for domain services
        aadds = azure.core.ResourceGroup("aadds", location="westeurope")
        example_service = azure.domainservices.Service("exampleService",
            location=aadds.location,
            resource_group_name=aadds.name,
            domain_name="widgetslogin.net",
            sku="Enterprise",
            filtered_sync_enabled=False,
            initial_replica_set=azure.domainservices.ServiceInitialReplicaSetArgs(
                location=primary_virtual_network.location,
                subnet_id=primary_subnet.id,
            ),
            notifications=azure.domainservices.ServiceNotificationsArgs(
                additional_recipients=[
                    "notifyA@example.net",
                    "notifyB@example.org",
                ],
                notify_dc_admins=True,
                notify_global_admins=True,
            ),
            security=azure.domainservices.ServiceSecurityArgs(
                sync_kerberos_passwords=True,
                sync_ntlm_passwords=True,
                sync_on_prem_passwords=True,
            ),
            tags={
                "Environment": "prod",
            },
            opts=pulumi.ResourceOptions(depends_on=[
                    example_service_principal,
                    primary_subnet_network_security_group_association,
                ]))
        replica_resource_group = azure.core.ResourceGroup("replicaResourceGroup", location="North Europe")
        replica_virtual_network = azure.network.VirtualNetwork("replicaVirtualNetwork",
            location=replica_resource_group.location,
            resource_group_name=replica_resource_group.name,
            address_spaces=["10.20.0.0/16"])
        aadds_replica_subnet = azure.network.Subnet("aaddsReplicaSubnet",
            resource_group_name=replica_resource_group.name,
            virtual_network_name=replica_virtual_network.name,
            address_prefixes=["10.20.0.0/24"])
        aadds_replica_network_security_group = azure.network.NetworkSecurityGroup("aaddsReplicaNetworkSecurityGroup",
            location=replica_resource_group.location,
            resource_group_name=replica_resource_group.name,
            security_rules=[
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowSyncWithAzureAD",
                    priority=101,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="443",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowRD",
                    priority=201,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="3389",
                    source_address_prefix="CorpNetSaw",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowPSRemoting",
                    priority=301,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="5986",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowLDAPS",
                    priority=401,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="636",
                    source_address_prefix="*",
                    destination_address_prefix="*",
                ),
            ])
        replica_subnet_network_security_group_association = azure.network.SubnetNetworkSecurityGroupAssociation("replicaSubnetNetworkSecurityGroupAssociation",
            subnet_id=aadds_replica_subnet.id,
            network_security_group_id=aadds_replica_network_security_group.id)
        primary_replica = azure.network.VirtualNetworkPeering("primaryReplica",
            resource_group_name=primary_virtual_network.resource_group_name,
            virtual_network_name=primary_virtual_network.name,
            remote_virtual_network_id=replica_virtual_network.id,
            allow_forwarded_traffic=True,
            allow_gateway_transit=False,
            allow_virtual_network_access=True,
            use_remote_gateways=False)
        replica_primary = azure.network.VirtualNetworkPeering("replicaPrimary",
            resource_group_name=replica_virtual_network.resource_group_name,
            virtual_network_name=replica_virtual_network.name,
            remote_virtual_network_id=primary_virtual_network.id,
            allow_forwarded_traffic=True,
            allow_gateway_transit=False,
            allow_virtual_network_access=True,
            use_remote_gateways=False)
        replica_virtual_network_dns_servers = azure.network.VirtualNetworkDnsServers("replicaVirtualNetworkDnsServers",
            virtual_network_id=replica_virtual_network.id,
            dns_servers=example_service.initial_replica_set.domain_controller_ip_addresses)
        replica_replica_set = azure.domainservices.ReplicaSet("replicaReplicaSet",
            domain_service_id=example_service.id,
            location=replica_resource_group.location,
            subnet_id=aadds_replica_subnet.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    replica_subnet_network_security_group_association,
                    primary_replica,
                    replica_primary,
                ]))
        ```

        ## Import

        Domain Service Replica Sets can be imported using the resource ID of the parent Domain Service and the Replica Set ID, e.g.

        ```sh
         $ pulumi import azure:domainservices/replicaSet:ReplicaSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_service_id: The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] subnet_id: The ID of the subnet in which to place this Replica Set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicaSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Replica Set for an Active Directory Domain Service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        primary_resource_group = azure.core.ResourceGroup("primaryResourceGroup", location="West Europe")
        primary_virtual_network = azure.network.VirtualNetwork("primaryVirtualNetwork",
            location=primary_resource_group.location,
            resource_group_name=primary_resource_group.name,
            address_spaces=["10.0.1.0/16"])
        primary_subnet = azure.network.Subnet("primarySubnet",
            resource_group_name=primary_resource_group.name,
            virtual_network_name=primary_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        primary_network_security_group = azure.network.NetworkSecurityGroup("primaryNetworkSecurityGroup",
            location=primary_resource_group.location,
            resource_group_name=primary_resource_group.name,
            security_rules=[
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowSyncWithAzureAD",
                    priority=101,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="443",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowRD",
                    priority=201,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="3389",
                    source_address_prefix="CorpNetSaw",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowPSRemoting",
                    priority=301,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="5986",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowLDAPS",
                    priority=401,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="636",
                    source_address_prefix="*",
                    destination_address_prefix="*",
                ),
            ])
        primary_subnet_network_security_group_association = azure.network.SubnetNetworkSecurityGroupAssociation("primarySubnetNetworkSecurityGroupAssociation",
            subnet_id=primary_subnet.id,
            network_security_group_id=primary_network_security_group.id)
        dc_admins = azuread.Group("dcAdmins")
        admin_user = azuread.User("adminUser",
            user_principal_name="dc-admin@$hashicorp-example.net",
            display_name="DC Administrator",
            password="Pa55w0Rd!!1")
        admin_group_member = azuread.GroupMember("adminGroupMember",
            group_object_id=dc_admins.object_id,
            member_object_id=admin_user.object_id)
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id="2565bd9d-da50-47d4-8b85-4c97f669dc36")
        # published app for domain services
        aadds = azure.core.ResourceGroup("aadds", location="westeurope")
        example_service = azure.domainservices.Service("exampleService",
            location=aadds.location,
            resource_group_name=aadds.name,
            domain_name="widgetslogin.net",
            sku="Enterprise",
            filtered_sync_enabled=False,
            initial_replica_set=azure.domainservices.ServiceInitialReplicaSetArgs(
                location=primary_virtual_network.location,
                subnet_id=primary_subnet.id,
            ),
            notifications=azure.domainservices.ServiceNotificationsArgs(
                additional_recipients=[
                    "notifyA@example.net",
                    "notifyB@example.org",
                ],
                notify_dc_admins=True,
                notify_global_admins=True,
            ),
            security=azure.domainservices.ServiceSecurityArgs(
                sync_kerberos_passwords=True,
                sync_ntlm_passwords=True,
                sync_on_prem_passwords=True,
            ),
            tags={
                "Environment": "prod",
            },
            opts=pulumi.ResourceOptions(depends_on=[
                    example_service_principal,
                    primary_subnet_network_security_group_association,
                ]))
        replica_resource_group = azure.core.ResourceGroup("replicaResourceGroup", location="North Europe")
        replica_virtual_network = azure.network.VirtualNetwork("replicaVirtualNetwork",
            location=replica_resource_group.location,
            resource_group_name=replica_resource_group.name,
            address_spaces=["10.20.0.0/16"])
        aadds_replica_subnet = azure.network.Subnet("aaddsReplicaSubnet",
            resource_group_name=replica_resource_group.name,
            virtual_network_name=replica_virtual_network.name,
            address_prefixes=["10.20.0.0/24"])
        aadds_replica_network_security_group = azure.network.NetworkSecurityGroup("aaddsReplicaNetworkSecurityGroup",
            location=replica_resource_group.location,
            resource_group_name=replica_resource_group.name,
            security_rules=[
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowSyncWithAzureAD",
                    priority=101,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="443",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowRD",
                    priority=201,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="3389",
                    source_address_prefix="CorpNetSaw",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowPSRemoting",
                    priority=301,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="5986",
                    source_address_prefix="AzureActiveDirectoryDomainServices",
                    destination_address_prefix="*",
                ),
                azure.network.NetworkSecurityGroupSecurityRuleArgs(
                    name="AllowLDAPS",
                    priority=401,
                    direction="Inbound",
                    access="Allow",
                    protocol="Tcp",
                    source_port_range="*",
                    destination_port_range="636",
                    source_address_prefix="*",
                    destination_address_prefix="*",
                ),
            ])
        replica_subnet_network_security_group_association = azure.network.SubnetNetworkSecurityGroupAssociation("replicaSubnetNetworkSecurityGroupAssociation",
            subnet_id=aadds_replica_subnet.id,
            network_security_group_id=aadds_replica_network_security_group.id)
        primary_replica = azure.network.VirtualNetworkPeering("primaryReplica",
            resource_group_name=primary_virtual_network.resource_group_name,
            virtual_network_name=primary_virtual_network.name,
            remote_virtual_network_id=replica_virtual_network.id,
            allow_forwarded_traffic=True,
            allow_gateway_transit=False,
            allow_virtual_network_access=True,
            use_remote_gateways=False)
        replica_primary = azure.network.VirtualNetworkPeering("replicaPrimary",
            resource_group_name=replica_virtual_network.resource_group_name,
            virtual_network_name=replica_virtual_network.name,
            remote_virtual_network_id=primary_virtual_network.id,
            allow_forwarded_traffic=True,
            allow_gateway_transit=False,
            allow_virtual_network_access=True,
            use_remote_gateways=False)
        replica_virtual_network_dns_servers = azure.network.VirtualNetworkDnsServers("replicaVirtualNetworkDnsServers",
            virtual_network_id=replica_virtual_network.id,
            dns_servers=example_service.initial_replica_set.domain_controller_ip_addresses)
        replica_replica_set = azure.domainservices.ReplicaSet("replicaReplicaSet",
            domain_service_id=example_service.id,
            location=replica_resource_group.location,
            subnet_id=aadds_replica_subnet.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    replica_subnet_network_security_group_association,
                    primary_replica,
                    replica_primary,
                ]))
        ```

        ## Import

        Domain Service Replica Sets can be imported using the resource ID of the parent Domain Service and the Replica Set ID, e.g.

        ```sh
         $ pulumi import azure:domainservices/replicaSet:ReplicaSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ReplicaSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicaSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_service_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReplicaSetArgs.__new__(ReplicaSetArgs)

            if domain_service_id is None and not opts.urn:
                raise TypeError("Missing required property 'domain_service_id'")
            __props__.__dict__["domain_service_id"] = domain_service_id
            __props__.__dict__["location"] = location
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["domain_controller_ip_addresses"] = None
            __props__.__dict__["external_access_ip_address"] = None
            __props__.__dict__["service_status"] = None
        super(ReplicaSet, __self__).__init__(
            'azure:domainservices/replicaSet:ReplicaSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_controller_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            domain_service_id: Optional[pulumi.Input[str]] = None,
            external_access_ip_address: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            service_status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'ReplicaSet':
        """
        Get an existing ReplicaSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_controller_ip_addresses: A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        :param pulumi.Input[str] domain_service_id: The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] external_access_ip_address: The publicly routable IP address for the domain controllers in this Replica Set.
        :param pulumi.Input[str] location: The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_status: The current service status for the replica set.
        :param pulumi.Input[str] subnet_id: The ID of the subnet in which to place this Replica Set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicaSetState.__new__(_ReplicaSetState)

        __props__.__dict__["domain_controller_ip_addresses"] = domain_controller_ip_addresses
        __props__.__dict__["domain_service_id"] = domain_service_id
        __props__.__dict__["external_access_ip_address"] = external_access_ip_address
        __props__.__dict__["location"] = location
        __props__.__dict__["service_status"] = service_status
        __props__.__dict__["subnet_id"] = subnet_id
        return ReplicaSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainControllerIpAddresses")
    def domain_controller_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
        """
        return pulumi.get(self, "domain_controller_ip_addresses")

    @property
    @pulumi.getter(name="domainServiceId")
    def domain_service_id(self) -> pulumi.Output[str]:
        """
        The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "domain_service_id")

    @property
    @pulumi.getter(name="externalAccessIpAddress")
    def external_access_ip_address(self) -> pulumi.Output[str]:
        """
        The publicly routable IP address for the domain controllers in this Replica Set.
        """
        return pulumi.get(self, "external_access_ip_address")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> pulumi.Output[str]:
        """
        The current service status for the replica set.
        """
        return pulumi.get(self, "service_status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet in which to place this Replica Set.
        """
        return pulumi.get(self, "subnet_id")

