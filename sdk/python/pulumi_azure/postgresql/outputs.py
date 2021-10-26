# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FlexibleServerHighAvailability',
    'FlexibleServerMaintenanceWindow',
    'ServerIdentity',
    'ServerStorageProfile',
    'ServerThreatDetectionPolicy',
    'GetServerIdentityResult',
]

@pulumi.output_type
class FlexibleServerHighAvailability(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "standbyAvailabilityZone":
            suggest = "standby_availability_zone"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlexibleServerHighAvailability. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlexibleServerHighAvailability.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlexibleServerHighAvailability.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mode: str,
                 standby_availability_zone: Optional[str] = None):
        """
        :param str mode: The high availability mode for the PostgreSQL Flexible Server. The only possible value is `ZoneRedundant`.
        :param str standby_availability_zone: The Availability Zone of the standby Flexible Server. Possible values are `1`, `2` and `3`.
        """
        pulumi.set(__self__, "mode", mode)
        if standby_availability_zone is not None:
            pulumi.set(__self__, "standby_availability_zone", standby_availability_zone)

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        The high availability mode for the PostgreSQL Flexible Server. The only possible value is `ZoneRedundant`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="standbyAvailabilityZone")
    def standby_availability_zone(self) -> Optional[str]:
        """
        The Availability Zone of the standby Flexible Server. Possible values are `1`, `2` and `3`.
        """
        return pulumi.get(self, "standby_availability_zone")


@pulumi.output_type
class FlexibleServerMaintenanceWindow(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dayOfWeek":
            suggest = "day_of_week"
        elif key == "startHour":
            suggest = "start_hour"
        elif key == "startMinute":
            suggest = "start_minute"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlexibleServerMaintenanceWindow. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlexibleServerMaintenanceWindow.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlexibleServerMaintenanceWindow.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 day_of_week: Optional[int] = None,
                 start_hour: Optional[int] = None,
                 start_minute: Optional[int] = None):
        """
        :param int day_of_week: The day of week for maintenance window. Defaults to `0`.
        :param int start_hour: The day of week for maintenance window. Defaults to `0`.
        :param int start_minute: The start minute for maintenance window. Defaults to `0`.
        """
        if day_of_week is not None:
            pulumi.set(__self__, "day_of_week", day_of_week)
        if start_hour is not None:
            pulumi.set(__self__, "start_hour", start_hour)
        if start_minute is not None:
            pulumi.set(__self__, "start_minute", start_minute)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> Optional[int]:
        """
        The day of week for maintenance window. Defaults to `0`.
        """
        return pulumi.get(self, "day_of_week")

    @property
    @pulumi.getter(name="startHour")
    def start_hour(self) -> Optional[int]:
        """
        The day of week for maintenance window. Defaults to `0`.
        """
        return pulumi.get(self, "start_hour")

    @property
    @pulumi.getter(name="startMinute")
    def start_minute(self) -> Optional[int]:
        """
        The start minute for maintenance window. Defaults to `0`.
        """
        return pulumi.get(self, "start_minute")


@pulumi.output_type
class ServerIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Identity which should be used for this PostgreSQL Server. At this time the only possible value is `SystemAssigned`.
        :param str principal_id: The Client ID of the Service Principal assigned to this PostgreSQL Server.
        :param str tenant_id: The ID of the Tenant the Service Principal is assigned in.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Identity which should be used for this PostgreSQL Server. At this time the only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The Client ID of the Service Principal assigned to this PostgreSQL Server.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The ID of the Tenant the Service Principal is assigned in.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class ServerStorageProfile(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "autoGrow":
            suggest = "auto_grow"
        elif key == "backupRetentionDays":
            suggest = "backup_retention_days"
        elif key == "geoRedundantBackup":
            suggest = "geo_redundant_backup"
        elif key == "storageMb":
            suggest = "storage_mb"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerStorageProfile. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerStorageProfile.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerStorageProfile.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auto_grow: Optional[str] = None,
                 backup_retention_days: Optional[int] = None,
                 geo_redundant_backup: Optional[str] = None,
                 storage_mb: Optional[int] = None):
        """
        :param int backup_retention_days: Backup retention days for the server, supported values are between `7` and `35` days.
        :param int storage_mb: Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
        """
        if auto_grow is not None:
            pulumi.set(__self__, "auto_grow", auto_grow)
        if backup_retention_days is not None:
            pulumi.set(__self__, "backup_retention_days", backup_retention_days)
        if geo_redundant_backup is not None:
            pulumi.set(__self__, "geo_redundant_backup", geo_redundant_backup)
        if storage_mb is not None:
            pulumi.set(__self__, "storage_mb", storage_mb)

    @property
    @pulumi.getter(name="autoGrow")
    def auto_grow(self) -> Optional[str]:
        return pulumi.get(self, "auto_grow")

    @property
    @pulumi.getter(name="backupRetentionDays")
    def backup_retention_days(self) -> Optional[int]:
        """
        Backup retention days for the server, supported values are between `7` and `35` days.
        """
        return pulumi.get(self, "backup_retention_days")

    @property
    @pulumi.getter(name="geoRedundantBackup")
    def geo_redundant_backup(self) -> Optional[str]:
        return pulumi.get(self, "geo_redundant_backup")

    @property
    @pulumi.getter(name="storageMb")
    def storage_mb(self) -> Optional[int]:
        """
        Max storage allowed for a server. Possible values are between `5120` MB(5GB) and `1048576` MB(1TB) for the Basic SKU and between `5120` MB(5GB) and `16777216` MB(16TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/postgresql/servers/create#StorageProfile).
        """
        return pulumi.get(self, "storage_mb")


