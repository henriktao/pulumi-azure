// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    /// <summary>
    /// Manages a ServiceBus Subscription Rule.
    /// 
    /// ## Example Usage
    /// ### SQL Filter)
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
    ///         var exampleNamespace = new Azure.ServiceBus.Namespace("exampleNamespace", new Azure.ServiceBus.NamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Tags = 
    ///             {
    ///                 { "source", "example" },
    ///             },
    ///         });
    ///         var exampleTopic = new Azure.ServiceBus.Topic("exampleTopic", new Azure.ServiceBus.TopicArgs
    ///         {
    ///             NamespaceId = exampleNamespace.Id,
    ///             EnablePartitioning = true,
    ///         });
    ///         var exampleSubscription = new Azure.ServiceBus.Subscription("exampleSubscription", new Azure.ServiceBus.SubscriptionArgs
    ///         {
    ///             TopicId = exampleTopic.Id,
    ///             MaxDeliveryCount = 1,
    ///         });
    ///         var exampleSubscriptionRule = new Azure.ServiceBus.SubscriptionRule("exampleSubscriptionRule", new Azure.ServiceBus.SubscriptionRuleArgs
    ///         {
    ///             SubscriptionId = exampleSubscription.Id,
    ///             FilterType = "SqlFilter",
    ///             SqlFilter = "colour = 'red'",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Correlation Filter)
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
    ///         var exampleNamespace = new Azure.ServiceBus.Namespace("exampleNamespace", new Azure.ServiceBus.NamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "Standard",
    ///             Tags = 
    ///             {
    ///                 { "source", "example" },
    ///             },
    ///         });
    ///         var exampleTopic = new Azure.ServiceBus.Topic("exampleTopic", new Azure.ServiceBus.TopicArgs
    ///         {
    ///             NamespaceId = exampleNamespace.Id,
    ///             EnablePartitioning = true,
    ///         });
    ///         var exampleSubscription = new Azure.ServiceBus.Subscription("exampleSubscription", new Azure.ServiceBus.SubscriptionArgs
    ///         {
    ///             TopicId = exampleTopic.Id,
    ///             MaxDeliveryCount = 1,
    ///         });
    ///         var exampleSubscriptionRule = new Azure.ServiceBus.SubscriptionRule("exampleSubscriptionRule", new Azure.ServiceBus.SubscriptionRuleArgs
    ///         {
    ///             SubscriptionId = exampleSubscription.Id,
    ///             FilterType = "CorrelationFilter",
    ///             CorrelationFilter = new Azure.ServiceBus.Inputs.SubscriptionRuleCorrelationFilterArgs
    ///             {
    ///                 CorrelationId = "high",
    ///                 Label = "red",
    ///                 Properties = 
    ///                 {
    ///                     { "customProperty", "value" },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service Bus Subscription Rule can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:eventhub/subscriptionRule:SubscriptionRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.servicebus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
    /// ```
    /// </summary>
    [Obsolete(@"azure.eventhub.SubscriptionRule has been deprecated in favor of azure.servicebus.SubscriptionRule")]
    [AzureResourceType("azure:eventhub/subscriptionRule:SubscriptionRule")]
    public partial class SubscriptionRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Output("correlationFilter")]
        public Output<Outputs.SubscriptionRuleCorrelationFilter?> CorrelationFilter { get; private set; } = null!;

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Output("filterType")]
        public Output<string> FilterType { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Output("sqlFilter")]
        public Output<string?> SqlFilter { get; private set; } = null!;

        /// <summary>
        /// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        [Output("subscriptionName")]
        public Output<string> SubscriptionName { get; private set; } = null!;

        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a SubscriptionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscriptionRule(string name, SubscriptionRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscriptionRule:SubscriptionRule", name, args ?? new SubscriptionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubscriptionRule(string name, Input<string> id, SubscriptionRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:eventhub/subscriptionRule:SubscriptionRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubscriptionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscriptionRule Get(string name, Input<string> id, SubscriptionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SubscriptionRule(name, id, state, options);
        }
    }

    public sealed class SubscriptionRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Input("correlationFilter")]
        public Input<Inputs.SubscriptionRuleCorrelationFilterArgs>? CorrelationFilter { get; set; }

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Input("filterType", required: true)]
        public Input<string> FilterType { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Input("sqlFilter")]
        public Input<string>? SqlFilter { get; set; }

        /// <summary>
        /// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscriptionRuleArgs()
        {
        }
    }

    public sealed class SubscriptionRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents set of actions written in SQL language-based syntax that is performed against a BrokeredMessage.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A `correlation_filter` block as documented below to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `CorrelationFilter`.
        /// </summary>
        [Input("correlationFilter")]
        public Input<Inputs.SubscriptionRuleCorrelationFilterGetArgs>? CorrelationFilter { get; set; }

        /// <summary>
        /// Type of filter to be applied to a BrokeredMessage. Possible values are `SqlFilter` and `CorrelationFilter`.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        /// <summary>
        /// Specifies the name of the ServiceBus Subscription Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Represents a filter written in SQL language-based syntax that to be evaluated against a BrokeredMessage. Required when `filter_type` is set to `SqlFilter`.
        /// </summary>
        [Input("sqlFilter")]
        public Input<string>? SqlFilter { get; set; }

        /// <summary>
        /// The ID of the ServiceBus Subscription in which this Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public SubscriptionRuleState()
        {
        }
    }
}
