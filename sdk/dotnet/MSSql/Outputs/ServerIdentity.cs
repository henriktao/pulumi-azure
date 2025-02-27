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
    public sealed class ServerIdentity
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Identity of this SQL Server.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The tenant id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// Specifies the identity type of the Microsoft SQL Server. Possible values are `SystemAssigned` (where Azure will generate a Service Principal for you) and `UserAssigned` where you can specify the Service Principal IDs in the `user_assigned_identity_ids` field.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Specifies a list of User Assigned Identity IDs to be assigned. Required if `type` is `UserAssigned` and should be combined with `primary_user_assigned_identity_id`.
        /// </summary>
        public readonly ImmutableArray<string> UserAssignedIdentityIds;

        [OutputConstructor]
        private ServerIdentity(
            string? principalId,

            string? tenantId,

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
