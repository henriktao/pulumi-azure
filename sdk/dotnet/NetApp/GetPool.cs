// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    public static class GetPool
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Pool.
        /// 
        /// 
        /// ## NetApp Pool Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.NetApp.GetPool.InvokeAsync(new Azure.NetApp.GetPoolArgs
        ///         {
        ///             ResourceGroupName = "acctestRG",
        ///             AccountName = "acctestnetappaccount",
        ///             Name = "acctestnetapppool",
        ///         }));
        ///         this.NetappPoolId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("netappPoolId")]
        ///     public Output&lt;string&gt; NetappPoolId { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("azure:netapp/getPool:getPool", args ?? new GetPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Uses this data source to access information about an existing NetApp Pool.
        /// 
        /// 
        /// ## NetApp Pool Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.NetApp.GetPool.InvokeAsync(new Azure.NetApp.GetPoolArgs
        ///         {
        ///             ResourceGroupName = "acctestRG",
        ///             AccountName = "acctestnetappaccount",
        ///             Name = "acctestnetapppool",
        ///         }));
        ///         this.NetappPoolId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("netappPoolId")]
        ///     public Output&lt;string&gt; NetappPoolId { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Output<GetPoolResult> Invoke(GetPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPoolResult>("azure:netapp/getPool:getPool", args ?? new GetPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account where the NetApp pool exists.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Pool exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPoolArgs()
        {
        }
    }

    public sealed class GetPoolInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account where the NetApp pool exists.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Pool.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Pool exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPoolResult
    {
        public readonly string AccountName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the NetApp Pool exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The service level of the file system.
        /// </summary>
        public readonly string ServiceLevel;
        /// <summary>
        /// Provisioned size of the pool in TB.
        /// </summary>
        public readonly int SizeInTb;

        [OutputConstructor]
        private GetPoolResult(
            string accountName,

            string id,

            string location,

            string name,

            string resourceGroupName,

            string serviceLevel,

            int sizeInTb)
        {
            AccountName = accountName;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ServiceLevel = serviceLevel;
            SizeInTb = sizeInTb;
        }
    }
}
