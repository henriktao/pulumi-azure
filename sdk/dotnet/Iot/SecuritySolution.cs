// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages an iot security solution.
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
    ///         var exampleIoTHub = new Azure.Iot.IoTHub("exampleIoTHub", new Azure.Iot.IoTHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IoTHubSkuArgs
    ///             {
    ///                 Name = "S1",
    ///                 Capacity = 1,
    ///             },
    ///         });
    ///         var exampleSecuritySolution = new Azure.Iot.SecuritySolution("exampleSecuritySolution", new Azure.Iot.SecuritySolutionArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             DisplayName = "Iot Security Solution",
    ///             IothubIds = 
    ///             {
    ///                 exampleIoTHub.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iot Security Solution can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:iot/securitySolution:SecuritySolution example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Security/IoTSecuritySolutions/solution1
    /// ```
    /// </summary>
    [AzureResourceType("azure:iot/securitySolution:SecuritySolution")]
    public partial class SecuritySolution : Pulumi.CustomResource
    {
        /// <summary>
        /// A `additional_workspace` block as defined below.
        /// </summary>
        [Output("additionalWorkspaces")]
        public Output<ImmutableArray<Outputs.SecuritySolutionAdditionalWorkspace>> AdditionalWorkspaces { get; private set; } = null!;

        /// <summary>
        /// A list of disabled data sources for the Iot Security Solution. Possible value is `TwinData`.
        /// </summary>
        [Output("disabledDataSources")]
        public Output<ImmutableArray<string>> DisabledDataSources { get; private set; } = null!;

        /// <summary>
        /// Specifies the Display Name for this Iot Security Solution.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Is the Iot Security Solution enabled? Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        /// </summary>
        [Output("eventsToExports")]
        public Output<ImmutableArray<string>> EventsToExports { get; private set; } = null!;

        /// <summary>
        /// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        /// </summary>
        [Output("iothubIds")]
        public Output<ImmutableArray<string>> IothubIds { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the Log Analytics Workspace ID to which the security data will be sent.
        /// </summary>
        [Output("logAnalyticsWorkspaceId")]
        public Output<string?> LogAnalyticsWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// Should ip addressed be unmasked in the log? Defaults to `false`.
        /// </summary>
        [Output("logUnmaskedIpsEnabled")]
        public Output<bool?> LogUnmaskedIpsEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An Azure Resource Graph query used to set the resources monitored.
        /// </summary>
        [Output("queryForResources")]
        public Output<string> QueryForResources { get; private set; } = null!;

        /// <summary>
        /// A list of subscription Ids on which the user defined resources query should be executed.
        /// </summary>
        [Output("querySubscriptionIds")]
        public Output<ImmutableArray<string>> QuerySubscriptionIds { get; private set; } = null!;

        /// <summary>
        /// A `recommendations_enabled` block of options to enable or disable as defined below.
        /// </summary>
        [Output("recommendationsEnabled")]
        public Output<Outputs.SecuritySolutionRecommendationsEnabled> RecommendationsEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a SecuritySolution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecuritySolution(string name, SecuritySolutionArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/securitySolution:SecuritySolution", name, args ?? new SecuritySolutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecuritySolution(string name, Input<string> id, SecuritySolutionState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/securitySolution:SecuritySolution", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecuritySolution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecuritySolution Get(string name, Input<string> id, SecuritySolutionState? state = null, CustomResourceOptions? options = null)
        {
            return new SecuritySolution(name, id, state, options);
        }
    }

    public sealed class SecuritySolutionArgs : Pulumi.ResourceArgs
    {
        [Input("additionalWorkspaces")]
        private InputList<Inputs.SecuritySolutionAdditionalWorkspaceArgs>? _additionalWorkspaces;

        /// <summary>
        /// A `additional_workspace` block as defined below.
        /// </summary>
        public InputList<Inputs.SecuritySolutionAdditionalWorkspaceArgs> AdditionalWorkspaces
        {
            get => _additionalWorkspaces ?? (_additionalWorkspaces = new InputList<Inputs.SecuritySolutionAdditionalWorkspaceArgs>());
            set => _additionalWorkspaces = value;
        }

        [Input("disabledDataSources")]
        private InputList<string>? _disabledDataSources;

        /// <summary>
        /// A list of disabled data sources for the Iot Security Solution. Possible value is `TwinData`.
        /// </summary>
        public InputList<string> DisabledDataSources
        {
            get => _disabledDataSources ?? (_disabledDataSources = new InputList<string>());
            set => _disabledDataSources = value;
        }

        /// <summary>
        /// Specifies the Display Name for this Iot Security Solution.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Is the Iot Security Solution enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventsToExports")]
        private InputList<string>? _eventsToExports;

        /// <summary>
        /// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        /// </summary>
        public InputList<string> EventsToExports
        {
            get => _eventsToExports ?? (_eventsToExports = new InputList<string>());
            set => _eventsToExports = value;
        }

        [Input("iothubIds", required: true)]
        private InputList<string>? _iothubIds;

        /// <summary>
        /// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        /// </summary>
        public InputList<string> IothubIds
        {
            get => _iothubIds ?? (_iothubIds = new InputList<string>());
            set => _iothubIds = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the Log Analytics Workspace ID to which the security data will be sent.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        /// <summary>
        /// Should ip addressed be unmasked in the log? Defaults to `false`.
        /// </summary>
        [Input("logUnmaskedIpsEnabled")]
        public Input<bool>? LogUnmaskedIpsEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An Azure Resource Graph query used to set the resources monitored.
        /// </summary>
        [Input("queryForResources")]
        public Input<string>? QueryForResources { get; set; }

        [Input("querySubscriptionIds")]
        private InputList<string>? _querySubscriptionIds;

        /// <summary>
        /// A list of subscription Ids on which the user defined resources query should be executed.
        /// </summary>
        public InputList<string> QuerySubscriptionIds
        {
            get => _querySubscriptionIds ?? (_querySubscriptionIds = new InputList<string>());
            set => _querySubscriptionIds = value;
        }

        /// <summary>
        /// A `recommendations_enabled` block of options to enable or disable as defined below.
        /// </summary>
        [Input("recommendationsEnabled")]
        public Input<Inputs.SecuritySolutionRecommendationsEnabledArgs>? RecommendationsEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
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

        public SecuritySolutionArgs()
        {
        }
    }

    public sealed class SecuritySolutionState : Pulumi.ResourceArgs
    {
        [Input("additionalWorkspaces")]
        private InputList<Inputs.SecuritySolutionAdditionalWorkspaceGetArgs>? _additionalWorkspaces;

        /// <summary>
        /// A `additional_workspace` block as defined below.
        /// </summary>
        public InputList<Inputs.SecuritySolutionAdditionalWorkspaceGetArgs> AdditionalWorkspaces
        {
            get => _additionalWorkspaces ?? (_additionalWorkspaces = new InputList<Inputs.SecuritySolutionAdditionalWorkspaceGetArgs>());
            set => _additionalWorkspaces = value;
        }

        [Input("disabledDataSources")]
        private InputList<string>? _disabledDataSources;

        /// <summary>
        /// A list of disabled data sources for the Iot Security Solution. Possible value is `TwinData`.
        /// </summary>
        public InputList<string> DisabledDataSources
        {
            get => _disabledDataSources ?? (_disabledDataSources = new InputList<string>());
            set => _disabledDataSources = value;
        }

        /// <summary>
        /// Specifies the Display Name for this Iot Security Solution.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Is the Iot Security Solution enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventsToExports")]
        private InputList<string>? _eventsToExports;

        /// <summary>
        /// A list of data which is to exported to analytic workspace. Valid values include `RawEvents`.
        /// </summary>
        public InputList<string> EventsToExports
        {
            get => _eventsToExports ?? (_eventsToExports = new InputList<string>());
            set => _eventsToExports = value;
        }

        [Input("iothubIds")]
        private InputList<string>? _iothubIds;

        /// <summary>
        /// Specifies the IoT Hub resource IDs to which this Iot Security Solution is applied.
        /// </summary>
        public InputList<string> IothubIds
        {
            get => _iothubIds ?? (_iothubIds = new InputList<string>());
            set => _iothubIds = value;
        }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the Log Analytics Workspace ID to which the security data will be sent.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        /// <summary>
        /// Should ip addressed be unmasked in the log? Defaults to `false`.
        /// </summary>
        [Input("logUnmaskedIpsEnabled")]
        public Input<bool>? LogUnmaskedIpsEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the Iot Security Solution. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An Azure Resource Graph query used to set the resources monitored.
        /// </summary>
        [Input("queryForResources")]
        public Input<string>? QueryForResources { get; set; }

        [Input("querySubscriptionIds")]
        private InputList<string>? _querySubscriptionIds;

        /// <summary>
        /// A list of subscription Ids on which the user defined resources query should be executed.
        /// </summary>
        public InputList<string> QuerySubscriptionIds
        {
            get => _querySubscriptionIds ?? (_querySubscriptionIds = new InputList<string>());
            set => _querySubscriptionIds = value;
        }

        /// <summary>
        /// A `recommendations_enabled` block of options to enable or disable as defined below.
        /// </summary>
        [Input("recommendationsEnabled")]
        public Input<Inputs.SecuritySolutionRecommendationsEnabledGetArgs>? RecommendationsEnabled { get; set; }

        /// <summary>
        /// Specifies the name of the resource group in which to create the Iot Security Solution. Changing this forces a new resource to be created.
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

        public SecuritySolutionState()
        {
        }
    }
}
