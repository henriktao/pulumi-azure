// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the Assignment of an API Management API Tag to an API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = azure.apimanagement.getServiceOutput({
 *     name: "example-apim",
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleApi = new azure.apimanagement.Api("exampleApi", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     apiManagementName: azurerm_api_management.example.name,
 *     revision: "1",
 * });
 * const exampleTag = new azure.apimanagement.Tag("exampleTag", {apiManagementId: exampleService.apply(exampleService => exampleService.id)});
 * const exampleApiTag = new azure.apimanagement.ApiTag("exampleApiTag", {apiId: exampleApi.id});
 * ```
 *
 * ## Import
 *
 * API Management API Tags can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/apiTag:ApiTag example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
 * ```
 */
export class ApiTag extends pulumi.CustomResource {
    /**
     * Get an existing ApiTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiTagState, opts?: pulumi.CustomResourceOptions): ApiTag {
        return new ApiTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/apiTag:ApiTag';

    /**
     * Returns true if the given object is an instance of ApiTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiTag.__pulumiType;
    }

    /**
     * The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ApiTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiTagArgs | ApiTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiTagState | undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ApiTagArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiTag resources.
 */
export interface ApiTagState {
    /**
     * The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
     */
    apiId?: pulumi.Input<string>;
    /**
     * The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiTag resource.
 */
export interface ApiTagArgs {
    /**
     * The ID of the API Management API. Changing this forces a new API Management API Tag to be created.
     */
    apiId: pulumi.Input<string>;
    /**
     * The name of the tag. It must be known in the API Management instance. Changing this forces a new API Management API Tag to be created.
     */
    name?: pulumi.Input<string>;
}
