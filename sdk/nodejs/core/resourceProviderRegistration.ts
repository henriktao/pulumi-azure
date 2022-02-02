// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages the registration of a Resource Provider - which allows access to the API's supported by this Resource Provider.
 *
 * > The Azure Provider will automatically register all of the Resource Providers which it supports on launch (unless opted-out using the `skipProviderRegistration` field within the provider block).
 *
 * !> **Note:** The errors returned from the Azure API when a Resource Provider is unregistered are unclear (example `API version '2019-01-01' was not found for 'Microsoft.Foo'`) - please ensure that all of the necessary Resource Providers you're using are registered - if in doubt **we strongly recommend letting the provider register these for you**.
 *
 * > **Note:** Adding or Removing a Preview Feature will re-register the Resource Provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceProviderRegistration("example", {});
 * ```
 * ### Registering A Preview Feature)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceProviderRegistration("example", {
 *     features: [{
 *         name: "AKS-DataPlaneAutoApprove",
 *         registered: true,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Resource Provider Registrations can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:core/resourceProviderRegistration:ResourceProviderRegistration example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights
 * ```
 */
export class ResourceProviderRegistration extends pulumi.CustomResource {
    /**
     * Get an existing ResourceProviderRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceProviderRegistrationState, opts?: pulumi.CustomResourceOptions): ResourceProviderRegistration {
        return new ResourceProviderRegistration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:core/resourceProviderRegistration:ResourceProviderRegistration';

    /**
     * Returns true if the given object is an instance of ResourceProviderRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceProviderRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceProviderRegistration.__pulumiType;
    }

    /**
     * A list of `feature` blocks as defined below.
     */
    public readonly features!: pulumi.Output<outputs.core.ResourceProviderRegistrationFeature[] | undefined>;
    /**
     * The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ResourceProviderRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceProviderRegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceProviderRegistrationArgs | ResourceProviderRegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceProviderRegistrationState | undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ResourceProviderRegistrationArgs | undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceProviderRegistration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceProviderRegistration resources.
 */
export interface ResourceProviderRegistrationState {
    /**
     * A list of `feature` blocks as defined below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.core.ResourceProviderRegistrationFeature>[]>;
    /**
     * The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceProviderRegistration resource.
 */
export interface ResourceProviderRegistrationArgs {
    /**
     * A list of `feature` blocks as defined below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.core.ResourceProviderRegistrationFeature>[]>;
    /**
     * The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
}
