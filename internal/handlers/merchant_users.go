package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
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

// @summary Update the merchant user role
// @desc Update the merchant user role using the role ID
// @id merchantUsersRoleСhangeRole
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @success 200 {string} Returns an empty response body if the user's role was successfully changed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /admin/api/v1/merchants/users/roles/{role_id} [put]
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

// @summary Get the merchant users list
// @desc Get the merchant users list
// @id merchantUsersGetMerchantUsers
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @success 200 {object} []billingpb.UserRole Returns the merchant users list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/users [get]
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

// @summary Send an invitation to the merchant user
// @desc Send an invitation to add the user as the merchant
// @id merchantInviteSendInvite
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @body billingpb.InviteUserMerchantRequest
// @success 200 {object} billingpb.InviteUserMerchantResponse Returns the merchant user role data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/invite [post]
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

// @summary Resend an invitation to the user
// @desc Resend an invitation to add the user as the merchant
// @id merchantInviteResendResendInvite
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @body billingpb.ResendInviteMerchantRequest
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns an empty response body if the user's invitation was successfully send
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/users/resend [post]
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

// @summary Get the merchant roles
// @desc Get the merchant roles
// @id merchantListRolesListRoles
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetRoleListResponse Returns the merchant roles data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/roles [get]
func (h *MerchantUsersRoute) listRoles(ctx echo.Context) error {
	req := &billingpb.GetRoleListRequest{Type: billingpb.RoleTypeMerchant}
	res, err := h.dispatch.Services.Billing.GetRoleList(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetRoleList")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the merchant user
// @desc Delete the merchant user
// @id merchantUsersRoleDeleteUser
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns an empty response body if the user was successfully removed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /admin/api/v1/merchants/users/roles/{role_id} [delete]
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

// @summary Get the merchant user role
// @desc Get the merchant user role
// @id merchantUsersRoleGetUser
// @tag Merchant user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.UserRoleResponse Returns the merchant user role data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /admin/api/v1/merchants/users/roles/{role_id} [get]
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
