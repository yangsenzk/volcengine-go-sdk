// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcEndpointAttributesCommon = "DescribeVpcEndpointAttributes"

// DescribeVpcEndpointAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointAttributesCommon operation. The "output" return
// value will be populated with the DescribeVpcEndpointAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointAttributesCommon Send returns without error.
//
// See DescribeVpcEndpointAttributesCommon for more information on using the DescribeVpcEndpointAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointAttributesCommonRequest method.
//    req, resp := client.DescribeVpcEndpointAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointAttributesCommon,
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

// DescribeVpcEndpointAttributesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointAttributesCommon for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointAttributesCommonWithContext is the same as DescribeVpcEndpointAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcEndpointAttributes = "DescribeVpcEndpointAttributes"

// DescribeVpcEndpointAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointAttributes operation. The "output" return
// value will be populated with the DescribeVpcEndpointAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointAttributesCommon Send returns without error.
//
// See DescribeVpcEndpointAttributes for more information on using the DescribeVpcEndpointAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointAttributesRequest method.
//    req, resp := client.DescribeVpcEndpointAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointAttributesRequest(input *DescribeVpcEndpointAttributesInput) (req *request.Request, output *DescribeVpcEndpointAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcEndpointAttributesInput{}
	}

	output = &DescribeVpcEndpointAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcEndpointAttributes API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointAttributes for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointAttributes(input *DescribeVpcEndpointAttributesInput) (*DescribeVpcEndpointAttributesOutput, error) {
	req, out := c.DescribeVpcEndpointAttributesRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointAttributesWithContext is the same as DescribeVpcEndpointAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointAttributesWithContext(ctx volcengine.Context, input *DescribeVpcEndpointAttributesInput, opts ...request.Option) (*DescribeVpcEndpointAttributesOutput, error) {
	req, out := c.DescribeVpcEndpointAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpcEndpointAttributesInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpcEndpointAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcEndpointAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpcEndpointAttributesInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *DescribeVpcEndpointAttributesInput) SetEndpointId(v string) *DescribeVpcEndpointAttributesInput {
	s.EndpointId = &v
	return s
}

type DescribeVpcEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Bps *int32 `type:"int32"`

	BusinessStatus *string `type:"string"`

	ConnectionStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	Endpoint *EndpointForDescribeVpcEndpointAttributesOutput `type:"structure"`

	EndpointDomain *string `type:"string"`

	EndpointId *string `type:"string"`

	EndpointIndex *int32 `type:"int32"`

	EndpointName *string `type:"string"`

	EndpointType *string `type:"string"`

	IpAddressVersions []*string `type:"list"`

	IpAddressVersionsN []*string `type:"list" json:"IpAddressVersions.N"`

	Payer *string `type:"string"`

	PrivateDNSEnabled *bool `type:"boolean"`

	PrivateDNSName *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceManaged *bool `type:"boolean"`

	ServiceName *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeVpcEndpointAttributesOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointAttributesOutput) GoString() string {
	return s.String()
}

