// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter
{
    /// <summary>
    /// Manages the subscription's Security Center Contact.
    /// 
    /// &gt; **NOTE:** Owner access permission is required.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/security_center_contact.html.markdown.
    /// </summary>
    public partial class Contact : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to send security alerts notifications to the security contact.
        /// </summary>
        [Output("alertNotifications")]
        public Output<bool> AlertNotifications { get; private set; } = null!;

        /// <summary>
        /// Whether to send security alerts notifications to subscription admins.
        /// </summary>
        [Output("alertsToAdmins")]
        public Output<bool> AlertsToAdmins { get; private set; } = null!;

        /// <summary>
        /// The email of the Security Center Contact.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The phone number of the Security Center Contact.
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;


        /// <summary>
        /// Create a Contact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Contact(string name, ContactArgs args, CustomResourceOptions? options = null)
            : base("azure:securitycenter/contact:Contact", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Contact(string name, Input<string> id, ContactState? state = null, CustomResourceOptions? options = null)
            : base("azure:securitycenter/contact:Contact", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Contact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Contact Get(string name, Input<string> id, ContactState? state = null, CustomResourceOptions? options = null)
        {
            return new Contact(name, id, state, options);
        }
    }

    public sealed class ContactArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to send security alerts notifications to the security contact.
        /// </summary>
        [Input("alertNotifications", required: true)]
        public Input<bool> AlertNotifications { get; set; } = null!;

        /// <summary>
        /// Whether to send security alerts notifications to subscription admins.
        /// </summary>
        [Input("alertsToAdmins", required: true)]
        public Input<bool> AlertsToAdmins { get; set; } = null!;

        /// <summary>
        /// The email of the Security Center Contact.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// The phone number of the Security Center Contact.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        public ContactArgs()
        {
        }
    }

    public sealed class ContactState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to send security alerts notifications to the security contact.
        /// </summary>
        [Input("alertNotifications")]
        public Input<bool>? AlertNotifications { get; set; }

        /// <summary>
        /// Whether to send security alerts notifications to subscription admins.
        /// </summary>
        [Input("alertsToAdmins")]
        public Input<bool>? AlertsToAdmins { get; set; }

        /// <summary>
        /// The email of the Security Center Contact.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The phone number of the Security Center Contact.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        public ContactState()
        {
        }
    }
}
