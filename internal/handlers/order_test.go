package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewOrderRoute(set.HandlerSet, set.GlobalConfig)
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

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + orderPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
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
