# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'VaultEncryption',
    'VaultIdentity',
]

@pulumi.output_type
class VaultEncryption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "infrastructureEncryptionEnabled":
            suggest = "infrastructure_encryption_enabled"
        elif key == "keyId":
            suggest = "key_id"
        elif key == "useSystemAssignedIdentity":
            suggest = "use_system_assigned_identity"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VaultEncryption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VaultEncryption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VaultEncryption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 infrastructure_encryption_enabled: bool,
                 key_id: str,
                 use_system_assigned_identity: Optional[bool] = None):
        """
        :param bool infrastructure_encryption_enabled: Enabling/Disabling the Double Encryption state.
        :param str key_id: The Key Vault key id used to encrypt this vault. Key managed by Vault Managed Hardware Security Module is also supported.
        :param bool use_system_assigned_identity: Indicate that system assigned identity should be used or not. At this time the only possible value is `true`. Defaults to `true`.
        """
        pulumi.set(__self__, "infrastructure_encryption_enabled", infrastructure_encryption_enabled)
        pulumi.set(__self__, "key_id", key_id)
        if use_system_assigned_identity is not None:
            pulumi.set(__self__, "use_system_assigned_identity", use_system_assigned_identity)

    @property
    @pulumi.getter(name="infrastructureEncryptionEnabled")
    def infrastructure_encryption_enabled(self) -> bool:
        """
        Enabling/Disabling the Double Encryption state.
        """
        return pulumi.get(self, "infrastructure_encryption_enabled")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        The Key Vault key id used to encrypt this vault. Key managed by Vault Managed Hardware Security Module is also supported.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="useSystemAssignedIdentity")
    def use_system_assigned_identity(self) -> Optional[bool]:
        """
        Indicate that system assigned identity should be used or not. At this time the only possible value is `true`. Defaults to `true`.
        """
        return pulumi.get(self, "use_system_assigned_identity")


@pulumi.output_type
class VaultIdentity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VaultIdentity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VaultIdentity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VaultIdentity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


