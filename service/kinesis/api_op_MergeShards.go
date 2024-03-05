// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Merges two adjacent shards in a Kinesis data stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data. This
// API is only supported for the data streams with the provisioned capacity mode.
// Two shards are considered adjacent if the union of the hash key ranges for the
// two shards form a contiguous set with no gaps. For example, if you have two
// shards, one with a hash key range of 276...381 and the other with a hash key
// range of 382...454, then you could merge these two shards into a single shard
// that would have a hash key range of 276...454. After the merge, the single child
// shard receives data for all hash key values covered by the two parent shards.
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API. MergeShards is called when there is a need
// to reduce the overall capacity of a stream because of excess capacity that is
// not being used. You must specify the shard to be merged and the adjacent shard
// for a stream. For more information about merging shards, see Merge Two Shards (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-merge.html)
// in the Amazon Kinesis Data Streams Developer Guide. If the stream is in the
// ACTIVE state, you can call MergeShards . If a stream is in the CREATING ,
// UPDATING , or DELETING state, MergeShards returns a ResourceInUseException . If
// the specified stream does not exist, MergeShards returns a
// ResourceNotFoundException . You can use DescribeStreamSummary to check the
// state of the stream, which is returned in StreamStatus . MergeShards is an
// asynchronous operation. Upon receiving a MergeShards request, Amazon Kinesis
// Data Streams immediately returns a response and sets the StreamStatus to
// UPDATING . After the operation is completed, Kinesis Data Streams sets the
// StreamStatus to ACTIVE . Read and write operations continue to work while the
// stream is in the UPDATING state. You use DescribeStreamSummary and the
// ListShards APIs to determine the shard IDs that are specified in the MergeShards
// request. If you try to operate on too many streams in parallel using
// CreateStream , DeleteStream , MergeShards , or SplitShard , you receive a
// LimitExceededException . MergeShards has a limit of five transactions per
// second per account.
func (c *Client) MergeShards(ctx context.Context, params *MergeShardsInput, optFns ...func(*Options)) (*MergeShardsOutput, error) {
	if params == nil {
		params = &MergeShardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergeShards", params, optFns, c.addOperationMergeShardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergeShardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for MergeShards .
type MergeShardsInput struct {

	// The shard ID of the adjacent shard for the merge.
	//
	// This member is required.
	AdjacentShardToMerge *string

	// The shard ID of the shard to combine with the adjacent shard for the merge.
	//
	// This member is required.
	ShardToMerge *string

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream for the merge.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *MergeShardsInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type MergeShardsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMergeShardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&smithyRpcv2cbor_serializeOpMergeShards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&smithyRpcv2cbor_deserializeOpMergeShards{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "MergeShards"); err != nil {
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
	if err = addOpMergeShardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergeShards(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergeShards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "MergeShards",
	}
}
