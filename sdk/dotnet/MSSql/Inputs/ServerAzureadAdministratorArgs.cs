// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql.Inputs
{

    public sealed class ServerAzureadAdministratorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Specifies whether only AD Users and administrators (like `azuread_administrator.0.login_username`) can be used to login or also local database users (like `administrator_login`).
        /// </summary>
        [Input("azureadAuthenticationOnly")]
        public Input<bool>? AzureadAuthenticationOnly { get; set; }

        /// <summary>
        /// (Required)  The login username of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("loginUsername", required: true)]
        public Input<string> LoginUsername { get; set; } = null!;

        /// <summary>
        /// (Required) The object id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// (Optional) The tenant id of the Azure AD Administrator of this SQL Server.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ServerAzureadAdministratorArgs()
        {
        }
    }
}
