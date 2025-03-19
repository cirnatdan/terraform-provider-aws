// Code generated by "internal/generate/listpages/main.go -ListOps=GetApis,GetApiMappings,GetDomainNames,GetIntegrations,GetRoutes,GetStages,GetVpcLinks"; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func getAPIMappingsPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetApiMappingsInput, fn func(*apigatewayv2.GetApiMappingsOutput, bool) bool) error {
	for {
		output, err := conn.GetApiMappings(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getAPIsPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetApisInput, fn func(*apigatewayv2.GetApisOutput, bool) bool) error {
	for {
		output, err := conn.GetApis(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getDomainNamesPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetDomainNamesInput, fn func(*apigatewayv2.GetDomainNamesOutput, bool) bool) error {
	for {
		output, err := conn.GetDomainNames(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getIntegrationsPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetIntegrationsInput, fn func(*apigatewayv2.GetIntegrationsOutput, bool) bool) error {
	for {
		output, err := conn.GetIntegrations(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getRoutesPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetRoutesInput, fn func(*apigatewayv2.GetRoutesOutput, bool) bool) error {
	for {
		output, err := conn.GetRoutes(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getStagesPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetStagesInput, fn func(*apigatewayv2.GetStagesOutput, bool) bool) error {
	for {
		output, err := conn.GetStages(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getVPCLinksPages(ctx context.Context, conn *apigatewayv2.Client, input *apigatewayv2.GetVpcLinksInput, fn func(*apigatewayv2.GetVpcLinksOutput, bool) bool) error {
	for {
		output, err := conn.GetVpcLinks(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
