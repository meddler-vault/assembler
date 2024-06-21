// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation enables you to delete multiple objects from a bucket using a
// single HTTP request. If you know the object keys that you want to delete, then
// this operation provides a suitable alternative to sending individual delete
// requests, reducing per-request overhead.
//
// The request can contain a list of up to 1000 keys that you want to delete. In
// the XML, you provide the object key names, and optionally, version IDs if you
// want to delete a specific version of the object from a versioning-enabled
// bucket. For each key, Amazon S3 performs a delete operation and returns the
// result of that delete, success or failure, in the response. Note that if the
// object specified in the request is not found, Amazon S3 returns the result as
// deleted.
//
//   - Directory buckets - S3 Versioning isn't enabled and supported for directory
//     buckets.
//
//   - Directory buckets - For directory buckets, you must make requests for this
//     API operation to the Zonal endpoint. These endpoints support
//     virtual-hosted-style requests in the format
//     https://bucket_name.s3express-az_id.region.amazonaws.com/key-name .
//     Path-style requests are not supported. For more information, see [Regional and Zonal endpoints]in the
//     Amazon S3 User Guide.
//
// The operation supports two modes for the response: verbose and quiet. By
// default, the operation uses verbose mode in which the response includes the
// result of deletion of each key in your request. In quiet mode the response
// includes only keys where the delete operation encountered an error. For a
// successful deletion in a quiet mode, the operation does not return any
// information about the delete in the response body.
//
// When performing this action on an MFA Delete enabled bucket, that attempts to
// delete any versioned objects, you must include an MFA token. If you do not
// provide one, the entire request will fail, even if there are non-versioned
// objects you are trying to delete. If you provide an invalid token, whether there
// are versioned keys in the request or not, the entire Multi-Object Delete request
// will fail. For information about MFA Delete, see [MFA Delete]in the Amazon S3 User Guide.
//
// Directory buckets - MFA delete is not supported by directory buckets.
//
// Permissions
//
//   - General purpose bucket permissions - The following permissions are required
//     in your policies when your DeleteObjects request includes specific headers.
//
//   - s3:DeleteObject - To delete an object from a bucket, you must always specify
//     the s3:DeleteObject permission.
//
//   - s3:DeleteObjectVersion - To delete a specific version of an object from a
//     versiong-enabled bucket, you must specify the s3:DeleteObjectVersion
//     permission.
//
//   - Directory bucket permissions - To grant access to this API operation on a
//     directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
//     for session-based authorization. Specifically, you grant the
//     s3express:CreateSession permission to the directory bucket in a bucket policy
//     or an IAM identity-based policy. Then, you make the CreateSession API call on
//     the bucket to obtain a session token. With the session token in your request
//     header, you can make API requests to this operation. After the session token
//     expires, you make another CreateSession API call to generate a new session
//     token for use. Amazon Web Services CLI or SDKs create session and refresh the
//     session token automatically to avoid service interruptions when a session
//     expires. For more information about authorization, see [CreateSession]CreateSession .
//
// Content-MD5 request header
//
//   - General purpose bucket - The Content-MD5 request header is required for all
//     Multi-Object Delete requests. Amazon S3 uses the header value to ensure that
//     your request body has not been altered in transit.
//
//   - Directory bucket - The Content-MD5 request header or a additional checksum
//     request header (including x-amz-checksum-crc32 , x-amz-checksum-crc32c ,
//     x-amz-checksum-sha1 , or x-amz-checksum-sha256 ) is required for all
//     Multi-Object Delete requests.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket_name.s3express-az_id.region.amazonaws.com .
//
// The following operations are related to DeleteObjects :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [AbortMultipartUpload]
//
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [Regional and Zonal endpoints]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-Regions-and-Zones.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [MFA Delete]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
func (c *Client) DeleteObjects(ctx context.Context, params *DeleteObjectsInput, optFns ...func(*Options)) (*DeleteObjectsOutput, error) {
	if params == nil {
		params = &DeleteObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteObjects", params, optFns, c.addOperationDeleteObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteObjectsInput struct {

	// The bucket name containing the objects to delete.
	//
	// Directory buckets - When you use this operation with a directory bucket, you
	// must use virtual-hosted-style requests in the format
	// Bucket_name.s3express-az_id.region.amazonaws.com . Path-style requests are not
	// supported. Directory bucket names must be unique in the chosen Availability
	// Zone. Bucket names must follow the format bucket_base_name--az-id--x-s3 (for
	// example, DOC-EXAMPLE-BUCKET--usw2-az1--x-s3 ). For information about bucket
	// naming restrictions, see [Directory bucket naming rules]in the Amazon S3 User Guide.
	//
	// Access points - When you use this action with an access point, you must provide
	// the alias of the access point in place of the bucket name or specify the access
	// point ARN. When using the access point ARN, you must direct requests to the
	// access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see [Using access points]in the Amazon S3 User Guide.
	//
	// Access points and Object Lambda access points are not supported by directory
	// buckets.
	//
	// S3 on Outposts - When you use this action with Amazon S3 on Outposts, you must
	// direct requests to the S3 on Outposts hostname. The S3 on Outposts hostname
	// takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see [What is S3 on Outposts?]in the Amazon S3 User Guide.
	//
	// [Directory bucket naming rules]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-bucket-naming-rules.html
	// [What is S3 on Outposts?]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
	// [Using access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html
	//
	// This member is required.
	Bucket *string

	// Container for the request.
	//
	// This member is required.
	Delete *types.Delete

	// Specifies whether you want to delete this object even if it has a
	// Governance-type Object Lock in place. To use this header, you must have the
	// s3:BypassGovernanceRetention permission.
	//
	// This functionality is not supported for directory buckets.
	BypassGovernanceRetention *bool

	// Indicates the algorithm used to create the checksum for the object when you use
	// the SDK. This header will not provide any additional functionality if you don't
	// use the SDK. When you send this header, there must be a corresponding
	// x-amz-checksum-algorithm or x-amz-trailer header sent. Otherwise, Amazon S3
	// fails the request with the HTTP status code 400 Bad Request .
	//
	// For the x-amz-checksum-algorithm  header, replace  algorithm  with the
	// supported algorithm from the following list:
	//
	//   - CRC32
	//
	//   - CRC32C
	//
	//   - SHA1
	//
	//   - SHA256
	//
	// For more information, see [Checking object integrity] in the Amazon S3 User Guide.
	//
	// If the individual checksum value you provide through x-amz-checksum-algorithm
	// doesn't match the checksum algorithm you set through
	// x-amz-sdk-checksum-algorithm , Amazon S3 ignores any provided ChecksumAlgorithm
	// parameter and uses the checksum algorithm that matches the provided value in
	// x-amz-checksum-algorithm .
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided
	// ChecksumAlgorithm parameter.
	//
	// [Checking object integrity]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// The concatenation of the authentication device's serial number, a space, and
	// the value that is displayed on your authentication device. Required to
	// permanently delete a versioned object if versioning is configured with MFA
	// delete enabled.
	//
	// When performing the DeleteObjects operation on an MFA delete enabled bucket,
	// which attempts to delete the specified versioned objects, you must include an
	// MFA token. If you don't provide an MFA token, the entire request will fail, even
	// if there are non-versioned objects that you are trying to delete. If you provide
	// an invalid token, whether there are versioned object keys in the request or not,
	// the entire Multi-Object Delete request will fail. For information about MFA
	// Delete, see [MFA Delete]in the Amazon S3 User Guide.
	//
	// This functionality is not supported for directory buckets.
	//
	// [MFA Delete]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete
	MFA *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. If either the
	// source or destination S3 bucket has Requester Pays enabled, the requester will
	// pay for corresponding charges to copy the object. For information about
	// downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays Buckets]in the Amazon S3 User
	// Guide.
	//
	// This functionality is not supported for directory buckets.
	//
	// [Downloading Objects in Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer types.RequestPayer

	noSmithyDocumentSerde
}

func (in *DeleteObjectsInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket

}

type DeleteObjectsOutput struct {

	// Container element for a successful delete. It identifies the object that was
	// successfully deleted.
	Deleted []types.DeletedObject

	// Container for a failed delete action that describes the object that Amazon S3
	// attempted to delete and the error it encountered.
	Errors []types.Error

	// If present, indicates that the requester was successfully charged for the
	// request.
	//
	// This functionality is not supported for directory buckets.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteObjects{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteObjects"); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpDeleteObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteObjects(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteObjectsInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addDeleteObjectsUpdateEndpoint(stack, options); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3cust.AddExpressDefaultChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (v *DeleteObjectsInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opDeleteObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteObjects",
	}
}

// getDeleteObjectsRequestAlgorithmMember gets the request checksum algorithm
// value provided as input.
func getDeleteObjectsRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*DeleteObjectsInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addDeleteObjectsInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getDeleteObjectsRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getDeleteObjectsBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getDeleteObjectsBucketMember(input interface{}) (*string, bool) {
	in := input.(*DeleteObjectsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addDeleteObjectsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getDeleteObjectsBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
