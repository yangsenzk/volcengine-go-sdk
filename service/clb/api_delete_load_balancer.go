// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteLoadBalancerCommon = "DeleteLoadBalancer"

// DeleteLoadBalancerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteLoadBalancerCommon operation. The "output" return
// value will be populated with the DeleteLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteLoadBalancerCommon Send returns without error.
//
// See DeleteLoadBalancerCommon for more information on using the DeleteLoadBalancerCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteLoadBalancerCommonRequest method.
//    req, resp := client.DeleteLoadBalancerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteLoadBalancerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteLoadBalancerCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteLoadBalancerCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteLoadBalancerCommon for usage and error information.
func (c *CLB) DeleteLoadBalancerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteLoadBalancerCommonRequest(input)
	return out, req.Send()
}

// DeleteLoadBalancerCommonWithContext is the same as DeleteLoadBalancerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteLoadBalancerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteLoadBalancerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteLoadBalancerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteLoadBalancer = "DeleteLoadBalancer"

// DeleteLoadBalancerRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteLoadBalancer operation. The "output" return
// value will be populated with the DeleteLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteLoadBalancerCommon Send returns without error.
//
// See DeleteLoadBalancer for more information on using the DeleteLoadBalancer
// API call, and error handling.
//
//    // Example sending a request using the DeleteLoadBalancerRequest method.
//    req, resp := client.DeleteLoadBalancerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteLoadBalancerRequest(input *DeleteLoadBalancerInput) (req *request.Request, output *DeleteLoadBalancerOutput) {
	op := &request.Operation{
		Name:       opDeleteLoadBalancer,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLoadBalancerInput{}
	}

	output = &DeleteLoadBalancerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteLoadBalancer API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteLoadBalancer for usage and error information.
func (c *CLB) DeleteLoadBalancer(input *DeleteLoadBalancerInput) (*DeleteLoadBalancerOutput, error) {
	req, out := c.DeleteLoadBalancerRequest(input)
	return out, req.Send()
}

// DeleteLoadBalancerWithContext is the same as DeleteLoadBalancer with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteLoadBalancer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteLoadBalancerWithContext(ctx volcengine.Context, input *DeleteLoadBalancerInput, opts ...request.Option) (*DeleteLoadBalancerOutput, error) {
	req, out := c.DeleteLoadBalancerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	ForceDelete *bool `type:"boolean"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLoadBalancerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoadBalancerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteLoadBalancerInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetForceDelete sets the ForceDelete field's value.
func (s *DeleteLoadBalancerInput) SetForceDelete(v bool) *DeleteLoadBalancerInput {
	s.ForceDelete = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DeleteLoadBalancerInput) SetLoadBalancerId(v string) *DeleteLoadBalancerInput {
	s.LoadBalancerId = &v
	return s
}

type DeleteLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteLoadBalancerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLoadBalancerOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteLoadBalancerOutput) SetRequestId(v string) *DeleteLoadBalancerOutput {
	s.RequestId = &v
	return s
}
