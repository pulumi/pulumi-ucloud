// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPN Gateway resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const fooVPC = new ucloud.vpc.VPC("fooVPC", {
 *     tag: "tf-acc",
 *     cidrBlocks: ["192.168.0.0/16"],
 * });
 * const fooEIP = new ucloud.unet.EIP("fooEIP", {
 *     bandwidth: 1,
 *     internetType: "bgp",
 *     chargeMode: "bandwidth",
 *     tag: "tf-acc",
 * });
 * const fooVPNGateway = new ucloud.ipsecvpn.VPNGateway("fooVPNGateway", {
 *     vpcId: fooVPC.id,
 *     grade: "enhanced",
 *     eipId: fooEIP.id,
 *     tag: "tf-acc",
 * });
 * ```
 *
 * ## Import
 *
 * VPN Gateway can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import ucloud:ipsecvpn/vPNGateway:VPNGateway example vpngw-abc123456
 * ```
 */
export class VPNGateway extends pulumi.CustomResource {
    /**
     * Get an existing VPNGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VPNGatewayState, opts?: pulumi.CustomResourceOptions): VPNGateway {
        return new VPNGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:ipsecvpn/vPNGateway:VPNGateway';

    /**
     * Returns true if the given object is an instance of VPNGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VPNGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VPNGateway.__pulumiType;
    }

    /**
     * The charge type of VPN Gateway, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    public readonly chargeType!: pulumi.Output<string>;
    /**
     * The creation time for VPN Gateway, formatted in RFC3339 time string.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The duration that you will buy the VPN Gateway (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    public readonly duration!: pulumi.Output<number | undefined>;
    /**
     * The ID of eip associate to the VPN Gateway.
     */
    public readonly eipId!: pulumi.Output<string>;
    /**
     * The expiration time for VPN Gateway, formatted in RFC3339 time string.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * The type of the VPN Gateway. Possible values: `standard`, `enhanced`. `standard` recommended application scenario: Applicable to services with bidirectional peak bandwidth of 1M~50M; `enhanced` recommended application scenario: Suitable for services with bidirectional peak bandwidths of 50M~100M.
     */
    public readonly grade!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The remarks of the VPN Gateway. (Default: `""`).
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * A tag assigned to VPN Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * The ID of VPC linked to the VPN Gateway.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VPNGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPNGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VPNGatewayArgs | VPNGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VPNGatewayState | undefined;
            inputs["chargeType"] = state ? state.chargeType : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["eipId"] = state ? state.eipId : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["grade"] = state ? state.grade : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["tag"] = state ? state.tag : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VPNGatewayArgs | undefined;
            if ((!args || args.eipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eipId'");
            }
            if ((!args || args.grade === undefined) && !opts.urn) {
                throw new Error("Missing required property 'grade'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["chargeType"] = args ? args.chargeType : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["eipId"] = args ? args.eipId : undefined;
            inputs["grade"] = args ? args.grade : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VPNGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VPNGateway resources.
 */
export interface VPNGatewayState {
    /**
     * The charge type of VPN Gateway, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The creation time for VPN Gateway, formatted in RFC3339 time string.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The duration that you will buy the VPN Gateway (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    duration?: pulumi.Input<number>;
    /**
     * The ID of eip associate to the VPN Gateway.
     */
    eipId?: pulumi.Input<string>;
    /**
     * The expiration time for VPN Gateway, formatted in RFC3339 time string.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * The type of the VPN Gateway. Possible values: `standard`, `enhanced`. `standard` recommended application scenario: Applicable to services with bidirectional peak bandwidth of 1M~50M; `enhanced` recommended application scenario: Suitable for services with bidirectional peak bandwidths of 50M~100M.
     */
    grade?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the VPN Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * A tag assigned to VPN Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the VPN Gateway.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VPNGateway resource.
 */
export interface VPNGatewayArgs {
    /**
     * The charge type of VPN Gateway, possible values are: `year`, `month` and `dynamic` as pay by hour (specific permission required). (Default: `month`).
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the VPN Gateway (Default: `1`). The value is `0` when pay by month and the instance will be valid till the last day of that month. It is not required when `dynamic` (pay by hour).
     */
    duration?: pulumi.Input<number>;
    /**
     * The ID of eip associate to the VPN Gateway.
     */
    eipId: pulumi.Input<string>;
    /**
     * The type of the VPN Gateway. Possible values: `standard`, `enhanced`. `standard` recommended application scenario: Applicable to services with bidirectional peak bandwidth of 1M~50M; `enhanced` recommended application scenario: Suitable for services with bidirectional peak bandwidths of 50M~100M.
     */
    grade: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the VPN Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * A tag assigned to VPN Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the VPN Gateway.
     */
    vpcId: pulumi.Input<string>;
}
