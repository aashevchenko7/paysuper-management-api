package handlers

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	billMock "github.com/paysuper/paysuper-proto/go/billingpb/mocks"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/mock"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type KeyProductTestSuite struct {
	suite.Suite
	router *KeyProductRoute
	caller *test.EchoReqResCaller
}

func Test_keyProduct(t *testing.T) {
	suite.Run(t, new(KeyProductTestSuite))
}

func (suite *KeyProductTestSuite) BeforeTest(suiteName, testName string) {
	switch testName {
	case "TestProject_CreateKeyProduct_GroupPrice_Error":
		suite.SetupTestForTestProject_CreateKeyProduct_GroupPrice_Error()
	case "TestProject_CreateKeyProduct_GroupPrice_Ok":
		suite.SetupTestForTestProject_CreateKeyProduct_GroupPrice_Ok()
	}
}

func (suite *KeyProductTestSuite) SetupTestForTestProject_CreateKeyProduct_GroupPrice_Ok() {
	billingService := &billMock.BillingService{}

	billingService.On("GetMerchantBy", mock2.Anything, mock2.Anything).Return(&billingpb.GetMerchantResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.Merchant{
			Id: bson.NewObjectId().Hex(),
		},
	}, nil)

	billingService.On("GetPriceGroupByRegion", mock2.Anything, mock2.Anything).Return(&billingpb.GetPriceGroupByRegionResponse{Status: 200, Group: &billingpb.PriceGroup{Id: "Some id"}}, nil)

	billingService.On("CreateOrUpdateKeyProduct", mock2.Anything, mock2.Anything).Return(&billingpb.KeyProductResponse{Status: 200, Product: &billingpb.KeyProduct{
		Id: "some_id",
	}}, nil)

	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: billingService,
		Geo:     mock.NewGeoIpServiceTestOk(),
	}
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}

	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewKeyProductRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *KeyProductTestSuite) SetupTestForTestProject_CreateKeyProduct_GroupPrice_Error() {
	billingService := &billMock.BillingService{}

	billingService.On("GetMerchantBy", mock2.Anything, mock2.Anything).Return(&billingpb.GetMerchantResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.Merchant{
			Id: bson.NewObjectId().Hex(),
		},
	}, nil)

	billingService.On("GetPriceGroupByRegion", mock2.Anything, mock2.Anything).Return(&billingpb.GetPriceGroupByRegionResponse{Status: 400, Group: nil, Message: &billingpb.ResponseErrorMessage{Message: "some error"}}, nil)
	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}
	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: billingService,
		Geo:     mock.NewGeoIpServiceTestOk(),
	}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewKeyProductRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *KeyProductTestSuite) SetupTest() {
	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: mock.NewBillingServerOkMock(),
		Geo:     mock.NewGeoIpServiceTestOk(),
	}

	user := &common.AuthUser{
		Id:         "ffffffffffffffffffffffff",
		MerchantId: "ffffffffffffffffffffffff",
		Role:       "owner",
	}

	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		mw.Pre(test.PreAuthUserMiddleware(user))
		suite.router = NewKeyProductRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *KeyProductTestSuite) TearDownTest() {}

func (suite *KeyProductTestSuite) TestProject_PublishKeyProduct_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + keyProductsPublishPath).
		Init(test.ReqInitJSON()).
		BodyString("{}").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_GetPlatformList_Error() {

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("limit", "qwe").
		Path(common.AuthUserGroupPath + platformsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
}

func (suite *KeyProductTestSuite) TestUnpublishKeyProduct_InternalError() {
	billingService := &billMock.BillingService{}
	billingService.On("UnPublishKeyProduct", mock2.Anything, mock2.Anything).Return(nil, errors.New("some error"))
	suite.router.dispatch.Services.Billing = billingService

	body := &billingpb.UnPublishKeyProductRequest{}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)
	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + keyProductsUnPublishPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	assert.EqualValues(suite.T(), 500, err.(*echo.HTTPError).Code)
}

