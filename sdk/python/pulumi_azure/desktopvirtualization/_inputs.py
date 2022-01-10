# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'HostPoolRegistrationInfoArgs',
    'ScalingPlanHostPoolArgs',
    'ScalingPlanScheduleArgs',
]

@pulumi.input_type
class HostPoolRegistrationInfoArgs:
    def __init__(__self__, *,
                 expiration_date: pulumi.Input[str],
                 reset_token: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expiration_date: A valid `RFC3339Time` for the expiration of the token.
        :param pulumi.Input[str] token: The registration token generated by the Virtual Desktop Host Pool.
        """
        pulumi.set(__self__, "expiration_date", expiration_date)
        if reset_token is not None:
            pulumi.set(__self__, "reset_token", reset_token)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Input[str]:
        """
        A valid `RFC3339Time` for the expiration of the token.
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: pulumi.Input[str]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="resetToken")
    def reset_token(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "reset_token")

    @reset_token.setter
    def reset_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_token", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The registration token generated by the Virtual Desktop Host Pool.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class ScalingPlanHostPoolArgs:
    def __init__(__self__, *,
                 hostpool_id: pulumi.Input[str],
                 scaling_plan_enabled: pulumi.Input[bool]):
        """
        :param pulumi.Input[str] hostpool_id: The ID of the HostPool to assign the Scaling Plan to.
        :param pulumi.Input[bool] scaling_plan_enabled: Specifies if the scaling plan is enabled or disabled for the HostPool.
        """
        pulumi.set(__self__, "hostpool_id", hostpool_id)
        pulumi.set(__self__, "scaling_plan_enabled", scaling_plan_enabled)

    @property
    @pulumi.getter(name="hostpoolId")
    def hostpool_id(self) -> pulumi.Input[str]:
        """
        The ID of the HostPool to assign the Scaling Plan to.
        """
        return pulumi.get(self, "hostpool_id")

    @hostpool_id.setter
    def hostpool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostpool_id", value)

    @property
    @pulumi.getter(name="scalingPlanEnabled")
    def scaling_plan_enabled(self) -> pulumi.Input[bool]:
        """
        Specifies if the scaling plan is enabled or disabled for the HostPool.
        """
        return pulumi.get(self, "scaling_plan_enabled")

    @scaling_plan_enabled.setter
    def scaling_plan_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "scaling_plan_enabled", value)


