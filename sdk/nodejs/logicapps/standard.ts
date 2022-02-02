// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Logic App (Standard / Single Tenant)
 *
 * > **Note:** To connect an Azure Logic App and a subnet within the same region `azure.appservice.VirtualNetworkSwiftConnection` can be used.
 * For an example, check the `azure.appservice.VirtualNetworkSwiftConnection` documentation.
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
 *         tier: "WorkflowStandard",
 *         size: "WS1",
 *     },
 * });
 * const exampleStandard = new azure.logicapps.Standard("exampleStandard", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 * });
 * ```
 * ### For Container Mode)
 *
 * > **Note:** You must set `azure.appservice.Plan` `kind` to `Linux` and `reserved` to `true` when used with `linuxFxVersion`
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
 *     kind: "Linux",
 *     reserved: true,
 *     sku: {
 *         tier: "WorkflowStandard",
 *         size: "WS1",
 *     },
 * });
 * const exampleStandard = new azure.logicapps.Standard("exampleStandard", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *     siteConfig: {
 *         linuxFxVersion: "DOCKER|mcr.microsoft.com/azure-functions/dotnet:3.0-appservice",
 *     },
 *     appSettings: {
 *         DOCKER_REGISTRY_SERVER_URL: "https://<server-name>.azurecr.io",
 *         DOCKER_REGISTRY_SERVER_USERNAME: "username",
 *         DOCKER_REGISTRY_SERVER_PASSWORD: "password",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Logic Apps can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:logicapps/standard:Standard logicapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/logicapp1
 * ```
 */
