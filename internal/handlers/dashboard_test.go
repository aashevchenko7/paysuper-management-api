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

type DashboardTestSuite struct {
	suite.Suite
	router *DashboardRoute
	caller *test.EchoReqResCaller
}

func Test_Dashboard(t *testing.T) {
	suite.Run(t, new(DashboardTestSuite))
}

func (suite *DashboardTestSuite) SetupTest() {
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}
	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardMainReport", mock.Anything, mock.Anything, mock.Anything).
		Return(&billingpb.GetDashboardMainResponse{Status: billingpb.ResponseStatusOk, Item: &billingpb.DashboardMainReport{}}, nil)
	bs.On("GetDashboardCustomersReport", mock.Anything, mock.Anything).
		Return(&billingpb.GetDashboardCustomerReportResponse{Status: billingpb.ResponseStatusOk, Item: &billingpb.DashboardCustomerReport{}}, nil)
	bs.On("GetDashboardRevenueDynamicsReport", mock.Anything, mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardRevenueDynamicsReportResponse{
				Status: billingpb.ResponseStatusOk,
				Item:   &billingpb.DashboardRevenueDynamicReport{},
			},
			nil,
		)
	bs.On("GetDashboardBaseReport", mock.Anything, mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardBaseReportResponse{
				Status: billingpb.ResponseStatusOk,
				Item:   &billingpb.DashboardBaseReports{},
			},
			nil,
		)

	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: bs,
	}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewDashboardRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *DashboardTestSuite) TearDownTest() {}

func (suite *DashboardTestSuite) TestDashboard_GetMainReports_Ok() {
	q := make(url.Values)
	q.Set("period", "current_month")
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardMainPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *DashboardTestSuite) TestDashboard_GetMainReports_ValidationError() {
	q := make(url.Values)
	q.Set("period", "123")
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardMainPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorIncorrectPeriod, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetMainReports_BillingServerSystemError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardMainReport", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardMainPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorUnknown, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetMainReports_BillingServerReturnError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardMainReport", mock.Anything, mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardMainResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardMainPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", msg.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetRevenueDynamicsReport_Ok() {
	q := make(url.Values)
	q.Set("period", "current_month")
	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardRevenueDynamicsPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *DashboardTestSuite) TestDashboard_GetRevenueDynamicsReport_ValidationError() {
	q := make(url.Values)
	q.Set("period", "123")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardRevenueDynamicsPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorIncorrectPeriod, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetRevenueDynamicsReport_BillingServerSystemError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardRevenueDynamicsReport", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardRevenueDynamicsPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorUnknown, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetRevenueDynamicsReport_BillingServerReturnError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardRevenueDynamicsReport", mock.Anything, mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardRevenueDynamicsReportResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardRevenueDynamicsPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", msg.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetBaseReports_Ok() {
	q := make(url.Values)
	q.Set("period", "current_month")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardBasePath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *DashboardTestSuite) TestDashboard_GetBaseReports_ValidationError() {
	q := make(url.Values)
	q.Set("period", "123")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardBasePath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorIncorrectPeriod, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetBaseReports_BillingServerSystemError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardBaseReport", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardBasePath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorUnknown, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetBaseReports_BillingServerReturnError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardBaseReport", mock.Anything, mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardBaseReportResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + dashboardBasePath).
		SetQueryParams(q).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", msg.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetCustomersReports_Ok() {
	q := make(url.Values)
	q.Set("period", "current_month")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardCustomersPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *DashboardTestSuite) TestDashboard_GetCustomersReports_ValidationError() {
	q := make(url.Values)
	q.Set("period", "123")

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardCustomersPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.NotEmpty(suite.T(), httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetCustomersReports_BillingServerSystemError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardCustomersReport", mock.Anything, mock.Anything).
		Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath+dashboardCustomersPath).
		SetQueryParams(q).
		Params(":"+common.RequestParameterId, bson.NewObjectId().Hex()).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorUnknown, httpErr.Message)
}

func (suite *DashboardTestSuite) TestDashboard_GetCustomersReports_BillingServerReturnError() {
	q := make(url.Values)
	q.Set("period", "current_month")

	bs := &billingMocks.BillingService{}
	bs.On("GetDashboardCustomersReport", mock.Anything, mock.Anything).
		Return(
			&billingpb.GetDashboardCustomerReportResponse{
				Status:  billingpb.ResponseStatusBadData,
				Message: &billingpb.ResponseErrorMessage{Message: "some error"},
			},
			nil,
		)
	suite.router.dispatch.Services.Billing = bs

	_, err := suite.caller.Builder().
		Path(common.AuthUserGroupPath + dashboardCustomersPath).
		SetQueryParams(q).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "some error", msg.Message)
}
