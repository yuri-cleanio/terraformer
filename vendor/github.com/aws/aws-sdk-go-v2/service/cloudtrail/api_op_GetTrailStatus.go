// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The name of a trail about which you want the current status.
type GetTrailStatusInput struct {
	_ struct{} `type:"structure"`

	// Specifies the name or the CloudTrail ARN of the trail for which you are requesting
	// status. To get the status of a shadow trail (a replication of the trail in
	// another region), you must specify its ARN. The format of a trail ARN is:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTrailStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTrailStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTrailStatusInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type GetTrailStatusOutput struct {
	_ struct{} `type:"structure"`

	// Whether the CloudTrail is currently logging AWS API calls.
	IsLogging *bool `type:"boolean"`

	// Displays any CloudWatch Logs error that CloudTrail encountered when attempting
	// to deliver logs to CloudWatch Logs.
	LatestCloudWatchLogsDeliveryError *string `type:"string"`

	// Displays the most recent date and time when CloudTrail delivered logs to
	// CloudWatch Logs.
	LatestCloudWatchLogsDeliveryTime *time.Time `type:"timestamp"`

	// This field is no longer in use.
	LatestDeliveryAttemptSucceeded *string `type:"string"`

	// This field is no longer in use.
	LatestDeliveryAttemptTime *string `type:"string"`

	// Displays any Amazon S3 error that CloudTrail encountered when attempting
	// to deliver log files to the designated bucket. For more information see the
	// topic Error Responses (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html)
	// in the Amazon S3 API Reference.
	//
	// This error occurs only when there is a problem with the destination S3 bucket
	// and will not occur for timeouts. To resolve the issue, create a new bucket
	// and call UpdateTrail to specify the new bucket, or fix the existing objects
	// so that CloudTrail can again write to the bucket.
	LatestDeliveryError *string `type:"string"`

	// Specifies the date and time that CloudTrail last delivered log files to an
	// account's Amazon S3 bucket.
	LatestDeliveryTime *time.Time `type:"timestamp"`

	// Displays any Amazon S3 error that CloudTrail encountered when attempting
	// to deliver a digest file to the designated bucket. For more information see
	// the topic Error Responses (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html)
	// in the Amazon S3 API Reference.
	//
	// This error occurs only when there is a problem with the destination S3 bucket
	// and will not occur for timeouts. To resolve the issue, create a new bucket
	// and call UpdateTrail to specify the new bucket, or fix the existing objects
	// so that CloudTrail can again write to the bucket.
	LatestDigestDeliveryError *string `type:"string"`

	// Specifies the date and time that CloudTrail last delivered a digest file
	// to an account's Amazon S3 bucket.
	LatestDigestDeliveryTime *time.Time `type:"timestamp"`

	// This field is no longer in use.
	LatestNotificationAttemptSucceeded *string `type:"string"`

	// This field is no longer in use.
	LatestNotificationAttemptTime *string `type:"string"`

	// Displays any Amazon SNS error that CloudTrail encountered when attempting
	// to send a notification. For more information about Amazon SNS errors, see
	// the Amazon SNS Developer Guide (https://docs.aws.amazon.com/sns/latest/dg/welcome.html).
	LatestNotificationError *string `type:"string"`

	// Specifies the date and time of the most recent Amazon SNS notification that
	// CloudTrail has written a new log file to an account's Amazon S3 bucket.
	LatestNotificationTime *time.Time `type:"timestamp"`

	// Specifies the most recent date and time when CloudTrail started recording
	// API calls for an AWS account.
	StartLoggingTime *time.Time `type:"timestamp"`

	// Specifies the most recent date and time when CloudTrail stopped recording
	// API calls for an AWS account.
	StopLoggingTime *time.Time `type:"timestamp"`

	// This field is no longer in use.
	TimeLoggingStarted *string `type:"string"`

	// This field is no longer in use.
	TimeLoggingStopped *string `type:"string"`
}

// String returns the string representation
func (s GetTrailStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTrailStatus = "GetTrailStatus"

// GetTrailStatusRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Returns a JSON-formatted list of information about the specified trail. Fields
// include information on delivery errors, Amazon SNS and Amazon S3 errors,
// and start and stop logging times for each trail. This operation returns trail
// status from a single region. To return trail status from all regions, you
// must call the operation on each region.
//
//    // Example sending a request using GetTrailStatusRequest.
//    req := client.GetTrailStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/GetTrailStatus
func (c *Client) GetTrailStatusRequest(input *GetTrailStatusInput) GetTrailStatusRequest {
	op := &aws.Operation{
		Name:       opGetTrailStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTrailStatusInput{}
	}

	req := c.newRequest(op, input, &GetTrailStatusOutput{})
	return GetTrailStatusRequest{Request: req, Input: input, Copy: c.GetTrailStatusRequest}
}

// GetTrailStatusRequest is the request type for the
// GetTrailStatus API operation.
type GetTrailStatusRequest struct {
	*aws.Request
	Input *GetTrailStatusInput
	Copy  func(*GetTrailStatusInput) GetTrailStatusRequest
}

// Send marshals and sends the GetTrailStatus API request.
func (r GetTrailStatusRequest) Send(ctx context.Context) (*GetTrailStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTrailStatusResponse{
		GetTrailStatusOutput: r.Request.Data.(*GetTrailStatusOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTrailStatusResponse is the response type for the
// GetTrailStatus API operation.
type GetTrailStatusResponse struct {
	*GetTrailStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTrailStatus request.
func (r *GetTrailStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
