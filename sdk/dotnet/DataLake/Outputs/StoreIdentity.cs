// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataLake.Outputs
{

    [OutputType]
    public sealed class StoreIdentity
    {
        public readonly string? PrincipalId;
        public readonly string? TenantId;
        /// <summary>
        /// The Type of Identity which should be used for this Data Lake Store Account. At this time the only possible value is `SystemAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private StoreIdentity(
            string? principalId,

            string? tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
