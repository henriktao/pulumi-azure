// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService
{
    public static class GetRegistry
    {
        /// <summary>
        /// Use this data source to access information about an existing Container Registry.
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
        ///         var example = Output.Create(Azure.ContainerService.GetRegistry.InvokeAsync(new Azure.ContainerService.GetRegistryArgs
        ///         {
        ///             Name = "testacr",
        ///             ResourceGroupName = "test",
        ///         }));
        ///         this.LoginServer = example.Apply(example =&gt; example.LoginServer);
        ///     }
        /// 
        ///     [Output("loginServer")]
        ///     public Output&lt;string&gt; LoginServer { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("azure:containerservice/getRegistry:getRegistry", args ?? new GetRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Container Registry.
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
        ///         var example = Output.Create(Azure.ContainerService.GetRegistry.InvokeAsync(new Azure.ContainerService.GetRegistryArgs
        ///         {
        ///             Name = "testacr",
        ///             ResourceGroupName = "test",
        ///         }));
        ///         this.LoginServer = example.Apply(example =&gt; example.LoginServer);
        ///     }
        /// 
        ///     [Output("loginServer")]
        ///     public Output&lt;string&gt; LoginServer { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegistryResult> Invoke(GetRegistryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryResult>("azure:containerservice/getRegistry:getRegistry", args ?? new GetRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Container Registry.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Container Registry exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
    }

    public sealed class GetRegistryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Container Registry.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Container Registry exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetRegistryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// Is the Administrator account enabled for this Container Registry.
        /// </summary>
        public readonly bool AdminEnabled;
        /// <summary>
        /// The Password associated with the Container Registry Admin account - if the admin account is enabled.
        /// </summary>
        public readonly string AdminPassword;
        /// <summary>
        /// The Username associated with the Container Registry Admin account - if the admin account is enabled.
        /// </summary>
        public readonly string AdminUsername;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which this Container Registry exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The URL that can be used to log into the container registry.
        /// </summary>
        public readonly string LoginServer;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SKU of this Container Registry, such as `Basic`.
        /// </summary>
        public readonly string Sku;
        public readonly string StorageAccountId;
        /// <summary>
        /// A map of tags assigned to the Container Registry.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetRegistryResult(
            bool adminEnabled,

            string adminPassword,

            string adminUsername,

            string id,

            string location,

            string loginServer,

            string name,

            string resourceGroupName,

            string sku,

            string storageAccountId,

            ImmutableDictionary<string, string> tags)
        {
            AdminEnabled = adminEnabled;
            AdminPassword = adminPassword;
            AdminUsername = adminUsername;
            Id = id;
            Location = location;
            LoginServer = loginServer;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Sku = sku;
            StorageAccountId = storageAccountId;
            Tags = tags;
        }
    }
}
