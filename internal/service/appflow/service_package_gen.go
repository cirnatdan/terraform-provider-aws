// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appflow

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	appflow_sdkv2 "github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceConnectorProfile,
			TypeName: "aws_appflow_connector_profile",
			Name:     "Connector Profile",
		},
		{
			Factory:  resourceFlow,
			TypeName: "aws_appflow_flow",
			Name:     "Flow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppFlow
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appflow_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return appflow_sdkv2.NewFromConfig(cfg, func(o *appflow_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
