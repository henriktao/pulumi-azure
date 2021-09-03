// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Application Insights component.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleInsights = new azure.appinsights.Insights("exampleInsights", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     applicationType: "web",
 * });
 * export const instrumentationKey = exampleInsights.instrumentationKey;
 * export const appId = exampleInsights.appId;
 * ```
 * ### Workspace Mode
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 *     retentionInDays: 30,
 * });
 * const exampleInsights = new azure.appinsights.Insights("exampleInsights", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceId: exampleAnalyticsWorkspace.id,
 *     applicationType: "web",
 * });
 * export const instrumentationKey = exampleInsights.instrumentationKey;
 * export const appId = exampleInsights.appId;
 * ```
 *
 * ## Import
 *
 * Application Insights instances can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appinsights/insights:Insights instance1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.insights/components/instance1
 * ```
 */
export class Insights extends pulumi.CustomResource {
    /**
     * Get an existing Insights resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InsightsState, opts?: pulumi.CustomResourceOptions): Insights {
        return new Insights(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appinsights/insights:Insights';

    /**
     * Returns true if the given object is an instance of Insights.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Insights {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Insights.__pulumiType;
    }

    /**
     * The App ID associated with this Application Insights component.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    public readonly applicationType!: pulumi.Output<string>;
    /**
     * The Connection String for this Application Insights component. (Sensitive)
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    public readonly dailyDataCapInGb!: pulumi.Output<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    public readonly dailyDataCapNotificationsDisabled!: pulumi.Output<boolean>;
    /**
     * By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
     */
    public readonly disableIpMasking!: pulumi.Output<boolean | undefined>;
    /**
     * The Instrumentation Key for this Application Insights component. (Sensitive)
     */
    public /*out*/ readonly instrumentationKey!: pulumi.Output<string>;
    /**
     * Disable Non-Azure AD based Auth. Defaults to `false`.
     */
    public readonly localAuthenticationDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    public readonly samplingPercentage!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the id of a log analytics workspace resource
     */
    public readonly workspaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a Insights resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InsightsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InsightsArgs | InsightsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InsightsState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["applicationType"] = state ? state.applicationType : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dailyDataCapInGb"] = state ? state.dailyDataCapInGb : undefined;
            inputs["dailyDataCapNotificationsDisabled"] = state ? state.dailyDataCapNotificationsDisabled : undefined;
            inputs["disableIpMasking"] = state ? state.disableIpMasking : undefined;
            inputs["instrumentationKey"] = state ? state.instrumentationKey : undefined;
            inputs["localAuthenticationDisabled"] = state ? state.localAuthenticationDisabled : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            inputs["samplingPercentage"] = state ? state.samplingPercentage : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as InsightsArgs | undefined;
            if ((!args || args.applicationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationType"] = args ? args.applicationType : undefined;
            inputs["dailyDataCapInGb"] = args ? args.dailyDataCapInGb : undefined;
            inputs["dailyDataCapNotificationsDisabled"] = args ? args.dailyDataCapNotificationsDisabled : undefined;
            inputs["disableIpMasking"] = args ? args.disableIpMasking : undefined;
            inputs["localAuthenticationDisabled"] = args ? args.localAuthenticationDisabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["samplingPercentage"] = args ? args.samplingPercentage : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceId"] = args ? args.workspaceId : undefined;
            inputs["appId"] = undefined /*out*/;
            inputs["connectionString"] = undefined /*out*/;
            inputs["instrumentationKey"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Insights.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Insights resources.
 */
export interface InsightsState {
    /**
     * The App ID associated with this Application Insights component.
     */
    appId?: pulumi.Input<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    applicationType?: pulumi.Input<string>;
    /**
     * The Connection String for this Application Insights component. (Sensitive)
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    dailyDataCapInGb?: pulumi.Input<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    dailyDataCapNotificationsDisabled?: pulumi.Input<boolean>;
    /**
     * By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
     */
    disableIpMasking?: pulumi.Input<boolean>;
    /**
     * The Instrumentation Key for this Application Insights component. (Sensitive)
     */
    instrumentationKey?: pulumi.Input<string>;
    /**
     * Disable Non-Azure AD based Auth. Defaults to `false`.
     */
    localAuthenticationDisabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    samplingPercentage?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the id of a log analytics workspace resource
     */
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Insights resource.
 */
export interface InsightsArgs {
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    applicationType: pulumi.Input<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    dailyDataCapInGb?: pulumi.Input<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    dailyDataCapNotificationsDisabled?: pulumi.Input<boolean>;
    /**
     * By default the real client ip is masked as `0.0.0.0` in the logs. Use this argument to disable masking and log the real client ip. Defaults to `false`.
     */
    disableIpMasking?: pulumi.Input<boolean>;
    /**
     * Disable Non-Azure AD based Auth. Defaults to `false`.
     */
    localAuthenticationDisabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`. Defaults to `90`.
     */
    retentionInDays?: pulumi.Input<number>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    samplingPercentage?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the id of a log analytics workspace resource
     */
    workspaceId?: pulumi.Input<string>;
}
