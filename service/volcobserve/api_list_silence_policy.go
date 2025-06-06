// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSilencePolicyCommon = "ListSilencePolicy"

// ListSilencePolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSilencePolicyCommon operation. The "output" return
// value will be populated with the ListSilencePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSilencePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSilencePolicyCommon Send returns without error.
//
// See ListSilencePolicyCommon for more information on using the ListSilencePolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSilencePolicyCommonRequest method.
//    req, resp := client.ListSilencePolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListSilencePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSilencePolicyCommon,
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

// ListSilencePolicyCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListSilencePolicyCommon for usage and error information.
func (c *VOLCOBSERVE) ListSilencePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSilencePolicyCommonRequest(input)
	return out, req.Send()
}

// ListSilencePolicyCommonWithContext is the same as ListSilencePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSilencePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListSilencePolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSilencePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSilencePolicy = "ListSilencePolicy"

// ListSilencePolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSilencePolicy operation. The "output" return
// value will be populated with the ListSilencePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSilencePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSilencePolicyCommon Send returns without error.
//
// See ListSilencePolicy for more information on using the ListSilencePolicy
// API call, and error handling.
//
//    // Example sending a request using the ListSilencePolicyRequest method.
//    req, resp := client.ListSilencePolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListSilencePolicyRequest(input *ListSilencePolicyInput) (req *request.Request, output *ListSilencePolicyOutput) {
	op := &request.Operation{
		Name:       opListSilencePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSilencePolicyInput{}
	}

	output = &ListSilencePolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListSilencePolicy API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListSilencePolicy for usage and error information.
func (c *VOLCOBSERVE) ListSilencePolicy(input *ListSilencePolicyInput) (*ListSilencePolicyOutput, error) {
	req, out := c.ListSilencePolicyRequest(input)
	return out, req.Send()
}

// ListSilencePolicyWithContext is the same as ListSilencePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ListSilencePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListSilencePolicyWithContext(ctx volcengine.Context, input *ListSilencePolicyInput, opts ...request.Option) (*ListSilencePolicyOutput, error) {
	req, out := c.ListSilencePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	EffectTime *EffectTimeForListSilencePolicyOutput `type:"structure" json:",omitempty"`

	EnableState *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	SilenceConditions *SilenceConditionsForListSilencePolicyOutput `type:"structure" json:",omitempty"`

	SilenceType *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`

	ValidReason *string `type:"string" json:",omitempty"`

	ValidState *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListSilencePolicyOutput) SetCreatedAt(v string) *DataForListSilencePolicyOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListSilencePolicyOutput) SetDescription(v string) *DataForListSilencePolicyOutput {
	s.Description = &v
	return s
}

// SetEffectTime sets the EffectTime field's value.
func (s *DataForListSilencePolicyOutput) SetEffectTime(v *EffectTimeForListSilencePolicyOutput) *DataForListSilencePolicyOutput {
	s.EffectTime = v
	return s
}

// SetEnableState sets the EnableState field's value.
func (s *DataForListSilencePolicyOutput) SetEnableState(v string) *DataForListSilencePolicyOutput {
	s.EnableState = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListSilencePolicyOutput) SetId(v string) *DataForListSilencePolicyOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListSilencePolicyOutput) SetName(v string) *DataForListSilencePolicyOutput {
	s.Name = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForListSilencePolicyOutput) SetNamespace(v string) *DataForListSilencePolicyOutput {
	s.Namespace = &v
	return s
}

// SetSilenceConditions sets the SilenceConditions field's value.
func (s *DataForListSilencePolicyOutput) SetSilenceConditions(v *SilenceConditionsForListSilencePolicyOutput) *DataForListSilencePolicyOutput {
	s.SilenceConditions = v
	return s
}

// SetSilenceType sets the SilenceType field's value.
func (s *DataForListSilencePolicyOutput) SetSilenceType(v string) *DataForListSilencePolicyOutput {
	s.SilenceType = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListSilencePolicyOutput) SetUpdatedAt(v string) *DataForListSilencePolicyOutput {
	s.UpdatedAt = &v
	return s
}

// SetValidReason sets the ValidReason field's value.
func (s *DataForListSilencePolicyOutput) SetValidReason(v string) *DataForListSilencePolicyOutput {
	s.ValidReason = &v
	return s
}

// SetValidState sets the ValidState field's value.
func (s *DataForListSilencePolicyOutput) SetValidState(v string) *DataForListSilencePolicyOutput {
	s.ValidState = &v
	return s
}

type EffectTimeForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Range []*RangeForListSilencePolicyOutput `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EffectTimeForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EffectTimeForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetRange sets the Range field's value.
func (s *EffectTimeForListSilencePolicyOutput) SetRange(v []*RangeForListSilencePolicyOutput) *EffectTimeForListSilencePolicyOutput {
	s.Range = v
	return s
}

