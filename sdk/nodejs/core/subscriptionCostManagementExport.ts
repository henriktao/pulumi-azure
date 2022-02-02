// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Cost Management Export for a Subscription.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleSubscription = azure.core.getSubscription({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleContainer = new azure.storage.Container("exampleContainer", {storageAccountName: azurerm_storage_account.test.name});
 * const exampleSubscriptionCostManagementExport = new azure.core.SubscriptionCostManagementExport("exampleSubscriptionCostManagementExport", {
 *     subscriptionId: azurerm_subscription.example.id,
 *     recurrenceType: "Monthly",
 *     recurrencePeriodStartDate: "2020-08-18T00:00:00Z",
 *     recurrencePeriodEndDate: "2020-09-18T00:00:00Z",
 *     exportDataStorageLocation: {
 *         containerId: exampleContainer.resourceManagerId,
 *         rootFolderPath: "/root/updated",
 *     },
 *     exportDataOptions: {
 *         type: "Usage",
 *         timeFrame: "WeekToDate",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Subscription Cost Management Exports can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:core/subscriptionCostManagementExport:SubscriptionCostManagementExport example /subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/exports/export1
 * ```
 */
export class SubscriptionCostManagementExport extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionCostManagementExport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionCostManagementExportState, opts?: pulumi.CustomResourceOptions): SubscriptionCostManagementExport {
        return new SubscriptionCostManagementExport(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:core/subscriptionCostManagementExport:SubscriptionCostManagementExport';

    /**
     * Returns true if the given object is an instance of SubscriptionCostManagementExport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionCostManagementExport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionCostManagementExport.__pulumiType;
    }

    /**
     * Is the cost management export active? Default is `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * A `exportDataOptions` block as defined below.
     */
    public readonly exportDataOptions!: pulumi.Output<outputs.core.SubscriptionCostManagementExportExportDataOptions>;
    /**
     * A `exportDataStorageLocation` block as defined below.
     */
    public readonly exportDataStorageLocation!: pulumi.Output<outputs.core.SubscriptionCostManagementExportExportDataStorageLocation>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly recurrencePeriodEndDate!: pulumi.Output<string>;
    /**
     * The date the export will start capturing information.
     */
    public readonly recurrencePeriodStartDate!: pulumi.Output<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    public readonly recurrenceType!: pulumi.Output<string>;
    /**
     * The id of the subscription on which to create an export.
     */
    public readonly subscriptionId!: pulumi.Output<string>;

    /**
     * Create a SubscriptionCostManagementExport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionCostManagementExportArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionCostManagementExportArgs | SubscriptionCostManagementExportState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionCostManagementExportState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["exportDataOptions"] = state ? state.exportDataOptions : undefined;
            resourceInputs["exportDataStorageLocation"] = state ? state.exportDataStorageLocation : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recurrencePeriodEndDate"] = state ? state.recurrencePeriodEndDate : undefined;
            resourceInputs["recurrencePeriodStartDate"] = state ? state.recurrencePeriodStartDate : undefined;
            resourceInputs["recurrenceType"] = state ? state.recurrenceType : undefined;
            resourceInputs["subscriptionId"] = state ? state.subscriptionId : undefined;
        } else {
            const args = argsOrState as SubscriptionCostManagementExportArgs | undefined;
            if ((!args || args.exportDataOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exportDataOptions'");
            }
            if ((!args || args.exportDataStorageLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exportDataStorageLocation'");
            }
            if ((!args || args.recurrencePeriodEndDate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recurrencePeriodEndDate'");
            }
            if ((!args || args.recurrencePeriodStartDate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recurrencePeriodStartDate'");
            }
            if ((!args || args.recurrenceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recurrenceType'");
            }
            if ((!args || args.subscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["exportDataOptions"] = args ? args.exportDataOptions : undefined;
            resourceInputs["exportDataStorageLocation"] = args ? args.exportDataStorageLocation : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recurrencePeriodEndDate"] = args ? args.recurrencePeriodEndDate : undefined;
            resourceInputs["recurrencePeriodStartDate"] = args ? args.recurrencePeriodStartDate : undefined;
            resourceInputs["recurrenceType"] = args ? args.recurrenceType : undefined;
            resourceInputs["subscriptionId"] = args ? args.subscriptionId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubscriptionCostManagementExport.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionCostManagementExport resources.
 */
export interface SubscriptionCostManagementExportState {
    /**
     * Is the cost management export active? Default is `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * A `exportDataOptions` block as defined below.
     */
    exportDataOptions?: pulumi.Input<inputs.core.SubscriptionCostManagementExportExportDataOptions>;
    /**
     * A `exportDataStorageLocation` block as defined below.
     */
    exportDataStorageLocation?: pulumi.Input<inputs.core.SubscriptionCostManagementExportExportDataStorageLocation>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    recurrencePeriodEndDate?: pulumi.Input<string>;
    /**
     * The date the export will start capturing information.
     */
    recurrencePeriodStartDate?: pulumi.Input<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    recurrenceType?: pulumi.Input<string>;
    /**
     * The id of the subscription on which to create an export.
     */
    subscriptionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionCostManagementExport resource.
 */
export interface SubscriptionCostManagementExportArgs {
    /**
     * Is the cost management export active? Default is `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * A `exportDataOptions` block as defined below.
     */
    exportDataOptions: pulumi.Input<inputs.core.SubscriptionCostManagementExportExportDataOptions>;
    /**
     * A `exportDataStorageLocation` block as defined below.
     */
    exportDataStorageLocation: pulumi.Input<inputs.core.SubscriptionCostManagementExportExportDataStorageLocation>;
    /**
     * Specifies the name of the Cost Management Export. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    recurrencePeriodEndDate: pulumi.Input<string>;
    /**
     * The date the export will start capturing information.
     */
    recurrencePeriodStartDate: pulumi.Input<string>;
    /**
     * How often the requested information will be exported. Valid values include `Annually`, `Daily`, `Monthly`, `Weekly`.
     */
    recurrenceType: pulumi.Input<string>;
    /**
     * The id of the subscription on which to create an export.
     */
    subscriptionId: pulumi.Input<string>;
}
