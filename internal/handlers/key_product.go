package handlers

import (
	"github.com/ProtocolONE/geoip-service/pkg/proto"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/client"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	keyProductsPath               = "/key-products"
	keyProductsIdPath             = "/key-products/:key_product_id"
	keyProductsPublishPath        = "/key-products/:key_product_id/publish"
	keyProductsUnPublishPath      = "/key-products/:key_product_id/unpublish"
	platformsPath                 = "/platforms"
	keyProductsPlatformsFilePath  = "/key-products/:key_product_id/platforms/:platform_id/file"
	keyProductsPlatformsCountPath = "/key-products/:key_product_id/platforms/:platform_id/count"
)

type KeyProductRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewKeyProductRoute(set common.HandlerSet, cfg *common.Config) *KeyProductRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "KeyProductRoute"})
	return &KeyProductRoute{
		dispatch: set,
		cfg:      *cfg,
		LMT:      &set.AwareSet,
	}
}

func (h *KeyProductRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(keyProductsPath, h.getKeyProductList)
	groups.AuthUser.POST(keyProductsPath, h.createKeyProduct)
	groups.AuthUser.GET(keyProductsIdPath, h.getKeyProductById)
	groups.AuthUser.PUT(keyProductsIdPath, h.changeKeyProduct)
	groups.AuthUser.POST(keyProductsPublishPath, h.publishKeyProduct)
	groups.AuthUser.POST(keyProductsUnPublishPath, h.unpublishKeyProduct)
	groups.AuthUser.DELETE(keyProductsIdPath, h.deleteKeyProductById)
	groups.AuthUser.GET(platformsPath, h.getPlatformsList)

	groups.AuthUser.POST(keyProductsPlatformsFilePath, h.uploadKeys)
	groups.AuthUser.GET(keyProductsPlatformsCountPath, h.getCountOfKeys)

	groups.AuthProject.GET(keyProductsIdPath, h.getKeyProduct)
}

// @summary Make the key-activated product inactive
// @desc Make the key-activated product inactive
// @id keyProductsUnPublishPathUnpublishKeyProduct
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/key-products/{key_product_id}/unpublish [post]
func (h *KeyProductRoute) unpublishKeyProduct(ctx echo.Context) error {
	req := &billingpb.UnPublishKeyProductRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.KeyProductId = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.UnPublishKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Product)
}

// @summary Make the key-activated product active
// @desc Make the key-activated product active
// @id keyProductsPublishPathPublishKeyProduct
// @tag Product
// @accept application/json
// @produce application/json
// @body billingpb.PublishKeyProductRequest
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/key-products/{key_product_id}/publish [post]
func (h *KeyProductRoute) publishKeyProduct(ctx echo.Context) error {
	req := &billingpb.PublishKeyProductRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.KeyProductId = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.PublishKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Message != nil {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Product)
}

