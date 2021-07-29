// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DomainServices.Inputs
{

    public sealed class ServiceSecureLdapArgs : Pulumi.ResourceArgs
    {
        [Input("certificateExpiry")]
        public Input<string>? CertificateExpiry { get; set; }

        [Input("certificateThumbprint")]
        public Input<string>? CertificateThumbprint { get; set; }

        /// <summary>
        /// Whether to enable secure LDAP for the managed domain. Defaults to `false`.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Whether to enable external access to LDAPS over the Internet. Defaults to `false`.
        /// </summary>
        [Input("externalAccessEnabled")]
        public Input<bool>? ExternalAccessEnabled { get; set; }

        /// <summary>
        /// The certificate/private key to use for LDAPS, as a base64-encoded TripleDES-SHA1 encrypted PKCS#12 bundle (PFX file).
        /// </summary>
        [Input("pfxCertificate", required: true)]
        public Input<string> PfxCertificate { get; set; } = null!;

        /// <summary>
        /// The password to use for decrypting the PKCS#12 bundle (PFX file).
        /// </summary>
        [Input("pfxCertificatePassword", required: true)]
        public Input<string> PfxCertificatePassword { get; set; } = null!;

        [Input("publicCertificate")]
        public Input<string>? PublicCertificate { get; set; }

        public ServiceSecureLdapArgs()
        {
        }
    }
}
