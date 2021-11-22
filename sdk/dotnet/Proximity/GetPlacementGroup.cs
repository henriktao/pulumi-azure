// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.Proximity
{
    public static class GetPlacementGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Proximity Placement Group.
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
        ///         var example = Output.Create(Azure.Proximity.GetPlacementGroup.InvokeAsync(new Azure.Proximity.GetPlacementGroupArgs
        ///         {
        ///             Name = "tf-appsecuritygroup",
        ///             ResourceGroupName = "my-resource-group",
        ///         }));
        ///         this.ProximityPlacementGroupId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("proximityPlacementGroupId")]
        ///     public Output&lt;string&gt; ProximityPlacementGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPlacementGroupResult> InvokeAsync(GetPlacementGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPlacementGroupResult>("azure:proximity/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Proximity Placement Group.
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
        ///         var example = Output.Create(Azure.Proximity.GetPlacementGroup.InvokeAsync(new Azure.Proximity.GetPlacementGroupArgs
        ///         {
        ///             Name = "tf-appsecuritygroup",
        ///             ResourceGroupName = "my-resource-group",
        ///         }));
        ///         this.ProximityPlacementGroupId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("proximityPlacementGroupId")]
        ///     public Output&lt;string&gt; ProximityPlacementGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPlacementGroupResult> Invoke(GetPlacementGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPlacementGroupResult>("azure:proximity/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPlacementGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Proximity Placement Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Proximity Placement Group exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPlacementGroupArgs()
        {
        }
    }

    public sealed class GetPlacementGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Proximity Placement Group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Proximity Placement Group exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPlacementGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPlacementGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetPlacementGroupResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
