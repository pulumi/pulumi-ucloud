// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package uhost

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an UHost Instance resource.
//
// > **Note** The instance will reboot automatically to make the change take effect when update `instanceType`, `rootPassword`, `bootDiskSize`, `dataDiskSize`. In addition, once the instance complete creation, it takes around 10 minutes for boot disk initialization for the running instance, and the updates will only be made to some specific attributes (`rootPassword`, `bootDiskSize`) if required once the instance initialization completed.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.ImageId == nil {
		return nil, errors.New("missing required argument 'ImageId'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["availabilityZone"] = nil
		inputs["bootDiskSize"] = nil
		inputs["bootDiskType"] = nil
		inputs["chargeType"] = nil
		inputs["dataDiskSize"] = nil
		inputs["dataDiskType"] = nil
		inputs["duration"] = nil
		inputs["imageId"] = nil
		inputs["instanceType"] = nil
		inputs["isolationGroup"] = nil
		inputs["name"] = nil
		inputs["remark"] = nil
		inputs["rootPassword"] = nil
		inputs["securityGroup"] = nil
		inputs["subnetId"] = nil
		inputs["tag"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["bootDiskSize"] = args.BootDiskSize
		inputs["bootDiskType"] = args.BootDiskType
		inputs["chargeType"] = args.ChargeType
		inputs["dataDiskSize"] = args.DataDiskSize
		inputs["dataDiskType"] = args.DataDiskType
		inputs["duration"] = args.Duration
		inputs["imageId"] = args.ImageId
		inputs["instanceType"] = args.InstanceType
		inputs["isolationGroup"] = args.IsolationGroup
		inputs["name"] = args.Name
		inputs["remark"] = args.Remark
		inputs["rootPassword"] = args.RootPassword
		inputs["securityGroup"] = args.SecurityGroup
		inputs["subnetId"] = args.SubnetId
		inputs["tag"] = args.Tag
		inputs["vpcId"] = args.VpcId
	}
	inputs["autoRenew"] = nil
	inputs["cpu"] = nil
	inputs["createTime"] = nil
	inputs["diskSets"] = nil
	inputs["expireTime"] = nil
	inputs["ipSets"] = nil
	inputs["memory"] = nil
	inputs["privateIp"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("ucloud:uhost/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoRenew"] = state.AutoRenew
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["bootDiskSize"] = state.BootDiskSize
		inputs["bootDiskType"] = state.BootDiskType
		inputs["chargeType"] = state.ChargeType
		inputs["cpu"] = state.Cpu
		inputs["createTime"] = state.CreateTime
		inputs["dataDiskSize"] = state.DataDiskSize
		inputs["dataDiskType"] = state.DataDiskType
		inputs["diskSets"] = state.DiskSets
		inputs["duration"] = state.Duration
		inputs["expireTime"] = state.ExpireTime
		inputs["imageId"] = state.ImageId
		inputs["instanceType"] = state.InstanceType
		inputs["ipSets"] = state.IpSets
		inputs["isolationGroup"] = state.IsolationGroup
		inputs["memory"] = state.Memory
		inputs["name"] = state.Name
		inputs["privateIp"] = state.PrivateIp
		inputs["remark"] = state.Remark
		inputs["rootPassword"] = state.RootPassword
		inputs["securityGroup"] = state.SecurityGroup
		inputs["status"] = state.Status
		inputs["subnetId"] = state.SubnetId
		inputs["tag"] = state.Tag
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("ucloud:uhost/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Whether to renew an instance automatically or not.
func (r *Instance) AutoRenew() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoRenew"])
}

// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
func (r *Instance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
func (r *Instance) BootDiskSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["bootDiskSize"])
}

// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
func (r *Instance) BootDiskType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bootDiskType"])
}

// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
func (r *Instance) ChargeType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["chargeType"])
}

// The number of cores of virtual CPU, measured in core.
func (r *Instance) Cpu() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpu"])
}

// The time of creation for instance, formatted in RFC3339 time string.
func (r *Instance) CreateTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTime"])
}

// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
func (r *Instance) DataDiskSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dataDiskSize"])
}

// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
func (r *Instance) DataDiskType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dataDiskType"])
}

