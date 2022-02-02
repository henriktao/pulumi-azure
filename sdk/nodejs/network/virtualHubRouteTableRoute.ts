// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Route in a Virtual Hub Route Table.
 *
 * > **Note:** Route table routes can managed with this resource, or in-line with the virtualHubRouteTable resource. Using both is not supported.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.5.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("exampleNetworkSecurityGroup", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.5.1.0/24"],
 * });
 * const exampleSubnetNetworkSecurityGroupAssociation = new azure.network.SubnetNetworkSecurityGroupAssociation("exampleSubnetNetworkSecurityGroupAssociation", {
 *     subnetId: exampleSubnet.id,
 *     networkSecurityGroupId: exampleNetworkSecurityGroup.id,
 * });
 * const exampleVirtualWan = new azure.network.VirtualWan("exampleVirtualWan", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 * });
 * const exampleVirtualHub = new azure.network.VirtualHub("exampleVirtualHub", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     virtualWanId: exampleVirtualWan.id,
 *     addressPrefix: "10.0.2.0/24",
 * });
 * const exampleVirtualHubRouteTable = new azure.network.VirtualHubRouteTable("exampleVirtualHubRouteTable", {
 *     virtualHubId: exampleVirtualHub.id,
 *     labels: ["label1"],
 * });
 * const exampleVirtualHubConnection = new azure.network.VirtualHubConnection("exampleVirtualHubConnection", {
 *     virtualHubId: exampleVirtualHub.id,
 *     remoteVirtualNetworkId: exampleVirtualNetwork.id,
 *     routing: {
 *         associatedRouteTableId: exampleVirtualHubRouteTable.id,
 *     },
 * });
 * const exampleVirtualHubRouteTableRoute = new azure.network.VirtualHubRouteTableRoute("exampleVirtualHubRouteTableRoute", {
 *     routeTableId: exampleVirtualHubRouteTable.id,
 *     destinationsType: "CIDR",
 *     destinations: ["10.0.0.0/16"],
 *     nextHopType: "ResourceId",
 *     nextHop: exampleVirtualHubConnection.id,
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Hub Route Table Routes can be imported using `<Route Table Resource Id>/routes/<Route Name>`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/virtualHubRouteTableRoute:VirtualHubRouteTableRoute example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1/routes/routeName
 * ```
 */
export class VirtualHubRouteTableRoute extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubRouteTableRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualHubRouteTableRouteState, opts?: pulumi.CustomResourceOptions): VirtualHubRouteTableRoute {
        return new VirtualHubRouteTableRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualHubRouteTableRoute:VirtualHubRouteTableRoute';

    /**
     * Returns true if the given object is an instance of VirtualHubRouteTableRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubRouteTableRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubRouteTableRoute.__pulumiType;
    }

    /**
     * A list of destination addresses for this route.
     */
    public readonly destinations!: pulumi.Output<string[]>;
    /**
     * The type of destinations. Possible values are `CIDR`, `ResourceId` and `Service`.
     */
    public readonly destinationsType!: pulumi.Output<string>;
    /**
     * The name which should be used for this route. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The next hop's resource ID.
     */
    public readonly nextHop!: pulumi.Output<string>;
    /**
     * The type of next hop. Currently the only possible value is `ResourceId`. Defaults to `ResourceId`.
     */
    public readonly nextHopType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Virtual Hub Route Table to link this route to. Changing this forces a new resource to be created.
     */
    public readonly routeTableId!: pulumi.Output<string>;

    /**
     * Create a VirtualHubRouteTableRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubRouteTableRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualHubRouteTableRouteArgs | VirtualHubRouteTableRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualHubRouteTableRouteState | undefined;
            resourceInputs["destinations"] = state ? state.destinations : undefined;
            resourceInputs["destinationsType"] = state ? state.destinationsType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextHop"] = state ? state.nextHop : undefined;
            resourceInputs["nextHopType"] = state ? state.nextHopType : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
        } else {
            const args = argsOrState as VirtualHubRouteTableRouteArgs | undefined;
            if ((!args || args.destinations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinations'");
            }
            if ((!args || args.destinationsType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationsType'");
            }
            if ((!args || args.nextHop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHop'");
            }
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            resourceInputs["destinations"] = args ? args.destinations : undefined;
            resourceInputs["destinationsType"] = args ? args.destinationsType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nextHop"] = args ? args.nextHop : undefined;
            resourceInputs["nextHopType"] = args ? args.nextHopType : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualHubRouteTableRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualHubRouteTableRoute resources.
 */
export interface VirtualHubRouteTableRouteState {
    /**
     * A list of destination addresses for this route.
     */
    destinations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of destinations. Possible values are `CIDR`, `ResourceId` and `Service`.
     */
    destinationsType?: pulumi.Input<string>;
    /**
     * The name which should be used for this route. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The next hop's resource ID.
     */
    nextHop?: pulumi.Input<string>;
    /**
     * The type of next hop. Currently the only possible value is `ResourceId`. Defaults to `ResourceId`.
     */
    nextHopType?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Hub Route Table to link this route to. Changing this forces a new resource to be created.
     */
    routeTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualHubRouteTableRoute resource.
 */
export interface VirtualHubRouteTableRouteArgs {
    /**
     * A list of destination addresses for this route.
     */
    destinations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of destinations. Possible values are `CIDR`, `ResourceId` and `Service`.
     */
    destinationsType: pulumi.Input<string>;
    /**
     * The name which should be used for this route. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The next hop's resource ID.
     */
    nextHop: pulumi.Input<string>;
    /**
     * The type of next hop. Currently the only possible value is `ResourceId`. Defaults to `ResourceId`.
     */
    nextHopType?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Hub Route Table to link this route to. Changing this forces a new resource to be created.
     */
    routeTableId: pulumi.Input<string>;
}
