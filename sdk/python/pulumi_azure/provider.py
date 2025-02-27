# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 auxiliary_tenant_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_correlation_request_id: Optional[pulumi.Input[bool]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 features: Optional[pulumi.Input['ProviderFeaturesArgs']] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_provider_registration: Optional[pulumi.Input[bool]] = None,
                 storage_use_azuread: Optional[pulumi.Input[bool]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_msal: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] client_certificate_password: The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
               Certificate
        :param pulumi.Input[str] client_certificate_path: The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
               Principal using a Client Certificate.
        :param pulumi.Input[str] client_id: The Client ID which should be used.
        :param pulumi.Input[str] client_secret: The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        :param pulumi.Input[bool] disable_correlation_request_id: This will disable the x-ms-correlation-request-id header.
        :param pulumi.Input[bool] disable_terraform_partner_id: This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        :param pulumi.Input[str] environment: The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
        :param pulumi.Input[str] metadata_host: The Hostname which should be used for the Azure Metadata Service.
        :param pulumi.Input[str] metadata_url: Deprecated - replaced by `metadata_host`.
        :param pulumi.Input[str] msi_endpoint: The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
               automatically.
        :param pulumi.Input[str] partner_id: A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        :param pulumi.Input[bool] skip_credentials_validation: [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        :param pulumi.Input[bool] skip_provider_registration: Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
               registered?
        :param pulumi.Input[bool] storage_use_azuread: Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        :param pulumi.Input[str] subscription_id: The Subscription ID which should be used.
        :param pulumi.Input[str] tenant_id: The Tenant ID which should be used.
        :param pulumi.Input[bool] use_msal: Should Terraform obtain MSAL auth tokens and no longer use Azure Active Directory Graph?
        :param pulumi.Input[bool] use_msi: Allowed Managed Service Identity be used for Authentication.
        """
        if auxiliary_tenant_ids is not None:
            pulumi.set(__self__, "auxiliary_tenant_ids", auxiliary_tenant_ids)
        if client_certificate_password is not None:
            pulumi.set(__self__, "client_certificate_password", client_certificate_password)
        if client_certificate_path is not None:
            pulumi.set(__self__, "client_certificate_path", client_certificate_path)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if disable_correlation_request_id is not None:
            pulumi.set(__self__, "disable_correlation_request_id", disable_correlation_request_id)
        if disable_terraform_partner_id is not None:
            pulumi.set(__self__, "disable_terraform_partner_id", disable_terraform_partner_id)
        if environment is None:
            environment = (_utilities.get_env('AZURE_ENVIRONMENT', 'ARM_ENVIRONMENT') or 'public')
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if metadata_host is None:
            metadata_host = _utilities.get_env('ARM_METADATA_HOSTNAME')
        if metadata_host is not None:
            pulumi.set(__self__, "metadata_host", metadata_host)
        if metadata_url is None:
            metadata_url = _utilities.get_env('ARM_METADATA_URL')
        if metadata_url is not None:
            warnings.warn("""use `metadata_host` instead""", DeprecationWarning)
            pulumi.log.warn("""metadata_url is deprecated: use `metadata_host` instead""")
        if metadata_url is not None:
            pulumi.set(__self__, "metadata_url", metadata_url)
        if msi_endpoint is not None:
            pulumi.set(__self__, "msi_endpoint", msi_endpoint)
        if partner_id is not None:
            pulumi.set(__self__, "partner_id", partner_id)
        if skip_credentials_validation is not None:
            warnings.warn("""This field is deprecated and will be removed in version 3.0 of the Azure Provider""", DeprecationWarning)
            pulumi.log.warn("""skip_credentials_validation is deprecated: This field is deprecated and will be removed in version 3.0 of the Azure Provider""")
        if skip_credentials_validation is not None:
            pulumi.set(__self__, "skip_credentials_validation", skip_credentials_validation)
        if skip_provider_registration is None:
            skip_provider_registration = (_utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
        if skip_provider_registration is not None:
            pulumi.set(__self__, "skip_provider_registration", skip_provider_registration)
        if storage_use_azuread is None:
            storage_use_azuread = (_utilities.get_env_bool('ARM_STORAGE_USE_AZUREAD') or False)
        if storage_use_azuread is not None:
            pulumi.set(__self__, "storage_use_azuread", storage_use_azuread)
        if subscription_id is None:
            subscription_id = (_utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if use_msal is not None:
            pulumi.set(__self__, "use_msal", use_msal)
        if use_msi is not None:
            pulumi.set(__self__, "use_msi", use_msi)

    @property
    @pulumi.getter(name="auxiliaryTenantIds")
    def auxiliary_tenant_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "auxiliary_tenant_ids")

    @auxiliary_tenant_ids.setter
    def auxiliary_tenant_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "auxiliary_tenant_ids", value)

    @property
    @pulumi.getter(name="clientCertificatePassword")
    def client_certificate_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return pulumi.get(self, "client_certificate_password")

    @client_certificate_password.setter
    def client_certificate_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate_password", value)

    @property
    @pulumi.getter(name="clientCertificatePath")
    def client_certificate_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return pulumi.get(self, "client_certificate_path")

    @client_certificate_path.setter
    def client_certificate_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate_path", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Client ID which should be used.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="disableCorrelationRequestId")
    def disable_correlation_request_id(self) -> Optional[pulumi.Input[bool]]:
        """
        This will disable the x-ms-correlation-request-id header.
        """
        return pulumi.get(self, "disable_correlation_request_id")

    @disable_correlation_request_id.setter
    def disable_correlation_request_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_correlation_request_id", value)

    @property
    @pulumi.getter(name="disableTerraformPartnerId")
    def disable_terraform_partner_id(self) -> Optional[pulumi.Input[bool]]:
        """
        This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        """
        return pulumi.get(self, "disable_terraform_partner_id")

    @disable_terraform_partner_id.setter
    def disable_terraform_partner_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_terraform_partner_id", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def features(self) -> Optional[pulumi.Input['ProviderFeaturesArgs']]:
        return pulumi.get(self, "features")

    @features.setter
    def features(self, value: Optional[pulumi.Input['ProviderFeaturesArgs']]):
        pulumi.set(self, "features", value)

    @property
    @pulumi.getter(name="metadataHost")
    def metadata_host(self) -> Optional[pulumi.Input[str]]:
        """
        The Hostname which should be used for the Azure Metadata Service.
        """
        return pulumi.get(self, "metadata_host")

    @metadata_host.setter
    def metadata_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_host", value)

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated - replaced by `metadata_host`.
        """
        return pulumi.get(self, "metadata_url")

    @metadata_url.setter
    def metadata_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_url", value)

    @property
    @pulumi.getter(name="msiEndpoint")
    def msi_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
        automatically.
        """
        return pulumi.get(self, "msi_endpoint")

    @msi_endpoint.setter
    def msi_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "msi_endpoint", value)

    @property
    @pulumi.getter(name="partnerId")
    def partner_id(self) -> Optional[pulumi.Input[str]]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return pulumi.get(self, "partner_id")

    @partner_id.setter
    def partner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partner_id", value)

    @property
    @pulumi.getter(name="skipCredentialsValidation")
    def skip_credentials_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        """
        return pulumi.get(self, "skip_credentials_validation")

    @skip_credentials_validation.setter
    def skip_credentials_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_credentials_validation", value)

    @property
    @pulumi.getter(name="skipProviderRegistration")
    def skip_provider_registration(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
        registered?
        """
        return pulumi.get(self, "skip_provider_registration")

    @skip_provider_registration.setter
    def skip_provider_registration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_provider_registration", value)

    @property
    @pulumi.getter(name="storageUseAzuread")
    def storage_use_azuread(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        """
        return pulumi.get(self, "storage_use_azuread")

    @storage_use_azuread.setter
    def storage_use_azuread(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "storage_use_azuread", value)

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Subscription ID which should be used.
        """
        return pulumi.get(self, "subscription_id")

    @subscription_id.setter
    def subscription_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID which should be used.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="useMsal")
    def use_msal(self) -> Optional[pulumi.Input[bool]]:
        """
        Should Terraform obtain MSAL auth tokens and no longer use Azure Active Directory Graph?
        """
        return pulumi.get(self, "use_msal")

    @use_msal.setter
    def use_msal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_msal", value)

    @property
    @pulumi.getter(name="useMsi")
    def use_msi(self) -> Optional[pulumi.Input[bool]]:
        """
        Allowed Managed Service Identity be used for Authentication.
        """
        return pulumi.get(self, "use_msi")

    @use_msi.setter
    def use_msi(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_msi", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auxiliary_tenant_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_correlation_request_id: Optional[pulumi.Input[bool]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 features: Optional[pulumi.Input[pulumi.InputType['ProviderFeaturesArgs']]] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_provider_registration: Optional[pulumi.Input[bool]] = None,
                 storage_use_azuread: Optional[pulumi.Input[bool]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_msal: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the azurerm package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_certificate_password: The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
               Certificate
        :param pulumi.Input[str] client_certificate_path: The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
               Principal using a Client Certificate.
        :param pulumi.Input[str] client_id: The Client ID which should be used.
        :param pulumi.Input[str] client_secret: The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        :param pulumi.Input[bool] disable_correlation_request_id: This will disable the x-ms-correlation-request-id header.
        :param pulumi.Input[bool] disable_terraform_partner_id: This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        :param pulumi.Input[str] environment: The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
        :param pulumi.Input[str] metadata_host: The Hostname which should be used for the Azure Metadata Service.
        :param pulumi.Input[str] metadata_url: Deprecated - replaced by `metadata_host`.
        :param pulumi.Input[str] msi_endpoint: The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
               automatically.
        :param pulumi.Input[str] partner_id: A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        :param pulumi.Input[bool] skip_credentials_validation: [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
        :param pulumi.Input[bool] skip_provider_registration: Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
               registered?
        :param pulumi.Input[bool] storage_use_azuread: Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
        :param pulumi.Input[str] subscription_id: The Subscription ID which should be used.
        :param pulumi.Input[str] tenant_id: The Tenant ID which should be used.
        :param pulumi.Input[bool] use_msal: Should Terraform obtain MSAL auth tokens and no longer use Azure Active Directory Graph?
        :param pulumi.Input[bool] use_msi: Allowed Managed Service Identity be used for Authentication.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the azurerm package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auxiliary_tenant_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_correlation_request_id: Optional[pulumi.Input[bool]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 features: Optional[pulumi.Input[pulumi.InputType['ProviderFeaturesArgs']]] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_provider_registration: Optional[pulumi.Input[bool]] = None,
                 storage_use_azuread: Optional[pulumi.Input[bool]] = None,
                 subscription_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_msal: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["auxiliary_tenant_ids"] = pulumi.Output.from_input(auxiliary_tenant_ids).apply(pulumi.runtime.to_json) if auxiliary_tenant_ids is not None else None
            __props__.__dict__["client_certificate_password"] = client_certificate_password
            __props__.__dict__["client_certificate_path"] = client_certificate_path
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = client_secret
            __props__.__dict__["disable_correlation_request_id"] = pulumi.Output.from_input(disable_correlation_request_id).apply(pulumi.runtime.to_json) if disable_correlation_request_id is not None else None
            __props__.__dict__["disable_terraform_partner_id"] = pulumi.Output.from_input(disable_terraform_partner_id).apply(pulumi.runtime.to_json) if disable_terraform_partner_id is not None else None
            if environment is None:
                environment = (_utilities.get_env('AZURE_ENVIRONMENT', 'ARM_ENVIRONMENT') or 'public')
            __props__.__dict__["environment"] = environment
            __props__.__dict__["features"] = pulumi.Output.from_input(features).apply(pulumi.runtime.to_json) if features is not None else None
            if metadata_host is None:
                metadata_host = _utilities.get_env('ARM_METADATA_HOSTNAME')
            __props__.__dict__["metadata_host"] = metadata_host
            if metadata_url is None:
                metadata_url = _utilities.get_env('ARM_METADATA_URL')
            if metadata_url is not None and not opts.urn:
                warnings.warn("""use `metadata_host` instead""", DeprecationWarning)
                pulumi.log.warn("""metadata_url is deprecated: use `metadata_host` instead""")
            __props__.__dict__["metadata_url"] = metadata_url
            __props__.__dict__["msi_endpoint"] = msi_endpoint
            __props__.__dict__["partner_id"] = partner_id
            if skip_credentials_validation is not None and not opts.urn:
                warnings.warn("""This field is deprecated and will be removed in version 3.0 of the Azure Provider""", DeprecationWarning)
                pulumi.log.warn("""skip_credentials_validation is deprecated: This field is deprecated and will be removed in version 3.0 of the Azure Provider""")
            __props__.__dict__["skip_credentials_validation"] = pulumi.Output.from_input(skip_credentials_validation).apply(pulumi.runtime.to_json) if skip_credentials_validation is not None else None
            if skip_provider_registration is None:
                skip_provider_registration = (_utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
            __props__.__dict__["skip_provider_registration"] = pulumi.Output.from_input(skip_provider_registration).apply(pulumi.runtime.to_json) if skip_provider_registration is not None else None
            if storage_use_azuread is None:
                storage_use_azuread = (_utilities.get_env_bool('ARM_STORAGE_USE_AZUREAD') or False)
            __props__.__dict__["storage_use_azuread"] = pulumi.Output.from_input(storage_use_azuread).apply(pulumi.runtime.to_json) if storage_use_azuread is not None else None
            if subscription_id is None:
                subscription_id = (_utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
            __props__.__dict__["subscription_id"] = subscription_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["use_msal"] = pulumi.Output.from_input(use_msal).apply(pulumi.runtime.to_json) if use_msal is not None else None
            __props__.__dict__["use_msi"] = pulumi.Output.from_input(use_msi).apply(pulumi.runtime.to_json) if use_msi is not None else None
        super(Provider, __self__).__init__(
            'azure',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="clientCertificatePassword")
    def client_certificate_password(self) -> pulumi.Output[Optional[str]]:
        """
        The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return pulumi.get(self, "client_certificate_password")

    @property
    @pulumi.getter(name="clientCertificatePath")
    def client_certificate_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return pulumi.get(self, "client_certificate_path")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Client ID which should be used.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="metadataHost")
    def metadata_host(self) -> pulumi.Output[Optional[str]]:
        """
        The Hostname which should be used for the Azure Metadata Service.
        """
        return pulumi.get(self, "metadata_host")

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> pulumi.Output[Optional[str]]:
        """
        Deprecated - replaced by `metadata_host`.
        """
        return pulumi.get(self, "metadata_url")

    @property
    @pulumi.getter(name="msiEndpoint")
    def msi_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
        automatically.
        """
        return pulumi.get(self, "msi_endpoint")

    @property
    @pulumi.getter(name="partnerId")
    def partner_id(self) -> pulumi.Output[Optional[str]]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return pulumi.get(self, "partner_id")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Subscription ID which should be used.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Tenant ID which should be used.
        """
        return pulumi.get(self, "tenant_id")

