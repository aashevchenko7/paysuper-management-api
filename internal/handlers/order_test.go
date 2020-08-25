package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	billMock "github.com/paysuper/paysuper-proto/go/billingpb/mocks"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/mock"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/url"
	"testing"
)

type OrderTestSuite struct {
	suite.Suite
	router *OrderRoute
	caller *test.EchoReqResCaller
}

func Test_Order(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

func (suite *OrderTestSuite) SetupTest() {
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
	}

	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: mock.NewBillingServerOkMock(),
	}
	cloudwatchMock := &mock.CloudWatchInterface{}
	cloudwatchMock.On("FilterLogEventsWithContext", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&cloudwatchlogs.FilterLogEventsOutput{
				Events: []*cloudwatchlogs.FilteredLogEvent{
					{
						Timestamp: aws.Int64(1586868621704),
						Message:   aws.String(`{"level":"info","ts":1586868647.0548847,"logger":"PAYONE_BILLING_SERVER","caller":"service/cardpay.go:867","msg":"/api/payments","request_headers":{"Authorization":["Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJxY0FzMjlSNUFPbEltSFRobWJPR3lCYXI4aHZOdEo5Ni1GZGwyS2pVa1ZrIn0.eyJqdGkiOiJmMGY4NGMwOC0wNzQ3LTQ0YjItOTBkYS1jYjk3ODI4YjBhNjYiLCJleHAiOjE1ODY4Njg5NDYsIm5iZiI6MCwiaWF0IjoxNTg2ODY4NjQ2LCJpc3MiOiJodHRwczovL3Nzby1zYW5kYm94LmNhcmRwYXkuY29tL2F1dGgvcmVhbG1zL2FwaSIsInN1YiI6ImI5ZjNkN2NlLWQzMGItNDhiMS1hZmU4LTU4ZmE0ZmQxY2U2YiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImFwaXYzLWF1dGgtYWdlbnQiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiIxODZjMmMwNi1iN2ZjLTRmYmYtYjc3Yi03M2I3YzNlYTExODAiLCJhY3IiOiIxIiwic2NvcGUiOiIiLCJ0ZXJtaW5hbCI6IjIxMTI1In0.PVk1tVSg29GyDS4KcCoX3TY4B_oTg-gCBRvim2XUGgJqmTkNJEvCgktXumx_nJhLk8nIWif7yXTuuoAalP-9bT2daToGnS-8ALP193Y21gZz3aJmUJiDguCf6BxvbP27P3-T-d0inrIp28iiXQx7SOSESGoRH8jHPrVf16KEAFC-2AgRIQPfEPfBps--M9r05TaP8MsJX5TiaWEum3p7juKNfPN-fDlBZbOnsNDlVq0aV1FcAPBaKGYuM58UVSDVGPeLGUMz5m5cJDs5TIvzjB60ao4SD-4CRtQOQeyVX6AaMnZKVgre502tLXyzj8WCZIe0QkrxiLplsQdfd5fGsA"],"Content-Type":["application/json"]},"request_body":"{\"request\":{\"id\":\"5e95b18d455b51545379c11a\",\"time\":\"2020-04-14T12:50:46Z\"},\"merchant_order\":{\"id\":\"5e95b18d455b51545379c11a\",\"description\":\"Payment by order # 5e95b18d455b51545379c11a\"},\"description\":\"Payment by order # 5e95b18d455b51545379c11a\",\"payment_method\":\"BANKCARD\",\"payment_data\":{\"currency\":\"USD\",\"amount\":30,\"dynamic_descriptor\":\"\",\"note\":\"\"},\"card_account\":{\"card\":{\"pan\":\"400000******0077\",\"holder\":\"CARDHOLDER\",\"security_code\":\"***\",\"expiration\":\"02/2022\"}},\"customer\":{\"email\":\"1CWutDD6dx1wpm9lDzrAzexKFlogkZHy@paysuper.com\",\"ip\":\"109.87.150.29\",\"id\":\"1CWutDD6dx1wpm9lDzrAzexKFlogkZHy\"},\"return_urls\":{\"cancel_url\":\"https://checkout.pay.super.com/pay/order/?result=fail\",\"decline_url\":\"https://checkout.pay.super.com/pay/order/?result=fail\",\"success_url\":\"https://checkout.pay.super.com/pay/order/?result=success\"}}","response_status":200,"response_headers":{"Cache-Control":["no-cache, no-store, max-age=0, must-revalidate"],"Cf-Cache-Status":["DYNAMIC"],"Cf-Ray":["583d8df28eeef40b-LHR"],"Content-Type":["application/json;charset=UTF-8"],"Date":["Tue, 14 Apr 2020 12:50:47 GMT"],"Expect-Ct":["max-age=604800, report-uri=\"https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct\""],"Expires":["0"],"Pragma":["no-cache"],"Server":["cloudflare"],"Set-Cookie":["__cfduid=d32ff7fb56a9824ead052a81739ca5d721586868646; expires=Thu, 14-May-20 12:50:46 GMT; path=/; domain=.cardpay.com; HttpOnly; SameSite=Lax"],"Strict-Transport-Security":["max-age=31536000 ; includeSubDomains"],"Vary":["Accept-Encoding,User-Agent"],"X-Content-Type-Options":["nosniff"],"X-Frame-Options":["DENY"],"X-Xss-Protection":["1; mode=block"]},"response_body":"{\"redirect_url\":\"https://sandbox.cardpay.com/MI/payments/redirect?token=FadFHG6cg5255eF047eBcCC7\",\"payment_data\":{\"id\":\"3549175\"}}"}`),
					},
				},
			},
			nil,
		)

	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewOrderRoute(set.HandlerSet, cloudwatchMock, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *OrderTestSuite) TearDownTest() {}

