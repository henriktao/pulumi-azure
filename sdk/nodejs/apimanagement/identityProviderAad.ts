// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management AAD Identity Provider.
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
 *     publisherEmail: "company@mycompany.io",
 *     skuName: "Developer_1",
 * });
 * const exampleIdentityProviderAad = new azure.apimanagement.IdentityProviderAad("exampleIdentityProviderAad", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     apiManagementName: exampleService.name,
 *     clientId: "00000000-0000-0000-0000-000000000000",
 *     clientSecret: "00000000000000000000000000000000",
 *     allowedTenants: ["00000000-0000-0000-0000-000000000000"],
 * });
 * ```
 *
 * ## Import
 *
 * API Management AAD Identity Provider can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/identityProviderAad:IdentityProviderAad example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/aad
 * ```
 */
export class IdentityProviderAad extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProviderAad resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderAadState, opts?: pulumi.CustomResourceOptions): IdentityProviderAad {
        return new IdentityProviderAad(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/identityProviderAad:IdentityProviderAad';

    /**
     * Returns true if the given object is an instance of IdentityProviderAad.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProviderAad {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProviderAad.__pulumiType;
    }

    /**
     * List of allowed AAD Tenants.
     */
    public readonly allowedTenants!: pulumi.Output<string[]>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * Client Id of the Application in the AAD Identity Provider.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client secret of the Application in the AAD Identity Provider.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The AAD Tenant to use instead of Common when logging into Active Directory
     */
    public readonly signinTenant!: pulumi.Output<string | undefined>;

    /**
     * Create a IdentityProviderAad resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderAadArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderAadArgs | IdentityProviderAadState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityProviderAadState | undefined;
            resourceInputs["allowedTenants"] = state ? state.allowedTenants : undefined;
            resourceInputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["signinTenant"] = state ? state.signinTenant : undefined;
        } else {
            const args = argsOrState as IdentityProviderAadArgs | undefined;
            if ((!args || args.allowedTenants === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedTenants'");
            }
            if ((!args || args.apiManagementName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementName'");
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
            resourceInputs["allowedTenants"] = args ? args.allowedTenants : undefined;
            resourceInputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["signinTenant"] = args ? args.signinTenant : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IdentityProviderAad.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProviderAad resources.
 */
export interface IdentityProviderAadState {
    /**
     * List of allowed AAD Tenants.
     */
    allowedTenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    apiManagementName?: pulumi.Input<string>;
    /**
     * Client Id of the Application in the AAD Identity Provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Client secret of the Application in the AAD Identity Provider.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The AAD Tenant to use instead of Common when logging into Active Directory
     */
    signinTenant?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityProviderAad resource.
 */
export interface IdentityProviderAadArgs {
    /**
     * List of allowed AAD Tenants.
     */
    allowedTenants: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name of the API Management Service where this AAD Identity Provider should be created. Changing this forces a new resource to be created.
     */
    apiManagementName: pulumi.Input<string>;
    /**
     * Client Id of the Application in the AAD Identity Provider.
     */
    clientId: pulumi.Input<string>;
    /**
     * Client secret of the Application in the AAD Identity Provider.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The AAD Tenant to use instead of Common when logging into Active Directory
     */
    signinTenant?: pulumi.Input<string>;
}
