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

type PaymentCostRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPaymentCostRoute(set common.HandlerSet, cfg *common.Config) *PaymentCostRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PaymentCostRoute"})
	return &PaymentCostRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

const (
	paymentCostsChannelSystemPath        = "/payment_costs/channel/system"
	paymentCostsChannelSystemAllPath     = "/payment_costs/channel/system/all"
	paymentCostsChannelMerchantPath      = "/payment_costs/channel/merchant/:merchant_id"
	paymentCostsChannelMerchantAllPath   = "/payment_costs/channel/merchant/:merchant_id/all"
	paymentCostsChannelSystemIdPath      = "/payment_costs/channel/system/:id"
	paymentCostsChannelMerchantIdsPath   = "/payment_costs/channel/merchant/:merchant_id/:rate_id"
	paymentCostsMoneyBackAllPath         = "/payment_costs/money_back/system/all"
	paymentCostsMoneyBackMerchantPath    = "/payment_costs/money_back/merchant/:merchant_id"
	paymentCostsMoneyBackMerchantAllPath = "/payment_costs/money_back/merchant/:merchant_id/all"
	paymentCostsMoneyBackSystemPath      = "/payment_costs/money_back/system"
	paymentCostsMoneyBackSystemIdPath    = "/payment_costs/money_back/system/:id"
	paymentCostsMoneyBackMerchantIdsPath = "/payment_costs/money_back/merchant/:merchant_id/:rate_id"
)

func (h *PaymentCostRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(paymentCostsChannelSystemAllPath, h.getAllPaymentChannelCostSystem)
	groups.SystemUser.GET(paymentCostsChannelMerchantAllPath, h.getAllPaymentChannelCostMerchant) //надо править
	groups.SystemUser.GET(paymentCostsMoneyBackAllPath, h.getAllMoneyBackCostSystem)
	groups.SystemUser.GET(paymentCostsMoneyBackMerchantAllPath, h.getAllMoneyBackCostMerchant) //надо править

	groups.SystemUser.GET(paymentCostsChannelSystemPath, h.getPaymentChannelCostSystem)
	groups.SystemUser.GET(paymentCostsChannelMerchantPath, h.getPaymentChannelCostMerchant)
	groups.SystemUser.GET(paymentCostsMoneyBackSystemPath, h.getMoneyBackCostSystem)
	groups.SystemUser.GET(paymentCostsMoneyBackMerchantPath, h.getMoneyBackCostMerchant)

	groups.SystemUser.DELETE(paymentCostsChannelSystemIdPath, h.deletePaymentChannelCostSystem)
	groups.SystemUser.DELETE(paymentCostsChannelMerchantPath, h.deletePaymentChannelCostMerchant)
	groups.SystemUser.DELETE(paymentCostsMoneyBackSystemIdPath, h.deleteMoneyBackCostSystem)
	groups.SystemUser.DELETE(paymentCostsMoneyBackMerchantPath, h.deleteMoneyBackCostMerchant)

	groups.SystemUser.POST(paymentCostsChannelSystemPath, h.setPaymentChannelCostSystem)
	groups.SystemUser.POST(paymentCostsChannelMerchantPath, h.setPaymentChannelCostMerchant)
	groups.SystemUser.POST(paymentCostsMoneyBackSystemPath, h.setMoneyBackCostSystem)
	groups.SystemUser.POST(paymentCostsMoneyBackMerchantPath, h.setMoneyBackCostMerchant)

	groups.SystemUser.PUT(paymentCostsChannelSystemIdPath, h.setPaymentChannelCostSystem)
	groups.SystemUser.PUT(paymentCostsChannelMerchantIdsPath, h.setPaymentChannelCostMerchant)
	groups.SystemUser.PUT(paymentCostsMoneyBackSystemIdPath, h.setMoneyBackCostSystem)
	groups.SystemUser.PUT(paymentCostsMoneyBackMerchantIdsPath, h.setMoneyBackCostMerchant)
}

