// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the association between a Network Interface and a Network Security Group.
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
 * const exampleNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
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
 * const exampleNetworkInterfaceSecurityGroupAssociation = new azure.network.NetworkInterfaceSecurityGroupAssociation("exampleNetworkInterfaceSecurityGroupAssociation", {
 *     networkInterfaceId: exampleNetworkInterface.id,
 *     networkSecurityGroupId: exampleNetworkSecurityGroup.id,
 * });
 * ```
 *
 * ## Import
 *
 * Associations between Network Interfaces and Network Security Group can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation association1 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkSecurityGroups/group1"
 * ```
 */
export class NetworkInterfaceSecurityGroupAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceSecurityGroupAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceSecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceSecurityGroupAssociation {
        return new NetworkInterfaceSecurityGroupAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceSecurityGroupAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceSecurityGroupAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceSecurityGroupAssociation.__pulumiType;
    }

    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
     */
    public readonly networkSecurityGroupId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceSecurityGroupAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceSecurityGroupAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceSecurityGroupAssociationArgs | NetworkInterfaceSecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceSecurityGroupAssociationState | undefined;
            resourceInputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            resourceInputs["networkSecurityGroupId"] = state ? state.networkSecurityGroupId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceSecurityGroupAssociationArgs | undefined;
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            if ((!args || args.networkSecurityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkSecurityGroupId'");
            }
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["networkSecurityGroupId"] = args ? args.networkSecurityGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkInterfaceSecurityGroupAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAssociation resources.
 */
export interface NetworkInterfaceSecurityGroupAssociationState {
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
     */
    networkSecurityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceSecurityGroupAssociation resource.
 */
export interface NetworkInterfaceSecurityGroupAssociationArgs {
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    networkInterfaceId: pulumi.Input<string>;
    /**
     * The ID of the Network Security Group which should be attached to the Network Interface. Changing this forces a new resource to be created.
     */
    networkSecurityGroupId: pulumi.Input<string>;
}
