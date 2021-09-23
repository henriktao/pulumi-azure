// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.LogicApps.Inputs
{

    public sealed class WorkflowAccessControlTriggerArgs : Pulumi.ResourceArgs
    {
        [Input("allowedCallerIpAddressRanges", required: true)]
        private InputList<string>? _allowedCallerIpAddressRanges;

        /// <summary>
        /// A list of the allowed caller IP address ranges.
        /// </summary>
        public InputList<string> AllowedCallerIpAddressRanges
        {
            get => _allowedCallerIpAddressRanges ?? (_allowedCallerIpAddressRanges = new InputList<string>());
            set => _allowedCallerIpAddressRanges = value;
        }

        public WorkflowAccessControlTriggerArgs()
        {
        }
    }
}
