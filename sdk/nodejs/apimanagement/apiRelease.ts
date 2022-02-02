// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a API Management API Release.
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
 *     skuName: "Developer_1",
 * });
 * const exampleApi = new azure.apimanagement.Api("exampleApi", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     apiManagementName: exampleService.name,
 *     revision: "1",
 *     displayName: "Example API",
 *     path: "example",
 *     protocols: ["https"],
 *     "import": {
 *         contentFormat: "swagger-link-json",
 *         contentValue: "http://conferenceapi.azurewebsites.net/?format=json",
 *     },
 * });
 * const exampleApiRelease = new azure.apimanagement.ApiRelease("exampleApiRelease", {apiId: exampleApi.id});
 * ```
 *
 * ## Import
 *
 * API Management API Releases can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/apiRelease:ApiRelease example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/releases/release1
 * ```
 */
export class ApiRelease extends pulumi.CustomResource {
    /**
     * Get an existing ApiRelease resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiReleaseState, opts?: pulumi.CustomResourceOptions): ApiRelease {
        return new ApiRelease(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/apiRelease:ApiRelease';

    /**
     * Returns true if the given object is an instance of ApiRelease.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiRelease {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiRelease.__pulumiType;
    }

    /**
     * The ID of the API Management API. Changing this forces a new API Management API Release to be created.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The name which should be used for this API Management API Release. Changing this forces a new API Management API Release to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Release Notes.
     */
    public readonly notes!: pulumi.Output<string | undefined>;

    /**
     * Create a ApiRelease resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiReleaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiReleaseArgs | ApiReleaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiReleaseState | undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
        } else {
            const args = argsOrState as ApiReleaseArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiRelease.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiRelease resources.
 */
export interface ApiReleaseState {
    /**
     * The ID of the API Management API. Changing this forces a new API Management API Release to be created.
     */
    apiId?: pulumi.Input<string>;
    /**
     * The name which should be used for this API Management API Release. Changing this forces a new API Management API Release to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The Release Notes.
     */
    notes?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiRelease resource.
 */
export interface ApiReleaseArgs {
    /**
     * The ID of the API Management API. Changing this forces a new API Management API Release to be created.
     */
    apiId: pulumi.Input<string>;
    /**
     * The name which should be used for this API Management API Release. Changing this forces a new API Management API Release to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The Release Notes.
     */
    notes?: pulumi.Input<string>;
}
