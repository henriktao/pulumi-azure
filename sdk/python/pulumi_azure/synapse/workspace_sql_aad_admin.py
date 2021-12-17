# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WorkspaceSqlAadAdminInitArgs', 'WorkspaceSqlAadAdmin']

@pulumi.input_type
class WorkspaceSqlAadAdminInitArgs:
    def __init__(__self__, *,
                 login: pulumi.Input[str],
                 object_id: pulumi.Input[str],
                 synapse_workspace_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a WorkspaceSqlAadAdmin resource.
        :param pulumi.Input[str] login: The login name of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] object_id: The object id of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        :param pulumi.Input[str] tenant_id: The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        pulumi.set(__self__, "login", login)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "synapse_workspace_id", synapse_workspace_id)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def login(self) -> pulumi.Input[str]:
        """
        The login name of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: pulumi.Input[str]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[str]:
        """
        The object id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="synapseWorkspaceId")
    def synapse_workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        """
        return pulumi.get(self, "synapse_workspace_id")

    @synapse_workspace_id.setter
    def synapse_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "synapse_workspace_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _WorkspaceSqlAadAdminState:
    def __init__(__self__, *,
                 login: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 synapse_workspace_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WorkspaceSqlAadAdmin resources.
        :param pulumi.Input[str] login: The login name of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] object_id: The object id of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        :param pulumi.Input[str] tenant_id: The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        if login is not None:
            pulumi.set(__self__, "login", login)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if synapse_workspace_id is not None:
            pulumi.set(__self__, "synapse_workspace_id", synapse_workspace_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[str]]:
        """
        The login name of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="synapseWorkspaceId")
    def synapse_workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        """
        return pulumi.get(self, "synapse_workspace_id")

    @synapse_workspace_id.setter
    def synapse_workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "synapse_workspace_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class WorkspaceSqlAadAdmin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 synapse_workspace_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Azure Active Directory Sql Administrator setting for a Synapse Workspace

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="StorageV2",
            is_hns_enabled=True)
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        current = azure.core.get_client_config()
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=azurerm_resource_group["exampl"]["location"],
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            purge_protection_enabled=True)
        deployer = azure.keyvault.AccessPolicy("deployer",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=current.object_id,
            key_permissions=[
                "create",
                "get",
                "delete",
                "purge",
            ])
        example_key = azure.keyvault.Key("exampleKey",
            key_vault_id=example_key_vault.id,
            key_type="RSA",
            key_size=2048,
            key_opts=[
                "unwrapKey",
                "wrapKey",
            ],
            opts=pulumi.ResourceOptions(depends_on=[deployer]))
        example_workspace = azure.synapse.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!",
            tags={
                "Env": "production",
            })
        test = azure.synapse.WorkspaceSqlAadAdmin("test",
            synapse_workspace_id=azurerm_synapse_workspace["test"]["id"],
            login="AzureAD Admin",
            object_id=current.object_id,
            tenant_id=current.tenant_id)
        ```

        ## Import

        Synapse Workspace Azure AD Administrator can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:synapse/workspaceSqlAadAdmin:WorkspaceSqlAadAdmin example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlAdministrators/activeDirectory
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] login: The login name of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] object_id: The object id of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        :param pulumi.Input[str] tenant_id: The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkspaceSqlAadAdminInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure Active Directory Sql Administrator setting for a Synapse Workspace

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="StorageV2",
            is_hns_enabled=True)
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        current = azure.core.get_client_config()
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=azurerm_resource_group["exampl"]["location"],
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            purge_protection_enabled=True)
        deployer = azure.keyvault.AccessPolicy("deployer",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=current.object_id,
            key_permissions=[
                "create",
                "get",
                "delete",
                "purge",
            ])
        example_key = azure.keyvault.Key("exampleKey",
            key_vault_id=example_key_vault.id,
            key_type="RSA",
            key_size=2048,
            key_opts=[
                "unwrapKey",
                "wrapKey",
            ],
            opts=pulumi.ResourceOptions(depends_on=[deployer]))
        example_workspace = azure.synapse.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!",
            tags={
                "Env": "production",
            })
        test = azure.synapse.WorkspaceSqlAadAdmin("test",
            synapse_workspace_id=azurerm_synapse_workspace["test"]["id"],
            login="AzureAD Admin",
            object_id=current.object_id,
            tenant_id=current.tenant_id)
        ```

        ## Import

        Synapse Workspace Azure AD Administrator can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:synapse/workspaceSqlAadAdmin:WorkspaceSqlAadAdmin example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlAdministrators/activeDirectory
        ```

        :param str resource_name: The name of the resource.
        :param WorkspaceSqlAadAdminInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkspaceSqlAadAdminInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 synapse_workspace_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkspaceSqlAadAdminInitArgs.__new__(WorkspaceSqlAadAdminInitArgs)

            if login is None and not opts.urn:
                raise TypeError("Missing required property 'login'")
            __props__.__dict__["login"] = login
            if object_id is None and not opts.urn:
                raise TypeError("Missing required property 'object_id'")
            __props__.__dict__["object_id"] = object_id
            if synapse_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'synapse_workspace_id'")
            __props__.__dict__["synapse_workspace_id"] = synapse_workspace_id
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
        super(WorkspaceSqlAadAdmin, __self__).__init__(
            'azure:synapse/workspaceSqlAadAdmin:WorkspaceSqlAadAdmin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            login: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            synapse_workspace_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'WorkspaceSqlAadAdmin':
        """
        Get an existing WorkspaceSqlAadAdmin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] login: The login name of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] object_id: The object id of the Azure AD Administrator of this Synapse Workspace.
        :param pulumi.Input[str] synapse_workspace_id: The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        :param pulumi.Input[str] tenant_id: The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkspaceSqlAadAdminState.__new__(_WorkspaceSqlAadAdminState)

        __props__.__dict__["login"] = login
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["synapse_workspace_id"] = synapse_workspace_id
        __props__.__dict__["tenant_id"] = tenant_id
        return WorkspaceSqlAadAdmin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[str]:
        """
        The login name of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The object id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="synapseWorkspaceId")
    def synapse_workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Synapse Workspace where the Azure AD Administrator should be configured.
        """
        return pulumi.get(self, "synapse_workspace_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The tenant id of the Azure AD Administrator of this Synapse Workspace.
        """
        return pulumi.get(self, "tenant_id")

