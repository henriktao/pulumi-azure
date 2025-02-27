// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub
{
    public static class GetAuthorizationRule
    {
        /// <summary>
        /// Use this data source to access information about an existing Event Hubs Authorization Rule within an Event Hub.
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
        ///         var test = Output.Create(Azure.EventHub.GetAuthorizationRule.InvokeAsync(new Azure.EventHub.GetAuthorizationRuleArgs
        ///         {
        ///             Name = "test",
        ///             NamespaceName = azurerm_eventhub_namespace.Test.Name,
        ///             EventhubName = azurerm_eventhub.Test.Name,
        ///             ResourceGroupName = azurerm_resource_group.Test.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthorizationRuleResult> InvokeAsync(GetAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationRuleResult>("azure:eventhub/getAuthorizationRule:getAuthorizationRule", args ?? new GetAuthorizationRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Event Hubs Authorization Rule within an Event Hub.
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
        ///         var test = Output.Create(Azure.EventHub.GetAuthorizationRule.InvokeAsync(new Azure.EventHub.GetAuthorizationRuleArgs
        ///         {
        ///             Name = "test",
        ///             NamespaceName = azurerm_eventhub_namespace.Test.Name,
        ///             EventhubName = azurerm_eventhub.Test.Name,
        ///             ResourceGroupName = azurerm_resource_group.Test.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthorizationRuleResult> Invoke(GetAuthorizationRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthorizationRuleResult>("azure:eventhub/getAuthorizationRule:getAuthorizationRule", args ?? new GetAuthorizationRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub.
        /// </summary>
        [Input("eventhubName", required: true)]
        public string EventhubName { get; set; } = null!;

        [Input("listen")]
        public bool? Listen { get; set; }

        [Input("manage")]
        public bool? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. be created.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("send")]
        public bool? Send { get; set; }

        public GetAuthorizationRuleArgs()
        {
        }
    }

    public sealed class GetAuthorizationRuleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the EventHub.
        /// </summary>
        [Input("eventhubName", required: true)]
        public Input<string> EventhubName { get; set; } = null!;

        [Input("listen")]
        public Input<bool>? Listen { get; set; }

        [Input("manage")]
        public Input<bool>? Manage { get; set; }

        /// <summary>
        /// Specifies the name of the EventHub Authorization Rule resource. be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the grandparent EventHub Namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("send")]
        public Input<bool>? Send { get; set; }

        public GetAuthorizationRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorizationRuleResult
    {
        public readonly string EventhubName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? Listen;
        public readonly string Location;
        public readonly bool? Manage;
        public readonly string Name;
        public readonly string NamespaceName;
        /// <summary>
        /// The Primary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The alias of the Primary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string PrimaryConnectionStringAlias;
        /// <summary>
        /// The Primary Key for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The alias of the Secondary Connection String for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string SecondaryConnectionStringAlias;
        /// <summary>
        /// The Secondary Key for the Event Hubs Authorization Rule.
        /// </summary>
        public readonly string SecondaryKey;
        public readonly bool? Send;

        [OutputConstructor]
        private GetAuthorizationRuleResult(
            string eventhubName,

            string id,

            bool? listen,

            string location,

            bool? manage,

            string name,

            string namespaceName,

            string primaryConnectionString,

            string primaryConnectionStringAlias,

            string primaryKey,

            string resourceGroupName,

            string secondaryConnectionString,

            string secondaryConnectionStringAlias,

            string secondaryKey,

            bool? send)
        {
            EventhubName = eventhubName;
            Id = id;
            Listen = listen;
            Location = location;
            Manage = manage;
            Name = name;
            NamespaceName = namespaceName;
            PrimaryConnectionString = primaryConnectionString;
            PrimaryConnectionStringAlias = primaryConnectionStringAlias;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryConnectionStringAlias = secondaryConnectionStringAlias;
            SecondaryKey = secondaryKey;
            Send = send;
        }
    }
}
