// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Synapse Workspace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     isHnsEnabled: "true",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleAccount.id});
 * const exampleWorkspace = new azure.synapse.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     storageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.id,
 *     sqlAdministratorLogin: "sqladminuser",
 *     sqlAdministratorLoginPassword: "H@Sh1CoR3!",
 *     aadAdmin: {
 *         login: "AzureAD Admin",
 *         objectId: "00000000-0000-0000-0000-000000000000",
 *         tenantId: "00000000-0000-0000-0000-000000000000",
 *     },
 *     tags: {
 *         Env: "production",
 *     },
 * });
 * ```
 * ### Creating A Workspace With Customer Managed Key And Azure AD Admin
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 *     accountKind: "StorageV2",
 *     isHnsEnabled: "true",
 * });
 * const exampleDataLakeGen2Filesystem = new azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", {storageAccountId: exampleAccount.id});
 * const exampleKeyVault = new azure.keyvault.KeyVault("exampleKeyVault", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     tenantId: current.then(current => current.tenantId),
 *     skuName: "standard",
 *     purgeProtectionEnabled: true,
 * });
 * const deployer = new azure.keyvault.AccessPolicy("deployer", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: current.then(current => current.tenantId),
 *     objectId: current.then(current => current.objectId),
 *     keyPermissions: [
 *         "create",
 *         "get",
 *         "delete",
 *         "purge",
 *     ],
 * });
 * const exampleKey = new azure.keyvault.Key("exampleKey", {
 *     keyVaultId: exampleKeyVault.id,
 *     keyType: "RSA",
 *     keySize: 2048,
 *     keyOpts: [
 *         "unwrapKey",
 *         "wrapKey",
 *     ],
 * }, {
 *     dependsOn: [deployer],
 * });
 * const exampleWorkspace = new azure.synapse.Workspace("exampleWorkspace", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     storageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.id,
 *     sqlAdministratorLogin: "sqladminuser",
 *     sqlAdministratorLoginPassword: "H@Sh1CoR3!",
 *     customerManagedKey: {
 *         keyVersionlessId: exampleKey.versionlessId,
 *         keyName: "enckey",
 *     },
 *     tags: {
 *         Env: "production",
 *     },
 * });
 * const workspacePolicy = new azure.keyvault.AccessPolicy("workspacePolicy", {
 *     keyVaultId: exampleKeyVault.id,
 *     tenantId: exampleWorkspace.identities.apply(identities => identities[0].tenantId),
 *     objectId: exampleWorkspace.identities.apply(identities => identities[0].principalId),
 *     keyPermissions: [
 *         "Get",
 *         "WrapKey",
 *         "UnwrapKey",
 *     ],
 * });
 * const exampleWorkspaceKey = new azure.synapse.WorkspaceKey("exampleWorkspaceKey", {
 *     customerManagedKeyVersionlessId: exampleKey.versionlessId,
 *     synapseWorkspaceId: exampleWorkspace.id,
 *     active: true,
 *     customerManagedKeyName: "enckey",
 * }, {
 *     dependsOn: [workspacePolicy],
 * });
 * const exampleWorkspaceAadAdmin = new azure.synapse.WorkspaceAadAdmin("exampleWorkspaceAadAdmin", {
 *     synapseWorkspaceId: exampleWorkspace.id,
 *     login: "AzureAD Admin",
 *     objectId: "00000000-0000-0000-0000-000000000000",
 *     tenantId: "00000000-0000-0000-0000-000000000000",
 * }, {
 *     dependsOn: [exampleWorkspaceKey],
 * });
 * ```
 *
 * ## Import
 *
 * Synapse Workspace can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:synapse/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
 * ```
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkspaceState, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:synapse/workspace:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * An `aadAdmin` block as defined below. Conflicts with `customerManagedKey`.
     */
    public readonly aadAdmin!: pulumi.Output<outputs.synapse.WorkspaceAadAdmin>;
    /**
     * An `azureDevopsRepo` block as defined below.
     */
    public readonly azureDevopsRepo!: pulumi.Output<outputs.synapse.WorkspaceAzureDevopsRepo | undefined>;
    /**
     * Subnet ID used for computes in workspace
     */
    public readonly computeSubnetId!: pulumi.Output<string | undefined>;
    /**
     * A list of Connectivity endpoints for this Synapse Workspace.
     */
    public /*out*/ readonly connectivityEndpoints!: pulumi.Output<{[key: string]: string}>;
    /**
     * A `customerManagedKey` block as defined below. Conflicts with `aadAdmin`.
     */
    public readonly customerManagedKey!: pulumi.Output<outputs.synapse.WorkspaceCustomerManagedKey | undefined>;
    /**
     * Is data exfiltration protection enabled in this workspace? If set to `true`, `managedVirtualNetworkEnabled` must also be set to `true`. Changing this forces a new resource to be created.
     */
    public readonly dataExfiltrationProtectionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A `githubRepo` block as defined below.
     */
    public readonly githubRepo!: pulumi.Output<outputs.synapse.WorkspaceGithubRepo | undefined>;
    /**
     * An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
     */
    public /*out*/ readonly identities!: pulumi.Output<outputs.synapse.WorkspaceIdentity[]>;
    /**
     * Allowed Aad Tenant Ids For Linking.
     */
    public readonly linkingAllowedForAadTenantIds!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Workspace managed resource group.
     */
    public readonly managedResourceGroupName!: pulumi.Output<string>;
    /**
     * Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
     */
    public readonly managedVirtualNetworkEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
     */
    public readonly publicNetworkAccessEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of purview account.
     */
    public readonly purviewId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * An `sqlAadAdmin` block as defined below.
     */
    public readonly sqlAadAdmin!: pulumi.Output<outputs.synapse.WorkspaceSqlAadAdmin>;
    /**
     * Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
     */
    public readonly sqlAdministratorLogin!: pulumi.Output<string>;
    /**
     * The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
     */
    public readonly sqlAdministratorLoginPassword!: pulumi.Output<string>;
    /**
     * Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
     */
    public readonly sqlIdentityControlEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
     */
    public readonly storageDataLakeGen2FilesystemId!: pulumi.Output<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Workspace.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceArgs | WorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkspaceState | undefined;
            inputs["aadAdmin"] = state ? state.aadAdmin : undefined;
            inputs["azureDevopsRepo"] = state ? state.azureDevopsRepo : undefined;
            inputs["computeSubnetId"] = state ? state.computeSubnetId : undefined;
            inputs["connectivityEndpoints"] = state ? state.connectivityEndpoints : undefined;
            inputs["customerManagedKey"] = state ? state.customerManagedKey : undefined;
            inputs["dataExfiltrationProtectionEnabled"] = state ? state.dataExfiltrationProtectionEnabled : undefined;
            inputs["githubRepo"] = state ? state.githubRepo : undefined;
            inputs["identities"] = state ? state.identities : undefined;
            inputs["linkingAllowedForAadTenantIds"] = state ? state.linkingAllowedForAadTenantIds : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managedResourceGroupName"] = state ? state.managedResourceGroupName : undefined;
            inputs["managedVirtualNetworkEnabled"] = state ? state.managedVirtualNetworkEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publicNetworkAccessEnabled"] = state ? state.publicNetworkAccessEnabled : undefined;
            inputs["purviewId"] = state ? state.purviewId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sqlAadAdmin"] = state ? state.sqlAadAdmin : undefined;
            inputs["sqlAdministratorLogin"] = state ? state.sqlAdministratorLogin : undefined;
            inputs["sqlAdministratorLoginPassword"] = state ? state.sqlAdministratorLoginPassword : undefined;
            inputs["sqlIdentityControlEnabled"] = state ? state.sqlIdentityControlEnabled : undefined;
            inputs["storageDataLakeGen2FilesystemId"] = state ? state.storageDataLakeGen2FilesystemId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as WorkspaceArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sqlAdministratorLogin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlAdministratorLogin'");
            }
            if ((!args || args.sqlAdministratorLoginPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlAdministratorLoginPassword'");
            }
            if ((!args || args.storageDataLakeGen2FilesystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageDataLakeGen2FilesystemId'");
            }
            inputs["aadAdmin"] = args ? args.aadAdmin : undefined;
            inputs["azureDevopsRepo"] = args ? args.azureDevopsRepo : undefined;
            inputs["computeSubnetId"] = args ? args.computeSubnetId : undefined;
            inputs["customerManagedKey"] = args ? args.customerManagedKey : undefined;
            inputs["dataExfiltrationProtectionEnabled"] = args ? args.dataExfiltrationProtectionEnabled : undefined;
            inputs["githubRepo"] = args ? args.githubRepo : undefined;
            inputs["linkingAllowedForAadTenantIds"] = args ? args.linkingAllowedForAadTenantIds : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedResourceGroupName"] = args ? args.managedResourceGroupName : undefined;
            inputs["managedVirtualNetworkEnabled"] = args ? args.managedVirtualNetworkEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicNetworkAccessEnabled"] = args ? args.publicNetworkAccessEnabled : undefined;
            inputs["purviewId"] = args ? args.purviewId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlAadAdmin"] = args ? args.sqlAadAdmin : undefined;
            inputs["sqlAdministratorLogin"] = args ? args.sqlAdministratorLogin : undefined;
            inputs["sqlAdministratorLoginPassword"] = args ? args.sqlAdministratorLoginPassword : undefined;
            inputs["sqlIdentityControlEnabled"] = args ? args.sqlIdentityControlEnabled : undefined;
            inputs["storageDataLakeGen2FilesystemId"] = args ? args.storageDataLakeGen2FilesystemId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["connectivityEndpoints"] = undefined /*out*/;
            inputs["identities"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workspace resources.
 */
export interface WorkspaceState {
    /**
     * An `aadAdmin` block as defined below. Conflicts with `customerManagedKey`.
     */
    aadAdmin?: pulumi.Input<inputs.synapse.WorkspaceAadAdmin>;
    /**
     * An `azureDevopsRepo` block as defined below.
     */
    azureDevopsRepo?: pulumi.Input<inputs.synapse.WorkspaceAzureDevopsRepo>;
    /**
     * Subnet ID used for computes in workspace
     */
    computeSubnetId?: pulumi.Input<string>;
    /**
     * A list of Connectivity endpoints for this Synapse Workspace.
     */
    connectivityEndpoints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `customerManagedKey` block as defined below. Conflicts with `aadAdmin`.
     */
    customerManagedKey?: pulumi.Input<inputs.synapse.WorkspaceCustomerManagedKey>;
    /**
     * Is data exfiltration protection enabled in this workspace? If set to `true`, `managedVirtualNetworkEnabled` must also be set to `true`. Changing this forces a new resource to be created.
     */
    dataExfiltrationProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * A `githubRepo` block as defined below.
     */
    githubRepo?: pulumi.Input<inputs.synapse.WorkspaceGithubRepo>;
    /**
     * An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
     */
    identities?: pulumi.Input<pulumi.Input<inputs.synapse.WorkspaceIdentity>[]>;
    /**
     * Allowed Aad Tenant Ids For Linking.
     */
    linkingAllowedForAadTenantIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Workspace managed resource group.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
     */
    managedVirtualNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of purview account.
     */
    purviewId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * An `sqlAadAdmin` block as defined below.
     */
    sqlAadAdmin?: pulumi.Input<inputs.synapse.WorkspaceSqlAadAdmin>;
    /**
     * Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
     */
    sqlAdministratorLogin?: pulumi.Input<string>;
    /**
     * The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
     */
    sqlAdministratorLoginPassword?: pulumi.Input<string>;
    /**
     * Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
     */
    sqlIdentityControlEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
     */
    storageDataLakeGen2FilesystemId?: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Workspace.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * An `aadAdmin` block as defined below. Conflicts with `customerManagedKey`.
     */
    aadAdmin?: pulumi.Input<inputs.synapse.WorkspaceAadAdmin>;
    /**
     * An `azureDevopsRepo` block as defined below.
     */
    azureDevopsRepo?: pulumi.Input<inputs.synapse.WorkspaceAzureDevopsRepo>;
    /**
     * Subnet ID used for computes in workspace
     */
    computeSubnetId?: pulumi.Input<string>;
    /**
     * A `customerManagedKey` block as defined below. Conflicts with `aadAdmin`.
     */
    customerManagedKey?: pulumi.Input<inputs.synapse.WorkspaceCustomerManagedKey>;
    /**
     * Is data exfiltration protection enabled in this workspace? If set to `true`, `managedVirtualNetworkEnabled` must also be set to `true`. Changing this forces a new resource to be created.
     */
    dataExfiltrationProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * A `githubRepo` block as defined below.
     */
    githubRepo?: pulumi.Input<inputs.synapse.WorkspaceGithubRepo>;
    /**
     * Allowed Aad Tenant Ids For Linking.
     */
    linkingAllowedForAadTenantIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Workspace managed resource group.
     */
    managedResourceGroupName?: pulumi.Input<string>;
    /**
     * Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
     */
    managedVirtualNetworkEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether public network access is allowed for the Cognitive Account. Defaults to `true`.
     */
    publicNetworkAccessEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of purview account.
     */
    purviewId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * An `sqlAadAdmin` block as defined below.
     */
    sqlAadAdmin?: pulumi.Input<inputs.synapse.WorkspaceSqlAadAdmin>;
    /**
     * Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
     */
    sqlAdministratorLogin: pulumi.Input<string>;
    /**
     * The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
     */
    sqlAdministratorLoginPassword: pulumi.Input<string>;
    /**
     * Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
     */
    sqlIdentityControlEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
     */
    storageDataLakeGen2FilesystemId: pulumi.Input<string>;
    /**
     * A mapping of tags which should be assigned to the Synapse Workspace.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
