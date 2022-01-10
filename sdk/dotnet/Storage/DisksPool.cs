// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages a Disk Pool.
    /// 
    /// !&gt; **Note:** This resource has been deprecated in favour of `azure.compute.DiskPool` and will be removed in version 3.0 of the Azure Provider.
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
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleVirtualNetwork.ResourceGroupName,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.0.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "diskspool",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/virtualNetworks/read",
    ///                         },
    ///                         Name = "Microsoft.StoragePool/diskPools",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleDisksPool = new Azure.Storage.DisksPool("exampleDisksPool", new Azure.Storage.DisksPoolArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             SubnetId = exampleSubnet.Id,
    ///             AvailabilityZones = 
    ///             {
    ///                 "1",
    ///             },
    ///             SkuName = "Basic_B1",
    ///             Tags = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Disk Pools can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:storage/disksPool:DisksPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/disksPool1
    /// ```
    /// </summary>
    [AzureResourceType("azure:storage/disksPool:DisksPool")]
    public partial class DisksPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Disks Pool.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DisksPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DisksPool(string name, DisksPoolArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/disksPool:DisksPool", name, args ?? new DisksPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DisksPool(string name, Input<string> id, DisksPoolState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/disksPool:DisksPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DisksPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DisksPool Get(string name, Input<string> id, DisksPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new DisksPool(name, id, state, options);
        }
    }

    public sealed class DisksPoolArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZones", required: true)]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        /// <summary>
        /// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Disks Pool.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DisksPoolArgs()
        {
        }
    }

    public sealed class DisksPoolState : Pulumi.ResourceArgs
    {
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// Specifies a list of logical zone (e.g. `["1"]`). Changing this forces a new Disk Pool to be created.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The Azure Region where the Disks Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Disks Pool. The name must begin with a letter or number, end with a letter, number or underscore, and may contain only letters, numbers, underscores, periods, or hyphens, and length should be in the range [7 - 30]. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Disk Pool should exist. Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The sku name of the Disk Pool. Possible values are "Basic_B1", "Standard_S1" and "Premium_P1". Changing this forces a new Disk Pool to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The ID of the Subnet for the Disk Pool. Changing this forces a new Disks Pool to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Disks Pool.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DisksPoolState()
        {
        }
    }
}
