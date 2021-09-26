// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Inputs
{

    public sealed class RulesEngineRuleActionArgs : Pulumi.ResourceArgs
    {
        [Input("requestHeaders")]
        private InputList<Inputs.RulesEngineRuleActionRequestHeaderArgs>? _requestHeaders;

        /// <summary>
        /// A `request_header` block as defined below.
        /// </summary>
        public InputList<Inputs.RulesEngineRuleActionRequestHeaderArgs> RequestHeaders
        {
            get => _requestHeaders ?? (_requestHeaders = new InputList<Inputs.RulesEngineRuleActionRequestHeaderArgs>());
            set => _requestHeaders = value;
        }

        [Input("responseHeaders")]
        private InputList<Inputs.RulesEngineRuleActionResponseHeaderArgs>? _responseHeaders;

        /// <summary>
        /// A `response_header` block as defined below.
        /// </summary>
        public InputList<Inputs.RulesEngineRuleActionResponseHeaderArgs> ResponseHeaders
        {
            get => _responseHeaders ?? (_responseHeaders = new InputList<Inputs.RulesEngineRuleActionResponseHeaderArgs>());
            set => _responseHeaders = value;
        }

        public RulesEngineRuleActionArgs()
        {
        }
    }
}
