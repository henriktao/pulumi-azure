// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation.Outputs
{

    [OutputType]
    public sealed class AccountIdentity
    {
        /// <summary>
        /// The ID of the User Assigned Identity which should be assigned to this Automation Account.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        public readonly string? PrincipalId;
        public readonly string? TenantId;
        /// <summary>
        /// The type of identity used for the automation account. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccountIdentity(
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
