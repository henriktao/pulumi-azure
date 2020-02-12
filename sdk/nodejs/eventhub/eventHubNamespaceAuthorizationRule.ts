// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Authorization Rule for an Event Hub Namespace.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub_namespace_authorization_rule.html.markdown.
 */
export class EventHubNamespaceAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing EventHubNamespaceAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventHubNamespaceAuthorizationRuleState, opts?: pulumi.CustomResourceOptions): EventHubNamespaceAuthorizationRule {
        return new EventHubNamespaceAuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventhub/eventHubNamespaceAuthorizationRule:EventHubNamespaceAuthorizationRule';

    /**
     * Returns true if the given object is an instance of EventHubNamespaceAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventHubNamespaceAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventHubNamespaceAuthorizationRule.__pulumiType;
    }

    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    public readonly listen!: pulumi.Output<boolean | undefined>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    public readonly manage!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The Primary Connection String for the Authorization Rule.
     */
    public /*out*/ readonly primaryConnectionString!: pulumi.Output<string>;
    /**
     * The Primary Key for the Authorization Rule.
     */
    public /*out*/ readonly primaryKey!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Connection String for the Authorization Rule.
     */
    public /*out*/ readonly secondaryConnectionString!: pulumi.Output<string>;
    /**
     * The Secondary Key for the Authorization Rule.
     */
    public /*out*/ readonly secondaryKey!: pulumi.Output<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    public readonly send!: pulumi.Output<boolean | undefined>;

    /**
     * Create a EventHubNamespaceAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventHubNamespaceAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventHubNamespaceAuthorizationRuleArgs | EventHubNamespaceAuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventHubNamespaceAuthorizationRuleState | undefined;
            inputs["listen"] = state ? state.listen : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["manage"] = state ? state.manage : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceName"] = state ? state.namespaceName : undefined;
            inputs["primaryConnectionString"] = state ? state.primaryConnectionString : undefined;
            inputs["primaryKey"] = state ? state.primaryKey : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryConnectionString"] = state ? state.secondaryConnectionString : undefined;
            inputs["secondaryKey"] = state ? state.secondaryKey : undefined;
            inputs["send"] = state ? state.send : undefined;
        } else {
            const args = argsOrState as EventHubNamespaceAuthorizationRuleArgs | undefined;
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["listen"] = args ? args.listen : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["manage"] = args ? args.manage : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["send"] = args ? args.send : undefined;
            inputs["primaryConnectionString"] = undefined /*out*/;
            inputs["primaryKey"] = undefined /*out*/;
            inputs["secondaryConnectionString"] = undefined /*out*/;
            inputs["secondaryKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventHubNamespaceAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventHubNamespaceAuthorizationRule resources.
 */
export interface EventHubNamespaceAuthorizationRuleState {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    readonly listen?: pulumi.Input<boolean>;
    readonly location?: pulumi.Input<string>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    readonly manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The Primary Connection String for the Authorization Rule.
     */
    readonly primaryConnectionString?: pulumi.Input<string>;
    /**
     * The Primary Key for the Authorization Rule.
     */
    readonly primaryKey?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Connection String for the Authorization Rule.
     */
    readonly secondaryConnectionString?: pulumi.Input<string>;
    /**
     * The Secondary Key for the Authorization Rule.
     */
    readonly secondaryKey?: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    readonly send?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a EventHubNamespaceAuthorizationRule resource.
 */
export interface EventHubNamespaceAuthorizationRuleArgs {
    /**
     * Grants listen access to this this Authorization Rule. Defaults to `false`.
     */
    readonly listen?: pulumi.Input<boolean>;
    readonly location?: pulumi.Input<string>;
    /**
     * Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
     */
    readonly manage?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Grants send access to this this Authorization Rule. Defaults to `false`.
     */
    readonly send?: pulumi.Input<boolean>;
}
