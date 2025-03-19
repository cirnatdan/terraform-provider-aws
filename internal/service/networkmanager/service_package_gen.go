// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newDirectConnectGatewayAttachmentResource,
			TypeName: "aws_networkmanager_dx_gateway_attachment",
			Name:     "Direct Connect Gateway Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceConnection,
			TypeName: "aws_networkmanager_connection",
			Name:     "Connection",
		},
		{
			Factory:  dataSourceConnections,
			TypeName: "aws_networkmanager_connections",
			Name:     "Connections",
		},
		{
			Factory:  dataSourceCoreNetworkPolicyDocument,
			TypeName: "aws_networkmanager_core_network_policy_document",
			Name:     "Core Network Policy Document",
		},
		{
			Factory:  dataSourceDevice,
			TypeName: "aws_networkmanager_device",
			Name:     "Device",
		},
		{
			Factory:  dataSourceDevices,
			TypeName: "aws_networkmanager_devices",
			Name:     "Devices",
		},
		{
			Factory:  dataSourceGlobalNetwork,
			TypeName: "aws_networkmanager_global_network",
			Name:     "Global Network",
		},
		{
			Factory:  dataSourceGlobalNetworks,
			TypeName: "aws_networkmanager_global_networks",
			Name:     "Global Networks",
		},
		{
			Factory:  dataSourceLink,
			TypeName: "aws_networkmanager_link",
			Name:     "Link",
		},
		{
			Factory:  dataSourceLinks,
			TypeName: "aws_networkmanager_links",
			Name:     "Links",
		},
		{
			Factory:  dataSourceSite,
			TypeName: "aws_networkmanager_site",
			Name:     "Site",
		},
		{
			Factory:  dataSourceSites,
			TypeName: "aws_networkmanager_sites",
			Name:     "Sites",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAttachmentAccepter,
			TypeName: "aws_networkmanager_attachment_accepter",
			Name:     "Attachment Accepter",
		},
		{
			Factory:  resourceConnectAttachment,
			TypeName: "aws_networkmanager_connect_attachment",
			Name:     "Connect Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnectPeer,
			TypeName: "aws_networkmanager_connect_peer",
			Name:     "Connect Peer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_networkmanager_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCoreNetwork,
			TypeName: "aws_networkmanager_core_network",
			Name:     "Core Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCoreNetworkPolicyAttachment,
			TypeName: "aws_networkmanager_core_network_policy_attachment",
			Name:     "Core Network Policy Attachment",
		},
		{
			Factory:  resourceCustomerGatewayAssociation,
			TypeName: "aws_networkmanager_customer_gateway_association",
			Name:     "Customer Gateway Association",
		},
		{
			Factory:  resourceDevice,
			TypeName: "aws_networkmanager_device",
			Name:     "Device",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalNetwork,
			TypeName: "aws_networkmanager_global_network",
			Name:     "Global Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceLink,
			TypeName: "aws_networkmanager_link",
			Name:     "Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceLinkAssociation,
			TypeName: "aws_networkmanager_link_association",
			Name:     "Link Association",
		},
		{
			Factory:  resourceSite,
			TypeName: "aws_networkmanager_site",
			Name:     "Site",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSiteToSiteVPNAttachment,
			TypeName: "aws_networkmanager_site_to_site_vpn_attachment",
			Name:     "Site To Site VPN Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTransitGatewayConnectPeerAssociation,
			TypeName: "aws_networkmanager_transit_gateway_connect_peer_association",
			Name:     "Transit Gateway Connect Peer Association",
		},
		{
			Factory:  resourceTransitGatewayPeering,
			TypeName: "aws_networkmanager_transit_gateway_peering",
			Name:     "Transit Gateway Peering",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTransitGatewayRegistration,
			TypeName: "aws_networkmanager_transit_gateway_registration",
			Name:     "Transit Gateway Registration",
		},
		{
			Factory:  resourceTransitGatewayRouteTableAttachment,
			TypeName: "aws_networkmanager_transit_gateway_route_table_attachment",
			Name:     "Transit Gateway Route Table Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceVPCAttachment,
			TypeName: "aws_networkmanager_vpc_attachment",
			Name:     "VPC Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.NetworkManager
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*networkmanager.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*networkmanager.Options){
		networkmanager.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return networkmanager.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*networkmanager.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*networkmanager.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *networkmanager.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*networkmanager.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
