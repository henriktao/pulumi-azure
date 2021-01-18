// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EventGrid Domain
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US 2"});
 * const exampleDomain = new azure.eventgrid.Domain("exampleDomain", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tags: {
 *         environment: "Production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * EventGrid Domains can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:eventgrid/domain:Domain domain1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:eventgrid/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * The Endpoint associated with the EventGrid Domain.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * One or more `inboundIpRule` blocks as defined below.
     */
    public readonly inboundIpRules!: pulumi.Output<outputs.eventgrid.DomainInboundIpRule[] | undefined>;
    /**
     * A `inputMappingDefaultValues` block as defined below.
     */
    public readonly inputMappingDefaultValues!: pulumi.Output<outputs.eventgrid.DomainInputMappingDefaultValues | undefined>;
    /**
     * A `inputMappingFields` block as defined below.
     */
    public readonly inputMappingFields!: pulumi.Output<outputs.eventgrid.DomainInputMappingFields | undefined>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    public readonly inputSchema!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Primary Shared Access Key associated with the EventGrid Domain.
     */
    public /*out*/ readonly primaryAccessKey!: pulumi.Output<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The Secondary Shared Access Key associated with the EventGrid Domain.
     */
    public /*out*/ readonly secondaryAccessKey!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DomainState | undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["inboundIpRules"] = state ? state.inboundIpRules : undefined;
            inputs["inputMappingDefaultValues"] = state ? state.inputMappingDefaultValues : undefined;
            inputs["inputMappingFields"] = state ? state.inputMappingFields : undefined;
            inputs["inputSchema"] = state ? state.inputSchema : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryAccessKey"] = state ? state.primaryAccessKey : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secondaryAccessKey"] = state ? state.secondaryAccessKey : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["inboundIpRules"] = args ? args.inboundIpRules : undefined;
            inputs["inputMappingDefaultValues"] = args ? args.inputMappingDefaultValues : undefined;
            inputs["inputMappingFields"] = args ? args.inputMappingFields : undefined;
            inputs["inputSchema"] = args ? args.inputSchema : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["primaryAccessKey"] = undefined /*out*/;
            inputs["secondaryAccessKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure:eventhub/domain:Domain" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The Endpoint associated with the EventGrid Domain.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * One or more `inboundIpRule` blocks as defined below.
     */
    readonly inboundIpRules?: pulumi.Input<pulumi.Input<inputs.eventgrid.DomainInboundIpRule>[]>;
    /**
     * A `inputMappingDefaultValues` block as defined below.
     */
    readonly inputMappingDefaultValues?: pulumi.Input<inputs.eventgrid.DomainInputMappingDefaultValues>;
    /**
     * A `inputMappingFields` block as defined below.
     */
    readonly inputMappingFields?: pulumi.Input<inputs.eventgrid.DomainInputMappingFields>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    readonly inputSchema?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Primary Shared Access Key associated with the EventGrid Domain.
     */
    readonly primaryAccessKey?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    readonly publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The Secondary Shared Access Key associated with the EventGrid Domain.
     */
    readonly secondaryAccessKey?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * One or more `inboundIpRule` blocks as defined below.
     */
    readonly inboundIpRules?: pulumi.Input<pulumi.Input<inputs.eventgrid.DomainInboundIpRule>[]>;
    /**
     * A `inputMappingDefaultValues` block as defined below.
     */
    readonly inputMappingDefaultValues?: pulumi.Input<inputs.eventgrid.DomainInputMappingDefaultValues>;
    /**
     * A `inputMappingFields` block as defined below.
     */
    readonly inputMappingFields?: pulumi.Input<inputs.eventgrid.DomainInputMappingFields>;
    /**
     * Specifies the schema in which incoming events will be published to this domain. Allowed values are `CloudEventSchemaV1_0`, `CustomEventSchema`, or `EventGridSchema`. Defaults to `eventgridschema`. Changing this forces a new resource to be created.
     */
    readonly inputSchema?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the EventGrid Domain resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether or not public network access is allowed for this server. Defaults to `true`.
     */
    readonly publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
