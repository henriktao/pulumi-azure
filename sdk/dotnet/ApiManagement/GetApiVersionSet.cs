// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    public static class GetApiVersionSet
    {
        /// <summary>
        /// Uses this data source to access information about an API Version Set within an API Management Service.
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
        ///         var example = Output.Create(Azure.ApiManagement.GetApiVersionSet.InvokeAsync(new Azure.ApiManagement.GetApiVersionSetArgs
        ///         {
        ///             ResourceGroupName = "example-resources",
        ///             ApiManagementName = "example-api",
        ///             Name = "example-api-version-set",
        ///         }));
        ///         this.ApiManagementApiVersionSetId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("apiManagementApiVersionSetId")]
        ///     public Output&lt;string&gt; ApiManagementApiVersionSetId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApiVersionSetResult> InvokeAsync(GetApiVersionSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiVersionSetResult>("azure:apimanagement/getApiVersionSet:getApiVersionSet", args ?? new GetApiVersionSetArgs(), options.WithDefaults());

        /// <summary>
        /// Uses this data source to access information about an API Version Set within an API Management Service.
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
        ///         var example = Output.Create(Azure.ApiManagement.GetApiVersionSet.InvokeAsync(new Azure.ApiManagement.GetApiVersionSetArgs
        ///         {
        ///             ResourceGroupName = "example-resources",
        ///             ApiManagementName = "example-api",
        ///             Name = "example-api-version-set",
        ///         }));
        ///         this.ApiManagementApiVersionSetId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("apiManagementApiVersionSetId")]
        ///     public Output&lt;string&gt; ApiManagementApiVersionSetId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApiVersionSetResult> Invoke(GetApiVersionSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApiVersionSetResult>("azure:apimanagement/getApiVersionSet:getApiVersionSet", args ?? new GetApiVersionSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiVersionSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management Service where the API Version Set exists.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public string ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The name of the API Version Set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the parent API Management Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApiVersionSetArgs()
        {
        }
    }

    public sealed class GetApiVersionSetInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management Service where the API Version Set exists.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// The name of the API Version Set.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the parent API Management Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApiVersionSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiVersionSetResult
    {
        public readonly string ApiManagementName;
        /// <summary>
        /// The description of API Version Set.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of this API Version Set.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The name of the Header which should be read from Inbound Requests which defines the API Version.
        /// </summary>
        public readonly string VersionHeaderName;
        /// <summary>
        /// The name of the Query String which should be read from Inbound Requests which defines the API Version.
        /// </summary>
        public readonly string VersionQueryName;
        public readonly string VersioningScheme;

        [OutputConstructor]
        private GetApiVersionSetResult(
            string apiManagementName,

            string description,

            string displayName,

            string id,

            string name,

            string resourceGroupName,

            string versionHeaderName,

            string versionQueryName,

            string versioningScheme)
        {
            ApiManagementName = apiManagementName;
            Description = description;
            DisplayName = displayName;
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            VersionHeaderName = versionHeaderName;
            VersionQueryName = versionQueryName;
            VersioningScheme = versioningScheme;
        }
    }
}
