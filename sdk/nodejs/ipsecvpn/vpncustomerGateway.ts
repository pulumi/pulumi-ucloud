// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPN Customer Gateway resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const foo = new ucloud.ipsecvpn.VPNCustomerGateway("foo", {
 *     ipAddress: "10.0.0.1",
 *     tag: "tf-acc",
 * });
 * ```
 *
 * ## Import
 *
 * VPN Customer Gateway can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import ucloud:ipsecvpn/vPNCustomerGateway:VPNCustomerGateway example remotevpngw-abc123456
 * ```
 */
export class VPNCustomerGateway extends pulumi.CustomResource {
    /**
     * Get an existing VPNCustomerGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VPNCustomerGatewayState, opts?: pulumi.CustomResourceOptions): VPNCustomerGateway {
        return new VPNCustomerGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:ipsecvpn/vPNCustomerGateway:VPNCustomerGateway';

    /**
     * Returns true if the given object is an instance of VPNCustomerGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VPNCustomerGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VPNCustomerGateway.__pulumiType;
    }

    /**
     * The creation time for VPN Customer Gateway, formatted in RFC3339 time string.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The ip address of the VPN Customer Gateway.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The remarks of the VPN Customer Gateway. (Default: `""`).
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    public readonly tag!: pulumi.Output<string | undefined>;

    /**
     * Create a VPNCustomerGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPNCustomerGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VPNCustomerGatewayArgs | VPNCustomerGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VPNCustomerGatewayState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["remark"] = state ? state.remark : undefined;
            inputs["tag"] = state ? state.tag : undefined;
        } else {
            const args = argsOrState as VPNCustomerGatewayArgs | undefined;
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["tag"] = args ? args.tag : undefined;
            inputs["createTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VPNCustomerGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VPNCustomerGateway resources.
 */
export interface VPNCustomerGatewayState {
    /**
     * The creation time for VPN Customer Gateway, formatted in RFC3339 time string.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ip address of the VPN Customer Gateway.
     */
    ipAddress?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the VPN Customer Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VPNCustomerGateway resource.
 */
export interface VPNCustomerGatewayArgs {
    /**
     * The ip address of the VPN Customer Gateway.
     */
    ipAddress: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * The remarks of the VPN Customer Gateway. (Default: `""`).
     */
    remark?: pulumi.Input<string>;
    /**
     * A tag assigned to VPN Customer Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
     * * ``
     */
    tag?: pulumi.Input<string>;
}
