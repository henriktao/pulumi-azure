// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a virtual network including any configured subnets. Each subnet can
    /// optionally be configured with a security group to be associated with the subnet.
    /// 
    /// &gt; **NOTE on Virtual Networks and Subnet's:** This provider currently
    /// provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
    /// At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
    /// **NOTE on Virtual Networks and DNS Servers:** This provider currently provides both a standalone virtual network DNS Servers resource, and allows for DNS servers to be defined in-line within the Virtual Network resource.
    /// At this time you cannot use a Virtual Network with in-line DNS servers in conjunction with any Virtual Network DNS Servers resources. Doing so will cause a conflict of Virtual Network DNS Servers configurations and will overwrite virtual networks DNS servers.
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
    ///         var exampleNetworkSecurityGroup = new Azure.Network.NetworkSecurityGroup("exampleNetworkSecurityGroup", new Azure.Network.NetworkSecurityGroupArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///         var exampleVirtualNetwork = new Azure.Network.VirtualNetwork("exampleVirtualNetwork", new Azure.Network.VirtualNetworkArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///             DnsServers = 
    ///             {
    ///                 "10.0.0.4",
    ///                 "10.0.0.5",
    ///             },
    ///             Subnets = 
    ///             {
    ///                 new Azure.Network.Inputs.VirtualNetworkSubnetArgs
    ///                 {
    ///                     Name = "subnet1",
    ///                     AddressPrefix = "10.0.1.0/24",
    ///                 },
    ///                 new Azure.Network.Inputs.VirtualNetworkSubnetArgs
    ///                 {
    ///                     Name = "subnet2",
    ///                     AddressPrefix = "10.0.2.0/24",
    ///                     SecurityGroup = exampleNetworkSecurityGroup.Id,
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "environment", "Production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual Networks can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/virtualNetwork:VirtualNetwork exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
    /// ```
    /// </summary>
    [AzureResourceType("azure:network/virtualNetwork:VirtualNetwork")]
    public partial class VirtualNetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// The address space that is used the virtual network. You can supply more than one address space.
        /// </summary>
        [Output("addressSpaces")]
        public Output<ImmutableArray<string>> AddressSpaces { get; private set; } = null!;

        /// <summary>
        /// The BGP community attribute in format `&lt;as-number&gt;:&lt;community-value&gt;`.
        /// </summary>
        [Output("bgpCommunity")]
        public Output<string?> BgpCommunity { get; private set; } = null!;

        /// <summary>
        /// A `ddos_protection_plan` block as documented below.
        /// </summary>
        [Output("ddosProtectionPlan")]
        public Output<Outputs.VirtualNetworkDdosProtectionPlan?> DdosProtectionPlan { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses of DNS servers
        /// </summary>
        [Output("dnsServers")]
        public Output<ImmutableArray<string>> DnsServers { get; private set; } = null!;

        /// <summary>
        /// The flow timeout in minutes for the Virtual Network, which is used to enable connection tracking for intra-VM flows. Possible values are between `4` and `30` minutes.
        /// </summary>
        [Output("flowTimeoutInMinutes")]
        public Output<int?> FlowTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// The GUID of the virtual network.
        /// </summary>
        [Output("guid")]
        public Output<string> Guid { get; private set; } = null!;

        /// <summary>
        /// The location/region where the virtual network is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the virtual network. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the virtual network.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.VirtualNetworkSubnet>> Subnets { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("vmProtectionEnabled")]
        public Output<bool?> VmProtectionEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetwork:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, VirtualNetworkState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/virtualNetwork:VirtualNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetwork Get(string name, Input<string> id, VirtualNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualNetwork(name, id, state, options);
        }
    }

    public sealed class VirtualNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("addressSpaces", required: true)]
        private InputList<string>? _addressSpaces;

        /// <summary>
        /// The address space that is used the virtual network. You can supply more than one address space.
        /// </summary>
        public InputList<string> AddressSpaces
        {
            get => _addressSpaces ?? (_addressSpaces = new InputList<string>());
            set => _addressSpaces = value;
        }

        /// <summary>
        /// The BGP community attribute in format `&lt;as-number&gt;:&lt;community-value&gt;`.
        /// </summary>
        [Input("bgpCommunity")]
        public Input<string>? BgpCommunity { get; set; }

        /// <summary>
        /// A `ddos_protection_plan` block as documented below.
        /// </summary>
        [Input("ddosProtectionPlan")]
        public Input<Inputs.VirtualNetworkDdosProtectionPlanArgs>? DdosProtectionPlan { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// List of IP addresses of DNS servers
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The flow timeout in minutes for the Virtual Network, which is used to enable connection tracking for intra-VM flows. Possible values are between `4` and `30` minutes.
        /// </summary>
        [Input("flowTimeoutInMinutes")]
        public Input<int>? FlowTimeoutInMinutes { get; set; }

        /// <summary>
        /// The location/region where the virtual network is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the virtual network.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("subnets")]
        private InputList<Inputs.VirtualNetworkSubnetArgs>? _subnets;

        /// <summary>
        /// Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.VirtualNetworkSubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.VirtualNetworkSubnetArgs>());
            set => _subnets = value;
        }

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

        [Input("vmProtectionEnabled")]
        public Input<bool>? VmProtectionEnabled { get; set; }

        public VirtualNetworkArgs()
        {
        }
    }

    public sealed class VirtualNetworkState : Pulumi.ResourceArgs
    {
        [Input("addressSpaces")]
        private InputList<string>? _addressSpaces;

        /// <summary>
        /// The address space that is used the virtual network. You can supply more than one address space.
        /// </summary>
        public InputList<string> AddressSpaces
        {
            get => _addressSpaces ?? (_addressSpaces = new InputList<string>());
            set => _addressSpaces = value;
        }

        /// <summary>
        /// The BGP community attribute in format `&lt;as-number&gt;:&lt;community-value&gt;`.
        /// </summary>
        [Input("bgpCommunity")]
        public Input<string>? BgpCommunity { get; set; }

        /// <summary>
        /// A `ddos_protection_plan` block as documented below.
        /// </summary>
        [Input("ddosProtectionPlan")]
        public Input<Inputs.VirtualNetworkDdosProtectionPlanGetArgs>? DdosProtectionPlan { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// List of IP addresses of DNS servers
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The flow timeout in minutes for the Virtual Network, which is used to enable connection tracking for intra-VM flows. Possible values are between `4` and `30` minutes.
        /// </summary>
        [Input("flowTimeoutInMinutes")]
        public Input<int>? FlowTimeoutInMinutes { get; set; }

        /// <summary>
        /// The GUID of the virtual network.
        /// </summary>
        [Input("guid")]
        public Input<string>? Guid { get; set; }

        /// <summary>
        /// The location/region where the virtual network is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual network. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the virtual network.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("subnets")]
        private InputList<Inputs.VirtualNetworkSubnetGetArgs>? _subnets;

        /// <summary>
        /// Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
        /// </summary>
        public InputList<Inputs.VirtualNetworkSubnetGetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.VirtualNetworkSubnetGetArgs>());
            set => _subnets = value;
        }

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

        [Input("vmProtectionEnabled")]
        public Input<bool>? VmProtectionEnabled { get; set; }

        public VirtualNetworkState()
        {
        }
    }
}
