// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.StreamAnalytics
{
    public static class GetJob
    {
        /// <summary>
        /// Use this data source to access information about an existing Stream Analytics Job.
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
        ///         var example = Output.Create(Azure.StreamAnalytics.GetJob.InvokeAsync(new Azure.StreamAnalytics.GetJobArgs
        ///         {
        ///             Name = "example-job",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.JobId = example.Apply(example =&gt; example.JobId);
        ///     }
        /// 
        ///     [Output("jobId")]
        ///     public Output&lt;string&gt; JobId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("azure:streamanalytics/getJob:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Stream Analytics Job.
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
        ///         var example = Output.Create(Azure.StreamAnalytics.GetJob.InvokeAsync(new Azure.StreamAnalytics.GetJobArgs
        ///         {
        ///             Name = "example-job",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.JobId = example.Apply(example =&gt; example.JobId);
        ///     }
        /// 
        ///     [Output("jobId")]
        ///     public Output&lt;string&gt; JobId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetJobResult>("azure:streamanalytics/getJob:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Stream Analytics Job.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Stream Analytics Job is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetJobArgs()
        {
        }
    }

    public sealed class GetJobInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Stream Analytics Job.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Stream Analytics Job is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// The compatibility level for this job.
        /// </summary>
        public readonly string CompatibilityLevel;
        /// <summary>
        /// The Data Locale of the Job.
        /// </summary>
        public readonly string DataLocale;
        /// <summary>
        /// The maximum tolerable delay in seconds where events arriving late could be included.
        /// </summary>
        public readonly int EventsLateArrivalMaxDelayInSeconds;
        /// <summary>
        /// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
        /// </summary>
        public readonly int EventsOutOfOrderMaxDelayInSeconds;
        /// <summary>
        /// The policy which should be applied to events which arrive out of order in the input event stream.
        /// </summary>
        public readonly string EventsOutOfOrderPolicy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Optional) An `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetJobIdentityResult> Identities;
        /// <summary>
        /// The Job ID assigned by the Stream Analytics Job.
        /// </summary>
        public readonly string JobId;
        /// <summary>
        /// The Azure location where the Stream Analytics Job exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).
        /// </summary>
        public readonly string OutputErrorPolicy;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The number of streaming units that the streaming job uses.
        /// </summary>
        public readonly int StreamingUnits;
        /// <summary>
        /// The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
        /// </summary>
        public readonly string TransformationQuery;

        [OutputConstructor]
        private GetJobResult(
            string compatibilityLevel,

            string dataLocale,

            int eventsLateArrivalMaxDelayInSeconds,

            int eventsOutOfOrderMaxDelayInSeconds,

            string eventsOutOfOrderPolicy,

            string id,

            ImmutableArray<Outputs.GetJobIdentityResult> identities,

            string jobId,

            string location,

            string name,

            string outputErrorPolicy,

            string resourceGroupName,

            int streamingUnits,

            string transformationQuery)
        {
            CompatibilityLevel = compatibilityLevel;
            DataLocale = dataLocale;
            EventsLateArrivalMaxDelayInSeconds = eventsLateArrivalMaxDelayInSeconds;
            EventsOutOfOrderMaxDelayInSeconds = eventsOutOfOrderMaxDelayInSeconds;
            EventsOutOfOrderPolicy = eventsOutOfOrderPolicy;
            Id = id;
            Identities = identities;
            JobId = jobId;
            Location = location;
            Name = name;
            OutputErrorPolicy = outputErrorPolicy;
            ResourceGroupName = resourceGroupName;
            StreamingUnits = streamingUnits;
            TransformationQuery = transformationQuery;
        }
    }
}
