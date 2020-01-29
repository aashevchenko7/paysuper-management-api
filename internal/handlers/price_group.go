package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"net/http"
)

const (
	priceGroupCountryPath    = "/price_group/country"
	priceGroupCurrenciesPath = "/price_group/currencies"
	priceGroupRegionPath     = "/price_group/region"
)

type PriceGroup struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPriceGroupRoute(set common.HandlerSet, cfg *common.Config) *PriceGroup {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PriceGroup"})
	return &PriceGroup{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *PriceGroup) Route(groups *common.Groups) {
	groups.Common.GET(priceGroupCountryPath, h.getPriceGroupByCountry)
	groups.Common.GET(priceGroupCurrenciesPath, h.getCurrencyList)
	groups.Common.GET(priceGroupRegionPath, h.getCurrencyByRegion)
}

// @summary Get the currency and region
// @desc Get the currency and region using the country's name
// @id priceGroupCountryPathGetPriceGroupByCountry
// @tag Price group
// @accept application/json
// @produce application/json
// @success 200 {object} billing.PriceGroup Returns the country's region and currency
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param country query {string} true The country's name.
// @router /api/v1/price_group/country [get]
func (h *PriceGroup) getPriceGroupByCountry(ctx echo.Context) error {
	req := &grpc.PriceGroupByCountryRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPriceGroupByCountry(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupByCountry)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the list of currencies
// @desc Get the full list of currencies with information about regions and countries using the country's name
// @id priceGroupCurrenciesPathGetCurrencyList
// @tag Price group
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.PriceGroupCurrenciesResponse Returns a full list of currencies with information about regions and countries
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param country query {string} true The country's name.
// @param zip query {string} true The postal code. Required for US.
// @param limit query {integer} false The number of objects returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /api/v1/price_group/currencies [get]
func (h *PriceGroup) getCurrencyList(ctx echo.Context) error {
	req := &grpc.EmptyRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPriceGroupCurrencies(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupCurrencyList)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the currency and the list of countries
// @desc Get the currency and the list of countries using the region
// @id priceGroupRegionPathGetCurrencyByRegion
// @tag Price group
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.PriceGroupCurrenciesResponse Returns the currency and the list of countries
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param region query {string} true The country region's name.
// @router /api/v1/price_group/region [get]
func (h *PriceGroup) getCurrencyByRegion(ctx echo.Context) error {
	req := &grpc.PriceGroupByRegionRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPriceGroupCurrencyByRegion(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupCurrencyByRegion)
	}

	return ctx.JSON(http.StatusOK, res)
}
