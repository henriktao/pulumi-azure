// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Customer Managed Key for a Storage Account.
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
 *     skuName: "standard",
 *     purgeProtectionEnabled: true,
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "GRS",
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * const storage = new azure.keyvault.AccessPolicy("storage", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: exampleAccount.identity.apply(identity => identity?.principalId),
 *     keyPermissions: [
 *         "get",
 *         "create",
 *         "list",
 *         "restore",
 *         "recover",
 *         "unwrapkey",
 *         "wrapkey",
 *         "purge",
 *         "encrypt",
 *         "decrypt",
 *         "sign",
 *         "verify",
 *     ],
 *     secretPermissions: ["get"],
 * });
 * const client = new azure.keyvault.AccessPolicy("client", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 *     keyPermissions: [
 *         "get",
 *         "create",
 *         "delete",
 *         "list",
 *         "restore",
 *         "recover",
 *         "unwrapkey",
 *         "wrapkey",
 *         "purge",
 *         "encrypt",
 *         "decrypt",
 *         "sign",
 *         "verify",
 *     ],
 *     secretPermissions: ["get"],
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
 *     dependsOn: [
 *         client,
 *         storage,
 *     ],
 * });
 * const exampleCustomerManagedKey = new azure.storage.CustomerManagedKey("exampleCustomerManagedKey", {
 *     storageAccountId: exampleAccount.id,
 *     keyVaultId: exampleKeyVault.id,
 *     keyName: exampleKey.name,
 * });
 * ```
 *
 * ## Import
 *
 * Customer Managed Keys for a Storage Account can be imported using the `resource id` of the Storage Account, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/customerManagedKey:CustomerManagedKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
 * ```
 */
export class CustomerManagedKey extends pulumi.CustomResource {
    /**
     * Get an existing CustomerManagedKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomerManagedKeyState, opts?: pulumi.CustomResourceOptions): CustomerManagedKey {
        return new CustomerManagedKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/customerManagedKey:CustomerManagedKey';

    /**
     * Returns true if the given object is an instance of CustomerManagedKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerManagedKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerManagedKey.__pulumiType;
    }

    /**
     * The name of Key Vault Key.
     */
    public readonly keyName!: pulumi.Output<string>;
    /**
     * The ID of the Key Vault. Changing this forces a new resource to be created.
     */
    public readonly keyVaultId!: pulumi.Output<string>;
    /**
     * The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
     */
    public readonly keyVersion!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Storage Account. Changing this forces a new resource to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * The ID of a user assigned identity.
     */
    public readonly userAssignedIdentityId!: pulumi.Output<string | undefined>;

    /**
     * Create a CustomerManagedKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerManagedKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerManagedKeyArgs | CustomerManagedKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomerManagedKeyState | undefined;
            resourceInputs["keyName"] = state ? state.keyName : undefined;
            resourceInputs["keyVaultId"] = state ? state.keyVaultId : undefined;
            resourceInputs["keyVersion"] = state ? state.keyVersion : undefined;
            resourceInputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            resourceInputs["userAssignedIdentityId"] = state ? state.userAssignedIdentityId : undefined;
        } else {
            const args = argsOrState as CustomerManagedKeyArgs | undefined;
            if ((!args || args.keyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyName'");
            }
            if ((!args || args.keyVaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyVaultId'");
            }
            if ((!args || args.storageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            resourceInputs["keyName"] = args ? args.keyName : undefined;
            resourceInputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            resourceInputs["keyVersion"] = args ? args.keyVersion : undefined;
            resourceInputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            resourceInputs["userAssignedIdentityId"] = args ? args.userAssignedIdentityId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomerManagedKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomerManagedKey resources.
 */
export interface CustomerManagedKeyState {
    /**
     * The name of Key Vault Key.
     */
    keyName?: pulumi.Input<string>;
    /**
     * The ID of the Key Vault. Changing this forces a new resource to be created.
     */
    keyVaultId?: pulumi.Input<string>;
    /**
     * The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
     */
    keyVersion?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account. Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * The ID of a user assigned identity.
     */
    userAssignedIdentityId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomerManagedKey resource.
 */
export interface CustomerManagedKeyArgs {
    /**
     * The name of Key Vault Key.
     */
    keyName: pulumi.Input<string>;
    /**
     * The ID of the Key Vault. Changing this forces a new resource to be created.
     */
    keyVaultId: pulumi.Input<string>;
    /**
     * The version of Key Vault Key. Remove or omit this argument to enable Automatic Key Rotation.
     */
    keyVersion?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account. Changing this forces a new resource to be created.
     */
    storageAccountId: pulumi.Input<string>;
    /**
     * The ID of a user assigned identity.
     */
    userAssignedIdentityId?: pulumi.Input<string>;
}
