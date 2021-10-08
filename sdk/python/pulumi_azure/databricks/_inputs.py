# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'WorkspaceCustomParametersArgs',
    'WorkspaceStorageAccountIdentityArgs',
]

@pulumi.input_type
class WorkspaceCustomParametersArgs:
    def __init__(__self__, *,
                 machine_learning_workspace_id: Optional[pulumi.Input[str]] = None,
                 nat_gateway_name: Optional[pulumi.Input[str]] = None,
                 no_public_ip: Optional[pulumi.Input[bool]] = None,
                 private_subnet_name: Optional[pulumi.Input[str]] = None,
                 private_subnet_network_security_group_association_id: Optional[pulumi.Input[str]] = None,
                 public_ip_name: Optional[pulumi.Input[str]] = None,
                 public_subnet_name: Optional[pulumi.Input[str]] = None,
                 public_subnet_network_security_group_association_id: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_account_sku_name: Optional[pulumi.Input[str]] = None,
                 virtual_network_id: Optional[pulumi.Input[str]] = None,
                 vnet_address_prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] nat_gateway_name: Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets. Defaults to `nat-gateway`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] no_public_ip: Are public IP Addresses not allowed? Possible values are `true` or `false`. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_subnet_name: The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_subnet_network_security_group_association_id: The resource ID of the `network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `private_subnet_name` field. Required if `virtual_network_id` is set.
        :param pulumi.Input[str] public_ip_name: Name of the Public IP for No Public IP workspace with managed vNet. Defaults to `nat-gw-public-ip`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] public_subnet_name: The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        :param pulumi.Input[str] public_subnet_network_security_group_association_id: The resource ID of the `network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `public_subnet_name` field. Required if `virtual_network_id` is set.
        :param pulumi.Input[str] storage_account_name: Default Databricks File Storage account name. Defaults to a randomized name(e.g. `dbstoragel6mfeghoe5kxu`). Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_sku_name: Storage account SKU name. Possible values include `Standard_LRS`, `Standard_GRS`, `Standard_RAGRS`, `Standard_GZRS`, `Standard_RAGZRS`, `Standard_ZRS`, `Premium_LRS` or `Premium_ZRS`. Defaults to `Standard_GRS`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_network_id: The ID of a Virtual Network where this Databricks Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] vnet_address_prefix: Address prefix for Managed virtual network. Defaults to `10.139`. Changing this forces a new resource to be created.
        """
        if machine_learning_workspace_id is not None:
            pulumi.set(__self__, "machine_learning_workspace_id", machine_learning_workspace_id)
        if nat_gateway_name is not None:
            pulumi.set(__self__, "nat_gateway_name", nat_gateway_name)
        if no_public_ip is not None:
            pulumi.set(__self__, "no_public_ip", no_public_ip)
        if private_subnet_name is not None:
            pulumi.set(__self__, "private_subnet_name", private_subnet_name)
        if private_subnet_network_security_group_association_id is not None:
            pulumi.set(__self__, "private_subnet_network_security_group_association_id", private_subnet_network_security_group_association_id)
        if public_ip_name is not None:
            pulumi.set(__self__, "public_ip_name", public_ip_name)
        if public_subnet_name is not None:
            pulumi.set(__self__, "public_subnet_name", public_subnet_name)
        if public_subnet_network_security_group_association_id is not None:
            pulumi.set(__self__, "public_subnet_network_security_group_association_id", public_subnet_network_security_group_association_id)
        if storage_account_name is not None:
            pulumi.set(__self__, "storage_account_name", storage_account_name)
        if storage_account_sku_name is not None:
            pulumi.set(__self__, "storage_account_sku_name", storage_account_sku_name)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)
        if vnet_address_prefix is not None:
            pulumi.set(__self__, "vnet_address_prefix", vnet_address_prefix)

    @property
    @pulumi.getter(name="machineLearningWorkspaceId")
    def machine_learning_workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "machine_learning_workspace_id")

    @machine_learning_workspace_id.setter
    def machine_learning_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_learning_workspace_id", value)

    @property
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets. Defaults to `nat-gateway`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "nat_gateway_name")

    @nat_gateway_name.setter
    def nat_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_name", value)

    @property
    @pulumi.getter(name="noPublicIp")
    def no_public_ip(self) -> Optional[pulumi.Input[bool]]:
        """
        Are public IP Addresses not allowed? Possible values are `true` or `false`. Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "no_public_ip")

    @no_public_ip.setter
    def no_public_ip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_public_ip", value)

    @property
    @pulumi.getter(name="privateSubnetName")
    def private_subnet_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "private_subnet_name")

    @private_subnet_name.setter
    def private_subnet_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_subnet_name", value)

    @property
    @pulumi.getter(name="privateSubnetNetworkSecurityGroupAssociationId")
    def private_subnet_network_security_group_association_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the `network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `private_subnet_name` field. Required if `virtual_network_id` is set.
        """
        return pulumi.get(self, "private_subnet_network_security_group_association_id")

    @private_subnet_network_security_group_association_id.setter
    def private_subnet_network_security_group_association_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_subnet_network_security_group_association_id", value)

    @property
    @pulumi.getter(name="publicIpName")
    def public_ip_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Public IP for No Public IP workspace with managed vNet. Defaults to `nat-gw-public-ip`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "public_ip_name")

    @public_ip_name.setter
    def public_ip_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_ip_name", value)

    @property
    @pulumi.getter(name="publicSubnetName")
    def public_subnet_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "public_subnet_name")

    @public_subnet_name.setter
    def public_subnet_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_subnet_name", value)

    @property
    @pulumi.getter(name="publicSubnetNetworkSecurityGroupAssociationId")
    def public_subnet_network_security_group_association_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the `network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `public_subnet_name` field. Required if `virtual_network_id` is set.
        """
        return pulumi.get(self, "public_subnet_network_security_group_association_id")

    @public_subnet_network_security_group_association_id.setter
    def public_subnet_network_security_group_association_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_subnet_network_security_group_association_id", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Default Databricks File Storage account name. Defaults to a randomized name(e.g. `dbstoragel6mfeghoe5kxu`). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter(name="storageAccountSkuName")
    def storage_account_sku_name(self) -> Optional[pulumi.Input[str]]:
        """
        Storage account SKU name. Possible values include `Standard_LRS`, `Standard_GRS`, `Standard_RAGRS`, `Standard_GZRS`, `Standard_RAGZRS`, `Standard_ZRS`, `Premium_LRS` or `Premium_ZRS`. Defaults to `Standard_GRS`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_account_sku_name")

    @storage_account_sku_name.setter
    def storage_account_sku_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_sku_name", value)

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a Virtual Network where this Databricks Cluster should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_network_id", value)

    @property
    @pulumi.getter(name="vnetAddressPrefix")
    def vnet_address_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Address prefix for Managed virtual network. Defaults to `10.139`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "vnet_address_prefix")

    @vnet_address_prefix.setter
    def vnet_address_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vnet_address_prefix", value)


@pulumi.input_type
class WorkspaceStorageAccountIdentityArgs:
    def __init__(__self__, *,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: The principal UUID for the internal databricks storage account needed to provide access to the workspace for enabling Customer Managed Keys.
        :param pulumi.Input[str] tenant_id: The UUID of the tenant where the internal databricks storage account was created.
        :param pulumi.Input[str] type: The type of the internal databricks storage account.
        """
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The principal UUID for the internal databricks storage account needed to provide access to the workspace for enabling Customer Managed Keys.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the tenant where the internal databricks storage account was created.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the internal databricks storage account.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


