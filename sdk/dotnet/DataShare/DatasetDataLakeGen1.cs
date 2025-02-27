// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataShare
{
    /// <summary>
    /// Manages a Data Share Data Lake Gen1 Dataset.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleAccount = new Azure.DataShare.Account("exampleAccount", new Azure.DataShare.AccountArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Identity = new Azure.DataShare.Inputs.AccountIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///         });
    ///         var exampleShare = new Azure.DataShare.Share("exampleShare", new Azure.DataShare.ShareArgs
    ///         {
    ///             AccountId = exampleAccount.Id,
    ///             Kind = "CopyBased",
    ///         });
    ///         var exampleStore = new Azure.DataLake.Store("exampleStore", new Azure.DataLake.StoreArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             FirewallState = "Disabled",
    ///         });
    ///         var exampleStoreFile = new Azure.DataLake.StoreFile("exampleStoreFile", new Azure.DataLake.StoreFileArgs
    ///         {
    ///             AccountName = exampleStore.Name,
    ///             LocalFilePath = "./example/myfile.txt",
    ///             RemoteFilePath = "/example/myfile.txt",
    ///         });
    ///         var exampleServicePrincipal = AzureAD.GetServicePrincipal.Invoke(new AzureAD.GetServicePrincipalInvokeArgs
    ///         {
    ///             DisplayName = exampleAccount.Name,
    ///         });
    ///         var exampleAssignment = new Azure.Authorization.Assignment("exampleAssignment", new Azure.Authorization.AssignmentArgs
    ///         {
    ///             Scope = exampleStore.Id,
    ///             RoleDefinitionName = "Owner",
    ///             PrincipalId = exampleServicePrincipal.Apply(exampleServicePrincipal =&gt; exampleServicePrincipal.ObjectId),
    ///         });
    ///         var exampleDatasetDataLakeGen1 = new Azure.DataShare.DatasetDataLakeGen1("exampleDatasetDataLakeGen1", new Azure.DataShare.DatasetDataLakeGen1Args
    ///         {
    ///             DataShareId = exampleShare.Id,
    ///             DataLakeStoreId = exampleStore.Id,
    ///             FileName = "myfile.txt",
    ///             FolderPath = "example",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleAssignment,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Data Share Data Lake Gen1 Datasets can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1 example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataShare/accounts/account1/shares/share1/dataSets/dataSet1
    /// ```
    /// </summary>
    [AzureResourceType("azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1")]
    public partial class DatasetDataLakeGen1 : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of the Data Lake Store to be shared with the receiver.
        /// </summary>
        [Output("dataLakeStoreId")]
        public Output<string> DataLakeStoreId { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Output("dataShareId")]
        public Output<string> DataShareId { get; private set; } = null!;

        /// <summary>
        /// The displayed name of the Data Share Dataset.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Output("fileName")]
        public Output<string?> FileName { get; private set; } = null!;

        /// <summary>
        /// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Output("folderPath")]
        public Output<string> FolderPath { get; private set; } = null!;

        /// <summary>
        /// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetDataLakeGen1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetDataLakeGen1(string name, DatasetDataLakeGen1Args args, CustomResourceOptions? options = null)
            : base("azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1", name, args ?? new DatasetDataLakeGen1Args(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetDataLakeGen1(string name, Input<string> id, DatasetDataLakeGen1State? state = null, CustomResourceOptions? options = null)
            : base("azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetDataLakeGen1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetDataLakeGen1 Get(string name, Input<string> id, DatasetDataLakeGen1State? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetDataLakeGen1(name, id, state, options);
        }
    }

    public sealed class DatasetDataLakeGen1Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the Data Lake Store to be shared with the receiver.
        /// </summary>
        [Input("dataLakeStoreId", required: true)]
        public Input<string> DataLakeStoreId { get; set; } = null!;

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("dataShareId", required: true)]
        public Input<string> DataShareId { get; set; } = null!;

        /// <summary>
        /// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("folderPath", required: true)]
        public Input<string> FolderPath { get; set; } = null!;

        /// <summary>
        /// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DatasetDataLakeGen1Args()
        {
        }
    }

    public sealed class DatasetDataLakeGen1State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the Data Lake Store to be shared with the receiver.
        /// </summary>
        [Input("dataLakeStoreId")]
        public Input<string>? DataLakeStoreId { get; set; }

        /// <summary>
        /// The resource ID of the Data Share where this Data Share Data Lake Gen1 Dataset should be created. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("dataShareId")]
        public Input<string>? DataShareId { get; set; }

        /// <summary>
        /// The displayed name of the Data Share Dataset.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The file name of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// The folder path of the data lake store to be shared with the receiver. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("folderPath")]
        public Input<string>? FolderPath { get; set; }

        /// <summary>
        /// The name of the Data Share Data Lake Gen1 Dataset. Changing this forces a new Data Share Data Lake Gen1 Dataset to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DatasetDataLakeGen1State()
        {
        }
    }
}
