// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Avs
{
    public static class GetPrivateCloud
    {
        public static Task<GetPrivateCloudResult> InvokeAsync(GetPrivateCloudArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateCloudResult>("azure:avs/getPrivateCloud:getPrivateCloud", args ?? new GetPrivateCloudArgs(), options.WithDefaults());

        public static Output<GetPrivateCloudResult> Invoke(GetPrivateCloudInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPrivateCloudResult>("azure:avs/getPrivateCloud:getPrivateCloud", args ?? new GetPrivateCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateCloudArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Vmware Private Cloud.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Vmware Private Cloud exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateCloudArgs()
        {
        }
    }

    public sealed class GetPrivateCloudInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this Vmware Private Cloud.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Vmware Private Cloud exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPrivateCloudInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateCloudResult
    {
        /// <summary>
        /// A `circuit` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPrivateCloudCircuitResult> Circuits;
        /// <summary>
        /// The endpoint for the HCX Cloud Manager.
        /// </summary>
        public readonly string HcxCloudManagerEndpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Is the Vmware Private Cluster connected to the internet?
        /// </summary>
        public readonly bool InternetConnectionEnabled;
        /// <summary>
        /// The Azure Region where the Vmware Private Cloud exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// A `management_cluster` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPrivateCloudManagementClusterResult> ManagementClusters;
        /// <summary>
        /// The network used to access vCenter Server and NSX-T Manager.
        /// </summary>
        public readonly string ManagementSubnetCidr;
        public readonly string Name;
        /// <summary>
        /// The subnet cidr of the Vmware Private Cloud.
        /// </summary>
        public readonly string NetworkSubnetCidr;
        /// <summary>
        /// The thumbprint of the NSX-T Manager SSL certificate.
        /// </summary>
        public readonly string NsxtCertificateThumbprint;
        /// <summary>
        /// The endpoint for the NSX-T Data Center manager.
        /// </summary>
        public readonly string NsxtManagerEndpoint;
        /// <summary>
        /// The network which isused for virtual machine cold migration, cloning, and snapshot migration.
        /// </summary>
        public readonly string ProvisioningSubnetCidr;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The Name of the SKU used for this Private Cloud.
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// A mapping of tags assigned to the Vmware Private Cloud.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The thumbprint of the vCenter Server SSL certificate.
        /// </summary>
        public readonly string VcenterCertificateThumbprint;
        /// <summary>
        /// The endpoint for Virtual Center Server Appliance.
        /// </summary>
        public readonly string VcsaEndpoint;
        /// <summary>
        /// The network which is used for live migration of virtual machines.
        /// </summary>
        public readonly string VmotionSubnetCidr;

        [OutputConstructor]
        private GetPrivateCloudResult(
            ImmutableArray<Outputs.GetPrivateCloudCircuitResult> circuits,

            string hcxCloudManagerEndpoint,

            string id,

            bool internetConnectionEnabled,

            string location,

            ImmutableArray<Outputs.GetPrivateCloudManagementClusterResult> managementClusters,

            string managementSubnetCidr,

            string name,

            string networkSubnetCidr,

            string nsxtCertificateThumbprint,

            string nsxtManagerEndpoint,

            string provisioningSubnetCidr,

            string resourceGroupName,

            string skuName,

            ImmutableDictionary<string, string> tags,

            string vcenterCertificateThumbprint,

            string vcsaEndpoint,

            string vmotionSubnetCidr)
        {
            Circuits = circuits;
            HcxCloudManagerEndpoint = hcxCloudManagerEndpoint;
            Id = id;
            InternetConnectionEnabled = internetConnectionEnabled;
            Location = location;
            ManagementClusters = managementClusters;
            ManagementSubnetCidr = managementSubnetCidr;
            Name = name;
            NetworkSubnetCidr = networkSubnetCidr;
            NsxtCertificateThumbprint = nsxtCertificateThumbprint;
            NsxtManagerEndpoint = nsxtManagerEndpoint;
            ProvisioningSubnetCidr = provisioningSubnetCidr;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            Tags = tags;
            VcenterCertificateThumbprint = vcenterCertificateThumbprint;
            VcsaEndpoint = vcsaEndpoint;
            VmotionSubnetCidr = vmotionSubnetCidr;
        }
    }
}
