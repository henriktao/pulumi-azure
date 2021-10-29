# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'Features',
    'FeaturesApiManagement',
    'FeaturesCognitiveAccount',
    'FeaturesKeyVault',
    'FeaturesLogAnalyticsWorkspace',
    'FeaturesNetwork',
    'FeaturesResourceGroup',
    'FeaturesTemplateDeployment',
    'FeaturesVirtualMachine',
    'FeaturesVirtualMachineScaleSet',
]

@pulumi.output_type
class Features(dict):
    def __init__(__self__, *,
                 api_management: Optional['outputs.FeaturesApiManagement'] = None,
                 cognitive_account: Optional['outputs.FeaturesCognitiveAccount'] = None,
                 key_vault: Optional['outputs.FeaturesKeyVault'] = None,
                 log_analytics_workspace: Optional['outputs.FeaturesLogAnalyticsWorkspace'] = None,
                 network: Optional['outputs.FeaturesNetwork'] = None,
                 resource_group: Optional['outputs.FeaturesResourceGroup'] = None,
                 template_deployment: Optional['outputs.FeaturesTemplateDeployment'] = None,
                 virtual_machine: Optional['outputs.FeaturesVirtualMachine'] = None,
                 virtual_machine_scale_set: Optional['outputs.FeaturesVirtualMachineScaleSet'] = None):
        if api_management is not None:
            pulumi.set(__self__, "api_management", api_management)
        if cognitive_account is not None:
            pulumi.set(__self__, "cognitive_account", cognitive_account)
        if key_vault is not None:
            pulumi.set(__self__, "key_vault", key_vault)
        if log_analytics_workspace is not None:
            pulumi.set(__self__, "log_analytics_workspace", log_analytics_workspace)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if resource_group is not None:
            pulumi.set(__self__, "resource_group", resource_group)
        if template_deployment is not None:
            pulumi.set(__self__, "template_deployment", template_deployment)
        if virtual_machine is not None:
            pulumi.set(__self__, "virtual_machine", virtual_machine)
        if virtual_machine_scale_set is not None:
            pulumi.set(__self__, "virtual_machine_scale_set", virtual_machine_scale_set)

    @property
    @pulumi.getter(name="apiManagement")
    def api_management(self) -> Optional['outputs.FeaturesApiManagement']:
        return pulumi.get(self, "api_management")

    @property
    @pulumi.getter(name="cognitiveAccount")
    def cognitive_account(self) -> Optional['outputs.FeaturesCognitiveAccount']:
        return pulumi.get(self, "cognitive_account")

    @property
    @pulumi.getter(name="keyVault")
    def key_vault(self) -> Optional['outputs.FeaturesKeyVault']:
        return pulumi.get(self, "key_vault")

    @property
    @pulumi.getter(name="logAnalyticsWorkspace")
    def log_analytics_workspace(self) -> Optional['outputs.FeaturesLogAnalyticsWorkspace']:
        return pulumi.get(self, "log_analytics_workspace")

    @property
    @pulumi.getter
    def network(self) -> Optional['outputs.FeaturesNetwork']:
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> Optional['outputs.FeaturesResourceGroup']:
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="templateDeployment")
    def template_deployment(self) -> Optional['outputs.FeaturesTemplateDeployment']:
        return pulumi.get(self, "template_deployment")

    @property
    @pulumi.getter(name="virtualMachine")
    def virtual_machine(self) -> Optional['outputs.FeaturesVirtualMachine']:
        return pulumi.get(self, "virtual_machine")

    @property
    @pulumi.getter(name="virtualMachineScaleSet")
    def virtual_machine_scale_set(self) -> Optional['outputs.FeaturesVirtualMachineScaleSet']:
        return pulumi.get(self, "virtual_machine_scale_set")


@pulumi.output_type
class FeaturesApiManagement(dict):
    def __init__(__self__, *,
                 purge_soft_delete_on_destroy: Optional[bool] = None):
        if purge_soft_delete_on_destroy is not None:
            pulumi.set(__self__, "purge_soft_delete_on_destroy", purge_soft_delete_on_destroy)

    @property
    @pulumi.getter(name="purgeSoftDeleteOnDestroy")
    def purge_soft_delete_on_destroy(self) -> Optional[bool]:
        return pulumi.get(self, "purge_soft_delete_on_destroy")


@pulumi.output_type
class FeaturesCognitiveAccount(dict):
    def __init__(__self__, *,
                 purge_soft_delete_on_destroy: Optional[bool] = None):
        if purge_soft_delete_on_destroy is not None:
            pulumi.set(__self__, "purge_soft_delete_on_destroy", purge_soft_delete_on_destroy)

    @property
    @pulumi.getter(name="purgeSoftDeleteOnDestroy")
    def purge_soft_delete_on_destroy(self) -> Optional[bool]:
        return pulumi.get(self, "purge_soft_delete_on_destroy")


