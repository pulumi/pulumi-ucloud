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
    /// Provides an VPC Peering Connection for establishing a connection between multiple VPC.
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
    ///         var foo = new Ucloud.Vpc.VPC("foo", new Ucloud.Vpc.VPCArgs
    ///         {
    ///             Tag = "tf-example",
    ///             CidrBlocks = 
    ///             {
    ///                 "192.168.0.0/16",
    ///             },
    ///         });
    ///         var bar = new Ucloud.Vpc.VPC("bar", new Ucloud.Vpc.VPCArgs
    ///         {
    ///             Tag = "tf-example",
    ///             CidrBlocks = 
    ///             {
    ///                 "10.10.0.0/16",
    ///             },
    ///         });
    ///         var connection = new Ucloud.Vpc.VPCPeeringConnection("connection", new Ucloud.Vpc.VPCPeeringConnectionArgs
    ///         {
    ///             VpcId = foo.Id,
    ///             PeerVpcId = bar.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [UcloudResourceType("ucloud:vpc/vPCPeeringConnection:VPCPeeringConnection")]
    public partial class VPCPeeringConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of accepter project of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Output("peerProjectId")]
        public Output<string> PeerProjectId { get; private set; } = null!;

        /// <summary>
        /// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Output("peerVpcId")]
        public Output<string> PeerVpcId { get; private set; } = null!;

        /// <summary>
        /// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VPCPeeringConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPCPeeringConnection(string name, VPCPeeringConnectionArgs args, CustomResourceOptions? options = null)
            : base("ucloud:vpc/vPCPeeringConnection:VPCPeeringConnection", name, args ?? new VPCPeeringConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPCPeeringConnection(string name, Input<string> id, VPCPeeringConnectionState? state = null, CustomResourceOptions? options = null)
            : base("ucloud:vpc/vPCPeeringConnection:VPCPeeringConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VPCPeeringConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPCPeeringConnection Get(string name, Input<string> id, VPCPeeringConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VPCPeeringConnection(name, id, state, options);
        }
    }

    public sealed class VPCPeeringConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of accepter project of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerProjectId")]
        public Input<string>? PeerProjectId { get; set; }

        /// <summary>
        /// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerVpcId", required: true)]
        public Input<string> PeerVpcId { get; set; } = null!;

        /// <summary>
        /// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VPCPeeringConnectionArgs()
        {
        }
    }

    public sealed class VPCPeeringConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of accepter project of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerProjectId")]
        public Input<string>? PeerProjectId { get; set; }

        /// <summary>
        /// The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerVpcId")]
        public Input<string>? PeerVpcId { get; set; }

        /// <summary>
        /// The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VPCPeeringConnectionState()
        {
        }
    }
}