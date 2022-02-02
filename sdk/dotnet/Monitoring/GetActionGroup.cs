// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    public static class GetActionGroup
    {
        /// <summary>
        /// Use this data source to access the properties of an Action Group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Monitoring.GetActionGroup.InvokeAsync(new Azure.Monitoring.GetActionGroupArgs
        ///         {
        ///             ResourceGroupName = "example-rg",
        ///             Name = "tfex-actiongroup",
        ///         }));
        ///         this.ActionGroupId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("actionGroupId")]
        ///     public Output&lt;string&gt; ActionGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetActionGroupResult> InvokeAsync(GetActionGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActionGroupResult>("azure:monitoring/getActionGroup:getActionGroup", args ?? new GetActionGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access the properties of an Action Group.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Monitoring.GetActionGroup.InvokeAsync(new Azure.Monitoring.GetActionGroupArgs
        ///         {
        ///             ResourceGroupName = "example-rg",
        ///             Name = "tfex-actiongroup",
        ///         }));
        ///         this.ActionGroupId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("actionGroupId")]
        ///     public Output&lt;string&gt; ActionGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetActionGroupResult> Invoke(GetActionGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetActionGroupResult>("azure:monitoring/getActionGroup:getActionGroup", args ?? new GetActionGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionGroupArgs : Pulumi.InvokeArgs
    {
        [Input("eventHubReceivers")]
        private List<Inputs.GetActionGroupEventHubReceiverArgs>? _eventHubReceivers;

        /// <summary>
        /// One or more `event_hub_receiver` blocks as defined below.
        /// </summary>
        public List<Inputs.GetActionGroupEventHubReceiverArgs> EventHubReceivers
        {
            get => _eventHubReceivers ?? (_eventHubReceivers = new List<Inputs.GetActionGroupEventHubReceiverArgs>());
            set => _eventHubReceivers = value;
        }

        /// <summary>
        /// Specifies the name of the Action Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Action Group is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetActionGroupArgs()
        {
        }
    }

    public sealed class GetActionGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("eventHubReceivers")]
        private InputList<Inputs.GetActionGroupEventHubReceiverInputArgs>? _eventHubReceivers;

        /// <summary>
        /// One or more `event_hub_receiver` blocks as defined below.
        /// </summary>
        public InputList<Inputs.GetActionGroupEventHubReceiverInputArgs> EventHubReceivers
        {
            get => _eventHubReceivers ?? (_eventHubReceivers = new InputList<Inputs.GetActionGroupEventHubReceiverInputArgs>());
            set => _eventHubReceivers = value;
        }

        /// <summary>
        /// Specifies the name of the Action Group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Action Group is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetActionGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetActionGroupResult
    {
        /// <summary>
        /// One or more `arm_role_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupArmRoleReceiverResult> ArmRoleReceivers;
        /// <summary>
        /// One or more `automation_runbook_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupAutomationRunbookReceiverResult> AutomationRunbookReceivers;
        /// <summary>
        /// One or more `azure_app_push_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupAzureAppPushReceiverResult> AzureAppPushReceivers;
        /// <summary>
        /// One or more `azure_function_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupAzureFunctionReceiverResult> AzureFunctionReceivers;
        /// <summary>
        /// One or more `email_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupEmailReceiverResult> EmailReceivers;
        /// <summary>
        /// Whether this action group is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// One or more `event_hub_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupEventHubReceiverResult> EventHubReceivers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more `itsm_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupItsmReceiverResult> ItsmReceivers;
        /// <summary>
        /// One or more `logic_app_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupLogicAppReceiverResult> LogicAppReceivers;
        /// <summary>
        /// The name of the webhook receiver.
        /// </summary>
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The short name of the action group.
        /// </summary>
        public readonly string ShortName;
        /// <summary>
        /// One or more `sms_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupSmsReceiverResult> SmsReceivers;
        /// <summary>
        /// One or more `voice_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupVoiceReceiverResult> VoiceReceivers;
        /// <summary>
        /// One or more `webhook_receiver` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionGroupWebhookReceiverResult> WebhookReceivers;

        [OutputConstructor]
        private GetActionGroupResult(
            ImmutableArray<Outputs.GetActionGroupArmRoleReceiverResult> armRoleReceivers,

            ImmutableArray<Outputs.GetActionGroupAutomationRunbookReceiverResult> automationRunbookReceivers,

            ImmutableArray<Outputs.GetActionGroupAzureAppPushReceiverResult> azureAppPushReceivers,

            ImmutableArray<Outputs.GetActionGroupAzureFunctionReceiverResult> azureFunctionReceivers,

            ImmutableArray<Outputs.GetActionGroupEmailReceiverResult> emailReceivers,

            bool enabled,

            ImmutableArray<Outputs.GetActionGroupEventHubReceiverResult> eventHubReceivers,

            string id,

            ImmutableArray<Outputs.GetActionGroupItsmReceiverResult> itsmReceivers,

            ImmutableArray<Outputs.GetActionGroupLogicAppReceiverResult> logicAppReceivers,

            string name,

            string resourceGroupName,

            string shortName,

            ImmutableArray<Outputs.GetActionGroupSmsReceiverResult> smsReceivers,

            ImmutableArray<Outputs.GetActionGroupVoiceReceiverResult> voiceReceivers,

            ImmutableArray<Outputs.GetActionGroupWebhookReceiverResult> webhookReceivers)
        {
            ArmRoleReceivers = armRoleReceivers;
            AutomationRunbookReceivers = automationRunbookReceivers;
            AzureAppPushReceivers = azureAppPushReceivers;
            AzureFunctionReceivers = azureFunctionReceivers;
            EmailReceivers = emailReceivers;
            Enabled = enabled;
            EventHubReceivers = eventHubReceivers;
            Id = id;
            ItsmReceivers = itsmReceivers;
            LogicAppReceivers = logicAppReceivers;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ShortName = shortName;
            SmsReceivers = smsReceivers;
            VoiceReceivers = voiceReceivers;
            WebhookReceivers = webhookReceivers;
        }
    }
}
