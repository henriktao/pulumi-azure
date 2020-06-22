// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateLink.Inputs
{

    public sealed class EndpointPrivateDnsZoneGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Private DNS Zone Config.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies the Name of the Private Service Connection. Changing this forces the a new `private_dns_zone_group` to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("privateDnsZoneIds", required: true)]
        private InputList<string>? _privateDnsZoneIds;

        /// <summary>
        /// Specifies the list of Private DNS Zones to include within the `private_dns_zone_group`.
        /// </summary>
        public InputList<string> PrivateDnsZoneIds
        {
            get => _privateDnsZoneIds ?? (_privateDnsZoneIds = new InputList<string>());
            set => _privateDnsZoneIds = value;
        }

        public EndpointPrivateDnsZoneGroupArgs()
        {
        }
    }
}
