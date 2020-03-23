package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	u "github.com/PuerkitoBio/purell"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	paylinksPath               = "/paylinks"
	paylinksIdPath             = "/paylinks/:id"
	paylinksUrlPath            = "/paylinks/:id/url"
	paylinksIdStatSummaryPath  = "/paylinks/:id/dashboard/summary"
	paylinksIdStatCountryPath  = "/paylinks/:id/dashboard/country"
	paylinksIdStatReferrerPath = "/paylinks/:id/dashboard/referrer"
	paylinksIdStatDatePath     = "/paylinks/:id/dashboard/date"
	paylinksIdStatUtmPath      = "/paylinks/:id/dashboard/utm"
	paylinksIdTransactionsPath = "/paylinks/:id/transactions"
)

type PayLinkRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPayLinkRoute(set common.HandlerSet, cfg *common.Config) *PayLinkRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PayLinkRoute"})
	return &PayLinkRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *PayLinkRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(paylinksPath, h.getPaylinksList)
	groups.AuthUser.GET(paylinksIdPath, h.getPaylink)
	groups.AuthUser.GET(paylinksUrlPath, h.getPaylinkUrl)
	groups.AuthUser.DELETE(paylinksIdPath, h.deletePaylink)
	groups.AuthUser.POST(paylinksPath, h.createPaylink)
	groups.AuthUser.PUT(paylinksIdPath, h.updatePaylink)
	groups.AuthUser.GET(paylinksIdStatSummaryPath, h.getPaylinkStatSummary)
	groups.AuthUser.GET(paylinksIdStatCountryPath, h.getPaylinkStatByCountry)
	groups.AuthUser.GET(paylinksIdStatReferrerPath, h.getPaylinkStatByReferrer)
	groups.AuthUser.GET(paylinksIdStatDatePath, h.getPaylinkStatByDate)
	groups.AuthUser.GET(paylinksIdStatUtmPath, h.getPaylinkStatByUtm)
	groups.AuthUser.GET(paylinksIdTransactionsPath, h.getPaylinkTransactions)
}

