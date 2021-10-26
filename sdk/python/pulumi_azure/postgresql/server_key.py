# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServerKeyArgs', 'ServerKey']

@pulumi.input_type
class ServerKeyArgs:
    def __init__(__self__, *,
                 key_vault_key_id: pulumi.Input[str],
                 server_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ServerKey resource.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "key_vault_key_id", key_vault_key_id)
        pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> pulumi.Input[str]:
        """
        The URL to a Key Vault Key.
        """
        return pulumi.get(self, "key_vault_key_id")

    @key_vault_key_id.setter
    def key_vault_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_vault_key_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)


@pulumi.input_type
class _ServerKeyState:
    def __init__(__self__, *,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerKey resources.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        if key_vault_key_id is not None:
            pulumi.set(__self__, "key_vault_key_id", key_vault_key_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to a Key Vault Key.
        """
        return pulumi.get(self, "key_vault_key_id")

    @key_vault_key_id.setter
    def key_vault_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_key_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)


class ServerKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Customer Managed Key for a PostgreSQL Server.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            purge_protection_enabled=True)
        example_server = azure.postgresql.Server("exampleServer",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            administrator_login="psqladmin",
            administrator_login_password="H@Sh1CoR3!",
            sku_name="GP_Gen5_2",
            version="11",
            storage_mb=51200,
            ssl_enforcement_enabled=True,
            identity=azure.postgresql.ServerIdentityArgs(
                type="SystemAssigned",
            ))
        server = azure.keyvault.AccessPolicy("server",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=example_server.identity.principal_id,
            key_permissions=[
                "get",
                "unwrapkey",
                "wrapkey",
            ],
            secret_permissions=["get"])
        client = azure.keyvault.AccessPolicy("client",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=current.object_id,
            key_permissions=[
                "get",
                "create",
                "delete",
                "list",
                "restore",
                "recover",
                "unwrapkey",
                "wrapkey",
                "purge",
                "encrypt",
                "decrypt",
                "sign",
                "verify",
            ],
            secret_permissions=["get"])
        example_key = azure.keyvault.Key("exampleKey",
            key_vault_id=example_key_vault.id,
            key_type="RSA",
            key_size=2048,
            key_opts=[
                "decrypt",
                "encrypt",
                "sign",
                "unwrapKey",
                "verify",
                "wrapKey",
            ],
            opts=pulumi.ResourceOptions(depends_on=[
                    client,
                    server,
                ]))
        example_server_key = azure.postgresql.ServerKey("exampleServerKey",
            server_id=example_server.id,
            key_vault_key_id=example_key.id)
        ```

        ## Import

        A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.

        ```sh
         $ pulumi import azure:postgresql/serverKey:ServerKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Customer Managed Key for a PostgreSQL Server.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            purge_protection_enabled=True)
        example_server = azure.postgresql.Server("exampleServer",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            administrator_login="psqladmin",
            administrator_login_password="H@Sh1CoR3!",
            sku_name="GP_Gen5_2",
            version="11",
            storage_mb=51200,
            ssl_enforcement_enabled=True,
            identity=azure.postgresql.ServerIdentityArgs(
                type="SystemAssigned",
            ))
        server = azure.keyvault.AccessPolicy("server",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=example_server.identity.principal_id,
            key_permissions=[
                "get",
                "unwrapkey",
                "wrapkey",
            ],
            secret_permissions=["get"])
        client = azure.keyvault.AccessPolicy("client",
            key_vault_id=example_key_vault.id,
            tenant_id=current.tenant_id,
            object_id=current.object_id,
            key_permissions=[
                "get",
                "create",
                "delete",
                "list",
                "restore",
                "recover",
                "unwrapkey",
                "wrapkey",
                "purge",
                "encrypt",
                "decrypt",
                "sign",
                "verify",
            ],
            secret_permissions=["get"])
        example_key = azure.keyvault.Key("exampleKey",
            key_vault_id=example_key_vault.id,
            key_type="RSA",
            key_size=2048,
            key_opts=[
                "decrypt",
                "encrypt",
                "sign",
                "unwrapKey",
                "verify",
                "wrapKey",
            ],
            opts=pulumi.ResourceOptions(depends_on=[
                    client,
                    server,
                ]))
        example_server_key = azure.postgresql.ServerKey("exampleServerKey",
            server_id=example_server.id,
            key_vault_key_id=example_key.id)
        ```

        ## Import

        A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.

        ```sh
         $ pulumi import azure:postgresql/serverKey:ServerKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
        ```

        :param str resource_name: The name of the resource.
        :param ServerKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServerKeyArgs.__new__(ServerKeyArgs)

            if key_vault_key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_vault_key_id'")
            __props__.__dict__["key_vault_key_id"] = key_vault_key_id
            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__.__dict__["server_id"] = server_id
        super(ServerKey, __self__).__init__(
            'azure:postgresql/serverKey:ServerKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_vault_key_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None) -> 'ServerKey':
        """
        Get an existing ServerKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerKeyState.__new__(_ServerKeyState)

        __props__.__dict__["key_vault_key_id"] = key_vault_key_id
        __props__.__dict__["server_id"] = server_id
        return ServerKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> pulumi.Output[str]:
        """
        The URL to a Key Vault Key.
        """
        return pulumi.get(self, "key_vault_key_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_id")

