# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WebhookArgs', 'Webhook']

@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 automation_account_name: pulumi.Input[str],
                 expiry_time: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 runbook_name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 run_on_worker_group: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Webhook resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] expiry_time: Timestamp when the webhook expires. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] runbook_name: Name of the Automation Runbook to execute by Webhook.
        :param pulumi.Input[bool] enabled: Controls if Webhook is enabled. Defaults to `true`.
        :param pulumi.Input[str] name: Specifies the name of the Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Map of input parameters passed to runbook.
        :param pulumi.Input[str] run_on_worker_group: Name of the hybrid worker group the Webhook job will run on.
        :param pulumi.Input[str] uri: URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        pulumi.set(__self__, "automation_account_name", automation_account_name)
        pulumi.set(__self__, "expiry_time", expiry_time)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "runbook_name", runbook_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if run_on_worker_group is not None:
            pulumi.set(__self__, "run_on_worker_group", run_on_worker_group)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> pulumi.Input[str]:
        """
        The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @automation_account_name.setter
    def automation_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "automation_account_name", value)

    @property
    @pulumi.getter(name="expiryTime")
    def expiry_time(self) -> pulumi.Input[str]:
        """
        Timestamp when the webhook expires. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "expiry_time")

    @expiry_time.setter
    def expiry_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "expiry_time", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="runbookName")
    def runbook_name(self) -> pulumi.Input[str]:
        """
        Name of the Automation Runbook to execute by Webhook.
        """
        return pulumi.get(self, "runbook_name")

    @runbook_name.setter
    def runbook_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "runbook_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls if Webhook is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Webhook. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of input parameters passed to runbook.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="runOnWorkerGroup")
    def run_on_worker_group(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the hybrid worker group the Webhook job will run on.
        """
        return pulumi.get(self, "run_on_worker_group")

    @run_on_worker_group.setter
    def run_on_worker_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "run_on_worker_group", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)


@pulumi.input_type
class _WebhookState:
    def __init__(__self__, *,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiry_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_on_worker_group: Optional[pulumi.Input[str]] = None,
                 runbook_name: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Webhook resources.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Controls if Webhook is enabled. Defaults to `true`.
        :param pulumi.Input[str] expiry_time: Timestamp when the webhook expires. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Map of input parameters passed to runbook.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] run_on_worker_group: Name of the hybrid worker group the Webhook job will run on.
        :param pulumi.Input[str] runbook_name: Name of the Automation Runbook to execute by Webhook.
        :param pulumi.Input[str] uri: URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        if automation_account_name is not None:
            pulumi.set(__self__, "automation_account_name", automation_account_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if expiry_time is not None:
            pulumi.set(__self__, "expiry_time", expiry_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if run_on_worker_group is not None:
            pulumi.set(__self__, "run_on_worker_group", run_on_worker_group)
        if runbook_name is not None:
            pulumi.set(__self__, "runbook_name", runbook_name)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @automation_account_name.setter
    def automation_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "automation_account_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls if Webhook is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="expiryTime")
    def expiry_time(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp when the webhook expires. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "expiry_time")

    @expiry_time.setter
    def expiry_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Webhook. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of input parameters passed to runbook.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="runOnWorkerGroup")
    def run_on_worker_group(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the hybrid worker group the Webhook job will run on.
        """
        return pulumi.get(self, "run_on_worker_group")

    @run_on_worker_group.setter
    def run_on_worker_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "run_on_worker_group", value)

    @property
    @pulumi.getter(name="runbookName")
    def runbook_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Automation Runbook to execute by Webhook.
        """
        return pulumi.get(self, "runbook_name")

    @runbook_name.setter
    def runbook_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runbook_name", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)


class Webhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiry_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_on_worker_group: Optional[pulumi.Input[str]] = None,
                 runbook_name: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Automation Runbook's Webhook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Basic")
        example_run_book = azure.automation.RunBook("exampleRunBook",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            log_verbose=True,
            log_progress=True,
            description="This is an example runbook",
            runbook_type="PowerShellWorkflow",
            publish_content_link=azure.automation.RunBookPublishContentLinkArgs(
                uri="https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1",
            ))
        example_webhook = azure.automation.Webhook("exampleWebhook",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            expiry_time="2021-12-31T00:00:00Z",
            enabled=True,
            runbook_name=example_run_book.name,
            parameters={
                "input": "parameter",
            })
        ```

        ## Import

        Automation Webhooks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:automation/webhook:Webhook TestRunbook_webhook /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webhooks/TestRunbook_webhook
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Controls if Webhook is enabled. Defaults to `true`.
        :param pulumi.Input[str] expiry_time: Timestamp when the webhook expires. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Map of input parameters passed to runbook.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] run_on_worker_group: Name of the hybrid worker group the Webhook job will run on.
        :param pulumi.Input[str] runbook_name: Name of the Automation Runbook to execute by Webhook.
        :param pulumi.Input[str] uri: URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Automation Runbook's Webhook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.automation.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Basic")
        example_run_book = azure.automation.RunBook("exampleRunBook",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            log_verbose=True,
            log_progress=True,
            description="This is an example runbook",
            runbook_type="PowerShellWorkflow",
            publish_content_link=azure.automation.RunBookPublishContentLinkArgs(
                uri="https://raw.githubusercontent.com/Azure/azure-quickstart-templates/c4935ffb69246a6058eb24f54640f53f69d3ac9f/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1",
            ))
        example_webhook = azure.automation.Webhook("exampleWebhook",
            resource_group_name=example_resource_group.name,
            automation_account_name=example_account.name,
            expiry_time="2021-12-31T00:00:00Z",
            enabled=True,
            runbook_name=example_run_book.name,
            parameters={
                "input": "parameter",
            })
        ```

        ## Import

        Automation Webhooks can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:automation/webhook:Webhook TestRunbook_webhook /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webhooks/TestRunbook_webhook
        ```

        :param str resource_name: The name of the resource.
        :param WebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiry_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_on_worker_group: Optional[pulumi.Input[str]] = None,
                 runbook_name: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookArgs.__new__(WebhookArgs)

            if automation_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__.__dict__["automation_account_name"] = automation_account_name
            __props__.__dict__["enabled"] = enabled
            if expiry_time is None and not opts.urn:
                raise TypeError("Missing required property 'expiry_time'")
            __props__.__dict__["expiry_time"] = expiry_time
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["run_on_worker_group"] = run_on_worker_group
            if runbook_name is None and not opts.urn:
                raise TypeError("Missing required property 'runbook_name'")
            __props__.__dict__["runbook_name"] = runbook_name
            __props__.__dict__["uri"] = uri
        super(Webhook, __self__).__init__(
            'azure:automation/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            automation_account_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            expiry_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            run_on_worker_group: Optional[pulumi.Input[str]] = None,
            runbook_name: Optional[pulumi.Input[str]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Controls if Webhook is enabled. Defaults to `true`.
        :param pulumi.Input[str] expiry_time: Timestamp when the webhook expires. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Webhook. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Map of input parameters passed to runbook.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] run_on_worker_group: Name of the hybrid worker group the Webhook job will run on.
        :param pulumi.Input[str] runbook_name: Name of the Automation Runbook to execute by Webhook.
        :param pulumi.Input[str] uri: URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebhookState.__new__(_WebhookState)

        __props__.__dict__["automation_account_name"] = automation_account_name
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["expiry_time"] = expiry_time
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["run_on_worker_group"] = run_on_worker_group
        __props__.__dict__["runbook_name"] = runbook_name
        __props__.__dict__["uri"] = uri
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> pulumi.Output[str]:
        """
        The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "automation_account_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Controls if Webhook is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="expiryTime")
    def expiry_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the webhook expires. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "expiry_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Webhook. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of input parameters passed to runbook.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="runOnWorkerGroup")
    def run_on_worker_group(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the hybrid worker group the Webhook job will run on.
        """
        return pulumi.get(self, "run_on_worker_group")

    @property
    @pulumi.getter(name="runbookName")
    def runbook_name(self) -> pulumi.Output[str]:
        """
        Name of the Automation Runbook to execute by Webhook.
        """
        return pulumi.get(self, "runbook_name")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        URI to initiate the webhook. Can be generated using [Generate URI API](https://docs.microsoft.com/en-us/rest/api/automation/webhook/generate-uri). By default, new URI is generated on each new resource creation.
        """
        return pulumi.get(self, "uri")

