// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ucloud.ipsecvpn
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source providers a list of VPN Customer Gateway resources according to their ID, name and tag.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/d/vpn_customer_gateways.html.markdown.
        /// </summary>
        public static Task<LookupVpnCustomerGatewaysResult> LookupVpnCustomerGateways(LookupVpnCustomerGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<LookupVpnCustomerGatewaysResult>("ucloud:ipsecvpn/lookupVpnCustomerGateways:lookupVpnCustomerGateways", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class LookupVpnCustomerGatewaysArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of VPN Customer Gateway IDs, all the VPN Customer Gateways belongs to the defined region will be retrieved if this argument is "".
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting VPN Customer Gateways by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A tag assigned to VPN Customer Gateway.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public LookupVpnCustomerGatewaysArgs()
        {
        }
    }

    [OutputType]
    public sealed class LookupVpnCustomerGatewaysResult
    {
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// A tag assigned to the VPN Customer Gateway.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// Total number of VPN Customer Gateways that satisfy the condition.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// It is a nested type. VPN Customer Gateways documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.LookupVpnCustomerGatewaysVpnCustomerGatewaysResult> VpnCustomerGateways;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LookupVpnCustomerGatewaysResult(
            ImmutableArray<string> ids,
            string? nameRegex,
            string? outputFile,
            string? tag,
            int totalCount,
            ImmutableArray<Outputs.LookupVpnCustomerGatewaysVpnCustomerGatewaysResult> vpnCustomerGateways,
            string id)
        {
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Tag = tag;
            TotalCount = totalCount;
            VpnCustomerGateways = vpnCustomerGateways;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LookupVpnCustomerGatewaysVpnCustomerGatewaysResult
    {
        /// <summary>
        /// The time of creation for VPN Customer Gateway, formatted in RFC3339 time string.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ID of VPN Customer Gateway.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ip address of the VPN Customer Gateway.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The name of the VPN Customer Gateway.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The remarks of VPN Customer Gateway.
        /// </summary>
        public readonly string Remark;
        /// <summary>
        /// A tag assigned to VPN Customer Gateway.
        /// </summary>
        public readonly string Tag;

        [OutputConstructor]
        private LookupVpnCustomerGatewaysVpnCustomerGatewaysResult(
            string createTime,
            string id,
            string ipAddress,
            string name,
            string remark,
            string tag)
        {
            CreateTime = createTime;
            Id = id;
            IpAddress = ipAddress;
            Name = name;
            Remark = remark;
            Tag = tag;
        }
    }
    }
}