func (suite *OrderTestSuite) TestOrder_GetRefund_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Params(":refund_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + orderRefundsIdsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())

	refund := &billingpb.JsonRefund{}
	err = json.Unmarshal(res.Body.Bytes(), refund)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), refund.Id)
	assert.NotEmpty(suite.T(), refund.Currency)
	assert.Len(suite.T(), refund.Currency, 3)
}

func (suite *OrderTestSuite) TestOrder_GetRefund_RefundIdEmpty_Error() {

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsIdsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Regexp(suite.T(), common.NewValidationError("RefundId"), httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetRefund_OrderIdEmpty_Error() {

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsIdsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Regexp(suite.T(), common.NewValidationError("OrderId"), httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetRefund_BillingServerError() {

	suite.router.dispatch.Services.Billing = mock.NewBillingServerSystemErrorMock()

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Params(":refund_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + orderRefundsIdsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetRefund_BillingServer_RefundNotFound_Error() {

	suite.router.dispatch.Services.Billing = mock.NewBillingServerErrorMock()

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Params(":refund_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + orderRefundsIdsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListRefunds_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_ListRefunds_BindError() {

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Regexp(suite.T(), common.NewValidationError("OrderId"), httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListRefunds_BillingServerError() {

	suite.router.dispatch.Services.Billing = mock.NewBillingServerSystemErrorMock()

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_CreateRefund_Ok() {
	data := `{"amount": 10, "reason": "test"}`

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		BodyString(data).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_CreateRefund_BindError() {
	data := `{"amount": "qwerty", "reason": "test"}`

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		BodyString(data).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorRequestParamsIncorrect, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_CreateRefund_ValidationError() {
	data := `{"amount": -10, "reason": "test"}`

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		BodyString(data).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Regexp(suite.T(), common.NewValidationError("Amount"), httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_CreateRefund_BillingServerError() {
	data := `{"amount": 10, "reason": "test"}`

	suite.router.dispatch.Services.Billing = mock.NewBillingServerSystemErrorMock()

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		BodyString(data).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_CreateRefund_BillingServer_CreateError() {
	data := `{"amount": 10, "reason": "test"}`

	suite.router.dispatch.Services.Billing = mock.NewBillingServerErrorMock()

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":order_id", uuid.New().String()).
		Path(common.AuthUserGroupPath + orderRefundsPath).
		Init(test.ReqInitJSON()).
		BodyString(data).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrders_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPublic_DateValidationError() {
	bs := &billMock.BillingService{}
	bs.On("FindAllOrdersPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.ListOrdersPublicResponse{
				Status: billingpb.ResponseStatusOk,
				Item: &billingpb.ListOrdersPublicResponseItem{
					Count: 1,
					Items: []*billingpb.OrderViewPublic{},
				},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+orderPath).
		SetQueryParam("project_date_from", "a1b2c3").
		Init(test.ReqInitJSON()).
		Exec(suite.T())
	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), common.ErrorMessageListOrdersRequestPmDateFrom.Code, msg.Code)
	assert.Equal(suite.T(), common.ErrorMessageListOrdersRequestPmDateFrom.Message, msg.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BindError() {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+orderPath).
		SetQueryParam("limit", "abc").
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorRequestParamsIncorrect, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BillingServerError() {
	bs := &billMock.BillingService{}
	bs.On("FindAllOrdersPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BillingServer_ResultError() {
	bs := &billMock.BillingService{}
	bs.On("FindAllOrdersPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.ListOrdersPublicResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Code: "000", Message: "message"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "000", msg.Code)
	assert.Equal(suite.T(), "message", msg.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPublic_RoundError() {
	suite.router.moneyPrecision = -1

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BindError_Id() {
	q := make(url.Values)
	q.Set(common.RequestParameterId, bson.NewObjectId().Hex())
	suite.testGetOrdersBindError(q, fmt.Sprintf(common.ErrorMessageMask, "Id", "uuid"))
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BindError_Project() {
	q := url.Values{common.RequestParameterProject: []string{"foo"}}
	suite.testGetOrdersBindError(q, fmt.Sprintf(common.ErrorMessageMask, "Project[0]", "hexadecimal"))
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BindError_PaymentMethod() {
	q := url.Values{common.RequestParameterPaymentMethod: []string{"foo"}}
	suite.testGetOrdersBindError(q, fmt.Sprintf(common.ErrorMessageMask, "PaymentMethod[0]", "hexadecimal"))
}

func (suite *OrderTestSuite) TestOrder_GetOrders_BindError_Country() {
	q := url.Values{common.RequestParameterCountries: []string{"foo"}}
	suite.testGetOrdersBindError(q, fmt.Sprintf(common.ErrorMessageMask, "Country[0]", "len"))
}

func (suite *OrderTestSuite) testGetOrdersBindError(q url.Values, error string) {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParams(q).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.NewValidationError(error), httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ChangeOrderCode_Ok() {
	shouldBe := require.New(suite.T())

	changeOrderRequest := &billingpb.ChangeCodeInOrderRequest{
		KeyProductId: bson.NewObjectId().Hex(),
	}
	b, err := json.Marshal(changeOrderRequest)
	assert.NoError(suite.T(), err)

	billingService := &billMock.BillingService{}
	billingService.On("ChangeCodeInOrder", mock2.Anything, mock2.Anything).Return(&billingpb.ChangeCodeInOrderResponse{
		Status: billingpb.ResponseStatusOk,
		Order:  &billingpb.Order{},
	}, nil)

	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodPut).
		Params(":order_id", bson.NewObjectId().Hex()).
		Path(common.SystemUserGroupPath + orderReplaceCodePath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	shouldBe.NoError(err)
	shouldBe.Nil(err)
	shouldBe.EqualValues(http.StatusOK, res.Code)
	shouldBe.NotEmpty(res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_ChangeOrderCode_ServiceError() {
	shouldBe := require.New(suite.T())

	changeOrderRequest := &billingpb.ChangeCodeInOrderRequest{
		KeyProductId: bson.NewObjectId().Hex(),
	}
	b, err := json.Marshal(changeOrderRequest)
	assert.NoError(suite.T(), err)

	billingService := &billMock.BillingService{}
	billingService.On("ChangeCodeInOrder", mock2.Anything, mock2.Anything).Return(nil, errors.New("some error"))

	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Params(":order_id", bson.NewObjectId().Hex()).
		Path(common.SystemUserGroupPath + orderReplaceCodePath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	shouldBe.NotNil(err)
	httpErr, ok := err.(*echo.HTTPError)
	shouldBe.True(ok)
	shouldBe.EqualValues(http.StatusInternalServerError, httpErr.Code)
}

func (suite *OrderTestSuite) TestOrder_ChangeOrderCode_ErrorInService() {
	shouldBe := require.New(suite.T())

	changeOrderRequest := &billingpb.ChangeCodeInOrderRequest{
		KeyProductId: bson.NewObjectId().Hex(),
	}
	b, err := json.Marshal(changeOrderRequest)
	assert.NoError(suite.T(), err)

	billingService := &billMock.BillingService{}
	billingService.On("ChangeCodeInOrder", mock2.Anything, mock2.Anything).Return(&billingpb.ChangeCodeInOrderResponse{
		Status:  400,
		Message: &billingpb.ResponseErrorMessage{Message: "Some error"},
	}, nil)

	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Params(":order_id", bson.NewObjectId().Hex()).
		Path(common.SystemUserGroupPath + orderReplaceCodePath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	shouldBe.NotNil(err)
	httpErr, ok := err.(*echo.HTTPError)
	shouldBe.True(ok)
	shouldBe.EqualValues(http.StatusBadRequest, httpErr.Code)
}

func (suite *OrderTestSuite) TestOrder_ChangeOrderCode_ValidationError() {
	shouldBe := require.New(suite.T())

	// Missing key product id
	changeOrderRequest := &billingpb.ChangeCodeInOrderRequest{
		KeyProductId: "",
	}
	b, err := json.Marshal(changeOrderRequest)
	assert.NoError(suite.T(), err)

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Params(":order_id", bson.NewObjectId().Hex()).
		Path(common.SystemUserGroupPath + orderReplaceCodePath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	shouldBe.NotNil(err)
	httpErr, ok := err.(*echo.HTTPError)
	shouldBe.True(ok)
	shouldBe.EqualValues(http.StatusBadRequest, httpErr.Code)

	// Wrong order id
	changeOrderRequest = &billingpb.ChangeCodeInOrderRequest{
		KeyProductId: bson.NewObjectId().Hex(),
	}
	b, err = json.Marshal(changeOrderRequest)
	assert.NoError(suite.T(), err)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.AuthUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_InvalidOrderId_Error() {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "1234567890").
		Path(common.AuthUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusBadRequest, httpErr.Code)
	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), message.Message, common.ErrorValidationFailed.Message)
	assert.Regexp(suite.T(), "OrderId", message.Details)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_GetOrderPublic_BadResult_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.GetOrderPublicResponse{
				Status:  billingpb.ResponseStatusNotFound,
				Message: &billingpb.ResponseErrorMessage{Code: "000", Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.AuthUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusNotFound, httpErr.Code)
	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", message.Message)
	assert.Equal(suite.T(), "000", message.Code)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_GetOrderPublic_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.AuthUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrderLogs_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderGetLogsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())

	logs := new(GetOrderLogsResponse)
	err = json.Unmarshal(res.Body.Bytes(), &logs)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), logs.Create)
	assert.NotEmpty(suite.T(), logs.Create)
	assert.NotNil(suite.T(), logs.Callback)
	assert.NotEmpty(suite.T(), logs.Callback)
	assert.NotNil(suite.T(), logs.Notify)
	assert.NotEmpty(suite.T(), logs.Notify)
}

func (suite *OrderTestSuite) TestOrder_GetOrderLogs_GetOrderPublic_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.GetOrderPublicResponse{
				Status:  billingpb.ResponseStatusNotFound,
				Message: &billingpb.ResponseErrorMessage{Code: "000", Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderGetLogsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusNotFound, httpErr.Code)
	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", message.Message)
	assert.Equal(suite.T(), "000", message.Code)
}

func (suite *OrderTestSuite) TestOrder_GetOrderLogs_GetOrderPublic_Error1() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderGetLogsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrderLogs_PTypesTimestamp_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPublic", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.GetOrderPublicResponse{
				Status: billingpb.ResponseStatusOk,
				Item:   &billingpb.OrderViewPublic{},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderGetLogsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_GetOrderLogs_CloudWatchLog_Error() {
	cloudwatchMock := &mock.CloudWatchInterface{}
	cloudwatchMock.On("FilterLogEventsWithContext", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.cloudWatch.instance = cloudwatchMock

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderGetLogsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())

	logs := new(GetOrderLogsResponse)
	err = json.Unmarshal(res.Body.Bytes(), &logs)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), logs.Create)
	assert.Nil(suite.T(), logs.Callback)
	assert.Nil(suite.T(), logs.Notify)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPrivate_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_GetOrderPrivate_InvalidOrderId_Error() {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "1234567890").
		Path(common.SystemUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusBadRequest, httpErr.Code)
	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), message.Message, common.ErrorValidationFailed.Message)
	assert.Regexp(suite.T(), "OrderId", message.Details)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_GetOrderPrivate_BadResult_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPrivate", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.GetOrderPrivateResponse{
				Status:  billingpb.ResponseStatusNotFound,
				Message: &billingpb.ResponseErrorMessage{Code: "000", Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusNotFound, httpErr.Code)
	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", message.Message)
	assert.Equal(suite.T(), "000", message.Code)
}

func (suite *OrderTestSuite) TestOrder_GetOrderPublic_GetOrderPrivate_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("GetOrderPrivate", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":order_id", "ace2fc5c-b8c2-4424-96e8-5b631a73b88a").
		Path(common.SystemUserGroupPath + orderIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.EqualValues(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivate_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivate_ValidationError() {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+orderPath).
		SetQueryParam("project_date_from", "a1b2c3").
		Init(test.ReqInitJSON()).
		Exec(suite.T())
	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), common.ErrorMessageListOrdersRequestPmDateFrom.Code, msg.Code)
	assert.Equal(suite.T(), common.ErrorMessageListOrdersRequestPmDateFrom.Message, msg.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivate_BindError() {
	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+orderPath).
		SetQueryParam("limit", "abc").
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorRequestParamsIncorrect, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivate_BillingServerError() {
	bs := &billMock.BillingService{}
	bs.On("FindAllOrdersPrivate", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivate_BillingServiceResultError() {
	bs := &billMock.BillingService{}
	bs.On("FindAllOrdersPrivate", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.ListOrdersPrivateResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Code: "000", Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())
	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "000", msg.Code)
	assert.Equal(suite.T(), "some error", msg.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersPrivateRoundError() {
	suite.router.moneyPrecision = -1

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersS2s_Ok() {
	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.MerchantS2SGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *OrderTestSuite) TestOrder_ListOrdersS2s_ListOrders_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("FindAllOrders", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(nil, errors.New("TestOrder_ListOrdersS2s_ListOrders_Error"))
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.MerchantS2SGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())
	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorInternal, httpErr.Message)
}

func (suite *OrderTestSuite) TestOrder_ListOrdersS2s_Billing_FindAllOrders_Result_Error() {
	billingMock := &billMock.BillingService{}
	billingMock.On("FindAllOrders", mock2.Anything, mock2.Anything, mock2.Anything).
		Return(
			&billingpb.ListOrdersResponse{
				Status: billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{
					Code:    "000",
					Message: "TestOrder_ListOrdersS2s_Billing_FindAllOrders_Result_Error",
				},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = billingMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.MerchantS2SGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())
	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	message, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "000", message.Code)
	assert.Equal(suite.T(), "TestOrder_ListOrdersS2s_Billing_FindAllOrders_Result_Error", message.Message)
}
