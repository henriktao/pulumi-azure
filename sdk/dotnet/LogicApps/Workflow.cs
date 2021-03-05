// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps
{
    /// <summary>
    /// Manages a Logic App Workflow.
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
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleWorkflow = new Azure.LogicApps.Workflow("exampleWorkflow", new Azure.LogicApps.WorkflowArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Logic App Workflows can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:logicapps/workflow:Workflow workflow1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
    /// ```
    /// </summary>
    [AzureResourceType("azure:logicapps/workflow:Workflow")]
    public partial class Workflow : Pulumi.CustomResource
    {
        /// <summary>
        /// The Access Endpoint for the Logic App Workflow.
        /// </summary>
        [Output("accessEndpoint")]
        public Output<string> AccessEndpoint { get; private set; } = null!;

        /// <summary>
        /// The list of access endpoint ip addresses of connector.
        /// </summary>
        [Output("connectorEndpointIpAddresses")]
        public Output<ImmutableArray<string>> ConnectorEndpointIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The list of outgoing ip addresses of connector.
        /// </summary>
        [Output("connectorOutboundIpAddresses")]
        public Output<ImmutableArray<string>> ConnectorOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        /// </summary>
        [Output("integrationServiceEnvironmentId")]
        public Output<string?> IntegrationServiceEnvironmentId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the integration account linked by this Logic App Workflow.
        /// </summary>
        [Output("logicAppIntegrationAccountId")]
        public Output<string?> LogicAppIntegrationAccountId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of Key-Value pairs.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The list of access endpoint ip addresses of workflow.
        /// </summary>
        [Output("workflowEndpointIpAddresses")]
        public Output<ImmutableArray<string>> WorkflowEndpointIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The list of outgoing ip addresses of workflow.
        /// </summary>
        [Output("workflowOutboundIpAddresses")]
        public Output<ImmutableArray<string>> WorkflowOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("workflowSchema")]
        public Output<string?> WorkflowSchema { get; private set; } = null!;

        /// <summary>
        /// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("workflowVersion")]
        public Output<string?> WorkflowVersion { get; private set; } = null!;


        /// <summary>
        /// Create a Workflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workflow(string name, WorkflowArgs args, CustomResourceOptions? options = null)
            : base("azure:logicapps/workflow:Workflow", name, args ?? new WorkflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workflow(string name, Input<string> id, WorkflowState? state = null, CustomResourceOptions? options = null)
            : base("azure:logicapps/workflow:Workflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workflow Get(string name, Input<string> id, WorkflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Workflow(name, id, state, options);
        }
    }

    public sealed class WorkflowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        /// </summary>
        [Input("integrationServiceEnvironmentId")]
        public Input<string>? IntegrationServiceEnvironmentId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the integration account linked by this Logic App Workflow.
        /// </summary>
        [Input("logicAppIntegrationAccountId")]
        public Input<string>? LogicAppIntegrationAccountId { get; set; }

        /// <summary>
        /// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of Key-Value pairs.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("workflowSchema")]
        public Input<string>? WorkflowSchema { get; set; }

        /// <summary>
        /// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("workflowVersion")]
        public Input<string>? WorkflowVersion { get; set; }

        public WorkflowArgs()
        {
        }
    }

    public sealed class WorkflowState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Endpoint for the Logic App Workflow.
        /// </summary>
        [Input("accessEndpoint")]
        public Input<string>? AccessEndpoint { get; set; }

        [Input("connectorEndpointIpAddresses")]
        private InputList<string>? _connectorEndpointIpAddresses;

        /// <summary>
        /// The list of access endpoint ip addresses of connector.
        /// </summary>
        public InputList<string> ConnectorEndpointIpAddresses
        {
            get => _connectorEndpointIpAddresses ?? (_connectorEndpointIpAddresses = new InputList<string>());
            set => _connectorEndpointIpAddresses = value;
        }

        [Input("connectorOutboundIpAddresses")]
        private InputList<string>? _connectorOutboundIpAddresses;

        /// <summary>
        /// The list of outgoing ip addresses of connector.
        /// </summary>
        public InputList<string> ConnectorOutboundIpAddresses
        {
            get => _connectorOutboundIpAddresses ?? (_connectorOutboundIpAddresses = new InputList<string>());
            set => _connectorOutboundIpAddresses = value;
        }

        /// <summary>
        /// The ID of the Integration Service Environment to which this Logic App Workflow belongs.  Changing this forces a new Logic App Workflow to be created.
        /// </summary>
        [Input("integrationServiceEnvironmentId")]
        public Input<string>? IntegrationServiceEnvironmentId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the integration account linked by this Logic App Workflow.
        /// </summary>
        [Input("logicAppIntegrationAccountId")]
        public Input<string>? LogicAppIntegrationAccountId { get; set; }

        /// <summary>
        /// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// A map of Key-Value pairs.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("workflowEndpointIpAddresses")]
        private InputList<string>? _workflowEndpointIpAddresses;

        /// <summary>
        /// The list of access endpoint ip addresses of workflow.
        /// </summary>
        public InputList<string> WorkflowEndpointIpAddresses
        {
            get => _workflowEndpointIpAddresses ?? (_workflowEndpointIpAddresses = new InputList<string>());
            set => _workflowEndpointIpAddresses = value;
        }

        [Input("workflowOutboundIpAddresses")]
        private InputList<string>? _workflowOutboundIpAddresses;

        /// <summary>
        /// The list of outgoing ip addresses of workflow.
        /// </summary>
        public InputList<string> WorkflowOutboundIpAddresses
        {
            get => _workflowOutboundIpAddresses ?? (_workflowOutboundIpAddresses = new InputList<string>());
            set => _workflowOutboundIpAddresses = value;
        }

        /// <summary>
        /// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("workflowSchema")]
        public Input<string>? WorkflowSchema { get; set; }

        /// <summary>
        /// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("workflowVersion")]
        public Input<string>? WorkflowVersion { get; set; }

        public WorkflowState()
        {
        }
    }
}
