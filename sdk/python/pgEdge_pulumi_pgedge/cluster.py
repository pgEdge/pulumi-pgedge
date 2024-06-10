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

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cloud_account_id: pulumi.Input[str],
                 regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]] = None,
                 node_location: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]] = None,
                 ssh_key_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cloud_account_id: ID of the target cloud account
        :param pulumi.Input[str] name: Name of the network
        :param pulumi.Input[str] node_location: Network location for nodes (public or private)
        :param pulumi.Input[str] ssh_key_id: ID of the SSH key to add to the cluster nodes
        """
        pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        pulumi.set(__self__, "regions", regions)
        if firewall_rules is not None:
            pulumi.set(__self__, "firewall_rules", firewall_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if node_location is not None:
            pulumi.set(__self__, "node_location", node_location)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if ssh_key_id is not None:
            pulumi.set(__self__, "ssh_key_id", ssh_key_id)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Input[str]:
        """
        ID of the target cloud account
        """
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]]:
        return pulumi.get(self, "firewall_rules")

    @firewall_rules.setter
    def firewall_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]]):
        pulumi.set(self, "firewall_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the network
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter(name="nodeLocation")
    def node_location(self) -> Optional[pulumi.Input[str]]:
        """
        Network location for nodes (public or private)
        """
        return pulumi.get(self, "node_location")

    @node_location.setter
    def node_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_location", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]]:
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter(name="sshKeyId")
    def ssh_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the SSH key to add to the cluster nodes
        """
        return pulumi.get(self, "ssh_key_id")

    @ssh_key_id.setter
    def ssh_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_key_id", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]] = None,
                 node_location: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ssh_key_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cloud_account_id: ID of the target cloud account
        :param pulumi.Input[str] created_at: Creation time of the cluster
        :param pulumi.Input[str] name: Name of the network
        :param pulumi.Input[str] node_location: Network location for nodes (public or private)
        :param pulumi.Input[str] ssh_key_id: ID of the SSH key to add to the cluster nodes
        :param pulumi.Input[str] status: Status of the cluster
        """
        if cloud_account_id is not None:
            pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if firewall_rules is not None:
            pulumi.set(__self__, "firewall_rules", firewall_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if node_location is not None:
            pulumi.set(__self__, "node_location", node_location)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if ssh_key_id is not None:
            pulumi.set(__self__, "ssh_key_id", ssh_key_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the target cloud account
        """
        return pulumi.get(self, "cloud_account_id")

    @cloud_account_id.setter
    def cloud_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_account_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of the cluster
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]]:
        return pulumi.get(self, "firewall_rules")

    @firewall_rules.setter
    def firewall_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterFirewallRuleArgs']]]]):
        pulumi.set(self, "firewall_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the network
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter(name="nodeLocation")
    def node_location(self) -> Optional[pulumi.Input[str]]:
        """
        Network location for nodes (public or private)
        """
        return pulumi.get(self, "node_location")

    @node_location.setter
    def node_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_location", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]]:
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="sshKeyId")
    def ssh_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the SSH key to add to the cluster nodes
        """
        return pulumi.get(self, "ssh_key_id")

    @ssh_key_id.setter
    def ssh_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_key_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the cluster
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]]]] = None,
                 node_location: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodeArgs']]]]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ssh_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Interface with the pgEdge service API for clusters.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_account_id: ID of the target cloud account
        :param pulumi.Input[str] name: Name of the network
        :param pulumi.Input[str] node_location: Network location for nodes (public or private)
        :param pulumi.Input[str] ssh_key_id: ID of the SSH key to add to the cluster nodes
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Interface with the pgEdge service API for clusters.

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_account_id: Optional[pulumi.Input[str]] = None,
                 firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]]]] = None,
                 node_location: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodeArgs']]]]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ssh_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            if cloud_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_account_id'")
            __props__.__dict__["cloud_account_id"] = cloud_account_id
            __props__.__dict__["firewall_rules"] = firewall_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["networks"] = networks
            __props__.__dict__["node_location"] = node_location
            __props__.__dict__["nodes"] = nodes
            if regions is None and not opts.urn:
                raise TypeError("Missing required property 'regions'")
            __props__.__dict__["regions"] = regions
            __props__.__dict__["ssh_key_id"] = ssh_key_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
        super(Cluster, __self__).__init__(
            'pgedge:index/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_account_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            firewall_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterFirewallRuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]]]] = None,
            node_location: Optional[pulumi.Input[str]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterNodeArgs']]]]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ssh_key_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_account_id: ID of the target cloud account
        :param pulumi.Input[str] created_at: Creation time of the cluster
        :param pulumi.Input[str] name: Name of the network
        :param pulumi.Input[str] node_location: Network location for nodes (public or private)
        :param pulumi.Input[str] ssh_key_id: ID of the SSH key to add to the cluster nodes
        :param pulumi.Input[str] status: Status of the cluster
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cloud_account_id"] = cloud_account_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["firewall_rules"] = firewall_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["networks"] = networks
        __props__.__dict__["node_location"] = node_location
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["regions"] = regions
        __props__.__dict__["ssh_key_id"] = ssh_key_id
        __props__.__dict__["status"] = status
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> pulumi.Output[str]:
        """
        ID of the target cloud account
        """
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation time of the cluster
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterFirewallRule']]]:
        return pulumi.get(self, "firewall_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the network
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence['outputs.ClusterNetwork']]:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="nodeLocation")
    def node_location(self) -> pulumi.Output[str]:
        """
        Network location for nodes (public or private)
        """
        return pulumi.get(self, "node_location")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence['outputs.ClusterNode']]:
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="sshKeyId")
    def ssh_key_id(self) -> pulumi.Output[str]:
        """
        ID of the SSH key to add to the cluster nodes
        """
        return pulumi.get(self, "ssh_key_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the cluster
        """
        return pulumi.get(self, "status")

