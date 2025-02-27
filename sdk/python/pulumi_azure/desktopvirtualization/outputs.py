# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'HostPoolRegistrationInfo',
    'ScalingPlanHostPool',
    'ScalingPlanSchedule',
]

@pulumi.output_type
class HostPoolRegistrationInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expirationDate":
            suggest = "expiration_date"
        elif key == "resetToken":
            suggest = "reset_token"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HostPoolRegistrationInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HostPoolRegistrationInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HostPoolRegistrationInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expiration_date: Optional[str] = None,
                 reset_token: Optional[bool] = None,
                 token: Optional[str] = None):
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if reset_token is not None:
            pulumi.set(__self__, "reset_token", reset_token)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[str]:
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="resetToken")
    def reset_token(self) -> Optional[bool]:
        return pulumi.get(self, "reset_token")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        return pulumi.get(self, "token")


@pulumi.output_type
class ScalingPlanHostPool(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hostpoolId":
            suggest = "hostpool_id"
        elif key == "scalingPlanEnabled":
            suggest = "scaling_plan_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalingPlanHostPool. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalingPlanHostPool.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalingPlanHostPool.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hostpool_id: str,
                 scaling_plan_enabled: bool):
        """
        :param str hostpool_id: The ID of the HostPool to assign the Scaling Plan to.
        :param bool scaling_plan_enabled: Specifies if the scaling plan is enabled or disabled for the HostPool.
        """
        pulumi.set(__self__, "hostpool_id", hostpool_id)
        pulumi.set(__self__, "scaling_plan_enabled", scaling_plan_enabled)

    @property
    @pulumi.getter(name="hostpoolId")
    def hostpool_id(self) -> str:
        """
        The ID of the HostPool to assign the Scaling Plan to.
        """
        return pulumi.get(self, "hostpool_id")

    @property
    @pulumi.getter(name="scalingPlanEnabled")
    def scaling_plan_enabled(self) -> bool:
        """
        Specifies if the scaling plan is enabled or disabled for the HostPool.
        """
        return pulumi.get(self, "scaling_plan_enabled")


