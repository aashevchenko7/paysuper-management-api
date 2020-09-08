package handlers

import (
	"encoding/json"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/golang/protobuf/ptypes"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

const (
	orderIdPath          = "/order/:order_id"
	orderPath            = "/order"
	orderDownloadPath    = "/order/download"
	orderRefundsPath     = "/order/:order_id/refunds"
	orderRefundsIdsPath  = "/order/:order_id/refunds/:refund_id"
	orderReplaceCodePath = "/order/:order_id/replace_code"
	orderGetLogsPath     = "/order/:order_id/logs"
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
	PmDateFrom string `json:"pm_date_from" validate:"omitempty,datetime"`
	// The end date when the payment was closed.
	PmDateTo string `json:"pm_date_to" validate:"omitempty,datetime"`
}

type cloudWatchLogSettings struct {
	group   string
	pattern func(order *billingpb.OrderViewPublic) string
	setter  func(result *GetOrderLogsResponse, value *LogOrder)
}

type cloudWatch struct {
	logSettings []*cloudWatchLogSettings
	instance    common.CloudWatchInterface
}

type LogRequest struct {
	Headers interface{} `json:"headers"`
	Body    interface{} `json:"body"`
}

type LogResponse struct {
	HttpStatus interface{} `json:"http_status"`
	LogRequest
}

type LogOrder struct {
	Date     time.Time    `json:"date"`
	Uri      interface{}  `json:"uri"`
	Request  *LogRequest  `json:"request"`
	Response *LogResponse `json:"response"`
}

type GetOrderLogsResponse struct {
	// The order's logs list of the payment creation process.
	Create []*LogOrder `json:"create"`
	// The order's logs list of the payment callback notification process received from the payment system.
	Callback []*LogOrder `json:"callback"`
	// The order's logs list of the notification about the payment status sent to the project.
	Notify []*LogOrder `json:"notify"`
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
	*cloudWatch
}

func NewOrderRoute(
	set common.HandlerSet,
	cloudWatchLog common.CloudWatchInterface,
	cfg *common.Config,
) *OrderRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "OrderRoute"})
	cloudWatch := &cloudWatch{
		logSettings: []*cloudWatchLogSettings{
			{
				group: cfg.AwsCloudWatchLogGroupBillingServer,
				pattern: func(order *billingpb.OrderViewPublic) string {
					return order.Id + " cardpay"
				},
				setter: func(result *GetOrderLogsResponse, value *LogOrder) {
					result.Create = append(result.Create, value)
				},
			},
			{
				group: cfg.AwsCloudWatchLogGroupManagementApi,
				pattern: func(order *billingpb.OrderViewPublic) string {
					return order.Id + " webhook"
				},
				setter: func(result *GetOrderLogsResponse, value *LogOrder) {
					result.Callback = append(result.Callback, value)
				},
			},
			{
				group: cfg.AwsCloudWatchLogGroupWebhookNotifier,
				pattern: func(order *billingpb.OrderViewPublic) string {
					return order.Uuid + " delivery_try"
				},
				setter: func(result *GetOrderLogsResponse, value *LogOrder) {
					result.Notify = append(result.Notify, value)
				},
			},
		},
		instance: cloudWatchLog,
	}

	return &OrderRoute{
		dispatch:   set,
		LMT:        &set.AwareSet,
		cloudWatch: cloudWatch,
		cfg:        *cfg,
	}
}

func (h *OrderRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(orderPath, h.listOrdersPublic)
	groups.AuthUser.GET(orderIdPath, h.getOrderPublic)
	groups.AuthUser.POST(orderDownloadPath, h.downloadOrdersPublic)
	groups.AuthUser.GET(orderRefundsPath, h.listRefunds)
	groups.AuthUser.GET(orderRefundsIdsPath, h.getRefund)
	groups.AuthUser.POST(orderRefundsPath, h.createRefund)

	groups.SystemUser.GET(orderPath, h.listOrdersPrivate)
	groups.SystemUser.GET(orderIdPath, h.getOrderPrivate)
	groups.SystemUser.GET(orderGetLogsPath, h.getOrderLogs)
	groups.SystemUser.GET(orderRefundsIdsPath, h.getRefund)
	groups.SystemUser.POST(orderRefundsPath, h.createRefund)
	groups.SystemUser.PUT(orderReplaceCodePath, h.replaceCode)

	groups.MerchantS2S.GET(orderPath, h.listOrdersS2s)
}

