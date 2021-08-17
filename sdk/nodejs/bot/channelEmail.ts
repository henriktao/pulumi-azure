// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Email integration for a Bot Channel
 *
 * > **Note** A bot can only have a single Email Channel associated with it.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleChannelsRegistration = new azure.bot.ChannelsRegistration("exampleChannelsRegistration", {
 *     location: "global",
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "F0",
 *     microsoftAppId: current.then(current => current.clientId),
 * });
 * const exampleChannelEmail = new azure.bot.ChannelEmail("exampleChannelEmail", {
 *     botName: exampleChannelsRegistration.name,
 *     location: exampleChannelsRegistration.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     clientId: "exampleId",
 *     clientSecret: "exampleSecret",
 *     verificationToken: "exampleVerificationToken",
 * });
 * ```
 *
 * ## Import
 *
 * The Email Integration for a Bot Channel can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:bot/channelEmail:ChannelEmail example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/EmailChannel
 * ```
 */
export class ChannelEmail extends pulumi.CustomResource {
    /**
     * Get an existing ChannelEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChannelEmailState, opts?: pulumi.CustomResourceOptions): ChannelEmail {
        return new ChannelEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:bot/channelEmail:ChannelEmail';

    /**
     * Returns true if the given object is an instance of ChannelEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelEmail.__pulumiType;
    }

    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    public readonly botName!: pulumi.Output<string>;
    /**
     * The email address that the Bot will authenticate with.
     */
    public readonly emailAddress!: pulumi.Output<string>;
    /**
     * The email password that the Bot will authenticate with.
     */
    public readonly emailPassword!: pulumi.Output<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a ChannelEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChannelEmailArgs | ChannelEmailState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ChannelEmailState | undefined;
            inputs["botName"] = state ? state.botName : undefined;
            inputs["emailAddress"] = state ? state.emailAddress : undefined;
            inputs["emailPassword"] = state ? state.emailPassword : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ChannelEmailArgs | undefined;
            if ((!args || args.botName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'botName'");
            }
            if ((!args || args.emailAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailAddress'");
            }
            if ((!args || args.emailPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailPassword'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["botName"] = args ? args.botName : undefined;
            inputs["emailAddress"] = args ? args.emailAddress : undefined;
            inputs["emailPassword"] = args ? args.emailPassword : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ChannelEmail.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChannelEmail resources.
 */
export interface ChannelEmailState {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    botName?: pulumi.Input<string>;
    /**
     * The email address that the Bot will authenticate with.
     */
    emailAddress?: pulumi.Input<string>;
    /**
     * The email password that the Bot will authenticate with.
     */
    emailPassword?: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ChannelEmail resource.
 */
export interface ChannelEmailArgs {
    /**
     * The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
     */
    botName: pulumi.Input<string>;
    /**
     * The email address that the Bot will authenticate with.
     */
    emailAddress: pulumi.Input<string>;
    /**
     * The email password that the Bot will authenticate with.
     */
    emailPassword: pulumi.Input<string>;
    /**
     * The supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Bot Channel. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
}
