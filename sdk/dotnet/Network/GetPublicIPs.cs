// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetPublicIPs
    {
        /// <summary>
        /// Use this data source to access information about a set of existing Public IP Addresses.
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
        ///         var example = Output.Create(Azure.Network.GetPublicIPs.InvokeAsync(new Azure.Network.GetPublicIPsArgs
        ///         {
        ///             AttachmentStatus = "Attached",
        ///             ResourceGroupName = "pip-test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPublicIPsResult> InvokeAsync(GetPublicIPsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIPsResult>("azure:network/getPublicIPs:getPublicIPs", args ?? new GetPublicIPsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about a set of existing Public IP Addresses.
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
        ///         var example = Output.Create(Azure.Network.GetPublicIPs.InvokeAsync(new Azure.Network.GetPublicIPsArgs
        ///         {
        ///             AttachmentStatus = "Attached",
        ///             ResourceGroupName = "pip-test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPublicIPsResult> Invoke(GetPublicIPsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPublicIPsResult>("azure:network/getPublicIPs:getPublicIPs", args ?? new GetPublicIPsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicIPsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Allocation Type for the Public IP Address. Possible values include `Static` or `Dynamic`.
        /// </summary>
        [Input("allocationType")]
        public string? AllocationType { get; set; }

        [Input("attached")]
        public bool? Attached { get; set; }

        /// <summary>
        /// Filter to include IP Addresses which are attached to a device, such as a VM/LB (`Attached`) or unattached (`Unattached`). To allow for both, use `All`.
        /// </summary>
        [Input("attachmentStatus")]
        public string? AttachmentStatus { get; set; }

        /// <summary>
        /// A prefix match used for the IP Addresses `name` field, case sensitive.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        /// <summary>
        /// Specifies the name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPublicIPsArgs()
        {
        }
    }

    public sealed class GetPublicIPsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Allocation Type for the Public IP Address. Possible values include `Static` or `Dynamic`.
        /// </summary>
        [Input("allocationType")]
        public Input<string>? AllocationType { get; set; }

        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        /// <summary>
        /// Filter to include IP Addresses which are attached to a device, such as a VM/LB (`Attached`) or unattached (`Unattached`). To allow for both, use `All`.
        /// </summary>
        [Input("attachmentStatus")]
        public Input<string>? AttachmentStatus { get; set; }

        /// <summary>
        /// A prefix match used for the IP Addresses `name` field, case sensitive.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Specifies the name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPublicIPsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicIPsResult
    {
        public readonly string? AllocationType;
        public readonly bool? Attached;
        public readonly string? AttachmentStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NamePrefix;
        /// <summary>
        /// A List of `public_ips` blocks as defined below filtered by the criteria above.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPublicIPsPublicIpResult> PublicIps;
        public readonly string ResourceGroupName;

        [OutputConstructor]
        private GetPublicIPsResult(
            string? allocationType,

            bool? attached,

            string? attachmentStatus,

            string id,

            string? namePrefix,

            ImmutableArray<Outputs.GetPublicIPsPublicIpResult> publicIps,

            string resourceGroupName)
        {
            AllocationType = allocationType;
            Attached = attached;
            AttachmentStatus = attachmentStatus;
            Id = id;
            NamePrefix = namePrefix;
            PublicIps = publicIps;
            ResourceGroupName = resourceGroupName;
        }
    }
}
