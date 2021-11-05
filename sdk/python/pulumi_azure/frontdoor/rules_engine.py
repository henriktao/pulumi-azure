# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RulesEngineArgs', 'RulesEngine']

@pulumi.input_type
class RulesEngineArgs:
    def __init__(__self__, *,
                 frontdoor_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]] = None):
        """
        The set of arguments for constructing a RulesEngine resource.
        :param pulumi.Input[str] frontdoor_name: The name of the Front Door instance. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Rules engine configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]] rules: A `rule` block as defined below.
        """
        pulumi.set(__self__, "frontdoor_name", frontdoor_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="frontdoorName")
    def frontdoor_name(self) -> pulumi.Input[str]:
        """
        The name of the Front Door instance. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "frontdoor_name")

    @frontdoor_name.setter
    def frontdoor_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "frontdoor_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Rules engine configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]]:
        """
        A `rule` block as defined below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class _RulesEngineState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 frontdoor_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering RulesEngine resources.
        :param pulumi.Input[str] frontdoor_name: The name of the Front Door instance. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Rules engine configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]] rules: A `rule` block as defined below.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if frontdoor_name is not None:
            pulumi.set(__self__, "frontdoor_name", frontdoor_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="frontdoorName")
    def frontdoor_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Front Door instance. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "frontdoor_name")

    @frontdoor_name.setter
    def frontdoor_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frontdoor_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Rules engine configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]]:
        """
        A `rule` block as defined below.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RulesEngineRuleArgs']]]]):
        pulumi.set(self, "rules", value)


class RulesEngine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 frontdoor_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RulesEngineRuleArgs']]]]] = None,
                 __props__=None):
        """
        Manages an Azure Front Door Rules Engine configuration and rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_rules_engine = azure.frontdoor.RulesEngine("exampleRulesEngine",
            frontdoor_name=azurerm_frontdoor["example"]["name"],
            resource_group_name=azurerm_frontdoor["example"]["resource_group_name"],
            rules=[
                azure.frontdoor.RulesEngineRuleArgs(
                    name="debuggingoutput",
                    priority=1,
                    action=azure.frontdoor.RulesEngineRuleActionArgs(
                        response_headers=[azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                            header_action_type="Append",
                            header_name="X-TEST-HEADER",
                            value="Append Header Rule",
                        )],
                    ),
                ),
                azure.frontdoor.RulesEngineRuleArgs(
                    name="overwriteorigin",
                    priority=2,
                    match_conditions=[azure.frontdoor.RulesEngineRuleMatchConditionArgs(
                        variable="RequestMethod",
                        operator="Equal",
                        values=[
                            "GET",
                            "POST",
                        ],
                    )],
                    action=azure.frontdoor.RulesEngineRuleActionArgs(
                        response_headers=[
                            azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                                header_action_type="Overwrite",
                                header_name="Access-Control-Allow-Origin",
                                value="*",
                            ),
                            azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                                header_action_type="Overwrite",
                                header_name="Access-Control-Allow-Credentials",
                                value="true",
                            ),
                        ],
                    ),
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] frontdoor_name: The name of the Front Door instance. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Rules engine configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RulesEngineRuleArgs']]]] rules: A `rule` block as defined below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RulesEngineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Front Door Rules Engine configuration and rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_rules_engine = azure.frontdoor.RulesEngine("exampleRulesEngine",
            frontdoor_name=azurerm_frontdoor["example"]["name"],
            resource_group_name=azurerm_frontdoor["example"]["resource_group_name"],
            rules=[
                azure.frontdoor.RulesEngineRuleArgs(
                    name="debuggingoutput",
                    priority=1,
                    action=azure.frontdoor.RulesEngineRuleActionArgs(
                        response_headers=[azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                            header_action_type="Append",
                            header_name="X-TEST-HEADER",
                            value="Append Header Rule",
                        )],
                    ),
                ),
                azure.frontdoor.RulesEngineRuleArgs(
                    name="overwriteorigin",
                    priority=2,
                    match_conditions=[azure.frontdoor.RulesEngineRuleMatchConditionArgs(
                        variable="RequestMethod",
                        operator="Equal",
                        values=[
                            "GET",
                            "POST",
                        ],
                    )],
                    action=azure.frontdoor.RulesEngineRuleActionArgs(
                        response_headers=[
                            azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                                header_action_type="Overwrite",
                                header_name="Access-Control-Allow-Origin",
                                value="*",
                            ),
                            azure.frontdoor.RulesEngineRuleActionResponseHeaderArgs(
                                header_action_type="Overwrite",
                                header_name="Access-Control-Allow-Credentials",
                                value="true",
                            ),
                        ],
                    ),
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param RulesEngineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RulesEngineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 frontdoor_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RulesEngineRuleArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RulesEngineArgs.__new__(RulesEngineArgs)

            __props__.__dict__["enabled"] = enabled
            if frontdoor_name is None and not opts.urn:
                raise TypeError("Missing required property 'frontdoor_name'")
            __props__.__dict__["frontdoor_name"] = frontdoor_name
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["rules"] = rules
            __props__.__dict__["location"] = None
        super(RulesEngine, __self__).__init__(
            'azure:frontdoor/rulesEngine:RulesEngine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            frontdoor_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RulesEngineRuleArgs']]]]] = None) -> 'RulesEngine':
        """
        Get an existing RulesEngine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] frontdoor_name: The name of the Front Door instance. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Rules engine configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RulesEngineRuleArgs']]]] rules: A `rule` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RulesEngineState.__new__(_RulesEngineState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["frontdoor_name"] = frontdoor_name
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["rules"] = rules
        return RulesEngine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="frontdoorName")
    def frontdoor_name(self) -> pulumi.Output[str]:
        """
        The name of the Front Door instance. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "frontdoor_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Rules engine configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.RulesEngineRule']]]:
        """
        A `rule` block as defined below.
        """
        return pulumi.get(self, "rules")

