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
    public sealed class GetServerIdentityResult
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The Tenant ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The identity type of the Microsoft SQL Server.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of the User Assigned Identities of this SQL Server.
        /// </summary>
        public readonly ImmutableArray<string> UserAssignedIdentityIds;

        [OutputConstructor]
        private GetServerIdentityResult(
            string principalId,

            string tenantId,

            string type,

            ImmutableArray<string> userAssignedIdentityIds)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentityIds = userAssignedIdentityIds;
        }
    }
}
