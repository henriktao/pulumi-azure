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
    public sealed class ServiceHostnameConfigurationScm
    {
        /// <summary>
        /// One or more (up to 10) `certificate` blocks as defined below.
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// The password for the certificate.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// The expiration date of the certificate in RFC3339 format: `2000-01-02T03:04:05Z`.
        /// </summary>
        public readonly string? Expiry;
        /// <summary>
        /// The Hostname to use for the Management API.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type `application/x-pkcs12`.
        /// </summary>
        public readonly string? KeyVaultId;
        /// <summary>
        /// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to `false`.
        /// </summary>
        public readonly bool? NegotiateClientCertificate;
        /// <summary>
        /// The client id of the System or User Assigned Managed identity generated by Azure AD, which has `GET` access to the keyVault containing the SSL certificate.
        /// </summary>
        public readonly string? SslKeyvaultIdentityClientId;
        /// <summary>
        /// The subject of the certificate.
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// The thumbprint of the certificate.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private ServiceHostnameConfigurationScm(
            string? certificate,

            string? certificatePassword,

            string? expiry,

            string hostName,

            string? keyVaultId,

            bool? negotiateClientCertificate,

            string? sslKeyvaultIdentityClientId,

            string? subject,

            string? thumbprint)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            Expiry = expiry;
            HostName = hostName;
            KeyVaultId = keyVaultId;
            NegotiateClientCertificate = negotiateClientCertificate;
            SslKeyvaultIdentityClientId = sslKeyvaultIdentityClientId;
            Subject = subject;
            Thumbprint = thumbprint;
        }
    }
}
