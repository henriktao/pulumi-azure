# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDedicatedHostResult',
    'AwaitableGetDedicatedHostResult',
    'get_dedicated_host',
    'get_dedicated_host_output',
]

@pulumi.output_type
class GetDedicatedHostResult:
    """
    A collection of values returned by getDedicatedHost.
    """
    def __init__(__self__, dedicated_host_group_name=None, id=None, location=None, name=None, resource_group_name=None, tags=None):
        if dedicated_host_group_name and not isinstance(dedicated_host_group_name, str):
            raise TypeError("Expected argument 'dedicated_host_group_name' to be a str")
        pulumi.set(__self__, "dedicated_host_group_name", dedicated_host_group_name)
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

    @property
    @pulumi.getter(name="dedicatedHostGroupName")
    def dedicated_host_group_name(self) -> str:
        return pulumi.get(self, "dedicated_host_group_name")

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
        """
        The location where the Dedicated Host exists.
        """
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
        A mapping of tags assigned to the Dedicated Host.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDedicatedHostResult(GetDedicatedHostResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDedicatedHostResult(
            dedicated_host_group_name=self.dedicated_host_group_name,
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags)


def get_dedicated_host(dedicated_host_group_name: Optional[str] = None,
                       name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDedicatedHostResult:
    """
    Use this data source to access information about an existing Dedicated Host.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.compute.get_dedicated_host(name="example-host",
        dedicated_host_group_name="example-host-group",
        resource_group_name="example-resources")
    pulumi.export("dedicatedHostId", example.id)
    ```


    :param str dedicated_host_group_name: Specifies the name of the Dedicated Host Group the Dedicated Host is located in.
    :param str name: Specifies the name of the Dedicated Host.
    :param str resource_group_name: Specifies the name of the resource group the Dedicated Host is located in.
    """
    __args__ = dict()
    __args__['dedicatedHostGroupName'] = dedicated_host_group_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getDedicatedHost:getDedicatedHost', __args__, opts=opts, typ=GetDedicatedHostResult).value

    return AwaitableGetDedicatedHostResult(
        dedicated_host_group_name=__ret__.dedicated_host_group_name,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_dedicated_host)
def get_dedicated_host_output(dedicated_host_group_name: Optional[pulumi.Input[str]] = None,
                              name: Optional[pulumi.Input[str]] = None,
                              resource_group_name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDedicatedHostResult]:
    """
    Use this data source to access information about an existing Dedicated Host.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.compute.get_dedicated_host(name="example-host",
        dedicated_host_group_name="example-host-group",
        resource_group_name="example-resources")
    pulumi.export("dedicatedHostId", example.id)
    ```


    :param str dedicated_host_group_name: Specifies the name of the Dedicated Host Group the Dedicated Host is located in.
    :param str name: Specifies the name of the Dedicated Host.
    :param str resource_group_name: Specifies the name of the resource group the Dedicated Host is located in.
    """
    ...
