# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DatabaseExtendedAuditingPolicyArgs',
    'DatabaseImportArgs',
    'DatabaseThreatDetectionPolicyArgs',
    'FailoverGroupPartnerServerArgs',
    'FailoverGroupReadWriteEndpointFailoverPolicyArgs',
    'FailoverGroupReadonlyEndpointFailoverPolicyArgs',
    'ManagedInstanceIdentityArgs',
    'SqlServerExtendedAuditingPolicyArgs',
    'SqlServerIdentityArgs',
    'SqlServerThreatDetectionPolicyArgs',
]

@pulumi.input_type
class DatabaseExtendedAuditingPolicyArgs:
    def __init__(__self__, *,
                 log_monitoring_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = None,
                 storage_endpoint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] retention_in_days: Specifies the number of days to retain logs for in the storage account.
        :param pulumi.Input[str] storage_account_access_key: Specifies the access key to use for the auditing storage account.
        :param pulumi.Input[bool] storage_account_access_key_is_secondary: Specifies whether `storage_account_access_key` value is the storage's secondary key.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
        """
        if log_monitoring_enabled is not None:
            pulumi.set(__self__, "log_monitoring_enabled", log_monitoring_enabled)
        if retention_in_days is not None:
            pulumi.set(__self__, "retention_in_days", retention_in_days)
        if storage_account_access_key is not None:
            pulumi.set(__self__, "storage_account_access_key", storage_account_access_key)
        if storage_account_access_key_is_secondary is not None:
            pulumi.set(__self__, "storage_account_access_key_is_secondary", storage_account_access_key_is_secondary)
        if storage_endpoint is not None:
            pulumi.set(__self__, "storage_endpoint", storage_endpoint)

    @property
    @pulumi.getter(name="logMonitoringEnabled")
    def log_monitoring_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "log_monitoring_enabled")

    @log_monitoring_enabled.setter
    def log_monitoring_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_monitoring_enabled", value)

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days to retain logs for in the storage account.
        """
        return pulumi.get(self, "retention_in_days")

    @retention_in_days.setter
    def retention_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_days", value)

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the access key to use for the auditing storage account.
        """
        return pulumi.get(self, "storage_account_access_key")

    @storage_account_access_key.setter
    def storage_account_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_access_key", value)

    @property
    @pulumi.getter(name="storageAccountAccessKeyIsSecondary")
    def storage_account_access_key_is_secondary(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether `storage_account_access_key` value is the storage's secondary key.
        """
        return pulumi.get(self, "storage_account_access_key_is_secondary")

    @storage_account_access_key_is_secondary.setter
    def storage_account_access_key_is_secondary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "storage_account_access_key_is_secondary", value)

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
        """
        return pulumi.get(self, "storage_endpoint")

    @storage_endpoint.setter
    def storage_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_endpoint", value)


@pulumi.input_type
class DatabaseImportArgs:
    def __init__(__self__, *,
                 administrator_login: pulumi.Input[str],
                 administrator_login_password: pulumi.Input[str],
                 authentication_type: pulumi.Input[str],
                 storage_key: pulumi.Input[str],
                 storage_key_type: pulumi.Input[str],
                 storage_uri: pulumi.Input[str],
                 operation_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] administrator_login: Specifies the name of the SQL administrator.
        :param pulumi.Input[str] administrator_login_password: Specifies the password of the SQL administrator.
        :param pulumi.Input[str] authentication_type: Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
        :param pulumi.Input[str] storage_key: Specifies the access key for the storage account.
        :param pulumi.Input[str] storage_key_type: Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
        :param pulumi.Input[str] storage_uri: Specifies the blob URI of the .bacpac file.
        :param pulumi.Input[str] operation_mode: Specifies the type of import operation being performed. The only allowable value is `Import`.
        """
        pulumi.set(__self__, "administrator_login", administrator_login)
        pulumi.set(__self__, "administrator_login_password", administrator_login_password)
        pulumi.set(__self__, "authentication_type", authentication_type)
        pulumi.set(__self__, "storage_key", storage_key)
        pulumi.set(__self__, "storage_key_type", storage_key_type)
        pulumi.set(__self__, "storage_uri", storage_uri)
        if operation_mode is not None:
            pulumi.set(__self__, "operation_mode", operation_mode)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Input[str]:
        """
        Specifies the name of the SQL administrator.
        """
        return pulumi.get(self, "administrator_login")

    @administrator_login.setter
    def administrator_login(self, value: pulumi.Input[str]):
        pulumi.set(self, "administrator_login", value)

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> pulumi.Input[str]:
        """
        Specifies the password of the SQL administrator.
        """
        return pulumi.get(self, "administrator_login_password")

    @administrator_login_password.setter
    def administrator_login_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "administrator_login_password", value)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Input[str]:
        """
        Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
        """
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="storageKey")
    def storage_key(self) -> pulumi.Input[str]:
        """
        Specifies the access key for the storage account.
        """
        return pulumi.get(self, "storage_key")

    @storage_key.setter
    def storage_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_key", value)

    @property
    @pulumi.getter(name="storageKeyType")
    def storage_key_type(self) -> pulumi.Input[str]:
        """
        Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
        """
        return pulumi.get(self, "storage_key_type")

    @storage_key_type.setter
    def storage_key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_key_type", value)

    @property
    @pulumi.getter(name="storageUri")
    def storage_uri(self) -> pulumi.Input[str]:
        """
        Specifies the blob URI of the .bacpac file.
        """
        return pulumi.get(self, "storage_uri")

    @storage_uri.setter
    def storage_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_uri", value)

    @property
    @pulumi.getter(name="operationMode")
    def operation_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of import operation being performed. The only allowable value is `Import`.
        """
        return pulumi.get(self, "operation_mode")

    @operation_mode.setter
    def operation_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_mode", value)


@pulumi.input_type
class DatabaseThreatDetectionPolicyArgs:
    def __init__(__self__, *,
                 disabled_alerts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_account_admins: Optional[pulumi.Input[str]] = None,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_endpoint: Optional[pulumi.Input[str]] = None,
                 use_server_default: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disabled_alerts: Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
        :param pulumi.Input[str] email_account_admins: Should the account administrators be emailed when this alert is triggered?
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: A list of email addresses which alerts should be sent to.
        :param pulumi.Input[int] retention_days: Specifies the number of days to keep in the Threat Detection audit logs.
        :param pulumi.Input[str] state: The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        if disabled_alerts is not None:
            pulumi.set(__self__, "disabled_alerts", disabled_alerts)
        if email_account_admins is not None:
            pulumi.set(__self__, "email_account_admins", email_account_admins)
        if email_addresses is not None:
            pulumi.set(__self__, "email_addresses", email_addresses)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if storage_account_access_key is not None:
            pulumi.set(__self__, "storage_account_access_key", storage_account_access_key)
        if storage_endpoint is not None:
            pulumi.set(__self__, "storage_endpoint", storage_endpoint)
        if use_server_default is not None:
            warnings.warn("""This field is now non-functional and thus will be removed in version 3.0 of the Azure Provider""", DeprecationWarning)
            pulumi.log.warn("""use_server_default is deprecated: This field is now non-functional and thus will be removed in version 3.0 of the Azure Provider""")
        if use_server_default is not None:
            pulumi.set(__self__, "use_server_default", use_server_default)

    @property
    @pulumi.getter(name="disabledAlerts")
    def disabled_alerts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
        """
        return pulumi.get(self, "disabled_alerts")

    @disabled_alerts.setter
    def disabled_alerts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disabled_alerts", value)

    @property
    @pulumi.getter(name="emailAccountAdmins")
    def email_account_admins(self) -> Optional[pulumi.Input[str]]:
        """
        Should the account administrators be emailed when this alert is triggered?
        """
        return pulumi.get(self, "email_account_admins")

    @email_account_admins.setter
    def email_account_admins(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_account_admins", value)

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses which alerts should be sent to.
        """
        return pulumi.get(self, "email_addresses")

    @email_addresses.setter
    def email_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "email_addresses", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days to keep in the Threat Detection audit logs.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_account_access_key")

    @storage_account_access_key.setter
    def storage_account_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_access_key", value)

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_endpoint")

    @storage_endpoint.setter
    def storage_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_endpoint", value)

    @property
    @pulumi.getter(name="useServerDefault")
    def use_server_default(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "use_server_default")

    @use_server_default.setter
    def use_server_default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_server_default", value)


@pulumi.input_type
class FailoverGroupPartnerServerArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: the SQL server ID
        :param pulumi.Input[str] location: the location of the failover group.
        :param pulumi.Input[str] role: local replication role of the failover group instance.
        """
        pulumi.set(__self__, "id", id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        the SQL server ID
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        the location of the failover group.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        local replication role of the failover group instance.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class FailoverGroupReadWriteEndpointFailoverPolicyArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str],
                 grace_minutes: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] mode: the failover mode. Possible values are `Manual`, `Automatic`
        :param pulumi.Input[int] grace_minutes: Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
        """
        pulumi.set(__self__, "mode", mode)
        if grace_minutes is not None:
            pulumi.set(__self__, "grace_minutes", grace_minutes)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        the failover mode. Possible values are `Manual`, `Automatic`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="graceMinutes")
    def grace_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Applies only if `mode` is `Automatic`. The grace period in minutes before failover with data loss is attempted
        """
        return pulumi.get(self, "grace_minutes")

    @grace_minutes.setter
    def grace_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "grace_minutes", value)


