// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Replica Set for an Active Directory Domain Service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuread from "@pulumi/azuread";
 *
 * const primaryResourceGroup = new azure.core.ResourceGroup("primaryResourceGroup", {location: "West Europe"});
 * const primaryVirtualNetwork = new azure.network.VirtualNetwork("primaryVirtualNetwork", {
 *     location: primaryResourceGroup.location,
 *     resourceGroupName: primaryResourceGroup.name,
 *     addressSpaces: ["10.0.1.0/16"],
 * });
 * const primarySubnet = new azure.network.Subnet("primarySubnet", {
 *     resourceGroupName: primaryResourceGroup.name,
 *     virtualNetworkName: primaryVirtualNetwork.name,
 *     addressPrefixes: ["10.0.1.0/24"],
 * });
 * const primaryNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("primaryNetworkSecurityGroup", {
 *     location: primaryResourceGroup.location,
 *     resourceGroupName: primaryResourceGroup.name,
 *     securityRules: [
 *         {
 *             name: "AllowSyncWithAzureAD",
 *             priority: 101,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "443",
 *             sourceAddressPrefix: "AzureActiveDirectoryDomainServices",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowRD",
 *             priority: 201,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "3389",
 *             sourceAddressPrefix: "CorpNetSaw",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowPSRemoting",
 *             priority: 301,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "5986",
 *             sourceAddressPrefix: "AzureActiveDirectoryDomainServices",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowLDAPS",
 *             priority: 401,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "636",
 *             sourceAddressPrefix: "*",
 *             destinationAddressPrefix: "*",
 *         },
 *     ],
 * });
 * const primarySubnetNetworkSecurityGroupAssociation = new azure.network.SubnetNetworkSecurityGroupAssociation("primarySubnetNetworkSecurityGroupAssociation", {
 *     subnetId: primarySubnet.id,
 *     networkSecurityGroupId: primaryNetworkSecurityGroup.id,
 * });
 * const dcAdmins = new azuread.Group("dcAdmins", {});
 * const adminUser = new azuread.User("adminUser", {
 *     userPrincipalName: `dc-admin@$hashicorp-example.net`,
 *     displayName: "DC Administrator",
 *     password: "Pa55w0Rd!!1",
 * });
 * const adminGroupMember = new azuread.GroupMember("adminGroupMember", {
 *     groupObjectId: dcAdmins.objectId,
 *     memberObjectId: adminUser.objectId,
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {applicationId: "2565bd9d-da50-47d4-8b85-4c97f669dc36"});
 * // published app for domain services
 * const aadds = new azure.core.ResourceGroup("aadds", {location: "westeurope"});
 * const exampleService = new azure.domainservices.Service("exampleService", {
 *     location: aadds.location,
 *     resourceGroupName: aadds.name,
 *     domainName: "widgetslogin.net",
 *     sku: "Enterprise",
 *     filteredSyncEnabled: false,
 *     initialReplicaSet: {
 *         location: primaryVirtualNetwork.location,
 *         subnetId: primarySubnet.id,
 *     },
 *     notifications: {
 *         additionalRecipients: [
 *             "notifyA@example.net",
 *             "notifyB@example.org",
 *         ],
 *         notifyDcAdmins: true,
 *         notifyGlobalAdmins: true,
 *     },
 *     security: {
 *         syncKerberosPasswords: true,
 *         syncNtlmPasswords: true,
 *         syncOnPremPasswords: true,
 *     },
 *     tags: {
 *         Environment: "prod",
 *     },
 * }, {
 *     dependsOn: [
 *         exampleServicePrincipal,
 *         primarySubnetNetworkSecurityGroupAssociation,
 *     ],
 * });
 * const replicaResourceGroup = new azure.core.ResourceGroup("replicaResourceGroup", {location: "North Europe"});
 * const replicaVirtualNetwork = new azure.network.VirtualNetwork("replicaVirtualNetwork", {
 *     location: replicaResourceGroup.location,
 *     resourceGroupName: replicaResourceGroup.name,
 *     addressSpaces: ["10.20.0.0/16"],
 * });
 * const aaddsReplicaSubnet = new azure.network.Subnet("aaddsReplicaSubnet", {
 *     resourceGroupName: replicaResourceGroup.name,
 *     virtualNetworkName: replicaVirtualNetwork.name,
 *     addressPrefixes: ["10.20.0.0/24"],
 * });
 * const aaddsReplicaNetworkSecurityGroup = new azure.network.NetworkSecurityGroup("aaddsReplicaNetworkSecurityGroup", {
 *     location: replicaResourceGroup.location,
 *     resourceGroupName: replicaResourceGroup.name,
 *     securityRules: [
 *         {
 *             name: "AllowSyncWithAzureAD",
 *             priority: 101,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "443",
 *             sourceAddressPrefix: "AzureActiveDirectoryDomainServices",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowRD",
 *             priority: 201,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "3389",
 *             sourceAddressPrefix: "CorpNetSaw",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowPSRemoting",
 *             priority: 301,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "5986",
 *             sourceAddressPrefix: "AzureActiveDirectoryDomainServices",
 *             destinationAddressPrefix: "*",
 *         },
 *         {
 *             name: "AllowLDAPS",
 *             priority: 401,
 *             direction: "Inbound",
 *             access: "Allow",
 *             protocol: "Tcp",
 *             sourcePortRange: "*",
 *             destinationPortRange: "636",
 *             sourceAddressPrefix: "*",
 *             destinationAddressPrefix: "*",
 *         },
 *     ],
 * });
 * const replicaSubnetNetworkSecurityGroupAssociation = new azure.network.SubnetNetworkSecurityGroupAssociation("replicaSubnetNetworkSecurityGroupAssociation", {
 *     subnetId: aaddsReplicaSubnet.id,
 *     networkSecurityGroupId: aaddsReplicaNetworkSecurityGroup.id,
 * });
 * const primaryReplica = new azure.network.VirtualNetworkPeering("primaryReplica", {
 *     resourceGroupName: primaryVirtualNetwork.resourceGroupName,
 *     virtualNetworkName: primaryVirtualNetwork.name,
 *     remoteVirtualNetworkId: replicaVirtualNetwork.id,
 *     allowForwardedTraffic: true,
 *     allowGatewayTransit: false,
 *     allowVirtualNetworkAccess: true,
 *     useRemoteGateways: false,
 * });
 * const replicaPrimary = new azure.network.VirtualNetworkPeering("replicaPrimary", {
 *     resourceGroupName: replicaVirtualNetwork.resourceGroupName,
 *     virtualNetworkName: replicaVirtualNetwork.name,
 *     remoteVirtualNetworkId: primaryVirtualNetwork.id,
 *     allowForwardedTraffic: true,
 *     allowGatewayTransit: false,
 *     allowVirtualNetworkAccess: true,
 *     useRemoteGateways: false,
 * });
 * const replicaVirtualNetworkDnsServers = new azure.network.VirtualNetworkDnsServers("replicaVirtualNetworkDnsServers", {
 *     virtualNetworkId: replicaVirtualNetwork.id,
 *     dnsServers: exampleService.initialReplicaSet.apply(initialReplicaSet => initialReplicaSet.domainControllerIpAddresses),
 * });
 * const replicaReplicaSet = new azure.domainservices.ReplicaSet("replicaReplicaSet", {
 *     domainServiceId: exampleService.id,
 *     location: replicaResourceGroup.location,
 *     subnetId: aaddsReplicaSubnet.id,
 * }, {
 *     dependsOn: [
 *         replicaSubnetNetworkSecurityGroupAssociation,
 *         primaryReplica,
 *         replicaPrimary,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Domain Service Replica Sets can be imported using the resource ID of the parent Domain Service and the Replica Set ID, e.g.
 *
 * ```sh
 *  $ pulumi import azure:domainservices/replicaSet:ReplicaSet example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ReplicaSet extends pulumi.CustomResource {
    /**
     * Get an existing ReplicaSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicaSetState, opts?: pulumi.CustomResourceOptions): ReplicaSet {
        return new ReplicaSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:domainservices/replicaSet:ReplicaSet';

    /**
     * Returns true if the given object is an instance of ReplicaSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicaSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicaSet.__pulumiType;
    }

    /**
     * A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
     */
    public /*out*/ readonly domainControllerIpAddresses!: pulumi.Output<string[]>;
    /**
     * The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
     */
    public readonly domainServiceId!: pulumi.Output<string>;
    /**
     * The publicly routable IP address for the domain controllers in this Replica Set.
     */
    public /*out*/ readonly externalAccessIpAddress!: pulumi.Output<string>;
    /**
     * The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The current service status for the replica set.
     */
    public /*out*/ readonly serviceStatus!: pulumi.Output<string>;
    /**
     * The ID of the subnet in which to place this Replica Set.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a ReplicaSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicaSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicaSetArgs | ReplicaSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicaSetState | undefined;
            resourceInputs["domainControllerIpAddresses"] = state ? state.domainControllerIpAddresses : undefined;
            resourceInputs["domainServiceId"] = state ? state.domainServiceId : undefined;
            resourceInputs["externalAccessIpAddress"] = state ? state.externalAccessIpAddress : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["serviceStatus"] = state ? state.serviceStatus : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as ReplicaSetArgs | undefined;
            if ((!args || args.domainServiceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainServiceId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["domainServiceId"] = args ? args.domainServiceId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["domainControllerIpAddresses"] = undefined /*out*/;
            resourceInputs["externalAccessIpAddress"] = undefined /*out*/;
            resourceInputs["serviceStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicaSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicaSet resources.
 */
export interface ReplicaSetState {
    /**
     * A list of subnet IP addresses for the domain controllers in this Replica Set, typically two.
     */
    domainControllerIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
     */
    domainServiceId?: pulumi.Input<string>;
    /**
     * The publicly routable IP address for the domain controllers in this Replica Set.
     */
    externalAccessIpAddress?: pulumi.Input<string>;
    /**
     * The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The current service status for the replica set.
     */
    serviceStatus?: pulumi.Input<string>;
    /**
     * The ID of the subnet in which to place this Replica Set.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicaSet resource.
 */
export interface ReplicaSetArgs {
    /**
     * The ID of the Domain Service for which to create this Replica Set. Changing this forces a new resource to be created.
     */
    domainServiceId: pulumi.Input<string>;
    /**
     * The Azure location where this Replica Set should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the subnet in which to place this Replica Set.
     */
    subnetId: pulumi.Input<string>;
}
