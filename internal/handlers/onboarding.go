package handlers

import (
	"fmt"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/client"
	awsWrapper "github.com/paysuper/paysuper-aws-manager"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	merchantsPath                      = "/merchants"
	merchantsIdPath                    = "/merchants/:merchant_id"
	merchantsCompanyPath               = "/merchants/company"
	merchantsContactsPath              = "/merchants/contacts"
	merchantsBankingPath               = "/merchants/banking"
	merchantsIdCompanyPath             = "/merchants/:merchant_id/company"
	merchantsIdContactsPath            = "/merchants/:merchant_id/contacts"
	merchantsIdBankingPath             = "/merchants/:merchant_id/banking"
	merchantsStatusCompanyPath         = "/merchants/status"
	merchantsIdChangeStatusCompanyPath = "/merchants/:merchant_id/change-status"
	merchantsNotificationsPath         = "/merchants/notifications"
	merchantsIdNotificationsPath       = "/merchants/:merchant_id/notifications"
	merchantsAgreementPath             = "/merchants/agreement"
	merchantsIdAgreementPath           = "/merchants/:merchant_id/agreement"
	merchantsAgreementDocumentPath     = "/merchants/agreement/document"
	merchantsIdAgreementDocumentPath   = "/merchants/:merchant_id/agreement/document"
	merchantsNotificationsIdPath       = "/merchants/notifications/:notification_id"
	merchantsNotificationsMarkReadPath = "/merchants/notifications/:notification_id/mark-as-read"
	merchantsTariffsPath               = "/merchants/tariffs"
	merchantsIdTariffsPath             = "/merchants/:merchant_id/tariffs"
	merchantsIdManualPayoutEnablePath  = "/merchants/manual_payout/enable"
	merchantsIdManualPayoutDisablePath = "/merchants/manual_payout/disable"
	merchantsIdSetOperatingCompanyPath = "/merchants/:merchant_id/set_operating_company"
)

const (
	agreementFileMask        = "agreement_%s.pdf"
	agreementContentType     = "application/pdf"
	agreementExtension       = "pdf"
	merchantAgreementUrlMask = "%s://%s/admin/api/v1/merchants/agreement/document"
	systemAgreementUrlMask   = "%s://%s/system/api/v1/merchants/%s/agreement/document"
	agreementUploadMaxSize   = 3145728
)

type OnboardingFileMetadata struct {
	// The agreement's file name.
	Name string `json:"name"`
	// The agreement's file extension.
	Extension string `json:"extension"`
	// The agreement's file content type.
	ContentType string `json:"content_type"`
	// The agreement's file size in KB.
	Size int64 `json:"size"`
}

type OnboardingFileData struct {
	// The URL for the printable agreement.
	Url string `json:"url"`
	// The metadata of the agreement file.
	Metadata *OnboardingFileMetadata `json:"metadata"`
}

type OnboardingRoute struct {
	dispatch   common.HandlerSet
	awsManager awsWrapper.AwsManagerInterface
	cfg        common.Config
	provider.LMT
}

func NewOnboardingRoute(set common.HandlerSet, initial config.Initial, awsManager awsWrapper.AwsManagerInterface, globalCfg *common.Config) *OnboardingRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "OnboardingRoute"})
	return &OnboardingRoute{
		dispatch:   set,
		LMT:        &set.AwareSet,
		cfg:        *globalCfg,
		awsManager: awsManager,
	}
}

