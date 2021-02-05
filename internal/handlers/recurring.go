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
	recurringPlanList   = "/projects/:project_id/subscriptions/plans"
	recurringPlan       = "/projects/:project_id/subscriptions/plans/:plan_id"
	recurringPlanDelete = "/projects/:project_id/subscriptions/plans/:plan_id/delete"
)

type RecurringRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewRecurringRoute(set common.HandlerSet, cfg *common.Config) *RecurringRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "RecurringRoute"})
	return &RecurringRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *RecurringRoute) Route(groups *common.Groups) {
	groups.AuthUser.POST(recurringPlanList, h.addRecurringPlan)
	groups.AuthUser.PUT(recurringPlan, h.updateRecurringPlan)
	groups.AuthUser.GET(recurringPlan, h.getRecurringPlan)
	groups.AuthUser.PATCH(recurringPlan, h.enableRecurringPlan)
	groups.AuthUser.DELETE(recurringPlan, h.disableRecurringPlan)
	groups.AuthUser.GET(recurringPlanList, h.getRecurringPlans)
	groups.AuthUser.DELETE(recurringPlanDelete, h.deleteRecurringPlan)
}

// @summary Add the recurring plan to merchant
// @desc Add the recurring plan for the merchant and project
// @id addRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @body billingpb.RecurringPlan
// @success 200 {object} billingpb.RecurringPlan Returns the added recurring plan
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans [post]
func (h *RecurringRoute) addRecurringPlan(ctx echo.Context) error {
	req := &billingpb.RecurringPlan{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.AddRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "AddRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Update the recurring plan of merchant
// @desc Update the recurring plan of merchant
// @id updateRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @body billingpb.RecurringPlan
// @success 204 {string} Returns an empty response body if the recurring plan successfully updated
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param plan_id path {string} Identity of recurring plan.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans/:plan_id [put]
func (h *RecurringRoute) updateRecurringPlan(ctx echo.Context) error {
	req := &billingpb.RecurringPlan{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.UpdateRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "UpdateRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Get the recurring plan of merchant
// @desc Get the recurring plan of merchant
// @id getRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @body billingpb.GetRecurringPlanRequest
// @success 200 {string} billingpb.RecurringPlan Returns the recurring plan.
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param plan_id path {string} Identity of recurring plan.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans/:plan_id [get]
func (h *RecurringRoute) getRecurringPlan(ctx echo.Context) error {
	req := &billingpb.GetRecurringPlanRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Enable the recurring plan
// @desc Enable the recurring plan
// @id enableRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the recurring plan successfully enabled
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param plan_id path {string} Identity of recurring plan.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans/:plan_id [patch]
func (h *RecurringRoute) enableRecurringPlan(ctx echo.Context) error {
	req := &billingpb.EnableRecurringPlanRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.EnableRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "EnableRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Disable the recurring plan
// @desc Disable the recurring plan
// @id disableRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the recurring plan successfully disabled
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param plan_id path {string} Identity of recurring plan.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans/:plan_id [patch]
func (h *RecurringRoute) disableRecurringPlan(ctx echo.Context) error {
	req := &billingpb.DisableRecurringPlanRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.DisableRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DisableRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Delete the recurring plan
// @desc Delete the recurring plan
// @id deleteRecurringPlan
// @tag Recurring
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the recurring plan successfully deleted
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param plan_id path {string} Identity of recurring plan.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans/:plan_id/delete [patch]
func (h *RecurringRoute) deleteRecurringPlan(ctx echo.Context) error {
	req := &billingpb.DeleteRecurringPlanRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.DeleteRecurringPlan(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "DeleteRecurringPlan", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Get the recurring plans of merchant
// @desc Get the recurring plans of merchant with search options and paging
// @id getRecurringPlans
// @tag Recurring
// @accept application/json
// @produce application/json
// @body billingpb.RecurringPlan
// @success 200 {object} billingpb.RecurringPlan Returns the added recurring plan
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 403 {object} billingpb.ResponseErrorMessage Access deny
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param project_id path {string} Identity of merchant project.
// @param external_id query {string} External plan identity.
// @param group_id query {string} Recurring plan group identity.
// @param query query {integer} Search string for finding by plan name.
// @param limit query {integer} The number of objects returned in one page. Default value is 100.
// @param offset query {integer} The ranking number of the first item on the page.
// @router /admin/api/v1/projects/:project_id/subscriptions/plans [get]
func (h *RecurringRoute) getRecurringPlans(ctx echo.Context) error {
	req := &billingpb.GetRecurringPlansRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetRecurringPlans(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetRecurringPlans", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}
