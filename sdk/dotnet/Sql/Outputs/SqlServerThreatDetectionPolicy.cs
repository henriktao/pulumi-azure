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
    public sealed class SqlServerThreatDetectionPolicy
    {
        /// <summary>
        /// Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Data_Exfiltration`, `Sql_Injection`, `Sql_Injection_Vulnerability` and `Unsafe_Action"`,.
        /// </summary>
        public readonly ImmutableArray<string> DisabledAlerts;
        /// <summary>
        /// Should the account administrators be emailed when this alert is triggered?
        /// </summary>
        public readonly bool? EmailAccountAdmins;
        /// <summary>
        /// A list of email addresses which alerts should be sent to.
        /// </summary>
        public readonly ImmutableArray<string> EmailAddresses;
        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs.
        /// </summary>
        public readonly int? RetentionDays;
        /// <summary>
        /// The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        /// </summary>
        public readonly string? StorageAccountAccessKey;
        /// <summary>
        /// Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        /// </summary>
        public readonly string? StorageEndpoint;

        [OutputConstructor]
        private SqlServerThreatDetectionPolicy(
            ImmutableArray<string> disabledAlerts,

            bool? emailAccountAdmins,

            ImmutableArray<string> emailAddresses,

            int? retentionDays,

            string? state,

            string? storageAccountAccessKey,

            string? storageEndpoint)
        {
            DisabledAlerts = disabledAlerts;
            EmailAccountAdmins = emailAccountAdmins;
            EmailAddresses = emailAddresses;
            RetentionDays = retentionDays;
            State = state;
            StorageAccountAccessKey = storageAccountAccessKey;
            StorageEndpoint = storageEndpoint;
        }
    }
}
