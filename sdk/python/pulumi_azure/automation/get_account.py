# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, endpoint=None, id=None, name=None, primary_key=None, resource_group_name=None, secondary_key=None):
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        pulumi.set(__self__, "primary_key", primary_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The Endpoint for this Automation Account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The Primary Access Key for the Automation Account.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The Secondary Access Key for the Automation Account.
        """
        return pulumi.get(self, "secondary_key")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            endpoint=self.endpoint,
            id=self.id,
            name=self.name,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_key=self.secondary_key)


def get_account(name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Use this data source to access information about an existing Automation Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.automation.get_account(name="example-account",
        resource_group_name="example-resources")
    pulumi.export("automationAccountId", example.id)
    ```


    :param str name: The name of the Automation Account.
    :param str resource_group_name: Specifies the name of the Resource Group where the Automation Account exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:automation/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        endpoint=__ret__.endpoint,
        id=__ret__.id,
        name=__ret__.name,
        primary_key=__ret__.primary_key,
        resource_group_name=__ret__.resource_group_name,
        secondary_key=__ret__.secondary_key)


@_utilities.lift_output_func(get_account)
def get_account_output(name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    Use this data source to access information about an existing Automation Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.automation.get_account(name="example-account",
        resource_group_name="example-resources")
    pulumi.export("automationAccountId", example.id)
    ```


    :param str name: The name of the Automation Account.
    :param str resource_group_name: Specifies the name of the Resource Group where the Automation Account exists.
    """
    ...
