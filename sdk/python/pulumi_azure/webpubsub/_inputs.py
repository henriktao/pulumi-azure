# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'HubEventHandlerArgs',
    'HubEventHandlerAuthArgs',
    'NetworkAclPrivateEndpointArgs',
    'NetworkAclPublicNetworkArgs',
    'ServiceLiveTraceArgs',
]

@pulumi.input_type
class HubEventHandlerArgs:
    def __init__(__self__, *,
                 url_template: pulumi.Input[str],
                 auth: Optional[pulumi.Input['HubEventHandlerAuthArgs']] = None,
                 system_events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_event_pattern: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] url_template: The Event Handler URL Template. Two predefined parameters `{hub}` and `{event}` are
               available to use in the template. The value of the EventHandler URL is dynamically calculated when the client request
               comes in. Example: `http://example.com/api/{hub}/{event}`.
        :param pulumi.Input['HubEventHandlerAuthArgs'] auth: An `auth` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] system_events: Specify the list of system events. Supported values are `connect`, `connected`
               and `disconnected`.
        :param pulumi.Input[str] user_event_pattern: Specify the matching event names. There are 3 kind of patterns supported:
               - `*` matches any event name
               - `,` Combine multiple events with `,` for example `event1,event2`, it matches event `event1` and `event2`
               - The single event name, for example `event1`, it matches `event1`.
        """
        pulumi.set(__self__, "url_template", url_template)
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if system_events is not None:
            pulumi.set(__self__, "system_events", system_events)
        if user_event_pattern is not None:
            pulumi.set(__self__, "user_event_pattern", user_event_pattern)

    @property
    @pulumi.getter(name="urlTemplate")
    def url_template(self) -> pulumi.Input[str]:
        """
        The Event Handler URL Template. Two predefined parameters `{hub}` and `{event}` are
        available to use in the template. The value of the EventHandler URL is dynamically calculated when the client request
        comes in. Example: `http://example.com/api/{hub}/{event}`.
        """
        return pulumi.get(self, "url_template")

    @url_template.setter
    def url_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_template", value)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input['HubEventHandlerAuthArgs']]:
        """
        An `auth` block as defined below.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input['HubEventHandlerAuthArgs']]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="systemEvents")
    def system_events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specify the list of system events. Supported values are `connect`, `connected`
        and `disconnected`.
        """
        return pulumi.get(self, "system_events")

    @system_events.setter
    def system_events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "system_events", value)

    @property
    @pulumi.getter(name="userEventPattern")
    def user_event_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the matching event names. There are 3 kind of patterns supported:
        - `*` matches any event name
        - `,` Combine multiple events with `,` for example `event1,event2`, it matches event `event1` and `event2`
        - The single event name, for example `event1`, it matches `event1`.
        """
        return pulumi.get(self, "user_event_pattern")

    @user_event_pattern.setter
    def user_event_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_event_pattern", value)


@pulumi.input_type
class HubEventHandlerAuthArgs:
    def __init__(__self__, *,
                 managed_identity_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] managed_identity_id: Specify the identity ID of the target resource.
        """
        pulumi.set(__self__, "managed_identity_id", managed_identity_id)

    @property
    @pulumi.getter(name="managedIdentityId")
    def managed_identity_id(self) -> pulumi.Input[str]:
        """
        Specify the identity ID of the target resource.
        """
        return pulumi.get(self, "managed_identity_id")

    @managed_identity_id.setter
    def managed_identity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_identity_id", value)


@pulumi.input_type
class NetworkAclPrivateEndpointArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 allowed_request_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 denied_request_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] id: The ID of the Private Endpoint which is based on the Web Pubsub service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_request_types: The allowed request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] denied_request_types: The denied request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        pulumi.set(__self__, "id", id)
        if allowed_request_types is not None:
            pulumi.set(__self__, "allowed_request_types", allowed_request_types)
        if denied_request_types is not None:
            pulumi.set(__self__, "denied_request_types", denied_request_types)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of the Private Endpoint which is based on the Web Pubsub service.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="allowedRequestTypes")
    def allowed_request_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The allowed request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "allowed_request_types")

    @allowed_request_types.setter
    def allowed_request_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_request_types", value)

    @property
    @pulumi.getter(name="deniedRequestTypes")
    def denied_request_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The denied request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "denied_request_types")

    @denied_request_types.setter
    def denied_request_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "denied_request_types", value)


@pulumi.input_type
class NetworkAclPublicNetworkArgs:
    def __init__(__self__, *,
                 allowed_request_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 denied_request_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_request_types: The allowed request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] denied_request_types: The denied request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        if allowed_request_types is not None:
            pulumi.set(__self__, "allowed_request_types", allowed_request_types)
        if denied_request_types is not None:
            pulumi.set(__self__, "denied_request_types", denied_request_types)

    @property
    @pulumi.getter(name="allowedRequestTypes")
    def allowed_request_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The allowed request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "allowed_request_types")

    @allowed_request_types.setter
    def allowed_request_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_request_types", value)

    @property
    @pulumi.getter(name="deniedRequestTypes")
    def denied_request_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The denied request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "denied_request_types")

    @denied_request_types.setter
    def denied_request_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "denied_request_types", value)


@pulumi.input_type
class ServiceLiveTraceArgs:
    def __init__(__self__, *,
                 connectivity_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 http_request_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 messaging_logs_enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] connectivity_logs_enabled: Whether the log category `ConnectivityLogs` is enabled? Defaults to `true`
        :param pulumi.Input[bool] enabled: Whether the live trace is enabled? Defaults to `true`.
        :param pulumi.Input[bool] http_request_logs_enabled: Whether the log category `HttpRequestLogs` is enabled? Defaults to `true`
        :param pulumi.Input[bool] messaging_logs_enabled: Whether the log category `MessagingLogs` is enabled? Defaults to `true`
        """
        if connectivity_logs_enabled is not None:
            pulumi.set(__self__, "connectivity_logs_enabled", connectivity_logs_enabled)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if http_request_logs_enabled is not None:
            pulumi.set(__self__, "http_request_logs_enabled", http_request_logs_enabled)
        if messaging_logs_enabled is not None:
            pulumi.set(__self__, "messaging_logs_enabled", messaging_logs_enabled)

    @property
    @pulumi.getter(name="connectivityLogsEnabled")
    def connectivity_logs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the log category `ConnectivityLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "connectivity_logs_enabled")

    @connectivity_logs_enabled.setter
    def connectivity_logs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "connectivity_logs_enabled", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the live trace is enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="httpRequestLogsEnabled")
    def http_request_logs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the log category `HttpRequestLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "http_request_logs_enabled")

    @http_request_logs_enabled.setter
    def http_request_logs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http_request_logs_enabled", value)

    @property
    @pulumi.getter(name="messagingLogsEnabled")
    def messaging_logs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the log category `MessagingLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "messaging_logs_enabled")

    @messaging_logs_enabled.setter
    def messaging_logs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "messaging_logs_enabled", value)


