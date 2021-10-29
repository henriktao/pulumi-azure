// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Synapse.Inputs
{

    public sealed class WorkspaceAzureDevopsRepoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Azure DevOps account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies the collaboration branch of the repository to get code from.
        /// </summary>
        [Input("branchName", required: true)]
        public Input<string> BranchName { get; set; } = null!;

        /// <summary>
        /// The last commit ID.
        /// </summary>
        [Input("lastCommitId")]
        public Input<string>? LastCommitId { get; set; }

        /// <summary>
        /// Specifies the name of the Azure DevOps project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the git repository.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        /// <summary>
        /// Specifies the root folder within the repository. Set to `/` for the top level.
        /// </summary>
        [Input("rootFolder", required: true)]
        public Input<string> RootFolder { get; set; } = null!;

        /// <summary>
        /// the ID of the tenant for the Azure DevOps account.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public WorkspaceAzureDevopsRepoArgs()
        {
        }
    }
}
