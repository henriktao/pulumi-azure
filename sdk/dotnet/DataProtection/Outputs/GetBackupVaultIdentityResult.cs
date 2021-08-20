// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataProtection.Outputs
{

    [OutputType]
    public sealed class GetBackupVaultIdentityResult
    {
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Specifies the identity type of the Backup Vault.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBackupVaultIdentityResult(
            string principalId,

            string tenantId,

            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
