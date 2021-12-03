// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class FunctionAppAuthSettingsTwitter
    {
        /// <summary>
        /// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
        /// </summary>
        public readonly string ConsumerKey;
        /// <summary>
        /// The OAuth 1.0a consumer secret of the Twitter application used for sign-in.
        /// </summary>
        public readonly string ConsumerSecret;

        [OutputConstructor]
        private FunctionAppAuthSettingsTwitter(
            string consumerKey,

            string consumerSecret)
        {
            ConsumerKey = consumerKey;
            ConsumerSecret = consumerSecret;
        }
    }
}
