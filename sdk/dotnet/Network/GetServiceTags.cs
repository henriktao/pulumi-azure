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
    public static class GetServiceTags
    {
        /// <summary>
        /// Use this data source to access information about Service Tags.
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
        ///         var example = Output.Create(Azure.Network.GetServiceTags.InvokeAsync(new Azure.Network.GetServiceTagsArgs
        ///         {
        ///             Location = "westcentralus",
        ///             Service = "AzureKeyVault",
        ///             LocationFilter = "northeurope",
        ///         }));
        ///         this.AddressPrefixes = example.Apply(example =&gt; example.AddressPrefixes);
        ///         this.Ipv4Cidrs = example.Apply(example =&gt; example.Ipv4Cidrs);
        ///     }
        /// 
        ///     [Output("addressPrefixes")]
        ///     public Output&lt;string&gt; AddressPrefixes { get; set; }
        ///     [Output("ipv4Cidrs")]
        ///     public Output&lt;string&gt; Ipv4Cidrs { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceTagsResult> InvokeAsync(GetServiceTagsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceTagsResult>("azure:network/getServiceTags:getServiceTags", args ?? new GetServiceTagsArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about Service Tags.
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
        ///         var example = Output.Create(Azure.Network.GetServiceTags.InvokeAsync(new Azure.Network.GetServiceTagsArgs
        ///         {
        ///             Location = "westcentralus",
        ///             Service = "AzureKeyVault",
        ///             LocationFilter = "northeurope",
        ///         }));
        ///         this.AddressPrefixes = example.Apply(example =&gt; example.AddressPrefixes);
        ///         this.Ipv4Cidrs = example.Apply(example =&gt; example.Ipv4Cidrs);
        ///     }
        /// 
        ///     [Output("addressPrefixes")]
        ///     public Output&lt;string&gt; AddressPrefixes { get; set; }
        ///     [Output("ipv4Cidrs")]
        ///     public Output&lt;string&gt; Ipv4Cidrs { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceTagsResult> Invoke(GetServiceTagsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceTagsResult>("azure:network/getServiceTags:getServiceTags", args ?? new GetServiceTagsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetServiceTagsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Region where the Service Tags exists. This value is not used to filter the results but for specifying the region to request. For filtering by region use `location_filter` instead.  More information can be found here: [Service Tags URL parameters](https://docs.microsoft.com/en-us/rest/api/virtualnetwork/servicetags/list#uri-parameters).
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// Changes the scope of the service tags. Can be any value that is also valid for `location`. If this field is empty then all address prefixes are considered instead of only location specific ones.
        /// </summary>
        [Input("locationFilter")]
        public string? LocationFilter { get; set; }

        /// <summary>
        /// The type of the service for which address prefixes will be fetched. Available service tags can be found here: [Available service tags](https://docs.microsoft.com/en-us/azure/virtual-network/service-tags-overview#available-service-tags).
        /// </summary>
        [Input("service", required: true)]
        public string Service { get; set; } = null!;

        public GetServiceTagsArgs()
        {
        }
    }

    public sealed class GetServiceTagsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Region where the Service Tags exists. This value is not used to filter the results but for specifying the region to request. For filtering by region use `location_filter` instead.  More information can be found here: [Service Tags URL parameters](https://docs.microsoft.com/en-us/rest/api/virtualnetwork/servicetags/list#uri-parameters).
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Changes the scope of the service tags. Can be any value that is also valid for `location`. If this field is empty then all address prefixes are considered instead of only location specific ones.
        /// </summary>
        [Input("locationFilter")]
        public Input<string>? LocationFilter { get; set; }

        /// <summary>
        /// The type of the service for which address prefixes will be fetched. Available service tags can be found here: [Available service tags](https://docs.microsoft.com/en-us/azure/virtual-network/service-tags-overview#available-service-tags).
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public GetServiceTagsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceTagsResult
    {
        /// <summary>
        /// List of address prefixes for the service type (and optionally a specific region).
        /// </summary>
        public readonly ImmutableArray<string> AddressPrefixes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of IPv4 addresses for the service type (and optionally a specific region)
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Cidrs;
        /// <summary>
        /// List of IPv6 addresses for the service type (and optionally a specific region)
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Cidrs;
        public readonly string Location;
        public readonly string? LocationFilter;
        public readonly string Service;

        [OutputConstructor]
        private GetServiceTagsResult(
            ImmutableArray<string> addressPrefixes,

            string id,

            ImmutableArray<string> ipv4Cidrs,

            ImmutableArray<string> ipv6Cidrs,

            string location,

            string? locationFilter,

            string service)
        {
            AddressPrefixes = addressPrefixes;
            Id = id;
            Ipv4Cidrs = ipv4Cidrs;
            Ipv6Cidrs = ipv6Cidrs;
            Location = location;
            LocationFilter = locationFilter;
            Service = service;
        }
    }
}
