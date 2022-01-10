# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CacheBlobTargetArgs', 'CacheBlobTarget']

@pulumi.input_type
class CacheBlobTargetArgs:
    def __init__(__self__, *,
                 cache_name: pulumi.Input[str],
                 namespace_path: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 storage_container_id: pulumi.Input[str],
                 access_policy_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CacheBlobTarget resource.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_path: The client-facing file path of the HPC Cache Blob Target.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_id: The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] access_policy_name: The name of the access policy applied to this target. Defaults to `default`.
        :param pulumi.Input[str] name: The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "cache_name", cache_name)
        pulumi.set(__self__, "namespace_path", namespace_path)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "storage_container_id", storage_container_id)
        if access_policy_name is not None:
            pulumi.set(__self__, "access_policy_name", access_policy_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="cacheName")
    def cache_name(self) -> pulumi.Input[str]:
        """
        The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cache_name")

    @cache_name.setter
    def cache_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cache_name", value)

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Input[str]:
        """
        The client-facing file path of the HPC Cache Blob Target.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_path", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="storageContainerId")
    def storage_container_id(self) -> pulumi.Input[str]:
        """
        The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_container_id")

    @storage_container_id.setter
    def storage_container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_container_id", value)

    @property
    @pulumi.getter(name="accessPolicyName")
    def access_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the access policy applied to this target. Defaults to `default`.
        """
        return pulumi.get(self, "access_policy_name")

    @access_policy_name.setter
    def access_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_policy_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CacheBlobTargetState:
    def __init__(__self__, *,
                 access_policy_name: Optional[pulumi.Input[str]] = None,
                 cache_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_container_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CacheBlobTarget resources.
        :param pulumi.Input[str] access_policy_name: The name of the access policy applied to this target. Defaults to `default`.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_path: The client-facing file path of the HPC Cache Blob Target.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_id: The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        if access_policy_name is not None:
            pulumi.set(__self__, "access_policy_name", access_policy_name)
        if cache_name is not None:
            pulumi.set(__self__, "cache_name", cache_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if storage_container_id is not None:
            pulumi.set(__self__, "storage_container_id", storage_container_id)

    @property
    @pulumi.getter(name="accessPolicyName")
    def access_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the access policy applied to this target. Defaults to `default`.
        """
        return pulumi.get(self, "access_policy_name")

    @access_policy_name.setter
    def access_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_policy_name", value)

    @property
    @pulumi.getter(name="cacheName")
    def cache_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cache_name")

    @cache_name.setter
    def cache_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[pulumi.Input[str]]:
        """
        The client-facing file path of the HPC Cache Blob Target.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_path", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="storageContainerId")
    def storage_container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_container_id")

    @storage_container_id.setter
    def storage_container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_container_id", value)


class CacheBlobTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_name: Optional[pulumi.Input[str]] = None,
                 cache_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_container_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Blob Target within a HPC Cache.

        > **NOTE:**: By request of the service team the provider no longer automatically registering the `Microsoft.StorageCache` Resource Provider for this resource. To register it you can run `az provider register --namespace 'Microsoft.StorageCache'`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        example_cache = azure.hpc.Cache("exampleCache",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            cache_size_in_gb=3072,
            subnet_id=example_subnet.id,
            sku_name="Standard_2G")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer", storage_account_name=example_account.name)
        example_service_principal = azuread.get_service_principal(display_name="HPC Cache Resource Provider")
        example_storage_account_contrib = azure.authorization.Assignment("exampleStorageAccountContrib",
            scope=example_account.id,
            role_definition_name="Storage Account Contributor",
            principal_id=example_service_principal.object_id)
        example_storage_blob_data_contrib = azure.authorization.Assignment("exampleStorageBlobDataContrib",
            scope=example_account.id,
            role_definition_name="Storage Blob Data Contributor",
            principal_id=example_service_principal.object_id)
        example_cache_blob_target = azure.hpc.CacheBlobTarget("exampleCacheBlobTarget",
            resource_group_name=example_resource_group.name,
            cache_name=example_cache.name,
            storage_container_id=example_container.resource_manager_id,
            namespace_path="/blob_storage")
        ```

        ## Import

        Blob Targets within an HPC Cache can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:hpc/cacheBlobTarget:CacheBlobTarget example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy_name: The name of the access policy applied to this target. Defaults to `default`.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_path: The client-facing file path of the HPC Cache Blob Target.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_id: The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CacheBlobTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Blob Target within a HPC Cache.

        > **NOTE:**: By request of the service team the provider no longer automatically registering the `Microsoft.StorageCache` Resource Provider for this resource. To register it you can run `az provider register --namespace 'Microsoft.StorageCache'`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_azuread as azuread

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            address_spaces=["10.0.0.0/16"],
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/24"])
        example_cache = azure.hpc.Cache("exampleCache",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            cache_size_in_gb=3072,
            subnet_id=example_subnet.id,
            sku_name="Standard_2G")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_container = azure.storage.Container("exampleContainer", storage_account_name=example_account.name)
        example_service_principal = azuread.get_service_principal(display_name="HPC Cache Resource Provider")
        example_storage_account_contrib = azure.authorization.Assignment("exampleStorageAccountContrib",
            scope=example_account.id,
            role_definition_name="Storage Account Contributor",
            principal_id=example_service_principal.object_id)
        example_storage_blob_data_contrib = azure.authorization.Assignment("exampleStorageBlobDataContrib",
            scope=example_account.id,
            role_definition_name="Storage Blob Data Contributor",
            principal_id=example_service_principal.object_id)
        example_cache_blob_target = azure.hpc.CacheBlobTarget("exampleCacheBlobTarget",
            resource_group_name=example_resource_group.name,
            cache_name=example_cache.name,
            storage_container_id=example_container.resource_manager_id,
            namespace_path="/blob_storage")
        ```

        ## Import

        Blob Targets within an HPC Cache can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:hpc/cacheBlobTarget:CacheBlobTarget example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageCache/caches/cache1/storageTargets/target1
        ```

        :param str resource_name: The name of the resource.
        :param CacheBlobTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CacheBlobTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_name: Optional[pulumi.Input[str]] = None,
                 cache_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_path: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_container_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CacheBlobTargetArgs.__new__(CacheBlobTargetArgs)

            __props__.__dict__["access_policy_name"] = access_policy_name
            if cache_name is None and not opts.urn:
                raise TypeError("Missing required property 'cache_name'")
            __props__.__dict__["cache_name"] = cache_name
            __props__.__dict__["name"] = name
            if namespace_path is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_path'")
            __props__.__dict__["namespace_path"] = namespace_path
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if storage_container_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_container_id'")
            __props__.__dict__["storage_container_id"] = storage_container_id
        super(CacheBlobTarget, __self__).__init__(
            'azure:hpc/cacheBlobTarget:CacheBlobTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_policy_name: Optional[pulumi.Input[str]] = None,
            cache_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_path: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            storage_container_id: Optional[pulumi.Input[str]] = None) -> 'CacheBlobTarget':
        """
        Get an existing CacheBlobTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy_name: The name of the access policy applied to this target. Defaults to `default`.
        :param pulumi.Input[str] cache_name: The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_path: The client-facing file path of the HPC Cache Blob Target.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_container_id: The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CacheBlobTargetState.__new__(_CacheBlobTargetState)

        __props__.__dict__["access_policy_name"] = access_policy_name
        __props__.__dict__["cache_name"] = cache_name
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["storage_container_id"] = storage_container_id
        return CacheBlobTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPolicyName")
    def access_policy_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the access policy applied to this target. Defaults to `default`.
        """
        return pulumi.get(self, "access_policy_name")

    @property
    @pulumi.getter(name="cacheName")
    def cache_name(self) -> pulumi.Output[str]:
        """
        The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cache_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Output[str]:
        """
        The client-facing file path of the HPC Cache Blob Target.
        """
        return pulumi.get(self, "namespace_path")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="storageContainerId")
    def storage_container_id(self) -> pulumi.Output[str]:
        """
        The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_container_id")

