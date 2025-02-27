// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages a Bastion Host.
    /// 
    /// ## Example Usage
    /// 
    /// This example deploys an Azure Bastion Host Instance to a target virtual network.
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
    ///                 "192.168.1.0/24",
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
    ///                 "192.168.1.224/27",
    ///             },
    ///         });
    ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AllocationMethod = "Static",
    ///             Sku = "Standard",
    ///         });
    ///         var exampleBastionHost = new Azure.Compute.BastionHost("exampleBastionHost", new Azure.Compute.BastionHostArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             IpConfiguration = new Azure.Compute.Inputs.BastionHostIpConfigurationArgs
    ///             {
    ///                 Name = "configuration",
    ///                 SubnetId = exampleSubnet.Id,
    ///                 PublicIpAddressId = examplePublicIp.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bastion Hosts can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:compute/bastionHost:BastionHost example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
    /// ```
    /// </summary>
    [AzureResourceType("azure:compute/bastionHost:BastionHost")]
    public partial class BastionHost : Pulumi.CustomResource
    {
        /// <summary>
        /// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
        /// </summary>
        [Output("copyPasteEnabled")]
        public Output<bool?> CopyPasteEnabled { get; private set; } = null!;

        /// <summary>
        /// The FQDN for the Bastion Host.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Output("fileCopyEnabled")]
        public Output<bool?> FileCopyEnabled { get; private set; } = null!;

        /// <summary>
        /// A `ip_configuration` block as defined below.
        /// </summary>
        [Output("ipConfiguration")]
        public Output<Outputs.BastionHostIpConfiguration?> IpConfiguration { get; private set; } = null!;

        /// <summary>
        /// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Output("ipConnectEnabled")]
        public Output<bool?> IpConnectEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Bastion Host.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
        /// </summary>
        [Output("scaleUnits")]
        public Output<int?> ScaleUnits { get; private set; } = null!;

        /// <summary>
        /// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Output("shareableLinkEnabled")]
        public Output<bool?> ShareableLinkEnabled { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Output("tunnelingEnabled")]
        public Output<bool?> TunnelingEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a BastionHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BastionHost(string name, BastionHostArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/bastionHost:BastionHost", name, args ?? new BastionHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BastionHost(string name, Input<string> id, BastionHostState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/bastionHost:BastionHost", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BastionHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BastionHost Get(string name, Input<string> id, BastionHostState? state = null, CustomResourceOptions? options = null)
        {
            return new BastionHost(name, id, state, options);
        }
    }

    public sealed class BastionHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
        /// </summary>
        [Input("copyPasteEnabled")]
        public Input<bool>? CopyPasteEnabled { get; set; }

        /// <summary>
        /// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("fileCopyEnabled")]
        public Input<bool>? FileCopyEnabled { get; set; }

        /// <summary>
        /// A `ip_configuration` block as defined below.
        /// </summary>
        [Input("ipConfiguration")]
        public Input<Inputs.BastionHostIpConfigurationArgs>? IpConfiguration { get; set; }

        /// <summary>
        /// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("ipConnectEnabled")]
        public Input<bool>? IpConnectEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bastion Host.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
        /// </summary>
        [Input("scaleUnits")]
        public Input<int>? ScaleUnits { get; set; }

        /// <summary>
        /// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("shareableLinkEnabled")]
        public Input<bool>? ShareableLinkEnabled { get; set; }

        /// <summary>
        /// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("tunnelingEnabled")]
        public Input<bool>? TunnelingEnabled { get; set; }

        public BastionHostArgs()
        {
        }
    }

    public sealed class BastionHostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is Copy/Paste feature enabled for the Bastion Host. Defaults to `true`.
        /// </summary>
        [Input("copyPasteEnabled")]
        public Input<bool>? CopyPasteEnabled { get; set; }

        /// <summary>
        /// The FQDN for the Bastion Host.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Is File Copy feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("fileCopyEnabled")]
        public Input<bool>? FileCopyEnabled { get; set; }

        /// <summary>
        /// A `ip_configuration` block as defined below.
        /// </summary>
        [Input("ipConfiguration")]
        public Input<Inputs.BastionHostIpConfigurationGetArgs>? IpConfiguration { get; set; }

        /// <summary>
        /// Is IP Connect feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("ipConnectEnabled")]
        public Input<bool>? IpConnectEnabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.  Review [Azure Bastion Host FAQ](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq) for supported locations.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Bastion Host.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The number of scale units with which to provision the Bastion Host. Possible values are between `2` and `50`. Defaults to `2`.
        /// </summary>
        [Input("scaleUnits")]
        public Input<int>? ScaleUnits { get; set; }

        /// <summary>
        /// Is Shareable Link feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("shareableLinkEnabled")]
        public Input<bool>? ShareableLinkEnabled { get; set; }

        /// <summary>
        /// The SKU of the Bastion Host. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Is Tunneling feature enabled for the Bastion Host. Defaults to `false`.
        /// </summary>
        [Input("tunnelingEnabled")]
        public Input<bool>? TunnelingEnabled { get; set; }

        public BastionHostState()
        {
        }
    }
}
