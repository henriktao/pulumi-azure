// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Inputs
{

    public sealed class LinkedServiceAzureSqlDatabaseKeyVaultConnectionStringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of an existing Key Vault Data Factory Linked Service.
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public Input<string> LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// Specifies the secret name in Azure Key Vault that stores SQL Server connection string.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public LinkedServiceAzureSqlDatabaseKeyVaultConnectionStringArgs()
        {
        }
    }
}
