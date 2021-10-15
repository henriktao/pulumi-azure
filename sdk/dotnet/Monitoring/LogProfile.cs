// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    /// <summary>
    /// Manages a [Log Profile](https://docs.microsoft.com/en-us/azure/monitoring-and-diagnostics/monitoring-overview-activity-logs#export-the-activity-log-with-a-log-profile). A Log Profile configures how Activity Logs are exported.
    /// 
    /// &gt; **NOTE:** It's only possible to configure one Log Profile per Subscription. If you are trying to create more than one Log Profile, an error with `StatusCode=409` will occur.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleEventHubNamespace = new Azure.EventHub.EventHubNamespace("exampleEventHubNamespace", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Capacity = 2,
    ///         });
    ///         var exampleLogProfile = new Azure.Monitoring.LogProfile("exampleLogProfile", new Azure.Monitoring.LogProfileArgs
    ///         {
    ///             Categories = 
    ///             {
    ///                 "Action",
    ///                 "Delete",
    ///                 "Write",
    ///             },
    ///             Locations = 
    ///             {
    ///                 "westus",
    ///                 "global",
    ///             },
    ///             ServicebusRuleId = exampleEventHubNamespace.Id.Apply(id =&gt; $"{id}/authorizationrules/RootManageSharedAccessKey"),
    ///             StorageAccountId = exampleAccount.Id,
    ///             RetentionPolicy = new Azure.Monitoring.Inputs.LogProfileRetentionPolicyArgs
    ///             {
    ///                 Enabled = true,
    ///                 Days = 7,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A Log Profile can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:monitoring/logProfile:LogProfile example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Insights/logProfiles/test
    /// ```
    /// </summary>
    [AzureResourceType("azure:monitoring/logProfile:LogProfile")]
    public partial class LogProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// List of categories of the logs.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<string>> Categories { get; private set; } = null!;

        /// <summary>
        /// List of regions for which Activity Log events are stored or streamed.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// The name of the Log Profile. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.LogProfileRetentionPolicy> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Output("servicebusRuleId")]
        public Output<string?> ServicebusRuleId { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Output("storageAccountId")]
        public Output<string?> StorageAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a LogProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogProfile(string name, LogProfileArgs args, CustomResourceOptions? options = null)
            : base("azure:monitoring/logProfile:LogProfile", name, args ?? new LogProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogProfile(string name, Input<string> id, LogProfileState? state = null, CustomResourceOptions? options = null)
            : base("azure:monitoring/logProfile:LogProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LogProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogProfile Get(string name, Input<string> id, LogProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new LogProfile(name, id, state, options);
        }
    }

    public sealed class LogProfileArgs : Pulumi.ResourceArgs
    {
        [Input("categories", required: true)]
        private InputList<string>? _categories;

        /// <summary>
        /// List of categories of the logs.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        [Input("locations", required: true)]
        private InputList<string>? _locations;

        /// <summary>
        /// List of regions for which Activity Log events are stored or streamed.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// The name of the Log Profile. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        /// </summary>
        [Input("retentionPolicy", required: true)]
        public Input<Inputs.LogProfileRetentionPolicyArgs> RetentionPolicy { get; set; } = null!;

        /// <summary>
        /// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Input("servicebusRuleId")]
        public Input<string>? ServicebusRuleId { get; set; }

        /// <summary>
        /// The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public LogProfileArgs()
        {
        }
    }

    public sealed class LogProfileState : Pulumi.ResourceArgs
    {
        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// List of categories of the logs.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// List of regions for which Activity Log events are stored or streamed.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// The name of the Log Profile. Changing this forces a
        /// new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `retention_policy` block as documented below. A retention policy for how long Activity Logs are retained in the storage account.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.LogProfileRetentionPolicyGetArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Input("servicebusRuleId")]
        public Input<string>? ServicebusRuleId { get; set; }

        /// <summary>
        /// The resource ID of the storage account in which the Activity Log is stored. At least one of `storage_account_id` or `servicebus_rule_id` must be set.
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public LogProfileState()
        {
        }
    }
}
