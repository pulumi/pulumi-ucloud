// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ucloud.Ipsecvpn
{
    public static class GetVPNConnection
    {
        /// <summary>
        /// This data source providers a list of VPN Connection resources according to their ID, name and tag.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ucloud = Pulumi.Ucloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ucloud.Ipsecvpn.GetVPNConnection.InvokeAsync());
        ///         this.First = example.Apply(example =&gt; example.VpnConnections[0].Id);
        ///     }
        /// 
        ///     [Output("first")]
        ///     public Output&lt;string&gt; First { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVPNConnectionResult> InvokeAsync(GetVPNConnectionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVPNConnectionResult>("ucloud:ipsecvpn/getVPNConnection:getVPNConnection", args ?? new GetVPNConnectionArgs(), options.WithVersion());
    }


    public sealed class GetVPNConnectionArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of VPN Connection IDs, all the VPN Connections belongs to the defined region will be retrieved if this argument is `[]`.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting VPN Connections by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A tag assigned to VPN Connection.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public GetVPNConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVPNConnectionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// A tag assigned to the VPN Connection.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Total number of VPN Connections that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// It is a nested type. VPN Connections documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVPNConnectionVpnConnectionResult> VpnConnections;

        [OutputConstructor]
        private GetVPNConnectionResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? tag,

            int totalCount,

            ImmutableArray<Outputs.GetVPNConnectionVpnConnectionResult> vpnConnections)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Tag = tag;
            TotalCount = totalCount;
            VpnConnections = vpnConnections;
        }
    }
}
