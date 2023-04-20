// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of the specified IoT policy. To update a policy, create a
// new policy version. A managed policy can have up to five versions. If the policy
// has five versions, you must use DeletePolicyVersion to delete an existing
// version before you create a new one. Optionally, you can set the new version as
// the policy's default version. The default version is the operative version (that
// is, the version that is in effect for the certificates to which the policy is
// attached). Requires permission to access the CreatePolicyVersion (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreatePolicyVersion(ctx context.Context, params *CreatePolicyVersionInput, optFns ...func(*Options)) (*CreatePolicyVersionOutput, error) {
	if params == nil {
		params = &CreatePolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicyVersion", params, optFns, c.addOperationCreatePolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreatePolicyVersion operation.
type CreatePolicyVersionInput struct {

	// The JSON document that describes the policy. Minimum length of 1. Maximum
	// length of 2048, excluding whitespace.
	//
	// This member is required.
	PolicyDocument *string

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// Specifies whether the policy version is set as the default. When this parameter
	// is true, the new policy version becomes the operative version (that is, the
	// version that is in effect for the certificates to which the policy is attached).
	SetAsDefault bool

	noSmithyDocumentSerde
}

// The output of the CreatePolicyVersion operation.
type CreatePolicyVersionOutput struct {

	// Specifies whether the policy version is the default.
	IsDefaultVersion bool

	// The policy ARN.
	PolicyArn *string

	// The JSON document that describes the policy.
	PolicyDocument *string

	// The policy version ID.
	PolicyVersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePolicyVersion{}, middleware.After)
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
	if err = addOpCreatePolicyVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicyVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePolicyVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "CreatePolicyVersion",
	}
}
