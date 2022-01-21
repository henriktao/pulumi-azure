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
    public sealed class GetFlexibleServerHighAvailabilityResult
    {
        /// <summary>
        /// The high availability mode of the MySQL Flexible Server.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The availability zone of the standby Flexible Server.
        /// </summary>
        public readonly string StandbyAvailabilityZone;

        [OutputConstructor]
        private GetFlexibleServerHighAvailabilityResult(
            string mode,

            string standbyAvailabilityZone)
        {
            Mode = mode;
            StandbyAvailabilityZone = standbyAvailabilityZone;
        }
    }
}
