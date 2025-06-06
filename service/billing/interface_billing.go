// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package billingiface provides an interface to enable mocking the BILLING service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// BILLINGAPI provides an interface to enable mocking the
// billing.BILLING service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // BILLING.
//    func myFunc(svc BILLINGAPI) bool {
//        // Make svc.CancelInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := billing.New(sess)
//
//        myFunc(svc)
//    }
//
type BILLINGAPI interface {
	CancelInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelInvitation(*CancelInvitationInput) (*CancelInvitationOutput, error)
	CancelInvitationWithContext(volcengine.Context, *CancelInvitationInput, ...request.Option) (*CancelInvitationOutput, error)
	CancelInvitationRequest(*CancelInvitationInput) (*request.Request, *CancelInvitationOutput)

	CancelOrderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelOrderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelOrderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelOrder(*CancelOrderInput) (*CancelOrderOutput, error)
	CancelOrderWithContext(volcengine.Context, *CancelOrderInput, ...request.Option) (*CancelOrderOutput, error)
	CancelOrderRequest(*CancelOrderInput) (*request.Request, *CancelOrderOutput)

	CleanUpFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CleanUpFinancialRelationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CleanUpFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CleanUpFinancialRelation(*CleanUpFinancialRelationInput) (*CleanUpFinancialRelationOutput, error)
	CleanUpFinancialRelationWithContext(volcengine.Context, *CleanUpFinancialRelationInput, ...request.Option) (*CleanUpFinancialRelationOutput, error)
	CleanUpFinancialRelationRequest(*CleanUpFinancialRelationInput) (*request.Request, *CleanUpFinancialRelationOutput)

	CreateFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateFinancialRelationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateFinancialRelation(*CreateFinancialRelationInput) (*CreateFinancialRelationOutput, error)
	CreateFinancialRelationWithContext(volcengine.Context, *CreateFinancialRelationInput, ...request.Option) (*CreateFinancialRelationOutput, error)
	CreateFinancialRelationRequest(*CreateFinancialRelationInput) (*request.Request, *CreateFinancialRelationOutput)

	DeleteFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteFinancialRelationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteFinancialRelation(*DeleteFinancialRelationInput) (*DeleteFinancialRelationOutput, error)
	DeleteFinancialRelationWithContext(volcengine.Context, *DeleteFinancialRelationInput, ...request.Option) (*DeleteFinancialRelationOutput, error)
	DeleteFinancialRelationRequest(*DeleteFinancialRelationInput) (*request.Request, *DeleteFinancialRelationOutput)

	GetOrderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetOrderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetOrderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetOrder(*GetOrderInput) (*GetOrderOutput, error)
	GetOrderWithContext(volcengine.Context, *GetOrderInput, ...request.Option) (*GetOrderOutput, error)
	GetOrderRequest(*GetOrderInput) (*request.Request, *GetOrderOutput)

	HandleInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	HandleInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HandleInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HandleInvitation(*HandleInvitationInput) (*HandleInvitationOutput, error)
	HandleInvitationWithContext(volcengine.Context, *HandleInvitationInput, ...request.Option) (*HandleInvitationOutput, error)
	HandleInvitationRequest(*HandleInvitationInput) (*request.Request, *HandleInvitationOutput)

	ListAmortizedCostBillDailyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAmortizedCostBillDailyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAmortizedCostBillDailyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAmortizedCostBillDaily(*ListAmortizedCostBillDailyInput) (*ListAmortizedCostBillDailyOutput, error)
	ListAmortizedCostBillDailyWithContext(volcengine.Context, *ListAmortizedCostBillDailyInput, ...request.Option) (*ListAmortizedCostBillDailyOutput, error)
	ListAmortizedCostBillDailyRequest(*ListAmortizedCostBillDailyInput) (*request.Request, *ListAmortizedCostBillDailyOutput)

	ListAmortizedCostBillDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAmortizedCostBillDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAmortizedCostBillDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAmortizedCostBillDetail(*ListAmortizedCostBillDetailInput) (*ListAmortizedCostBillDetailOutput, error)
	ListAmortizedCostBillDetailWithContext(volcengine.Context, *ListAmortizedCostBillDetailInput, ...request.Option) (*ListAmortizedCostBillDetailOutput, error)
	ListAmortizedCostBillDetailRequest(*ListAmortizedCostBillDetailInput) (*request.Request, *ListAmortizedCostBillDetailOutput)

	ListAmortizedCostBillMonthlyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAmortizedCostBillMonthlyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAmortizedCostBillMonthlyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAmortizedCostBillMonthly(*ListAmortizedCostBillMonthlyInput) (*ListAmortizedCostBillMonthlyOutput, error)
	ListAmortizedCostBillMonthlyWithContext(volcengine.Context, *ListAmortizedCostBillMonthlyInput, ...request.Option) (*ListAmortizedCostBillMonthlyOutput, error)
	ListAmortizedCostBillMonthlyRequest(*ListAmortizedCostBillMonthlyInput) (*request.Request, *ListAmortizedCostBillMonthlyOutput)

	ListAvailableInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAvailableInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAvailableInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAvailableInstances(*ListAvailableInstancesInput) (*ListAvailableInstancesOutput, error)
	ListAvailableInstancesWithContext(volcengine.Context, *ListAvailableInstancesInput, ...request.Option) (*ListAvailableInstancesOutput, error)
	ListAvailableInstancesRequest(*ListAvailableInstancesInput) (*request.Request, *ListAvailableInstancesOutput)

	ListBillCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBill(*ListBillInput) (*ListBillOutput, error)
	ListBillWithContext(volcengine.Context, *ListBillInput, ...request.Option) (*ListBillOutput, error)
	ListBillRequest(*ListBillInput) (*request.Request, *ListBillOutput)

	ListBillDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillDetail(*ListBillDetailInput) (*ListBillDetailOutput, error)
	ListBillDetailWithContext(volcengine.Context, *ListBillDetailInput, ...request.Option) (*ListBillDetailOutput, error)
	ListBillDetailRequest(*ListBillDetailInput) (*request.Request, *ListBillDetailOutput)

	ListBillOverviewByCategoryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillOverviewByCategoryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillOverviewByCategoryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillOverviewByCategory(*ListBillOverviewByCategoryInput) (*ListBillOverviewByCategoryOutput, error)
	ListBillOverviewByCategoryWithContext(volcengine.Context, *ListBillOverviewByCategoryInput, ...request.Option) (*ListBillOverviewByCategoryOutput, error)
	ListBillOverviewByCategoryRequest(*ListBillOverviewByCategoryInput) (*request.Request, *ListBillOverviewByCategoryOutput)

	ListBillOverviewByProdCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillOverviewByProdCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillOverviewByProdCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillOverviewByProd(*ListBillOverviewByProdInput) (*ListBillOverviewByProdOutput, error)
	ListBillOverviewByProdWithContext(volcengine.Context, *ListBillOverviewByProdInput, ...request.Option) (*ListBillOverviewByProdOutput, error)
	ListBillOverviewByProdRequest(*ListBillOverviewByProdInput) (*request.Request, *ListBillOverviewByProdOutput)

	ListFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListFinancialRelationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListFinancialRelation(*ListFinancialRelationInput) (*ListFinancialRelationOutput, error)
	ListFinancialRelationWithContext(volcengine.Context, *ListFinancialRelationInput, ...request.Option) (*ListFinancialRelationOutput, error)
	ListFinancialRelationRequest(*ListFinancialRelationInput) (*request.Request, *ListFinancialRelationOutput)

	ListInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListInvitationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListInvitation(*ListInvitationInput) (*ListInvitationOutput, error)
	ListInvitationWithContext(volcengine.Context, *ListInvitationInput, ...request.Option) (*ListInvitationOutput, error)
	ListInvitationRequest(*ListInvitationInput) (*request.Request, *ListInvitationOutput)

	ListOrderProductDetailsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListOrderProductDetailsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListOrderProductDetailsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListOrderProductDetails(*ListOrderProductDetailsInput) (*ListOrderProductDetailsOutput, error)
	ListOrderProductDetailsWithContext(volcengine.Context, *ListOrderProductDetailsInput, ...request.Option) (*ListOrderProductDetailsOutput, error)
	ListOrderProductDetailsRequest(*ListOrderProductDetailsInput) (*request.Request, *ListOrderProductDetailsOutput)

	ListOrdersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListOrdersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListOrdersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListOrders(*ListOrdersInput) (*ListOrdersOutput, error)
	ListOrdersWithContext(volcengine.Context, *ListOrdersInput, ...request.Option) (*ListOrdersOutput, error)
	ListOrdersRequest(*ListOrdersInput) (*request.Request, *ListOrdersOutput)

	ListPackageUsageDetailsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPackageUsageDetailsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPackageUsageDetailsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPackageUsageDetails(*ListPackageUsageDetailsInput) (*ListPackageUsageDetailsOutput, error)
	ListPackageUsageDetailsWithContext(volcengine.Context, *ListPackageUsageDetailsInput, ...request.Option) (*ListPackageUsageDetailsOutput, error)
	ListPackageUsageDetailsRequest(*ListPackageUsageDetailsInput) (*request.Request, *ListPackageUsageDetailsOutput)

	ListResourcePackagesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListResourcePackagesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListResourcePackagesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListResourcePackages(*ListResourcePackagesInput) (*ListResourcePackagesOutput, error)
	ListResourcePackagesWithContext(volcengine.Context, *ListResourcePackagesInput, ...request.Option) (*ListResourcePackagesOutput, error)
	ListResourcePackagesRequest(*ListResourcePackagesInput) (*request.Request, *ListResourcePackagesOutput)

	ListSplitBillDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListSplitBillDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListSplitBillDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListSplitBillDetail(*ListSplitBillDetailInput) (*ListSplitBillDetailOutput, error)
	ListSplitBillDetailWithContext(volcengine.Context, *ListSplitBillDetailInput, ...request.Option) (*ListSplitBillDetailOutput, error)
	ListSplitBillDetailRequest(*ListSplitBillDetailInput) (*request.Request, *ListSplitBillDetailOutput)

	PayOrderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	PayOrderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	PayOrderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	PayOrder(*PayOrderInput) (*PayOrderOutput, error)
	PayOrderWithContext(volcengine.Context, *PayOrderInput, ...request.Option) (*PayOrderOutput, error)
	PayOrderRequest(*PayOrderInput) (*request.Request, *PayOrderOutput)

	QueryBalanceAcctCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryBalanceAcctCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryBalanceAcctCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryBalanceAcct(*QueryBalanceAcctInput) (*QueryBalanceAcctOutput, error)
	QueryBalanceAcctWithContext(volcengine.Context, *QueryBalanceAcctInput, ...request.Option) (*QueryBalanceAcctOutput, error)
	QueryBalanceAcctRequest(*QueryBalanceAcctInput) (*request.Request, *QueryBalanceAcctOutput)

	QueryPriceForPayAsYouGoCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryPriceForPayAsYouGoCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryPriceForPayAsYouGoCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryPriceForPayAsYouGo(*QueryPriceForPayAsYouGoInput) (*QueryPriceForPayAsYouGoOutput, error)
	QueryPriceForPayAsYouGoWithContext(volcengine.Context, *QueryPriceForPayAsYouGoInput, ...request.Option) (*QueryPriceForPayAsYouGoOutput, error)
	QueryPriceForPayAsYouGoRequest(*QueryPriceForPayAsYouGoInput) (*request.Request, *QueryPriceForPayAsYouGoOutput)

	QueryPriceForRenewCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryPriceForRenewCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryPriceForRenewCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryPriceForRenew(*QueryPriceForRenewInput) (*QueryPriceForRenewOutput, error)
	QueryPriceForRenewWithContext(volcengine.Context, *QueryPriceForRenewInput, ...request.Option) (*QueryPriceForRenewOutput, error)
	QueryPriceForRenewRequest(*QueryPriceForRenewInput) (*request.Request, *QueryPriceForRenewOutput)

	QueryPriceForSubscriptionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryPriceForSubscriptionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryPriceForSubscriptionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryPriceForSubscription(*QueryPriceForSubscriptionInput) (*QueryPriceForSubscriptionOutput, error)
	QueryPriceForSubscriptionWithContext(volcengine.Context, *QueryPriceForSubscriptionInput, ...request.Option) (*QueryPriceForSubscriptionOutput, error)
	QueryPriceForSubscriptionRequest(*QueryPriceForSubscriptionInput) (*request.Request, *QueryPriceForSubscriptionOutput)

	RenewInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RenewInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenewInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenewInstance(*RenewInstanceInput) (*RenewInstanceOutput, error)
	RenewInstanceWithContext(volcengine.Context, *RenewInstanceInput, ...request.Option) (*RenewInstanceOutput, error)
	RenewInstanceRequest(*RenewInstanceInput) (*request.Request, *RenewInstanceOutput)

	SetRenewalTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetRenewalTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetRenewalTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetRenewalType(*SetRenewalTypeInput) (*SetRenewalTypeOutput, error)
	SetRenewalTypeWithContext(volcengine.Context, *SetRenewalTypeInput, ...request.Option) (*SetRenewalTypeOutput, error)
	SetRenewalTypeRequest(*SetRenewalTypeInput) (*request.Request, *SetRenewalTypeOutput)

	UnsubscribeInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UnsubscribeInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UnsubscribeInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UnsubscribeInstance(*UnsubscribeInstanceInput) (*UnsubscribeInstanceOutput, error)
	UnsubscribeInstanceWithContext(volcengine.Context, *UnsubscribeInstanceInput, ...request.Option) (*UnsubscribeInstanceOutput, error)
	UnsubscribeInstanceRequest(*UnsubscribeInstanceInput) (*request.Request, *UnsubscribeInstanceOutput)

	UpdateAuthCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAuthCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAuthCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAuth(*UpdateAuthInput) (*UpdateAuthOutput, error)
	UpdateAuthWithContext(volcengine.Context, *UpdateAuthInput, ...request.Option) (*UpdateAuthOutput, error)
	UpdateAuthRequest(*UpdateAuthInput) (*request.Request, *UpdateAuthOutput)
}

var _ BILLINGAPI = (*BILLING)(nil)
