# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEnterpriseDatabaseResult',
    'AwaitableGetEnterpriseDatabaseResult',
    'get_enterprise_database',
]

@pulumi.output_type
class GetEnterpriseDatabaseResult:
    """
    A collection of values returned by getEnterpriseDatabase.
    """
    def __init__(__self__, cluster_id=None, id=None, name=None, primary_access_key=None, resource_group_name=None, secondary_access_key=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if primary_access_key and not isinstance(primary_access_key, str):
            raise TypeError("Expected argument 'primary_access_key' to be a str")
        pulumi.set(__self__, "primary_access_key", primary_access_key)
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if secondary_access_key and not isinstance(secondary_access_key, str):
            raise TypeError("Expected argument 'secondary_access_key' to be a str")
        pulumi.set(__self__, "secondary_access_key", secondary_access_key)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The Redis Enterprise Cluster ID that is hosting the Redis Enterprise Database.
        """
        return pulumi.get(self, "cluster_id")

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
        """
        The Redis Enterprise Database name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryAccessKey")
    def primary_access_key(self) -> str:
        """
        The Primary Access Key for the Redis Enterprise Database instance.
        """
        return pulumi.get(self, "primary_access_key")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryAccessKey")
    def secondary_access_key(self) -> str:
        """
        The Secondary Access Key for the Redis Enterprise Database instance.
        """
        return pulumi.get(self, "secondary_access_key")


class AwaitableGetEnterpriseDatabaseResult(GetEnterpriseDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnterpriseDatabaseResult(
            cluster_id=self.cluster_id,
            id=self.id,
            name=self.name,
            primary_access_key=self.primary_access_key,
            resource_group_name=self.resource_group_name,
            secondary_access_key=self.secondary_access_key)


def get_enterprise_database(cluster_id: Optional[str] = None,
                            name: Optional[str] = None,
                            resource_group_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnterpriseDatabaseResult:
    """
    Use this data source to access information about an existing Redis Enterprise Database

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.redis.get_enterprise_database(name="default",
        resource_group_name=azurerm_resource_group["example"]["name"],
        cluster_id=azurerm_redis_enterprise_cluster["example"]["id"])
    pulumi.export("redisEnterpriseDatabasePrimaryKey", example.primary_access_key)
    pulumi.export("redisEnterpriseDatabaseSecondaryKey", example.secondary_access_key)
    ```


    :param str cluster_id: The resource ID of Redis Enterprise Cluster which hosts the Redis Enterprise Database instance.
    :param str name: The name of the Redis Enterprise Database.
    :param str resource_group_name: The name of the resource group the Redis Enterprise Database instance is located in.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:redis/getEnterpriseDatabase:getEnterpriseDatabase', __args__, opts=opts, typ=GetEnterpriseDatabaseResult).value

    return AwaitableGetEnterpriseDatabaseResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        name=__ret__.name,
        primary_access_key=__ret__.primary_access_key,
        resource_group_name=__ret__.resource_group_name,
        secondary_access_key=__ret__.secondary_access_key)
