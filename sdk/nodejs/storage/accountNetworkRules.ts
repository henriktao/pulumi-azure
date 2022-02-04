// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages network rules inside of a Azure Storage Account.
 *
 * > **NOTE:** Network Rules can be defined either directly on the `azure.storage.Account` resource, or using the `azure.storage.AccountNetworkRules` resource - but the two cannot be used together. Spurious changes will occur if both are used against the same Storage Account.
 *
 * > **NOTE:** Only one `azure.storage.AccountNetworkRules` can be tied to an `azure.storage.Account`. Spurious changes will occur if more than `azure.storage.AccountNetworkRules` is tied to the same `azure.storage.Account`.
 *
 * > **NOTE:** Deleting this resource updates the storage account back to the default values it had when the storage account was created.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
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
 *     serviceEndpoints: ["Microsoft.Storage"],
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "GRS",
 *     tags: {
 *         environment: "staging",
 *     },
 * });
 * const test = new azure.storage.AccountNetworkRules("test", {
 *     storageAccountId: azurerm_storage_account.test.id,
 *     defaultAction: "Allow",
 *     ipRules: ["127.0.0.1"],
 *     virtualNetworkSubnetIds: [azurerm_subnet.test.id],
 *     bypasses: ["Metrics"],
 * });
 * ```
 *
 * ## Import
 *
 * Storage Account Network Rules can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:storage/accountNetworkRules:AccountNetworkRules storageAcc1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
 * ```
 */
export class AccountNetworkRules extends pulumi.CustomResource {
    /**
     * Get an existing AccountNetworkRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountNetworkRulesState, opts?: pulumi.CustomResourceOptions): AccountNetworkRules {
        return new AccountNetworkRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/accountNetworkRules:AccountNetworkRules';

    /**
     * Returns true if the given object is an instance of AccountNetworkRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountNetworkRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountNetworkRules.__pulumiType;
    }

    /**
     * Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
     */
    public readonly bypasses!: pulumi.Output<string[]>;
    /**
     * Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
     */
    public readonly defaultAction!: pulumi.Output<string>;
    /**
     * List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
     */
    public readonly ipRules!: pulumi.Output<string[]>;
    /**
     * One or More `privateLinkAccess` block as defined below.
     */
    public readonly privateLinkAccessRules!: pulumi.Output<outputs.storage.AccountNetworkRulesPrivateLinkAccessRule[] | undefined>;
    /**
     * The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the ID of the storage account. Changing this forces a new resource to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;
    /**
     * Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * A list of virtual network subnet ids to to secure the storage account.
     */
    public readonly virtualNetworkSubnetIds!: pulumi.Output<string[]>;

    /**
     * Create a AccountNetworkRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountNetworkRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountNetworkRulesArgs | AccountNetworkRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountNetworkRulesState | undefined;
            resourceInputs["bypasses"] = state ? state.bypasses : undefined;
            resourceInputs["defaultAction"] = state ? state.defaultAction : undefined;
            resourceInputs["ipRules"] = state ? state.ipRules : undefined;
            resourceInputs["privateLinkAccessRules"] = state ? state.privateLinkAccessRules : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["storageAccountId"] = state ? state.storageAccountId : undefined;
            resourceInputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            resourceInputs["virtualNetworkSubnetIds"] = state ? state.virtualNetworkSubnetIds : undefined;
        } else {
            const args = argsOrState as AccountNetworkRulesArgs | undefined;
            if ((!args || args.defaultAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultAction'");
            }
            resourceInputs["bypasses"] = args ? args.bypasses : undefined;
            resourceInputs["defaultAction"] = args ? args.defaultAction : undefined;
            resourceInputs["ipRules"] = args ? args.ipRules : undefined;
            resourceInputs["privateLinkAccessRules"] = args ? args.privateLinkAccessRules : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            resourceInputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            resourceInputs["virtualNetworkSubnetIds"] = args ? args.virtualNetworkSubnetIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountNetworkRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountNetworkRules resources.
 */
export interface AccountNetworkRulesState {
    /**
     * Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
     */
    bypasses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
     */
    defaultAction?: pulumi.Input<string>;
    /**
     * List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
     */
    ipRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or More `privateLinkAccess` block as defined below.
     */
    privateLinkAccessRules?: pulumi.Input<pulumi.Input<inputs.storage.AccountNetworkRulesPrivateLinkAccessRule>[]>;
    /**
     * The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the ID of the storage account. Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    storageAccountName?: pulumi.Input<string>;
    /**
     * A list of virtual network subnet ids to to secure the storage account.
     */
    virtualNetworkSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AccountNetworkRules resource.
 */
export interface AccountNetworkRulesArgs {
    /**
     * Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
     */
    bypasses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
     */
    defaultAction: pulumi.Input<string>;
    /**
     * List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
     */
    ipRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or More `privateLinkAccess` block as defined below.
     */
    privateLinkAccessRules?: pulumi.Input<pulumi.Input<inputs.storage.AccountNetworkRulesPrivateLinkAccessRule>[]>;
    /**
     * The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the ID of the storage account. Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
     *
     * @deprecated Deprecated in favour of `storage_account_id`
     */
    storageAccountName?: pulumi.Input<string>;
    /**
     * A list of virtual network subnet ids to to secure the storage account.
     */
    virtualNetworkSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
}
