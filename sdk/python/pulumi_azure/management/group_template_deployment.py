# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['GroupTemplateDeployment']


class GroupTemplateDeployment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 debug_level: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters_content: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_content: Optional[pulumi.Input[str]] = None,
                 template_spec_version_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Template Deployment at a Management Group Scope.

        > **Note:** Deleting a Deployment at the Management Group Scope will not delete any resources created by the deployment.

        > **Note:** Deployments to a Management Group are always Incrementally applied. Existing resources that are not part of the template will not be removed.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_group = azure.management.get_group(name="00000000-0000-0000-0000-000000000000")
        example_group_template_deployment = azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment",
            location="West Europe",
            management_group_id=example_group.id,
            template_content=\"\"\"{
          "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
          "contentVersion": "1.0.0.0",
          "parameters": {
            "policyAssignmentName": {
              "type": "string",
              "defaultValue": "[guid(parameters('policyDefinitionID'), resourceGroup().name)]",
              "metadata": {
                "description": "Specifies the name of the policy assignment, can be used defined or an idempotent name as the defaultValue provides."
              }
            },
            "policyDefinitionID": {
              "type": "string",
              "metadata": {
                "description": "Specifies the ID of the policy definition or policy set definition being assigned."
              }
            }
          },
          "resources": [
            {
              "type": "Microsoft.Authorization/policyAssignments",
              "name": "[parameters('policyAssignmentName')]",
              "apiVersion": "2019-09-01",
              "properties": {
                "scope": "[subscriptionResourceId('Microsoft.Resources/resourceGroups', resourceGroup().name)]",
                "policyDefinitionId": "[parameters('policyDefinitionID')]"
              }
            }
          ]
        }
        \"\"\",
            parameters_content=\"\"\"{
          "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#",
          "contentVersion": "1.0.0.0",
          "parameters": {
            "policyDefinitionID": {
              "value": "/providers/Microsoft.Authorization/policyDefinitions/0a914e76-4921-4c19-b460-a2d36003525a"
            }
          }
        }
        \"\"\")
        ```

        ```python
        import pulumi
        import pulumi_azure as azure

        example_group = azure.management.get_group(name="00000000-0000-0000-0000-000000000000")
        example_group_template_deployment = azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment",
            location="West Europe",
            management_group_id=example_group.id,
            template_content=(lambda path: open(path).read())("templates/example-deploy-template.json"),
            parameters_content=(lambda path: open(path).read())("templates/example-deploy-params.json"))
        ```

        ```python
        import pulumi
        import pulumi_azure as azure

        example_group = azure.management.get_group(name="00000000-0000-0000-0000-000000000000")
        example_template_spec_version = azure.core.get_template_spec_version(name="exampleTemplateForManagementGroup",
            resource_group_name="exampleResourceGroup",
            version="v1.0.9")
        example_group_template_deployment = azure.management.GroupTemplateDeployment("exampleGroupTemplateDeployment",
            location="West Europe",
            management_group_id=example_group.id,
            template_spec_version_id=example_template_spec_version.id)
        ```

        ## Import

        Management Group Template Deployments can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:management/groupTemplateDeployment:GroupTemplateDeployment example /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] debug_level: The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
        :param pulumi.Input[str] location: The Azure Region where the Template should exist. Changing this forces a new Template to be created.
        :param pulumi.Input[str] name: The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
        :param pulumi.Input[str] parameters_content: The contents of the ARM Template parameters file - containing a JSON list of parameters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Template.
        :param pulumi.Input[str] template_content: The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `template_spec_version_id`.
        :param pulumi.Input[str] template_spec_version_id: The ID of the Template Spec Version to deploy. Cannot be specified with `template_content`.
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

            __props__['debug_level'] = debug_level
            __props__['location'] = location
            if management_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'management_group_id'")
            __props__['management_group_id'] = management_group_id
            __props__['name'] = name
            __props__['parameters_content'] = parameters_content
            __props__['tags'] = tags
            __props__['template_content'] = template_content
            __props__['template_spec_version_id'] = template_spec_version_id
            __props__['output_content'] = None
        super(GroupTemplateDeployment, __self__).__init__(
            'azure:management/groupTemplateDeployment:GroupTemplateDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            debug_level: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            management_group_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            output_content: Optional[pulumi.Input[str]] = None,
            parameters_content: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            template_content: Optional[pulumi.Input[str]] = None,
            template_spec_version_id: Optional[pulumi.Input[str]] = None) -> 'GroupTemplateDeployment':
        """
        Get an existing GroupTemplateDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] debug_level: The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
        :param pulumi.Input[str] location: The Azure Region where the Template should exist. Changing this forces a new Template to be created.
        :param pulumi.Input[str] name: The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
        :param pulumi.Input[str] output_content: The JSON Content of the Outputs of the ARM Template Deployment.
        :param pulumi.Input[str] parameters_content: The contents of the ARM Template parameters file - containing a JSON list of parameters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Template.
        :param pulumi.Input[str] template_content: The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `template_spec_version_id`.
        :param pulumi.Input[str] template_spec_version_id: The ID of the Template Spec Version to deploy. Cannot be specified with `template_content`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["debug_level"] = debug_level
        __props__["location"] = location
        __props__["management_group_id"] = management_group_id
        __props__["name"] = name
        __props__["output_content"] = output_content
        __props__["parameters_content"] = parameters_content
        __props__["tags"] = tags
        __props__["template_content"] = template_content
        __props__["template_spec_version_id"] = template_spec_version_id
        return GroupTemplateDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="debugLevel")
    def debug_level(self) -> pulumi.Output[Optional[str]]:
        """
        The Debug Level which should be used for this Resource Group Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
        """
        return pulumi.get(self, "debug_level")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Template should exist. Changing this forces a new Template to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementGroupId")
    def management_group_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "management_group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Template Deployment. Changing this forces a new Template Deployment to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputContent")
    def output_content(self) -> pulumi.Output[str]:
        """
        The JSON Content of the Outputs of the ARM Template Deployment.
        """
        return pulumi.get(self, "output_content")

    @property
    @pulumi.getter(name="parametersContent")
    def parameters_content(self) -> pulumi.Output[str]:
        """
        The contents of the ARM Template parameters file - containing a JSON list of parameters.
        """
        return pulumi.get(self, "parameters_content")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Template.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateContent")
    def template_content(self) -> pulumi.Output[str]:
        """
        The contents of the ARM Template which should be deployed into this Resource Group. Cannot be specified with `template_spec_version_id`.
        """
        return pulumi.get(self, "template_content")

    @property
    @pulumi.getter(name="templateSpecVersionId")
    def template_spec_version_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Template Spec Version to deploy. Cannot be specified with `template_content`.
        """
        return pulumi.get(self, "template_spec_version_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

