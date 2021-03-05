// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Firewall.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Static",
 *     sku: "Standard",
 * });
 * const exampleFirewall = new azure.network.Firewall("exampleFirewall", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipConfigurations: [{
 *         name: "configuration",
 *         subnetId: exampleSubnet.id,
 *         publicIpAddressId: examplePublicIp.id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Azure Firewalls can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/firewall:Firewall example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/azureFirewalls/testfirewall
 * ```
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/firewall:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
     */
    public readonly dnsServers!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the Firewall Policy applied to this Firewall.
     */
    public readonly firewallPolicyId!: pulumi.Output<string | undefined>;
    /**
     * An `ipConfiguration` block as documented below.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.FirewallIpConfiguration[] | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A `managementIpConfiguration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnetId` in an existing block forces a new resource to be created.
     */
    public readonly managementIpConfiguration!: pulumi.Output<outputs.network.FirewallManagementIpConfiguration | undefined>;
    /**
     * Specifies the name of the Firewall. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
     */
    public readonly skuTier!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
     */
    public readonly threatIntelMode!: pulumi.Output<string | undefined>;
    /**
     * A `virtualHub` block as documented below.
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.FirewallVirtualHub | undefined>;
    /**
     * Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallArgs | FirewallState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallState | undefined;
            inputs["dnsServers"] = state ? state.dnsServers : undefined;
            inputs["firewallPolicyId"] = state ? state.firewallPolicyId : undefined;
            inputs["ipConfigurations"] = state ? state.ipConfigurations : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managementIpConfiguration"] = state ? state.managementIpConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["skuTier"] = state ? state.skuTier : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["threatIntelMode"] = state ? state.threatIntelMode : undefined;
            inputs["virtualHub"] = state ? state.virtualHub : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as FirewallArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dnsServers"] = args ? args.dnsServers : undefined;
            inputs["firewallPolicyId"] = args ? args.firewallPolicyId : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementIpConfiguration"] = args ? args.managementIpConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["skuTier"] = args ? args.skuTier : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatIntelMode"] = args ? args.threatIntelMode : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["zones"] = args ? args.zones : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Firewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Firewall Policy applied to this Firewall.
     */
    readonly firewallPolicyId?: pulumi.Input<string>;
    /**
     * An `ipConfiguration` block as documented below.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.FirewallIpConfiguration>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `managementIpConfiguration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnetId` in an existing block forces a new resource to be created.
     */
    readonly managementIpConfiguration?: pulumi.Input<inputs.network.FirewallManagementIpConfiguration>;
    /**
     * Specifies the name of the Firewall. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
     */
    readonly skuTier?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
     */
    readonly threatIntelMode?: pulumi.Input<string>;
    /**
     * A `virtualHub` block as documented below.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.FirewallVirtualHub>;
    /**
     * Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * A list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.
     */
    readonly dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Firewall Policy applied to this Firewall.
     */
    readonly firewallPolicyId?: pulumi.Input<string>;
    /**
     * An `ipConfiguration` block as documented below.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.FirewallIpConfiguration>[]>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A `managementIpConfiguration` block as documented below, which allows force-tunnelling of traffic to be performed by the firewall. Adding or removing this block or changing the `subnetId` in an existing block forces a new resource to be created.
     */
    readonly managementIpConfiguration?: pulumi.Input<inputs.network.FirewallManagementIpConfiguration>;
    /**
     * Specifies the name of the Firewall. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Sku name of the Firewall. Possible values are `AZFW_Hub` and `AZFW_VNet`.  Changing this forces a new resource to be created.
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * Sku tier of the Firewall. Possible values are `Premium` and `Standard`.  Changing this forces a new resource to be created.
     */
    readonly skuTier?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The operation mode for threat intelligence-based filtering. Possible values are: `Off`, `Alert`,`Deny` and `""`(empty string). Defaults to `Alert`.
     */
    readonly threatIntelMode?: pulumi.Input<string>;
    /**
     * A `virtualHub` block as documented below.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.FirewallVirtualHub>;
    /**
     * Specifies the availability zones in which the Azure Firewall should be created. Changing this forces a new resource to be created.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
