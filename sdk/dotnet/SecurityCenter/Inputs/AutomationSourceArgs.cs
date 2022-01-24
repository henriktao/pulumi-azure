// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter.Inputs
{

    public sealed class AutomationSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of data that will trigger this automation. Must be one of `Alerts`, `Assessments`, `AssessmentsSnapshot`, `RegulatoryComplianceAssessment`, `RegulatoryComplianceAssessmentSnapshot`, `SecureScoreControls`, `SecureScoreControlsSnapshot`, `SecureScores`, `SecureScoresSnapshot`, `SubAssessments` or `SubAssessmentsSnapshot`. Note. assessments are also referred to as recommendations
        /// </summary>
        [Input("eventSource", required: true)]
        public Input<string> EventSource { get; set; } = null!;

        [Input("ruleSets")]
        private InputList<Inputs.AutomationSourceRuleSetArgs>? _ruleSets;

        /// <summary>
        /// A set of rules which evaluate upon event and data interception. This is defined in one or more `rule_set` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AutomationSourceRuleSetArgs> RuleSets
        {
            get => _ruleSets ?? (_ruleSets = new InputList<Inputs.AutomationSourceRuleSetArgs>());
            set => _ruleSets = value;
        }

        public AutomationSourceArgs()
        {
        }
    }
}
