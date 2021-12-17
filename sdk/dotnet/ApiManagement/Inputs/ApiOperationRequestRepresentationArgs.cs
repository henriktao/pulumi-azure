// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ApiOperationRequestRepresentationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Content Type of this representation, such as `application/json`.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        [Input("examples")]
        private InputList<Inputs.ApiOperationRequestRepresentationExampleArgs>? _examples;

        /// <summary>
        /// One or more `example` blocks as defined above.
        /// </summary>
        public InputList<Inputs.ApiOperationRequestRepresentationExampleArgs> Examples
        {
            get => _examples ?? (_examples = new InputList<Inputs.ApiOperationRequestRepresentationExampleArgs>());
            set => _examples = value;
        }

        [Input("formParameters")]
        private InputList<Inputs.ApiOperationRequestRepresentationFormParameterArgs>? _formParameters;

        /// <summary>
        /// One or more `form_parameter` block as defined above.
        /// </summary>
        public InputList<Inputs.ApiOperationRequestRepresentationFormParameterArgs> FormParameters
        {
            get => _formParameters ?? (_formParameters = new InputList<Inputs.ApiOperationRequestRepresentationFormParameterArgs>());
            set => _formParameters = value;
        }

        /// <summary>
        /// An example of this representation.
        /// </summary>
        [Input("sample")]
        public Input<string>? Sample { get; set; }

        /// <summary>
        /// The ID of an API Management Schema which represents this Response.
        /// </summary>
        [Input("schemaId")]
        public Input<string>? SchemaId { get; set; }

        /// <summary>
        /// The Type Name defined by the Schema.
        /// </summary>
        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        public ApiOperationRequestRepresentationArgs()
        {
        }
    }
}
