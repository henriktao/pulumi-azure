// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.WebPubSub.Inputs
{

    public sealed class NetworkAclPublicNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("allowedRequestTypes")]
        private InputList<string>? _allowedRequestTypes;

        /// <summary>
        /// The allowed request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        /// </summary>
        public InputList<string> AllowedRequestTypes
        {
            get => _allowedRequestTypes ?? (_allowedRequestTypes = new InputList<string>());
            set => _allowedRequestTypes = value;
        }

        [Input("deniedRequestTypes")]
        private InputList<string>? _deniedRequestTypes;

        /// <summary>
        /// The denied request types for the public network. Possible values are `ClientConnection`, `ServerConnection`, `RESTAPI` and `Trace`.
        /// </summary>
        public InputList<string> DeniedRequestTypes
        {
            get => _deniedRequestTypes ?? (_deniedRequestTypes = new InputList<string>());
            set => _deniedRequestTypes = value;
        }

        public NetworkAclPublicNetworkArgs()
        {
        }
    }
}
