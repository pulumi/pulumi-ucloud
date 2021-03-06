// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer SSL attachment resource for attaching SSL certificate to Load Balancer Listener.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 * import * from "fs";
 *
 * const fooLB = new ucloud.ulb.LB("fooLB", {tag: "tf-example"});
 * const fooLBListener = new ucloud.ulb.LBListener("fooLBListener", {
 *     loadBalancerId: fooLB.id,
 *     protocol: "https",
 * });
 * const fooLBSsl = new ucloud.ulb.LBSsl("fooLBSsl", {
 *     privateKey: fs.readFileSync("private.key"),
 *     userCert: fs.readFileSync("user.crt"),
 *     caCert: fs.readFileSync("ca.crt"),
 * });
 * const fooLBSslAttachment = new ucloud.ulb.LBSslAttachment("fooLBSslAttachment", {
 *     loadBalancerId: fooLB.id,
 *     listenerId: fooLBListener.id,
 *     sslId: fooLBSsl.id,
 * });
 * ```
 */
export class LBSslAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LBSslAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LBSslAttachmentState, opts?: pulumi.CustomResourceOptions): LBSslAttachment {
        return new LBSslAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:ulb/lBSslAttachment:LBSslAttachment';

    /**
     * Returns true if the given object is an instance of LBSslAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LBSslAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LBSslAttachment.__pulumiType;
    }

    /**
     * The ID of listener servers.
     */
    public readonly listenerId!: pulumi.Output<string>;
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The ID of SSL certificate.
     */
    public readonly sslId!: pulumi.Output<string>;

    /**
     * Create a LBSslAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LBSslAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LBSslAttachmentArgs | LBSslAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LBSslAttachmentState | undefined;
            inputs["listenerId"] = state ? state.listenerId : undefined;
            inputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            inputs["sslId"] = state ? state.sslId : undefined;
        } else {
            const args = argsOrState as LBSslAttachmentArgs | undefined;
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.sslId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslId'");
            }
            inputs["listenerId"] = args ? args.listenerId : undefined;
            inputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            inputs["sslId"] = args ? args.sslId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LBSslAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LBSslAttachment resources.
 */
export interface LBSslAttachmentState {
    /**
     * The ID of listener servers.
     */
    listenerId?: pulumi.Input<string>;
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The ID of SSL certificate.
     */
    sslId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LBSslAttachment resource.
 */
export interface LBSslAttachmentArgs {
    /**
     * The ID of listener servers.
     */
    listenerId: pulumi.Input<string>;
    loadBalancerId: pulumi.Input<string>;
    /**
     * The ID of SSL certificate.
     */
    sslId: pulumi.Input<string>;
}
