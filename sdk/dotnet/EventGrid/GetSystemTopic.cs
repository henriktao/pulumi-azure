// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.EventGrid
{
    public static class GetSystemTopic
    {
        /// <summary>
        /// Use this data source to access information about an existing EventGrid System Topic
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
        ///         var example = Output.Create(Azure.EventGrid.GetSystemTopic.InvokeAsync(new Azure.EventGrid.GetSystemTopicArgs
        ///         {
        ///             Name = "eventgrid-system-topic",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSystemTopicResult> InvokeAsync(GetSystemTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSystemTopicResult>("azure:eventgrid/getSystemTopic:getSystemTopic", args ?? new GetSystemTopicArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing EventGrid System Topic
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
        ///         var example = Output.Create(Azure.EventGrid.GetSystemTopic.InvokeAsync(new Azure.EventGrid.GetSystemTopicArgs
        ///         {
        ///             Name = "eventgrid-system-topic",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSystemTopicResult> Invoke(GetSystemTopicInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSystemTopicResult>("azure:eventgrid/getSystemTopic:getSystemTopic", args ?? new GetSystemTopicInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSystemTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
        /// </summary>
        [Input("identity")]
        public Inputs.GetSystemTopicIdentityArgs? Identity { get; set; }

        /// <summary>
        /// The name of the EventGrid System Topic resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid System Topic exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSystemTopicArgs()
        {
        }
    }

    public sealed class GetSystemTopicInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.GetSystemTopicIdentityInputArgs>? Identity { get; set; }

        /// <summary>
        /// The name of the EventGrid System Topic resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventGrid System Topic exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSystemTopicInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSystemTopicResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identity` block as defined below, which contains the Managed Service Identity information for this Event Grid System Topic.
        /// </summary>
        public readonly Outputs.GetSystemTopicIdentityResult Identity;
        public readonly string Location;
        /// <summary>
        /// The Metric ARM Resource ID of the Event Grid System Topic.
        /// </summary>
        public readonly string MetricArmResourceId;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The ID of the Event Grid System Topic ARM Source.
        /// </summary>
        public readonly string SourceArmResourceId;
        /// <summary>
        /// A mapping of tags which are assigned to the Event Grid System Topic.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The Topic Type of the Event Grid System Topic.
        /// </summary>
        public readonly string TopicType;

        [OutputConstructor]
        private GetSystemTopicResult(
            string id,

            Outputs.GetSystemTopicIdentityResult identity,

            string location,

            string metricArmResourceId,

            string name,

            string resourceGroupName,

            string sourceArmResourceId,

            ImmutableDictionary<string, string> tags,

            string topicType)
        {
            Id = id;
            Identity = identity;
            Location = location;
            MetricArmResourceId = metricArmResourceId;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SourceArmResourceId = sourceArmResourceId;
            Tags = tags;
            TopicType = topicType;
        }
    }
}
