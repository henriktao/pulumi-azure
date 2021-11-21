# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CustomProviderActionArgs',
    'CustomProviderResourceTypeArgs',
    'CustomProviderValidationArgs',
    'ResourceGroupCostManagementExportExportDataOptionsArgs',
    'ResourceGroupCostManagementExportExportDataStorageLocationArgs',
    'ResourceGroupPolicyAssignmentIdentityArgs',
    'ResourcePolicyAssignmentIdentityArgs',
    'ResourceProviderRegistrationFeatureArgs',
    'SubscriptionCostManagementExportExportDataOptionsArgs',
    'SubscriptionCostManagementExportExportDataStorageLocationArgs',
    'SubscriptionPolicyAssignmentIdentityArgs',
]

@pulumi.input_type
class CustomProviderActionArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] endpoint: Specifies the endpoint of the action.
        :param pulumi.Input[str] name: Specifies the name of the action.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        Specifies the endpoint of the action.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the action.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class CustomProviderResourceTypeArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 name: pulumi.Input[str],
                 routing_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] endpoint: Specifies the endpoint of the route definition.
        :param pulumi.Input[str] name: Specifies the name of the route definition.
        :param pulumi.Input[str] routing_type: The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        Specifies the endpoint of the route definition.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the route definition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
        """
        return pulumi.get(self, "routing_type")

    @routing_type.setter
    def routing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_type", value)


@pulumi.input_type
class CustomProviderValidationArgs:
    def __init__(__self__, *,
                 specification: pulumi.Input[str]):
        """
        :param pulumi.Input[str] specification: The endpoint where the validation specification is located.
        """
        pulumi.set(__self__, "specification", specification)

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Input[str]:
        """
        The endpoint where the validation specification is located.
        """
        return pulumi.get(self, "specification")

    @specification.setter
    def specification(self, value: pulumi.Input[str]):
        pulumi.set(self, "specification", value)


@pulumi.input_type
class ResourceGroupCostManagementExportExportDataOptionsArgs:
    def __init__(__self__, *,
                 time_frame: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] time_frame: The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        :param pulumi.Input[str] type: The type of the query.
        """
        pulumi.set(__self__, "time_frame", time_frame)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="timeFrame")
    def time_frame(self) -> pulumi.Input[str]:
        """
        The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        """
        return pulumi.get(self, "time_frame")

    @time_frame.setter
    def time_frame(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_frame", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the query.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ResourceGroupCostManagementExportExportDataStorageLocationArgs:
    def __init__(__self__, *,
                 container_id: pulumi.Input[str],
                 root_folder_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] container_id: The Resource Manager ID of the container where exports will be uploaded.
        :param pulumi.Input[str] root_folder_path: The path of the directory where exports will be uploaded.
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "root_folder_path", root_folder_path)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[str]:
        """
        The Resource Manager ID of the container where exports will be uploaded.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> pulumi.Input[str]:
        """
        The path of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")

    @root_folder_path.setter
    def root_folder_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_folder_path", value)


@pulumi.input_type
class ResourceGroupPolicyAssignmentIdentityArgs:
    def __init__(__self__, *,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: The Principal ID of the Policy Assignment for this Resource Group.
        :param pulumi.Input[str] tenant_id: The Tenant ID of the Policy Assignment for this Resource Group.
        :param pulumi.Input[str] type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
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
        The Principal ID of the Policy Assignment for this Resource Group.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID of the Policy Assignment for this Resource Group.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ResourcePolicyAssignmentIdentityArgs:
    def __init__(__self__, *,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: The Principal ID of the Policy Assignment for this Resource.
        :param pulumi.Input[str] tenant_id: The Tenant ID of the Policy Assignment for this Resource.
        :param pulumi.Input[str] type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
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
        The Principal ID of the Policy Assignment for this Resource.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID of the Policy Assignment for this Resource.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ResourceProviderRegistrationFeatureArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 registered: pulumi.Input[bool]):
        """
        :param pulumi.Input[str] name: Specifies the name of the feature to register.
        :param pulumi.Input[bool] registered: Should this feature be Registered or Unregistered?
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "registered", registered)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the feature to register.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def registered(self) -> pulumi.Input[bool]:
        """
        Should this feature be Registered or Unregistered?
        """
        return pulumi.get(self, "registered")

    @registered.setter
    def registered(self, value: pulumi.Input[bool]):
        pulumi.set(self, "registered", value)


@pulumi.input_type
class SubscriptionCostManagementExportExportDataOptionsArgs:
    def __init__(__self__, *,
                 time_frame: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] time_frame: The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        :param pulumi.Input[str] type: The type of the query.
        """
        pulumi.set(__self__, "time_frame", time_frame)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="timeFrame")
    def time_frame(self) -> pulumi.Input[str]:
        """
        The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        """
        return pulumi.get(self, "time_frame")

    @time_frame.setter
    def time_frame(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_frame", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the query.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SubscriptionCostManagementExportExportDataStorageLocationArgs:
    def __init__(__self__, *,
                 container_id: pulumi.Input[str],
                 root_folder_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] container_id: The Resource Manager ID of the container where exports will be uploaded.
        :param pulumi.Input[str] root_folder_path: The path of the directory where exports will be uploaded.
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "root_folder_path", root_folder_path)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[str]:
        """
        The Resource Manager ID of the container where exports will be uploaded.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> pulumi.Input[str]:
        """
        The path of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")

    @root_folder_path.setter
    def root_folder_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_folder_path", value)


@pulumi.input_type
class SubscriptionPolicyAssignmentIdentityArgs:
    def __init__(__self__, *,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: The Principal ID of the Policy Assignment for this Subscription.
        :param pulumi.Input[str] tenant_id: The Tenant ID of the Policy Assignment for this Subscription.
        :param pulumi.Input[str] type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
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
        The Principal ID of the Policy Assignment for this Subscription.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID of the Policy Assignment for this Subscription.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


