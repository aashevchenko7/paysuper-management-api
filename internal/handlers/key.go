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

const (
	keysIdPath = "/keys/:key_id"
)

type KeyRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewKeyRoute(set common.HandlerSet, cfg *common.Config) *KeyRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "KeyRoute"})
	return &KeyRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *KeyRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(keysIdPath, h.getKeyInfo)
}

// @summary Get the key's data
// @desc Get the key's data
// @id keysIdPathGetKeyInfo
// @tag Key, Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.Key Returns the key's data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param key_id path {string} true The unique identifier for the key.
// @router /admin/api/v1/keys/{key_id} [get]
func (h *KeyRoute) getKeyInfo(ctx echo.Context) error {
	req := &billingpb.KeyForOrderRequest{
		KeyId: ctx.Param("key_id"),
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetKeyByID(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Key)
}
