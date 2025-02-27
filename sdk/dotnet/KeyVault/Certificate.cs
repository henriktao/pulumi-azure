// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault
{
    /// <summary>
    /// Manages a Key Vault Certificate.
    /// 
    /// ## Example Usage
    /// ### Importing a PFX
    /// 
    /// &gt; **Note:** this example assumed the PFX file is located in the same directory at `certificate-to-import.pfx`.
    /// 
    /// ```csharp
    /// using System;
    /// using System.IO;
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)))
    /// 	}
    /// 
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "premium",
    ///             AccessPolicies = 
    ///             {
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///                     CertificatePermissions = 
    ///                     {
    ///                         "create",
    ///                         "delete",
    ///                         "deleteissuers",
    ///                         "get",
    ///                         "getissuers",
    ///                         "import",
    ///                         "list",
    ///                         "listissuers",
    ///                         "managecontacts",
    ///                         "manageissuers",
    ///                         "setissuers",
    ///                         "update",
    ///                     },
    ///                     KeyPermissions = 
    ///                     {
    ///                         "backup",
    ///                         "create",
    ///                         "decrypt",
    ///                         "delete",
    ///                         "encrypt",
    ///                         "get",
    ///                         "import",
    ///                         "list",
    ///                         "purge",
    ///                         "recover",
    ///                         "restore",
    ///                         "sign",
    ///                         "unwrapKey",
    ///                         "update",
    ///                         "verify",
    ///                         "wrapKey",
    ///                     },
    ///                     SecretPermissions = 
    ///                     {
    ///                         "backup",
    ///                         "delete",
    ///                         "get",
    ///                         "list",
    ///                         "purge",
    ///                         "recover",
    ///                         "restore",
    ///                         "set",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleCertificate = new Azure.KeyVault.Certificate("exampleCertificate", new Azure.KeyVault.CertificateArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             Certificate = new Azure.KeyVault.Inputs.CertificateCertificateArgs
    ///             {
    ///                 Contents = ReadFileBase64("certificate-to-import.pfx"),
    ///                 Password = "",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Generating a new certificate
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Azure.Core.GetClientConfig.InvokeAsync());
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleKeyVault = new Azure.KeyVault.KeyVault("exampleKeyVault", new Azure.KeyVault.KeyVaultArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             TenantId = current.Apply(current =&gt; current.TenantId),
    ///             SkuName = "standard",
    ///             SoftDeleteRetentionDays = 7,
    ///             AccessPolicies = 
    ///             {
    ///                 new Azure.KeyVault.Inputs.KeyVaultAccessPolicyArgs
    ///                 {
    ///                     TenantId = current.Apply(current =&gt; current.TenantId),
    ///                     ObjectId = current.Apply(current =&gt; current.ObjectId),
    ///                     CertificatePermissions = 
    ///                     {
    ///                         "create",
    ///                         "delete",
    ///                         "deleteissuers",
    ///                         "get",
    ///                         "getissuers",
    ///                         "import",
    ///                         "list",
    ///                         "listissuers",
    ///                         "managecontacts",
    ///                         "manageissuers",
    ///                         "purge",
    ///                         "setissuers",
    ///                         "update",
    ///                     },
    ///                     KeyPermissions = 
    ///                     {
    ///                         "backup",
    ///                         "create",
    ///                         "decrypt",
    ///                         "delete",
    ///                         "encrypt",
    ///                         "get",
    ///                         "import",
    ///                         "list",
    ///                         "purge",
    ///                         "recover",
    ///                         "restore",
    ///                         "sign",
    ///                         "unwrapKey",
    ///                         "update",
    ///                         "verify",
    ///                         "wrapKey",
    ///                     },
    ///                     SecretPermissions = 
    ///                     {
    ///                         "backup",
    ///                         "delete",
    ///                         "get",
    ///                         "list",
    ///                         "purge",
    ///                         "recover",
    ///                         "restore",
    ///                         "set",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleCertificate = new Azure.KeyVault.Certificate("exampleCertificate", new Azure.KeyVault.CertificateArgs
    ///         {
    ///             KeyVaultId = exampleKeyVault.Id,
    ///             CertificatePolicy = new Azure.KeyVault.Inputs.CertificateCertificatePolicyArgs
    ///             {
    ///                 IssuerParameters = new Azure.KeyVault.Inputs.CertificateCertificatePolicyIssuerParametersArgs
    ///                 {
    ///                     Name = "Self",
    ///                 },
    ///                 KeyProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicyKeyPropertiesArgs
    ///                 {
    ///                     Exportable = true,
    ///                     KeySize = 2048,
    ///                     KeyType = "RSA",
    ///                     ReuseKey = true,
    ///                 },
    ///                 LifetimeActions = 
    ///                 {
    ///                     new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionArgs
    ///                     {
    ///                         Action = new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionActionArgs
    ///                         {
    ///                             ActionType = "AutoRenew",
    ///                         },
    ///                         Trigger = new Azure.KeyVault.Inputs.CertificateCertificatePolicyLifetimeActionTriggerArgs
    ///                         {
    ///                             DaysBeforeExpiry = 30,
    ///                         },
    ///                     },
    ///                 },
    ///                 SecretProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicySecretPropertiesArgs
    ///                 {
    ///                     ContentType = "application/x-pkcs12",
    ///                 },
    ///                 X509CertificateProperties = new Azure.KeyVault.Inputs.CertificateCertificatePolicyX509CertificatePropertiesArgs
    ///                 {
    ///                     ExtendedKeyUsages = 
    ///                     {
    ///                         "1.3.6.1.5.5.7.3.1",
    ///                     },
    ///                     KeyUsages = 
    ///                     {
    ///                         "cRLSign",
    ///                         "dataEncipherment",
    ///                         "digitalSignature",
    ///                         "keyAgreement",
    ///                         "keyCertSign",
    ///                         "keyEncipherment",
    ///                     },
    ///                     SubjectAlternativeNames = new Azure.KeyVault.Inputs.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs
    ///                     {
    ///                         DnsNames = 
    ///                         {
    ///                             "internal.contoso.com",
    ///                             "domain.hello.world",
    ///                         },
    ///                     },
    ///                     Subject = "CN=hello-world",
    ///                     ValidityInMonths = 12,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key Vault Certificates can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:keyvault/certificate:Certificate example "https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217"
    /// ```
    /// </summary>
    [AzureResourceType("azure:keyvault/certificate:Certificate")]
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Output("certificate")]
        public Output<Outputs.CertificateCertificate?> KeyVaultCertificate { get; private set; } = null!;

        /// <summary>
        /// A `certificate_attribute` block as defined below.
        /// </summary>
        [Output("certificateAttributes")]
        public Output<ImmutableArray<Outputs.CertificateCertificateAttribute>> CertificateAttributes { get; private set; } = null!;

        /// <summary>
        /// The raw Key Vault Certificate data represented as a hexadecimal string.
        /// </summary>
        [Output("certificateData")]
        public Output<string> CertificateData { get; private set; } = null!;

        /// <summary>
        /// The Base64 encoded Key Vault Certificate data.
        /// </summary>
        [Output("certificateDataBase64")]
        public Output<string> CertificateDataBase64 { get; private set; } = null!;

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Output("certificatePolicy")]
        public Output<Outputs.CertificateCertificatePolicy> CertificatePolicy { get; private set; } = null!;

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Key Vault Secret.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        /// </summary>
        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;

        /// <summary>
        /// The current version of the Key Vault Certificate.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The Base ID of the Key Vault Certificate.
        /// </summary>
        [Output("versionlessId")]
        public Output<string> VersionlessId { get; private set; } = null!;

        /// <summary>
        /// The Base ID of the Key Vault Secret.
        /// </summary>
        [Output("versionlessSecretId")]
        public Output<string> VersionlessSecretId { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("azure:keyvault/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("azure:keyvault/certificate:Certificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure:keyvault/certifiate:Certifiate"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertificateCertificateArgs>? KeyVaultCertificate { get; set; }

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Input("certificatePolicy")]
        public Input<Inputs.CertificateCertificatePolicyArgs>? CertificatePolicy { get; set; }

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `certificate` block as defined below, used to Import an existing certificate.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertificateCertificateGetArgs>? KeyVaultCertificate { get; set; }

        [Input("certificateAttributes")]
        private InputList<Inputs.CertificateCertificateAttributeGetArgs>? _certificateAttributes;

        /// <summary>
        /// A `certificate_attribute` block as defined below.
        /// </summary>
        public InputList<Inputs.CertificateCertificateAttributeGetArgs> CertificateAttributes
        {
            get => _certificateAttributes ?? (_certificateAttributes = new InputList<Inputs.CertificateCertificateAttributeGetArgs>());
            set => _certificateAttributes = value;
        }

        /// <summary>
        /// The raw Key Vault Certificate data represented as a hexadecimal string.
        /// </summary>
        [Input("certificateData")]
        public Input<string>? CertificateData { get; set; }

        /// <summary>
        /// The Base64 encoded Key Vault Certificate data.
        /// </summary>
        [Input("certificateDataBase64")]
        public Input<string>? CertificateDataBase64 { get; set; }

        /// <summary>
        /// A `certificate_policy` block as defined below.
        /// </summary>
        [Input("certificatePolicy")]
        public Input<Inputs.CertificateCertificatePolicyGetArgs>? CertificatePolicy { get; set; }

        /// <summary>
        /// The ID of the Key Vault where the Certificate should be created.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the associated Key Vault Secret.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        /// <summary>
        /// The current version of the Key Vault Certificate.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The Base ID of the Key Vault Certificate.
        /// </summary>
        [Input("versionlessId")]
        public Input<string>? VersionlessId { get; set; }

        /// <summary>
        /// The Base ID of the Key Vault Secret.
        /// </summary>
        [Input("versionlessSecretId")]
        public Input<string>? VersionlessSecretId { get; set; }

        public CertificateState()
        {
        }
    }
}
