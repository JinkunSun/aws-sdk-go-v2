// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// For asynchronous function invocation, use Invoke . Invokes a function
// asynchronously. If you do use the InvokeAsync action, note that it doesn't
// support the use of X-Ray active tracing. Trace ID is not propagated to the
// function, even if X-Ray active tracing is turned on.
//
// Deprecated: This operation has been deprecated.
func (c *Client) InvokeAsync(ctx context.Context, params *InvokeAsyncInput, optFns ...func(*Options)) (*InvokeAsyncOutput, error) {
	if params == nil {
		params = &InvokeAsyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeAsync", params, optFns, c.addOperationInvokeAsyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeAsyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeAsyncInput struct {

	// The name or ARN of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The JSON that you want to provide to your Lambda function as input.
	//
	// This member is required.
	InvokeArgs io.Reader

	noSmithyDocumentSerde
}

// A success response ( 202 Accepted ) indicates that the request is queued for
// invocation.
type InvokeAsyncOutput struct {

	// The status code.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeAsyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeAsync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeAsync{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeAsync"); err != nil {
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
	if err = addOpInvokeAsyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeAsync(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeAsync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeAsync",
	}
}
