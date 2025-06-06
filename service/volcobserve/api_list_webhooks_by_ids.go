// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListWebhooksByIdsCommon = "ListWebhooksByIds"

// ListWebhooksByIdsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWebhooksByIdsCommon operation. The "output" return
// value will be populated with the ListWebhooksByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWebhooksByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWebhooksByIdsCommon Send returns without error.
//
// See ListWebhooksByIdsCommon for more information on using the ListWebhooksByIdsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListWebhooksByIdsCommonRequest method.
//    req, resp := client.ListWebhooksByIdsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListWebhooksByIdsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListWebhooksByIdsCommon,
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

// ListWebhooksByIdsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListWebhooksByIdsCommon for usage and error information.
func (c *VOLCOBSERVE) ListWebhooksByIdsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListWebhooksByIdsCommonRequest(input)
	return out, req.Send()
}

// ListWebhooksByIdsCommonWithContext is the same as ListWebhooksByIdsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListWebhooksByIdsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListWebhooksByIdsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListWebhooksByIdsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListWebhooksByIds = "ListWebhooksByIds"

// ListWebhooksByIdsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWebhooksByIds operation. The "output" return
// value will be populated with the ListWebhooksByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWebhooksByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWebhooksByIdsCommon Send returns without error.
//
// See ListWebhooksByIds for more information on using the ListWebhooksByIds
// API call, and error handling.
//
//    // Example sending a request using the ListWebhooksByIdsRequest method.
//    req, resp := client.ListWebhooksByIdsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListWebhooksByIdsRequest(input *ListWebhooksByIdsInput) (req *request.Request, output *ListWebhooksByIdsOutput) {
	op := &request.Operation{
		Name:       opListWebhooksByIds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListWebhooksByIdsInput{}
	}

	output = &ListWebhooksByIdsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListWebhooksByIds API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListWebhooksByIds for usage and error information.
func (c *VOLCOBSERVE) ListWebhooksByIds(input *ListWebhooksByIdsInput) (*ListWebhooksByIdsOutput, error) {
	req, out := c.ListWebhooksByIdsRequest(input)
	return out, req.Send()
}

// ListWebhooksByIdsWithContext is the same as ListWebhooksByIds with the addition of
// the ability to pass a context and additional request options.
//
// See ListWebhooksByIds for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListWebhooksByIdsWithContext(ctx volcengine.Context, input *ListWebhooksByIdsInput, opts ...request.Option) (*ListWebhooksByIdsOutput, error) {
	req, out := c.ListWebhooksByIdsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListWebhooksByIdsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	EventRuleIds []*string `type:"list" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	RuleIds []*string `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListWebhooksByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListWebhooksByIdsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListWebhooksByIdsOutput) SetCreatedAt(v string) *DataForListWebhooksByIdsOutput {
	s.CreatedAt = &v
	return s
}

// SetEventRuleIds sets the EventRuleIds field's value.
func (s *DataForListWebhooksByIdsOutput) SetEventRuleIds(v []*string) *DataForListWebhooksByIdsOutput {
	s.EventRuleIds = v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListWebhooksByIdsOutput) SetId(v string) *DataForListWebhooksByIdsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListWebhooksByIdsOutput) SetName(v string) *DataForListWebhooksByIdsOutput {
	s.Name = &v
	return s
}

// SetRuleIds sets the RuleIds field's value.
func (s *DataForListWebhooksByIdsOutput) SetRuleIds(v []*string) *DataForListWebhooksByIdsOutput {
	s.RuleIds = v
	return s
}

// SetType sets the Type field's value.
func (s *DataForListWebhooksByIdsOutput) SetType(v string) *DataForListWebhooksByIdsOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListWebhooksByIdsOutput) SetUpdatedAt(v string) *DataForListWebhooksByIdsOutput {
	s.UpdatedAt = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *DataForListWebhooksByIdsOutput) SetUrl(v string) *DataForListWebhooksByIdsOutput {
	s.Url = &v
	return s
}

type ListWebhooksByIdsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListWebhooksByIdsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWebhooksByIdsInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *ListWebhooksByIdsInput) SetIds(v []*string) *ListWebhooksByIdsInput {
	s.Ids = v
	return s
}

type ListWebhooksByIdsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListWebhooksByIdsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListWebhooksByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWebhooksByIdsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListWebhooksByIdsOutput) SetData(v []*DataForListWebhooksByIdsOutput) *ListWebhooksByIdsOutput {
	s.Data = v
	return s
}
