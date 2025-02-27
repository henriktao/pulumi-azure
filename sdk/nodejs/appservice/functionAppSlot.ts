// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Function App deployment Slot.
 *
 * ## Example Usage
 * ### With App Service Plan)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const examplePlan = new azure.appservice.Plan("examplePlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         tier: "Standard",
 *         size: "S1",
 *     },
 * });
 * const exampleFunctionApp = new azure.appservice.FunctionApp("exampleFunctionApp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 * });
 * const exampleFunctionAppSlot = new azure.appservice.FunctionAppSlot("exampleFunctionAppSlot", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     functionAppName: exampleFunctionApp.name,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 * });
 * ```
 *
 * ## Import
 *
 * Function Apps Deployment Slots can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/functionAppSlot:FunctionAppSlot functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1/slots/staging
 * ```
 */
export class FunctionAppSlot extends pulumi.CustomResource {
    /**
     * Get an existing FunctionAppSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionAppSlotState, opts?: pulumi.CustomResourceOptions): FunctionAppSlot {
        return new FunctionAppSlot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/functionAppSlot:FunctionAppSlot';

    /**
     * Returns true if the given object is an instance of FunctionAppSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionAppSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionAppSlot.__pulumiType;
    }

    /**
     * The ID of the App Service Plan within which to create this Function App Slot.
     */
    public readonly appServicePlanId!: pulumi.Output<string>;
    /**
     * A key-value pair of App Settings.
     */
    public readonly appSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * An `authSettings` block as defined below.
     */
    public readonly authSettings!: pulumi.Output<outputs.appservice.FunctionAppSlotAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     *
     * @deprecated This property is no longer configurable in the service and has been deprecated. It will be removed in 3.0 of the provider.
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean>;
    /**
     * A `connectionString` block as defined below.
     */
    public readonly connectionStrings!: pulumi.Output<outputs.appservice.FunctionAppSlotConnectionString[]>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    public readonly dailyMemoryTimeQuota!: pulumi.Output<number | undefined>;
    /**
     * The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * Should the built-in logging of the Function App be enabled? Defaults to `true`.
     */
    public readonly enableBuiltinLogging!: pulumi.Output<boolean | undefined>;
    /**
     * Is the Function App enabled?
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
     */
    public readonly functionAppName!: pulumi.Output<string>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.appservice.FunctionAppSlotIdentity>;
    /**
     * The Function App kind - such as `functionapp,linux,container`
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    public /*out*/ readonly outboundIpAddresses!: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    public /*out*/ readonly possibleOutboundIpAddresses!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Function App Slot.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    public readonly siteConfig!: pulumi.Output<outputs.appservice.FunctionAppSlotSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
     */
    public /*out*/ readonly siteCredentials!: pulumi.Output<outputs.appservice.FunctionAppSlotSiteCredential[]>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string>;
    /**
     * The backend storage account name which will be used by the Function App (such as the dashboard, logs).
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a FunctionAppSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionAppSlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionAppSlotArgs | FunctionAppSlotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionAppSlotState | undefined;
            resourceInputs["appServicePlanId"] = state ? state.appServicePlanId : undefined;
            resourceInputs["appSettings"] = state ? state.appSettings : undefined;
            resourceInputs["authSettings"] = state ? state.authSettings : undefined;
            resourceInputs["clientAffinityEnabled"] = state ? state.clientAffinityEnabled : undefined;
            resourceInputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            resourceInputs["dailyMemoryTimeQuota"] = state ? state.dailyMemoryTimeQuota : undefined;
            resourceInputs["defaultHostname"] = state ? state.defaultHostname : undefined;
            resourceInputs["enableBuiltinLogging"] = state ? state.enableBuiltinLogging : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["functionAppName"] = state ? state.functionAppName : undefined;
            resourceInputs["httpsOnly"] = state ? state.httpsOnly : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["outboundIpAddresses"] = state ? state.outboundIpAddresses : undefined;
            resourceInputs["possibleOutboundIpAddresses"] = state ? state.possibleOutboundIpAddresses : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["siteConfig"] = state ? state.siteConfig : undefined;
            resourceInputs["siteCredentials"] = state ? state.siteCredentials : undefined;
            resourceInputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FunctionAppSlotArgs | undefined;
            if ((!args || args.appServicePlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appServicePlanId'");
            }
            if ((!args || args.functionAppName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionAppName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.storageAccountAccessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountAccessKey'");
            }
            if ((!args || args.storageAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            resourceInputs["appServicePlanId"] = args ? args.appServicePlanId : undefined;
            resourceInputs["appSettings"] = args ? args.appSettings : undefined;
            resourceInputs["authSettings"] = args ? args.authSettings : undefined;
            resourceInputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            resourceInputs["connectionStrings"] = args ? args.connectionStrings : undefined;
            resourceInputs["dailyMemoryTimeQuota"] = args ? args.dailyMemoryTimeQuota : undefined;
            resourceInputs["enableBuiltinLogging"] = args ? args.enableBuiltinLogging : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["functionAppName"] = args ? args.functionAppName : undefined;
            resourceInputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osType"] = args ? args.osType : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["siteConfig"] = args ? args.siteConfig : undefined;
            resourceInputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["defaultHostname"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["outboundIpAddresses"] = undefined /*out*/;
            resourceInputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            resourceInputs["siteCredentials"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionAppSlot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionAppSlot resources.
 */
export interface FunctionAppSlotState {
    /**
     * The ID of the App Service Plan within which to create this Function App Slot.
     */
    appServicePlanId?: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `authSettings` block as defined below.
     */
    authSettings?: pulumi.Input<inputs.appservice.FunctionAppSlotAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     *
     * @deprecated This property is no longer configurable in the service and has been deprecated. It will be removed in 3.0 of the provider.
     */
    clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * A `connectionString` block as defined below.
     */
    connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppSlotConnectionString>[]>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
     */
    defaultHostname?: pulumi.Input<string>;
    /**
     * Should the built-in logging of the Function App be enabled? Defaults to `true`.
     */
    enableBuiltinLogging?: pulumi.Input<boolean>;
    /**
     * Is the Function App enabled?
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
     */
    functionAppName?: pulumi.Input<string>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.appservice.FunctionAppSlotIdentity>;
    /**
     * The Function App kind - such as `functionapp,linux,container`
     */
    kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    osType?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    outboundIpAddresses?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    possibleOutboundIpAddresses?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Function App Slot.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    siteConfig?: pulumi.Input<inputs.appservice.FunctionAppSlotSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
     */
    siteCredentials?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppSlotSiteCredential>[]>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by the Function App (such as the dashboard, logs).
     */
    storageAccountName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionAppSlot resource.
 */
export interface FunctionAppSlotArgs {
    /**
     * The ID of the App Service Plan within which to create this Function App Slot.
     */
    appServicePlanId: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An `authSettings` block as defined below.
     */
    authSettings?: pulumi.Input<inputs.appservice.FunctionAppSlotAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     *
     * @deprecated This property is no longer configurable in the service and has been deprecated. It will be removed in 3.0 of the provider.
     */
    clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * A `connectionString` block as defined below.
     */
    connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppSlotConnectionString>[]>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * Should the built-in logging of the Function App be enabled? Defaults to `true`.
     */
    enableBuiltinLogging?: pulumi.Input<boolean>;
    /**
     * Is the Function App enabled?
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
     */
    functionAppName: pulumi.Input<string>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.appservice.FunctionAppSlotIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    osType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Function App Slot.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    siteConfig?: pulumi.Input<inputs.appservice.FunctionAppSlotSiteConfig>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    storageAccountAccessKey: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by the Function App (such as the dashboard, logs).
     */
    storageAccountName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    version?: pulumi.Input<string>;
}
