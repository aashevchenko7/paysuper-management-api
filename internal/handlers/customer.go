package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/helpers"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	"net/http"
)

const (
	customerListing  = "/customers"
	customerDetailed = "/customers/:id"
	customerSubscriptions = "/customers/:id/subscriptions"
	customerSubscription = "/customers/:id/subscription/:subscription_id"
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
	groups.AuthUser.GET(customerListing, h.getCustomers)
	groups.AuthUser.GET(customerDetailed, h.getCustomerDetails)
	groups.SystemUser.GET(customerListing, h.getCustomers)
	groups.SystemUser.GET(customerDetailed, h.getCustomerDetails)

	groups.AuthUser.GET(customerSubscriptions, h.getCustomerSubscriptions)
	groups.SystemUser.GET(customerSubscriptions, h.getCustomerSubscriptions)
	groups.Common.GET(customerSubscriptions, h.getCustomerSubscriptions)

	groups.AuthUser.GET(customerSubscription, h.getCustomerSubscription)
	groups.SystemUser.GET(customerSubscription, h.getCustomerSubscription)
	groups.Common.GET(customerSubscription, h.getCustomerSubscription)
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
	req := &recurringpb.FindSubscriptionsRequest{}
	err := ctx.Bind(req)

	customerId := ctx.Param(common.RequestParameterId)
	req.CustomerUuid = customerId
	user := common.ExtractUserContext(ctx)

	if user == nil {
		cookies := helpers.GetRequestCookie(ctx, common.CustomerTokenCookiesName)
		rsp, err := h.dispatch.Services.Billing.DeserializeCookie(ctx.Request().Context(), &billingpb.DeserializeCookieRequest{Cookie: cookies})
		if err != nil {
			common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeserializeCookie", req)
			return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
		}

		if rsp.Status != billingpb.ResponseStatusOk {
			return echo.NewHTTPError(int(rsp.Status), rsp.Message)
		}

		if rsp.Item.CustomerId != customerId {
			return echo.NewHTTPError(http.StatusForbidden, common.ErrorUnknown)
		}
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Repository.FindSubscriptions(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, "recurringpb", "FindSubscriptions", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res.List)
}

func (h *CustomerRoute) getCustomerSubscription(ctx echo.Context) error {
	req := &recurringpb.GetSubscriptionRequest{}

	req.Id = ctx.Param("subscription_id")
	err := h.dispatch.Validate.Struct(req)
	customerId := ctx.Param(common.RequestParameterId)

	user := common.ExtractUserContext(ctx)
	if user == nil {
		cookies := helpers.GetRequestCookie(ctx, common.CustomerTokenCookiesName)
		rsp, err := h.dispatch.Services.Billing.DeserializeCookie(ctx.Request().Context(), &billingpb.DeserializeCookieRequest{Cookie: cookies})
		if err != nil {
			common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeserializeCookie", req)
			return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
		}

		if rsp.Status != billingpb.ResponseStatusOk {
			return echo.NewHTTPError(int(rsp.Status), rsp.Message)
		}

		if rsp.Item.CustomerId != customerId {
			return echo.NewHTTPError(http.StatusForbidden, common.ErrorUnknown)
		}
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Repository.GetSubscription(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, "recurringpb", "GetSubscription", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	if res.Subscription.CustomerUuid != customerId {
		return echo.NewHTTPError(http.StatusForbidden, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res.Subscription)
}