// @summary Get system costs for payment operations
// @desc Get system costs for payment operations
// @id paymentCostsChannelSystemPathGetPaymentChannelCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.PaymentChannelCostSystemList Returns system costs for payment operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param name query {string} true The payment method's name.
// @param region query {string} true The region name. Available values: CIS, Russia, West Asia, EU, North America, Central America, South America, United Kingdom, Worldwide, South Pacific.
// @param country query {string} false The country code.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @router /system/api/v1/payment_costs/channel/system [get]
func (h *PaymentCostRoute) getPaymentChannelCostSystem(ctx echo.Context) error {
	req := &billingpb.PaymentChannelCostSystemRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaymentChannelCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaymentChannelCostSystem", req)

		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get system costs for payment operations
// @desc Get system costs for payment operations
// @id paymentCostsChannelMerchantPathGetPaymentChannelCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.PaymentChannelCostMerchant Returns system costs for payment operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param payout_currency query {string} true The payout currency.
// @param amount query {integer} true The payout amount.
// @param name query {string} true The payment method's name.
// @param region query {string} true The region name. Available values: CIS, Russia, West Asia, EU, North America, Central America, South America, United Kingdom, Worldwide, South Pacific.
// @param country query {string} false The country code.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @router /system/api/v1/payment_costs/channel/merchant/{merchant_id} [get]
func (h *PaymentCostRoute) getPaymentChannelCostMerchant(ctx echo.Context) error {
	req := &billingpb.PaymentChannelCostMerchantRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.MerchantId = ctx.Param(common.RequestParameterMerchantId)
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaymentChannelCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaymentChannelCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get system costs for money back operations
// @desc Get system costs for money back operations
// @id paymentCostsMoneyBackSystemPathGetMoneyBackCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.MoneyBackCostSystem Returns system costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param name query {string} true The payment method's name.
// @param undo_reason query {string} true The return reason. Available values: refund, reversal, chargeback.
// @param payout_currency query {string} true The payout currency.
// @param payment_stage query {integer} true The payout stage.
// @param days query {integer} true The number of days after the payment operation.
// @param region query {string} true The region name. Available values: CIS, Russia, West Asia, EU, North America, Central America, South America, United Kingdom, Worldwide, South Pacific.
// @param country query {string} false The country code.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @router /system/api/v1/payment_costs/money_back/system [get]
func (h *PaymentCostRoute) getMoneyBackCostSystem(ctx echo.Context) error {
	req := &billingpb.MoneyBackCostSystemRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetMoneyBackCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMoneyBackCostSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get merchant costs for money back operations
// @desc Get merchant costs for money back operations
// @id paymentCostsMoneyBackMerchantPathGetMoneyBackCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.MoneyBackCostMerchant Returns merchant costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param name query {string} true The payment method's name.
// @param undo_reason query {string} true The return reason. Available values: refund, reversal, chargeback.
// @param payout_currency query {string} true The payout currency.
// @param payment_stage query {integer} true The payout stage.
// @param days query {integer} true The number of days after the payment operation.
// @param region query {string} true The region name. Available values: CIS, Russia, West Asia, EU, North America, Central America, South America, United Kingdom, Worldwide, South Pacific.
// @param country query {string} false The country code.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @router /system/api/v1/payment_costs/money_back/merchant/{merchant_id} [get]
func (h *PaymentCostRoute) getMoneyBackCostMerchant(ctx echo.Context) error {
	req := &billingpb.MoneyBackCostMerchantRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.MerchantId = ctx.Param(common.RequestParameterMerchantId)
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetMoneyBackCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMoneyBackCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Delete the system cost for payment operations
// @desc Delete the system cost for payment operations
// @id paymentCostsChannelSystemIdPathDeletePaymentChannelCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the cost was successfully removed
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/channel/system/{id} [delete]
func (h *PaymentCostRoute) deletePaymentChannelCostSystem(ctx echo.Context) error {
	req := &billingpb.PaymentCostDeleteRequest{Id: ctx.Param(common.RequestParameterId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeletePaymentChannelCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeletePaymentChannelCostSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Delete merchant costs for payment operations
// @desc Delete merchant costs for payment operations
// @id paymentCostsChannelMerchantPathDeletePaymentChannelCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the cost was successfully removed
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/channel/merchant/{merchant_id} [delete]
func (h *PaymentCostRoute) deletePaymentChannelCostMerchant(ctx echo.Context) error {
	req := &billingpb.PaymentCostDeleteRequest{Id: ctx.Param(common.RequestParameterMerchantId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeletePaymentChannelCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeletePaymentChannelCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Delete the system cost for money back operations
// @desc Delete the system cost for money back operations
// @id paymentCostsMoneyBackSystemIdPathDeleteMoneyBackCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the cost was successfully removed
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/money_back/system/{id} [delete]
func (h *PaymentCostRoute) deleteMoneyBackCostSystem(ctx echo.Context) error {
	req := &billingpb.PaymentCostDeleteRequest{Id: ctx.Param(common.RequestParameterId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeleteMoneyBackCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeleteMoneyBackCostSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Delete merchant costs for money back operations
// @desc Delete merchant costs for money back operations
// @id paymentCostsMoneyBackMerchantPathDeleteMoneyBackCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the cost was successfully removed
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/money_back/merchant/{merchant_id} [delete]
func (h *PaymentCostRoute) deleteMoneyBackCostMerchant(ctx echo.Context) error {
	req := &billingpb.PaymentCostDeleteRequest{Id: ctx.Param(common.RequestParameterMerchantId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeleteMoneyBackCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeleteMoneyBackCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Create system costs for payments operations
// @desc Create system costs for payments operations
// @id paymentCostsChannelSystemPathSetPaymentChannelCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.PaymentChannelCostSystem
// @success 200 {object} billing.PaymentChannelCostSystem Returns system costs for payments operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/payment_costs/channel/system [post]

// @summary Update system costs for payments operations
// @desc Update system costs for payments operations
// @id paymentCostsChannelSystemIdPathSetPaymentChannelCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.PaymentChannelCostSystem
// @success 200 {object} billing.PaymentChannelCostSystem Returns system costs for payments operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/channel/system/{id} [put]
func (h *PaymentCostRoute) setPaymentChannelCostSystem(ctx echo.Context) error {
	req := &billingpb.PaymentChannelCostSystem{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	if pcId := ctx.Param(common.RequestParameterId); pcId != "" {
		req.Id = pcId
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetPaymentChannelCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "SetPaymentChannelCostSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Create merchant costs for payments operations
// @desc Create merchant costs for payments operations
// @id paymentCostsChannelMerchantPathSetPaymentChannelCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.PaymentChannelCostMerchant
// @success 200 {object} billing.PaymentChannelCostMerchant Returns merchant costs for payments operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/channel/merchant/{merchant_id} [post]

// @summary Update merchant costs for payments operations
// @desc Update merchant costs for payments operations
// @id paymentCostsChannelMerchantIdsPathSetPaymentChannelCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.PaymentChannelCostMerchant
// @success 200 {object} billing.PaymentChannelCostMerchant Returns merchant costs for payments operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param rate_id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/channel/merchant/{merchant_id}/{rate_id} [put]
func (h *PaymentCostRoute) setPaymentChannelCostMerchant(ctx echo.Context) error {
	req := &billingpb.PaymentChannelCostMerchant{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.MerchantId = ctx.Param(common.RequestParameterMerchantId)

	if ctx.Request().Method == http.MethodPut {
		req.Id = ctx.Param(common.RequestParameterRateId)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetPaymentChannelCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "SetPaymentChannelCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Create system costs for money back operations
// @desc Create system costs for money back operations
// @id paymentCostsMoneyBackSystemPathSetMoneyBackCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MoneyBackCostSystem
// @success 200 {object} billing.MoneyBackCostSystem Returns system costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/payment_costs/money_back/system [post]

// @summary Create system costs for money back operations
// @desc Create system costs for money back operations
// @id paymentCostsMoneyBackSystemIdPathSetMoneyBackCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MoneyBackCostSystem
// @success 200 {object} billing.MoneyBackCostSystem Returns system costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/money_back/system/{id} [post]
func (h *PaymentCostRoute) setMoneyBackCostSystem(ctx echo.Context) error {
	req := &billingpb.MoneyBackCostSystem{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	if pcId := ctx.Param(common.RequestParameterId); pcId != "" {
		req.Id = pcId
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetMoneyBackCostSystem(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "SetMoneyBackCostSystem", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Create merchant costs for money back operations
// @desc Create merchant costs for money back operations
// @id paymentCostsMoneyBackMerchantPathSetMoneyBackCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MoneyBackCostMerchant
// @success 200 {object} billing.MoneyBackCostMerchant Returns merchant costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/money_back/merchant/{merchant_id} [post]

// @summary Update merchant costs for money back operations
// @desc Update merchant costs for money back operations
// @id paymentCostsMoneyBackMerchantIdsPathSetMoneyBackCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MoneyBackCostMerchant
// @success 200 {object} billing.MoneyBackCostMerchant Returns merchant costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param rate_id path {string} true The unique identifier for the cost.
// @router /system/api/v1/payment_costs/money_back/merchant/{merchant_id}/{rate_id} [put]
func (h *PaymentCostRoute) setMoneyBackCostMerchant(ctx echo.Context) error {
	req := &billingpb.MoneyBackCostMerchant{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.MerchantId = ctx.Param(common.RequestParameterMerchantId)

	if ctx.Request().Method == http.MethodPut {
		req.Id = ctx.Param(common.RequestParameterRateId)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetMoneyBackCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "SetMoneyBackCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get all system costs for payments
// @desc Get all system costs for payments
// @id paymentCostsChannelSystemAllPathGetAllPaymentChannelCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.PaymentChannelCostSystemList Returns the all system costs for payments
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/payment_costs/channel/system/all [get]
func (h *PaymentCostRoute) getAllPaymentChannelCostSystem(ctx echo.Context) error {
	res, err := h.dispatch.Services.Billing.GetAllPaymentChannelCostSystem(ctx.Request().Context(), &billingpb.EmptyRequest{})

	if err != nil {
		h.L().Error(billingpb.ErrorGrpcServiceCallFailed, logger.PairArgs("err", err.Error(), common.ErrorFieldService, billingpb.ServiceName, common.ErrorFieldMethod, "GetAllPaymentChannelCostSystem"))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get all merchant costs for payments
// @desc Get all merchant costs for payments
// @id paymentCostsChannelMerchantAllPathGetAllPaymentChannelCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.PaymentChannelCostMerchantList Returns all merchant costs for payments
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/channel/merchant/{merchant_id}/all [get]
func (h *PaymentCostRoute) getAllPaymentChannelCostMerchant(ctx echo.Context) error {
	req := &billingpb.PaymentChannelCostMerchantListRequest{MerchantId: ctx.Param(common.RequestParameterMerchantId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetAllPaymentChannelCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetAllPaymentChannelCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get all system costs for the money back operations
// @desc Get all system costs for money back operations
// @id paymentCostsMoneyBackAllPathGetAllMoneyBackCostSystem
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.MoneyBackCostSystemList Returns all system costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/payment_costs/money_back/system/all [get]
func (h *PaymentCostRoute) getAllMoneyBackCostSystem(ctx echo.Context) error {
	res, err := h.dispatch.Services.Billing.GetAllMoneyBackCostSystem(ctx.Request().Context(), &billingpb.EmptyRequest{})

	if err != nil {
		h.L().Error(billingpb.ErrorGrpcServiceCallFailed, logger.PairArgs("err", err.Error(), common.ErrorFieldService, billingpb.ServiceName, common.ErrorFieldMethod, "GetAllMoneyBackCostSystem"))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get all merchant costs for money back operations
// @desc Get all merchant costs for money back operations
// @id paymentCostsMoneyBackMerchantAllPathGetAllMoneyBackCostMerchant
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.MoneyBackCostMerchantList Returns all merchant costs for money back operations
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/payment_costs/money_back/merchant/{merchant_id}/all [get]
func (h *PaymentCostRoute) getAllMoneyBackCostMerchant(ctx echo.Context) error {
	req := &billingpb.MoneyBackCostMerchantListRequest{MerchantId: ctx.Param(common.RequestParameterMerchantId)}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetAllMoneyBackCostMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetAllMoneyBackCostMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}