@pulumi.output_type
class ScalingPlanSchedule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "daysOfWeeks":
            suggest = "days_of_weeks"
        elif key == "offPeakLoadBalancingAlgorithm":
            suggest = "off_peak_load_balancing_algorithm"
        elif key == "offPeakStartTime":
            suggest = "off_peak_start_time"
        elif key == "peakLoadBalancingAlgorithm":
            suggest = "peak_load_balancing_algorithm"
        elif key == "peakStartTime":
            suggest = "peak_start_time"
        elif key == "rampDownCapacityThresholdPercent":
            suggest = "ramp_down_capacity_threshold_percent"
        elif key == "rampDownForceLogoffUsers":
            suggest = "ramp_down_force_logoff_users"
        elif key == "rampDownLoadBalancingAlgorithm":
            suggest = "ramp_down_load_balancing_algorithm"
        elif key == "rampDownMinimumHostsPercent":
            suggest = "ramp_down_minimum_hosts_percent"
        elif key == "rampDownNotificationMessage":
            suggest = "ramp_down_notification_message"
        elif key == "rampDownStartTime":
            suggest = "ramp_down_start_time"
        elif key == "rampDownStopHostsWhen":
            suggest = "ramp_down_stop_hosts_when"
        elif key == "rampDownWaitTimeMinutes":
            suggest = "ramp_down_wait_time_minutes"
        elif key == "rampUpLoadBalancingAlgorithm":
            suggest = "ramp_up_load_balancing_algorithm"
        elif key == "rampUpStartTime":
            suggest = "ramp_up_start_time"
        elif key == "rampUpCapacityThresholdPercent":
            suggest = "ramp_up_capacity_threshold_percent"
        elif key == "rampUpMinimumHostsPercent":
            suggest = "ramp_up_minimum_hosts_percent"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScalingPlanSchedule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScalingPlanSchedule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScalingPlanSchedule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 days_of_weeks: Sequence[str],
                 name: str,
                 off_peak_load_balancing_algorithm: str,
                 off_peak_start_time: str,
                 peak_load_balancing_algorithm: str,
                 peak_start_time: str,
                 ramp_down_capacity_threshold_percent: int,
                 ramp_down_force_logoff_users: bool,
                 ramp_down_load_balancing_algorithm: str,
                 ramp_down_minimum_hosts_percent: int,
                 ramp_down_notification_message: str,
                 ramp_down_start_time: str,
                 ramp_down_stop_hosts_when: str,
                 ramp_down_wait_time_minutes: int,
                 ramp_up_load_balancing_algorithm: str,
                 ramp_up_start_time: str,
                 ramp_up_capacity_threshold_percent: Optional[int] = None,
                 ramp_up_minimum_hosts_percent: Optional[int] = None):
        """
        :param Sequence[str] days_of_weeks: A list of Days of the Week on which this schedule will be used..Possible values are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`
        :param str name: The name of the schedule.
        :param str off_peak_load_balancing_algorithm: The load Balancing Algorithm to use during Off-Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        :param str off_peak_start_time: The time at which Off-Peak scaling will begin. This is also the end-time for the Ramp-Down period. The time must be specified in "HH:MM" format.
        :param str peak_load_balancing_algorithm: The load Balancing Algorithm to use during Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        :param str peak_start_time: The time at which Peak scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param int ramp_down_capacity_threshold_percent: This is the value in percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-down and off-peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        :param bool ramp_down_force_logoff_users: Whether users will be forced to log-off session hosts once the `ramp_down_wait_time_minutes` value has been exceeded during the Ramp-Down period. Possible
        :param str ramp_down_load_balancing_algorithm: The load Balancing Algorithm to use during the Ramp-Down period. Possible values are `DepthFirst` and `BreadthFirst`.
        :param int ramp_down_minimum_hosts_percent: The minimum percentage of session host virtual machines that you would like to get to for ramp-down and off-peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        :param str ramp_down_notification_message: The notification message to send to users during Ramp-Down period when they are required to log-off.
        :param str ramp_down_start_time: The time at which Ramp-Down scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param str ramp_down_stop_hosts_when: Controls Session Host shutdown behaviour during Ramp-Down period. Session Hosts can either be shutdown when all sessions on the Session Host have ended, or when there are no Active sessions left on the Session Host. Possible values are `ZeroSessions` and `ZeroActiveSessions`.
        :param int ramp_down_wait_time_minutes: The number of minutes during Ramp-Down period that autoscale will wait after setting the session host VMs to drain mode, notifying any currently signed in users to save their work before forcing the users to logoff. Once all user sessions on the session host VM have been logged off, Autoscale will shut down the VM.
        :param str ramp_up_load_balancing_algorithm: The load Balancing Algorithm to use during the Ramp-Up period. Possible values are `DepthFirst` and `BreadthFirst`.
        :param str ramp_up_start_time: The time at which Ramp-Up scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        :param int ramp_up_capacity_threshold_percent: Specify minimum percentage of session host virtual machines to start for ramp-up and peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        :param int ramp_up_minimum_hosts_percent: This is the value of percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-up and peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
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
    def days_of_weeks(self) -> Sequence[str]:
        """
        A list of Days of the Week on which this schedule will be used..Possible values are `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, and `Sunday`
        """
        return pulumi.get(self, "days_of_weeks")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the schedule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="offPeakLoadBalancingAlgorithm")
    def off_peak_load_balancing_algorithm(self) -> str:
        """
        The load Balancing Algorithm to use during Off-Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "off_peak_load_balancing_algorithm")

    @property
    @pulumi.getter(name="offPeakStartTime")
    def off_peak_start_time(self) -> str:
        """
        The time at which Off-Peak scaling will begin. This is also the end-time for the Ramp-Down period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "off_peak_start_time")

    @property
    @pulumi.getter(name="peakLoadBalancingAlgorithm")
    def peak_load_balancing_algorithm(self) -> str:
        """
        The load Balancing Algorithm to use during Peak Hours. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "peak_load_balancing_algorithm")

    @property
    @pulumi.getter(name="peakStartTime")
    def peak_start_time(self) -> str:
        """
        The time at which Peak scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "peak_start_time")

    @property
    @pulumi.getter(name="rampDownCapacityThresholdPercent")
    def ramp_down_capacity_threshold_percent(self) -> int:
        """
        This is the value in percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-down and off-peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        """
        return pulumi.get(self, "ramp_down_capacity_threshold_percent")

    @property
    @pulumi.getter(name="rampDownForceLogoffUsers")
    def ramp_down_force_logoff_users(self) -> bool:
        """
        Whether users will be forced to log-off session hosts once the `ramp_down_wait_time_minutes` value has been exceeded during the Ramp-Down period. Possible
        """
        return pulumi.get(self, "ramp_down_force_logoff_users")

    @property
    @pulumi.getter(name="rampDownLoadBalancingAlgorithm")
    def ramp_down_load_balancing_algorithm(self) -> str:
        """
        The load Balancing Algorithm to use during the Ramp-Down period. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "ramp_down_load_balancing_algorithm")

    @property
    @pulumi.getter(name="rampDownMinimumHostsPercent")
    def ramp_down_minimum_hosts_percent(self) -> int:
        """
        The minimum percentage of session host virtual machines that you would like to get to for ramp-down and off-peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        """
        return pulumi.get(self, "ramp_down_minimum_hosts_percent")

    @property
    @pulumi.getter(name="rampDownNotificationMessage")
    def ramp_down_notification_message(self) -> str:
        """
        The notification message to send to users during Ramp-Down period when they are required to log-off.
        """
        return pulumi.get(self, "ramp_down_notification_message")

    @property
    @pulumi.getter(name="rampDownStartTime")
    def ramp_down_start_time(self) -> str:
        """
        The time at which Ramp-Down scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "ramp_down_start_time")

    @property
    @pulumi.getter(name="rampDownStopHostsWhen")
    def ramp_down_stop_hosts_when(self) -> str:
        """
        Controls Session Host shutdown behaviour during Ramp-Down period. Session Hosts can either be shutdown when all sessions on the Session Host have ended, or when there are no Active sessions left on the Session Host. Possible values are `ZeroSessions` and `ZeroActiveSessions`.
        """
        return pulumi.get(self, "ramp_down_stop_hosts_when")

    @property
    @pulumi.getter(name="rampDownWaitTimeMinutes")
    def ramp_down_wait_time_minutes(self) -> int:
        """
        The number of minutes during Ramp-Down period that autoscale will wait after setting the session host VMs to drain mode, notifying any currently signed in users to save their work before forcing the users to logoff. Once all user sessions on the session host VM have been logged off, Autoscale will shut down the VM.
        """
        return pulumi.get(self, "ramp_down_wait_time_minutes")

    @property
    @pulumi.getter(name="rampUpLoadBalancingAlgorithm")
    def ramp_up_load_balancing_algorithm(self) -> str:
        """
        The load Balancing Algorithm to use during the Ramp-Up period. Possible values are `DepthFirst` and `BreadthFirst`.
        """
        return pulumi.get(self, "ramp_up_load_balancing_algorithm")

    @property
    @pulumi.getter(name="rampUpStartTime")
    def ramp_up_start_time(self) -> str:
        """
        The time at which Ramp-Up scaling will begin. This is also the end-time for the Ramp-Up period. The time must be specified in "HH:MM" format.
        """
        return pulumi.get(self, "ramp_up_start_time")

    @property
    @pulumi.getter(name="rampUpCapacityThresholdPercent")
    def ramp_up_capacity_threshold_percent(self) -> Optional[int]:
        """
        Specify minimum percentage of session host virtual machines to start for ramp-up and peak hours. For example, if Minimum percentage of hosts is specified as 10% and total number of session hosts in your host pool is 10, autoscale will ensure a minimum of 1 session host is available to take user connections.
        """
        return pulumi.get(self, "ramp_up_capacity_threshold_percent")

    @property
    @pulumi.getter(name="rampUpMinimumHostsPercent")
    def ramp_up_minimum_hosts_percent(self) -> Optional[int]:
        """
        This is the value of percentage of used host pool capacity that will be considered to evaluate whether to turn on/off virtual machines during the ramp-up and peak hours. For example, if capacity threshold is specified as 60% and your total host pool capacity is 100 sessions, autoscale will turn on additional session hosts once the host pool exceeds a load of 60 sessions.
        """
        return pulumi.get(self, "ramp_up_minimum_hosts_percent")


