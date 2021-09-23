# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetWorkspacePrivateEndpointConnectionResult',
    'AwaitableGetWorkspacePrivateEndpointConnectionResult',
    'get_workspace_private_endpoint_connection',
    'get_workspace_private_endpoint_connection_output',
]

@pulumi.output_type
class GetWorkspacePrivateEndpointConnectionResult:
    """
    A collection of values returned by getWorkspacePrivateEndpointConnection.
    """
    def __init__(__self__, connections=None, id=None, private_endpoint_id=None, workspace_id=None):
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if private_endpoint_id and not isinstance(private_endpoint_id, str):
            raise TypeError("Expected argument 'private_endpoint_id' to be a str")
        pulumi.set(__self__, "private_endpoint_id", private_endpoint_id)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.GetWorkspacePrivateEndpointConnectionConnectionResult']:
        """
        A `connections` block as documented below.
        """
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="privateEndpointId")
    def private_endpoint_id(self) -> str:
        """
        The resource ID of the Private Endpoint.
        """
        return pulumi.get(self, "private_endpoint_id")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        """
        The resource ID of the Databricks Workspace.
        """
        return pulumi.get(self, "workspace_id")


class AwaitableGetWorkspacePrivateEndpointConnectionResult(GetWorkspacePrivateEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspacePrivateEndpointConnectionResult(
            connections=self.connections,
            id=self.id,
            private_endpoint_id=self.private_endpoint_id,
            workspace_id=self.workspace_id)


def get_workspace_private_endpoint_connection(private_endpoint_id: Optional[str] = None,
                                              workspace_id: Optional[str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspacePrivateEndpointConnectionResult:
    """
    Use this data source to access information on an existing Databricks Workspace private endpoint connection state.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.databricks.get_workspace_private_endpoint_connection(workspace_id=azurerm_databricks_workspace["example"]["id"],
        private_endpoint_id=azurerm_private_endpoint["example"]["id"])
    pulumi.export("databricksWorkspacePrivateEndpointConnectionStatus", example.connections[0].status)
    ```


    :param str private_endpoint_id: The resource ID of the Private Endpoint.
    :param str workspace_id: The resource ID of the Databricks Workspace.
    """
    __args__ = dict()
    __args__['privateEndpointId'] = private_endpoint_id
    __args__['workspaceId'] = workspace_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:databricks/getWorkspacePrivateEndpointConnection:getWorkspacePrivateEndpointConnection', __args__, opts=opts, typ=GetWorkspacePrivateEndpointConnectionResult).value

    return AwaitableGetWorkspacePrivateEndpointConnectionResult(
        connections=__ret__.connections,
        id=__ret__.id,
        private_endpoint_id=__ret__.private_endpoint_id,
        workspace_id=__ret__.workspace_id)


@_utilities.lift_output_func(get_workspace_private_endpoint_connection)
def get_workspace_private_endpoint_connection_output(private_endpoint_id: Optional[pulumi.Input[str]] = None,
                                                     workspace_id: Optional[pulumi.Input[str]] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkspacePrivateEndpointConnectionResult]:
    """
    Use this data source to access information on an existing Databricks Workspace private endpoint connection state.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azure as azure

    example = azure.databricks.get_workspace_private_endpoint_connection(workspace_id=azurerm_databricks_workspace["example"]["id"],
        private_endpoint_id=azurerm_private_endpoint["example"]["id"])
    pulumi.export("databricksWorkspacePrivateEndpointConnectionStatus", example.connections[0].status)
    ```


    :param str private_endpoint_id: The resource ID of the Private Endpoint.
    :param str workspace_id: The resource ID of the Databricks Workspace.
    """
    ...
