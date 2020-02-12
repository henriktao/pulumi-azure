// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an API Management User.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_user.html.markdown.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
     */
    public readonly confirmation!: pulumi.Output<string | undefined>;
    /**
     * The email address associated with this user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The first name for this user.
     */
    public readonly firstName!: pulumi.Output<string>;
    /**
     * The last name for this user.
     */
    public readonly lastName!: pulumi.Output<string>;
    /**
     * A note about this user.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * The password associated with this user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The state of this user. Possible values are `active`, `blocked` and `pending`.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["confirmation"] = state ? state.confirmation : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["firstName"] = state ? state.firstName : undefined;
            inputs["lastName"] = state ? state.lastName : undefined;
            inputs["note"] = state ? state.note : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.email === undefined) {
                throw new Error("Missing required property 'email'");
            }
            if (!args || args.firstName === undefined) {
                throw new Error("Missing required property 'firstName'");
            }
            if (!args || args.lastName === undefined) {
                throw new Error("Missing required property 'lastName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.userId === undefined) {
                throw new Error("Missing required property 'userId'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["confirmation"] = args ? args.confirmation : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["firstName"] = args ? args.firstName : undefined;
            inputs["lastName"] = args ? args.lastName : undefined;
            inputs["note"] = args ? args.note : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["userId"] = args ? args.userId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
     */
    readonly confirmation?: pulumi.Input<string>;
    /**
     * The email address associated with this user.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The first name for this user.
     */
    readonly firstName?: pulumi.Input<string>;
    /**
     * The last name for this user.
     */
    readonly lastName?: pulumi.Input<string>;
    /**
     * A note about this user.
     */
    readonly note?: pulumi.Input<string>;
    /**
     * The password associated with this user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The state of this user. Possible values are `active`, `blocked` and `pending`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The name of the API Management Service in which the User should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The kind of confirmation email which will be sent to this user. Possible values are `invite` and `signup`. Changing this forces a new resource to be created.
     */
    readonly confirmation?: pulumi.Input<string>;
    /**
     * The email address associated with this user.
     */
    readonly email: pulumi.Input<string>;
    /**
     * The first name for this user.
     */
    readonly firstName: pulumi.Input<string>;
    /**
     * The last name for this user.
     */
    readonly lastName: pulumi.Input<string>;
    /**
     * A note about this user.
     */
    readonly note?: pulumi.Input<string>;
    /**
     * The password associated with this user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The state of this user. Possible values are `active`, `blocked` and `pending`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The Identifier for this User, which must be unique within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly userId: pulumi.Input<string>;
}
