package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"
)

const (
	royaltyReportsPath                     = "/royalty_reports"
	royaltyReportsIdPath                   = "/royalty_reports/:report_id"
	royaltyReportsIdDownloadPath           = "/royalty_reports/:report_id/download"
	royaltyReportsTransactionsPath         = "/royalty_reports/:report_id/transactions"
	royaltyReportsTransactionsDownloadPath = "/royalty_reports/:report_id/transactions/download"
	royaltyReportsAcceptPath               = "/royalty_reports/:report_id/accept"
	royaltyReportsDeclinePath              = "/royalty_reports/:report_id/decline"
	royaltyReportsChangePath               = "/royalty_reports/:report_id/change"
)

type RoyaltyReportRequestFile struct {
	Id         string `json:"id" validate:"required,hexadecimal,len=24"`
	MerchantId string `json:"merchant_id" validate:"required,hexadecimal,len=24"`
}

type RoyaltyReportsRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewRoyaltyReportsRoute(set common.HandlerSet, cfg *common.Config) *RoyaltyReportsRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "RoyaltyReportsRoute"})
	return &RoyaltyReportsRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *RoyaltyReportsRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(royaltyReportsPath, h.getRoyaltyReportsList)
	groups.AuthUser.GET(royaltyReportsIdPath, h.getRoyaltyReport)
	groups.AuthUser.POST(royaltyReportsIdDownloadPath, h.downloadRoyaltyReport)
	groups.AuthUser.GET(royaltyReportsTransactionsPath, h.listRoyaltyReportOrders)
	groups.AuthUser.POST(royaltyReportsTransactionsDownloadPath, h.downloadRoyaltyReportOrders)
	groups.AuthUser.POST(royaltyReportsAcceptPath, h.merchantReviewRoyaltyReport)
	groups.AuthUser.POST(royaltyReportsDeclinePath, h.merchantDeclineRoyaltyReport)
	groups.SystemUser.POST(royaltyReportsChangePath, h.changeRoyaltyReport)
}

