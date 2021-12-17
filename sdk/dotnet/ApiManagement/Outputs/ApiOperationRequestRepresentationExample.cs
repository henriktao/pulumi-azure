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
    public sealed class ApiOperationRequestRepresentationExample
    {
        /// <summary>
        /// A long description for this example.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A URL that points to the literal example.
        /// </summary>
        public readonly string? ExternalValue;
        /// <summary>
        /// The name of this example.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A short description for this example.
        /// </summary>
        public readonly string? Summary;
        /// <summary>
        /// The example of the representation.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ApiOperationRequestRepresentationExample(
            string? description,

            string? externalValue,

            string name,

            string? summary,

            string? value)
        {
            Description = description;
            ExternalValue = externalValue;
            Name = name;
            Summary = summary;
            Value = value;
        }
    }
}
