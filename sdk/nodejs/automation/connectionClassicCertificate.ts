// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Automation Connection with type `AzureClassicCertificate`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleClientConfig = azure.core.getClientConfig({});
 * const exampleAccount = new azure.automation.Account("exampleAccount", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: [{
 *         name: "Basic",
 *     }],
 * });
 * const exampleConnectionClassicCertificate = new azure.automation.ConnectionClassicCertificate("exampleConnectionClassicCertificate", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     automationAccountName: exampleAccount.name,
 *     certificateAssetName: "cert1",
 *     subscriptionName: "subs1",
 *     subscriptionId: exampleClientConfig.then(exampleClientConfig => exampleClientConfig.subscriptionId),
 * });
 * ```
 *
 * ## Import
 *
 * Automation Connection can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:automation/connectionClassicCertificate:ConnectionClassicCertificate conn1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
 * ```
 */
export class ConnectionClassicCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionClassicCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionClassicCertificateState, opts?: pulumi.CustomResourceOptions): ConnectionClassicCertificate {
        return new ConnectionClassicCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:automation/connectionClassicCertificate:ConnectionClassicCertificate';

    /**
     * Returns true if the given object is an instance of ConnectionClassicCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionClassicCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionClassicCertificate.__pulumiType;
    }

    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    public readonly automationAccountName!: pulumi.Output<string>;
    /**
     * The name of the certificate asset.
     */
    public readonly certificateAssetName!: pulumi.Output<string>;
    /**
     * A description for this Connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The id of subscription.
     */
    public readonly subscriptionId!: pulumi.Output<string>;
    /**
     * The name of subscription.
     */
    public readonly subscriptionName!: pulumi.Output<string>;

    /**
     * Create a ConnectionClassicCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionClassicCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionClassicCertificateArgs | ConnectionClassicCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionClassicCertificateState | undefined;
            resourceInputs["automationAccountName"] = state ? state.automationAccountName : undefined;
            resourceInputs["certificateAssetName"] = state ? state.certificateAssetName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["subscriptionId"] = state ? state.subscriptionId : undefined;
            resourceInputs["subscriptionName"] = state ? state.subscriptionName : undefined;
        } else {
            const args = argsOrState as ConnectionClassicCertificateArgs | undefined;
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.certificateAssetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAssetName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.subscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            if ((!args || args.subscriptionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            resourceInputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            resourceInputs["certificateAssetName"] = args ? args.certificateAssetName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            resourceInputs["subscriptionName"] = args ? args.subscriptionName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectionClassicCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionClassicCertificate resources.
 */
export interface ConnectionClassicCertificateState {
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    automationAccountName?: pulumi.Input<string>;
    /**
     * The name of the certificate asset.
     */
    certificateAssetName?: pulumi.Input<string>;
    /**
     * A description for this Connection.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * The id of subscription.
     */
    subscriptionId?: pulumi.Input<string>;
    /**
     * The name of subscription.
     */
    subscriptionName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectionClassicCertificate resource.
 */
export interface ConnectionClassicCertificateArgs {
    /**
     * The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * The name of the certificate asset.
     */
    certificateAssetName: pulumi.Input<string>;
    /**
     * A description for this Connection.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the name of the Connection. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The id of subscription.
     */
    subscriptionId: pulumi.Input<string>;
    /**
     * The name of subscription.
     */
    subscriptionName: pulumi.Input<string>;
}
