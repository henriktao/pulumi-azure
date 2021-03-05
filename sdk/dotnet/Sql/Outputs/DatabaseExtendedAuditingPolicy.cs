// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql.Outputs
{

    [OutputType]
    public sealed class DatabaseExtendedAuditingPolicy
    {
        public readonly bool? LogMonitoringEnabled;
        /// <summary>
        /// Specifies the number of days to retain logs for in the storage account.
        /// </summary>
        public readonly int? RetentionInDays;
        /// <summary>
        /// Specifies the access key to use for the auditing storage account.
        /// </summary>
        public readonly string? StorageAccountAccessKey;
        /// <summary>
        /// Specifies whether `storage_account_access_key` value is the storage's secondary key.
        /// </summary>
        public readonly bool? StorageAccountAccessKeyIsSecondary;
        /// <summary>
        /// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
        /// </summary>
        public readonly string? StorageEndpoint;

        [OutputConstructor]
        private DatabaseExtendedAuditingPolicy(
            bool? logMonitoringEnabled,

            int? retentionInDays,

            string? storageAccountAccessKey,

            bool? storageAccountAccessKeyIsSecondary,

            string? storageEndpoint)
        {
            LogMonitoringEnabled = logMonitoringEnabled;
            RetentionInDays = retentionInDays;
            StorageAccountAccessKey = storageAccountAccessKey;
            StorageAccountAccessKeyIsSecondary = storageAccountAccessKeyIsSecondary;
            StorageEndpoint = storageEndpoint;
        }
    }
}
