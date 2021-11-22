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
    public static class GetAvailabilitySet
    {
        /// <summary>
        /// Use this data source to access information about an existing Availability Set.
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
        ///         var example = Output.Create(Azure.Compute.GetAvailabilitySet.InvokeAsync(new Azure.Compute.GetAvailabilitySetArgs
        ///         {
        ///             Name = "tf-appsecuritygroup",
        ///             ResourceGroupName = "my-resource-group",
        ///         }));
        ///         this.AvailabilitySetId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("availabilitySetId")]
        ///     public Output&lt;string&gt; AvailabilitySetId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAvailabilitySetResult> InvokeAsync(GetAvailabilitySetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilitySetResult>("azure:compute/getAvailabilitySet:getAvailabilitySet", args ?? new GetAvailabilitySetArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Availability Set.
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
        ///         var example = Output.Create(Azure.Compute.GetAvailabilitySet.InvokeAsync(new Azure.Compute.GetAvailabilitySetArgs
        ///         {
        ///             Name = "tf-appsecuritygroup",
        ///             ResourceGroupName = "my-resource-group",
        ///         }));
        ///         this.AvailabilitySetId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("availabilitySetId")]
        ///     public Output&lt;string&gt; AvailabilitySetId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAvailabilitySetResult> Invoke(GetAvailabilitySetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAvailabilitySetResult>("azure:compute/getAvailabilitySet:getAvailabilitySet", args ?? new GetAvailabilitySetInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAvailabilitySetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Availability Set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Availability Set exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetArgs()
        {
        }
    }

    public sealed class GetAvailabilitySetInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Availability Set.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Availability Set exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAvailabilitySetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAvailabilitySetResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The supported Azure location where the Availability Set exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Whether the availability set is managed or not.
        /// </summary>
        public readonly bool Managed;
        public readonly string Name;
        /// <summary>
        /// The number of fault domains that are used.
        /// </summary>
        public readonly string PlatformFaultDomainCount;
        /// <summary>
        /// The number of update domains that are used.
        /// </summary>
        public readonly string PlatformUpdateDomainCount;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetAvailabilitySetResult(
            string id,

            string location,

            bool managed,

            string name,

            string platformFaultDomainCount,

            string platformUpdateDomainCount,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Managed = managed;
            Name = name;
            PlatformFaultDomainCount = platformFaultDomainCount;
            PlatformUpdateDomainCount = platformUpdateDomainCount;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
