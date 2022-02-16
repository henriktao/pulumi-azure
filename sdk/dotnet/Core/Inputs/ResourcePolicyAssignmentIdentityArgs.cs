// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core.Inputs
{

    public sealed class ResourcePolicyAssignmentIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Principal ID of the Policy Assignment for this Resource.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The Tenant ID of the Policy Assignment for this Resource.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The Type of Managed Identity which should be added to this Policy Definition. The only possible value is `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourcePolicyAssignmentIdentityArgs()
        {
        }
    }
}
