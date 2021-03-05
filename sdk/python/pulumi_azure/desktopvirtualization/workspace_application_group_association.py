# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['WorkspaceApplicationGroupAssociation']


class WorkspaceApplicationGroupAssociation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_group_id: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Virtual Desktop Workspace Application Group Association.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example = azure.core.ResourceGroup("example", location="West Europe")
        pooledbreadthfirst = azure.desktopvirtualization.HostPool("pooledbreadthfirst",
            location=example.location,
            resource_group_name=example.name,
            type="Pooled",
            load_balancer_type="BreadthFirst")
        remoteapp = azure.desktopvirtualization.ApplicationGroup("remoteapp",
            location=example.location,
            resource_group_name=example.name,
            type="RemoteApp",
            host_pool_id=pooledbreadthfirst.id)
        workspace = azure.desktopvirtualization.Workspace("workspace",
            location=example.location,
            resource_group_name=example.name)
        workspaceremoteapp = azure.desktopvirtualization.WorkspaceApplicationGroupAssociation("workspaceremoteapp",
            workspace_id=workspace.id,
            application_group_id=remoteapp.id)
        ```

        ## Import

        Associations between Virtual Desktop Workspaces and Virtual Desktop Application Groups can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myGroup1/providers/Microsoft.DesktopVirtualization/workspaces/myworkspace|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/myapplicationgroup"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_group_id: The resource ID for the Virtual Desktop Application Group.
        :param pulumi.Input[str] workspace_id: The resource ID for the Virtual Desktop Workspace.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if application_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_group_id'")
            __props__['application_group_id'] = application_group_id
            if workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_id'")
            __props__['workspace_id'] = workspace_id
        super(WorkspaceApplicationGroupAssociation, __self__).__init__(
            'azure:desktopvirtualization/workspaceApplicationGroupAssociation:WorkspaceApplicationGroupAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_group_id: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'WorkspaceApplicationGroupAssociation':
        """
        Get an existing WorkspaceApplicationGroupAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_group_id: The resource ID for the Virtual Desktop Application Group.
        :param pulumi.Input[str] workspace_id: The resource ID for the Virtual Desktop Workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_group_id"] = application_group_id
        __props__["workspace_id"] = workspace_id
        return WorkspaceApplicationGroupAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationGroupId")
    def application_group_id(self) -> pulumi.Output[str]:
        """
        The resource ID for the Virtual Desktop Application Group.
        """
        return pulumi.get(self, "application_group_id")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        The resource ID for the Virtual Desktop Workspace.
        """
        return pulumi.get(self, "workspace_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

