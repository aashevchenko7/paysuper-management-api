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
	paymentMethodPath           = "/payment_method"
	paymentMethodIdPath         = "/payment_method/:id"
	paymentMethodProductionPath = "/payment_method/:id/production"
	paymentMethodTestPath       = "/payment_method/:id/test"
)

type PaymentMethodApiV1 struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPaymentMethodApiV1(set common.HandlerSet, cfg *common.Config) *PaymentMethodApiV1 {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PaymentMethodApiV1"})
	return &PaymentMethodApiV1{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *PaymentMethodApiV1) Route(groups *common.Groups) {
	groups.SystemUser.POST(paymentMethodPath, h.create)
	groups.SystemUser.PUT(paymentMethodIdPath, h.update)
	groups.SystemUser.POST(paymentMethodProductionPath, h.createProductionSettings)
	groups.SystemUser.PUT(paymentMethodProductionPath, h.updateProductionSettings)
	groups.SystemUser.GET(paymentMethodProductionPath, h.getProductionSettings)
	groups.SystemUser.DELETE(paymentMethodProductionPath, h.deleteProductionSettings)
	groups.SystemUser.POST(paymentMethodTestPath, h.createTestSettings)
	groups.SystemUser.PUT(paymentMethodTestPath, h.updateTestSettings)
	groups.SystemUser.GET(paymentMethodTestPath, h.getTestSettings)
	groups.SystemUser.DELETE(paymentMethodTestPath, h.deleteTestSettings)
}

// @summary Create a payment method of the payment system
// @desc Create a payment method for the payment system
// @id paymentMethodPathCreate
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethod
// @success 200 {object} grpc.ChangePaymentMethodResponse Returns the status of creation
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/payment_method [post]
func (h *PaymentMethodApiV1) create(ctx echo.Context) error {
	return h.createOrUpdatePaymentMethod(ctx)
}

// @summary Update the payment method of the payment system
// @desc Update the payment method for the payment system
// @id paymentMethodIdPathUpdate
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethod
// @success 200 {object} grpc.ChangePaymentMethodResponse Returns the status of update
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @router /system/api/v1/payment_method/{id} [put]
func (h *PaymentMethodApiV1) update(ctx echo.Context) error {
	return h.createOrUpdatePaymentMethod(ctx)
}

func (h *PaymentMethodApiV1) createOrUpdatePaymentMethod(ctx echo.Context) error {
	req := &billingpb.PaymentMethod{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdatePaymentMethod(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the production settings
// @desc Get the production settings
// @id paymentMethodProductionPathGetProductionSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.GetPaymentMethodSettingsResponse Returns the production settings
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @param currency_a3 query {string} false Three-letter currency code by ISO 4217, in uppercase.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @param operating_company_id query {string} false The unique identifier for the operation company.
// @router /system/api/v1/payment_method/{id}/production [get]
func (h *PaymentMethodApiV1) getProductionSettings(ctx echo.Context) error {
	req := &billingpb.GetPaymentMethodSettingsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaymentMethodProductionSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Add production settings of the payment method
// @desc Add production settings of the payment method
// @id paymentMethodProductionPathCreateProductionSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethodParams
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of creation
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @router /system/api/v1/payment_method/{id}/production [post]
func (h *PaymentMethodApiV1) createProductionSettings(ctx echo.Context) error {
	return h.createOrUpdateProductionSettings(ctx)
}

// @summary Update the production settings
// @desc Update the production settings
// @id paymentMethodProductionPathUpdateProductionSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethodParams
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of update
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @router /system/api/v1/payment_method/{id}/production [put]
func (h *PaymentMethodApiV1) updateProductionSettings(ctx echo.Context) error {
	return h.createOrUpdateProductionSettings(ctx)
}

func (h *PaymentMethodApiV1) createOrUpdateProductionSettings(ctx echo.Context) error {

	req := &billingpb.ChangePaymentMethodParamsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdatePaymentMethodProductionSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the production settings of the payment method
// @desc Delete the production settings of the payment method
// @id paymentMethodProductionPathDeleteProductionSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of deletion
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @param currency_a3 query {string} false Three-letter currency code by ISO 4217, in uppercase.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @param operating_company_id query {string} false The unique identifier for the operation company.
// @router /system/api/v1/payment_method/{id}/production [delete]
func (h *PaymentMethodApiV1) deleteProductionSettings(ctx echo.Context) error {
	req := &billingpb.GetPaymentMethodSettingsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeletePaymentMethodProductionSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the testing settings
// @desc Get the testing settings
// @id paymentMethodTestPathGetTestSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.GetPaymentMethodSettingsResponse Returns the testing settings
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @param currency_a3 query {string} false Three-letter currency code by ISO 4217, in uppercase.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @param operating_company_id query {string} false The unique identifier for the operation company.
// @router /system/api/v1/payment_method/{id}/test [get]
func (h *PaymentMethodApiV1) getTestSettings(ctx echo.Context) error {
	req := &billingpb.GetPaymentMethodSettingsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaymentMethodTestSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Add testing settings of the payment method
// @desc Add testing settings of the payment method
// @id paymentMethodTestPathCreateTestSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethodParams
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of creation
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @router /system/api/v1/payment_method/{id}/test [post]
func (h *PaymentMethodApiV1) createTestSettings(ctx echo.Context) error {
	return h.createOrUpdateTestSettings(ctx)
}

// @summary Update the testing settings of the payment method
// @desc Update the testing settings of the payment method
// @id paymentMethodTestPathCreateTestSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethodParams
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of update
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @router /system/api/v1/payment_method/{id}/test [put]
func (h *PaymentMethodApiV1) updateTestSettings(ctx echo.Context) error {
	return h.createOrUpdateTestSettings(ctx)
}

func (h *PaymentMethodApiV1) createOrUpdateTestSettings(ctx echo.Context) error {
	req := &billingpb.ChangePaymentMethodParamsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdatePaymentMethodTestSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the testing settings of the payment method
// @desc Delete the testing settings of the payment method
// @id paymentMethodTestPathCreateTestSettings
// @tag Payment method
// @accept application/json
// @produce application/json
// @body billing.PaymentMethodParams
// @success 200 {object} grpc.ChangePaymentMethodParamsResponse Returns the status of deletion
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment method.
// @param currency_a3 query {string} false Three-letter currency code by ISO 4217, in uppercase.
// @param mcc_code query {string} true The Merchant Category Code (MCC) is a four-digit number listed in ISO 18245.
// @param operating_company_id query {string} false The unique identifier for the operation company.
// @router /system/api/v1/payment_method/{id}/test [delete]
func (h *PaymentMethodApiV1) deleteTestSettings(ctx echo.Context) error {
	req := &billingpb.GetPaymentMethodSettingsRequest{
		PaymentMethodId: ctx.Param("id"),
	}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeletePaymentMethodTestSettings(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}
