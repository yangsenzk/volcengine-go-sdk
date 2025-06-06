// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateAccountCommon = "CreateAccount"

// CreateAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccountCommon operation. The "output" return
// value will be populated with the CreateAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccountCommon Send returns without error.
//
// See CreateAccountCommon for more information on using the CreateAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateAccountCommonRequest method.
//    req, resp := client.CreateAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) CreateAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateAccountCommon,
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

// CreateAccountCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation CreateAccountCommon for usage and error information.
func (c *ORGANIZATION) CreateAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateAccountCommonRequest(input)
	return out, req.Send()
}

// CreateAccountCommonWithContext is the same as CreateAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) CreateAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateAccount = "CreateAccount"

// CreateAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccount operation. The "output" return
// value will be populated with the CreateAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccountCommon Send returns without error.
//
// See CreateAccount for more information on using the CreateAccount
// API call, and error handling.
//
//    // Example sending a request using the CreateAccountRequest method.
//    req, resp := client.CreateAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) CreateAccountRequest(input *CreateAccountInput) (req *request.Request, output *CreateAccountOutput) {
	op := &request.Operation{
		Name:       opCreateAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAccountInput{}
	}

	output = &CreateAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateAccount API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation CreateAccount for usage and error information.
func (c *ORGANIZATION) CreateAccount(input *CreateAccountInput) (*CreateAccountOutput, error) {
	req, out := c.CreateAccountRequest(input)
	return out, req.Send()
}

// CreateAccountWithContext is the same as CreateAccount with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) CreateAccountWithContext(ctx volcengine.Context, input *CreateAccountInput, opts ...request.Option) (*CreateAccountOutput, error) {
	req, out := c.CreateAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateAccountInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountName is a required field
	AccountName *string `type:"string" json:",omitempty" required:"true"`

	Description *string `type:"string" json:",omitempty"`

	InheritProfile *bool `type:"boolean" json:",omitempty"`

	OrgUnitId *string `type:"string" json:",omitempty"`

	// ShowName is a required field
	ShowName *string `type:"string" json:",omitempty" required:"true"`

	VerificationRelationId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateAccountInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
	if s.ShowName == nil {
		invalidParams.Add(request.NewErrParamRequired("ShowName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *CreateAccountInput) SetAccountName(v string) *CreateAccountInput {
	s.AccountName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateAccountInput) SetDescription(v string) *CreateAccountInput {
	s.Description = &v
	return s
}

// SetInheritProfile sets the InheritProfile field's value.
func (s *CreateAccountInput) SetInheritProfile(v bool) *CreateAccountInput {
	s.InheritProfile = &v
	return s
}

// SetOrgUnitId sets the OrgUnitId field's value.
func (s *CreateAccountInput) SetOrgUnitId(v string) *CreateAccountInput {
	s.OrgUnitId = &v
	return s
}

// SetShowName sets the ShowName field's value.
func (s *CreateAccountInput) SetShowName(v string) *CreateAccountInput {
	s.ShowName = &v
	return s
}

// SetVerificationRelationId sets the VerificationRelationId field's value.
func (s *CreateAccountInput) SetVerificationRelationId(v string) *CreateAccountInput {
	s.VerificationRelationId = &v
	return s
}

type CreateAccountOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccountOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CreateAccountOutput) SetAccountId(v string) *CreateAccountOutput {
	s.AccountId = &v
	return s
}