// @summary Get available platforms list
// @desc Get available platforms list
// @id platformsPathGetPlatformsList
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListPlatformsResponse Returns the available platforms list
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param limit query {integer} true The number of platforms returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/platforms [get]
func (h *KeyProductRoute) getPlatformsList(ctx echo.Context) error {
	req := &billingpb.ListPlatformsRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit == 0 {
		req.Limit = h.cfg.LimitDefault
	}

	if req.Limit > h.cfg.LimitMax {
		req.Limit = h.cfg.LimitMax
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetPlatforms(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Remove the key-activated product
// @desc Remove the key-activated product using the product ID
// @id keyProductsIdPathDeleteKeyProductById
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/key-products/{key_product_id} [delete]
func (h *KeyProductRoute) deleteKeyProductById(ctx echo.Context) error {
	req := &billingpb.RequestKeyProductMerchant{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Id = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.DeleteKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

// @summary Update the key-activated product
// @desc Update the  key-activated product
// @id keyProductsIdPathChangeKeyProduct
// @tag Product
// @accept application/json
// @produce application/json
// @body billingpb.CreateOrUpdateKeyProductRequest
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/key-products/{key_product_id} [put]
func (h *KeyProductRoute) changeKeyProduct(ctx echo.Context) error {
	req := &billingpb.CreateOrUpdateKeyProductRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Id = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdateKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Product)
}

// @summary Get the key-activated product using the product ID
// @desc Get the key-activated product using the product ID
// @id keyProductsIdPathGetKeyProductById
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/key-products/{key_product_id} [get]
func (h *KeyProductRoute) getKeyProductById(ctx echo.Context) error {
	req := &billingpb.RequestKeyProductMerchant{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.Id = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Product)
}

// @summary Create a new key-activated product
// @desc Create a new key-activated product
// @id keyProductsPathCreateKeyProduct
// @tag Product
// @accept application/json
// @produce application/json
// @body billingpb.CreateOrUpdateKeyProductRequest
// @success 200 {object} billingpb.KeyProduct Returns the key-activated product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/key-products [post]
func (h *KeyProductRoute) createKeyProduct(ctx echo.Context) error {
	req := &billingpb.CreateOrUpdateKeyProductRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	h.L().Info("createKeyProduct", logger.PairArgs("req", req))

	res, err := h.dispatch.Services.Billing.CreateOrUpdateKeyProduct(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Message != nil {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusCreated, res.Product)
}

// @summary Get the list of the key-activated products
// @desc Get the list of the key-activated products. This list can be filtered by name, sku and other parameters.
// @id keyProductsPathGetKeyProductList
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListKeyProductsResponse Returns the list of the key-activated products
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param name query {string} true The unique identifier for the key-activated product.
// @param sku query {string} false The SKU of the product.
// @param project_id query {string} false The unique identifier for the project.
// @param enabled query {string} false The status of whether the product is enabled. Available values: all, true, false.
// @param limit query {integer} true The number of products returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/key-products [get]
func (h *KeyProductRoute) getKeyProductList(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &billingpb.ListKeyProductsRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	if req.Limit > int64(h.cfg.LimitMax) {
		req.Limit = int64(h.cfg.LimitMax)
	}

	if req.Limit <= 0 {
		req.Limit = int64(h.cfg.LimitDefault)
	}

	req.MerchantId = authUser.MerchantId

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetKeyProducts(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Message != nil {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the list of platforms and prices for the key-activated product
// @desc Get the list of platforms and prices for the key-activated product. This list can be filtered by country, language or currency.
// @id keyProductsIdPathGetKeyProduct
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.KeyProductInfo Returns the product data (platforms, prices and others)
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the key-activated product.
// @param country query {string} false The country code to calculate the price for.
// @param language query {string} false The language.
// @param currency query {string} false The price currency.
// @router /auth/api/v1/key-products/{key_product_id} [get]
func (h *KeyProductRoute) getKeyProduct(ctx echo.Context) error {
	req := &billingpb.GetKeyProductInfoRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.KeyProductId = ctx.Param("key_product_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	if req.Currency == "" && req.Country == "" {
		res, err := h.dispatch.Services.Geo.GetIpData(ctx.Request().Context(), &proto.GeoIpDataRequest{IP: ctx.RealIP()})
		if err != nil {
			h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		} else {
			req.Country = res.Country.IsoCode
		}
	}

	if req.Language == "" {
		req.Language, _ = h.getCountryFromAcceptLanguage(ctx.Request().Header.Get(common.HeaderAcceptLanguage))
	}

	res, err := h.dispatch.Services.Billing.GetKeyProductInfo(ctx.Request().Context(), req)
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.KeyProduct)
}

// @summary Send the file with the list of keys
// @desc Send the file with the list of keys to process them
// @id keyProductsPlatformsFilePathUploadKeys
// @tag Product
// @accept application/octet-stream
// @produce application/json
// @body billingpb.PlatformKeysFileRequest
// @success 200 {object} billingpb.PlatformKeysFileResponse Returns the number of the processed keys
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the key-activated product.
// @param platform_id path {string} true The platform's name. Available values: steam, gog, uplay, origin, psn, xbox, nintendo, itch, egs.
// @router /admin/api/v1/key-products/{key_product_id}/platforms/{platform_id}/file [post]
func (h *KeyProductRoute) uploadKeys(ctx echo.Context) error {
	req := &billingpb.PlatformKeysFileRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		h.L().Error(common.ErrorMessageFileNotFound.String(), logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorMessageFileNotFound)
	}

	src, err := file.Open()
	if err != nil {
		h.L().Error(common.ErrorMessageCantReadFile.String(), logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorMessageCantReadFile)
	}
	defer src.Close()

	req.File, err = ioutil.ReadAll(src)

	if err != nil {
		h.L().Error(common.ErrorMessageCantReadFile.String(), logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorMessageCantReadFile)
	}

	req.KeyProductId = ctx.Param("key_product_id")
	req.PlatformId = ctx.Param("platform_id")

	keyProductRes, err := h.dispatch.Services.Billing.GetKeyProduct(ctx.Request().Context(), &billingpb.RequestKeyProductMerchant{Id: req.KeyProductId, MerchantId: req.MerchantId})
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if keyProductRes.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(keyProductRes.Status), keyProductRes.Message)
	}

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.UploadKeysFile(ctx.Request().Context(), req, client.WithRequestTimeout(time.Minute*10))
	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the number of keys for the specified platform and product
// @desc Get the number of keys for the specified platform and product
// @id keyProductsPlatformsCountPathGetCountOfKeys
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetPlatformKeyCountResponse Returns the number of keys for the specified platform and product
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 404 {object} billingpb.ResponseErrorMessage Not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param key_product_id path {string} true The unique identifier for the key-activated product.
// @param platform_id path {string} true The platform's name. Available values: steam, gog, uplay, origin, psn, xbox, nintendo, itch, egs.
// @router /auth/api/v1/key-products/{key_product_id}/platforms/{platform_id}/count [get]
func (h *KeyProductRoute) getCountOfKeys(ctx echo.Context) error {
	req := &billingpb.GetPlatformKeyCountRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.KeyProductId = ctx.Param("key_product_id")
	req.PlatformId = ctx.Param("platform_id")

	if err := h.dispatch.Validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetAvailableKeysCount(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.PairArgs("err", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorInternal)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *KeyProductRoute) getCountryFromAcceptLanguage(acceptLanguage string) (string, string) {
	it := strings.Split(acceptLanguage, ",")

	if len(it) == 0 {
		return "", ""
	}

	if strings.Index(it[0], "-") == -1 {
		return "", ""
	}

	it = strings.Split(it[0], "-")

	return strings.ToLower(it[0]), strings.ToUpper(it[1])
}