@pulumi.input_type
class ScalingPlanScheduleArgs:
    def __init__(__self__, *,
                 days_of_weeks: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: pulumi.Input[str],
                 off_peak_load_balancing_algorithm: pulumi.Input[str],
                 off_peak_start_time: pulumi.Input[str],
                 peak_load_balancing_algorithm: pulumi.Input[str],
                 peak_start_time: pulumi.Input[str],
                 ramp_down_capacity_threshold_percent: pulumi.Input[int],
                 ramp_down_force_logoff_users: pulumi.Input[bool],
                 ramp_down_load_balancing_algorithm: pulumi.Input[str],
                 ramp_down_minimum_hosts_percent: pulumi.Input[int],
                 ramp_down_notification_message: pulumi.Input[str],
                 ramp_down_start_time: pulumi.Input[str],
                 ramp_down_stop_hosts_when: pulumi.Input[str],
                 ramp_down_wait_time_minutes: pulumi.Input[int],
                 ramp_up_load_balancing_algorithm: pulumi.Input[str],
                 ramp_up_start_time: pulumi.Input[str],
                 ramp_up_capacity_threshold_percent: Optional[pulumi.Input[int]] = None,
                 ramp_up_minimum_hosts_percent: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] days_of_weeks: A list of Days of the Week on which this schedule will be used..Possible values are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`
        :param pulumi.Input[str] name: The name of the schedule.
        :param pulumi.Input[str] off_peak_load_balancing_algorithm: The load Balancing Algorithm to use during Off-Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        :param pulumi.Input[str] off_peak_start_time: The time at which Off-Peak scaling will begin. This is also the end-time for the Ramp-Down period. The time must be specified in "HH:MM" format.
        :param pulumi.Input[str] peak_load_balancing_algorithm: The load Balancing Algorithm to use during Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        :param pulumi.Input[str] peak_start_time: The time at which Peak scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param pulumi.Input[int] ramp_down_capacity_threshold_percent: This is the value in percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-down and off-peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        :param pulumi.Input[bool] ramp_down_force_logoff_users: Whether users will be forced to log-off session hosts once the `ramp_down_wait_time_minutes` value has been exceeded during the Ramp-Down period. Possible
        :param pulumi.Input[str] ramp_down_load_balancing_algorithm: The load Balancing Algorithm to use during the Ramp-Down period. Possible values are `DepthFirst` and `BreadthFirst`.
        :param pulumi.Input[int] ramp_down_minimum_hosts_percent: The minimum percentage of session host virtual machines that you would like to get to for ramp-down and off-peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        :param pulumi.Input[str] ramp_down_notification_message: The notification message to send to users during Ramp-Down period when they are required to log-off.
        :param pulumi.Input[str] ramp_down_start_time: The time at which Ramp-Down scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param pulumi.Input[str] ramp_down_stop_hosts_when: Controls Session Host shutdown behaviour during Ramp-Down period. Session Hosts can either be shutdown when all sessions on the Session Host have ended, or when there are no Active sessions left on the Session Host. Possible values are `ZeroSessions` and `ZeroActiveSessions`.
        :param pulumi.Input[int] ramp_down_wait_time_minutes: The number of minutes during Ramp-Down period that autoscale will wait after setting the session host VMs to drain mode, notifying any currently signed in users to save their work before forcing the users to logoff. Once all user sessions on the session host VM have been logged off, Autoscale will shut down the VM.
        :param pulumi.Input[str] ramp_up_load_balancing_algorithm: The load Balancing Algorithm to use during the Ramp-Up period. Possible values are `DepthFirst` and `BreadthFirst`.
        :param pulumi.Input[str] ramp_up_start_time: The time at which Ramp-Up scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param pulumi.Input[int] ramp_up_capacity_threshold_percent: Specify minimum percentage of session host virtual machines to start for ramp-up and peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        :param pulumi.Input[int] ramp_up_minimum_hosts_percent: This is the value of percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-up and peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        """
        pulumi.set(__self__, "days_of_weeks", days_of_weeks)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "off_peak_load_balancing_algorithm", off_peak_load_balancing_algorithm)
        pulumi.set(__self__, "off_peak_start_time", off_peak_start_time)
        pulumi.set(__self__, "peak_load_balancing_algorithm", peak_load_balancing_algorithm)
        pulumi.set(__self__, "peak_start_time", peak_start_time)
        pulumi.set(__self__, "ramp_down_capacity_threshold_percent", ramp_down_capacity_threshold_percent)
        pulumi.set(__self__, "ramp_down_force_logoff_users", ramp_down_force_logoff_users)
        pulumi.set(__self__, "ramp_down_load_balancing_algorithm", ramp_down_load_balancing_algorithm)
        pulumi.set(__self__, "ramp_down_minimum_hosts_percent", ramp_down_minimum_hosts_percent)
        pulumi.set(__self__, "ramp_down_notification_message", ramp_down_notification_message)
        pulumi.set(__self__, "ramp_down_start_time", ramp_down_start_time)
        pulumi.set(__self__, "ramp_down_stop_hosts_when", ramp_down_stop_hosts_when)
        pulumi.set(__self__, "ramp_down_wait_time_minutes", ramp_down_wait_time_minutes)
        pulumi.set(__self__, "ramp_up_load_balancing_algorithm", ramp_up_load_balancing_algorithm)
        pulumi.set(__self__, "ramp_up_start_time", ramp_up_start_time)
        if ramp_up_capacity_threshold_percent is not None:
            pulumi.set(__self__, "ramp_up_capacity_threshold_percent", ramp_up_capacity_threshold_percent)
        if ramp_up_minimum_hosts_percent is not None:
            pulumi.set(__self__, "ramp_up_minimum_hosts_percent", ramp_up_minimum_hosts_percent)

    @property
    @pulumi.getter(name="daysOfWeeks")
    def days_of_weeks(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of Days of the Week on which this schedule will be used..Possible values are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`
        """
        return pulumi.get(self, "days_of_weeks")

    @days_of_weeks.setter
    def days_of_weeks(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "days_of_weeks", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the schedule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="offPeakLoadBalancingAlgorithm")
    def off_peak_load_balancing_algorithm(self) -> pulumi.Input[str]:
        """
        The load Balancing Algorithm to use during Off-Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "off_peak_load_balancing_algorithm")

    @off_peak_load_balancing_algorithm.setter
    def off_peak_load_balancing_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "off_peak_load_balancing_algorithm", value)

    @property
    @pulumi.getter(name="offPeakStartTime")
    def off_peak_start_time(self) -> pulumi.Input[str]:
        """
        The time at which Off-Peak scaling will begin. This is also the end-time for the Ramp-Down period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "off_peak_start_time")

    @off_peak_start_time.setter
    def off_peak_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "off_peak_start_time", value)

    @property
    @pulumi.getter(name="peakLoadBalancingAlgorithm")
    def peak_load_balancing_algorithm(self) -> pulumi.Input[str]:
        """
        The load Balancing Algorithm to use during Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "peak_load_balancing_algorithm")

    @peak_load_balancing_algorithm.setter
    def peak_load_balancing_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "peak_load_balancing_algorithm", value)

    @property
    @pulumi.getter(name="peakStartTime")
    def peak_start_time(self) -> pulumi.Input[str]:
        """
        The time at which Peak scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "peak_start_time")

    @peak_start_time.setter
    def peak_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "peak_start_time", value)

    @property
    @pulumi.getter(name="rampDownCapacityThresholdPercent")
    def ramp_down_capacity_threshold_percent(self) -> pulumi.Input[int]:
        """
        This is the value in percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-down and off-peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        """
        return pulumi.get(self, "ramp_down_capacity_threshold_percent")

    @ramp_down_capacity_threshold_percent.setter
    def ramp_down_capacity_threshold_percent(self, value: pulumi.Input[int]):
        pulumi.set(self, "ramp_down_capacity_threshold_percent", value)

    @property
    @pulumi.getter(name="rampDownForceLogoffUsers")
    def ramp_down_force_logoff_users(self) -> pulumi.Input[bool]:
        """
        Whether users will be forced to log-off session hosts once the `ramp_down_wait_time_minutes` value has been exceeded during the Ramp-Down period. Possible
        """
        return pulumi.get(self, "ramp_down_force_logoff_users")

    @ramp_down_force_logoff_users.setter
    def ramp_down_force_logoff_users(self, value: pulumi.Input[bool]):
        pulumi.set(self, "ramp_down_force_logoff_users", value)

    @property
    @pulumi.getter(name="rampDownLoadBalancingAlgorithm")
    def ramp_down_load_balancing_algorithm(self) -> pulumi.Input[str]:
        """
        The load Balancing Algorithm to use during the Ramp-Down period. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "ramp_down_load_balancing_algorithm")

    @ramp_down_load_balancing_algorithm.setter
    def ramp_down_load_balancing_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_down_load_balancing_algorithm", value)

    @property
    @pulumi.getter(name="rampDownMinimumHostsPercent")
    def ramp_down_minimum_hosts_percent(self) -> pulumi.Input[int]:
        """
        The minimum percentage of session host virtual machines that you would like to get to for ramp-down and off-peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        """
        return pulumi.get(self, "ramp_down_minimum_hosts_percent")

    @ramp_down_minimum_hosts_percent.setter
    def ramp_down_minimum_hosts_percent(self, value: pulumi.Input[int]):
        pulumi.set(self, "ramp_down_minimum_hosts_percent", value)

    @property
    @pulumi.getter(name="rampDownNotificationMessage")
    def ramp_down_notification_message(self) -> pulumi.Input[str]:
        """
        The notification message to send to users during Ramp-Down period when they are required to log-off.
        """
        return pulumi.get(self, "ramp_down_notification_message")

    @ramp_down_notification_message.setter
    def ramp_down_notification_message(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_down_notification_message", value)

    @property
    @pulumi.getter(name="rampDownStartTime")
    def ramp_down_start_time(self) -> pulumi.Input[str]:
        """
        The time at which Ramp-Down scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "ramp_down_start_time")

    @ramp_down_start_time.setter
    def ramp_down_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_down_start_time", value)

    @property
    @pulumi.getter(name="rampDownStopHostsWhen")
    def ramp_down_stop_hosts_when(self) -> pulumi.Input[str]:
        """
        Controls Session Host shutdown behaviour during Ramp-Down period. Session Hosts can either be shutdown when all sessions on the Session Host have ended, or when there are no Active sessions left on the Session Host. Possible values are `ZeroSessions` and `ZeroActiveSessions`.
        """
        return pulumi.get(self, "ramp_down_stop_hosts_when")

    @ramp_down_stop_hosts_when.setter
    def ramp_down_stop_hosts_when(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_down_stop_hosts_when", value)

    @property
    @pulumi.getter(name="rampDownWaitTimeMinutes")
    def ramp_down_wait_time_minutes(self) -> pulumi.Input[int]:
        """
        The number of minutes during Ramp-Down period that autoscale will wait after setting the session host VMs to drain mode, notifying any currently signed in users to save their work before forcing the users to logoff. Once all user sessions on the session host VM have been logged off, Autoscale will shut down the VM.
        """
        return pulumi.get(self, "ramp_down_wait_time_minutes")

    @ramp_down_wait_time_minutes.setter
    def ramp_down_wait_time_minutes(self, value: pulumi.Input[int]):
        pulumi.set(self, "ramp_down_wait_time_minutes", value)

    @property
    @pulumi.getter(name="rampUpLoadBalancingAlgorithm")
    def ramp_up_load_balancing_algorithm(self) -> pulumi.Input[str]:
        """
        The load Balancing Algorithm to use during the Ramp-Up period. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "ramp_up_load_balancing_algorithm")

    @ramp_up_load_balancing_algorithm.setter
    def ramp_up_load_balancing_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_up_load_balancing_algorithm", value)

    @property
    @pulumi.getter(name="rampUpStartTime")
    def ramp_up_start_time(self) -> pulumi.Input[str]:
        """
        The time at which Ramp-Up scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "ramp_up_start_time")

    @ramp_up_start_time.setter
    def ramp_up_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "ramp_up_start_time", value)

    @property
    @pulumi.getter(name="rampUpCapacityThresholdPercent")
    def ramp_up_capacity_threshold_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Specify minimum percentage of session host virtual machines to start for ramp-up and peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        """
        return pulumi.get(self, "ramp_up_capacity_threshold_percent")

    @ramp_up_capacity_threshold_percent.setter
    def ramp_up_capacity_threshold_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ramp_up_capacity_threshold_percent", value)

    @property
    @pulumi.getter(name="rampUpMinimumHostsPercent")
    def ramp_up_minimum_hosts_percent(self) -> Optional[pulumi.Input[int]]:
        """
        This is the value of percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-up and peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        """
        return pulumi.get(self, "ramp_up_minimum_hosts_percent")

    @ramp_up_minimum_hosts_percent.setter
    def ramp_up_minimum_hosts_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ramp_up_minimum_hosts_percent", value)


