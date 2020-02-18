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
	operatingCompanyPath   = "/operating_company"
	operatingCompanyIdPath = "/operating_company/:id"
)

type OperatingCompanyRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewOperatingCompanyRoute(set common.HandlerSet, cfg *common.Config) *OperatingCompanyRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "OperatingCompanyRoute"})
	return &OperatingCompanyRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *OperatingCompanyRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(operatingCompanyPath, h.getOperatingCompanyList)
	groups.SystemUser.GET(operatingCompanyIdPath, h.getOperatingCompany)
	groups.SystemUser.POST(operatingCompanyPath, h.addOperatingCompany)
	groups.SystemUser.POST(operatingCompanyIdPath, h.updateOperatingCompany)

}

// @summary Get the operating companies list
// @desc Get the operating companies list
// @id operatingCompanyPathGetOperatingCompanyList
// @tag Operating company
// @accept application/json
// @produce application/json
// @success 200 {object} []billing.OperatingCompany Returns the operating companies list
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/operating_company [get]
func (h *OperatingCompanyRoute) getOperatingCompanyList(ctx echo.Context) error {
	req := &billingpb.EmptyRequest{}

	res, err := h.dispatch.Services.Billing.GetOperatingCompaniesList(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetOperatingCompaniesList", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Items)
}

// @summary Get the operating company
// @desc Get the operating company data using the operating company ID
// @id operatingCompanyIdPathGetOperatingCompany
// @tag Operating company
// @accept application/json
// @produce application/json
// @success 200 {object} billing.OperatingCompany Returns the operating company data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the operating company.
// @router /system/api/v1/operating_company/{id} [get]
func (h *OperatingCompanyRoute) getOperatingCompany(ctx echo.Context) error {
	req := &billingpb.GetOperatingCompanyRequest{
		Id: ctx.Param(common.RequestParameterId),
	}

	res, err := h.dispatch.Services.Billing.GetOperatingCompany(ctx.Request().Context(), req)
	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetOperatingCompaniesList", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}
	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}
	return ctx.JSON(http.StatusOK, res.Company)
}

// @summary Create a new operating company
// @desc Create a new operating company
// @id operatingCompanyPathAddOperatingCompany
// @tag Operating company
// @accept application/json
// @produce application/json
// @body billing.OperatingCompany
// @success 204 {string} Returns an empty response body if the operating company has been successfully added
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/operating_company [post]
func (h *OperatingCompanyRoute) addOperatingCompany(ctx echo.Context) error {
	return h.addOrUpdateOperatingCompany(ctx, "")
}

// @summary Update the operating company
// @desc Update the operating company
// @id operatingCompanyIdPathUpdateOperatingCompany
// @tag Operating company
// @accept application/json
// @produce application/json
// @body billing.OperatingCompany
// @success 204 {string} Returns an empty response body if the operating company has been successfully updated
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique identifier for the operating company.
// @router /system/api/v1/operating_company/{id} [post]
func (h *OperatingCompanyRoute) updateOperatingCompany(ctx echo.Context) error {
	return h.addOrUpdateOperatingCompany(ctx, ctx.Param(common.RequestParameterId))
}

func (h *OperatingCompanyRoute) addOrUpdateOperatingCompany(ctx echo.Context, operatingCompanyId string) error {
	req := &billingpb.OperatingCompany{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Id = operatingCompanyId

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.AddOperatingCompany(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "AddOperatingCompany", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusNoContent)
}
