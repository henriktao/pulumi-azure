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
    public static class GetImage
    {
        /// <summary>
        /// Use this data source to access information about an existing Image.
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
        ///         var search = Output.Create(Azure.Compute.GetImage.InvokeAsync(new Azure.Compute.GetImageArgs
        ///         {
        ///             Name = "search-api",
        ///             ResourceGroupName = "packerimages",
        ///         }));
        ///         this.ImageId = search.Apply(search =&gt; search.Id);
        ///     }
        /// 
        ///     [Output("imageId")]
        ///     public Output&lt;string&gt; ImageId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("azure:compute/getImage:getImage", args ?? new GetImageArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Image.
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
        ///         var search = Output.Create(Azure.Compute.GetImage.InvokeAsync(new Azure.Compute.GetImageArgs
        ///         {
        ///             Name = "search-api",
        ///             ResourceGroupName = "packerimages",
        ///         }));
        ///         this.ImageId = search.Apply(search =&gt; search.Id);
        ///     }
        /// 
        ///     [Output("imageId")]
        ///     public Output&lt;string&gt; ImageId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetImageResult>("azure:compute/getImage:getImage", args ?? new GetImageInvokeArgs(), options.WithVersion());
    }


    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Image.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Regex pattern of the image to match.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The Name of the Resource Group where this Image exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
        /// </summary>
        [Input("sortDescending")]
        public bool? SortDescending { get; set; }

        public GetImageArgs()
        {
        }
    }

    public sealed class GetImageInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Regex pattern of the image to match.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The Name of the Resource Group where this Image exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
        /// </summary>
        [Input("sortDescending")]
        public Input<bool>? SortDescending { get; set; }

        public GetImageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// a collection of `data_disk` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImageDataDiskResult> DataDisks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the Azure Location where this Image exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// the name of the Image.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        /// <summary>
        /// a `os_disk` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImageOsDiskResult> OsDisks;
        public readonly string ResourceGroupName;
        public readonly bool? SortDescending;
        /// <summary>
        /// a mapping of tags to assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// is zone resiliency enabled?
        /// </summary>
        public readonly bool ZoneResilient;

        [OutputConstructor]
        private GetImageResult(
            ImmutableArray<Outputs.GetImageDataDiskResult> dataDisks,

            string id,

            string location,

            string? name,

            string? nameRegex,

            ImmutableArray<Outputs.GetImageOsDiskResult> osDisks,

            string resourceGroupName,

            bool? sortDescending,

            ImmutableDictionary<string, string> tags,

            bool zoneResilient)
        {
            DataDisks = dataDisks;
            Id = id;
            Location = location;
            Name = name;
            NameRegex = nameRegex;
            OsDisks = osDisks;
            ResourceGroupName = resourceGroupName;
            SortDescending = sortDescending;
            Tags = tags;
            ZoneResilient = zoneResilient;
        }
    }
}
