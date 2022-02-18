# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'HubEventHandler',
    'HubEventHandlerAuth',
    'NetworkAclPrivateEndpoint',
    'NetworkAclPublicNetwork',
    'ServiceIdentity',
    'ServiceLiveTrace',
]

@pulumi.output_type
class HubEventHandler(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "urlTemplate":
            suggest = "url_template"
        elif key == "systemEvents":
            suggest = "system_events"
        elif key == "userEventPattern":
            suggest = "user_event_pattern"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HubEventHandler. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HubEventHandler.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HubEventHandler.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 url_template: str,
                 auth: Optional['outputs.HubEventHandlerAuth'] = None,
                 system_events: Optional[Sequence[str]] = None,
                 user_event_pattern: Optional[str] = None):
        """
        :param str url_template: The Event Handler URL Template. Two predefined parameters `{hub}` and `{event}` are
               available to use in the template. The value of the EventHandler URL is dynamically calculated when the client request
               comes in. Example: `http://example.com/api/{hub}/{event}`.
        :param 'HubEventHandlerAuthArgs' auth: An `auth` block as defined below.
        :param Sequence[str] system_events: Specify the list of system events. Supported values are `connect`, `connected`
               and `disconnected`.
        :param str user_event_pattern: Specify the matching event names. There are 3 kind of patterns supported:
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
    def url_template(self) -> str:
        """
        The Event Handler URL Template. Two predefined parameters `{hub}` and `{event}` are
        available to use in the template. The value of the EventHandler URL is dynamically calculated when the client request
        comes in. Example: `http://example.com/api/{hub}/{event}`.
        """
        return pulumi.get(self, "url_template")

    @property
    @pulumi.getter
    def auth(self) -> Optional['outputs.HubEventHandlerAuth']:
        """
        An `auth` block as defined below.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="systemEvents")
    def system_events(self) -> Optional[Sequence[str]]:
        """
        Specify the list of system events. Supported values are `connect`, `connected`
        and `disconnected`.
        """
        return pulumi.get(self, "system_events")

    @property
    @pulumi.getter(name="userEventPattern")
    def user_event_pattern(self) -> Optional[str]:
        """
        Specify the matching event names. There are 3 kind of patterns supported:
        - `*` matches any event name
        - `,` Combine multiple events with `,` for example `event1,event2`, it matches event `event1` and `event2`
        - The single event name, for example `event1`, it matches `event1`.
        """
        return pulumi.get(self, "user_event_pattern")


@pulumi.output_type
class HubEventHandlerAuth(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "managedIdentityId":
            suggest = "managed_identity_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HubEventHandlerAuth. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HubEventHandlerAuth.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HubEventHandlerAuth.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 managed_identity_id: str):
        """
        :param str managed_identity_id: Specify the identity ID of the target resource.
        """
        pulumi.set(__self__, "managed_identity_id", managed_identity_id)

    @property
    @pulumi.getter(name="managedIdentityId")
    def managed_identity_id(self) -> str:
        """
        Specify the identity ID of the target resource.
        """
        return pulumi.get(self, "managed_identity_id")


@pulumi.output_type
class NetworkAclPrivateEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedRequestTypes":
            suggest = "allowed_request_types"
        elif key == "deniedRequestTypes":
            suggest = "denied_request_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkAclPrivateEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkAclPrivateEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkAclPrivateEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 allowed_request_types: Optional[Sequence[str]] = None,
                 denied_request_types: Optional[Sequence[str]] = None):
        """
        :param str id: The ID of the Private Endpoint which is based on the Web Pubsub service.
        :param Sequence[str] allowed_request_types: The allowed request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        :param Sequence[str] denied_request_types: The denied request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        pulumi.set(__self__, "id", id)
        if allowed_request_types is not None:
            pulumi.set(__self__, "allowed_request_types", allowed_request_types)
        if denied_request_types is not None:
            pulumi.set(__self__, "denied_request_types", denied_request_types)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Private Endpoint which is based on the Web Pubsub service.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="allowedRequestTypes")
    def allowed_request_types(self) -> Optional[Sequence[str]]:
        """
        The allowed request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "allowed_request_types")

    @property
    @pulumi.getter(name="deniedRequestTypes")
    def denied_request_types(self) -> Optional[Sequence[str]]:
        """
        The denied request types for the Private Endpoint Connection. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "denied_request_types")


@pulumi.output_type
class NetworkAclPublicNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedRequestTypes":
            suggest = "allowed_request_types"
        elif key == "deniedRequestTypes":
            suggest = "denied_request_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkAclPublicNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkAclPublicNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkAclPublicNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed_request_types: Optional[Sequence[str]] = None,
                 denied_request_types: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] allowed_request_types: The allowed request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        :param Sequence[str] denied_request_types: The denied request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        if allowed_request_types is not None:
            pulumi.set(__self__, "allowed_request_types", allowed_request_types)
        if denied_request_types is not None:
            pulumi.set(__self__, "denied_request_types", denied_request_types)

    @property
    @pulumi.getter(name="allowedRequestTypes")
    def allowed_request_types(self) -> Optional[Sequence[str]]:
        """
        The allowed request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "allowed_request_types")

    @property
    @pulumi.getter(name="deniedRequestTypes")
    def denied_request_types(self) -> Optional[Sequence[str]]:
        """
        The denied request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        """
        return pulumi.get(self, "denied_request_types")


