# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DatabaseArgs', 'Database']

@pulumi.input_type
class DatabaseArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 backups: Optional[pulumi.Input['DatabaseBackupsArgs']] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input['DatabaseExtensionsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] cluster_id: ID of the cluster to place the database on
        :param pulumi.Input['DatabaseBackupsArgs'] backups: Backup configuration for the database.
        :param pulumi.Input[str] config_version: Config version of the database
        :param pulumi.Input['DatabaseExtensionsArgs'] extensions: Extensions configuration for the database.
        :param pulumi.Input[str] name: Name of the database
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]] nodes: List of nodes in the database.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] options: Options for creating the database
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]] roles: List of roles in the database.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if backups is not None:
            pulumi.set(__self__, "backups", backups)
        if config_version is not None:
            pulumi.set(__self__, "config_version", config_version)
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        ID of the cluster to place the database on
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def backups(self) -> Optional[pulumi.Input['DatabaseBackupsArgs']]:
        """
        Backup configuration for the database.
        """
        return pulumi.get(self, "backups")

    @backups.setter
    def backups(self, value: Optional[pulumi.Input['DatabaseBackupsArgs']]):
        pulumi.set(self, "backups", value)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[pulumi.Input[str]]:
        """
        Config version of the database
        """
        return pulumi.get(self, "config_version")

    @config_version.setter
    def config_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version", value)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[pulumi.Input['DatabaseExtensionsArgs']]:
        """
        Extensions configuration for the database.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input['DatabaseExtensionsArgs']]):
        pulumi.set(self, "extensions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]]:
        """
        List of nodes in the database.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Options for creating the database
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]]:
        """
        List of roles in the database.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]]):
        pulumi.set(self, "roles", value)


@pulumi.input_type
class _DatabaseState:
    def __init__(__self__, *,
                 backups: Optional[pulumi.Input['DatabaseBackupsArgs']] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 components: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseComponentArgs']]]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input['DatabaseExtensionsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 pg_version: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 storage_used: Optional[pulumi.Input[int]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Database resources.
        :param pulumi.Input['DatabaseBackupsArgs'] backups: Backup configuration for the database.
        :param pulumi.Input[str] cluster_id: ID of the cluster to place the database on
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseComponentArgs']]] components: List of components in the database.
        :param pulumi.Input[str] config_version: Config version of the database
        :param pulumi.Input[str] created_at: Created at of the database
        :param pulumi.Input[str] domain: Domain of the database
        :param pulumi.Input['DatabaseExtensionsArgs'] extensions: Extensions configuration for the database.
        :param pulumi.Input[str] name: Name of the database
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]] nodes: List of nodes in the database.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] options: Options for creating the database
        :param pulumi.Input[str] pg_version: Postgres version of the database
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]] roles: List of roles in the database.
        :param pulumi.Input[str] status: Status of the database
        :param pulumi.Input[int] storage_used: Storage used of the database
        :param pulumi.Input[str] updated_at: Updated at of the database
        """
        if backups is not None:
            pulumi.set(__self__, "backups", backups)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if components is not None:
            pulumi.set(__self__, "components", components)
        if config_version is not None:
            pulumi.set(__self__, "config_version", config_version)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if pg_version is not None:
            pulumi.set(__self__, "pg_version", pg_version)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if storage_used is not None:
            pulumi.set(__self__, "storage_used", storage_used)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def backups(self) -> Optional[pulumi.Input['DatabaseBackupsArgs']]:
        """
        Backup configuration for the database.
        """
        return pulumi.get(self, "backups")

    @backups.setter
    def backups(self, value: Optional[pulumi.Input['DatabaseBackupsArgs']]):
        pulumi.set(self, "backups", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the cluster to place the database on
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def components(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseComponentArgs']]]]:
        """
        List of components in the database.
        """
        return pulumi.get(self, "components")

    @components.setter
    def components(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseComponentArgs']]]]):
        pulumi.set(self, "components", value)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[pulumi.Input[str]]:
        """
        Config version of the database
        """
        return pulumi.get(self, "config_version")

    @config_version.setter
    def config_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Created at of the database
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain of the database
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[pulumi.Input['DatabaseExtensionsArgs']]:
        """
        Extensions configuration for the database.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input['DatabaseExtensionsArgs']]):
        pulumi.set(self, "extensions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]]:
        """
        List of nodes in the database.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Options for creating the database
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="pgVersion")
    def pg_version(self) -> Optional[pulumi.Input[str]]:
        """
        Postgres version of the database
        """
        return pulumi.get(self, "pg_version")

    @pg_version.setter
    def pg_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pg_version", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]]:
        """
        List of roles in the database.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseRoleArgs']]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the database
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="storageUsed")
    def storage_used(self) -> Optional[pulumi.Input[int]]:
        """
        Storage used of the database
        """
        return pulumi.get(self, "storage_used")

    @storage_used.setter
    def storage_used(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_used", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Updated at of the database
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backups: Optional[pulumi.Input[pulumi.InputType['DatabaseBackupsArgs']]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[pulumi.InputType['DatabaseExtensionsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseNodeArgs']]]]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseRoleArgs']]]]] = None,
                 __props__=None):
        """
        Interface with the pgEdge service API.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DatabaseBackupsArgs']] backups: Backup configuration for the database.
        :param pulumi.Input[str] cluster_id: ID of the cluster to place the database on
        :param pulumi.Input[str] config_version: Config version of the database
        :param pulumi.Input[pulumi.InputType['DatabaseExtensionsArgs']] extensions: Extensions configuration for the database.
        :param pulumi.Input[str] name: Name of the database
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseNodeArgs']]]] nodes: List of nodes in the database.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] options: Options for creating the database
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseRoleArgs']]]] roles: List of roles in the database.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Interface with the pgEdge service API.

        :param str resource_name: The name of the resource.
        :param DatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backups: Optional[pulumi.Input[pulumi.InputType['DatabaseBackupsArgs']]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[pulumi.InputType['DatabaseExtensionsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseNodeArgs']]]]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseRoleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseArgs.__new__(DatabaseArgs)

            __props__.__dict__["backups"] = backups
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["config_version"] = config_version
            __props__.__dict__["extensions"] = extensions
            __props__.__dict__["name"] = name
            __props__.__dict__["nodes"] = nodes
            __props__.__dict__["options"] = options
            __props__.__dict__["roles"] = roles
            __props__.__dict__["components"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["domain"] = None
            __props__.__dict__["pg_version"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["storage_used"] = None
            __props__.__dict__["updated_at"] = None
        super(Database, __self__).__init__(
            'pgedge:index/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backups: Optional[pulumi.Input[pulumi.InputType['DatabaseBackupsArgs']]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            components: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseComponentArgs']]]]] = None,
            config_version: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            extensions: Optional[pulumi.Input[pulumi.InputType['DatabaseExtensionsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseNodeArgs']]]]] = None,
            options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            pg_version: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseRoleArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            storage_used: Optional[pulumi.Input[int]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DatabaseBackupsArgs']] backups: Backup configuration for the database.
        :param pulumi.Input[str] cluster_id: ID of the cluster to place the database on
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseComponentArgs']]]] components: List of components in the database.
        :param pulumi.Input[str] config_version: Config version of the database
        :param pulumi.Input[str] created_at: Created at of the database
        :param pulumi.Input[str] domain: Domain of the database
        :param pulumi.Input[pulumi.InputType['DatabaseExtensionsArgs']] extensions: Extensions configuration for the database.
        :param pulumi.Input[str] name: Name of the database
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseNodeArgs']]]] nodes: List of nodes in the database.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] options: Options for creating the database
        :param pulumi.Input[str] pg_version: Postgres version of the database
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseRoleArgs']]]] roles: List of roles in the database.
        :param pulumi.Input[str] status: Status of the database
        :param pulumi.Input[int] storage_used: Storage used of the database
        :param pulumi.Input[str] updated_at: Updated at of the database
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseState.__new__(_DatabaseState)

        __props__.__dict__["backups"] = backups
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["components"] = components
        __props__.__dict__["config_version"] = config_version
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["domain"] = domain
        __props__.__dict__["extensions"] = extensions
        __props__.__dict__["name"] = name
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["options"] = options
        __props__.__dict__["pg_version"] = pg_version
        __props__.__dict__["roles"] = roles
        __props__.__dict__["status"] = status
        __props__.__dict__["storage_used"] = storage_used
        __props__.__dict__["updated_at"] = updated_at
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backups(self) -> pulumi.Output['outputs.DatabaseBackups']:
        """
        Backup configuration for the database.
        """
        return pulumi.get(self, "backups")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        ID of the cluster to place the database on
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def components(self) -> pulumi.Output[Sequence['outputs.DatabaseComponent']]:
        """
        List of components in the database.
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> pulumi.Output[str]:
        """
        Config version of the database
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Created at of the database
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain of the database
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def extensions(self) -> pulumi.Output['outputs.DatabaseExtensions']:
        """
        Extensions configuration for the database.
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the database
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence['outputs.DatabaseNode']]:
        """
        List of nodes in the database.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Options for creating the database
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="pgVersion")
    def pg_version(self) -> pulumi.Output[str]:
        """
        Postgres version of the database
        """
        return pulumi.get(self, "pg_version")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence['outputs.DatabaseRole']]:
        """
        List of roles in the database.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the database
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageUsed")
    def storage_used(self) -> pulumi.Output[int]:
        """
        Storage used of the database
        """
        return pulumi.get(self, "storage_used")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Updated at of the database
        """
        return pulumi.get(self, "updated_at")

