// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the repositories in the added repositories list of the specified
// restriction type for a package group. For more information about restriction
// types and added repository lists, see Package group origin controls (https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-origin-controls.html)
// in the CodeArtifact User Guide.
func (c *Client) ListAllowedRepositoriesForGroup(ctx context.Context, params *ListAllowedRepositoriesForGroupInput, optFns ...func(*Options)) (*ListAllowedRepositoriesForGroupOutput, error) {
	if params == nil {
		params = &ListAllowedRepositoriesForGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAllowedRepositoriesForGroup", params, optFns, c.addOperationListAllowedRepositoriesForGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAllowedRepositoriesForGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAllowedRepositoriesForGroupInput struct {

	// The name of the domain that contains the package group from which to list
	// allowed repositories.
	//
	// This member is required.
	Domain *string

	// The origin configuration restriction type of which to list allowed repositories.
	//
	// This member is required.
	OriginRestrictionType types.PackageGroupOriginRestrictionType

	// The pattern of the package group from which to list allowed repositories.
	//
	// This member is required.
	PackageGroup *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAllowedRepositoriesForGroupOutput struct {

	// The list of allowed repositories for the package group and origin configuration
	// restriction type.
	AllowedRepositories []string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAllowedRepositoriesForGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAllowedRepositoriesForGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAllowedRepositoriesForGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAllowedRepositoriesForGroup"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAllowedRepositoriesForGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAllowedRepositoriesForGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListAllowedRepositoriesForGroupAPIClient is a client that implements the
// ListAllowedRepositoriesForGroup operation.
type ListAllowedRepositoriesForGroupAPIClient interface {
	ListAllowedRepositoriesForGroup(context.Context, *ListAllowedRepositoriesForGroupInput, ...func(*Options)) (*ListAllowedRepositoriesForGroupOutput, error)
}

var _ ListAllowedRepositoriesForGroupAPIClient = (*Client)(nil)

// ListAllowedRepositoriesForGroupPaginatorOptions is the paginator options for
// ListAllowedRepositoriesForGroup
type ListAllowedRepositoriesForGroupPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAllowedRepositoriesForGroupPaginator is a paginator for
// ListAllowedRepositoriesForGroup
type ListAllowedRepositoriesForGroupPaginator struct {
	options   ListAllowedRepositoriesForGroupPaginatorOptions
	client    ListAllowedRepositoriesForGroupAPIClient
	params    *ListAllowedRepositoriesForGroupInput
	nextToken *string
	firstPage bool
}

// NewListAllowedRepositoriesForGroupPaginator returns a new
// ListAllowedRepositoriesForGroupPaginator
func NewListAllowedRepositoriesForGroupPaginator(client ListAllowedRepositoriesForGroupAPIClient, params *ListAllowedRepositoriesForGroupInput, optFns ...func(*ListAllowedRepositoriesForGroupPaginatorOptions)) *ListAllowedRepositoriesForGroupPaginator {
	if params == nil {
		params = &ListAllowedRepositoriesForGroupInput{}
	}

	options := ListAllowedRepositoriesForGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAllowedRepositoriesForGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAllowedRepositoriesForGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAllowedRepositoriesForGroup page.
func (p *ListAllowedRepositoriesForGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAllowedRepositoriesForGroupOutput, error) {
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

	result, err := p.client.ListAllowedRepositoriesForGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAllowedRepositoriesForGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAllowedRepositoriesForGroup",
	}
}
