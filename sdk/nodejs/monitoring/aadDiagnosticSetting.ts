// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Azure Active Directory Diagnostic Setting for Azure Monitor.
 *
 * !> **Authentication** The API for this resource does not support service principal authentication. This resource can only be used with Azure CLI authentication.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "west europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountKind: "StorageV2",
 *     accountReplicationType: "LRS",
 * });
 * const exampleAadDiagnosticSetting = new azure.monitoring.AadDiagnosticSetting("exampleAadDiagnosticSetting", {
 *     storageAccountId: exampleAccount.id,
 *     logs: [
 *         {
 *             category: "SignInLogs",
 *             enabled: true,
 *             retentionPolicy: {
 *                 enabled: true,
 *                 days: 1,
 *             },
 *         },
 *         {
 *             category: "AuditLogs",
 *             enabled: true,
 *             retentionPolicy: {
 *                 enabled: true,
 *                 days: 1,
 *             },
 *         },
 *         {
 *             category: "NonInteractiveUserSignInLogs",
 *             enabled: true,
 *             retentionPolicy: {
 *                 enabled: true,
 *                 days: 1,
 *             },
 *         },
 *         {
 *             category: "ServicePrincipalSignInLogs",
 *             enabled: true,
 *             retentionPolicy: {
 *                 enabled: true,
 *                 days: 1,
 *             },
 *         },
 *         {
 *             category: "ManagedIdentitySignInLogs",
 *             enabled: false,
 *             retentionPolicy: {},
 *         },
 *         {
 *             category: "ProvisioningLogs",
 *             enabled: false,
 *             retentionPolicy: {},
 *         },
 *         {
 *             category: "ADFSSignInLogs",
 *             enabled: false,
 *             retentionPolicy: {},
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Monitor Azure Active Directory Diagnostic Settings can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:monitoring/aadDiagnosticSetting:AadDiagnosticSetting example /providers/Microsoft.AADIAM/diagnosticSettings/setting1
 * ```
 */
export class AadDiagnosticSetting extends pulumi.CustomResource {
    /**
     * Get an existing AadDiagnosticSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AadDiagnosticSettingState, opts?: pulumi.CustomResourceOptions): AadDiagnosticSetting {
        return new AadDiagnosticSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:monitoring/aadDiagnosticSetting:AadDiagnosticSetting';

    /**
     * Returns true if the given object is an instance of AadDiagnosticSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AadDiagnosticSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AadDiagnosticSetting.__pulumiType;
    }

    /**
     * Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
     */
    public readonly eventhubAuthorizationRuleId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Event Hub where Diagnostics Data should be sent. If not specified, the default Event Hub will be used. Changing this forces a new resource to be created.
     */
    public readonly eventhubName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string | undefined>;
    /**
     * One or more `log` blocks as defined below.
     */
    public readonly logs!: pulumi.Output<outputs.monitoring.AadDiagnosticSettingLog[]>;
    /**
     * The name which should be used for this Monitor Azure Active Directory Diagnostic Setting. Changing this forces a new Monitor Azure Active Directory Diagnostic Setting to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;

    /**
     * Create a AadDiagnosticSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AadDiagnosticSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AadDiagnosticSettingArgs | AadDiagnosticSettingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AadDiagnosticSettingState | undefined;
            inputs["eventhubAuthorizationRuleId"] = state ? state.eventhubAuthorizationRuleId : undefined;
            inputs["eventhubName"] = state ? state.eventhubName : undefined;
            inputs["logAnalyticsWorkspaceId"] = state ? state.logAnalyticsWorkspaceId : undefined;
            inputs["logs"] = state ? state.logs : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as AadDiagnosticSettingArgs | undefined;
            if ((!args || args.logs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logs'");
            }
            inputs["eventhubAuthorizationRuleId"] = args ? args.eventhubAuthorizationRuleId : undefined;
            inputs["eventhubName"] = args ? args.eventhubName : undefined;
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["logs"] = args ? args.logs : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AadDiagnosticSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AadDiagnosticSetting resources.
 */
export interface AadDiagnosticSettingState {
    /**
     * Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
     */
    eventhubAuthorizationRuleId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Event Hub where Diagnostics Data should be sent. If not specified, the default Event Hub will be used. Changing this forces a new resource to be created.
     */
    eventhubName?: pulumi.Input<string>;
    /**
     * Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
     */
    logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * One or more `log` blocks as defined below.
     */
    logs?: pulumi.Input<pulumi.Input<inputs.monitoring.AadDiagnosticSettingLog>[]>;
    /**
     * The name which should be used for this Monitor Azure Active Directory Diagnostic Setting. Changing this forces a new Monitor Azure Active Directory Diagnostic Setting to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AadDiagnosticSetting resource.
 */
export interface AadDiagnosticSettingArgs {
    /**
     * Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
     */
    eventhubAuthorizationRuleId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Event Hub where Diagnostics Data should be sent. If not specified, the default Event Hub will be used. Changing this forces a new resource to be created.
     */
    eventhubName?: pulumi.Input<string>;
    /**
     * Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent.
     */
    logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * One or more `log` blocks as defined below.
     */
    logs: pulumi.Input<pulumi.Input<inputs.monitoring.AadDiagnosticSettingLog>[]>;
    /**
     * The name which should be used for this Monitor Azure Active Directory Diagnostic Setting. Changing this forces a new Monitor Azure Active Directory Diagnostic Setting to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Storage Account where logs should be sent. Changing this forces a new resource to be created.
     */
    storageAccountId?: pulumi.Input<string>;
}
