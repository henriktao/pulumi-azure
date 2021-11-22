// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Azure.DataBricks
{
    public static class GetWorkspace
    {
        /// <summary>
        /// Use this data source to access information about an existing Databricks workspace.
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
        ///         var example = Output.Create(Azure.DataBricks.GetWorkspace.InvokeAsync(new Azure.DataBricks.GetWorkspaceArgs
        ///         {
        ///             Name = "example-workspace",
        ///             ResourceGroupName = "example-rg",
        ///         }));
        ///         this.DatabricksWorkspaceId = example.Apply(example =&gt; example.WorkspaceId);
        ///     }
        /// 
        ///     [Output("databricksWorkspaceId")]
        ///     public Output&lt;string&gt; DatabricksWorkspaceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azure:databricks/getWorkspace:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to access information about an existing Databricks workspace.
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
        ///         var example = Output.Create(Azure.DataBricks.GetWorkspace.InvokeAsync(new Azure.DataBricks.GetWorkspaceArgs
        ///         {
        ///             Name = "example-workspace",
        ///             ResourceGroupName = "example-rg",
        ///         }));
        ///         this.DatabricksWorkspaceId = example.Apply(example =&gt; example.WorkspaceId);
        ///     }
        /// 
        ///     [Output("databricksWorkspaceId")]
        ///     public Output&lt;string&gt; DatabricksWorkspaceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("azure:databricks/getWorkspace:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Databricks Workspace.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the Databricks Workspace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Databricks Workspace.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetWorkspaceArgs()
        {
        }
    }

    public sealed class GetWorkspaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Databricks Workspace.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the Databricks Workspace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Databricks Workspace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetWorkspaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// SKU of this Databricks Workspace.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// A mapping of tags to assign to the Databricks Workspace.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Unique ID of this Databricks Workspace in Databricks management plane.
        /// </summary>
        public readonly string WorkspaceId;
        /// <summary>
        /// URL this Databricks Workspace is accessible on.
        /// </summary>
        public readonly string WorkspaceUrl;

        [OutputConstructor]
        private GetWorkspaceResult(
            string id,

            string name,

            string resourceGroupName,

            string sku,

            ImmutableDictionary<string, string>? tags,

            string workspaceId,

            string workspaceUrl)
        {
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            Tags = tags;
            WorkspaceId = workspaceId;
            WorkspaceUrl = workspaceUrl;
        }
    }
}
