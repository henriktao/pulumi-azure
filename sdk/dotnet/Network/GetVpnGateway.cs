// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetVpnGateway
    {
        /// <summary>
        /// Use this data source to access information about an existing VPN Gateway within a Virtual Hub.
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
        ///         var example = Output.Create(Azure.Network.GetVpnGateway.InvokeAsync(new Azure.Network.GetVpnGatewayArgs
        ///         {
        ///             Name = "existing-local-vpn_gateway",
        ///             ResourceGroupName = "existing-vpn_gateway",
        ///         }));
        ///         this.AzurermVpnGatewayId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("azurermVpnGatewayId")]
        ///     public Output&lt;string&gt; AzurermVpnGatewayId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpnGatewayResult> InvokeAsync(GetVpnGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnGatewayResult>("azure:network/getVpnGateway:getVpnGateway", args ?? new GetVpnGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing VPN Gateway within a Virtual Hub.
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
        ///         var example = Output.Create(Azure.Network.GetVpnGateway.InvokeAsync(new Azure.Network.GetVpnGatewayArgs
        ///         {
        ///             Name = "existing-local-vpn_gateway",
        ///             ResourceGroupName = "existing-vpn_gateway",
        ///         }));
        ///         this.AzurermVpnGatewayId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("azurermVpnGatewayId")]
        ///     public Output&lt;string&gt; AzurermVpnGatewayId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpnGatewayResult> Invoke(GetVpnGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpnGatewayResult>("azure:network/getVpnGateway:getVpnGateway", args ?? new GetVpnGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpnGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the VPN Gateway.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the VPN Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVpnGatewayArgs()
        {
        }
    }

    public sealed class GetVpnGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the VPN Gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the VPN Gateway exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetVpnGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpnGatewayResult
    {
        /// <summary>
        /// A `bgp_settings` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpnGatewayBgpSettingResult> BgpSettings;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the VPN Gateway exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Scale Unit of this VPN Gateway.
        /// </summary>
        public readonly int ScaleUnit;
        /// <summary>
        /// A mapping of tags assigned to the VPN Gateway.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The ID of the Virtual Hub within which this VPN Gateway has been created.
        /// </summary>
        public readonly string VirtualHubId;

        [OutputConstructor]
        private GetVpnGatewayResult(
            ImmutableArray<Outputs.GetVpnGatewayBgpSettingResult> bgpSettings,

            string id,

            string location,

            string name,

            string resourceGroupName,

            int scaleUnit,

            ImmutableDictionary<string, string> tags,

            string virtualHubId)
        {
            BgpSettings = bgpSettings;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ScaleUnit = scaleUnit;
            Tags = tags;
            VirtualHubId = virtualHubId;
        }
    }
}
