// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Mongo Collection within a Cosmos DB Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleAccount = azure.cosmosdb.getAccount({
 *     name: "tfex-cosmosdb-account",
 *     resourceGroupName: "tfex-cosmosdb-account-rg",
 * });
 * const exampleMongoDatabase = new azure.cosmosdb.MongoDatabase("exampleMongoDatabase", {
 *     resourceGroupName: exampleAccount.then(exampleAccount => exampleAccount.resourceGroupName),
 *     accountName: exampleAccount.then(exampleAccount => exampleAccount.name),
 * });
 * const exampleMongoCollection = new azure.cosmosdb.MongoCollection("exampleMongoCollection", {
 *     resourceGroupName: exampleAccount.then(exampleAccount => exampleAccount.resourceGroupName),
 *     accountName: exampleAccount.then(exampleAccount => exampleAccount.name),
 *     databaseName: exampleMongoDatabase.name,
 *     defaultTtlSeconds: "777",
 *     shardKey: "uniqueKey",
 *     throughput: 400,
 * });
 * ```
 *
 * ## Import
 *
 * CosmosDB Mongo Collection can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:cosmosdb/mongoCollection:MongoCollection collection1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1/collections/collection1
 * ```
 */
export class MongoCollection extends pulumi.CustomResource {
    /**
     * Get an existing MongoCollection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MongoCollectionState, opts?: pulumi.CustomResourceOptions): MongoCollection {
        return new MongoCollection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:cosmosdb/mongoCollection:MongoCollection';

    /**
     * Returns true if the given object is an instance of MongoCollection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MongoCollection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MongoCollection.__pulumiType;
    }

    public readonly accountName!: pulumi.Output<string>;
    /**
     * The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
     */
    public readonly analyticalStorageTtl!: pulumi.Output<number | undefined>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
     */
    public readonly autoscaleSettings!: pulumi.Output<outputs.cosmosdb.MongoCollectionAutoscaleSettings | undefined>;
    /**
     * The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
     */
    public readonly defaultTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * One or more `index` blocks as defined below.
     */
    public readonly indices!: pulumi.Output<outputs.cosmosdb.MongoCollectionIndex[] | undefined>;
    /**
     * Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The name of the key to partition on for sharding. There must not be any other unique index keys.
     */
    public readonly shardKey!: pulumi.Output<string | undefined>;
    /**
     * One or more `systemIndexes` blocks as defined below.
     */
    public /*out*/ readonly systemIndexes!: pulumi.Output<outputs.cosmosdb.MongoCollectionSystemIndex[]>;
    /**
     * The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    public readonly throughput!: pulumi.Output<number>;

    /**
     * Create a MongoCollection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MongoCollectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MongoCollectionArgs | MongoCollectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MongoCollectionState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["analyticalStorageTtl"] = state ? state.analyticalStorageTtl : undefined;
            inputs["autoscaleSettings"] = state ? state.autoscaleSettings : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["defaultTtlSeconds"] = state ? state.defaultTtlSeconds : undefined;
            inputs["indices"] = state ? state.indices : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["shardKey"] = state ? state.shardKey : undefined;
            inputs["systemIndexes"] = state ? state.systemIndexes : undefined;
            inputs["throughput"] = state ? state.throughput : undefined;
        } else {
            const args = argsOrState as MongoCollectionArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["analyticalStorageTtl"] = args ? args.analyticalStorageTtl : undefined;
            inputs["autoscaleSettings"] = args ? args.autoscaleSettings : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["defaultTtlSeconds"] = args ? args.defaultTtlSeconds : undefined;
            inputs["indices"] = args ? args.indices : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shardKey"] = args ? args.shardKey : undefined;
            inputs["throughput"] = args ? args.throughput : undefined;
            inputs["systemIndexes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MongoCollection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MongoCollection resources.
 */
export interface MongoCollectionState {
    readonly accountName?: pulumi.Input<string>;
    /**
     * The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
     */
    readonly analyticalStorageTtl?: pulumi.Input<number>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
     */
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.MongoCollectionAutoscaleSettings>;
    /**
     * The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
     */
    readonly defaultTtlSeconds?: pulumi.Input<number>;
    /**
     * One or more `index` blocks as defined below.
     */
    readonly indices?: pulumi.Input<pulumi.Input<inputs.cosmosdb.MongoCollectionIndex>[]>;
    /**
     * Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The name of the key to partition on for sharding. There must not be any other unique index keys.
     */
    readonly shardKey?: pulumi.Input<string>;
    /**
     * One or more `systemIndexes` blocks as defined below.
     */
    readonly systemIndexes?: pulumi.Input<pulumi.Input<inputs.cosmosdb.MongoCollectionSystemIndex>[]>;
    /**
     * The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    readonly throughput?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a MongoCollection resource.
 */
export interface MongoCollectionArgs {
    readonly accountName: pulumi.Input<string>;
    /**
     * The default time to live of Analytical Storage for this Mongo Collection. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
     */
    readonly analyticalStorageTtl?: pulumi.Input<number>;
    /**
     * An `autoscaleSettings` block as defined below. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply. Requires `shardKey` to be set.
     */
    readonly autoscaleSettings?: pulumi.Input<inputs.cosmosdb.MongoCollectionAutoscaleSettings>;
    /**
     * The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The default Time To Live in seconds. If the value is `-1` or `0`, items are not automatically expired.
     */
    readonly defaultTtlSeconds?: pulumi.Input<number>;
    /**
     * One or more `index` blocks as defined below.
     */
    readonly indices?: pulumi.Input<pulumi.Input<inputs.cosmosdb.MongoCollectionIndex>[]>;
    /**
     * Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the key to partition on for sharding. There must not be any other unique index keys.
     */
    readonly shardKey?: pulumi.Input<string>;
    /**
     * The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual manual destroy-apply.
     */
    readonly throughput?: pulumi.Input<number>;
}
