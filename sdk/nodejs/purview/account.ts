// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Purview Account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.purview.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * ```
 *
 * ## Import
 *
 * Purview Accounts can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:purview/account:Account example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Purview/accounts/account1
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:purview/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * Atlas Kafka endpoint primary connection string.
     */
    public /*out*/ readonly atlasKafkaEndpointPrimaryConnectionString!: pulumi.Output<string>;
    /**
     * Atlas Kafka endpoint secondary connection string.
     */
    public /*out*/ readonly atlasKafkaEndpointSecondaryConnectionString!: pulumi.Output<string>;
    /**
     * Catalog endpoint.
     */
    public /*out*/ readonly catalogEndpoint!: pulumi.Output<string>;
    /**
     * Guardian endpoint.
     */
    public /*out*/ readonly guardianEndpoint!: pulumi.Output<string>;
    /**
     * A `identity` block as defined below.
     */
    public /*out*/ readonly identities!: pulumi.Output<outputs.purview.AccountIdentity[]>;
    /**
     * The Azure Region where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for the new Resource Group where Purview Account creates the managed resources. Changing this forces a new Purview Account to be created.
     */
    public readonly managedResourceGroupName!: pulumi.Output<string>;
    /**
     * A `managedResources` block as defined below.
     */
    public /*out*/ readonly managedResources!: pulumi.Output<outputs.purview.AccountManagedResource[]>;
    /**
     * The name which should be used for this Purview Account. Changing this forces a new Purview Account to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Should the Purview Account be visible to the public network? Defaults to `true`.
     */
    public readonly publicNetworkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Resource Group where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Scan endpoint.
     */
    public /*out*/ readonly scanEndpoint!: pulumi.Output<string>;
    /**
     * @deprecated This property can no longer be specified on create/update, it can only be updated by creating a support ticket at Azure
     */
    public readonly skuName!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags which should be assigned to the Purview Account.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            resourceInputs["atlasKafkaEndpointPrimaryConnectionString"] = state ? state.atlasKafkaEndpointPrimaryConnectionString : undefined;
            resourceInputs["atlasKafkaEndpointSecondaryConnectionString"] = state ? state.atlasKafkaEndpointSecondaryConnectionString : undefined;
            resourceInputs["catalogEndpoint"] = state ? state.catalogEndpoint : undefined;
            resourceInputs["guardianEndpoint"] = state ? state.guardianEndpoint : undefined;
            resourceInputs["identities"] = state ? state.identities : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["managedResourceGroupName"] = state ? state.managedResourceGroupName : undefined;
            resourceInputs["managedResources"] = state ? state.managedResources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicNetworkEnabled"] = state ? state.publicNetworkEnabled : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["scanEndpoint"] = state ? state.scanEndpoint : undefined;
            resourceInputs["skuName"] = state ? state.skuName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedResourceGroupName"] = args ? args.managedResourceGroupName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicNetworkEnabled"] = args ? args.publicNetworkEnabled : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["skuName"] = args ? args.skuName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["atlasKafkaEndpointPrimaryConnectionString"] = undefined /*out*/;
            resourceInputs["atlasKafkaEndpointSecondaryConnectionString"] = undefined /*out*/;
            resourceInputs["catalogEndpoint"] = undefined /*out*/;
            resourceInputs["guardianEndpoint"] = undefined /*out*/;
            resourceInputs["identities"] = undefined /*out*/;
            resourceInputs["managedResources"] = undefined /*out*/;
            resourceInputs["scanEndpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Atlas Kafka endpoint primary connection string.
     */
    atlasKafkaEndpointPrimaryConnectionString?: pulumi.Input<string>;
    /**
     * Atlas Kafka endpoint secondary connection string.
     */
    atlasKafkaEndpointSecondaryConnectionString?: pulumi.Input<string>;
    /**
     * Catalog endpoint.
     */
    catalogEndpoint?: pulumi.Input<string>;
    /**
     * Guardian endpoint.
     */
    guardianEndpoint?: pulumi.Input<string>;
    /**
     * A `identity` block as defined below.
     */
    identities?: pulumi.Input<pulumi.Input<inputs.purview.AccountIdentity>[]>;
    /**
     * The Azure Region where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name which should be used for the new Resource Group where Purview Account creates the managed resources. Changing this forces a new Purview Account to be created.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * A `managedResources` block as defined below.
     */
    managedResources?: pulumi.Input<pulumi.Input<inputs.purview.AccountManagedResource>[]>;
    /**
     * The name which should be used for this Purview Account. Changing this forces a new Purview Account to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Should the Purview Account be visible to the public network? Defaults to `true`.
     */
    publicNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Scan endpoint.
     */
    scanEndpoint?: pulumi.Input<string>;
    /**
     * @deprecated This property can no longer be specified on create/update, it can only be updated by creating a support ticket at Azure
     */
    skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Purview Account.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The Azure Region where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name which should be used for the new Resource Group where Purview Account creates the managed resources. Changing this forces a new Purview Account to be created.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * The name which should be used for this Purview Account. Changing this forces a new Purview Account to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Should the Purview Account be visible to the public network? Defaults to `true`.
     */
    publicNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Resource Group where the Purview Account should exist. Changing this forces a new Purview Account to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * @deprecated This property can no longer be specified on create/update, it can only be updated by creating a support ticket at Azure
     */
    skuName?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Purview Account.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
