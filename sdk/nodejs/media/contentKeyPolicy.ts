// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Content Key Policy.
 *
 * ## Example Usage
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
 *     accountReplicationType: "GRS",
 * });
 * const exampleServiceAccount = new azure.media.ServiceAccount("exampleServiceAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     storageAccounts: [{
 *         id: exampleAccount.id,
 *         isPrimary: true,
 *     }],
 * });
 * const exampleContentKeyPolicy = new azure.media.ContentKeyPolicy("exampleContentKeyPolicy", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     mediaServicesAccountName: exampleServiceAccount.name,
 *     policyOptions: [
 *         {
 *             name: "fairPlay",
 *             fairplayConfiguration: {
 *                 ask: "bb566284cc124a21c435a92cd3c108c4",
 *                 pfx: "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",
 *                 pfxPassword: "password",
 *                 rentalDurationSeconds: 2249,
 *                 rentalAndLeaseKeyType: "PersistentUnlimited",
 *             },
 *             openRestrictionEnabled: true,
 *         },
 *         {
 *             name: "playReady",
 *             playreadyConfigurationLicenses: [{
 *                 allowTestDevices: true,
 *                 beginDate: "2017-10-16T18:22:53Z",
 *                 playRight: {
 *                     scmsRestriction: 2,
 *                     digitalVideoOnlyContentRestriction: false,
 *                     imageConstraintForAnalogComponentVideoRestriction: false,
 *                     imageConstraintForAnalogComputerMonitorRestriction: false,
 *                     allowPassingVideoContentToUnknownOutput: "NotAllowed",
 *                     uncompressedDigitalVideoOpl: 100,
 *                     uncompressedDigitalAudioOpl: 100,
 *                     analogVideoOpl: 150,
 *                     compressedDigitalAudioOpl: 150,
 *                 },
 *                 licenseType: "Persistent",
 *                 contentType: "UltraVioletDownload",
 *                 contentKeyLocationFromHeaderEnabled: true,
 *             }],
 *             openRestrictionEnabled: true,
 *         },
 *         {
 *             name: "clearKey",
 *             clearKeyConfigurationEnabled: true,
 *             tokenRestriction: {
 *                 issuer: "urn:issuer",
 *                 audience: "urn:audience",
 *                 tokenType: "Swt",
 *                 primarySymmetricTokenKey: "AAAAAAAAAAAAAAAAAAAAAA==",
 *             },
 *         },
 *         {
 *             name: "widevine",
 *             widevineConfigurationTemplate: JSON.stringify({
 *                 allowed_track_types: "SD_HD",
 *                 content_key_specs: [{
 *                     track_type: "SD",
 *                     security_level: 1,
 *                     required_output_protection: {
 *                         hdcp: "HDCP_V2",
 *                     },
 *                 }],
 *                 policy_overrides: {
 *                     can_play: true,
 *                     can_persist: true,
 *                     can_renew: false,
 *                 },
 *             }),
 *             openRestrictionEnabled: true,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Resource Groups can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:media/contentKeyPolicy:ContentKeyPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/account1/contentkeypolicies/policy1
 * ```
 */
export class ContentKeyPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ContentKeyPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContentKeyPolicyState, opts?: pulumi.CustomResourceOptions): ContentKeyPolicy {
        return new ContentKeyPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:media/contentKeyPolicy:ContentKeyPolicy';

    /**
     * Returns true if the given object is an instance of ContentKeyPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentKeyPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentKeyPolicy.__pulumiType;
    }

    /**
     * A description for the Policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Media Services account name. Changing this forces a new Content Key Policy to be created.
     */
    public readonly mediaServicesAccountName!: pulumi.Output<string>;
    /**
     * The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `policyOption` blocks as defined below.
     */
    public readonly policyOptions!: pulumi.Output<outputs.media.ContentKeyPolicyPolicyOption[]>;
    /**
     * The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a ContentKeyPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentKeyPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContentKeyPolicyArgs | ContentKeyPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContentKeyPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["mediaServicesAccountName"] = state ? state.mediaServicesAccountName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyOptions"] = state ? state.policyOptions : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ContentKeyPolicyArgs | undefined;
            if ((!args || args.mediaServicesAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mediaServicesAccountName'");
            }
            if ((!args || args.policyOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyOptions'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["mediaServicesAccountName"] = args ? args.mediaServicesAccountName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyOptions"] = args ? args.policyOptions : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContentKeyPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContentKeyPolicy resources.
 */
export interface ContentKeyPolicyState {
    /**
     * A description for the Policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The Media Services account name. Changing this forces a new Content Key Policy to be created.
     */
    mediaServicesAccountName?: pulumi.Input<string>;
    /**
     * The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more `policyOption` blocks as defined below.
     */
    policyOptions?: pulumi.Input<pulumi.Input<inputs.media.ContentKeyPolicyPolicyOption>[]>;
    /**
     * The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContentKeyPolicy resource.
 */
export interface ContentKeyPolicyArgs {
    /**
     * A description for the Policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The Media Services account name. Changing this forces a new Content Key Policy to be created.
     */
    mediaServicesAccountName: pulumi.Input<string>;
    /**
     * The name which should be used for this Content Key Policy. Changing this forces a new Content Key Policy to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more `policyOption` blocks as defined below.
     */
    policyOptions: pulumi.Input<pulumi.Input<inputs.media.ContentKeyPolicyPolicyOption>[]>;
    /**
     * The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
     */
    resourceGroupName: pulumi.Input<string>;
}
