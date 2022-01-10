// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cdn.Inputs
{

    public sealed class EndpointCustomDomainUserManagedHttpsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Key Vault Certificate that contains the HTTPS certificate.
        /// </summary>
        [Input("keyVaultCertificateId", required: true)]
        public Input<string> KeyVaultCertificateId { get; set; } = null!;

        /// <summary>
        /// The TLS protocol version that is used for HTTPS. Possible values are `TLS10` (representing TLS 1.0/1.1) and `TLS12` (representing TLS 1.2). Defaults to `TLS12`.
        /// </summary>
        [Input("tlsVersion")]
        public Input<string>? TlsVersion { get; set; }

        public EndpointCustomDomainUserManagedHttpsGetArgs()
        {
        }
    }
}
