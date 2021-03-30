// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Subscription Template Deployment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.SubscriptionTemplateDeployment("example", {
 *     location: "West Europe",
 *     templateContent: ` {
 *    "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
 *    "contentVersion": "1.0.0.0",
 *    "parameters": {},
 *    "variables": {},
 *    "resources": [
 *      {
 *        "type": "Microsoft.Resources/resourceGroups",
 *        "apiVersion": "2018-05-01",
 *        "location": "West Europe",
 *        "name": "some-resource-group",
 *        "properties": {}
 *      }
 *    ]
 *  }
 *  `,
 * });
 * ```
 *
 * ## Import
 *
 * Subscription Template Deployments can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:core/subscriptionTemplateDeployment:SubscriptionTemplateDeployment example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/template1
 * ```
 */
export class SubscriptionTemplateDeployment extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionTemplateDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionTemplateDeploymentState, opts?: pulumi.CustomResourceOptions): SubscriptionTemplateDeployment {
        return new SubscriptionTemplateDeployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:core/subscriptionTemplateDeployment:SubscriptionTemplateDeployment';

    /**
     * Returns true if the given object is an instance of SubscriptionTemplateDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionTemplateDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionTemplateDeployment.__pulumiType;
    }

    /**
     * The Debug Level which should be used for this Subscription Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    public readonly debugLevel!: pulumi.Output<string | undefined>;
    /**
     * The Azure Region where the Subscription Template Deployment should exist. Changing this forces a new Subscription Template Deployment to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Subscription Template Deployment. Changing this forces a new Subscription Template Deployment to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The JSON Content of the Outputs of the ARM Template Deployment.
     */
    public /*out*/ readonly outputContent!: pulumi.Output<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    public readonly parametersContent!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Subscription Template Deployment.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The contents of the ARM Template which should be deployed into this Subscription.
     */
    public readonly templateContent!: pulumi.Output<string>;
    /**
     * The ID of the Template Spec Version to deploy into the Subscription. Cannot be specified with `templateContent`.
     */
    public readonly templateSpecVersionId!: pulumi.Output<string | undefined>;

    /**
     * Create a SubscriptionTemplateDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SubscriptionTemplateDeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionTemplateDeploymentArgs | SubscriptionTemplateDeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionTemplateDeploymentState | undefined;
            inputs["debugLevel"] = state ? state.debugLevel : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outputContent"] = state ? state.outputContent : undefined;
            inputs["parametersContent"] = state ? state.parametersContent : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["templateContent"] = state ? state.templateContent : undefined;
            inputs["templateSpecVersionId"] = state ? state.templateSpecVersionId : undefined;
        } else {
            const args = argsOrState as SubscriptionTemplateDeploymentArgs | undefined;
            inputs["debugLevel"] = args ? args.debugLevel : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parametersContent"] = args ? args.parametersContent : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateContent"] = args ? args.templateContent : undefined;
            inputs["templateSpecVersionId"] = args ? args.templateSpecVersionId : undefined;
            inputs["outputContent"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SubscriptionTemplateDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionTemplateDeployment resources.
 */
export interface SubscriptionTemplateDeploymentState {
    /**
     * The Debug Level which should be used for this Subscription Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    readonly debugLevel?: pulumi.Input<string>;
    /**
     * The Azure Region where the Subscription Template Deployment should exist. Changing this forces a new Subscription Template Deployment to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Subscription Template Deployment. Changing this forces a new Subscription Template Deployment to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The JSON Content of the Outputs of the ARM Template Deployment.
     */
    readonly outputContent?: pulumi.Input<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    readonly parametersContent?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Subscription Template Deployment.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The contents of the ARM Template which should be deployed into this Subscription.
     */
    readonly templateContent?: pulumi.Input<string>;
    /**
     * The ID of the Template Spec Version to deploy into the Subscription. Cannot be specified with `templateContent`.
     */
    readonly templateSpecVersionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionTemplateDeployment resource.
 */
export interface SubscriptionTemplateDeploymentArgs {
    /**
     * The Debug Level which should be used for this Subscription Template Deployment. Possible values are `none`, `requestContent`, `responseContent` and `requestContent, responseContent`.
     */
    readonly debugLevel?: pulumi.Input<string>;
    /**
     * The Azure Region where the Subscription Template Deployment should exist. Changing this forces a new Subscription Template Deployment to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Subscription Template Deployment. Changing this forces a new Subscription Template Deployment to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The contents of the ARM Template parameters file - containing a JSON list of parameters.
     */
    readonly parametersContent?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Subscription Template Deployment.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The contents of the ARM Template which should be deployed into this Subscription.
     */
    readonly templateContent?: pulumi.Input<string>;
    /**
     * The ID of the Template Spec Version to deploy into the Subscription. Cannot be specified with `templateContent`.
     */
    readonly templateSpecVersionId?: pulumi.Input<string>;
}
