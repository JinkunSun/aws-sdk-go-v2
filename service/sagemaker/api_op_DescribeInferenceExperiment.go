// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about an inference experiment.
func (c *Client) DescribeInferenceExperiment(ctx context.Context, params *DescribeInferenceExperimentInput, optFns ...func(*Options)) (*DescribeInferenceExperimentOutput, error) {
	if params == nil {
		params = &DescribeInferenceExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInferenceExperiment", params, optFns, c.addOperationDescribeInferenceExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInferenceExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInferenceExperimentInput struct {

	// The name of the inference experiment to describe.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeInferenceExperimentOutput struct {

	// The ARN of the inference experiment being described.
	//
	// This member is required.
	Arn *string

	// The metadata of the endpoint on which the inference experiment ran.
	//
	// This member is required.
	EndpointMetadata *types.EndpointMetadata

	// An array of ModelVariantConfigSummary objects. There is one for each variant in
	// the inference experiment. Each ModelVariantConfigSummary object in the array
	// describes the infrastructure configuration for deploying the corresponding
	// variant.
	//
	// This member is required.
	ModelVariants []types.ModelVariantConfigSummary

	// The name of the inference experiment.
	//
	// This member is required.
	Name *string

	// The status of the inference experiment. The following are the possible statuses
	// for an inference experiment:
	//   - Creating - Amazon SageMaker is creating your experiment.
	//   - Created - Amazon SageMaker has finished the creation of your experiment and
	//   will begin the experiment at the scheduled time.
	//   - Updating - When you make changes to your experiment, your experiment shows
	//   as updating.
	//   - Starting - Amazon SageMaker is beginning your experiment.
	//   - Running - Your experiment is in progress.
	//   - Stopping - Amazon SageMaker is stopping your experiment.
	//   - Completed - Your experiment has completed.
	//   - Cancelled - When you conclude your experiment early using the
	//   StopInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopInferenceExperiment.html)
	//   API, or if any operation fails with an unexpected error, it shows as cancelled.
	//
	// This member is required.
	Status types.InferenceExperimentStatus

	// The type of the inference experiment.
	//
	// This member is required.
	Type types.InferenceExperimentType

	// The timestamp at which the inference experiment was completed.
	CompletionTime *time.Time

	// The timestamp at which you created the inference experiment.
	CreationTime *time.Time

	// The Amazon S3 location and configuration for storing inference request and
	// response data.
	DataStorageConfig *types.InferenceExperimentDataStorageConfig

	// The description of the inference experiment.
	Description *string

	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key
	// that Amazon SageMaker uses to encrypt data on the storage volume attached to the
	// ML compute instance that hosts the endpoint. For more information, see
	// CreateInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateInferenceExperiment.html)
	// .
	KmsKey *string

	// The timestamp at which you last modified the inference experiment.
	LastModifiedTime *time.Time

	// The ARN of the IAM role that Amazon SageMaker can assume to access model
	// artifacts and container images, and manage Amazon SageMaker Inference endpoints
	// for model deployment.
	RoleArn *string

	// The duration for which the inference experiment ran or will run.
	Schedule *types.InferenceExperimentSchedule

	// The configuration of ShadowMode inference experiment type, which shows the
	// production variant that takes all the inference requests, and the shadow variant
	// to which Amazon SageMaker replicates a percentage of the inference requests. For
	// the shadow variant it also shows the percentage of requests that Amazon
	// SageMaker replicates.
	ShadowModeConfig *types.ShadowModeConfig

	// The error message or client-specified Reason from the StopInferenceExperiment (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopInferenceExperiment.html)
	// API, that explains the status of the inference experiment.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInferenceExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInferenceExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInferenceExperiment{}, middleware.After)
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
	if err = addOpDescribeInferenceExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInferenceExperiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInferenceExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeInferenceExperiment",
	}
}
