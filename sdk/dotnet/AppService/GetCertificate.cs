// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source to access information about an App Service Certificate.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.AppService.GetCertificate.InvokeAsync(new Azure.AppService.GetCertificateArgs
        ///         {
        ///             Name = "example-app-service-certificate",
        ///             ResourceGroupName = "example-rg",
        ///         }));
        ///         this.AppServiceCertificateId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("appServiceCertificateId")]
        ///     public Output&lt;string&gt; AppServiceCertificateId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure:appservice/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an App Service Certificate.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.AppService.GetCertificate.InvokeAsync(new Azure.AppService.GetCertificateArgs
        ///         {
        ///             Name = "example-app-service-certificate",
        ///             ResourceGroupName = "example-rg",
        ///         }));
        ///         this.AppServiceCertificateId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("appServiceCertificateId")]
        ///     public Output&lt;string&gt; AppServiceCertificateId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("azure:appservice/getCertificate:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the certificate.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetCertificateArgs()
        {
        }
    }

    public sealed class GetCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the certificate.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The expiration date for the certificate.
        /// </summary>
        public readonly string ExpirationDate;
        /// <summary>
        /// The friendly name of the certificate.
        /// </summary>
        public readonly string FriendlyName;
        /// <summary>
        /// List of host names the certificate applies to.
        /// </summary>
        public readonly ImmutableArray<string> HostNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The issue date for the certificate.
        /// </summary>
        public readonly string IssueDate;
        /// <summary>
        /// The name of the certificate issuer.
        /// </summary>
        public readonly string Issuer;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The subject name of the certificate.
        /// </summary>
        public readonly string SubjectName;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The thumbprint for the certificate.
        /// </summary>
        public readonly string Thumbprint;

        [OutputConstructor]
        private GetCertificateResult(
            string expirationDate,

            string friendlyName,

            ImmutableArray<string> hostNames,

            string id,

            string issueDate,

            string issuer,

            string location,

            string name,

            string resourceGroupName,

            string subjectName,

            ImmutableDictionary<string, string>? tags,

            string thumbprint)
        {
            ExpirationDate = expirationDate;
            FriendlyName = friendlyName;
            HostNames = hostNames;
            Id = id;
            IssueDate = issueDate;
            Issuer = issuer;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SubjectName = subjectName;
            Tags = tags;
            Thumbprint = thumbprint;
        }
    }
}
