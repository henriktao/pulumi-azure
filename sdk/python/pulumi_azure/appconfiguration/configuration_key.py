# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigurationKeyArgs', 'ConfigurationKey']

@pulumi.input_type
class ConfigurationKeyArgs:
    def __init__(__self__, *,
                 configuration_store_id: pulumi.Input[str],
                 key: pulumi.Input[str],
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vault_key_reference: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConfigurationKey resource.
        :param pulumi.Input[str] configuration_store_id: Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] key: The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_type: The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] etag: The ETag of the key.
        :param pulumi.Input[str] label: The label of the App Configuration Key.  Changing this forces a new resource to be created.
        :param pulumi.Input[bool] locked: Should this App Configuration Key be Locked to prevent changes?
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        :param pulumi.Input[str] value: The value of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] vault_key_reference: The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        pulumi.set(__self__, "configuration_store_id", configuration_store_id)
        pulumi.set(__self__, "key", key)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if vault_key_reference is not None:
            pulumi.set(__self__, "vault_key_reference", vault_key_reference)

    @property
    @pulumi.getter(name="configurationStoreId")
    def configuration_store_id(self) -> pulumi.Input[str]:
        """
        Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "configuration_store_id")

    @configuration_store_id.setter
    def configuration_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "configuration_store_id", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        The ETag of the key.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The label of the App Configuration Key.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this App Configuration Key be Locked to prevent changes?
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="vaultKeyReference")
    def vault_key_reference(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        return pulumi.get(self, "vault_key_reference")

    @vault_key_reference.setter
    def vault_key_reference(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_key_reference", value)


@pulumi.input_type
class _ConfigurationKeyState:
    def __init__(__self__, *,
                 configuration_store_id: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vault_key_reference: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigurationKey resources.
        :param pulumi.Input[str] configuration_store_id: Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_type: The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] etag: The ETag of the key.
        :param pulumi.Input[str] key: The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] label: The label of the App Configuration Key.  Changing this forces a new resource to be created.
        :param pulumi.Input[bool] locked: Should this App Configuration Key be Locked to prevent changes?
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        :param pulumi.Input[str] value: The value of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] vault_key_reference: The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        if configuration_store_id is not None:
            pulumi.set(__self__, "configuration_store_id", configuration_store_id)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if vault_key_reference is not None:
            pulumi.set(__self__, "vault_key_reference", vault_key_reference)

    @property
    @pulumi.getter(name="configurationStoreId")
    def configuration_store_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "configuration_store_id")

    @configuration_store_id.setter
    def configuration_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_store_id", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        The ETag of the key.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The label of the App Configuration Key.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this App Configuration Key be Locked to prevent changes?
        """
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="vaultKeyReference")
    def vault_key_reference(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        return pulumi.get(self, "vault_key_reference")

    @vault_key_reference.setter
    def vault_key_reference(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_key_reference", value)


class ConfigurationKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_store_id: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vault_key_reference: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Azure App Configuration Key.

        ## Example Usage
        ### `Kv` Type

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        appconf = azure.appconfiguration.ConfigurationStore("appconf",
            resource_group_name=rg.name,
            location=rg.location)
        test = azure.appconfiguration.ConfigurationKey("test",
            configuration_store_id=appconf.id,
            key="appConfKey1",
            label="somelabel",
            value="a test")
        ```
        ### `Vault` Type
        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        appconf = azure.appconfiguration.ConfigurationStore("appconf",
            resource_group_name=rg.name,
            location=rg.location)
        current = azure.core.get_client_config()
        kv = azure.keyvault.KeyVault("kv",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            tenant_id=current.tenant_id,
            sku_name="premium",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                key_permissions=[
                    "create",
                    "get",
                ],
                secret_permissions=[
                    "set",
                    "get",
                    "delete",
                    "purge",
                    "recover",
                ],
            )])
        kvs = azure.keyvault.Secret("kvs",
            value="szechuan",
            key_vault_id=kv.id)
        test = azure.appconfiguration.ConfigurationKey("test",
            configuration_store_id=azurerm_app_configuration["test"]["id"],
            key="key1",
            type="vault",
            label="label1",
            vault_key_reference=kvs.id)
        ```

        ## Import

        App Configuration Keys can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/label1
        ```

         If you wish to import a key with an empty label then sustitute the label's name with `%00`, like this

        ```sh
         $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/%00
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration_store_id: Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_type: The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] etag: The ETag of the key.
        :param pulumi.Input[str] key: The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] label: The label of the App Configuration Key.  Changing this forces a new resource to be created.
        :param pulumi.Input[bool] locked: Should this App Configuration Key be Locked to prevent changes?
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        :param pulumi.Input[str] value: The value of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] vault_key_reference: The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Azure App Configuration Key.

        ## Example Usage
        ### `Kv` Type

        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        appconf = azure.appconfiguration.ConfigurationStore("appconf",
            resource_group_name=rg.name,
            location=rg.location)
        test = azure.appconfiguration.ConfigurationKey("test",
            configuration_store_id=appconf.id,
            key="appConfKey1",
            label="somelabel",
            value="a test")
        ```
        ### `Vault` Type
        ```python
        import pulumi
        import pulumi_azure as azure

        rg = azure.core.ResourceGroup("rg", location="West Europe")
        appconf = azure.appconfiguration.ConfigurationStore("appconf",
            resource_group_name=rg.name,
            location=rg.location)
        current = azure.core.get_client_config()
        kv = azure.keyvault.KeyVault("kv",
            location=azurerm_resource_group["test"]["location"],
            resource_group_name=azurerm_resource_group["test"]["name"],
            tenant_id=current.tenant_id,
            sku_name="premium",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                key_permissions=[
                    "create",
                    "get",
                ],
                secret_permissions=[
                    "set",
                    "get",
                    "delete",
                    "purge",
                    "recover",
                ],
            )])
        kvs = azure.keyvault.Secret("kvs",
            value="szechuan",
            key_vault_id=kv.id)
        test = azure.appconfiguration.ConfigurationKey("test",
            configuration_store_id=azurerm_app_configuration["test"]["id"],
            key="key1",
            type="vault",
            label="label1",
            vault_key_reference=kvs.id)
        ```

        ## Import

        App Configuration Keys can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/label1
        ```

         If you wish to import a key with an empty label then sustitute the label's name with `%00`, like this

        ```sh
         $ pulumi import azure:appconfiguration/configurationKey:ConfigurationKey test /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/%00
        ```

        :param str resource_name: The name of the resource.
        :param ConfigurationKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_store_id: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vault_key_reference: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConfigurationKeyArgs.__new__(ConfigurationKeyArgs)

            if configuration_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'configuration_store_id'")
            __props__.__dict__["configuration_store_id"] = configuration_store_id
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["etag"] = etag
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["label"] = label
            __props__.__dict__["locked"] = locked
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["value"] = value
            __props__.__dict__["vault_key_reference"] = vault_key_reference
        super(ConfigurationKey, __self__).__init__(
            'azure:appconfiguration/configurationKey:ConfigurationKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration_store_id: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            locked: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            vault_key_reference: Optional[pulumi.Input[str]] = None) -> 'ConfigurationKey':
        """
        Get an existing ConfigurationKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration_store_id: Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] content_type: The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] etag: The ETag of the key.
        :param pulumi.Input[str] key: The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] label: The label of the App Configuration Key.  Changing this forces a new resource to be created.
        :param pulumi.Input[bool] locked: Should this App Configuration Key be Locked to prevent changes?
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] type: The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        :param pulumi.Input[str] value: The value of the App Configuration Key. This should only be set when type is set to `kv`.
        :param pulumi.Input[str] vault_key_reference: The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationKeyState.__new__(_ConfigurationKeyState)

        __props__.__dict__["configuration_store_id"] = configuration_store_id
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["etag"] = etag
        __props__.__dict__["key"] = key
        __props__.__dict__["label"] = label
        __props__.__dict__["locked"] = locked
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        __props__.__dict__["vault_key_reference"] = vault_key_reference
        return ConfigurationKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configurationStoreId")
    def configuration_store_id(self) -> pulumi.Output[str]:
        """
        Specifies the id of the App Configuration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "configuration_store_id")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        """
        The content type of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The ETag of the key.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The name of the App Configuration Key to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[Optional[str]]:
        """
        The label of the App Configuration Key.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def locked(self) -> pulumi.Output[Optional[bool]]:
        """
        Should this App Configuration Key be Locked to prevent changes?
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the App Configuration Key. It can either be `kv` (simple [key/value](https://docs.microsoft.com/en-us/azure/azure-app-configuration/concept-key-value)) or `vault` (where the value is a reference to a [Key Vault Secret](https://azure.microsoft.com/en-gb/services/key-vault/).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the App Configuration Key. This should only be set when type is set to `kv`.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="vaultKeyReference")
    def vault_key_reference(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the vault secret this App Configuration Key refers to, when `type` is set to `vault`.
        """
        return pulumi.get(self, "vault_key_reference")

