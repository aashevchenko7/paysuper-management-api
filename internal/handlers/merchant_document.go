package handlers

import (
	"context"
	"fmt"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
	awsWrapper "github.com/paysuper/paysuper-aws-manager"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	merchantDocumentsPath          = "/documents"
	merchantDocumentPath           = "/documents/:id"
	merchantDocumentDownloadPath   = "/documents/:id/download"
	merchantIdDocumentsPath        = "/merchants/:merchant_id/documents"
	merchantIdDocumentPath         = "/merchants/:merchant_id/documents/:id"
	merchantIdDocumentDownloadPath = "/merchants/:merchant_id/documents/:id/download"
)

var (
	availableContentTypes = map[string]bool{
		"application/msword": true,
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/vnd.ms-excel": true,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true,
		"application/rtf":             true,
		"text/plain":                  true,
		"image/bmp":                   true,
		"image/jpeg":                  true,
		"image/png":                   true,
		"image/gif":                   true,
		"image/tiff":                  true,
		"application/vnd.rar":         true,
		"application/zip":             true,
		"application/x-7z-compressed": true,
		"application/pdf":             true,
	}
	fileExtensions = map[string]string{
		"application/msword": "doc",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
		"application/vnd.ms-excel": "xls",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": "xlsx",
		"application/rtf":             "rtf",
		"text/plain":                  "txt",
		"image/bmp":                   "bmp",
		"image/jpeg":                  "jpg",
		"image/png":                   "png",
		"image/gif":                   "gif",
		"image/tiff":                  "tiff",
		"application/vnd.rar":         "rar",
		"application/zip":             "zip",
		"application/x-7z-compressed": "7z",
		"application/pdf":             "pdf",
	}
	merchantDocumentMaxSize = int64(31457280) //30MB
)

type MerchantDocumentRoute struct {
	dispatch   common.HandlerSet
	cfg        common.Config
	awsManager awsWrapper.AwsManagerInterface
	provider.LMT
}

func NewMerchantDocumentRoute(set common.HandlerSet, awsManager awsWrapper.AwsManagerInterface, cfg *common.Config) *MerchantDocumentRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "MerchantDocumentRoute"})
	return &MerchantDocumentRoute{
		dispatch:   set,
		LMT:        &set.AwareSet,
		cfg:        *cfg,
		awsManager: awsManager,
	}
}

func (h *MerchantDocumentRoute) Route(groups *common.Groups) {
	groups.AuthUser.POST(merchantDocumentsPath, h.uploadDocument)
	groups.SystemUser.POST(merchantIdDocumentsPath, h.uploadDocument)
	groups.AuthUser.GET(merchantDocumentsPath, h.listDocuments)
	groups.SystemUser.GET(merchantIdDocumentsPath, h.listDocuments)
	groups.AuthUser.GET(merchantDocumentPath, h.getDocument)
	groups.SystemUser.GET(merchantIdDocumentPath, h.getDocument)
	groups.AuthUser.GET(merchantDocumentDownloadPath, h.downloadDocument)
	groups.SystemUser.GET(merchantIdDocumentDownloadPath, h.downloadDocument)
}

// @summary Upload document for merchant
// @desc Upload document for merchant
// @id merchantDocumentUpload
// @tag Document
// @accept multipart/form-data
// @produce application/json
// @success 200 {object} billingpb.AddMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param file form {array} true Uploaded file byte array.
// @router /admin/api/v1/documents [post]