func (suite *KeyProductTestSuite) TestUnpublishKeyProduct_Error() {
	billingService := &billMock.BillingService{}
	billingService.On("UnPublishKeyProduct", mock2.Anything, mock2.Anything).Return(&billingpb.KeyProductResponse{
		Status: billingpb.ResponseStatusBadData,
		Message: &billingpb.ResponseErrorMessage{
			Code:    "Some code",
			Message: "Some error",
		},
	}, nil)
	suite.router.dispatch.Services.Billing = billingService

	body := &billingpb.UnPublishKeyProductRequest{}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)
	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + keyProductsUnPublishPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	assert.EqualValues(suite.T(), 400, err.(*echo.HTTPError).Code)
}

func (suite *KeyProductTestSuite) TestUnpublishKeyProduct_Ok() {
	billingService := &billMock.BillingService{}
	billingService.On("UnPublishKeyProduct", mock2.Anything, mock2.Anything).Return(&billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusOk,
		Product: &billingpb.KeyProduct{},
	}, nil)
	suite.router.dispatch.Services.Billing = billingService

	body := &billingpb.UnPublishKeyProductRequest{}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)
	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + keyProductsUnPublishPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *KeyProductTestSuite) TestProject_GetPlatformList_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("limit", "10").
		Path(common.AuthUserGroupPath + platformsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())

	res, err = suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("limit", "").
		Path(common.AuthUserGroupPath + platformsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())

	res, err = suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("limit", "300000").
		Path(common.AuthUserGroupPath + platformsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_GetListKeyProduct_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthUserGroupPath + keyProductsPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_GetKeyProduct_ValidationError() {

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":key_product_id", " ").
		Path(common.AuthUserGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *KeyProductTestSuite) TestProject_GetKeyProduct_Ok() {

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthUserGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_CreateKeyProduct_Ok() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		MerchantId:      bson.NewObjectId().Hex(),
		Name:            map[string]string{"en": "A", "ru": "А"},
		Description:     map[string]string{"en": "A", "ru": "А"},
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Pricing:         "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Params(":key_product_id", body.Id).
		Path(common.AuthUserGroupPath + keyProductsPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_ChangeKeyProduct_ValidationError() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		Id:              bson.NewObjectId().Hex(),
		MerchantId:      bson.NewObjectId().Hex(),
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Pricing:         "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	_, err = suite.caller.Builder().
		Method(http.MethodPut).
		Params(":key_product_id", body.Id).
		Path(common.AuthUserGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *KeyProductTestSuite) TestProject_ChangeKeyProduct_Ok() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		Id:              bson.NewObjectId().Hex(),
		MerchantId:      bson.NewObjectId().Hex(),
		Name:            map[string]string{"en": "A", "ru": "А"},
		Description:     map[string]string{"en": "A", "ru": "А"},
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Pricing:         "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	res, err := suite.caller.Builder().
		Method(http.MethodPut).
		Params(":key_product_id", body.Id).
		Path(common.AuthUserGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_CreateKeyProduct_GroupPrice_Ok() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		Name:            map[string]string{"en": "A", "ru": "А"},
		MerchantId:      bson.NewObjectId().Hex(),
		Description:     map[string]string{"en": "A", "ru": "А"},
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Platforms: []*billingpb.PlatformPrice{
			{
				Id:   "gog",
				Name: "Gog",
				Prices: []*billingpb.ProductPrice{
					{
						Currency: "RUB",
						Region:   "RUB",
						Amount:   666,
					},
				}},
		},
		Pricing: "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath + keyProductsPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *KeyProductTestSuite) TestProject_CreateKeyProduct_GroupPrice_Error() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		Name:            map[string]string{"en": "A", "ru": "А"},
		MerchantId:      bson.NewObjectId().Hex(),
		Description:     map[string]string{"en": "A", "ru": "А"},
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Platforms: []*billingpb.PlatformPrice{
			{
				Id:   "gog",
				Name: "Gog",
				Prices: []*billingpb.ProductPrice{
					{
						Currency: "RUB",
						Region:   "TestRegion",
						Amount:   666,
					},
				}},
		},
		Pricing: "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath + keyProductsPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *KeyProductTestSuite) TestProject_CreateKeyProduct_ValidationError() {
	body := &billingpb.CreateOrUpdateKeyProductRequest{
		MerchantId:      bson.NewObjectId().Hex(),
		Description:     map[string]string{"en": "A", "ru": "А"},
		DefaultCurrency: "RUB",
		ProjectId:       bson.NewObjectId().Hex(),
		Sku:             "some_sku",
		Object:          "key_product",
		Pricing:         "manual",
	}

	b, err := json.Marshal(&body)
	assert.NoError(suite.T(), err)

	_, err = suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath + keyProductsPath).
		Init(test.ReqInitJSON()).
		BodyBytes(b).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	e, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 400, e.Code)
	assert.NotEmpty(suite.T(), e.Message)
}

func (suite *KeyProductTestSuite) TestProject_getKeyProduct_ValidationError() {

	billingService := &billMock.BillingService{}
	billingService.On("GetKeyProduct", mock2.Anything, mock2.Anything).Return(&billingpb.KeyProductResponse{}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.AuthProjectGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	err2, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, err2.Code)
	assert.NotEmpty(suite.T(), err2.Message)
}

func (suite *KeyProductTestSuite) TestProject_getKeyProduct_BillingServer() {

	billingService := &billMock.BillingService{}
	billingService.On("GetKeyProductInfo", mock2.Anything, mock2.Anything).Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthProjectGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.Error(suite.T(), err)
	err2, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, err2.Code)
	assert.NotEmpty(suite.T(), err2.Message)
}

