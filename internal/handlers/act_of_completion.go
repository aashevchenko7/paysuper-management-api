package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
)

const (
	actOfCompletionListPath             = "/act_of_completion/list"
	actOfCompletionListMerchantPath     = "/act_of_completion/:merchant_id/list"
	actOfCompletionDownloadPath         = "/act_of_completion/download"
	actOfCompletionDownloadMerchantPath = "/act_of_completion/:merchant_id/download"
)

type ActOfCompletionApiV1 struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewActOfCompletionApiV1(set common.HandlerSet, cfg *common.Config) *ActOfCompletionApiV1 {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "ActOfCompletionApiV1"})
	return &ActOfCompletionApiV1{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *ActOfCompletionApiV1) Route(groups *common.Groups) {
	groups.AuthUser.GET(actOfCompletionListPath, h.list)
	groups.SystemUser.GET(actOfCompletionListMerchantPath, h.list)
	groups.AuthUser.POST(actOfCompletionDownloadPath, h.download)
	groups.SystemUser.POST(actOfCompletionDownloadMerchantPath, h.download)
}

// @summary Get list of acts of completion
// @desc Returns list of acts of completion from merchant's first payment date to last month
// @id actsOfCompletionList
// @tag Act Of Completion
// @accept application/json
// @produce application/json
// @body billingpb.ActsOfCompletionListRequest
// @success 200 {object} billingpb.ActsOfCompletionListResponse Returns the list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/act_of_completion/list [get]
// @router /system/api/v1/act_of_completion/{merchant_id}/list [get]
func (h *ActOfCompletionApiV1) list(ctx echo.Context) error {
	req := &billingpb.ActsOfCompletionListRequest{}
	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetActsOfCompletionList(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetActsOfCompletionList", req)
		return ctx.Render(http.StatusBadRequest, errorTemplateName, map[string]interface{}{})
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	// prevent null in response if no Items found
	if len(res.Items) == 0 {
		return ctx.JSON(http.StatusOK, []string{})
	}

	return ctx.JSON(http.StatusOK, res.Items)
}

// @summary Create an act of completion
// @desc Create an act of completed work for the specified period
// @id actOfCompletionCreate
// @tag Act Of Completion
// @accept application/json
// @produce application/json
// @body billingpb.ActOfCompletionRequest
// @success 200 {object} reporterpb.CreateFileResponse Returns the document file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/act_of_completion/download [post]
// @router /system/api/v1/act_of_completion/{merchant_id}/download [post]
func (h *ActOfCompletionApiV1) download(ctx echo.Context) error {
	req := &billingpb.ActOfCompletionRequest{}
	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	file := &reporterpb.ReportFile{
		MerchantId:       req.MerchantId,
		ReportType:       reporterpb.ReportTypeActOfCompletion,
		FileType:         req.FileType,
		SendNotification: true,
		SkipPostProcess:  true,
	}

	if file.FileType == "" {
		file.FileType = reporterpb.OutputExtensionPdf
	}

	params := map[string]interface{}{
		reporterpb.ParamsFieldDateFrom: req.DateFrom,
		reporterpb.ParamsFieldDateTo:   req.DateTo,
	}

	return h.dispatch.RequestReportFile(ctx, file, params)
}
