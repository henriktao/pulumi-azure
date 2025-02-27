// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.StreamAnalytics
{
    /// <summary>
    /// Manages a Stream Analytics Output to an EventHub.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
    ///         {
    ///             Name = "example-resources",
    ///         }));
    ///         var exampleJob = Output.Create(Azure.StreamAnalytics.GetJob.InvokeAsync(new Azure.StreamAnalytics.GetJobArgs
    ///         {
    ///             Name = "example-job",
    ///             ResourceGroupName = azurerm_resource_group.Example.Name,
    ///         }));
    ///         var exampleEventHubNamespace = new Azure.EventHub.EventHubNamespace("exampleEventHubNamespace", new Azure.EventHub.EventHubNamespaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Apply(exampleResourceGroup =&gt; exampleResourceGroup.Location),
    ///             ResourceGroupName = exampleResourceGroup.Apply(exampleResourceGroup =&gt; exampleResourceGroup.Name),
    ///             Sku = "Standard",
    ///             Capacity = 1,
    ///         });
    ///         var exampleEventHub = new Azure.EventHub.EventHub("exampleEventHub", new Azure.EventHub.EventHubArgs
    ///         {
    ///             NamespaceName = exampleEventHubNamespace.Name,
    ///             ResourceGroupName = exampleResourceGroup.Apply(exampleResourceGroup =&gt; exampleResourceGroup.Name),
    ///             PartitionCount = 2,
    ///             MessageRetention = 1,
    ///         });
    ///         var exampleOutputEventHub = new Azure.StreamAnalytics.OutputEventHub("exampleOutputEventHub", new Azure.StreamAnalytics.OutputEventHubArgs
    ///         {
    ///             StreamAnalyticsJobName = exampleJob.Apply(exampleJob =&gt; exampleJob.Name),
    ///             ResourceGroupName = exampleJob.Apply(exampleJob =&gt; exampleJob.ResourceGroupName),
    ///             EventhubName = exampleEventHub.Name,
    ///             ServicebusNamespace = exampleEventHubNamespace.Name,
    ///             SharedAccessPolicyKey = exampleEventHubNamespace.DefaultPrimaryKey,
    ///             SharedAccessPolicyName = "RootManageSharedAccessKey",
    ///             Serialization = new Azure.StreamAnalytics.Inputs.OutputEventHubSerializationArgs
    ///             {
    ///                 Type = "Avro",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Stream Analytics Outputs to an EventHub can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:streamanalytics/outputEventHub:OutputEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
    /// ```
    /// </summary>
    [AzureResourceType("azure:streamanalytics/outputEventHub:OutputEventHub")]
    public partial class OutputEventHub : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Event Hub.
        /// </summary>
        [Output("eventhubName")]
        public Output<string> EventhubName { get; private set; } = null!;

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The column that is used for the Event Hub partition key.
        /// </summary>
        [Output("partitionKey")]
        public Output<string?> PartitionKey { get; private set; } = null!;

        /// <summary>
        /// A list of property columns to add to the Event Hub output.
        /// </summary>
        [Output("propertyColumns")]
        public Output<ImmutableArray<string>> PropertyColumns { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Output("serialization")]
        public Output<Outputs.OutputEventHubSerialization> Serialization { get; private set; } = null!;

        /// <summary>
        /// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Output("servicebusNamespace")]
        public Output<string> ServicebusNamespace { get; private set; } = null!;

        /// <summary>
        /// The shared access policy key for the specified shared access policy.
        /// </summary>
        [Output("sharedAccessPolicyKey")]
        public Output<string> SharedAccessPolicyKey { get; private set; } = null!;

        /// <summary>
        /// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Output("sharedAccessPolicyName")]
        public Output<string> SharedAccessPolicyName { get; private set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Output("streamAnalyticsJobName")]
        public Output<string> StreamAnalyticsJobName { get; private set; } = null!;


        /// <summary>
        /// Create a OutputEventHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutputEventHub(string name, OutputEventHubArgs args, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/outputEventHub:OutputEventHub", name, args ?? new OutputEventHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OutputEventHub(string name, Input<string> id, OutputEventHubState? state = null, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/outputEventHub:OutputEventHub", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OutputEventHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutputEventHub Get(string name, Input<string> id, OutputEventHubState? state = null, CustomResourceOptions? options = null)
        {
            return new OutputEventHub(name, id, state, options);
        }
    }

    public sealed class OutputEventHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Event Hub.
        /// </summary>
        [Input("eventhubName", required: true)]
        public Input<string> EventhubName { get; set; } = null!;

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The column that is used for the Event Hub partition key.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        [Input("propertyColumns")]
        private InputList<string>? _propertyColumns;

        /// <summary>
        /// A list of property columns to add to the Event Hub output.
        /// </summary>
        public InputList<string> PropertyColumns
        {
            get => _propertyColumns ?? (_propertyColumns = new InputList<string>());
            set => _propertyColumns = value;
        }

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Input("serialization", required: true)]
        public Input<Inputs.OutputEventHubSerializationArgs> Serialization { get; set; } = null!;

        /// <summary>
        /// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Input("servicebusNamespace", required: true)]
        public Input<string> ServicebusNamespace { get; set; } = null!;

        /// <summary>
        /// The shared access policy key for the specified shared access policy.
        /// </summary>
        [Input("sharedAccessPolicyKey", required: true)]
        public Input<string> SharedAccessPolicyKey { get; set; } = null!;

        /// <summary>
        /// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Input("sharedAccessPolicyName", required: true)]
        public Input<string> SharedAccessPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName", required: true)]
        public Input<string> StreamAnalyticsJobName { get; set; } = null!;

        public OutputEventHubArgs()
        {
        }
    }

    public sealed class OutputEventHubState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Event Hub.
        /// </summary>
        [Input("eventhubName")]
        public Input<string>? EventhubName { get; set; }

        /// <summary>
        /// The name of the Stream Output. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The column that is used for the Event Hub partition key.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        [Input("propertyColumns")]
        private InputList<string>? _propertyColumns;

        /// <summary>
        /// A list of property columns to add to the Event Hub output.
        /// </summary>
        public InputList<string> PropertyColumns
        {
            get => _propertyColumns ?? (_propertyColumns = new InputList<string>());
            set => _propertyColumns = value;
        }

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `serialization` block as defined below.
        /// </summary>
        [Input("serialization")]
        public Input<Inputs.OutputEventHubSerializationGetArgs>? Serialization { get; set; }

        /// <summary>
        /// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Input("servicebusNamespace")]
        public Input<string>? ServicebusNamespace { get; set; }

        /// <summary>
        /// The shared access policy key for the specified shared access policy.
        /// </summary>
        [Input("sharedAccessPolicyKey")]
        public Input<string>? SharedAccessPolicyKey { get; set; }

        /// <summary>
        /// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        /// </summary>
        [Input("sharedAccessPolicyName")]
        public Input<string>? SharedAccessPolicyName { get; set; }

        /// <summary>
        /// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName")]
        public Input<string>? StreamAnalyticsJobName { get; set; }

        public OutputEventHubState()
        {
        }
    }
}
