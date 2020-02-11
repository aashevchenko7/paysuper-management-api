package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billing "github.com/paysuper/paysuper-proto/go/billingpb"
	grpc "github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const testMerchantWebhook = "/projects/:id/webhook/testing"

type WebHookRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewWebHookRoute(set common.HandlerSet, cfg *common.Config) *WebHookRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "WebHookRoute"})
	return &WebHookRoute{
		dispatch: set,
		cfg:      *cfg,
		LMT:      &set.AwareSet,
	}
}

func (h *WebHookRoute) Route(groups *common.Groups) {
	groups.AuthUser.POST(testMerchantWebhook, h.sendWebhookTest)
}

// @summary Test the project's webhook settings
// @desc Test the project's webhook settings
// @id testMerchantWebhookPathSendWebhookTest
// @tag Project
// @accept application/json
// @produce application/json
// @body billing.OrderCreateRequest
// @success 200 {object} grpc.SendWebhookToMerchantResponse Returns the unique identifier for the order
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage Not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/projects/{project_id}/webhook/testing [post]
func (h *WebHookRoute) sendWebhookTest(ctx echo.Context) error {
	req := &billingpb.OrderCreateRequest{}
	errBind := h.dispatch.BindAndValidate(req, ctx)

	if errBind != nil {
		return errBind
	}

	if len(req.TestingCase) == 0 {
		h.L().Error(common.BindingErrorTemplate, logger.PairArgs("err", "testing case is empty"))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	res, err := h.dispatch.Services.Billing.SendWebhookToMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "SendWebhookToMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}
