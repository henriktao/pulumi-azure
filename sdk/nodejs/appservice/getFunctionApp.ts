// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about a Function App.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = azure.appservice.getFunctionApp({
 *     name: "test-azure-functions",
 *     resourceGroupName: azurerm_resource_group.example.name,
 * });
 * ```
 */
export function getFunctionApp(args: GetFunctionAppArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionAppResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:appservice/getFunctionApp:getFunctionApp", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionApp.
 */
export interface GetFunctionAppArgs {
    /**
     * The name of the Function App resource.
     */
    name: string;
    /**
     * The name of the Resource Group where the Function App exists.
     */
    resourceGroupName: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getFunctionApp.
 */
export interface GetFunctionAppResult {
    /**
     * The ID of the App Service Plan within which to create this Function App.
     */
    readonly appServicePlanId: string;
    /**
     * A key-value pair of App Settings.
     */
    readonly appSettings: {[key: string]: string};
    /**
     * The mode of the Function App's client certificates requirement for incoming requests.
     */
    readonly clientCertMode: string;
    /**
     * An `connectionString` block as defined below.
     */
    readonly connectionStrings: outputs.appservice.GetFunctionAppConnectionString[];
    /**
     * An identifier used by App Service to perform domain ownership verification via DNS TXT record.
     */
    readonly customDomainVerificationId: string;
    /**
     * The default hostname associated with the Function App.
     */
    readonly defaultHostname: string;
    /**
     * Is the Function App enabled?
     */
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A `identity` block as defined below.
     */
    readonly identities: outputs.appservice.GetFunctionAppIdentity[];
    readonly location: string;
    /**
     * The name for this IP Restriction.
     */
    readonly name: string;
    /**
     * A string indicating the Operating System type for this function app.
     */
    readonly osType: string;
    /**
     * A comma separated list of outbound IP addresses.
     */
    readonly outboundIpAddresses: string;
    /**
     * A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    readonly possibleOutboundIpAddresses: string;
    readonly resourceGroupName: string;
    readonly siteConfigs: outputs.appservice.GetFunctionAppSiteConfig[];
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    readonly siteCredentials: outputs.appservice.GetFunctionAppSiteCredential[];
    /**
     * A `sourceControl` block as defined below.
     */
    readonly sourceControls: outputs.appservice.GetFunctionAppSourceControl[];
    readonly tags?: {[key: string]: string};
}

export function getFunctionAppOutput(args: GetFunctionAppOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionAppResult> {
    return pulumi.output(args).apply(a => getFunctionApp(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionApp.
 */
export interface GetFunctionAppOutputArgs {
    /**
     * The name of the Function App resource.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Resource Group where the Function App exists.
     */
    resourceGroupName: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