export class Standard extends pulumi.CustomResource {
    /**
     * Get an existing Standard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StandardState, opts?: pulumi.CustomResourceOptions): Standard {
        return new Standard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:logicapps/standard:Standard';

    /**
     * Returns true if the given object is an instance of Standard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Standard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Standard.__pulumiType;
    }

    /**
     * The ID of the App Service Plan within which to create this Logic App
     */
    public readonly appServicePlanId!: pulumi.Output<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    public readonly appSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * If `useExtensionBundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
     */
    public readonly bundleVersion!: pulumi.Output<string | undefined>;
    /**
     * Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean>;
    /**
     * The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
     */
    public readonly clientCertificateMode!: pulumi.Output<string | undefined>;
    /**
     * An `connectionString` block as defined below.
     */
    public readonly connectionStrings!: pulumi.Output<outputs.logicapps.StandardConnectionString[]>;
    /**
     * An identifier used by App Service to perform domain ownership verification via DNS TXT record.
     */
    public /*out*/ readonly customDomainVerificationId!: pulumi.Output<string>;
    /**
     * The default hostname associated with the Logic App - such as `mysite.azurewebsites.net`
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * Is the Logic App enabled?
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Can the Logic App only be accessed via HTTPS? Defaults to `false`.
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.logicapps.StandardIdentity>;
    /**
     * The Logic App kind - will be `functionapp,workflowapp`
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Logic App Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    public /*out*/ readonly outboundIpAddresses!: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    public /*out*/ readonly possibleOutboundIpAddresses!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Logic App
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    public readonly siteConfig!: pulumi.Output<outputs.logicapps.StandardSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    public /*out*/ readonly siteCredentials!: pulumi.Output<outputs.logicapps.StandardSiteCredential[]>;
    /**
     * The access key which will be used to access the backend storage account for the Logic App
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string>;
    /**
     * The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
     */
    public readonly storageAccountShareName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
     */
    public readonly useExtensionBundle!: pulumi.Output<boolean | undefined>;
    /**
     * The runtime version associated with the Logic App Defaults to `~1`.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Standard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StandardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StandardArgs | StandardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StandardState | undefined;
            resourceInputs["appServicePlanId"] = state ? state.appServicePlanId : undefined;
            resourceInputs["appSettings"] = state ? state.appSettings : undefined;
            resourceInputs["bundleVersion"] = state ? state.bundleVersion : undefined;
            resourceInputs["clientAffinityEnabled"] = state ? state.clientAffinityEnabled : undefined;
            resourceInputs["clientCertificateMode"] = state ? state.clientCertificateMode : undefined;
            resourceInputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            resourceInputs["customDomainVerificationId"] = state ? state.customDomainVerificationId : undefined;
            resourceInputs["defaultHostname"] = state ? state.defaultHostname : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["httpsOnly"] = state ? state.httpsOnly : undefined;
            resourceInputs["identity"] = state ? state.identity : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outboundIpAddresses"] = state ? state.outboundIpAddresses : undefined;
            resourceInputs["possibleOutboundIpAddresses"] = state ? state.possibleOutboundIpAddresses : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["siteConfig"] = state ? state.siteConfig : undefined;
            resourceInputs["siteCredentials"] = state ? state.siteCredentials : undefined;
            resourceInputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            resourceInputs["storageAccountShareName"] = state ? state.storageAccountShareName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["useExtensionBundle"] = state ? state.useExtensionBundle : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as StandardArgs | undefined;
            if ((!args || args.appServicePlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appServicePlanId'");
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
            resourceInputs["bundleVersion"] = args ? args.bundleVersion : undefined;
            resourceInputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            resourceInputs["clientCertificateMode"] = args ? args.clientCertificateMode : undefined;
            resourceInputs["connectionStrings"] = args ? args.connectionStrings : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["siteConfig"] = args ? args.siteConfig : undefined;
            resourceInputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            resourceInputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            resourceInputs["storageAccountShareName"] = args ? args.storageAccountShareName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useExtensionBundle"] = args ? args.useExtensionBundle : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["customDomainVerificationId"] = undefined /*out*/;
            resourceInputs["defaultHostname"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["outboundIpAddresses"] = undefined /*out*/;
            resourceInputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            resourceInputs["siteCredentials"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Standard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Standard resources.
 */
export interface StandardState {
    /**
     * The ID of the App Service Plan within which to create this Logic App
     */
    appServicePlanId?: pulumi.Input<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If `useExtensionBundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
     */
    bundleVersion?: pulumi.Input<string>;
    /**
     * Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
     */
    clientCertificateMode?: pulumi.Input<string>;
    /**
     * An `connectionString` block as defined below.
     */
    connectionStrings?: pulumi.Input<pulumi.Input<inputs.logicapps.StandardConnectionString>[]>;
    /**
     * An identifier used by App Service to perform domain ownership verification via DNS TXT record.
     */
    customDomainVerificationId?: pulumi.Input<string>;
    /**
     * The default hostname associated with the Logic App - such as `mysite.azurewebsites.net`
     */
    defaultHostname?: pulumi.Input<string>;
    /**
     * Is the Logic App enabled?
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Can the Logic App only be accessed via HTTPS? Defaults to `false`.
     */
    httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.logicapps.StandardIdentity>;
    /**
     * The Logic App kind - will be `functionapp,workflowapp`
     */
    kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Logic App Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    outboundIpAddresses?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    possibleOutboundIpAddresses?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Logic App
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    siteConfig?: pulumi.Input<inputs.logicapps.StandardSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    siteCredentials?: pulumi.Input<pulumi.Input<inputs.logicapps.StandardSiteCredential>[]>;
    /**
     * The access key which will be used to access the backend storage account for the Logic App
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
     */
    storageAccountName?: pulumi.Input<string>;
    /**
     * The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
     */
    storageAccountShareName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
     */
    useExtensionBundle?: pulumi.Input<boolean>;
    /**
     * The runtime version associated with the Logic App Defaults to `~1`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Standard resource.
 */
export interface StandardArgs {
    /**
     * The ID of the App Service Plan within which to create this Logic App
     */
    appServicePlanId: pulumi.Input<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If `useExtensionBundle` then controls the allowed range for bundle versions. Default `[1.*, 2.0.0)`
     */
    bundleVersion?: pulumi.Input<string>;
    /**
     * Should the Logic App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * The mode of the Logic App's client certificates requirement for incoming requests. Possible values are `Required` and `Optional`.
     */
    clientCertificateMode?: pulumi.Input<string>;
    /**
     * An `connectionString` block as defined below.
     */
    connectionStrings?: pulumi.Input<pulumi.Input<inputs.logicapps.StandardConnectionString>[]>;
    /**
     * Is the Logic App enabled?
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Can the Logic App only be accessed via HTTPS? Defaults to `false`.
     */
    httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    identity?: pulumi.Input<inputs.logicapps.StandardIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Logic App Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Logic App
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    siteConfig?: pulumi.Input<inputs.logicapps.StandardSiteConfig>;
    /**
     * The access key which will be used to access the backend storage account for the Logic App
     */
    storageAccountAccessKey: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by this Logic App (e.g. for Stateful workflows data)
     */
    storageAccountName: pulumi.Input<string>;
    /**
     * The name of the share used by the logic app, if you want to use a custom name. This corresponds to the WEBSITE_CONTENTSHARE appsetting, which this resource will create for you. If you don't specify a name, then this resource will generate a dynamic name.  This setting is useful if you want to provision a storage account and create a share using azurerm_storage_share
     */
    storageAccountShareName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Should the logic app use the bundled extension package? If true, then application settings for `AzureFunctionsJobHost__extensionBundle__id` and `AzureFunctionsJobHost__extensionBundle__version` will be created. Default true
     */
    useExtensionBundle?: pulumi.Input<boolean>;
    /**
     * The runtime version associated with the Logic App Defaults to `~1`.
     */
    version?: pulumi.Input<string>;
}
