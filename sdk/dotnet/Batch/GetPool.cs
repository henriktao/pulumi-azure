// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Batch
{
    public static class GetPool
    {
        /// <summary>
        /// Use this data source to access information about an existing Batch pool
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
        ///         var example = Output.Create(Azure.Batch.GetPool.InvokeAsync(new Azure.Batch.GetPoolArgs
        ///         {
        ///             AccountName = "testbatchaccount",
        ///             Name = "testbatchpool",
        ///             ResourceGroupName = "test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("azure:batch/getPool:getPool", args ?? new GetPoolArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Batch pool
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
        ///         var example = Output.Create(Azure.Batch.GetPool.InvokeAsync(new Azure.Batch.GetPoolArgs
        ///         {
        ///             AccountName = "testbatchaccount",
        ///             Name = "testbatchpool",
        ///             ResourceGroupName = "test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPoolResult> Invoke(GetPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPoolResult>("azure:batch/getPool:getPool", args ?? new GetPoolInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPoolArgs()
        {
        }
    }

    public sealed class GetPoolInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPoolResult
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// A `auto_scale` block that describes the scale settings when using auto scale.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolAutoScaleResult> AutoScales;
        /// <summary>
        /// One or more `certificate` blocks that describe the certificates installed on each compute node in the pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolCertificateResult> Certificates;
        /// <summary>
        /// The container configuration used in the pool's VMs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolContainerConfigurationResult> ContainerConfigurations;
        public readonly string DisplayName;
        /// <summary>
        /// A `fixed_scale` block that describes the scale settings when using fixed scale.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolFixedScaleResult> FixedScales;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The maximum number of tasks that can run concurrently on a single compute node in the pool.
        /// </summary>
        public readonly int MaxTasksPerNode;
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.GetPoolNetworkConfigurationResult NetworkConfiguration;
        /// <summary>
        /// The Sku of the node agents in the Batch pool.
        /// </summary>
        public readonly string NodeAgentSkuId;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A `start_task` block that describes the start task settings for the Batch pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolStartTaskResult> StartTasks;
        /// <summary>
        /// The reference of the storage image used by the nodes in the Batch pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoolStorageImageReferenceResult> StorageImageReferences;
        /// <summary>
        /// The size of the VM created in the Batch pool.
        /// </summary>
        public readonly string VmSize;

        [OutputConstructor]
        private GetPoolResult(
            string accountName,

            ImmutableArray<Outputs.GetPoolAutoScaleResult> autoScales,

            ImmutableArray<Outputs.GetPoolCertificateResult> certificates,

            ImmutableArray<Outputs.GetPoolContainerConfigurationResult> containerConfigurations,

            string displayName,

            ImmutableArray<Outputs.GetPoolFixedScaleResult> fixedScales,

            string id,

            int maxTasksPerNode,

            ImmutableDictionary<string, string> metadata,

            string name,

            Outputs.GetPoolNetworkConfigurationResult networkConfiguration,

            string nodeAgentSkuId,

            string resourceGroupName,

            ImmutableArray<Outputs.GetPoolStartTaskResult> startTasks,

            ImmutableArray<Outputs.GetPoolStorageImageReferenceResult> storageImageReferences,

            string vmSize)
        {
            AccountName = accountName;
            AutoScales = autoScales;
            Certificates = certificates;
            ContainerConfigurations = containerConfigurations;
            DisplayName = displayName;
            FixedScales = fixedScales;
            Id = id;
            MaxTasksPerNode = maxTasksPerNode;
            Metadata = metadata;
            Name = name;
            NetworkConfiguration = networkConfiguration;
            NodeAgentSkuId = nodeAgentSkuId;
            ResourceGroupName = resourceGroupName;
            StartTasks = startTasks;
            StorageImageReferences = storageImageReferences;
            VmSize = vmSize;
        }
    }
}