// @summary Upload document for merchant using the merchant ID
// @desc Upload document for merchant using the merchant ID
// @id merchantDocumentUploadAdmin
// @tag Document
// @accept multipart/form-data
// @produce application/json
// @success 200 {object} billingpb.AddMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param file form {array} true Uploaded file byte array.
// @router /system/api/v1/merchants/{merchant_id}/documents [post]
func (h *MerchantDocumentRoute) uploadDocument(ctx echo.Context) error {
	document := &billingpb.MerchantDocument{}
	err := ctx.Bind(document)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	document.Id = bson.NewObjectId().Hex()
	document.UserId = common.ExtractUserContext(ctx).Id

	file, err := h.validateUploadedFile(ctx, document)

	if err != nil {
		h.L().Error(
			"failed validate merchant uploaded document",
			logger.PairArgs("err", err.Error()),
		)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.dispatch.Validate.Struct(document); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	in := &awsWrapper.UploadInput{
		Body:     file,
		FileName: document.FilePath,
	}
	ctxUpload, _ := context.WithTimeout(context.Background(), time.Second*10)
	_, err = h.awsManager.Upload(ctxUpload, in)

	if err != nil {
		h.L().Error(
			"unable to upload merchant document into S3",
			logger.PairArgs("err", err.Error()),
		)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessageMerchantDocumentUploadFailed)
	}

	res, err := h.dispatch.Services.Billing.AddMerchantDocument(ctx.Request().Context(), document)

	if err != nil {
		return h.dispatch.SrvCallHandler(document, err, billingpb.ServiceName, "AddMerchantDocument")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *MerchantDocumentRoute) validateUploadedFile(ctx echo.Context, document *billingpb.MerchantDocument) (multipart.File, error) {
	ctx.Request().ParseMultipartForm(10 << 30)
	file, handler, err := ctx.Request().FormFile("file")

	if err != nil {
		h.L().Error(
			"unable to find file in merchant document upload request",
			logger.PairArgs("err", err.Error()),
		)
		return nil, common.ErrorMessageMerchantDocumentNotFoundInRequest
	}

	defer file.Close()

	if handler.Size > merchantDocumentMaxSize {
		return nil, common.ErrorMessageMerchantDocumentInvalidSize
	}

	if _, ok := availableContentTypes[handler.Header.Get("Content-Type")]; !ok {
		return nil, common.ErrorMessageMerchantDocumentInvalidType
	}

	ext, ok := fileExtensions[handler.Header.Get("Content-Type")]

	if !ok {
		return nil, common.ErrorMessageMerchantDocumentInvalidType
	}

	document.OriginalName = handler.Filename
	document.FilePath = fmt.Sprintf("%s_%s.%s", document.MerchantId, document.Id, ext)

	return file, nil
}

// @summary List documents for merchant
// @desc List documents for merchant
// @id merchantDocumentsList
// @tag Document
// @accept application/json
// @produce application/json
// @body billingpb.GetMerchantDocumentsRequest
// @success 200 {object} billingpb.GetMerchantDocumentsResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /admin/api/v1/documents [get]

// @summary List documents of merchant using the merchant ID
// @desc List documents of merchant using the merchant ID
// @id merchantDocumentsListAdmin
// @tag Document
// @accept application/json
// @produce application/json
// @body billingpb.GetMerchantDocumentsRequest
// @success 200 {object} billingpb.GetMerchantDocumentsResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @router /system/api/v1/merchants/{merchant_id}/documents [get]
func (h *MerchantDocumentRoute) listDocuments(ctx echo.Context) error {
	req := &billingpb.GetMerchantDocumentsRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantDocuments(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMerchantDocuments", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get document for merchant
// @desc Get document for merchant
// @id merchantDocumentView
// @tag Document
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique document identifier.
// @router /admin/api/v1/document/{id} [get]

// @summary Get document of merchant using the merchant ID
// @desc Get document of merchant using the merchant ID
// @id merchantDocumentViewAdmin
// @tag Document
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param id path {string} true The unique document identifier.
// @router /system/api/v1/merchants/{merchant_id}/documents/{id} [get]
func (h *MerchantDocumentRoute) getDocument(ctx echo.Context) error {
	req := &billingpb.GetMerchantDocumentRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantDocument(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMerchantDocument", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Download document for merchant
// @desc Download document for merchant
// @id merchantDocumentDownload
// @tag Document
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param id path {string} true The unique document identifier.
// @router /admin/api/v1/document/{id}/download [get]

// @summary Download document of merchant using the merchant ID
// @desc Download document of merchant using the merchant ID
// @id merchantDocumentDownloadAdmin
// @tag Document
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetMerchantDocumentResponse Returns the uploaded document data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param id path {string} true The unique document identifier.
// @router /system/api/v1/merchants/{merchant_id}/documents/{id}/download [get]
func (h *MerchantDocumentRoute) downloadDocument(ctx echo.Context) error {
	req := &billingpb.GetMerchantDocumentRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantDocument(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetMerchantDocument", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	filePath := os.TempDir() + string(os.PathSeparator) + res.Item.Id
	_, err = h.awsManager.Download(ctx.Request().Context(), filePath, &awsWrapper.DownloadInput{FileName: res.Item.FilePath})

	if err != nil {
		h.L().Error(
			"unable to download the merchant document",
			logger.PairArgs("err", err),
			logger.PairArgs("document", res.Item),
		)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessageMerchantDocumentDownloadFailed)
	}

	defer os.Remove(filePath)

	return ctx.Inline(filePath, res.Item.OriginalName)
}
