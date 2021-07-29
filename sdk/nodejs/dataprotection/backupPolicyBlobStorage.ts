// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Backup Policy Blob Storage.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const exampleBackupVault = new azure.dataprotection.BackupVault("exampleBackupVault", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     datastoreType: "VaultStore",
 *     redundancy: "LocallyRedundant",
 * });
 * const exampleBackupPolicyBlobStorage = new azure.dataprotection.BackupPolicyBlobStorage("exampleBackupPolicyBlobStorage", {
 *     vaultId: exampleBackupVault.id,
 *     retentionDuration: "P30D",
 * });
 * ```
 *
 * ## Import
 *
 * Backup Policy Blob Storages can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:dataprotection/backupPolicyBlobStorage:BackupPolicyBlobStorage example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
 * ```
 */
export class BackupPolicyBlobStorage extends pulumi.CustomResource {
    /**
     * Get an existing BackupPolicyBlobStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPolicyBlobStorageState, opts?: pulumi.CustomResourceOptions): BackupPolicyBlobStorage {
        return new BackupPolicyBlobStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:dataprotection/backupPolicyBlobStorage:BackupPolicyBlobStorage';

    /**
     * Returns true if the given object is an instance of BackupPolicyBlobStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPolicyBlobStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPolicyBlobStorage.__pulumiType;
    }

    /**
     * The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    public readonly retentionDuration!: pulumi.Output<string>;
    /**
     * The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    public readonly vaultId!: pulumi.Output<string>;

    /**
     * Create a BackupPolicyBlobStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPolicyBlobStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPolicyBlobStorageArgs | BackupPolicyBlobStorageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupPolicyBlobStorageState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["retentionDuration"] = state ? state.retentionDuration : undefined;
            inputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as BackupPolicyBlobStorageArgs | undefined;
            if ((!args || args.retentionDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionDuration'");
            }
            if ((!args || args.vaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["retentionDuration"] = args ? args.retentionDuration : undefined;
            inputs["vaultId"] = args ? args.vaultId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackupPolicyBlobStorage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPolicyBlobStorage resources.
 */
export interface BackupPolicyBlobStorageState {
    /**
     * The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    retentionDuration?: pulumi.Input<string>;
    /**
     * The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupPolicyBlobStorage resource.
 */
export interface BackupPolicyBlobStorageArgs {
    /**
     * The name which should be used for this Backup Policy Blob Storage. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Duration of deletion after given timespan. It should follow `ISO 8601` duration format. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    retentionDuration: pulumi.Input<string>;
    /**
     * The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
     */
    vaultId: pulumi.Input<string>;
}
