package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"go.uber.org/zap"
	"net/http"
)

const (
	customerListing       = "/customers"
	customerDetailed      = "/customers/:id"
	customerSubscriptions = "/customers/:id/subscriptions"
	customerSubscription  = "/customers/:id/subscription/:subscription_id"
	customerCard          = "/customers/:id/card/:card_id"
)

type CustomerRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewCustomerRoute(set common.HandlerSet, cfg *common.Config) *CustomerRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "CustomerRoute"})
	return &CustomerRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *CustomerRoute) Route(groups *common.Groups) {
	groups.AuthUser.POST(customerListing, h.getCustomers)
	groups.AuthUser.GET(customerDetailed, h.getCustomerDetails)
	groups.SystemUser.POST(customerListing, h.getCustomers)
	groups.SystemUser.GET(customerDetailed, h.getCustomerDetails)

	groups.AuthUser.GET(customerSubscriptions, h.getCustomerSubscriptions)
	groups.SystemUser.GET(customerSubscriptions, h.getCustomerSubscriptions)

	groups.AuthUser.GET(customerSubscription, h.getCustomerSubscription)
	groups.SystemUser.GET(customerSubscription, h.getCustomerSubscription)

	groups.AuthUser.DELETE(customerSubscription, h.deleteCustomerSubscription)
	groups.SystemUser.DELETE(customerSubscription, h.deleteCustomerSubscription)

	groups.SystemUser.DELETE(customerCard, h.deleteCustomerCardAdmin)
}

func (h *CustomerRoute) getCustomerDetails(ctx echo.Context) error {
	req := &billingpb.GetCustomerInfoRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = ctx.Param(common.RequestParameterId)
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetCustomerInfo(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetCustomerInfo", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

func (h *CustomerRoute) getCustomers(ctx echo.Context) error {
	req := &billingpb.ListCustomersRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetCustomerList(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetCustomerList", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Items)
}

func (h *CustomerRoute) getCustomerSubscriptions(ctx echo.Context) error {
	req := &billingpb.FindSubscriptionsRequest{}
	err := ctx.Bind(req)

	customerId := ctx.Param(common.RequestParameterId)
	req.CustomerId = customerId

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.FindSubscriptions(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, "recurringpb", "FindSubscriptions", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res.List)
}

func (h *CustomerRoute) getCustomerSubscription(ctx echo.Context) error {
	req := &billingpb.GetSubscriptionRequest{}

	req.Id = ctx.Param("subscription_id")
	req.CustomerId = ctx.Param(common.RequestParameterId)
	err := h.dispatch.Validate.Struct(req)
	customerId := ctx.Param(common.RequestParameterId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetCustomerSubscription(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, "recurringpb", "GetSubscription", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	if res.Subscription.CustomerUuid != customerId {
		zap.L().Error("trying to get wrong subscription", zap.String("request_customer_id", customerId), zap.String("subscription_customer_id", res.Subscription.CustomerId))
		return echo.NewHTTPError(http.StatusForbidden, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res.Subscription)
}

func (h *CustomerRoute) deleteCustomerCardAdmin(ctx echo.Context) error {
	req := &billingpb.DeleteCustomerCardRequest{}

	req.CustomerId = ctx.Param(common.RequestParameterId)
	req.Id = ctx.Param("card_id")

	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeleteCustomerCard(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeleteCustomerCard", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *CustomerRoute) deleteCustomerSubscription(ctx echo.Context) error {
	req := &billingpb.DeleteRecurringSubscriptionRequest{}

	req.CustomerId = ctx.Param(common.RequestParameterId)
	req.SubscriptionId = ctx.Param("subscription_id")

	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
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
