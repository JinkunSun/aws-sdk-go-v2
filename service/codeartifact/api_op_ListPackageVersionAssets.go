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

// Returns a list of AssetSummary (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html)
// objects for assets in a package version.
func (c *Client) ListPackageVersionAssets(ctx context.Context, params *ListPackageVersionAssetsInput, optFns ...func(*Options)) (*ListPackageVersionAssetsOutput, error) {
	if params == nil {
		params = &ListPackageVersionAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageVersionAssets", params, optFns, c.addOperationListPackageVersionAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageVersionAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionAssetsInput struct {

	// The name of the domain that contains the repository associated with the package
	// version assets.
	//
	// This member is required.
	Domain *string

	// The format of the package that contains the requested package version assets.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package that contains the requested package version assets.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2 ).
	//
	// This member is required.
	PackageVersion *string

	// The name of the repository that contains the package that contains the
	// requested package version assets.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The namespace of the package version that contains the requested package
	// version assets. The package component that specifies its namespace depends on
	// its type. For example: The namespace is required requesting assets from package
	// versions of the following formats:
	//   - Maven
	//   - Swift
	//   - generic
	//
	//   - The namespace of a Maven package version is its groupId .
	//   - The namespace of an npm or Swift package version is its scope .
	//   - The namespace of a generic package is its namespace .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPackageVersionAssetsOutput struct {

	// The returned list of AssetSummary (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html)
	// objects.
	Assets []types.AssetSummary

	// The format of the package that contains the requested package version assets.
	Format types.PackageFormat

	// The namespace of the package version that contains the requested package
	// version assets. The package component that specifies its namespace depends on
	// its type. For example:
	//   - The namespace of a Maven package version is its groupId .
	//   - The namespace of an npm or Swift package version is its scope .
	//   - The namespace of a generic package is its namespace .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	Namespace *string

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The name of the package that contains the requested package version assets.
	Package *string

	// The version of the package associated with the requested assets.
	Version *string

	// The current revision associated with the package version.
	VersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackageVersionAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersionAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersionAssets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPackageVersionAssets"); err != nil {
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
	if err = addOpListPackageVersionAssetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersionAssets(options.Region), middleware.Before); err != nil {
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

// ListPackageVersionAssetsAPIClient is a client that implements the
// ListPackageVersionAssets operation.
type ListPackageVersionAssetsAPIClient interface {
	ListPackageVersionAssets(context.Context, *ListPackageVersionAssetsInput, ...func(*Options)) (*ListPackageVersionAssetsOutput, error)
}

var _ ListPackageVersionAssetsAPIClient = (*Client)(nil)

// ListPackageVersionAssetsPaginatorOptions is the paginator options for
// ListPackageVersionAssets
type ListPackageVersionAssetsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackageVersionAssetsPaginator is a paginator for ListPackageVersionAssets
type ListPackageVersionAssetsPaginator struct {
	options   ListPackageVersionAssetsPaginatorOptions
	client    ListPackageVersionAssetsAPIClient
	params    *ListPackageVersionAssetsInput
	nextToken *string
	firstPage bool
}

// NewListPackageVersionAssetsPaginator returns a new
// ListPackageVersionAssetsPaginator
func NewListPackageVersionAssetsPaginator(client ListPackageVersionAssetsAPIClient, params *ListPackageVersionAssetsInput, optFns ...func(*ListPackageVersionAssetsPaginatorOptions)) *ListPackageVersionAssetsPaginator {
	if params == nil {
		params = &ListPackageVersionAssetsInput{}
	}

	options := ListPackageVersionAssetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackageVersionAssetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackageVersionAssetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackageVersionAssets page.
func (p *ListPackageVersionAssetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackageVersionAssetsOutput, error) {
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

	result, err := p.client.ListPackageVersionAssets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackageVersionAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPackageVersionAssets",
	}
}
