// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Trigger Schedule inside a Azure Data Factory.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "northeurope",
 * });
 * const exampleFactory = new azure.datafactory.Factory("example", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const testPipeline = new azure.datafactory.Pipeline("test", {
 *     dataFactoryName: azurerm_data_factory_test.name,
 *     resourceGroupName: azurerm_resource_group_test.name,
 * });
 * const testTriggerSchedule = new azure.datafactory.TriggerSchedule("test", {
 *     dataFactoryName: azurerm_data_factory_test.name,
 *     frequency: "Day",
 *     interval: 5,
 *     pipelineName: testPipeline.name,
 *     resourceGroupName: azurerm_resource_group_test.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_factory_trigger_schedule.html.markdown.
 */
export class TriggerSchedule extends pulumi.CustomResource {
    /**
     * Get an existing TriggerSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerScheduleState, opts?: pulumi.CustomResourceOptions): TriggerSchedule {
        return new TriggerSchedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:datafactory/triggerSchedule:TriggerSchedule';

    /**
     * Returns true if the given object is an instance of TriggerSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TriggerSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TriggerSchedule.__pulumiType;
    }

    /**
     * List of tags that can be used for describing the Data Factory Schedule Trigger.
     */
    public readonly annotations!: pulumi.Output<string[] | undefined>;
    /**
     * The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
     */
    public readonly dataFactoryName!: pulumi.Output<string>;
    /**
     * The time the Schedule Trigger should end. The time will be represented in UTC.
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
     */
    public readonly frequency!: pulumi.Output<string | undefined>;
    /**
     * The interval for how often the trigger occurs. This defaults to 1.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Data Factory Pipeline name that the trigger will act on.
     */
    public readonly pipelineName!: pulumi.Output<string>;
    /**
     * The pipeline parameters that the trigger will act upon.
     */
    public readonly pipelineParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
     */
    public readonly startTime!: pulumi.Output<string>;

    /**
     * Create a TriggerSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerScheduleArgs | TriggerScheduleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TriggerScheduleState | undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["dataFactoryName"] = state ? state.dataFactoryName : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["frequency"] = state ? state.frequency : undefined;
            inputs["interval"] = state ? state.interval : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pipelineName"] = state ? state.pipelineName : undefined;
            inputs["pipelineParameters"] = state ? state.pipelineParameters : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as TriggerScheduleArgs | undefined;
            if (!args || args.dataFactoryName === undefined) {
                throw new Error("Missing required property 'dataFactoryName'");
            }
            if (!args || args.pipelineName === undefined) {
                throw new Error("Missing required property 'pipelineName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["dataFactoryName"] = args ? args.dataFactoryName : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["frequency"] = args ? args.frequency : undefined;
            inputs["interval"] = args ? args.interval : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pipelineName"] = args ? args.pipelineName : undefined;
            inputs["pipelineParameters"] = args ? args.pipelineParameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TriggerSchedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TriggerSchedule resources.
 */
export interface TriggerScheduleState {
    /**
     * List of tags that can be used for describing the Data Factory Schedule Trigger.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
     */
    readonly dataFactoryName?: pulumi.Input<string>;
    /**
     * The time the Schedule Trigger should end. The time will be represented in UTC.
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
     */
    readonly frequency?: pulumi.Input<string>;
    /**
     * The interval for how often the trigger occurs. This defaults to 1.
     */
    readonly interval?: pulumi.Input<number>;
    /**
     * Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Data Factory Pipeline name that the trigger will act on.
     */
    readonly pipelineName?: pulumi.Input<string>;
    /**
     * The pipeline parameters that the trigger will act upon.
     */
    readonly pipelineParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
     */
    readonly startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TriggerSchedule resource.
 */
export interface TriggerScheduleArgs {
    /**
     * List of tags that can be used for describing the Data Factory Schedule Trigger.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
     */
    readonly dataFactoryName: pulumi.Input<string>;
    /**
     * The time the Schedule Trigger should end. The time will be represented in UTC.
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
     */
    readonly frequency?: pulumi.Input<string>;
    /**
     * The interval for how often the trigger occurs. This defaults to 1.
     */
    readonly interval?: pulumi.Input<number>;
    /**
     * Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Data Factory Pipeline name that the trigger will act on.
     */
    readonly pipelineName: pulumi.Input<string>;
    /**
     * The pipeline parameters that the trigger will act upon.
     */
    readonly pipelineParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
     */
    readonly startTime?: pulumi.Input<string>;
}
