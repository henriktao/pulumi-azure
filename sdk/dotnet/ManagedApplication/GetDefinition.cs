// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.ManagedApplication
{
    public static class GetDefinition
    {
        /// <summary>
        /// Uses this data source to access information about an existing Managed Application Definition.
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
        ///         var example = Output.Create(Azure.ManagedApplication.GetDefinition.InvokeAsync(new Azure.ManagedApplication.GetDefinitionArgs
        ///         {
        ///             Name = "example-managedappdef",
        ///             ResourceGroupName = "example-resources",
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
        public static Task<GetDefinitionResult> InvokeAsync(GetDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefinitionResult>("azure:managedapplication/getDefinition:getDefinition", args ?? new GetDefinitionArgs(), options.WithVersion());

        /// <summary>
        /// Uses this data source to access information about an existing Managed Application Definition.
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
        ///         var example = Output.Create(Azure.ManagedApplication.GetDefinition.InvokeAsync(new Azure.ManagedApplication.GetDefinitionArgs
        ///         {
        ///             Name = "example-managedappdef",
        ///             ResourceGroupName = "example-resources",
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
        public static Output<GetDefinitionResult> Invoke(GetDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDefinitionResult>("azure:managedapplication/getDefinition:getDefinition", args ?? new GetDefinitionInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Managed Application Definition.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Managed Application Definition exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDefinitionArgs()
        {
        }
    }

    public sealed class GetDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Managed Application Definition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Managed Application Definition exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDefinitionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetDefinitionResult(
            string id,

            string location,

            string name,

            string resourceGroupName)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
        }
    }
}
