// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MySql.Outputs
{

    [OutputType]
    public sealed class FlexibleServerHighAvailability
    {
        /// <summary>
        /// The high availability mode for the MySQL Flexible Server. Possibles values are `SameZone` and `ZoneRedundant`.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The availability zone of the standby Flexible Server. Possible values are `1`, `2` and `3`.
        /// </summary>
        public readonly string? StandbyAvailabilityZone;

        [OutputConstructor]
        private FlexibleServerHighAvailability(
            string mode,

            string? standbyAvailabilityZone)
        {
            Mode = mode;
            StandbyAvailabilityZone = standbyAvailabilityZone;
        }
    }
}
