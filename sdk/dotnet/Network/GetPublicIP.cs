// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetPublicIP
    {
        /// <summary>
        /// Use this data source to access information about an existing Public IP Address.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Reference An Existing)
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Network.GetPublicIP.InvokeAsync(new Azure.Network.GetPublicIPArgs
        ///         {
        ///             Name = "name_of_public_ip",
        ///             ResourceGroupName = "name_of_resource_group",
        ///         }));
        ///         this.DomainNameLabel = example.Apply(example =&gt; example.DomainNameLabel);
        ///         this.PublicIpAddress = example.Apply(example =&gt; example.IpAddress);
        ///     }
        /// 
        ///     [Output("domainNameLabel")]
        ///     public Output&lt;string&gt; DomainNameLabel { get; set; }
        ///     [Output("publicIpAddress")]
        ///     public Output&lt;string&gt; PublicIpAddress { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Retrieve The Dynamic Public IP Of A New VM)
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
        ///                 "10.0.2.0/24",
        ///             },
        ///         });
        ///         var examplePublicIp = new Azure.Network.PublicIp("examplePublicIp", new Azure.Network.PublicIpArgs
        ///         {
        ///             Location = exampleResourceGroup.Location,
        ///             ResourceGroupName = exampleResourceGroup.Name,
        ///             AllocationMethod = "Dynamic",
        ///             IdleTimeoutInMinutes = 30,
        ///             Tags = 
        ///             {
        ///                 { "environment", "test" },
        ///             },
        ///         });
        ///         var exampleNetworkInterface = new Azure.Network.NetworkInterface("exampleNetworkInterface", new Azure.Network.NetworkInterfaceArgs
        ///         {
        ///             Location = exampleResourceGroup.Location,
        ///             ResourceGroupName = exampleResourceGroup.Name,
        ///             IpConfigurations = 
        ///             {
        ///                 new Azure.Network.Inputs.NetworkInterfaceIpConfigurationArgs
        ///                 {
        ///                     Name = "testconfiguration1",
        ///                     SubnetId = exampleSubnet.Id,
        ///                     PrivateIpAddressAllocation = "Static",
        ///                     PrivateIpAddress = "10.0.2.5",
        ///                     PublicIpAddressId = examplePublicIp.Id,
        ///                 },
        ///             },
        ///         });
        ///         var exampleVirtualMachine = new Azure.Compute.VirtualMachine("exampleVirtualMachine", new Azure.Compute.VirtualMachineArgs
        ///         {
        ///             Location = exampleResourceGroup.Location,
        ///             ResourceGroupName = exampleResourceGroup.Name,
        ///             NetworkInterfaceIds = 
        ///             {
        ///                 exampleNetworkInterface.Id,
        ///             },
        ///         });
        ///         // ...
        ///         var examplePublicIP = Output.Tuple(examplePublicIp.Name, exampleVirtualMachine.ResourceGroupName).Apply(values =&gt;
        ///         {
        ///             var name = values.Item1;
        ///             var resourceGroupName = values.Item2;
        ///             return Azure.Network.GetPublicIP.InvokeAsync(new Azure.Network.GetPublicIPArgs
        ///             {
        ///                 Name = name,
        ///                 ResourceGroupName = resourceGroupName,
        ///             });
        ///         });
        ///         this.PublicIpAddress = examplePublicIp.IpAddress;
        ///     }
        /// 
        ///     [Output("publicIpAddress")]
        ///     public Output&lt;string&gt; PublicIpAddress { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPublicIPResult> InvokeAsync(GetPublicIPArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIPResult>("azure:network/getPublicIP:getPublicIP", args ?? new GetPublicIPArgs(), options.WithVersion());
    }


    public sealed class GetPublicIPArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the public IP address.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetPublicIPArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicIPResult
    {
        public readonly string AllocationMethod;
        /// <summary>
        /// The label for the Domain Name.
        /// </summary>
        public readonly string DomainNameLabel;
        /// <summary>
        /// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the timeout for the TCP idle connection.
        /// </summary>
        public readonly int IdleTimeoutInMinutes;
        /// <summary>
        /// The IP address value that was allocated.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> IpTags;
        /// <summary>
        /// The IP version being used, for example `IPv4` or `IPv6`.
        /// </summary>
        public readonly string IpVersion;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly string ReverseFqdn;
        /// <summary>
        /// The SKU of the Public IP.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetPublicIPResult(
            string allocationMethod,

            string domainNameLabel,

            string fqdn,

            string id,

            int idleTimeoutInMinutes,

            string ipAddress,

            ImmutableDictionary<string, string> ipTags,

            string ipVersion,

            string location,

            string name,

            string resourceGroupName,

            string reverseFqdn,

            string sku,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<string> zones)
        {
            AllocationMethod = allocationMethod;
            DomainNameLabel = domainNameLabel;
            Fqdn = fqdn;
            Id = id;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpAddress = ipAddress;
            IpTags = ipTags;
            IpVersion = ipVersion;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ReverseFqdn = reverseFqdn;
            Sku = sku;
            Tags = tags;
            Zones = zones;
        }
    }
}
