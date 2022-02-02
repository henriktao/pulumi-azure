// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid
{
    public static class GetDomain
    {
        /// <summary>
        /// Use this data source to access information about an existing EventGrid Domain
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
        ///         var example = Output.Create(Azure.EventGrid.GetDomain.InvokeAsync(new Azure.EventGrid.GetDomainArgs
        ///         {
        ///             Name = "my-eventgrid-domain",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.EventgridDomainMappingTopic = example.Apply(example =&gt; example.InputMappingFields?[0]?.Topic);
        ///     }
        /// 
        ///     [Output("eventgridDomainMappingTopic")]
        ///     public Output&lt;string&gt; EventgridDomainMappingTopic { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("azure:eventgrid/getDomain:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing EventGrid Domain
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
        ///         var example = Output.Create(Azure.EventGrid.GetDomain.InvokeAsync(new Azure.EventGrid.GetDomainArgs
        ///         {
        ///             Name = "my-eventgrid-domain",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.EventgridDomainMappingTopic = example.Apply(example =&gt; example.InputMappingFields?[0]?.Topic);
        ///     }
        /// 
        ///     [Output("eventgridDomainMappingTopic")]
        ///     public Output&lt;string&gt; EventgridDomainMappingTopic { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainResult>("azure:eventgrid/getDomain:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        [Input("inboundIpRules")]
        private List<Inputs.GetDomainInboundIpRuleArgs>? _inboundIpRules;

        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        public List<Inputs.GetDomainInboundIpRuleArgs> InboundIpRules
        {
            get => _inboundIpRules ?? (_inboundIpRules = new List<Inputs.GetDomainInboundIpRuleArgs>());
            set => _inboundIpRules = value;
        }

        /// <summary>
        /// The name of the EventGrid Domain resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this server.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public bool? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the EventGrid Domain.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDomainArgs()
        {
        }
    }

    public sealed class GetDomainInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("inboundIpRules")]
        private InputList<Inputs.GetDomainInboundIpRuleInputArgs>? _inboundIpRules;

        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.GetDomainInboundIpRuleInputArgs> InboundIpRules
        {
            get => _inboundIpRules ?? (_inboundIpRules = new InputList<Inputs.GetDomainInboundIpRuleInputArgs>());
            set => _inboundIpRules = value;
        }

        /// <summary>
        /// The name of the EventGrid Domain resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this server.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which the EventGrid Domain exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the EventGrid Domain.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetDomainInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The Endpoint associated with the EventGrid Domain.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more `inbound_ip_rule` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainInboundIpRuleResult> InboundIpRules;
        /// <summary>
        /// A `input_mapping_default_values` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainInputMappingDefaultValueResult> InputMappingDefaultValues;
        /// <summary>
        /// A `input_mapping_fields` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainInputMappingFieldResult> InputMappingFields;
        /// <summary>
        /// The schema in which incoming events will be published to this domain. Possible values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`.
        /// </summary>
        public readonly string InputSchema;
        /// <summary>
        /// The Azure Region in which this EventGrid Domain exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The primary access key associated with the EventGrid Domain.
        /// </summary>
        public readonly string PrimaryAccessKey;
        /// <summary>
        /// Whether or not public network access is allowed for this server.
        /// </summary>
        public readonly bool? PublicNetworkAccessEnabled;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The secondary access key associated with the EventGrid Domain.
        /// </summary>
        public readonly string SecondaryAccessKey;
        /// <summary>
        /// A mapping of tags assigned to the EventGrid Domain.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetDomainResult(
            string endpoint,

            string id,

            ImmutableArray<Outputs.GetDomainInboundIpRuleResult> inboundIpRules,

            ImmutableArray<Outputs.GetDomainInputMappingDefaultValueResult> inputMappingDefaultValues,

            ImmutableArray<Outputs.GetDomainInputMappingFieldResult> inputMappingFields,

            string inputSchema,

            string location,

            string name,

            string primaryAccessKey,

            bool? publicNetworkAccessEnabled,

            string resourceGroupName,

            string secondaryAccessKey,

            ImmutableDictionary<string, string>? tags)
        {
            Endpoint = endpoint;
            Id = id;
            InboundIpRules = inboundIpRules;
            InputMappingDefaultValues = inputMappingDefaultValues;
            InputMappingFields = inputMappingFields;
            InputSchema = inputSchema;
            Location = location;
            Name = name;
            PrimaryAccessKey = primaryAccessKey;
            PublicNetworkAccessEnabled = publicNetworkAccessEnabled;
            ResourceGroupName = resourceGroupName;
            SecondaryAccessKey = secondaryAccessKey;
            Tags = tags;
        }
    }
}
