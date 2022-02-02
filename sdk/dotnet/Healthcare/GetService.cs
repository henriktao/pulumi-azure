// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Healthcare
{
    public static class GetService
    {
        /// <summary>
        /// Use this data source to access information about an existing Healthcare Service
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
        ///         var example = Output.Create(Azure.Healthcare.GetService.InvokeAsync(new Azure.Healthcare.GetServiceArgs
        ///         {
        ///             Name = "example-healthcare_service",
        ///             ResourceGroupName = "example-resources",
        ///             Location = "westus2",
        ///         }));
        ///         this.HealthcareServiceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("healthcareServiceId")]
        ///     public Output&lt;string&gt; HealthcareServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azure:healthcare/getService:getService", args ?? new GetServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing Healthcare Service
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
        ///         var example = Output.Create(Azure.Healthcare.GetService.InvokeAsync(new Azure.Healthcare.GetServiceArgs
        ///         {
        ///             Name = "example-healthcare_service",
        ///             ResourceGroupName = "example-resources",
        ///             Location = "westus2",
        ///         }));
        ///         this.HealthcareServiceId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("healthcareServiceId")]
        ///     public Output&lt;string&gt; HealthcareServiceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceResult> Invoke(GetServiceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceResult>("azure:healthcare/getService:getService", args ?? new GetServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Region where the Service is located.
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Healthcare Service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Healthcare Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }

    public sealed class GetServiceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure Region where the Service is located.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Healthcare Service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Healthcare Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetServiceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        public readonly ImmutableArray<string> AccessPolicyObjectIds;
        /// <summary>
        /// An `authentication_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceAuthenticationConfigurationResult> AuthenticationConfigurations;
        /// <summary>
        /// A `cors_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceCorsConfigurationResult> CorsConfigurations;
        /// <summary>
        /// The versionless Key Vault Key ID for CMK encryption of the backing database.
        /// </summary>
        public readonly string CosmosdbKeyVaultKeyVersionlessId;
        /// <summary>
        /// The provisioned throughput for the backing database.
        /// </summary>
        public readonly int CosmosdbThroughput;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The type of the service.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The Azure Region where the Service is located.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetServiceResult(
            ImmutableArray<string> accessPolicyObjectIds,

            ImmutableArray<Outputs.GetServiceAuthenticationConfigurationResult> authenticationConfigurations,

            ImmutableArray<Outputs.GetServiceCorsConfigurationResult> corsConfigurations,

            string cosmosdbKeyVaultKeyVersionlessId,

            int cosmosdbThroughput,

            string id,

            string kind,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            AccessPolicyObjectIds = accessPolicyObjectIds;
            AuthenticationConfigurations = authenticationConfigurations;
            CorsConfigurations = corsConfigurations;
            CosmosdbKeyVaultKeyVersionlessId = cosmosdbKeyVaultKeyVersionlessId;
            CosmosdbThroughput = cosmosdbThroughput;
            Id = id;
            Kind = kind;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
