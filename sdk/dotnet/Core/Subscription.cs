// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    /// <summary>
    /// Manages an Alias for a Subscription - which adds an Alias to an existing Subscription, allowing it to be managed in the provider - or create a new Subscription with a new Alias.
    /// 
    /// &gt; **NOTE:** Destroying a Subscription controlled by this resource will place the Subscription into a cancelled state. It is possible to re-activate a subscription within 90-days of cancellation, after which time the Subscription is irrevocably deleted, and the Subscription ID cannot be re-used. For further information see [here](https://docs.microsoft.com/en-us/azure/cost-management-billing/manage/cancel-azure-subscription#what-happens-after-subscription-cancellation). Users can optionally delete a Subscription once 72 hours have passed, however, this functionality is not suitable for this provider. A `Deleted` subscription cannot be reactivated.
    /// 
    /// &gt; **NOTE:** It is not possible to destroy (cancel) a subscription if it contains resources. If resources are present that are not managed by the provider then these will need to be removed before the Subscription can be destroyed.
    /// 
    /// &gt; **NOTE:** Azure supports Multiple Aliases per Subscription, however, to reliably manage this resource in this provider only a single Alias is supported.
    /// 
    /// ## Example Usage
    /// ### Creating A New Alias And Subscription For An Enrollment Account
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEnrollmentAccountScope = Output.Create(Azure.Billing.GetEnrollmentAccountScope.InvokeAsync(new Azure.Billing.GetEnrollmentAccountScopeArgs
    ///         {
    ///             BillingAccountName = "1234567890",
    ///             EnrollmentAccountName = "0123456",
    ///         }));
    ///         var exampleSubscription = new Azure.Core.Subscription("exampleSubscription", new Azure.Core.SubscriptionArgs
    ///         {
    ///             SubscriptionName = "My Example EA Subscription",
    ///             BillingScopeId = exampleEnrollmentAccountScope.Apply(exampleEnrollmentAccountScope =&gt; exampleEnrollmentAccountScope.Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Creating A New Alias And Subscription For A Microsoft Customer Account
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleMcaAccountScope = Output.Create(Azure.Billing.GetMcaAccountScope.InvokeAsync(new Azure.Billing.GetMcaAccountScopeArgs
    ///         {
    ///             BillingAccountName = "e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
    ///             BillingProfileName = "PE2Q-NOIT-BG7-TGB",
    ///             InvoiceSectionName = "MTT4-OBS7-PJA-TGB",
    ///         }));
    ///         var exampleSubscription = new Azure.Core.Subscription("exampleSubscription", new Azure.Core.SubscriptionArgs
    ///         {
    ///             SubscriptionName = "My Example MCA Subscription",
    ///             BillingScopeId = exampleMcaAccountScope.Apply(exampleMcaAccountScope =&gt; exampleMcaAccountScope.Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Creating A New Alias And Subscription For A Microsoft Partner Account
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleMpaAccountScope = Output.Create(Azure.Billing.GetMpaAccountScope.InvokeAsync(new Azure.Billing.GetMpaAccountScopeArgs
    ///         {
    ///             BillingAccountName = "e879cf0f-2b4d-5431-109a-f72fc9868693:024cabf4-7321-4cf9-be59-df0c77ca51de_2019-05-31",
    ///             CustomerName = "2281f543-7321-4cf9-1e23-edb4Oc31a31c",
    ///         }));
    ///         var exampleSubscription = new Azure.Core.Subscription("exampleSubscription", new Azure.Core.SubscriptionArgs
    ///         {
    ///             SubscriptionName = "My Example MPA Subscription",
    ///             BillingScopeId = exampleMpaAccountScope.Apply(exampleMpaAccountScope =&gt; exampleMpaAccountScope.Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Adding An Alias To An Existing Subscription
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Azure.Core.Subscription("example", new Azure.Core.SubscriptionArgs
    ///         {
    ///             Alias = "examplesub",
    ///             SubscriptionId = "12345678-12234-5678-9012-123456789012",
    ///             SubscriptionName = "My Example Subscription",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Subscriptions can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:core/subscription:Subscription example "/providers/Microsoft.Subscription/aliases/subscription1"
    /// ```
    /// 
    ///  In this scenario, the `subscription_id` property can be completed and the provider will assume control of the existing subscription by creating an Alias. See the `adding an Alias to an existing Subscription` above. This provider requires an alias to correctly manage Subscription resources due to Azure Subscription API design.
    /// </summary>
    [AzureResourceType("azure:core/subscription:Subscription")]
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// The Azure Billing Scope ID. Can be a Microsoft Customer Account Billing Scope ID, a Microsoft Partner Account Billing Scope ID or an Enrollment Billing Scope ID.
        /// </summary>
        [Output("billingScopeId")]
        public Output<string?> BillingScopeId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subscription. Changing this forces a new Subscription to be created.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The Name of the Subscription. This is the Display Name in the portal.
        /// </summary>
        [Output("subscriptionName")]
        public Output<string> SubscriptionName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Subscription.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the Tenant to which the subscription belongs.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        /// </summary>
        [Output("workload")]
        public Output<string?> Workload { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure:core/subscription:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("azure:core/subscription:Subscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, SubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, state, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The Azure Billing Scope ID. Can be a Microsoft Customer Account Billing Scope ID, a Microsoft Partner Account Billing Scope ID or an Enrollment Billing Scope ID.
        /// </summary>
        [Input("billingScopeId")]
        public Input<string>? BillingScopeId { get; set; }

        /// <summary>
        /// The ID of the Subscription. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The Name of the Subscription. This is the Display Name in the portal.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Subscription.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("workload")]
        public Input<string>? Workload { get; set; }

        public SubscriptionArgs()
        {
        }
    }

    public sealed class SubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alias name for the subscription. This provider will generate a new GUID if this is not supplied. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The Azure Billing Scope ID. Can be a Microsoft Customer Account Billing Scope ID, a Microsoft Partner Account Billing Scope ID or an Enrollment Billing Scope ID.
        /// </summary>
        [Input("billingScopeId")]
        public Input<string>? BillingScopeId { get; set; }

        /// <summary>
        /// The ID of the Subscription. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The Name of the Subscription. This is the Display Name in the portal.
        /// </summary>
        [Input("subscriptionName")]
        public Input<string>? SubscriptionName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Subscription.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the Tenant to which the subscription belongs.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The workload type of the Subscription.  Possible values are `Production` (default) and `DevTest`. Changing this forces a new Subscription to be created.
        /// </summary>
        [Input("workload")]
        public Input<string>? Workload { get; set; }

        public SubscriptionState()
        {
        }
    }
}
