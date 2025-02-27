// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static class GetSharedImage
    {
        /// <summary>
        /// Use this data source to access information about an existing Shared Image within a Shared Image Gallery.
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
        ///         var example = Output.Create(Azure.Compute.GetSharedImage.InvokeAsync(new Azure.Compute.GetSharedImageArgs
        ///         {
        ///             GalleryName = "my-image-gallery",
        ///             Name = "my-image",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSharedImageResult> InvokeAsync(GetSharedImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSharedImageResult>("azure:compute/getSharedImage:getSharedImage", args ?? new GetSharedImageArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Shared Image within a Shared Image Gallery.
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
        ///         var example = Output.Create(Azure.Compute.GetSharedImage.InvokeAsync(new Azure.Compute.GetSharedImageArgs
        ///         {
        ///             GalleryName = "my-image-gallery",
        ///             Name = "my-image",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSharedImageResult> Invoke(GetSharedImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSharedImageResult>("azure:compute/getSharedImage:getSharedImage", args ?? new GetSharedImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSharedImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery in which the Shared Image exists.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSharedImageArgs()
        {
        }
    }

    public sealed class GetSharedImageInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery in which the Shared Image exists.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSharedImageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSharedImageResult
    {
        /// <summary>
        /// The description of this Shared Image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The End User Licence Agreement for the Shared Image.
        /// </summary>
        public readonly string Eula;
        public readonly string GalleryName;
        /// <summary>
        /// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on.
        /// </summary>
        public readonly string HyperVGeneration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identifier` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSharedImageIdentifierResult> Identifiers;
        /// <summary>
        /// The supported Azure location where the Shared Image Gallery exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The type of Operating System present in this Shared Image.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The URI containing the Privacy Statement for this Shared Image.
        /// </summary>
        public readonly string PrivacyStatementUri;
        /// <summary>
        /// The URI containing the Release Notes for this Shared Image.
        /// </summary>
        public readonly string ReleaseNoteUri;
        public readonly string ResourceGroupName;
        /// <summary>
        /// Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run).
        /// </summary>
        public readonly bool Specialized;
        /// <summary>
        /// A mapping of tags assigned to the Shared Image.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetSharedImageResult(
            string description,

            string eula,

            string galleryName,

            string hyperVGeneration,

            string id,

            ImmutableArray<Outputs.GetSharedImageIdentifierResult> identifiers,

            string location,

            string name,

            string osType,

            string privacyStatementUri,

            string releaseNoteUri,

            string resourceGroupName,

            bool specialized,

            ImmutableDictionary<string, string> tags)
        {
            Description = description;
            Eula = eula;
            GalleryName = galleryName;
            HyperVGeneration = hyperVGeneration;
            Id = id;
            Identifiers = identifiers;
            Location = location;
            Name = name;
            OsType = osType;
            PrivacyStatementUri = privacyStatementUri;
            ReleaseNoteUri = releaseNoteUri;
            ResourceGroupName = resourceGroupName;
            Specialized = specialized;
            Tags = tags;
        }
    }
}
