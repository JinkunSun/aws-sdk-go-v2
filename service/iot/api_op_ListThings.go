// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your things. Use the attributeName and attributeValue parameters to
// filter your things. For example, calling ListThings with attributeName=Color
// and attributeValue=Red retrieves all things in the registry that contain an
// attribute Color with the value Red. For more information, see List Things (https://docs.aws.amazon.com/iot/latest/developerguide/thing-registry.html#list-things)
// from the Amazon Web Services IoT Core Developer Guide. Requires permission to
// access the ListThings (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action. You will not be charged for calling this API if an Access denied error
// is returned. You will also not be charged if no attributes or pagination token
// was provided in request and no pagination token and no results were returned.
func (c *Client) ListThings(ctx context.Context, params *ListThingsInput, optFns ...func(*Options)) (*ListThingsOutput, error) {
	if params == nil {
		params = &ListThingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThings", params, optFns, c.addOperationListThingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThings operation.
type ListThingsInput struct {

	// The attribute name used to search for things.
	AttributeName *string

	// The attribute value used to search for things.
	AttributeValue *string

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// The name of the thing type used to search for things.
	ThingTypeName *string

	// When true , the action returns the thing resources with attribute values that
	// start with the attributeValue provided. When false , or not present, the action
	// returns only the thing resources with attribute values that match the entire
	// attributeValue provided.
	UsePrefixAttributeValue bool

	noSmithyDocumentSerde
}

// The output from the ListThings operation.
type ListThingsOutput struct {

	// The token to use to get the next set of results. Will not be returned if
	// operation has returned all results.
	NextToken *string

	// The things.
	Things []types.ThingAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThings(options.Region), middleware.Before); err != nil {
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

// ListThingsAPIClient is a client that implements the ListThings operation.
type ListThingsAPIClient interface {
	ListThings(context.Context, *ListThingsInput, ...func(*Options)) (*ListThingsOutput, error)
}

var _ ListThingsAPIClient = (*Client)(nil)

// ListThingsPaginatorOptions is the paginator options for ListThings
type ListThingsPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingsPaginator is a paginator for ListThings
type ListThingsPaginator struct {
	options   ListThingsPaginatorOptions
	client    ListThingsAPIClient
	params    *ListThingsInput
	nextToken *string
	firstPage bool
}

// NewListThingsPaginator returns a new ListThingsPaginator
func NewListThingsPaginator(client ListThingsAPIClient, params *ListThingsInput, optFns ...func(*ListThingsPaginatorOptions)) *ListThingsPaginator {
	if params == nil {
		params = &ListThingsInput{}
	}

	options := ListThingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThings page.
func (p *ListThingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListThings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListThings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "ListThings",
	}
}
