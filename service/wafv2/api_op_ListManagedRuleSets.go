// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the managed rule sets that you own. This is intended for use only by
// vendors of managed rule sets. Vendors are Amazon Web Services and Amazon Web
// Services Marketplace sellers. Vendors, you can use the managed rule set APIs to
// provide controlled rollout of your versioned managed rule group offerings for
// your customers. The APIs are ListManagedRuleSets , GetManagedRuleSet ,
// PutManagedRuleSetVersions , and UpdateManagedRuleSetVersionExpiryDate .
func (c *Client) ListManagedRuleSets(ctx context.Context, params *ListManagedRuleSetsInput, optFns ...func(*Options)) (*ListManagedRuleSetsOutput, error) {
	if params == nil {
		params = &ListManagedRuleSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedRuleSets", params, optFns, c.addOperationListManagedRuleSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedRuleSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedRuleSetsInput struct {

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, or an App Runner service. To work with CloudFront, you must also
	// specify the Region US East (N. Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The maximum number of objects that you want WAF to return for this request. If
	// more objects are available, in the response, WAF provides a NextMarker value
	// that you can use in a subsequent call to get the next batch of objects.
	Limit *int32

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	noSmithyDocumentSerde
}

type ListManagedRuleSetsOutput struct {

	// Your managed rule sets. If you specified a Limit in your request, this might
	// not be the full list.
	ManagedRuleSets []types.ManagedRuleSetSummary

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedRuleSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListManagedRuleSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListManagedRuleSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListManagedRuleSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedRuleSets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListManagedRuleSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "ListManagedRuleSets",
	}
}
