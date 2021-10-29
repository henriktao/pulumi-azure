// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel.Outputs
{

    [OutputType]
    public sealed class AuthomationRuleActionIncident
    {
        /// <summary>
        /// The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        /// </summary>
        public readonly string? Classification;
        /// <summary>
        /// The comment why the incident is to be closed.
        /// </summary>
        public readonly string? ClassificationComment;
        /// <summary>
        /// Specifies a list of labels to add to the incident.
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The execution order of this action.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// The object ID of the entity this incident is assigned to.
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// The severity to add to the incident.
        /// </summary>
        public readonly string? Severity;
        /// <summary>
        /// The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private AuthomationRuleActionIncident(
            string? classification,

            string? classificationComment,

            ImmutableArray<string> labels,

            int order,

            string? ownerId,

            string? severity,

            string? status)
        {
            Classification = classification;
            ClassificationComment = classificationComment;
            Labels = labels;
            Order = order;
            OwnerId = ownerId;
            Severity = severity;
            Status = status;
        }
    }
}
