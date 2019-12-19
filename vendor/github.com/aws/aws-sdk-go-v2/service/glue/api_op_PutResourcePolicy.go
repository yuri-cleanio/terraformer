// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// A value of MUST_EXIST is used to update a policy. A value of NOT_EXIST is
	// used to create a new policy. If a value of NONE or a null value is used,
	// the call will not depend on the existence of a policy.
	PolicyExistsCondition ExistCondition `type:"string" enum:"true"`

	// The hash value returned when the previous policy was set using PutResourcePolicy.
	// Its purpose is to prevent concurrent modifications of a policy. Do not use
	// this parameter if no previous policy has been set.
	PolicyHashCondition *string `min:"1" type:"string"`

	// Contains the policy document to set, in JSON format.
	//
	// PolicyInJson is a required field
	PolicyInJson *string `min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourcePolicyInput"}
	if s.PolicyHashCondition != nil && len(*s.PolicyHashCondition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyHashCondition", 1))
	}

	if s.PolicyInJson == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyInJson"))
	}
	if s.PolicyInJson != nil && len(*s.PolicyInJson) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyInJson", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// A hash of the policy that has just been set. This must be included in a subsequent
	// call that overwrites or updates this policy.
	PolicyHash *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutResourcePolicy = "PutResourcePolicy"

// PutResourcePolicyRequest returns a request value for making API operation for
// AWS Glue.
//
// Sets the Data Catalog resource policy for access control.
//
//    // Example sending a request using PutResourcePolicyRequest.
//    req := client.PutResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/PutResourcePolicy
func (c *Client) PutResourcePolicyRequest(input *PutResourcePolicyInput) PutResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opPutResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &PutResourcePolicyOutput{})
	return PutResourcePolicyRequest{Request: req, Input: input, Copy: c.PutResourcePolicyRequest}
}

// PutResourcePolicyRequest is the request type for the
// PutResourcePolicy API operation.
type PutResourcePolicyRequest struct {
	*aws.Request
	Input *PutResourcePolicyInput
	Copy  func(*PutResourcePolicyInput) PutResourcePolicyRequest
}

// Send marshals and sends the PutResourcePolicy API request.
func (r PutResourcePolicyRequest) Send(ctx context.Context) (*PutResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResourcePolicyResponse{
		PutResourcePolicyOutput: r.Request.Data.(*PutResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResourcePolicyResponse is the response type for the
// PutResourcePolicy API operation.
type PutResourcePolicyResponse struct {
	*PutResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResourcePolicy request.
func (r *PutResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
