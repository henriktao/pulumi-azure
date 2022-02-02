// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Maintenance
{
    public static class GetConfiguration
    {
        /// <summary>
        /// Use this data source to access information about an existing Maintenance Configuration.
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
        ///         var existing = Output.Create(Azure.Maintenance.GetConfiguration.InvokeAsync(new Azure.Maintenance.GetConfigurationArgs
        ///         {
        ///             Name = "example-mc",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.Id = azurerm_maintenance_configuration.Existing.Id;
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConfigurationResult> InvokeAsync(GetConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationResult>("azure:maintenance/getConfiguration:getConfiguration", args ?? new GetConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Maintenance Configuration.
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
        ///         var existing = Output.Create(Azure.Maintenance.GetConfiguration.InvokeAsync(new Azure.Maintenance.GetConfigurationArgs
        ///         {
        ///             Name = "example-mc",
        ///             ResourceGroupName = "example-resources",
        ///         }));
        ///         this.Id = azurerm_maintenance_configuration.Existing.Id;
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("azure:maintenance/getConfiguration:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Maintenance Configuration.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Maintenance Configuration exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConfigurationArgs()
        {
        }
    }

    public sealed class GetConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Maintenance Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where this Maintenance Configuration exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The properties assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The scope of the Maintenance Configuration.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The visibility of the Maintenance Configuration.
        /// </summary>
        public readonly string Visibility;
        /// <summary>
        /// A `window` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationWindowResult> Windows;

        [OutputConstructor]
        private GetConfigurationResult(
            string id,

            string location,

            string name,

            ImmutableDictionary<string, string> properties,

            string resourceGroupName,

            string scope,

            ImmutableDictionary<string, string> tags,

            string visibility,

            ImmutableArray<Outputs.GetConfigurationWindowResult> windows)
        {
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            ResourceGroupName = resourceGroupName;
            Scope = scope;
            Tags = tags;
            Visibility = visibility;
            Windows = windows;
        }
    }
}
