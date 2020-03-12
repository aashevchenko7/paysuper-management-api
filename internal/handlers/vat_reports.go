package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"
	"strings"
)

const (
	vatReportsPath                = "/vat_reports"
	vatReportsCountryPath         = "/vat_reports/country/:country"
	vatReportsCountryDownloadPath = "/vat_reports/country/:country/download"
	vatReportsDetailsPath         = "/vat_reports/details/:id"
	vatReportsDetailsDownloadPath = "/vat_reports/details/:id/download"
	vatReportsStatusPath          = "/vat_reports/status/:id"
)

type VatReportsRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewVatReportsRoute(set common.HandlerSet, cfg *common.Config) *VatReportsRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "VatReportsRoute"})
	return &VatReportsRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *VatReportsRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(vatReportsPath, h.getVatReportsDashboard)
	groups.SystemUser.GET(vatReportsCountryPath, h.getVatReportsForCountry)
	groups.SystemUser.POST(vatReportsCountryDownloadPath, h.downloadVatReportsForCountry)
	groups.SystemUser.GET(vatReportsDetailsPath, h.getVatReportTransactions)
	groups.SystemUser.POST(vatReportsDetailsDownloadPath, h.downloadVatReportTransactions)
	groups.SystemUser.POST(vatReportsStatusPath, h.updateVatReportStatus)
}

// @summary Get the VAT reports list for the Dashboard
// @desc Get the VAT reports list for the Dashboard
// @id vatReportsPathGetVatReportsDashboard
// @tag VAT reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.VatReportsPaginate Returns the the VAT reports list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/vat_reports [get]
func (h *VatReportsRoute) getVatReportsDashboard(ctx echo.Context) error {

	res, err := h.dispatch.Services.Billing.GetVatReportsDashboard(ctx.Request().Context(), &billingpb.EmptyRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Get the VAT reports list by country
// @desc Get the VAT reports list by country
// @id vatReportsCountryPathGetVatReportsForCountry
// @tag VAT reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.VatReportsPaginate Returns the the VAT reports list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param country path {string} true The country code.
// @param sort query {[]string} false The list of VAT fields for sorting.
// @param limit query {integer} true The number of reports returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /system/api/v1/vat_reports/country/{country} [get]
func (h *VatReportsRoute) getVatReportsForCountry(ctx echo.Context) error {
	req := &billingpb.VatReportsRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.Country = strings.ToUpper(ctx.Param(common.RequestParameterCountry))

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetVatReportsForCountry(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Export the VAT reports list filtered by country
// @desc Export the VAT reports list filtered by country
// @id vatReportsCountryDownloadPathDownloadVatReportsForCountry
// @tag VAT reports
// @accept application/json
// @produce application/json
// @body reporterpb.ReportFile
// @success 200 {object} reporterpb.CreateFileResponse Returns the file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param country path {string} true The country code.
// @router /system/api/v1/vat_reports/country/{country}/download [post]
func (h *VatReportsRoute) downloadVatReportsForCountry(ctx echo.Context) error {
	req := &reporterpb.ReportFile{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	req.ReportType = reporterpb.ReportTypeVat
	params := map[string]interface{}{
		reporterpb.ParamsFieldCountry: ctx.Param(common.RequestParameterCountry),
	}

	return h.dispatch.RequestReportFile(ctx, req, params)
}

// @summary Get the VAT report transactions
// @desc Get the VAT report details transactions
// @id vatReportsDetailsPathGetVatReportTransactions
// @tag VAT reports
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.PrivateTransactionsPaginate Returns the VAT report transactions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the VAT report.
// @param sort query {[]string} false The list of transaction fields for sorting.
// @param limit query {integer} true The number of transactions returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /system/api/v1/vat_reports/details/{id} [get]
func (h *VatReportsRoute) getVatReportTransactions(ctx echo.Context) error {
	req := &billingpb.VatTransactionsRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.VatReportId = ctx.Param(common.RequestParameterId)

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetVatReportTransactions(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Export the VAT report transactions
// @desc Export the VAT report details transactions
// @id vatReportsDetailsDownloadPathDownloadVatReportTransactions
// @tag VAT reports
// @accept application/json
// @produce application/json
// @success 200 {object} reporterpb.CreateFileResponse Returns the file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the VAT report.
// @router /system/api/v1/vat_reports/details/{id}/download [post]
func (h *VatReportsRoute) downloadVatReportTransactions(ctx echo.Context) error {
	req := &reporterpb.ReportFile{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	req.ReportType = reporterpb.ReportTypeVatTransactions
	params := map[string]interface{}{
		reporterpb.ParamsFieldId: ctx.Param(common.RequestParameterId),
	}

	return h.dispatch.RequestReportFile(ctx, req, params)
}

// @summary Update the VAT report status
// @desc Update the VAT report status
// @id vatReportsStatusPathUpdateVatReportStatus
// @tag VAT reports
// @accept application/json
// @produce application/json
// @body billingpb.UpdateVatReportStatusRequest
// @success 204 {string} Returns an empty response body if the VAT report status was successfully changed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the VAT report.
// @router /system/api/v1/vat_reports/status/{id} [post]
func (h *VatReportsRoute) updateVatReportStatus(ctx echo.Context) error {

	req := &billingpb.UpdateVatReportStatusRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.NewValidationError(err.Error()))
	}

	req.Id = ctx.Param(common.RequestParameterId)

	if err = h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.UpdateVatReportStatus(ctx.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.NoContent(http.StatusNoContent)
}
