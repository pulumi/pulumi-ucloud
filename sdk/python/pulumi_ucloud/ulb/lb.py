# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Lb(pulumi.CustomResource):
    charge_type: pulumi.Output[str]
    """
    **Deprecated**, argument `charge_type` is deprecated for optimizing parameters.
    """
    create_time: pulumi.Output[str]
    """
    The time of creation for load balancer, formatted in RFC3339 time string.
    """
    expire_time: pulumi.Output[str]
    """
    **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
    """
    internal: pulumi.Output[bool]
    """
    Indicate whether the load balancer is intranet mode.(Default: `"false"`)
    """
    ip_sets: pulumi.Output[list]
    """
    It is a nested type which documented below.
    
      * `internetType` (`str`) - Type of Elastic IP routes.
      * `ip` (`str`) - Elastic IP address.
    """
    name: pulumi.Output[str]
    private_ip: pulumi.Output[str]
    """
    The IP address of intranet IP. It is `""` if `internal` is `false`.
    """
    remark: pulumi.Output[str]
    """
    The remarks of the load balancer. (Default: `""`).
    """
    subnet_id: pulumi.Output[str]
    """
    The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
    """
    tag: pulumi.Output[str]
    """
    A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
    """
    def __init__(__self__, resource_name, opts=None, charge_type=None, internal=None, name=None, remark=None, subnet_id=None, tag=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: **Deprecated**, argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/lb.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['charge_type'] = charge_type
            __props__['internal'] = internal
            __props__['name'] = name
            __props__['remark'] = remark
            __props__['subnet_id'] = subnet_id
            __props__['tag'] = tag
            __props__['vpc_id'] = vpc_id
            __props__['create_time'] = None
            __props__['expire_time'] = None
            __props__['ip_sets'] = None
            __props__['private_ip'] = None
        super(Lb, __self__).__init__(
            'ucloud:ulb/lb:Lb',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, charge_type=None, create_time=None, expire_time=None, internal=None, ip_sets=None, name=None, private_ip=None, remark=None, subnet_id=None, tag=None, vpc_id=None):
        """
        Get an existing Lb resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] charge_type: **Deprecated**, argument `charge_type` is deprecated for optimizing parameters.
        :param pulumi.Input[str] create_time: The time of creation for load balancer, formatted in RFC3339 time string.
        :param pulumi.Input[str] expire_time: **Deprecated** attribute `expire_time` is deprecated for optimizing outputs.
        :param pulumi.Input[bool] internal: Indicate whether the load balancer is intranet mode.(Default: `"false"`)
        :param pulumi.Input[list] ip_sets: It is a nested type which documented below.
        :param pulumi.Input[str] private_ip: The IP address of intranet IP. It is `""` if `internal` is `false`.
        :param pulumi.Input[str] remark: The remarks of the load balancer. (Default: `""`).
        :param pulumi.Input[str] subnet_id: The ID of subnet that intranet load balancer belongs to. This argument is not required if default subnet.
        :param pulumi.Input[str] tag: A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        :param pulumi.Input[str] vpc_id: The ID of the VPC linked to the Load balancer, This argument is not required if default VPC.
        
        The **ip_sets** object supports the following:
        
          * `internetType` (`pulumi.Input[str]`) - Type of Elastic IP routes.
          * `ip` (`pulumi.Input[str]`) - Elastic IP address.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/lb.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["charge_type"] = charge_type
        __props__["create_time"] = create_time
        __props__["expire_time"] = expire_time
        __props__["internal"] = internal
        __props__["ip_sets"] = ip_sets
        __props__["name"] = name
        __props__["private_ip"] = private_ip
        __props__["remark"] = remark
        __props__["subnet_id"] = subnet_id
        __props__["tag"] = tag
        __props__["vpc_id"] = vpc_id
        return Lb(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
