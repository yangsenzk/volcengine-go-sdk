// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceAZConfigureCommon = "ModifyDBInstanceAZConfigure"

// ModifyDBInstanceAZConfigureCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceAZConfigureCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceAZConfigureCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceAZConfigureCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceAZConfigureCommon Send returns without error.
//
// See ModifyDBInstanceAZConfigureCommon for more information on using the ModifyDBInstanceAZConfigureCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceAZConfigureCommonRequest method.
//    req, resp := client.ModifyDBInstanceAZConfigureCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceAZConfigureCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceAZConfigureCommon,
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

// ModifyDBInstanceAZConfigureCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceAZConfigureCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceAZConfigureCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceAZConfigureCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceAZConfigureCommonWithContext is the same as ModifyDBInstanceAZConfigureCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceAZConfigureCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceAZConfigureCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceAZConfigureCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceAZConfigure = "ModifyDBInstanceAZConfigure"

// ModifyDBInstanceAZConfigureRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceAZConfigure operation. The "output" return
// value will be populated with the ModifyDBInstanceAZConfigureCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceAZConfigureCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceAZConfigureCommon Send returns without error.
//
// See ModifyDBInstanceAZConfigure for more information on using the ModifyDBInstanceAZConfigure
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceAZConfigureRequest method.
//    req, resp := client.ModifyDBInstanceAZConfigureRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceAZConfigureRequest(input *ModifyDBInstanceAZConfigureInput) (req *request.Request, output *ModifyDBInstanceAZConfigureOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceAZConfigure,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceAZConfigureInput{}
	}

	output = &ModifyDBInstanceAZConfigureOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceAZConfigure API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceAZConfigure for usage and error information.
func (c *REDIS) ModifyDBInstanceAZConfigure(input *ModifyDBInstanceAZConfigureInput) (*ModifyDBInstanceAZConfigureOutput, error) {
	req, out := c.ModifyDBInstanceAZConfigureRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceAZConfigureWithContext is the same as ModifyDBInstanceAZConfigure with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceAZConfigure for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceAZConfigureWithContext(ctx volcengine.Context, input *ModifyDBInstanceAZConfigureInput, opts ...request.Option) (*ModifyDBInstanceAZConfigureOutput, error) {
	req, out := c.ModifyDBInstanceAZConfigureRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigureNodeForModifyDBInstanceAZConfigureInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AZ *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConfigureNodeForModifyDBInstanceAZConfigureInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigureNodeForModifyDBInstanceAZConfigureInput) GoString() string {
	return s.String()
}

// SetAZ sets the AZ field's value.
func (s *ConfigureNodeForModifyDBInstanceAZConfigureInput) SetAZ(v string) *ConfigureNodeForModifyDBInstanceAZConfigureInput {
	s.AZ = &v
	return s
}

type ModifyDBInstanceAZConfigureInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" json:",omitempty" required:"true"`

	BackupPointName *string `type:"string" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	ConfigureNodes []*ConfigureNodeForModifyDBInstanceAZConfigureInput `type:"list" json:",omitempty"`

	CreateBackup *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// MultiAZ is a required field
	MultiAZ *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceAZConfigureInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceAZConfigureInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceAZConfigureInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceAZConfigureInput"}
	if s.ApplyImmediately == nil {
		invalidParams.Add(request.NewErrParamRequired("ApplyImmediately"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.MultiAZ == nil {
		invalidParams.Add(request.NewErrParamRequired("MultiAZ"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApplyImmediately sets the ApplyImmediately field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetApplyImmediately(v bool) *ModifyDBInstanceAZConfigureInput {
	s.ApplyImmediately = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetBackupPointName(v string) *ModifyDBInstanceAZConfigureInput {
	s.BackupPointName = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetClientToken(v string) *ModifyDBInstanceAZConfigureInput {
	s.ClientToken = &v
	return s
}

// SetConfigureNodes sets the ConfigureNodes field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetConfigureNodes(v []*ConfigureNodeForModifyDBInstanceAZConfigureInput) *ModifyDBInstanceAZConfigureInput {
	s.ConfigureNodes = v
	return s
}

// SetCreateBackup sets the CreateBackup field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetCreateBackup(v bool) *ModifyDBInstanceAZConfigureInput {
	s.CreateBackup = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetInstanceId(v string) *ModifyDBInstanceAZConfigureInput {
	s.InstanceId = &v
	return s
}

// SetMultiAZ sets the MultiAZ field's value.
func (s *ModifyDBInstanceAZConfigureInput) SetMultiAZ(v string) *ModifyDBInstanceAZConfigureInput {
	s.MultiAZ = &v
	return s
}

type ModifyDBInstanceAZConfigureOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderNO *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBInstanceAZConfigureOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceAZConfigureOutput) GoString() string {
	return s.String()
}

// SetOrderNO sets the OrderNO field's value.
func (s *ModifyDBInstanceAZConfigureOutput) SetOrderNO(v string) *ModifyDBInstanceAZConfigureOutput {
	s.OrderNO = &v
	return s
}
