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
    /// Manages a VPN Gateway Nat Rule.
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
    ///         var exampleVirtualWan = new Azure.Network.VirtualWan("exampleVirtualWan", new Azure.Network.VirtualWanArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///         });
    ///         var exampleVirtualHub = new Azure.Network.VirtualHub("exampleVirtualHub", new Azure.Network.VirtualHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AddressPrefix = "10.0.1.0/24",
    ///             VirtualWanId = exampleVirtualWan.Id,
    ///         });
    ///         var exampleVpnGateway = new Azure.Network.VpnGateway("exampleVpnGateway", new Azure.Network.VpnGatewayArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VirtualHubId = exampleVirtualHub.Id,
    ///         });
    ///         var exampleVnpGatewayNatRule = new Azure.Network.VnpGatewayNatRule("exampleVnpGatewayNatRule", new Azure.Network.VnpGatewayNatRuleArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             VpnGatewayId = exampleVpnGateway.Id,
    ///             ExternalAddressSpaceMappings = 
    ///             {
    ///                 "192.168.21.0/26",
    ///             },
    ///             InternalAddressSpaceMappings = 
    ///             {
    ///                 "10.4.0.0/26",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPN Gateway Nat Rules can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:network/vnpGatewayNatRule:VnpGatewayNatRule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnGateways/vpnGateway1/natRules/natRule1
    /// ```
    /// </summary>
    [AzureResourceType("azure:network/vnpGatewayNatRule:VnpGatewayNatRule")]
    public partial class VnpGatewayNatRule : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
        /// </summary>
        [Output("externalAddressSpaceMappings")]
        public Output<ImmutableArray<string>> ExternalAddressSpaceMappings { get; private set; } = null!;

        /// <summary>
        /// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
        /// </summary>
        [Output("internalAddressSpaceMappings")]
        public Output<ImmutableArray<string>> InternalAddressSpaceMappings { get; private set; } = null!;

        /// <summary>
        /// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
        /// </summary>
        [Output("ipConfigurationId")]
        public Output<string?> IpConfigurationId { get; private set; } = null!;

        /// <summary>
        /// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a VnpGatewayNatRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VnpGatewayNatRule(string name, VnpGatewayNatRuleArgs args, CustomResourceOptions? options = null)
            : base("azure:network/vnpGatewayNatRule:VnpGatewayNatRule", name, args ?? new VnpGatewayNatRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VnpGatewayNatRule(string name, Input<string> id, VnpGatewayNatRuleState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/vnpGatewayNatRule:VnpGatewayNatRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VnpGatewayNatRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VnpGatewayNatRule Get(string name, Input<string> id, VnpGatewayNatRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new VnpGatewayNatRule(name, id, state, options);
        }
    }

    public sealed class VnpGatewayNatRuleArgs : Pulumi.ResourceArgs
    {
        [Input("externalAddressSpaceMappings", required: true)]
        private InputList<string>? _externalAddressSpaceMappings;

        /// <summary>
        /// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
        /// </summary>
        public InputList<string> ExternalAddressSpaceMappings
        {
            get => _externalAddressSpaceMappings ?? (_externalAddressSpaceMappings = new InputList<string>());
            set => _externalAddressSpaceMappings = value;
        }

        [Input("internalAddressSpaceMappings", required: true)]
        private InputList<string>? _internalAddressSpaceMappings;

        /// <summary>
        /// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
        /// </summary>
        public InputList<string> InternalAddressSpaceMappings
        {
            get => _internalAddressSpaceMappings ?? (_internalAddressSpaceMappings = new InputList<string>());
            set => _internalAddressSpaceMappings = value;
        }

        /// <summary>
        /// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
        /// </summary>
        [Input("ipConfigurationId")]
        public Input<string>? IpConfigurationId { get; set; }

        /// <summary>
        /// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public VnpGatewayNatRuleArgs()
        {
        }
    }

    public sealed class VnpGatewayNatRuleState : Pulumi.ResourceArgs
    {
        [Input("externalAddressSpaceMappings")]
        private InputList<string>? _externalAddressSpaceMappings;

        /// <summary>
        /// A list of CIDR Ranges which are used for external mapping of the VPN Gateway Nat Rule.
        /// </summary>
        public InputList<string> ExternalAddressSpaceMappings
        {
            get => _externalAddressSpaceMappings ?? (_externalAddressSpaceMappings = new InputList<string>());
            set => _externalAddressSpaceMappings = value;
        }

        [Input("internalAddressSpaceMappings")]
        private InputList<string>? _internalAddressSpaceMappings;

        /// <summary>
        /// A list of CIDR Ranges which are used for internal mapping of the VPN Gateway Nat Rule.
        /// </summary>
        public InputList<string> InternalAddressSpaceMappings
        {
            get => _internalAddressSpaceMappings ?? (_internalAddressSpaceMappings = new InputList<string>());
            set => _internalAddressSpaceMappings = value;
        }

        /// <summary>
        /// The ID of the IP Configuration this VPN Gateway Nat Rule applies to. Possible values are `Instance0` and `Instance1`.
        /// </summary>
        [Input("ipConfigurationId")]
        public Input<string>? IpConfigurationId { get; set; }

        /// <summary>
        /// The source Nat direction of the VPN Nat. Possible values are `EgressSnat` and `IngressSnat`. Defaults to `EgressSnat`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name which should be used for this VPN Gateway Nat Rule. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Name of the Resource Group in which this VPN Gateway Nat Rule should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The type of the VPN Gateway Nat Rule. Possible values are `Dynamic` and `Static`. Defaults to `Static`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The ID of the VPN Gateway that this VPN Gateway Nat Rule belongs to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public VnpGatewayNatRuleState()
        {
        }
    }
}
