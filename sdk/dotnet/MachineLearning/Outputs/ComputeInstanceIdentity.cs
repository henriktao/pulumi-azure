// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MachineLearning.Outputs
{

    [OutputType]
    public sealed class ComputeInstanceIdentity
    {
        /// <summary>
        /// A list of User Managed Identity ID's which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Instance.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// User’s AAD Tenant Id.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The Type of Identity which should be used for this Machine Learning Synapse Spark. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned, UserAssigned` (to enable both).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ComputeInstanceIdentity(
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
