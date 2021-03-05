// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages a Logger within an API Management Service.
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
    ///         var exampleInsights = new Azure.AppInsights.Insights("exampleInsights", new Azure.AppInsights.InsightsArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ApplicationType = "other",
    ///         });
    ///         var exampleService = new Azure.ApiManagement.Service("exampleService", new Azure.ApiManagement.ServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PublisherName = "My Company",
    ///             PublisherEmail = "company@exmaple.com",
    ///             SkuName = "Developer_1",
    ///         });
    ///         var exampleLogger = new Azure.ApiManagement.Logger("exampleLogger", new Azure.ApiManagement.LoggerArgs
    ///         {
    ///             ApiManagementName = exampleService.Name,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             ResourceId = exampleInsights.Id,
    ///             ApplicationInsights = new Azure.ApiManagement.Inputs.LoggerApplicationInsightsArgs
    ///             {
    ///                 InstrumentationKey = exampleInsights.InstrumentationKey,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management Loggers can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/logger:Logger example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/example-rg/Microsoft.ApiManagement/service/example-apim/loggers/example-logger
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/logger:Logger")]
    public partial class Logger : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiManagementName")]
        public Output<string> ApiManagementName { get; private set; } = null!;

        /// <summary>
        /// An `application_insights` block as documented below.
        /// </summary>
        [Output("applicationInsights")]
        public Output<Outputs.LoggerApplicationInsights?> ApplicationInsights { get; private set; } = null!;

        /// <summary>
        /// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
        /// </summary>
        [Output("buffered")]
        public Output<bool?> Buffered { get; private set; } = null!;

        /// <summary>
        /// A description of this Logger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// An `eventhub` block as documented below.
        /// </summary>
        [Output("eventhub")]
        public Output<Outputs.LoggerEventhub?> Eventhub { get; private set; } = null!;

        /// <summary>
        /// The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The target resource id which will be linked in the API-Management portal page.
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a Logger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Logger(string name, LoggerArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/logger:Logger", name, args ?? new LoggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Logger(string name, Input<string> id, LoggerState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/logger:Logger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Logger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Logger Get(string name, Input<string> id, LoggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Logger(name, id, state, options);
        }
    }

    public sealed class LoggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName", required: true)]
        public Input<string> ApiManagementName { get; set; } = null!;

        /// <summary>
        /// An `application_insights` block as documented below.
        /// </summary>
        [Input("applicationInsights")]
        public Input<Inputs.LoggerApplicationInsightsArgs>? ApplicationInsights { get; set; }

        /// <summary>
        /// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
        /// </summary>
        [Input("buffered")]
        public Input<bool>? Buffered { get; set; }

        /// <summary>
        /// A description of this Logger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An `eventhub` block as documented below.
        /// </summary>
        [Input("eventhub")]
        public Input<Inputs.LoggerEventhubArgs>? Eventhub { get; set; }

        /// <summary>
        /// The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The target resource id which will be linked in the API-Management portal page.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public LoggerArgs()
        {
        }
    }

    public sealed class LoggerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiManagementName")]
        public Input<string>? ApiManagementName { get; set; }

        /// <summary>
        /// An `application_insights` block as documented below.
        /// </summary>
        [Input("applicationInsights")]
        public Input<Inputs.LoggerApplicationInsightsGetArgs>? ApplicationInsights { get; set; }

        /// <summary>
        /// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to `true`.
        /// </summary>
        [Input("buffered")]
        public Input<bool>? Buffered { get; set; }

        /// <summary>
        /// A description of this Logger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An `eventhub` block as documented below.
        /// </summary>
        [Input("eventhub")]
        public Input<Inputs.LoggerEventhubGetArgs>? Eventhub { get; set; }

        /// <summary>
        /// The name of this Logger, which must be unique within the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The target resource id which will be linked in the API-Management portal page.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public LoggerState()
        {
        }
    }
}
