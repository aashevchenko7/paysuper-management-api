package mock

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/micro/go-micro/client"

	"github.com/paysuper/paysuper-proto/go/billingpb"
)

type BillingServerOkTemporaryMock struct{}

func (s *BillingServerOkTemporaryMock) GetDashboardCustomersReport(ctx context.Context, in *billingpb.DashboardCustomerReportRequest, opts ...client.CallOption) (*billingpb.GetDashboardCustomerReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetDashboardCustomerArpu(ctx context.Context, in *billingpb.DashboardCustomerReportArpuRequest, opts ...client.CallOption) (*billingpb.DashboardCustomerReportArpuResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetCustomerList(ctx context.Context, in *billingpb.ListCustomersRequest, opts ...client.CallOption) (*billingpb.ListCustomersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetCustomerInfo(ctx context.Context, in *billingpb.GetCustomerInfoRequest, opts ...client.CallOption) (*billingpb.GetCustomerInfoResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetVatReportTransactions(ctx context.Context, in *billingpb.VatTransactionsRequest, opts ...client.CallOption) (*billingpb.PrivateTransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetVatReport(ctx context.Context, in *billingpb.VatReportRequest, opts ...client.CallOption) (*billingpb.VatReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) OrderReCreateProcess(ctx context.Context, in *billingpb.OrderReCreateProcessRequest, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func NewBillingServerOkTemporaryMock() billingpb.BillingService {
	return &BillingServerOkTemporaryMock{}
}

func (s *BillingServerOkTemporaryMock) GetProductsForOrder(
	ctx context.Context,
	in *billingpb.GetProductsForOrderRequest,
	opts ...client.CallOption,
) (*billingpb.ListProductsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) OrderCreateProcess(
	ctx context.Context,
	in *billingpb.OrderCreateRequest,
	opts ...client.CallOption,
) (*billingpb.OrderCreateProcessResponse, error) {
	return &billingpb.OrderCreateProcessResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Order{},
	}, nil
}

func (s *BillingServerOkTemporaryMock) PaymentFormJsonDataProcess(
	ctx context.Context,
	in *billingpb.PaymentFormJsonDataRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormJsonDataResponse, error) {
	return &billingpb.PaymentFormJsonDataResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) PaymentCreateProcess(
	ctx context.Context,
	in *billingpb.PaymentCreateRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentCreateResponse, error) {
	return &billingpb.PaymentCreateResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) PaymentCallbackProcess(
	ctx context.Context,
	in *billingpb.PaymentNotifyRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) RebuildCache(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) UpdateOrder(
	ctx context.Context,
	in *billingpb.Order,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) UpdateMerchant(
	ctx context.Context,
	in *billingpb.Merchant,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetMerchantBy(
	ctx context.Context,
	in *billingpb.GetMerchantByRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantResponse, error) {
	rsp := &billingpb.GetMerchantResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: &billingpb.ResponseErrorMessage{},
		Item:    OnboardingMerchantMock,
	}

	return rsp, nil
}

func (s *BillingServerOkTemporaryMock) ListMerchants(
	ctx context.Context,
	in *billingpb.MerchantListingRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantListingResponse, error) {
	return &billingpb.MerchantListingResponse{
		Count: 3,
		Items: []*billingpb.Merchant{OnboardingMerchantMock, OnboardingMerchantMock, OnboardingMerchantMock},
	}, nil
}

func (s *BillingServerOkTemporaryMock) ChangeMerchant(
	ctx context.Context,
	in *billingpb.OnboardingRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantResponse, error) {
	m := &billingpb.Merchant{
		Company:  in.Company,
		Contacts: in.Contacts,
		Banking: &billingpb.MerchantBanking{
			Currency:      "RUB",
			Name:          in.Banking.Name,
			Address:       in.Banking.Address,
			AccountNumber: in.Banking.AccountNumber,
			Swift:         in.Banking.Swift,
			Details:       in.Banking.Details,
		},
		Status: billingpb.MerchantStatusDraft,
	}

	if in.Id != "" {
		m.Id = in.Id
	} else {
		m.Id = bson.NewObjectId().Hex()
	}

	return &billingpb.ChangeMerchantResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   m,
	}, nil
}

func (s *BillingServerOkTemporaryMock) ChangeMerchantStatus(
	ctx context.Context,
	in *billingpb.MerchantChangeStatusRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantStatusResponse, error) {
	return &billingpb.ChangeMerchantStatusResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Merchant{Id: in.MerchantId, Status: in.Status},
	}, nil
}

func (s *BillingServerOkTemporaryMock) CreateNotification(
	ctx context.Context,
	in *billingpb.NotificationRequest,
	opts ...client.CallOption,
) (*billingpb.CreateNotificationResponse, error) {
	return &billingpb.CreateNotificationResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Notification{},
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetNotification(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerOkTemporaryMock) ListNotifications(
	ctx context.Context,
	in *billingpb.ListingNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notifications, error) {
	return &billingpb.Notifications{}, nil
}

func (s *BillingServerOkTemporaryMock) MarkNotificationAsRead(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerOkTemporaryMock) ListMerchantPaymentMethods(
	ctx context.Context,
	in *billingpb.ListMerchantPaymentMethodsRequest,
	opts ...client.CallOption,
) (*billingpb.ListingMerchantPaymentMethod, error) {
	return &billingpb.ListingMerchantPaymentMethod{}, nil
}

func (s *BillingServerOkTemporaryMock) GetMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.GetMerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantPaymentMethodResponse, error) {
	return &billingpb.GetMerchantPaymentMethodResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) ChangeMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.MerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantPaymentMethodResponse, error) {
	return &billingpb.MerchantPaymentMethodResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.MerchantPaymentMethod{
			PaymentMethod: &billingpb.MerchantPaymentMethodIdentification{
				Id:   in.PaymentMethod.Id,
				Name: in.PaymentMethod.Name,
			},
			Commission:  in.Commission,
			Integration: in.Integration,
			IsActive:    in.IsActive,
		},
	}, nil
}

func (s *BillingServerOkTemporaryMock) CreateRefund(
	ctx context.Context,
	in *billingpb.CreateRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Refund{},
	}, nil
}

