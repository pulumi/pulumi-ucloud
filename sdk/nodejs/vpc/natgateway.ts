// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Nat Gateway resource.
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
 * const fooSubnet = new ucloud.vpc.Subnet("fooSubnet", {
 *     tag: "tf-acc",
 *     cidrBlock: "192.168.1.0/24",
 *     vpcId: fooVPC.id,
 * });
 * const fooEIP = new ucloud.unet.EIP("fooEIP", {
 *     bandwidth: 1,
 *     internetType: "bgp",
 *     chargeMode: "bandwidth",
 *     tag: "tf-acc",
 * });
 * const fooSecurityGroup = ucloud.unet.getSecurityGroup({
 *     type: "recommend_web",
 * });
 * const fooNATGateway = new ucloud.vpc.NATGateway("fooNATGateway", {
 *     vpcId: fooVPC.id,
 *     subnetIds: [fooSubnet.id],
 *     eipId: fooEIP.id,
 *     tag: "tf-acc",
 *     securityGroup: fooSecurityGroup.then(fooSecurityGroup => fooSecurityGroup.securityGroups[0].id),
 * });
 * ```
 *
 * ## Import
 *
 * Nat Gateway can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import ucloud:vpc/nATGateway:NATGateway example natgw-abc123456
 * ```
 */
export class NATGateway extends pulumi.CustomResource {
    /**
     * Get an existing NATGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NATGatewayState, opts?: pulumi.CustomResourceOptions): NATGateway {
        return new NATGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:vpc/nATGateway:NATGateway';

    /**
     * Returns true if the given object is an instance of NATGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NATGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NATGateway.__pulumiType;
    }

    /**
     * The time of creation of Nat Gateway, formatted in RFC3339 time string.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    public readonly eipId!: pulumi.Output<string>;
    /**
     * The boolean value to Controls whether or not start the whitelist mode.
     */
    public readonly enableWhiteList!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The remarks of the Nat Gateway. (Default: `""`).
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * The ID of the associated security group.
     */
    public readonly securityGroup!: pulumi.Output<string>;
    /**
     * The list of subnet ID under the VPC.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * The ID of VPC linked to the Nat Gateway.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The white list of instance under the Nat Gateway.
     */
    public readonly whiteLists!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NATGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NATGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NATGatewayArgs | NATGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NATGatewayState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["eipId"] = state ? state.eipId : undefined;
            inputs["enableWhiteList"] = state ? state.enableWhiteList : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["securityGroup"] = state ? state.securityGroup : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tag"] = state ? state.tag : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["whiteLists"] = state ? state.whiteLists : undefined;
        } else {
            const args = argsOrState as NATGatewayArgs | undefined;
            if ((!args || args.eipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eipId'");
            }
            if ((!args || args.enableWhiteList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enableWhiteList'");
            }
            if ((!args || args.securityGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroup'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["eipId"] = args ? args.eipId : undefined;
            inputs["enableWhiteList"] = args ? args.enableWhiteList : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["securityGroup"] = args ? args.securityGroup : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["whiteLists"] = args ? args.whiteLists : undefined;
            inputs["createTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NATGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NATGateway resources.
 */
export interface NATGatewayState {
    /**
     * The time of creation of Nat Gateway, formatted in RFC3339 time string.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    eipId?: pulumi.Input<string>;
    /**
     * The boolean value to Controls whether or not start the whitelist mode.
     */
    enableWhiteList?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the Nat Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * The ID of the associated security group.
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The list of subnet ID under the VPC.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the Nat Gateway.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The white list of instance under the Nat Gateway.
     */
    whiteLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NATGateway resource.
 */
export interface NATGatewayArgs {
    /**
     * The ID of eip associate to the Nat Gateway.
     */
    eipId: pulumi.Input<string>;
    /**
     * The boolean value to Controls whether or not start the whitelist mode.
     */
    enableWhiteList: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the Nat Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * The ID of the associated security group.
     */
    securityGroup: pulumi.Input<string>;
    /**
     * The list of subnet ID under the VPC.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
    /**
     * The ID of VPC linked to the Nat Gateway.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The white list of instance under the Nat Gateway.
     */
    whiteLists?: pulumi.Input<pulumi.Input<string>[]>;
}