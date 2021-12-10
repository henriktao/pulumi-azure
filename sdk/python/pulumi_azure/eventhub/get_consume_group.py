# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetConsumeGroupResult',
    'AwaitableGetConsumeGroupResult',
    'get_consume_group',
    'get_consume_group_output',
]

@pulumi.output_type
class GetConsumeGroupResult:
    """
    A collection of values returned by getConsumeGroup.
    """
    def __init__(__self__, eventhub_name=None, id=None, location=None, name=None, namespace_name=None, resource_group_name=None, user_metadata=None):
        if eventhub_name and not isinstance(eventhub_name, str):
            raise TypeError("Expected argument 'eventhub_name' to be a str")
        pulumi.set(__self__, "eventhub_name", eventhub_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        pulumi.set(__self__, "namespace_name", namespace_name)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if user_metadata and not isinstance(user_metadata, str):
            raise TypeError("Expected argument 'user_metadata' to be a str")
        pulumi.set(__self__, "user_metadata", user_metadata)

    @property
    @pulumi.getter(name="eventhubName")
    def eventhub_name(self) -> str:
        return pulumi.get(self, "eventhub_name")

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
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> str:
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="userMetadata")
    def user_metadata(self) -> str:
        """
        Specifies the user metadata.
        """
        return pulumi.get(self, "user_metadata")


class AwaitableGetConsumeGroupResult(GetConsumeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsumeGroupResult(
            eventhub_name=self.eventhub_name,
            id=self.id,
            location=self.location,
            name=self.name,
            namespace_name=self.namespace_name,
            resource_group_name=self.resource_group_name,
            user_metadata=self.user_metadata)


def get_consume_group(eventhub_name: Optional[str] = None,
                      name: Optional[str] = None,
                      namespace_name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsumeGroupResult:
    """
    Use this data source to access information about an existing Event Hubs Consumer Group within an Event Hub.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    test = azure.eventhub.get_consume_group(name=azurerm_eventhub_consumer_group["test"]["name"],
        namespace_name=azurerm_eventhub_namespace["test"]["name"],
        eventhub_name=azurerm_eventhub["test"]["name"],
        resource_group_name=azurerm_resource_group["test"]["name"])
    ```


    :param str eventhub_name: Specifies the name of the EventHub.
    :param str name: Specifies the name of the EventHub Consumer Group resource.
    :param str namespace_name: Specifies the name of the grandparent EventHub Namespace.
    :param str resource_group_name: The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists.
    """
    __args__ = dict()
    __args__['eventhubName'] = eventhub_name
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:eventhub/getConsumeGroup:getConsumeGroup', __args__, opts=opts, typ=GetConsumeGroupResult).value

    return AwaitableGetConsumeGroupResult(
        eventhub_name=__ret__.eventhub_name,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        namespace_name=__ret__.namespace_name,
        resource_group_name=__ret__.resource_group_name,
        user_metadata=__ret__.user_metadata)


@_utilities.lift_output_func(get_consume_group)
def get_consume_group_output(eventhub_name: Optional[pulumi.Input[str]] = None,
                             name: Optional[pulumi.Input[str]] = None,
                             namespace_name: Optional[pulumi.Input[str]] = None,
                             resource_group_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConsumeGroupResult]:
    """
    Use this data source to access information about an existing Event Hubs Consumer Group within an Event Hub.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    test = azure.eventhub.get_consume_group(name=azurerm_eventhub_consumer_group["test"]["name"],
        namespace_name=azurerm_eventhub_namespace["test"]["name"],
        eventhub_name=azurerm_eventhub["test"]["name"],
        resource_group_name=azurerm_resource_group["test"]["name"])
    ```


    :param str eventhub_name: Specifies the name of the EventHub.
    :param str name: Specifies the name of the EventHub Consumer Group resource.
    :param str namespace_name: Specifies the name of the grandparent EventHub Namespace.
    :param str resource_group_name: The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists.
    """
    ...
