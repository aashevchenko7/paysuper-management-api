package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/taxpb"
	"net/http"
	"strconv"
)

const (
	taxesPath   = "/taxes"
	taxesIDPath = "/taxes/:id"
)

type TaxesRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewTaxesRoute(set common.HandlerSet, cfg *common.Config) *TaxesRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "TaxesRoute"})
	return &TaxesRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *TaxesRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(taxesPath, h.getTaxes)
	groups.SystemUser.POST(taxesPath, h.setTax)
	groups.SystemUser.DELETE(taxesIDPath, h.deleteTax)
}

// @summary Get the system tax rates list
// @desc Get the system tax rates list
// @id taxesPathGetTaxes
// @tag Tax
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetPaymentMethodSettingsResponse Returns the production settings
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param zip query {string} false The postal code. Required for US.
// @param country query {string} false The country code.
// @param city query {string} false The city's name.
// @param state query {string} false The state's name.
// @param limit query {integer} true The number of taxes returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /system/api/v1/taxes [get]
func (h *TaxesRoute) getTaxes(ctx echo.Context) error {
	req := h.bindGetTaxes(ctx)
	res, err := h.dispatch.Services.Tax.GetRates(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	return ctx.JSON(http.StatusOK, res.Rates)
}

func (h *TaxesRoute) bindGetTaxes(ctx echo.Context) *taxpb.GetRatesRequest {
	structure := &taxpb.GetRatesRequest{}

	params := ctx.QueryParams()

	if v, ok := params["country"]; ok {
		structure.Country = string(v[0])
	}

	if v, ok := params["city"]; ok {
		structure.City = string(v[0])
	}

	if v, ok := params["state"]; ok {
		structure.State = string(v[0])
	}

	if v, ok := params["zip"]; ok {
		structure.Zip = string(v[0])
	}

	if v, ok := params[common.RequestParameterLimit]; ok {
		if i, err := strconv.ParseInt(v[0], 10, 32); err == nil {
			structure.Limit = int32(i)
		}
	} else {
		structure.Limit = int32(h.cfg.LimitDefault)
	}

	if v, ok := params[common.RequestParameterOffset]; ok {
		if i, err := strconv.ParseInt(v[0], 10, 32); err == nil {
			structure.Offset = int32(i)
		}
	} else {
		structure.Offset = int32(h.cfg.OffsetDefault)
	}

	return structure
}

// @summary Edit the system tax rate
// @desc Create or update the system tax rate
// @id taxesPathSetTax
// @tag Tax
// @accept application/json
// @produce application/json
// @body taxpb.TaxRate
// @success 200 {object} taxpb.TaxRate Returns the tax rate
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/taxes [post]
func (h *TaxesRoute) setTax(ctx echo.Context) error {
	if ctx.Request().ContentLength == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req := &taxpb.TaxRate{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	res, err := h.dispatch.Services.Tax.CreateOrUpdate(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the system tax rate
// @desc Mark the system tax rate as removed
// @id taxesIDPathDeleteTax
// @tag Tax
// @accept application/json
// @produce application/json
// @success 200 {string} Returns an empty response body if the tax rate has been successfully removed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the tax rate.
// @router /system/api/v1/taxes/{id} [delete]
func (h *TaxesRoute) deleteTax(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	value, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	res, err := h.dispatch.Services.Tax.DeleteRateById(ctx.Request().Context(), &taxpb.DeleteRateRequest{Id: uint32(value)})
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	return ctx.JSON(http.StatusOK, res)
}
