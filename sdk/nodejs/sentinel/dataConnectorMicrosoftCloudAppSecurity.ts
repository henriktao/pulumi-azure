// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Microsoft Cloud App Security Data Connector.
 *
 *  !> **NOTE:** This resource requires that [Enterprise Mobility + Security E5](https://www.microsoft.com/en-us/microsoft-365/enterprise-mobility-security) is enabled on the tenant being connected to.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "west europe"});
 * const exampleAnalyticsWorkspace = new azure.operationalinsights.AnalyticsWorkspace("exampleAnalyticsWorkspace", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: "PerGB2018",
 * });
 * const exampleAnalyticsSolution = new azure.operationalinsights.AnalyticsSolution("exampleAnalyticsSolution", {
 *     solutionName: "SecurityInsights",
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     workspaceResourceId: exampleAnalyticsWorkspace.id,
 *     workspaceName: exampleAnalyticsWorkspace.name,
 *     plan: {
 *         publisher: "Microsoft",
 *         product: "OMSGallery/SecurityInsights",
 *     },
 * });
 * const exampleDataConnectorMicrosoftCloudAppSecurity = new azure.sentinel.DataConnectorMicrosoftCloudAppSecurity("exampleDataConnectorMicrosoftCloudAppSecurity", {logAnalyticsWorkspaceId: exampleAnalyticsSolution.workspaceResourceId});
 * ```
 *
 * ## Import
 *
 * Microsoft Cloud App Security Data Connectors can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:sentinel/dataConnectorMicrosoftCloudAppSecurity:DataConnectorMicrosoftCloudAppSecurity example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
 * ```
 */
export class DataConnectorMicrosoftCloudAppSecurity extends pulumi.CustomResource {
    /**
     * Get an existing DataConnectorMicrosoftCloudAppSecurity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataConnectorMicrosoftCloudAppSecurityState, opts?: pulumi.CustomResourceOptions): DataConnectorMicrosoftCloudAppSecurity {
        return new DataConnectorMicrosoftCloudAppSecurity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:sentinel/dataConnectorMicrosoftCloudAppSecurity:DataConnectorMicrosoftCloudAppSecurity';

    /**
     * Returns true if the given object is an instance of DataConnectorMicrosoftCloudAppSecurity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataConnectorMicrosoftCloudAppSecurity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataConnectorMicrosoftCloudAppSecurity.__pulumiType;
    }

    /**
     * Should the alerts be enabled? Defaults to `true`.
     */
    public readonly alertsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Should the Discovery Logs be enabled? Defaults to `true`.
     */
    public readonly discoveryLogsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Log Analytics Workspace that this Microsoft Cloud App Security Data Connector resides in. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    public readonly logAnalyticsWorkspaceId!: pulumi.Output<string>;
    /**
     * The name which should be used for this Microsoft Cloud App Security Data Connector. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Tenant that this Microsoft Cloud App Security Data Connector connects to. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a DataConnectorMicrosoftCloudAppSecurity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataConnectorMicrosoftCloudAppSecurityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataConnectorMicrosoftCloudAppSecurityArgs | DataConnectorMicrosoftCloudAppSecurityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataConnectorMicrosoftCloudAppSecurityState | undefined;
            inputs["alertsEnabled"] = state ? state.alertsEnabled : undefined;
            inputs["discoveryLogsEnabled"] = state ? state.discoveryLogsEnabled : undefined;
            inputs["logAnalyticsWorkspaceId"] = state ? state.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as DataConnectorMicrosoftCloudAppSecurityArgs | undefined;
            if ((!args || args.logAnalyticsWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logAnalyticsWorkspaceId'");
            }
            inputs["alertsEnabled"] = args ? args.alertsEnabled : undefined;
            inputs["discoveryLogsEnabled"] = args ? args.discoveryLogsEnabled : undefined;
            inputs["logAnalyticsWorkspaceId"] = args ? args.logAnalyticsWorkspaceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataConnectorMicrosoftCloudAppSecurity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataConnectorMicrosoftCloudAppSecurity resources.
 */
export interface DataConnectorMicrosoftCloudAppSecurityState {
    /**
     * Should the alerts be enabled? Defaults to `true`.
     */
    alertsEnabled?: pulumi.Input<boolean>;
    /**
     * Should the Discovery Logs be enabled? Defaults to `true`.
     */
    discoveryLogsEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Log Analytics Workspace that this Microsoft Cloud App Security Data Connector resides in. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    logAnalyticsWorkspaceId?: pulumi.Input<string>;
    /**
     * The name which should be used for this Microsoft Cloud App Security Data Connector. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Tenant that this Microsoft Cloud App Security Data Connector connects to. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataConnectorMicrosoftCloudAppSecurity resource.
 */
export interface DataConnectorMicrosoftCloudAppSecurityArgs {
    /**
     * Should the alerts be enabled? Defaults to `true`.
     */
    alertsEnabled?: pulumi.Input<boolean>;
    /**
     * Should the Discovery Logs be enabled? Defaults to `true`.
     */
    discoveryLogsEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the Log Analytics Workspace that this Microsoft Cloud App Security Data Connector resides in. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    logAnalyticsWorkspaceId: pulumi.Input<string>;
    /**
     * The name which should be used for this Microsoft Cloud App Security Data Connector. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Tenant that this Microsoft Cloud App Security Data Connector connects to. Changing this forces a new Microsoft Cloud App Security Data Connector to be created.
     */
    tenantId?: pulumi.Input<string>;
}
