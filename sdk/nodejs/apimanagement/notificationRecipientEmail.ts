// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a API Management Notification Recipient Email.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     publisherName: "My Company",
 *     publisherEmail: "company@terraform.io",
 *     skuName: "Developer_1",
 * });
 * const exampleNotificationRecipientEmail = new azure.apimanagement.NotificationRecipientEmail("exampleNotificationRecipientEmail", {
 *     apiManagementId: exampleService.id,
 *     notificationType: "AccountClosedPublisher",
 *     email: "foo@bar.com",
 * });
 * ```
 *
 * ## Import
 *
 * API Management Notification Recipient Emails can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/notificationRecipientEmail:NotificationRecipientEmail example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientEmails/email1
 * ```
 */
export class NotificationRecipientEmail extends pulumi.CustomResource {
    /**
     * Get an existing NotificationRecipientEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationRecipientEmailState, opts?: pulumi.CustomResourceOptions): NotificationRecipientEmail {
        return new NotificationRecipientEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/notificationRecipientEmail:NotificationRecipientEmail';

    /**
     * Returns true if the given object is an instance of NotificationRecipientEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationRecipientEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationRecipientEmail.__pulumiType;
    }

    /**
     * The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    public readonly apiManagementId!: pulumi.Output<string>;
    /**
     * The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
     */
    public readonly notificationType!: pulumi.Output<string>;

    /**
     * Create a NotificationRecipientEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationRecipientEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationRecipientEmailArgs | NotificationRecipientEmailState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationRecipientEmailState | undefined;
            inputs["apiManagementId"] = state ? state.apiManagementId : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["notificationType"] = state ? state.notificationType : undefined;
        } else {
            const args = argsOrState as NotificationRecipientEmailArgs | undefined;
            if ((!args || args.apiManagementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementId'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.notificationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationType'");
            }
            inputs["apiManagementId"] = args ? args.apiManagementId : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["notificationType"] = args ? args.notificationType : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationRecipientEmail.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationRecipientEmail resources.
 */
export interface NotificationRecipientEmailState {
    /**
     * The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    apiManagementId?: pulumi.Input<string>;
    /**
     * The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    email?: pulumi.Input<string>;
    /**
     * The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
     */
    notificationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotificationRecipientEmail resource.
 */
export interface NotificationRecipientEmailArgs {
    /**
     * The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    apiManagementId: pulumi.Input<string>;
    /**
     * The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
     */
    email: pulumi.Input<string>;
    /**
     * The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are `AccountClosedPublisher`, `BCC`, `NewApplicationNotificationMessage`, `NewIssuePublisherNotificationMessage`, `PurchasePublisherNotificationMessage`, `QuotaLimitApproachingPublisherNotificationMessage`, and `RequestPublisherNotificationMessage`.
     */
    notificationType: pulumi.Input<string>;
}