func (s *BillingServerOkTemporaryMock) ListRefunds(
	ctx context.Context,
	in *billingpb.ListRefundsRequest,
	opts ...client.CallOption,
) (*billingpb.ListRefundsResponse, error) {
	return &billingpb.ListRefundsResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetRefund(
	ctx context.Context,
	in *billingpb.GetRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Refund{},
	}, nil
}

func (s *BillingServerOkTemporaryMock) ProcessRefundCallback(
	ctx context.Context,
	in *billingpb.CallbackRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{
		Status: billingpb.ResponseStatusOk,
		Error:  SomeError.Message,
	}, nil
}

func (s *BillingServerOkTemporaryMock) ChangeProject(
	ctx context.Context,
	in *billingpb.Project,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) DeleteProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return &billingpb.ChangeProjectResponse{
		Status:  billingpb.ResponseStatusBadData,
		Message: SomeError,
	}, nil
}

func (s *BillingServerOkTemporaryMock) CreateToken(
	ctx context.Context,
	in *billingpb.TokenRequest,
	opts ...client.CallOption,
) (*billingpb.TokenResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CheckProjectRequestSignature(
	ctx context.Context,
	in *billingpb.CheckProjectRequestSignatureRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdateProduct(ctx context.Context, in *billingpb.Product, opts ...client.CallOption) (*billingpb.Product, error) {
	return Product, nil
}

func (s *BillingServerOkTemporaryMock) ListProducts(ctx context.Context, in *billingpb.ListProductsRequest, opts ...client.CallOption) (*billingpb.ListProductsResponse, error) {
	return &billingpb.ListProductsResponse{
		Limit:  1,
		Offset: 0,
		Total:  200,
		Products: []*billingpb.Product{
			Product,
		},
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.GetProductResponse, error) {
	return GetProductResponse, nil
}

func (s *BillingServerOkTemporaryMock) DeleteProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) PaymentFormLanguageChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangeLangRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) PaymentFormPaymentAccountChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePaymentAccountRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) ProcessBillingAddress(
	ctx context.Context,
	in *billingpb.ProcessBillingAddressRequest,
	opts ...client.CallOption,
) (*billingpb.ProcessBillingAddressResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) ChangeMerchantData(
	ctx context.Context,
	in *billingpb.ChangeMerchantDataRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return &billingpb.ChangeMerchantDataResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) SetMerchantS3Agreement(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	return &billingpb.ChangeMerchantDataResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) ListProjects(ctx context.Context, in *billingpb.ListProjectsRequest, opts ...client.CallOption) (*billingpb.ListProjectsResponse, error) {
	return &billingpb.ListProjectsResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetOrder(ctx context.Context, in *billingpb.GetOrderRequest, opts ...client.CallOption) (*billingpb.Order, error) {
	return &billingpb.Order{}, nil
}

func (s *BillingServerOkTemporaryMock) IsOrderCanBePaying(
	ctx context.Context,
	in *billingpb.IsOrderCanBePayingRequest,
	opts ...client.CallOption,
) (*billingpb.IsOrderCanBePayingResponse, error) {
	return &billingpb.IsOrderCanBePayingResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Order{},
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetCountry(ctx context.Context, in *billingpb.GetCountryRequest, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdateCountry(ctx context.Context, in *billingpb.Country, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPriceGroup(ctx context.Context, in *billingpb.GetPriceGroupRequest, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdatePriceGroup(ctx context.Context, in *billingpb.PriceGroup, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetUserNotifySales(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetUserNotifyNewRegion(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}
func (s *BillingServerOkTemporaryMock) GetCountriesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.CountriesList, error) {
	panic("implement me")
}
func (s *BillingServerOkTemporaryMock) GetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystemRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystem, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeletePaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchant, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeletePaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystemRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystem, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteMoneyBackCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchant, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteMoneyBackCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAllPaymentChannelCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemListResponse, error) {
	return &billingpb.PaymentChannelCostSystemListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetAllPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantListRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantListResponse, error) {
	return &billingpb.PaymentChannelCostMerchantListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetAllMoneyBackCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemListResponse, error) {
	return &billingpb.MoneyBackCostSystemListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkTemporaryMock) GetAllMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantListRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantListResponse, error) {
	return &billingpb.MoneyBackCostMerchantListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdatePaymentMethodTestSettings(ctx context.Context, in *billingpb.ChangePaymentMethodParamsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeletePaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) FindByZipCode(
	ctx context.Context,
	in *billingpb.FindByZipCodeRequest,
	opts ...client.CallOption,
) (*billingpb.FindByZipCodeResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdatePaymentMethod(
	ctx context.Context,
	in *billingpb.PaymentMethod,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdatePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.ChangePaymentMethodParamsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) DeletePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.GetPaymentMethodSettingsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreateAccountingEntry(ctx context.Context, in *billingpb.CreateAccountingEntryRequest, opts ...client.CallOption) (*billingpb.CreateAccountingEntryResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CreateRoyaltyReport(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ListRoyaltyReports(ctx context.Context, in *billingpb.ListRoyaltyReportsRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeRoyaltyReportStatus(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ListRoyaltyReportOrders(ctx context.Context, in *billingpb.ListRoyaltyReportOrdersRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetVatReportsDashboard(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetVatReportsForCountry(ctx context.Context, in *billingpb.VatReportsRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CalcAnnualTurnovers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}
func (s *BillingServerOkTemporaryMock) ProcessVatReports(ctx context.Context, in *billingpb.ProcessVatReportsRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdateVatReportStatus(ctx context.Context, in *billingpb.UpdateVatReportStatusRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdateProductPrices(
	ctx context.Context,
	in *billingpb.UpdateProductPricesRequest,
	opts ...client.CallOption,
) (*billingpb.ResponseError, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetProductPrices(
	ctx context.Context,
	in *billingpb.RequestProduct,
	opts ...client.CallOption,
) (*billingpb.ProductPricesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPriceGroupRecommendedPrice(
	ctx context.Context,
	in *billingpb.RecommendedPriceRequest,
	opts ...client.CallOption,
) (*billingpb.RecommendedPriceResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPriceGroupCurrencyByRegion(
	ctx context.Context,
	in *billingpb.PriceGroupByRegionRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPriceGroupCurrencies(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPriceGroupByCountry(
	ctx context.Context,
	in *billingpb.PriceGroupByCountryRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroup, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPaymentMethodProductionSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetPaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeRoyaltyReport(ctx context.Context, in *billingpb.ChangeRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) AutoAcceptRoyaltyReports(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetUserProfile(
	ctx context.Context,
	in *billingpb.GetUserProfileRequest,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdateUserProfile(
	ctx context.Context,
	in *billingpb.UserProfile,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) ConfirmUserEmail(
	ctx context.Context,
	in *billingpb.ConfirmUserEmailRequest,
	opts ...client.CallOption,
) (*billingpb.ConfirmUserEmailResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) CreatePageReview(
	ctx context.Context,
	in *billingpb.CreatePageReviewRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) MerchantReviewRoyaltyReport(ctx context.Context, in *billingpb.MerchantReviewRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantOnboardingCompleteData(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantOnboardingCompleteDataResponse, error) {
	return nil, SomeError
}

func (s *BillingServerOkTemporaryMock) GetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.GetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantTariffRatesResponse, error) {
	return &billingpb.GetMerchantTariffRatesResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) SetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.SetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdateKeyProduct(ctx context.Context, in *billingpb.CreateOrUpdateKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetKeyProducts(ctx context.Context, in *billingpb.ListKeyProductsRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) PublishKeyProduct(ctx context.Context, in *billingpb.PublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetKeyProductsForOrder(ctx context.Context, in *billingpb.GetKeyProductsForOrderRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPlatforms(ctx context.Context, in *billingpb.ListPlatformsRequest, opts ...client.CallOption) (*billingpb.ListPlatformsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeletePlatformFromProduct(ctx context.Context, in *billingpb.RemovePlatformRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAvailableKeysCount(ctx context.Context, in *billingpb.GetPlatformKeyCountRequest, opts ...client.CallOption) (*billingpb.GetPlatformKeyCountResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UploadKeysFile(ctx context.Context, in *billingpb.PlatformKeysFileRequest, opts ...client.CallOption) (*billingpb.PlatformKeysFileResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetKeyByID(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ReserveKeyForOrder(ctx context.Context, in *billingpb.PlatformKeyReserveRequest, opts ...client.CallOption) (*billingpb.PlatformKeyReserveResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) FinishRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CancelRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetKeyProductInfo(ctx context.Context, in *billingpb.GetKeyProductInfoRequest, opts ...client.CallOption) (*billingpb.GetKeyProductInfoResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeCodeInOrder(ctx context.Context, in *billingpb.ChangeCodeInOrderRequest, opts ...client.CallOption) (*billingpb.ChangeCodeInOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetDashboardMainReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardMainResponse, error) {
	return &billingpb.GetDashboardMainResponse{}, nil
}
func (s *BillingServerOkTemporaryMock) GetDashboardRevenueDynamicsReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardRevenueDynamicsReportResponse, error) {
	return &billingpb.GetDashboardRevenueDynamicsReportResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetDashboardBaseReport(
	ctx context.Context,
	in *billingpb.GetDashboardBaseReportRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardBaseReportResponse, error) {
	return &billingpb.GetDashboardBaseReportResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetOrderPublic(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPublicResponse, error) {
	return &billingpb.GetOrderPublicResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) GetOrderPrivate(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPrivateResponse, error) {
	return &billingpb.GetOrderPrivateResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) FindAllOrdersPublic(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPublicResponse, error) {
	return &billingpb.ListOrdersPublicResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) FindAllOrdersPrivate(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPrivateResponse, error) {
	return &billingpb.ListOrdersPrivateResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) CreatePayoutDocument(ctx context.Context, in *billingpb.CreatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.CreatePayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdatePayoutDocument(ctx context.Context, in *billingpb.UpdatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPayoutDocuments(ctx context.Context, in *billingpb.GetPayoutDocumentsRequest, opts ...client.CallOption) (*billingpb.GetPayoutDocumentsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UpdatePayoutDocumentSignatures(ctx context.Context, in *billingpb.UpdatePayoutDocumentSignaturesRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantBalance(ctx context.Context, in *billingpb.GetMerchantBalanceRequest, opts ...client.CallOption) (*billingpb.GetMerchantBalanceResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) PayoutDocumentPdfUploaded(ctx context.Context, in *billingpb.PayoutDocumentPdfUploadedRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetRoyaltyReport(ctx context.Context, in *billingpb.GetRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.GetRoyaltyReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) UnPublishKeyProduct(ctx context.Context, in *billingpb.UnPublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) PaymentFormPlatformChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePlatformRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) OrderReceipt(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) OrderReceiptRefund(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetRecommendedPriceByPriceGroup(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetRecommendedPriceByConversion(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CheckSkuAndKeyProject(ctx context.Context, in *billingpb.CheckSkuAndKeyProjectRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPriceGroupByRegion(ctx context.Context, in *billingpb.GetPriceGroupByRegionRequest, opts ...client.CallOption) (*billingpb.GetPriceGroupByRegionResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantUsers(ctx context.Context, in *billingpb.GetMerchantUsersRequest, opts ...client.CallOption) (*billingpb.GetMerchantUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) FindAllOrders(ctx context.Context, in *billingpb.ListOrdersRequest, opts ...client.CallOption) (*billingpb.ListOrdersResponse, error) {
	return &billingpb.ListOrdersResponse{}, nil
}

func (s *BillingServerOkTemporaryMock) OrderCreateByPaylink(ctx context.Context, in *billingpb.OrderCreateByPaylink, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinks(ctx context.Context, in *billingpb.GetPaylinksRequest, opts ...client.CallOption) (*billingpb.GetPaylinksResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) IncrPaylinkVisits(ctx context.Context, in *billingpb.PaylinkRequestById, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkURL(ctx context.Context, in *billingpb.GetPaylinkURLRequest, opts ...client.CallOption) (*billingpb.GetPaylinkUrlResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CreateOrUpdatePaylink(ctx context.Context, in *billingpb.CreatePaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeletePaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkStatTotal(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkStatByCountry(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkStatByReferrer(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkStatByDate(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkStatByUtm(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetRecommendedPriceTable(ctx context.Context, in *billingpb.RecommendedPriceTableRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceTableResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) RoyaltyReportPdfUploaded(ctx context.Context, in *billingpb.RoyaltyReportPdfUploadedRequest, opts ...client.CallOption) (*billingpb.RoyaltyReportPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPayoutDocument(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPayoutDocumentRoyaltyReports(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) AutoCreatePayoutDocuments(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAdminUsers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetAdminUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantsForUser(ctx context.Context, in *billingpb.GetMerchantsForUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantsForUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) InviteUserMerchant(ctx context.Context, in *billingpb.InviteUserMerchantRequest, opts ...client.CallOption) (*billingpb.InviteUserMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) InviteUserAdmin(ctx context.Context, in *billingpb.InviteUserAdminRequest, opts ...client.CallOption) (*billingpb.InviteUserAdminResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ResendInviteMerchant(ctx context.Context, in *billingpb.ResendInviteMerchantRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ResendInviteAdmin(ctx context.Context, in *billingpb.ResendInviteAdminRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantUser(ctx context.Context, in *billingpb.GetMerchantUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAdminUser(ctx context.Context, in *billingpb.GetAdminUserRequest, opts ...client.CallOption) (*billingpb.GetAdminUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) AcceptInvite(ctx context.Context, in *billingpb.AcceptInviteRequest, opts ...client.CallOption) (*billingpb.AcceptInviteResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) CheckInviteToken(ctx context.Context, in *billingpb.CheckInviteTokenRequest, opts ...client.CallOption) (*billingpb.CheckInviteTokenResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeRoleForMerchantUser(ctx context.Context, in *billingpb.ChangeRoleForMerchantUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeRoleForAdminUser(ctx context.Context, in *billingpb.ChangeRoleForAdminUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetRoleList(ctx context.Context, in *billingpb.GetRoleListRequest, opts ...client.CallOption) (*billingpb.GetRoleListResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) ChangeMerchantManualPayouts(ctx context.Context, in *billingpb.ChangeMerchantManualPayoutsRequest, opts ...client.CallOption) (*billingpb.ChangeMerchantManualPayoutsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteMerchantUser(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteAdminUser(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAdminUserRole(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantUserRole(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetCommonUserProfile(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.CommonUserProfileResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) DeleteSavedCard(ctx context.Context, in *billingpb.DeleteSavedCardRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetMerchantOperatingCompany(ctx context.Context, in *billingpb.SetMerchantOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.SetMerchantOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetOperatingCompaniesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompaniesListResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) AddOperatingCompany(ctx context.Context, in *billingpb.OperatingCompany, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaymentMinLimitsSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetPaymentMinLimitsSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetPaymentMinLimitSystem(ctx context.Context, in *billingpb.PaymentMinLimitSystem, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetOperatingCompany(ctx context.Context, in *billingpb.GetOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SendWebhookToMerchant(ctx context.Context, in *billingpb.OrderCreateRequest, opts ...client.CallOption) (*billingpb.SendWebhookToMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) NotifyWebhookTestResults(ctx context.Context, in *billingpb.NotifyWebhookTestResultsRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}
func (s *BillingServerOkTemporaryMock) GetCountriesListForOrder(ctx context.Context, in *billingpb.GetCountriesListForOrderRequest, opts ...client.CallOption) (*billingpb.GetCountriesListForOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetPaylinkTransactions(ctx context.Context, in *billingpb.GetPaylinkTransactionsRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetAdminByUserId(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) RoyaltyReportFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) PayoutFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetActOfCompletion(ctx context.Context, in *billingpb.ActOfCompletionRequest, opts ...client.CallOption) (*billingpb.ActOfCompletionResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) SetCustomerPaymentActivity(ctx context.Context, in *billingpb.SetCustomerPaymentActivityRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetDashboardCustomersReport(ctx context.Context, in *billingpb.DashboardCustomerReportRequest, opts ...client.CallOption) (*billingpb.GetDashboardCustomerReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetDashboardCustomerArpu(ctx context.Context, in *billingpb.DashboardCustomerReportArpuRequest, opts ...client.CallOption) (*billingpb.DashboardCustomerReportArpuResponse, error) {
	panic("implement me")
}
