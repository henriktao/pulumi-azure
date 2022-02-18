// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Outputs
{

    [OutputType]
    public sealed class ManagedInstanceFailoverGroupPartnerRegion
    {
        /// <summary>
        /// The Azure Region where the Managed Instance Failover Group should exist. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The partner replication role of the Managed Instance Failover Group.
        /// </summary>
        public readonly string? Role;

        [OutputConstructor]
        private ManagedInstanceFailoverGroupPartnerRegion(
            string? location,

            string? role)
        {
            Location = location;
            Role = role;
        }
    }
}
