package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	merchantListRoles           = "/merchants/roles"
	merchantUsers               = "/merchants/users"
	merchantInvite              = "/merchants/invite"
	merchantInviteResend        = "/merchants/users/resend"
	merchantUsersRole           = "/merchants/users/roles/:role_id"
	merchantSubscriptions       = "/merchants/subscriptions"
	merchantSubscriptionDetails = "/merchants/subscriptions/:subscription_id"
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
	groups.AuthUser.GET(merchantSubscriptions, h.getMerchantSubscriptions)
	groups.AuthUser.GET(merchantSubscriptionDetails, h.getMerchantSubscriptionDetails)

	groups.SystemUser.GET(merchantSubscriptions, h.getMerchantSubscriptions)
	groups.SystemUser.GET(merchantSubscriptionDetails, h.getMerchantSubscriptionDetails)
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

// @summary Get subscriptions for merchant
// @desc Get subscriptions for merchant
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantSubscriptionsResponse Returns the merchant subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/subscriptions [get]
func (h *MerchantUsersRoute) getMerchantSubscriptions(ctx echo.Context) error {
	req := &billingpb.GetMerchantSubscriptionsRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantSubscriptions(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetMerchantSubscriptions")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get orders for subscriptions
// @desc Get orders for subscriptions
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetSubscriptionOrdersResponse Returns list of orders for subscriptions
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/subscriptions/:id [get]
func (h *MerchantUsersRoute) getMerchantSubscriptionDetails(ctx echo.Context) error {
	req := &billingpb.GetSubscriptionOrdersRequest{}
	req.SubscriptionId = ctx.Param(common.RequestParameterSubscriptionId)

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetSubscriptionOrders(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetSubscriptionOrders")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	req1 := &billingpb.GetSubscriptionRequest{
		Id:     req.SubscriptionId,
		Cookie: req.Cookie,
	}

	res1, err := h.dispatch.Services.Billing.GetCustomerSubscription(ctx.Request().Context(), req1)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetCustomerSubscription", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res1.Status != billingpb.ResponseStatusOk {
		zap.L().Error("response code is not OK", zap.Int32("status", res1.Status), zap.String("method", "GetCustomerSubscription"))
		return echo.NewHTTPError(int(res1.Status), res1)
	}

	projRes, err := h.dispatch.Services.Billing.GetProject(ctx.Request().Context(), &billingpb.GetProjectRequest{
		MerchantId: res1.Subscription.MerchantId,
		ProjectId:  res1.Subscription.ProjectId,
	})
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetProject", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if projRes.Status != billingpb.ResponseStatusOk {
		zap.L().Error("response code is not OK", zap.Int32("status", projRes.Status), zap.String("method", "GetProject"))
		return echo.NewHTTPError(int(projRes.Status), projRes)
	}

	type subscriptionDetails struct {
		Amount       float32                 `json:"amount"`
		Currency     string                  `json:"currency"`
		Id           string                  `json:"id"`
		IsActive     bool                    `json:"is_active"`
		MaskedPan    string                  `json:"masked_pan"`
		ProjectName  string                  `json:"project_name"`
		StartDate    *timestamp.Timestamp    `json:"start_date"`
		Transactions []*billingpb.ShortOrder `json:"transactions"`
		Count        int32                   `json:"count"`
		CustomerId   string                  `json:"customer_id"`
	}

	subscription := res1.Subscription
	result := subscriptionDetails{
		Amount:       float32(subscription.Amount),
		Currency:     subscription.Currency,
		Id:           subscription.Id,
		IsActive:     subscription.IsActive,
		MaskedPan:    subscription.MaskedPan,
		ProjectName:  projRes.Item.Name["en"],
		StartDate:    subscription.CreatedAt,
		Transactions: res.List,
		Count:        res.Count,
		CustomerId:   subscription.CustomerUuid,
	}

	return ctx.JSON(http.StatusOK, result)
}
