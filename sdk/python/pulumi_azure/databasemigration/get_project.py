# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, id=None, location=None, name=None, resource_group_name=None, service_name=None, source_platform=None, tags=None, target_platform=None):
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
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if source_platform and not isinstance(source_platform, str):
            raise TypeError("Expected argument 'source_platform' to be a str")
        pulumi.set(__self__, "source_platform", source_platform)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_platform and not isinstance(target_platform, str):
            raise TypeError("Expected argument 'target_platform' to be a str")
        pulumi.set(__self__, "target_platform", target_platform)

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
        Azure location where the resource exists.
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
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="sourcePlatform")
    def source_platform(self) -> str:
        """
        The platform type of the migration source.
        """
        return pulumi.get(self, "source_platform")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A mapping of tags to assigned to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetPlatform")
    def target_platform(self) -> str:
        """
        The platform type of the migration target.
        """
        return pulumi.get(self, "target_platform")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            id=self.id,
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            service_name=self.service_name,
            source_platform=self.source_platform,
            tags=self.tags,
            target_platform=self.target_platform)


def get_project(name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                service_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Use this data source to access information about an existing Database Migration Project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.databasemigration.get_project(name="example-dbms-project",
        resource_group_name="example-rg",
        service_name="example-dbms")
    pulumi.export("name", example.name)
    ```


    :param str name: Name of the database migration project.
    :param str resource_group_name: Name of the resource group where resource belongs to.
    :param str service_name: Name of the database migration service where resource belongs to.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:databasemigration/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        resource_group_name=__ret__.resource_group_name,
        service_name=__ret__.service_name,
        source_platform=__ret__.source_platform,
        tags=__ret__.tags,
        target_platform=__ret__.target_platform)


@_utilities.lift_output_func(get_project)
def get_project_output(name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[str]] = None,
                       service_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Use this data source to access information about an existing Database Migration Project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.databasemigration.get_project(name="example-dbms-project",
        resource_group_name="example-rg",
        service_name="example-dbms")
    pulumi.export("name", example.name)
    ```


    :param str name: Name of the database migration project.
    :param str resource_group_name: Name of the resource group where resource belongs to.
    :param str service_name: Name of the database migration service where resource belongs to.
    """
    ...
