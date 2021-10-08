// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hpc
{
    /// <summary>
    /// Manages a HPC Cache Access Policy.
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
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.1.0/24",
    ///             },
    ///         });
    ///         var exampleCache = new Azure.Hpc.Cache("exampleCache", new Azure.Hpc.CacheArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             CacheSizeInGb = 3072,
    ///             SubnetId = exampleSubnet.Id,
    ///             SkuName = "Standard_2G",
    ///         });
    ///         var exampleCacheAccessPolicy = new Azure.Hpc.CacheAccessPolicy("exampleCacheAccessPolicy", new Azure.Hpc.CacheAccessPolicyArgs
    ///         {
    ///             HpcCacheId = exampleCache.Id,
    ///             AccessRules = 
    ///             {
    ///                 new Azure.Hpc.Inputs.CacheAccessPolicyAccessRuleArgs
    ///                 {
    ///                     Scope = "default",
    ///                     Access = "rw",
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
    /// HPC Cache Access Policies can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:hpc/cacheAccessPolicy:CacheAccessPolicy example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/cacheAccessPolicies/policy1
    /// ```
    /// </summary>
    [AzureResourceType("azure:hpc/cacheAccessPolicy:CacheAccessPolicy")]
    public partial class CacheAccessPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Up to three `access_rule` blocks as defined below.
        /// </summary>
        [Output("accessRules")]
        public Output<ImmutableArray<Outputs.CacheAccessPolicyAccessRule>> AccessRules { get; private set; } = null!;

        /// <summary>
        /// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Output("hpcCacheId")]
        public Output<string> HpcCacheId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a CacheAccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CacheAccessPolicy(string name, CacheAccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:hpc/cacheAccessPolicy:CacheAccessPolicy", name, args ?? new CacheAccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CacheAccessPolicy(string name, Input<string> id, CacheAccessPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:hpc/cacheAccessPolicy:CacheAccessPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CacheAccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CacheAccessPolicy Get(string name, Input<string> id, CacheAccessPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new CacheAccessPolicy(name, id, state, options);
        }
    }

    public sealed class CacheAccessPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("accessRules", required: true)]
        private InputList<Inputs.CacheAccessPolicyAccessRuleArgs>? _accessRules;

        /// <summary>
        /// Up to three `access_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CacheAccessPolicyAccessRuleArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.CacheAccessPolicyAccessRuleArgs>());
            set => _accessRules = value;
        }

        /// <summary>
        /// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Input("hpcCacheId", required: true)]
        public Input<string> HpcCacheId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CacheAccessPolicyArgs()
        {
        }
    }

    public sealed class CacheAccessPolicyState : Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.CacheAccessPolicyAccessRuleGetArgs>? _accessRules;

        /// <summary>
        /// Up to three `access_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.CacheAccessPolicyAccessRuleGetArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.CacheAccessPolicyAccessRuleGetArgs>());
            set => _accessRules = value;
        }

        /// <summary>
        /// The ID of the HPC Cache that this HPC Cache Access Policy resides in. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Input("hpcCacheId")]
        public Input<string>? HpcCacheId { get; set; }

        /// <summary>
        /// The name which should be used for this HPC Cache Access Policy. Changing this forces a new HPC Cache Access Policy to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CacheAccessPolicyState()
        {
        }
    }
}
