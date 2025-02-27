// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB
{
    /// <summary>
    /// Manages a CosmosDB (formally DocumentDB) Account.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using Random = Pulumi.Random;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rg = new Azure.Core.ResourceGroup("rg", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "westus",
    ///         });
    ///         var ri = new Random.RandomInteger("ri", new Random.RandomIntegerArgs
    ///         {
    ///             Min = 10000,
    ///             Max = 99999,
    ///         });
    ///         var db = new Azure.CosmosDB.Account("db", new Azure.CosmosDB.AccountArgs
    ///         {
    ///             Location = rg.Location,
    ///             ResourceGroupName = rg.Name,
    ///             OfferType = "Standard",
    ///             Kind = "MongoDB",
    ///             EnableAutomaticFailover = true,
    ///             Capabilities = 
    ///             {
    ///                 new Azure.CosmosDB.Inputs.AccountCapabilityArgs
    ///                 {
    ///                     Name = "EnableAggregationPipeline",
    ///                 },
    ///                 new Azure.CosmosDB.Inputs.AccountCapabilityArgs
    ///                 {
    ///                     Name = "mongoEnableDocLevelTTL",
    ///                 },
    ///                 new Azure.CosmosDB.Inputs.AccountCapabilityArgs
    ///                 {
    ///                     Name = "MongoDBv3.4",
    ///                 },
    ///                 new Azure.CosmosDB.Inputs.AccountCapabilityArgs
    ///                 {
    ///                     Name = "EnableMongo",
    ///                 },
    ///             },
    ///             ConsistencyPolicy = new Azure.CosmosDB.Inputs.AccountConsistencyPolicyArgs
    ///             {
    ///                 ConsistencyLevel = "BoundedStaleness",
    ///                 MaxIntervalInSeconds = 300,
    ///                 MaxStalenessPrefix = 100000,
    ///             },
    ///             GeoLocations = 
    ///             {
    ///                 new Azure.CosmosDB.Inputs.AccountGeoLocationArgs
    ///                 {
    ///                     Location = @var.Failover_location,
    ///                     FailoverPriority = 1,
    ///                 },
    ///                 new Azure.CosmosDB.Inputs.AccountGeoLocationArgs
    ///                 {
    ///                     Location = rg.Location,
    ///                     FailoverPriority = 0,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CosmosDB Accounts can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:cosmosdb/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1
    /// ```
    /// </summary>
    [AzureResourceType("azure:cosmosdb/account:Account")]
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
        /// </summary>
        [Output("accessKeyMetadataWritesEnabled")]
        public Output<bool?> AccessKeyMetadataWritesEnabled { get; private set; } = null!;

        /// <summary>
        /// An `analytical_storage` block as defined below.
        /// </summary>
        [Output("analyticalStorage")]
        public Output<Outputs.AccountAnalyticalStorage> AnalyticalStorage { get; private set; } = null!;

        /// <summary>
        /// Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("analyticalStorageEnabled")]
        public Output<bool?> AnalyticalStorageEnabled { get; private set; } = null!;

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Output("backup")]
        public Output<Outputs.AccountBackup> Backup { get; private set; } = null!;

        /// <summary>
        /// The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("capabilities")]
        public Output<ImmutableArray<Outputs.AccountCapability>> Capabilities { get; private set; } = null!;

        /// <summary>
        /// A `capacity` block as defined below.
        /// </summary>
        [Output("capacity")]
        public Output<Outputs.AccountCapacity> Capacity { get; private set; } = null!;

        /// <summary>
        /// A list of connection strings available for this CosmosDB account.
        /// </summary>
        [Output("connectionStrings")]
        public Output<ImmutableArray<string>> ConnectionStrings { get; private set; } = null!;

        /// <summary>
        /// Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        /// </summary>
        [Output("consistencyPolicy")]
        public Output<Outputs.AccountConsistencyPolicy> ConsistencyPolicy { get; private set; } = null!;

        /// <summary>
        /// A `cors_rule` block as defined below.
        /// </summary>
        [Output("corsRule")]
        public Output<Outputs.AccountCorsRule?> CorsRule { get; private set; } = null!;

        /// <summary>
        /// The creation mode for the CosmosDB Account. Possible values are `Default` and `Restore`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("createMode")]
        public Output<string> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The default identity for accessing Key Vault. Possible values are `FirstPartyIdentity`, `SystemAssignedIdentity` or start with `UserAssignedIdentity`. Defaults to `FirstPartyIdentity`.
        /// </summary>
        [Output("defaultIdentityType")]
        public Output<string?> DefaultIdentityType { get; private set; } = null!;

        /// <summary>
        /// Enable automatic fail over for this Cosmos DB account.
        /// </summary>
        [Output("enableAutomaticFailover")]
        public Output<bool?> EnableAutomaticFailover { get; private set; } = null!;

        /// <summary>
        /// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("enableFreeTier")]
        public Output<bool?> EnableFreeTier { get; private set; } = null!;

        /// <summary>
        /// Enable multiple write locations for this Cosmos DB account.
        /// </summary>
        [Output("enableMultipleWriteLocations")]
        public Output<bool?> EnableMultipleWriteLocations { get; private set; } = null!;

        /// <summary>
        /// The endpoint used to connect to the CosmosDB account.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location. Value is a `geo_location` block as defined below.
        /// </summary>
        [Output("geoLocations")]
        public Output<ImmutableArray<Outputs.AccountGeoLocation>> GeoLocations { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.AccountIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        /// </summary>
        [Output("ipRangeFilter")]
        public Output<string?> IpRangeFilter { get; private set; } = null!;

        /// <summary>
        /// Enables virtual network filtering for this Cosmos DB account.
        /// </summary>
        [Output("isVirtualNetworkFilterEnabled")]
        public Output<bool?> IsVirtualNetworkFilterEnabled { get; private set; } = null!;

        /// <summary>
        /// A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        /// </summary>
        [Output("keyVaultKeyId")]
        public Output<string?> KeyVaultKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
        /// </summary>
        [Output("localAuthenticationDisabled")]
        public Output<bool?> LocalAuthenticationDisabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`.
        /// </summary>
        [Output("mongoServerVersion")]
        public Output<string> MongoServerVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If azure services can bypass ACLs. Defaults to `false`.
        /// </summary>
        [Output("networkAclBypassForAzureServices")]
        public Output<bool?> NetworkAclBypassForAzureServices { get; private set; } = null!;

        /// <summary>
        /// The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
        /// </summary>
        [Output("networkAclBypassIds")]
        public Output<ImmutableArray<string>> NetworkAclBypassIds { get; private set; } = null!;

        /// <summary>
        /// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        /// </summary>
        [Output("offerType")]
        public Output<string> OfferType { get; private set; } = null!;

        /// <summary>
        /// The Primary key for the CosmosDB Account.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        [Output("primaryMasterKey")]
        public Output<string> PrimaryMasterKey { get; private set; } = null!;

        /// <summary>
        /// The Primary read-only Key for the CosmosDB Account.
        /// </summary>
        [Output("primaryReadonlyKey")]
        public Output<string> PrimaryReadonlyKey { get; private set; } = null!;

        [Output("primaryReadonlyMasterKey")]
        public Output<string> PrimaryReadonlyMasterKey { get; private set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this CosmosDB account.
        /// </summary>
        [Output("publicNetworkAccessEnabled")]
        public Output<bool?> PublicNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of read endpoints available for this CosmosDB account.
        /// </summary>
        [Output("readEndpoints")]
        public Output<ImmutableArray<string>> ReadEndpoints { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `restore` block as defined below.
        /// </summary>
        [Output("restore")]
        public Output<Outputs.AccountRestore?> Restore { get; private set; } = null!;

        /// <summary>
        /// The Secondary key for the CosmosDB Account.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        [Output("secondaryMasterKey")]
        public Output<string> SecondaryMasterKey { get; private set; } = null!;

        /// <summary>
        /// The Secondary read-only key for the CosmosDB Account.
        /// </summary>
        [Output("secondaryReadonlyKey")]
        public Output<string> SecondaryReadonlyKey { get; private set; } = null!;

        [Output("secondaryReadonlyMasterKey")]
        public Output<string> SecondaryReadonlyMasterKey { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
        /// </summary>
        [Output("virtualNetworkRules")]
        public Output<ImmutableArray<Outputs.AccountVirtualNetworkRule>> VirtualNetworkRules { get; private set; } = null!;

        /// <summary>
        /// A list of write endpoints available for this CosmosDB account.
        /// </summary>
        [Output("writeEndpoints")]
        public Output<ImmutableArray<string>> WriteEndpoints { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azure:cosmosdb/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:cosmosdb/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
        /// </summary>
        [Input("accessKeyMetadataWritesEnabled")]
        public Input<bool>? AccessKeyMetadataWritesEnabled { get; set; }

        /// <summary>
        /// An `analytical_storage` block as defined below.
        /// </summary>
        [Input("analyticalStorage")]
        public Input<Inputs.AccountAnalyticalStorageArgs>? AnalyticalStorage { get; set; }

        /// <summary>
        /// Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("analyticalStorageEnabled")]
        public Input<bool>? AnalyticalStorageEnabled { get; set; }

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Input("backup")]
        public Input<Inputs.AccountBackupArgs>? Backup { get; set; }

        [Input("capabilities")]
        private InputList<Inputs.AccountCapabilityArgs>? _capabilities;

        /// <summary>
        /// The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<Inputs.AccountCapabilityArgs> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<Inputs.AccountCapabilityArgs>());
            set => _capabilities = value;
        }

        /// <summary>
        /// A `capacity` block as defined below.
        /// </summary>
        [Input("capacity")]
        public Input<Inputs.AccountCapacityArgs>? Capacity { get; set; }

        /// <summary>
        /// Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        /// </summary>
        [Input("consistencyPolicy", required: true)]
        public Input<Inputs.AccountConsistencyPolicyArgs> ConsistencyPolicy { get; set; } = null!;

        /// <summary>
        /// A `cors_rule` block as defined below.
        /// </summary>
        [Input("corsRule")]
        public Input<Inputs.AccountCorsRuleArgs>? CorsRule { get; set; }

        /// <summary>
        /// The creation mode for the CosmosDB Account. Possible values are `Default` and `Restore`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The default identity for accessing Key Vault. Possible values are `FirstPartyIdentity`, `SystemAssignedIdentity` or start with `UserAssignedIdentity`. Defaults to `FirstPartyIdentity`.
        /// </summary>
        [Input("defaultIdentityType")]
        public Input<string>? DefaultIdentityType { get; set; }

        /// <summary>
        /// Enable automatic fail over for this Cosmos DB account.
        /// </summary>
        [Input("enableAutomaticFailover")]
        public Input<bool>? EnableAutomaticFailover { get; set; }

        /// <summary>
        /// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("enableFreeTier")]
        public Input<bool>? EnableFreeTier { get; set; }

        /// <summary>
        /// Enable multiple write locations for this Cosmos DB account.
        /// </summary>
        [Input("enableMultipleWriteLocations")]
        public Input<bool>? EnableMultipleWriteLocations { get; set; }

        [Input("geoLocations", required: true)]
        private InputList<Inputs.AccountGeoLocationArgs>? _geoLocations;

        /// <summary>
        /// Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location. Value is a `geo_location` block as defined below.
        /// </summary>
        public InputList<Inputs.AccountGeoLocationArgs> GeoLocations
        {
            get => _geoLocations ?? (_geoLocations = new InputList<Inputs.AccountGeoLocationArgs>());
            set => _geoLocations = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AccountIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        /// </summary>
        [Input("ipRangeFilter")]
        public Input<string>? IpRangeFilter { get; set; }

        /// <summary>
        /// Enables virtual network filtering for this Cosmos DB account.
        /// </summary>
        [Input("isVirtualNetworkFilterEnabled")]
        public Input<bool>? IsVirtualNetworkFilterEnabled { get; set; }

        /// <summary>
        /// A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
        /// </summary>
        [Input("localAuthenticationDisabled")]
        public Input<bool>? LocalAuthenticationDisabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`.
        /// </summary>
        [Input("mongoServerVersion")]
        public Input<string>? MongoServerVersion { get; set; }

        /// <summary>
        /// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If azure services can bypass ACLs. Defaults to `false`.
        /// </summary>
        [Input("networkAclBypassForAzureServices")]
        public Input<bool>? NetworkAclBypassForAzureServices { get; set; }

        [Input("networkAclBypassIds")]
        private InputList<string>? _networkAclBypassIds;

        /// <summary>
        /// The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
        /// </summary>
        public InputList<string> NetworkAclBypassIds
        {
            get => _networkAclBypassIds ?? (_networkAclBypassIds = new InputList<string>());
            set => _networkAclBypassIds = value;
        }

        /// <summary>
        /// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        /// </summary>
        [Input("offerType", required: true)]
        public Input<string> OfferType { get; set; } = null!;

        /// <summary>
        /// Whether or not public network access is allowed for this CosmosDB account.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        /// <summary>
        /// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `restore` block as defined below.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.AccountRestoreArgs>? Restore { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("virtualNetworkRules")]
        private InputList<Inputs.AccountVirtualNetworkRuleArgs>? _virtualNetworkRules;

        /// <summary>
        /// Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
        /// </summary>
        public InputList<Inputs.AccountVirtualNetworkRuleArgs> VirtualNetworkRules
        {
            get => _virtualNetworkRules ?? (_virtualNetworkRules = new InputList<Inputs.AccountVirtualNetworkRuleArgs>());
            set => _virtualNetworkRules = value;
        }

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
        /// </summary>
        [Input("accessKeyMetadataWritesEnabled")]
        public Input<bool>? AccessKeyMetadataWritesEnabled { get; set; }

        /// <summary>
        /// An `analytical_storage` block as defined below.
        /// </summary>
        [Input("analyticalStorage")]
        public Input<Inputs.AccountAnalyticalStorageGetArgs>? AnalyticalStorage { get; set; }

        /// <summary>
        /// Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("analyticalStorageEnabled")]
        public Input<bool>? AnalyticalStorageEnabled { get; set; }

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Input("backup")]
        public Input<Inputs.AccountBackupGetArgs>? Backup { get; set; }

        [Input("capabilities")]
        private InputList<Inputs.AccountCapabilityGetArgs>? _capabilities;

        /// <summary>
        /// The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<Inputs.AccountCapabilityGetArgs> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<Inputs.AccountCapabilityGetArgs>());
            set => _capabilities = value;
        }

        /// <summary>
        /// A `capacity` block as defined below.
        /// </summary>
        [Input("capacity")]
        public Input<Inputs.AccountCapacityGetArgs>? Capacity { get; set; }

        [Input("connectionStrings")]
        private InputList<string>? _connectionStrings;

        /// <summary>
        /// A list of connection strings available for this CosmosDB account.
        /// </summary>
        public InputList<string> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<string>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        /// </summary>
        [Input("consistencyPolicy")]
        public Input<Inputs.AccountConsistencyPolicyGetArgs>? ConsistencyPolicy { get; set; }

        /// <summary>
        /// A `cors_rule` block as defined below.
        /// </summary>
        [Input("corsRule")]
        public Input<Inputs.AccountCorsRuleGetArgs>? CorsRule { get; set; }

        /// <summary>
        /// The creation mode for the CosmosDB Account. Possible values are `Default` and `Restore`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The default identity for accessing Key Vault. Possible values are `FirstPartyIdentity`, `SystemAssignedIdentity` or start with `UserAssignedIdentity`. Defaults to `FirstPartyIdentity`.
        /// </summary>
        [Input("defaultIdentityType")]
        public Input<string>? DefaultIdentityType { get; set; }

        /// <summary>
        /// Enable automatic fail over for this Cosmos DB account.
        /// </summary>
        [Input("enableAutomaticFailover")]
        public Input<bool>? EnableAutomaticFailover { get; set; }

        /// <summary>
        /// Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("enableFreeTier")]
        public Input<bool>? EnableFreeTier { get; set; }

        /// <summary>
        /// Enable multiple write locations for this Cosmos DB account.
        /// </summary>
        [Input("enableMultipleWriteLocations")]
        public Input<bool>? EnableMultipleWriteLocations { get; set; }

        /// <summary>
        /// The endpoint used to connect to the CosmosDB account.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("geoLocations")]
        private InputList<Inputs.AccountGeoLocationGetArgs>? _geoLocations;

        /// <summary>
        /// Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location. Value is a `geo_location` block as defined below.
        /// </summary>
        public InputList<Inputs.AccountGeoLocationGetArgs> GeoLocations
        {
            get => _geoLocations ?? (_geoLocations = new InputList<Inputs.AccountGeoLocationGetArgs>());
            set => _geoLocations = value;
        }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AccountIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        /// </summary>
        [Input("ipRangeFilter")]
        public Input<string>? IpRangeFilter { get; set; }

        /// <summary>
        /// Enables virtual network filtering for this Cosmos DB account.
        /// </summary>
        [Input("isVirtualNetworkFilterEnabled")]
        public Input<bool>? IsVirtualNetworkFilterEnabled { get; set; }

        /// <summary>
        /// A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
        /// </summary>
        [Input("localAuthenticationDisabled")]
        public Input<bool>? LocalAuthenticationDisabled { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`.
        /// </summary>
        [Input("mongoServerVersion")]
        public Input<string>? MongoServerVersion { get; set; }

        /// <summary>
        /// Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If azure services can bypass ACLs. Defaults to `false`.
        /// </summary>
        [Input("networkAclBypassForAzureServices")]
        public Input<bool>? NetworkAclBypassForAzureServices { get; set; }

        [Input("networkAclBypassIds")]
        private InputList<string>? _networkAclBypassIds;

        /// <summary>
        /// The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
        /// </summary>
        public InputList<string> NetworkAclBypassIds
        {
            get => _networkAclBypassIds ?? (_networkAclBypassIds = new InputList<string>());
            set => _networkAclBypassIds = value;
        }

        /// <summary>
        /// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        /// </summary>
        [Input("offerType")]
        public Input<string>? OfferType { get; set; }

        /// <summary>
        /// The Primary key for the CosmosDB Account.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        [Input("primaryMasterKey")]
        public Input<string>? PrimaryMasterKey { get; set; }

        /// <summary>
        /// The Primary read-only Key for the CosmosDB Account.
        /// </summary>
        [Input("primaryReadonlyKey")]
        public Input<string>? PrimaryReadonlyKey { get; set; }

        [Input("primaryReadonlyMasterKey")]
        public Input<string>? PrimaryReadonlyMasterKey { get; set; }

        /// <summary>
        /// Whether or not public network access is allowed for this CosmosDB account.
        /// </summary>
        [Input("publicNetworkAccessEnabled")]
        public Input<bool>? PublicNetworkAccessEnabled { get; set; }

        [Input("readEndpoints")]
        private InputList<string>? _readEndpoints;

        /// <summary>
        /// A list of read endpoints available for this CosmosDB account.
        /// </summary>
        public InputList<string> ReadEndpoints
        {
            get => _readEndpoints ?? (_readEndpoints = new InputList<string>());
            set => _readEndpoints = value;
        }

        /// <summary>
        /// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `restore` block as defined below.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.AccountRestoreGetArgs>? Restore { get; set; }

        /// <summary>
        /// The Secondary key for the CosmosDB Account.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        [Input("secondaryMasterKey")]
        public Input<string>? SecondaryMasterKey { get; set; }

        /// <summary>
        /// The Secondary read-only key for the CosmosDB Account.
        /// </summary>
        [Input("secondaryReadonlyKey")]
        public Input<string>? SecondaryReadonlyKey { get; set; }

        [Input("secondaryReadonlyMasterKey")]
        public Input<string>? SecondaryReadonlyMasterKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("virtualNetworkRules")]
        private InputList<Inputs.AccountVirtualNetworkRuleGetArgs>? _virtualNetworkRules;

        /// <summary>
        /// Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
        /// </summary>
        public InputList<Inputs.AccountVirtualNetworkRuleGetArgs> VirtualNetworkRules
        {
            get => _virtualNetworkRules ?? (_virtualNetworkRules = new InputList<Inputs.AccountVirtualNetworkRuleGetArgs>());
            set => _virtualNetworkRules = value;
        }

        [Input("writeEndpoints")]
        private InputList<string>? _writeEndpoints;

        /// <summary>
        /// A list of write endpoints available for this CosmosDB account.
        /// </summary>
        public InputList<string> WriteEndpoints
        {
            get => _writeEndpoints ?? (_writeEndpoints = new InputList<string>());
            set => _writeEndpoints = value;
        }

        public AccountState()
        {
        }
    }
}
