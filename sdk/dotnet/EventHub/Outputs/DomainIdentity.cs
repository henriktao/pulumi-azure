// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventHub.Outputs
{

    [OutputType]
    public sealed class DomainIdentity
    {
        /// <summary>
        /// Specifies a list of user managed identity ids to be assigned. Required if `type` is `UserAssigned`.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        /// <summary>
        /// Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this Event Grid Domain.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this Event Grid Domain.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// Specifies the identity type of Event Grid Domain. Possible values are `SystemAssigned` (where Azure will generate a Principal for you) or `UserAssigned` where you can specify the User Assigned Managed Identity IDs in the `identity_ids` field.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DomainIdentity(
            ImmutableArray<string> identityIds,

            string? principalId,

            string? tenantId,

            string type)
        {
            IdentityIds = identityIds;
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
