// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transfers the specified certificate to the specified Amazon Web Services
// account. Requires permission to access the TransferCertificate (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action. You can cancel the transfer until it is acknowledged by the recipient.
// No notification is sent to the transfer destination's account. It is up to the
// caller to notify the transfer target. The certificate being transferred must not
// be in the ACTIVE state. You can use the UpdateCertificate action to deactivate
// it. The certificate must not have any policies attached to it. You can use the
// DetachPolicy action to detach them.
func (c *Client) TransferCertificate(ctx context.Context, params *TransferCertificateInput, optFns ...func(*Options)) (*TransferCertificateOutput, error) {
	if params == nil {
		params = &TransferCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TransferCertificate", params, optFns, c.addOperationTransferCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TransferCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the TransferCertificate operation.
type TransferCertificateInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string

	// The Amazon Web Services account.
	//
	// This member is required.
	TargetAwsAccount *string

	// The transfer message.
	TransferMessage *string

	noSmithyDocumentSerde
}

// The output from the TransferCertificate operation.
type TransferCertificateOutput struct {

	// The ARN of the certificate.
	TransferredCertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTransferCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTransferCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTransferCertificate{}, middleware.After)
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
	if err = addOpTransferCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTransferCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTransferCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot",
		OperationName: "TransferCertificate",
	}
}
