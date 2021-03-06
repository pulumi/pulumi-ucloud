// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EIP Association resource for associating Elastic IP to UHost Instance, Load Balancer, etc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ucloud from "@pulumi/ucloud";
 *
 * const defaultZone = ucloud.uaccount.getZone({});
 * const defaultImage = defaultZone.then(defaultZone => ucloud.uhost.getImage({
 *     availabilityZone: defaultZone.zones[0].id,
 *     nameRegex: "^CentOS 7.[1-2] 64",
 *     imageType: "base",
 * }));
 * // Create security group
 * const defaultSecurityGroup = new ucloud.unet.SecurityGroup("defaultSecurityGroup", {
 *     tag: "tf-example",
 *     rules: [{
 *         portRange: "80",
 *         protocol: "tcp",
 *         cidrBlock: "0.0.0.0/0",
 *         policy: "accept",
 *     }],
 * });
 * // Create an eip
 * const defaultEIP = new ucloud.unet.EIP("defaultEIP", {
 *     bandwidth: 2,
 *     chargeMode: "bandwidth",
 *     tag: "tf-example",
 *     internetType: "bgp",
 * });
 * // Create a web server
 * const web = new ucloud.uhost.Instance("web", {
 *     instanceType: "n-basic-2",
 *     availabilityZone: defaultZone.then(defaultZone => defaultZone.zones[0].id),
 *     imageId: defaultImage.then(defaultImage => defaultImage.images[0].id),
 *     dataDiskSize: 50,
 *     rootPassword: "wA1234567",
 *     securityGroup: defaultSecurityGroup.id,
 *     tag: "tf-example",
 * });
 * // Bind eip to instance
 * const defaultEIPAssociation = new ucloud.unet.EIPAssociation("defaultEIPAssociation", {
 *     resourceId: web.id,
 *     eipId: defaultEIP.id,
 * });
 * ```
 */
export class EIPAssociation extends pulumi.CustomResource {
    /**
     * Get an existing EIPAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EIPAssociationState, opts?: pulumi.CustomResourceOptions): EIPAssociation {
        return new EIPAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ucloud:unet/eIPAssociation:EIPAssociation';

    /**
     * Returns true if the given object is an instance of EIPAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EIPAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EIPAssociation.__pulumiType;
    }

    /**
     * The ID of EIP.
     */
    public readonly eipId!: pulumi.Output<string>;
    /**
     * The ID of resource with EIP attached.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * , attribute `resourceType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `resource_type` is deprecated for optimizing parameters
     */
    public readonly resourceType!: pulumi.Output<string>;

    /**
     * Create a EIPAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EIPAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EIPAssociationArgs | EIPAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EIPAssociationState | undefined;
            inputs["eipId"] = state ? state.eipId : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["resourceType"] = state ? state.resourceType : undefined;
        } else {
            const args = argsOrState as EIPAssociationArgs | undefined;
            if ((!args || args.eipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eipId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            inputs["eipId"] = args ? args.eipId : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EIPAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EIPAssociation resources.
 */
export interface EIPAssociationState {
    /**
     * The ID of EIP.
     */
    eipId?: pulumi.Input<string>;
    /**
     * The ID of resource with EIP attached.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * , attribute `resourceType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `resource_type` is deprecated for optimizing parameters
     */
    resourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EIPAssociation resource.
 */
export interface EIPAssociationArgs {
    /**
     * The ID of EIP.
     */
    eipId: pulumi.Input<string>;
    /**
     * The ID of resource with EIP attached.
     */
    resourceId: pulumi.Input<string>;
    /**
     * , attribute `resourceType` is deprecated for optimizing parameters.
     *
     * @deprecated attribute `resource_type` is deprecated for optimizing parameters
     */
    resourceType?: pulumi.Input<string>;
}
