// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Must be the name of an existing DB cluster parameter group.
	//
	//    * You can't delete a default DB cluster parameter group.
	//
	//    * Can't be associated with any DB clusters.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBClusterParameterGroupInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBClusterParameterGroup = "DeleteDBClusterParameterGroup"

// DeleteDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes a specified DB cluster parameter group. The DB cluster parameter
// group to be deleted can't be associated with any DB clusters.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DeleteDBClusterParameterGroupRequest.
//    req := client.DeleteDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBClusterParameterGroup
func (c *Client) DeleteDBClusterParameterGroupRequest(input *DeleteDBClusterParameterGroupInput) DeleteDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBClusterParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteDBClusterParameterGroupRequest}
}

// DeleteDBClusterParameterGroupRequest is the request type for the
// DeleteDBClusterParameterGroup API operation.
type DeleteDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *DeleteDBClusterParameterGroupInput
	Copy  func(*DeleteDBClusterParameterGroupInput) DeleteDBClusterParameterGroupRequest
}

// Send marshals and sends the DeleteDBClusterParameterGroup API request.
func (r DeleteDBClusterParameterGroupRequest) Send(ctx context.Context) (*DeleteDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBClusterParameterGroupResponse{
		DeleteDBClusterParameterGroupOutput: r.Request.Data.(*DeleteDBClusterParameterGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBClusterParameterGroupResponse is the response type for the
// DeleteDBClusterParameterGroup API operation.
type DeleteDBClusterParameterGroupResponse struct {
	*DeleteDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBClusterParameterGroup request.
func (r *DeleteDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
