// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the transparent data encryption configuration for a MSSQL Server
 *
 * > **NOTE:** Once transparent data encryption is enabled on a MS SQL instance, it is not possible to remove TDE. You will be able to switch between 'ServiceManaged' and 'CustomerManaged' keys, but will not be able to remove encryption. For safety when this resource is deleted, the TDE mode will automatically be set to 'ServiceManaged'. See `keyVaultUri` for more information on how to specify the key types. As SQL Server only supports a single configuration for encryption settings, this resource will replace the current encryption settings on the server.
 *
 * > **Note:** See [documentation](https://docs.microsoft.com/en-us/azure/azure-sql/database/transparent-data-encryption-byok-overview) for important information on how handle lifecycle management of the keys to prevent data lockout.
 *
 * ## Example Usage
 * ### With Service Managed Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "EastUs"});
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "thisIsKat11",
 *     minimumTlsVersion: "1.2",
 *     azureadAdministrator: {
 *         loginUsername: "AzureAD Admin",
 *         objectId: "00000000-0000-0000-0000-000000000000",
 *     },
 *     extendedAuditingPolicy: {
 *         storageEndpoint: azurerm_storage_account.example.primary_blob_endpoint,
 *         storageAccountAccessKey: azurerm_storage_account.example.primary_access_key,
 *         storageAccountAccessKeyIsSecondary: true,
 *         retentionInDays: 6,
 *     },
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * const exampleServerTransparentDataEncryption = new azure.mssql.ServerTransparentDataEncryption("exampleServerTransparentDataEncryption", {serverId: exampleServer.id});
 * ```
 * ### With Customer Managed Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "EastUs"});
 * const exampleServer = new azure.mssql.Server("exampleServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     version: "12.0",
 *     administratorLogin: "missadministrator",
 *     administratorLoginPassword: "thisIsKat11",
 *     minimumTlsVersion: "1.2",
 *     azureadAdministrator: {
 *         loginUsername: "AzureAD Admin",
 *         objectId: "00000000-0000-0000-0000-000000000000",
 *     },
 *     extendedAuditingPolicy: {
 *         storageEndpoint: azurerm_storage_account.example.primary_blob_endpoint,
 *         storageAccountAccessKey: azurerm_storage_account.example.primary_access_key,
 *         storageAccountAccessKeyIsSecondary: true,
 *         retentionInDays: 6,
 *     },
 *     tags: {
 *         environment: "production",
 *     },
 *     identity: {
 *         type: "SystemAssigned",
 *     },
 * });
 * // Create a key vault with policies for the deployer to create a key & SQL Server to wrap/unwrap/get key
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     enabledForDiskEncryption: true,
 *     tenantId: current.then(current => current.tenantId),
 *     softDeleteRetentionDays: 7,
 *     purgeProtectionEnabled: false,
 *     skuName: "standard",
 *     accessPolicies: [
 *         {
 *             tenantId: current.then(current => current.tenantId),
 *             objectId: current.then(current => current.objectId),
 *             keyPermissions: [
 *                 "Get",
 *                 "List",
 *                 "Create",
 *                 "Delete",
 *                 "Update",
 *                 "Recover",
 *                 "Purge",
 *             ],
 *         },
 *         {
 *             tenantId: exampleServer.identity.apply(identity => identity?.tenantId),
 *             objectId: exampleServer.identity.apply(identity => identity?.principalId),
 *             keyPermissions: [
 *                 "Get",
 *                 "WrapKey",
 *                 "UnwrapKey",
 *             ],
 *         },
 *     ],
 * });
 * const exampleKey = new azure.keyvault.Key("exampleKey", {
 *     keyVaultId: exampleKeyVault.id,
 *     keyType: "RSA",
 *     keySize: 2048,
 *     keyOpts: [
 *         "unwrapKey",
 *         "wrapKey",
 *     ],
 * }, {
 *     dependsOn: [exampleKeyVault],
 * });
 * const exampleServerTransparentDataEncryption = new azure.mssql.ServerTransparentDataEncryption("exampleServerTransparentDataEncryption", {
 *     serverId: exampleServer.id,
 *     keyVaultKeyId: exampleKey.id,
 * });
 * ```
 *
 * ## Import
 *
 * SQL Server Transparent Data Encryption can be imported using the resource id, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption example /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/encryptionProtector/current
 * ```
 */
export class ServerTransparentDataEncryption extends pulumi.CustomResource {
    /**
     * Get an existing ServerTransparentDataEncryption resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerTransparentDataEncryptionState, opts?: pulumi.CustomResourceOptions): ServerTransparentDataEncryption {
        return new ServerTransparentDataEncryption(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/serverTransparentDataEncryption:ServerTransparentDataEncryption';

    /**
     * Returns true if the given object is an instance of ServerTransparentDataEncryption.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerTransparentDataEncryption {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerTransparentDataEncryption.__pulumiType;
    }

    /**
     * To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
     */
    public readonly keyVaultKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the MS SQL Server.
     */
    public readonly serverId!: pulumi.Output<string>;

    /**
     * Create a ServerTransparentDataEncryption resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerTransparentDataEncryptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerTransparentDataEncryptionArgs | ServerTransparentDataEncryptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerTransparentDataEncryptionState | undefined;
            resourceInputs["keyVaultKeyId"] = state ? state.keyVaultKeyId : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
        } else {
            const args = argsOrState as ServerTransparentDataEncryptionArgs | undefined;
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            resourceInputs["keyVaultKeyId"] = args ? args.keyVaultKeyId : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerTransparentDataEncryption.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerTransparentDataEncryption resources.
 */
export interface ServerTransparentDataEncryptionState {
    /**
     * To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the name of the MS SQL Server.
     */
    serverId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerTransparentDataEncryption resource.
 */
export interface ServerTransparentDataEncryptionArgs {
    /**
     * To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
     */
    keyVaultKeyId?: pulumi.Input<string>;
    /**
     * Specifies the name of the MS SQL Server.
     */
    serverId: pulumi.Input<string>;
}
