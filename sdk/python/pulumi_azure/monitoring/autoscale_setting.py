# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AutoscaleSetting']


class AutoscaleSetting(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification: Optional[pulumi.Input[pulumi.InputType['AutoscaleSettingNotificationArgs']]] = None,
                 profiles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscaleSettingProfileArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a AutoScale Setting which can be applied to Virtual Machine Scale Sets, App Services and other scalable resources.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_scale_set = azure.compute.ScaleSet("exampleScaleSet")
        # ...
        example_autoscale_setting = azure.monitoring.AutoscaleSetting("exampleAutoscaleSetting",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            target_resource_id=example_scale_set.id,
            profiles=[azure.monitoring.AutoscaleSettingProfileArgs(
                name="defaultProfile",
                capacity=azure.monitoring.AutoscaleSettingProfileCapacityArgs(
                    default=1,
                    minimum=1,
                    maximum=10,
                ),
                rules=[
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="GreaterThan",
                            threshold=75,
                            metric_namespace="microsoft.compute/virtualmachinescalesets",
                            dimensions=[azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerDimensionArgs(
                                name="AppName",
                                operator="Equals",
                                values=["App1"],
                            )],
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Increase",
                            type="ChangeCount",
                            value=1,
                            cooldown="PT1M",
                        ),
                    ),
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="LessThan",
                            threshold=25,
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Decrease",
                            type="ChangeCount",
                            value=1,
                            cooldown="PT1M",
                        ),
                    ),
                ],
            )],
            notification=azure.monitoring.AutoscaleSettingNotificationArgs(
                email=azure.monitoring.AutoscaleSettingNotificationEmailArgs(
                    send_to_subscription_administrator=True,
                    send_to_subscription_co_administrator=True,
                    custom_emails=["admin@contoso.com"],
                ),
            ))
        ```
        ### Repeating On Weekends)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_scale_set = azure.compute.ScaleSet("exampleScaleSet")
        # ...
        example_autoscale_setting = azure.monitoring.AutoscaleSetting("exampleAutoscaleSetting",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            target_resource_id=example_scale_set.id,
            profiles=[azure.monitoring.AutoscaleSettingProfileArgs(
                name="Weekends",
                capacity=azure.monitoring.AutoscaleSettingProfileCapacityArgs(
                    default=1,
                    minimum=1,
                    maximum=10,
                ),
                rules=[
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="GreaterThan",
                            threshold=90,
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Increase",
                            type="ChangeCount",
                            value=2,
                            cooldown="PT1M",
                        ),
                    ),
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="LessThan",
                            threshold=10,
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Decrease",
                            type="ChangeCount",
                            value=2,
                            cooldown="PT1M",
                        ),
                    ),
                ],
                recurrence=azure.monitoring.AutoscaleSettingProfileRecurrenceArgs(
                    frequency="Week",
                    timezone="Pacific Standard Time",
                    days=[
                        "Saturday",
                        "Sunday",
                    ],
                    hours=[12],
                    minutes=[0],
                ),
            )],
            notification=azure.monitoring.AutoscaleSettingNotificationArgs(
                email=azure.monitoring.AutoscaleSettingNotificationEmailArgs(
                    send_to_subscription_administrator=True,
                    send_to_subscription_co_administrator=True,
                    custom_emails=["admin@contoso.com"],
                ),
            ))
        ```
        ### For Fixed Dates)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_scale_set = azure.compute.ScaleSet("exampleScaleSet")
        # ...
        example_autoscale_setting = azure.monitoring.AutoscaleSetting("exampleAutoscaleSetting",
            enabled=True,
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            target_resource_id=example_scale_set.id,
            profiles=[azure.monitoring.AutoscaleSettingProfileArgs(
                name="forJuly",
                capacity=azure.monitoring.AutoscaleSettingProfileCapacityArgs(
                    default=1,
                    minimum=1,
                    maximum=10,
                ),
                rules=[
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="GreaterThan",
                            threshold=90,
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Increase",
                            type="ChangeCount",
                            value=2,
                            cooldown="PT1M",
                        ),
                    ),
                    azure.monitoring.AutoscaleSettingProfileRuleArgs(
                        metric_trigger=azure.monitoring.AutoscaleSettingProfileRuleMetricTriggerArgs(
                            metric_name="Percentage CPU",
                            metric_resource_id=example_scale_set.id,
                            time_grain="PT1M",
                            statistic="Average",
                            time_window="PT5M",
                            time_aggregation="Average",
                            operator="LessThan",
                            threshold=10,
                        ),
                        scale_action=azure.monitoring.AutoscaleSettingProfileRuleScaleActionArgs(
                            direction="Decrease",
                            type="ChangeCount",
                            value=2,
                            cooldown="PT1M",
                        ),
                    ),
                ],
                fixed_date=azure.monitoring.AutoscaleSettingProfileFixedDateArgs(
                    timezone="Pacific Standard Time",
                    start="2020-07-01T00:00:00Z",
                    end="2020-07-31T23:59:59Z",
                ),
            )],
            notification=azure.monitoring.AutoscaleSettingNotificationArgs(
                email=azure.monitoring.AutoscaleSettingNotificationEmailArgs(
                    send_to_subscription_administrator=True,
                    send_to_subscription_co_administrator=True,
                    custom_emails=["admin@contoso.com"],
                ),
            ))
        ```

        ## Import

        AutoScale Setting can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:monitoring/autoscaleSetting:AutoscaleSetting example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/microsoft.insights/autoscalesettings/setting1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the AutoScale Setting. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['AutoscaleSettingNotificationArgs']] notification: Specifies a `notification` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscaleSettingProfileArgs']]]] profiles: Specifies one or more (up to 20) `profile` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_resource_id: Specifies the resource ID of the resource that the autoscale setting should be added to.
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

            __props__['enabled'] = enabled
            __props__['location'] = location
            __props__['name'] = name
            __props__['notification'] = notification
            if profiles is None and not opts.urn:
                raise TypeError("Missing required property 'profiles'")
            __props__['profiles'] = profiles
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if target_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_id'")
            __props__['target_resource_id'] = target_resource_id
        super(AutoscaleSetting, __self__).__init__(
            'azure:monitoring/autoscaleSetting:AutoscaleSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification: Optional[pulumi.Input[pulumi.InputType['AutoscaleSettingNotificationArgs']]] = None,
            profiles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscaleSettingProfileArgs']]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_resource_id: Optional[pulumi.Input[str]] = None) -> 'AutoscaleSetting':
        """
        Get an existing AutoscaleSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the AutoScale Setting. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['AutoscaleSettingNotificationArgs']] notification: Specifies a `notification` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoscaleSettingProfileArgs']]]] profiles: Specifies one or more (up to 20) `profile` blocks as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_resource_id: Specifies the resource ID of the resource that the autoscale setting should be added to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enabled"] = enabled
        __props__["location"] = location
        __props__["name"] = name
        __props__["notification"] = notification
        __props__["profiles"] = profiles
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["target_resource_id"] = target_resource_id
        return AutoscaleSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the AutoScale Setting. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notification(self) -> pulumi.Output[Optional['outputs.AutoscaleSettingNotification']]:
        """
        Specifies a `notification` block as defined below.
        """
        return pulumi.get(self, "notification")

    @property
    @pulumi.getter
    def profiles(self) -> pulumi.Output[Sequence['outputs.AutoscaleSettingProfile']]:
        """
        Specifies one or more (up to 20) `profile` blocks as defined below.
        """
        return pulumi.get(self, "profiles")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> pulumi.Output[str]:
        """
        Specifies the resource ID of the resource that the autoscale setting should be added to.
        """
        return pulumi.get(self, "target_resource_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