// @summary Get the royalty reports list
// @desc Get the royalty reports list. This list can be filtered.
// @id royaltyReportsPathGetRoyaltyReportsList
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.RoyaltyReportsPaginate Returns the royalty reports list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param period_from query {integer} false The first date of the period for which the report results are calculated.
// @param period_to query {integer} false The last date of the period for which the report results are calculated.
// @param status query {[]string} false The list of the reports' statuses. Available values: pending, confirmed, paying, paid, dispute, canceled.
// @param limit query {integer} false The number of reports returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/royalty_reports [get]
func (h *RoyaltyReportsRoute) getRoyaltyReportsList(ctx echo.Context) error {
	req := &billingpb.ListRoyaltyReportsRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListRoyaltyReports(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "ListRoyaltyReports", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Get the royalty report
// @desc Get the royalty report using the report ID
// @id royaltyReportsIdPathGetRoyaltyReport
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.RoyaltyReport Returns the royalty report
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /admin/api/v1/royalty_reports/{report_id} [get]
func (h *RoyaltyReportsRoute) getRoyaltyReport(ctx echo.Context) error {
	req := &billingpb.GetRoyaltyReportRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetRoyaltyReport(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetRoyaltyReport")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Export the royalty report
// @desc Export the royalty report using the report ID
// @id royaltyReportsIdDownloadPathDownloadRoyaltyReport
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @body reporterpb.ReportFile
// @success 200 {object} reporterpb.CreateFileResponse Returns the report file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /admin/api/v1/royalty_reports/{report_id}/download [post]
func (h *RoyaltyReportsRoute) downloadRoyaltyReport(ctx echo.Context) error {
	req := &reporterpb.ReportFile{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = common.ExtractUserContext(ctx).Id
	req.ReportType = reporterpb.ReportTypeRoyalty
	params := map[string]interface{}{
		reporterpb.ParamsFieldId: ctx.Param(common.RequestParameterReportId),
	}

	return h.dispatch.RequestReportFile(ctx, req, params)
}

// @summary Get the transactions list included in the royalty report
// @desc Get the transactions list included in the royalty report
// @id royaltyReportsTransactionsPathListRoyaltyReportOrders
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.TransactionsPaginate Returns the transactions list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @param limit query {integer} false The number of transactions returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/royalty_reports/{report_id}/transactions [get]
func (h *RoyaltyReportsRoute) listRoyaltyReportOrders(ctx echo.Context) error {
	req := &billingpb.ListRoyaltyReportOrdersRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListRoyaltyReportOrders(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "ListRoyaltyReportOrders", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Export the file of the transactions list included in the royalty report
// @desc Export the file of the transactions list included in the royalty report
// @id royaltyReportsTransactionsDownloadPathDownloadRoyaltyReportOrders
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @body reporterpb.ReportFile
// @success 200 {object} reporterpb.CreateFileResponse Returns the file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /admin/api/v1/royalty_reports/{report_id}/transactions/download [post]
func (h *RoyaltyReportsRoute) downloadRoyaltyReportOrders(ctx echo.Context) error {
	req := &reporterpb.ReportFile{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = common.ExtractUserContext(ctx).Id
	req.ReportType = reporterpb.ReportTypeRoyaltyTransactions
	params := map[string]interface{}{
		reporterpb.ParamsFieldId: ctx.Param(common.RequestParameterReportId),
	}

	return h.dispatch.RequestReportFile(ctx, req, params)
}

// @summary Accept the royalty report by the merchant
// @desc Accept the royalty report by the merchant
// @id royaltyReportsAcceptPathMerchantReviewRoyaltyReport
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @body billingpb.MerchantReviewRoyaltyReportRequest
// @success 204 {string} Returns an empty response body if the royalty report has been successfully accepted
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /admin/api/v1/royalty_reports/{report_id}/accept [post]
func (h *RoyaltyReportsRoute) merchantReviewRoyaltyReport(ctx echo.Context) error {
	req := &billingpb.MerchantReviewRoyaltyReportRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.IsAccepted = true
	req.Ip = ctx.RealIP()

	res, err := h.dispatch.Services.Billing.MerchantReviewRoyaltyReport(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "MerchantReviewRoyaltyReport")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Dispute the royalty report by the merchant
// @desc Dispute the royalty report by the merchant
// @id royaltyReportsDeclinePathMerchantDeclineRoyaltyReport
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @body billingpb.MerchantReviewRoyaltyReportRequest
// @success 204 {string} Returns an empty response body if the royalty report has been declined
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /admin/api/v1/royalty_reports/{report_id}/decline [post]
func (h *RoyaltyReportsRoute) merchantDeclineRoyaltyReport(ctx echo.Context) error {
	req := &billingpb.MerchantReviewRoyaltyReportRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.IsAccepted = false
	req.Ip = ctx.RealIP()

	res, err := h.dispatch.Services.Billing.MerchantReviewRoyaltyReport(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "MerchantReviewRoyaltyReport")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Update the royalty report
// @desc Update the royalty report
// @id royaltyReportsChangePathChangeRoyaltyReport
// @tag Royalty reports
// @accept application/json
// @produce application/json
// @body billingpb.ChangeRoyaltyReportRequest
// @success 204 {string} Returns an empty response body if the royalty report has been successfully changed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param report_id path {string} true The unique identifier for the royalty report.
// @router /system/api/v1/royalty_reports/{report_id}/change [post]
func (h *RoyaltyReportsRoute) changeRoyaltyReport(ctx echo.Context) error {
	req := &billingpb.ChangeRoyaltyReportRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.Ip = ctx.RealIP()

	res, err := h.dispatch.Services.Billing.ChangeRoyaltyReport(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ChangeRoyaltyReport")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}
