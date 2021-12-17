// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps.Outputs
{

    [OutputType]
    public sealed class WorkflowAccessControlTriggerOpenAuthenticationPolicyClaim
    {
        /// <summary>
        /// The name of the OAuth policy claim for the Logic App Workflow.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the OAuth policy claim for the Logic App Workflow.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private WorkflowAccessControlTriggerOpenAuthenticationPolicyClaim(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
