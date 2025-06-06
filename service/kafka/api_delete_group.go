// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteGroupCommon = "DeleteGroup"

// DeleteGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteGroupCommon operation. The "output" return
// value will be populated with the DeleteGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteGroupCommon Send returns without error.
//
// See DeleteGroupCommon for more information on using the DeleteGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteGroupCommonRequest method.
//    req, resp := client.DeleteGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DeleteGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteGroupCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteGroupCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DeleteGroupCommon for usage and error information.
func (c *KAFKA) DeleteGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteGroupCommonRequest(input)
	return out, req.Send()
}

// DeleteGroupCommonWithContext is the same as DeleteGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DeleteGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteGroup = "DeleteGroup"

// DeleteGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteGroup operation. The "output" return
// value will be populated with the DeleteGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteGroupCommon Send returns without error.
//
// See DeleteGroup for more information on using the DeleteGroup
// API call, and error handling.
//
//    // Example sending a request using the DeleteGroupRequest method.
//    req, resp := client.DeleteGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) DeleteGroupRequest(input *DeleteGroupInput) (req *request.Request, output *DeleteGroupOutput) {
	op := &request.Operation{
		Name:       opDeleteGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGroupInput{}
	}

	output = &DeleteGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteGroup API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DeleteGroup for usage and error information.
func (c *KAFKA) DeleteGroup(input *DeleteGroupInput) (*DeleteGroupOutput, error) {
	req, out := c.DeleteGroupRequest(input)
	return out, req.Send()
}

// DeleteGroupWithContext is the same as DeleteGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DeleteGroupWithContext(ctx volcengine.Context, input *DeleteGroupInput, opts ...request.Option) (*DeleteGroupOutput, error) {
	req, out := c.DeleteGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// GroupId is a required field
	GroupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteGroupInput"}
	if s.GroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGroupId sets the GroupId field's value.
func (s *DeleteGroupInput) SetGroupId(v string) *DeleteGroupInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteGroupInput) SetInstanceId(v string) *DeleteGroupInput {
	s.InstanceId = &v
	return s
}

type DeleteGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteGroupOutput) GoString() string {
	return s.String()
}
