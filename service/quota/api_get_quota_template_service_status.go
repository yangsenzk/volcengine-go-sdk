// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package quota

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetQuotaTemplateServiceStatusCommon = "GetQuotaTemplateServiceStatus"

// GetQuotaTemplateServiceStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetQuotaTemplateServiceStatusCommon operation. The "output" return
// value will be populated with the GetQuotaTemplateServiceStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetQuotaTemplateServiceStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetQuotaTemplateServiceStatusCommon Send returns without error.
//
// See GetQuotaTemplateServiceStatusCommon for more information on using the GetQuotaTemplateServiceStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the GetQuotaTemplateServiceStatusCommonRequest method.
//    req, resp := client.GetQuotaTemplateServiceStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) GetQuotaTemplateServiceStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetQuotaTemplateServiceStatusCommon,
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

// GetQuotaTemplateServiceStatusCommon API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation GetQuotaTemplateServiceStatusCommon for usage and error information.
func (c *QUOTA) GetQuotaTemplateServiceStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetQuotaTemplateServiceStatusCommonRequest(input)
	return out, req.Send()
}

// GetQuotaTemplateServiceStatusCommonWithContext is the same as GetQuotaTemplateServiceStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetQuotaTemplateServiceStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) GetQuotaTemplateServiceStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetQuotaTemplateServiceStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetQuotaTemplateServiceStatus = "GetQuotaTemplateServiceStatus"

// GetQuotaTemplateServiceStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the GetQuotaTemplateServiceStatus operation. The "output" return
// value will be populated with the GetQuotaTemplateServiceStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetQuotaTemplateServiceStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetQuotaTemplateServiceStatusCommon Send returns without error.
//
// See GetQuotaTemplateServiceStatus for more information on using the GetQuotaTemplateServiceStatus
// API call, and error handling.
//
//    // Example sending a request using the GetQuotaTemplateServiceStatusRequest method.
//    req, resp := client.GetQuotaTemplateServiceStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) GetQuotaTemplateServiceStatusRequest(input *GetQuotaTemplateServiceStatusInput) (req *request.Request, output *GetQuotaTemplateServiceStatusOutput) {
	op := &request.Operation{
		Name:       opGetQuotaTemplateServiceStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQuotaTemplateServiceStatusInput{}
	}

	output = &GetQuotaTemplateServiceStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetQuotaTemplateServiceStatus API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation GetQuotaTemplateServiceStatus for usage and error information.
func (c *QUOTA) GetQuotaTemplateServiceStatus(input *GetQuotaTemplateServiceStatusInput) (*GetQuotaTemplateServiceStatusOutput, error) {
	req, out := c.GetQuotaTemplateServiceStatusRequest(input)
	return out, req.Send()
}

// GetQuotaTemplateServiceStatusWithContext is the same as GetQuotaTemplateServiceStatus with the addition of
// the ability to pass a context and additional request options.
//
// See GetQuotaTemplateServiceStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) GetQuotaTemplateServiceStatusWithContext(ctx volcengine.Context, input *GetQuotaTemplateServiceStatusInput, opts ...request.Option) (*GetQuotaTemplateServiceStatusOutput, error) {
	req, out := c.GetQuotaTemplateServiceStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetQuotaTemplateServiceStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetQuotaTemplateServiceStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQuotaTemplateServiceStatusInput) GoString() string {
	return s.String()
}

type GetQuotaTemplateServiceStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TemplateStatus *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GetQuotaTemplateServiceStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQuotaTemplateServiceStatusOutput) GoString() string {
	return s.String()
}

// SetTemplateStatus sets the TemplateStatus field's value.
func (s *GetQuotaTemplateServiceStatusOutput) SetTemplateStatus(v int32) *GetQuotaTemplateServiceStatusOutput {
	s.TemplateStatus = &v
	return s
}
