package handlers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CustomerTestSuite struct {
	suite.Suite
	router *CustomerRoute
	caller *test.EchoReqResCaller
}

func Test_Customers(t *testing.T) {
	suite.Run(t, new(CustomerTestSuite))
}

func (suite *CustomerTestSuite) TearDownTest() {}

func (suite *CustomerTestSuite) SetupTest() {
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}
	bs := &billingMocks.BillingService{}
	bs.On("GetCustomerList", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.ListCustomersResponse{Status: billingpb.ResponseStatusOk, Items: []*billingpb.ShortCustomerInfo{
			&billingpb.ShortCustomerInfo{
				Id:         "",
				ExternalId: "",
				Country:    "",
				Language:   "",
				LastOrder:  nil,
				Orders:     0,
				Revenue:    0,
			},
	}}, nil)

	bs.On("GetCustomerInfo", mock.Anything, mock.Anything).
		Return(&billingpb.GetCustomerInfoResponse{Status: billingpb.ResponseStatusOk, Item: &billingpb.Customer{
		Id:                    "",
		TechEmail:             "",
		ExternalId:            "",
		Email:                 "",
		EmailVerified:         false,
		Phone:                 "",
		PhoneVerified:         false,
		Name:                  "",
		Ip:                    nil,
		Locale:                "",
		AcceptLanguage:        "",
		UserAgent:             "",
		Address:               nil,
		Identity:              nil,
		IpHistory:             nil,
		AddressHistory:        nil,
		LocaleHistory:         nil,
		AcceptLanguageHistory: nil,
		Metadata:              nil,
		CreatedAt:             nil,
		UpdatedAt:             nil,
		NotifySale:            false,
		NotifySaleEmail:       "",
		NotifyNewRegion:       false,
		NotifyNewRegionEmail:  "",
		IpString:              "",
		PaymentActivity:       nil,
		Uuid:                  "",
	}}, nil)

	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: bs,
	}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewCustomerRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *CustomerTestSuite) TestCustomer_GetDetails_Ok() {
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+customerDetailed).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *CustomerTestSuite) TestCustomer_GetDetails_Error() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetCustomerInfo", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetCustomerInfoResponse{Status: billingpb.ResponseStatusBadData, Item: nil, Message: &billingpb.ResponseErrorMessage{Message: "asd", Code: "123"}}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+customerDetailed).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *CustomerTestSuite) TestCustomer_GetListing_Ok() {
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+customerListing).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *CustomerTestSuite) TestCustomer_GetListing_Error() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetCustomerList", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.ListCustomersResponse{Status: billingpb.ResponseStatusBadData, Items: nil, Message: &billingpb.ResponseErrorMessage{Message: "asd", Code: "123"}}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+customerListing).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}