// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Storage Object Replication.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const srcResourceGroup = new azure.core.ResourceGroup("srcResourceGroup", {location: "West Europe"});
 * const srcAccount = new azure.storage.Account("srcAccount", {
 *     resourceGroupName: srcResourceGroup.name,
 *     location: srcResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     blobProperties: {
 *         versioningEnabled: true,
 *         changeFeedEnabled: true,
 *     },
 * });
 * const srcContainer = new azure.storage.Container("srcContainer", {
 *     storageAccountName: srcAccount.name,
 *     containerAccessType: "private",
 * });
 * const dstResourceGroup = new azure.core.ResourceGroup("dstResourceGroup", {location: "East US"});
 * const dstAccount = new azure.storage.Account("dstAccount", {
 *     resourceGroupName: dstResourceGroup.name,
 *     location: dstResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     blobProperties: {
 *         versioningEnabled: true,
 *         changeFeedEnabled: true,
 *     },
 * });
 * const dstContainer = new azure.storage.Container("dstContainer", {
 *     storageAccountName: dstAccount.name,
 *     containerAccessType: "private",
 * });
 * const example = new azure.storage.ObjectReplication("example", {
 *     sourceStorageAccountId: srcAccount.id,
 *     destinationStorageAccountId: dstAccount.id,
 *     rules: [{
 *         sourceContainerName: srcContainer.name,
 *         destinationContainerName: dstContainer.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Storage Object Replication Policies can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/objectReplication:ObjectReplication example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/storageAccount1/objectReplicationPolicies/objectReplicationPolicy1;/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group2/providers/Microsoft.Storage/storageAccounts/storageAccount2/objectReplicationPolicies/objectReplicationPolicy2
 * ```
 */
export class ObjectReplication extends pulumi.CustomResource {
    /**
     * Get an existing ObjectReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectReplicationState, opts?: pulumi.CustomResourceOptions): ObjectReplication {
        return new ObjectReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/objectReplication:ObjectReplication';

    /**
     * Returns true if the given object is an instance of ObjectReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectReplication.__pulumiType;
    }

    /**
     * The ID of the Object Replication in the destination storage account.
     */
    public /*out*/ readonly destinationObjectReplicationId!: pulumi.Output<string>;
    /**
     * The ID of the destination storage account. Changing this forces a new Storage Object Replication to be created.
     */
    public readonly destinationStorageAccountId!: pulumi.Output<string>;
    /**
     * One or more `rules` blocks as defined below.
     */
    public readonly rules!: pulumi.Output<outputs.storage.ObjectReplicationRule[]>;
    /**
     * The ID of the Object Replication in the source storage account.
     */
    public /*out*/ readonly sourceObjectReplicationId!: pulumi.Output<string>;
    /**
     * The ID of the source storage account. Changing this forces a new Storage Object Replication to be created.
     */
    public readonly sourceStorageAccountId!: pulumi.Output<string>;

    /**
     * Create a ObjectReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectReplicationArgs | ObjectReplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectReplicationState | undefined;
            inputs["destinationObjectReplicationId"] = state ? state.destinationObjectReplicationId : undefined;
            inputs["destinationStorageAccountId"] = state ? state.destinationStorageAccountId : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["sourceObjectReplicationId"] = state ? state.sourceObjectReplicationId : undefined;
            inputs["sourceStorageAccountId"] = state ? state.sourceStorageAccountId : undefined;
        } else {
            const args = argsOrState as ObjectReplicationArgs | undefined;
            if ((!args || args.destinationStorageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationStorageAccountId'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.sourceStorageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceStorageAccountId'");
            }
            inputs["destinationStorageAccountId"] = args ? args.destinationStorageAccountId : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["sourceStorageAccountId"] = args ? args.sourceStorageAccountId : undefined;
            inputs["destinationObjectReplicationId"] = undefined /*out*/;
            inputs["sourceObjectReplicationId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ObjectReplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectReplication resources.
 */
export interface ObjectReplicationState {
    /**
     * The ID of the Object Replication in the destination storage account.
     */
    destinationObjectReplicationId?: pulumi.Input<string>;
    /**
     * The ID of the destination storage account. Changing this forces a new Storage Object Replication to be created.
     */
    destinationStorageAccountId?: pulumi.Input<string>;
    /**
     * One or more `rules` blocks as defined below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.storage.ObjectReplicationRule>[]>;
    /**
     * The ID of the Object Replication in the source storage account.
     */
    sourceObjectReplicationId?: pulumi.Input<string>;
    /**
     * The ID of the source storage account. Changing this forces a new Storage Object Replication to be created.
     */
    sourceStorageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectReplication resource.
 */
export interface ObjectReplicationArgs {
    /**
     * The ID of the destination storage account. Changing this forces a new Storage Object Replication to be created.
     */
    destinationStorageAccountId: pulumi.Input<string>;
    /**
     * One or more `rules` blocks as defined below.
     */
    rules: pulumi.Input<pulumi.Input<inputs.storage.ObjectReplicationRule>[]>;
    /**
     * The ID of the source storage account. Changing this forces a new Storage Object Replication to be created.
     */
    sourceStorageAccountId: pulumi.Input<string>;
}
