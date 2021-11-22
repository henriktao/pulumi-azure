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
    public static class GetApplication
    {
        /// <summary>
        /// Use this data source to access information about an existing Batch Application instance.
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
        ///         var example = Output.Create(Azure.Batch.GetApplication.InvokeAsync(new Azure.Batch.GetApplicationArgs
        ///         {
        ///             Name = "testapplication",
        ///             ResourceGroupName = "test",
        ///             AccountName = "testbatchaccount",
        ///         }));
        ///         this.BatchApplicationId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("batchApplicationId")]
        ///     public Output&lt;string&gt; BatchApplicationId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azure:batch/getApplication:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Batch Application instance.
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
        ///         var example = Output.Create(Azure.Batch.GetApplication.InvokeAsync(new Azure.Batch.GetApplicationArgs
        ///         {
        ///             Name = "testapplication",
        ///             ResourceGroupName = "test",
        ///             AccountName = "testbatchaccount",
        ///         }));
        ///         this.BatchApplicationId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("batchApplicationId")]
        ///     public Output&lt;string&gt; BatchApplicationId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("azure:batch/getApplication:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
    }

    public sealed class GetApplicationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        public readonly string AccountName;
        /// <summary>
        /// May packages within the application be overwritten using the same version string.
        /// </summary>
        public readonly bool AllowUpdates;
        /// <summary>
        /// The package to use if a client requests the application but does not specify a version.
        /// </summary>
        public readonly string DefaultVersion;
        /// <summary>
        /// The display name for the application.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Batch application name.
        /// </summary>
        public readonly string Name;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetApplicationResult(
            string accountName,

            bool allowUpdates,

            string defaultVersion,

            string displayName,

            string id,

            string name,

            string resourceGroupName)
        {
            AccountName = accountName;
            AllowUpdates = allowUpdates;
            DefaultVersion = defaultVersion;
            DisplayName = displayName;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
        }
    }
}
