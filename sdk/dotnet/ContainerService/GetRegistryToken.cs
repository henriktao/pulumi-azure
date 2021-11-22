// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.ContainerService
{
    public static class GetRegistryToken
    {
        /// <summary>
        /// Use this data source to access information about an existing Container Registry token.
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
        ///         var example = Output.Create(Azure.ContainerService.GetRegistryToken.InvokeAsync(new Azure.ContainerService.GetRegistryTokenArgs
        ///         {
        ///             Name = "exampletoken",
        ///             ResourceGroupName = "example-resource-group",
        ///             ContainerRegistryName = "example-registry",
        ///         }));
        ///         this.ScopeMapId = example.Apply(example =&gt; example.ScopeMapId);
        ///     }
        /// 
        ///     [Output("scopeMapId")]
        ///     public Output&lt;string&gt; ScopeMapId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryTokenResult> InvokeAsync(GetRegistryTokenArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryTokenResult>("azure:containerservice/getRegistryToken:getRegistryToken", args ?? new GetRegistryTokenArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Container Registry token.
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
        ///         var example = Output.Create(Azure.ContainerService.GetRegistryToken.InvokeAsync(new Azure.ContainerService.GetRegistryTokenArgs
        ///         {
        ///             Name = "exampletoken",
        ///             ResourceGroupName = "example-resource-group",
        ///             ContainerRegistryName = "example-registry",
        ///         }));
        ///         this.ScopeMapId = example.Apply(example =&gt; example.ScopeMapId);
        ///     }
        /// 
        ///     [Output("scopeMapId")]
        ///     public Output&lt;string&gt; ScopeMapId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegistryTokenResult> Invoke(GetRegistryTokenInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryTokenResult>("azure:containerservice/getRegistryToken:getRegistryToken", args ?? new GetRegistryTokenInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRegistryTokenArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the Container Registry where the token exists.
        /// </summary>
        [Input("containerRegistryName", required: true)]
        public string ContainerRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the Container Registry token.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Container Registry token exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRegistryTokenArgs()
        {
        }
    }

    public sealed class GetRegistryTokenInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the Container Registry where the token exists.
        /// </summary>
        [Input("containerRegistryName", required: true)]
        public Input<string> ContainerRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the Container Registry token.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Container Registry token exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetRegistryTokenInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryTokenResult
    {
        public readonly string ContainerRegistryName;
        /// <summary>
        /// Whether this Token is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Scope Map ID used by the token.
        /// </summary>
        public readonly string ScopeMapId;

        [OutputConstructor]
        private GetRegistryTokenResult(
            string containerRegistryName,

            bool enabled,

            string id,

            string name,

            string resourceGroupName,

            string scopeMapId)
        {
            ContainerRegistryName = containerRegistryName;
            Enabled = enabled;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            ScopeMapId = scopeMapId;
        }
    }
}
