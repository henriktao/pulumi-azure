// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to access information about an existing HDInsight Cluster.
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
        ///         var example = Output.Create(Azure.HDInsight.GetCluster.InvokeAsync(new Azure.HDInsight.GetClusterArgs
        ///         {
        ///             Name = "example",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.HttpsEndpoint = example.Apply(example =&gt; example.HttpsEndpoint);
        ///     }
        /// 
        ///     [Output("httpsEndpoint")]
        ///     public Output&lt;string&gt; HttpsEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azure:hdinsight/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing HDInsight Cluster.
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
        ///         var example = Output.Create(Azure.HDInsight.GetCluster.InvokeAsync(new Azure.HDInsight.GetClusterArgs
        ///         {
        ///             Name = "example",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.HttpsEndpoint = example.Apply(example =&gt; example.HttpsEndpoint);
        ///     }
        /// 
        ///     [Output("httpsEndpoint")]
        ///     public Output&lt;string&gt; HttpsEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("azure:hdinsight/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of this HDInsight Cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
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
        /// Specifies the name of this HDInsight Cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight Cluster exists.
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
        /// The version of HDInsights which is used on this HDInsight Cluster.
        /// </summary>
        public readonly string ClusterVersion;
        /// <summary>
        /// A map of versions of software used on this HDInsights Cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ComponentVersions;
        /// <summary>
        /// The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.
        /// </summary>
        public readonly string EdgeSshEndpoint;
        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterGatewayResult> Gateways;
        /// <summary>
        /// The HTTPS Endpoint for this HDInsight Cluster.
        /// </summary>
        public readonly string HttpsEndpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Kafka Rest Proxy Endpoint for this HDInsight Cluster.
        /// </summary>
        public readonly string KafkaRestProxyEndpoint;
        /// <summary>
        /// The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The Azure Region in which this HDInsight Cluster exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SSH Endpoint for this HDInsight Cluster.
        /// </summary>
        public readonly string SshEndpoint;
        /// <summary>
        /// A map of tags assigned to the HDInsight Cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The SKU / Tier of this HDInsight Cluster.
        /// </summary>
        public readonly string Tier;
        /// <summary>
        /// The minimal supported tls version.
        /// </summary>
        public readonly string TlsMinVersion;

        [OutputConstructor]
        private GetClusterResult(
            string clusterVersion,

            ImmutableDictionary<string, string> componentVersions,

            string edgeSshEndpoint,

            ImmutableArray<Outputs.GetClusterGatewayResult> gateways,

            string httpsEndpoint,

            string id,

            string kafkaRestProxyEndpoint,

            string kind,

            string location,

            string name,

            string resourceGroupName,

            string sshEndpoint,

            ImmutableDictionary<string, string> tags,

            string tier,

            string tlsMinVersion)
        {
            ClusterVersion = clusterVersion;
            ComponentVersions = componentVersions;
            EdgeSshEndpoint = edgeSshEndpoint;
            Gateways = gateways;
            HttpsEndpoint = httpsEndpoint;
            Id = id;
            KafkaRestProxyEndpoint = kafkaRestProxyEndpoint;
            Kind = kind;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SshEndpoint = sshEndpoint;
            Tags = tags;
            Tier = tier;
            TlsMinVersion = tlsMinVersion;
        }
    }
}