func (h *OnboardingRoute) Route(groups *common.Groups) {
	groups.SystemUser.GET(merchantsPath, h.listMerchants)
	groups.SystemUser.GET(merchantsIdPath, h.getMerchant)

	groups.AuthUser.PUT(merchantsCompanyPath, h.setMerchantCompany)
	groups.AuthUser.PUT(merchantsContactsPath, h.setMerchantContacts)
	groups.AuthUser.PUT(merchantsBankingPath, h.setMerchantBanking)
	groups.SystemUser.PUT(merchantsIdCompanyPath, h.setMerchantCompany)
	groups.SystemUser.PUT(merchantsIdContactsPath, h.setMerchantContacts)
	groups.SystemUser.PUT(merchantsIdBankingPath, h.setMerchantBanking)
	groups.AuthUser.GET(merchantsStatusCompanyPath, h.getMerchantStatus)

	groups.SystemUser.PUT(merchantsIdChangeStatusCompanyPath, h.changeMerchantStatus)
	groups.AuthUser.PATCH(merchantsPath, h.changeAgreement)

	groups.AuthUser.GET(merchantsAgreementPath, h.getMerchantAgreementData)
	groups.SystemUser.GET(merchantsIdAgreementPath, h.getSystemAgreementData)
	groups.AuthUser.GET(merchantsAgreementDocumentPath, h.getAgreementDocument)
	groups.SystemUser.GET(merchantsIdAgreementDocumentPath, h.getAgreementDocument)

	groups.SystemUser.POST(merchantsIdNotificationsPath, h.createNotification)
	groups.SystemUser.GET(merchantsIdNotificationsPath, h.listNotifications)
	groups.AuthUser.GET(merchantsNotificationsIdPath, h.getNotification)
	groups.AuthUser.GET(merchantsNotificationsPath, h.listNotifications)
	groups.AuthUser.PUT(merchantsNotificationsMarkReadPath, h.markAsReadNotification)

	groups.AuthUser.GET(merchantsTariffsPath, h.getTariffRates)
	groups.AuthUser.POST(merchantsTariffsPath, h.setTariffRates)
	groups.SystemUser.POST(merchantsIdTariffsPath, h.setTariffRates)

	groups.AuthUser.PUT(merchantsIdManualPayoutEnablePath, h.enableMerchantManualPayout)
	groups.AuthUser.PUT(merchantsIdManualPayoutDisablePath, h.disableMerchantManualPayout)

	groups.SystemUser.POST(merchantsIdSetOperatingCompanyPath, h.setOperatingCompany)
}

