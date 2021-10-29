// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sql.Inputs
{

    public sealed class SqlServerThreatDetectionPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("disabledAlerts")]
        private InputList<string>? _disabledAlerts;

        /// <summary>
        /// Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Data_Exfiltration`, `Sql_Injection`, `Sql_Injection_Vulnerability` and `Unsafe_Action"`,.
        /// </summary>
        public InputList<string> DisabledAlerts
        {
            get => _disabledAlerts ?? (_disabledAlerts = new InputList<string>());
            set => _disabledAlerts = value;
        }

        /// <summary>
        /// Should the account administrators be emailed when this alert is triggered?
        /// </summary>
        [Input("emailAccountAdmins")]
        public Input<bool>? EmailAccountAdmins { get; set; }

        [Input("emailAddresses")]
        private InputList<string>? _emailAddresses;

        /// <summary>
        /// A list of email addresses which alerts should be sent to.
        /// </summary>
        public InputList<string> EmailAddresses
        {
            get => _emailAddresses ?? (_emailAddresses = new InputList<string>());
            set => _emailAddresses = value;
        }

        /// <summary>
        /// Specifies the number of days to keep in the Threat Detection audit logs.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The State of the Policy. Possible values are `Enabled` or `Disabled`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Specifies the blob storage endpoint (e.g. `https://MyAccount.blob.core.windows.net`). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public SqlServerThreatDetectionPolicyArgs()
        {
        }
    }
}
