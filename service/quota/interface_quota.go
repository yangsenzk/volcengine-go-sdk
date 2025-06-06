// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package quotaiface provides an interface to enable mocking the QUOTA service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package quota

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// QUOTAAPI provides an interface to enable mocking the
// quota.QUOTA service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // QUOTA.
//    func myFunc(svc QUOTAAPI) bool {
//        // Make svc.CreateAlarmRule request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := quota.New(sess)
//
//        myFunc(svc)
//    }
//
type QUOTAAPI interface {
	CreateAlarmRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlarmRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlarmRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlarmRule(*CreateAlarmRuleInput) (*CreateAlarmRuleOutput, error)
	CreateAlarmRuleWithContext(volcengine.Context, *CreateAlarmRuleInput, ...request.Option) (*CreateAlarmRuleOutput, error)
	CreateAlarmRuleRequest(*CreateAlarmRuleInput) (*request.Request, *CreateAlarmRuleOutput)

	CreateQuotaApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateQuotaApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateQuotaApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateQuotaApplication(*CreateQuotaApplicationInput) (*CreateQuotaApplicationOutput, error)
	CreateQuotaApplicationWithContext(volcengine.Context, *CreateQuotaApplicationInput, ...request.Option) (*CreateQuotaApplicationOutput, error)
	CreateQuotaApplicationRequest(*CreateQuotaApplicationInput) (*request.Request, *CreateQuotaApplicationOutput)

	CreateTemplateQuotaItemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTemplateQuotaItemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTemplateQuotaItemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTemplateQuotaItem(*CreateTemplateQuotaItemInput) (*CreateTemplateQuotaItemOutput, error)
	CreateTemplateQuotaItemWithContext(volcengine.Context, *CreateTemplateQuotaItemInput, ...request.Option) (*CreateTemplateQuotaItemOutput, error)
	CreateTemplateQuotaItemRequest(*CreateTemplateQuotaItemInput) (*request.Request, *CreateTemplateQuotaItemOutput)

	DeleteAlarmRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlarmRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlarmRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlarmRules(*DeleteAlarmRulesInput) (*DeleteAlarmRulesOutput, error)
	DeleteAlarmRulesWithContext(volcengine.Context, *DeleteAlarmRulesInput, ...request.Option) (*DeleteAlarmRulesOutput, error)
	DeleteAlarmRulesRequest(*DeleteAlarmRulesInput) (*request.Request, *DeleteAlarmRulesOutput)

	DeleteTemplateQuotaItemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTemplateQuotaItemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTemplateQuotaItemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTemplateQuotaItem(*DeleteTemplateQuotaItemInput) (*DeleteTemplateQuotaItemOutput, error)
	DeleteTemplateQuotaItemWithContext(volcengine.Context, *DeleteTemplateQuotaItemInput, ...request.Option) (*DeleteTemplateQuotaItemOutput, error)
	DeleteTemplateQuotaItemRequest(*DeleteTemplateQuotaItemInput) (*request.Request, *DeleteTemplateQuotaItemOutput)

	GetAlarmRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetAlarmRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetAlarmRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetAlarmRule(*GetAlarmRuleInput) (*GetAlarmRuleOutput, error)
	GetAlarmRuleWithContext(volcengine.Context, *GetAlarmRuleInput, ...request.Option) (*GetAlarmRuleOutput, error)
	GetAlarmRuleRequest(*GetAlarmRuleInput) (*request.Request, *GetAlarmRuleOutput)

	GetProductQuotaCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetProductQuotaCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetProductQuotaCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetProductQuota(*GetProductQuotaInput) (*GetProductQuotaOutput, error)
	GetProductQuotaWithContext(volcengine.Context, *GetProductQuotaInput, ...request.Option) (*GetProductQuotaOutput, error)
	GetProductQuotaRequest(*GetProductQuotaInput) (*request.Request, *GetProductQuotaOutput)

	GetQuotaApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetQuotaApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetQuotaApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetQuotaApplication(*GetQuotaApplicationInput) (*GetQuotaApplicationOutput, error)
	GetQuotaApplicationWithContext(volcengine.Context, *GetQuotaApplicationInput, ...request.Option) (*GetQuotaApplicationOutput, error)
	GetQuotaApplicationRequest(*GetQuotaApplicationInput) (*request.Request, *GetQuotaApplicationOutput)

	GetQuotaTemplateServiceStatusCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetQuotaTemplateServiceStatusCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetQuotaTemplateServiceStatusCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetQuotaTemplateServiceStatus(*GetQuotaTemplateServiceStatusInput) (*GetQuotaTemplateServiceStatusOutput, error)
	GetQuotaTemplateServiceStatusWithContext(volcengine.Context, *GetQuotaTemplateServiceStatusInput, ...request.Option) (*GetQuotaTemplateServiceStatusOutput, error)
	GetQuotaTemplateServiceStatusRequest(*GetQuotaTemplateServiceStatusInput) (*request.Request, *GetQuotaTemplateServiceStatusOutput)

	ListAlarmHistoryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAlarmHistoryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAlarmHistoryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAlarmHistory(*ListAlarmHistoryInput) (*ListAlarmHistoryOutput, error)
	ListAlarmHistoryWithContext(volcengine.Context, *ListAlarmHistoryInput, ...request.Option) (*ListAlarmHistoryOutput, error)
	ListAlarmHistoryRequest(*ListAlarmHistoryInput) (*request.Request, *ListAlarmHistoryOutput)

	ListProductQuotaDimensionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListProductQuotaDimensionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListProductQuotaDimensionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListProductQuotaDimensions(*ListProductQuotaDimensionsInput) (*ListProductQuotaDimensionsOutput, error)
	ListProductQuotaDimensionsWithContext(volcengine.Context, *ListProductQuotaDimensionsInput, ...request.Option) (*ListProductQuotaDimensionsOutput, error)
	ListProductQuotaDimensionsRequest(*ListProductQuotaDimensionsInput) (*request.Request, *ListProductQuotaDimensionsOutput)

	ListProductQuotasCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListProductQuotasCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListProductQuotasCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListProductQuotas(*ListProductQuotasInput) (*ListProductQuotasOutput, error)
	ListProductQuotasWithContext(volcengine.Context, *ListProductQuotasInput, ...request.Option) (*ListProductQuotasOutput, error)
	ListProductQuotasRequest(*ListProductQuotasInput) (*request.Request, *ListProductQuotasOutput)

	ListProductsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListProductsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListProductsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListProducts(*ListProductsInput) (*ListProductsOutput, error)
	ListProductsWithContext(volcengine.Context, *ListProductsInput, ...request.Option) (*ListProductsOutput, error)
	ListProductsRequest(*ListProductsInput) (*request.Request, *ListProductsOutput)

	ListQuotaAlarmRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListQuotaAlarmRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListQuotaAlarmRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListQuotaAlarmRules(*ListQuotaAlarmRulesInput) (*ListQuotaAlarmRulesOutput, error)
	ListQuotaAlarmRulesWithContext(volcengine.Context, *ListQuotaAlarmRulesInput, ...request.Option) (*ListQuotaAlarmRulesOutput, error)
	ListQuotaAlarmRulesRequest(*ListQuotaAlarmRulesInput) (*request.Request, *ListQuotaAlarmRulesOutput)

	ListQuotaApplicationTemplatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListQuotaApplicationTemplatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListQuotaApplicationTemplatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListQuotaApplicationTemplates(*ListQuotaApplicationTemplatesInput) (*ListQuotaApplicationTemplatesOutput, error)
	ListQuotaApplicationTemplatesWithContext(volcengine.Context, *ListQuotaApplicationTemplatesInput, ...request.Option) (*ListQuotaApplicationTemplatesOutput, error)
	ListQuotaApplicationTemplatesRequest(*ListQuotaApplicationTemplatesInput) (*request.Request, *ListQuotaApplicationTemplatesOutput)

	ListQuotaApplicationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListQuotaApplicationsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListQuotaApplicationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListQuotaApplications(*ListQuotaApplicationsInput) (*ListQuotaApplicationsOutput, error)
	ListQuotaApplicationsWithContext(volcengine.Context, *ListQuotaApplicationsInput, ...request.Option) (*ListQuotaApplicationsOutput, error)
	ListQuotaApplicationsRequest(*ListQuotaApplicationsInput) (*request.Request, *ListQuotaApplicationsOutput)

	ModifyQuotaTemplateServiceStatusCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyQuotaTemplateServiceStatusCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyQuotaTemplateServiceStatusCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyQuotaTemplateServiceStatus(*ModifyQuotaTemplateServiceStatusInput) (*ModifyQuotaTemplateServiceStatusOutput, error)
	ModifyQuotaTemplateServiceStatusWithContext(volcengine.Context, *ModifyQuotaTemplateServiceStatusInput, ...request.Option) (*ModifyQuotaTemplateServiceStatusOutput, error)
	ModifyQuotaTemplateServiceStatusRequest(*ModifyQuotaTemplateServiceStatusInput) (*request.Request, *ModifyQuotaTemplateServiceStatusOutput)

	ModifyTemplateQuotaItemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyTemplateQuotaItemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyTemplateQuotaItemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyTemplateQuotaItem(*ModifyTemplateQuotaItemInput) (*ModifyTemplateQuotaItemOutput, error)
	ModifyTemplateQuotaItemWithContext(volcengine.Context, *ModifyTemplateQuotaItemInput, ...request.Option) (*ModifyTemplateQuotaItemOutput, error)
	ModifyTemplateQuotaItemRequest(*ModifyTemplateQuotaItemInput) (*request.Request, *ModifyTemplateQuotaItemOutput)

	UpdateAlarmRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAlarmRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAlarmRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAlarmRule(*UpdateAlarmRuleInput) (*UpdateAlarmRuleOutput, error)
	UpdateAlarmRuleWithContext(volcengine.Context, *UpdateAlarmRuleInput, ...request.Option) (*UpdateAlarmRuleOutput, error)
	UpdateAlarmRuleRequest(*UpdateAlarmRuleInput) (*request.Request, *UpdateAlarmRuleOutput)
}

var _ QUOTAAPI = (*QUOTA)(nil)
