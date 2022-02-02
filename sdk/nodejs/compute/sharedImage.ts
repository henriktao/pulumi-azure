// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Shared Image within a Shared Image Gallery.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleSharedImageGallery = new azure.compute.SharedImageGallery("exampleSharedImageGallery", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     description: "Shared images and things.",
 *     tags: {
 *         Hello: "There",
 *         World: "Example",
 *     },
 * });
 * const exampleSharedImage = new azure.compute.SharedImage("exampleSharedImage", {
 *     galleryName: exampleSharedImageGallery.name,
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     osType: "Linux",
 *     identifier: {
 *         publisher: "PublisherName",
 *         offer: "OfferName",
 *         sku: "ExampleSku",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Shared Images can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:compute/sharedImage:SharedImage image1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
 * ```
 */
export class SharedImage extends pulumi.CustomResource {
    /**
     * Get an existing SharedImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedImageState, opts?: pulumi.CustomResourceOptions): SharedImage {
        return new SharedImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/sharedImage:SharedImage';

    /**
     * Returns true if the given object is an instance of SharedImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SharedImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SharedImage.__pulumiType;
    }

    /**
     * A description of this Shared Image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    public readonly eula!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    public readonly galleryName!: pulumi.Output<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * An `identifier` block as defined below.
     */
    public readonly identifier!: pulumi.Output<outputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    public readonly privacyStatementUri!: pulumi.Output<string | undefined>;
    /**
     * A `purchasePlan` block as defined below.
     */
    public readonly purchasePlan!: pulumi.Output<outputs.compute.SharedImagePurchasePlan | undefined>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    public readonly releaseNoteUri!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly specialized!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly trustedLaunchEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SharedImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedImageArgs | SharedImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SharedImageState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["eula"] = state ? state.eula : undefined;
            resourceInputs["galleryName"] = state ? state.galleryName : undefined;
            resourceInputs["hyperVGeneration"] = state ? state.hyperVGeneration : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["privacyStatementUri"] = state ? state.privacyStatementUri : undefined;
            resourceInputs["purchasePlan"] = state ? state.purchasePlan : undefined;
            resourceInputs["releaseNoteUri"] = state ? state.releaseNoteUri : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["specialized"] = state ? state.specialized : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["trustedLaunchEnabled"] = state ? state.trustedLaunchEnabled : undefined;
        } else {
            const args = argsOrState as SharedImageArgs | undefined;
            if ((!args || args.galleryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'galleryName'");
            }
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.osType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'osType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eula"] = args ? args.eula : undefined;
            resourceInputs["galleryName"] = args ? args.galleryName : undefined;
            resourceInputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osType"] = args ? args.osType : undefined;
            resourceInputs["privacyStatementUri"] = args ? args.privacyStatementUri : undefined;
            resourceInputs["purchasePlan"] = args ? args.purchasePlan : undefined;
            resourceInputs["releaseNoteUri"] = args ? args.releaseNoteUri : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["specialized"] = args ? args.specialized : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trustedLaunchEnabled"] = args ? args.trustedLaunchEnabled : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SharedImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedImage resources.
 */
export interface SharedImageState {
    /**
     * A description of this Shared Image.
     */
    description?: pulumi.Input<string>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    eula?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    galleryName?: pulumi.Input<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    hyperVGeneration?: pulumi.Input<string>;
    /**
     * An `identifier` block as defined below.
     */
    identifier?: pulumi.Input<inputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    osType?: pulumi.Input<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    privacyStatementUri?: pulumi.Input<string>;
    /**
     * A `purchasePlan` block as defined below.
     */
    purchasePlan?: pulumi.Input<inputs.compute.SharedImagePurchasePlan>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
     */
    specialized?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Defaults to `false`. Changing this forces a new resource to be created.
     */
    trustedLaunchEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SharedImage resource.
 */
export interface SharedImageArgs {
    /**
     * A description of this Shared Image.
     */
    description?: pulumi.Input<string>;
    /**
     * The End User Licence Agreement for the Shared Image.
     */
    eula?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
     */
    galleryName: pulumi.Input<string>;
    /**
     * The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are `V1` and `V2`. Defaults to `V1`. Changing this forces a new resource to be created.
     */
    hyperVGeneration?: pulumi.Input<string>;
    /**
     * An `identifier` block as defined below.
     */
    identifier: pulumi.Input<inputs.compute.SharedImageIdentifier>;
    /**
     * Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Shared Image. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`. Changing this forces a new resource to be created.
     */
    osType: pulumi.Input<string>;
    /**
     * The URI containing the Privacy Statement associated with this Shared Image.
     */
    privacyStatementUri?: pulumi.Input<string>;
    /**
     * A `purchasePlan` block as defined below.
     */
    purchasePlan?: pulumi.Input<inputs.compute.SharedImagePurchasePlan>;
    /**
     * The URI containing the Release Notes associated with this Shared Image.
     */
    releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies that the Operating System used inside this Image has not been Generalized (for example, `sysprep` on Windows has not been run). Defaults to `false`. Changing this forces a new resource to be created.
     */
    specialized?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the Shared Image.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Defaults to `false`. Changing this forces a new resource to be created.
     */
    trustedLaunchEnabled?: pulumi.Input<boolean>;
}
