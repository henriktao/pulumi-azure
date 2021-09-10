// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayTrustedClientCertificate
    {
        /// <summary>
        /// The base-64 encoded certificate.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the Trusted Client Certificate that is unique within this Application Gateway.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ApplicationGatewayTrustedClientCertificate(
            string data,

            string? id,

            string name)
        {
            Data = data;
            Id = id;
            Name = name;
        }
    }
}
