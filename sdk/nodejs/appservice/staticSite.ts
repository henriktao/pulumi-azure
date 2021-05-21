// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service Static Site.
 *
 * ->**NOTE**: After the Static Site is provisioned, you'll need to associate your target repository, which contains your web app, to the Static Site, by following the [Azure Static Site document](https://docs.microsoft.com/en-us/azure/static-web-apps/github-actions-workflow).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.appservice.StaticSite("example", {
 *     location: "West Europe",
 *     resourceGroupName: "example",
 * });
 * ```
 *
 * ## Import
 *
 * Static Web Apps can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/staticSite:StaticSite example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
 * ```
 */
export class StaticSite extends pulumi.CustomResource {
    /**
     * Get an existing StaticSite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StaticSiteState, opts?: pulumi.CustomResourceOptions): StaticSite {
        return new StaticSite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/staticSite:StaticSite';

    /**
     * Returns true if the given object is an instance of StaticSite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StaticSite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StaticSite.__pulumiType;
    }

    /**
     * The API key of this Static Web App, which is used for later interacting with this Static Web App from other clients, e.g. Github Action.
     */
    public /*out*/ readonly apiKey!: pulumi.Output<string>;
    /**
     * The default host name of the Static Web App.
     */
    public /*out*/ readonly defaultHostName!: pulumi.Output<string>;
    /**
     * The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    public readonly skuSize!: pulumi.Output<string | undefined>;
    public readonly skuTier!: pulumi.Output<string | undefined>;

    /**
     * Create a StaticSite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StaticSiteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StaticSiteArgs | StaticSiteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StaticSiteState | undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["defaultHostName"] = state ? state.defaultHostName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuSize"] = state ? state.skuSize : undefined;
            inputs["skuTier"] = state ? state.skuTier : undefined;
        } else {
            const args = argsOrState as StaticSiteArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuSize"] = args ? args.skuSize : undefined;
            inputs["skuTier"] = args ? args.skuTier : undefined;
            inputs["apiKey"] = undefined /*out*/;
            inputs["defaultHostName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StaticSite.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StaticSite resources.
 */
export interface StaticSiteState {
    /**
     * The API key of this Static Web App, which is used for later interacting with this Static Web App from other clients, e.g. Github Action.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The default host name of the Static Web App.
     */
    readonly defaultHostName?: pulumi.Input<string>;
    /**
     * The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    readonly skuSize?: pulumi.Input<string>;
    readonly skuTier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StaticSite resource.
 */
export interface StaticSiteArgs {
    /**
     * The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Static Web App. Changing this forces a new Static Web App to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    readonly skuSize?: pulumi.Input<string>;
    readonly skuTier?: pulumi.Input<string>;
}
