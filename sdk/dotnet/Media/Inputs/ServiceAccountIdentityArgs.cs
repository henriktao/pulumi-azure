// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class ServiceAccountIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Principal ID associated with this Managed Service Identity.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The Tenant ID associated with this Managed Service Identity.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Specifies the type of Managed Service Identity that should be configured on this Media Services Account. Possible value is  `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServiceAccountIdentityArgs()
        {
        }
    }
}
