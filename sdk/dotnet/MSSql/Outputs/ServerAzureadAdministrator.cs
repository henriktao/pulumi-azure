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
    public sealed class ServerAzureadAdministrator
    {
        /// <summary>
        /// Specifies whether only AD Users and administrators (like `azuread_administrator.0.login_username`) can be used to login or also local database users (like `administrator_login`).
        /// </summary>
        public readonly bool? AzureadAuthenticationOnly;
        /// <summary>
        /// The login username of the Azure AD Administrator of this SQL Server.
        /// </summary>
        public readonly string LoginUsername;
        /// <summary>
        /// The object id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The tenant id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        public readonly string? TenantId;

        [OutputConstructor]
        private ServerAzureadAdministrator(
            bool? azureadAuthenticationOnly,

            string loginUsername,

            string objectId,

            string? tenantId)
        {
            AzureadAuthenticationOnly = azureadAuthenticationOnly;
            LoginUsername = loginUsername;
            ObjectId = objectId;
            TenantId = tenantId;
        }
    }
}
