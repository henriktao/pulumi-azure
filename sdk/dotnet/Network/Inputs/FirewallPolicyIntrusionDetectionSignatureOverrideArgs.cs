// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class FirewallPolicyIntrusionDetectionSignatureOverrideArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// 12-digit number (id) which identifies your signature.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// state can be any of "Off", "Alert" or "Deny".
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public FirewallPolicyIntrusionDetectionSignatureOverrideArgs()
        {
        }
    }
}
