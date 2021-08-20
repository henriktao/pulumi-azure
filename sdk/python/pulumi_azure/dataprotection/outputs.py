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
    'BackupPolicyDiskRetentionRule',
    'BackupPolicyDiskRetentionRuleCriteria',
    'BackupPolicyPostgresqlRetentionRule',
    'BackupPolicyPostgresqlRetentionRuleCriteria',
    'BackupVaultIdentity',
    'GetBackupVaultIdentityResult',
]

@pulumi.output_type
class BackupPolicyDiskRetentionRule(dict):
    def __init__(__self__, *,
                 criteria: 'outputs.BackupPolicyDiskRetentionRuleCriteria',
                 duration: str,
                 name: str,
                 priority: int):
        """
        :param 'BackupPolicyDiskRetentionRuleCriteriaArgs' criteria: A `criteria` block as defined below. Changing this forces a new Backup Policy Disk to be created.
        :param str duration: Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Disk to be created.
        :param str name: The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
        :param int priority: Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
        """
        pulumi.set(__self__, "criteria", criteria)
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter
    def criteria(self) -> 'outputs.BackupPolicyDiskRetentionRuleCriteria':
        """
        A `criteria` block as defined below. Changing this forces a new Backup Policy Disk to be created.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Disk to be created.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name which should be used for this retention rule. Changing this forces a new Backup Policy Disk to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Retention Tag priority. Changing this forces a new Backup Policy Disk to be created.
        """
        return pulumi.get(self, "priority")


@pulumi.output_type
class BackupPolicyDiskRetentionRuleCriteria(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "absoluteCriteria":
            suggest = "absolute_criteria"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BackupPolicyDiskRetentionRuleCriteria. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BackupPolicyDiskRetentionRuleCriteria.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BackupPolicyDiskRetentionRuleCriteria.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 absolute_criteria: Optional[str] = None):
        """
        :param str absolute_criteria: Possible values are `FirstOfDay` and `FirstOfWeek`. Changing this forces a new Backup Policy Disk to be created.
        """
        if absolute_criteria is not None:
            pulumi.set(__self__, "absolute_criteria", absolute_criteria)

    @property
    @pulumi.getter(name="absoluteCriteria")
    def absolute_criteria(self) -> Optional[str]:
        """
        Possible values are `FirstOfDay` and `FirstOfWeek`. Changing this forces a new Backup Policy Disk to be created.
        """
        return pulumi.get(self, "absolute_criteria")


@pulumi.output_type
class BackupPolicyPostgresqlRetentionRule(dict):
    def __init__(__self__, *,
                 criteria: 'outputs.BackupPolicyPostgresqlRetentionRuleCriteria',
                 duration: str,
                 name: str,
                 priority: int):
        """
        :param 'BackupPolicyPostgresqlRetentionRuleCriteriaArgs' criteria: A `criteria` block as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param str duration: Duration after which the backup is deleted. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param str name: The name which should be used for this retention rule. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param int priority: Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Postgre Sql to be created.
        """
        pulumi.set(__self__, "criteria", criteria)
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter
    def criteria(self) -> 'outputs.BackupPolicyPostgresqlRetentionRuleCriteria':
        """
        A `criteria` block as defined below. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        Duration after which the backup is deleted. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name which should be used for this retention rule. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Specifies the priority of the rule. The priority number must be unique for each rule. The lower the priority number, the higher the priority of the rule. Changing this forces a new Backup Policy Postgre Sql to be created.
        """
        return pulumi.get(self, "priority")


