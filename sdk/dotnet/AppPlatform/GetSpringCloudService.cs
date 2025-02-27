// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform
{
    public static class GetSpringCloudService
    {
        /// <summary>
        /// Use this data source to access information about an existing Spring Cloud Service.
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
        ///         var example = Output.Create(Azure.AppPlatform.GetSpringCloudService.InvokeAsync(new Azure.AppPlatform.GetSpringCloudServiceArgs
        ///         {
        ///             Name = azurerm_spring_cloud_service.Example.Name,
        ///             ResourceGroupName = azurerm_spring_cloud_service.Example.Resource_group_name,
        ///         }));
        ///         this.SpringCloudServiceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("springCloudServiceId")]
        ///     public Output&lt;string&gt; SpringCloudServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSpringCloudServiceResult> InvokeAsync(GetSpringCloudServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpringCloudServiceResult>("azure:appplatform/getSpringCloudService:getSpringCloudService", args ?? new GetSpringCloudServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Spring Cloud Service.
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
        ///         var example = Output.Create(Azure.AppPlatform.GetSpringCloudService.InvokeAsync(new Azure.AppPlatform.GetSpringCloudServiceArgs
        ///         {
        ///             Name = azurerm_spring_cloud_service.Example.Name,
        ///             ResourceGroupName = azurerm_spring_cloud_service.Example.Resource_group_name,
        ///         }));
        ///         this.SpringCloudServiceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("springCloudServiceId")]
        ///     public Output&lt;string&gt; SpringCloudServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSpringCloudServiceResult> Invoke(GetSpringCloudServiceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSpringCloudServiceResult>("azure:appplatform/getSpringCloudService:getSpringCloudService", args ?? new GetSpringCloudServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSpringCloudServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies The name of the Spring Cloud Service resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Spring Cloud Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSpringCloudServiceArgs()
        {
        }
    }

    public sealed class GetSpringCloudServiceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies The name of the Spring Cloud Service resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Spring Cloud Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSpringCloudServiceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSpringCloudServiceResult
    {
        /// <summary>
        /// A `config_server_git_setting` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSpringCloudServiceConfigServerGitSettingResult> ConfigServerGitSettings;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location of Spring Cloud Service.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name to identify on the Git repository.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of the outbound Public IP Addresses used by this Spring Cloud Service.
        /// </summary>
        public readonly ImmutableArray<string> OutboundPublicIpAddresses;
        /// <summary>
        /// A list of `required_network_traffic_rules` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSpringCloudServiceRequiredNetworkTrafficRuleResult> RequiredNetworkTrafficRules;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to Spring Cloud Service.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetSpringCloudServiceResult(
            ImmutableArray<Outputs.GetSpringCloudServiceConfigServerGitSettingResult> configServerGitSettings,

            string id,

            string location,

            string name,

            ImmutableArray<string> outboundPublicIpAddresses,

            ImmutableArray<Outputs.GetSpringCloudServiceRequiredNetworkTrafficRuleResult> requiredNetworkTrafficRules,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            ConfigServerGitSettings = configServerGitSettings;
            Id = id;
            Location = location;
            Name = name;
            OutboundPublicIpAddresses = outboundPublicIpAddresses;
            RequiredNetworkTrafficRules = requiredNetworkTrafficRules;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
