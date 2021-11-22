// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Compute
{
    public static class GetImages
    {
        /// <summary>
        /// Use this data source to access information about existing Images within a Resource Group.
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
        ///         var example = Output.Create(Azure.Compute.GetImages.InvokeAsync(new Azure.Compute.GetImagesArgs
        ///         {
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("azure:compute/getImages:getImages", args ?? new GetImagesArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about existing Images within a Resource Group.
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
        ///         var example = Output.Create(Azure.Compute.GetImages.InvokeAsync(new Azure.Compute.GetImagesArgs
        ///         {
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImagesResult> Invoke(GetImagesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetImagesResult>("azure:compute/getImages:getImages", args ?? new GetImagesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource Group in which the Image exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tagsFilter")]
        private Dictionary<string, string>? _tagsFilter;

        /// <summary>
        /// A mapping of tags to filter the list of images against.
        /// </summary>
        public Dictionary<string, string> TagsFilter
        {
            get => _tagsFilter ?? (_tagsFilter = new Dictionary<string, string>());
            set => _tagsFilter = value;
        }

        public GetImagesArgs()
        {
        }
    }

    public sealed class GetImagesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource Group in which the Image exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tagsFilter")]
        private InputMap<string>? _tagsFilter;

        /// <summary>
        /// A mapping of tags to filter the list of images against.
        /// </summary>
        public InputMap<string> TagsFilter
        {
            get => _tagsFilter ?? (_tagsFilter = new InputMap<string>());
            set => _tagsFilter = value;
        }

        public GetImagesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more `images` blocks as defined below:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImageResult> Images;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string>? TagsFilter;

        [OutputConstructor]
        private GetImagesResult(
            string id,

            ImmutableArray<Outputs.GetImagesImageResult> images,

            string resourceGroupName,

            ImmutableDictionary<string, string>? tagsFilter)
        {
            Id = id;
            Images = images;
            ResourceGroupName = resourceGroupName;
            TagsFilter = tagsFilter;
        }
    }
}
