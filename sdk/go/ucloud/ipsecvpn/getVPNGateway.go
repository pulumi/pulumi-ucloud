// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ipsecvpn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source providers a list of VPN Gateway resources according to their ID, name, vpc and tag.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/ipsecvpn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ipsecvpn.LookupVPNGateway(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("first", example.VpnGateways[0].Id)
// 		return nil
// 	})
// }
// ```
func LookupVPNGateway(ctx *pulumi.Context, args *LookupVPNGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVPNGatewayResult, error) {
	var rv LookupVPNGatewayResult
	err := ctx.Invoke("ucloud:ipsecvpn/getVPNGateway:getVPNGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVPNGateway.
type LookupVPNGatewayArgs struct {
	// A list of VPN Gateway IDs, all the VPN Gateways belongs to the defined region will be retrieved if this argument is `[]`.
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting VPN Gateways by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A tag assigned to VPN Gateway.
	Tag *string `pulumi:"tag"`
	// The ID of VPC linked to the VPN Gateway.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVPNGateway.
type LookupVPNGatewayResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// A tag assigned to the VPN Gateway.
	Tag *string `pulumi:"tag"`
	// Total number of VPN Gateways that satisfy the condition.
	TotalCount int `pulumi:"totalCount"`
	// The ID of VPC linked to the VPN Gateway.
	VpcId *string `pulumi:"vpcId"`
	// It is a nested type. VPN Gateways documented below.
	VpnGateways []GetVPNGatewayVpnGateway `pulumi:"vpnGateways"`
}