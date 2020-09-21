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
	dashboardMainPath            = "/merchants/dashboard/main"
	dashboardRevenueDynamicsPath = "/merchants/dashboard/revenue_dynamics"
	dashboardBasePath            = "/merchants/dashboard/base"
	dashboardCustomersPath       = "/merchants/dashboard/customers"
)

type DashboardRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewDashboardRoute(set common.HandlerSet, cfg *common.Config) *DashboardRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "DashboardRoute"})
	return &DashboardRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *DashboardRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(dashboardMainPath, h.getMainReports)
	groups.AuthUser.GET(dashboardRevenueDynamicsPath, h.getRevenueDynamicsReport)
	groups.AuthUser.GET(dashboardBasePath, h.getBaseReports)

	groups.AuthUser.GET(dashboardCustomersPath, h.getCustomers)
}

func (h *DashboardRoute) getCustomers(ctx echo.Context) error {
	req := &billingpb.DashboardCustomerReportRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetDashboardCustomersReport(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetDashboardCustomersReport", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the main reports for the Dashboard
// @desc Get the main reports for the Dashboard such as Gross revenue, Total transactions, VAT, Average revenue per user (ARPU)
// @id dashboardMainPathGetMainReports
// @tag Dashboard
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.DashboardMainReport Returns the main reports data for the Dashboard
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param period query {string} false The fixed period. Available values: current_month, previous_month, current_quarter, previous_quarter, current_year, previous_year.
// @router /admin/api/v1/merchants/dashboard/main [get]
func (h *DashboardRoute) getMainReports(ctx echo.Context) error {
	req := &billingpb.GetDashboardMainRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetDashboardMainReport(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetDashboardMainReport", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the revenue dynamic report for the Dashboard
// @desc Get the revenue dynamic report for the Dashboard
// @id dashboardRevenueDynamicsPathGetRevenueDynamicsReport
// @tag Dashboard
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.DashboardRevenueDynamicReport Returns the revenue dynamic report data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param period query {string} false The fixed period. Available values: current_month, previous_month, current_quarter, previous_quarter, current_year, previous_year.
// @router /admin/api/v1/merchants/dashboard/revenue_dynamics [get]
func (h *DashboardRoute) getRevenueDynamicsReport(ctx echo.Context) error {
	req := &billingpb.GetDashboardMainRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetDashboardRevenueDynamicsReport(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetDashboardMainReport", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the base report for the Dashboard
// @desc Get the base report for the Dashboard such as Revenue by country, Sales today, Sources.
// @id dashboardBasePathGetBaseReports
// @tag Dashboard
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.DashboardBaseReports Returns the base report data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param period query {string} false The fixed period. Available values: current_day, previous_day, current_week, previous_week, current_month, previous_month, current_quarter, previous_quarter, current_year, previous_year.
// @router /admin/api/v1/merchants/dashboard/base [get]
func (h *DashboardRoute) getBaseReports(ctx echo.Context) error {
	req := &billingpb.GetDashboardBaseReportRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetDashboardBaseReport(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetDashboardMainReport", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}
