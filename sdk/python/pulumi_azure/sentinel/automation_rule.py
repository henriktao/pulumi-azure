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

__all__ = ['AutomationRuleArgs', 'AutomationRule']

@pulumi.input_type
class AutomationRuleArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 log_analytics_workspace_id: pulumi.Input[str],
                 order: pulumi.Input[int],
                 action_incidents: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]] = None,
                 action_playbooks: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AutomationRule resource.
        :param pulumi.Input[str] display_name: The display name which should be used for this Sentinel Automation Rule.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[int] order: The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]] action_incidents: One or more `action_incident` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]] action_playbooks: One or more `action_playbook` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]] conditions: One or more `condition` blocks as defined below.
        :param pulumi.Input[bool] enabled: Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        :param pulumi.Input[str] expiration: The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        :param pulumi.Input[str] name: The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        pulumi.set(__self__, "order", order)
        if action_incidents is not None:
            pulumi.set(__self__, "action_incidents", action_incidents)
        if action_playbooks is not None:
            pulumi.set(__self__, "action_playbooks", action_playbooks)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name which should be used for this Sentinel Automation Rule.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @log_analytics_workspace_id.setter
    def log_analytics_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_analytics_workspace_id", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        """
        The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="actionIncidents")
    def action_incidents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]]:
        """
        One or more `action_incident` blocks as defined below.
        """
        return pulumi.get(self, "action_incidents")

    @action_incidents.setter
    def action_incidents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]]):
        pulumi.set(self, "action_incidents", value)

    @property
    @pulumi.getter(name="actionPlaybooks")
    def action_playbooks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]]:
        """
        One or more `action_playbook` blocks as defined below.
        """
        return pulumi.get(self, "action_playbooks")

    @action_playbooks.setter
    def action_playbooks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]]):
        pulumi.set(self, "action_playbooks", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]]:
        """
        One or more `condition` blocks as defined below.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        """
        The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AutomationRuleState:
    def __init__(__self__, *,
                 action_incidents: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]] = None,
                 action_playbooks: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AutomationRule resources.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]] action_incidents: One or more `action_incident` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]] action_playbooks: One or more `action_playbook` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]] conditions: One or more `condition` blocks as defined below.
        :param pulumi.Input[str] display_name: The display name which should be used for this Sentinel Automation Rule.
        :param pulumi.Input[bool] enabled: Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        :param pulumi.Input[str] expiration: The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[str] name: The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[int] order: The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        if action_incidents is not None:
            pulumi.set(__self__, "action_incidents", action_incidents)
        if action_playbooks is not None:
            pulumi.set(__self__, "action_playbooks", action_playbooks)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if log_analytics_workspace_id is not None:
            pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter(name="actionIncidents")
    def action_incidents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]]:
        """
        One or more `action_incident` blocks as defined below.
        """
        return pulumi.get(self, "action_incidents")

    @action_incidents.setter
    def action_incidents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionIncidentArgs']]]]):
        pulumi.set(self, "action_incidents", value)

    @property
    @pulumi.getter(name="actionPlaybooks")
    def action_playbooks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]]:
        """
        One or more `action_playbook` blocks as defined below.
        """
        return pulumi.get(self, "action_playbooks")

    @action_playbooks.setter
    def action_playbooks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleActionPlaybookArgs']]]]):
        pulumi.set(self, "action_playbooks", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]]:
        """
        One or more `condition` blocks as defined below.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name which should be used for this Sentinel Automation Rule.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        """
        The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @log_analytics_workspace_id.setter
    def log_analytics_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_analytics_workspace_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)


class AutomationRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_incidents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionIncidentArgs']]]]] = None,
                 action_playbooks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionPlaybookArgs']]]]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleConditionArgs']]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a Sentinel Automation Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="west europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="pergb2018")
        sentinel = azure.operationalinsights.AnalyticsSolution("sentinel",
            solution_name="SecurityInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            workspace_resource_id=example_analytics_workspace.id,
            workspace_name=example_analytics_workspace.name,
            plan=azure.operationalinsights.AnalyticsSolutionPlanArgs(
                publisher="Microsoft",
                product="OMSGallery/SecurityInsights",
            ))
        example_automation_rule = azure.sentinel.AutomationRule("exampleAutomationRule",
            log_analytics_workspace_id=sentinel.workspace_resource_id,
            display_name="automation_rule1",
            order=1,
            action_incidents=[azure.sentinel.AutomationRuleActionIncidentArgs(
                order=1,
                status="Active",
            )])
        ```

        ## Import

        Sentinel Automation Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:sentinel/automationRule:AutomationRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/AutomationRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionIncidentArgs']]]] action_incidents: One or more `action_incident` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionPlaybookArgs']]]] action_playbooks: One or more `action_playbook` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleConditionArgs']]]] conditions: One or more `condition` blocks as defined below.
        :param pulumi.Input[str] display_name: The display name which should be used for this Sentinel Automation Rule.
        :param pulumi.Input[bool] enabled: Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        :param pulumi.Input[str] expiration: The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[str] name: The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[int] order: The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutomationRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Sentinel Automation Rule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="west europe")
        example_analytics_workspace = azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="pergb2018")
        sentinel = azure.operationalinsights.AnalyticsSolution("sentinel",
            solution_name="SecurityInsights",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            workspace_resource_id=example_analytics_workspace.id,
            workspace_name=example_analytics_workspace.name,
            plan=azure.operationalinsights.AnalyticsSolutionPlanArgs(
                publisher="Microsoft",
                product="OMSGallery/SecurityInsights",
            ))
        example_automation_rule = azure.sentinel.AutomationRule("exampleAutomationRule",
            log_analytics_workspace_id=sentinel.workspace_resource_id,
            display_name="automation_rule1",
            order=1,
            action_incidents=[azure.sentinel.AutomationRuleActionIncidentArgs(
                order=1,
                status="Active",
            )])
        ```

        ## Import

        Sentinel Automation Rules can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:sentinel/automationRule:AutomationRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/AutomationRules/rule1
        ```

        :param str resource_name: The name of the resource.
        :param AutomationRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutomationRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_incidents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionIncidentArgs']]]]] = None,
                 action_playbooks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionPlaybookArgs']]]]] = None,
                 conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleConditionArgs']]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
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
            __props__ = AutomationRuleArgs.__new__(AutomationRuleArgs)

            __props__.__dict__["action_incidents"] = action_incidents
            __props__.__dict__["action_playbooks"] = action_playbooks
            __props__.__dict__["conditions"] = conditions
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["expiration"] = expiration
            if log_analytics_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_analytics_workspace_id'")
            __props__.__dict__["log_analytics_workspace_id"] = log_analytics_workspace_id
            __props__.__dict__["name"] = name
            if order is None and not opts.urn:
                raise TypeError("Missing required property 'order'")
            __props__.__dict__["order"] = order
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:sentinel/authomationRule:AuthomationRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AutomationRule, __self__).__init__(
            'azure:sentinel/automationRule:AutomationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_incidents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionIncidentArgs']]]]] = None,
            action_playbooks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionPlaybookArgs']]]]] = None,
            conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleConditionArgs']]]]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            log_analytics_workspace_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            order: Optional[pulumi.Input[int]] = None) -> 'AutomationRule':
        """
        Get an existing AutomationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionIncidentArgs']]]] action_incidents: One or more `action_incident` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleActionPlaybookArgs']]]] action_playbooks: One or more `action_playbook` blocks as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRuleConditionArgs']]]] conditions: One or more `condition` blocks as defined below.
        :param pulumi.Input[str] display_name: The display name which should be used for this Sentinel Automation Rule.
        :param pulumi.Input[bool] enabled: Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        :param pulumi.Input[str] expiration: The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        :param pulumi.Input[str] log_analytics_workspace_id: The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[str] name: The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        :param pulumi.Input[int] order: The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutomationRuleState.__new__(_AutomationRuleState)

        __props__.__dict__["action_incidents"] = action_incidents
        __props__.__dict__["action_playbooks"] = action_playbooks
        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["expiration"] = expiration
        __props__.__dict__["log_analytics_workspace_id"] = log_analytics_workspace_id
        __props__.__dict__["name"] = name
        __props__.__dict__["order"] = order
        return AutomationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionIncidents")
    def action_incidents(self) -> pulumi.Output[Optional[Sequence['outputs.AutomationRuleActionIncident']]]:
        """
        One or more `action_incident` blocks as defined below.
        """
        return pulumi.get(self, "action_incidents")

    @property
    @pulumi.getter(name="actionPlaybooks")
    def action_playbooks(self) -> pulumi.Output[Optional[Sequence['outputs.AutomationRuleActionPlaybook']]]:
        """
        One or more `action_playbook` blocks as defined below.
        """
        return pulumi.get(self, "action_playbooks")

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output[Optional[Sequence['outputs.AutomationRuleCondition']]]:
        """
        One or more `condition` blocks as defined below.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name which should be used for this Sentinel Automation Rule.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this Sentinel Automation Rule is enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def expiration(self) -> pulumi.Output[Optional[str]]:
        """
        The time in RFC3339 format of kind `UTC` that determines when this Automation Rule should expire and be disabled.
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[int]:
        """
        The order of this Sentinel Automation Rule. Possible values varies between `1` and `1000`.
        """
        return pulumi.get(self, "order")

