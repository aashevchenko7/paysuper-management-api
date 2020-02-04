package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
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
	Id              string                    `json:"id"`
	PaymentFormUrl  string                    `json:"payment_form_url"`
	PaymentFormData *billingpb.PaymentFormJsonData `json:"payment_form_data,omitempty"`
}

type ListOrdersRequest struct {
	MerchantId    string   `json:"merchant_id" validate:"required,hexadecimal,len=24"`
	FileType      string   `json:"file_type" validate:"required"`
	Template      string   `json:"template" validate:"omitempty,hexadecimal"`
	Id            string   `json:"id" validate:"omitempty,uuid"`
	Project       []string `json:"project" validate:"omitempty,dive,hexadecimal,len=24"`
	PaymentMethod []string `json:"payment_method" validate:"omitempty,dive,hexadecimal,len=24"`
	Country       []string `json:"country" validate:"omitempty,dive,alpha,len=2"`
	Status        []string `json:"status," validate:"omitempty,dive,alpha,oneof=created processed canceled rejected refunded chargeback pending"`
	PmDateFrom    int64    `json:"pm_date_from" validate:"omitempty,numeric,gt=0"`
	PmDateTo      int64    `json:"pm_date_to" validate:"omitempty,numeric,gt=0"`
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
