// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a CosmosDB (formally DocumentDB) Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as random from "@pulumi/random";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "westus"});
 * const ri = new random.RandomInteger("ri", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const db = new azure.cosmosdb.Account("db", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     offerType: "Standard",
 *     kind: "GlobalDocumentDB",
 *     enableAutomaticFailover: true,
 *     capabilities: [
 *         {
 *             name: "EnableAggregationPipeline",
 *         },
 *         {
 *             name: "mongoEnableDocLevelTTL",
 *         },
 *         {
 *             name: "MongoDBv3.4",
 *         },
 *     ],
 *     consistencyPolicy: {
 *         consistencyLevel: "BoundedStaleness",
 *         maxIntervalInSeconds: 10,
 *         maxStalenessPrefix: 200,
 *     },
 *     geoLocations: [
 *         {
 *             location: _var.failover_location,
 *             failoverPriority: 1,
 *         },
 *         {
 *             location: rg.location,
 *             failoverPriority: 0,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * CosmosDB Accounts can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:cosmosdb/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cosmosdb/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
     */
    public readonly accessKeyMetadataWritesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly analyticalStorageEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A `backup` block as defined below.
     */
    public readonly backup!: pulumi.Output<outputs.cosmosdb.AccountBackup>;
    /**
     * The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below.
     */
    public readonly capabilities!: pulumi.Output<outputs.cosmosdb.AccountCapability[] | undefined>;
    /**
     * A list of connection strings available for this CosmosDB account.
     */
    public /*out*/ readonly connectionStrings!: pulumi.Output<string[]>;
    /**
     * Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
     */
    public readonly consistencyPolicy!: pulumi.Output<outputs.cosmosdb.AccountConsistencyPolicy>;
    /**
     * A `corsRule` block as defined below.
     */
    public readonly corsRule!: pulumi.Output<outputs.cosmosdb.AccountCorsRule | undefined>;
    /**
     * Enable automatic fail over for this Cosmos DB account.
     */
    public readonly enableAutomaticFailover!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly enableFreeTier!: pulumi.Output<boolean | undefined>;
    /**
     * Enable multiple write locations for this Cosmos DB account.
     */
    public readonly enableMultipleWriteLocations!: pulumi.Output<boolean | undefined>;
    /**
     * The endpoint used to connect to the CosmosDB account.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location. Value is a `geoLocation` block as defined below.
     */
    public readonly geoLocations!: pulumi.Output<outputs.cosmosdb.AccountGeoLocation[]>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.cosmosdb.AccountIdentity | undefined>;
    /**
     * CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
     */
    public readonly ipRangeFilter!: pulumi.Output<string | undefined>;
    /**
     * Enables virtual network filtering for this Cosmos DB account.
     */
    public readonly isVirtualNetworkFilterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
     */
    public readonly keyVaultKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
     */
    public readonly localAuthenticationDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`. Changing this forces a new resource to be created.
     */
    public readonly mongoServerVersion!: pulumi.Output<string>;
    /**
     * Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If azure services can bypass ACLs. Defaults to `false`.
     */
    public readonly networkAclBypassForAzureServices!: pulumi.Output<boolean | undefined>;
    /**
     * The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
     */
    public readonly networkAclBypassIds!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
     */
    public readonly offerType!: pulumi.Output<string>;
    /**
     * The Primary key for the CosmosDB Account.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * @deprecated This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    public /*out*/ readonly primaryMasterKey!: pulumi.Output<string>;
    /**
     * The Primary read-only Key for the CosmosDB Account.
     */
    public /*out*/ readonly primaryReadonlyKey!: pulumi.Output<string>;
    /**
     * @deprecated This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    public /*out*/ readonly primaryReadonlyMasterKey!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this CosmosDB account.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of read endpoints available for this CosmosDB account.
     */
    public /*out*/ readonly readEndpoints!: pulumi.Output<string[]>;
    /**
     * The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary key for the CosmosDB Account.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * @deprecated This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    public /*out*/ readonly secondaryMasterKey!: pulumi.Output<string>;
    /**
     * The Secondary read-only key for the CosmosDB Account.
     */
    public /*out*/ readonly secondaryReadonlyKey!: pulumi.Output<string>;
    /**
     * @deprecated This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    public /*out*/ readonly secondaryReadonlyMasterKey!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
     */
    public readonly virtualNetworkRules!: pulumi.Output<outputs.cosmosdb.AccountVirtualNetworkRule[] | undefined>;
    /**
     * A list of write endpoints available for this CosmosDB account.
     */
    public /*out*/ readonly writeEndpoints!: pulumi.Output<string[]>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            inputs["accessKeyMetadataWritesEnabled"] = state ? state.accessKeyMetadataWritesEnabled : undefined;
            inputs["analyticalStorageEnabled"] = state ? state.analyticalStorageEnabled : undefined;
            inputs["backup"] = state ? state.backup : undefined;
            inputs["capabilities"] = state ? state.capabilities : undefined;
            inputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            inputs["consistencyPolicy"] = state ? state.consistencyPolicy : undefined;
            inputs["corsRule"] = state ? state.corsRule : undefined;
            inputs["enableAutomaticFailover"] = state ? state.enableAutomaticFailover : undefined;
            inputs["enableFreeTier"] = state ? state.enableFreeTier : undefined;
            inputs["enableMultipleWriteLocations"] = state ? state.enableMultipleWriteLocations : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["geoLocations"] = state ? state.geoLocations : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["ipRangeFilter"] = state ? state.ipRangeFilter : undefined;
            inputs["isVirtualNetworkFilterEnabled"] = state ? state.isVirtualNetworkFilterEnabled : undefined;
            inputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["localAuthenticationDisabled"] = state ? state.localAuthenticationDisabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["mongoServerVersion"] = state ? state.mongoServerVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkAclBypassForAzureServices"] = state ? state.networkAclBypassForAzureServices : undefined;
            inputs["networkAclBypassIds"] = state ? state.networkAclBypassIds : undefined;
            inputs["offerType"] = state ? state.offerType : undefined;
            inputs["primaryKey"] = state ? state.primaryKey : undefined;
            inputs["primaryMasterKey"] = state ? state.primaryMasterKey : undefined;
            inputs["primaryReadonlyKey"] = state ? state.primaryReadonlyKey : undefined;
            inputs["primaryReadonlyMasterKey"] = state ? state.primaryReadonlyMasterKey : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["readEndpoints"] = state ? state.readEndpoints : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            inputs["secondaryMasterKey"] = state ? state.secondaryMasterKey : undefined;
            inputs["secondaryReadonlyKey"] = state ? state.secondaryReadonlyKey : undefined;
            inputs["secondaryReadonlyMasterKey"] = state ? state.secondaryReadonlyMasterKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualNetworkRules"] = state ? state.virtualNetworkRules : undefined;
            inputs["writeEndpoints"] = state ? state.writeEndpoints : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.consistencyPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consistencyPolicy'");
            }
            if ((!args || args.geoLocations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'geoLocations'");
            }
            if ((!args || args.offerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'offerType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessKeyMetadataWritesEnabled"] = args ? args.accessKeyMetadataWritesEnabled : undefined;
            inputs["analyticalStorageEnabled"] = args ? args.analyticalStorageEnabled : undefined;
            inputs["backup"] = args ? args.backup : undefined;
            inputs["capabilities"] = args ? args.capabilities : undefined;
            inputs["consistencyPolicy"] = args ? args.consistencyPolicy : undefined;
            inputs["corsRule"] = args ? args.corsRule : undefined;
            inputs["enableAutomaticFailover"] = args ? args.enableAutomaticFailover : undefined;
            inputs["enableFreeTier"] = args ? args.enableFreeTier : undefined;
            inputs["enableMultipleWriteLocations"] = args ? args.enableMultipleWriteLocations : undefined;
            inputs["geoLocations"] = args ? args.geoLocations : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["ipRangeFilter"] = args ? args.ipRangeFilter : undefined;
            inputs["isVirtualNetworkFilterEnabled"] = args ? args.isVirtualNetworkFilterEnabled : undefined;
            inputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["localAuthenticationDisabled"] = args ? args.localAuthenticationDisabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mongoServerVersion"] = args ? args.mongoServerVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkAclBypassForAzureServices"] = args ? args.networkAclBypassForAzureServices : undefined;
            inputs["networkAclBypassIds"] = args ? args.networkAclBypassIds : undefined;
            inputs["offerType"] = args ? args.offerType : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkRules"] = args ? args.virtualNetworkRules : undefined;
            inputs["connectionStrings"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["primaryMasterKey"] = undefined /*out*/;
            inputs["primaryReadonlyKey"] = undefined /*out*/;
            inputs["primaryReadonlyMasterKey"] = undefined /*out*/;
            inputs["readEndpoints"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
            inputs["secondaryMasterKey"] = undefined /*out*/;
            inputs["secondaryReadonlyKey"] = undefined /*out*/;
            inputs["secondaryReadonlyMasterKey"] = undefined /*out*/;
            inputs["writeEndpoints"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Account.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
     */
    accessKeyMetadataWritesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    analyticalStorageEnabled?: pulumi.Input<boolean>;
    /**
     * A `backup` block as defined below.
     */
    backup?: pulumi.Input<inputs.cosmosdb.AccountBackup>;
    /**
     * The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below.
     */
    capabilities?: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountCapability>[]>;
    /**
     * A list of connection strings available for this CosmosDB account.
     */
    connectionStrings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
     */
    consistencyPolicy?: pulumi.Input<inputs.cosmosdb.AccountConsistencyPolicy>;
    /**
     * A `corsRule` block as defined below.
     */
    corsRule?: pulumi.Input<inputs.cosmosdb.AccountCorsRule>;
    /**
     * Enable automatic fail over for this Cosmos DB account.
     */
    enableAutomaticFailover?: pulumi.Input<boolean>;
    /**
     * Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    enableFreeTier?: pulumi.Input<boolean>;
    /**
     * Enable multiple write locations for this Cosmos DB account.
     */
    enableMultipleWriteLocations?: pulumi.Input<boolean>;
    /**
     * The endpoint used to connect to the CosmosDB account.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location. Value is a `geoLocation` block as defined below.
     */
    geoLocations?: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountGeoLocation>[]>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.cosmosdb.AccountIdentity>;
    /**
     * CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
     */
    ipRangeFilter?: pulumi.Input<string>;
    /**
     * Enables virtual network filtering for this Cosmos DB account.
     */
    isVirtualNetworkFilterEnabled?: pulumi.Input<boolean>;
    /**
     * A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
     */
    kind?: pulumi.Input<string>;
    /**
     * Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
     */
    localAuthenticationDisabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`. Changing this forces a new resource to be created.
     */
    mongoServerVersion?: pulumi.Input<string>;
    /**
     * Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * If azure services can bypass ACLs. Defaults to `false`.
     */
    networkAclBypassForAzureServices?: pulumi.Input<boolean>;
    /**
     * The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
     */
    networkAclBypassIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
     */
    offerType?: pulumi.Input<string>;
    /**
     * The Primary key for the CosmosDB Account.
     */
    primaryKey?: pulumi.Input<string>;
    /**
     * @deprecated This property has been renamed to `primary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    primaryMasterKey?: pulumi.Input<string>;
    /**
     * The Primary read-only Key for the CosmosDB Account.
     */
    primaryReadonlyKey?: pulumi.Input<string>;
    /**
     * @deprecated This property has been renamed to `primary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    primaryReadonlyMasterKey?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this CosmosDB account.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * A list of read endpoints available for this CosmosDB account.
     */
    readEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary key for the CosmosDB Account.
     */
    secondaryKey?: pulumi.Input<string>;
    /**
     * @deprecated This property has been renamed to `secondary_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    secondaryMasterKey?: pulumi.Input<string>;
    /**
     * The Secondary read-only key for the CosmosDB Account.
     */
    secondaryReadonlyKey?: pulumi.Input<string>;
    /**
     * @deprecated This property has been renamed to `secondary_readonly_key` and will be removed in v3.0 of the provider in support of HashiCorp's inclusive language policy which can be found here: https://discuss.hashicorp.com/t/inclusive-language-changes
     */
    secondaryReadonlyMasterKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
     */
    virtualNetworkRules?: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountVirtualNetworkRule>[]>;
    /**
     * A list of write endpoints available for this CosmosDB account.
     */
    writeEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Is write operations on metadata resources (databases, containers, throughput) via account keys enabled? Defaults to `true`.
     */
    accessKeyMetadataWritesEnabled?: pulumi.Input<boolean>;
    /**
     * Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    analyticalStorageEnabled?: pulumi.Input<boolean>;
    /**
     * A `backup` block as defined below.
     */
    backup?: pulumi.Input<inputs.cosmosdb.AccountBackup>;
    /**
     * The capabilities which should be enabled for this Cosmos DB account. Value is a `capabilities` block as defined below.
     */
    capabilities?: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountCapability>[]>;
    /**
     * Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
     */
    consistencyPolicy: pulumi.Input<inputs.cosmosdb.AccountConsistencyPolicy>;
    /**
     * A `corsRule` block as defined below.
     */
    corsRule?: pulumi.Input<inputs.cosmosdb.AccountCorsRule>;
    /**
     * Enable automatic fail over for this Cosmos DB account.
     */
    enableAutomaticFailover?: pulumi.Input<boolean>;
    /**
     * Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
     */
    enableFreeTier?: pulumi.Input<boolean>;
    /**
     * Enable multiple write locations for this Cosmos DB account.
     */
    enableMultipleWriteLocations?: pulumi.Input<boolean>;
    /**
     * Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location. Value is a `geoLocation` block as defined below.
     */
    geoLocations: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountGeoLocation>[]>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.cosmosdb.AccountIdentity>;
    /**
     * CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
     */
    ipRangeFilter?: pulumi.Input<string>;
    /**
     * Enables virtual network filtering for this Cosmos DB account.
     */
    isVirtualNetworkFilterEnabled?: pulumi.Input<boolean>;
    /**
     * A versionless Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
     */
    kind?: pulumi.Input<string>;
    /**
     * Disable local authentication and ensure only MSI and AAD can be used exclusively for authentication. Defaults to `false`. Can be set only when using the SQL API.
     */
    localAuthenticationDisabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The Server Version of a MongoDB account. Possible values are `4.0`, `3.6`, and `3.2`. Changing this forces a new resource to be created.
     */
    mongoServerVersion?: pulumi.Input<string>;
    /**
     * Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * If azure services can bypass ACLs. Defaults to `false`.
     */
    networkAclBypassForAzureServices?: pulumi.Input<boolean>;
    /**
     * The list of resource Ids for Network Acl Bypass for this Cosmos DB account.
     */
    networkAclBypassIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
     */
    offerType: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this CosmosDB account.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
     */
    virtualNetworkRules?: pulumi.Input<pulumi.Input<inputs.cosmosdb.AccountVirtualNetworkRule>[]>;
}
