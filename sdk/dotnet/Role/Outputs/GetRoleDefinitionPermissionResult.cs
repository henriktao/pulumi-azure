// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Role.Outputs
{

    [OutputType]
    public sealed class GetRoleDefinitionPermissionResult
    {
        /// <summary>
        /// a list of actions supported by this role
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        public readonly ImmutableArray<string> DataActions;
        /// <summary>
        /// a list of actions which are denied by this role
        /// </summary>
        public readonly ImmutableArray<string> NotActions;
        public readonly ImmutableArray<string> NotDataActions;

        [OutputConstructor]
        private GetRoleDefinitionPermissionResult(
            ImmutableArray<string> actions,

            ImmutableArray<string> dataActions,

            ImmutableArray<string> notActions,

            ImmutableArray<string> notDataActions)
        {
            Actions = actions;
            DataActions = dataActions;
            NotActions = notActions;
            NotDataActions = notDataActions;
        }
    }
}
