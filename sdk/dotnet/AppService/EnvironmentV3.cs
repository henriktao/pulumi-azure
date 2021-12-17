// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Manages a 3rd Generation (v3) App Service Environment.
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
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AddressSpaces = 
    ///             {
    ///                 "10.0.0.0/16",
    ///             },
    ///         });
    ///         var exampleSubnet = new Azure.Network.Subnet("exampleSubnet", new Azure.Network.SubnetArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualNetworkName = exampleVirtualNetwork.Name,
    ///             AddressPrefixes = 
    ///             {
    ///                 "10.0.2.0/24",
    ///             },
    ///             Delegations = 
    ///             {
    ///                 new Azure.Network.Inputs.SubnetDelegationArgs
    ///                 {
    ///                     Name = "delegation",
    ///                     ServiceDelegation = new Azure.Network.Inputs.SubnetDelegationServiceDelegationArgs
    ///                     {
    ///                         Name = "Microsoft.Web/hostingEnvironments",
    ///                         Actions = 
    ///                         {
    ///                             "Microsoft.Network/virtualNetworks/subnets/action",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleEnvironmentV3 = new Azure.AppService.EnvironmentV3("exampleEnvironmentV3", new Azure.AppService.EnvironmentV3Args
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             SubnetId = exampleSubnet.Id,
    ///             ClusterSettings = 
    ///             {
    ///                 new Azure.AppService.Inputs.EnvironmentV3ClusterSettingArgs
    ///                 {
    ///                     Name = "DisableTls1.0",
    ///                     Value = "1",
    ///                 },
    ///                 new Azure.AppService.Inputs.EnvironmentV3ClusterSettingArgs
    ///                 {
    ///                     Name = "InternalEncryption",
    ///                     Value = "true",
    ///                 },
    ///                 new Azure.AppService.Inputs.EnvironmentV3ClusterSettingArgs
    ///                 {
    ///                     Name = "FrontEndSSLCipherSuiteOrder",
    ///                     Value = "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "env", "production" },
    ///                 { "terraformed", "true" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A 3rd Generation (v3) App Service Environment can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:appservice/environmentV3:EnvironmentV3 myAppServiceEnv /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Web/hostingEnvironments/myAppServiceEnv
    /// ```
    /// </summary>
    [AzureResourceType("azure:appservice/environmentV3:EnvironmentV3")]
    public partial class EnvironmentV3 : Pulumi.CustomResource
    {
        /// <summary>
        /// Should new Private Endpoint Connections be allowed. Defaults to `true`.
        /// </summary>
        [Output("allowNewPrivateEndpointConnections")]
        public Output<bool?> AllowNewPrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Zero or more `cluster_setting` blocks as defined below.
        /// </summary>
        [Output("clusterSettings")]
        public Output<ImmutableArray<Outputs.EnvironmentV3ClusterSetting>> ClusterSettings { get; private set; } = null!;

        /// <summary>
        /// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dedicatedHostCount")]
        public Output<int?> DedicatedHostCount { get; private set; } = null!;

        /// <summary>
        /// the DNS suffix for this App Service Environment V3.
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// The external inbound IP addresses of the App Service Environment V3.
        /// </summary>
        [Output("externalInboundIpAddresses")]
        public Output<ImmutableArray<string>> ExternalInboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// An Inbound Network Dependencies block as defined below.
        /// </summary>
        [Output("inboundNetworkDependencies")]
        public Output<ImmutableArray<Outputs.EnvironmentV3InboundNetworkDependency>> InboundNetworkDependencies { get; private set; } = null!;

        /// <summary>
        /// The internal inbound IP addresses of the App Service Environment V3.
        /// </summary>
        [Output("internalInboundIpAddresses")]
        public Output<ImmutableArray<string>> InternalInboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
        /// </summary>
        [Output("internalLoadBalancingMode")]
        public Output<string?> InternalLoadBalancingMode { get; private set; } = null!;

        /// <summary>
        /// The number of IP SSL addresses reserved for the App Service Environment V3.
        /// </summary>
        [Output("ipSslAddressCount")]
        public Output<int> IpSslAddressCount { get; private set; } = null!;

        /// <summary>
        /// Outbound addresses of Linux based Apps in this App Service Environment V3
        /// </summary>
        [Output("linuxOutboundIpAddresses")]
        public Output<ImmutableArray<string>> LinuxOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The location where the App Service Environment exists.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the App Service Environment. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Pricing tier for the front end instances.
        /// </summary>
        [Output("pricingTier")]
        public Output<string> PricingTier { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Outbound addresses of Windows based Apps in this App Service Environment V3.
        /// </summary>
        [Output("windowsOutboundIpAddresses")]
        public Output<ImmutableArray<string>> WindowsOutboundIpAddresses { get; private set; } = null!;

        [Output("zoneRedundant")]
        public Output<bool?> ZoneRedundant { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentV3(string name, EnvironmentV3Args args, CustomResourceOptions? options = null)
            : base("azure:appservice/environmentV3:EnvironmentV3", name, args ?? new EnvironmentV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentV3(string name, Input<string> id, EnvironmentV3State? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/environmentV3:EnvironmentV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnvironmentV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentV3 Get(string name, Input<string> id, EnvironmentV3State? state = null, CustomResourceOptions? options = null)
        {
            return new EnvironmentV3(name, id, state, options);
        }
    }

    public sealed class EnvironmentV3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should new Private Endpoint Connections be allowed. Defaults to `true`.
        /// </summary>
        [Input("allowNewPrivateEndpointConnections")]
        public Input<bool>? AllowNewPrivateEndpointConnections { get; set; }

        [Input("clusterSettings")]
        private InputList<Inputs.EnvironmentV3ClusterSettingArgs>? _clusterSettings;

        /// <summary>
        /// Zero or more `cluster_setting` blocks as defined below.
        /// </summary>
        public InputList<Inputs.EnvironmentV3ClusterSettingArgs> ClusterSettings
        {
            get => _clusterSettings ?? (_clusterSettings = new InputList<Inputs.EnvironmentV3ClusterSettingArgs>());
            set => _clusterSettings = value;
        }

        /// <summary>
        /// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostCount")]
        public Input<int>? DedicatedHostCount { get; set; }

        /// <summary>
        /// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
        /// </summary>
        [Input("internalLoadBalancingMode")]
        public Input<string>? InternalLoadBalancingMode { get; set; }

        /// <summary>
        /// The name of the App Service Environment. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public EnvironmentV3Args()
        {
        }
    }

    public sealed class EnvironmentV3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should new Private Endpoint Connections be allowed. Defaults to `true`.
        /// </summary>
        [Input("allowNewPrivateEndpointConnections")]
        public Input<bool>? AllowNewPrivateEndpointConnections { get; set; }

        [Input("clusterSettings")]
        private InputList<Inputs.EnvironmentV3ClusterSettingGetArgs>? _clusterSettings;

        /// <summary>
        /// Zero or more `cluster_setting` blocks as defined below.
        /// </summary>
        public InputList<Inputs.EnvironmentV3ClusterSettingGetArgs> ClusterSettings
        {
            get => _clusterSettings ?? (_clusterSettings = new InputList<Inputs.EnvironmentV3ClusterSettingGetArgs>());
            set => _clusterSettings = value;
        }

        /// <summary>
        /// This ASEv3 should use dedicated Hosts. Possible vales are `2`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dedicatedHostCount")]
        public Input<int>? DedicatedHostCount { get; set; }

        /// <summary>
        /// the DNS suffix for this App Service Environment V3.
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        [Input("externalInboundIpAddresses")]
        private InputList<string>? _externalInboundIpAddresses;

        /// <summary>
        /// The external inbound IP addresses of the App Service Environment V3.
        /// </summary>
        public InputList<string> ExternalInboundIpAddresses
        {
            get => _externalInboundIpAddresses ?? (_externalInboundIpAddresses = new InputList<string>());
            set => _externalInboundIpAddresses = value;
        }

        [Input("inboundNetworkDependencies")]
        private InputList<Inputs.EnvironmentV3InboundNetworkDependencyGetArgs>? _inboundNetworkDependencies;

        /// <summary>
        /// An Inbound Network Dependencies block as defined below.
        /// </summary>
        public InputList<Inputs.EnvironmentV3InboundNetworkDependencyGetArgs> InboundNetworkDependencies
        {
            get => _inboundNetworkDependencies ?? (_inboundNetworkDependencies = new InputList<Inputs.EnvironmentV3InboundNetworkDependencyGetArgs>());
            set => _inboundNetworkDependencies = value;
        }

        [Input("internalInboundIpAddresses")]
        private InputList<string>? _internalInboundIpAddresses;

        /// <summary>
        /// The internal inbound IP addresses of the App Service Environment V3.
        /// </summary>
        public InputList<string> InternalInboundIpAddresses
        {
            get => _internalInboundIpAddresses ?? (_internalInboundIpAddresses = new InputList<string>());
            set => _internalInboundIpAddresses = value;
        }

        /// <summary>
        /// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment. Possible values are `None` (for an External VIP Type), and `"Web, Publishing"` (for an Internal VIP Type). Defaults to `None`.
        /// </summary>
        [Input("internalLoadBalancingMode")]
        public Input<string>? InternalLoadBalancingMode { get; set; }

        /// <summary>
        /// The number of IP SSL addresses reserved for the App Service Environment V3.
        /// </summary>
        [Input("ipSslAddressCount")]
        public Input<int>? IpSslAddressCount { get; set; }

        [Input("linuxOutboundIpAddresses")]
        private InputList<string>? _linuxOutboundIpAddresses;

        /// <summary>
        /// Outbound addresses of Linux based Apps in this App Service Environment V3
        /// </summary>
        public InputList<string> LinuxOutboundIpAddresses
        {
            get => _linuxOutboundIpAddresses ?? (_linuxOutboundIpAddresses = new InputList<string>());
            set => _linuxOutboundIpAddresses = value;
        }

        /// <summary>
        /// The location where the App Service Environment exists.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the App Service Environment. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Pricing tier for the front end instances.
        /// </summary>
        [Input("pricingTier")]
        public Input<string>? PricingTier { get; set; }

        /// <summary>
        /// The name of the Resource Group where the App Service Environment exists. Defaults to the Resource Group of the Subnet (specified by `subnet_id`).
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the Subnet which the App Service Environment should be connected to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("windowsOutboundIpAddresses")]
        private InputList<string>? _windowsOutboundIpAddresses;

        /// <summary>
        /// Outbound addresses of Windows based Apps in this App Service Environment V3.
        /// </summary>
        public InputList<string> WindowsOutboundIpAddresses
        {
            get => _windowsOutboundIpAddresses ?? (_windowsOutboundIpAddresses = new InputList<string>());
            set => _windowsOutboundIpAddresses = value;
        }

        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public EnvironmentV3State()
        {
        }
    }
}
