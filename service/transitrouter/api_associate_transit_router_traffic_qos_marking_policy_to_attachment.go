// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon = "AssociateTransitRouterTrafficQosMarkingPolicyToAttachment"

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon operation. The "output" return
// value will be populated with the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon Send returns without error.
//
// See AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon for more information on using the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon
// API call, and error handling.
//
//    // Example sending a request using the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest method.
//    req, resp := client.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon,
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

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon for usage and error information.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest(input)
	return out, req.Send()
}

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonWithContext is the same as AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateTransitRouterTrafficQosMarkingPolicyToAttachment = "AssociateTransitRouterTrafficQosMarkingPolicyToAttachment"

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateTransitRouterTrafficQosMarkingPolicyToAttachment operation. The "output" return
// value will be populated with the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentCommon Send returns without error.
//
// See AssociateTransitRouterTrafficQosMarkingPolicyToAttachment for more information on using the AssociateTransitRouterTrafficQosMarkingPolicyToAttachment
// API call, and error handling.
//
//    // Example sending a request using the AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest method.
//    req, resp := client.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest(input *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) (req *request.Request, output *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput) {
	op := &request.Operation{
		Name:       opAssociateTransitRouterTrafficQosMarkingPolicyToAttachment,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput{}
	}

	output = &AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachment API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation AssociateTransitRouterTrafficQosMarkingPolicyToAttachment for usage and error information.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachment(input *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) (*AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput, error) {
	req, out := c.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest(input)
	return out, req.Send()
}

// AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentWithContext is the same as AssociateTransitRouterTrafficQosMarkingPolicyToAttachment with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateTransitRouterTrafficQosMarkingPolicyToAttachment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentWithContext(ctx volcengine.Context, input *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput, opts ...request.Option) (*AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput, error) {
	req, out := c.AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterAttachmentId is a required field
	TransitRouterAttachmentId *string `type:"string" required:"true"`

	// TransitRouterTrafficQosMarkingPolicyId is a required field
	TransitRouterTrafficQosMarkingPolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput"}
	if s.TransitRouterAttachmentId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterAttachmentId"))
	}
	if s.TransitRouterTrafficQosMarkingPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosMarkingPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterAttachmentId sets the TransitRouterAttachmentId field's value.
func (s *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) SetTransitRouterAttachmentId(v string) *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput {
	s.TransitRouterAttachmentId = &v
	return s
}

// SetTransitRouterTrafficQosMarkingPolicyId sets the TransitRouterTrafficQosMarkingPolicyId field's value.
func (s *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput) SetTransitRouterTrafficQosMarkingPolicyId(v string) *AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentInput {
	s.TransitRouterTrafficQosMarkingPolicyId = &v
	return s
}

type AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateTransitRouterTrafficQosMarkingPolicyToAttachmentOutput) GoString() string {
	return s.String()
}
