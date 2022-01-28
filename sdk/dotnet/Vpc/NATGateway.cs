// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Vpc
{
    /// <summary>
    /// Provides a Nat Gateway resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Ucloud = Pulumi.Ucloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooVPC = new Ucloud.Vpc.VPC("fooVPC", new Ucloud.Vpc.VPCArgs
    ///         {
    ///             Tag = "tf-acc",
    ///             CidrBlocks = 
    ///             {
    ///                 "192.168.0.0/16",
    ///             },
    ///         });
    ///         var fooSubnet = new Ucloud.Vpc.Subnet("fooSubnet", new Ucloud.Vpc.SubnetArgs
    ///         {
    ///             Tag = "tf-acc",
    ///             CidrBlock = "192.168.1.0/24",
    ///             VpcId = fooVPC.Id,
    ///         });
    ///         var fooEIP = new Ucloud.Unet.EIP("fooEIP", new Ucloud.Unet.EIPArgs
    ///         {
    ///             Bandwidth = 1,
    ///             InternetType = "bgp",
    ///             ChargeMode = "bandwidth",
    ///             Tag = "tf-acc",
    ///         });
    ///         var fooSecurityGroup = Output.Create(Ucloud.Unet.GetSecurityGroup.InvokeAsync(new Ucloud.Unet.GetSecurityGroupArgs
    ///         {
    ///             Type = "recommend_web",
    ///         }));
    ///         var fooNATGateway = new Ucloud.Vpc.NATGateway("fooNATGateway", new Ucloud.Vpc.NATGatewayArgs
    ///         {
    ///             VpcId = fooVPC.Id,
    ///             SubnetIds = 
    ///             {
    ///                 fooSubnet.Id,
    ///             },
    ///             EipId = fooEIP.Id,
    ///             Tag = "tf-acc",
    ///             SecurityGroup = fooSecurityGroup.Apply(fooSecurityGroup =&gt; fooSecurityGroup.SecurityGroups[0].Id),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Nat Gateway can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import ucloud:vpc/nATGateway:NATGateway example natgw-abc123456
    /// ```
    /// </summary>
    [UcloudResourceType("ucloud:vpc/nATGateway:NATGateway")]
    public partial class NATGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The time of creation of Nat Gateway, formatted in RFC3339 time string.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of eip associate to the Nat Gateway.
        /// </summary>
        [Output("eipId")]
        public Output<string> EipId { get; private set; } = null!;

        /// <summary>
        /// The boolean value to Controls whether or not start the whitelist mode.
        /// </summary>
        [Output("enableWhiteList")]
        public Output<bool> EnableWhiteList { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The remarks of the Nat Gateway. (Default: `""`).
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Output("securityGroup")]
        public Output<string> SecurityGroup { get; private set; } = null!;

        /// <summary>
        /// The list of subnet ID under the VPC.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// * ``
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC linked to the Nat Gateway.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The white list of instance under the Nat Gateway.
        /// </summary>
        [Output("whiteLists")]
        public Output<ImmutableArray<string>> WhiteLists { get; private set; } = null!;


        /// <summary>
        /// Create a NATGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NATGateway(string name, NATGatewayArgs args, CustomResourceOptions? options = null)
            : base("ucloud:vpc/nATGateway:NATGateway", name, args ?? new NATGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NATGateway(string name, Input<string> id, NATGatewayState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:vpc/nATGateway:NATGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NATGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NATGateway Get(string name, Input<string> id, NATGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new NATGateway(name, id, state, options);
        }
    }

    public sealed class NATGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of eip associate to the Nat Gateway.
        /// </summary>
        [Input("eipId", required: true)]
        public Input<string> EipId { get; set; } = null!;

        /// <summary>
        /// The boolean value to Controls whether or not start the whitelist mode.
        /// </summary>
        [Input("enableWhiteList", required: true)]
        public Input<bool> EnableWhiteList { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The remarks of the Nat Gateway. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Input("securityGroup", required: true)]
        public Input<string> SecurityGroup { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The list of subnet ID under the VPC.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// * ``
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The ID of VPC linked to the Nat Gateway.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        [Input("whiteLists")]
        private InputList<string>? _whiteLists;

        /// <summary>
        /// The white list of instance under the Nat Gateway.
        /// </summary>
        public InputList<string> WhiteLists
        {
            get => _whiteLists ?? (_whiteLists = new InputList<string>());
            set => _whiteLists = value;
        }

        public NATGatewayArgs()
        {
        }
    }

    public sealed class NATGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time of creation of Nat Gateway, formatted in RFC3339 time string.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The ID of eip associate to the Nat Gateway.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The boolean value to Controls whether or not start the whitelist mode.
        /// </summary>
        [Input("enableWhiteList")]
        public Input<bool>? EnableWhiteList { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The remarks of the Nat Gateway. (Default: `""`).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The ID of the associated security group.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The list of subnet ID under the VPC.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// A tag assigned to Nat Gateway, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: `Default`).
        /// * ``
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The ID of VPC linked to the Nat Gateway.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("whiteLists")]
        private InputList<string>? _whiteLists;

        /// <summary>
        /// The white list of instance under the Nat Gateway.
        /// </summary>
        public InputList<string> WhiteLists
        {
            get => _whiteLists ?? (_whiteLists = new InputList<string>());
            set => _whiteLists = value;
        }

        public NATGatewayState()
        {
        }
    }
}