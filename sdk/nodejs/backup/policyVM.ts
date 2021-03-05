// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Backup VM Backup Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVault = new azure.recoveryservices.Vault("exampleVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "Standard",
 * });
 * const examplePolicyVM = new azure.backup.PolicyVM("examplePolicyVM", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     recoveryVaultName: exampleVault.name,
 *     timezone: "UTC",
 *     backup: {
 *         frequency: "Daily",
 *         time: "23:00",
 *     },
 *     retentionDaily: {
 *         count: 10,
 *     },
 *     retentionWeekly: {
 *         count: 42,
 *         weekdays: [
 *             "Sunday",
 *             "Wednesday",
 *             "Friday",
 *             "Saturday",
 *         ],
 *     },
 *     retentionMonthly: {
 *         count: 7,
 *         weekdays: [
 *             "Sunday",
 *             "Wednesday",
 *         ],
 *         weeks: [
 *             "First",
 *             "Last",
 *         ],
 *     },
 *     retentionYearly: {
 *         count: 77,
 *         weekdays: ["Sunday"],
 *         weeks: ["Last"],
 *         months: ["January"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * VM Backup Policies can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:backup/policyVM:PolicyVM policy1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
 * ```
 */
export class PolicyVM extends pulumi.CustomResource {
    /**
     * Get an existing PolicyVM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyVMState, opts?: pulumi.CustomResourceOptions): PolicyVM {
        return new PolicyVM(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:backup/policyVM:PolicyVM';

    /**
     * Returns true if the given object is an instance of PolicyVM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyVM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyVM.__pulumiType;
    }

    /**
     * Configures the Policy backup frequency, times & days as documented in the `backup` block below.
     */
    public readonly backup!: pulumi.Output<outputs.backup.PolicyVMBackup>;
    /**
     * Specifies the instant restore retention range in days.
     */
    public readonly instantRestoreRetentionDays!: pulumi.Output<number>;
    /**
     * Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    public readonly recoveryVaultName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
     */
    public readonly retentionDaily!: pulumi.Output<outputs.backup.PolicyVMRetentionDaily | undefined>;
    /**
     * Configures the policy monthly retention as documented in the `retentionMonthly` block below.
     */
    public readonly retentionMonthly!: pulumi.Output<outputs.backup.PolicyVMRetentionMonthly | undefined>;
    /**
     * Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
     */
    public readonly retentionWeekly!: pulumi.Output<outputs.backup.PolicyVMRetentionWeekly | undefined>;
    /**
     * Configures the policy yearly retention as documented in the `retentionYearly` block below.
     */
    public readonly retentionYearly!: pulumi.Output<outputs.backup.PolicyVMRetentionYearly | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
     */
    public readonly timezone!: pulumi.Output<string | undefined>;

    /**
     * Create a PolicyVM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyVMArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyVMArgs | PolicyVMState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyVMState | undefined;
            inputs["backup"] = state ? state.backup : undefined;
            inputs["instantRestoreRetentionDays"] = state ? state.instantRestoreRetentionDays : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recoveryVaultName"] = state ? state.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionDaily"] = state ? state.retentionDaily : undefined;
            inputs["retentionMonthly"] = state ? state.retentionMonthly : undefined;
            inputs["retentionWeekly"] = state ? state.retentionWeekly : undefined;
            inputs["retentionYearly"] = state ? state.retentionYearly : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timezone"] = state ? state.timezone : undefined;
        } else {
            const args = argsOrState as PolicyVMArgs | undefined;
            if ((!args || args.backup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backup'");
            }
            if ((!args || args.recoveryVaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recoveryVaultName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backup"] = args ? args.backup : undefined;
            inputs["instantRestoreRetentionDays"] = args ? args.instantRestoreRetentionDays : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recoveryVaultName"] = args ? args.recoveryVaultName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionDaily"] = args ? args.retentionDaily : undefined;
            inputs["retentionMonthly"] = args ? args.retentionMonthly : undefined;
            inputs["retentionWeekly"] = args ? args.retentionWeekly : undefined;
            inputs["retentionYearly"] = args ? args.retentionYearly : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timezone"] = args ? args.timezone : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PolicyVM.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyVM resources.
 */
export interface PolicyVMState {
    /**
     * Configures the Policy backup frequency, times & days as documented in the `backup` block below.
     */
    readonly backup?: pulumi.Input<inputs.backup.PolicyVMBackup>;
    /**
     * Specifies the instant restore retention range in days.
     */
    readonly instantRestoreRetentionDays?: pulumi.Input<number>;
    /**
     * Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    readonly recoveryVaultName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
     */
    readonly retentionDaily?: pulumi.Input<inputs.backup.PolicyVMRetentionDaily>;
    /**
     * Configures the policy monthly retention as documented in the `retentionMonthly` block below.
     */
    readonly retentionMonthly?: pulumi.Input<inputs.backup.PolicyVMRetentionMonthly>;
    /**
     * Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
     */
    readonly retentionWeekly?: pulumi.Input<inputs.backup.PolicyVMRetentionWeekly>;
    /**
     * Configures the policy yearly retention as documented in the `retentionYearly` block below.
     */
    readonly retentionYearly?: pulumi.Input<inputs.backup.PolicyVMRetentionYearly>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
     */
    readonly timezone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyVM resource.
 */
export interface PolicyVMArgs {
    /**
     * Configures the Policy backup frequency, times & days as documented in the `backup` block below.
     */
    readonly backup: pulumi.Input<inputs.backup.PolicyVMBackup>;
    /**
     * Specifies the instant restore retention range in days.
     */
    readonly instantRestoreRetentionDays?: pulumi.Input<number>;
    /**
     * Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
     */
    readonly recoveryVaultName: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
     */
    readonly retentionDaily?: pulumi.Input<inputs.backup.PolicyVMRetentionDaily>;
    /**
     * Configures the policy monthly retention as documented in the `retentionMonthly` block below.
     */
    readonly retentionMonthly?: pulumi.Input<inputs.backup.PolicyVMRetentionMonthly>;
    /**
     * Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
     */
    readonly retentionWeekly?: pulumi.Input<inputs.backup.PolicyVMRetentionWeekly>;
    /**
     * Configures the policy yearly retention as documented in the `retentionYearly` block below.
     */
    readonly retentionYearly?: pulumi.Input<inputs.backup.PolicyVMRetentionYearly>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the timezone. [the possible values are defined here](http://jackstromberg.com/2017/01/list-of-time-zones-consumed-by-azure/). Defaults to `UTC`
     */
    readonly timezone?: pulumi.Input<string>;
}
