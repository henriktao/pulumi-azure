# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CustomProviderAction',
    'CustomProviderResourceType',
    'CustomProviderValidation',
    'ResourceGroupCostManagementExportExportDataOptions',
    'ResourceGroupCostManagementExportExportDataStorageLocation',
    'ResourceGroupPolicyAssignmentIdentity',
    'ResourceGroupPolicyAssignmentNonComplianceMessage',
    'ResourcePolicyAssignmentIdentity',
    'ResourcePolicyAssignmentNonComplianceMessage',
    'ResourceProviderRegistrationFeature',
    'SubscriptionCostManagementExportExportDataOptions',
    'SubscriptionCostManagementExportExportDataStorageLocation',
    'SubscriptionPolicyAssignmentIdentity',
    'SubscriptionPolicyAssignmentNonComplianceMessage',
    'GetResourcesResourceResult',
    'GetSubscriptionsSubscriptionResult',
]

@pulumi.output_type
class CustomProviderAction(dict):
    def __init__(__self__, *,
                 endpoint: str,
                 name: str):
        """
        :param str endpoint: Specifies the endpoint of the action.
        :param str name: Specifies the name of the action.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        Specifies the endpoint of the action.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the action.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class CustomProviderResourceType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "routingType":
            suggest = "routing_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomProviderResourceType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomProviderResourceType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomProviderResourceType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 endpoint: str,
                 name: str,
                 routing_type: Optional[str] = None):
        """
        :param str endpoint: Specifies the endpoint of the route definition.
        :param str name: Specifies the name of the route definition.
        :param str routing_type: The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "name", name)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        Specifies the endpoint of the route definition.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the route definition.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[str]:
        """
        The routing type that is supported for the resource request. Valid values are `ResourceTypeRoutingProxy` or `ResourceTypeRoutingProxyCache`. This value defaults to `ResourceTypeRoutingProxy`.
        """
        return pulumi.get(self, "routing_type")


@pulumi.output_type
class CustomProviderValidation(dict):
    def __init__(__self__, *,
                 specification: str):
        """
        :param str specification: The endpoint where the validation specification is located.
        """
        pulumi.set(__self__, "specification", specification)

    @property
    @pulumi.getter
    def specification(self) -> str:
        """
        The endpoint where the validation specification is located.
        """
        return pulumi.get(self, "specification")


@pulumi.output_type
class ResourceGroupCostManagementExportExportDataOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "timeFrame":
            suggest = "time_frame"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceGroupCostManagementExportExportDataOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceGroupCostManagementExportExportDataOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceGroupCostManagementExportExportDataOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 time_frame: str,
                 type: str):
        """
        :param str time_frame: The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        :param str type: The type of the query.
        """
        pulumi.set(__self__, "time_frame", time_frame)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="timeFrame")
    def time_frame(self) -> str:
        """
        The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        """
        return pulumi.get(self, "time_frame")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the query.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ResourceGroupCostManagementExportExportDataStorageLocation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "containerId":
            suggest = "container_id"
        elif key == "rootFolderPath":
            suggest = "root_folder_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceGroupCostManagementExportExportDataStorageLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceGroupCostManagementExportExportDataStorageLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceGroupCostManagementExportExportDataStorageLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 container_id: str,
                 root_folder_path: str):
        """
        :param str container_id: The Resource Manager ID of the container where exports will be uploaded.
        :param str root_folder_path: The path of the directory where exports will be uploaded.
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "root_folder_path", root_folder_path)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> str:
        """
        The Resource Manager ID of the container where exports will be uploaded.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> str:
        """
        The path of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")


@pulumi.output_type
class ResourceGroupPolicyAssignmentIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceGroupPolicyAssignmentIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceGroupPolicyAssignmentIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceGroupPolicyAssignmentIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        :param str principal_id: The Principal ID of the Policy Assignment for this Resource Group.
        :param str tenant_id: The Tenant ID of the Policy Assignment for this Resource Group.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The Principal ID of the Policy Assignment for this Resource Group.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID of the Policy Assignment for this Resource Group.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class ResourceGroupPolicyAssignmentNonComplianceMessage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyDefinitionReferenceId":
            suggest = "policy_definition_reference_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceGroupPolicyAssignmentNonComplianceMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceGroupPolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceGroupPolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content: str,
                 policy_definition_reference_id: Optional[str] = None):
        """
        :param str content: The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        :param str policy_definition_reference_id: When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        pulumi.set(__self__, "content", content)
        if policy_definition_reference_id is not None:
            pulumi.set(__self__, "policy_definition_reference_id", policy_definition_reference_id)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[str]:
        """
        When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        return pulumi.get(self, "policy_definition_reference_id")


@pulumi.output_type
class ResourcePolicyAssignmentIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourcePolicyAssignmentIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourcePolicyAssignmentIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourcePolicyAssignmentIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        :param str principal_id: The Principal ID of the Policy Assignment for this Resource.
        :param str tenant_id: The Tenant ID of the Policy Assignment for this Resource.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The Principal ID of the Policy Assignment for this Resource.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID of the Policy Assignment for this Resource.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class ResourcePolicyAssignmentNonComplianceMessage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyDefinitionReferenceId":
            suggest = "policy_definition_reference_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourcePolicyAssignmentNonComplianceMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourcePolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourcePolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content: str,
                 policy_definition_reference_id: Optional[str] = None):
        """
        :param str content: The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        :param str policy_definition_reference_id: When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        pulumi.set(__self__, "content", content)
        if policy_definition_reference_id is not None:
            pulumi.set(__self__, "policy_definition_reference_id", policy_definition_reference_id)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[str]:
        """
        When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        return pulumi.get(self, "policy_definition_reference_id")


@pulumi.output_type
class ResourceProviderRegistrationFeature(dict):
    def __init__(__self__, *,
                 name: str,
                 registered: bool):
        """
        :param str name: Specifies the name of the feature to register.
        :param bool registered: Should this feature be Registered or Unregistered?
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "registered", registered)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the feature to register.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def registered(self) -> bool:
        """
        Should this feature be Registered or Unregistered?
        """
        return pulumi.get(self, "registered")