// @summary Get the list of payment links
// @desc Get the list of payment links for the authorized merchant
// @id paylinksPathGetPaylinksList
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.PaylinksPaginate Returns the list of payment links
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param limit query {integer} false The number of payment links returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/paylinks [get]
func (h *PayLinkRoute) getPaylinksList(ctx echo.Context) error {
	req := &billingpb.GetPaylinksRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.ProjectId = ""

	if req.Limit == 0 {
		req.Limit = int64(h.cfg.LimitDefault)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinks(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinks", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Get the payment link data
// @desc Get the payment link data using the payment link ID
// @id paylinksIdPathGetPaylink
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.Paylink Returns the payment link data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @router /admin/api/v1/paylinks/{id} [get]
func (h *PayLinkRoute) getPaylink(ctx echo.Context) error {
	req := &billingpb.PaylinkRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	err := h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylink(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylink", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link URL
// @desc Get the payment link URL using the payment link ID
// @id paylinksUrlPathGetPaylinkUrl
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {string} Returns the payment link URL with UTM parameters (if any)
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param utm_source query {string} false The UTM-tag of the advertising system, for example: Bing Ads, Google Adwords.
// @param utm_medium query {string} false The UTM-tag of the traffic type, e.g.: cpc, cpm, email newsletter.
// @param utm_campaign query {string} false The UTM-tag of the advertising campaign, for example: Online games, Simulation game.
// @router /admin/api/v1/paylinks/{id}/url [get]
func (h *PayLinkRoute) getPaylinkUrl(ctx echo.Context) error {
	req := &billingpb.GetPaylinkURLRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	req.Id = ctx.Param(common.RequestParameterId)
	req.UrlMask = h.cfg.OrderInlineFormUrlMask + "/?paylink_id=%s"

	res, err := h.dispatch.Services.Billing.GetPaylinkURL(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkURL", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	url, err := u.NormalizeURLString(res.Url, u.FlagsUsuallySafeGreedy|u.FlagRemoveDuplicateSlashes)

	if err != nil {
		h.L().Error("NormalizeURLString failed", logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, url)
}

// @summary Delete the payment link
// @desc Delete the payment link using the payment link ID
// @id paylinksIdPathDeletePaylink
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the payment link was successfully removed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 404 {object} billingpb.ResponseErrorMessage The payment link not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @router /admin/api/v1/paylinks/{id} [delete]
func (h *PayLinkRoute) deletePaylink(ctx echo.Context) error {
	req := &billingpb.PaylinkRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeletePaylink(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeletePaylink", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Create a payment link
// @desc Create a payment link
// @id paylinksPathCreatePaylink
// @tag Payment link
// @accept application/json
// @produce application/json
// @body billingpb.CreatePaylinkRequest
// @success 200 {object} billingpb.Paylink Returns the created payment link data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/paylinks [post]
func (h *PayLinkRoute) createPaylink(ctx echo.Context) error {
	return h.createOrUpdatePaylink(ctx, "")
}

// @summary Update the payment link
// @desc Update the payment link using the payment link ID
// @id paylinksIdPathUpdatePaylink
// @tag Payment link
// @accept application/json
// @produce application/json
// @body billingpb.CreatePaylinkRequest
// @success 200 {object} billingpb.Paylink Returns the created payment link data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @router /admin/api/v1/paylinks/{id} [put]
func (h *PayLinkRoute) updatePaylink(ctx echo.Context) error {
	return h.createOrUpdatePaylink(ctx, ctx.Param(common.RequestParameterId))
}

func (h *PayLinkRoute) createOrUpdatePaylink(ctx echo.Context, paylinkId string) error {
	req := &billingpb.CreatePaylinkRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Id = paylinkId

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdatePaylink(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "CreateOrUpdatePaylink", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link summary for the Dashboard
// @desc Get payment link statistical results for the period using the payment link ID
// @id paylinksIdStatSummaryPathGetPaylinkStatSummary
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.StatCommon Returns the payment link summary
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param period_from query {integer} false The first date of the period for which the statistical results are calculated.
// @param period_to query {integer} false The last date of the period for which the statistical results are calculated.
// @router /admin/api/v1/paylinks/{id}/dashboard/summary [get]
func (h *PayLinkRoute) getPaylinkStatSummary(ctx echo.Context) error {
	req := &billingpb.GetPaylinkStatCommonRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkStatTotal(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkStatTotal", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link summary grouped by the country
// @desc Get payment link statistical results for the period grouped by the country using the payment link ID
// @id paylinksIdStatCountryPathGetPaylinkStatByCountry
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GroupStatCommon Returns the payment link summary
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param period_from query {integer} false The first date of the period for which the statistical results are calculated.
// @param period_to query {integer} false The last date of the period for which the statistical results are calculated.
// @router /admin/api/v1/paylinks/{id}/dashboard/country [get]
func (h *PayLinkRoute) getPaylinkStatByCountry(ctx echo.Context) error {
	req := &billingpb.GetPaylinkStatCommonRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkStatByCountry(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkStatByCountry", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link summary grouped by the referrer
// @desc Get payment link statistical results for the period grouped by the referrer using the payment link ID
// @id paylinksIdStatReferrerPathGetPaylinkStatByReferrer
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GroupStatCommon Returns the payment link summary
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param period_from query {integer} false The first date of the period for which the statistical results are calculated.
// @param period_to query {integer} false The last date of the period for which the statistical results are calculated.
// @router /admin/api/v1/paylinks/{id}/dashboard/referrer [get]
func (h *PayLinkRoute) getPaylinkStatByReferrer(ctx echo.Context) error {
	req := &billingpb.GetPaylinkStatCommonRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkStatByReferrer(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkStatByReferrer", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link summary grouped by the date
// @desc Get payment link statistical results for the period grouped by the date using the payment link ID
// @id paylinksIdStatDatePathGetPaylinkStatByDate
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GroupStatCommon Returns the payment link summary
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param period_from query {integer} false The first date of the period for which the statistical results are calculated.
// @param period_to query {integer} false The last date of the period for which the statistical results are calculated.
// @router /admin/api/v1/paylinks/{id}/dashboard/date [get]
func (h *PayLinkRoute) getPaylinkStatByDate(ctx echo.Context) error {
	req := &billingpb.GetPaylinkStatCommonRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkStatByDate(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkStatByDate", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the payment link summary grouped by the UTM-tag
// @desc Get payment link statistical results for the period grouped by the UTM-tag using the payment link ID
// @id paylinksIdStatUtmPathGetPaylinkStatByUtm
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GroupStatCommon Returns the payment link summary
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param period_from query {integer} false The first date of the period for which the statistical results are calculated.
// @param period_to query {integer} false The last date of the period for which the statistical results are calculated.
// @router /admin/api/v1/paylinks/{id}/dashboard/utm [get]
func (h *PayLinkRoute) getPaylinkStatByUtm(ctx echo.Context) error {
	req := &billingpb.GetPaylinkStatCommonRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkStatByUtm(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkStatByUtm", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the list of payment link's transactions
// @desc Get the list of payment link's transactions using the payment link ID
// @id paylinksIdTransactionsPathGetPaylinkTransactions
// @tag Payment link
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.TransactionsPaginate Returns the list of payment link's transactions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the payment link.
// @param limit query {integer} false The number of transactions returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/paylinks/{id}/transactions [get]
func (h *PayLinkRoute) getPaylinkTransactions(ctx echo.Context) error {
	req := &billingpb.GetPaylinkTransactionsRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	err := h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPaylinkTransactions(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetPaylinkTransactions", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data)
}
