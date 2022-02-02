// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetFunctionAppHostKeys
    {
        /// <summary>
        /// Use this data source to fetch the Host Keys of an existing Function App
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
        ///         var example = Output.Create(Azure.AppService.GetFunctionAppHostKeys.InvokeAsync(new Azure.AppService.GetFunctionAppHostKeysArgs
        ///         {
        ///             Name = "example-function",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionAppHostKeysResult> InvokeAsync(GetFunctionAppHostKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionAppHostKeysResult>("azure:appservice/getFunctionAppHostKeys:getFunctionAppHostKeys", args ?? new GetFunctionAppHostKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to fetch the Host Keys of an existing Function App
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
        ///         var example = Output.Create(Azure.AppService.GetFunctionAppHostKeys.InvokeAsync(new Azure.AppService.GetFunctionAppHostKeysArgs
        ///         {
        ///             Name = "example-function",
        ///             ResourceGroupName = azurerm_resource_group.Example.Name,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFunctionAppHostKeysResult> Invoke(GetFunctionAppHostKeysInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionAppHostKeysResult>("azure:appservice/getFunctionAppHostKeys:getFunctionAppHostKeys", args ?? new GetFunctionAppHostKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionAppHostKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Function App.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Function App exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFunctionAppHostKeysArgs()
        {
        }
    }

    public sealed class GetFunctionAppHostKeysInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Function App.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Function App exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFunctionAppHostKeysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionAppHostKeysResult
    {
        /// <summary>
        /// Function App resource's default function key.
        /// </summary>
        public readonly string DefaultFunctionKey;
        /// <summary>
        /// Function App resource's Durable Task Extension system key.
        /// </summary>
        public readonly string DurabletaskExtensionKey;
        /// <summary>
        /// Function App resource's Event Grid Extension Config system key.
        /// </summary>
        public readonly string EventGridExtensionConfigKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Function App resource's secret key
        /// </summary>
        public readonly string MasterKey;
        public readonly string Name;
        public readonly string PrimaryKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// Function App resource's SignalR Extension system key.
        /// </summary>
        public readonly string SignalrExtensionKey;

        [OutputConstructor]
        private GetFunctionAppHostKeysResult(
            string defaultFunctionKey,

            string durabletaskExtensionKey,

            string eventGridExtensionConfigKey,

            string id,

            string masterKey,

            string name,

            string primaryKey,

            string resourceGroupName,

            string signalrExtensionKey)
        {
            DefaultFunctionKey = defaultFunctionKey;
            DurabletaskExtensionKey = durabletaskExtensionKey;
            EventGridExtensionConfigKey = eventGridExtensionConfigKey;
            Id = id;
            MasterKey = masterKey;
            Name = name;
            PrimaryKey = primaryKey;
            ResourceGroupName = resourceGroupName;
            SignalrExtensionKey = signalrExtensionKey;
        }
    }
}
