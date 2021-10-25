// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./applicationGateway";
export * from "./applicationSecurityGroup";
export * from "./bgpConnection";
export * from "./ddosProtectionPlan";
export * from "./expressRouteCircuit";
export * from "./expressRouteCircuitAuthorization";
export * from "./expressRouteCircuitConnection";
export * from "./expressRouteCircuitPeering";
export * from "./expressRouteConnection";
export * from "./expressRouteGateway";
export * from "./expressRoutePort";
export * from "./firewall";
export * from "./firewallApplicationRuleCollection";
export * from "./firewallNatRuleCollection";
export * from "./firewallNetworkRuleCollection";
export * from "./firewallPolicy";
export * from "./firewallPolicyRuleCollectionGroup";
export * from "./getApplicationGateway";
export * from "./getApplicationSecurityGroup";
export * from "./getExpressRouteCircuit";
export * from "./getFirewall";
export * from "./getFirewallPolicy";
export * from "./getGatewayConnection";
export * from "./getIpGroup";
export * from "./getLocalNetworkGateway";
export * from "./getNatGateway";
export * from "./getNetworkDdosProtectionPlan";
export * from "./getNetworkInterface";
export * from "./getNetworkSecurityGroup";
export * from "./getNetworkWatcher";
export * from "./getPublicIP";
export * from "./getPublicIPs";
export * from "./getPublicIpPrefix";
export * from "./getRouteFilter";
export * from "./getRouteTable";
export * from "./getServiceTags";
export * from "./getSubnet";
export * from "./getTrafficManager";
export * from "./getTrafficManagerProfile";
export * from "./getVirtualHub";
export * from "./getVirtualNetwork";
export * from "./getVirtualNetworkGateway";
export * from "./getVirtualWan";
export * from "./getVpnGateway";
export * from "./ipgroup";
export * from "./localNetworkGateway";
export * from "./natGateway";
export * from "./natGatewayPublicIpAssociation";
export * from "./natGatewayPublicIpPrefixAssociation";
export * from "./networkConnectionMonitor";
export * from "./networkInterface";
export * from "./networkInterfaceApplicationGatewayBackendAddressPoolAssociation";
export * from "./networkInterfaceApplicationSecurityGroupAssociation";
export * from "./networkInterfaceBackendAddressPoolAssociation";
export * from "./networkInterfaceNatRuleAssociation";
export * from "./networkInterfaceSecurityGroupAssociation";
export * from "./networkPacketCapture";
export * from "./networkSecurityGroup";
export * from "./networkSecurityRule";
export * from "./networkWatcher";
export * from "./networkWatcherFlowLog";
export * from "./packetCapture";
export * from "./pointToPointVpnGateway";
export * from "./profile";
export * from "./publicIp";
export * from "./publicIpPrefix";
export * from "./route";
export * from "./routeFilter";
export * from "./routeTable";
export * from "./securityPartnerProvider";
export * from "./subnet";
export * from "./subnetNatGatewayAssociation";
export * from "./subnetNetworkSecurityGroupAssociation";
export * from "./subnetRouteTableAssociation";
export * from "./subnetServiceEndpointStoragePolicy";
export * from "./trafficManagerEndpoint";
export * from "./trafficManagerProfile";
export * from "./virtualHub";
export * from "./virtualHubConnection";
export * from "./virtualHubIp";
export * from "./virtualHubRouteTable";
export * from "./virtualHubRouteTableRoute";
export * from "./virtualNetwork";
export * from "./virtualNetworkDnsServers";
export * from "./virtualNetworkGateway";
export * from "./virtualNetworkGatewayConnection";
export * from "./virtualNetworkPeering";
export * from "./virtualWan";
export * from "./vpnGateway";
export * from "./vpnGatewayConnection";
export * from "./vpnServerConfiguration";
export * from "./vpnSite";

