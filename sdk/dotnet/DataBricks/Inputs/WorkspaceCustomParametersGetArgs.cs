// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataBricks.Inputs
{

    public sealed class WorkspaceCustomParametersGetArgs : Pulumi.ResourceArgs
    {
        [Input("machineLearningWorkspaceId")]
        public Input<string>? MachineLearningWorkspaceId { get; set; }

        /// <summary>
        /// Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets. Defaults to `nat-gateway`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("natGatewayName")]
        public Input<string>? NatGatewayName { get; set; }

        /// <summary>
        /// Are public IP Addresses not allowed? Possible values are `true` or `false`. Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("noPublicIp")]
        public Input<bool>? NoPublicIp { get; set; }

        /// <summary>
        /// The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("privateSubnetName")]
        public Input<string>? PrivateSubnetName { get; set; }

        /// <summary>
        /// The resource ID of the `azure.network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `private_subnet_name` field. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("privateSubnetNetworkSecurityGroupAssociationId")]
        public Input<string>? PrivateSubnetNetworkSecurityGroupAssociationId { get; set; }

        /// <summary>
        /// Name of the Public IP for No Public IP workspace with managed vNet. Defaults to `nat-gw-public-ip`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publicIpName")]
        public Input<string>? PublicIpName { get; set; }

        /// <summary>
        /// The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publicSubnetName")]
        public Input<string>? PublicSubnetName { get; set; }

        /// <summary>
        /// The resource ID of the `azure.network.SubnetNetworkSecurityGroupAssociation` resource which is referred to by the `public_subnet_name` field. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("publicSubnetNetworkSecurityGroupAssociationId")]
        public Input<string>? PublicSubnetNetworkSecurityGroupAssociationId { get; set; }

        /// <summary>
        /// Default Databricks File Storage account name. Defaults to a randomized name(e.g. `dbstoragel6mfeghoe5kxu`). Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        /// <summary>
        /// Storage account SKU name. Possible values include `Standard_LRS`, `Standard_GRS`, `Standard_RAGRS`, `Standard_GZRS`, `Standard_RAGZRS`, `Standard_ZRS`, `Premium_LRS` or `Premium_ZRS`. Defaults to `Standard_GRS`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountSkuName")]
        public Input<string>? StorageAccountSkuName { get; set; }

        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        /// <summary>
        /// Address prefix for Managed virtual network. Defaults to `10.139`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("vnetAddressPrefix")]
        public Input<string>? VnetAddressPrefix { get; set; }

        public WorkspaceCustomParametersGetArgs()
        {
        }
    }
}
