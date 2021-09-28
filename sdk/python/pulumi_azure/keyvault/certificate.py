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

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 certificate_policy: pulumi.Input['CertificateCertificatePolicyArgs'],
                 key_vault_id: pulumi.Input[str],
                 certificate: Optional[pulumi.Input['CertificateCertificateArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input['CertificateCertificatePolicyArgs'] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input['CertificateCertificateArgs'] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "certificate_policy", certificate_policy)
        pulumi.set(__self__, "key_vault_id", key_vault_id)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="certificatePolicy")
    def certificate_policy(self) -> pulumi.Input['CertificateCertificatePolicyArgs']:
        """
        A `certificate_policy` block as defined below.
        """
        return pulumi.get(self, "certificate_policy")

    @certificate_policy.setter
    def certificate_policy(self, value: pulumi.Input['CertificateCertificatePolicyArgs']):
        pulumi.set(self, "certificate_policy", value)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Input[str]:
        """
        The ID of the Key Vault where the Certificate should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input['CertificateCertificateArgs']]:
        """
        A `certificate` block as defined below, used to Import an existing certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input['CertificateCertificateArgs']]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


@pulumi.input_type
class _CertificateState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input['CertificateCertificateArgs']] = None,
                 certificate_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateCertificateAttributeArgs']]]] = None,
                 certificate_data: Optional[pulumi.Input[str]] = None,
                 certificate_data_base64: Optional[pulumi.Input[str]] = None,
                 certificate_policy: Optional[pulumi.Input['CertificateCertificatePolicyArgs']] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Certificate resources.
        :param pulumi.Input['CertificateCertificateArgs'] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[Sequence[pulumi.Input['CertificateCertificateAttributeArgs']]] certificate_attributes: A `certificate_attribute` block as defined below.
        :param pulumi.Input[str] certificate_data: The raw Key Vault Certificate data represented as a hexadecimal string.
        :param pulumi.Input[str] certificate_data_base64: The Base64 encoded Key Vault Certificate data.
        :param pulumi.Input['CertificateCertificatePolicyArgs'] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secret_id: The ID of the associated Key Vault Secret.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] thumbprint: The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        :param pulumi.Input[str] version: The current version of the Key Vault Certificate.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if certificate_attributes is not None:
            pulumi.set(__self__, "certificate_attributes", certificate_attributes)
        if certificate_data is not None:
            pulumi.set(__self__, "certificate_data", certificate_data)
        if certificate_data_base64 is not None:
            pulumi.set(__self__, "certificate_data_base64", certificate_data_base64)
        if certificate_policy is not None:
            pulumi.set(__self__, "certificate_policy", certificate_policy)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input['CertificateCertificateArgs']]:
        """
        A `certificate` block as defined below, used to Import an existing certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input['CertificateCertificateArgs']]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="certificateAttributes")
    def certificate_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CertificateCertificateAttributeArgs']]]]:
        """
        A `certificate_attribute` block as defined below.
        """
        return pulumi.get(self, "certificate_attributes")

    @certificate_attributes.setter
    def certificate_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateCertificateAttributeArgs']]]]):
        pulumi.set(self, "certificate_attributes", value)

    @property
    @pulumi.getter(name="certificateData")
    def certificate_data(self) -> Optional[pulumi.Input[str]]:
        """
        The raw Key Vault Certificate data represented as a hexadecimal string.
        """
        return pulumi.get(self, "certificate_data")

    @certificate_data.setter
    def certificate_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_data", value)

    @property
    @pulumi.getter(name="certificateDataBase64")
    def certificate_data_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The Base64 encoded Key Vault Certificate data.
        """
        return pulumi.get(self, "certificate_data_base64")

    @certificate_data_base64.setter
    def certificate_data_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_data_base64", value)

    @property
    @pulumi.getter(name="certificatePolicy")
    def certificate_policy(self) -> Optional[pulumi.Input['CertificateCertificatePolicyArgs']]:
        """
        A `certificate_policy` block as defined below.
        """
        return pulumi.get(self, "certificate_policy")

    @certificate_policy.setter
    def certificate_policy(self, value: Optional[pulumi.Input['CertificateCertificatePolicyArgs']]):
        pulumi.set(self, "certificate_policy", value)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault where the Certificate should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated Key Vault Secret.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

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
    def thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The current version of the Key Vault Certificate.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Certificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[pulumi.InputType['CertificateCertificateArgs']]] = None,
                 certificate_policy: Optional[pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a Key Vault Certificate.

        ## Example Usage
        ### Importing a PFX

        > **Note:** this example assumed the PFX file is located in the same directory at `certificate-to-import.pfx`.

        ```python
        import pulumi
        import base64
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                certificate_permissions=[
                    "create",
                    "delete",
                    "deleteissuers",
                    "get",
                    "getissuers",
                    "import",
                    "list",
                    "listissuers",
                    "managecontacts",
                    "manageissuers",
                    "setissuers",
                    "update",
                ],
                key_permissions=[
                    "backup",
                    "create",
                    "decrypt",
                    "delete",
                    "encrypt",
                    "get",
                    "import",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "sign",
                    "unwrapKey",
                    "update",
                    "verify",
                    "wrapKey",
                ],
                secret_permissions=[
                    "backup",
                    "delete",
                    "get",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "set",
                ],
            )])
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=example_key_vault.id,
            certificate=azure.keyvault.CertificateCertificateArgs(
                contents=(lambda path: base64.b64encode(open(path).read().encode()).decode())("certificate-to-import.pfx"),
                password="",
            ),
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties=azure.keyvault.CertificateCertificatePolicyKeyPropertiesArgs(
                    exportable=True,
                    key_size=2048,
                    key_type="RSA",
                    reuse_key=False,
                ),
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
            ))
        ```
        ### Generating a new certificate

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                certificate_permissions=[
                    "create",
                    "delete",
                    "deleteissuers",
                    "get",
                    "getissuers",
                    "import",
                    "list",
                    "listissuers",
                    "managecontacts",
                    "manageissuers",
                    "purge",
                    "setissuers",
                    "update",
                ],
                key_permissions=[
                    "backup",
                    "create",
                    "decrypt",
                    "delete",
                    "encrypt",
                    "get",
                    "import",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "sign",
                    "unwrapKey",
                    "update",
                    "verify",
                    "wrapKey",
                ],
                secret_permissions=[
                    "backup",
                    "delete",
                    "get",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "set",
                ],
            )])
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=example_key_vault.id,
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties=azure.keyvault.CertificateCertificatePolicyKeyPropertiesArgs(
                    exportable=True,
                    key_size=2048,
                    key_type="RSA",
                    reuse_key=True,
                ),
                lifetime_actions=[azure.keyvault.CertificateCertificatePolicyLifetimeActionArgs(
                    action=azure.keyvault.CertificateCertificatePolicyLifetimeActionActionArgs(
                        action_type="AutoRenew",
                    ),
                    trigger=azure.keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs(
                        days_before_expiry=30,
                    ),
                )],
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
                x509_certificate_properties=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs(
                    extended_key_usages=["1.3.6.1.5.5.7.3.1"],
                    key_usages=[
                        "cRLSign",
                        "dataEncipherment",
                        "digitalSignature",
                        "keyAgreement",
                        "keyCertSign",
                        "keyEncipherment",
                    ],
                    subject_alternative_names=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs(
                        dns_names=[
                            "internal.contoso.com",
                            "domain.hello.world",
                        ],
                    ),
                    subject="CN=hello-world",
                    validity_in_months=12,
                ),
            ))
        ```

        ## Import

        Key Vault Certificates can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:keyvault/certificate:Certificate example "https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateCertificateArgs']] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Key Vault Certificate.

        ## Example Usage
        ### Importing a PFX

        > **Note:** this example assumed the PFX file is located in the same directory at `certificate-to-import.pfx`.

        ```python
        import pulumi
        import base64
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="premium",
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                certificate_permissions=[
                    "create",
                    "delete",
                    "deleteissuers",
                    "get",
                    "getissuers",
                    "import",
                    "list",
                    "listissuers",
                    "managecontacts",
                    "manageissuers",
                    "setissuers",
                    "update",
                ],
                key_permissions=[
                    "backup",
                    "create",
                    "decrypt",
                    "delete",
                    "encrypt",
                    "get",
                    "import",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "sign",
                    "unwrapKey",
                    "update",
                    "verify",
                    "wrapKey",
                ],
                secret_permissions=[
                    "backup",
                    "delete",
                    "get",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "set",
                ],
            )])
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=example_key_vault.id,
            certificate=azure.keyvault.CertificateCertificateArgs(
                contents=(lambda path: base64.b64encode(open(path).read().encode()).decode())("certificate-to-import.pfx"),
                password="",
            ),
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties=azure.keyvault.CertificateCertificatePolicyKeyPropertiesArgs(
                    exportable=True,
                    key_size=2048,
                    key_type="RSA",
                    reuse_key=False,
                ),
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
            ))
        ```
        ### Generating a new certificate

        ```python
        import pulumi
        import pulumi_azure as azure

        current = azure.core.get_client_config()
        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_key_vault = azure.keyvault.KeyVault("exampleKeyVault",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            tenant_id=current.tenant_id,
            sku_name="standard",
            soft_delete_retention_days=7,
            access_policies=[azure.keyvault.KeyVaultAccessPolicyArgs(
                tenant_id=current.tenant_id,
                object_id=current.object_id,
                certificate_permissions=[
                    "create",
                    "delete",
                    "deleteissuers",
                    "get",
                    "getissuers",
                    "import",
                    "list",
                    "listissuers",
                    "managecontacts",
                    "manageissuers",
                    "purge",
                    "setissuers",
                    "update",
                ],
                key_permissions=[
                    "backup",
                    "create",
                    "decrypt",
                    "delete",
                    "encrypt",
                    "get",
                    "import",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "sign",
                    "unwrapKey",
                    "update",
                    "verify",
                    "wrapKey",
                ],
                secret_permissions=[
                    "backup",
                    "delete",
                    "get",
                    "list",
                    "purge",
                    "recover",
                    "restore",
                    "set",
                ],
            )])
        example_certificate = azure.keyvault.Certificate("exampleCertificate",
            key_vault_id=example_key_vault.id,
            certificate_policy=azure.keyvault.CertificateCertificatePolicyArgs(
                issuer_parameters=azure.keyvault.CertificateCertificatePolicyIssuerParametersArgs(
                    name="Self",
                ),
                key_properties=azure.keyvault.CertificateCertificatePolicyKeyPropertiesArgs(
                    exportable=True,
                    key_size=2048,
                    key_type="RSA",
                    reuse_key=True,
                ),
                lifetime_actions=[azure.keyvault.CertificateCertificatePolicyLifetimeActionArgs(
                    action=azure.keyvault.CertificateCertificatePolicyLifetimeActionActionArgs(
                        action_type="AutoRenew",
                    ),
                    trigger=azure.keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs(
                        days_before_expiry=30,
                    ),
                )],
                secret_properties=azure.keyvault.CertificateCertificatePolicySecretPropertiesArgs(
                    content_type="application/x-pkcs12",
                ),
                x509_certificate_properties=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs(
                    extended_key_usages=["1.3.6.1.5.5.7.3.1"],
                    key_usages=[
                        "cRLSign",
                        "dataEncipherment",
                        "digitalSignature",
                        "keyAgreement",
                        "keyCertSign",
                        "keyEncipherment",
                    ],
                    subject_alternative_names=azure.keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs(
                        dns_names=[
                            "internal.contoso.com",
                            "domain.hello.world",
                        ],
                    ),
                    subject="CN=hello-world",
                    validity_in_months=12,
                ),
            ))
        ```

        ## Import

        Key Vault Certificates can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:keyvault/certificate:Certificate example "https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217"
        ```

        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[pulumi.InputType['CertificateCertificateArgs']]] = None,
                 certificate_policy: Optional[pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = CertificateArgs.__new__(CertificateArgs)

            __props__.__dict__["certificate"] = certificate
            if certificate_policy is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_policy'")
            __props__.__dict__["certificate_policy"] = certificate_policy
            if key_vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_vault_id'")
            __props__.__dict__["key_vault_id"] = key_vault_id
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["certificate_attributes"] = None
            __props__.__dict__["certificate_data"] = None
            __props__.__dict__["certificate_data_base64"] = None
            __props__.__dict__["secret_id"] = None
            __props__.__dict__["thumbprint"] = None
            __props__.__dict__["version"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure:keyvault/certifiate:Certifiate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Certificate, __self__).__init__(
            'azure:keyvault/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[pulumi.InputType['CertificateCertificateArgs']]] = None,
            certificate_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateCertificateAttributeArgs']]]]] = None,
            certificate_data: Optional[pulumi.Input[str]] = None,
            certificate_data_base64: Optional[pulumi.Input[str]] = None,
            certificate_policy: Optional[pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateCertificateArgs']] certificate: A `certificate` block as defined below, used to Import an existing certificate.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateCertificateAttributeArgs']]]] certificate_attributes: A `certificate_attribute` block as defined below.
        :param pulumi.Input[str] certificate_data: The raw Key Vault Certificate data represented as a hexadecimal string.
        :param pulumi.Input[str] certificate_data_base64: The Base64 encoded Key Vault Certificate data.
        :param pulumi.Input[pulumi.InputType['CertificateCertificatePolicyArgs']] certificate_policy: A `certificate_policy` block as defined below.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Certificate should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secret_id: The ID of the associated Key Vault Secret.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] thumbprint: The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        :param pulumi.Input[str] version: The current version of the Key Vault Certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateState.__new__(_CertificateState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["certificate_attributes"] = certificate_attributes
        __props__.__dict__["certificate_data"] = certificate_data
        __props__.__dict__["certificate_data_base64"] = certificate_data_base64
        __props__.__dict__["certificate_policy"] = certificate_policy
        __props__.__dict__["key_vault_id"] = key_vault_id
        __props__.__dict__["name"] = name
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["thumbprint"] = thumbprint
        __props__.__dict__["version"] = version
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional['outputs.CertificateCertificate']]:
        """
        A `certificate` block as defined below, used to Import an existing certificate.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateAttributes")
    def certificate_attributes(self) -> pulumi.Output[Sequence['outputs.CertificateCertificateAttribute']]:
        """
        A `certificate_attribute` block as defined below.
        """
        return pulumi.get(self, "certificate_attributes")

    @property
    @pulumi.getter(name="certificateData")
    def certificate_data(self) -> pulumi.Output[str]:
        """
        The raw Key Vault Certificate data represented as a hexadecimal string.
        """
        return pulumi.get(self, "certificate_data")

    @property
    @pulumi.getter(name="certificateDataBase64")
    def certificate_data_base64(self) -> pulumi.Output[str]:
        """
        The Base64 encoded Key Vault Certificate data.
        """
        return pulumi.get(self, "certificate_data_base64")

    @property
    @pulumi.getter(name="certificatePolicy")
    def certificate_policy(self) -> pulumi.Output['outputs.CertificateCertificatePolicy']:
        """
        A `certificate_policy` block as defined below.
        """
        return pulumi.get(self, "certificate_policy")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the Key Vault where the Certificate should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated Key Vault Secret.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[str]:
        """
        The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the Key Vault Certificate.
        """
        return pulumi.get(self, "version")

