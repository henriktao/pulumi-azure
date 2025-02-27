// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Disk Encryption Set.
 *
 * > **NOTE:** At this time the Key Vault used to store the Active Key for this Disk Encryption Set must have both Soft Delete & Purge Protection enabled - which are not yet supported by this provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: current.then(current => current.tenantId),
 *     skuName: "premium",
 *     enabledForDiskEncryption: true,
 *     softDeleteEnabled: true,
 *     purgeProtectionEnabled: true,
 * });
 * const example_user = new azure.keyvault.AccessPolicy("example-user", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 *     keyPermissions: [
 *         "get",
 *         "create",
 *         "delete",
 *     ],
 * });
 * const exampleKey = new azure.keyvault.Key("exampleKey", {
 *     keyVaultId: exampleKeyVault.id,
 *     keyType: "RSA",
 *     keySize: 2048,
 *     keyOpts: [
 *         "decrypt",
 *         "encrypt",
 *         "sign",
 *         "unwrapKey",
 *         "verify",
 *         "wrapKey",
 *     ],
 * }, {
 *     dependsOn: [example_user],
 * });
 * const exampleDiskEncryptionSet = new azure.compute.DiskEncryptionSet("exampleDiskEncryptionSet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     keyVaultKeyId: exampleKey.id,
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const example_disk = new azure.keyvault.AccessPolicy("example-disk", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: exampleDiskEncryptionSet.identity.apply(identity => identity.tenantId),
 *     objectId: exampleDiskEncryptionSet.identity.apply(identity => identity.principalId),
 *     keyPermissions: [
 *         "Get",
 *         "WrapKey",
 *         "UnwrapKey",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Disk Encryption Sets can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/diskEncryptionSet:DiskEncryptionSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
 * ```
 */
export class DiskEncryptionSet extends pulumi.CustomResource {
    /**
     * Get an existing DiskEncryptionSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskEncryptionSetState, opts?: pulumi.CustomResourceOptions): DiskEncryptionSet {
        return new DiskEncryptionSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/diskEncryptionSet:DiskEncryptionSet';

    /**
     * Returns true if the given object is an instance of DiskEncryptionSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskEncryptionSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskEncryptionSet.__pulumiType;
    }

    /**
     * Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
     */
    public readonly autoKeyRotationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
     */
    public readonly encryptionType!: pulumi.Output<string | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    public readonly keyVaultKeyId!: pulumi.Output<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DiskEncryptionSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskEncryptionSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskEncryptionSetArgs | DiskEncryptionSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskEncryptionSetState | undefined;
            resourceInputs["autoKeyRotationEnabled"] = state ? state.autoKeyRotationEnabled : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DiskEncryptionSetArgs | undefined;
            if ((!args || args.identity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identity'");
            }
            if ((!args || args.keyVaultKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyVaultKeyId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["autoKeyRotationEnabled"] = args ? args.autoKeyRotationEnabled : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DiskEncryptionSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskEncryptionSet resources.
 */
export interface DiskEncryptionSetState {
    /**
     * Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
     */
    autoKeyRotationEnabled?: pulumi.Input<boolean>;
    /**
     * The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DiskEncryptionSet resource.
 */
export interface DiskEncryptionSetArgs {
    /**
     * Boolean flag to specify whether Azure Disk Encryption Set automatically rotates encryption Key to latest version. Defaults to `false`.
     */
    autoKeyRotationEnabled?: pulumi.Input<boolean>;
    /**
     * The type of key used to encrypt the data of the disk. Possible values are `EncryptionAtRestWithCustomerKey` and `EncryptionAtRestWithPlatformAndCustomerKeys`. Defaults to `EncryptionAtRestWithCustomerKey`.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    identity: pulumi.Input<inputs.compute.DiskEncryptionSetIdentity>;
    /**
     * Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
     */
    keyVaultKeyId: pulumi.Input<string>;
    /**
     * Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Disk Encryption Set. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the Disk Encryption Set.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