@pulumi.output_type
class FeaturesKeyVault(dict):
    def __init__(__self__, *,
                 purge_soft_delete_on_destroy: Optional[bool] = None,
                 recover_soft_deleted_key_vaults: Optional[bool] = None):
        if purge_soft_delete_on_destroy is not None:
            pulumi.set(__self__, "purge_soft_delete_on_destroy", purge_soft_delete_on_destroy)
        if recover_soft_deleted_key_vaults is not None:
            pulumi.set(__self__, "recover_soft_deleted_key_vaults", recover_soft_deleted_key_vaults)

    @property
    @pulumi.getter(name="purgeSoftDeleteOnDestroy")
    def purge_soft_delete_on_destroy(self) -> Optional[bool]:
        return pulumi.get(self, "purge_soft_delete_on_destroy")

    @property
    @pulumi.getter(name="recoverSoftDeletedKeyVaults")
    def recover_soft_deleted_key_vaults(self) -> Optional[bool]:
        return pulumi.get(self, "recover_soft_deleted_key_vaults")


@pulumi.output_type
class FeaturesLogAnalyticsWorkspace(dict):
    def __init__(__self__, *,
                 permanently_delete_on_destroy: bool):
        pulumi.set(__self__, "permanently_delete_on_destroy", permanently_delete_on_destroy)

    @property
    @pulumi.getter(name="permanentlyDeleteOnDestroy")
    def permanently_delete_on_destroy(self) -> bool:
        return pulumi.get(self, "permanently_delete_on_destroy")


@pulumi.output_type
class FeaturesNetwork(dict):
    def __init__(__self__, *,
                 relaxed_locking: bool):
        pulumi.set(__self__, "relaxed_locking", relaxed_locking)

    @property
    @pulumi.getter(name="relaxedLocking")
    def relaxed_locking(self) -> bool:
        return pulumi.get(self, "relaxed_locking")


@pulumi.output_type
class FeaturesResourceGroup(dict):
    def __init__(__self__, *,
                 prevent_deletion_if_contains_resources: Optional[bool] = None):
        if prevent_deletion_if_contains_resources is not None:
            pulumi.set(__self__, "prevent_deletion_if_contains_resources", prevent_deletion_if_contains_resources)

    @property
    @pulumi.getter(name="preventDeletionIfContainsResources")
    def prevent_deletion_if_contains_resources(self) -> Optional[bool]:
        return pulumi.get(self, "prevent_deletion_if_contains_resources")


@pulumi.output_type
class FeaturesTemplateDeployment(dict):
    def __init__(__self__, *,
                 delete_nested_items_during_deletion: bool):
        pulumi.set(__self__, "delete_nested_items_during_deletion", delete_nested_items_during_deletion)

    @property
    @pulumi.getter(name="deleteNestedItemsDuringDeletion")
    def delete_nested_items_during_deletion(self) -> bool:
        return pulumi.get(self, "delete_nested_items_during_deletion")


@pulumi.output_type
class FeaturesVirtualMachine(dict):
    def __init__(__self__, *,
                 delete_os_disk_on_deletion: Optional[bool] = None,
                 graceful_shutdown: Optional[bool] = None,
                 skip_shutdown_and_force_delete: Optional[bool] = None):
        if delete_os_disk_on_deletion is not None:
            pulumi.set(__self__, "delete_os_disk_on_deletion", delete_os_disk_on_deletion)
        if graceful_shutdown is not None:
            pulumi.set(__self__, "graceful_shutdown", graceful_shutdown)
        if skip_shutdown_and_force_delete is not None:
            pulumi.set(__self__, "skip_shutdown_and_force_delete", skip_shutdown_and_force_delete)

    @property
    @pulumi.getter(name="deleteOsDiskOnDeletion")
    def delete_os_disk_on_deletion(self) -> Optional[bool]:
        return pulumi.get(self, "delete_os_disk_on_deletion")

    @property
    @pulumi.getter(name="gracefulShutdown")
    def graceful_shutdown(self) -> Optional[bool]:
        return pulumi.get(self, "graceful_shutdown")

    @property
    @pulumi.getter(name="skipShutdownAndForceDelete")
    def skip_shutdown_and_force_delete(self) -> Optional[bool]:
        return pulumi.get(self, "skip_shutdown_and_force_delete")


@pulumi.output_type
class FeaturesVirtualMachineScaleSet(dict):
    def __init__(__self__, *,
                 roll_instances_when_required: bool,
                 force_delete: Optional[bool] = None,
                 scale_to_zero_before_deletion: Optional[bool] = None):
        pulumi.set(__self__, "roll_instances_when_required", roll_instances_when_required)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if scale_to_zero_before_deletion is not None:
            pulumi.set(__self__, "scale_to_zero_before_deletion", scale_to_zero_before_deletion)

    @property
    @pulumi.getter(name="rollInstancesWhenRequired")
    def roll_instances_when_required(self) -> bool:
        return pulumi.get(self, "roll_instances_when_required")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[bool]:
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter(name="scaleToZeroBeforeDeletion")
    def scale_to_zero_before_deletion(self) -> Optional[bool]:
        return pulumi.get(self, "scale_to_zero_before_deletion")


