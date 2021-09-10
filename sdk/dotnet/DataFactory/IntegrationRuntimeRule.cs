// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    /// <summary>
    /// Manages a Data Factory Azure Integration Runtime.
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
    ///         var exampleFactory = new Azure.DataFactory.Factory("exampleFactory", new Azure.DataFactory.FactoryArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleIntegrationRuntimeRule = new Azure.DataFactory.IntegrationRuntimeRule("exampleIntegrationRuntimeRule", new Azure.DataFactory.IntegrationRuntimeRuleArgs
    ///         {
    ///             DataFactoryName = exampleFactory.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Factory Azure Integration Runtimes can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
    /// ```
    /// </summary>
    [AzureResourceType("azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule")]
    public partial class IntegrationRuntimeRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster will not be recycled and it will be used in next data flow activity run until TTL (time to live) is reached if this is set as `false`. Default is `true`.
        /// </summary>
        [Output("cleanupEnabled")]
        public Output<bool> CleanupEnabled { get; private set; } = null!;

        /// <summary>
        /// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
        /// </summary>
        [Output("computeType")]
        public Output<string?> ComputeType { get; private set; } = null!;

        /// <summary>
        /// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
        /// </summary>
        [Output("coreCount")]
        public Output<int?> CoreCount { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dataFactoryName")]
        public Output<string> DataFactoryName { get; private set; } = null!;

        /// <summary>
        /// Integration runtime description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
        /// </summary>
        [Output("timeToLiveMin")]
        public Output<int?> TimeToLiveMin { get; private set; } = null!;

        /// <summary>
        /// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Output("virtualNetworkEnabled")]
        public Output<bool?> VirtualNetworkEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationRuntimeRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationRuntimeRule(string name, IntegrationRuntimeRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule", name, args ?? new IntegrationRuntimeRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationRuntimeRule(string name, Input<string> id, IntegrationRuntimeRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationRuntimeRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationRuntimeRule Get(string name, Input<string> id, IntegrationRuntimeRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationRuntimeRule(name, id, state, options);
        }
    }

    public sealed class IntegrationRuntimeRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster will not be recycled and it will be used in next data flow activity run until TTL (time to live) is reached if this is set as `false`. Default is `true`.
        /// </summary>
        [Input("cleanupEnabled")]
        public Input<bool>? CleanupEnabled { get; set; }

        /// <summary>
        /// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
        /// </summary>
        [Input("computeType")]
        public Input<string>? ComputeType { get; set; }

        /// <summary>
        /// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
        /// </summary>
        [Input("coreCount")]
        public Input<int>? CoreCount { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryName", required: true)]
        public Input<string> DataFactoryName { get; set; } = null!;

        /// <summary>
        /// Integration runtime description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
        /// </summary>
        [Input("timeToLiveMin")]
        public Input<int>? TimeToLiveMin { get; set; }

        /// <summary>
        /// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkEnabled")]
        public Input<bool>? VirtualNetworkEnabled { get; set; }

        public IntegrationRuntimeRuleArgs()
        {
        }
    }

    public sealed class IntegrationRuntimeRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster will not be recycled and it will be used in next data flow activity run until TTL (time to live) is reached if this is set as `false`. Default is `true`.
        /// </summary>
        [Input("cleanupEnabled")]
        public Input<bool>? CleanupEnabled { get; set; }

        /// <summary>
        /// Compute type of the cluster which will execute data flow job. Valid values are `General`, `ComputeOptimized` and `MemoryOptimized`. Defaults to `General`.
        /// </summary>
        [Input("computeType")]
        public Input<string>? ComputeType { get; set; }

        /// <summary>
        /// Core count of the cluster which will execute data flow job. Valid values are `8`, `16`, `32`, `48`, `80`, `144` and `272`. Defaults to `8`.
        /// </summary>
        [Input("coreCount")]
        public Input<int>? CoreCount { get; set; }

        /// <summary>
        /// Specifies the name of the Data Factory the Managed Integration Runtime belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataFactoryName")]
        public Input<string>? DataFactoryName { get; set; }

        /// <summary>
        /// Integration runtime description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Managed Integration Runtime. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Managed Integration Runtime. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job. Defaults to `0`.
        /// </summary>
        [Input("timeToLiveMin")]
        public Input<int>? TimeToLiveMin { get; set; }

        /// <summary>
        /// Is Integration Runtime compute provisioned within Managed Virtual Network? Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkEnabled")]
        public Input<bool>? VirtualNetworkEnabled { get; set; }

        public IntegrationRuntimeRuleState()
        {
        }
    }
}
