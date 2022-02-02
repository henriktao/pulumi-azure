// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Management Lock which is scoped to a Subscription, Resource Group or Resource.
 *
 * ## Example Usage
 * ### Subscription Level Lock)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getSubscription({});
 * const subscription_level = new azure.management.Lock("subscription-level", {
 *     scope: current.then(current => current.id),
 *     lockLevel: "CanNotDelete",
 *     notes: "Items can't be deleted in this subscription!",
 * });
 * ```
 *
 * ## Example Usage (Resource Group Level Lock)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.core.ResourceGroup("example", {location: "West Europe"});
 * const resource_group_level = new azure.management.Lock("resource-group-level", {
 *     scope: example.id,
 *     lockLevel: "ReadOnly",
 *     notes: "This Resource Group is Read-Only",
 * });
 * ```
 * ### Resource Level Lock)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const examplePublicIp = new azure.network.PublicIp("examplePublicIp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     allocationMethod: "Static",
 *     idleTimeoutInMinutes: 30,
 * });
 * const public_ip = new azure.management.Lock("public-ip", {
 *     scope: examplePublicIp.id,
 *     lockLevel: "CanNotDelete",
 *     notes: "Locked because it's needed by a third-party",
 * });
 * ```
 *
 * ## Import
 *
 * Management Locks can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:management/lock:Lock lock1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Authorization/locks/lock1
 * ```
 */
export class Lock extends pulumi.CustomResource {
    /**
     * Get an existing Lock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LockState, opts?: pulumi.CustomResourceOptions): Lock {
        return new Lock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:management/lock:Lock';

    /**
     * Returns true if the given object is an instance of Lock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lock.__pulumiType;
    }

    /**
     * Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
     */
    public readonly lockLevel!: pulumi.Output<string>;
    /**
     * Specifies the name of the Management Lock. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
     */
    public readonly scope!: pulumi.Output<string>;

    /**
     * Create a Lock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LockArgs | LockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LockState | undefined;
            resourceInputs["lockLevel"] = state ? state.lockLevel : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as LockArgs | undefined;
            if ((!args || args.lockLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lockLevel'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["lockLevel"] = args ? args.lockLevel : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure:managementresource/manangementLock:ManangementLock" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Lock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lock resources.
 */
export interface LockState {
    /**
     * Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
     */
    lockLevel?: pulumi.Input<string>;
    /**
     * Specifies the name of the Management Lock. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
     */
    notes?: pulumi.Input<string>;
    /**
     * Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
     */
    scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Lock resource.
 */
export interface LockArgs {
    /**
     * Specifies the Level to be used for this Lock. Possible values are `CanNotDelete` and `ReadOnly`. Changing this forces a new resource to be created.
     */
    lockLevel: pulumi.Input<string>;
    /**
     * Specifies the name of the Management Lock. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies some notes about the lock. Maximum of 512 characters. Changing this forces a new resource to be created.
     */
    notes?: pulumi.Input<string>;
    /**
     * Specifies the scope at which the Management Lock should be created. Changing this forces a new resource to be created.
     */
    scope: pulumi.Input<string>;
}
