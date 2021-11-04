// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management Azure AD B2C Identity Provider.
 *
 * ## Import
 *
 * API Management Azure AD B2C Identity Providers can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/identityProviderAadb2c:IdentityProviderAadb2c example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service1/identityProviders/AadB2C
 * ```
 */
export class IdentityProviderAadb2c extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProviderAadb2c resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderAadb2cState, opts?: pulumi.CustomResourceOptions): IdentityProviderAadb2c {
        return new IdentityProviderAadb2c(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/identityProviderAadb2c:IdentityProviderAadb2c';

    /**
     * Returns true if the given object is an instance of IdentityProviderAadb2c.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProviderAadb2c {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProviderAadb2c.__pulumiType;
    }

    /**
     * The allowed AAD tenant, usually your B2C tenant domain.
     */
    public readonly allowedTenant!: pulumi.Output<string>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * OpenID Connect discovery endpoint hostname, usually your b2clogin.com domain.
     */
    public readonly authority!: pulumi.Output<string>;
    /**
     * Client ID of the Application in your B2C tenant.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client secret of the Application in your B2C tenant.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * Password reset Policy Name.
     */
    public readonly passwordResetPolicy!: pulumi.Output<string | undefined>;
    /**
     * Profile editing Policy Name.
     */
    public readonly profileEditingPolicy!: pulumi.Output<string | undefined>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Signin Policy Name.
     */
    public readonly signinPolicy!: pulumi.Output<string>;
    /**
     * The tenant to use instead of Common when logging into Active Directory, usually your B2C tenant domain.
     */
    public readonly signinTenant!: pulumi.Output<string>;
    /**
     * Signup Policy Name.
     */
    public readonly signupPolicy!: pulumi.Output<string>;

    /**
     * Create a IdentityProviderAadb2c resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderAadb2cArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderAadb2cArgs | IdentityProviderAadb2cState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityProviderAadb2cState | undefined;
            inputs["allowedTenant"] = state ? state.allowedTenant : undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["authority"] = state ? state.authority : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["passwordResetPolicy"] = state ? state.passwordResetPolicy : undefined;
            inputs["profileEditingPolicy"] = state ? state.profileEditingPolicy : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["signinPolicy"] = state ? state.signinPolicy : undefined;
            inputs["signinTenant"] = state ? state.signinTenant : undefined;
            inputs["signupPolicy"] = state ? state.signupPolicy : undefined;
        } else {
            const args = argsOrState as IdentityProviderAadb2cArgs | undefined;
            if ((!args || args.allowedTenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedTenant'");
            }
            if ((!args || args.apiManagementName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if ((!args || args.authority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authority'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.signinPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signinPolicy'");
            }
            if ((!args || args.signinTenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signinTenant'");
            }
            if ((!args || args.signupPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signupPolicy'");
            }
            inputs["allowedTenant"] = args ? args.allowedTenant : undefined;
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["authority"] = args ? args.authority : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["passwordResetPolicy"] = args ? args.passwordResetPolicy : undefined;
            inputs["profileEditingPolicy"] = args ? args.profileEditingPolicy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["signinPolicy"] = args ? args.signinPolicy : undefined;
            inputs["signinTenant"] = args ? args.signinTenant : undefined;
            inputs["signupPolicy"] = args ? args.signupPolicy : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IdentityProviderAadb2c.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProviderAadb2c resources.
 */
export interface IdentityProviderAadb2cState {
    /**
     * The allowed AAD tenant, usually your B2C tenant domain.
     */
    allowedTenant?: pulumi.Input<string>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    apiManagementName?: pulumi.Input<string>;
    /**
     * OpenID Connect discovery endpoint hostname, usually your b2clogin.com domain.
     */
    authority?: pulumi.Input<string>;
    /**
     * Client ID of the Application in your B2C tenant.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Client secret of the Application in your B2C tenant.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Password reset Policy Name.
     */
    passwordResetPolicy?: pulumi.Input<string>;
    /**
     * Profile editing Policy Name.
     */
    profileEditingPolicy?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Signin Policy Name.
     */
    signinPolicy?: pulumi.Input<string>;
    /**
     * The tenant to use instead of Common when logging into Active Directory, usually your B2C tenant domain.
     */
    signinTenant?: pulumi.Input<string>;
    /**
     * Signup Policy Name.
     */
    signupPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityProviderAadb2c resource.
 */
export interface IdentityProviderAadb2cArgs {
    /**
     * The allowed AAD tenant, usually your B2C tenant domain.
     */
    allowedTenant: pulumi.Input<string>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    apiManagementName: pulumi.Input<string>;
    /**
     * OpenID Connect discovery endpoint hostname, usually your b2clogin.com domain.
     */
    authority: pulumi.Input<string>;
    /**
     * Client ID of the Application in your B2C tenant.
     */
    clientId: pulumi.Input<string>;
    /**
     * Client secret of the Application in your B2C tenant.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * Password reset Policy Name.
     */
    passwordResetPolicy?: pulumi.Input<string>;
    /**
     * Profile editing Policy Name.
     */
    profileEditingPolicy?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Signin Policy Name.
     */
    signinPolicy: pulumi.Input<string>;
    /**
     * The tenant to use instead of Common when logging into Active Directory, usually your B2C tenant domain.
     */
    signinTenant: pulumi.Input<string>;
    /**
     * Signup Policy Name.
     */
    signupPolicy: pulumi.Input<string>;
}
