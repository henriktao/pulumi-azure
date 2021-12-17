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
    public sealed class ApiOperationRequestRepresentation
    {
        /// <summary>
        /// The Content Type of this representation, such as `application/json`.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// One or more `example` blocks as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiOperationRequestRepresentationExample> Examples;
        /// <summary>
        /// One or more `form_parameter` block as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiOperationRequestRepresentationFormParameter> FormParameters;
        /// <summary>
        /// An example of this representation.
        /// </summary>
        public readonly string? Sample;
        /// <summary>
        /// The ID of an API Management Schema which represents this Response.
        /// </summary>
        public readonly string? SchemaId;
        /// <summary>
        /// The Type Name defined by the Schema.
        /// </summary>
        public readonly string? TypeName;

        [OutputConstructor]
        private ApiOperationRequestRepresentation(
            string contentType,

            ImmutableArray<Outputs.ApiOperationRequestRepresentationExample> examples,

            ImmutableArray<Outputs.ApiOperationRequestRepresentationFormParameter> formParameters,

            string? sample,

            string? schemaId,

            string? typeName)
        {
            ContentType = contentType;
            Examples = examples;
            FormParameters = formParameters;
            Sample = sample;
            SchemaId = schemaId;
            TypeName = typeName;
        }
    }
}
