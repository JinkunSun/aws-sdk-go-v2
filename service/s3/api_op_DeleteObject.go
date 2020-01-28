// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type DeleteObjectInput struct {
	_ struct{} `type:"structure"`

	// The bucket name of the bucket containing the object.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Indicates whether S3 Object Lock should bypass Governance-mode restrictions
	// to process this operation.
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// Key name of the object to delete.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device. Required to
	// permanently delete a versioned object if versioning is configured with MFA
	// delete enabled.
	MFA *string `location:"header" locationName:"x-amz-mfa" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// For information about downloading objects from Requester Pays buckets, see
	// Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/http:/docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// VersionId used to reference a specific version of the object.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s DeleteObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteObjectInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BypassGovernanceRetention != nil {
		v := *s.BypassGovernanceRetention

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bypass-governance-retention", protocol.BoolValue(v), metadata)
	}
	if s.MFA != nil {
		v := *s.MFA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mfa", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeleteObjectInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteObjectInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteObjectOutput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the versioned object that was permanently deleted was (true)
	// or was not (false) a delete marker.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// Returns the version ID of the delete marker created as a result of the DELETE
	// operation.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s DeleteObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeleteMarker != nil {
		v := *s.DeleteMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-delete-marker", protocol.BoolValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opDeleteObject = "DeleteObject"

// DeleteObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Removes the null version (if there is one) of an object and inserts a delete
// marker, which becomes the latest version of the object. If there isn't a
// null version, Amazon S3 does not remove any objects.
//
// To remove a specific version, you must be the bucket owner and you must use
// the version Id subresource. Using this subresource permanently deletes the
// version. If the object deleted is a delete marker, Amazon S3 sets the response
// header, x-amz-delete-marker, to true.
//
// If the object you want to delete is in a bucket where the bucket versioning
// configuration is MFA Delete enabled, you must include the x-amz-mfa request
// header in the DELETE versionId request. Requests that include x-amz-mfa must
// use HTTPS.
//
// For more information about MFA Delete, see Using MFA Delete (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMFADelete.html).
// To see sample requests that use versioning, see Sample Request (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html#ExampleVersionObjectDelete).
//
// You can delete objects by explicitly calling the DELETE Object API or configure
// its lifecycle (PutBucketLifecycle) to enable Amazon S3 to remove them for
// you. If you want to block users or accounts from removing or deleting objects
// from your bucket, you must deny them the s3:DeleteObject, s3:DeleteObjectVersion,
// and s3:PutLifeCycleConfiguration actions.
//
// The following operation is related to DeleteObject:
//
//    * PutObject
//
//    // Example sending a request using DeleteObjectRequest.
//    req := client.DeleteObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteObject
func (c *Client) DeleteObjectRequest(input *DeleteObjectInput) DeleteObjectRequest {
	op := &aws.Operation{
		Name:       opDeleteObject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &DeleteObjectInput{}
	}

	req := c.newRequest(op, input, &DeleteObjectOutput{})
	return DeleteObjectRequest{Request: req, Input: input, Copy: c.DeleteObjectRequest}
}

// DeleteObjectRequest is the request type for the
// DeleteObject API operation.
type DeleteObjectRequest struct {
	*aws.Request
	Input *DeleteObjectInput
	Copy  func(*DeleteObjectInput) DeleteObjectRequest
}

// Send marshals and sends the DeleteObject API request.
func (r DeleteObjectRequest) Send(ctx context.Context) (*DeleteObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteObjectResponse{
		DeleteObjectOutput: r.Request.Data.(*DeleteObjectOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteObjectResponse is the response type for the
// DeleteObject API operation.
type DeleteObjectResponse struct {
	*DeleteObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteObject request.
func (r *DeleteObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
