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
	balancePath         = "/balance"
	balanceMerchantPath = "/balance/:merchant_id"
)

type BalanceRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewBalanceRoute(set common.HandlerSet, cfg *common.Config) *BalanceRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "BalanceRoute"})
	return &BalanceRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *BalanceRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(balancePath, h.getBalance)
	groups.SystemUser.GET(balanceMerchantPath, h.getBalance)
}

// @summary Get the merchant's balance
// @desc Get the merchant's balance
// @id balancePathGetBalance
// @tag Balance
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.MerchantBalance Returns the merchant's balance data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/balance [get]

// @summary Get the merchant's balance using the merchant ID
// @desc Get the merchant's balance using the merchant ID
// @id balanceMerchantPathGetBalance
// @tag Balance
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.MerchantBalance Returns the merchant's balance data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/balance/{merchant_id} [get]
func (h *BalanceRoute) getBalance(ctx echo.Context) error {
	req := &billingpb.GetMerchantBalanceRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	res, err := h.dispatch.Services.Billing.GetMerchantBalance(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMerchantBalance", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}
