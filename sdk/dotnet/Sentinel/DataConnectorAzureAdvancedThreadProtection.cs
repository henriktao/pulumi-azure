// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel
{
    /// <summary>
    /// Manages a Azure Advanced Threat Protection Data Connector.
    /// 
    /// !&gt; **NOTE:** This resource requires that [Enterprise Mobility + Security E5](https://www.microsoft.com/en-us/microsoft-365/enterprise-mobility-security) is enabled on the tenant being connected to.
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
    ///             Location = "west europe",
    ///         });
    ///         var exampleAnalyticsWorkspace = new Azure.OperationalInsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", new Azure.OperationalInsights.AnalyticsWorkspaceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = "PerGB2018",
    ///         });
    ///         var exampleAnalyticsSolution = new Azure.OperationalInsights.AnalyticsSolution("exampleAnalyticsSolution", new Azure.OperationalInsights.AnalyticsSolutionArgs
    ///         {
    ///             SolutionName = "SecurityInsights",
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             WorkspaceResourceId = exampleAnalyticsWorkspace.Id,
    ///             WorkspaceName = exampleAnalyticsWorkspace.Name,
    ///             Plan = new Azure.OperationalInsights.Inputs.AnalyticsSolutionPlanArgs
    ///             {
    ///                 Publisher = "Microsoft",
    ///                 Product = "OMSGallery/SecurityInsights",
    ///             },
    ///         });
    ///         var exampleDataConnectorAzureAdvancedThreadProtection = new Azure.Sentinel.DataConnectorAzureAdvancedThreadProtection("exampleDataConnectorAzureAdvancedThreadProtection", new Azure.Sentinel.DataConnectorAzureAdvancedThreadProtectionArgs
    ///         {
    ///             LogAnalyticsWorkspaceId = exampleAnalyticsSolution.WorkspaceResourceId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Azure Advanced Threat Protection Data Connectors can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:sentinel/dataConnectorAzureAdvancedThreadProtection:DataConnectorAzureAdvancedThreadProtection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
    /// ```
    /// </summary>
    [AzureResourceType("azure:sentinel/dataConnectorAzureAdvancedThreadProtection:DataConnectorAzureAdvancedThreadProtection")]
    public partial class DataConnectorAzureAdvancedThreadProtection : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Log Analytics Workspace that this Azure Advanced Threat Protection Data Connector resides in. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Output("logAnalyticsWorkspaceId")]
        public Output<string> LogAnalyticsWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The name which should be used for this Azure Advanced Threat Protection Data Connector. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the tenant that this Azure Advanced Threat Protection Data Connector connects to. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a DataConnectorAzureAdvancedThreadProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataConnectorAzureAdvancedThreadProtection(string name, DataConnectorAzureAdvancedThreadProtectionArgs args, CustomResourceOptions? options = null)
            : base("azure:sentinel/dataConnectorAzureAdvancedThreadProtection:DataConnectorAzureAdvancedThreadProtection", name, args ?? new DataConnectorAzureAdvancedThreadProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataConnectorAzureAdvancedThreadProtection(string name, Input<string> id, DataConnectorAzureAdvancedThreadProtectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:sentinel/dataConnectorAzureAdvancedThreadProtection:DataConnectorAzureAdvancedThreadProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataConnectorAzureAdvancedThreadProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataConnectorAzureAdvancedThreadProtection Get(string name, Input<string> id, DataConnectorAzureAdvancedThreadProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new DataConnectorAzureAdvancedThreadProtection(name, id, state, options);
        }
    }

    public sealed class DataConnectorAzureAdvancedThreadProtectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Log Analytics Workspace that this Azure Advanced Threat Protection Data Connector resides in. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId", required: true)]
        public Input<string> LogAnalyticsWorkspaceId { get; set; } = null!;

        /// <summary>
        /// The name which should be used for this Azure Advanced Threat Protection Data Connector. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the tenant that this Azure Advanced Threat Protection Data Connector connects to. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public DataConnectorAzureAdvancedThreadProtectionArgs()
        {
        }
    }

    public sealed class DataConnectorAzureAdvancedThreadProtectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Log Analytics Workspace that this Azure Advanced Threat Protection Data Connector resides in. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        /// <summary>
        /// The name which should be used for this Azure Advanced Threat Protection Data Connector. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the tenant that this Azure Advanced Threat Protection Data Connector connects to. Changing this forces a new Azure Advanced Threat Protection Data Connector to be created.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public DataConnectorAzureAdvancedThreadProtectionState()
        {
        }
    }
}
