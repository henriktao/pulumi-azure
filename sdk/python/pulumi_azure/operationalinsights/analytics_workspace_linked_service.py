# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class AnalyticsWorkspaceLinkedService(pulumi.CustomResource):
    """
    Links a Log Analytics (formally Operational Insights) Workspace to another resource. The (currently) only linkable service is an Azure Automation Account.
    """
    def __init__(__self__, __name__, __opts__=None, linked_service_name=None, linked_service_properties=None, resource_group_name=None, tags=None, workspace_name=None):
        """Create a AnalyticsWorkspaceLinkedService resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['linked_service_name'] = linked_service_name

        if not linked_service_properties:
            raise TypeError('Missing required property linked_service_properties')
        __props__['linked_service_properties'] = linked_service_properties

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['tags'] = tags

        if not workspace_name:
            raise TypeError('Missing required property workspace_name')
        __props__['workspace_name'] = workspace_name

        __props__['name'] = None

        super(AnalyticsWorkspaceLinkedService, __self__).__init__(
            'azure:operationalinsights/analyticsWorkspaceLinkedService:AnalyticsWorkspaceLinkedService',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