// SetType sets the Type field's value.
func (s *EffectTimeForListSilencePolicyOutput) SetType(v string) *EffectTimeForListSilencePolicyOutput {
	s.Type = &v
	return s
}

type ListSilencePolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EnableState []*string `type:"list" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Namespaces []*string `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	SilenceType []*string `type:"list" json:",omitempty"`

	ValidState *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListSilencePolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSilencePolicyInput) GoString() string {
	return s.String()
}

// SetEnableState sets the EnableState field's value.
func (s *ListSilencePolicyInput) SetEnableState(v []*string) *ListSilencePolicyInput {
	s.EnableState = v
	return s
}

// SetIds sets the Ids field's value.
func (s *ListSilencePolicyInput) SetIds(v []*string) *ListSilencePolicyInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *ListSilencePolicyInput) SetName(v string) *ListSilencePolicyInput {
	s.Name = &v
	return s
}

// SetNamespaces sets the Namespaces field's value.
func (s *ListSilencePolicyInput) SetNamespaces(v []*string) *ListSilencePolicyInput {
	s.Namespaces = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListSilencePolicyInput) SetPageNumber(v int64) *ListSilencePolicyInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListSilencePolicyInput) SetPageSize(v int64) *ListSilencePolicyInput {
	s.PageSize = &v
	return s
}

// SetSilenceType sets the SilenceType field's value.
func (s *ListSilencePolicyInput) SetSilenceType(v []*string) *ListSilencePolicyInput {
	s.SilenceType = v
	return s
}

// SetValidState sets the ValidState field's value.
func (s *ListSilencePolicyInput) SetValidState(v string) *ListSilencePolicyInput {
	s.ValidState = &v
	return s
}

type ListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListSilencePolicyOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListSilencePolicyOutput) SetData(v []*DataForListSilencePolicyOutput) *ListSilencePolicyOutput {
	s.Data = v
	return s
}

type MetaConditionForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`

	Metas []*MetaForListSilencePolicyOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaConditionForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaConditionForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetCondition sets the Condition field's value.
func (s *MetaConditionForListSilencePolicyOutput) SetCondition(v string) *MetaConditionForListSilencePolicyOutput {
	s.Condition = &v
	return s
}

// SetMetas sets the Metas field's value.
func (s *MetaConditionForListSilencePolicyOutput) SetMetas(v []*MetaForListSilencePolicyOutput) *MetaConditionForListSilencePolicyOutput {
	s.Metas = v
	return s
}

type MetaForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comparator *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetComparator sets the Comparator field's value.
func (s *MetaForListSilencePolicyOutput) SetComparator(v string) *MetaForListSilencePolicyOutput {
	s.Comparator = &v
	return s
}

// SetKey sets the Key field's value.
func (s *MetaForListSilencePolicyOutput) SetKey(v string) *MetaForListSilencePolicyOutput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *MetaForListSilencePolicyOutput) SetValues(v []*string) *MetaForListSilencePolicyOutput {
	s.Values = v
	return s
}

type RangeForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndDate *string `type:"string" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	StartDate *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RangeForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RangeForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetEndDate sets the EndDate field's value.
func (s *RangeForListSilencePolicyOutput) SetEndDate(v string) *RangeForListSilencePolicyOutput {
	s.EndDate = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *RangeForListSilencePolicyOutput) SetEndTime(v string) *RangeForListSilencePolicyOutput {
	s.EndTime = &v
	return s
}

// SetStartDate sets the StartDate field's value.
func (s *RangeForListSilencePolicyOutput) SetStartDate(v string) *RangeForListSilencePolicyOutput {
	s.StartDate = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *RangeForListSilencePolicyOutput) SetStartTime(v string) *RangeForListSilencePolicyOutput {
	s.StartTime = &v
	return s
}

type SilenceConditionsForListSilencePolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EventType []*string `type:"list" json:",omitempty"`

	MetaCondition *MetaConditionForListSilencePolicyOutput `type:"structure" json:",omitempty"`

	RuleId *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SilenceConditionsForListSilencePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SilenceConditionsForListSilencePolicyOutput) GoString() string {
	return s.String()
}

// SetEventType sets the EventType field's value.
func (s *SilenceConditionsForListSilencePolicyOutput) SetEventType(v []*string) *SilenceConditionsForListSilencePolicyOutput {
	s.EventType = v
	return s
}

// SetMetaCondition sets the MetaCondition field's value.
func (s *SilenceConditionsForListSilencePolicyOutput) SetMetaCondition(v *MetaConditionForListSilencePolicyOutput) *SilenceConditionsForListSilencePolicyOutput {
	s.MetaCondition = v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *SilenceConditionsForListSilencePolicyOutput) SetRuleId(v string) *SilenceConditionsForListSilencePolicyOutput {
	s.RuleId = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *SilenceConditionsForListSilencePolicyOutput) SetRuleName(v string) *SilenceConditionsForListSilencePolicyOutput {
	s.RuleName = &v
	return s
}
