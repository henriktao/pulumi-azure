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
    /// Allows you to add, update, or remove an Azure Data Lake Store to a subnet of a virtual network.
    /// 
    /// &gt; **Note:** This resoruce manages an `Azure Data Lake Storage Gen1`, previously known as `Azure Data Lake Store`.
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
    ///             Location = "northeurope",
    ///         });
    ///         var vnet = new Azure.Network.VirtualNetwork("vnet", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.7.29.0/29",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var subnet = new Azure.Network.Subnet("subnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = vnet.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.7.29.0/29",
    ///             },
    ///             ServiceEndpoints = 
    ///             {
    ///                 "Microsoft.AzureActiveDirectory",
    ///             },
    ///         });
    ///         var exampleStore = new Azure.DataLake.Store("exampleStore", new Azure.DataLake.StoreArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var adlsvnetrule = new Azure.DataLake.StoreVirtualNetworkRule("adlsvnetrule", new Azure.DataLake.StoreVirtualNetworkRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AccountName = exampleStore.Name,
    ///             SubnetId = subnet.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Lake Store Virtual Network Rules can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datalake/storeVirtualNetworkRule:StoreVirtualNetworkRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DataLakeStore/accounts/myaccount/virtualNetworkRules/vnetrulename
    /// ```
    /// </summary>
    [AzureResourceType("azure:datalake/storeVirtualNetworkRule:StoreVirtualNetworkRule")]
    public partial class StoreVirtualNetworkRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Data Lake Store to which this Data Lake Store virtual network rule will be applied to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the Data Lake Store virtual network rule. Changing this forces a new resource to be created. Cannot be empty and must only contain alphanumeric characters, underscores, periods and hyphens. Cannot start with a period, underscore or hyphen, and cannot end with a period and a hyphen.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where the Data Lake Store resides. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet that the Data Lake Store will be connected to.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a StoreVirtualNetworkRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StoreVirtualNetworkRule(string name, StoreVirtualNetworkRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:datalake/storeVirtualNetworkRule:StoreVirtualNetworkRule", name, args ?? new StoreVirtualNetworkRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StoreVirtualNetworkRule(string name, Input<string> id, StoreVirtualNetworkRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:datalake/storeVirtualNetworkRule:StoreVirtualNetworkRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StoreVirtualNetworkRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StoreVirtualNetworkRule Get(string name, Input<string> id, StoreVirtualNetworkRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new StoreVirtualNetworkRule(name, id, state, options);
        }
    }

    public sealed class StoreVirtualNetworkRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Store to which this Data Lake Store virtual network rule will be applied to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Data Lake Store virtual network rule. Changing this forces a new resource to be created. Cannot be empty and must only contain alphanumeric characters, underscores, periods and hyphens. Cannot start with a period, underscore or hyphen, and cannot end with a period and a hyphen.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group where the Data Lake Store resides. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet that the Data Lake Store will be connected to.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public StoreVirtualNetworkRuleArgs()
        {
        }
    }

    public sealed class StoreVirtualNetworkRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Store to which this Data Lake Store virtual network rule will be applied to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The name of the Data Lake Store virtual network rule. Changing this forces a new resource to be created. Cannot be empty and must only contain alphanumeric characters, underscores, periods and hyphens. Cannot start with a period, underscore or hyphen, and cannot end with a period and a hyphen.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group where the Data Lake Store resides. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the subnet that the Data Lake Store will be connected to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public StoreVirtualNetworkRuleState()
        {
        }
    }
}
