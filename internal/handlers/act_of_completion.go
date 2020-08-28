package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
)

const (
	actOfCompletionDownloadPath         = "/act_of_completion"
	actOfCompletionDownloadMerchantPath = "/act_of_completion/:merchant_id"
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
	groups.AuthUser.POST(actOfCompletionDownloadPath, h.download)
	groups.SystemUser.POST(actOfCompletionDownloadMerchantPath, h.download)
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
// @router /admin/api/v1/act_of_completion [post]
// @router /system/api/v1/act_of_completion/{merchant_id} [post]
func (h *ActOfCompletionApiV1) download(ctx echo.Context) error {
	req := &billingpb.ActOfCompletionRequest{}
	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	file := &reporterpb.ReportFile{
		MerchantId:       req.MerchantId,
		ReportType:       reporterpb.ReportTypeActOfCompletion,
		SendNotification: true,
		SkipPostProcess:  true,
	}

	if req.Format == "" {
		file.FileType = reporterpb.OutputExtensionPdf
	}

	params := map[string]interface{}{
		reporterpb.ParamsFieldDateFrom: req.DateFrom,
		reporterpb.ParamsFieldDateTo:   req.DateTo,
	}

	return h.dispatch.RequestReportFile(ctx, file, params)
}
