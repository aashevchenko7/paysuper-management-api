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
	projectsPath    = "/projects"
	projectsIdPath  = "/projects/:project_id"
	projectsSkuPath = "/projects/:project_id/sku"
)

type ProjectRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewProjectRoute(set common.HandlerSet, cfg *common.Config) *ProjectRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "ProjectRoute"})
	return &ProjectRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *ProjectRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(projectsPath, h.listProjects)
	groups.AuthUser.GET(projectsIdPath, h.getProject)
	groups.AuthUser.POST(projectsPath, h.createProject)
	groups.AuthUser.PATCH(projectsIdPath, h.updateProject)
	groups.AuthUser.DELETE(projectsIdPath, h.deleteProject)
	groups.AuthUser.POST(projectsSkuPath, h.checkSku)
}

// @summary Create a new project
// @desc Create a new project
// @id projectsPathCreateProject
// @tag Project
// @accept application/json
// @produce application/json
// @body billing.Project
// @success 201 {object} billing.Project Returns the project data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/projects [post]
func (h *ProjectRoute) createProject(ctx echo.Context) error {
	req := &billingpb.Project{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	authUser := common.ExtractUserContext(ctx)
	if req.MerchantId != authUser.MerchantId {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorIncorrectMerchantId)
	}

	if len(req.CallbackProtocol) == 0 {
		req.CallbackProtocol = billingpb.ProjectCallbackProtocolEmpty
	}

	// vat payer is seller by default on project creation
	req.VatPayer = billingpb.VatPayerSeller

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ChangeProject(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusCreated, res.Item)
}

// @summary Update the project
// @desc Update the project using the project ID
// @id projectsIdPathUpdateProject
// @tag Project
// @accept application/json
// @produce application/json
// @body billing.Project
// @success 200 {object} billing.Project Returns the project data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param project_id path {string} true The unique identifier for the project.
// @router /admin/api/v1/projects/{project_id} [patch]
func (h *ProjectRoute) updateProject(ctx echo.Context) error {
	req := &billingpb.Project{}
	binder := common.NewChangeProjectRequestBinder(h.dispatch, h.cfg)

	if err := binder.Bind(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	authUser := common.ExtractUserContext(ctx)
	if req.MerchantId != authUser.MerchantId {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorIncorrectMerchantId)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ChangeProject(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the project data
// @desc Get the project data using the project ID
// @id projectsIdPathGetProject
// @tag Project
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ChangeProjectResponse Returns the project data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param project_id path {string} true The unique identifier for the project.
// @router /admin/api/v1/projects/{project_id} [get]
func (h *ProjectRoute) getProject(ctx echo.Context) error {
	req := &billingpb.GetProjectRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetProject(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the list of projects
// @desc Get the list of projects for the authorized merchant
// @id projectsPathListProjects
// @tag Project
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ListProjectsResponse Returns the list of projects. The list can be filtered by the project's name, status, and sorted by the project's fields.
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param limit query {integer} true The number of projects returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @param quick_search query {string} false The project's name for the quick search.
// @param status query {[]string} false The list of the project's statuses.
// @param sort query {[]integer} false The list of the project's fields for sorting.
// @router /admin/api/v1/projects [get]
func (h *ProjectRoute) listProjects(ctx echo.Context) error {
	req := &billingpb.ListProjectsRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit <= 0 {
		req.Limit = h.cfg.LimitDefault
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListProjects(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the project
// @desc Delete the project using the project ID
// @id projectsIdPathDeleteProject
// @tag Project
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ChangeProjectResponse Returns the project data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param project_id path {string} true The unique identifier for the project.
// @router /admin/api/v1/projects/{project_id} [delete]
func (h *ProjectRoute) deleteProject(ctx echo.Context) error {
	req := &billingpb.GetProjectRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeleteProject(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Check the project contains the SKU
// @desc Check the project contains the SKU using the project ID
// @id projectsSkuPathCheckSku
// @tag Project
// @accept application/json
// @produce application/json
// @body grpc.CheckSkuAndKeyProjectRequest
// @success 200 {string} Returns an empty response body if the SKU was found in this project
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param project_id path {string} true The unique identifier for the project.
// @router /admin/api/v1/projects/{project_id}/sku [post]
func (h *ProjectRoute) checkSku(ctx echo.Context) error {
	req := &billingpb.CheckSkuAndKeyProjectRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CheckSkuAndKeyProject(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}
