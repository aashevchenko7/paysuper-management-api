package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"net/http"
)

const (
	merchantListRoles    = "/merchants/roles"
	merchantUsers        = "/merchants/users"
	merchantInvite       = "/merchants/invite"
	merchantInviteResend = "/merchants/users/resend"
	merchantUsersRole    = "/merchants/users/roles/:role_id"
)

type MerchantUsersRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewMerchantUsersRoute(set common.HandlerSet, cfg *common.Config) *MerchantUsersRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "MerchantUsersRoute"})
	return &MerchantUsersRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *MerchantUsersRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(merchantUsers, h.getMerchantUsers)
	groups.AuthUser.PUT(merchantUsersRole, h.changeRole)
	groups.AuthUser.POST(merchantInvite, h.sendInvite)
	groups.AuthUser.POST(merchantInviteResend, h.resendInvite)
	groups.AuthUser.GET(merchantListRoles, h.listRoles)
	groups.AuthUser.DELETE(merchantUsersRole, h.deleteUser)
	groups.AuthUser.GET(merchantUsersRole, h.getUser)
}

func (h *MerchantUsersRoute) changeRole(ctx echo.Context) error {
	req := &billingpb.ChangeRoleForMerchantUserRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ChangeRoleForMerchantUser(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ChangeRoleForMerchantUser")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *MerchantUsersRoute) getMerchantUsers(ctx echo.Context) error {
	req := &billingpb.GetMerchantUsersRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantUsers(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetMerchantUsers")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Users)
}

func (h *MerchantUsersRoute) sendInvite(ctx echo.Context) error {
	req := &billingpb.InviteUserMerchantRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.InviteUserMerchant(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "InviteUserMerchant")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *MerchantUsersRoute) resendInvite(ctx echo.Context) error {
	req := &billingpb.ResendInviteMerchantRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ResendInviteMerchant(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ResendInviteMerchant")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *MerchantUsersRoute) listRoles(ctx echo.Context) error {
	req := &billingpb.GetRoleListRequest{Type: billingpb.RoleTypeMerchant}
	res, err := h.dispatch.Services.Billing.GetRoleList(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetRoleList")
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *MerchantUsersRoute) deleteUser(ctx echo.Context) error {
	req := &billingpb.MerchantRoleRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.DeleteMerchantUser(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "DeleteMerchantUser")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *MerchantUsersRoute) getUser(ctx echo.Context) error {
	req := &billingpb.MerchantRoleRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantUserRole(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetMerchantUserRole")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}
