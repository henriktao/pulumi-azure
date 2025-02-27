// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NotificationHub
{
    public static class GetHub
    {
        /// <summary>
        /// Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.
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
        ///         var example = Output.Create(Azure.NotificationHub.GetHub.InvokeAsync(new Azure.NotificationHub.GetHubArgs
        ///         {
        ///             Name = "notification-hub",
        ///             NamespaceName = "namespace-name",
        ///             ResourceGroupName = "resource-group-name",
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
        public static Task<GetHubResult> InvokeAsync(GetHubArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHubResult>("azure:notificationhub/getHub:getHub", args ?? new GetHubArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.
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
        ///         var example = Output.Create(Azure.NotificationHub.GetHub.InvokeAsync(new Azure.NotificationHub.GetHubArgs
        ///         {
        ///             Name = "notification-hub",
        ///             NamespaceName = "namespace-name",
        ///             ResourceGroupName = "resource-group-name",
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
        public static Output<GetHubResult> Invoke(GetHubInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHubResult>("azure:notificationhub/getHub:getHub", args ?? new GetHubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHubArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Notification Hub.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Notification Hub Namespace which contains the Notification Hub.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Notification Hub exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHubArgs()
        {
        }
    }

    public sealed class GetHubInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Name of the Notification Hub.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Notification Hub Namespace which contains the Notification Hub.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Specifies the Name of the Resource Group within which the Notification Hub exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetHubInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHubResult
    {
        /// <summary>
        /// A `apns_credential` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHubApnsCredentialResult> ApnsCredentials;
        /// <summary>
        /// A `gcm_credential` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHubGcmCredentialResult> GcmCredentials;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which this Notification Hub exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string NamespaceName;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetHubResult(
            ImmutableArray<Outputs.GetHubApnsCredentialResult> apnsCredentials,

            ImmutableArray<Outputs.GetHubGcmCredentialResult> gcmCredentials,

            string id,

            string location,

            string name,

            string namespaceName,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            ApnsCredentials = apnsCredentials;
            GcmCredentials = gcmCredentials;
            Id = id;
            Location = location;
            Name = name;
            NamespaceName = namespaceName;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
