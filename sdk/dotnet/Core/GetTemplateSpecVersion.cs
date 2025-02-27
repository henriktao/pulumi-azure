// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    public static class GetTemplateSpecVersion
    {
        /// <summary>
        /// Use this data source to access information about an existing Template Spec Version.
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
        ///         var example = Output.Create(Azure.Core.GetTemplateSpecVersion.InvokeAsync(new Azure.Core.GetTemplateSpecVersionArgs
        ///         {
        ///             Name = "exampleTemplateSpec",
        ///             ResourceGroupName = "MyResourceGroup",
        ///             Version = "v1.0.4",
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
        public static Task<GetTemplateSpecVersionResult> InvokeAsync(GetTemplateSpecVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTemplateSpecVersionResult>("azure:core/getTemplateSpecVersion:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Template Spec Version.
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
        ///         var example = Output.Create(Azure.Core.GetTemplateSpecVersion.InvokeAsync(new Azure.Core.GetTemplateSpecVersionArgs
        ///         {
        ///             Name = "exampleTemplateSpec",
        ///             ResourceGroupName = "MyResourceGroup",
        ///             Version = "v1.0.4",
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
        public static Output<GetTemplateSpecVersionResult> Invoke(GetTemplateSpecVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTemplateSpecVersionResult>("azure:core/getTemplateSpecVersion:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateSpecVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Template Spec.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Template Spec exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Version Name of the Template Spec.
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetTemplateSpecVersionArgs()
        {
        }
    }

    public sealed class GetTemplateSpecVersionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Template Spec.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Template Spec exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Version Name of the Template Spec.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetTemplateSpecVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTemplateSpecVersionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the Template.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The ARM Template body of the Template Spec Version.
        /// </summary>
        public readonly string TemplateBody;
        public readonly string Version;

        [OutputConstructor]
        private GetTemplateSpecVersionResult(
            string id,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            string templateBody,

            string version)
        {
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            TemplateBody = templateBody;
            Version = version;
        }
    }
}
