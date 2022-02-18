// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class GetKubernetesClusterAgentPoolProfileResult
    {
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// The number of Agents (VM's) in the Pool.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// If the auto-scaler is enabled.
        /// </summary>
        public readonly bool EnableAutoScaling;
        /// <summary>
        /// If the Public IPs for the nodes in this Agent Pool are enabled.
        /// </summary>
        public readonly bool EnableNodePublicIp;
        /// <summary>
        /// Maximum number of nodes for auto-scaling
        /// </summary>
        public readonly int MaxCount;
        /// <summary>
        /// The maximum number of pods that can run on each agent.
        /// </summary>
        public readonly int MaxPods;
        /// <summary>
        /// Minimum number of nodes for auto-scaling
        /// </summary>
        public readonly int MinCount;
        /// <summary>
        /// The name of the managed Kubernetes Cluster.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableDictionary<string, string> NodeLabels;
        /// <summary>
        /// Resource ID for the Public IP Addresses Prefix for the nodes in this Agent Pool.
        /// </summary>
        public readonly string NodePublicIpPrefixId;
        public readonly ImmutableArray<string> NodeTaints;
        /// <summary>
        /// Kubernetes version used for the Agents.
        /// </summary>
        public readonly string OrchestratorVersion;
        /// <summary>
        /// The size of the Agent VM's Operating System Disk in GB.
        /// </summary>
        public readonly int OsDiskSizeGb;
        /// <summary>
        /// The Operating System used for the Agents.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The type of identity used for the managed cluster.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A `upgrade_settings` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterAgentPoolProfileUpgradeSettingResult> UpgradeSettings;
        /// <summary>
        /// The size of each VM in the Agent Pool (e.g. `Standard_F1`).
        /// </summary>
        public readonly string VmSize;
        /// <summary>
        /// The ID of the Subnet where the Agents in the Pool are provisioned.
        /// </summary>
        public readonly string VnetSubnetId;
        /// <summary>
        /// Specifies the Availability Zones where the Nodes within this Agent Pool exist.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetKubernetesClusterAgentPoolProfileResult(
            ImmutableArray<string> availabilityZones,

            int count,

            bool enableAutoScaling,

            bool enableNodePublicIp,

            int maxCount,

            int maxPods,

            int minCount,

            string name,

            ImmutableDictionary<string, string> nodeLabels,

            string nodePublicIpPrefixId,

            ImmutableArray<string> nodeTaints,

            string orchestratorVersion,

            int osDiskSizeGb,

            string osType,

            ImmutableDictionary<string, string> tags,

            string type,

            ImmutableArray<Outputs.GetKubernetesClusterAgentPoolProfileUpgradeSettingResult> upgradeSettings,

            string vmSize,

            string vnetSubnetId,

            ImmutableArray<string> zones)
        {
            AvailabilityZones = availabilityZones;
            Count = count;
            EnableAutoScaling = enableAutoScaling;
            EnableNodePublicIp = enableNodePublicIp;
            MaxCount = maxCount;
            MaxPods = maxPods;
            MinCount = minCount;
            Name = name;
            NodeLabels = nodeLabels;
            NodePublicIpPrefixId = nodePublicIpPrefixId;
            NodeTaints = nodeTaints;
            OrchestratorVersion = orchestratorVersion;
            OsDiskSizeGb = osDiskSizeGb;
            OsType = osType;
            Tags = tags;
            Type = type;
            UpgradeSettings = upgradeSettings;
            VmSize = vmSize;
            VnetSubnetId = vnetSubnetId;
            Zones = zones;
        }
    }
}
