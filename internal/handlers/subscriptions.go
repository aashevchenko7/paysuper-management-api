package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	subscriptionsPath       = "/subscriptions"
	subscriptionDetailsPath = "/subscriptions/:id"
	subscriptionOrdersPath  = "/subscriptions/:id/orders"

	merchantIdSubscriptionsPath       = "/merchants/:merchant_id/subscriptions"
	merchantIdSubscriptionDetailsPath = "/merchants/:merchant_id/subscriptions/:id"
	merchantIdSubscriptionOrdersPath  = "/merchants/:merchant_id/subscriptions/:id/orders"
)

type SubscriptionsRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewSubscriptionsRoute(set common.HandlerSet, cfg *common.Config) *SubscriptionsRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "SubscriptionsRoute"})
	return &SubscriptionsRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *SubscriptionsRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(subscriptionsPath, h.getSubscriptions)
	groups.AuthUser.GET(subscriptionDetailsPath, h.getSubscription)
	groups.AuthUser.DELETE(subscriptionDetailsPath, h.deleteSubscription)
	groups.AuthUser.GET(subscriptionOrdersPath, h.getSubscriptionOrders)

	groups.SystemUser.GET(merchantIdSubscriptionsPath, h.getSubscriptions)
	groups.SystemUser.GET(merchantIdSubscriptionDetailsPath, h.getSubscription)
	groups.SystemUser.DELETE(merchantIdSubscriptionDetailsPath, h.deleteSubscription)
	groups.SystemUser.GET(merchantIdSubscriptionOrdersPath, h.getSubscriptionOrders)
}

// @summary Get subscriptions for merchant
// @desc Get subscriptions for merchant
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantSubscriptionsResponse Returns the merchant subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/subscriptions [get]

// @summary Get subscriptions by merchant for admin
// @desc Get subscriptions by merchant for admin
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantSubscriptionsResponse Returns the merchant subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/merchants/{merchant_id}/subscriptions [get]
func (h *SubscriptionsRoute) getSubscriptions(ctx echo.Context) error {
	req := &billingpb.FindSubscriptionsRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.FindSubscriptions(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "FindSubscriptions")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get subscription for admin
// @desc Get subscription for admin
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetSubscriptionResponse Returns subscription info
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/merchants/:merchant_id/subscriptions/:id [get]

// @summary Get subscription for merchant
// @desc Get subscription for merchant
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetSubscriptionResponse Returns subscription info
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/subscriptions/:id [get]
func (h *SubscriptionsRoute) getSubscription(ctx echo.Context) error {
	req := &billingpb.GetSubscriptionRequest{
		Id: ctx.Param(common.RequestParameterId),
	}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetSubscription(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, "recurringpb", "GetSubscription", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get subscription orders for admin
// @desc Get subscription orders for admin
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetSubscriptionOrdersResponse Returns list of orders for subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/merchants/:merchant_id/subscriptions/:id/orders [get]

// @summary Get subscription orders for merchant
// @desc Get subscription orders for merchant
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetSubscriptionOrdersResponse Returns list of orders for subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/subscriptions/:id/orders [get]
func (h *SubscriptionsRoute) getSubscriptionOrders(ctx echo.Context) error {
	req := &billingpb.GetSubscriptionOrdersRequest{
		Id: ctx.Param(common.RequestParameterId),
	}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetSubscriptionOrders(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetSubscriptionOrders")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete subscription for admin
// @desc Delete subscription for admin
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns list of orders for subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/merchants/:merchant_id/subscriptions/:id [delete]

// @summary Delete subscription for merchant
// @desc Delete subscription for merchant
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns result for deleting subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/subscriptions/:id [delete]
func (h *SubscriptionsRoute) deleteSubscription(ctx echo.Context) error {
	req := &billingpb.DeleteRecurringSubscriptionRequest{
		Id: ctx.Param(common.RequestParameterId),
	}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.DeleteRecurringSubscription(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeleteRecurringSubscription", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}
