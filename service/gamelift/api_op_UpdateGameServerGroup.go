// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the Amazon GameLift FleetIQ solution and game
// server groups. Updates Amazon GameLift FleetIQ-specific properties for a game
// server group. Many Auto Scaling group properties are updated on the Auto Scaling
// group directly, including the launch template, Auto Scaling policies, and
// maximum/minimum/desired instance counts. To update the game server group,
// specify the game server group ID and provide the updated values. Before applying
// the updates, the new values are validated to ensure that Amazon GameLift FleetIQ
// can continue to perform instance balancing activity. If successful, a
// GameServerGroup object is returned. Learn more Amazon GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
func (c *Client) UpdateGameServerGroup(ctx context.Context, params *UpdateGameServerGroupInput, optFns ...func(*Options)) (*UpdateGameServerGroupOutput, error) {
	if params == nil {
		params = &UpdateGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGameServerGroup", params, optFns, c.addOperationUpdateGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGameServerGroupInput struct {

	// A unique identifier for the game server group. Use either the name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// Indicates how Amazon GameLift FleetIQ balances the use of Spot Instances and
	// On-Demand Instances in the game server group. Method options include the
	// following:
	//   - SPOT_ONLY - Only Spot Instances are used in the game server group. If Spot
	//   Instances are unavailable or not viable for game hosting, the game server group
	//   provides no hosting capacity until Spot Instances can again be used. Until then,
	//   no new instances are started, and the existing nonviable Spot Instances are
	//   terminated (after current gameplay ends) and are not replaced.
	//   - SPOT_PREFERRED - (default value) Spot Instances are used whenever available
	//   in the game server group. If Spot Instances are unavailable, the game server
	//   group continues to provide hosting capacity by falling back to On-Demand
	//   Instances. Existing nonviable Spot Instances are terminated (after current
	//   gameplay ends) and are replaced with new On-Demand Instances.
	//   - ON_DEMAND_ONLY - Only On-Demand Instances are used in the game server group.
	//   No Spot Instances are used, even when available, while this balancing strategy
	//   is in force.
	BalancingStrategy types.BalancingStrategy

	// A flag that indicates whether instances in the game server group are protected
	// from early termination. Unprotected instances that have active game servers
	// running might be terminated during a scale-down event, causing players to be
	// dropped from the game. Protected instances cannot be terminated while there are
	// active game servers running except in the event of a forced game server group
	// deletion (see ). An exception to this is with Spot Instances, which can be
	// terminated by Amazon Web Services regardless of protection status. This property
	// is set to NO_PROTECTION by default.
	GameServerProtectionPolicy types.GameServerProtectionPolicy

	// An updated list of Amazon EC2 instance types to use in the Auto Scaling group.
	// The instance definitions must specify at least two different instance types that
	// are supported by Amazon GameLift FleetIQ. This updated list replaces the entire
	// current list of instance definitions for the game server group. For more
	// information on instance types, see EC2 Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon EC2 User Guide. You can optionally specify capacity weighting for
	// each instance type. If no weight value is specified for an instance type, it is
	// set to the default value "1". For more information about capacity weighting, see
	// Instance Weighting for Amazon EC2 Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceDefinitions []types.InstanceDefinition

	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) for an IAM role that allows Amazon GameLift to access your Amazon EC2 Auto
	// Scaling groups.
	RoleArn *string

	noSmithyDocumentSerde
}

type UpdateGameServerGroupOutput struct {

	// An object that describes the game server group resource with updated properties.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameServerGroup{}, middleware.After)
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
	if err = addOpUpdateGameServerGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameServerGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGameServerGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameServerGroup",
	}
}
