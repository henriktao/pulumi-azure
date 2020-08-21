# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Definition(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the policy definition.
    """
    display_name: pulumi.Output[str]
    """
    The display name of the policy definition.
    """
    management_group_id: pulumi.Output[str]
    """
    The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
    """
    management_group_name: pulumi.Output[str]
    """
    The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
    """
    metadata: pulumi.Output[str]
    """
    The metadata for the policy definition. This
    is a json object representing additional metadata that should be stored
    with the policy definition.
    """
    mode: pulumi.Output[str]
    """
    The policy mode that allows you to specify which resource
    types will be evaluated.  The value can be "All", "Indexed" or
    "NotSpecified".
    """
    name: pulumi.Output[str]
    """
    The name of the policy definition. Changing this forces a
    new resource to be created.
    """
    parameters: pulumi.Output[str]
    """
    Parameters for the policy definition. This field
    is a json object that allows you to parameterize your policy definition.
    """
    policy_rule: pulumi.Output[str]
    """
    The policy rule for the policy definition. This
    is a json object representing the rule that contains an if and
    a then block.
    """
    policy_type: pulumi.Output[str]
    """
    The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, management_group_id=None, management_group_name=None, metadata=None, mode=None, name=None, parameters=None, policy_rule=None, policy_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a policy rule definition on a management group or your provider subscription.

        Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        policy = azure.policy.Definition("policy",
            display_name="acceptance test policy definition",
            metadata=\"\"\"    {
            "category": "General"
            }


        \"\"\",
            mode="Indexed",
            parameters=\"\"\"	{
            "allowedLocations": {
              "type": "Array",
              "metadata": {
                "description": "The list of allowed locations for resources.",
                "displayName": "Allowed locations",
                "strongType": "location"
              }
            }
          }

        \"\"\",
            policy_rule=\"\"\"	{
            "if": {
              "not": {
                "field": "location",
                "in": "[parameters('allowedLocations')]"
              }
            },
            "then": {
              "effect": "audit"
            }
          }

        \"\"\",
            policy_type="Custom")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy definition.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy definition. This
               is a json object representing additional metadata that should be stored
               with the policy definition.
        :param pulumi.Input[str] mode: The policy mode that allows you to specify which resource
               types will be evaluated.  The value can be "All", "Indexed" or
               "NotSpecified".
        :param pulumi.Input[str] name: The name of the policy definition. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field
               is a json object that allows you to parameterize your policy definition.
        :param pulumi.Input[str] policy_rule: The policy rule for the policy definition. This
               is a json object representing the rule that contains an if and
               a then block.
        :param pulumi.Input[str] policy_type: The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
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

            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if management_group_id is not None:
                warnings.warn("Deprecated in favour of `management_group_name`", DeprecationWarning)
                pulumi.log.warn("management_group_id is deprecated: Deprecated in favour of `management_group_name`")
            __props__['management_group_id'] = management_group_id
            __props__['management_group_name'] = management_group_name
            __props__['metadata'] = metadata
            if mode is None:
                raise TypeError("Missing required property 'mode'")
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_rule'] = policy_rule
            if policy_type is None:
                raise TypeError("Missing required property 'policy_type'")
            __props__['policy_type'] = policy_type
        super(Definition, __self__).__init__(
            'azure:policy/definition:Definition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, display_name=None, management_group_id=None, management_group_name=None, metadata=None, mode=None, name=None, parameters=None, policy_rule=None, policy_type=None):
        """
        Get an existing Definition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy definition.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[str] management_group_id: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_group_name: The name of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
        :param pulumi.Input[str] metadata: The metadata for the policy definition. This
               is a json object representing additional metadata that should be stored
               with the policy definition.
        :param pulumi.Input[str] mode: The policy mode that allows you to specify which resource
               types will be evaluated.  The value can be "All", "Indexed" or
               "NotSpecified".
        :param pulumi.Input[str] name: The name of the policy definition. Changing this forces a
               new resource to be created.
        :param pulumi.Input[str] parameters: Parameters for the policy definition. This field
               is a json object that allows you to parameterize your policy definition.
        :param pulumi.Input[str] policy_rule: The policy rule for the policy definition. This
               is a json object representing the rule that contains an if and
               a then block.
        :param pulumi.Input[str] policy_type: The policy type. Possible values are `BuiltIn`, `Custom` and `NotSpecified`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["management_group_id"] = management_group_id
        __props__["management_group_name"] = management_group_name
        __props__["metadata"] = metadata
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["parameters"] = parameters
        __props__["policy_rule"] = policy_rule
        __props__["policy_type"] = policy_type
        return Definition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
