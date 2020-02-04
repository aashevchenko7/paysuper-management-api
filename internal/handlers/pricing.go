package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"net/http"
)

const (
	pricingRecommendedConversionPath = "/pricing/recommended/conversion"
	pricingRecommendedSteamPath      = "/pricing/recommended/steam"
	pricingRecommendedTablePath      = "/pricing/recommended/table"
)

type Pricing struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPricingRoute(set common.HandlerSet, cfg *common.Config) *Pricing {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PriceGroup"})
	return &Pricing{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *Pricing) Route(groups *common.Groups) {
	groups.Common.GET(pricingRecommendedConversionPath, h.getRecommendedByConversion)
	groups.Common.GET(pricingRecommendedSteamPath, h.getRecommendedBySteam)
	groups.Common.GET(pricingRecommendedTablePath, h.getRecommendedTable)
}

// @summary Get recommended currency conversion prices based on exchange rates
// @desc Calculation of recommended currency conversion prices for regions based on the exchange rates
// @id pricingRecommendedConversionPathGetRecommendedByConversion
// @tag Pricing
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.RecommendedPriceResponse Returns the list of recommended currency conversion prices
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param amount query {string} true The amount of price.
// @param currency query {string} true Three-letter currency code by ISO 4217, in uppercase.
// @router /api/v1/pricing/recommended/conversion [get]
func (h *Pricing) getRecommendedByConversion(ctx echo.Context) error {
	req := &billingpb.RecommendedPriceRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetRecommendedPriceByConversion(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupRecommendedList)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get recommended currency conversion prices based on based on the Steam price ranges
// @desc Calculation of recommended currency conversion prices based on the Steam price ranges taking the regional factors into account
// @id pricingRecommendedSteamPathGetRecommendedBySteam
// @tag Pricing
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.RecommendedPriceResponse Returns the list of recommended currency conversion prices
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param amount query {string} true The amount of price.
// @param currency query {string} true Three-letter currency code by ISO 4217, in uppercase.
// @router /api/v1/pricing/recommended/steam [get]
func (h *Pricing) getRecommendedBySteam(ctx echo.Context) error {
	req := &billingpb.RecommendedPriceRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetRecommendedPriceByPriceGroup(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupRecommendedList)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get ranges of recommended currency conversion prices
// @desc Get the table of recommended currency conversion prices ranges
// @id pricingRecommendedTablePathGetRecommendedTable
// @tag Pricing
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.RecommendedPriceTableResponse Returns the table of recommended currency conversion prices ranges
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param currency query {string} true Three-letter currency code by ISO 4217, in uppercase.
// @router /api/v1/pricing/recommended/table [get]
func (h *Pricing) getRecommendedTable(ctx echo.Context) error {
	req := &billingpb.RecommendedPriceTableRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetRecommendedPriceTable(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessagePriceGroupRecommendedList)
	}

	return ctx.JSON(http.StatusOK, res)
}
