# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AlertRuleScheduledEventGroupingArgs',
    'AlertRuleScheduledIncidentConfigurationArgs',
    'AlertRuleScheduledIncidentConfigurationGroupingArgs',
    'AuthomationRuleActionIncidentArgs',
    'AuthomationRuleActionPlaybookArgs',
    'AuthomationRuleConditionArgs',
]

@pulumi.input_type
class AlertRuleScheduledEventGroupingArgs:
    def __init__(__self__, *,
                 aggregation_method: pulumi.Input[str]):
        """
        :param pulumi.Input[str] aggregation_method: The aggregation type of grouping the events.
        """
        pulumi.set(__self__, "aggregation_method", aggregation_method)

    @property
    @pulumi.getter(name="aggregationMethod")
    def aggregation_method(self) -> pulumi.Input[str]:
        """
        The aggregation type of grouping the events.
        """
        return pulumi.get(self, "aggregation_method")

    @aggregation_method.setter
    def aggregation_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "aggregation_method", value)


@pulumi.input_type
class AlertRuleScheduledIncidentConfigurationArgs:
    def __init__(__self__, *,
                 create_incident: pulumi.Input[bool],
                 grouping: pulumi.Input['AlertRuleScheduledIncidentConfigurationGroupingArgs']):
        """
        :param pulumi.Input[bool] create_incident: Whether to create an incident from alerts triggered by this Sentinel Scheduled Alert Rule?
        :param pulumi.Input['AlertRuleScheduledIncidentConfigurationGroupingArgs'] grouping: A `grouping` block as defined below.
        """
        pulumi.set(__self__, "create_incident", create_incident)
        pulumi.set(__self__, "grouping", grouping)

    @property
    @pulumi.getter(name="createIncident")
    def create_incident(self) -> pulumi.Input[bool]:
        """
        Whether to create an incident from alerts triggered by this Sentinel Scheduled Alert Rule?
        """
        return pulumi.get(self, "create_incident")

    @create_incident.setter
    def create_incident(self, value: pulumi.Input[bool]):
        pulumi.set(self, "create_incident", value)

    @property
    @pulumi.getter
    def grouping(self) -> pulumi.Input['AlertRuleScheduledIncidentConfigurationGroupingArgs']:
        """
        A `grouping` block as defined below.
        """
        return pulumi.get(self, "grouping")

    @grouping.setter
    def grouping(self, value: pulumi.Input['AlertRuleScheduledIncidentConfigurationGroupingArgs']):
        pulumi.set(self, "grouping", value)