// SetBps sets the Bps field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetBps(v int32) *DescribeVpcEndpointAttributesOutput {
	s.Bps = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetBusinessStatus(v string) *DescribeVpcEndpointAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionStatus sets the ConnectionStatus field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetConnectionStatus(v string) *DescribeVpcEndpointAttributesOutput {
	s.ConnectionStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetCreationTime(v string) *DescribeVpcEndpointAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetDeletedTime(v string) *DescribeVpcEndpointAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetDescription(v string) *DescribeVpcEndpointAttributesOutput {
	s.Description = &v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpoint(v *EndpointForDescribeVpcEndpointAttributesOutput) *DescribeVpcEndpointAttributesOutput {
	s.Endpoint = v
	return s
}

// SetEndpointDomain sets the EndpointDomain field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpointDomain(v string) *DescribeVpcEndpointAttributesOutput {
	s.EndpointDomain = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpointId(v string) *DescribeVpcEndpointAttributesOutput {
	s.EndpointId = &v
	return s
}

// SetEndpointIndex sets the EndpointIndex field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpointIndex(v int32) *DescribeVpcEndpointAttributesOutput {
	s.EndpointIndex = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpointName(v string) *DescribeVpcEndpointAttributesOutput {
	s.EndpointName = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetEndpointType(v string) *DescribeVpcEndpointAttributesOutput {
	s.EndpointType = &v
	return s
}

// SetIpAddressVersions sets the IpAddressVersions field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetIpAddressVersions(v []*string) *DescribeVpcEndpointAttributesOutput {
	s.IpAddressVersions = v
	return s
}

// SetIpAddressVersionsN sets the IpAddressVersionsN field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetIpAddressVersionsN(v []*string) *DescribeVpcEndpointAttributesOutput {
	s.IpAddressVersionsN = v
	return s
}

// SetPayer sets the Payer field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetPayer(v string) *DescribeVpcEndpointAttributesOutput {
	s.Payer = &v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetPrivateDNSEnabled(v bool) *DescribeVpcEndpointAttributesOutput {
	s.PrivateDNSEnabled = &v
	return s
}

// SetPrivateDNSName sets the PrivateDNSName field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetPrivateDNSName(v string) *DescribeVpcEndpointAttributesOutput {
	s.PrivateDNSName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetProjectName(v string) *DescribeVpcEndpointAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetRequestId(v string) *DescribeVpcEndpointAttributesOutput {
	s.RequestId = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetServiceId(v string) *DescribeVpcEndpointAttributesOutput {
	s.ServiceId = &v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetServiceManaged(v bool) *DescribeVpcEndpointAttributesOutput {
	s.ServiceManaged = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetServiceName(v string) *DescribeVpcEndpointAttributesOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetStatus(v string) *DescribeVpcEndpointAttributesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetTags(v []*TagForDescribeVpcEndpointAttributesOutput) *DescribeVpcEndpointAttributesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetUpdateTime(v string) *DescribeVpcEndpointAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeVpcEndpointAttributesOutput) SetVpcId(v string) *DescribeVpcEndpointAttributesOutput {
	s.VpcId = &v
	return s
}

type EndpointForDescribeVpcEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	Bps *int32 `type:"int32"`

	BusinessStatus *string `type:"string"`

	ConnectionStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EndpointDomain *string `type:"string"`

	EndpointId *string `type:"string"`

	EndpointIndex *int32 `type:"int32"`

	EndpointName *string `type:"string"`

	EndpointType *string `type:"string"`

	IpAddressVersions []*string `type:"list"`

	Payer *string `type:"string"`

	PrivateDNSEnabled *bool `type:"boolean"`

	PrivateDNSName *string `type:"string"`

	ProjectName *string `type:"string"`

	ServiceId *string `type:"string"`

	ServiceManaged *bool `type:"boolean"`

	ServiceName *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeVpcEndpointAttributesOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s EndpointForDescribeVpcEndpointAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointForDescribeVpcEndpointAttributesOutput) GoString() string {
	return s.String()
}

// SetBps sets the Bps field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetBps(v int32) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.Bps = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetBusinessStatus(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionStatus sets the ConnectionStatus field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetConnectionStatus(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.ConnectionStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetCreationTime(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetDeletedTime(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetDescription(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.Description = &v
	return s
}

// SetEndpointDomain sets the EndpointDomain field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetEndpointDomain(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.EndpointDomain = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetEndpointId(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.EndpointId = &v
	return s
}

// SetEndpointIndex sets the EndpointIndex field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetEndpointIndex(v int32) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.EndpointIndex = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetEndpointName(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.EndpointName = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetEndpointType(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.EndpointType = &v
	return s
}

// SetIpAddressVersions sets the IpAddressVersions field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetIpAddressVersions(v []*string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.IpAddressVersions = v
	return s
}

// SetPayer sets the Payer field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetPayer(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.Payer = &v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetPrivateDNSEnabled(v bool) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.PrivateDNSEnabled = &v
	return s
}

// SetPrivateDNSName sets the PrivateDNSName field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetPrivateDNSName(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.PrivateDNSName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetProjectName(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetServiceId(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.ServiceId = &v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetServiceManaged(v bool) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.ServiceManaged = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetServiceName(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetStatus(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetTags(v []*TagForDescribeVpcEndpointAttributesOutput) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetUpdateTime(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *EndpointForDescribeVpcEndpointAttributesOutput) SetVpcId(v string) *EndpointForDescribeVpcEndpointAttributesOutput {
	s.VpcId = &v
	return s
}

type TagForDescribeVpcEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeVpcEndpointAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeVpcEndpointAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeVpcEndpointAttributesOutput) SetKey(v string) *TagForDescribeVpcEndpointAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeVpcEndpointAttributesOutput) SetValue(v string) *TagForDescribeVpcEndpointAttributesOutput {
	s.Value = &v
	return s
}
