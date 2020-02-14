package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billing "github.com/paysuper/paysuper-proto/go/billingpb"
	grpc "github.com/paysuper/paysuper-proto/go/billingpb"
	reporterPkg "github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"
)

const (
	orderIdPath          = "/order/:order_id"
	orderPath            = "/order"
	orderDownloadPath    = "/order/download"
	orderRefundsPath     = "/order/:order_id/refunds"
	orderRefundsIdsPath  = "/order/:order_id/refunds/:refund_id"
	orderReplaceCodePath = "/order/:order_id/replace_code"
)

const (
	errorTemplateName = "error.html"
)

type CreateOrderJsonProjectResponse struct {
	Id              string                         `json:"id"`
	PaymentFormUrl  string                         `json:"payment_form_url"`
	PaymentFormData *billingpb.PaymentFormJsonData `json:"payment_form_data,omitempty"`
}

type ListOrdersRequest struct {
	// The unique identifier for the merchant.
	MerchantId string `json:"merchant_id" validate:"required,hexadecimal,len=24"`
	// The supported file format. Available values: PDF, CSV, XLSX.
	FileType string `json:"file_type" validate:"required"`
	// The file template.
	Template string `json:"template" validate:"omitempty,hexadecimal"`
	// The unique identifier for the order.
	Id string `json:"id" validate:"omitempty,uuid"`
	// The list of projects.
	Project []string `json:"project" validate:"omitempty,dive,hexadecimal,len=24"`
	// The list of payment methods.
	PaymentMethod []string `json:"payment_method" validate:"omitempty,dive,hexadecimal,len=24"`
	// The list of the payer's countries.
	Country []string `json:"country" validate:"omitempty,dive,alpha,len=2"`
	// The list of orders' statuses. Available values: created, processed, canceled, rejected, refunded, chargeback, pending.
	Status []string `json:"status," validate:"omitempty,dive,alpha,oneof=created processed canceled rejected refunded chargeback pending"`
	// The start date when the payment was created.
	PmDateFrom int64 `json:"pm_date_from" validate:"omitempty,numeric,gt=0"`
	// The end date when the payment was closed.
	PmDateTo int64 `json:"pm_date_to" validate:"omitempty,numeric,gt=0"`
}

type OrderListRefundsBinder struct {
	dispatch common.HandlerSet
	provider.LMT
	cfg common.Config
}

func (b *OrderListRefundsBinder) Bind(i interface{}, ctx echo.Context) error {
	db := new(echo.DefaultBinder)
	err := db.Bind(i, ctx)

	if err != nil {
		return err
	}

	structure := i.(*billingpb.ListRefundsRequest)
	structure.OrderId = ctx.Param(common.RequestParameterOrderId)

	if structure.Limit <= 0 {
		structure.Limit = int64(b.cfg.LimitDefault)
	}

	return nil
}

type OrderRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewOrderRoute(set common.HandlerSet, cfg *common.Config) *OrderRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "OrderRoute"})
	return &OrderRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *OrderRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(orderPath, h.listOrdersPublic)
	groups.AuthUser.POST(orderDownloadPath, h.downloadOrdersPublic)
	groups.AuthUser.GET(orderIdPath, h.getOrderPublic) // TODO: Need a test

	groups.AuthUser.GET(orderRefundsPath, h.listRefunds)
	groups.AuthUser.GET(orderRefundsIdsPath, h.getRefund)
	groups.AuthUser.POST(orderRefundsPath, h.createRefund)
	groups.SystemUser.PUT(orderReplaceCodePath, h.replaceCode)
}

