// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql.Inputs
{

    public sealed class FlexibleServerHighAvailabilityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The high availability mode for the MySQL Flexible Server. Possibles values are `SameZone` and `ZoneRedundant`.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The availability zone of the standby Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        [Input("standbyAvailabilityZone")]
        public Input<string>? StandbyAvailabilityZone { get; set; }

        public FlexibleServerHighAvailabilityArgs()
        {
        }
    }
}
