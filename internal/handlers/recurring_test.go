package handlers

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	awsWrapperMocks "github.com/paysuper/paysuper-aws-manager/pkg/mocks"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/mock"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RecurringTestSuite struct {
	suite.Suite
	router *RecurringRoute
	caller *test.EchoReqResCaller
}

func Test_Recurring(t *testing.T) {
	suite.Run(t, new(RecurringTestSuite))
}

func (suite *RecurringTestSuite) SetupTest() {
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Email:      "test@unit.test",
	}
	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: mock.NewBillingServerOkMock(),
	}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))

		awsManagerMock := &awsWrapperMocks.AwsManagerInterface{}
		awsManagerMock.On("Upload", mock2.Anything, mock2.Anything, mock2.Anything).Return(&s3manager.UploadOutput{}, nil)

		suite.router = NewRecurringRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *RecurringTestSuite) TearDownTest() {}

func (suite *RecurringTestSuite) TestAddRecurringPlan_Ok() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("AddRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddRecurringPlanResponse{Item: &billingpb.RecurringPlan{Id: "id"}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.RecurringPlan{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), obj)
}

func (suite *RecurringTestSuite) TestAddRecurringPlan_NameEmpty() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("AddRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddRecurringPlanResponse{Item: &billingpb.RecurringPlan{Id: "id"}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Regexp(suite.T(), "Name", msg.Details)
}

func (suite *RecurringTestSuite) TestAddRecurringPlan_ChargeEmpty() {
	plan := &billingpb.RecurringPlan{
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("AddRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddRecurringPlanResponse{Item: &billingpb.RecurringPlan{Id: "id"}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Regexp(suite.T(), "Charge", msg.Details)
}

func (suite *RecurringTestSuite) TestAddRecurringPlan_BillingResponseError() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("AddRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddRecurringPlanResponse{Status: billingpb.ResponseStatusBadData, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestAddRecurringPlan_BillingError() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("AddRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestUpdateRecurringPlan_Ok() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("UpdateRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.UpdateRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodPut).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNoContent, res.Code)
}

func (suite *RecurringTestSuite) TestUpdateRecurringPlan_NameEmpty() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("UpdateRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.UpdateRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Regexp(suite.T(), "Name", msg.Details)
}

func (suite *RecurringTestSuite) TestUpdateRecurringPlan_ChargeEmpty() {
	plan := &billingpb.RecurringPlan{
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("UpdateRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.UpdateRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)

	msg, ok := httpErr.Message.(*billingpb.ResponseErrorMessage)
	assert.True(suite.T(), ok)
	assert.Regexp(suite.T(), "Charge", msg.Details)
}

func (suite *RecurringTestSuite) TestUpdateRecurringPlan_BillingResponseError() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("UpdateRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.UpdateRecurringPlanResponse{Status: billingpb.ResponseStatusBadData, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestUpdateRecurringPlan_BillingError() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("UpdateRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestGetRecurringPlan_Ok() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}
	b, err := json.Marshal(plan)
	assert.NoError(suite.T(), err)

	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetRecurringPlanResponse{Item: plan, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.RecurringPlan{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), obj)
}

func (suite *RecurringTestSuite) TestGetRecurringPlan_NotFound() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetRecurringPlanResponse{Status: billingpb.ResponseStatusNotFound, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestGetRecurringPlan_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestEnableRecurringPlan_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("EnableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.EnableRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodPatch).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNoContent, res.Code)
}

func (suite *RecurringTestSuite) TestEnableRecurringPlan_NotFound() {
	billingService := &billingMocks.BillingService{}
	billingService.On("EnableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.EnableRecurringPlanResponse{Status: billingpb.ResponseStatusNotFound, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodPatch).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestEnableRecurringPlan_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("EnableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodPatch).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestDisableRecurringPlan_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DisableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.DisableRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNoContent, res.Code)
}

func (suite *RecurringTestSuite) TestDisableRecurringPlan_NotFound() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DisableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.DisableRecurringPlanResponse{Status: billingpb.ResponseStatusNotFound, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestDisableRecurringPlan_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DisableRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlan).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestDeleteRecurringPlan_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DeleteRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.DeleteRecurringPlanResponse{Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlanDelete).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNoContent, res.Code)
}

func (suite *RecurringTestSuite) TestDeleteRecurringPlan_NotFound() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DeleteRecurringPlan", mock2.Anything, mock2.Anything).
		Return(&billingpb.DeleteRecurringPlanResponse{Status: billingpb.ResponseStatusNotFound, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlanDelete).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestDeleteRecurringPlan_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("DeleteRecurringPlan", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodDelete).
		Path(common.AuthUserGroupPath+recurringPlanDelete).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff", ":"+common.RequestParameterPlanId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *RecurringTestSuite) TestGetRecurringPlans_Ok() {
	plan := &billingpb.RecurringPlan{
		Charge: &billingpb.RecurringPlanCharge{
			Period: &billingpb.RecurringPlanPeriod{
				Type:  recurringpb.RecurringPeriodDay,
				Value: 1,
			},
			Amount:   1,
			Currency: "RUB",
		},
		Name: map[string]string{"en": "en"},
	}

	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlans", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetRecurringPlansResponse{List: []*billingpb.RecurringPlan{plan}, Count: int32(1), Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.GetRecurringPlansResponse{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), obj.List)
	assert.NotEmpty(suite.T(), obj.Count)
}

func (suite *RecurringTestSuite) TestGetRecurringPlans_BillingResponseError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlans", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetRecurringPlansResponse{Status: billingpb.ResponseStatusBadData, Message: mock.SomeError}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), mock.SomeError, httpErr.Message)
}

func (suite *RecurringTestSuite) TestGetRecurringPlans_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetRecurringPlans", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath+recurringPlanList).
		Init(test.ReqInitJSON()).
		Params(":"+common.RequestParameterProjectId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}
