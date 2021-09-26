// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Outputs
{

    [OutputType]
    public sealed class RulesEngineRuleAction
    {
        /// <summary>
        /// A `request_header` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RulesEngineRuleActionRequestHeader> RequestHeaders;
        /// <summary>
        /// A `response_header` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RulesEngineRuleActionResponseHeader> ResponseHeaders;

        [OutputConstructor]
        private RulesEngineRuleAction(
            ImmutableArray<Outputs.RulesEngineRuleActionRequestHeader> requestHeaders,

            ImmutableArray<Outputs.RulesEngineRuleActionResponseHeader> responseHeaders)
        {
            RequestHeaders = requestHeaders;
            ResponseHeaders = responseHeaders;
        }
    }
}
