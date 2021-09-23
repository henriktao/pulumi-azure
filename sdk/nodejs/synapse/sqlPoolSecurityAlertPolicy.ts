// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Security Alert Policy for a Synapse SQL Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     isHnsEnabled: "true",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleAccount.id});
 * const exampleWorkspace = new azure.synapse.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     storageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.id,
 *     sqlAdministratorLogin: "sqladminuser",
 *     sqlAdministratorLoginPassword: "H@Sh1CoR3!",
 *     aadAdmin: {
 *         login: "AzureAD Admin",
 *         objectId: "00000000-0000-0000-0000-000000000000",
 *         tenantId: "00000000-0000-0000-0000-000000000000",
 *     },
 *     tags: {
 *         Env: "production",
 *     },
 * });
 * const exampleSqlPool = new azure.synapse.SqlPool("exampleSqlPool", {
 *     synapseWorkspaceId: exampleWorkspace.id,
 *     skuName: "DW100c",
 *     createMode: "Default",
 * });
 * const auditLogs = new azure.storage.Account("auditLogs", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleSqlPoolSecurityAlertPolicy = new azure.synapse.SqlPoolSecurityAlertPolicy("exampleSqlPoolSecurityAlertPolicy", {
 *     sqlPoolId: exampleSqlPool.id,
 *     policyState: "Enabled",
 *     storageEndpoint: auditLogs.primaryBlobEndpoint,
 *     storageAccountAccessKey: auditLogs.primaryAccessKey,
 *     disabledAlerts: [
 *         "Sql_Injection",
 *         "Data_Exfiltration",
 *     ],
 *     retentionDays: 20,
 * });
 * ```
 *
 * ## Import
 *
 * Synapse SQL Pool Security Alert Policies can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:synapse/sqlPoolSecurityAlertPolicy:SqlPoolSecurityAlertPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/securityAlertPolicies/default
 * ```
 */
export class SqlPoolSecurityAlertPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SqlPoolSecurityAlertPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlPoolSecurityAlertPolicyState, opts?: pulumi.CustomResourceOptions): SqlPoolSecurityAlertPolicy {
        return new SqlPoolSecurityAlertPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/sqlPoolSecurityAlertPolicy:SqlPoolSecurityAlertPolicy';

    /**
     * Returns true if the given object is an instance of SqlPoolSecurityAlertPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlPoolSecurityAlertPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlPoolSecurityAlertPolicy.__pulumiType;
    }

    /**
     * Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
     */
    public readonly disabledAlerts!: pulumi.Output<string[] | undefined>;
    /**
     * Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
     */
    public readonly emailAccountAdminsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies an array of e-mail addresses to which the alert is sent.
     */
    public readonly emailAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific SQL pool. Allowed values are: `Disabled`, `Enabled`.
     */
    public readonly policyState!: pulumi.Output<string>;
    /**
     * Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    /**
     * Specifies the ID of the Synapse SQL Pool. Changing this forces a new resource to be created.
     */
    public readonly sqlPoolId!: pulumi.Output<string>;
    /**
     * Specifies the identifier key of the Threat Detection audit storage account.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
     */
    public readonly storageEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a SqlPoolSecurityAlertPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlPoolSecurityAlertPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlPoolSecurityAlertPolicyArgs | SqlPoolSecurityAlertPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SqlPoolSecurityAlertPolicyState | undefined;
            inputs["disabledAlerts"] = state ? state.disabledAlerts : undefined;
            inputs["emailAccountAdminsEnabled"] = state ? state.emailAccountAdminsEnabled : undefined;
            inputs["emailAddresses"] = state ? state.emailAddresses : undefined;
            inputs["policyState"] = state ? state.policyState : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
            inputs["sqlPoolId"] = state ? state.sqlPoolId : undefined;
            inputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            inputs["storageEndpoint"] = state ? state.storageEndpoint : undefined;
        } else {
            const args = argsOrState as SqlPoolSecurityAlertPolicyArgs | undefined;
            if ((!args || args.policyState === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyState'");
            }
            if ((!args || args.sqlPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlPoolId'");
            }
            inputs["disabledAlerts"] = args ? args.disabledAlerts : undefined;
            inputs["emailAccountAdminsEnabled"] = args ? args.emailAccountAdminsEnabled : undefined;
            inputs["emailAddresses"] = args ? args.emailAddresses : undefined;
            inputs["policyState"] = args ? args.policyState : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
            inputs["sqlPoolId"] = args ? args.sqlPoolId : undefined;
            inputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            inputs["storageEndpoint"] = args ? args.storageEndpoint : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SqlPoolSecurityAlertPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SqlPoolSecurityAlertPolicy resources.
 */
export interface SqlPoolSecurityAlertPolicyState {
    /**
     * Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
     */
    disabledAlerts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
     */
    emailAccountAdminsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies an array of e-mail addresses to which the alert is sent.
     */
    emailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific SQL pool. Allowed values are: `Disabled`, `Enabled`.
     */
    policyState?: pulumi.Input<string>;
    /**
     * Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * Specifies the ID of the Synapse SQL Pool. Changing this forces a new resource to be created.
     */
    sqlPoolId?: pulumi.Input<string>;
    /**
     * Specifies the identifier key of the Threat Detection audit storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SqlPoolSecurityAlertPolicy resource.
 */
export interface SqlPoolSecurityAlertPolicyArgs {
    /**
     * Specifies an array of alerts that are disabled. Allowed values are: `Sql_Injection`, `Sql_Injection_Vulnerability`, `Access_Anomaly`, `Data_Exfiltration`, `Unsafe_Action`.
     */
    disabledAlerts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to `false`.
     */
    emailAccountAdminsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies an array of e-mail addresses to which the alert is sent.
     */
    emailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific SQL pool. Allowed values are: `Disabled`, `Enabled`.
     */
    policyState: pulumi.Input<string>;
    /**
     * Specifies the number of days to keep in the Threat Detection audit logs. Defaults to `0`.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * Specifies the ID of the Synapse SQL Pool. Changing this forces a new resource to be created.
     */
    sqlPoolId: pulumi.Input<string>;
    /**
     * Specifies the identifier key of the Threat Detection audit storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}
