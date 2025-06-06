// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVpcGatewayEndpointCommon = "CreateVpcGatewayEndpoint"

// CreateVpcGatewayEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpcGatewayEndpointCommon operation. The "output" return
// value will be populated with the CreateVpcGatewayEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpcGatewayEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpcGatewayEndpointCommon Send returns without error.
//
// See CreateVpcGatewayEndpointCommon for more information on using the CreateVpcGatewayEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateVpcGatewayEndpointCommonRequest method.
//    req, resp := client.CreateVpcGatewayEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) CreateVpcGatewayEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVpcGatewayEndpointCommon,
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

// CreateVpcGatewayEndpointCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation CreateVpcGatewayEndpointCommon for usage and error information.
func (c *PRIVATELINK) CreateVpcGatewayEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVpcGatewayEndpointCommonRequest(input)
	return out, req.Send()
}

// CreateVpcGatewayEndpointCommonWithContext is the same as CreateVpcGatewayEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpcGatewayEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) CreateVpcGatewayEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVpcGatewayEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVpcGatewayEndpoint = "CreateVpcGatewayEndpoint"

// CreateVpcGatewayEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpcGatewayEndpoint operation. The "output" return
// value will be populated with the CreateVpcGatewayEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpcGatewayEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpcGatewayEndpointCommon Send returns without error.
//
// See CreateVpcGatewayEndpoint for more information on using the CreateVpcGatewayEndpoint
// API call, and error handling.
//
//    // Example sending a request using the CreateVpcGatewayEndpointRequest method.
//    req, resp := client.CreateVpcGatewayEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) CreateVpcGatewayEndpointRequest(input *CreateVpcGatewayEndpointInput) (req *request.Request, output *CreateVpcGatewayEndpointOutput) {
	op := &request.Operation{
		Name:       opCreateVpcGatewayEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcGatewayEndpointInput{}
	}

	output = &CreateVpcGatewayEndpointOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVpcGatewayEndpoint API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation CreateVpcGatewayEndpoint for usage and error information.
func (c *PRIVATELINK) CreateVpcGatewayEndpoint(input *CreateVpcGatewayEndpointInput) (*CreateVpcGatewayEndpointOutput, error) {
	req, out := c.CreateVpcGatewayEndpointRequest(input)
	return out, req.Send()
}

// CreateVpcGatewayEndpointWithContext is the same as CreateVpcGatewayEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpcGatewayEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) CreateVpcGatewayEndpointWithContext(ctx volcengine.Context, input *CreateVpcGatewayEndpointInput, opts ...request.Option) (*CreateVpcGatewayEndpointOutput, error) {
	req, out := c.CreateVpcGatewayEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVpcGatewayEndpointInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	EndpointName *string `type:"string"`

	ProjectName *string `type:"string"`

	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`

	Tags []*TagForCreateVpcGatewayEndpointInput `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	VpcPolicy *string `type:"string"`
}

// String returns the string representation
func (s CreateVpcGatewayEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpcGatewayEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcGatewayEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVpcGatewayEndpointInput"}
	if s.ServiceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateVpcGatewayEndpointInput) SetClientToken(v string) *CreateVpcGatewayEndpointInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateVpcGatewayEndpointInput) SetDescription(v string) *CreateVpcGatewayEndpointInput {
	s.Description = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *CreateVpcGatewayEndpointInput) SetEndpointName(v string) *CreateVpcGatewayEndpointInput {
	s.EndpointName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateVpcGatewayEndpointInput) SetProjectName(v string) *CreateVpcGatewayEndpointInput {
	s.ProjectName = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *CreateVpcGatewayEndpointInput) SetServiceId(v string) *CreateVpcGatewayEndpointInput {
	s.ServiceId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateVpcGatewayEndpointInput) SetTags(v []*TagForCreateVpcGatewayEndpointInput) *CreateVpcGatewayEndpointInput {
	s.Tags = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateVpcGatewayEndpointInput) SetVpcId(v string) *CreateVpcGatewayEndpointInput {
	s.VpcId = &v
	return s
}

// SetVpcPolicy sets the VpcPolicy field's value.
func (s *CreateVpcGatewayEndpointInput) SetVpcPolicy(v string) *CreateVpcGatewayEndpointInput {
	s.VpcPolicy = &v
	return s
}

type CreateVpcGatewayEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EndpointId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpcGatewayEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpcGatewayEndpointOutput) GoString() string {
	return s.String()
}

// SetEndpointId sets the EndpointId field's value.
func (s *CreateVpcGatewayEndpointOutput) SetEndpointId(v string) *CreateVpcGatewayEndpointOutput {
	s.EndpointId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateVpcGatewayEndpointOutput) SetRequestId(v string) *CreateVpcGatewayEndpointOutput {
	s.RequestId = &v
	return s
}

type TagForCreateVpcGatewayEndpointInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateVpcGatewayEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateVpcGatewayEndpointInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateVpcGatewayEndpointInput) SetKey(v string) *TagForCreateVpcGatewayEndpointInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateVpcGatewayEndpointInput) SetValue(v string) *TagForCreateVpcGatewayEndpointInput {
	s.Value = &v
	return s
}
