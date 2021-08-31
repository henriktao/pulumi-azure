// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Firewall Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.network.FirewallPolicy("example", {
 *     location: "West Europe",
 *     resourceGroupName: "example",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Policies can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/firewallPolicy:FirewallPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/firewallPolicies/policy1
 * ```
 */
export class FirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicyState, opts?: pulumi.CustomResourceOptions): FirewallPolicy {
        return new FirewallPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/firewallPolicy:FirewallPolicy';

    /**
     * Returns true if the given object is an instance of FirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy.__pulumiType;
    }

    /**
     * The ID of the base Firewall Policy.
     */
    public readonly basePolicyId!: pulumi.Output<string | undefined>;
    /**
     * A list of reference to child Firewall Policies of this Firewall Policy.
     */
    public /*out*/ readonly childPolicies!: pulumi.Output<string[]>;
    /**
     * A `dns` block as defined below.
     */
    public readonly dns!: pulumi.Output<outputs.network.FirewallPolicyDns | undefined>;
    /**
     * A list of references to Azure Firewalls that this Firewall Policy is associated with.
     */
    public /*out*/ readonly firewalls!: pulumi.Output<string[]>;
    /**
     * An `identity` block as defined below. Changing this forces a new Firewall Policy to be created.
     */
    public readonly identity!: pulumi.Output<outputs.network.FirewallPolicyIdentity | undefined>;
    /**
     * A `intrusionDetection` block as defined below.
     */
    public readonly intrusionDetection!: pulumi.Output<outputs.network.FirewallPolicyIntrusionDetection | undefined>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of private IP ranges to which traffic will not be SNAT.
     */
    public readonly privateIpRanges!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
     */
    public /*out*/ readonly ruleCollectionGroups!: pulumi.Output<string[]>;
    /**
     * The SKU Tier of the Firewall Policy. Possible values are `Standard`, `Premium`. Changing this forces a new Firewall Policy to be created.
     */
    public readonly sku!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    public readonly threatIntelligenceAllowlist!: pulumi.Output<outputs.network.FirewallPolicyThreatIntelligenceAllowlist | undefined>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    public readonly threatIntelligenceMode!: pulumi.Output<string | undefined>;
    /**
     * A `tlsCertificate` block as defined below.
     */
    public readonly tlsCertificate!: pulumi.Output<outputs.network.FirewallPolicyTlsCertificate | undefined>;

    /**
     * Create a FirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicyArgs | FirewallPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallPolicyState | undefined;
            inputs["basePolicyId"] = state ? state.basePolicyId : undefined;
            inputs["childPolicies"] = state ? state.childPolicies : undefined;
            inputs["dns"] = state ? state.dns : undefined;
            inputs["firewalls"] = state ? state.firewalls : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["intrusionDetection"] = state ? state.intrusionDetection : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateIpRanges"] = state ? state.privateIpRanges : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["ruleCollectionGroups"] = state ? state.ruleCollectionGroups : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["threatIntelligenceAllowlist"] = state ? state.threatIntelligenceAllowlist : undefined;
            inputs["threatIntelligenceMode"] = state ? state.threatIntelligenceMode : undefined;
            inputs["tlsCertificate"] = state ? state.tlsCertificate : undefined;
        } else {
            const args = argsOrState as FirewallPolicyArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["basePolicyId"] = args ? args.basePolicyId : undefined;
            inputs["dns"] = args ? args.dns : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["intrusionDetection"] = args ? args.intrusionDetection : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateIpRanges"] = args ? args.privateIpRanges : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatIntelligenceAllowlist"] = args ? args.threatIntelligenceAllowlist : undefined;
            inputs["threatIntelligenceMode"] = args ? args.threatIntelligenceMode : undefined;
            inputs["tlsCertificate"] = args ? args.tlsCertificate : undefined;
            inputs["childPolicies"] = undefined /*out*/;
            inputs["firewalls"] = undefined /*out*/;
            inputs["ruleCollectionGroups"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicy resources.
 */
export interface FirewallPolicyState {
    /**
     * The ID of the base Firewall Policy.
     */
    basePolicyId?: pulumi.Input<string>;
    /**
     * A list of reference to child Firewall Policies of this Firewall Policy.
     */
    childPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `dns` block as defined below.
     */
    dns?: pulumi.Input<inputs.network.FirewallPolicyDns>;
    /**
     * A list of references to Azure Firewalls that this Firewall Policy is associated with.
     */
    firewalls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An `identity` block as defined below. Changing this forces a new Firewall Policy to be created.
     */
    identity?: pulumi.Input<inputs.network.FirewallPolicyIdentity>;
    /**
     * A `intrusionDetection` block as defined below.
     */
    intrusionDetection?: pulumi.Input<inputs.network.FirewallPolicyIntrusionDetection>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of private IP ranges to which traffic will not be SNAT.
     */
    privateIpRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * A list of references to Firewall Policy Rule Collection Groups that belongs to this Firewall Policy.
     */
    ruleCollectionGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The SKU Tier of the Firewall Policy. Possible values are `Standard`, `Premium`. Changing this forces a new Firewall Policy to be created.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    threatIntelligenceAllowlist?: pulumi.Input<inputs.network.FirewallPolicyThreatIntelligenceAllowlist>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    threatIntelligenceMode?: pulumi.Input<string>;
    /**
     * A `tlsCertificate` block as defined below.
     */
    tlsCertificate?: pulumi.Input<inputs.network.FirewallPolicyTlsCertificate>;
}

/**
 * The set of arguments for constructing a FirewallPolicy resource.
 */
export interface FirewallPolicyArgs {
    /**
     * The ID of the base Firewall Policy.
     */
    basePolicyId?: pulumi.Input<string>;
    /**
     * A `dns` block as defined below.
     */
    dns?: pulumi.Input<inputs.network.FirewallPolicyDns>;
    /**
     * An `identity` block as defined below. Changing this forces a new Firewall Policy to be created.
     */
    identity?: pulumi.Input<inputs.network.FirewallPolicyIdentity>;
    /**
     * A `intrusionDetection` block as defined below.
     */
    intrusionDetection?: pulumi.Input<inputs.network.FirewallPolicyIntrusionDetection>;
    /**
     * The Azure Region where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name which should be used for this Firewall Policy. Changing this forces a new Firewall Policy to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of private IP ranges to which traffic will not be SNAT.
     */
    privateIpRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource Group where the Firewall Policy should exist. Changing this forces a new Firewall Policy to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU Tier of the Firewall Policy. Possible values are `Standard`, `Premium`. Changing this forces a new Firewall Policy to be created.
     */
    sku?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Firewall Policy.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `threatIntelligenceAllowlist` block as defined below.
     */
    threatIntelligenceAllowlist?: pulumi.Input<inputs.network.FirewallPolicyThreatIntelligenceAllowlist>;
    /**
     * The operation mode for Threat Intelligence. Possible values are `Alert`, `Deny` and `Off`. Defaults to `Alert`.
     */
    threatIntelligenceMode?: pulumi.Input<string>;
    /**
     * A `tlsCertificate` block as defined below.
     */
    tlsCertificate?: pulumi.Input<inputs.network.FirewallPolicyTlsCertificate>;
}
