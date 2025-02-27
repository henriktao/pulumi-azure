// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyIntrusionDetectionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// In which mode you want to run intrusion detection: "Off", "Alert" or "Deny".
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("signatureOverrides")]
        private InputList<Inputs.FirewallPolicyIntrusionDetectionSignatureOverrideGetArgs>? _signatureOverrides;

        /// <summary>
        /// One or more `signature_overrides` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallPolicyIntrusionDetectionSignatureOverrideGetArgs> SignatureOverrides
        {
            get => _signatureOverrides ?? (_signatureOverrides = new InputList<Inputs.FirewallPolicyIntrusionDetectionSignatureOverrideGetArgs>());
            set => _signatureOverrides = value;
        }

        [Input("trafficBypasses")]
        private InputList<Inputs.FirewallPolicyIntrusionDetectionTrafficBypassGetArgs>? _trafficBypasses;

        /// <summary>
        /// One or more `traffic_bypass` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FirewallPolicyIntrusionDetectionTrafficBypassGetArgs> TrafficBypasses
        {
            get => _trafficBypasses ?? (_trafficBypasses = new InputList<Inputs.FirewallPolicyIntrusionDetectionTrafficBypassGetArgs>());
            set => _trafficBypasses = value;
        }

        public FirewallPolicyIntrusionDetectionGetArgs()
        {
        }
    }
}
