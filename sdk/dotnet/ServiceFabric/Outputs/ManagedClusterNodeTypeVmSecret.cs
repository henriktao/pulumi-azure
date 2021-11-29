// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Outputs
{

    [OutputType]
    public sealed class ManagedClusterNodeTypeVmSecret
    {
        /// <summary>
        /// One or more `certificates` blocks as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedClusterNodeTypeVmSecretCertificate> Certificates;
        /// <summary>
        /// The ID of the Vault that contain the certificates.
        /// </summary>
        public readonly string VaultId;

        [OutputConstructor]
        private ManagedClusterNodeTypeVmSecret(
            ImmutableArray<Outputs.ManagedClusterNodeTypeVmSecretCertificate> certificates,

            string vaultId)
        {
            Certificates = certificates;
            VaultId = vaultId;
        }
    }
}
