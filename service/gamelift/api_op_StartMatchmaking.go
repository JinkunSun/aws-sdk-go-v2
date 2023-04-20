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

// Uses FlexMatch to create a game match for a group of players based on custom
// matchmaking rules. With games that use Amazon GameLift managed hosting, this
// operation also triggers Amazon GameLift to find hosting resources and start a
// new game session for the new match. Each matchmaking request includes
// information on one or more players and specifies the FlexMatch matchmaker to
// use. When a request is for multiple players, FlexMatch attempts to build a match
// that includes all players in the request, placing them in the same team and
// finding additional players as needed to fill the match. To start matchmaking,
// provide a unique ticket ID, specify a matchmaking configuration, and include the
// players to be matched. You must also include any player attributes that are
// required by the matchmaking configuration's rule set. If successful, a
// matchmaking ticket is returned with status set to QUEUED . Track matchmaking
// events to respond as needed and acquire game session connection information for
// successfully completed matches. Ticket status updates are tracked using event
// notification through Amazon Simple Notification Service, which is defined in the
// matchmaking configuration. Learn more Add FlexMatch to a game client (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html)
// Set Up FlexMatch event notification (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html)
// How Amazon GameLift FlexMatch works (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/gamelift-match.html)
func (c *Client) StartMatchmaking(ctx context.Context, params *StartMatchmakingInput, optFns ...func(*Options)) (*StartMatchmakingOutput, error) {
	if params == nil {
		params = &StartMatchmakingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMatchmaking", params, optFns, c.addOperationStartMatchmakingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMatchmakingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMatchmakingInput struct {

	// Name of the matchmaking configuration to use for this request. Matchmaking
	// configurations must exist in the same Region as this request. You can use either
	// the configuration name or ARN value.
	//
	// This member is required.
	ConfigurationName *string

	// Information on each player to be matched. This information must include a
	// player ID, and may contain player attributes and latency data to be used in the
	// matchmaking process. After a successful match, Player objects contain the name
	// of the team the player is assigned to. You can include up to 10 Players in a
	// StartMatchmaking request.
	//
	// This member is required.
	Players []types.Player

	// A unique identifier for a matchmaking ticket. If no ticket ID is specified
	// here, Amazon GameLift will generate one in the form of a UUID. Use this
	// identifier to track the matchmaking ticket status and retrieve match results.
	TicketId *string

	noSmithyDocumentSerde
}

type StartMatchmakingOutput struct {

	// Ticket representing the matchmaking request. This object include the
	// information included in the request, ticket status, and match results as
	// generated during the matchmaking process.
	MatchmakingTicket *types.MatchmakingTicket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMatchmakingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMatchmaking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMatchmaking{}, middleware.After)
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
	if err = addOpStartMatchmakingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMatchmaking(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMatchmaking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StartMatchmaking",
	}
}