// Import resources to register:
import { ApplicationGateway } from "./applicationGateway";
import { ApplicationSecurityGroup } from "./applicationSecurityGroup";
import { BgpConnection } from "./bgpConnection";
import { DdosProtectionPlan } from "./ddosProtectionPlan";
import { ExpressRouteCircuit } from "./expressRouteCircuit";
import { ExpressRouteCircuitAuthorization } from "./expressRouteCircuitAuthorization";
import { ExpressRouteCircuitConnection } from "./expressRouteCircuitConnection";
import { ExpressRouteCircuitPeering } from "./expressRouteCircuitPeering";
import { ExpressRouteConnection } from "./expressRouteConnection";
import { ExpressRouteGateway } from "./expressRouteGateway";
import { ExpressRoutePort } from "./expressRoutePort";
import { Firewall } from "./firewall";
import { FirewallApplicationRuleCollection } from "./firewallApplicationRuleCollection";
import { FirewallNatRuleCollection } from "./firewallNatRuleCollection";
import { FirewallNetworkRuleCollection } from "./firewallNetworkRuleCollection";
import { FirewallPolicy } from "./firewallPolicy";
import { FirewallPolicyRuleCollectionGroup } from "./firewallPolicyRuleCollectionGroup";
import { IPGroup } from "./ipgroup";
import { LocalNetworkGateway } from "./localNetworkGateway";
import { NatGateway } from "./natGateway";
import { NatGatewayPublicIpAssociation } from "./natGatewayPublicIpAssociation";
import { NatGatewayPublicIpPrefixAssociation } from "./natGatewayPublicIpPrefixAssociation";
import { NetworkConnectionMonitor } from "./networkConnectionMonitor";
import { NetworkInterface } from "./networkInterface";
import { NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation } from "./networkInterfaceApplicationGatewayBackendAddressPoolAssociation";
import { NetworkInterfaceApplicationSecurityGroupAssociation } from "./networkInterfaceApplicationSecurityGroupAssociation";
import { NetworkInterfaceBackendAddressPoolAssociation } from "./networkInterfaceBackendAddressPoolAssociation";
import { NetworkInterfaceNatRuleAssociation } from "./networkInterfaceNatRuleAssociation";
import { NetworkInterfaceSecurityGroupAssociation } from "./networkInterfaceSecurityGroupAssociation";
import { NetworkPacketCapture } from "./networkPacketCapture";
import { NetworkSecurityGroup } from "./networkSecurityGroup";
import { NetworkSecurityRule } from "./networkSecurityRule";
import { NetworkWatcher } from "./networkWatcher";
import { NetworkWatcherFlowLog } from "./networkWatcherFlowLog";
import { PacketCapture } from "./packetCapture";
import { PointToPointVpnGateway } from "./pointToPointVpnGateway";
import { Profile } from "./profile";
import { PublicIp } from "./publicIp";
import { PublicIpPrefix } from "./publicIpPrefix";
import { Route } from "./route";
import { RouteFilter } from "./routeFilter";
import { RouteTable } from "./routeTable";
import { SecurityPartnerProvider } from "./securityPartnerProvider";
import { Subnet } from "./subnet";
import { SubnetNatGatewayAssociation } from "./subnetNatGatewayAssociation";
import { SubnetNetworkSecurityGroupAssociation } from "./subnetNetworkSecurityGroupAssociation";
import { SubnetRouteTableAssociation } from "./subnetRouteTableAssociation";
import { SubnetServiceEndpointStoragePolicy } from "./subnetServiceEndpointStoragePolicy";
import { TrafficManagerEndpoint } from "./trafficManagerEndpoint";
import { TrafficManagerProfile } from "./trafficManagerProfile";
import { VirtualHub } from "./virtualHub";
import { VirtualHubConnection } from "./virtualHubConnection";
import { VirtualHubIp } from "./virtualHubIp";
import { VirtualHubRouteTable } from "./virtualHubRouteTable";
import { VirtualHubRouteTableRoute } from "./virtualHubRouteTableRoute";
import { VirtualNetwork } from "./virtualNetwork";
import { VirtualNetworkDnsServers } from "./virtualNetworkDnsServers";
import { VirtualNetworkGateway } from "./virtualNetworkGateway";
import { VirtualNetworkGatewayConnection } from "./virtualNetworkGatewayConnection";
import { VirtualNetworkPeering } from "./virtualNetworkPeering";
import { VirtualWan } from "./virtualWan";
import { VpnGateway } from "./vpnGateway";
import { VpnGatewayConnection } from "./vpnGatewayConnection";
import { VpnServerConfiguration } from "./vpnServerConfiguration";
import { VpnSite } from "./vpnSite";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:network/applicationGateway:ApplicationGateway":
                return new ApplicationGateway(name, <any>undefined, { urn })
            case "azure:network/applicationSecurityGroup:ApplicationSecurityGroup":
                return new ApplicationSecurityGroup(name, <any>undefined, { urn })
            case "azure:network/bgpConnection:BgpConnection":
                return new BgpConnection(name, <any>undefined, { urn })
            case "azure:network/ddosProtectionPlan:DdosProtectionPlan":
                return new DdosProtectionPlan(name, <any>undefined, { urn })
            case "azure:network/expressRouteCircuit:ExpressRouteCircuit":
                return new ExpressRouteCircuit(name, <any>undefined, { urn })
            case "azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization":
                return new ExpressRouteCircuitAuthorization(name, <any>undefined, { urn })
            case "azure:network/expressRouteCircuitConnection:ExpressRouteCircuitConnection":
                return new ExpressRouteCircuitConnection(name, <any>undefined, { urn })
            case "azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering":
                return new ExpressRouteCircuitPeering(name, <any>undefined, { urn })
            case "azure:network/expressRouteConnection:ExpressRouteConnection":
                return new ExpressRouteConnection(name, <any>undefined, { urn })
            case "azure:network/expressRouteGateway:ExpressRouteGateway":
                return new ExpressRouteGateway(name, <any>undefined, { urn })
            case "azure:network/expressRoutePort:ExpressRoutePort":
                return new ExpressRoutePort(name, <any>undefined, { urn })
            case "azure:network/firewall:Firewall":
                return new Firewall(name, <any>undefined, { urn })
            case "azure:network/firewallApplicationRuleCollection:FirewallApplicationRuleCollection":
                return new FirewallApplicationRuleCollection(name, <any>undefined, { urn })
            case "azure:network/firewallNatRuleCollection:FirewallNatRuleCollection":
                return new FirewallNatRuleCollection(name, <any>undefined, { urn })
            case "azure:network/firewallNetworkRuleCollection:FirewallNetworkRuleCollection":
                return new FirewallNetworkRuleCollection(name, <any>undefined, { urn })
            case "azure:network/firewallPolicy:FirewallPolicy":
                return new FirewallPolicy(name, <any>undefined, { urn })
            case "azure:network/firewallPolicyRuleCollectionGroup:FirewallPolicyRuleCollectionGroup":
                return new FirewallPolicyRuleCollectionGroup(name, <any>undefined, { urn })
            case "azure:network/iPGroup:IPGroup":
                return new IPGroup(name, <any>undefined, { urn })
            case "azure:network/localNetworkGateway:LocalNetworkGateway":
                return new LocalNetworkGateway(name, <any>undefined, { urn })
            case "azure:network/natGateway:NatGateway":
                return new NatGateway(name, <any>undefined, { urn })
            case "azure:network/natGatewayPublicIpAssociation:NatGatewayPublicIpAssociation":
                return new NatGatewayPublicIpAssociation(name, <any>undefined, { urn })
            case "azure:network/natGatewayPublicIpPrefixAssociation:NatGatewayPublicIpPrefixAssociation":
                return new NatGatewayPublicIpPrefixAssociation(name, <any>undefined, { urn })
            case "azure:network/networkConnectionMonitor:NetworkConnectionMonitor":
                return new NetworkConnectionMonitor(name, <any>undefined, { urn })
            case "azure:network/networkInterface:NetworkInterface":
                return new NetworkInterface(name, <any>undefined, { urn })
            case "azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation":
                return new NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(name, <any>undefined, { urn })
            case "azure:network/networkInterfaceApplicationSecurityGroupAssociation:NetworkInterfaceApplicationSecurityGroupAssociation":
                return new NetworkInterfaceApplicationSecurityGroupAssociation(name, <any>undefined, { urn })
            case "azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation":
                return new NetworkInterfaceBackendAddressPoolAssociation(name, <any>undefined, { urn })
            case "azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation":
                return new NetworkInterfaceNatRuleAssociation(name, <any>undefined, { urn })
            case "azure:network/networkInterfaceSecurityGroupAssociation:NetworkInterfaceSecurityGroupAssociation":
                return new NetworkInterfaceSecurityGroupAssociation(name, <any>undefined, { urn })
            case "azure:network/networkPacketCapture:NetworkPacketCapture":
                return new NetworkPacketCapture(name, <any>undefined, { urn })
            case "azure:network/networkSecurityGroup:NetworkSecurityGroup":
                return new NetworkSecurityGroup(name, <any>undefined, { urn })
            case "azure:network/networkSecurityRule:NetworkSecurityRule":
                return new NetworkSecurityRule(name, <any>undefined, { urn })
            case "azure:network/networkWatcher:NetworkWatcher":
                return new NetworkWatcher(name, <any>undefined, { urn })
            case "azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog":
                return new NetworkWatcherFlowLog(name, <any>undefined, { urn })
            case "azure:network/packetCapture:PacketCapture":
                return new PacketCapture(name, <any>undefined, { urn })
            case "azure:network/pointToPointVpnGateway:PointToPointVpnGateway":
                return new PointToPointVpnGateway(name, <any>undefined, { urn })
            case "azure:network/profile:Profile":
                return new Profile(name, <any>undefined, { urn })
            case "azure:network/publicIp:PublicIp":
                return new PublicIp(name, <any>undefined, { urn })
            case "azure:network/publicIpPrefix:PublicIpPrefix":
                return new PublicIpPrefix(name, <any>undefined, { urn })
            case "azure:network/route:Route":
                return new Route(name, <any>undefined, { urn })
            case "azure:network/routeFilter:RouteFilter":
                return new RouteFilter(name, <any>undefined, { urn })
            case "azure:network/routeTable:RouteTable":
                return new RouteTable(name, <any>undefined, { urn })
            case "azure:network/securityPartnerProvider:SecurityPartnerProvider":
                return new SecurityPartnerProvider(name, <any>undefined, { urn })
            case "azure:network/subnet:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "azure:network/subnetNatGatewayAssociation:SubnetNatGatewayAssociation":
                return new SubnetNatGatewayAssociation(name, <any>undefined, { urn })
            case "azure:network/subnetNetworkSecurityGroupAssociation:SubnetNetworkSecurityGroupAssociation":
                return new SubnetNetworkSecurityGroupAssociation(name, <any>undefined, { urn })
            case "azure:network/subnetRouteTableAssociation:SubnetRouteTableAssociation":
                return new SubnetRouteTableAssociation(name, <any>undefined, { urn })
            case "azure:network/subnetServiceEndpointStoragePolicy:SubnetServiceEndpointStoragePolicy":
                return new SubnetServiceEndpointStoragePolicy(name, <any>undefined, { urn })
            case "azure:network/trafficManagerEndpoint:TrafficManagerEndpoint":
                return new TrafficManagerEndpoint(name, <any>undefined, { urn })
            case "azure:network/trafficManagerProfile:TrafficManagerProfile":
                return new TrafficManagerProfile(name, <any>undefined, { urn })
            case "azure:network/virtualHub:VirtualHub":
                return new VirtualHub(name, <any>undefined, { urn })
            case "azure:network/virtualHubConnection:VirtualHubConnection":
                return new VirtualHubConnection(name, <any>undefined, { urn })
            case "azure:network/virtualHubIp:VirtualHubIp":
                return new VirtualHubIp(name, <any>undefined, { urn })
            case "azure:network/virtualHubRouteTable:VirtualHubRouteTable":
                return new VirtualHubRouteTable(name, <any>undefined, { urn })
            case "azure:network/virtualHubRouteTableRoute:VirtualHubRouteTableRoute":
                return new VirtualHubRouteTableRoute(name, <any>undefined, { urn })
            case "azure:network/virtualNetwork:VirtualNetwork":
                return new VirtualNetwork(name, <any>undefined, { urn })
            case "azure:network/virtualNetworkDnsServers:VirtualNetworkDnsServers":
                return new VirtualNetworkDnsServers(name, <any>undefined, { urn })
            case "azure:network/virtualNetworkGateway:VirtualNetworkGateway":
                return new VirtualNetworkGateway(name, <any>undefined, { urn })
            case "azure:network/virtualNetworkGatewayConnection:VirtualNetworkGatewayConnection":
                return new VirtualNetworkGatewayConnection(name, <any>undefined, { urn })
            case "azure:network/virtualNetworkPeering:VirtualNetworkPeering":
                return new VirtualNetworkPeering(name, <any>undefined, { urn })
            case "azure:network/virtualWan:VirtualWan":
                return new VirtualWan(name, <any>undefined, { urn })
            case "azure:network/vpnGateway:VpnGateway":
                return new VpnGateway(name, <any>undefined, { urn })
            case "azure:network/vpnGatewayConnection:VpnGatewayConnection":
                return new VpnGatewayConnection(name, <any>undefined, { urn })
            case "azure:network/vpnServerConfiguration:VpnServerConfiguration":
                return new VpnServerConfiguration(name, <any>undefined, { urn })
            case "azure:network/vpnSite:VpnSite":
                return new VpnSite(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "network/applicationGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/applicationSecurityGroup", _module)
