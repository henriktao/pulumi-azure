// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring
{
    /// <summary>
    /// Manages a logz Monitor.
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
    ///         var exampleLogzMonitor = new Azure.Monitoring.LogzMonitor("exampleLogzMonitor", new Azure.Monitoring.LogzMonitorArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Plan = new Azure.Monitoring.Inputs.LogzMonitorPlanArgs
    ///             {
    ///                 BillingCycle = "Monthly",
    ///                 EffectiveDate = "2022-06-06T00:00:00Z",
    ///                 PlanId = "100gb14days",
    ///                 UsageType = "Committed",
    ///             },
    ///             User = new Azure.Monitoring.Inputs.LogzMonitorUserArgs
    ///             {
    ///                 Email = "user@example.com",
    ///                 FirstName = "Example",
    ///                 LastName = "User",
    ///                 PhoneNumber = "+12313803556",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// logz Monitors can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:monitoring/logzMonitor:LogzMonitor example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1
    /// ```
    /// </summary>
    [AzureResourceType("azure:monitoring/logzMonitor:LogzMonitor")]
    public partial class LogzMonitor : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Output("companyName")]
        public Output<string?> CompanyName { get; private set; } = null!;

        /// <summary>
        /// Whether the resource monitoring is enabled?
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Output("enterpriseAppId")]
        public Output<string?> EnterpriseAppId { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID associated with the logz organization of this logz Monitor.
        /// </summary>
        [Output("logzOrganizationId")]
        public Output<string> LogzOrganizationId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.LogzMonitorPlan> Plan { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The single sign on url associated with the logz organization of this logz Monitor.
        /// </summary>
        [Output("singleSignOnUrl")]
        public Output<string> SingleSignOnUrl { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags which should be assigned to the logz Monitor.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A `user` block as defined below.
        /// </summary>
        [Output("user")]
        public Output<Outputs.LogzMonitorUser> User { get; private set; } = null!;


        /// <summary>
        /// Create a LogzMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogzMonitor(string name, LogzMonitorArgs args, CustomResourceOptions? options = null)
            : base("azure:monitoring/logzMonitor:LogzMonitor", name, args ?? new LogzMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogzMonitor(string name, Input<string> id, LogzMonitorState? state = null, CustomResourceOptions? options = null)
            : base("azure:monitoring/logzMonitor:LogzMonitor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogzMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogzMonitor Get(string name, Input<string> id, LogzMonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new LogzMonitor(name, id, state, options);
        }
    }

    public sealed class LogzMonitorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("companyName")]
        public Input<string>? CompanyName { get; set; }

        /// <summary>
        /// Whether the resource monitoring is enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("enterpriseAppId")]
        public Input<string>? EnterpriseAppId { get; set; }

        /// <summary>
        /// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Input("plan", required: true)]
        public Input<Inputs.LogzMonitorPlanArgs> Plan { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the logz Monitor.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `user` block as defined below.
        /// </summary>
        [Input("user", required: true)]
        public Input<Inputs.LogzMonitorUserArgs> User { get; set; } = null!;

        public LogzMonitorArgs()
        {
        }
    }

    public sealed class LogzMonitorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("companyName")]
        public Input<string>? CompanyName { get; set; }

        /// <summary>
        /// Whether the resource monitoring is enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("enterpriseAppId")]
        public Input<string>? EnterpriseAppId { get; set; }

        /// <summary>
        /// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID associated with the logz organization of this logz Monitor.
        /// </summary>
        [Input("logzOrganizationId")]
        public Input<string>? LogzOrganizationId { get; set; }

        /// <summary>
        /// The name which should be used for this logz Monitor. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A `plan` block as defined below.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.LogzMonitorPlanGetArgs>? Plan { get; set; }

        /// <summary>
        /// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The single sign on url associated with the logz organization of this logz Monitor.
        /// </summary>
        [Input("singleSignOnUrl")]
        public Input<string>? SingleSignOnUrl { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the logz Monitor.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `user` block as defined below.
        /// </summary>
        [Input("user")]
        public Input<Inputs.LogzMonitorUserGetArgs>? User { get; set; }

        public LogzMonitorState()
        {
        }
    }
}
