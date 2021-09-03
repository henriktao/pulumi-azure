// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppConfiguration.Outputs
{

    [OutputType]
    public sealed class ConfigurationStoreIdentity
    {
        /// <summary>
        /// A list of IDs for User Assigned Managed Identity resources to be assigned.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        /// <summary>
        /// The ID of the Principal (Client) in Azure Active Directory.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The ID of the Azure Active Directory Tenant.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// Specifies the type of Managed Service Identity that should be configured on this API Management Service. Possible values are `SystemAssigned`, `UserAssigned`, `SystemAssigned, UserAssigned` (to enable both).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ConfigurationStoreIdentity(
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
