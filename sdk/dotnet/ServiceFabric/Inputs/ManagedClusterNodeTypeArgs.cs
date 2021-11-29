// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Inputs
{

    public sealed class ManagedClusterNodeTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets the port range available for applications. Format is `&lt;from_port&gt;-&lt;to_port&gt;`, for example `10000-20000`.
        /// </summary>
        [Input("applicationPortRange", required: true)]
        public Input<string> ApplicationPortRange { get; set; } = null!;

        [Input("capacities")]
        private InputMap<string>? _capacities;

        /// <summary>
        /// Specifies a list of key/value pairs used to set capacity tags for this node type.
        /// </summary>
        public InputMap<string> Capacities
        {
            get => _capacities ?? (_capacities = new InputMap<string>());
            set => _capacities = value;
        }

        /// <summary>
        /// The size of the data disk in gigabytes..
        /// </summary>
        [Input("dataDiskSizeGb", required: true)]
        public Input<int> DataDiskSizeGb { get; set; } = null!;

        /// <summary>
        /// The type of the disk to use for storing data. It can be one of `Premium_LRS`, `Standard_LRS`, or `StandardSSD_LRS`.
        /// </summary>
        [Input("dataDiskType")]
        public Input<string>? DataDiskType { get; set; }

        /// <summary>
        /// Sets the port range available for the OS. Format is `&lt;from_port&gt;-&lt;to_port&gt;`, for example `10000-20000`. There has to be at least 255 ports available and cannot overlap with `application_port_range`..
        /// </summary>
        [Input("ephemeralPortRange", required: true)]
        public Input<string> EphemeralPortRange { get; set; } = null!;

        /// <summary>
        /// The ID of the Resource Group.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// If set the node type can be composed of multiple placement groups.
        /// </summary>
        [Input("multiplePlacementGroupsEnabled")]
        public Input<bool>? MultiplePlacementGroupsEnabled { get; set; }

        /// <summary>
        /// The name which should be used for this node type.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("placementProperties")]
        private InputMap<string>? _placementProperties;

        /// <summary>
        /// Specifies a list of placement tags that can be used to indicate where services should run..
        /// </summary>
        public InputMap<string> PlacementProperties
        {
            get => _placementProperties ?? (_placementProperties = new InputMap<string>());
            set => _placementProperties = value;
        }

        /// <summary>
        /// If set to true, system services will run on this node type. Only one node type should be marked as primary. Primary node type cannot be deleted or changed once they're created.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// If set to true, only stateless workloads can run on this node type.
        /// </summary>
        [Input("stateless")]
        public Input<bool>? Stateless { get; set; }

        /// <summary>
        /// The offer type of the marketplace image cluster VMs will use.
        /// </summary>
        [Input("vmImageOffer", required: true)]
        public Input<string> VmImageOffer { get; set; } = null!;

        /// <summary>
        /// The publisher of the marketplace image cluster VMs will use.
        /// </summary>
        [Input("vmImagePublisher", required: true)]
        public Input<string> VmImagePublisher { get; set; } = null!;

        /// <summary>
        /// The SKU of the marketplace image cluster VMs will use.
        /// </summary>
        [Input("vmImageSku", required: true)]
        public Input<string> VmImageSku { get; set; } = null!;

        /// <summary>
        /// The version of the marketplace image cluster VMs will use.
        /// </summary>
        [Input("vmImageVersion", required: true)]
        public Input<string> VmImageVersion { get; set; } = null!;

        /// <summary>
        /// The number of instances this node type will launch.
        /// </summary>
        [Input("vmInstanceCount", required: true)]
        public Input<int> VmInstanceCount { get; set; } = null!;

        [Input("vmSecrets")]
        private InputList<Inputs.ManagedClusterNodeTypeVmSecretArgs>? _vmSecrets;

        /// <summary>
        /// One or more `vm_secrets` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ManagedClusterNodeTypeVmSecretArgs> VmSecrets
        {
            get => _vmSecrets ?? (_vmSecrets = new InputList<Inputs.ManagedClusterNodeTypeVmSecretArgs>());
            set => _vmSecrets = value;
        }

        /// <summary>
        /// The size of the instances in this node type.
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public ManagedClusterNodeTypeArgs()
        {
        }
    }
}
