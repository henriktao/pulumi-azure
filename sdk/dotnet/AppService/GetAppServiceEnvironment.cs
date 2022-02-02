// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetAppServiceEnvironment
    {
        /// <summary>
        /// Use this data source to access information about an existing App Service Environment.
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
        ///         var example = Output.Create(Azure.AppService.GetAppServiceEnvironment.InvokeAsync(new Azure.AppService.GetAppServiceEnvironmentArgs
        ///         {
        ///             Name = "existing-ase",
        ///             ResourceGroupName = "existing-rg",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppServiceEnvironmentResult> InvokeAsync(GetAppServiceEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppServiceEnvironmentResult>("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", args ?? new GetAppServiceEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing App Service Environment.
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
        ///         var example = Output.Create(Azure.AppService.GetAppServiceEnvironment.InvokeAsync(new Azure.AppService.GetAppServiceEnvironmentArgs
        ///         {
        ///             Name = "existing-ase",
        ///             ResourceGroupName = "existing-rg",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppServiceEnvironmentResult> Invoke(GetAppServiceEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppServiceEnvironmentResult>("azure:appservice/getAppServiceEnvironment:getAppServiceEnvironment", args ?? new GetAppServiceEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppServiceEnvironmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this App Service Environment.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the App Service Environment exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppServiceEnvironmentArgs()
        {
        }
    }

    public sealed class GetAppServiceEnvironmentInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this App Service Environment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the App Service Environment exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAppServiceEnvironmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppServiceEnvironmentResult
    {
        /// <summary>
        /// Zero or more `cluster_setting` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppServiceEnvironmentClusterSettingResult> ClusterSettings;
        /// <summary>
        /// The number of app instances per App Service Environment Front End.
        /// </summary>
        public readonly int FrontEndScaleFactor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address of internal load balancer of the App Service Environment.
        /// </summary>
        public readonly string InternalIpAddress;
        /// <summary>
        /// The Azure Region where the App Service Environment exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Cluster Setting.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of outbound IP addresses of the App Service Environment.
        /// </summary>
        public readonly ImmutableArray<string> OutboundIpAddresses;
        /// <summary>
        /// The Pricing Tier (Isolated SKU) of the App Service Environment.
        /// </summary>
        public readonly string PricingTier;
        public readonly string ResourceGroupName;
        /// <summary>
        /// IP address of service endpoint of the App Service Environment.
        /// </summary>
        public readonly string ServiceIpAddress;
        /// <summary>
        /// A mapping of tags assigned to the App Service Environment.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAppServiceEnvironmentResult(
            ImmutableArray<Outputs.GetAppServiceEnvironmentClusterSettingResult> clusterSettings,

            int frontEndScaleFactor,

            string id,

            string internalIpAddress,

            string location,

            string name,

            ImmutableArray<string> outboundIpAddresses,

            string pricingTier,

            string resourceGroupName,

            string serviceIpAddress,

            ImmutableDictionary<string, string> tags)
        {
            ClusterSettings = clusterSettings;
            FrontEndScaleFactor = frontEndScaleFactor;
            Id = id;
            InternalIpAddress = internalIpAddress;
            Location = location;
            Name = name;
            OutboundIpAddresses = outboundIpAddresses;
            PricingTier = pricingTier;
            ResourceGroupName = resourceGroupName;
            ServiceIpAddress = serviceIpAddress;
            Tags = tags;
        }
    }
}
