// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure App Configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const rg = new azure.core.ResourceGroup("rg", {location: "West Europe"});
 * const appconf = new azure.appconfiguration.ConfigurationStore("appconf", {
 *     resourceGroupName: rg.name,
 *     location: rg.location,
 * });
 * ```
 *
 * ## Import
 *
 * App Configurations can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appconfiguration/configurationStore:ConfigurationStore appconf /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1
 * ```
 */
export class ConfigurationStore extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationStoreState, opts?: pulumi.CustomResourceOptions): ConfigurationStore {
        return new ConfigurationStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appconfiguration/configurationStore:ConfigurationStore';

    /**
     * Returns true if the given object is an instance of ConfigurationStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationStore.__pulumiType;
    }

    /**
     * The URL of the App Configuration.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.appconfiguration.ConfigurationStoreIdentity | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A `primaryReadKey` block as defined below containing the primary read access key.
     */
    public /*out*/ readonly primaryReadKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStorePrimaryReadKey[]>;
    /**
     * A `primaryWriteKey` block as defined below containing the primary write access key.
     */
    public /*out*/ readonly primaryWriteKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStorePrimaryWriteKey[]>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `secondaryReadKey` block as defined below containing the secondary read access key.
     */
    public /*out*/ readonly secondaryReadKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStoreSecondaryReadKey[]>;
    /**
     * A `secondaryWriteKey` block as defined below containing the secondary write access key.
     */
    public /*out*/ readonly secondaryWriteKeys!: pulumi.Output<outputs.appconfiguration.ConfigurationStoreSecondaryWriteKey[]>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    public readonly sku!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ConfigurationStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationStoreArgs | ConfigurationStoreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigurationStoreState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["primaryReadKeys"] = state ? state.primaryReadKeys : undefined;
            resourceInputs["primaryWriteKeys"] = state ? state.primaryWriteKeys : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["secondaryReadKeys"] = state ? state.secondaryReadKeys : undefined;
            resourceInputs["secondaryWriteKeys"] = state ? state.secondaryWriteKeys : undefined;
            resourceInputs["sku"] = state ? state.sku : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ConfigurationStoreArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["primaryReadKeys"] = undefined /*out*/;
            resourceInputs["primaryWriteKeys"] = undefined /*out*/;
            resourceInputs["secondaryReadKeys"] = undefined /*out*/;
            resourceInputs["secondaryWriteKeys"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigurationStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigurationStore resources.
 */
export interface ConfigurationStoreState {
    /**
     * The URL of the App Configuration.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.appconfiguration.ConfigurationStoreIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A `primaryReadKey` block as defined below containing the primary read access key.
     */
    primaryReadKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStorePrimaryReadKey>[]>;
    /**
     * A `primaryWriteKey` block as defined below containing the primary write access key.
     */
    primaryWriteKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStorePrimaryWriteKey>[]>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `secondaryReadKey` block as defined below containing the secondary read access key.
     */
    secondaryReadKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStoreSecondaryReadKey>[]>;
    /**
     * A `secondaryWriteKey` block as defined below containing the secondary write access key.
     */
    secondaryWriteKeys?: pulumi.Input<pulumi.Input<inputs.appconfiguration.ConfigurationStoreSecondaryWriteKey>[]>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ConfigurationStore resource.
 */
export interface ConfigurationStoreArgs {
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.appconfiguration.ConfigurationStoreIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Configuration. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU name of the the App Configuration. Possible values are `free` and `standard`.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