@pulumi.output_type
class ServiceIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "identityIds":
            suggest = "identity_ids"
        elif key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 identity_ids: Optional[Sequence[str]] = None,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The type of identity used for the Web PubSub service. Possible values are `SystemAssigned` and `UserAssigned`. If `UserAssigned` is set, a `user_assigned_identity_id` must be set as well.
        :param Sequence[str] identity_ids: A list of User Assigned Identity IDs which should be assigned to this Web PubSub service.
        """
        pulumi.set(__self__, "type", type)
        if identity_ids is not None:
            pulumi.set(__self__, "identity_ids", identity_ids)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of identity used for the Web PubSub service. Possible values are `SystemAssigned` and `UserAssigned`. If `UserAssigned` is set, a `user_assigned_identity_id` must be set as well.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="identityIds")
    def identity_ids(self) -> Optional[Sequence[str]]:
        """
        A list of User Assigned Identity IDs which should be assigned to this Web PubSub service.
        """
        return pulumi.get(self, "identity_ids")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class ServiceLiveTrace(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "connectivityLogsEnabled":
            suggest = "connectivity_logs_enabled"
        elif key == "httpRequestLogsEnabled":
            suggest = "http_request_logs_enabled"
        elif key == "messagingLogsEnabled":
            suggest = "messaging_logs_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceLiveTrace. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceLiveTrace.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceLiveTrace.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 connectivity_logs_enabled: Optional[bool] = None,
                 enabled: Optional[bool] = None,
                 http_request_logs_enabled: Optional[bool] = None,
                 messaging_logs_enabled: Optional[bool] = None):
        """
        :param bool connectivity_logs_enabled: Whether the log category `ConnectivityLogs` is enabled? Defaults to `true`
        :param bool enabled: Whether the live trace is enabled? Defaults to `true`.
        :param bool http_request_logs_enabled: Whether the log category `HttpRequestLogs` is enabled? Defaults to `true`
        :param bool messaging_logs_enabled: Whether the log category `MessagingLogs` is enabled? Defaults to `true`
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
    def connectivity_logs_enabled(self) -> Optional[bool]:
        """
        Whether the log category `ConnectivityLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "connectivity_logs_enabled")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether the live trace is enabled? Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="httpRequestLogsEnabled")
    def http_request_logs_enabled(self) -> Optional[bool]:
        """
        Whether the log category `HttpRequestLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "http_request_logs_enabled")

    @property
    @pulumi.getter(name="messagingLogsEnabled")
    def messaging_logs_enabled(self) -> Optional[bool]:
        """
        Whether the log category `MessagingLogs` is enabled? Defaults to `true`
        """
        return pulumi.get(self, "messaging_logs_enabled")


