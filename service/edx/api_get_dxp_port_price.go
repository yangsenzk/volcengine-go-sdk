// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetDXPPortPriceCommon = "GetDXPPortPrice"

// GetDXPPortPriceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPPortPriceCommon operation. The "output" return
// value will be populated with the GetDXPPortPriceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPPortPriceCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPPortPriceCommon Send returns without error.
//
// See GetDXPPortPriceCommon for more information on using the GetDXPPortPriceCommon
// API call, and error handling.
//
//    // Example sending a request using the GetDXPPortPriceCommonRequest method.
//    req, resp := client.GetDXPPortPriceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPPortPriceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetDXPPortPriceCommon,
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

// GetDXPPortPriceCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPPortPriceCommon for usage and error information.
func (c *EDX) GetDXPPortPriceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetDXPPortPriceCommonRequest(input)
	return out, req.Send()
}

// GetDXPPortPriceCommonWithContext is the same as GetDXPPortPriceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPPortPriceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPPortPriceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetDXPPortPriceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetDXPPortPrice = "GetDXPPortPrice"

// GetDXPPortPriceRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPPortPrice operation. The "output" return
// value will be populated with the GetDXPPortPriceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPPortPriceCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPPortPriceCommon Send returns without error.
//
// See GetDXPPortPrice for more information on using the GetDXPPortPrice
// API call, and error handling.
//
//    // Example sending a request using the GetDXPPortPriceRequest method.
//    req, resp := client.GetDXPPortPriceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPPortPriceRequest(input *GetDXPPortPriceInput) (req *request.Request, output *GetDXPPortPriceOutput) {
	op := &request.Operation{
		Name:       opGetDXPPortPrice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDXPPortPriceInput{}
	}

	output = &GetDXPPortPriceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetDXPPortPrice API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPPortPrice for usage and error information.
func (c *EDX) GetDXPPortPrice(input *GetDXPPortPriceInput) (*GetDXPPortPriceOutput, error) {
	req, out := c.GetDXPPortPriceRequest(input)
	return out, req.Send()
}

// GetDXPPortPriceWithContext is the same as GetDXPPortPrice with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPPortPrice for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPPortPriceWithContext(ctx volcengine.Context, input *GetDXPPortPriceInput, opts ...request.Option) (*GetDXPPortPriceOutput, error) {
	req, out := c.GetDXPPortPriceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetDXPPortPriceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ModuleType is a required field
	ModuleType *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetDXPPortPriceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPPortPriceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDXPPortPriceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetDXPPortPriceInput"}
	if s.ModuleType == nil {
		invalidParams.Add(request.NewErrParamRequired("ModuleType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetModuleType sets the ModuleType field's value.
func (s *GetDXPPortPriceInput) SetModuleType(v string) *GetDXPPortPriceInput {
	s.ModuleType = &v
	return s
}

type GetDXPPortPriceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CouponAmount *float64 `type:"double" json:",omitempty"`

	Currency *string `type:"string" json:",omitempty"`

	DiscountAmount *float64 `type:"double" json:",omitempty"`

	OriginalAmount *float64 `type:"double" json:",omitempty"`

	PayableAmount *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s GetDXPPortPriceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPPortPriceOutput) GoString() string {
	return s.String()
}

// SetCouponAmount sets the CouponAmount field's value.
func (s *GetDXPPortPriceOutput) SetCouponAmount(v float64) *GetDXPPortPriceOutput {
	s.CouponAmount = &v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *GetDXPPortPriceOutput) SetCurrency(v string) *GetDXPPortPriceOutput {
	s.Currency = &v
	return s
}

// SetDiscountAmount sets the DiscountAmount field's value.
func (s *GetDXPPortPriceOutput) SetDiscountAmount(v float64) *GetDXPPortPriceOutput {
	s.DiscountAmount = &v
	return s
}

// SetOriginalAmount sets the OriginalAmount field's value.
func (s *GetDXPPortPriceOutput) SetOriginalAmount(v float64) *GetDXPPortPriceOutput {
	s.OriginalAmount = &v
	return s
}

// SetPayableAmount sets the PayableAmount field's value.
func (s *GetDXPPortPriceOutput) SetPayableAmount(v float64) *GetDXPPortPriceOutput {
	s.PayableAmount = &v
	return s
}
