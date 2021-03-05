// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a virtual network including any configured subnets. Each subnet can
 * optionally be configured with a security group to be associated with the subnet.
 *
 * > **NOTE on Virtual Networks and Subnet's:** This provider currently
 * provides both a standalone Subnet resource, and allows for Subnets to be defined in-line within the Virtual Network resource.
 * At this time you cannot use a Virtual Network with in-line Subnets in conjunction with any Subnet resources. Doing so will cause a conflict of Subnet configurations and will overwrite Subnet's.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleDdosProtectionPlan = new azure.network.DdosProtectionPlan("exampleDdosProtectionPlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 *     dnsServers: [
 *         "10.0.0.4",
 *         "10.0.0.5",
 *     ],
 *     ddosProtectionPlan: {
 *         id: exampleDdosProtectionPlan.id,
 *         enable: true,
 *     },
 *     subnets: [
 *         {
 *             name: "subnet1",
 *             addressPrefix: "10.0.1.0/24",
 *         },
 *         {
 *             name: "subnet2",
 *             addressPrefix: "10.0.2.0/24",
 *         },
 *         {
 *             name: "subnet3",
 *             addressPrefix: "10.0.3.0/24",
 *             securityGroup: exampleNetworkSecurityGroup.id,
 *         },
 *     ],
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Networks can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/virtualNetwork:VirtualNetwork exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
 * ```
 */
export class VirtualNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkState, opts?: pulumi.CustomResourceOptions): VirtualNetwork {
        return new VirtualNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualNetwork:VirtualNetwork';

    /**
     * Returns true if the given object is an instance of VirtualNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetwork.__pulumiType;
    }

    /**
     * The address space that is used the virtual network. You can supply more than one address space.
     */
    public readonly addressSpaces!: pulumi.Output<string[]>;
    /**
     * The BGP community attribute in format `<as-number>:<community-value>`.
     */
    public readonly bgpCommunity!: pulumi.Output<string | undefined>;
    /**
     * A `ddosProtectionPlan` block as documented below.
     */
    public readonly ddosProtectionPlan!: pulumi.Output<outputs.network.VirtualNetworkDdosProtectionPlan | undefined>;
    /**
     * List of IP addresses of DNS servers
     */
    public readonly dnsServers!: pulumi.Output<string[] | undefined>;
    /**
     * The GUID of the virtual network.
     */
    public /*out*/ readonly guid!: pulumi.Output<string>;
    /**
     * The location/region where the virtual network is created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the virtual network. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the virtual network.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
     */
    public readonly subnets!: pulumi.Output<outputs.network.VirtualNetworkSubnet[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
     */
    public readonly vmProtectionEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VirtualNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkArgs | VirtualNetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualNetworkState | undefined;
            inputs["addressSpaces"] = state ? state.addressSpaces : undefined;
            inputs["bgpCommunity"] = state ? state.bgpCommunity : undefined;
            inputs["ddosProtectionPlan"] = state ? state.ddosProtectionPlan : undefined;
            inputs["dnsServers"] = state ? state.dnsServers : undefined;
            inputs["guid"] = state ? state.guid : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["subnets"] = state ? state.subnets : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vmProtectionEnabled"] = state ? state.vmProtectionEnabled : undefined;
        } else {
            const args = argsOrState as VirtualNetworkArgs | undefined;
            if ((!args || args.addressSpaces === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressSpaces'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["addressSpaces"] = args ? args.addressSpaces : undefined;
            inputs["bgpCommunity"] = args ? args.bgpCommunity : undefined;
            inputs["ddosProtectionPlan"] = args ? args.ddosProtectionPlan : undefined;
            inputs["dnsServers"] = args ? args.dnsServers : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vmProtectionEnabled"] = args ? args.vmProtectionEnabled : undefined;
            inputs["guid"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VirtualNetwork.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetwork resources.
 */
export interface VirtualNetworkState {
    /**
     * The address space that is used the virtual network. You can supply more than one address space.
     */
    readonly addressSpaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The BGP community attribute in format `<as-number>:<community-value>`.
     */
    readonly bgpCommunity?: pulumi.Input<string>;
    /**
     * A `ddosProtectionPlan` block as documented below.
     */
    readonly ddosProtectionPlan?: pulumi.Input<inputs.network.VirtualNetworkDdosProtectionPlan>;
    /**
     * List of IP addresses of DNS servers
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The GUID of the virtual network.
     */
    readonly guid?: pulumi.Input<string>;
    /**
     * The location/region where the virtual network is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the virtual network. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the virtual network.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<inputs.network.VirtualNetworkSubnet>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
     */
    readonly vmProtectionEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VirtualNetwork resource.
 */
export interface VirtualNetworkArgs {
    /**
     * The address space that is used the virtual network. You can supply more than one address space.
     */
    readonly addressSpaces: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The BGP community attribute in format `<as-number>:<community-value>`.
     */
    readonly bgpCommunity?: pulumi.Input<string>;
    /**
     * A `ddosProtectionPlan` block as documented below.
     */
    readonly ddosProtectionPlan?: pulumi.Input<inputs.network.VirtualNetworkDdosProtectionPlan>;
    /**
     * List of IP addresses of DNS servers
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location/region where the virtual network is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the virtual network. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the virtual network.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Can be specified multiple times to define multiple subnets. Each `subnet` block supports fields documented below.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<inputs.network.VirtualNetworkSubnet>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to enable VM protection for all the subnets in this Virtual Network. Defaults to `false`.
     */
    readonly vmProtectionEnabled?: pulumi.Input<boolean>;
}
