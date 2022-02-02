// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Kusto
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to access information about an existing Kusto (also known as Azure Data Explorer) Cluster
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
        ///         var example = Output.Create(Azure.Kusto.GetCluster.InvokeAsync(new Azure.Kusto.GetClusterArgs
        ///         {
        ///             Name = "kustocluster",
        ///             ResourceGroupName = "test_resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure:kusto/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Kusto (also known as Azure Data Explorer) Cluster
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
        ///         var example = Output.Create(Azure.Kusto.GetCluster.InvokeAsync(new Azure.Kusto.GetClusterArgs
        ///         {
        ///             Name = "kustocluster",
        ///             ResourceGroupName = "test_resource_group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure:kusto/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Kusto Cluster exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Kusto Cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Kusto Cluster exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The Kusto Cluster URI to be used for data ingestion.
        /// </summary>
        public readonly string DataIngestionUri;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The FQDN of the Azure Kusto Cluster.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetClusterResult(
            string dataIngestionUri,

            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            string uri)
        {
            DataIngestionUri = dataIngestionUri;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            Uri = uri;
        }
    }
}
