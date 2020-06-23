package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

type UserRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

const (
	inviteCheck   = "/user/invite/check"
	inviteApprove = "/user/invite/approve"
	getMerchants  = "/user/merchants"
)

func NewUserRoute(set common.HandlerSet, cfg *common.Config) *UserRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "UsersRoute"})
	return &UserRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *UserRoute) Route(groups *common.Groups) {
	groups.AuthProject.POST(inviteCheck, h.checkInvite)
	groups.AuthProject.POST(inviteApprove, h.approveInvite)
	groups.AuthProject.GET(getMerchants, h.getMerchants)

}

// @summary Check the invitation token
// @desc Check the invitation token
// @id inviteCheckCheckInvite
// @tag User
// @accept application/json
// @produce application/json
// @body billingpb.CheckInviteTokenRequest
// @success 200 {object} billingpb.CheckInviteTokenResponse Returns the user's role ID and type
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/invite/check [post]
func (h *UserRoute) checkInvite(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)

	req := &billingpb.CheckInviteTokenRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.Email = authUser.Email

	err := h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CheckInviteToken(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "CheckInviteToken", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessageUnableToCheckInviteToken)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Approve the user invitation
// @desc Approve the user invitation
// @id inviteApproveApproveInvite
// @tag User
// @accept application/json
// @produce application/json
// @body billingpb.AcceptInviteRequest
// @success 200 {object} billingpb.AcceptInviteResponse Returns the user's role data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/invite/approve [post]
func (h *UserRoute) approveInvite(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)

	req := &billingpb.AcceptInviteRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestDataInvalid)
	}

	req.UserId = authUser.Id
	req.Email = authUser.Email

	err := h.dispatch.Validate.Struct(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.AcceptInvite(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "AcceptInvite", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessageUnableToAcceptInvite)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the list of merchants for this user
// @desc Get the list of merchants for this user
// @id getMerchantsGetMerchants
// @tag User
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantsForUserResponse Returns the list of merchants
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 403 {object} billingpb.ResponseErrorMessage Access denied
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/merchants [get]
func (h *UserRoute) getMerchants(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)

	req := &billingpb.GetMerchantsForUserRequest{UserId: authUser.Id}

	res, err := h.dispatch.Services.Billing.GetMerchantsForUser(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMerchantsForUser", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}