// @summary Get the full data about the order
// @desc Get the full data about the order using the order ID
// @id adminOrderIdPathGetOrderPublic
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.OrderViewPublic Returns the order data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /admin/api/v1/order/{order_id} [get]
func (h *OrderRoute) getOrderPublic(ctx echo.Context) error {
	res, err := h.getOrder(ctx, h.dispatch.Services.Billing.GetOrderPublic)

	if err != nil {
		return err
	}

	typed := res.(*billingpb.GetOrderPublicResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	return ctx.JSON(http.StatusOK, typed.Item)
}

// @summary Get the full private data about the order
// @desc Get the full private data about the order using the order ID
// @id systemOrderIdPathGetOrderPrivate
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.OrderViewPublic Returns the order data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /system/api/v1/order/{order_id} [get]
func (h *OrderRoute) getOrderPrivate(ctx echo.Context) error {
	res, err := h.getOrder(ctx, h.dispatch.Services.Billing.GetOrderPrivate)

	if err != nil {
		return err
	}

	typed := res.(*billingpb.GetOrderPrivateResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	return ctx.JSON(http.StatusOK, typed.Item)
}

// @summary Get the orders list
// @desc Get the orders list. This list can be filtered by the order's parameters.
// @id adminOrderPathListOrdersPublic
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListOrdersPublicResponseItem Returns the orders list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
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
	rsp, err := h.listOrders(ctx, h.dispatch.Services.Billing.FindAllOrdersPublic)

	if err != nil {
		return err
	}

	typed := rsp.(*billingpb.ListOrdersPublicResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	return ctx.JSON(http.StatusOK, typed.Item)
}

// @summary Get the private orders list
// @desc Get the private orders list. This list can be filtered by the order's parameters.
// @id systemOrderPathListOrdersPrivate
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListOrdersPublicResponseItem Returns the orders list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
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
// @router /system/api/v1/order [get]
func (h *OrderRoute) listOrdersPrivate(ctx echo.Context) error {
	rsp, err := h.listOrders(ctx, h.dispatch.Services.Billing.FindAllOrdersPrivate)

	if err != nil {
		return err
	}

	typed := rsp.(*billingpb.ListOrdersPrivateResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	return ctx.JSON(http.StatusOK, typed.Item)
}

// @summary Get the orders list
// @desc Get the orders list. This list can be filtered by the order's parameters.
// @id merchantS2SOrderPathListOrders
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListOrdersPublicResponseItem Returns the orders list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id query {string} false The unique identifier for the order in PaySuper's billing system.
// @param project query {[]string} false The list of projects.
// @param status query {[]string} false The list of orders' statuses. Available values: created, processed, canceled, rejected, refunded, chargeback, pending.
// @param account query {string} false The payer account (for instance an account in the merchant's project, the account in the payment system, the payer email, etc.)
// @param project_date_from query {integer} false The end date when the payment was created in the project.
// @param project_date_to query {integer} false The end date when the payment was closed in the project.
// @param invoice_id {string} false The unique identifier for the order in merchant's billing system.
// @param limit query {integer} true The number of orders returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @param sort query {[]string} false The list of the order's fields for sorting.
// @router /merchant/s2s/api/v1/order [get]
func (h *OrderRoute) listOrdersS2s(ctx echo.Context) error {
	rsp, err := h.listOrders(ctx, h.dispatch.Services.Billing.FindAllOrders)

	if err != nil {
		return err
	}

	typed := rsp.(*billingpb.ListOrdersResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	return ctx.JSON(http.StatusOK, typed.Item)
}

// @summary Export the orders list
// @desc Export the orders list
// @id orderDownloadPathDownloadOrdersPublic
// @tag Order
// @accept application/json
// @produce application/json
// @body ListOrdersRequest
// @success 200 {object} reporterpb.CreateFileResponse Returns the file ID
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/order/download [post]
func (h *OrderRoute) downloadOrdersPublic(ctx echo.Context) error {
	req := &ListOrdersRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	var (
		pmDateFrom = int64(0)
		pmDateTo   = int64(0)
	)

	if req.PmDateFrom != "" {
		dateFrom, err := time.Parse(billingpb.FilterDatetimeFormat, req.PmDateFrom)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
		}

		pmDateFrom = dateFrom.Unix()
	}

	if req.PmDateTo != "" {
		dateTo, err := time.Parse(billingpb.FilterDatetimeFormat, req.PmDateTo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
		}

		pmDateTo = dateTo.Unix()
	}

	file := &reporterpb.ReportFile{
		ReportType: reporterpb.ReportTypeTransactions,
		FileType:   req.FileType,
		MerchantId: req.MerchantId,
	}
	params := map[string]interface{}{
		reporterpb.ParamsFieldStatus:        req.Status,
		reporterpb.ParamsFieldPaymentMethod: req.PaymentMethod,
		reporterpb.ParamsFieldDateFrom:      pmDateFrom,
		reporterpb.ParamsFieldDateTo:        pmDateTo,
	}

	return h.dispatch.RequestReportFile(ctx, file, params)
}

// @summary Get the refund data
// @desc Get the refund data using the order and refund IDs
// @id orderRefundsIdsPathGetRefund
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.Refund Returns the refund data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @param refund_id path {string} true The unique identifier for the refund.
// @router /admin/api/v1/order/{order_id}/refunds/{refund_id} [get]
//
// @summary Get the refund data
// @desc Get the refund data using the order and refund IDs
// @id orderRefundsIdsPathGetRefundSystem
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.Refund Returns the refund data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @param refund_id path {string} true The unique identifier for the refund.
// @router /system/api/v1/order/{order_id}/refunds/{refund_id} [get]
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
// @success 200 {object} billingpb.ListRefundsResponse Returns the order's refunds list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
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
// @body billingpb.ChangeCodeInOrderRequest
// @success 200 {object} billingpb.Order Returns the order data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
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
// @body billingpb.CreateRefundRequest
// @success 200 {object} billingpb.Refund Returns the refund data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /admin/api/v1/order/{order_id}/refunds [post]
//
// @summary Create a refund
// @desc Create a refund using the order ID
// @id orderRefundsPathCreateRefundSystem
// @tag Order
// @accept application/json
// @produce application/json
// @body billingpb.CreateRefundRequest
// @success 200 {object} billingpb.Refund Returns the refund data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /system/api/v1/order/{order_id}/refunds [post]
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

// @summary Get the order's logs list
// @desc Get the order's logs list using the order ID
// @id orderLogsPathListLogs
// @tag Order
// @accept application/json
// @produce application/json
// @success 200 {object} GetOrderLogsResponse Returns the order's logs list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param order_id path {string} true The unique identifier for the order.
// @router /system/api/v1/order/{order_id}/logs [get]
func (h *OrderRoute) getOrderLogs(ctx echo.Context) error {
	res, err := h.getOrder(ctx, h.dispatch.Services.Billing.GetOrderPublic)

	if err != nil {
		return err
	}

	typed := res.(*billingpb.GetOrderPublicResponse)

	if typed.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(typed.Status), typed.Message)
	}

	order := typed.Item
	createdAt, err := ptypes.Timestamp(order.CreatedAt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	result := new(GetOrderLogsResponse)

	for _, val := range h.cloudWatch.logSettings {
		pattern := val.pattern(order)

		rsp, err := h.cloudWatch.instance.FilterLogEventsWithContext(
			ctx.Request().Context(),
			&cloudwatchlogs.FilterLogEventsInput{
				Limit:         aws.Int64(100),
				LogGroupName:  aws.String(val.group),
				StartTime:     aws.Int64(aws.TimeUnixMilli(createdAt)),
				EndTime:       aws.Int64(aws.TimeUnixMilli(createdAt.AddDate(0, 0, 7))),
				FilterPattern: aws.String(pattern),
			},
		)

		if err != nil {
			h.dispatch.AwareSet.L().Error(
				"get logs form amazon cloudwatch failed",
				logger.PairArgs(
					"group", val.group,
					"pattern", pattern,
				),
				logger.WithPrettyFields(logger.Fields{"err": err}),
			)
			continue
		}

		for _, event := range rsp.Events {
			log := make(map[string]interface{})
			err = json.Unmarshal([]byte(*event.Message), &log)

			if err != nil {
				continue
			}

			logOrder := &LogOrder{
				Date: aws.MillisecondsTimeValue(event.Timestamp),
				Uri:  log["msg"],
				Request: &LogRequest{
					Headers: log["request_headers"],
					Body:    log["request_body"],
				},
				Response: &LogResponse{
					HttpStatus: log["response_status"],
					LogRequest: LogRequest{
						Headers: log["response_headers"],
						Body:    log["response_body"],
					},
				},
			}
			val.setter(result, logOrder)
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

func (h *OrderRoute) listOrders(ctx echo.Context, fn interface{}) (interface{}, error) {
	req := &billingpb.ListOrdersRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit <= 0 {
		req.Limit = int64(h.cfg.LimitDefault)
	}

	if req.Offset <= 0 {
		req.Offset = int64(h.cfg.OffsetDefault)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	refFn := reflect.ValueOf(fn)
	fnName := runtime.FuncForPC(refFn.Pointer()).Name()
	returnValues := refFn.Call([]reflect.Value{reflect.ValueOf(ctx.Request().Context()), reflect.ValueOf(req)})

	if err := returnValues[1].Interface(); err != nil {
		return nil, h.dispatch.SrvCallHandler(req, err.(error), billingpb.ServiceName, fnName)
	}

	return returnValues[0].Interface(), nil
}

func (h *OrderRoute) getOrder(ctx echo.Context, fn interface{}) (interface{}, error) {
	req := &billingpb.GetOrderRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return nil, err
	}

	refFn := reflect.ValueOf(fn)
	fnName := runtime.FuncForPC(refFn.Pointer()).Name()
	returnValues := refFn.Call([]reflect.Value{reflect.ValueOf(ctx.Request().Context()), reflect.ValueOf(req)})

	if err := returnValues[1].Interface(); err != nil {
		return nil, h.dispatch.SrvCallHandler(req, err.(error), billingpb.ServiceName, fnName)
	}

	return returnValues[0].Interface(), nil
}