func (suite *KeyProductTestSuite) TestProject_getKeyProductWithCurrency_Ok() {

	billingService := &billMock.BillingService{}
	billingService.On("GetKeyProductInfo", mock2.Anything, mock2.Anything).Return(&billingpb.GetKeyProductInfoResponse{
		Status: 200,
		KeyProduct: &billingpb.KeyProductInfo{
			LongDescription: "Description",
			Name:            "Name",
			Platforms: []*billingpb.PlatformPriceInfo{
				{Name: "Steam", Id: "steam", Price: &billingpb.ProductPriceInfo{Currency: "USD", Amount: 10, Region: "USD"}},
			},
		},
	}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("currency", "USD").
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthProjectGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_getKeyProductWithCountry_Ok() {

	billingService := &billMock.BillingService{}
	billingService.On("GetKeyProductInfo", mock2.Anything, mock2.Anything).Return(&billingpb.GetKeyProductInfoResponse{
		Status: 200,
		KeyProduct: &billingpb.KeyProductInfo{
			LongDescription: "Description",
			Name:            "Name",
			Platforms: []*billingpb.PlatformPriceInfo{
				{Name: "Steam", Id: "steam", Price: &billingpb.ProductPriceInfo{Currency: "USD", Amount: 10, Region: "USD"}},
			},
		},
	}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		SetQueryParam("country", "RUS").
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthProjectGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}

func (suite *KeyProductTestSuite) TestProject_getKeyProduct_Ok() {

	billingService := &billMock.BillingService{}
	billingService.On("GetKeyProductInfo", mock2.Anything, mock2.Anything).Return(&billingpb.GetKeyProductInfoResponse{
		Status: 200,
		KeyProduct: &billingpb.KeyProductInfo{
			LongDescription: "Description",
			Name:            "Name",
			Platforms: []*billingpb.PlatformPriceInfo{
				{Name: "Steam", Id: "steam", Price: &billingpb.ProductPriceInfo{Currency: "USD", Amount: 10, Region: "USD"}},
			},
		},
	}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Params(":key_product_id", bson.NewObjectId().Hex()).
		Path(common.AuthProjectGroupPath + keyProductsIdPath).
		Init(test.ReqInitJSON()).
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEmpty(suite.T(), res.Body.String())
}
