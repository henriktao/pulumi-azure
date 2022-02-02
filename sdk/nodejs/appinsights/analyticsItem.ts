// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Application Insights Analytics Item component.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleInsights = new azure.appinsights.Insights("exampleInsights", {
 *     location: "West Europe",
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationType: "web",
 * });
 * const exampleAnalyticsItem = new azure.appinsights.AnalyticsItem("exampleAnalyticsItem", {
 *     applicationInsightsId: exampleInsights.id,
 *     content: "requests //simple example query",
 *     scope: "shared",
 *     type: "query",
 * });
 * ```
 *
 * ## Import
 *
 * Application Insights Analytics Items can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appinsights/analyticsItem:AnalyticsItem example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/mycomponent1/analyticsItems/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  To find the Analytics Item ID you can query the REST API using the [`az rest` CLI command](https://docs.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest#az-rest), e.g. az rest --method GET --uri "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.insights/components/appinsightstest/analyticsItems?api-version=2015-05-01"
 */
export class AnalyticsItem extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsItemState, opts?: pulumi.CustomResourceOptions): AnalyticsItem {
        return new AnalyticsItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appinsights/analyticsItem:AnalyticsItem';

    /**
     * Returns true if the given object is an instance of AnalyticsItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsItem.__pulumiType;
    }

    /**
     * The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
     */
    public readonly applicationInsightsId!: pulumi.Output<string>;
    /**
     * The content for the Analytics Item, for example the query text if `type` is `query`.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The alias to use for the function. Required when `type` is `function`.
     */
    public readonly functionAlias!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * A string containing the time the Analytics Item was created.
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * A string containing the time the Analytics Item was last modified.
     */
    public /*out*/ readonly timeModified!: pulumi.Output<string>;
    /**
     * The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A string indicating the version of the query format
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a AnalyticsItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsItemArgs | AnalyticsItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnalyticsItemState | undefined;
            resourceInputs["applicationInsightsId"] = state ? state.applicationInsightsId : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["functionAlias"] = state ? state.functionAlias : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["timeCreated"] = state ? state.timeCreated : undefined;
            resourceInputs["timeModified"] = state ? state.timeModified : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AnalyticsItemArgs | undefined;
            if ((!args || args.applicationInsightsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationInsightsId'");
            }
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["applicationInsightsId"] = args ? args.applicationInsightsId : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["functionAlias"] = args ? args.functionAlias : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["timeCreated"] = undefined /*out*/;
            resourceInputs["timeModified"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnalyticsItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsItem resources.
 */
export interface AnalyticsItemState {
    /**
     * The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
     */
    applicationInsightsId?: pulumi.Input<string>;
    /**
     * The content for the Analytics Item, for example the query text if `type` is `query`.
     */
    content?: pulumi.Input<string>;
    /**
     * The alias to use for the function. Required when `type` is `function`.
     */
    functionAlias?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
     */
    scope?: pulumi.Input<string>;
    /**
     * A string containing the time the Analytics Item was created.
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * A string containing the time the Analytics Item was last modified.
     */
    timeModified?: pulumi.Input<string>;
    /**
     * The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * A string indicating the version of the query format
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnalyticsItem resource.
 */
export interface AnalyticsItemArgs {
    /**
     * The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
     */
    applicationInsightsId: pulumi.Input<string>;
    /**
     * The content for the Analytics Item, for example the query text if `type` is `query`.
     */
    content: pulumi.Input<string>;
    /**
     * The alias to use for the function. Required when `type` is `function`.
     */
    functionAlias?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
     */
    scope: pulumi.Input<string>;
    /**
     * The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
     */
    type: pulumi.Input<string>;
}
