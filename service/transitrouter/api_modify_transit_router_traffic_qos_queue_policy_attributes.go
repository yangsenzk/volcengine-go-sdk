// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTransitRouterTrafficQosQueuePolicyAttributesCommon = "ModifyTransitRouterTrafficQosQueuePolicyAttributes"

// ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon operation. The "output" return
// value will be populated with the ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon Send returns without error.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon for more information on using the ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest method.
//    req, resp := client.ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTransitRouterTrafficQosQueuePolicyAttributesCommon,
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

// ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonWithContext is the same as ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTransitRouterTrafficQosQueuePolicyAttributes = "ModifyTransitRouterTrafficQosQueuePolicyAttributes"

// ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterTrafficQosQueuePolicyAttributes operation. The "output" return
// value will be populated with the ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterTrafficQosQueuePolicyAttributesCommon Send returns without error.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAttributes for more information on using the ModifyTransitRouterTrafficQosQueuePolicyAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest method.
//    req, resp := client.ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest(input *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) (req *request.Request, output *ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTransitRouterTrafficQosQueuePolicyAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitRouterTrafficQosQueuePolicyAttributesInput{}
	}

	output = &ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTransitRouterTrafficQosQueuePolicyAttributes API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterTrafficQosQueuePolicyAttributes for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributes(input *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) (*ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterTrafficQosQueuePolicyAttributesWithContext is the same as ModifyTransitRouterTrafficQosQueuePolicyAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterTrafficQosQueuePolicyAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterTrafficQosQueuePolicyAttributesWithContext(ctx volcengine.Context, input *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput, opts ...request.Option) (*ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput, error) {
	req, out := c.ModifyTransitRouterTrafficQosQueuePolicyAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTransitRouterTrafficQosQueuePolicyAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// TransitRouterTrafficQosQueuePolicyId is a required field
	TransitRouterTrafficQosQueuePolicyId *string `type:"string" required:"true"`

	TransitRouterTrafficQosQueuePolicyName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTransitRouterTrafficQosQueuePolicyAttributesInput"}
	if s.TransitRouterTrafficQosQueuePolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosQueuePolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) SetDescription(v string) *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput {
	s.Description = &v
	return s
}

// SetTransitRouterTrafficQosQueuePolicyId sets the TransitRouterTrafficQosQueuePolicyId field's value.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) SetTransitRouterTrafficQosQueuePolicyId(v string) *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput {
	s.TransitRouterTrafficQosQueuePolicyId = &v
	return s
}

// SetTransitRouterTrafficQosQueuePolicyName sets the TransitRouterTrafficQosQueuePolicyName field's value.
func (s *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput) SetTransitRouterTrafficQosQueuePolicyName(v string) *ModifyTransitRouterTrafficQosQueuePolicyAttributesInput {
	s.TransitRouterTrafficQosQueuePolicyName = &v
	return s
}

type ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterTrafficQosQueuePolicyAttributesOutput) GoString() string {
	return s.String()
}
