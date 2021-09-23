# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDatasetKustoDatabaseResult',
    'AwaitableGetDatasetKustoDatabaseResult',
    'get_dataset_kusto_database',
    'get_dataset_kusto_database_output',
]

@pulumi.output_type
class GetDatasetKustoDatabaseResult:
    """
    A collection of values returned by getDatasetKustoDatabase.
    """
    def __init__(__self__, display_name=None, id=None, kusto_cluster_location=None, kusto_database_id=None, name=None, share_id=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kusto_cluster_location and not isinstance(kusto_cluster_location, str):
            raise TypeError("Expected argument 'kusto_cluster_location' to be a str")
        pulumi.set(__self__, "kusto_cluster_location", kusto_cluster_location)
        if kusto_database_id and not isinstance(kusto_database_id, str):
            raise TypeError("Expected argument 'kusto_database_id' to be a str")
        pulumi.set(__self__, "kusto_database_id", kusto_database_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if share_id and not isinstance(share_id, str):
            raise TypeError("Expected argument 'share_id' to be a str")
        pulumi.set(__self__, "share_id", share_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of the Data Share Dataset.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kustoClusterLocation")
    def kusto_cluster_location(self) -> str:
        """
        The location of the Kusto Cluster.
        """
        return pulumi.get(self, "kusto_cluster_location")

    @property
    @pulumi.getter(name="kustoDatabaseId")
    def kusto_database_id(self) -> str:
        """
        The resource ID of the Kusto Cluster Database to be shared with the receiver.
        """
        return pulumi.get(self, "kusto_database_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> str:
        return pulumi.get(self, "share_id")


class AwaitableGetDatasetKustoDatabaseResult(GetDatasetKustoDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetKustoDatabaseResult(
            display_name=self.display_name,
            id=self.id,
            kusto_cluster_location=self.kusto_cluster_location,
            kusto_database_id=self.kusto_database_id,
            name=self.name,
            share_id=self.share_id)


def get_dataset_kusto_database(name: Optional[str] = None,
                               share_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetKustoDatabaseResult:
    """
    Use this data source to access information about an existing Data Share Kusto Database Dataset.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.datashare.get_dataset_kusto_database(name="example-dskdds",
        share_id="example-share-id")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this Data Share Kusto Database Dataset.
    :param str share_id: The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['shareId'] = share_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:datashare/getDatasetKustoDatabase:getDatasetKustoDatabase', __args__, opts=opts, typ=GetDatasetKustoDatabaseResult).value

    return AwaitableGetDatasetKustoDatabaseResult(
        display_name=__ret__.display_name,
        id=__ret__.id,
        kusto_cluster_location=__ret__.kusto_cluster_location,
        kusto_database_id=__ret__.kusto_database_id,
        name=__ret__.name,
        share_id=__ret__.share_id)


@_utilities.lift_output_func(get_dataset_kusto_database)
def get_dataset_kusto_database_output(name: Optional[pulumi.Input[str]] = None,
                                      share_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatasetKustoDatabaseResult]:
    """
    Use this data source to access information about an existing Data Share Kusto Database Dataset.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.datashare.get_dataset_kusto_database(name="example-dskdds",
        share_id="example-share-id")
    pulumi.export("id", example.id)
    ```


    :param str name: The name of this Data Share Kusto Database Dataset.
    :param str share_id: The resource ID of the Data Share where this Data Share Kusto Database Dataset should be created.
    """
    ...
