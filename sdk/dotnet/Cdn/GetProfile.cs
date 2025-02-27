// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cdn
{
    public static class GetProfile
    {
        /// <summary>
        /// Use this data source to access information about an existing CDN Profile.
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
        ///         var example = Output.Create(Azure.Cdn.GetProfile.InvokeAsync(new Azure.Cdn.GetProfileArgs
        ///         {
        ///             Name = "myfirstcdnprofile",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.CdnProfileId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("cdnProfileId")]
        ///     public Output&lt;string&gt; CdnProfileId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProfileResult> InvokeAsync(GetProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProfileResult>("azure:cdn/getProfile:getProfile", args ?? new GetProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing CDN Profile.
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
        ///         var example = Output.Create(Azure.Cdn.GetProfile.InvokeAsync(new Azure.Cdn.GetProfileArgs
        ///         {
        ///             Name = "myfirstcdnprofile",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.CdnProfileId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("cdnProfileId")]
        ///     public Output&lt;string&gt; CdnProfileId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProfileResult> Invoke(GetProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProfileResult>("azure:cdn/getProfile:getProfile", args ?? new GetProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CDN Profile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the CDN Profile exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetProfileArgs()
        {
        }
    }

    public sealed class GetProfileInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the CDN Profile.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the CDN Profile exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetProfileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProfileResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The pricing related information of current CDN profile.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetProfileResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
        }
    }
}
