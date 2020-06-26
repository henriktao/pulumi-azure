# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class LinkedService(pulumi.CustomResource):
    linked_service_name: pulumi.Output[str]
    """
    Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspace_name`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
    """
    resource_id: pulumi.Output[str]
    """
    The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    workspace_name: pulumi.Output[str]
    """
    Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, linked_service_name=None, resource_group_name=None, resource_id=None, tags=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Links a Log Analytics (formally Operational Insights) Workspace to another resource. The (currently) only linkable service is an Azure Automation Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku=[{
                "name": "Basic",
            }],
            tags={
                "environment": "development",
            })
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="PerGB2018",
            retention_in_days=30)
        example_linked_service = azure.loganalytics.LinkedService("exampleLinkedService",
            resource_group_name=example_resource_group.name,
            workspace_name=example_analytics_workspace.name,
            resource_id=example_account.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] linked_service_name: Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspace_name`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_id: The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workspace_name: Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['linked_service_name'] = linked_service_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['tags'] = tags
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
        super(LinkedService, __self__).__init__(
            'azure:loganalytics/linkedService:LinkedService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, linked_service_name=None, name=None, resource_group_name=None, resource_id=None, tags=None, workspace_name=None):
        """
        Get an existing LinkedService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] linked_service_name: Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspace_name`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_id: The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] workspace_name: Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["linked_service_name"] = linked_service_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["resource_id"] = resource_id
        __props__["tags"] = tags
        __props__["workspace_name"] = workspace_name
        return LinkedService(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
