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
	zipCodePath = "/zip"
)

type ZipCodeRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewZipCodeRoute(set common.HandlerSet, cfg *common.Config) *ZipCodeRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "ZipCodeRoute"})
	return &ZipCodeRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *ZipCodeRoute) Route(groups *common.Groups) {
	groups.Common.GET(zipCodePath, h.checkZip)
}

// @summary Search the city
// @desc Search the city using the country (and the ZIP code for US)
// @id zipCodePathCheckZip
// @tag Country
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.FindByZipCodeResponse Returns the country data (region, city, and others)
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage The country not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param country query {string} true The country code.
// @param zip query {string} false The postal code. It's required for US.
// @param limit query {integer} false The number of objects returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /api/v1/zip [get]
func (h *ZipCodeRoute) checkZip(ctx echo.Context) error {
	req := &billingpb.FindByZipCodeRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit <= 0 {
		req.Limit = int64(h.cfg.LimitDefault)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.FindByZipCode(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}
