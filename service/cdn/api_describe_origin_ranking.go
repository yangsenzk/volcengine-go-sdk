// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeOriginRankingCommon = "DescribeOriginRanking"

// DescribeOriginRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginRankingCommon operation. The "output" return
// value will be populated with the DescribeOriginRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginRankingCommon Send returns without error.
//
// See DescribeOriginRankingCommon for more information on using the DescribeOriginRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginRankingCommonRequest method.
//    req, resp := client.DescribeOriginRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeOriginRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeOriginRankingCommon,
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

// DescribeOriginRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeOriginRankingCommon for usage and error information.
func (c *CDN) DescribeOriginRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeOriginRankingCommonWithContext is the same as DescribeOriginRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeOriginRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeOriginRanking = "DescribeOriginRanking"

// DescribeOriginRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginRanking operation. The "output" return
// value will be populated with the DescribeOriginRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginRankingCommon Send returns without error.
//
// See DescribeOriginRanking for more information on using the DescribeOriginRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginRankingRequest method.
//    req, resp := client.DescribeOriginRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeOriginRankingRequest(input *DescribeOriginRankingInput) (req *request.Request, output *DescribeOriginRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeOriginRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOriginRankingInput{}
	}

	output = &DescribeOriginRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeOriginRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeOriginRanking for usage and error information.
func (c *CDN) DescribeOriginRanking(input *DescribeOriginRankingInput) (*DescribeOriginRankingOutput, error) {
	req, out := c.DescribeOriginRankingRequest(input)
	return out, req.Send()
}

// DescribeOriginRankingWithContext is the same as DescribeOriginRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeOriginRankingWithContext(ctx volcengine.Context, input *DescribeOriginRankingInput, opts ...request.Option) (*DescribeOriginRankingOutput, error) {
	req, out := c.DescribeOriginRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeOriginRankingInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" json:",omitempty" required:"true"`

	Interval *string `type:"string" json:",omitempty"`

	// Item is a required field
	Item *string `type:"string" json:",omitempty" required:"true"`

	// Metric is a required field
	Metric *string `type:"string" json:",omitempty" required:"true"`

	Project *string `type:"string" json:",omitempty"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeOriginRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOriginRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeOriginRankingInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.Item == nil {
		invalidParams.Add(request.NewErrParamRequired("Item"))
	}
	if s.Metric == nil {
		invalidParams.Add(request.NewErrParamRequired("Metric"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomain sets the Domain field's value.
func (s *DescribeOriginRankingInput) SetDomain(v string) *DescribeOriginRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeOriginRankingInput) SetEndTime(v int64) *DescribeOriginRankingInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeOriginRankingInput) SetInterval(v string) *DescribeOriginRankingInput {
	s.Interval = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeOriginRankingInput) SetItem(v string) *DescribeOriginRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeOriginRankingInput) SetMetric(v string) *DescribeOriginRankingInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeOriginRankingInput) SetProject(v string) *DescribeOriginRankingInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeOriginRankingInput) SetStartTime(v int64) *DescribeOriginRankingInput {
	s.StartTime = &v
	return s
}

type DescribeOriginRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string" json:",omitempty"`

	TopDataDetails []*TopDataDetailForDescribeOriginRankingOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeOriginRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeOriginRankingOutput) SetItem(v string) *DescribeOriginRankingOutput {
	s.Item = &v
	return s
}

// SetTopDataDetails sets the TopDataDetails field's value.
func (s *DescribeOriginRankingOutput) SetTopDataDetails(v []*TopDataDetailForDescribeOriginRankingOutput) *DescribeOriginRankingOutput {
	s.TopDataDetails = v
	return s
}

type TopDataDetailForDescribeOriginRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metric *string `type:"string" json:",omitempty"`

	ValueDetails []*ValueDetailForDescribeOriginRankingOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TopDataDetailForDescribeOriginRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataDetailForDescribeOriginRankingOutput) GoString() string {
	return s.String()
}

// SetMetric sets the Metric field's value.
func (s *TopDataDetailForDescribeOriginRankingOutput) SetMetric(v string) *TopDataDetailForDescribeOriginRankingOutput {
	s.Metric = &v
	return s
}

// SetValueDetails sets the ValueDetails field's value.
func (s *TopDataDetailForDescribeOriginRankingOutput) SetValueDetails(v []*ValueDetailForDescribeOriginRankingOutput) *TopDataDetailForDescribeOriginRankingOutput {
	s.ValueDetails = v
	return s
}

type ValueDetailForDescribeOriginRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ItemKey *string `type:"string" json:",omitempty"`

	Ratio *float64 `type:"double" json:",omitempty"`

	Timestamp *int64 `type:"int64" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s ValueDetailForDescribeOriginRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueDetailForDescribeOriginRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *ValueDetailForDescribeOriginRankingOutput) SetItemKey(v string) *ValueDetailForDescribeOriginRankingOutput {
	s.ItemKey = &v
	return s
}

// SetRatio sets the Ratio field's value.
func (s *ValueDetailForDescribeOriginRankingOutput) SetRatio(v float64) *ValueDetailForDescribeOriginRankingOutput {
	s.Ratio = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *ValueDetailForDescribeOriginRankingOutput) SetTimestamp(v int64) *ValueDetailForDescribeOriginRankingOutput {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ValueDetailForDescribeOriginRankingOutput) SetValue(v float64) *ValueDetailForDescribeOriginRankingOutput {
	s.Value = &v
	return s
}
