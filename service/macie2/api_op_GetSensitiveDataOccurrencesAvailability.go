// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks whether occurrences of sensitive data can be retrieved (revealed) for a
// finding.
func (c *Client) GetSensitiveDataOccurrencesAvailability(ctx context.Context, params *GetSensitiveDataOccurrencesAvailabilityInput, optFns ...func(*Options)) (*GetSensitiveDataOccurrencesAvailabilityOutput, error) {
	if params == nil {
		params = &GetSensitiveDataOccurrencesAvailabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSensitiveDataOccurrencesAvailability", params, optFns, c.addOperationGetSensitiveDataOccurrencesAvailabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSensitiveDataOccurrencesAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSensitiveDataOccurrencesAvailabilityInput struct {

	// The unique identifier for the finding.
	//
	// This member is required.
	FindingId *string

	noSmithyDocumentSerde
}

type GetSensitiveDataOccurrencesAvailabilityOutput struct {

	// Specifies whether occurrences of sensitive data can be retrieved for the
	// finding. Possible values are: AVAILABLE, the sensitive data can be retrieved;
	// and, UNAVAILABLE, the sensitive data can't be retrieved. If this value is
	// UNAVAILABLE, the reasons array indicates why the data can't be retrieved.
	Code types.AvailabilityCode

	// Specifies why occurrences of sensitive data can't be retrieved for the finding.
	// Possible values are:
	//
	// * INVALID_CLASSIFICATION_RESULT - Amazon Macie can't
	// verify the location of the sensitive data to retrieve. There isn't a
	// corresponding sensitive data discovery result for the finding. Or the sensitive
	// data discovery result specified by the
	// ClassificationDetails.detailedResultsLocation field of the finding isn't
	// available, is malformed or corrupted, or uses an unsupported storage format.
	//
	// *
	// OBJECT_EXCEEDS_SIZE_QUOTA - The storage size of the affected S3 object exceeds
	// the size quota for retrieving occurrences of sensitive data.
	//
	// *
	// OBJECT_UNAVAILABLE - The affected S3 object isn't available. The object might
	// have been renamed, moved, or deleted. Or the object was changed after Amazon
	// Macie created the finding.
	//
	// * UNSUPPORTED_FINDING_TYPE - The specified finding
	// isn't a sensitive data finding.
	//
	// * UNSUPPORTED_OBJECT_TYPE - The affected S3
	// object uses a file or storage format that Macie doesn't support for retrieving
	// occurrences of sensitive data.
	//
	// This value is null if sensitive data can be
	// retrieved for the finding.
	Reasons []types.UnavailabilityReasonCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSensitiveDataOccurrencesAvailabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSensitiveDataOccurrencesAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSensitiveDataOccurrencesAvailability{}, middleware.After)
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
	if err = addOpGetSensitiveDataOccurrencesAvailabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSensitiveDataOccurrencesAvailability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSensitiveDataOccurrencesAvailability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetSensitiveDataOccurrencesAvailability",
	}
}
