# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetStoreResult',
    'AwaitableGetStoreResult',
    'get_store',
    'get_store_output',
]

@pulumi.output_type
class GetStoreResult:
    """
    A collection of values returned by getStore.
    """
    def __init__(__self__, encryption_state=None, encryption_type=None, firewall_allow_azure_ips=None, firewall_state=None, id=None, location=None, name=None, resource_group_name=None, tags=None, tier=None):
        if encryption_state and not isinstance(encryption_state, str):
            raise TypeError("Expected argument 'encryption_state' to be a str")
        pulumi.set(__self__, "encryption_state", encryption_state)
        if encryption_type and not isinstance(encryption_type, str):
            raise TypeError("Expected argument 'encryption_type' to be a str")
        pulumi.set(__self__, "encryption_type", encryption_type)
        if firewall_allow_azure_ips and not isinstance(firewall_allow_azure_ips, str):
            raise TypeError("Expected argument 'firewall_allow_azure_ips' to be a str")
        pulumi.set(__self__, "firewall_allow_azure_ips", firewall_allow_azure_ips)
        if firewall_state and not isinstance(firewall_state, str):
            raise TypeError("Expected argument 'firewall_state' to be a str")
        pulumi.set(__self__, "firewall_state", firewall_state)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter(name="encryptionState")
    def encryption_state(self) -> str:
        """
        the Encryption State of this Data Lake Store Account, such as `Enabled` or `Disabled`.
        """
        return pulumi.get(self, "encryption_state")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> str:
        """
        the Encryption Type used for this Data Lake Store Account.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="firewallAllowAzureIps")
    def firewall_allow_azure_ips(self) -> str:
        """
        are Azure Service IP's allowed through the firewall?
        """
        return pulumi.get(self, "firewall_allow_azure_ips")

    @property
    @pulumi.getter(name="firewallState")
    def firewall_state(self) -> str:
        """
        the state of the firewall, such as `Enabled` or `Disabled`.
        """
        return pulumi.get(self, "firewall_state")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assign to the Data Lake Store.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        Current monthly commitment tier for the account.
        """
        return pulumi.get(self, "tier")


class AwaitableGetStoreResult(GetStoreResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStoreResult(
            encryption_state=self.encryption_state,
            encryption_type=self.encryption_type,
            firewall_allow_azure_ips=self.firewall_allow_azure_ips,
            firewall_state=self.firewall_state,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            tier=self.tier)


def get_store(name: Optional[str] = None,
              resource_group_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStoreResult:
    """
    Use this data source to access information about an existing Data Lake Store.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.datalake.get_store(name="testdatalake",
        resource_group_name="testdatalake")
    pulumi.export("dataLakeStoreId", example.id)
    ```


    :param str name: The name of the Data Lake Store.
    :param str resource_group_name: The Name of the Resource Group where the Data Lake Store exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:datalake/getStore:getStore', __args__, opts=opts, typ=GetStoreResult).value

    return AwaitableGetStoreResult(
        encryption_state=__ret__.encryption_state,
        encryption_type=__ret__.encryption_type,
        firewall_allow_azure_ips=__ret__.firewall_allow_azure_ips,
        firewall_state=__ret__.firewall_state,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags,
        tier=__ret__.tier)


@_utilities.lift_output_func(get_store)
def get_store_output(name: Optional[pulumi.Input[str]] = None,
                     resource_group_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStoreResult]:
    """
    Use this data source to access information about an existing Data Lake Store.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.datalake.get_store(name="testdatalake",
        resource_group_name="testdatalake")
    pulumi.export("dataLakeStoreId", example.id)
    ```


    :param str name: The name of the Data Lake Store.
    :param str resource_group_name: The Name of the Resource Group where the Data Lake Store exists.
    """
    ...
