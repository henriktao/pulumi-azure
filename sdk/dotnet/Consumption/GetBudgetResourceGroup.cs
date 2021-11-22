// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Consumption
{
    public static class GetBudgetResourceGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Consumption Budget for a specific resource group.
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
        ///         var example = Output.Create(Azure.Consumption.GetBudgetResourceGroup.InvokeAsync(new Azure.Consumption.GetBudgetResourceGroupArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupId = azurerm_resource_group.Example.Id,
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBudgetResourceGroupResult> InvokeAsync(GetBudgetResourceGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetResourceGroupResult>("azure:consumption/getBudgetResourceGroup:getBudgetResourceGroup", args ?? new GetBudgetResourceGroupArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Consumption Budget for a specific resource group.
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
        ///         var example = Output.Create(Azure.Consumption.GetBudgetResourceGroup.InvokeAsync(new Azure.Consumption.GetBudgetResourceGroupArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupId = azurerm_resource_group.Example.Id,
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBudgetResourceGroupResult> Invoke(GetBudgetResourceGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBudgetResourceGroupResult>("azure:consumption/getBudgetResourceGroup:getBudgetResourceGroup", args ?? new GetBudgetResourceGroupInvokeArgs(), options.WithVersion());
    }


    public sealed class GetBudgetResourceGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Consumption Budget.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the subscription.
        /// </summary>
        [Input("resourceGroupId", required: true)]
        public string ResourceGroupId { get; set; } = null!;

        public GetBudgetResourceGroupArgs()
        {
        }
    }

    public sealed class GetBudgetResourceGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Consumption Budget.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the subscription.
        /// </summary>
        [Input("resourceGroupId", required: true)]
        public Input<string> ResourceGroupId { get; set; } = null!;

        public GetBudgetResourceGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetResourceGroupResult
    {
        /// <summary>
        /// The total amount of cost to track with the budget.
        /// </summary>
        public readonly double Amount;
        /// <summary>
        /// A `filter` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetResourceGroupFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the tag used for the filter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A `notification` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetResourceGroupNotificationResult> Notifications;
        public readonly string ResourceGroupId;
        /// <summary>
        /// The time covered by a budget.
        /// </summary>
        public readonly string TimeGrain;
        /// <summary>
        /// A `time_period` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBudgetResourceGroupTimePeriodResult> TimePeriods;

        [OutputConstructor]
        private GetBudgetResourceGroupResult(
            double amount,

            ImmutableArray<Outputs.GetBudgetResourceGroupFilterResult> filters,

            string id,

            string name,

            ImmutableArray<Outputs.GetBudgetResourceGroupNotificationResult> notifications,

            string resourceGroupId,

            string timeGrain,

            ImmutableArray<Outputs.GetBudgetResourceGroupTimePeriodResult> timePeriods)
        {
            Amount = amount;
            Filters = filters;
            Id = id;
            Name = name;
            Notifications = notifications;
            ResourceGroupId = resourceGroupId;
            TimeGrain = timeGrain;
            TimePeriods = timePeriods;
        }
    }
}