@pulumi.input_type
class AlertRuleScheduledIncidentConfigurationGroupingArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 entity_matching_method: Optional[pulumi.Input[str]] = None,
                 group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lookback_duration: Optional[pulumi.Input[str]] = None,
                 reopen_closed_incidents: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: Enable grouping incidents created from alerts triggered by this Sentinel Scheduled Alert Rule. Defaults to `true`.
        :param pulumi.Input[str] entity_matching_method: The method used to group incidents. Possible values are `All`, `Custom` and `None`. Defaults to `None`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_bies: A list of entity types to group by, only when the `entity_matching_method` is `Custom`. Possible values are `Account`, `Host`, `Url`, `Ip`.
        :param pulumi.Input[str] lookback_duration: Limit the group to alerts created within the lookback duration (in ISO 8601 duration format). Defaults to `PT5M`.
        :param pulumi.Input[bool] reopen_closed_incidents: Whether to re-open closed matching incidents? Defaults to `false`.
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
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable grouping incidents created from alerts triggered by this Sentinel Scheduled Alert Rule. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="entityMatchingMethod")
    def entity_matching_method(self) -> Optional[pulumi.Input[str]]:
        """
        The method used to group incidents. Possible values are `All`, `Custom` and `None`. Defaults to `None`.
        """
        return pulumi.get(self, "entity_matching_method")

    @entity_matching_method.setter
    def entity_matching_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_matching_method", value)

    @property
    @pulumi.getter(name="groupBies")
    def group_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of entity types to group by, only when the `entity_matching_method` is `Custom`. Possible values are `Account`, `Host`, `Url`, `Ip`.
        """
        return pulumi.get(self, "group_bies")

    @group_bies.setter
    def group_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_bies", value)

    @property
    @pulumi.getter(name="lookbackDuration")
    def lookback_duration(self) -> Optional[pulumi.Input[str]]:
        """
        Limit the group to alerts created within the lookback duration (in ISO 8601 duration format). Defaults to `PT5M`.
        """
        return pulumi.get(self, "lookback_duration")

    @lookback_duration.setter
    def lookback_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lookback_duration", value)

    @property
    @pulumi.getter(name="reopenClosedIncidents")
    def reopen_closed_incidents(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to re-open closed matching incidents? Defaults to `false`.
        """
        return pulumi.get(self, "reopen_closed_incidents")

    @reopen_closed_incidents.setter
    def reopen_closed_incidents(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reopen_closed_incidents", value)


@pulumi.input_type
class AuthomationRuleActionIncidentArgs:
    def __init__(__self__, *,
                 order: pulumi.Input[int],
                 classification: Optional[pulumi.Input[str]] = None,
                 classification_comment: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] order: The execution order of this action.
        :param pulumi.Input[str] classification: The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        :param pulumi.Input[str] classification_comment: The comment why the incident is to be closed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: Specifies a list of labels to add to the incident.
        :param pulumi.Input[str] owner_id: The object ID of the entity this incident is assigned to.
        :param pulumi.Input[str] severity: The severity to add to the incident.
        :param pulumi.Input[str] status: The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
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
    def order(self) -> pulumi.Input[int]:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def classification(self) -> Optional[pulumi.Input[str]]:
        """
        The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        """
        return pulumi.get(self, "classification")

    @classification.setter
    def classification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "classification", value)

    @property
    @pulumi.getter(name="classificationComment")
    def classification_comment(self) -> Optional[pulumi.Input[str]]:
        """
        The comment why the incident is to be closed.
        """
        return pulumi.get(self, "classification_comment")

    @classification_comment.setter
    def classification_comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "classification_comment", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies a list of labels to add to the incident.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the entity this incident is assigned to.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        The severity to add to the incident.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class AuthomationRuleActionPlaybookArgs:
    def __init__(__self__, *,
                 logic_app_id: pulumi.Input[str],
                 order: pulumi.Input[int],
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] logic_app_id: The ID of the Logic App that defines the playbook's logic.
        :param pulumi.Input[int] order: The execution order of this action.
        :param pulumi.Input[str] tenant_id: The ID of the Tenant that owns the playbook.
        """
        pulumi.set(__self__, "logic_app_id", logic_app_id)
        pulumi.set(__self__, "order", order)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="logicAppId")
    def logic_app_id(self) -> pulumi.Input[str]:
        """
        The ID of the Logic App that defines the playbook's logic.
        """
        return pulumi.get(self, "logic_app_id")

    @logic_app_id.setter
    def logic_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "logic_app_id", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        """
        The execution order of this action.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Tenant that owns the playbook.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class AuthomationRuleConditionArgs:
    def __init__(__self__, *,
                 operator: pulumi.Input[str],
                 property: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] operator: The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        :param pulumi.Input[str] property: The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: Specifies a list of values to use for evaluate the condition.
        """
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "property", property)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[str]:
        """
        The operator to use for evaluate the condition. Possible values include: `Equals`, `NotEquals`, `Contains`, `NotContains`, `StartsWith`, `NotStartsWith`, `EndsWith`, `NotEndsWith`.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies a list of values to use for evaluate the condition.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def property(self) -> pulumi.Input[str]:
        """
        The property to use for evaluate the condition. Possible values include: `AccountAadTenantId`, `AccountAadUserId`, `AccountNTDomain`, `AccountName`, `AccountObjectGuid`, `AccountPUID`, `AccountSid`, `AccountUPNSuffix`, `AzureResourceResourceId`, `AzureResourceSubscriptionId`, `CloudApplicationAppId`, `CloudApplicationAppName`, `DNSDomainName`, `FileDirectory`, `FileHashValue`, `FileName`, `HostAzureID`, `HostNTDomain`, `HostName`, `HostNetBiosName`, `HostOSVersion`, `IPAddress`, `IncidentDescription`, `IncidentProviderName`, `IncidentRelatedAnalyticRuleIds`, `IncidentSeverity`, `IncidentStatus`, `IncidentTactics`, `IncidentTitle`, `IoTDeviceId`, `IoTDeviceModel`, `IoTDeviceName`, `IoTDeviceOperatingSystem`, `IoTDeviceType`, `IoTDeviceVendor`, `MailMessageDeliveryAction`, `MailMessageDeliveryLocation`, `MailMessageP1Sender`, `MailMessageP2Sender`, `MailMessageRecipient`, `MailMessageSenderIP`, `MailMessageSubject`, `MailboxDisplayName`, `MailboxPrimaryAddress`, `MailboxUPN`, `MalwareCategory`, `MalwareName`, `ProcessCommandLine`, `ProcessId`, `RegistryKey`, `RegistryValueData`, `Url`.
        """
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: pulumi.Input[str]):
        pulumi.set(self, "property", value)


