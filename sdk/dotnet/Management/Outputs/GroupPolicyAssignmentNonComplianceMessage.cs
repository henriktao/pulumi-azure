// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Management.Outputs
{

    [OutputType]
    public sealed class GroupPolicyAssignmentNonComplianceMessage
    {
        /// <summary>
        /// The non-compliance message text. When assigning policy sets (initiatives), unless `policy_definition_reference_id` is specified then this message will be the default for all policies.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
        /// </summary>
        public readonly string? PolicyDefinitionReferenceId;

        [OutputConstructor]
        private GroupPolicyAssignmentNonComplianceMessage(
            string content,

            string? policyDefinitionReferenceId)
        {
            Content = content;
            PolicyDefinitionReferenceId = policyDefinitionReferenceId;
        }
    }
}
