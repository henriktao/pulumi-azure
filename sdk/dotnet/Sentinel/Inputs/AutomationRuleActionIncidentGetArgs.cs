// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Sentinel.Inputs
{

    public sealed class AutomationRuleActionIncidentGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The classification of the incident, when closing it. Possible values are: `BenignPositive_SuspiciousButExpected`, `FalsePositive_InaccurateData`, `FalsePositive_IncorrectAlertLogic`, `TruePositive_SuspiciousActivity` and `Undetermined`.
        /// </summary>
        [Input("classification")]
        public Input<string>? Classification { get; set; }

        /// <summary>
        /// The comment why the incident is to be closed.
        /// </summary>
        [Input("classificationComment")]
        public Input<string>? ClassificationComment { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// Specifies a list of labels to add to the incident.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The execution order of this action.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// The object ID of the entity this incident is assigned to.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The severity to add to the incident.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// The status to set to the incident. Possible values are: `Active`, `Closed`, `New`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AutomationRuleActionIncidentGetArgs()
        {
        }
    }
}
