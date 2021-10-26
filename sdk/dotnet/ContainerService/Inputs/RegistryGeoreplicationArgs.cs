// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class RegistryGeoreplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A location where the container registry should be geo-replicated.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Whether regional endpoint is enabled for this Container Registry? Defaults to `false`.
        /// </summary>
        [Input("regionalEndpointEnabled")]
        public Input<bool>? RegionalEndpointEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to this replication location.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether zone redundancy is enabled for this replication location? Defaults to `false`.
        /// </summary>
        [Input("zoneRedundancyEnabled")]
        public Input<bool>? ZoneRedundancyEnabled { get; set; }

        public RegistryGeoreplicationArgs()
        {
        }
    }
}
