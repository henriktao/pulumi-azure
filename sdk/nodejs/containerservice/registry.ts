// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Container Registry.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const acr = new azure.containerservice.Registry("acr", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     sku: "Premium",
 *     adminEnabled: false,
 *     georeplications: [
 *         {
 *             location: "East US",
 *             zoneRedundancyEnabled: true,
 *             tags: {},
 *         },
 *         {
 *             location: "westeurope",
 *             zoneRedundancyEnabled: true,
 *             tags: {},
 *         },
 *     ],
 * });
 * ```
 * ### Encryption)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const exampleUserAssignedIdentity = new azure.authorization.UserAssignedIdentity("exampleUserAssignedIdentity", {
 *     resourceGroupName: azurerm_resource_group.example.name,
 *     location: azurerm_resource_group.example.location,
 * });
 * const exampleKey = azure.keyvault.getKey({
 *     name: "super-secret",
 *     keyVaultId: data.azurerm_key_vault.existing.id,
 * });
 * const acr = new azure.containerservice.Registry("acr", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 *     sku: "Premium",
 *     identity: {
 *         type: "UserAssigned",
 *         identityIds: [exampleUserAssignedIdentity.id],
 *     },
 *     encryption: {
 *         enabled: true,
 *         keyVaultKeyId: exampleKey.then(exampleKey => exampleKey.id),
 *         identityClientId: exampleUserAssignedIdentity.clientId,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Container Registries can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:containerservice/registry:Registry example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1
 * ```
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryState, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/registry:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    public readonly adminEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Password associated with the Container Registry Admin account - if the admin account is enabled.
     */
    public /*out*/ readonly adminPassword!: pulumi.Output<string>;
    /**
     * The Username associated with the Container Registry Admin account - if the admin account is enabled.
     */
    public /*out*/ readonly adminUsername!: pulumi.Output<string>;
    /**
     * An `encryption` block as documented below.
     */
    public readonly encryption!: pulumi.Output<outputs.containerservice.RegistryEncryption>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     *
     * @deprecated Deprecated in favour of `georeplications`
     */
    public readonly georeplicationLocations!: pulumi.Output<string[]>;
    /**
     * A `georeplications` block as documented below.
     */
    public readonly georeplications!: pulumi.Output<outputs.containerservice.RegistryGeoreplication[]>;
    /**
     * An `identity` block as documented below.
     */
    public readonly identity!: pulumi.Output<outputs.containerservice.RegistryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The URL that can be used to log into the container registry.
     */
    public /*out*/ readonly loginServer!: pulumi.Output<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * Whether public network access is allowed for the container registry. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
     */
    public readonly quarantinePolicyEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.containerservice.RegistryRetentionPolicy>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * @deprecated this attribute is no longer recognized by the API and is not functional anymore, thus this property will be removed in v3.0
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `trustPolicy` block as documented below.
     */
    public readonly trustPolicy!: pulumi.Output<outputs.containerservice.RegistryTrustPolicy>;
    /**
     * Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to `false`.
     */
    public readonly zoneRedundancyEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs | RegistryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryState | undefined;
            inputs["adminEnabled"] = state ? state.adminEnabled : undefined;
            inputs["adminPassword"] = state ? state.adminPassword : undefined;
            inputs["adminUsername"] = state ? state.adminUsername : undefined;
            inputs["encryption"] = state ? state.encryption : undefined;
            inputs["georeplicationLocations"] = state ? state.georeplicationLocations : undefined;
            inputs["georeplications"] = state ? state.georeplications : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["loginServer"] = state ? state.loginServer : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkRuleSet"] = state ? state.networkRuleSet : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["quarantinePolicyEnabled"] = state ? state.quarantinePolicyEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionPolicy"] = state ? state.retentionPolicy : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["trustPolicy"] = state ? state.trustPolicy : undefined;
            inputs["zoneRedundancyEnabled"] = state ? state.zoneRedundancyEnabled : undefined;
        } else {
            const args = argsOrState as RegistryArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["adminEnabled"] = args ? args.adminEnabled : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["georeplicationLocations"] = args ? args.georeplicationLocations : undefined;
            inputs["georeplications"] = args ? args.georeplications : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["quarantinePolicyEnabled"] = args ? args.quarantinePolicyEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trustPolicy"] = args ? args.trustPolicy : undefined;
            inputs["zoneRedundancyEnabled"] = args ? args.zoneRedundancyEnabled : undefined;
            inputs["adminPassword"] = undefined /*out*/;
            inputs["adminUsername"] = undefined /*out*/;
            inputs["loginServer"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Registry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registry resources.
 */
export interface RegistryState {
    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    adminEnabled?: pulumi.Input<boolean>;
    /**
     * The Password associated with the Container Registry Admin account - if the admin account is enabled.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * The Username associated with the Container Registry Admin account - if the admin account is enabled.
     */
    adminUsername?: pulumi.Input<string>;
    /**
     * An `encryption` block as documented below.
     */
    encryption?: pulumi.Input<inputs.containerservice.RegistryEncryption>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     *
     * @deprecated Deprecated in favour of `georeplications`
     */
    georeplicationLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `georeplications` block as documented below.
     */
    georeplications?: pulumi.Input<pulumi.Input<inputs.containerservice.RegistryGeoreplication>[]>;
    /**
     * An `identity` block as documented below.
     */
    identity?: pulumi.Input<inputs.containerservice.RegistryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The URL that can be used to log into the container registry.
     */
    loginServer?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    networkRuleSet?: pulumi.Input<inputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * Whether public network access is allowed for the container registry. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
     */
    quarantinePolicyEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    retentionPolicy?: pulumi.Input<inputs.containerservice.RegistryRetentionPolicy>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    sku?: pulumi.Input<string>;
    /**
     * @deprecated this attribute is no longer recognized by the API and is not functional anymore, thus this property will be removed in v3.0
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `trustPolicy` block as documented below.
     */
    trustPolicy?: pulumi.Input<inputs.containerservice.RegistryTrustPolicy>;
    /**
     * Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to `false`.
     */
    zoneRedundancyEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * Specifies whether the admin user is enabled. Defaults to `false`.
     */
    adminEnabled?: pulumi.Input<boolean>;
    /**
     * An `encryption` block as documented below.
     */
    encryption?: pulumi.Input<inputs.containerservice.RegistryEncryption>;
    /**
     * A list of Azure locations where the container registry should be geo-replicated.
     *
     * @deprecated Deprecated in favour of `georeplications`
     */
    georeplicationLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `georeplications` block as documented below.
     */
    georeplications?: pulumi.Input<pulumi.Input<inputs.containerservice.RegistryGeoreplication>[]>;
    /**
     * An `identity` block as documented below.
     */
    identity?: pulumi.Input<inputs.containerservice.RegistryIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `networkRuleSet` block as documented below.
     */
    networkRuleSet?: pulumi.Input<inputs.containerservice.RegistryNetworkRuleSet>;
    /**
     * Whether public network access is allowed for the container registry. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * Boolean value that indicates whether quarantine policy is enabled. Defaults to `false`.
     */
    quarantinePolicyEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `retentionPolicy` block as documented below.
     */
    retentionPolicy?: pulumi.Input<inputs.containerservice.RegistryRetentionPolicy>;
    /**
     * The SKU name of the container registry. Possible values are  `Basic`, `Standard` and `Premium`. `Classic` (which was previously `Basic`) is supported only for existing resources.
     */
    sku?: pulumi.Input<string>;
    /**
     * @deprecated this attribute is no longer recognized by the API and is not functional anymore, thus this property will be removed in v3.0
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `trustPolicy` block as documented below.
     */
    trustPolicy?: pulumi.Input<inputs.containerservice.RegistryTrustPolicy>;
    /**
     * Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to `false`.
     */
    zoneRedundancyEnabled?: pulumi.Input<boolean>;
}
