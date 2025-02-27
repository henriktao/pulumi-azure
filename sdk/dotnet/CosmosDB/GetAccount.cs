// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB
{
    public static class GetAccount
    {
        /// <summary>
        /// Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
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
        ///         var example = Output.Create(Azure.CosmosDB.GetAccount.InvokeAsync(new Azure.CosmosDB.GetAccountArgs
        ///         {
        ///             Name = "tfex-cosmosdb-account",
        ///             ResourceGroupName = "tfex-cosmosdb-account-rg",
        ///         }));
        ///         this.CosmosdbAccountEndpoint = example.Apply(example =&gt; example.Endpoint);
        ///     }
        /// 
        ///     [Output("cosmosdbAccountEndpoint")]
        ///     public Output&lt;string&gt; CosmosdbAccountEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:cosmosdb/getAccount:getAccount", args ?? new GetAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.
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
        ///         var example = Output.Create(Azure.CosmosDB.GetAccount.InvokeAsync(new Azure.CosmosDB.GetAccountArgs
        ///         {
        ///             Name = "tfex-cosmosdb-account",
        ///             ResourceGroupName = "tfex-cosmosdb-account-rg",
        ///         }));
        ///         this.CosmosdbAccountEndpoint = example.Apply(example =&gt; example.Endpoint);
        ///     }
        /// 
        ///     [Output("cosmosdbAccountEndpoint")]
        ///     public Output&lt;string&gt; CosmosdbAccountEndpoint { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccountResult>("azure:cosmosdb/getAccount:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the CosmosDB Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group in which the CosmosDB Account resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
    }

    public sealed class GetAccountInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the CosmosDB Account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the resource group in which the CosmosDB Account resides.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// Capabilities enabled on this Cosmos DB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountCapabilityResult> Capabilities;
        public readonly ImmutableArray<Outputs.GetAccountConsistencyPolicyResult> ConsistencyPolicies;
        /// <summary>
        /// If automatic failover is enabled for this CosmosDB Account.
        /// </summary>
        public readonly bool EnableAutomaticFailover;
        /// <summary>
        /// If Free Tier pricing option is enabled for this CosmosDB Account.
        /// </summary>
        public readonly bool EnableFreeTier;
        /// <summary>
        /// If multiple write locations are enabled for this Cosmos DB account.
        /// </summary>
        public readonly bool EnableMultipleWriteLocations;
        /// <summary>
        /// The endpoint used to connect to the CosmosDB account.
        /// </summary>
        public readonly string Endpoint;
        public readonly ImmutableArray<Outputs.GetAccountGeoLocationResult> GeoLocations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The current IP Filter for this CosmosDB account
        /// </summary>
        public readonly string IpRangeFilter;
        /// <summary>
        /// If virtual network filtering is enabled for this Cosmos DB account.
        /// </summary>
        public readonly bool IsVirtualNetworkFilterEnabled;
        /// <summary>
        /// The Key Vault key URI for CMK encryption.
        /// </summary>
        public readonly string KeyVaultKeyId;
        /// <summary>
        /// The Kind of the CosmosDB account.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the Azure region hosting replicated data.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// The Offer Type to used by this CosmosDB Account.
        /// </summary>
        public readonly string OfferType;
        /// <summary>
        /// The Primary key for the CosmosDB Account.
        /// </summary>
        public readonly string PrimaryKey;
        public readonly string PrimaryMasterKey;
        /// <summary>
        /// The Primary read-only Key for the CosmosDB Account.
        /// </summary>
        public readonly string PrimaryReadonlyKey;
        public readonly string PrimaryReadonlyMasterKey;
        /// <summary>
        /// A list of read endpoints available for this CosmosDB account.
        /// </summary>
        public readonly ImmutableArray<string> ReadEndpoints;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Secondary key for the CosmosDB Account.
        /// </summary>
        public readonly string SecondaryKey;
        public readonly string SecondaryMasterKey;
        /// <summary>
        /// The Secondary read-only key for the CosmosDB Account.
        /// </summary>
        public readonly string SecondaryReadonlyKey;
        public readonly string SecondaryReadonlyMasterKey;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Subnets that are allowed to access this CosmosDB account.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountVirtualNetworkRuleResult> VirtualNetworkRules;
        /// <summary>
        /// A list of write endpoints available for this CosmosDB account.
        /// </summary>
        public readonly ImmutableArray<string> WriteEndpoints;

        [OutputConstructor]
        private GetAccountResult(
            ImmutableArray<Outputs.GetAccountCapabilityResult> capabilities,

            ImmutableArray<Outputs.GetAccountConsistencyPolicyResult> consistencyPolicies,

            bool enableAutomaticFailover,

            bool enableFreeTier,

            bool enableMultipleWriteLocations,

            string endpoint,

            ImmutableArray<Outputs.GetAccountGeoLocationResult> geoLocations,

            string id,

            string ipRangeFilter,

            bool isVirtualNetworkFilterEnabled,

            string keyVaultKeyId,

            string kind,

            string location,

            string name,

            string offerType,

            string primaryKey,

            string primaryMasterKey,

            string primaryReadonlyKey,

            string primaryReadonlyMasterKey,

            ImmutableArray<string> readEndpoints,

            string resourceGroupName,

            string secondaryKey,

            string secondaryMasterKey,

            string secondaryReadonlyKey,

            string secondaryReadonlyMasterKey,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetAccountVirtualNetworkRuleResult> virtualNetworkRules,

            ImmutableArray<string> writeEndpoints)
        {
            Capabilities = capabilities;
            ConsistencyPolicies = consistencyPolicies;
            EnableAutomaticFailover = enableAutomaticFailover;
            EnableFreeTier = enableFreeTier;
            EnableMultipleWriteLocations = enableMultipleWriteLocations;
            Endpoint = endpoint;
            GeoLocations = geoLocations;
            Id = id;
            IpRangeFilter = ipRangeFilter;
            IsVirtualNetworkFilterEnabled = isVirtualNetworkFilterEnabled;
            KeyVaultKeyId = keyVaultKeyId;
            Kind = kind;
            Location = location;
            Name = name;
            OfferType = offerType;
            PrimaryKey = primaryKey;
            PrimaryMasterKey = primaryMasterKey;
            PrimaryReadonlyKey = primaryReadonlyKey;
            PrimaryReadonlyMasterKey = primaryReadonlyMasterKey;
            ReadEndpoints = readEndpoints;
            ResourceGroupName = resourceGroupName;
            SecondaryKey = secondaryKey;
            SecondaryMasterKey = secondaryMasterKey;
            SecondaryReadonlyKey = secondaryReadonlyKey;
            SecondaryReadonlyMasterKey = secondaryReadonlyMasterKey;
            Tags = tags;
            VirtualNetworkRules = virtualNetworkRules;
            WriteEndpoints = writeEndpoints;
        }
    }
}
