package handlers

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/url"
	"testing"
)

type SubscriptionsTestSuite struct {
	suite.Suite
	router *SubscriptionsRoute
	user   *common.AuthUser
	caller *test.EchoReqResCaller
}

func Test_Subscriptions(t *testing.T) {
	suite.Run(t, new(SubscriptionsTestSuite))
}

func (suite *SubscriptionsTestSuite) TearDownTest() {}

func (suite *SubscriptionsTestSuite) SetupTest() {
	suite.user = &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}

	var e error
	settings := test.DefaultSettings()
	srv := common.Services{}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(suite.user))
		suite.router = NewSubscriptionsRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptions_Ok() {
	service := &billingMocks.BillingService{}
	service.On("FindSubscriptions", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.FindSubscriptionsResponse{List: []*billingpb.RecurringSubscription{}, Status: 200}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + subscriptionsPath).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptions_Validation_Error() {
	service := &billingMocks.BillingService{}
	service.On("FindSubscriptions", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.FindSubscriptionsResponse{List: []*billingpb.RecurringSubscription{}, Status: 200}, nil)
	suite.router.dispatch.Services.Billing = service

	q := make(url.Values)
	q.Set("limit", "a")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + subscriptionsPath).
		SetQueryParams(q).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptions_BillingFindSubscriptions_Error() {
	service := &billingMocks.BillingService{}
	service.On("FindSubscriptions", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.FindSubscriptionsResponse{}, errors.New("some error"))
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + subscriptionsPath).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 500, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptions_BillingFindSubscriptions_ErrorStatus() {
	service := &billingMocks.BillingService{}
	service.On("FindSubscriptions", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.FindSubscriptionsResponse{Status: billingpb.ResponseStatusBadData, Message: &billingpb.ResponseErrorMessage{Message: "some message"}}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + subscriptionsPath).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscription_Ok() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionResponse{Subscription: &billingpb.RecurringSubscription{}, Status: 200}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscription_ServiceError() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionResponse{Message: &billingpb.ResponseErrorMessage{Message: "some message"}, Status: 400}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscription_Error() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionResponse{}, errors.New("some error"))
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 500, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscription_ForbiddenError() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionResponse{Subscription: nil, Status: 403, Message: &billingpb.ResponseErrorMessage{Message: "some message"}}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 403, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptionOrders_Ok() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscriptionOrders", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionOrdersResponse{List: []*billingpb.SubscriptionOrder{}, Count: 0, Status: 200}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionOrdersPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptionOrders_ServiceError() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscriptionOrders", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionOrdersResponse{Message: &billingpb.ResponseErrorMessage{Message: "some message"}, Status: 400}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionOrdersPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptionOrders_Error() {
	service := &billingMocks.BillingService{}
	service.On("GetSubscriptionOrders", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetSubscriptionOrdersResponse{}, errors.New("some error"))
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionOrdersPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 500, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_GetSubscriptionOrders_ValidationError() {
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	q := make(url.Values)
	q.Set("limit", "a")

	_, err = suite.caller.Builder().
		Path(common.AuthUserGroupPath + subscriptionOrdersPath).
		SetQueryParams(q).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_DeleteSubscription_Ok() {
	id := bson.NewObjectId().Hex()

	service := &billingMocks.BillingService{}
	service.On("DeleteRecurringSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.EmptyResponseWithStatus{Status: 200}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, id).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *SubscriptionsTestSuite) TestCustomer_DeleteSubscription_ServiceError() {
	service := &billingMocks.BillingService{}
	service.On("DeleteRecurringSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.EmptyResponseWithStatus{Message: &billingpb.ResponseErrorMessage{Message: "some message"}, Status: 400}, nil)
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *SubscriptionsTestSuite) TestCustomer_DeleteSubscription_Error() {
	service := &billingMocks.BillingService{}
	service.On("DeleteRecurringSubscription", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.EmptyResponseWithStatus{}, errors.New("some error"))
	suite.router.dispatch.Services.Billing = service

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+subscriptionDetailsPath).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 500, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}
