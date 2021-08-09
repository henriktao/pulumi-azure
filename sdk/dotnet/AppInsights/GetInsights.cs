// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppInsights
{
    public static class GetInsights
    {
        /// <summary>
        /// Use this data source to access information about an existing Application Insights component.
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
        ///         var example = Output.Create(Azure.AppInsights.GetInsights.InvokeAsync(new Azure.AppInsights.GetInsightsArgs
        ///         {
        ///             Name = "production",
        ///             ResourceGroupName = "networking",
        ///         }));
        ///         this.ApplicationInsightsInstrumentationKey = example.Apply(example =&gt; example.InstrumentationKey);
        ///     }
        /// 
        ///     [Output("applicationInsightsInstrumentationKey")]
        ///     public Output&lt;string&gt; ApplicationInsightsInstrumentationKey { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInsightsResult> InvokeAsync(GetInsightsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInsightsResult>("azure:appinsights/getInsights:getInsights", args ?? new GetInsightsArgs(), options.WithVersion());
    }


    public sealed class GetInsightsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Application Insights component.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group the Application Insights component is located in.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetInsightsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInsightsResult
    {
        /// <summary>
        /// The App ID associated with this Application Insights component.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The type of the component.
        /// </summary>
        public readonly string ApplicationType;
        /// <summary>
        /// The connection string of the Application Insights component. (Sensitive)
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instrumentation key of the Application Insights component.
        /// </summary>
        public readonly string InstrumentationKey;
        /// <summary>
        /// The Azure location where the component exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The retention period in days.
        /// </summary>
        public readonly int RetentionInDays;
        /// <summary>
        /// Tags applied to the component.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The id of the associated Log Analytics workspace
        /// </summary>
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetInsightsResult(
            string appId,

            string applicationType,

            string connectionString,

            string id,

            string instrumentationKey,

            string location,

            string name,

            string resourceGroupName,

            int retentionInDays,

            ImmutableDictionary<string, string> tags,

            string workspaceId)
        {
            AppId = appId;
            ApplicationType = applicationType;
            ConnectionString = connectionString;
            Id = id;
            InstrumentationKey = instrumentationKey;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            RetentionInDays = retentionInDays;
            Tags = tags;
            WorkspaceId = workspaceId;
        }
    }
}
