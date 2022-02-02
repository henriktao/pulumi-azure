// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Ms Sql Server Extended Auditing Policy.
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
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleServerExtendedAuditingPolicy = new azure.mssql.ServerExtendedAuditingPolicy("exampleServerExtendedAuditingPolicy", {
 *     serverId: exampleServer.id,
 *     storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *     storageAccountAccessKeyIsSecondary: false,
 *     retentionInDays: 6,
 * });
 * ```
 * ### With Storage Account Behind VNet And Firewall
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const primary = azure.core.getSubscription({});
 * const exampleClientConfig = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.2.0/24"],
 *     serviceEndpoints: [
 *         "Microsoft.Sql",
 *         "Microsoft.Storage",
 *     ],
 *     enforcePrivateLinkEndpointNetworkPolicies: true,
 * });
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "AdminPassword123!",
 *     minimumTlsVersion: "1.2",
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleAssignment = new azure.authorization.Assignment("exampleAssignment", {
 *     scope: primary.then(primary => primary.id),
 *     roleDefinitionName: "Storage Blob Data Contributor",
 *     principalId: exampleServer.identity.apply(identity => identity?.principalId),
 * });
 * const sqlvnetrule = new azure.sql.VirtualNetworkRule("sqlvnetrule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleServer.name,
 *     subnetId: exampleSubnet.id,
 * });
 * const exampleFirewallRule = new azure.sql.FirewallRule("exampleFirewallRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     serverName: exampleServer.name,
 *     startIpAddress: "0.0.0.0",
 *     endIpAddress: "0.0.0.0",
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     allowBlobPublicAccess: false,
 *     networkRules: {
 *         defaultAction: "Deny",
 *         ipRules: ["127.0.0.1"],
 *         virtualNetworkSubnetIds: [exampleSubnet.id],
 *         bypasses: ["AzureServices"],
 *     },
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const exampleServerExtendedAuditingPolicy = new azure.mssql.ServerExtendedAuditingPolicy("exampleServerExtendedAuditingPolicy", {
 *     storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *     serverId: exampleServer.id,
 *     retentionInDays: 6,
 *     logMonitoringEnabled: false,
 *     storageAccountSubscriptionId: azurerm_subscription.primary.subscription_id,
 * }, {
 *     dependsOn: [
 *         exampleAssignment,
 *         exampleAccount,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Ms Sql Server Extended Auditing Policys can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/extendedAuditingSettings/default
 * ```
 */
export class ServerExtendedAuditingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServerExtendedAuditingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions): ServerExtendedAuditingPolicy {
        return new ServerExtendedAuditingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy';

    /**
     * Returns true if the given object is an instance of ServerExtendedAuditingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerExtendedAuditingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerExtendedAuditingPolicy.__pulumiType;
    }

    /**
     * Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
     */
    public readonly logMonitoringEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    public readonly serverId!: pulumi.Output<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    public readonly storageAccountAccessKeyIsSecondary!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Subscription containing the Storage Account.
     */
    public readonly storageAccountSubscriptionId!: pulumi.Output<string | undefined>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    public readonly storageEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a ServerExtendedAuditingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerExtendedAuditingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerExtendedAuditingPolicyArgs | ServerExtendedAuditingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerExtendedAuditingPolicyState | undefined;
            resourceInputs["logMonitoringEnabled"] = state ? state.logMonitoringEnabled : undefined;
            resourceInputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountAccessKeyIsSecondary"] = state ? state.storageAccountAccessKeyIsSecondary : undefined;
            resourceInputs["storageAccountSubscriptionId"] = state ? state.storageAccountSubscriptionId : undefined;
            resourceInputs["storageEndpoint"] = state ? state.storageEndpoint : undefined;
        } else {
            const args = argsOrState as ServerExtendedAuditingPolicyArgs | undefined;
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            resourceInputs["logMonitoringEnabled"] = args ? args.logMonitoringEnabled : undefined;
            resourceInputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
            resourceInputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountAccessKeyIsSecondary"] = args ? args.storageAccountAccessKeyIsSecondary : undefined;
            resourceInputs["storageAccountSubscriptionId"] = args ? args.storageAccountSubscriptionId : undefined;
            resourceInputs["storageEndpoint"] = args ? args.storageEndpoint : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerExtendedAuditingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerExtendedAuditingPolicy resources.
 */
export interface ServerExtendedAuditingPolicyState {
    /**
     * Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
     */
    logMonitoringEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    serverId?: pulumi.Input<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The ID of the Subscription containing the Storage Account.
     */
    storageAccountSubscriptionId?: pulumi.Input<string>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerExtendedAuditingPolicy resource.
 */
export interface ServerExtendedAuditingPolicyArgs {
    /**
     * Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its main database audit events to Azure Monitor.
     */
    logMonitoringEnabled?: pulumi.Input<boolean>;
    /**
     * The number of days to retain logs for in the storage account.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
     */
    serverId: pulumi.Input<string>;
    /**
     * The access key to use for the auditing storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Is `storageAccountAccessKey` value the storage's secondary key?
     */
    storageAccountAccessKeyIsSecondary?: pulumi.Input<boolean>;
    /**
     * The ID of the Subscription containing the Storage Account.
     */
    storageAccountSubscriptionId?: pulumi.Input<string>;
    /**
     * The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}
