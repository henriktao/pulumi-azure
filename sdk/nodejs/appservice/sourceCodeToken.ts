// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service source control token.
 *
 * > **NOTE:** Source Control Tokens are configured at the subscription level, not on each App Service - as such this can only be configured Subscription-wide
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.appservice.SourceCodeToken("example", {
 *     token: "7e57735e77e577e57",
 *     type: "GitHub",
 * });
 * ```
 *
 * ## Import
 *
 * App Service Source Control Token's can be imported using the `type`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/sourceCodeToken:SourceCodeToken example GitHub
 * ```
 */
export class SourceCodeToken extends pulumi.CustomResource {
    /**
     * Get an existing SourceCodeToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SourceCodeTokenState, opts?: pulumi.CustomResourceOptions): SourceCodeToken {
        return new SourceCodeToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/sourceCodeToken:SourceCodeToken';

    /**
     * Returns true if the given object is an instance of SourceCodeToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SourceCodeToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SourceCodeToken.__pulumiType;
    }

    /**
     * The OAuth access token.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The OAuth access token secret.
     */
    public readonly tokenSecret!: pulumi.Output<string | undefined>;
    /**
     * The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SourceCodeToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SourceCodeTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SourceCodeTokenArgs | SourceCodeTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SourceCodeTokenState | undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["tokenSecret"] = state ? state.tokenSecret : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SourceCodeTokenArgs | undefined;
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["tokenSecret"] = args ? args.tokenSecret : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SourceCodeToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SourceCodeToken resources.
 */
export interface SourceCodeTokenState {
    /**
     * The OAuth access token.
     */
    token?: pulumi.Input<string>;
    /**
     * The OAuth access token secret.
     */
    tokenSecret?: pulumi.Input<string>;
    /**
     * The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SourceCodeToken resource.
 */
export interface SourceCodeTokenArgs {
    /**
     * The OAuth access token.
     */
    token: pulumi.Input<string>;
    /**
     * The OAuth access token secret.
     */
    tokenSecret?: pulumi.Input<string>;
    /**
     * The source control type. Possible values are `BitBucket`, `Dropbox`, `GitHub` and `OneDrive`.
     */
    type: pulumi.Input<string>;
}
