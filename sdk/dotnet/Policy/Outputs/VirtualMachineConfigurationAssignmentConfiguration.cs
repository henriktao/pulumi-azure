// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy.Outputs
{

    [OutputType]
    public sealed class VirtualMachineConfigurationAssignmentConfiguration
    {
        /// <summary>
        /// The assignment type for the Guest Configuration Assignment. Possible values are `Audit`, `ApplyAndAutoCorrect`, `ApplyAndMonitor` and `DeployAndAutoCorrect`.
        /// </summary>
        public readonly string? AssignmentType;
        /// <summary>
        /// The content hash for the Guest Configuration package.
        /// </summary>
        public readonly string? ContentHash;
        /// <summary>
        /// The content URI where the Guest Configuration package is stored.
        /// </summary>
        public readonly string? ContentUri;
        /// <summary>
        /// This field is no longer used and will be removed in the next major version of the Azure Provider.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// One or more `parameter` blocks which define what configuration parameters and values against.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineConfigurationAssignmentConfigurationParameter> Parameters;
        /// <summary>
        /// The version of the Guest Configuration that will be assigned in this Guest Configuration Assignment.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private VirtualMachineConfigurationAssignmentConfiguration(
            string? assignmentType,

            string? contentHash,

            string? contentUri,

            string? name,

            ImmutableArray<Outputs.VirtualMachineConfigurationAssignmentConfigurationParameter> parameters,

            string? version)
        {
            AssignmentType = assignmentType;
            ContentHash = contentHash;
            ContentUri = contentUri;
            Name = name;
            Parameters = parameters;
            Version = version;
        }
    }
}
