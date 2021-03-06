// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Subnet resource under VPC resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vpc.NewVPC(ctx, "_default", &vpc.VPCArgs{
// 			Tag: pulumi.String("tf-example"),
// 			CidrBlocks: pulumi.StringArray{
// 				pulumi.String("192.168.0.0/16"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewSubnet(ctx, "example", &vpc.SubnetArgs{
// 			Tag:       pulumi.String("tf-example"),
// 			CidrBlock: pulumi.String("192.168.1.0/24"),
// 			VpcId:     _default.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Subnet can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import ucloud:vpc/subnet:Subnet example subnet-abc123456
// ```
type Subnet struct {
	pulumi.CustomResourceState

	// The cidr block of the desired subnet, format in "0.0.0.0/0", such as: `192.168.0.0/24`.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The time of creation of subnet, formatted in RFC3339 time string.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	Name       pulumi.StringOutput `pulumi:"name"`
	// The remarks of the subnet. (Default: `""`).
	Remark pulumi.StringOutput `pulumi:"remark"`
	// A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The id of the VPC that the desired subnet belongs to.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Subnet
	err := ctx.RegisterResource("ucloud:vpc/subnet:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("ucloud:vpc/subnet:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// The cidr block of the desired subnet, format in "0.0.0.0/0", such as: `192.168.0.0/24`.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The time of creation of subnet, formatted in RFC3339 time string.
	CreateTime *string `pulumi:"createTime"`
	Name       *string `pulumi:"name"`
	// The remarks of the subnet. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The id of the VPC that the desired subnet belongs to.
	VpcId *string `pulumi:"vpcId"`
}

type SubnetState struct {
	// The cidr block of the desired subnet, format in "0.0.0.0/0", such as: `192.168.0.0/24`.
	CidrBlock pulumi.StringPtrInput
	// The time of creation of subnet, formatted in RFC3339 time string.
	CreateTime pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	// The remarks of the subnet. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The id of the VPC that the desired subnet belongs to.
	VpcId pulumi.StringPtrInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// The cidr block of the desired subnet, format in "0.0.0.0/0", such as: `192.168.0.0/24`.
	CidrBlock string  `pulumi:"cidrBlock"`
	Name      *string `pulumi:"name"`
	// The remarks of the subnet. (Default: `""`).
	Remark *string `pulumi:"remark"`
	// A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag *string `pulumi:"tag"`
	// The id of the VPC that the desired subnet belongs to.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// The cidr block of the desired subnet, format in "0.0.0.0/0", such as: `192.168.0.0/24`.
	CidrBlock pulumi.StringInput
	Name      pulumi.StringPtrInput
	// The remarks of the subnet. (Default: `""`).
	Remark pulumi.StringPtrInput
	// A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
	Tag pulumi.StringPtrInput
	// The id of the VPC that the desired subnet belongs to.
	VpcId pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (*Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil))
}

func (i *Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i *Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i *Subnet) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *Subnet) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

type SubnetPtrInput interface {
	pulumi.Input

	ToSubnetPtrOutput() SubnetPtrOutput
	ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput
}

type subnetPtrType SubnetArgs

func (*subnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil))
}

func (i *subnetPtrType) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *subnetPtrType) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

// SubnetArrayInput is an input type that accepts SubnetArray and SubnetArrayOutput values.
// You can construct a concrete instance of `SubnetArrayInput` via:
//
//          SubnetArray{ SubnetArgs{...} }
type SubnetArrayInput interface {
	pulumi.Input

	ToSubnetArrayOutput() SubnetArrayOutput
	ToSubnetArrayOutputWithContext(context.Context) SubnetArrayOutput
}

type SubnetArray []SubnetInput

func (SubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subnet)(nil))
}

func (i SubnetArray) ToSubnetArrayOutput() SubnetArrayOutput {
	return i.ToSubnetArrayOutputWithContext(context.Background())
}

func (i SubnetArray) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetArrayOutput)
}

// SubnetMapInput is an input type that accepts SubnetMap and SubnetMapOutput values.
// You can construct a concrete instance of `SubnetMapInput` via:
//
//          SubnetMap{ "key": SubnetArgs{...} }
type SubnetMapInput interface {
	pulumi.Input

	ToSubnetMapOutput() SubnetMapOutput
	ToSubnetMapOutputWithContext(context.Context) SubnetMapOutput
}

type SubnetMap map[string]SubnetInput

func (SubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subnet)(nil))
}

func (i SubnetMap) ToSubnetMapOutput() SubnetMapOutput {
	return i.ToSubnetMapOutputWithContext(context.Background())
}

func (i SubnetMap) ToSubnetMapOutputWithContext(ctx context.Context) SubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetMapOutput)
}

type SubnetOutput struct {
	*pulumi.OutputState
}

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil))
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o.ToSubnetPtrOutputWithContext(context.Background())
}

func (o SubnetOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o.ApplyT(func(v Subnet) *Subnet {
		return &v
	}).(SubnetPtrOutput)
}

type SubnetPtrOutput struct {
	*pulumi.OutputState
}

func (SubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil))
}

func (o SubnetPtrOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o
}

type SubnetArrayOutput struct{ *pulumi.OutputState }

func (SubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil))
}

func (o SubnetArrayOutput) ToSubnetArrayOutput() SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) Index(i pulumi.IntInput) SubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subnet {
		return vs[0].([]Subnet)[vs[1].(int)]
	}).(SubnetOutput)
}

type SubnetMapOutput struct{ *pulumi.OutputState }

func (SubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subnet)(nil))
}

func (o SubnetMapOutput) ToSubnetMapOutput() SubnetMapOutput {
	return o
}

func (o SubnetMapOutput) ToSubnetMapOutputWithContext(ctx context.Context) SubnetMapOutput {
	return o
}

func (o SubnetMapOutput) MapIndex(k pulumi.StringInput) SubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subnet {
		return vs[0].(map[string]Subnet)[vs[1].(string)]
	}).(SubnetOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetArrayOutput{})
	pulumi.RegisterOutputType(SubnetMapOutput{})
}
