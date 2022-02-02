// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a connection in an existing Virtual Network Gateway.
 *
 * ## Example Usage
 * ### Site-to-Site connection
 *
 * The following example shows a connection between an Azure virtual network
 * and an on-premises VPN device and network.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 * });
 * const onpremiseLocalNetworkGateway = new azure.network.LocalNetworkGateway("onpremiseLocalNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     gatewayAddress: "168.62.225.23",
 *     addressSpaces: ["10.1.1.0/24"],
 * });
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const exampleVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("exampleVirtualNetworkGateway", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     activeActive: false,
 *     enableBgp: false,
 *     sku: "Basic",
 *     ipConfigurations: [{
 *         publicIpAddressId: examplePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: exampleSubnet.id,
 *     }],
 * });
 * const onpremiseVirtualNetworkGatewayConnection = new azure.network.VirtualNetworkGatewayConnection("onpremiseVirtualNetworkGatewayConnection", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     type: "IPsec",
 *     virtualNetworkGatewayId: exampleVirtualNetworkGateway.id,
 *     localNetworkGatewayId: onpremiseLocalNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * ```
 * ### VNet-to-VNet connection
 *
 * The following example shows a connection between two Azure virtual network
 * in different locations/regions.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const usResourceGroup = new azure.core.ResourceGroup("usResourceGroup", {location: "East US"});
 * const usVirtualNetwork = new azure.network.VirtualNetwork("usVirtualNetwork", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     addressSpaces: ["10.0.0.0/16"],
 * });
 * const usGateway = new azure.network.Subnet("usGateway", {
 *     resourceGroupName: usResourceGroup.name,
 *     virtualNetworkName: usVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 * });
 * const usPublicIp = new azure.network.PublicIp("usPublicIp", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const usVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("usVirtualNetworkGateway", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     sku: "Basic",
 *     ipConfigurations: [{
 *         publicIpAddressId: usPublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: usGateway.id,
 *     }],
 * });
 * const europeResourceGroup = new azure.core.ResourceGroup("europeResourceGroup", {location: "West Europe"});
 * const europeVirtualNetwork = new azure.network.VirtualNetwork("europeVirtualNetwork", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     addressSpaces: ["10.1.0.0/16"],
 * });
 * const europeGateway = new azure.network.Subnet("europeGateway", {
 *     resourceGroupName: europeResourceGroup.name,
 *     virtualNetworkName: europeVirtualNetwork.name,
 *     addressPrefixes: ["10.1.1.0/24"],
 * });
 * const europePublicIp = new azure.network.PublicIp("europePublicIp", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     allocationMethod: "Dynamic",
 * });
 * const europeVirtualNetworkGateway = new azure.network.VirtualNetworkGateway("europeVirtualNetworkGateway", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     type: "Vpn",
 *     vpnType: "RouteBased",
 *     sku: "Basic",
 *     ipConfigurations: [{
 *         publicIpAddressId: europePublicIp.id,
 *         privateIpAddressAllocation: "Dynamic",
 *         subnetId: europeGateway.id,
 *     }],
 * });
 * const usToEurope = new azure.network.VirtualNetworkGatewayConnection("usToEurope", {
 *     location: usResourceGroup.location,
 *     resourceGroupName: usResourceGroup.name,
 *     type: "Vnet2Vnet",
 *     virtualNetworkGatewayId: usVirtualNetworkGateway.id,
 *     peerVirtualNetworkGatewayId: europeVirtualNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * const europeToUs = new azure.network.VirtualNetworkGatewayConnection("europeToUs", {
 *     location: europeResourceGroup.location,
 *     resourceGroupName: europeResourceGroup.name,
 *     type: "Vnet2Vnet",
 *     virtualNetworkGatewayId: europeVirtualNetworkGateway.id,
 *     peerVirtualNetworkGatewayId: usVirtualNetworkGateway.id,
 *     sharedKey: "4-v3ry-53cr37-1p53c-5h4r3d-k3y",
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Network Gateway Connections can be imported using their `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection exampleConnection /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1
 * ```
 */
export class VirtualNetworkGatewayConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualNetworkGatewayConnectionState, opts?: pulumi.CustomResourceOptions): VirtualNetworkGatewayConnection {
        return new VirtualNetworkGatewayConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection';

    /**
     * Returns true if the given object is an instance of VirtualNetworkGatewayConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkGatewayConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkGatewayConnection.__pulumiType;
    }

    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    public readonly authorizationKey!: pulumi.Output<string | undefined>;
    /**
     * Connection mode to use. Possible
     * values are `Default`, `InitiatorOnly` and `ResponderOnly`. Defaults to `Default`.
     * Changing this value will force a resource to be created.
     */
    public readonly connectionMode!: pulumi.Output<string | undefined>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note:** Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    public readonly connectionProtocol!: pulumi.Output<string>;
    /**
     * The dead peer detection timeout of this connection in seconds. Changing this forces a new resource to be created.
     */
    public readonly dpdTimeoutSeconds!: pulumi.Output<number | undefined>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    public readonly enableBgp!: pulumi.Output<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    public readonly expressRouteCircuitId!: pulumi.Output<string | undefined>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    public readonly expressRouteGatewayBypass!: pulumi.Output<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    public readonly ipsecPolicy!: pulumi.Output<outputs.network.VirtualNetworkGatewayConnectionIpsecPolicy | undefined>;
    /**
     * Use private local Azure IP for the connection. Changing this forces a new resource to be created.
     */
    public readonly localAzureIpAddressEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    public readonly localNetworkGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    public readonly peerVirtualNetworkGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    public readonly routingWeight!: pulumi.Output<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    public readonly sharedKey!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
     * Only one block can be defined for a connection.
     * For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
     */
    public readonly trafficSelectorPolicy!: pulumi.Output<outputs.network.VirtualNetworkGatewayConnectionTrafficSelectorPolicy | undefined>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    public readonly usePolicyBasedTrafficSelectors!: pulumi.Output<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    public readonly virtualNetworkGatewayId!: pulumi.Output<string>;

    /**
     * Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkGatewayConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkGatewayConnectionArgs | VirtualNetworkGatewayConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualNetworkGatewayConnectionState | undefined;
            resourceInputs["authorizationKey"] = state ? state.authorizationKey : undefined;
            resourceInputs["connectionMode"] = state ? state.connectionMode : undefined;
            resourceInputs["connectionProtocol"] = state ? state.connectionProtocol : undefined;
            resourceInputs["dpdTimeoutSeconds"] = state ? state.dpdTimeoutSeconds : undefined;
            resourceInputs["enableBgp"] = state ? state.enableBgp : undefined;
            resourceInputs["expressRouteCircuitId"] = state ? state.expressRouteCircuitId : undefined;
            resourceInputs["expressRouteGatewayBypass"] = state ? state.expressRouteGatewayBypass : undefined;
            resourceInputs["ipsecPolicy"] = state ? state.ipsecPolicy : undefined;
            resourceInputs["localAzureIpAddressEnabled"] = state ? state.localAzureIpAddressEnabled : undefined;
            resourceInputs["localNetworkGatewayId"] = state ? state.localNetworkGatewayId : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peerVirtualNetworkGatewayId"] = state ? state.peerVirtualNetworkGatewayId : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["routingWeight"] = state ? state.routingWeight : undefined;
            resourceInputs["sharedKey"] = state ? state.sharedKey : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["trafficSelectorPolicy"] = state ? state.trafficSelectorPolicy : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["usePolicyBasedTrafficSelectors"] = state ? state.usePolicyBasedTrafficSelectors : undefined;
            resourceInputs["virtualNetworkGatewayId"] = state ? state.virtualNetworkGatewayId : undefined;
        } else {
            const args = argsOrState as VirtualNetworkGatewayConnectionArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.virtualNetworkGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualNetworkGatewayId'");
            }
            resourceInputs["authorizationKey"] = args ? args.authorizationKey : undefined;
            resourceInputs["connectionMode"] = args ? args.connectionMode : undefined;
            resourceInputs["connectionProtocol"] = args ? args.connectionProtocol : undefined;
            resourceInputs["dpdTimeoutSeconds"] = args ? args.dpdTimeoutSeconds : undefined;
            resourceInputs["enableBgp"] = args ? args.enableBgp : undefined;
            resourceInputs["expressRouteCircuitId"] = args ? args.expressRouteCircuitId : undefined;
            resourceInputs["expressRouteGatewayBypass"] = args ? args.expressRouteGatewayBypass : undefined;
            resourceInputs["ipsecPolicy"] = args ? args.ipsecPolicy : undefined;
            resourceInputs["localAzureIpAddressEnabled"] = args ? args.localAzureIpAddressEnabled : undefined;
            resourceInputs["localNetworkGatewayId"] = args ? args.localNetworkGatewayId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peerVirtualNetworkGatewayId"] = args ? args.peerVirtualNetworkGatewayId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["routingWeight"] = args ? args.routingWeight : undefined;
            resourceInputs["sharedKey"] = args ? args.sharedKey : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficSelectorPolicy"] = args ? args.trafficSelectorPolicy : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["usePolicyBasedTrafficSelectors"] = args ? args.usePolicyBasedTrafficSelectors : undefined;
            resourceInputs["virtualNetworkGatewayId"] = args ? args.virtualNetworkGatewayId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualNetworkGatewayConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualNetworkGatewayConnection resources.
 */
export interface VirtualNetworkGatewayConnectionState {
    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    authorizationKey?: pulumi.Input<string>;
    /**
     * Connection mode to use. Possible
     * values are `Default`, `InitiatorOnly` and `ResponderOnly`. Defaults to `Default`.
     * Changing this value will force a resource to be created.
     */
    connectionMode?: pulumi.Input<string>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note:** Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    connectionProtocol?: pulumi.Input<string>;
    /**
     * The dead peer detection timeout of this connection in seconds. Changing this forces a new resource to be created.
     */
    dpdTimeoutSeconds?: pulumi.Input<number>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    enableBgp?: pulumi.Input<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    expressRouteCircuitId?: pulumi.Input<string>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    expressRouteGatewayBypass?: pulumi.Input<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    ipsecPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionIpsecPolicy>;
    /**
     * Use private local Azure IP for the connection. Changing this forces a new resource to be created.
     */
    localAzureIpAddressEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    localNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    peerVirtualNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    routingWeight?: pulumi.Input<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    sharedKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
     * Only one block can be defined for a connection.
     * For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
     */
    trafficSelectorPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionTrafficSelectorPolicy>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    usePolicyBasedTrafficSelectors?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    virtualNetworkGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
 */
export interface VirtualNetworkGatewayConnectionArgs {
    /**
     * The authorization key associated with the
     * Express Route Circuit. This field is required only if the type is an
     * ExpressRoute connection.
     */
    authorizationKey?: pulumi.Input<string>;
    /**
     * Connection mode to use. Possible
     * values are `Default`, `InitiatorOnly` and `ResponderOnly`. Defaults to `Default`.
     * Changing this value will force a resource to be created.
     */
    connectionMode?: pulumi.Input<string>;
    /**
     * The IKE protocol version to use. Possible
     * values are `IKEv1` and `IKEv2`. Defaults to `IKEv2`.
     * Changing this value will force a resource to be created.
     * > **Note:** Only valid for `IPSec` connections on virtual network gateways with SKU `VpnGw1`, `VpnGw2`, `VpnGw3`, `VpnGw1AZ`, `VpnGw2AZ` or `VpnGw3AZ`.
     */
    connectionProtocol?: pulumi.Input<string>;
    /**
     * The dead peer detection timeout of this connection in seconds. Changing this forces a new resource to be created.
     */
    dpdTimeoutSeconds?: pulumi.Input<number>;
    /**
     * If `true`, BGP (Border Gateway Protocol) is enabled
     * for this connection. Defaults to `false`.
     */
    enableBgp?: pulumi.Input<boolean>;
    /**
     * The ID of the Express Route Circuit
     * when creating an ExpressRoute connection (i.e. when `type` is `ExpressRoute`).
     * The Express Route Circuit can be in the same or in a different subscription.
     */
    expressRouteCircuitId?: pulumi.Input<string>;
    /**
     * If `true`, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
     */
    expressRouteGatewayBypass?: pulumi.Input<boolean>;
    /**
     * A `ipsecPolicy` block which is documented below.
     * Only a single policy can be defined for a connection. For details on
     * custom policies refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-ipsecikepolicy-rm-powershell).
     */
    ipsecPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionIpsecPolicy>;
    /**
     * Use private local Azure IP for the connection. Changing this forces a new resource to be created.
     */
    localAzureIpAddressEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the local network gateway
     * when creating Site-to-Site connection (i.e. when `type` is `IPsec`).
     */
    localNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The location/region where the connection is
     * located. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the connection. Changing the name forces a
     * new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the peer virtual
     * network gateway when creating a VNet-to-VNet connection (i.e. when `type`
     * is `Vnet2Vnet`). The peer Virtual Network Gateway can be in the same or
     * in a different subscription.
     */
    peerVirtualNetworkGatewayId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the connection Changing the name forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The routing weight. Defaults to `10`.
     */
    routingWeight?: pulumi.Input<number>;
    /**
     * The shared IPSec key. A key could be provided if a
     * Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
     */
    sharedKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `trafficSelectorPolicy` which allows to specify traffic selector policy proposal to be used in a virtual network gateway connection.
     * Only one block can be defined for a connection.
     * For details about traffic selectors refer to [the relevant section in the Azure documentation](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-connect-multiple-policybased-rm-ps).
     */
    trafficSelectorPolicy?: pulumi.Input<inputs.network.VirtualNetworkGatewayConnectionTrafficSelectorPolicy>;
    /**
     * The type of connection. Valid options are `IPsec`
     * (Site-to-Site), `ExpressRoute` (ExpressRoute), and `Vnet2Vnet` (VNet-to-VNet).
     * Each connection type requires different mandatory arguments (refer to the
     * examples above). Changing the connection type will force a new connection
     * to be created.
     */
    type: pulumi.Input<string>;
    /**
     * If `true`, policy-based traffic
     * selectors are enabled for this connection. Enabling policy-based traffic
     * selectors requires an `ipsecPolicy` block. Defaults to `false`.
     */
    usePolicyBasedTrafficSelectors?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Network Gateway
     * in which the connection will be created. Changing the gateway forces a new
     * resource to be created.
     */
    virtualNetworkGatewayId: pulumi.Input<string>;
}
