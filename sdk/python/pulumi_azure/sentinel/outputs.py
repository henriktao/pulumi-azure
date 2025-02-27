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
    'AlertRuleScheduledEventGrouping',
    'AlertRuleScheduledIncidentConfiguration',
    'AlertRuleScheduledIncidentConfigurationGrouping',
    'AuthomationRuleActionIncident',
    'AuthomationRuleActionPlaybook',
    'AuthomationRuleCondition',
    'AutomationRuleActionIncident',
    'AutomationRuleActionPlaybook',
    'AutomationRuleCondition',
    'GetAlertRuleTemplateScheduledTemplateResult',
    'GetAlertRuleTemplateSecurityIncidentTemplateResult',
]

@pulumi.output_type
class AlertRuleScheduledEventGrouping(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "aggregationMethod":
            suggest = "aggregation_method"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AlertRuleScheduledEventGrouping. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AlertRuleScheduledEventGrouping.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AlertRuleScheduledEventGrouping.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 aggregation_method: str):
        """
        :param str aggregation_method: The aggregation type of grouping the events.
        """
        pulumi.set(__self__, "aggregation_method", aggregation_method)

    @property
    @pulumi.getter(name="aggregationMethod")
    def aggregation_method(self) -> str:
        """
        The aggregation type of grouping the events.
        """
        return pulumi.get(self, "aggregation_method")


@pulumi.output_type
class AlertRuleScheduledIncidentConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createIncident":
            suggest = "create_incident"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AlertRuleScheduledIncidentConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AlertRuleScheduledIncidentConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AlertRuleScheduledIncidentConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 create_incident: bool,
                 grouping: 'outputs.AlertRuleScheduledIncidentConfigurationGrouping'):
        """
        :param bool create_incident: Whether to create an incident from alerts triggered by this Sentinel Scheduled Alert Rule?
        :param 'AlertRuleScheduledIncidentConfigurationGroupingArgs' grouping: A `grouping` block as defined below.
        """
        pulumi.set(__self__, "create_incident", create_incident)
        pulumi.set(__self__, "grouping", grouping)

    @property
    @pulumi.getter(name="createIncident")
    def create_incident(self) -> bool:
        """
        Whether to create an incident from alerts triggered by this Sentinel Scheduled Alert Rule?
        """
        return pulumi.get(self, "create_incident")

    @property
    @pulumi.getter
    def grouping(self) -> 'outputs.AlertRuleScheduledIncidentConfigurationGrouping':
        """
        A `grouping` block as defined below.
        """
        return pulumi.get(self, "grouping")


