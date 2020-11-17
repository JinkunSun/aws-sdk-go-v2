// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the versioning state of an existing bucket. To set the versioning state,
// you must be the bucket owner. You can set the versioning state with one of the
// following values: Enabled—Enables versioning for the objects in the bucket. All
// objects added to the bucket receive a unique version ID. Suspended—Disables
// versioning for the objects in the bucket. All objects added to the bucket
// receive the version ID null. If the versioning state has never been set on a
// bucket, it has no versioning state; a GetBucketVersioning
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html)
// request does not return a versioning state value. If the bucket owner enables
// MFA Delete in the bucket versioning configuration, the bucket owner must include
// the x-amz-mfa request header and the Status and the MfaDelete request elements
// in a request to set the versioning state of the bucket. If you have an object
// expiration lifecycle policy in your non-versioned bucket and you want to
// maintain the same permanent delete behavior when you enable versioning, you must
// add a noncurrent expiration policy. The noncurrent expiration lifecycle policy
// will manage the deletes of the noncurrent object versions in the version-enabled
// bucket. (A version-enabled bucket maintains one current and zero or more
// noncurrent object versions.) For more information, see Lifecycle and Versioning
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config).
// Related Resources
//
// * CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//
// *
// DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// *
// GetBucketVersioning
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html)
func (c *Client) PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error) {
	if params == nil {
		params = &PutBucketVersioningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketVersioning", params, optFns, addOperationPutBucketVersioningMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketVersioningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketVersioningInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for setting the versioning state.
	//
	// This member is required.
	VersioningConfiguration *types.VersioningConfiguration

	// >The base64-encoded 128-bit MD5 digest of the data. You must use this header as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// The concatenation of the authentication device's serial number, a space, and the
	// value that is displayed on your authentication device.
	MFA *string
}

type PutBucketVersioningOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketVersioningMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketVersioning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketVersioning{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpPutBucketVersioningValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketVersioning(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketVersioningUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketVersioning(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketVersioning",
	}
}

// getPutBucketVersioningBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getPutBucketVersioningBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketVersioningInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketVersioningUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketVersioningBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
