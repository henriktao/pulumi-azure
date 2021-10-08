// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Key Vault Managed Storage Account SAS Definition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleClientConfig = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleAccountSAS = exampleAccount.primaryConnectionString.apply(primaryConnectionString => azure.storage.getAccountSAS({
 *     connectionString: primaryConnectionString,
 *     httpsOnly: true,
 *     resourceTypes: {
 *         service: true,
 *         container: false,
 *         object: false,
 *     },
 *     services: {
 *         blob: true,
 *         queue: false,
 *         table: false,
 *         file: false,
 *     },
 *     start: "2021-04-30T00:00:00Z",
 *     expiry: "2023-04-30T00:00:00Z",
 *     permissions: {
 *         read: true,
 *         write: true,
 *         "delete": false,
 *         list: false,
 *         add: true,
 *         create: true,
 *         update: false,
 *         process: false,
 *     },
 * }));
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: data.azurerm_client_config.current.tenant_id,
 *     skuName: "standard",
 *     accessPolicies: [{
 *         tenantId: data.azurerm_client_config.current.tenant_id,
 *         objectId: data.azurerm_client_config.current.object_id,
 *         secretPermissions: [
 *             "Get",
 *             "Delete",
 *         ],
 *         storagePermissions: [
 *             "Get",
 *             "List",
 *             "Set",
 *             "SetSAS",
 *             "GetSAS",
 *             "DeleteSAS",
 *             "Update",
 *             "RegenerateKey",
 *         ],
 *     }],
 * });
 * const test = new azure.keyvault.ManagedStorageAccount("test", {
 *     keyVaultId: exampleKeyVault.id,
 *     storageAccountId: exampleAccount.id,
 *     storageAccountKey: "key1",
 *     regenerateKeyAutomatically: false,
 *     regenerationPeriod: "P1D",
 * });
 * const exampleManagedStorageAccountSasTokenDefinition = new azure.keyvault.ManagedStorageAccountSasTokenDefinition("exampleManagedStorageAccountSasTokenDefinition", {
 *     validityPeriod: "P1D",
 *     managedStorageAccountId: azurerm_key_vault_managed_storage_account.example.id,
 *     sasTemplateUri: exampleAccountSAS.apply(exampleAccountSAS => exampleAccountSAS.sas),
 *     sasType: "account",
 * });
 * ```
 *
 * ## Import
 *
 * Key Vaults can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:keyvault/managedStorageAccountSasTokenDefinition:ManagedStorageAccountSasTokenDefinition example https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01/sas/exampleSasDefinition01
 * ```
 */
export class ManagedStorageAccountSasTokenDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ManagedStorageAccountSasTokenDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedStorageAccountSasTokenDefinitionState, opts?: pulumi.CustomResourceOptions): ManagedStorageAccountSasTokenDefinition {
        return new ManagedStorageAccountSasTokenDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:keyvault/managedStorageAccountSasTokenDefinition:ManagedStorageAccountSasTokenDefinition';

    /**
     * Returns true if the given object is an instance of ManagedStorageAccountSasTokenDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedStorageAccountSasTokenDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedStorageAccountSasTokenDefinition.__pulumiType;
    }

    /**
     * The ID of the Managed Storage Account.
     */
    public readonly managedStorageAccountId!: pulumi.Output<string>;
    /**
     * The name which should be used for this SAS Definition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The SAS definition token template signed with an arbitrary key. Tokens created according to the SAS definition will have the same properties as the template, but regenerated with a new validity period.
     */
    public readonly sasTemplateUri!: pulumi.Output<string>;
    /**
     * The type of SAS token the SAS definition will create. Possible values are `account` and `service`.
     */
    public readonly sasType!: pulumi.Output<string>;
    /**
     * The ID of the Secret that is created by Managed Storage Account SAS Definition.
     */
    public /*out*/ readonly secretId!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the SAS Definition.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Validity period of SAS token. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
     */
    public readonly validityPeriod!: pulumi.Output<string>;

    /**
     * Create a ManagedStorageAccountSasTokenDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedStorageAccountSasTokenDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedStorageAccountSasTokenDefinitionArgs | ManagedStorageAccountSasTokenDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedStorageAccountSasTokenDefinitionState | undefined;
            inputs["managedStorageAccountId"] = state ? state.managedStorageAccountId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sasTemplateUri"] = state ? state.sasTemplateUri : undefined;
            inputs["sasType"] = state ? state.sasType : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["validityPeriod"] = state ? state.validityPeriod : undefined;
        } else {
            const args = argsOrState as ManagedStorageAccountSasTokenDefinitionArgs | undefined;
            if ((!args || args.managedStorageAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedStorageAccountId'");
            }
            if ((!args || args.sasTemplateUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sasTemplateUri'");
            }
            if ((!args || args.sasType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sasType'");
            }
            if ((!args || args.validityPeriod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validityPeriod'");
            }
            inputs["managedStorageAccountId"] = args ? args.managedStorageAccountId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sasTemplateUri"] = args ? args.sasTemplateUri : undefined;
            inputs["sasType"] = args ? args.sasType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["validityPeriod"] = args ? args.validityPeriod : undefined;
            inputs["secretId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ManagedStorageAccountSasTokenDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedStorageAccountSasTokenDefinition resources.
 */
export interface ManagedStorageAccountSasTokenDefinitionState {
    /**
     * The ID of the Managed Storage Account.
     */
    managedStorageAccountId?: pulumi.Input<string>;
    /**
     * The name which should be used for this SAS Definition.
     */
    name?: pulumi.Input<string>;
    /**
     * The SAS definition token template signed with an arbitrary key. Tokens created according to the SAS definition will have the same properties as the template, but regenerated with a new validity period.
     */
    sasTemplateUri?: pulumi.Input<string>;
    /**
     * The type of SAS token the SAS definition will create. Possible values are `account` and `service`.
     */
    sasType?: pulumi.Input<string>;
    /**
     * The ID of the Secret that is created by Managed Storage Account SAS Definition.
     */
    secretId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the SAS Definition.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Validity period of SAS token. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
     */
    validityPeriod?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedStorageAccountSasTokenDefinition resource.
 */
export interface ManagedStorageAccountSasTokenDefinitionArgs {
    /**
     * The ID of the Managed Storage Account.
     */
    managedStorageAccountId: pulumi.Input<string>;
    /**
     * The name which should be used for this SAS Definition.
     */
    name?: pulumi.Input<string>;
    /**
     * The SAS definition token template signed with an arbitrary key. Tokens created according to the SAS definition will have the same properties as the template, but regenerated with a new validity period.
     */
    sasTemplateUri: pulumi.Input<string>;
    /**
     * The type of SAS token the SAS definition will create. Possible values are `account` and `service`.
     */
    sasType: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the SAS Definition.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Validity period of SAS token. Value needs to be in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
     */
    validityPeriod: pulumi.Input<string>;
}