@pulumi.output_type
class SubscriptionCostManagementExportExportDataOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "timeFrame":
            suggest = "time_frame"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionCostManagementExportExportDataOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionCostManagementExportExportDataOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionCostManagementExportExportDataOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 time_frame: str,
                 type: str):
        """
        :param str time_frame: The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        :param str type: The type of the query.
        """
        pulumi.set(__self__, "time_frame", time_frame)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="timeFrame")
    def time_frame(self) -> str:
        """
        The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: `WeekToDate`, `MonthToDate`, `BillingMonthToDate`, `TheLastWeek`, `TheLastMonth`, `TheLastBillingMonth`, `Custom`.
        """
        return pulumi.get(self, "time_frame")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the query.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class SubscriptionCostManagementExportExportDataStorageLocation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "containerId":
            suggest = "container_id"
        elif key == "rootFolderPath":
            suggest = "root_folder_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionCostManagementExportExportDataStorageLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionCostManagementExportExportDataStorageLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionCostManagementExportExportDataStorageLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 container_id: str,
                 root_folder_path: str):
        """
        :param str container_id: The Resource Manager ID of the container where exports will be uploaded.
        :param str root_folder_path: The path of the directory where exports will be uploaded.
        """
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "root_folder_path", root_folder_path)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> str:
        """
        The Resource Manager ID of the container where exports will be uploaded.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> str:
        """
        The path of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")


@pulumi.output_type
class SubscriptionPolicyAssignmentIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionPolicyAssignmentIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionPolicyAssignmentIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionPolicyAssignmentIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        :param str principal_id: The Principal ID of the Policy Assignment for this Subscription.
        :param str tenant_id: The Tenant ID of the Policy Assignment for this Subscription.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The Principal ID of the Policy Assignment for this Subscription.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID of the Policy Assignment for this Subscription.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class SubscriptionPolicyAssignmentNonComplianceMessage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyDefinitionReferenceId":
            suggest = "policy_definition_reference_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionPolicyAssignmentNonComplianceMessage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionPolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionPolicyAssignmentNonComplianceMessage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content: str,
                 policy_definition_reference_id: Optional[str] = None):
        """
        :param str content: The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        :param str policy_definition_reference_id: When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        pulumi.set(__self__, "content", content)
        if policy_definition_reference_id is not None:
            pulumi.set(__self__, "policy_definition_reference_id", policy_definition_reference_id)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[str]:
        """
        When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        """
        return pulumi.get(self, "policy_definition_reference_id")


@pulumi.output_type
class GetResourcesResourceResult(dict):
    def __init__(__self__, *,
                 id: str,
                 location: str,
                 name: str,
                 tags: Mapping[str, str],
                 type: str):
        """
        :param str id: The ID of this Resource.
        :param str location: The Azure Region in which this Resource exists.
        :param str name: The name of the Resource.
        :param Mapping[str, str] tags: A map of tags assigned to this Resource.
        :param str type: The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this Resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The Azure Region in which this Resource exists.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to this Resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetSubscriptionsSubscriptionResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 location_placement_id: str,
                 quota_id: str,
                 spending_limit: str,
                 state: str,
                 subscription_id: str,
                 tags: Mapping[str, str],
                 tenant_id: str):
        """
        :param str display_name: The subscription display name.
        :param str id: The ID of this subscription.
        :param str location_placement_id: The subscription location placement ID.
        :param str quota_id: The subscription quota ID.
        :param str spending_limit: The subscription spending limit.
        :param str state: The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
        :param str subscription_id: The subscription GUID.
        :param Mapping[str, str] tags: A mapping of tags assigned to the resource.
        :param str tenant_id: The subscription tenant ID.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "location_placement_id", location_placement_id)
        pulumi.set(__self__, "quota_id", quota_id)
        pulumi.set(__self__, "spending_limit", spending_limit)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "subscription_id", subscription_id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The subscription display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this subscription.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="locationPlacementId")
    def location_placement_id(self) -> str:
        """
        The subscription location placement ID.
        """
        return pulumi.get(self, "location_placement_id")

    @property
    @pulumi.getter(name="quotaId")
    def quota_id(self) -> str:
        """
        The subscription quota ID.
        """
        return pulumi.get(self, "quota_id")

    @property
    @pulumi.getter(name="spendingLimit")
    def spending_limit(self) -> str:
        """
        The subscription spending limit.
        """
        return pulumi.get(self, "spending_limit")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> str:
        """
        The subscription GUID.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The subscription tenant ID.
        """
        return pulumi.get(self, "tenant_id")


