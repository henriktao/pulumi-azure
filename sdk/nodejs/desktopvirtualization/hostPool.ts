// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Virtual Desktop Host Pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleHostPool = new azure.desktopvirtualization.HostPool("exampleHostPool", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     friendlyName: "pooleddepthfirst",
 *     validateEnvironment: true,
 *     startVmOnConnect: true,
 *     customRdpProperties: "audiocapturemode:i:1;audiomode:i:0;",
 *     description: "Acceptance Test: A pooled host pool - pooleddepthfirst",
 *     type: "Pooled",
 *     maximumSessionsAllowed: 50,
 *     loadBalancerType: "DepthFirst",
 * });
 * ```
 *
 * ## Import
 *
 * Virtual Desktop Host Pools can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:desktopvirtualization/hostPool:HostPool example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DesktopVirtualization/hostpools/myhostpool
 * ```
 */
export class HostPool extends pulumi.CustomResource {
    /**
     * Get an existing HostPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostPoolState, opts?: pulumi.CustomResourceOptions): HostPool {
        return new HostPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:desktopvirtualization/hostPool:HostPool';

    /**
     * Returns true if the given object is an instance of HostPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostPool.__pulumiType;
    }

    /**
     * A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
     */
    public readonly customRdpProperties!: pulumi.Output<string | undefined>;
    /**
     * A description for the Virtual Desktop Host Pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    public readonly loadBalancerType!: pulumi.Output<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    public readonly maximumSessionsAllowed!: pulumi.Output<number | undefined>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    public readonly personalDesktopAssignmentType!: pulumi.Output<string | undefined>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    public readonly preferredAppGroupType!: pulumi.Output<string | undefined>;
    /**
     * This block is now non-functional and will be removed in version 3.0 of the Azure Provider - use the
     * `azurerm_virtual_desktop_host_pool_registration_info` resource instead.
     */
    public readonly registrationInfo!: pulumi.Output<outputs.desktopvirtualization.HostPoolRegistrationInfo | undefined>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Enables or disables the Start VM on Connection Feature. Defaults to `false`.
     */
    public readonly startVmOnConnect!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Allows you to test service changes before they are deployed to production. Defaults to `false`.
     */
    public readonly validateEnvironment!: pulumi.Output<boolean | undefined>;

    /**
     * Create a HostPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostPoolArgs | HostPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostPoolState | undefined;
            resourceInputs["customRdpProperties"] = state ? state.customRdpProperties : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["maximumSessionsAllowed"] = state ? state.maximumSessionsAllowed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["personalDesktopAssignmentType"] = state ? state.personalDesktopAssignmentType : undefined;
            resourceInputs["preferredAppGroupType"] = state ? state.preferredAppGroupType : undefined;
            resourceInputs["registrationInfo"] = state ? state.registrationInfo : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["startVmOnConnect"] = state ? state.startVmOnConnect : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["validateEnvironment"] = state ? state.validateEnvironment : undefined;
        } else {
            const args = argsOrState as HostPoolArgs | undefined;
            if ((!args || args.loadBalancerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["customRdpProperties"] = args ? args.customRdpProperties : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maximumSessionsAllowed"] = args ? args.maximumSessionsAllowed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["personalDesktopAssignmentType"] = args ? args.personalDesktopAssignmentType : undefined;
            resourceInputs["preferredAppGroupType"] = args ? args.preferredAppGroupType : undefined;
            resourceInputs["registrationInfo"] = args ? args.registrationInfo : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["startVmOnConnect"] = args ? args.startVmOnConnect : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validateEnvironment"] = args ? args.validateEnvironment : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostPool resources.
 */
export interface HostPoolState {
    /**
     * A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
     */
    customRdpProperties?: pulumi.Input<string>;
    /**
     * A description for the Virtual Desktop Host Pool.
     */
    description?: pulumi.Input<string>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    friendlyName?: pulumi.Input<string>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    loadBalancerType?: pulumi.Input<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    maximumSessionsAllowed?: pulumi.Input<number>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    personalDesktopAssignmentType?: pulumi.Input<string>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    preferredAppGroupType?: pulumi.Input<string>;
    /**
     * This block is now non-functional and will be removed in version 3.0 of the Azure Provider - use the
     * `azurerm_virtual_desktop_host_pool_registration_info` resource instead.
     */
    registrationInfo?: pulumi.Input<inputs.desktopvirtualization.HostPoolRegistrationInfo>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Enables or disables the Start VM on Connection Feature. Defaults to `false`.
     */
    startVmOnConnect?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * Allows you to test service changes before they are deployed to production. Defaults to `false`.
     */
    validateEnvironment?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a HostPool resource.
 */
export interface HostPoolArgs {
    /**
     * A valid custom RDP properties string for the Virtual Desktop Host Pool, available properties can be [found in this article](https://docs.microsoft.com/en-us/windows-server/remote/remote-desktop-services/clients/rdp-files).
     */
    customRdpProperties?: pulumi.Input<string>;
    /**
     * A description for the Virtual Desktop Host Pool.
     */
    description?: pulumi.Input<string>;
    /**
     * A friendly name for the Virtual Desktop Host Pool.
     */
    friendlyName?: pulumi.Input<string>;
    /**
     * `BreadthFirst` load balancing distributes new user sessions across all available session hosts in the host pool.
     * `DepthFirst` load balancing distributes new user sessions to an available session host with the highest number of connections but has not reached its maximum session limit threshold.
     * `Persistent` should be used if the host pool type is `Personal`
     */
    loadBalancerType: pulumi.Input<string>;
    /**
     * The location/region where the Virtual Desktop Host Pool is
     * located. Changing the location/region forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * A valid integer value from 0 to 999999 for the maximum number of users that have concurrent sessions on a session host.
     * Should only be set if the `type` of your Virtual Desktop Host Pool is `Pooled`.
     */
    maximumSessionsAllowed?: pulumi.Input<number>;
    /**
     * The name of the Virtual Desktop Host Pool. Changing the name
     * forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * `Automatic` assignment – The service will select an available host and assign it to an user.
     * `Direct` Assignment – Admin selects a specific host to assign to an user.
     */
    personalDesktopAssignmentType?: pulumi.Input<string>;
    /**
     * Option to specify the preferred Application Group type for the Virtual Desktop Host Pool.
     * Valid options are `None`, `Desktop` or `RailApplications`. Default is `None`.
     */
    preferredAppGroupType?: pulumi.Input<string>;
    /**
     * This block is now non-functional and will be removed in version 3.0 of the Azure Provider - use the
     * `azurerm_virtual_desktop_host_pool_registration_info` resource instead.
     */
    registrationInfo?: pulumi.Input<inputs.desktopvirtualization.HostPoolRegistrationInfo>;
    /**
     * The name of the resource group in which to
     * create the Virtual Desktop Host Pool. Changing the resource group name forces
     * a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Enables or disables the Start VM on Connection Feature. Defaults to `false`.
     */
    startVmOnConnect?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Virtual Desktop Host Pool. Valid options are
     * `Personal` or `Pooled`. Changing the type forces a new resource to be created.
     */
    type: pulumi.Input<string>;
    /**
     * Allows you to test service changes before they are deployed to production. Defaults to `false`.
     */
    validateEnvironment?: pulumi.Input<boolean>;
}