@pulumi.output_type
class BackupPolicyPostgresqlRetentionRuleCriteria(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "absoluteCriteria":
            suggest = "absolute_criteria"
        elif key == "daysOfWeeks":
            suggest = "days_of_weeks"
        elif key == "monthsOfYears":
            suggest = "months_of_years"
        elif key == "scheduledBackupTimes":
            suggest = "scheduled_backup_times"
        elif key == "weeksOfMonths":
            suggest = "weeks_of_months"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BackupPolicyPostgresqlRetentionRuleCriteria. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BackupPolicyPostgresqlRetentionRuleCriteria.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BackupPolicyPostgresqlRetentionRuleCriteria.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 absolute_criteria: Optional[str] = None,
                 days_of_weeks: Optional[Sequence[str]] = None,
                 months_of_years: Optional[Sequence[str]] = None,
                 scheduled_backup_times: Optional[Sequence[str]] = None,
                 weeks_of_months: Optional[Sequence[str]] = None):
        """
        :param str absolute_criteria: Possible values are `AllBackup`, `FirstOfDay`, `FirstOfWeek`, `FirstOfMonth` and `FirstOfYear`. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param Sequence[str] days_of_weeks: Possible values are `Monday`, `Tuesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param Sequence[str] months_of_years: Possible values are `January`, `February`, `March`, `April`, `May`, `June`, `July`, `August`, `September`, `October`, `November` and `December`. Changing this forces a new Backup Policy PostgreSQL to be created.
        :param Sequence[str] scheduled_backup_times: Specifies a list of backup times for backup in the `RFC3339` format. Changing this forces a new Backup Policy Postgre Sql to be created.
        :param Sequence[str] weeks_of_months: Possible values are `First`, `Second`, `Third`, `Fourth` and `Last`. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        if absolute_criteria is not None:
            pulumi.set(__self__, "absolute_criteria", absolute_criteria)
        if days_of_weeks is not None:
            pulumi.set(__self__, "days_of_weeks", days_of_weeks)
        if months_of_years is not None:
            pulumi.set(__self__, "months_of_years", months_of_years)
        if scheduled_backup_times is not None:
            pulumi.set(__self__, "scheduled_backup_times", scheduled_backup_times)
        if weeks_of_months is not None:
            pulumi.set(__self__, "weeks_of_months", weeks_of_months)

    @property
    @pulumi.getter(name="absoluteCriteria")
    def absolute_criteria(self) -> Optional[str]:
        """
        Possible values are `AllBackup`, `FirstOfDay`, `FirstOfWeek`, `FirstOfMonth` and `FirstOfYear`. These values mean the first successful backup of the day/week/month/year. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "absolute_criteria")

    @property
    @pulumi.getter(name="daysOfWeeks")
    def days_of_weeks(self) -> Optional[Sequence[str]]:
        """
        Possible values are `Monday`, `Tuesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "days_of_weeks")

    @property
    @pulumi.getter(name="monthsOfYears")
    def months_of_years(self) -> Optional[Sequence[str]]:
        """
        Possible values are `January`, `February`, `March`, `April`, `May`, `June`, `July`, `August`, `September`, `October`, `November` and `December`. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "months_of_years")

    @property
    @pulumi.getter(name="scheduledBackupTimes")
    def scheduled_backup_times(self) -> Optional[Sequence[str]]:
        """
        Specifies a list of backup times for backup in the `RFC3339` format. Changing this forces a new Backup Policy Postgre Sql to be created.
        """
        return pulumi.get(self, "scheduled_backup_times")

    @property
    @pulumi.getter(name="weeksOfMonths")
    def weeks_of_months(self) -> Optional[Sequence[str]]:
        """
        Possible values are `First`, `Second`, `Third`, `Fourth` and `Last`. Changing this forces a new Backup Policy PostgreSQL to be created.
        """
        return pulumi.get(self, "weeks_of_months")


@pulumi.output_type
class BackupVaultIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BackupVaultIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BackupVaultIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BackupVaultIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str principal_id: The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
        :param str tenant_id: The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
        :param str type: Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
        """
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Specifies the identity type of the Backup Vault. Possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetBackupVaultIdentityResult(dict):
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        :param str principal_id: The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
        :param str tenant_id: The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
        :param str type: Specifies the identity type of the Backup Vault.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the identity type of the Backup Vault.
        """
        return pulumi.get(self, "type")