@pulumi.output_type
class AlertRuleScheduledIncidentConfigurationGrouping(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "entityMatchingMethod":
            suggest = "entity_matching_method"
        elif key == "groupBies":
            suggest = "group_bies"
        elif key == "lookbackDuration":
            suggest = "lookback_duration"
        elif key == "reopenClosedIncidents":
            suggest = "reopen_closed_incidents"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AlertRuleScheduledIncidentConfigurationGrouping. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AlertRuleScheduledIncidentConfigurationGrouping.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AlertRuleScheduledIncidentConfigurationGrouping.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 entity_matching_method: Optional[str] = None,
                 group_bies: Optional[Sequence[str]] = None,
                 lookback_duration: Optional[str] = None,
                 reopen_closed_incidents: Optional[bool] = None):
        """
        :param bool enabled: Enable grouping incidents created from alerts triggered by this Sentinel Scheduled Alert Rule. Defaults to `true`.
        :param str entity_matching_method: The method used to group incidents. Possible values are `All`, `Custom` and `None`. Defaults to `None`.
        :param Sequence[str] group_bies: A list of entity types to group by, only when the `entity_matching_method` is `Custom`. Possible values are `Account`, `Host`, `Url`, `Ip`.
        :param str lookback_duration: Limit the group to alerts created within the lookback duration (in ISO 8601 duration format). Defaults to `PT5M`.
        :param bool reopen_closed_incidents: Whether to re-open closed matching incidents? Defaults to `false`.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if entity_matching_method is not None:
            pulumi.set(__self__, "entity_matching_method", entity_matching_method)
        if group_bies is not None:
            pulumi.set(__self__, "group_bies", group_bies)
        if lookback_duration is not None:
            pulumi.set(__self__, "lookback_duration", lookback_duration)
        if reopen_closed_incidents is not None:
            pulumi.set(__self__, "reopen_closed_incidents", reopen_closed_incidents)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Enable grouping incidents created from alerts triggered by this Sentinel Scheduled Alert Rule. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="entityMatchingMethod")
    def entity_matching_method(self) -> Optional[str]:
        """
        The method used to group incidents. Possible values are `All`, `Custom` and `None`. Defaults to `None`.
        """
        return pulumi.get(self, "entity_matching_method")

    @property
    @pulumi.getter(name="groupBies")
    def group_bies(self) -> Optional[Sequence[str]]:
        """
        A list of entity types to group by, only when the `entity_matching_method` is `Custom`. Possible values are `Account`, `Host`, `Url`, `Ip`.
        """
        return pulumi.get(self, "group_bies")

    @property
    @pulumi.getter(name="lookbackDuration")
    def lookback_duration(self) -> Optional[str]:
        """
        Limit the group to alerts created within the lookback duration (in ISO 8601 duration format). Defaults to `PT5M`.
        """
        return pulumi.get(self, "lookback_duration")

    @property
    @pulumi.getter(name="reopenClosedIncidents")
    def reopen_closed_incidents(self) -> Optional[bool]:
        """
        Whether to re-open closed matching incidents? Defaults to `false`.
        """
        return pulumi.get(self, "reopen_closed_incidents")


@pulumi.output_type
class AuthomationRuleActionIncident(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "classificationComment":
            suggest = "classification_comment"
        elif key == "ownerId":
            suggest = "owner_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuthomationRuleActionIncident. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuthomationRuleActionIncident.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuthomationRuleActionIncident.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 order: int,
                 classification: Optional[str] = None,
                 classification_comment: Optional[str] = None,
                 labels: Optional[Sequence[str]] = None,
                 owner_id: Optional[str] = None,
                 severity: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param int order: The execution order of this action.
        :param str classification: The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        :param str classification_comment: The comment why the incident is to be closed.
        :param Sequence[str] labels: Specifies a list of labels to add to the incident.
        :param str owner_id: The object ID of the entity this incident is assigned to.
        :param str severity: The severity to add to the incident.
        :param str status: The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        """
        pulumi.set(__self__, "order", order)
        if classification is not None:
            pulumi.set(__self__, "classification", classification)
        if classification_comment is not None:
            pulumi.set(__self__, "classification_comment", classification_comment)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def classification(self) -> Optional[str]:
        """
        The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        """
        return pulumi.get(self, "classification")

    @property
    @pulumi.getter(name="classificationComment")
    def classification_comment(self) -> Optional[str]:
        """
        The comment why the incident is to be closed.
        """
        return pulumi.get(self, "classification_comment")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[str]]:
        """
        Specifies a list of labels to add to the incident.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[str]:
        """
        The object ID of the entity this incident is assigned to.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def severity(self) -> Optional[str]:
        """
        The severity to add to the incident.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class AuthomationRuleActionPlaybook(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logicAppId":
            suggest = "logic_app_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuthomationRuleActionPlaybook. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuthomationRuleActionPlaybook.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuthomationRuleActionPlaybook.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 logic_app_id: str,
                 order: int,
                 tenant_id: Optional[str] = None):
        """
        :param str logic_app_id: The ID of the Logic App that defines the playbook's logic.
        :param int order: The execution order of this action.
        :param str tenant_id: The ID of the Tenant that owns the playbook.
        """
        pulumi.set(__self__, "logic_app_id", logic_app_id)
        pulumi.set(__self__, "order", order)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> str:
        """
        The ID of the Logic App that defines the playbook's logic.
        """
        return pulumi.get(self, "logic_app_id")

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The ID of the Tenant that owns the playbook.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class AuthomationRuleCondition(dict):
    def __init__(__self__, *,
                 operator: str,
                 property: str,
                 values: Sequence[str]):
        """
        :param str operator: The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        :param str property: The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        :param Sequence[str] values: Specifies a list of values to use for evaluate the condition.
        """
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "property", property)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def operator(self) -> str:
        """
        The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Specifies a list of values to use for evaluate the condition.
        """
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def property(self) -> str:
        """
        The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        """
        return pulumi.get(self, "property")


@pulumi.output_type
class AutomationRuleActionIncident(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "classificationComment":
            suggest = "classification_comment"
        elif key == "ownerId":
            suggest = "owner_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationRuleActionIncident. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationRuleActionIncident.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationRuleActionIncident.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 order: int,
                 classification: Optional[str] = None,
                 classification_comment: Optional[str] = None,
                 labels: Optional[Sequence[str]] = None,
                 owner_id: Optional[str] = None,
                 severity: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param int order: The execution order of this action.
        :param str classification: The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        :param str classification_comment: The comment why the incident is to be closed.
        :param Sequence[str] labels: Specifies a list of labels to add to the incident.
        :param str owner_id: The object ID of the entity this incident is assigned to.
        :param str severity: The severity to add to the incident.
        :param str status: The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        """
        pulumi.set(__self__, "order", order)
        if classification is not None:
            pulumi.set(__self__, "classification", classification)
        if classification_comment is not None:
            pulumi.set(__self__, "classification_comment", classification_comment)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def classification(self) -> Optional[str]:
        """
        The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        """
        return pulumi.get(self, "classification")

    @property
    @pulumi.getter(name="classificationComment")
    def classification_comment(self) -> Optional[str]:
        """
        The comment why the incident is to be closed.
        """
        return pulumi.get(self, "classification_comment")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[str]]:
        """
        Specifies a list of labels to add to the incident.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[str]:
        """
        The object ID of the entity this incident is assigned to.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def severity(self) -> Optional[str]:
        """
        The severity to add to the incident.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class AutomationRuleActionPlaybook(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logicAppId":
            suggest = "logic_app_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AutomationRuleActionPlaybook. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AutomationRuleActionPlaybook.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AutomationRuleActionPlaybook.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 logic_app_id: str,
                 order: int,
                 tenant_id: Optional[str] = None):
        """
        :param str logic_app_id: The ID of the Logic App that defines the playbook's logic.
        :param int order: The execution order of this action.
        :param str tenant_id: The ID of the Tenant that owns the playbook.
        """
        pulumi.set(__self__, "logic_app_id", logic_app_id)
        pulumi.set(__self__, "order", order)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> str:
        """
        The ID of the Logic App that defines the playbook's logic.
        """
        return pulumi.get(self, "logic_app_id")

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        The ID of the Tenant that owns the playbook.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class AutomationRuleCondition(dict):
    def __init__(__self__, *,
                 operator: str,
                 property: str,
                 values: Sequence[str]):
        """
        :param str operator: The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        :param str property: The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        :param Sequence[str] values: Specifies a list of values to use for evaluate the condition.
        """
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "property", property)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def operator(self) -> str:
        """
        The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Specifies a list of values to use for evaluate the condition.
        """
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def property(self) -> str:
        """
        The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        """
        return pulumi.get(self, "property")


@pulumi.output_type
class GetAlertRuleTemplateScheduledTemplateResult(dict):
    def __init__(__self__, *,
                 description: str,
                 query: str,
                 query_frequency: str,
                 query_period: str,
                 severity: str,
                 tactics: Sequence[str],
                 trigger_operator: str,
                 trigger_threshold: int):
        """
        :param str description: The description of this Sentinel Scheduled Alert Rule Template.
        :param str query: The query of this Sentinel Scheduled Alert Rule Template.
        :param str query_frequency: The ISO 8601 timespan duration between two consecutive queries.
        :param str query_period: The ISO 8601 timespan duration, which determine the time period of the data covered by the query.
        :param str severity: The alert severity of this Sentinel Scheduled Alert Rule Template.
        :param Sequence[str] tactics: A list of categories of attacks by which to classify the rule.
        :param str trigger_operator: The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule Template.
        :param int trigger_threshold: The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule Template.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "query", query)
        pulumi.set(__self__, "query_frequency", query_frequency)
        pulumi.set(__self__, "query_period", query_period)
        pulumi.set(__self__, "severity", severity)
        pulumi.set(__self__, "tactics", tactics)
        pulumi.set(__self__, "trigger_operator", trigger_operator)
        pulumi.set(__self__, "trigger_threshold", trigger_threshold)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        The query of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryFrequency")
    def query_frequency(self) -> str:
        """
        The ISO 8601 timespan duration between two consecutive queries.
        """
        return pulumi.get(self, "query_frequency")

    @property
    @pulumi.getter(name="queryPeriod")
    def query_period(self) -> str:
        """
        The ISO 8601 timespan duration, which determine the time period of the data covered by the query.
        """
        return pulumi.get(self, "query_period")

    @property
    @pulumi.getter
    def severity(self) -> str:
        """
        The alert severity of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tactics(self) -> Sequence[str]:
        """
        A list of categories of attacks by which to classify the rule.
        """
        return pulumi.get(self, "tactics")

    @property
    @pulumi.getter(name="triggerOperator")
    def trigger_operator(self) -> str:
        """
        The alert trigger operator, combined with `trigger_threshold`, setting alert threshold of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "trigger_operator")

    @property
    @pulumi.getter(name="triggerThreshold")
    def trigger_threshold(self) -> int:
        """
        The baseline number of query results generated, combined with `trigger_operator`, setting alert threshold of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "trigger_threshold")


@pulumi.output_type
class GetAlertRuleTemplateSecurityIncidentTemplateResult(dict):
    def __init__(__self__, *,
                 description: str,
                 product_filter: str):
        """
        :param str description: The description of this Sentinel Scheduled Alert Rule Template.
        :param str product_filter: The Microsoft Security Service from where the alert will be generated.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "product_filter", product_filter)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of this Sentinel Scheduled Alert Rule Template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="productFilter")
    def product_filter(self) -> str:
        """
        The Microsoft Security Service from where the alert will be generated.
        """
        return pulumi.get(self, "product_filter")


