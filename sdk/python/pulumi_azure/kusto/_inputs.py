# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ClusterIdentityArgs',
    'ClusterOptimizedAutoScaleArgs',
    'ClusterSkuArgs',
    'ClusterVirtualNetworkConfigurationArgs',
]

@pulumi.input_type
class ClusterIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 identity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] identity_ids: A list of IDs for User Assigned Managed Identity resources to be assigned.
        :param pulumi.Input[str] principal_id: The Principal ID associated with this System Assigned Managed Service Identity.
        :param pulumi.Input[str] tenant_id: The Tenant ID associated with this System Assigned Managed Service Identity.
        """
        pulumi.set(__self__, "type", type)
        if identity_ids is not None:
            pulumi.set(__self__, "identity_ids", identity_ids)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of Managed Service Identity that is configured on this Kusto Cluster. Possible values are: `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="identityIds")
    def identity_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of IDs for User Assigned Managed Identity resources to be assigned.
        """
        return pulumi.get(self, "identity_ids")

    @identity_ids.setter
    def identity_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "identity_ids", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Principal ID associated with this System Assigned Managed Service Identity.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID associated with this System Assigned Managed Service Identity.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class ClusterOptimizedAutoScaleArgs:
    def __init__(__self__, *,
                 maximum_instances: pulumi.Input[int],
                 minimum_instances: pulumi.Input[int]):
        """
        :param pulumi.Input[int] maximum_instances: The maximum number of allowed instances. Must between `0` and `1000`.
        :param pulumi.Input[int] minimum_instances: The minimum number of allowed instances. Must between `0` and `1000`.
        """
        pulumi.set(__self__, "maximum_instances", maximum_instances)
        pulumi.set(__self__, "minimum_instances", minimum_instances)

    @property
    @pulumi.getter(name="maximumInstances")
    def maximum_instances(self) -> pulumi.Input[int]:
        """
        The maximum number of allowed instances. Must between `0` and `1000`.
        """
        return pulumi.get(self, "maximum_instances")

    @maximum_instances.setter
    def maximum_instances(self, value: pulumi.Input[int]):
        pulumi.set(self, "maximum_instances", value)

    @property
    @pulumi.getter(name="minimumInstances")
    def minimum_instances(self) -> pulumi.Input[int]:
        """
        The minimum number of allowed instances. Must between `0` and `1000`.
        """
        return pulumi.get(self, "minimum_instances")

    @minimum_instances.setter
    def minimum_instances(self, value: pulumi.Input[int]):
        pulumi.set(self, "minimum_instances", value)


@pulumi.input_type
class ClusterSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] name: The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`.
        :param pulumi.Input[int] capacity: Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the SKU. Valid values are: `Dev(No SLA)_Standard_D11_v2`, `Dev(No SLA)_Standard_E2a_v4`, `Standard_D11_v2`, `Standard_D12_v2`, `Standard_D13_v2`, `Standard_D14_v2`, `Standard_DS13_v2+1TB_PS`, `Standard_DS13_v2+2TB_PS`, `Standard_DS14_v2+3TB_PS`, `Standard_DS14_v2+4TB_PS`, `Standard_E16as_v4+3TB_PS`, `Standard_E16as_v4+4TB_PS`, `Standard_E16a_v4`, `Standard_E2a_v4`, `Standard_E4a_v4`, `Standard_E64i_v3`, `Standard_E8as_v4+1TB_PS`, `Standard_E8as_v4+2TB_PS`, `Standard_E8a_v4`, `Standard_L16s`, `Standard_L4s` and `Standard_L8s`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the node count for the cluster. Boundaries depend on the sku name.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)


@pulumi.input_type
class ClusterVirtualNetworkConfigurationArgs:
    def __init__(__self__, *,
                 data_management_public_ip_id: pulumi.Input[str],
                 engine_public_ip_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] data_management_public_ip_id: Data management's service public IP address resource id.
        :param pulumi.Input[str] engine_public_ip_id: Engine service's public IP address resource id.
        :param pulumi.Input[str] subnet_id: The subnet resource id.
        """
        pulumi.set(__self__, "data_management_public_ip_id", data_management_public_ip_id)
        pulumi.set(__self__, "engine_public_ip_id", engine_public_ip_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="dataManagementPublicIpId")
    def data_management_public_ip_id(self) -> pulumi.Input[str]:
        """
        Data management's service public IP address resource id.
        """
        return pulumi.get(self, "data_management_public_ip_id")

    @data_management_public_ip_id.setter
    def data_management_public_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_management_public_ip_id", value)

    @property
    @pulumi.getter(name="enginePublicIpId")
    def engine_public_ip_id(self) -> pulumi.Input[str]:
        """
        Engine service's public IP address resource id.
        """
        return pulumi.get(self, "engine_public_ip_id")

    @engine_public_ip_id.setter
    def engine_public_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_public_ip_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet resource id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


