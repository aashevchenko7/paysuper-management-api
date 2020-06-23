package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

type CountryApiV1 struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewCountryApiV1(set common.HandlerSet, cfg *common.Config) *CountryApiV1 {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "CountryApiV1"})
	return &CountryApiV1{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *CountryApiV1) Route(groups *common.Groups) {
	groups.Common.GET("/country", h.get)
	groups.Common.GET("/country/:code", h.getById)
}

// @summary Get the list of currencies
// @desc Get the full list of currencies using the country's name
// @id countryGet
// @tag Country
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.CountriesList Returns the list of currencies
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param name query {string} false The country's name or two-letter country code by ISO 3166-1.
// @router /api/v1/country [get]
func (h *CountryApiV1) get(ctx echo.Context) error {

	res, err := h.dispatch.Services.Billing.GetCountriesList(ctx.Request().Context(), &billingpb.EmptyRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError /*ErrorCountriesListError*/, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the country data
// @desc Get the country data by two-letter country code in ISO 3166-1
// @id countryGetById
// @tag Country
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.Country Returns the country data (taxes, currency, region and others)
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param code path {string} true Two-letter country code in ISO 3166-1
// @router /api/v1/country/{code} [get]
func (h *CountryApiV1) getById(ctx echo.Context) error {
	code := ctx.Param("code")

	if len(code) != 2 {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorIncorrectCountryIdentifier)
	}

	req := &billingpb.GetCountryRequest{
		IsoCode: code,
	}
	err := h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetCountry(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorCountryNotFound)
	}

	return ctx.JSON(http.StatusOK, res)
}
