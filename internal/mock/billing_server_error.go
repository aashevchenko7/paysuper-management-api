package mock

import (
	"context"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/paysuper/paysuper-proto/go/billingpb"

	"net/http"
)

type BillingServerErrorMock struct{}

func (s *BillingServerErrorMock) GetVatReportTransactions(ctx context.Context, in *billingpb.VatTransactionsRequest, opts ...client.CallOption) (*billingpb.PrivateTransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetVatReport(ctx context.Context, in *billingpb.VatReportRequest, opts ...client.CallOption) (*billingpb.VatReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) OrderReCreateProcess(ctx context.Context, in *billingpb.OrderReCreateProcessRequest, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func NewBillingServerErrorMock() billingpb.BillingService {
	return &BillingServerErrorMock{}
}

func (s *BillingServerErrorMock) GetProductsForOrder(
	ctx context.Context,
	in *billingpb.GetProductsForOrderRequest,
	opts ...client.CallOption,
) (*billingpb.ListProductsResponse, error) {
	return &billingpb.ListProductsResponse{}, nil
}

func (s *BillingServerErrorMock) OrderCreateProcess(
	ctx context.Context,
	in *billingpb.OrderCreateRequest,
	opts ...client.CallOption,
) (*billingpb.OrderCreateProcessResponse, error) {
	return &billingpb.OrderCreateProcessResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Order{},
	}, nil
}

func (s *BillingServerErrorMock) PaymentFormJsonDataProcess(
	ctx context.Context,
	in *billingpb.PaymentFormJsonDataRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormJsonDataResponse, error) {
	return &billingpb.PaymentFormJsonDataResponse{}, nil
}

func (s *BillingServerErrorMock) PaymentCreateProcess(
	ctx context.Context,
	in *billingpb.PaymentCreateRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentCreateResponse, error) {
	return &billingpb.PaymentCreateResponse{}, nil
}

func (s *BillingServerErrorMock) PaymentCallbackProcess(
	ctx context.Context,
	in *billingpb.PaymentNotifyRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{}, nil
}

func (s *BillingServerErrorMock) RebuildCache(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerErrorMock) UpdateOrder(
	ctx context.Context,
	in *billingpb.Order,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerErrorMock) UpdateMerchant(
	ctx context.Context,
	in *billingpb.Merchant,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerErrorMock) GetMerchantBy(
	ctx context.Context,
	in *billingpb.GetMerchantByRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantResponse, error) {
	return &billingpb.GetMerchantResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ListMerchants(
	ctx context.Context,
	in *billingpb.MerchantListingRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantListingResponse, error) {
	return &billingpb.MerchantListingResponse{}, nil
}

func (s *BillingServerErrorMock) ChangeMerchant(
	ctx context.Context,
	in *billingpb.OnboardingRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantResponse, error) {
	return &billingpb.ChangeMerchantResponse{
		Status:  http.StatusBadRequest,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ChangeMerchantStatus(
	ctx context.Context,
	in *billingpb.MerchantChangeStatusRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantStatusResponse, error) {
	return &billingpb.ChangeMerchantStatusResponse{
		Status:  http.StatusBadRequest,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreateNotification(
	ctx context.Context,
	in *billingpb.NotificationRequest,
	opts ...client.CallOption,
) (*billingpb.CreateNotificationResponse, error) {
	return &billingpb.CreateNotificationResponse{
		Status:  http.StatusBadRequest,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetNotification(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ListNotifications(
	ctx context.Context,
	in *billingpb.ListingNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notifications, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) MarkNotificationAsRead(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ListMerchantPaymentMethods(
	ctx context.Context,
	in *billingpb.ListMerchantPaymentMethodsRequest,
	opts ...client.CallOption,
) (*billingpb.ListingMerchantPaymentMethod, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.GetMerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantPaymentMethodResponse, error) {
	return &billingpb.GetMerchantPaymentMethodResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ChangeMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.MerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantPaymentMethodResponse, error) {
	return &billingpb.MerchantPaymentMethodResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreateRefund(
	ctx context.Context,
	in *billingpb.CreateRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ListRefunds(
	ctx context.Context,
	in *billingpb.ListRefundsRequest,
	opts ...client.CallOption,
) (*billingpb.ListRefundsResponse, error) {
	return &billingpb.ListRefundsResponse{}, nil
}

func (s *BillingServerErrorMock) GetRefund(
	ctx context.Context,
	in *billingpb.GetRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status:  billingpb.ResponseStatusNotFound,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ProcessRefundCallback(
	ctx context.Context,
	in *billingpb.CallbackRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{
		Status: billingpb.ResponseStatusNotFound,
		Error:  SomeError.Message,
	}, nil
}

func (s *BillingServerErrorMock) PaymentFormLanguageChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangeLangRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return &billingpb.PaymentFormDataChangeResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) PaymentFormPaymentAccountChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePaymentAccountRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return &billingpb.PaymentFormDataChangeResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ProcessBillingAddress(
	ctx context.Context,
	in *billingpb.ProcessBillingAddressRequest,
	opts ...client.CallOption,
) (*billingpb.ProcessBillingAddressResponse, error) {
	return &billingpb.ProcessBillingAddressResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ChangeMerchantData(
	ctx context.Context,
	in *billingpb.ChangeMerchantDataRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return &billingpb.ChangeMerchantDataResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) SetMerchantS3Agreement(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return &billingpb.ChangeMerchantDataResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ChangeProject(
	ctx context.Context,
	in *billingpb.Project,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return &billingpb.ChangeProjectResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	if in.ProjectId == SomeMerchantId {
		return &billingpb.ChangeProjectResponse{
			Status: billingpb.ResponseStatusOk,
			Item: &billingpb.Project{
				MerchantId:         "ffffffffffffffffffffffff",
				Name:               map[string]string{"en": "A", "ru": "А"},
				CallbackCurrency:   "RUB",
				CallbackProtocol:   billingpb.ProjectCallbackProtocolEmpty,
				LimitsCurrency:     "RUB",
				MinPaymentAmount:   0,
				MaxPaymentAmount:   15000,
				IsProductsCheckout: false,
				VatPayer:           billingpb.VatPayerBuyer,
			},
		}, nil
	}

	return &billingpb.ChangeProjectResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) DeleteProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return &billingpb.ChangeProjectResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreateToken(
	ctx context.Context,
	in *billingpb.TokenRequest,
	opts ...client.CallOption,
) (*billingpb.TokenResponse, error) {
	return &billingpb.TokenResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CheckProjectRequestSignature(
	ctx context.Context,
	in *billingpb.CheckProjectRequestSignatureRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreateOrUpdateProduct(ctx context.Context, in *billingpb.Product, opts ...client.CallOption) (*billingpb.Product, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ListProducts(ctx context.Context, in *billingpb.ListProductsRequest, opts ...client.CallOption) (*billingpb.ListProductsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.GetProductResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) DeleteProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ListProjects(ctx context.Context, in *billingpb.ListProjectsRequest, opts ...client.CallOption) (*billingpb.ListProjectsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetOrder(ctx context.Context, in *billingpb.GetOrderRequest, opts ...client.CallOption) (*billingpb.Order, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) IsOrderCanBePaying(
	ctx context.Context,
	in *billingpb.IsOrderCanBePayingRequest,
	opts ...client.CallOption,
) (*billingpb.IsOrderCanBePayingResponse, error) {
	return &billingpb.IsOrderCanBePayingResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetCountry(ctx context.Context, in *billingpb.GetCountryRequest, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UpdateCountry(ctx context.Context, in *billingpb.Country, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPriceGroup(ctx context.Context, in *billingpb.GetPriceGroupRequest, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UpdatePriceGroup(ctx context.Context, in *billingpb.PriceGroup, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetUserNotifySales(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetUserNotifyNewRegion(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetCountriesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.CountriesList, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystemRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystem, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeletePaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchant, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeletePaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystemRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystem, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeleteMoneyBackCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchant, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeleteMoneyBackCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetAllPaymentChannelCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerErrorMock) GetAllPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantListRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerErrorMock) GetAllMoneyBackCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerErrorMock) GetAllMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantListRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerErrorMock) CreateOrUpdatePaymentMethodTestSettings(ctx context.Context, in *billingpb.ChangePaymentMethodParamsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeletePaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) FindByZipCode(
	ctx context.Context,
	in *billingpb.FindByZipCodeRequest,
	opts ...client.CallOption,
) (*billingpb.FindByZipCodeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) CreateOrUpdatePaymentMethod(
	ctx context.Context,
	in *billingpb.PaymentMethod,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) CreateOrUpdatePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.ChangePaymentMethodParamsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) DeletePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.GetPaymentMethodSettingsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) CreateAccountingEntry(ctx context.Context, in *billingpb.CreateAccountingEntryRequest, opts ...client.CallOption) (*billingpb.CreateAccountingEntryResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) CreateRoyaltyReport(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ListRoyaltyReports(ctx context.Context, in *billingpb.ListRoyaltyReportsRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ChangeRoyaltyReportStatus(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ListRoyaltyReportOrders(ctx context.Context, in *billingpb.ListRoyaltyReportOrdersRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetVatReportsDashboard(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetVatReportsForCountry(ctx context.Context, in *billingpb.VatReportsRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) CalcAnnualTurnovers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ProcessVatReports(ctx context.Context, in *billingpb.ProcessVatReportsRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UpdateVatReportStatus(ctx context.Context, in *billingpb.UpdateVatReportStatusRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPriceGroupByCountry(
	ctx context.Context,
	in *billingpb.PriceGroupByCountryRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroup, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetPriceGroupCurrencies(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetPriceGroupCurrencyByRegion(
	ctx context.Context,
	in *billingpb.PriceGroupByRegionRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetPriceGroupRecommendedPrice(
	ctx context.Context,
	in *billingpb.RecommendedPriceRequest,
	opts ...client.CallOption,
) (*billingpb.RecommendedPriceResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetProductPrices(
	ctx context.Context,
	in *billingpb.RequestProduct,
	opts ...client.CallOption,
) (*billingpb.ProductPricesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) UpdateProductPrices(
	ctx context.Context,
	in *billingpb.UpdateProductPricesRequest,
	opts ...client.CallOption,
) (*billingpb.ResponseError, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetPaymentMethodProductionSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) GetPaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ChangeRoyaltyReport(ctx context.Context, in *billingpb.ChangeRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) AutoAcceptRoyaltyReports(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetUserProfile(
	ctx context.Context,
	in *billingpb.GetUserProfileRequest,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return &billingpb.GetUserProfileResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreateOrUpdateUserProfile(
	ctx context.Context,
	in *billingpb.UserProfile,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return &billingpb.GetUserProfileResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ConfirmUserEmail(
	ctx context.Context,
	in *billingpb.ConfirmUserEmailRequest,
	opts ...client.CallOption,
) (*billingpb.ConfirmUserEmailResponse, error) {
	return &billingpb.ConfirmUserEmailResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CreatePageReview(
	ctx context.Context,
	in *billingpb.CreatePageReviewRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) MerchantReviewRoyaltyReport(ctx context.Context, in *billingpb.MerchantReviewRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantOnboardingCompleteData(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantOnboardingCompleteDataResponse, error) {
	return &billingpb.GetMerchantOnboardingCompleteDataResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.GetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantTariffRatesResponse, error) {
	return &billingpb.GetMerchantTariffRatesResponse{}, nil
}

func (s *BillingServerErrorMock) SetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.SetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{}, nil
}

func (s *BillingServerErrorMock) CreateOrUpdateKeyProduct(ctx context.Context, in *billingpb.CreateOrUpdateKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetKeyProducts(ctx context.Context, in *billingpb.ListKeyProductsRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return &billingpb.ListKeyProductsResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) DeleteKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) PublishKeyProduct(ctx context.Context, in *billingpb.PublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetKeyProductsForOrder(ctx context.Context, in *billingpb.GetKeyProductsForOrderRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return &billingpb.ListKeyProductsResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetPlatforms(ctx context.Context, in *billingpb.ListPlatformsRequest, opts ...client.CallOption) (*billingpb.ListPlatformsResponse, error) {
	return &billingpb.ListPlatformsResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) DeletePlatformFromProduct(ctx context.Context, in *billingpb.RemovePlatformRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetAvailableKeysCount(ctx context.Context, in *billingpb.GetPlatformKeyCountRequest, opts ...client.CallOption) (*billingpb.GetPlatformKeyCountResponse, error) {
	return &billingpb.GetPlatformKeyCountResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) UploadKeysFile(ctx context.Context, in *billingpb.PlatformKeysFileRequest, opts ...client.CallOption) (*billingpb.PlatformKeysFileResponse, error) {
	return &billingpb.PlatformKeysFileResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetKeyByID(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return &billingpb.GetKeyForOrderRequestResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ReserveKeyForOrder(ctx context.Context, in *billingpb.PlatformKeyReserveRequest, opts ...client.CallOption) (*billingpb.PlatformKeyReserveResponse, error) {
	return &billingpb.PlatformKeyReserveResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) FinishRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return &billingpb.GetKeyForOrderRequestResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) CancelRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) GetKeyProductInfo(ctx context.Context, in *billingpb.GetKeyProductInfoRequest, opts ...client.CallOption) (*billingpb.GetKeyProductInfoResponse, error) {
	return &billingpb.GetKeyProductInfoResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerErrorMock) ChangeCodeInOrder(ctx context.Context, in *billingpb.ChangeCodeInOrderRequest, opts ...client.CallOption) (*billingpb.ChangeCodeInOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetDashboardMainReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardMainResponse, error) {
	return &billingpb.GetDashboardMainResponse{}, nil
}
func (s *BillingServerErrorMock) GetDashboardRevenueDynamicsReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardRevenueDynamicsReportResponse, error) {
	return &billingpb.GetDashboardRevenueDynamicsReportResponse{}, nil
}

func (s *BillingServerErrorMock) GetDashboardBaseReport(
	ctx context.Context,
	in *billingpb.GetDashboardBaseReportRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardBaseReportResponse, error) {
	return &billingpb.GetDashboardBaseReportResponse{}, nil
}

func (s *BillingServerErrorMock) GetOrderPublic(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPublicResponse, error) {
	return &billingpb.GetOrderPublicResponse{}, nil
}

func (s *BillingServerErrorMock) GetOrderPrivate(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPrivateResponse, error) {
	return &billingpb.GetOrderPrivateResponse{}, nil
}

func (s *BillingServerErrorMock) FindAllOrdersPublic(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPublicResponse, error) {
	return &billingpb.ListOrdersPublicResponse{}, nil
}

func (s *BillingServerErrorMock) FindAllOrdersPrivate(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPrivateResponse, error) {
	return &billingpb.ListOrdersPrivateResponse{}, nil
}

func (s *BillingServerErrorMock) CreatePayoutDocument(ctx context.Context, in *billingpb.CreatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.CreatePayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UpdatePayoutDocument(ctx context.Context, in *billingpb.UpdatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPayoutDocuments(ctx context.Context, in *billingpb.GetPayoutDocumentsRequest, opts ...client.CallOption) (*billingpb.GetPayoutDocumentsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UpdatePayoutDocumentSignatures(ctx context.Context, in *billingpb.UpdatePayoutDocumentSignaturesRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantBalance(ctx context.Context, in *billingpb.GetMerchantBalanceRequest, opts ...client.CallOption) (*billingpb.GetMerchantBalanceResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) PayoutDocumentPdfUploaded(ctx context.Context, in *billingpb.PayoutDocumentPdfUploadedRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetRoyaltyReport(ctx context.Context, in *billingpb.GetRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.GetRoyaltyReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) UnPublishKeyProduct(ctx context.Context, in *billingpb.UnPublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) PaymentFormPlatformChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePlatformRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) OrderReceipt(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) OrderReceiptRefund(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetRecommendedPriceByPriceGroup(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetRecommendedPriceByConversion(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) CheckSkuAndKeyProject(ctx context.Context, in *billingpb.CheckSkuAndKeyProjectRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPriceGroupByRegion(ctx context.Context, in *billingpb.GetPriceGroupByRegionRequest, opts ...client.CallOption) (*billingpb.GetPriceGroupByRegionResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantUsers(ctx context.Context, in *billingpb.GetMerchantUsersRequest, opts ...client.CallOption) (*billingpb.GetMerchantUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) FindAllOrders(ctx context.Context, in *billingpb.ListOrdersRequest, opts ...client.CallOption) (*billingpb.ListOrdersResponse, error) {
	return nil, SomeError
}

func (s *BillingServerErrorMock) ChangeMerchantManualPayouts(ctx context.Context, in *billingpb.ChangeMerchantManualPayoutsRequest, opts ...client.CallOption) (*billingpb.ChangeMerchantManualPayoutsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetAdminUsers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetAdminUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantsForUser(ctx context.Context, in *billingpb.GetMerchantsForUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantsForUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) InviteUserMerchant(ctx context.Context, in *billingpb.InviteUserMerchantRequest, opts ...client.CallOption) (*billingpb.InviteUserMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) InviteUserAdmin(ctx context.Context, in *billingpb.InviteUserAdminRequest, opts ...client.CallOption) (*billingpb.InviteUserAdminResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ResendInviteMerchant(ctx context.Context, in *billingpb.ResendInviteMerchantRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ResendInviteAdmin(ctx context.Context, in *billingpb.ResendInviteAdminRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantUser(ctx context.Context, in *billingpb.GetMerchantUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetAdminUser(ctx context.Context, in *billingpb.GetAdminUserRequest, opts ...client.CallOption) (*billingpb.GetAdminUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) AcceptInvite(ctx context.Context, in *billingpb.AcceptInviteRequest, opts ...client.CallOption) (*billingpb.AcceptInviteResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) CheckInviteToken(ctx context.Context, in *billingpb.CheckInviteTokenRequest, opts ...client.CallOption) (*billingpb.CheckInviteTokenResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ChangeRoleForMerchantUser(ctx context.Context, in *billingpb.ChangeRoleForMerchantUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) ChangeRoleForAdminUser(ctx context.Context, in *billingpb.ChangeRoleForAdminUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetRoleList(ctx context.Context, in *billingpb.GetRoleListRequest, opts ...client.CallOption) (*billingpb.GetRoleListResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeleteMerchantUser(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeleteAdminUser(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) OrderCreateByPaylink(ctx context.Context, in *billingpb.OrderCreateByPaylink, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinks(ctx context.Context, in *billingpb.GetPaylinksRequest, opts ...client.CallOption) (*billingpb.GetPaylinksResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) IncrPaylinkVisits(ctx context.Context, in *billingpb.PaylinkRequestById, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkURL(ctx context.Context, in *billingpb.GetPaylinkURLRequest, opts ...client.CallOption) (*billingpb.GetPaylinkUrlResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) CreateOrUpdatePaylink(ctx context.Context, in *billingpb.CreatePaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeletePaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkStatTotal(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkStatByCountry(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkStatByReferrer(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkStatByDate(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkStatByUtm(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetRecommendedPriceTable(ctx context.Context, in *billingpb.RecommendedPriceTableRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceTableResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) RoyaltyReportPdfUploaded(ctx context.Context, in *billingpb.RoyaltyReportPdfUploadedRequest, opts ...client.CallOption) (*billingpb.RoyaltyReportPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPayoutDocument(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPayoutDocumentRoyaltyReports(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) AutoCreatePayoutDocuments(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetAdminUserRole(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantUserRole(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetCommonUserProfile(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.CommonUserProfileResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) DeleteSavedCard(ctx context.Context, in *billingpb.DeleteSavedCardRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetMerchantOperatingCompany(ctx context.Context, in *billingpb.SetMerchantOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.SetMerchantOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetOperatingCompaniesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompaniesListResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) AddOperatingCompany(ctx context.Context, in *billingpb.OperatingCompany, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaymentMinLimitsSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetPaymentMinLimitsSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SetPaymentMinLimitSystem(ctx context.Context, in *billingpb.PaymentMinLimitSystem, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetOperatingCompany(ctx context.Context, in *billingpb.GetOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetCountriesListForOrder(ctx context.Context, in *billingpb.GetCountriesListForOrderRequest, opts ...client.CallOption) (*billingpb.GetCountriesListForOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetPaylinkTransactions(ctx context.Context, in *billingpb.GetPaylinkTransactionsRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) SendWebhookToMerchant(ctx context.Context, in *billingpb.OrderCreateRequest, opts ...client.CallOption) (*billingpb.SendWebhookToMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) NotifyWebhookTestResults(ctx context.Context, in *billingpb.NotifyWebhookTestResultsRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetAdminByUserId(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) RoyaltyReportFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) PayoutFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}
