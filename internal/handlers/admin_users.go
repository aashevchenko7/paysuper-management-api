package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

type AdminUsersRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

const (
	users             = "/users"
	adminListRoles    = "/users/roles"
	adminUserInvite   = "/users/invite"
	adminResendInvite = "/users/resend"
	adminUserRole     = "/users/roles/:role_id"
)

func NewAdminUsersRoute(set common.HandlerSet, cfg *common.Config) *AdminUsersRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "AdminUsersRoute"})
	return &AdminUsersRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *AdminUsersRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(users, h.listUsers)
	groups.SystemUser.PUT(adminUserRole, h.changeRole)
	groups.SystemUser.POST(adminUserInvite, h.sendInvite)
	groups.SystemUser.POST(adminResendInvite, h.resendInvite)
	groups.SystemUser.GET(adminListRoles, h.listRoles)
	groups.SystemUser.DELETE(adminUserRole, h.deleteUser)
	groups.SystemUser.GET(adminUserRole, h.getUser)
}

// @summary Update the admin user role
// @desc Update the admin user role using the role ID
// @id adminUserRoleChangeRole
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @body billingpb.ChangeRoleForAdminUserRequest
// @success 200 {string} Returns an empty response body if the user's role was successfully changed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /system/api/v1/users/roles/{role_id} [put]
func (h *AdminUsersRoute) changeRole(ctx echo.Context) error {
	req := &billingpb.ChangeRoleForAdminUserRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ChangeRoleForAdminUser(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ChangeRoleForAdminUser")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

// @summary Get the admin users list
// @desc Get the admin users list
// @id usersListUsers
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @success 200 {object} []billingpb.UserRole Returns the admin users list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/users [get]
func (h *AdminUsersRoute) listUsers(ctx echo.Context) error {
	res, err := h.dispatch.Services.Billing.GetAdminUsers(ctx.Request().Context(), &billingpb.EmptyRequest{})

	if err != nil {
		return h.dispatch.SrvCallHandler(&billingpb.EmptyRequest{}, err, billingpb.ServiceName, "GetAdminUsers")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Users)
}

// @summary Send an invitation to the admin user
// @desc Send an invitation to add the user as the administrator
// @id adminUserInviteSendInvite
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @body billingpb.InviteUserAdminRequest
// @success 200 {object} billingpb.InviteUserAdminResponse Returns the admin user role data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/users/roles/invite [post]
func (h *AdminUsersRoute) sendInvite(ctx echo.Context) error {
	req := &billingpb.InviteUserAdminRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.InviteUserAdmin(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "InviteUserAdmin")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Resend an invitation to the user
// @desc Resend an invitation to add the user as the administrator
// @id adminResendInviteResendInvite
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @body billingpb.ResendInviteAdminRequest
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns an empty response body if the user's invitation was successfully send
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/users/roles/resend [post]
func (h *AdminUsersRoute) resendInvite(ctx echo.Context) error {
	req := &billingpb.ResendInviteAdminRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ResendInviteAdmin(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ResendInviteAdmin")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the admin roles
// @desc Get the admin roles
// @id adminListRolesListRoles
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetRoleListResponse Returns the admin roles data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/roles [get]
func (h *AdminUsersRoute) listRoles(ctx echo.Context) error {
	req := &billingpb.GetRoleListRequest{Type: billingpb.RoleTypeSystem}
	res, err := h.dispatch.Services.Billing.GetRoleList(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetRoleList")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the admin user
// @desc Delete the admin user
// @id adminUserRoleDeleteUser
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.EmptyResponseWithStatus Returns an empty response body if the user was successfully removed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /system/api/v1/users/roles/{role_id} [delete]
func (h *AdminUsersRoute) deleteUser(ctx echo.Context) error {
	req := &billingpb.AdminRoleRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.DeleteAdminUser(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "DeleteAdminUser")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the admin user role
// @desc Get the admin user role data
// @id adminUserRoleGetUser
// @tag Admin user roles
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.UserRoleResponse Returns the admin user role data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param role_id path {string} true The unique identifier for the role.
// @router /system/api/v1/users/roles/{role_id} [get]
func (h *AdminUsersRoute) getUser(ctx echo.Context) error {
	req := &billingpb.AdminRoleRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetAdminUserRole(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetAdminUserRole")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}
