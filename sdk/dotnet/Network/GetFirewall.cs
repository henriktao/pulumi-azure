// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Network
{
    public static class GetFirewall
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Firewall.
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
        ///         var example = Output.Create(Azure.Network.GetFirewall.InvokeAsync(new Azure.Network.GetFirewallArgs
        ///         {
        ///             Name = "firewall1",
        ///             ResourceGroupName = "firewall-RG",
        ///         }));
        ///         this.FirewallPrivateIp = example.Apply(example =&gt; example.IpConfigurations?[0]?.PrivateIpAddress);
        ///     }
        /// 
        ///     [Output("firewallPrivateIp")]
        ///     public Output&lt;string&gt; FirewallPrivateIp { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFirewallResult> InvokeAsync(GetFirewallArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFirewallResult>("azure:network/getFirewall:getFirewall", args ?? new GetFirewallArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Azure Firewall.
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
        ///         var example = Output.Create(Azure.Network.GetFirewall.InvokeAsync(new Azure.Network.GetFirewallArgs
        ///         {
        ///             Name = "firewall1",
        ///             ResourceGroupName = "firewall-RG",
        ///         }));
        ///         this.FirewallPrivateIp = example.Apply(example =&gt; example.IpConfigurations?[0]?.PrivateIpAddress);
        ///     }
        /// 
        ///     [Output("firewallPrivateIp")]
        ///     public Output&lt;string&gt; FirewallPrivateIp { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFirewallResult> Invoke(GetFirewallInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFirewallResult>("azure:network/getFirewall:getFirewall", args ?? new GetFirewallInvokeArgs(), options.WithVersion());
    }


    public sealed class GetFirewallArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure Firewall.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Azure Firewall exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFirewallArgs()
        {
        }
    }

    public sealed class GetFirewallInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure Firewall.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Azure Firewall exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFirewallInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFirewallResult
    {
        /// <summary>
        /// The list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// The ID of the Firewall Policy applied to the Azure Firewall.
        /// </summary>
        public readonly string FirewallPolicyId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A `ip_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallIpConfigurationResult> IpConfigurations;
        /// <summary>
        /// The Azure location where the Azure Firewall exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// A `management_ip_configuration` block as defined below, which allows force-tunnelling of traffic to be performed by the firewall.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallManagementIpConfigurationResult> ManagementIpConfigurations;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The sku name of the Azure Firewall.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// The sku tier of the Azure Firewall.
        /// </summary>
        public readonly string SkuTier;
        /// <summary>
        /// A mapping of tags assigned to the Azure Firewall.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The operation mode for threat intelligence-based filtering.
        /// </summary>
        public readonly string ThreatIntelMode;
        /// <summary>
        /// A `virtual_hub` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallVirtualHubResult> VirtualHubs;
        /// <summary>
        /// The availability zones in which the Azure Firewall is created.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetFirewallResult(
            ImmutableArray<string> dnsServers,

            string firewallPolicyId,

            string id,

            ImmutableArray<Outputs.GetFirewallIpConfigurationResult> ipConfigurations,

            string location,

            ImmutableArray<Outputs.GetFirewallManagementIpConfigurationResult> managementIpConfigurations,

            string name,

            string resourceGroupName,

            string skuName,

            string skuTier,

            ImmutableDictionary<string, string> tags,

            string threatIntelMode,

            ImmutableArray<Outputs.GetFirewallVirtualHubResult> virtualHubs,

            ImmutableArray<string> zones)
        {
            DnsServers = dnsServers;
            FirewallPolicyId = firewallPolicyId;
            Id = id;
            IpConfigurations = ipConfigurations;
            Location = location;
            ManagementIpConfigurations = managementIpConfigurations;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            SkuTier = skuTier;
            Tags = tags;
            ThreatIntelMode = threatIntelMode;
            VirtualHubs = virtualHubs;
            Zones = zones;
        }
    }
}