// @summary Get the full data about the order
// @desc Get the full data about the order using the order ID
// @id orderIdPathGetOrderPublic
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billing.OrderViewPublic Returns the order data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /admin/api/v1/order/{order_id} [get]
func (h *OrderRoute) getOrderPublic(ctx echo.Context) error {
	req := &billingpb.GetOrderRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetOrderPublic(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetOrderPublic")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the orders list
// @desc Get the orders list. This list can be filtered by the order's parameters.
// @id orderPathListOrdersPublic
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ListOrdersPublicResponseItem Returns the orders list
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param id query {string} false The unique identifier for the order.
// @param project query {[]string} false The list of projects.
// @param payment_method query {[]string} false The list of payment methods.
// @param country query {[]string} false The list of the payer's countries.
// @param status query {[]string} false The list of orders' statuses. Available values: created, processed, canceled, rejected, refunded, chargeback, pending.
// @param account query {string} false The payer account (for instance an account in the merchant's project, the account in the payment system, the payer email, etc.)
// @param pm_date_from query {integer} false The start date when the payment was created.
// @param pm_date_to query {integer} false The end date when the payment was closed.
// @param project_date_from query {integer} false The end date when the payment was created in the project.
// @param project_date_to query {integer} false The end date when the payment was closed in the project.
// @param quick_search query {string} false The search string that contains multiple fields - the unique identifier for the order, the user external identifier, the project order identifier, the project's name, the payment method's name.
// @param limit query {integer} true The number of orders returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @param sort query {[]string} false The list of the order's fields for sorting.
// @param type query {string} false The sales type. Available values: simple, product, key.
// @param hide_test query {boolean} false Has a true value for getting only production orders.
// @router /admin/api/v1/order [get]
func (h *OrderRoute) listOrdersPublic(ctx echo.Context) error {
	req := &billingpb.ListOrdersRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit <= 0 {
		req.Limit = int64(h.cfg.LimitDefault)
	}

	if req.Offset <= 0 {
		req.Offset = int64(h.cfg.OffsetDefault)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.FindAllOrdersPublic(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "FindAllOrdersPublic")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Export the orders list
// @desc Export the orders list into a PDF, CSV, XLSX
// @id orderDownloadPathDownloadOrdersPublic
// @tag Order
// @accept application/json
// @produce application/pdf, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @body ListOrdersRequest
// @success 200 {string} Returns the file with the orders list
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/order/download [post]
func (h *OrderRoute) downloadOrdersPublic(ctx echo.Context) error {
	req := &ListOrdersRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	file := &common.ReportFileRequest{
		ReportType: reporterPkg.ReportTypeTransactions,
		FileType:   req.FileType,
		MerchantId: req.MerchantId,
		Params: map[string]interface{}{
			reporterPkg.ParamsFieldStatus:        req.Status,
			reporterPkg.ParamsFieldPaymentMethod: req.PaymentMethod,
			reporterPkg.ParamsFieldDateFrom:      req.PmDateFrom,
			reporterPkg.ParamsFieldDateTo:        req.PmDateTo,
		},
	}

	return h.dispatch.RequestReportFile(ctx, file)
}

// @summary Get the refund data
// @desc Get the refund data using the order and refund IDs
// @id orderRefundsIdsPathGetRefund
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billing.Refund Returns the refund data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @param refund_id path {string} true The unique identifier for the refund.
// @router /admin/api/v1/order/{order_id}/refunds/{refund_id} [get]
func (h *OrderRoute) getRefund(ctx echo.Context) error {
	req := &billingpb.GetRefundRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetRefund(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetRefund")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the order's refunds list
// @desc Get the order's refunds list using the order ID
// @id orderRefundsPathListRefunds
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.ListRefundsResponse Returns the order's refunds list
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @param limit query {integer} true The number of refunds returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/order/{order_id}/refunds [get]
func (h *OrderRoute) listRefunds(ctx echo.Context) error {
	req := &billingpb.ListRefundsRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListRefunds(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ListRefunds")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Replaces the activation code in the order
// @desc Replaces the activation code in the order
// @id orderReplaceCodePathReplaceCode
// @tag Order
// @accept application/json
// @produce application/json
// @body grpc.ChangeCodeInOrderRequest
// @success 200 {object} billing.Order Returns the order data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /system/api/v1/order/{order_id}/replace_code [put]
func (h *OrderRoute) replaceCode(ctx echo.Context) error {
	req := &billingpb.ChangeCodeInOrderRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.OrderId = ctx.Param("order_id")
	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res := &billingpb.ChangeCodeInOrderResponse{}

	res, err := h.dispatch.Services.Billing.ChangeCodeInOrder(ctx.Request().Context(), req)
	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ChangeCodeInOrder")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Order)
}

// @summary Create a refund
// @desc Create a refund using the order ID
// @id orderRefundsPathCreateRefund
// @tag Order
// @accept application/json
// @produce application/json
// @body grpc.CreateRefundRequest
// @success 200 {object} billing.Refund Returns the refund data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The country not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /admin/api/v1/order/{order_id}/refunds [post]
func (h *OrderRoute) createRefund(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &billingpb.CreateRefundRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.OrderId = ctx.Param(common.RequestParameterOrderId)
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	req.CreatorId = authUser.Id
	res, err := h.dispatch.Services.Billing.CreateRefund(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "CreateRefund")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusCreated, res.Item)
}
