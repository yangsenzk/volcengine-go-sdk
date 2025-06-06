// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vepfs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCancelDataFlowTaskCommon = "CancelDataFlowTask"

// CancelDataFlowTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelDataFlowTaskCommon operation. The "output" return
// value will be populated with the CancelDataFlowTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelDataFlowTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelDataFlowTaskCommon Send returns without error.
//
// See CancelDataFlowTaskCommon for more information on using the CancelDataFlowTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the CancelDataFlowTaskCommonRequest method.
//    req, resp := client.CancelDataFlowTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CancelDataFlowTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCancelDataFlowTaskCommon,
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

// CancelDataFlowTaskCommon API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CancelDataFlowTaskCommon for usage and error information.
func (c *VEPFS) CancelDataFlowTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CancelDataFlowTaskCommonRequest(input)
	return out, req.Send()
}

// CancelDataFlowTaskCommonWithContext is the same as CancelDataFlowTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CancelDataFlowTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CancelDataFlowTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CancelDataFlowTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCancelDataFlowTask = "CancelDataFlowTask"

// CancelDataFlowTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelDataFlowTask operation. The "output" return
// value will be populated with the CancelDataFlowTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelDataFlowTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelDataFlowTaskCommon Send returns without error.
//
// See CancelDataFlowTask for more information on using the CancelDataFlowTask
// API call, and error handling.
//
//    // Example sending a request using the CancelDataFlowTaskRequest method.
//    req, resp := client.CancelDataFlowTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CancelDataFlowTaskRequest(input *CancelDataFlowTaskInput) (req *request.Request, output *CancelDataFlowTaskOutput) {
	op := &request.Operation{
		Name:       opCancelDataFlowTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelDataFlowTaskInput{}
	}

	output = &CancelDataFlowTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CancelDataFlowTask API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CancelDataFlowTask for usage and error information.
func (c *VEPFS) CancelDataFlowTask(input *CancelDataFlowTaskInput) (*CancelDataFlowTaskOutput, error) {
	req, out := c.CancelDataFlowTaskRequest(input)
	return out, req.Send()
}

// CancelDataFlowTaskWithContext is the same as CancelDataFlowTask with the addition of
// the ability to pass a context and additional request options.
//
// See CancelDataFlowTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CancelDataFlowTaskWithContext(ctx volcengine.Context, input *CancelDataFlowTaskInput, opts ...request.Option) (*CancelDataFlowTaskOutput, error) {
	req, out := c.CancelDataFlowTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CancelDataFlowTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DataFlowTaskId is a required field
	DataFlowTaskId *string `type:"string" json:",omitempty" required:"true"`

	// FileSystemId is a required field
	FileSystemId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CancelDataFlowTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelDataFlowTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelDataFlowTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CancelDataFlowTaskInput"}
	if s.DataFlowTaskId == nil {
		invalidParams.Add(request.NewErrParamRequired("DataFlowTaskId"))
	}
	if s.FileSystemId == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDataFlowTaskId sets the DataFlowTaskId field's value.
func (s *CancelDataFlowTaskInput) SetDataFlowTaskId(v string) *CancelDataFlowTaskInput {
	s.DataFlowTaskId = &v
	return s
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *CancelDataFlowTaskInput) SetFileSystemId(v string) *CancelDataFlowTaskInput {
	s.FileSystemId = &v
	return s
}

type CancelDataFlowTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CancelDataFlowTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelDataFlowTaskOutput) GoString() string {
	return s.String()
}
