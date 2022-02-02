// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the association between a Network Interface and a Load Balancer's NAT Rule.
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
 *     addressPrefixes: ["10.0.2.0/24"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Static",
 * });
 * const exampleLoadBalancer = new azure.lb.LoadBalancer("exampleLoadBalancer", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     frontendIpConfigurations: [{
 *         name: "primary",
 *         publicIpAddressId: examplePublicIp.id,
 *     }],
 * });
 * const exampleNatRule = new azure.lb.NatRule("exampleNatRule", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     loadbalancerId: exampleLoadBalancer.id,
 *     protocol: "Tcp",
 *     frontendPort: 3389,
 *     backendPort: 3389,
 *     frontendIpConfigurationName: "primary",
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("exampleNetworkInterface", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipConfigurations: [{
 *         name: "testconfiguration1",
 *         subnetId: exampleSubnet.id,
 *         privateIpAddressAllocation: "Dynamic",
 *     }],
 * });
 * const exampleNetworkInterfaceNatRuleAssociation = new azure.network.NetworkInterfaceNatRuleAssociation("exampleNetworkInterfaceNatRuleAssociation", {
 *     networkInterfaceId: exampleNetworkInterface.id,
 *     ipConfigurationName: "testconfiguration1",
 *     natRuleId: exampleNatRule.id,
 * });
 * ```
 *
 * ## Import
 *
 * Associations between Network Interfaces and Load Balancer NAT Rule can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
 * ```
 */
export class NetworkInterfaceNatRuleAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceNatRuleAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceNatRuleAssociationState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceNatRuleAssociation {
        return new NetworkInterfaceNatRuleAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceNatRuleAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceNatRuleAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceNatRuleAssociation.__pulumiType;
    }

    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
     */
    public readonly ipConfigurationName!: pulumi.Output<string>;
    /**
     * The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    public readonly natRuleId!: pulumi.Output<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceNatRuleAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceNatRuleAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceNatRuleAssociationArgs | NetworkInterfaceNatRuleAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceNatRuleAssociationState | undefined;
            resourceInputs["ipConfigurationName"] = state ? state.ipConfigurationName : undefined;
            resourceInputs["natRuleId"] = state ? state.natRuleId : undefined;
            resourceInputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceNatRuleAssociationArgs | undefined;
            if ((!args || args.ipConfigurationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipConfigurationName'");
            }
            if ((!args || args.natRuleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natRuleId'");
            }
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            resourceInputs["ipConfigurationName"] = args ? args.ipConfigurationName : undefined;
            resourceInputs["natRuleId"] = args ? args.natRuleId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkInterfaceNatRuleAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceNatRuleAssociation resources.
 */
export interface NetworkInterfaceNatRuleAssociationState {
    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
     */
    ipConfigurationName?: pulumi.Input<string>;
    /**
     * The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    natRuleId?: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceNatRuleAssociation resource.
 */
export interface NetworkInterfaceNatRuleAssociationArgs {
    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
     */
    ipConfigurationName: pulumi.Input<string>;
    /**
     * The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    natRuleId: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId: pulumi.Input<string>;
}
