// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Consumption
{
    public static class GetBudgetSubscription
    {
        /// <summary>
        /// Use this data source to access information about an existing Consumption Budget for a specific subscription.
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
        ///         var example = Output.Create(Azure.Consumption.GetBudgetSubscription.InvokeAsync(new Azure.Consumption.GetBudgetSubscriptionArgs
        ///         {
        ///             Name = "existing",
        ///             SubscriptionId = "/subscriptions/00000000-0000-0000-0000-000000000000/",
        ///         }));
        ///         this.Id = data.Azurerm_consumption_budget.Example.Id;
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBudgetSubscriptionResult> InvokeAsync(GetBudgetSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetSubscriptionResult>("azure:consumption/getBudgetSubscription:getBudgetSubscription", args ?? new GetBudgetSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Consumption Budget for a specific subscription.
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
        ///         var example = Output.Create(Azure.Consumption.GetBudgetSubscription.InvokeAsync(new Azure.Consumption.GetBudgetSubscriptionArgs
        ///         {
        ///             Name = "existing",
        ///             SubscriptionId = "/subscriptions/00000000-0000-0000-0000-000000000000/",
        ///         }));
        ///         this.Id = data.Azurerm_consumption_budget.Example.Id;
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBudgetSubscriptionResult> Invoke(GetBudgetSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBudgetSubscriptionResult>("azure:consumption/getBudgetSubscription:getBudgetSubscription", args ?? new GetBudgetSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBudgetSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Consumption Budget.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the subscription.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        public GetBudgetSubscriptionArgs()
        {
        }
    }

    public sealed class GetBudgetSubscriptionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Consumption Budget.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the subscription.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public GetBudgetSubscriptionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetSubscriptionResult
    {
        /// <summary>
        /// The total amount of cost to track with the budget.
        /// </summary>
        public readonly double Amount;
        /// <summary>
        /// A `filter` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetSubscriptionFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the tag to use for the filter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `notification` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetSubscriptionNotificationResult> Notifications;
        public readonly string SubscriptionId;
        /// <summary>
        /// The time covered by a budget.
        /// </summary>
        public readonly string TimeGrain;
        /// <summary>
        /// A `time_period` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetSubscriptionTimePeriodResult> TimePeriods;

        [OutputConstructor]
        private GetBudgetSubscriptionResult(
            double amount,

            ImmutableArray<Outputs.GetBudgetSubscriptionFilterResult> filters,

            string id,

            string name,

            ImmutableArray<Outputs.GetBudgetSubscriptionNotificationResult> notifications,

            string subscriptionId,

            string timeGrain,

            ImmutableArray<Outputs.GetBudgetSubscriptionTimePeriodResult> timePeriods)
        {
            Amount = amount;
            Filters = filters;
            Id = id;
            Name = name;
            Notifications = notifications;
            SubscriptionId = subscriptionId;
            TimeGrain = timeGrain;
            TimePeriods = timePeriods;
        }
    }
}
