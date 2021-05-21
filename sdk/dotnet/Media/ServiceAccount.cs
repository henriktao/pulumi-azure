// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media
{
    /// <summary>
    /// Manages a Media Services Account.
    /// 
    /// ## Example Usage
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
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "GRS",
    ///         });
    ///         var exampleServiceAccount = new Azure.Media.ServiceAccount("exampleServiceAccount", new Azure.Media.ServiceAccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             StorageAccounts = 
    ///             {
    ///                 new Azure.Media.Inputs.ServiceAccountStorageAccountArgs
    ///                 {
    ///                     Id = exampleAccount.Id,
    ///                     IsPrimary = true,
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
    /// Media Services Accounts can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:media/serviceAccount:ServiceAccount account /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/mediaservices/account1
    /// ```
    /// </summary>
    [AzureResourceType("azure:media/serviceAccount:ServiceAccount")]
    public partial class ServiceAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServiceAccountIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// An `key_delivery_access_control` block is documented below.
        /// </summary>
        [Output("keyDeliveryAccessControl")]
        public Output<Outputs.ServiceAccountKeyDeliveryAccessControl> KeyDeliveryAccessControl { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.ServiceAccountStorageAccount>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage authentication type. 
        /// Possible value is  `ManagedIdentity` or `System`.
        /// </summary>
        [Output("storageAuthenticationType")]
        public Output<string> StorageAuthenticationType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAccount(string name, ServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("azure:media/serviceAccount:ServiceAccount", name, args ?? new ServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAccount(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:media/serviceAccount:ServiceAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure:mediaservices/account:Account"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAccount Get(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceAccount(name, id, state, options);
        }
    }

    public sealed class ServiceAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceAccountIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// An `key_delivery_access_control` block is documented below.
        /// </summary>
        [Input("keyDeliveryAccessControl")]
        public Input<Inputs.ServiceAccountKeyDeliveryAccessControlArgs>? KeyDeliveryAccessControl { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("storageAccounts", required: true)]
        private InputList<Inputs.ServiceAccountStorageAccountArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceAccountStorageAccountArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.ServiceAccountStorageAccountArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// Specifies the storage authentication type. 
        /// Possible value is  `ManagedIdentity` or `System`.
        /// </summary>
        [Input("storageAuthenticationType")]
        public Input<string>? StorageAuthenticationType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceAccountArgs()
        {
        }
    }

    public sealed class ServiceAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceAccountIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// An `key_delivery_access_control` block is documented below.
        /// </summary>
        [Input("keyDeliveryAccessControl")]
        public Input<Inputs.ServiceAccountKeyDeliveryAccessControlGetArgs>? KeyDeliveryAccessControl { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Media Services Account. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.ServiceAccountStorageAccountGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceAccountStorageAccountGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.ServiceAccountStorageAccountGetArgs>());
            set => _storageAccounts = value;
        }

        /// <summary>
        /// Specifies the storage authentication type. 
        /// Possible value is  `ManagedIdentity` or `System`.
        /// </summary>
        [Input("storageAuthenticationType")]
        public Input<string>? StorageAuthenticationType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceAccountState()
        {
        }
    }
}
