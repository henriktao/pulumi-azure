// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    [Obsolete(@"azure.eventhub.getEventhubNamespace has been deprecated in favor of azure.eventhub.getNamespace")]
    public static class GetEventhubNamespace
    {
        /// <summary>
        /// Use this data source to access information about an existing EventHub Namespace.
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
        ///         var example = Output.Create(Azure.EventHub.GetNamespace.InvokeAsync(new Azure.EventHub.GetNamespaceArgs
        ///         {
        ///             Name = "search-eventhubns",
        ///             ResourceGroupName = "search-service",
        ///         }));
        ///         this.EventhubNamespaceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("eventhubNamespaceId")]
        ///     public Output&lt;string&gt; EventhubNamespaceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEventhubNamespaceResult> InvokeAsync(GetEventhubNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventhubNamespaceResult>("azure:eventhub/getEventhubNamespace:getEventhubNamespace", args ?? new GetEventhubNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing EventHub Namespace.
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
        ///         var example = Output.Create(Azure.EventHub.GetNamespace.InvokeAsync(new Azure.EventHub.GetNamespaceArgs
        ///         {
        ///             Name = "search-eventhubns",
        ///             ResourceGroupName = "search-service",
        ///         }));
        ///         this.EventhubNamespaceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("eventhubNamespaceId")]
        ///     public Output&lt;string&gt; EventhubNamespaceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEventhubNamespaceResult> Invoke(GetEventhubNamespaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEventhubNamespaceResult>("azure:eventhub/getEventhubNamespace:getEventhubNamespace", args ?? new GetEventhubNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventhubNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventHub Namespace.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the EventHub Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEventhubNamespaceArgs()
        {
        }
    }

    public sealed class GetEventhubNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the EventHub Namespace.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the EventHub Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetEventhubNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventhubNamespaceResult
    {
        /// <summary>
        /// Is Auto Inflate enabled for the EventHub Namespace?
        /// </summary>
        public readonly bool AutoInflateEnabled;
        /// <summary>
        /// The Capacity / Throughput Units for a `Standard` SKU namespace.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// The ID of the EventHub Dedicated Cluster where this Namespace exists.
        /// </summary>
        public readonly string DedicatedClusterId;
        /// <summary>
        /// The primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultPrimaryConnectionString;
        /// <summary>
        /// The alias of the primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultPrimaryConnectionStringAlias;
        /// <summary>
        /// The primary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultPrimaryKey;
        /// <summary>
        /// The secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultSecondaryConnectionString;
        /// <summary>
        /// The alias of the secondary connection string for the
        /// authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultSecondaryConnectionStringAlias;
        /// <summary>
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultSecondaryKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool KafkaEnabled;
        /// <summary>
        /// The Azure location where the EventHub Namespace exists
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Specifies the maximum number of throughput units when Auto Inflate is Enabled.
        /// </summary>
        public readonly int MaximumThroughputUnits;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// Defines which tier to use.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags to assign to the EventHub Namespace.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Is this EventHub Namespace deployed across Availability Zones?
        /// </summary>
        public readonly bool ZoneRedundant;

        [OutputConstructor]
        private GetEventhubNamespaceResult(
            bool autoInflateEnabled,

            int capacity,

            string dedicatedClusterId,

            string defaultPrimaryConnectionString,

            string defaultPrimaryConnectionStringAlias,

            string defaultPrimaryKey,

            string defaultSecondaryConnectionString,

            string defaultSecondaryConnectionStringAlias,

            string defaultSecondaryKey,

            string id,

            bool kafkaEnabled,

            string location,

            int maximumThroughputUnits,

            string name,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags,

            bool zoneRedundant)
        {
            AutoInflateEnabled = autoInflateEnabled;
            Capacity = capacity;
            DedicatedClusterId = dedicatedClusterId;
            DefaultPrimaryConnectionString = defaultPrimaryConnectionString;
            DefaultPrimaryConnectionStringAlias = defaultPrimaryConnectionStringAlias;
            DefaultPrimaryKey = defaultPrimaryKey;
            DefaultSecondaryConnectionString = defaultSecondaryConnectionString;
            DefaultSecondaryConnectionStringAlias = defaultSecondaryConnectionStringAlias;
            DefaultSecondaryKey = defaultSecondaryKey;
            Id = id;
            KafkaEnabled = kafkaEnabled;
            Location = location;
            MaximumThroughputUnits = maximumThroughputUnits;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            ZoneRedundant = zoneRedundant;
        }
    }
}
