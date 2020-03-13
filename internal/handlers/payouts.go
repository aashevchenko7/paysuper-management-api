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
	payoutsPath           = "/payout_documents"
	payoutsIdPath         = "/payout_documents/:payout_document_id"
	payoutsIdDownloadPath = "/payout_documents/:payout_document_id/download"
	payoutsIdReportsPath  = "/payout_documents/:payout_document_id/reports"
)

type PayoutDocumentsRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewPayoutDocumentsRoute(set common.HandlerSet, cfg *common.Config) *PayoutDocumentsRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "PayoutDocumentsRoute"})
	return &PayoutDocumentsRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *PayoutDocumentsRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(payoutsPath, h.getPayoutDocumentsList)
	groups.AuthUser.GET(payoutsIdPath, h.getPayoutDocument)
	groups.AuthUser.POST(payoutsIdDownloadPath, h.downloadPayoutDocument)
	groups.AuthUser.GET(payoutsIdReportsPath, h.getPayoutRoyaltyReports)
	groups.AuthUser.POST(payoutsPath, h.createPayoutDocument)
	groups.SystemUser.POST(payoutsIdPath, h.updatePayoutDocument)

}

// @summary Get the list of payout documents
// @desc Get the list of payout documents. This list can be filtered.
// @id payoutsPathGetPayoutDocumentsList
// @tag Payouts
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.PayoutDocumentsPaginate Returns the list of the payout documents
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param payout_document_id query {string} false The unique identifier for the payout document.
// @param status query {[]string} false The list of documents' statuses. Available values: skip, pending, in_progress, paid, canceled, failed.
// @param date_from query {integer} false The payout period start date.
// @param date_to query {integer} false The payout period end date.
// @param limit query {integer} true The number of documents returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/payout_documents [get]
func (h *PayoutDocumentsRoute) getPayoutDocumentsList(ctx echo.Context) error {
	req := &billingpb.GetPayoutDocumentsRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetPayoutDocuments(ctx.Request().Context(), req)
	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetPayoutDocuments")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data)
}

// @summary Get the payout document
// @desc Get the payout document using the payout document ID
// @id payoutsIdPathGetPayoutDocument
// @tag Payouts
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.PayoutDocument Returns the payout document
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param payout_document_id path {string} true The unique identifier for the payout document.
// @router /admin/api/v1/payout_documents/{payout_document_id} [get]
func (h *PayoutDocumentsRoute) getPayoutDocument(ctx echo.Context) error {
	req := &billingpb.GetPayoutDocumentRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetPayoutDocument(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetPayoutDocument")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Export the payout document
// @desc Export the payout document using the payout document ID
// @id payoutsIdDownloadPathDownloadPayoutDocument
// @tag Payouts
// @accept application/json
// @produce application/json
// @body reporterpb.ReportFile
// @success 200 {object} reporterpb.CreateFileResponse Returns the payout document file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param payout_document_id path {string} true The unique identifier for the payout document.
// @router /admin/api/v1/payout_documents/{payout_document_id}/download [post]
func (h *PayoutDocumentsRoute) downloadPayoutDocument(ctx echo.Context) error {
	req := &reporterpb.ReportFile{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = common.ExtractUserContext(ctx).Id
	req.ReportType = reporterpb.ReportTypePayout
	params := map[string]interface{}{
		reporterpb.ParamsFieldId: ctx.Param(common.RequestPayoutDocumentId),
	}

	return h.dispatch.RequestReportFile(ctx, req, params)
}

// @summary Create the payout documents
// @desc Create the payout documents
// @id payoutsPathCreatePayoutDocument
// @tag Payouts
// @accept application/json
// @produce application/json
// @body billingpb.CreatePayoutDocumentRequest
// @success 200 {object} []billingpb.PayoutDocument Returns the list of the payout documents
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/payout_documents [post]
func (h *PayoutDocumentsRoute) createPayoutDocument(ctx echo.Context) error {
	req := &billingpb.CreatePayoutDocumentRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Ip = ctx.RealIP()
	req.Initiator = billingpb.RoyaltyReportChangeSourceMerchant

	err = h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreatePayoutDocument(ctx.Request().Context(), req)
	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "CreatePayoutDocument")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Items)
}

// @summary Update the payout document
// @desc Update the payout document using the payout document ID
// @id payoutsIdPathUpdatePayoutDocument
// @tag Payouts
// @accept application/json
// @produce application/json
// @body billingpb.UpdatePayoutDocumentRequest
// @success 200 {object} billingpb.PayoutDocument Returns the payout document
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param payout_document_id path {string} true The unique identifier for the payout document.
// @router /system/api/v1/payout_documents/{payout_document_id} [post]
func (h *PayoutDocumentsRoute) updatePayoutDocument(ctx echo.Context) error {
	req := &billingpb.UpdatePayoutDocumentRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}
	req.Ip = ctx.RealIP()

	res, err := h.dispatch.Services.Billing.UpdatePayoutDocument(ctx.Request().Context(), req)
	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "UpdatePayoutDocument")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the royalty reports in the payout documents
// @desc Get the royalty reports in the payout documents
// @id payoutsIdReportsPathGetPayoutRoyaltyReports
// @tag Payouts
// @accept application/json
// @produce application/json
// @success 200 {object} []billingpb.RoyaltyReport Returns the list of the payout documents
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param payout_document_id path {string} true The unique identifier for the payout document.
// @router /admin/api/v1/payout_documents/{payout_document_id}/reports [get]
func (h *PayoutDocumentsRoute) getPayoutRoyaltyReports(ctx echo.Context) error {
	req := &billingpb.GetPayoutDocumentRequest{}
	req.PayoutDocumentId = ctx.Param(common.RequestParameterId)

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetPayoutDocumentRoyaltyReports(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetPayoutDocumentRoyaltyReports")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Data.Items)
}
