// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a API Management Tag.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "My Company",
 *     publisherEmail: "company@terraform.io",
 *     skuName: "Consumption_0",
 * });
 * const exampleTag = new azure.apimanagement.Tag("exampleTag", {apiManagementId: exampleService.id});
 * ```
 *
 * ## Import
 *
 * API Management Tags can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/tag:Tag example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/tags/tag1
 * ```
 */
export class Tag extends pulumi.CustomResource {
    /**
     * Get an existing Tag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagState, opts?: pulumi.CustomResourceOptions): Tag {
        return new Tag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/tag:Tag';

    /**
     * Returns true if the given object is an instance of Tag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tag.__pulumiType;
    }

    /**
     * The ID of the API Management. Changing this forces a new API Management Tag to be created.
     */
    public readonly apiManagementId!: pulumi.Output<string>;
    /**
     * The display name of the API Management Tag. Defaults to the `name`.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The name which should be used for this API Management Tag. Changing this forces a new API Management Tag to be created. The name must be unique in the API Management Service.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Tag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagArgs | TagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagState | undefined;
            resourceInputs["apiManagementId"] = state ? state.apiManagementId : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TagArgs | undefined;
            if ((!args || args.apiManagementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementId'");
            }
            resourceInputs["apiManagementId"] = args ? args.apiManagementId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tag resources.
 */
export interface TagState {
    /**
     * The ID of the API Management. Changing this forces a new API Management Tag to be created.
     */
    apiManagementId?: pulumi.Input<string>;
    /**
     * The display name of the API Management Tag. Defaults to the `name`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The name which should be used for this API Management Tag. Changing this forces a new API Management Tag to be created. The name must be unique in the API Management Service.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tag resource.
 */
export interface TagArgs {
    /**
     * The ID of the API Management. Changing this forces a new API Management Tag to be created.
     */
    apiManagementId: pulumi.Input<string>;
    /**
     * The display name of the API Management Tag. Defaults to the `name`.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The name which should be used for this API Management Tag. Changing this forces a new API Management Tag to be created. The name must be unique in the API Management Service.
     */
    name?: pulumi.Input<string>;
}