@pulumi.output_type
class ServerThreatDetectionPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "disabledAlerts":
            suggest = "disabled_alerts"
        elif key == "emailAccountAdmins":
            suggest = "email_account_admins"
        elif key == "emailAddresses":
            suggest = "email_addresses"
        elif key == "retentionDays":
            suggest = "retention_days"
        elif key == "storageAccountAccessKey":
            suggest = "storage_account_access_key"
        elif key == "storageEndpoint":
            suggest = "storage_endpoint"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerThreatDetectionPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerThreatDetectionPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerThreatDetectionPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disabled_alerts: Optional[Sequence[str]] = None,
                 email_account_admins: Optional[bool] = None,
                 email_addresses: Optional[Sequence[str]] = None,
                 enabled: Optional[bool] = None,
                 retention_days: Optional[int] = None,
                 storage_account_access_key: Optional[str] = None,
                 storage_endpoint: Optional[str] = None):
        """
        :param Sequence[str] disabled_alerts: Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
        :param bool email_account_admins: Should the account administrators be emailed when this alert is triggered?
        :param Sequence[str] email_addresses: A list of email addresses which alerts should be sent to.
        :param bool enabled: Is the policy enabled?
        :param int retention_days: Specifies the number of days to keep in the Threat Detection audit logs.
        :param str storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account.
        :param str storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        """
        if disabled_alerts is not None:
            pulumi.set(__self__, "disabled_alerts", disabled_alerts)
        if email_account_admins is not None:
            pulumi.set(__self__, "email_account_admins", email_account_admins)
        if email_addresses is not None:
            pulumi.set(__self__, "email_addresses", email_addresses)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if storage_account_access_key is not None:
            pulumi.set(__self__, "storage_account_access_key", storage_account_access_key)
        if storage_endpoint is not None:
            pulumi.set(__self__, "storage_endpoint", storage_endpoint)

    @property
    @pulumi.getter(name="disabledAlerts")
    def disabled_alerts(self) -> Optional[Sequence[str]]:
        """
        Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
        """
        return pulumi.get(self, "disabled_alerts")

    @property
    @pulumi.getter(name="emailAccountAdmins")
    def email_account_admins(self) -> Optional[bool]:
        """
        Should the account administrators be emailed when this alert is triggered?
        """
        return pulumi.get(self, "email_account_admins")

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> Optional[Sequence[str]]:
        """
        A list of email addresses which alerts should be sent to.
        """
        return pulumi.get(self, "email_addresses")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Is the policy enabled?
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[int]:
        """
        Specifies the number of days to keep in the Threat Detection audit logs.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> Optional[str]:
        """
        Specifies the identifier key of the Threat Detection audit storage account.
        """
        return pulumi.get(self, "storage_account_access_key")

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> Optional[str]:
        """
        Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
        """
        return pulumi.get(self, "storage_endpoint")


@pulumi.output_type
class GetServerIdentityResult(dict):
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        :param str principal_id: The ID of the System Managed Service Principal assigned to the PostgreSQL Server.
        :param str tenant_id: The ID of the Tenant of the System Managed Service Principal assigned to the PostgreSQL Server.
        :param str type: The identity type of the Managed Identity assigned to the PostgreSQL Server.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The ID of the System Managed Service Principal assigned to the PostgreSQL Server.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The ID of the Tenant of the System Managed Service Principal assigned to the PostgreSQL Server.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The identity type of the Managed Identity assigned to the PostgreSQL Server.
        """
        return pulumi.get(self, "type")


