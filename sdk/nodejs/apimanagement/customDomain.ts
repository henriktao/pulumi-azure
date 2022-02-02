// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a API Management Custom Domain.
 *
 * ## Disclaimers
 *
 * > **Note:** It's possible to define Custom Domains both within the `azure.apimanagement.Service` resource via the `hostnameConfigurations` block and by using this resource. However it's not possible to use both methods to manage Custom Domains within an API Management Service, since there will be conflicts.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleService = new azure.apimanagement.Service("exampleService", {
 *     location: azurerm_resource_group.test.location,
 *     resourceGroupName: azurerm_resource_group.test.name,
 *     publisherName: "pub1",
 *     publisherEmail: "pub1@email.com",
 *     skuName: "Developer_1",
 * });
 * const exampleCertificate = new azure.keyvault.Certificate("exampleCertificate", {
 *     keyVaultId: azurerm_key_vault.test.id,
 *     certificatePolicy: {
 *         issuerParameters: {
 *             name: "Self",
 *         },
 *         keyProperties: {
 *             exportable: true,
 *             keySize: 2048,
 *             keyType: "RSA",
 *             reuseKey: true,
 *         },
 *         lifetimeActions: [{
 *             action: {
 *                 actionType: "AutoRenew",
 *             },
 *             trigger: {
 *                 daysBeforeExpiry: 30,
 *             },
 *         }],
 *         secretProperties: {
 *             contentType: "application/x-pkcs12",
 *         },
 *         x509CertificateProperties: {
 *             keyUsages: [
 *                 "cRLSign",
 *                 "dataEncipherment",
 *                 "digitalSignature",
 *                 "keyAgreement",
 *                 "keyCertSign",
 *                 "keyEncipherment",
 *             ],
 *             subject: "CN=api.example.com",
 *             validityInMonths: 12,
 *             subjectAlternativeNames: {
 *                 dnsNames: [
 *                     "api.example.com",
 *                     "portal.example.com",
 *                 ],
 *             },
 *         },
 *     },
 * });
 * const exampleCustomDomain = new azure.apimanagement.CustomDomain("exampleCustomDomain", {
 *     apiManagementId: exampleService.id,
 *     proxies: [{
 *         hostName: "api.example.com",
 *         keyVaultId: azurerm_key_vault_certificate.test.secret_id,
 *     }],
 *     developerPortals: [{
 *         hostName: "portal.example.com",
 *         keyVaultId: azurerm_key_vault_certificate.test.secret_id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * API Management Custom Domains can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/customDomain:CustomDomain example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
 * ```
 */
export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomDomainState, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/customDomain:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    /**
     * The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
     */
    public readonly apiManagementId!: pulumi.Output<string>;
    /**
     * One or more `developerPortal` blocks as defined below.
     */
    public readonly developerPortals!: pulumi.Output<outputs.apimanagement.CustomDomainDeveloperPortal[] | undefined>;
    /**
     * One or more `management` blocks as defined below.
     */
    public readonly managements!: pulumi.Output<outputs.apimanagement.CustomDomainManagement[] | undefined>;
    /**
     * One or more `portal` blocks as defined below.
     */
    public readonly portals!: pulumi.Output<outputs.apimanagement.CustomDomainPortal[] | undefined>;
    /**
     * One or more `proxy` blocks as defined below.
     *
     * @deprecated `proxy` is deprecated and will be renamed to `gateway` in version 3.0 of the AzureRM provider
     */
    public readonly proxies!: pulumi.Output<outputs.apimanagement.CustomDomainProxy[] | undefined>;
    /**
     * One or more `scm` blocks as defined below.
     */
    public readonly scms!: pulumi.Output<outputs.apimanagement.CustomDomainScm[] | undefined>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomDomainArgs | CustomDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomDomainState | undefined;
            resourceInputs["apiManagementId"] = state ? state.apiManagementId : undefined;
            resourceInputs["developerPortals"] = state ? state.developerPortals : undefined;
            resourceInputs["managements"] = state ? state.managements : undefined;
            resourceInputs["portals"] = state ? state.portals : undefined;
            resourceInputs["proxies"] = state ? state.proxies : undefined;
            resourceInputs["scms"] = state ? state.scms : undefined;
        } else {
            const args = argsOrState as CustomDomainArgs | undefined;
            if ((!args || args.apiManagementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiManagementId'");
            }
            resourceInputs["apiManagementId"] = args ? args.apiManagementId : undefined;
            resourceInputs["developerPortals"] = args ? args.developerPortals : undefined;
            resourceInputs["managements"] = args ? args.managements : undefined;
            resourceInputs["portals"] = args ? args.portals : undefined;
            resourceInputs["proxies"] = args ? args.proxies : undefined;
            resourceInputs["scms"] = args ? args.scms : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomDomain resources.
 */
export interface CustomDomainState {
    /**
     * The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
     */
    apiManagementId?: pulumi.Input<string>;
    /**
     * One or more `developerPortal` blocks as defined below.
     */
    developerPortals?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainDeveloperPortal>[]>;
    /**
     * One or more `management` blocks as defined below.
     */
    managements?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainManagement>[]>;
    /**
     * One or more `portal` blocks as defined below.
     */
    portals?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainPortal>[]>;
    /**
     * One or more `proxy` blocks as defined below.
     *
     * @deprecated `proxy` is deprecated and will be renamed to `gateway` in version 3.0 of the AzureRM provider
     */
    proxies?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainProxy>[]>;
    /**
     * One or more `scm` blocks as defined below.
     */
    scms?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainScm>[]>;
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    /**
     * The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
     */
    apiManagementId: pulumi.Input<string>;
    /**
     * One or more `developerPortal` blocks as defined below.
     */
    developerPortals?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainDeveloperPortal>[]>;
    /**
     * One or more `management` blocks as defined below.
     */
    managements?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainManagement>[]>;
    /**
     * One or more `portal` blocks as defined below.
     */
    portals?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainPortal>[]>;
    /**
     * One or more `proxy` blocks as defined below.
     *
     * @deprecated `proxy` is deprecated and will be renamed to `gateway` in version 3.0 of the AzureRM provider
     */
    proxies?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainProxy>[]>;
    /**
     * One or more `scm` blocks as defined below.
     */
    scms?: pulumi.Input<pulumi.Input<inputs.apimanagement.CustomDomainScm>[]>;
}
