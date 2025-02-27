# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'VaultEncryptionArgs',
    'VaultIdentityArgs',
]

@pulumi.input_type
class VaultEncryptionArgs:
    def __init__(__self__, *,
                 infrastructure_encryption_enabled: pulumi.Input[bool],
                 key_id: pulumi.Input[str],
                 use_system_assigned_identity: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] infrastructure_encryption_enabled: Enabling/Disabling the Double Encryption state.
        :param pulumi.Input[str] key_id: The Key Vault key id used to encrypt this vault. Key managed by Vault Managed Hardware Security Module is also supported.
        :param pulumi.Input[bool] use_system_assigned_identity: Indicate that system assigned identity should be used or not. At this time the only possible value is `true`. Defaults to `true`.
        """
        pulumi.set(__self__, "infrastructure_encryption_enabled", infrastructure_encryption_enabled)
        pulumi.set(__self__, "key_id", key_id)
        if use_system_assigned_identity is not None:
            pulumi.set(__self__, "use_system_assigned_identity", use_system_assigned_identity)

    @property
    @pulumi.getter(name="infrastructureEncryptionEnabled")
    def infrastructure_encryption_enabled(self) -> pulumi.Input[bool]:
        """
        Enabling/Disabling the Double Encryption state.
        """
        return pulumi.get(self, "infrastructure_encryption_enabled")

    @infrastructure_encryption_enabled.setter
    def infrastructure_encryption_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "infrastructure_encryption_enabled", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        The Key Vault key id used to encrypt this vault. Key managed by Vault Managed Hardware Security Module is also supported.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="useSystemAssignedIdentity")
    def use_system_assigned_identity(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate that system assigned identity should be used or not. At this time the only possible value is `true`. Defaults to `true`.
        """
        return pulumi.get(self, "use_system_assigned_identity")

    @use_system_assigned_identity.setter
    def use_system_assigned_identity(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_system_assigned_identity", value)


@pulumi.input_type
class VaultIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


