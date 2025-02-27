// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.StreamAnalytics.Outputs
{

    [OutputType]
    public sealed class JobIdentity
    {
        /// <summary>
        /// The ID of the Principal (Client) in Azure Active Directory.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The ID of the Azure Active Directory Tenant.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The type of identity used for the Stream Analytics Job. The only possible value is `SystemAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private JobIdentity(
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