pulumi.runtime.registerResourceModule("azure", "network/bgpConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/ddosProtectionPlan", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteCircuit", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteCircuitAuthorization", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteCircuitConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteCircuitPeering", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRouteGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/expressRoutePort", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewall", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewallApplicationRuleCollection", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewallNatRuleCollection", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewallNetworkRuleCollection", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewallPolicy", _module)
pulumi.runtime.registerResourceModule("azure", "network/firewallPolicyRuleCollectionGroup", _module)
pulumi.runtime.registerResourceModule("azure", "network/iPGroup", _module)
pulumi.runtime.registerResourceModule("azure", "network/localNetworkGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/natGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/natGatewayPublicIpAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/natGatewayPublicIpPrefixAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkConnectionMonitor", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterface", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterfaceApplicationSecurityGroupAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterfaceBackendAddressPoolAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterfaceNatRuleAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkInterfaceSecurityGroupAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkPacketCapture", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkSecurityGroup", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkSecurityRule", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkWatcher", _module)
pulumi.runtime.registerResourceModule("azure", "network/networkWatcherFlowLog", _module)
pulumi.runtime.registerResourceModule("azure", "network/packetCapture", _module)
pulumi.runtime.registerResourceModule("azure", "network/pointToPointVpnGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/profile", _module)
pulumi.runtime.registerResourceModule("azure", "network/publicIp", _module)
pulumi.runtime.registerResourceModule("azure", "network/publicIpPrefix", _module)
pulumi.runtime.registerResourceModule("azure", "network/route", _module)
pulumi.runtime.registerResourceModule("azure", "network/routeFilter", _module)
pulumi.runtime.registerResourceModule("azure", "network/routeTable", _module)
pulumi.runtime.registerResourceModule("azure", "network/securityPartnerProvider", _module)
pulumi.runtime.registerResourceModule("azure", "network/subnet", _module)
pulumi.runtime.registerResourceModule("azure", "network/subnetNatGatewayAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/subnetNetworkSecurityGroupAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/subnetRouteTableAssociation", _module)
pulumi.runtime.registerResourceModule("azure", "network/subnetServiceEndpointStoragePolicy", _module)
pulumi.runtime.registerResourceModule("azure", "network/trafficManagerEndpoint", _module)
pulumi.runtime.registerResourceModule("azure", "network/trafficManagerProfile", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualHub", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualHubConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualHubIp", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualHubRouteTable", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualHubRouteTableRoute", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualNetwork", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualNetworkDnsServers", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualNetworkGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualNetworkGatewayConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualNetworkPeering", _module)
pulumi.runtime.registerResourceModule("azure", "network/virtualWan", _module)
pulumi.runtime.registerResourceModule("azure", "network/vpnGateway", _module)
pulumi.runtime.registerResourceModule("azure", "network/vpnGatewayConnection", _module)
pulumi.runtime.registerResourceModule("azure", "network/vpnServerConfiguration", _module)
pulumi.runtime.registerResourceModule("azure", "network/vpnSite", _module)
