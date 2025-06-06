// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSubmitOverseaPreloadTaskCommon = "SubmitOverseaPreloadTask"

// SubmitOverseaPreloadTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitOverseaPreloadTaskCommon operation. The "output" return
// value will be populated with the SubmitOverseaPreloadTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitOverseaPreloadTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitOverseaPreloadTaskCommon Send returns without error.
//
// See SubmitOverseaPreloadTaskCommon for more information on using the SubmitOverseaPreloadTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the SubmitOverseaPreloadTaskCommonRequest method.
//    req, resp := client.SubmitOverseaPreloadTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) SubmitOverseaPreloadTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSubmitOverseaPreloadTaskCommon,
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

// SubmitOverseaPreloadTaskCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation SubmitOverseaPreloadTaskCommon for usage and error information.
func (c *MCDN) SubmitOverseaPreloadTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SubmitOverseaPreloadTaskCommonRequest(input)
	return out, req.Send()
}

// SubmitOverseaPreloadTaskCommonWithContext is the same as SubmitOverseaPreloadTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitOverseaPreloadTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) SubmitOverseaPreloadTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SubmitOverseaPreloadTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitOverseaPreloadTask = "SubmitOverseaPreloadTask"

// SubmitOverseaPreloadTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitOverseaPreloadTask operation. The "output" return
// value will be populated with the SubmitOverseaPreloadTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitOverseaPreloadTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitOverseaPreloadTaskCommon Send returns without error.
//
// See SubmitOverseaPreloadTask for more information on using the SubmitOverseaPreloadTask
// API call, and error handling.
//
//    // Example sending a request using the SubmitOverseaPreloadTaskRequest method.
//    req, resp := client.SubmitOverseaPreloadTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) SubmitOverseaPreloadTaskRequest(input *SubmitOverseaPreloadTaskInput) (req *request.Request, output *SubmitOverseaPreloadTaskOutput) {
	op := &request.Operation{
		Name:       opSubmitOverseaPreloadTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitOverseaPreloadTaskInput{}
	}

	output = &SubmitOverseaPreloadTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SubmitOverseaPreloadTask API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation SubmitOverseaPreloadTask for usage and error information.
func (c *MCDN) SubmitOverseaPreloadTask(input *SubmitOverseaPreloadTaskInput) (*SubmitOverseaPreloadTaskOutput, error) {
	req, out := c.SubmitOverseaPreloadTaskRequest(input)
	return out, req.Send()
}

// SubmitOverseaPreloadTaskWithContext is the same as SubmitOverseaPreloadTask with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitOverseaPreloadTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) SubmitOverseaPreloadTaskWithContext(ctx volcengine.Context, input *SubmitOverseaPreloadTaskInput, opts ...request.Option) (*SubmitOverseaPreloadTaskOutput, error) {
	req, out := c.SubmitOverseaPreloadTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SubmitOverseaPreloadTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Area is a required field
	Area *string `type:"string" json:",omitempty" required:"true"`

	// Urls is a required field
	Urls *string `type:"string" json:",omitempty" required:"true"`

	// Vendor is a required field
	Vendor *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s SubmitOverseaPreloadTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitOverseaPreloadTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitOverseaPreloadTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SubmitOverseaPreloadTaskInput"}
	if s.Area == nil {
		invalidParams.Add(request.NewErrParamRequired("Area"))
	}
	if s.Urls == nil {
		invalidParams.Add(request.NewErrParamRequired("Urls"))
	}
	if s.Vendor == nil {
		invalidParams.Add(request.NewErrParamRequired("Vendor"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *SubmitOverseaPreloadTaskInput) SetArea(v string) *SubmitOverseaPreloadTaskInput {
	s.Area = &v
	return s
}

// SetUrls sets the Urls field's value.
func (s *SubmitOverseaPreloadTaskInput) SetUrls(v string) *SubmitOverseaPreloadTaskInput {
	s.Urls = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *SubmitOverseaPreloadTaskInput) SetVendor(v string) *SubmitOverseaPreloadTaskInput {
	s.Vendor = &v
	return s
}

type SubmitOverseaPreloadTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TaskId *string `type:"string" json:",omitempty"`

	TaskIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s SubmitOverseaPreloadTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitOverseaPreloadTaskOutput) GoString() string {
	return s.String()
}

// SetTaskId sets the TaskId field's value.
func (s *SubmitOverseaPreloadTaskOutput) SetTaskId(v string) *SubmitOverseaPreloadTaskOutput {
	s.TaskId = &v
	return s
}

// SetTaskIds sets the TaskIds field's value.
func (s *SubmitOverseaPreloadTaskOutput) SetTaskIds(v []*string) *SubmitOverseaPreloadTaskOutput {
	s.TaskIds = v
	return s
}
