// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Purview.Inputs
{

    public sealed class AccountIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Principal (Client) in Azure Active Directory.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The ID of the Azure Active Directory Tenant.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of Managed Identity assigned to this Purview Account.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccountIdentityArgs()
        {
        }
    }
}
