// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Load Balancer NAT Rule. 
 * 
 * > **NOTE:** This resource cannot be used with with virtual machine scale sets, instead use the `azure.lb.NatPool` resource.
 * 
 * > **NOTE** When using this resource, the Load Balancer needs to have a FrontEnd IP Configuration Attached
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West US",
 *     name: "LoadBalancerRG",
 * });
 * const testPublicIp = new azure.network.PublicIp("test", {
 *     allocationMethod: "Static",
 *     location: "West US",
 *     name: "PublicIPForLB",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testLoadBalancer = new azure.lb.LoadBalancer("test", {
 *     frontendIpConfigurations: [{
 *         name: "PublicIPAddress",
 *         publicIpAddressId: testPublicIp.id,
 *     }],
 *     location: "West US",
 *     name: "TestLoadBalancer",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * const testNatRule = new azure.lb.NatRule("test", {
 *     backendPort: 3389,
 *     frontendIpConfigurationName: "PublicIPAddress",
 *     frontendPort: 3389,
 *     loadbalancerId: testLoadBalancer.id,
 *     name: "RDPAccess",
 *     protocol: "Tcp",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/lb_nat_rule.html.markdown.
 */
export class NatRule extends pulumi.CustomResource {
    /**
     * Get an existing NatRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatRuleState, opts?: pulumi.CustomResourceOptions): NatRule {
        return new NatRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:lb/natRule:NatRule';

    /**
     * Returns true if the given object is an instance of NatRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatRule.__pulumiType;
    }

    public /*out*/ readonly backendIpConfigurationId!: pulumi.Output<string>;
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    public readonly backendPort!: pulumi.Output<number>;
    /**
     * Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group.
     */
    public readonly enableFloatingIp!: pulumi.Output<boolean>;
    public /*out*/ readonly frontendIpConfigurationId!: pulumi.Output<string>;
    /**
     * The name of the frontend IP configuration exposing this rule.
     */
    public readonly frontendIpConfigurationName!: pulumi.Output<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    public readonly frontendPort!: pulumi.Output<number>;
    /**
     * The ID of the Load Balancer in which to create the NAT Rule.
     */
    public readonly loadbalancerId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the NAT Rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a NatRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatRuleArgs | NatRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NatRuleState | undefined;
            inputs["backendIpConfigurationId"] = state ? state.backendIpConfigurationId : undefined;
            inputs["backendPort"] = state ? state.backendPort : undefined;
            inputs["enableFloatingIp"] = state ? state.enableFloatingIp : undefined;
            inputs["frontendIpConfigurationId"] = state ? state.frontendIpConfigurationId : undefined;
            inputs["frontendIpConfigurationName"] = state ? state.frontendIpConfigurationName : undefined;
            inputs["frontendPort"] = state ? state.frontendPort : undefined;
            inputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as NatRuleArgs | undefined;
            if (!args || args.backendPort === undefined) {
                throw new Error("Missing required property 'backendPort'");
            }
            if (!args || args.frontendIpConfigurationName === undefined) {
                throw new Error("Missing required property 'frontendIpConfigurationName'");
            }
            if (!args || args.frontendPort === undefined) {
                throw new Error("Missing required property 'frontendPort'");
            }
            if (!args || args.loadbalancerId === undefined) {
                throw new Error("Missing required property 'loadbalancerId'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendPort"] = args ? args.backendPort : undefined;
            inputs["enableFloatingIp"] = args ? args.enableFloatingIp : undefined;
            inputs["frontendIpConfigurationName"] = args ? args.frontendIpConfigurationName : undefined;
            inputs["frontendPort"] = args ? args.frontendPort : undefined;
            inputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["backendIpConfigurationId"] = undefined /*out*/;
            inputs["frontendIpConfigurationId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NatRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatRule resources.
 */
export interface NatRuleState {
    readonly backendIpConfigurationId?: pulumi.Input<string>;
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    readonly backendPort?: pulumi.Input<number>;
    /**
     * Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group.
     */
    readonly enableFloatingIp?: pulumi.Input<boolean>;
    readonly frontendIpConfigurationId?: pulumi.Input<string>;
    /**
     * The name of the frontend IP configuration exposing this rule.
     */
    readonly frontendIpConfigurationName?: pulumi.Input<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    readonly frontendPort?: pulumi.Input<number>;
    /**
     * The ID of the Load Balancer in which to create the NAT Rule.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the NAT Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NatRule resource.
 */
export interface NatRuleArgs {
    /**
     * The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.
     */
    readonly backendPort: pulumi.Input<number>;
    /**
     * Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group.
     */
    readonly enableFloatingIp?: pulumi.Input<boolean>;
    /**
     * The name of the frontend IP configuration exposing this rule.
     */
    readonly frontendIpConfigurationName: pulumi.Input<string>;
    /**
     * The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.
     */
    readonly frontendPort: pulumi.Input<number>;
    /**
     * The ID of the Load Balancer in which to create the NAT Rule.
     */
    readonly loadbalancerId: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the NAT Rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The transport protocol for the external endpoint. Possible values are `Udp`, `Tcp` or `All`.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
