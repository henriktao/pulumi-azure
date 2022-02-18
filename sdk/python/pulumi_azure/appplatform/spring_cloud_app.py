# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SpringCloudAppArgs', 'SpringCloudApp']

@pulumi.input_type
class SpringCloudAppArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 custom_persistent_disks: Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input['SpringCloudAppIdentityArgs']] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistent_disk: Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SpringCloudApp resource.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]] custom_persistent_disks: A `custom_persistent_disk` block as defined below.
        :param pulumi.Input[bool] https_only: Is only https allowed? Defaults to `false`.
        :param pulumi.Input['SpringCloudAppIdentityArgs'] identity: An `identity` block as defined below.
        :param pulumi.Input[bool] is_public: Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input['SpringCloudAppPersistentDiskArgs'] persistent_disk: An `persistent_disk` block as defined below.
        :param pulumi.Input[bool] tls_enabled: Is End to End TLS Enabled? Defaults to `false`.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "service_name", service_name)
        if custom_persistent_disks is not None:
            pulumi.set(__self__, "custom_persistent_disks", custom_persistent_disks)
        if https_only is not None:
            pulumi.set(__self__, "https_only", https_only)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistent_disk is not None:
            pulumi.set(__self__, "persistent_disk", persistent_disk)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="customPersistentDisks")
    def custom_persistent_disks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]]:
        """
        A `custom_persistent_disk` block as defined below.
        """
        return pulumi.get(self, "custom_persistent_disks")

    @custom_persistent_disks.setter
    def custom_persistent_disks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]]):
        pulumi.set(self, "custom_persistent_disks", value)

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Is only https allowed? Defaults to `false`.
        """
        return pulumi.get(self, "https_only")

    @https_only.setter
    def https_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_only", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['SpringCloudAppIdentityArgs']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['SpringCloudAppIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="persistentDisk")
    def persistent_disk(self) -> Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']]:
        """
        An `persistent_disk` block as defined below.
        """
        return pulumi.get(self, "persistent_disk")

    @persistent_disk.setter
    def persistent_disk(self, value: Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']]):
        pulumi.set(self, "persistent_disk", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is End to End TLS Enabled? Defaults to `false`.
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)


@pulumi.input_type
class _SpringCloudAppState:
    def __init__(__self__, *,
                 custom_persistent_disks: Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input['SpringCloudAppIdentityArgs']] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistent_disk: Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SpringCloudApp resources.
        :param pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]] custom_persistent_disks: A `custom_persistent_disk` block as defined below.
        :param pulumi.Input[str] fqdn: The Fully Qualified DNS Name of the Spring Application in the service.
        :param pulumi.Input[bool] https_only: Is only https allowed? Defaults to `false`.
        :param pulumi.Input['SpringCloudAppIdentityArgs'] identity: An `identity` block as defined below.
        :param pulumi.Input[bool] is_public: Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input['SpringCloudAppPersistentDiskArgs'] persistent_disk: An `persistent_disk` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] tls_enabled: Is End to End TLS Enabled? Defaults to `false`.
        :param pulumi.Input[str] url: The public endpoint of the Spring Cloud Application.
        """
        if custom_persistent_disks is not None:
            pulumi.set(__self__, "custom_persistent_disks", custom_persistent_disks)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if https_only is not None:
            pulumi.set(__self__, "https_only", https_only)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistent_disk is not None:
            pulumi.set(__self__, "persistent_disk", persistent_disk)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="customPersistentDisks")
    def custom_persistent_disks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]]:
        """
        A `custom_persistent_disk` block as defined below.
        """
        return pulumi.get(self, "custom_persistent_disks")

    @custom_persistent_disks.setter
    def custom_persistent_disks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SpringCloudAppCustomPersistentDiskArgs']]]]):
        pulumi.set(self, "custom_persistent_disks", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        The Fully Qualified DNS Name of the Spring Application in the service.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Is only https allowed? Defaults to `false`.
        """
        return pulumi.get(self, "https_only")

    @https_only.setter
    def https_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_only", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['SpringCloudAppIdentityArgs']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['SpringCloudAppIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="persistentDisk")
    def persistent_disk(self) -> Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']]:
        """
        An `persistent_disk` block as defined below.
        """
        return pulumi.get(self, "persistent_disk")

    @persistent_disk.setter
    def persistent_disk(self, value: Optional[pulumi.Input['SpringCloudAppPersistentDiskArgs']]):
        pulumi.set(self, "persistent_disk", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is End to End TLS Enabled? Defaults to `false`.
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The public endpoint of the Spring Cloud Application.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class SpringCloudApp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_persistent_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpringCloudAppCustomPersistentDiskArgs']]]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppIdentityArgs']]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistent_disk: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppPersistentDiskArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manage an Azure Spring Cloud Application.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_spring_cloud_service = azure.appplatform.SpringCloudService("exampleSpringCloudService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_spring_cloud_app = azure.appplatform.SpringCloudApp("exampleSpringCloudApp",
            resource_group_name=example_resource_group.name,
            service_name=example_spring_cloud_service.name,
            identity=azure.appplatform.SpringCloudAppIdentityArgs(
                type="SystemAssigned",
            ))
        ```

        ## Import

        Spring Cloud Application can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appplatform/springCloudApp:SpringCloudApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpringCloudAppCustomPersistentDiskArgs']]]] custom_persistent_disks: A `custom_persistent_disk` block as defined below.
        :param pulumi.Input[bool] https_only: Is only https allowed? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['SpringCloudAppIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[bool] is_public: Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['SpringCloudAppPersistentDiskArgs']] persistent_disk: An `persistent_disk` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] tls_enabled: Is End to End TLS Enabled? Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SpringCloudAppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage an Azure Spring Cloud Application.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_spring_cloud_service = azure.appplatform.SpringCloudService("exampleSpringCloudService",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location)
        example_spring_cloud_app = azure.appplatform.SpringCloudApp("exampleSpringCloudApp",
            resource_group_name=example_resource_group.name,
            service_name=example_spring_cloud_service.name,
            identity=azure.appplatform.SpringCloudAppIdentityArgs(
                type="SystemAssigned",
            ))
        ```

        ## Import

        Spring Cloud Application can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:appplatform/springCloudApp:SpringCloudApp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp
        ```

        :param str resource_name: The name of the resource.
        :param SpringCloudAppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SpringCloudAppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_persistent_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpringCloudAppCustomPersistentDiskArgs']]]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppIdentityArgs']]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistent_disk: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppPersistentDiskArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
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
            __props__ = SpringCloudAppArgs.__new__(SpringCloudAppArgs)

            __props__.__dict__["custom_persistent_disks"] = custom_persistent_disks
            __props__.__dict__["https_only"] = https_only
            __props__.__dict__["identity"] = identity
            __props__.__dict__["is_public"] = is_public
            __props__.__dict__["name"] = name
            __props__.__dict__["persistent_disk"] = persistent_disk
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["tls_enabled"] = tls_enabled
            __props__.__dict__["fqdn"] = None
            __props__.__dict__["url"] = None
        super(SpringCloudApp, __self__).__init__(
            'azure:appplatform/springCloudApp:SpringCloudApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_persistent_disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpringCloudAppCustomPersistentDiskArgs']]]]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            https_only: Optional[pulumi.Input[bool]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppIdentityArgs']]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            persistent_disk: Optional[pulumi.Input[pulumi.InputType['SpringCloudAppPersistentDiskArgs']]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            tls_enabled: Optional[pulumi.Input[bool]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'SpringCloudApp':
        """
        Get an existing SpringCloudApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpringCloudAppCustomPersistentDiskArgs']]]] custom_persistent_disks: A `custom_persistent_disk` block as defined below.
        :param pulumi.Input[str] fqdn: The Fully Qualified DNS Name of the Spring Application in the service.
        :param pulumi.Input[bool] https_only: Is only https allowed? Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['SpringCloudAppIdentityArgs']] identity: An `identity` block as defined below.
        :param pulumi.Input[bool] is_public: Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['SpringCloudAppPersistentDiskArgs']] persistent_disk: An `persistent_disk` block as defined below.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_name: Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] tls_enabled: Is End to End TLS Enabled? Defaults to `false`.
        :param pulumi.Input[str] url: The public endpoint of the Spring Cloud Application.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SpringCloudAppState.__new__(_SpringCloudAppState)

        __props__.__dict__["custom_persistent_disks"] = custom_persistent_disks
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["https_only"] = https_only
        __props__.__dict__["identity"] = identity
        __props__.__dict__["is_public"] = is_public
        __props__.__dict__["name"] = name
        __props__.__dict__["persistent_disk"] = persistent_disk
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["tls_enabled"] = tls_enabled
        __props__.__dict__["url"] = url
        return SpringCloudApp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customPersistentDisks")
    def custom_persistent_disks(self) -> pulumi.Output[Optional[Sequence['outputs.SpringCloudAppCustomPersistentDisk']]]:
        """
        A `custom_persistent_disk` block as defined below.
        """
        return pulumi.get(self, "custom_persistent_disks")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The Fully Qualified DNS Name of the Spring Application in the service.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> pulumi.Output[Optional[bool]]:
        """
        Is only https allowed? Defaults to `false`.
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.SpringCloudAppIdentity']]:
        """
        An `identity` block as defined below.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[Optional[bool]]:
        """
        Does the Spring Cloud Application have public endpoint? Defaults to `false`.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="persistentDisk")
    def persistent_disk(self) -> pulumi.Output['outputs.SpringCloudAppPersistentDisk']:
        """
        An `persistent_disk` block as defined below.
        """
        return pulumi.get(self, "persistent_disk")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is End to End TLS Enabled? Defaults to `false`.
        """
        return pulumi.get(self, "tls_enabled")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The public endpoint of the Spring Cloud Application.
        """
        return pulumi.get(self, "url")

