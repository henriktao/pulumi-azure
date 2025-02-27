// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hsm
{
    /// <summary>
    /// Manages a Dedicated Hardware Security Module.
    /// 
    /// &gt; **Note:** Before using this resource, it's required to submit the request of registering the providers and features with Azure CLI `az provider register --namespace Microsoft.HardwareSecurityModules &amp;&amp; az feature register --namespace Microsoft.HardwareSecurityModules --name AzureDedicatedHSM &amp;&amp; az provider register --namespace Microsoft.Network &amp;&amp; az feature register --namespace Microsoft.Network --name AllowBaremetalServers` and ask service team (hsmrequest@microsoft.com) to approve. See more details from https://docs.microsoft.com/en-us/azure/dedicated-hsm/tutorial-deploy-hsm-cli#prerequisites.
    /// 
    /// &gt; **Note:** If the quota is not enough in some region, please submit the quota request to service team.
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
    ///                 "10.2.0.0/16",
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
    ///                 "10.2.0.0/24",
    ///             },
    ///         });
    ///         var example2 = new Azure.Network.Subnet("example2", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.2.1.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "first",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.HardwareSecurityModules/dedicatedHSMs",
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/networkinterfaces/*",
    ///                             "Microsoft.Network/virtualNetworks/subnets/join/action",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var example3 = new Azure.Network.Subnet("example3", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.2.255.0/26",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Dynamic",
    ///         });
    ///         var exampleVirtualNetworkGateway = new Azure.Network.VirtualNetworkGateway("exampleVirtualNetworkGateway", new Azure.Network.VirtualNetworkGatewayArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Type = "ExpressRoute",
    ///             VpnType = "PolicyBased",
    ///             Sku = "Standard",
    ///             IpConfigurations = 
    ///             {
    ///                 new Azure.Network.Inputs.VirtualNetworkGatewayIpConfigurationArgs
    ///                 {
    ///                     PublicIpAddressId = examplePublicIp.Id,
    ///                     PrivateIpAddressAllocation = "Dynamic",
    ///                     SubnetId = example3.Id,
    ///                 },
    ///             },
    ///         });
    ///         var exampleModule = new Azure.Hsm.Module("exampleModule", new Azure.Hsm.ModuleArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SkuName = "SafeNet Luna Network HSM A790",
    ///             NetworkProfile = new Azure.Hsm.Inputs.ModuleNetworkProfileArgs
    ///             {
    ///                 NetworkInterfacePrivateIpAddresses = 
    ///                 {
    ///                     "10.2.1.8",
    ///                 },
    ///                 SubnetId = example2.Id,
    ///             },
    ///             StampId = "stamp2",
    ///             Tags = 
    ///             {
    ///                 { "env", "Test" },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleVirtualNetworkGateway,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dedicated Hardware Security Module can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:hsm/module:Module example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
    /// ```
    /// </summary>
    [AzureResourceType("azure:hsm/module:Module")]
    public partial class Module : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.ModuleNetworkProfile> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("stampId")]
        public Output<string?> StampId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Module resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Module(string name, ModuleArgs args, CustomResourceOptions? options = null)
            : base("azure:hsm/module:Module", name, args ?? new ModuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Module(string name, Input<string> id, ModuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:hsm/module:Module", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Module resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Module Get(string name, Input<string> id, ModuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Module(name, id, state, options);
        }
    }

    public sealed class ModuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Input("networkProfile", required: true)]
        public Input<Inputs.ModuleNetworkProfileArgs> NetworkProfile { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        /// <summary>
        /// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("stampId")]
        public Input<string>? StampId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ModuleArgs()
        {
        }
    }

    public sealed class ModuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Region where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this Dedicated Hardware Security Module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `network_profile` block as defined below.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.ModuleNetworkProfileGetArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Dedicated Hardware Security Module should exist. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The sku name of the dedicated hardware security module. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        /// <summary>
        /// The ID of the stamp. Possible values are `stamp1` or `stamp2`. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        [Input("stampId")]
        public Input<string>? StampId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the Dedicated Hardware Security Module.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// The Dedicated Hardware Security Module zones. Changing this forces a new Dedicated Hardware Security Module to be created.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ModuleState()
        {
        }
    }
}
