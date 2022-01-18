// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.RecoveryServices.Outputs
{

    [OutputType]
    public sealed class VaultEncryption
    {
        /// <summary>
        /// Enabling/Disabling the Double Encryption state.
        /// </summary>
        public readonly bool InfrastructureEncryptionEnabled;
        /// <summary>
        /// The Key Vault key id used to encrypt this vault. Key managed by Vault Managed Hardware Security Module is also supported.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Indicate that system assigned identity should be used or not. At this time the only possible value is `true`. Defaults to `true`.
        /// </summary>
        public readonly bool? UseSystemAssignedIdentity;

        [OutputConstructor]
        private VaultEncryption(
            bool infrastructureEncryptionEnabled,

            string keyId,

            bool? useSystemAssignedIdentity)
        {
            InfrastructureEncryptionEnabled = infrastructureEncryptionEnabled;
            KeyId = keyId;
            UseSystemAssignedIdentity = useSystemAssignedIdentity;
        }
    }
}