func (h *OnboardingRoute) getMerchant(ctx echo.Context) error {
	req := &grpc.GetMerchantByRequest{}
	err := ctx.Bind(req)

	res, err := h.dispatch.Services.Billing.GetMerchantBy(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "GetMerchantBy", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

func (h *OnboardingRoute) getMerchantByUser(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	if authUser.Id == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, common.ErrorMessageAccessDenied)
	}

	req := &grpc.GetMerchantByRequest{UserId: authUser.Id}
	res, err := h.dispatch.Services.Billing.GetMerchantBy(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "GetMerchantBy", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

func (h *OnboardingRoute) listMerchants(ctx echo.Context) error {
	req := &grpc.MerchantListingRequest{}
	err := (&common.OnboardingMerchantListingBinder{
		LimitDefault:  int64(h.cfg.LimitDefault),
		OffsetDefault: int64(h.cfg.OffsetDefault),
	}).Bind(req, ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListMerchants(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ListMerchants", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *OnboardingRoute) changeMerchantStatus(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &grpc.MerchantChangeStatusRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	req.UserId = authUser.Id
	res, err := h.dispatch.Services.Billing.ChangeMerchantStatus(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ChangeMerchantStatus", req)
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

func (h *OnboardingRoute) createNotification(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &grpc.NotificationRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	req.UserId = authUser.Id
	res, err := h.dispatch.Services.Billing.CreateNotification(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "CreateNotification", req)
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusCreated, res.Item)
}

// @summary Get the merchant's notification
// @desc Get the merchant's notification using the notification ID
// @id merchantsNotificationsIdPathGetNotification
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.Notification Returns the notification data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param notification_id path {string} true The unique identifier for the notification.
// @router /admin/api/v1/merchants/notifications/{notification_id} [get]
func (h *OnboardingRoute) getNotification(ctx echo.Context) error {
	req := &grpc.GetNotificationRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetNotification(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "GetNotification")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the list of merchant's notifications
// @desc Get the list of merchant's notifications
// @id merchantsNotificationsPathListNotifications
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.Notifications Returns the list of notifications
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param user query {string} false The unique identifier for the user who is the sender of the notification.
// @param is_system query {integer} false Available values: 1 - the notifications between the merchant's owner and the PaySuper admin, 2 - the notifications generated automatically.
// @param limit query {integer} false The number of notifications returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @param sort query {[]string} false The list of the notification's fields for sorting.
// @router /admin/api/v1/merchants/notifications [get]
func (h *OnboardingRoute) listNotifications(ctx echo.Context) error {
	req := &grpc.ListingNotificationRequest{}
	err := (&common.OnboardingNotificationsListBinder{
		LimitDefault:  int64(h.cfg.LimitDefault),
		OffsetDefault: int64(h.cfg.OffsetDefault),
	}).Bind(req, ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ListNotifications(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ListNotifications", req)
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorUnknown)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Mark the notification status as read
// @desc Mark the notification status as read using the notification ID
// @id merchantsNotificationsMarkReadPathMarkAsReadNotification
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} billing.Notification Returns the notification data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param notification_id path {string} true The unique identifier for the notification.
// @router /admin/api/v1/merchants/notifications/{notification_id}/mark-as-read [put]
func (h *OnboardingRoute) markAsReadNotification(ctx echo.Context) error {
	req := &grpc.GetNotificationRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.MarkNotificationAsRead(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "MarkNotificationAsRead")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Update the merchant's agreement signature
// @desc Update the merchant's agreement signature using the merchant ID
// @id merchantsPathChangeAgreement
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body grpc.ChangeMerchantDataRequest
// @success 200 {object} billing.Merchant Returns the merchant data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param project_id path {string} true The unique identifier for the project.
// @router /admin/api/v1/merchants [patch]
func (h *OnboardingRoute) changeAgreement(ctx echo.Context) error {
	req := &grpc.ChangeMerchantDataRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ChangeMerchantData(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "ChangeMerchantData")
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Download the merchant's agreement document
// @desc Download the merchant's agreement document
// @id merchantsAgreementDocumentPathGetAgreementDocument
// @tag Onboarding
// @accept application/json
// @produce application/pdf
// @success 200 {string} Returns the merchant's agreement file
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/agreement/document [get]
func (h *OnboardingRoute) getAgreementDocument(ctx echo.Context) error {
	req := &grpc.GetMerchantByRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	res, err := h.dispatch.Services.Billing.GetMerchantBy(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "GetMerchantBy")
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	if res.Item.S3AgreementName == "" {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorMessageAgreementNotGenerated)
	}

	filePath := os.TempDir() + string(os.PathSeparator) + res.Item.S3AgreementName
	_, err = h.awsManager.Download(ctx.Request().Context(), filePath, &awsWrapper.DownloadInput{FileName: res.Item.S3AgreementName})

	if err != nil {
		h.L().Error("AWS api call to download file failed", logger.PairArgs("err", err.Error(), "file_name", res.Item.S3AgreementName))

		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorAgreementFileNotExist)
	}

	return ctx.File(filePath)
}

func (h *OnboardingRoute) getAgreementStructure(
	ctx echo.Context,
	merchantId, ext, ct, fPath string,
	signerType int32,
) (interface{}, error) {
	file, err := os.Open(fPath)

	if err != nil {
		return nil, common.ErrorMessageAgreementNotFound
	}

	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	fi, err := file.Stat()

	if err != nil {
		return nil, common.ErrorMessageAgreementNotFound
	}

	url := fmt.Sprintf(systemAgreementUrlMask, h.cfg.HttpScheme, ctx.Request().Host, merchantId)

	if signerType == pkg.SignerTypeMerchant {
		url = fmt.Sprintf(merchantAgreementUrlMask, h.cfg.HttpScheme, ctx.Request().Host)
	}

	data := &OnboardingFileData{
		Url: url,
		Metadata: &OnboardingFileMetadata{
			Name:        fi.Name(),
			Extension:   ext,
			ContentType: ct,
			Size:        fi.Size(),
		},
	}

	return data, nil
}

func (h *OnboardingRoute) validateUpload(file *multipart.FileHeader) (multipart.File, error) {
	if file.Size > agreementUploadMaxSize {
		return nil, common.ErrorMessageAgreementUploadMaxSize
	}

	src, err := file.Open()

	if err != nil {
		h.L().Error("validate upload error", logger.PairArgs("err", err.Error()))
		return nil, common.ErrorUnknown
	}

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)

	if err != nil {
		h.L().Error("validate upload error", logger.PairArgs("err", err.Error()))
		return nil, common.ErrorUnknown
	}

	_, err = src.Seek(0, 0)

	if err != nil {
		h.L().Error("validate upload error", logger.PairArgs("err", err.Error()))
		return nil, common.ErrorUnknown
	}

	ct := http.DetectContentType(buffer)

	if ct != agreementContentType {
		return nil, common.ErrorMessageAgreementContentType
	}

	return src, nil
}

// @summary Update the merchant's company information
// @desc Update the merchant's company information for the authorized merchant
// @id merchantsCompanyPathSetMerchantCompany
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MerchantCompanyInfo
// @success 200 {object} billing.Merchant Returns the merchant data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error

// @router /admin/api/v1/merchants/company [put]
func (h *OnboardingRoute) setMerchantCompany(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	in := &billing.MerchantCompanyInfo{}
	err := ctx.Bind(in)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req := &grpc.OnboardingRequest{
		Company: in,
		Id:      ctx.Param(common.RequestParameterMerchantId),
		User: &billing.MerchantUser{
			Id:    authUser.Id,
			Email: authUser.Email,
		},
	}
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ChangeMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ChangeMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Update the merchant's contacts
// @desc Update the merchant's contacts for the authorized merchant
// @id merchantsContactsPathSetMerchantContacts
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MerchantContact
// @success 200 {object} billing.Merchant Returns the merchant data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/contacts [put]
func (h *OnboardingRoute) setMerchantContacts(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	in := &billing.MerchantContact{}
	err := ctx.Bind(in)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req := &grpc.OnboardingRequest{
		Contacts: in,
		Id:       ctx.Param(common.RequestParameterMerchantId),
		User: &billing.MerchantUser{
			Id:    authUser.Id,
			Email: authUser.Email,
		},
	}
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ChangeMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ChangeMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Update the merchant's banking data
// @desc Update the merchant's banking data for the authorized merchant
// @id merchantsBankingPathSetMerchantBanking
// @tag Onboarding
// @accept application/json
// @produce application/json
// @body billing.MerchantBanking
// @success 200 {object} billing.Merchant Returns the merchant data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/banking [put]
func (h *OnboardingRoute) setMerchantBanking(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	in := &billing.MerchantBanking{}
	err := ctx.Bind(in)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req := &grpc.OnboardingRequest{
		Banking: in,
		Id:      ctx.Param(common.RequestParameterMerchantId),
		User: &billing.MerchantUser{
			Id:    authUser.Id,
			Email: authUser.Email,
		},
	}
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.ChangeMerchant(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "ChangeMerchant", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the merchant profile filling out status
// @desc Get the merchant profile filling out status for every steps
// @id merchantsStatusCompanyPathGetMerchantStatus
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.GetMerchantOnboardingCompleteDataResponseItem Returns the merchant profile filling out status
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/status [get]
func (h *OnboardingRoute) getMerchantStatus(ctx echo.Context) error {
	req := &grpc.SetMerchantS3AgreementRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantOnboardingCompleteData(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "GetMerchantOnboardingCompleteData")
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the tariff rates
// @desc Get the tariff rates
// @id merchantsTariffsPathGetTariffRates
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.GetMerchantTariffRatesResponseItems Returns the tariff rates for the specified region
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param region query {string} true The merchant's home region name. Available values: asia, europe, latin_america, russia_and_cis, worldwide.
// @param payer_region query {string} false The payer's region name. Available values: asia, europe, latin_america, russia_and_cis, worldwide.
// @param min_amount query {integer} false The minimum payment amount.
// @param max_amount query {integer} false The maximum payment amount.
// @param merchant_operations_type query {string} true The merchant's operations type. Available values: low-risk, high-risk.
// @router /admin/api/v1/merchants/tariffs [get]
func (h *OnboardingRoute) getTariffRates(ctx echo.Context) error {
	req := &grpc.GetMerchantTariffRatesRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetMerchantTariffRates(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "GetMerchantTariffRates", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Items)
}

// @summary Set the tariff rates
// @desc Set the tariff rates using the merchant ID
// @id merchantsTariffsPathSetTariffRates
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {string} Returns an empty response body if the tariff was successfully set
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @param merchant_operations_type query {string} true The merchant's operations type. Available values: low-risk, high-risk.
// @router /admin/api/v1/merchants/tariffs [post]
func (h *OnboardingRoute) setTariffRates(ctx echo.Context) error {
	req := &grpc.SetMerchantTariffRatesRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetMerchantTariffRates(
		ctx.Request().Context(),
		req,
		client.WithRequestTimeout(time.Minute*10),
	)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "SetMerchantTariffRates", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

// @summary Create the agreement
// @desc Create the agreement
// @id merchantsAgreementPathGetMerchantAgreementData
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {object} OnboardingFileData Returns the printable agreement document
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/agreement [get]
func (h *OnboardingRoute) getMerchantAgreementData(ctx echo.Context) error {
	return h.getAgreementData(ctx, pkg.SignerTypeMerchant)
}

func (h *OnboardingRoute) getSystemAgreementData(ctx echo.Context) error {
	return h.getAgreementData(ctx, pkg.SignerTypePs)
}

func (h *OnboardingRoute) getAgreementData(ctx echo.Context, signerType int32) error {
	req := &grpc.GetMerchantByRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetMerchantBy(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "GetMerchantBy")
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	if res.Item.S3AgreementName == "" {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorMessageAgreementNotFound)
	}

	filePath := os.TempDir() + string(os.PathSeparator) + res.Item.S3AgreementName
	_, err = h.awsManager.Download(ctx.Request().Context(), filePath, &awsWrapper.DownloadInput{FileName: res.Item.S3AgreementName})

	if err != nil {
		h.L().Error("AWS api call to download file failed", logger.PairArgs("err", err.Error(), "file_name", res.Item.S3AgreementName))

		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	fData, err := h.getAgreementStructure(ctx, req.MerchantId, agreementExtension, agreementContentType, filePath, signerType)

	if err != nil {
		h.L().Error("Get agreement structure failed", logger.PairArgs("err", err.Error(), "merchant_id", req.MerchantId))

		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	return ctx.JSON(http.StatusOK, fData)
}

// @summary Enable the manual payouts for the merchant
// @desc Enable the manual payouts for the merchant
// @id merchantsIdManualPayoutEnablePathEnableMerchantManualPayout
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {string} Returns an empty response body if the manual payouts was enabled
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/manual_payout/enable [put]
func (h *OnboardingRoute) enableMerchantManualPayout(ctx echo.Context) error {
	return h.changeMerchantManualPayout(ctx, true)
}

// @summary Disable the manual payouts for the merchant
// @desc Disable the manual payouts for the merchant
// @id merchantsIdManualPayoutDisablePathDisableMerchantManualPayout
// @tag Onboarding
// @accept application/json
// @produce application/json
// @success 200 {string} Returns an empty response body if the manual payouts was disabled
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/merchants/manual_payout/disable [put]
func (h *OnboardingRoute) disableMerchantManualPayout(ctx echo.Context) error {
	return h.changeMerchantManualPayout(ctx, false)
}

func (h *OnboardingRoute) changeMerchantManualPayout(ctx echo.Context, enableManualPayout bool) error {
	req := &grpc.ChangeMerchantManualPayoutsRequest{ManualPayoutsEnabled: enableManualPayout}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ChangeMerchantManualPayouts(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, pkg.ServiceName, "ChangeMerchantManualPayouts")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

func (h *OnboardingRoute) setOperatingCompany(ctx echo.Context) error {
	req := &grpc.SetMerchantOperatingCompanyRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.MerchantId = ctx.Param(common.RequestParameterMerchantId)
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.SetMerchantOperatingCompany(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, pkg.ServiceName, "SetMerchantOperatingCompany", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != pkg.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}
