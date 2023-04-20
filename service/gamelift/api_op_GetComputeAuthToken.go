// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Requests an authentication token from Amazon GameLift. The authentication token
// is used by your game server to authenticate with Amazon GameLift. Each
// authentication token has an expiration time. To continue using the compute
// resource to host your game server, regularly retrieve a new authorization token.
func (c *Client) GetComputeAuthToken(ctx context.Context, params *GetComputeAuthTokenInput, optFns ...func(*Options)) (*GetComputeAuthTokenOutput, error) {
	if params == nil {
		params = &GetComputeAuthTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComputeAuthToken", params, optFns, c.addOperationGetComputeAuthTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComputeAuthTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComputeAuthTokenInput struct {

	// The name of the compute resource you are requesting the authentication token
	// for.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet that the compute is registered to.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type GetComputeAuthTokenOutput struct {

	// The authentication token that your game server uses to authenticate with Amazon
	// GameLift.
	AuthToken *string

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift compute resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::compute/compute-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912
	ComputeArn *string

	// The name of the compute resource you are requesting the authentication token
	// for.
	ComputeName *string

	// The amount of time until the authentication token is no longer valid. To
	// continue using the compute resource for game server hosting, renew the
	// authentication token by using this operation again.
	ExpirationTimestamp *time.Time

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift fleet resource and uniquely identifies
	// it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	FleetArn *string

	// A unique identifier for the fleet that the compute is registered to.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComputeAuthTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComputeAuthToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComputeAuthToken{}, middleware.After)
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
	if err = addOpGetComputeAuthTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComputeAuthToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetComputeAuthToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "GetComputeAuthToken",
	}
}
