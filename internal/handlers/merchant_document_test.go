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
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

var (
	fixturesPath    = "../fixtures/"
	validFilePath   = fixturesPath + "merchant_document_valid.png"
	invalidFilePath = fixturesPath + "merchant_document_invalid.webp"
)

type MerchantDocumentTestSuite struct {
	suite.Suite
	router *MerchantDocumentRoute
	caller *test.EchoReqResCaller
}

func Test_MerchantDocument(t *testing.T) {
	suite.Run(t, new(MerchantDocumentTestSuite))
}

func (suite *MerchantDocumentTestSuite) SetupTest() {
	user := &common.AuthUser{
		Id:    "ffffffffffffffffffffffff",
		Email: "test@unit.test",
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

		suite.router = NewMerchantDocumentRoute(set.HandlerSet, awsManagerMock, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *MerchantDocumentTestSuite) TearDownTest() {}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("AddMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddMerchantDocumentResponse{Item: &billingpb.MerchantDocument{Id: "id"}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		ExecFileUpload(suite.T(), map[string]string{}, "file", validFilePath)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.AddMerchantDocumentResponse{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), obj.Item)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_UploadedFileType() {
	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		ExecFileUpload(suite.T(), map[string]string{}, "file", invalidFilePath)

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorMessageMerchantDocumentInvalidType, httpErr.Message)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_NoUploadedFile() {
	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
	assert.Equal(suite.T(), common.ErrorMessageMerchantDocumentNotFoundInRequest, httpErr.Message)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_InvalidMerchantId() {
	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "id").
		ExecFileUpload(suite.T(), map[string]string{}, "file", validFilePath)

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("AddMerchantDocument", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		ExecFileUpload(suite.T(), map[string]string{}, "file", validFilePath)

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_AddMerchantDocumentStatus() {
	billingService := &billingMocks.BillingService{}
	billingService.On("AddMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.AddMerchantDocumentResponse{Status: billingpb.ResponseStatusNotFound}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		ExecFileUpload(suite.T(), map[string]string{}, "file", validFilePath)

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusNotFound, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_UploadDocumentAdmin_Error_S3Upload() {
	awsManagerMock := &awsWrapperMocks.AwsManagerInterface{}
	awsManagerMock.On("Upload", mock2.Anything, mock2.Anything, mock2.Anything).Return(nil, errors.New("error"))
	suite.router.awsManager = awsManagerMock

	_, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Init(test.ReqInitMultipartForm()).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		ExecFileUpload(suite.T(), map[string]string{}, "file", validFilePath)

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DocumentListAdmin_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocuments", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentsResponse{List: []*billingpb.MerchantDocument{}, Count: 0, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.GetMerchantDocumentsResponse{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), []*billingpb.MerchantDocument{}, obj.List)
	assert.IsType(suite.T(), int64(0), obj.Count)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DocumentListAdmin_Error_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocuments", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DocumentListAdmin_Error_GetMerchantDocumentsResponse() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocuments", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentsResponse{Status: billingpb.ResponseStatusBadData}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentsPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_GetDocument_Ok() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentResponse{Item: &billingpb.MerchantDocument{}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)

	obj := &billingpb.GetMerchantDocumentResponse{}
	err = json.Unmarshal(res.Body.Bytes(), obj)
	assert.NoError(suite.T(), err)
	assert.IsType(suite.T(), &billingpb.MerchantDocument{}, obj.Item)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_GetDocument_Error_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_GetDocument_Error_GetMerchantDocumentsResponse() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentResponse{Status: billingpb.ResponseStatusBadData}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DownloadDocument_Ok() {
	document := &billingpb.MerchantDocument{Id: "id"}

	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentResponse{Item: document, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	awsManagerMock := &awsWrapperMocks.AwsManagerInterface{}
	awsManagerMock.On("Download", mock2.Anything, mock2.Anything, mock2.Anything).Return(int64(1), nil)
	suite.router.awsManager = awsManagerMock

	_, err := os.Create(os.TempDir() + string(os.PathSeparator) + document.Id)
	assert.NoError(suite.T(), err)

	res, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentDownloadPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DownloadDocument_Error_BillingError() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(nil, errors.New("error"))
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentDownloadPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DownloadDocument_Error_GetMerchantDocumentsResponse() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentResponse{Status: billingpb.ResponseStatusBadData}, nil)
	suite.router.dispatch.Services.Billing = billingService

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentDownloadPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.Code)
}

func (suite *MerchantDocumentTestSuite) TestMerchantDocument_DownloadDocument_Error_AwsDownload() {
	billingService := &billingMocks.BillingService{}
	billingService.On("GetMerchantDocument", mock2.Anything, mock2.Anything).
		Return(&billingpb.GetMerchantDocumentResponse{Item: &billingpb.MerchantDocument{}, Status: billingpb.ResponseStatusOk}, nil)
	suite.router.dispatch.Services.Billing = billingService

	awsManagerMock := &awsWrapperMocks.AwsManagerInterface{}
	awsManagerMock.On("Download", mock2.Anything, mock2.Anything, mock2.Anything).Return(int64(0), errors.New("error"))
	suite.router.awsManager = awsManagerMock

	_, err := suite.caller.Builder().
		Method(http.MethodGet).
		Path(common.SystemUserGroupPath+merchantIdDocumentDownloadPath).
		Params(":"+common.RequestParameterMerchantId, "ffffffffffffffffffffffff", ":"+common.RequestParameterId, "id").
		Exec(suite.T())

	assert.Error(suite.T(), err)

	httpErr, ok := err.(*echo.HTTPError)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), http.StatusInternalServerError, httpErr.Code)
}
