package mock

import (
	"context"
	"errors"
	"github.com/micro/go-micro/client"


	"github.com/paysuper/paysuper-proto/go/billingpb"

	"net/http"
)

type BillingServerSystemErrorMock struct{}

func (s *BillingServerSystemErrorMock) GetVatReportTransactions(ctx context.Context, in *billingpb.VatTransactionsRequest, opts ...client.CallOption) (*billingpb.PrivateTransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetVatReport(ctx context.Context, in *billingpb.VatReportRequest, opts ...client.CallOption) (*billingpb.VatReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) OrderReCreateProcess(ctx context.Context, in *billingpb.OrderReCreateProcessRequest, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func NewBillingServerSystemErrorMock() billingpb.BillingService {
	return &BillingServerSystemErrorMock{}
}

func (s *BillingServerSystemErrorMock) GetProductsForOrder(
	ctx context.Context,
	in *billingpb.GetProductsForOrderRequest,
	opts ...client.CallOption,
) (*billingpb.ListProductsResponse, error) {
	return &billingpb.ListProductsResponse{}, nil
}

func (s *BillingServerSystemErrorMock) OrderCreateProcess(
	ctx context.Context,
	in *billingpb.OrderCreateRequest,
	opts ...client.CallOption,
) (*billingpb.OrderCreateProcessResponse, error) {
	return &billingpb.OrderCreateProcessResponse{
		Status:  http.StatusBadRequest,
		Message: SomeError,
	}, nil
}

func (s *BillingServerSystemErrorMock) PaymentFormJsonDataProcess(
	ctx context.Context,
	in *billingpb.PaymentFormJsonDataRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormJsonDataResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) PaymentCreateProcess(
	ctx context.Context,
	in *billingpb.PaymentCreateRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentCreateResponse, error) {
	return &billingpb.PaymentCreateResponse{}, nil
}

func (s *BillingServerSystemErrorMock) PaymentCallbackProcess(
	ctx context.Context,
	in *billingpb.PaymentNotifyRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{}, nil
}

func (s *BillingServerSystemErrorMock) RebuildCache(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerSystemErrorMock) UpdateOrder(
	ctx context.Context,
	in *billingpb.Order,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerSystemErrorMock) UpdateMerchant(
	ctx context.Context,
	in *billingpb.Merchant,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerSystemErrorMock) GetMerchantBy(
	ctx context.Context,
	in *billingpb.GetMerchantByRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantResponse, error) {
	return nil, errors.New("some error")
}

func (s *BillingServerSystemErrorMock) ListMerchants(
	ctx context.Context,
	in *billingpb.MerchantListingRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantListingResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ChangeMerchant(
	ctx context.Context,
	in *billingpb.OnboardingRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ChangeMerchantStatus(
	ctx context.Context,
	in *billingpb.MerchantChangeStatusRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantStatusResponse, error) {
	return &billingpb.ChangeMerchantStatusResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Merchant{},
	}, nil
}

func (s *BillingServerSystemErrorMock) CreateNotification(
	ctx context.Context,
	in *billingpb.NotificationRequest,
	opts ...client.CallOption,
) (*billingpb.CreateNotificationResponse, error) {
	return &billingpb.CreateNotificationResponse{
		Status:  http.StatusBadRequest,
		Message: SomeError,
	}, nil
}

func (s *BillingServerSystemErrorMock) GetNotification(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerSystemErrorMock) ListNotifications(
	ctx context.Context,
	in *billingpb.ListingNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notifications, error) {
	return &billingpb.Notifications{}, nil
}

func (s *BillingServerSystemErrorMock) MarkNotificationAsRead(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerSystemErrorMock) ListMerchantPaymentMethods(
	ctx context.Context,
	in *billingpb.ListMerchantPaymentMethodsRequest,
	opts ...client.CallOption,
) (*billingpb.ListingMerchantPaymentMethod, error) {
	return &billingpb.ListingMerchantPaymentMethod{}, nil
}

func (s *BillingServerSystemErrorMock) GetMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.GetMerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantPaymentMethodResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ChangeMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.MerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantPaymentMethodResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateRefund(
	ctx context.Context,
	in *billingpb.CreateRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ListRefunds(
	ctx context.Context,
	in *billingpb.ListRefundsRequest,
	opts ...client.CallOption,
) (*billingpb.ListRefundsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetRefund(
	ctx context.Context,
	in *billingpb.GetRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ProcessRefundCallback(
	ctx context.Context,
	in *billingpb.CallbackRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) PaymentFormLanguageChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangeLangRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) PaymentFormPaymentAccountChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePaymentAccountRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ProcessBillingAddress(
	ctx context.Context,
	in *billingpb.ProcessBillingAddressRequest,
	opts ...client.CallOption,
) (*billingpb.ProcessBillingAddressResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ChangeMerchantData(
	ctx context.Context,
	in *billingpb.ChangeMerchantDataRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) SetMerchantS3Agreement(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ChangeProject(
	ctx context.Context,
	in *billingpb.Project,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetProject(
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

	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) DeleteProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateToken(
	ctx context.Context,
	in *billingpb.TokenRequest,
	opts ...client.CallOption,
) (*billingpb.TokenResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CheckProjectRequestSignature(
	ctx context.Context,
	in *billingpb.CheckProjectRequestSignatureRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateOrUpdateProduct(ctx context.Context, in *billingpb.Product, opts ...client.CallOption) (*billingpb.Product, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ListProducts(ctx context.Context, in *billingpb.ListProductsRequest, opts ...client.CallOption) (*billingpb.ListProductsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.GetProductResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) DeleteProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) IsOrderCanBePaying(
	ctx context.Context,
	in *billingpb.IsOrderCanBePayingRequest,
	opts ...client.CallOption,
) (*billingpb.IsOrderCanBePayingResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetCountry(ctx context.Context, in *billingpb.GetCountryRequest, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdateCountry(ctx context.Context, in *billingpb.Country, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPriceGroup(ctx context.Context, in *billingpb.GetPriceGroupRequest, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdatePriceGroup(ctx context.Context, in *billingpb.PriceGroup, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetUserNotifySales(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetUserNotifyNewRegion(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetCountriesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.CountriesList, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystemRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystem, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeletePaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchant, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeletePaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystemRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystem, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeleteMoneyBackCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchant, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeleteMoneyBackCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetAllPaymentChannelCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerSystemErrorMock) GetAllPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantListRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerSystemErrorMock) GetAllMoneyBackCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerSystemErrorMock) GetAllMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantListRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantListResponse, error) {
	return nil, errors.New("Some error")
}

func (s *BillingServerSystemErrorMock) CreateOrUpdatePaymentMethodTestSettings(ctx context.Context, in *billingpb.ChangePaymentMethodParamsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeletePaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) FindByZipCode(
	ctx context.Context,
	in *billingpb.FindByZipCodeRequest,
	opts ...client.CallOption,
) (*billingpb.FindByZipCodeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateOrUpdatePaymentMethod(
	ctx context.Context,
	in *billingpb.PaymentMethod,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateOrUpdatePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.ChangePaymentMethodParamsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) DeletePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.GetPaymentMethodSettingsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateAccountingEntry(ctx context.Context, in *billingpb.CreateAccountingEntryRequest, opts ...client.CallOption) (*billingpb.CreateAccountingEntryResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) CreateRoyaltyReport(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ListRoyaltyReports(ctx context.Context, in *billingpb.ListRoyaltyReportsRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeRoyaltyReportStatus(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ListRoyaltyReportOrders(ctx context.Context, in *billingpb.ListRoyaltyReportOrdersRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetVatReportsDashboard(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetVatReportsForCountry(ctx context.Context, in *billingpb.VatReportsRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) CalcAnnualTurnovers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ListProjects(ctx context.Context, in *billingpb.ListProjectsRequest, opts ...client.CallOption) (*billingpb.ListProjectsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ProcessVatReports(ctx context.Context, in *billingpb.ProcessVatReportsRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdateVatReportStatus(ctx context.Context, in *billingpb.UpdateVatReportStatusRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdateProductPrices(
	ctx context.Context,
	in *billingpb.UpdateProductPricesRequest,
	opts ...client.CallOption,
) (*billingpb.ResponseError, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetProductPrices(
	ctx context.Context,
	in *billingpb.RequestProduct,
	opts ...client.CallOption,
) (*billingpb.ProductPricesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPriceGroupRecommendedPrice(
	ctx context.Context,
	in *billingpb.RecommendedPriceRequest,
	opts ...client.CallOption,
) (*billingpb.RecommendedPriceResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPriceGroupCurrencyByRegion(
	ctx context.Context,
	in *billingpb.PriceGroupByRegionRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPriceGroupCurrencies(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPriceGroupByCountry(
	ctx context.Context,
	in *billingpb.PriceGroupByCountryRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroup, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPaymentMethodProductionSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeRoyaltyReport(ctx context.Context, in *billingpb.ChangeRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) AutoAcceptRoyaltyReports(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetUserProfile(
	ctx context.Context,
	in *billingpb.GetUserProfileRequest,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreateOrUpdateUserProfile(
	ctx context.Context,
	in *billingpb.UserProfile,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ConfirmUserEmail(
	ctx context.Context,
	in *billingpb.ConfirmUserEmailRequest,
	opts ...client.CallOption,
) (*billingpb.ConfirmUserEmailResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CreatePageReview(
	ctx context.Context,
	in *billingpb.CreatePageReviewRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) MerchantReviewRoyaltyReport(ctx context.Context, in *billingpb.MerchantReviewRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantOnboardingCompleteData(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantOnboardingCompleteDataResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.GetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantTariffRatesResponse, error) {
	return &billingpb.GetMerchantTariffRatesResponse{}, nil
}

func (s *BillingServerSystemErrorMock) SetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.SetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{}, nil
}

func (s *BillingServerSystemErrorMock) CreateOrUpdateKeyProduct(ctx context.Context, in *billingpb.CreateOrUpdateKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetKeyProducts(ctx context.Context, in *billingpb.ListKeyProductsRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) DeleteKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) PublishKeyProduct(ctx context.Context, in *billingpb.PublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetKeyProductsForOrder(ctx context.Context, in *billingpb.GetKeyProductsForOrderRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetPlatforms(ctx context.Context, in *billingpb.ListPlatformsRequest, opts ...client.CallOption) (*billingpb.ListPlatformsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) DeletePlatformFromProduct(ctx context.Context, in *billingpb.RemovePlatformRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetAvailableKeysCount(ctx context.Context, in *billingpb.GetPlatformKeyCountRequest, opts ...client.CallOption) (*billingpb.GetPlatformKeyCountResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) UploadKeysFile(ctx context.Context, in *billingpb.PlatformKeysFileRequest, opts ...client.CallOption) (*billingpb.PlatformKeysFileResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetKeyByID(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) ReserveKeyForOrder(ctx context.Context, in *billingpb.PlatformKeyReserveRequest, opts ...client.CallOption) (*billingpb.PlatformKeyReserveResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) FinishRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) CancelRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return nil, SomeError
}

func (s *BillingServerSystemErrorMock) GetKeyProductInfo(ctx context.Context, in *billingpb.GetKeyProductInfoRequest, opts ...client.CallOption) (*billingpb.GetKeyProductInfoResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeCodeInOrder(ctx context.Context, in *billingpb.ChangeCodeInOrderRequest, opts ...client.CallOption) (*billingpb.ChangeCodeInOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetDashboardMainReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardMainResponse, error) {
	return &billingpb.GetDashboardMainResponse{}, nil
}
func (s *BillingServerSystemErrorMock) GetDashboardRevenueDynamicsReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardRevenueDynamicsReportResponse, error) {
	return &billingpb.GetDashboardRevenueDynamicsReportResponse{}, nil
}

func (s *BillingServerSystemErrorMock) GetDashboardBaseReport(
	ctx context.Context,
	in *billingpb.GetDashboardBaseReportRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardBaseReportResponse, error) {
	return &billingpb.GetDashboardBaseReportResponse{}, nil
}

func (s *BillingServerSystemErrorMock) GetOrderPublic(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPublicResponse, error) {
	return &billingpb.GetOrderPublicResponse{}, nil
}

func (s *BillingServerSystemErrorMock) GetOrderPrivate(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPrivateResponse, error) {
	return &billingpb.GetOrderPrivateResponse{}, nil
}

func (s *BillingServerSystemErrorMock) FindAllOrdersPublic(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPublicResponse, error) {
	return &billingpb.ListOrdersPublicResponse{}, nil
}

func (s *BillingServerSystemErrorMock) FindAllOrdersPrivate(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPrivateResponse, error) {
	return &billingpb.ListOrdersPrivateResponse{}, nil
}

func (s *BillingServerSystemErrorMock) CreatePayoutDocument(ctx context.Context, in *billingpb.CreatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.CreatePayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdatePayoutDocument(ctx context.Context, in *billingpb.UpdatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPayoutDocuments(ctx context.Context, in *billingpb.GetPayoutDocumentsRequest, opts ...client.CallOption) (*billingpb.GetPayoutDocumentsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UpdatePayoutDocumentSignatures(ctx context.Context, in *billingpb.UpdatePayoutDocumentSignaturesRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantBalance(ctx context.Context, in *billingpb.GetMerchantBalanceRequest, opts ...client.CallOption) (*billingpb.GetMerchantBalanceResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) PayoutDocumentPdfUploaded(ctx context.Context, in *billingpb.PayoutDocumentPdfUploadedRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetRoyaltyReport(ctx context.Context, in *billingpb.GetRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.GetRoyaltyReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) UnPublishKeyProduct(ctx context.Context, in *billingpb.UnPublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) PaymentFormPlatformChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePlatformRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) OrderReceipt(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) OrderReceiptRefund(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetRecommendedPriceByPriceGroup(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetRecommendedPriceByConversion(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) CheckSkuAndKeyProject(ctx context.Context, in *billingpb.CheckSkuAndKeyProjectRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPriceGroupByRegion(ctx context.Context, in *billingpb.GetPriceGroupByRegionRequest, opts ...client.CallOption) (*billingpb.GetPriceGroupByRegionResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantUsers(ctx context.Context, in *billingpb.GetMerchantUsersRequest, opts ...client.CallOption) (*billingpb.GetMerchantUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) FindAllOrders(ctx context.Context, in *billingpb.ListOrdersRequest, opts ...client.CallOption) (*billingpb.ListOrdersResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) OrderCreateByPaylink(ctx context.Context, in *billingpb.OrderCreateByPaylink, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinks(ctx context.Context, in *billingpb.GetPaylinksRequest, opts ...client.CallOption) (*billingpb.GetPaylinksResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) IncrPaylinkVisits(ctx context.Context, in *billingpb.PaylinkRequestById, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkURL(ctx context.Context, in *billingpb.GetPaylinkURLRequest, opts ...client.CallOption) (*billingpb.GetPaylinkUrlResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) CreateOrUpdatePaylink(ctx context.Context, in *billingpb.CreatePaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeletePaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkStatTotal(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkStatByCountry(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkStatByReferrer(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkStatByDate(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkStatByUtm(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetRecommendedPriceTable(ctx context.Context, in *billingpb.RecommendedPriceTableRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceTableResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) RoyaltyReportPdfUploaded(ctx context.Context, in *billingpb.RoyaltyReportPdfUploadedRequest, opts ...client.CallOption) (*billingpb.RoyaltyReportPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPayoutDocument(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPayoutDocumentRoyaltyReports(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) AutoCreatePayoutDocuments(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetAdminUsers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetAdminUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantsForUser(ctx context.Context, in *billingpb.GetMerchantsForUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantsForUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) InviteUserMerchant(ctx context.Context, in *billingpb.InviteUserMerchantRequest, opts ...client.CallOption) (*billingpb.InviteUserMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) InviteUserAdmin(ctx context.Context, in *billingpb.InviteUserAdminRequest, opts ...client.CallOption) (*billingpb.InviteUserAdminResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ResendInviteMerchant(ctx context.Context, in *billingpb.ResendInviteMerchantRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ResendInviteAdmin(ctx context.Context, in *billingpb.ResendInviteAdminRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantUser(ctx context.Context, in *billingpb.GetMerchantUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetAdminUser(ctx context.Context, in *billingpb.GetAdminUserRequest, opts ...client.CallOption) (*billingpb.GetAdminUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) AcceptInvite(ctx context.Context, in *billingpb.AcceptInviteRequest, opts ...client.CallOption) (*billingpb.AcceptInviteResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) CheckInviteToken(ctx context.Context, in *billingpb.CheckInviteTokenRequest, opts ...client.CallOption) (*billingpb.CheckInviteTokenResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeRoleForMerchantUser(ctx context.Context, in *billingpb.ChangeRoleForMerchantUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeRoleForAdminUser(ctx context.Context, in *billingpb.ChangeRoleForAdminUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetRoleList(ctx context.Context, in *billingpb.GetRoleListRequest, opts ...client.CallOption) (*billingpb.GetRoleListResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) ChangeMerchantManualPayouts(ctx context.Context, in *billingpb.ChangeMerchantManualPayoutsRequest, opts ...client.CallOption) (*billingpb.ChangeMerchantManualPayoutsResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeleteMerchantUser(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeleteAdminUser(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetAdminUserRole(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetMerchantUserRole(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetCommonUserProfile(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.CommonUserProfileResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) DeleteSavedCard(ctx context.Context, in *billingpb.DeleteSavedCardRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetMerchantOperatingCompany(ctx context.Context, in *billingpb.SetMerchantOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.SetMerchantOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetOperatingCompaniesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompaniesListResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) AddOperatingCompany(ctx context.Context, in *billingpb.OperatingCompany, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaymentMinLimitsSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetPaymentMinLimitsSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) SetPaymentMinLimitSystem(ctx context.Context, in *billingpb.PaymentMinLimitSystem, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetOperatingCompany(ctx context.Context, in *billingpb.GetOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetCountriesListForOrder(ctx context.Context, in *billingpb.GetCountriesListForOrderRequest, opts ...client.CallOption) (*billingpb.GetCountriesListForOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerSystemErrorMock) GetPaylinkTransactions(ctx context.Context, in *billingpb.GetPaylinkTransactionsRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}
