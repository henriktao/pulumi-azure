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
    public sealed class ComputeClusterIdentity
    {
        /// <summary>
        /// A list of User Managed Identity ID's which should be assigned to the Machine Learning Compute Cluster. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public readonly ImmutableArray<string> IdentityIds;
        /// <summary>
        /// The Principal ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Cluster.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Cluster.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The Type of Identity which should be used for this Machine Learning Compute Cluster. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ComputeClusterIdentity(
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
