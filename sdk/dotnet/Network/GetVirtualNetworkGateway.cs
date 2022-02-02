// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetVirtualNetworkGateway
    {
        /// <summary>
        /// Use this data source to access information about an existing Virtual Network Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Network.GetVirtualNetworkGateway.InvokeAsync(new Azure.Network.GetVirtualNetworkGatewayArgs
        ///         {
        ///             Name = "production",
        ///             ResourceGroupName = "networking",
        ///         }));
        ///         this.VirtualNetworkGatewayId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("virtualNetworkGatewayId")]
        ///     public Output&lt;string&gt; VirtualNetworkGatewayId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualNetworkGatewayResult> InvokeAsync(GetVirtualNetworkGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkGatewayResult>("azure:network/getVirtualNetworkGateway:getVirtualNetworkGateway", args ?? new GetVirtualNetworkGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Virtual Network Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Network.GetVirtualNetworkGateway.InvokeAsync(new Azure.Network.GetVirtualNetworkGatewayArgs
        ///         {
        ///             Name = "production",
        ///             ResourceGroupName = "networking",
        ///         }));
        ///         this.VirtualNetworkGatewayId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("virtualNetworkGatewayId")]
        ///     public Output&lt;string&gt; VirtualNetworkGatewayId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVirtualNetworkGatewayResult> Invoke(GetVirtualNetworkGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayResult>("azure:network/getVirtualNetworkGateway:getVirtualNetworkGateway", args ?? new GetVirtualNetworkGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Virtual Network Gateway.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network Gateway is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkGatewayArgs()
        {
        }
    }

    public sealed class GetVirtualNetworkGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Virtual Network Gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Virtual Network Gateway is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVirtualNetworkGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkGatewayResult
    {
        /// <summary>
        /// Is this an Active-Active Gateway?
        /// </summary>
        public readonly bool ActiveActive;
        public readonly ImmutableArray<Outputs.GetVirtualNetworkGatewayBgpSettingResult> BgpSettings;
        public readonly ImmutableArray<Outputs.GetVirtualNetworkGatewayCustomRouteResult> CustomRoutes;
        /// <summary>
        /// The ID of the local network gateway
        /// through which outbound Internet traffic from the virtual network in which the
        /// gateway is created will be routed (*forced tunneling*). Refer to the
        /// [Azure documentation on forced tunneling](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-forced-tunneling-rm).
        /// </summary>
        public readonly string DefaultLocalNetworkGatewayId;
        /// <summary>
        /// Will BGP (Border Gateway Protocol) will be enabled
        /// for this Virtual Network Gateway.
        /// </summary>
        public readonly bool EnableBgp;
        /// <summary>
        /// The Generation of the Virtual Network Gateway.
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or two `ip_configuration` blocks documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualNetworkGatewayIpConfigurationResult> IpConfigurations;
        /// <summary>
        /// The location/region where the Virtual Network Gateway is located.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The user-defined name of the revoked certificate.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether a private IP will be used for this  gateway for connections.
        /// </summary>
        public readonly bool PrivateIpAddressEnabled;
        public readonly string ResourceGroupName;
        /// <summary>
        /// Configuration of the size and capacity of the Virtual Network Gateway.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The type of the Virtual Network Gateway.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A `vpn_client_configuration` block which is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualNetworkGatewayVpnClientConfigurationResult> VpnClientConfigurations;
        /// <summary>
        /// The routing type of the Virtual Network Gateway.
        /// </summary>
        public readonly string VpnType;

        [OutputConstructor]
        private GetVirtualNetworkGatewayResult(
            bool activeActive,

            ImmutableArray<Outputs.GetVirtualNetworkGatewayBgpSettingResult> bgpSettings,

            ImmutableArray<Outputs.GetVirtualNetworkGatewayCustomRouteResult> customRoutes,

            string defaultLocalNetworkGatewayId,

            bool enableBgp,

            string generation,

            string id,

            ImmutableArray<Outputs.GetVirtualNetworkGatewayIpConfigurationResult> ipConfigurations,

            string location,

            string name,

            bool privateIpAddressEnabled,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags,

            string type,

            ImmutableArray<Outputs.GetVirtualNetworkGatewayVpnClientConfigurationResult> vpnClientConfigurations,

            string vpnType)
        {
            ActiveActive = activeActive;
            BgpSettings = bgpSettings;
            CustomRoutes = customRoutes;
            DefaultLocalNetworkGatewayId = defaultLocalNetworkGatewayId;
            EnableBgp = enableBgp;
            Generation = generation;
            Id = id;
            IpConfigurations = ipConfigurations;
            Location = location;
            Name = name;
            PrivateIpAddressEnabled = privateIpAddressEnabled;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            Type = type;
            VpnClientConfigurations = vpnClientConfigurations;
            VpnType = vpnType;
        }
    }
}
