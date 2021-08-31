// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class ServiceCertificate
    {
        /// <summary>
        /// The password for the certificate.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// The Base64 Encoded PFX or Base64 Encoded X.509 Certificate.
        /// </summary>
        public readonly string EncodedCertificate;
        /// <summary>
        /// The expiration date of the certificate in RFC3339 format: `2000-01-02T03:04:05Z`.
        /// </summary>
        public readonly string? Expiry;
        /// <summary>
        /// The name of the Certificate Store where this certificate should be stored. Possible values are `CertificateAuthority` and `Root`.
        /// </summary>
        public readonly string StoreName;
        /// <summary>
        /// The subject of the certificate.
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// The thumbprint of the certificate.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private ServiceCertificate(
            string? certificatePassword,

            string encodedCertificate,

            string? expiry,

            string storeName,

            string? subject,

            string? thumbprint)
        {
            CertificatePassword = certificatePassword;
            EncodedCertificate = encodedCertificate;
            Expiry = expiry;
            StoreName = storeName;
            Subject = subject;
            Thumbprint = thumbprint;
        }
    }
}
