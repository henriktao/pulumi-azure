# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CertificateBindingArgs', 'CertificateBinding']

@pulumi.input_type
class CertificateBindingArgs:
    def __init__(__self__, *,
                 certificate_id: pulumi.Input[str],
                 hostname_binding_id: pulumi.Input[str],
                 ssl_state: pulumi.Input[str]):
        """
        The set of arguments for constructing a CertificateBinding resource.
        :param pulumi.Input[str] certificate_id: The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] hostname_binding_id: The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] ssl_state: The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        """
        pulumi.set(__self__, "certificate_id", certificate_id)
        pulumi.set(__self__, "hostname_binding_id", hostname_binding_id)
        pulumi.set(__self__, "ssl_state", ssl_state)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Input[str]:
        """
        The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="hostnameBindingId")
    def hostname_binding_id(self) -> pulumi.Input[str]:
        """
        The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "hostname_binding_id")

    @hostname_binding_id.setter
    def hostname_binding_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname_binding_id", value)

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> pulumi.Input[str]:
        """
        The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "ssl_state")

    @ssl_state.setter
    def ssl_state(self, value: pulumi.Input[str]):
        pulumi.set(self, "ssl_state", value)


@pulumi.input_type
class _CertificateBindingState:
    def __init__(__self__, *,
                 app_service_name: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 hostname_binding_id: Optional[pulumi.Input[str]] = None,
                 ssl_state: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CertificateBinding resources.
        :param pulumi.Input[str] app_service_name: The name of the App Service to which the certificate was bound.
        :param pulumi.Input[str] certificate_id: The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] hostname: The hostname of the bound certificate.
        :param pulumi.Input[str] hostname_binding_id: The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] ssl_state: The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] thumbprint: The certificate thumbprint.
        """
        if app_service_name is not None:
            pulumi.set(__self__, "app_service_name", app_service_name)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if hostname_binding_id is not None:
            pulumi.set(__self__, "hostname_binding_id", hostname_binding_id)
        if ssl_state is not None:
            pulumi.set(__self__, "ssl_state", ssl_state)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)

    @property
    @pulumi.getter(name="appServiceName")
    def app_service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the App Service to which the certificate was bound.
        """
        return pulumi.get(self, "app_service_name")

    @app_service_name.setter
    def app_service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service_name", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname of the bound certificate.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="hostnameBindingId")
    def hostname_binding_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "hostname_binding_id")

    @hostname_binding_id.setter
    def hostname_binding_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname_binding_id", value)

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> Optional[pulumi.Input[str]]:
        """
        The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "ssl_state")

    @ssl_state.setter
    def ssl_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_state", value)

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate thumbprint.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint", value)


class CertificateBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 hostname_binding_id: Optional[pulumi.Input[str]] = None,
                 ssl_state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an App Service Certificate Binding.

        ## Import

        App Service Certificate Bindings can be imported using the `hostname_binding_id` and the `app_service_certificate_id` , e.g.

        ```sh
         $ pulumi import azure:appservice/certificateBinding:CertificateBinding example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/mywebsite.com"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] hostname_binding_id: The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] ssl_state: The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an App Service Certificate Binding.

        ## Import

        App Service Certificate Bindings can be imported using the `hostname_binding_id` and the `app_service_certificate_id` , e.g.

        ```sh
         $ pulumi import azure:appservice/certificateBinding:CertificateBinding example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/mywebsite.com"
        ```

        :param str resource_name: The name of the resource.
        :param CertificateBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 hostname_binding_id: Optional[pulumi.Input[str]] = None,
                 ssl_state: Optional[pulumi.Input[str]] = None,
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
            __props__ = CertificateBindingArgs.__new__(CertificateBindingArgs)

            if certificate_id is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_id'")
            __props__.__dict__["certificate_id"] = certificate_id
            if hostname_binding_id is None and not opts.urn:
                raise TypeError("Missing required property 'hostname_binding_id'")
            __props__.__dict__["hostname_binding_id"] = hostname_binding_id
            if ssl_state is None and not opts.urn:
                raise TypeError("Missing required property 'ssl_state'")
            __props__.__dict__["ssl_state"] = ssl_state
            __props__.__dict__["app_service_name"] = None
            __props__.__dict__["hostname"] = None
            __props__.__dict__["thumbprint"] = None
        super(CertificateBinding, __self__).__init__(
            'azure:appservice/certificateBinding:CertificateBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service_name: Optional[pulumi.Input[str]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            hostname_binding_id: Optional[pulumi.Input[str]] = None,
            ssl_state: Optional[pulumi.Input[str]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None) -> 'CertificateBinding':
        """
        Get an existing CertificateBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: The name of the App Service to which the certificate was bound.
        :param pulumi.Input[str] certificate_id: The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] hostname: The hostname of the bound certificate.
        :param pulumi.Input[str] hostname_binding_id: The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] ssl_state: The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        :param pulumi.Input[str] thumbprint: The certificate thumbprint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateBindingState.__new__(_CertificateBindingState)

        __props__.__dict__["app_service_name"] = app_service_name
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["hostname_binding_id"] = hostname_binding_id
        __props__.__dict__["ssl_state"] = ssl_state
        __props__.__dict__["thumbprint"] = thumbprint
        return CertificateBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServiceName")
    def app_service_name(self) -> pulumi.Output[str]:
        """
        The name of the App Service to which the certificate was bound.
        """
        return pulumi.get(self, "app_service_name")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[str]:
        """
        The ID of the certificate to bind to the custom domain. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The hostname of the bound certificate.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="hostnameBindingId")
    def hostname_binding_id(self) -> pulumi.Output[str]:
        """
        The ID of the Custom Domain/Hostname Binding. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "hostname_binding_id")

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> pulumi.Output[str]:
        """
        The type of certificate binding. Allowed values are `IpBasedEnabled` or `SniEnabled`. Changing this forces a new App Service Certificate Binding to be created.
        """
        return pulumi.get(self, "ssl_state")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[str]:
        """
        The certificate thumbprint.
        """
        return pulumi.get(self, "thumbprint")

