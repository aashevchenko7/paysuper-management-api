package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"

	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	productsPath         = "/products"
	productsMerchantPath = "/products/merchant/:merchant_id"
	productsIdPath       = "/products/:product_id"
	productsPricesPath   = "/products/:product_id/prices"
)

type ProductRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewProductRoute(set common.HandlerSet, cfg *common.Config) *ProductRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "ProductRoute"})
	return &ProductRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *ProductRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(productsPath, h.getProductsList)
	groups.SystemUser.GET(productsMerchantPath, h.getProductsList)
	groups.AuthUser.POST(productsPath, h.createProduct)
	groups.AuthUser.GET(productsIdPath, h.getProduct)
	groups.AuthUser.PUT(productsIdPath, h.updateProduct)
	groups.AuthUser.DELETE(productsIdPath, h.deleteProduct)
	groups.AuthUser.GET(productsPricesPath, h.getProductPrices)    // TODO: Need test
	groups.AuthUser.PUT(productsPricesPath, h.updateProductPrices) // TODO: Need test
}

// @summary Get the list of products
// @desc Get the list of products for the authorized user
// @id productsPathGetProductsList
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListProductsResponse Returns the list of products for the authorized user. This list can be filtered by the product's name, SKU, status and the project ID.
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param name query {string} false The product's name.
// @param sku query {string} false The SKU of the product.
// @param project_id query {string} false The unique identifier for the project.
// @param enabled query {string} false The status of whether the product is enabled. Available values: all, true, false.
// @param limit query {integer} true The number of products returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /admin/api/v1/products [get]

// @summary Get the list of products using the merchant ID
// @desc Get the list of products using the merchant ID
// @id productsMerchantPathGetProductsList
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ListProductsResponse Returns the list of merchant's products. This list can be filtered by the product's name, SKU, status and the project ID.
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param merchant_id path {string} true The unique identifier for the merchant.
// @param name query {string} false The product's name.
// @param sku query {string} false The SKU of the product.
// @param project_id query {string} false The unique identifier for the project.
// @param enabled query {string} false The status of whether the product is enabled. Available values: all, true, false.
// @param limit query {integer} true The number of products returned in one page. Default value is 100.
// @param offset query {integer} false The ranking number of the first item on the page.
// @router /system/api/v1/products/merchant/{merchant_id} [get]
func (h *ProductRoute) getProductsList(ctx echo.Context) error {

	req := &billingpb.ListProductsRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.ListProducts(ctx.Request().Context(), req)
	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ListProducts")
	}
	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the product
// @desc Get the product using the product ID
// @id productsIdPathGetProduct
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.GetProductResponse Returns the product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/products/{product_id} [get]
func (h *ProductRoute) getProduct(ctx echo.Context) error {

	req := &billingpb.RequestProduct{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetProduct(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetProduct")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Delete the product
// @desc Delete the product using the product ID
// @id productsIdPathDeleteProduct
// @tag Product
// @accept application/json
// @produce application/json
// @success 204 {string} Returns an empty response body if the product was successfully removed
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 401 {object} billingpb.ResponseErrorMessage Unauthorized request
// @failure 403 {object} billingpb.ResponseErrorMessage Access denied
// @failure 404 {object} billingpb.ResponseErrorMessage The product not found
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/products/{product_id} [delete]
func (h *ProductRoute) deleteProduct(ctx echo.Context) error {

	req := &billingpb.RequestProduct{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	_, err := h.dispatch.Services.Billing.DeleteProduct(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "DeleteProduct")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// @summary Create a product
// @desc Create a new product in the project
// @id productsPathCreateProduct
// @tag Product
// @accept application/json
// @produce application/json
// @body billingpb.Product
// @success 200 {object} billingpb.Product Returns the created product data
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @router /admin/api/v1/products [post]
func (h *ProductRoute) createProduct(ctx echo.Context) error {
	return h.createOrUpdateProduct(ctx, &common.ProductsCreateProductBinder{})
}

// @summary Update the product
// @desc Update the product using the product ID
// @id productsIdPathUpdateProduct
// @tag Product
// @accept application/json
// @produce application/json
// @body billingpb.Product
// @success 200 {object} billingpb.Product Returns the updated product
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/products/{product_id} [put]
func (h *ProductRoute) updateProduct(ctx echo.Context) error {
	return h.createOrUpdateProduct(ctx, &common.ProductsUpdateProductBinder{})
}

func (h *ProductRoute) createOrUpdateProduct(ctx echo.Context, binder echo.Binder) error {

	req := &billingpb.Product{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdateProduct(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "CreateOrUpdateProduct")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Get the product's prices
// @desc Get the product's prices using the product ID
// @id productsPricesPathGetProductPrices
// @tag Product
// @accept application/json
// @produce application/json
// @success 200 {object} billingpb.ProductPricesResponse Returns the list of the product's prices
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/products/{product_id}/prices [get]
func (h *ProductRoute) getProductPrices(ctx echo.Context) error {

	req := &billingpb.RequestProduct{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.GetProductPrices(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetProductPrices")
	}

	return ctx.JSON(http.StatusOK, res)
}

// @summary Set the product's price
// @desc Set the product's price using the product ID
// @id productsPricesPathUpdateProductPrices
// @tag Product
// @accept application/json
// @produce application/json
// @body []billingpb.ProductPrice
// @success 200 {string} Returns an empty response body if the product's price was successfully set
// @failure 400 {object} billingpb.ResponseErrorMessage Invalid request data
// @failure 500 {object} billingpb.ResponseErrorMessage Internal Server Error
// @param product_id path {string} true The unique identifier for the product.
// @router /admin/api/v1/products/{product_id}/prices [put]
func (h *ProductRoute) updateProductPrices(ctx echo.Context) error {

	req := &billingpb.UpdateProductPricesRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return err
	}

	res, err := h.dispatch.Services.Billing.UpdateProductPrices(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "UpdateProductPrices")
	}

	return ctx.JSON(http.StatusOK, res)
}
