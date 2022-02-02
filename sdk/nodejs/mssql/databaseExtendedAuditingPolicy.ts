// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Ms Sql Database Extended Auditing Policy.
 *
 * > **NOTE:** The Database Extended Auditing Policy can also be set in the `extendedAuditingPolicy` block in the azure.mssql.Database resource. You can only use one or the other and using both will cause a conflict.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "AdminPassword123!",
 * });
 * const exampleDatabase = new azure.mssql.Database("exampleDatabase", {serverId: exampleServer.id});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleDatabaseExtendedAuditingPolicy = new azure.mssql.DatabaseExtendedAuditingPolicy("exampleDatabaseExtendedAuditingPolicy", {
 *     databaseId: exampleDatabase.id,
 *     storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *     storageAccountAccessKeyIsSecondary: false,
 *     retentionInDays: 6,
 * });
 * ```
 *
 * ## Import
 *
 * Ms Sql Database Extended Auditing Policys can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
 * ```
 */
export class DatabaseExtendedAuditingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseExtendedAuditingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions): DatabaseExtendedAuditingPolicy {
        return new DatabaseExtendedAuditingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/databaseExtendedAuditingPolicy:DatabaseExtendedAuditingPolicy';

    /**
     * Returns true if the given object is an instance of DatabaseExtendedAuditingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseExtendedAuditingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseExtendedAuditingPolicy.__pulumiType;
    }

    /**
     * The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * Enable audit events to Azure Monitor?
     */
    public readonly logMonitoringEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The access key to use for the auditing storage account.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    public readonly storageAccountAccessKeyIsSecondary!: pulumi.Output<boolean | undefined>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    public readonly storageEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a DatabaseExtendedAuditingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseExtendedAuditingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseExtendedAuditingPolicyArgs | DatabaseExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseExtendedAuditingPolicyState | undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["logMonitoringEnabled"] = state ? state.logMonitoringEnabled : undefined;
            resourceInputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            resourceInputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountAccessKeyIsSecondary"] = state ? state.storageAccountAccessKeyIsSecondary : undefined;
            resourceInputs["storageEndpoint"] = state ? state.storageEndpoint : undefined;
        } else {
            const args = argsOrState as DatabaseExtendedAuditingPolicyArgs | undefined;
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["logMonitoringEnabled"] = args ? args.logMonitoringEnabled : undefined;
            resourceInputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            resourceInputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountAccessKeyIsSecondary"] = args ? args.storageAccountAccessKeyIsSecondary : undefined;
            resourceInputs["storageEndpoint"] = args ? args.storageEndpoint : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseExtendedAuditingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseExtendedAuditingPolicy resources.
 */
export interface DatabaseExtendedAuditingPolicyState {
    /**
     * The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Enable audit events to Azure Monitor?
     */
    logMonitoringEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * The access key to use for the auditing storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseExtendedAuditingPolicy resource.
 */
export interface DatabaseExtendedAuditingPolicyArgs {
    /**
     * The ID of the sql database to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    databaseId: pulumi.Input<string>;
    /**
     * Enable audit events to Azure Monitor?
     */
    logMonitoringEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * The access key to use for the auditing storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}
