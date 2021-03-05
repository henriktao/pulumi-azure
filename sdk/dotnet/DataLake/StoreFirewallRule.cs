// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataLake
{
    /// <summary>
    /// Manages a Azure Data Lake Store Firewall Rule.
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
    ///         var exampleStore = new Azure.DataLake.Store("exampleStore", new Azure.DataLake.StoreArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var exampleStoreFirewallRule = new Azure.DataLake.StoreFirewallRule("exampleStoreFirewallRule", new Azure.DataLake.StoreFirewallRuleArgs
    ///         {
    ///             AccountName = exampleStore.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StartIpAddress = "1.2.3.4",
    ///             EndIpAddress = "2.3.4.5",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Lake Store Firewall Rules can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datalake/storeFirewallRule:StoreFirewallRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DataLakeStore/accounts/mydatalakeaccount/firewallRules/rule1
    /// ```
    /// </summary>
    [AzureResourceType("azure:datalake/storeFirewallRule:StoreFirewallRule")]
    public partial class StoreFirewallRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The End IP Address for the firewall rule.
        /// </summary>
        [Output("endIpAddress")]
        public Output<string> EndIpAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Data Lake Store.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The Start IP address for the firewall rule.
        /// </summary>
        [Output("startIpAddress")]
        public Output<string> StartIpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a StoreFirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StoreFirewallRule(string name, StoreFirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:datalake/storeFirewallRule:StoreFirewallRule", name, args ?? new StoreFirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StoreFirewallRule(string name, Input<string> id, StoreFirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:datalake/storeFirewallRule:StoreFirewallRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StoreFirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StoreFirewallRule Get(string name, Input<string> id, StoreFirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new StoreFirewallRule(name, id, state, options);
        }
    }

    public sealed class StoreFirewallRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The End IP Address for the firewall rule.
        /// </summary>
        [Input("endIpAddress", required: true)]
        public Input<string> EndIpAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Data Lake Store.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Start IP address for the firewall rule.
        /// </summary>
        [Input("startIpAddress", required: true)]
        public Input<string> StartIpAddress { get; set; } = null!;

        public StoreFirewallRuleArgs()
        {
        }
    }

    public sealed class StoreFirewallRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Data Lake Store for which the Firewall Rule should take effect.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The End IP Address for the firewall rule.
        /// </summary>
        [Input("endIpAddress")]
        public Input<string>? EndIpAddress { get; set; }

        /// <summary>
        /// Specifies the name of the Data Lake Store. Changing this forces a new resource to be created. Has to be between 3 to 24 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Data Lake Store.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The Start IP address for the firewall rule.
        /// </summary>
        [Input("startIpAddress")]
        public Input<string>? StartIpAddress { get; set; }

        public StoreFirewallRuleState()
        {
        }
    }
}