@pulumi.input_type
class FailoverGroupReadonlyEndpointFailoverPolicyArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str]):
        """
        :param pulumi.Input[str] mode: Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
        """
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        Failover policy for the read-only endpoint. Possible values are `Enabled`, and `Disabled`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class ManagedInstanceIdentityArgs:
    def __init__(__self__, *,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] principal_id: The Principal ID for the Service Principal associated with the Identity of this SQL Managed Instance.
        :param pulumi.Input[str] tenant_id: The Tenant ID for the Service Principal associated with the Identity of this SQL Managed Instance.
        :param pulumi.Input[str] type: The identity type of the SQL Managed Instance. Only possible values is `SystemAssigned`.
        """
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Principal ID for the Service Principal associated with the Identity of this SQL Managed Instance.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID for the Service Principal associated with the Identity of this SQL Managed Instance.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type of the SQL Managed Instance. Only possible values is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SqlServerExtendedAuditingPolicyArgs:
    def __init__(__self__, *,
                 log_monitoring_enabled: Optional[pulumi.Input[bool]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_account_access_key_is_secondary: Optional[pulumi.Input[bool]] = None,
                 storage_endpoint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        if log_monitoring_enabled is not None:
            pulumi.set(__self__, "log_monitoring_enabled", log_monitoring_enabled)
        if retention_in_days is not None:
            pulumi.set(__self__, "retention_in_days", retention_in_days)
        if storage_account_access_key is not None:
            pulumi.set(__self__, "storage_account_access_key", storage_account_access_key)
        if storage_account_access_key_is_secondary is not None:
            pulumi.set(__self__, "storage_account_access_key_is_secondary", storage_account_access_key_is_secondary)
        if storage_endpoint is not None:
            pulumi.set(__self__, "storage_endpoint", storage_endpoint)

    @property
    @pulumi.getter(name="logMonitoringEnabled")
    def log_monitoring_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "log_monitoring_enabled")

    @log_monitoring_enabled.setter
    def log_monitoring_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_monitoring_enabled", value)

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "retention_in_days")

    @retention_in_days.setter
    def retention_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_days", value)

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_account_access_key")

    @storage_account_access_key.setter
    def storage_account_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_access_key", value)

    @property
    @pulumi.getter(name="storageAccountAccessKeyIsSecondary")
    def storage_account_access_key_is_secondary(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "storage_account_access_key_is_secondary")

    @storage_account_access_key_is_secondary.setter
    def storage_account_access_key_is_secondary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "storage_account_access_key_is_secondary", value)

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_endpoint")

    @storage_endpoint.setter
    def storage_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_endpoint", value)


@pulumi.input_type
class SqlServerIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
        :param pulumi.Input[str] principal_id: The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        :param pulumi.Input[str] tenant_id: The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the identity type of the Microsoft SQL Server. At this time the only allowed value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class SqlServerThreatDetectionPolicyArgs:
    def __init__(__self__, *,
                 disabled_alerts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_account_admins: Optional[pulumi.Input[bool]] = None,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_endpoint: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disabled_alerts: Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Data_Exfiltration`, `Sql_Injection`, `Sql_Injection_Vulnerability` and `Unsafe_Action"`,.
        :param pulumi.Input[bool] email_account_admins: Should the account administrators be emailed when this alert is triggered?
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: A list of email addresses which alerts should be sent to.
        :param pulumi.Input[int] retention_days: Specifies the number of days to keep in the Threat Detection audit logs.
        :param pulumi.Input[str] state: The State of the Policy. Possible values are `Enabled` or `Disabled`.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        if disabled_alerts is not None:
            pulumi.set(__self__, "disabled_alerts", disabled_alerts)
        if email_account_admins is not None:
            pulumi.set(__self__, "email_account_admins", email_account_admins)
        if email_addresses is not None:
            pulumi.set(__self__, "email_addresses", email_addresses)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if storage_account_access_key is not None:
            pulumi.set(__self__, "storage_account_access_key", storage_account_access_key)
        if storage_endpoint is not None:
            pulumi.set(__self__, "storage_endpoint", storage_endpoint)

    @property
    @pulumi.getter(name="disabledAlerts")
    def disabled_alerts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Data_Exfiltration`, `Sql_Injection`, `Sql_Injection_Vulnerability` and `Unsafe_Action"`,.
        """
        return pulumi.get(self, "disabled_alerts")

    @disabled_alerts.setter
    def disabled_alerts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disabled_alerts", value)

    @property
    @pulumi.getter(name="emailAccountAdmins")
    def email_account_admins(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the account administrators be emailed when this alert is triggered?
        """
        return pulumi.get(self, "email_account_admins")

    @email_account_admins.setter
    def email_account_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email_account_admins", value)

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of email addresses which alerts should be sent to.
        """
        return pulumi.get(self, "email_addresses")

    @email_addresses.setter
    def email_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "email_addresses", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days to keep in the Threat Detection audit logs.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The State of the Policy. Possible values are `Enabled` or `Disabled`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_account_access_key")

    @storage_account_access_key.setter
    def storage_account_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_access_key", value)

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        """
        return pulumi.get(self, "storage_endpoint")

    @storage_endpoint.setter
    def storage_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_endpoint", value)


