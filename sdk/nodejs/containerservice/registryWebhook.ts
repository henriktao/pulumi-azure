// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure Container Registry Webhook.
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
 *     sku: "Standard",
 *     adminEnabled: false,
 * });
 * const webhook = new azure.containerservice.RegistryWebhook("webhook", {
 *     resourceGroupName: rg.name,
 *     registryName: acr.name,
 *     location: rg.location,
 *     serviceUri: "https://mywebhookreceiver.example/mytag",
 *     status: "enabled",
 *     scope: "mytag:*",
 *     actions: ["push"],
 *     customHeaders: {
 *         "Content-Type": "application/json",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Container Registry Webhooks can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:containerservice/registryWebhook:RegistryWebhook example /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1/webhooks/mywebhook1
 * ```
 */
export class RegistryWebhook extends pulumi.CustomResource {
    /**
     * Get an existing RegistryWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryWebhookState, opts?: pulumi.CustomResourceOptions): RegistryWebhook {
        return new RegistryWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:containerservice/registryWebhook:RegistryWebhook';

    /**
     * Returns true if the given object is an instance of RegistryWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryWebhook.__pulumiType;
    }

    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    public readonly actions!: pulumi.Output<string[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    public readonly customHeaders!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    public readonly registryName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    public readonly serviceUri!: pulumi.Output<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a RegistryWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryWebhookArgs | RegistryWebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryWebhookState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["customHeaders"] = state ? state.customHeaders : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["registryName"] = state ? state.registryName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["serviceUri"] = state ? state.serviceUri : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RegistryWebhookArgs | undefined;
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.registryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceUri'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["customHeaders"] = args ? args.customHeaders : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["serviceUri"] = args ? args.serviceUri : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure:containerservice/registryWebook:RegistryWebook" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(RegistryWebhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryWebhook resources.
 */
export interface RegistryWebhookState {
    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    readonly actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    readonly customHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    readonly registryName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    readonly serviceUri?: pulumi.Input<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RegistryWebhook resource.
 */
export interface RegistryWebhookArgs {
    /**
     * A list of actions that trigger the Webhook to post notifications. At least one action needs to be specified. Valid values are: `push`, `delete`, `quarantine`, `chartPush`, `chartDelete`
     */
    readonly actions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom headers that will be added to the webhook notifications request.
     */
    readonly customHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Name of Container registry this Webhook belongs to. Changing this forces a new resource to be created.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Container Registry Webhook. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the scope of repositories that can trigger an event. For example, `foo:*` means events for all tags under repository `foo`. `foo:bar` means events for 'foo:bar' only. `foo` is equivalent to `foo:latest`. Empty means all events.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * Specifies the service URI for the Webhook to post notifications.
     */
    readonly serviceUri: pulumi.Input<string>;
    /**
     * Specifies if this Webhook triggers notifications or not. Valid values: `enabled` and `disabled`. Default is `enabled`.
     */
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
