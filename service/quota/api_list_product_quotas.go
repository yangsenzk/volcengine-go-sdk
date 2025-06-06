// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package quota

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListProductQuotasCommon = "ListProductQuotas"

// ListProductQuotasCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProductQuotasCommon operation. The "output" return
// value will be populated with the ListProductQuotasCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProductQuotasCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProductQuotasCommon Send returns without error.
//
// See ListProductQuotasCommon for more information on using the ListProductQuotasCommon
// API call, and error handling.
//
//    // Example sending a request using the ListProductQuotasCommonRequest method.
//    req, resp := client.ListProductQuotasCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ListProductQuotasCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListProductQuotasCommon,
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

// ListProductQuotasCommon API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ListProductQuotasCommon for usage and error information.
func (c *QUOTA) ListProductQuotasCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListProductQuotasCommonRequest(input)
	return out, req.Send()
}

// ListProductQuotasCommonWithContext is the same as ListProductQuotasCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListProductQuotasCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ListProductQuotasCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListProductQuotasCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListProductQuotas = "ListProductQuotas"

// ListProductQuotasRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProductQuotas operation. The "output" return
// value will be populated with the ListProductQuotasCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProductQuotasCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProductQuotasCommon Send returns without error.
//
// See ListProductQuotas for more information on using the ListProductQuotas
// API call, and error handling.
//
//    // Example sending a request using the ListProductQuotasRequest method.
//    req, resp := client.ListProductQuotasRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *QUOTA) ListProductQuotasRequest(input *ListProductQuotasInput) (req *request.Request, output *ListProductQuotasOutput) {
	op := &request.Operation{
		Name:       opListProductQuotas,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProductQuotasInput{}
	}

	output = &ListProductQuotasOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListProductQuotas API operation for QUOTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for QUOTA's
// API operation ListProductQuotas for usage and error information.
func (c *QUOTA) ListProductQuotas(input *ListProductQuotasInput) (*ListProductQuotasOutput, error) {
	req, out := c.ListProductQuotasRequest(input)
	return out, req.Send()
}

// ListProductQuotasWithContext is the same as ListProductQuotas with the addition of
// the ability to pass a context and additional request options.
//
// See ListProductQuotas for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QUOTA) ListProductQuotasWithContext(ctx volcengine.Context, input *ListProductQuotasInput, opts ...request.Option) (*ListProductQuotasOutput, error) {
	req, out := c.ListProductQuotasRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DimensionForListProductQuotasInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForListProductQuotasInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForListProductQuotasInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForListProductQuotasInput) SetName(v string) *DimensionForListProductQuotasInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForListProductQuotasInput) SetValue(v string) *DimensionForListProductQuotasInput {
	s.Value = &v
	return s
}

type DimensionForListProductQuotasOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	NameCn *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`

	ValueCn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionForListProductQuotasOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForListProductQuotasOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForListProductQuotasOutput) SetName(v string) *DimensionForListProductQuotasOutput {
	s.Name = &v
	return s
}

// SetNameCn sets the NameCn field's value.
func (s *DimensionForListProductQuotasOutput) SetNameCn(v string) *DimensionForListProductQuotasOutput {
	s.NameCn = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForListProductQuotasOutput) SetValue(v string) *DimensionForListProductQuotasOutput {
	s.Value = &v
	return s
}

// SetValueCn sets the ValueCn field's value.
func (s *DimensionForListProductQuotasOutput) SetValueCn(v string) *DimensionForListProductQuotasOutput {
	s.ValueCn = &v
	return s
}

type ListProductQuotasInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Dimensions []*DimensionForListProductQuotasInput `type:"list" json:",omitempty"`

	MaxResults *int64 `type:"int64" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	// ProviderCode is a required field
	ProviderCode *string `type:"string" json:",omitempty" required:"true"`

	QuotaCode *string `type:"string" json:",omitempty"`

	QuotaType *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListProductQuotasInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProductQuotasInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProductQuotasInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListProductQuotasInput"}
	if s.ProviderCode == nil {
		invalidParams.Add(request.NewErrParamRequired("ProviderCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDimensions sets the Dimensions field's value.
func (s *ListProductQuotasInput) SetDimensions(v []*DimensionForListProductQuotasInput) *ListProductQuotasInput {
	s.Dimensions = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListProductQuotasInput) SetMaxResults(v int64) *ListProductQuotasInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListProductQuotasInput) SetNextToken(v string) *ListProductQuotasInput {
	s.NextToken = &v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *ListProductQuotasInput) SetProviderCode(v string) *ListProductQuotasInput {
	s.ProviderCode = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *ListProductQuotasInput) SetQuotaCode(v string) *ListProductQuotasInput {
	s.QuotaCode = &v
	return s
}

// SetQuotaType sets the QuotaType field's value.
func (s *ListProductQuotasInput) SetQuotaType(v string) *ListProductQuotasInput {
	s.QuotaType = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListProductQuotasInput) SetSortOrder(v string) *ListProductQuotasInput {
	s.SortOrder = &v
	return s
}

type ListProductQuotasOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string" json:",omitempty"`

	Quotas []*QuotaForListProductQuotasOutput `type:"list" json:",omitempty"`

	ResultsNum *int64 `type:"int64" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListProductQuotasOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProductQuotasOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListProductQuotasOutput) SetNextToken(v string) *ListProductQuotasOutput {
	s.NextToken = &v
	return s
}

// SetQuotas sets the Quotas field's value.
func (s *ListProductQuotasOutput) SetQuotas(v []*QuotaForListProductQuotasOutput) *ListProductQuotasOutput {
	s.Quotas = v
	return s
}

// SetResultsNum sets the ResultsNum field's value.
func (s *ListProductQuotasOutput) SetResultsNum(v int64) *ListProductQuotasOutput {
	s.ResultsNum = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListProductQuotasOutput) SetTotalCount(v int32) *ListProductQuotasOutput {
	s.TotalCount = &v
	return s
}

type QuotaForListProductQuotasOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApplicableType *string `type:"string" json:",omitempty"`

	ApplicableValue *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Dimensions []*DimensionForListProductQuotasOutput `type:"list" json:",omitempty"`

	ProviderCode *string `type:"string" json:",omitempty"`

	QuotaCode *string `type:"string" json:",omitempty"`

	QuotaType *string `type:"string" json:",omitempty"`

	TotalQuota *float64 `type:"double" json:",omitempty"`

	TotalUsage *TotalUsageForListProductQuotasOutput `type:"structure" json:",omitempty"`

	Trn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QuotaForListProductQuotasOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QuotaForListProductQuotasOutput) GoString() string {
	return s.String()
}

// SetApplicableType sets the ApplicableType field's value.
func (s *QuotaForListProductQuotasOutput) SetApplicableType(v string) *QuotaForListProductQuotasOutput {
	s.ApplicableType = &v
	return s
}

// SetApplicableValue sets the ApplicableValue field's value.
func (s *QuotaForListProductQuotasOutput) SetApplicableValue(v string) *QuotaForListProductQuotasOutput {
	s.ApplicableValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *QuotaForListProductQuotasOutput) SetDescription(v string) *QuotaForListProductQuotasOutput {
	s.Description = &v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *QuotaForListProductQuotasOutput) SetDimensions(v []*DimensionForListProductQuotasOutput) *QuotaForListProductQuotasOutput {
	s.Dimensions = v
	return s
}

// SetProviderCode sets the ProviderCode field's value.
func (s *QuotaForListProductQuotasOutput) SetProviderCode(v string) *QuotaForListProductQuotasOutput {
	s.ProviderCode = &v
	return s
}

// SetQuotaCode sets the QuotaCode field's value.
func (s *QuotaForListProductQuotasOutput) SetQuotaCode(v string) *QuotaForListProductQuotasOutput {
	s.QuotaCode = &v
	return s
}

// SetQuotaType sets the QuotaType field's value.
func (s *QuotaForListProductQuotasOutput) SetQuotaType(v string) *QuotaForListProductQuotasOutput {
	s.QuotaType = &v
	return s
}

// SetTotalQuota sets the TotalQuota field's value.
func (s *QuotaForListProductQuotasOutput) SetTotalQuota(v float64) *QuotaForListProductQuotasOutput {
	s.TotalQuota = &v
	return s
}

// SetTotalUsage sets the TotalUsage field's value.
func (s *QuotaForListProductQuotasOutput) SetTotalUsage(v *TotalUsageForListProductQuotasOutput) *QuotaForListProductQuotasOutput {
	s.TotalUsage = v
	return s
}

// SetTrn sets the Trn field's value.
func (s *QuotaForListProductQuotasOutput) SetTrn(v string) *QuotaForListProductQuotasOutput {
	s.Trn = &v
	return s
}

type TotalUsageForListProductQuotasOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Unit *string `type:"string" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s TotalUsageForListProductQuotasOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TotalUsageForListProductQuotasOutput) GoString() string {
	return s.String()
}

// SetUnit sets the Unit field's value.
func (s *TotalUsageForListProductQuotasOutput) SetUnit(v string) *TotalUsageForListProductQuotasOutput {
	s.Unit = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TotalUsageForListProductQuotasOutput) SetValue(v float64) *TotalUsageForListProductQuotasOutput {
	s.Value = &v
	return s
}