// It is a nested type which documented below.
func (r *Instance) DiskSets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["diskSets"])
}

// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
func (r *Instance) Duration() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["duration"])
}

// The expiration time for instance, formatted in RFC3339 time string.
func (r *Instance) ExpireTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expireTime"])
}

// The ID for the image to use for the instance.
func (r *Instance) ImageId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["imageId"])
}

func (r *Instance) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

// It is a nested type which documented below.
func (r *Instance) IpSets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipSets"])
}

// The ID of the associated isolation group.
func (r *Instance) IsolationGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["isolationGroup"])
}

// The size of memory, measured in MB (Megabyte).
func (r *Instance) Memory() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["memory"])
}

func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The private IP address assigned to the instance.
func (r *Instance) PrivateIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIp"])
}

// The remarks of instance. (Default: `""`).
func (r *Instance) Remark() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["remark"])
}

func (r *Instance) RootPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["rootPassword"])
}

// The ID of the associated security group.
func (r *Instance) SecurityGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["securityGroup"])
}

// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
func (r *Instance) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
func (r *Instance) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
func (r *Instance) Tag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tag"])
}

// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
func (r *Instance) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// Whether to renew an instance automatically or not.
	AutoRenew interface{}
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone interface{}
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize interface{}
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType interface{}
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType interface{}
	// The number of cores of virtual CPU, measured in core.
	Cpu interface{}
	// The time of creation for instance, formatted in RFC3339 time string.
	CreateTime interface{}
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize interface{}
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType interface{}
	// It is a nested type which documented below.
	DiskSets interface{}
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration interface{}
	// The expiration time for instance, formatted in RFC3339 time string.
	ExpireTime interface{}
	// The ID for the image to use for the instance.
	ImageId      interface{}
	InstanceType interface{}
	// It is a nested type which documented below.
	IpSets interface{}
	// The ID of the associated isolation group.
	IsolationGroup interface{}
	// The size of memory, measured in MB (Megabyte).
	Memory interface{}
	Name   interface{}
	// The private IP address assigned to the instance.
	PrivateIp interface{}
	// The remarks of instance. (Default: `""`).
	Remark       interface{}
	RootPassword interface{}
	// The ID of the associated security group.
	SecurityGroup interface{}
	// Instance current status. Possible values are `Initializing`, `Starting`, `Running`, `Stopping`, `Stopped`, `Install Fail`, `ResizeFail` and `Rebooting`.
	Status interface{}
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId interface{}
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag interface{}
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Availability zone where instance is located. such as: `cn-bj2-02`. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)
	AvailabilityZone interface{}
	// The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to `20`(default: `20`) for Linux and `40` (default: `40`) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported.
	BootDiskSize interface{}
	// The type of boot disk. Possible values are: `localNormal` and `localSsd` for local boot disk, `cloudSsd` for cloud SSD boot disk. (Default: `localNormal`). The `localSsd` and `cloudSsd` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.
	BootDiskType interface{}
	// The charge type of instance, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
	ChargeType interface{}
	// The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: `20`), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported.
	DataDiskSize interface{}
	// The type of local data disk. Possible values are: `localNormal` and `localSsd` for local data disk. (Default: `localNormal`). The `localSsd` is not fully support by all regions as data disk type, please proceed to UCloud console for more details. In addition, the `dataDiskType` must be same as `bootDiskType` if specified.
	DataDiskType interface{}
	// The duration that you will buy the instance (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
	Duration interface{}
	// The ID for the image to use for the instance.
	ImageId      interface{}
	InstanceType interface{}
	// The ID of the associated isolation group.
	IsolationGroup interface{}
	Name           interface{}
	// The remarks of instance. (Default: `""`).
	Remark       interface{}
	RootPassword interface{}
	// The ID of the associated security group.
	SecurityGroup interface{}
	// The ID of subnet. If defined `vpcId`, the `subnetId` is Required. If not defined `vpcId` and `subnetId`, the instance will use the default subnet in the current region.
	SubnetId interface{}
	// A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag interface{}
	// The ID of VPC linked to the instance. If not defined `vpcId`, the instance will use the default VPC in the current region.
	VpcId interface{}
}