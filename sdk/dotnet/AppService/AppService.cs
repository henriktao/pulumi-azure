// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService
{
    /// <summary>
    /// Manages an App Service (within an App Service Plan).
    /// 
    /// &gt; **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azure.appservice.AppService` resource will be overwritten when promoting a Slot using the `azure.appservice.ActiveSlot` resource.
    /// 
    /// ## Example Usage
    /// 
    /// This example provisions a Windows App Service.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var examplePlan = new Azure.AppService.Plan("examplePlan", new Azure.AppService.PlanArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Sku = new Azure.AppService.Inputs.PlanSkuArgs
    ///             {
    ///                 Tier = "Standard",
    ///                 Size = "S1",
    ///             },
    ///         });
    ///         var exampleAppService = new Azure.AppService.AppService("exampleAppService", new Azure.AppService.AppServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             AppServicePlanId = examplePlan.Id,
    ///             SiteConfig = new Azure.AppService.Inputs.AppServiceSiteConfigArgs
    ///             {
    ///                 DotnetFrameworkVersion = "v4.0",
    ///                 ScmType = "LocalGit",
    ///             },
    ///             AppSettings = 
    ///             {
    ///                 { "SOME_KEY", "some-value" },
    ///             },
    ///             ConnectionStrings = 
    ///             {
    ///                 new Azure.AppService.Inputs.AppServiceConnectionStringArgs
    ///                 {
    ///                     Name = "Database",
    ///                     Type = "SQLServer",
    ///                     Value = "Server=some-server.mydomain.com;Integrated Security=SSPI",
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
    /// App Services can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:appservice/appService:AppService instance1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1
    /// ```
    /// </summary>
    [AzureResourceType("azure:appservice/appService:AppService")]
    public partial class AppService : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this App Service.
        /// </summary>
        [Output("appServicePlanId")]
        public Output<string> AppServicePlanId { get; private set; } = null!;

        /// <summary>
        /// A key-value pair of App Settings.
        /// </summary>
        [Output("appSettings")]
        public Output<ImmutableDictionary<string, string>> AppSettings { get; private set; } = null!;

        /// <summary>
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Output("authSettings")]
        public Output<Outputs.AppServiceAuthSettings> AuthSettings { get; private set; } = null!;

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Output("backup")]
        public Output<Outputs.AppServiceBackup?> Backup { get; private set; } = null!;

        /// <summary>
        /// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Output("clientAffinityEnabled")]
        public Output<bool?> ClientAffinityEnabled { get; private set; } = null!;

        /// <summary>
        /// Does the App Service require client certificates for incoming requests? Defaults to `false`.
        /// </summary>
        [Output("clientCertEnabled")]
        public Output<bool?> ClientCertEnabled { get; private set; } = null!;

        /// <summary>
        /// One or more `connection_string` blocks as defined below.
        /// </summary>
        [Output("connectionStrings")]
        public Output<ImmutableArray<Outputs.AppServiceConnectionString>> ConnectionStrings { get; private set; } = null!;

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Output("customDomainVerificationId")]
        public Output<string> CustomDomainVerificationId { get; private set; } = null!;

        /// <summary>
        /// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
        /// </summary>
        [Output("defaultSiteHostname")]
        public Output<string> DefaultSiteHostname { get; private set; } = null!;

        /// <summary>
        /// Is the App Service Enabled?
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Can the App Service only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Output("httpsOnly")]
        public Output<bool?> HttpsOnly { get; private set; } = null!;

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.AppServiceIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. [For more information see - Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity)
        /// </summary>
        [Output("keyVaultReferenceIdentityId")]
        public Output<string> KeyVaultReferenceIdentityId { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A `logs` block as defined below.
        /// </summary>
        [Output("logs")]
        public Output<Outputs.AppServiceLogs> Logs { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12"]`
        /// </summary>
        [Output("outboundIpAddressLists")]
        public Output<ImmutableArray<string>> OutboundIpAddressLists { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        /// </summary>
        [Output("outboundIpAddresses")]
        public Output<string> OutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` - not all of which are necessarily in use. Superset of `outbound_ip_address_list`.
        /// </summary>
        [Output("possibleOutboundIpAddressLists")]
        public Output<ImmutableArray<string>> PossibleOutboundIpAddressLists { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        [Output("possibleOutboundIpAddresses")]
        public Output<string> PossibleOutboundIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the App Service.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `site_config` block as defined below.
        /// </summary>
        [Output("siteConfig")]
        public Output<Outputs.AppServiceSiteConfig> SiteConfig { get; private set; } = null!;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        [Output("siteCredentials")]
        public Output<ImmutableArray<Outputs.AppServiceSiteCredential>> SiteCredentials { get; private set; } = null!;

        /// <summary>
        /// A Source Control block as defined below
        /// </summary>
        [Output("sourceControl")]
        public Output<Outputs.AppServiceSourceControl> SourceControl { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.AppServiceStorageAccount>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AppService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppService(string name, AppServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:appservice/appService:AppService", name, args ?? new AppServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppService(string name, Input<string> id, AppServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:appservice/appService:AppService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppService Get(string name, Input<string> id, AppServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new AppService(name, id, state, options);
        }
    }

    public sealed class AppServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this App Service.
        /// </summary>
        [Input("appServicePlanId", required: true)]
        public Input<string> AppServicePlanId { get; set; } = null!;

        [Input("appSettings")]
        private InputMap<string>? _appSettings;

        /// <summary>
        /// A key-value pair of App Settings.
        /// </summary>
        public InputMap<string> AppSettings
        {
            get => _appSettings ?? (_appSettings = new InputMap<string>());
            set => _appSettings = value;
        }

        /// <summary>
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Input("authSettings")]
        public Input<Inputs.AppServiceAuthSettingsArgs>? AuthSettings { get; set; }

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Input("backup")]
        public Input<Inputs.AppServiceBackupArgs>? Backup { get; set; }

        /// <summary>
        /// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// Does the App Service require client certificates for incoming requests? Defaults to `false`.
        /// </summary>
        [Input("clientCertEnabled")]
        public Input<bool>? ClientCertEnabled { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.AppServiceConnectionStringArgs>? _connectionStrings;

        /// <summary>
        /// One or more `connection_string` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceConnectionStringArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.AppServiceConnectionStringArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// Is the App Service Enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the App Service only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AppServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. [For more information see - Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity)
        /// </summary>
        [Input("keyVaultReferenceIdentityId")]
        public Input<string>? KeyVaultReferenceIdentityId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `logs` block as defined below.
        /// </summary>
        [Input("logs")]
        public Input<Inputs.AppServiceLogsArgs>? Logs { get; set; }

        /// <summary>
        /// Specifies the name of the App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the App Service.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `site_config` block as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.AppServiceSiteConfigArgs>? SiteConfig { get; set; }

        /// <summary>
        /// A Source Control block as defined below
        /// </summary>
        [Input("sourceControl")]
        public Input<Inputs.AppServiceSourceControlArgs>? SourceControl { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.AppServiceStorageAccountArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceStorageAccountArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.AppServiceStorageAccountArgs>());
            set => _storageAccounts = value;
        }

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

        public AppServiceArgs()
        {
        }
    }

    public sealed class AppServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the App Service Plan within which to create this App Service.
        /// </summary>
        [Input("appServicePlanId")]
        public Input<string>? AppServicePlanId { get; set; }

        [Input("appSettings")]
        private InputMap<string>? _appSettings;

        /// <summary>
        /// A key-value pair of App Settings.
        /// </summary>
        public InputMap<string> AppSettings
        {
            get => _appSettings ?? (_appSettings = new InputMap<string>());
            set => _appSettings = value;
        }

        /// <summary>
        /// A `auth_settings` block as defined below.
        /// </summary>
        [Input("authSettings")]
        public Input<Inputs.AppServiceAuthSettingsGetArgs>? AuthSettings { get; set; }

        /// <summary>
        /// A `backup` block as defined below.
        /// </summary>
        [Input("backup")]
        public Input<Inputs.AppServiceBackupGetArgs>? Backup { get; set; }

        /// <summary>
        /// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
        /// </summary>
        [Input("clientAffinityEnabled")]
        public Input<bool>? ClientAffinityEnabled { get; set; }

        /// <summary>
        /// Does the App Service require client certificates for incoming requests? Defaults to `false`.
        /// </summary>
        [Input("clientCertEnabled")]
        public Input<bool>? ClientCertEnabled { get; set; }

        [Input("connectionStrings")]
        private InputList<Inputs.AppServiceConnectionStringGetArgs>? _connectionStrings;

        /// <summary>
        /// One or more `connection_string` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceConnectionStringGetArgs> ConnectionStrings
        {
            get => _connectionStrings ?? (_connectionStrings = new InputList<Inputs.AppServiceConnectionStringGetArgs>());
            set => _connectionStrings = value;
        }

        /// <summary>
        /// An identifier used by App Service to perform domain ownership verification via DNS TXT record.
        /// </summary>
        [Input("customDomainVerificationId")]
        public Input<string>? CustomDomainVerificationId { get; set; }

        /// <summary>
        /// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
        /// </summary>
        [Input("defaultSiteHostname")]
        public Input<string>? DefaultSiteHostname { get; set; }

        /// <summary>
        /// Is the App Service Enabled?
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Can the App Service only be accessed via HTTPS? Defaults to `false`.
        /// </summary>
        [Input("httpsOnly")]
        public Input<bool>? HttpsOnly { get; set; }

        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.AppServiceIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The User Assigned Identity Id used for looking up KeyVault secrets. The identity must be assigned to the application. [For more information see - Access vaults with a user-assigned identity](https://docs.microsoft.com/en-us/azure/app-service/app-service-key-vault-references#access-vaults-with-a-user-assigned-identity)
        /// </summary>
        [Input("keyVaultReferenceIdentityId")]
        public Input<string>? KeyVaultReferenceIdentityId { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A `logs` block as defined below.
        /// </summary>
        [Input("logs")]
        public Input<Inputs.AppServiceLogsGetArgs>? Logs { get; set; }

        /// <summary>
        /// Specifies the name of the App Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outboundIpAddressLists")]
        private InputList<string>? _outboundIpAddressLists;

        /// <summary>
        /// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12"]`
        /// </summary>
        public InputList<string> OutboundIpAddressLists
        {
            get => _outboundIpAddressLists ?? (_outboundIpAddressLists = new InputList<string>());
            set => _outboundIpAddressLists = value;
        }

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
        /// </summary>
        [Input("outboundIpAddresses")]
        public Input<string>? OutboundIpAddresses { get; set; }

        [Input("possibleOutboundIpAddressLists")]
        private InputList<string>? _possibleOutboundIpAddressLists;

        /// <summary>
        /// A list of outbound IP addresses - such as `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` - not all of which are necessarily in use. Superset of `outbound_ip_address_list`.
        /// </summary>
        public InputList<string> PossibleOutboundIpAddressLists
        {
            get => _possibleOutboundIpAddressLists ?? (_possibleOutboundIpAddressLists = new InputList<string>());
            set => _possibleOutboundIpAddressLists = value;
        }

        /// <summary>
        /// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
        /// </summary>
        [Input("possibleOutboundIpAddresses")]
        public Input<string>? PossibleOutboundIpAddresses { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the App Service.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `site_config` block as defined below.
        /// </summary>
        [Input("siteConfig")]
        public Input<Inputs.AppServiceSiteConfigGetArgs>? SiteConfig { get; set; }

        [Input("siteCredentials")]
        private InputList<Inputs.AppServiceSiteCredentialGetArgs>? _siteCredentials;

        /// <summary>
        /// A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
        /// </summary>
        public InputList<Inputs.AppServiceSiteCredentialGetArgs> SiteCredentials
        {
            get => _siteCredentials ?? (_siteCredentials = new InputList<Inputs.AppServiceSiteCredentialGetArgs>());
            set => _siteCredentials = value;
        }

        /// <summary>
        /// A Source Control block as defined below
        /// </summary>
        [Input("sourceControl")]
        public Input<Inputs.AppServiceSourceControlGetArgs>? SourceControl { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.AppServiceStorageAccountGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        public InputList<Inputs.AppServiceStorageAccountGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.AppServiceStorageAccountGetArgs>());
            set => _storageAccounts = value;
        }

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

        public AppServiceState()
        {
        }
    }
}
