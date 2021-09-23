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
    public sealed class ComputeClusterSsh
    {
        /// <summary>
        /// Password of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public readonly string? AdminPassword;
        /// <summary>
        /// Name of the administrator user account which can be used to SSH to nodes. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public readonly string AdminUsername;
        /// <summary>
        /// SSH public key of the administrator user account. Changing this forces a new Machine Learning Compute Cluster to be created.
        /// </summary>
        public readonly string? KeyValue;

        [OutputConstructor]
        private ComputeClusterSsh(
            string? adminPassword,

            string adminUsername,

            string? keyValue)
        {
            AdminPassword = adminPassword;
            AdminUsername = adminUsername;
            KeyValue = keyValue;
        }
    }
}
