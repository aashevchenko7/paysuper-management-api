package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billing "github.com/paysuper/paysuper-proto/go/billingpb"
	grpc "github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	paymentMinLimitSystemPath = "/payment_min_limit_system"
)

type PaymentMinLimitSystemRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPaymentMinLimitSystemRoute(set common.HandlerSet, cfg *common.Config) *PaymentMinLimitSystemRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PaymentMinLimitSystemRoute"})
	return &PaymentMinLimitSystemRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *PaymentMinLimitSystemRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(paymentMinLimitSystemPath, h.getPaymentMinLimitSystemList)
	groups.AuthUser.POST(paymentMinLimitSystemPath, h.setPaymentMinLimitSystem)
}

// @summary Get the list of the payment system limits
// @desc Get the list of the payment system limits
// @id paymentMinLimitSystemPathGetPaymentMinLimitSystemList
// @tag Limits
// @accept application/json
// @produce application/json
// @success 200 {object} []billing.OperatingCompany Returns the operating company's payment system limits
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/payment_min_limit_system [get]
func (h *PaymentMinLimitSystemRoute) getPaymentMinLimitSystemList(ctx echo.Context) error {
	req := &billingpb.EmptyRequest{}

	res, err := h.dispatch.Services.Billing.GetOperatingCompaniesList(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetOperatingCompaniesList", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Items)
}

// @summary Set the payment system limits
// @desc Set the payment system limits
// @id paymentMinLimitSystemPathSetPaymentMinLimitSystem
// @tag Limits
// @accept application/json
// @produce application/json
// @body billing.PaymentMinLimitSystem
// @success 200 {string} Returns an empty response body if the system limits were successfully set
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/payment_min_limit_system [post]
func (h *PaymentMinLimitSystemRoute) setPaymentMinLimitSystem(ctx echo.Context) error {
	req := &billingpb.PaymentMinLimitSystem{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetPaymentMinLimitSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "AddPaymentMinLimitSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}
