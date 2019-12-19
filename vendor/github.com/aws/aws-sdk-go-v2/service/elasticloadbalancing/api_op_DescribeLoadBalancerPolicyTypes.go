// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeLoadBalancerPolicyTypes.
type DescribeLoadBalancerPolicyTypesInput struct {
	_ struct{} `type:"structure"`

	// The names of the policy types. If no names are specified, describes all policy
	// types defined by Elastic Load Balancing.
	PolicyTypeNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeLoadBalancerPolicyTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeLoadBalancerPolicyTypes.
type DescribeLoadBalancerPolicyTypesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the policy types.
	PolicyTypeDescriptions []PolicyTypeDescription `type:"list"`
}

// String returns the string representation
func (s DescribeLoadBalancerPolicyTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLoadBalancerPolicyTypes = "DescribeLoadBalancerPolicyTypes"

// DescribeLoadBalancerPolicyTypesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the specified load balancer policy types or all load balancer policy
// types.
//
// The description of each type indicates how it can be used. For example, some
// policies can be used only with layer 7 listeners, some policies can be used
// only with layer 4 listeners, and some policies can be used only with your
// EC2 instances.
//
// You can use CreateLoadBalancerPolicy to create a policy configuration for
// any of these policy types. Then, depending on the policy type, use either
// SetLoadBalancerPoliciesOfListener or SetLoadBalancerPoliciesForBackendServer
// to set the policy.
//
//    // Example sending a request using DescribeLoadBalancerPolicyTypesRequest.
//    req := client.DescribeLoadBalancerPolicyTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeLoadBalancerPolicyTypes
func (c *Client) DescribeLoadBalancerPolicyTypesRequest(input *DescribeLoadBalancerPolicyTypesInput) DescribeLoadBalancerPolicyTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBalancerPolicyTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancerPolicyTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBalancerPolicyTypesOutput{})
	return DescribeLoadBalancerPolicyTypesRequest{Request: req, Input: input, Copy: c.DescribeLoadBalancerPolicyTypesRequest}
}

// DescribeLoadBalancerPolicyTypesRequest is the request type for the
// DescribeLoadBalancerPolicyTypes API operation.
type DescribeLoadBalancerPolicyTypesRequest struct {
	*aws.Request
	Input *DescribeLoadBalancerPolicyTypesInput
	Copy  func(*DescribeLoadBalancerPolicyTypesInput) DescribeLoadBalancerPolicyTypesRequest
}

// Send marshals and sends the DescribeLoadBalancerPolicyTypes API request.
func (r DescribeLoadBalancerPolicyTypesRequest) Send(ctx context.Context) (*DescribeLoadBalancerPolicyTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBalancerPolicyTypesResponse{
		DescribeLoadBalancerPolicyTypesOutput: r.Request.Data.(*DescribeLoadBalancerPolicyTypesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoadBalancerPolicyTypesResponse is the response type for the
// DescribeLoadBalancerPolicyTypes API operation.
type DescribeLoadBalancerPolicyTypesResponse struct {
	*DescribeLoadBalancerPolicyTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBalancerPolicyTypes request.
func (r *DescribeLoadBalancerPolicyTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
