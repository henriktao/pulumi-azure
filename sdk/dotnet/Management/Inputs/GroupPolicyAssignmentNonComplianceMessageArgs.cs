// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Management.Inputs
{

    public sealed class GroupPolicyAssignmentNonComplianceMessageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        /// </summary>
        [Input("policyDefinitionReferenceId")]
        public Input<string>? PolicyDefinitionReferenceId { get; set; }

        public GroupPolicyAssignmentNonComplianceMessageArgs()
        {
        }
    }
}
