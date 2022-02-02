// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    [Obsolete(@"azure.eventhub.getServiceBusNamespace has been deprecated in favor of azure.servicebus.getNamespace")]
    public static class GetServiceBusNamespace
    {
        /// <summary>
        /// Use this data source to access information about an existing ServiceBus Namespace.
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
        ///         var example = Output.Create(Azure.ServiceBus.GetNamespace.InvokeAsync(new Azure.ServiceBus.GetNamespaceArgs
        ///         {
        ///             Name = "examplenamespace",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.Location = example.Apply(example =&gt; example.Location);
        ///     }
        /// 
        ///     [Output("location")]
        ///     public Output&lt;string&gt; Location { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceBusNamespaceResult> InvokeAsync(GetServiceBusNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceBusNamespaceResult>("azure:eventhub/getServiceBusNamespace:getServiceBusNamespace", args ?? new GetServiceBusNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing ServiceBus Namespace.
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
        ///         var example = Output.Create(Azure.ServiceBus.GetNamespace.InvokeAsync(new Azure.ServiceBus.GetNamespaceArgs
        ///         {
        ///             Name = "examplenamespace",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.Location = example.Apply(example =&gt; example.Location);
        ///     }
        /// 
        ///     [Output("location")]
        ///     public Output&lt;string&gt; Location { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceBusNamespaceResult> Invoke(GetServiceBusNamespaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceBusNamespaceResult>("azure:eventhub/getServiceBusNamespace:getServiceBusNamespace", args ?? new GetServiceBusNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceBusNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the ServiceBus Namespace.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceBusNamespaceArgs()
        {
        }
    }

    public sealed class GetServiceBusNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the ServiceBus Namespace.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetServiceBusNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceBusNamespaceResult
    {
        /// <summary>
        /// The capacity of the ServiceBus Namespace.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// The primary connection string for the authorization
        /// rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultPrimaryConnectionString;
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
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        public readonly string DefaultSecondaryKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of the Resource Group in which the ServiceBus Namespace exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Tier used for the ServiceBus Namespace.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Whether or not this ServiceBus Namespace is zone redundant.
        /// </summary>
        public readonly bool ZoneRedundant;

        [OutputConstructor]
        private GetServiceBusNamespaceResult(
            int capacity,

            string defaultPrimaryConnectionString,

            string defaultPrimaryKey,

            string defaultSecondaryConnectionString,

            string defaultSecondaryKey,

            string id,

            string location,

            string name,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string> tags,

            bool zoneRedundant)
        {
            Capacity = capacity;
            DefaultPrimaryConnectionString = defaultPrimaryConnectionString;
            DefaultPrimaryKey = defaultPrimaryKey;
            DefaultSecondaryConnectionString = defaultSecondaryConnectionString;
            DefaultSecondaryKey = defaultSecondaryKey;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            ZoneRedundant = zoneRedundant;
        }
    }
}
