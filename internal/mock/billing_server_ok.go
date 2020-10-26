package mock

import (
	"context"
	"github.com/bxcodec/faker"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"reflect"

	"github.com/paysuper/paysuper-proto/go/billingpb"

	"net/http"
)

type BillingServerOkMock struct {
}

func (s *BillingServerOkMock) DeleteRecurringSubscription(ctx context.Context, in *billingpb.DeleteRecurringSubscriptionRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCustomerShortInfo(ctx context.Context, in *billingpb.GetCustomerShortInfoRequest, opts ...client.CallOption) (*billingpb.GetCustomerShortInfoResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetSubscriptionOrders(ctx context.Context, in *billingpb.GetSubscriptionOrdersRequest, opts ...client.CallOption) (*billingpb.GetSubscriptionOrdersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetSubscription(ctx context.Context, in *billingpb.GetSubscriptionRequest, opts ...client.CallOption) (*billingpb.GetSubscriptionResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) FindSubscriptions(ctx context.Context, in *billingpb.FindSubscriptionsRequest, opts ...client.CallOption) (*billingpb.FindSubscriptionsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeleteCustomerCard(ctx context.Context, in *billingpb.DeleteCustomerCardRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetDashboardCustomersReport(ctx context.Context, in *billingpb.DashboardCustomerReportRequest, opts ...client.CallOption) (*billingpb.GetDashboardCustomerReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetDashboardCustomerArpu(ctx context.Context, in *billingpb.DashboardCustomerReportArpuRequest, opts ...client.CallOption) (*billingpb.DashboardCustomerReportArpuResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCustomerList(ctx context.Context, in *billingpb.ListCustomersRequest, opts ...client.CallOption) (*billingpb.ListCustomersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCustomerInfo(ctx context.Context, in *billingpb.GetCustomerInfoRequest, opts ...client.CallOption) (*billingpb.GetCustomerInfoResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetVatReportTransactions(ctx context.Context, in *billingpb.VatTransactionsRequest, opts ...client.CallOption) (*billingpb.PrivateTransactionsResponse, error) {
	return &billingpb.PrivateTransactionsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.PrivateTransactionsPaginate{
			Count: 100,
			Items: []*billingpb.OrderViewPrivate{},
		},
	}, nil
}

func (s *BillingServerOkMock) GetVatReport(ctx context.Context, in *billingpb.VatReportRequest, opts ...client.CallOption) (*billingpb.VatReportResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) OrderReCreateProcess(ctx context.Context, in *billingpb.OrderReCreateProcessRequest, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func NewBillingServerOkMock() billingpb.BillingService {
	_ = faker.AddProvider("objectIdString", func(_ reflect.Value) (interface{}, error) {
		return "ffffffffffffffffffffffff", nil
	})

	return &BillingServerOkMock{}
}

func (s *BillingServerOkMock) GetProductsForOrder(
	ctx context.Context,
	in *billingpb.GetProductsForOrderRequest,
	opts ...client.CallOption,
) (*billingpb.ListProductsResponse, error) {
	return &billingpb.ListProductsResponse{}, nil
}

func (s *BillingServerOkMock) OrderCreateProcess(
	ctx context.Context,
	in *billingpb.OrderCreateRequest,
	opts ...client.CallOption,
) (*billingpb.OrderCreateProcessResponse, error) {
	return &billingpb.OrderCreateProcessResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.Order{
			Uuid: uuid.New().String(),
		},
	}, nil
}

func (s *BillingServerOkMock) PaymentFormJsonDataProcess(
	ctx context.Context,
	in *billingpb.PaymentFormJsonDataRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormJsonDataResponse, error) {
	cookie := in.Cookie
	if cookie == "" {
		cookie = bson.NewObjectId().Hex()
	}
	return &billingpb.PaymentFormJsonDataResponse{
		Status: billingpb.ResponseStatusOk,
		Cookie: cookie,
		Item:   &billingpb.PaymentFormJsonData{},
	}, nil
}

func (s *BillingServerOkMock) PaymentCreateProcess(
	ctx context.Context,
	in *billingpb.PaymentCreateRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentCreateResponse, error) {
	return &billingpb.PaymentCreateResponse{}, nil
}

func (s *BillingServerOkMock) PaymentCallbackProcess(
	ctx context.Context,
	in *billingpb.PaymentNotifyRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{}, nil
}

func (s *BillingServerOkMock) RebuildCache(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkMock) UpdateOrder(
	ctx context.Context,
	in *billingpb.Order,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkMock) UpdateMerchant(
	ctx context.Context,
	in *billingpb.Merchant,
	opts ...client.CallOption,
) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkMock) GetMerchantBy(
	ctx context.Context,
	in *billingpb.GetMerchantByRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantResponse, error) {
	if in.MerchantId == OnboardingMerchantMock.Id {
		OnboardingMerchantMock.S3AgreementName = SomeAgreementName
	} else if in.MerchantId == "ffffffffffffffffffffffff" {
		OnboardingMerchantMock.S3AgreementName = SomeAgreementName1
	} else {
		if in.MerchantId == SomeMerchantId1 {
			OnboardingMerchantMock.S3AgreementName = SomeAgreementName1
		} else {
			if in.MerchantId == SomeMerchantId2 {
				OnboardingMerchantMock.S3AgreementName = SomeAgreementName2
			} else {
				OnboardingMerchantMock.S3AgreementName = ""
			}
		}
	}

	if in.MerchantId == SomeMerchantId3 {
		OnboardingMerchantMock.Status = billingpb.MerchantStatusDraft
	} else {
		OnboardingMerchantMock.Status = billingpb.MerchantStatusAgreementSigning
	}

	rsp := &billingpb.GetMerchantResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: &billingpb.ResponseErrorMessage{},
		Item:    OnboardingMerchantMock,
	}

	return rsp, nil
}

func (s *BillingServerOkMock) ListMerchants(
	ctx context.Context,
	in *billingpb.MerchantListingRequest,
	opts ...client.CallOption,
) (*billingpb.MerchantListingResponse, error) {
	return &billingpb.MerchantListingResponse{
		Count: 3,
		Items: []*billingpb.Merchant{OnboardingMerchantMock, OnboardingMerchantMock, OnboardingMerchantMock},
	}, nil
}

func (s *BillingServerOkMock) ChangeMerchant(
	ctx context.Context,
	in *billingpb.OnboardingRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantResponse, error) {
	m := &billingpb.Merchant{
		User: &billingpb.MerchantUser{
			Id:    bson.NewObjectId().Hex(),
			Email: "test@unit.test",
		},
		Company:  in.Company,
		Contacts: in.Contacts,
		Banking:  in.Banking,
		Status:   billingpb.MerchantStatusDraft,
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

func (s *BillingServerOkMock) ChangeMerchantStatus(
	ctx context.Context,
	in *billingpb.MerchantChangeStatusRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantStatusResponse, error) {
	return &billingpb.ChangeMerchantStatusResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Merchant{Id: in.MerchantId, Status: in.Status},
	}, nil
}

func (s *BillingServerOkMock) CreateNotification(
	ctx context.Context,
	in *billingpb.NotificationRequest,
	opts ...client.CallOption,
) (*billingpb.CreateNotificationResponse, error) {
	return &billingpb.CreateNotificationResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetNotification(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerOkMock) ListNotifications(
	ctx context.Context,
	in *billingpb.ListingNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notifications, error) {
	return &billingpb.Notifications{}, nil
}

func (s *BillingServerOkMock) MarkNotificationAsRead(
	ctx context.Context,
	in *billingpb.GetNotificationRequest,
	opts ...client.CallOption,
) (*billingpb.Notification, error) {
	return &billingpb.Notification{}, nil
}

func (s *BillingServerOkMock) ListMerchantPaymentMethods(
	ctx context.Context,
	in *billingpb.ListMerchantPaymentMethodsRequest,
	opts ...client.CallOption,
) (*billingpb.ListingMerchantPaymentMethod, error) {
	return &billingpb.ListingMerchantPaymentMethod{}, nil
}

func (s *BillingServerOkMock) GetMerchantPaymentMethod(
	ctx context.Context,
	in *billingpb.GetMerchantPaymentMethodRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantPaymentMethodResponse, error) {
	return &billingpb.GetMerchantPaymentMethodResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) ChangeMerchantPaymentMethod(
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

func (s *BillingServerOkMock) CreateRefund(
	ctx context.Context,
	in *billingpb.CreateRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.Refund{
			Id:            bson.NewObjectId().Hex(),
			OriginalOrder: &billingpb.RefundOrder{Id: bson.NewObjectId().Hex(), Uuid: uuid.New().String()},
			ExternalId:    "",
			Amount:        10,
			CreatorId:     "",
			Reason:        SomeError.Message,
			Currency:      "RUB",
			Status:        0,
		},
	}, nil
}

func (s *BillingServerOkMock) ListRefunds(
	ctx context.Context,
	in *billingpb.ListRefundsRequest,
	opts ...client.CallOption,
) (*billingpb.ListRefundsResponse, error) {
	return &billingpb.ListRefundsResponse{
		Count: 2,
		Items: []*billingpb.Refund{
			{
				Id:            bson.NewObjectId().Hex(),
				OriginalOrder: &billingpb.RefundOrder{Id: bson.NewObjectId().Hex(), Uuid: uuid.New().String()},
				ExternalId:    "",
				Amount:        10,
				CreatorId:     "",
				Reason:        SomeError.Message,
				Currency:      "RUB",
				Status:        0,
			},
			{
				Id:            bson.NewObjectId().Hex(),
				OriginalOrder: &billingpb.RefundOrder{Id: bson.NewObjectId().Hex(), Uuid: uuid.New().String()},
				ExternalId:    "",
				Amount:        10,
				CreatorId:     "",
				Reason:        SomeError.Message,
				Currency:      "RUB",
				Status:        0,
			},
		},
	}, nil
}

func (s *BillingServerOkMock) GetRefund(
	ctx context.Context,
	in *billingpb.GetRefundRequest,
	opts ...client.CallOption,
) (*billingpb.CreateRefundResponse, error) {
	return &billingpb.CreateRefundResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.Refund{
			Id:            bson.NewObjectId().Hex(),
			OriginalOrder: &billingpb.RefundOrder{Id: bson.NewObjectId().Hex(), Uuid: uuid.New().String()},
			ExternalId:    "",
			Amount:        10,
			CreatorId:     "",
			Reason:        SomeError.Message,
			Currency:      "RUB",
			Status:        0,
		},
	}, nil
}

func (s *BillingServerOkMock) ProcessRefundCallback(
	ctx context.Context,
	in *billingpb.CallbackRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentNotifyResponse, error) {
	return &billingpb.PaymentNotifyResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) PaymentFormLanguageChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangeLangRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return &billingpb.PaymentFormDataChangeResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.PaymentFormDataChangeResponseItem{
			UserAddressDataRequired: true,
			UserIpData: &billingpb.UserIpData{
				Country: "RU",
				City:    "St.Petersburg",
				Zip:     "190000",
			},
		},
	}, nil
}

func (s *BillingServerOkMock) PaymentFormPaymentAccountChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePaymentAccountRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	return &billingpb.PaymentFormDataChangeResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.PaymentFormDataChangeResponseItem{
			UserAddressDataRequired: true,
			UserIpData: &billingpb.UserIpData{
				Country: "RU",
				City:    "St.Petersburg",
				Zip:     "190000",
			},
		},
	}, nil
}

func (s *BillingServerOkMock) ProcessBillingAddress(
	ctx context.Context,
	in *billingpb.ProcessBillingAddressRequest,
	opts ...client.CallOption,
) (*billingpb.ProcessBillingAddressResponse, error) {
	return &billingpb.ProcessBillingAddressResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.ProcessBillingAddressResponseItem{
			HasVat:      true,
			Vat:         10,
			Amount:      10,
			TotalAmount: 20,
		},
	}, nil
}

func (s *BillingServerOkMock) ChangeMerchantData(
	ctx context.Context,
	in *billingpb.ChangeMerchantDataRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	rsp := &billingpb.ChangeMerchantDataResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   OnboardingMerchantMock,
	}

	if in.MerchantId == SomeMerchantId {
		return nil, SomeError
	}

	return rsp, nil
}

func (s *BillingServerOkMock) SetMerchantS3Agreement(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeMerchantDataResponse, error) {
	rsp := &billingpb.ChangeMerchantDataResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   OnboardingMerchantMock,
	}

	if in.MerchantId == SomeMerchantId {
		return nil, SomeError
	}

	return rsp, nil
}

func (s *BillingServerOkMock) ChangeProject(
	ctx context.Context,
	in *billingpb.Project,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return &billingpb.ChangeProjectResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
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

func (s *BillingServerOkMock) DeleteProject(
	ctx context.Context,
	in *billingpb.GetProjectRequest,
	opts ...client.CallOption,
) (*billingpb.ChangeProjectResponse, error) {
	return &billingpb.ChangeProjectResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) CreateToken(
	ctx context.Context,
	in *billingpb.TokenRequest,
	opts ...client.CallOption,
) (*billingpb.TokenResponse, error) {
	return &billingpb.TokenResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) CheckProjectRequestSignature(
	ctx context.Context,
	in *billingpb.CheckProjectRequestSignatureRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) CreateOrUpdateProduct(ctx context.Context, in *billingpb.Product, opts ...client.CallOption) (*billingpb.Product, error) {
	return Product, nil
}

func (s *BillingServerOkMock) ListProducts(ctx context.Context, in *billingpb.ListProductsRequest, opts ...client.CallOption) (*billingpb.ListProductsResponse, error) {
	return &billingpb.ListProductsResponse{
		Limit:  1,
		Offset: 0,
		Total:  200,
		Products: []*billingpb.Product{
			Product,
		},
	}, nil
}

func (s *BillingServerOkMock) GetProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.GetProductResponse, error) {
	return GetProductResponse, nil
}

func (s *BillingServerOkMock) DeleteProduct(ctx context.Context, in *billingpb.RequestProduct, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkMock) ListProjects(ctx context.Context, in *billingpb.ListProjectsRequest, opts ...client.CallOption) (*billingpb.ListProjectsResponse, error) {
	return &billingpb.ListProjectsResponse{Count: 1, Items: []*billingpb.Project{{Id: "id"}}}, nil
}
func (s *BillingServerOkMock) GetOrder(ctx context.Context, in *billingpb.GetOrderRequest, opts ...client.CallOption) (*billingpb.Order, error) {
	return &billingpb.Order{}, nil
}

func (s *BillingServerOkMock) IsOrderCanBePaying(
	ctx context.Context,
	in *billingpb.IsOrderCanBePayingRequest,
	opts ...client.CallOption,
) (*billingpb.IsOrderCanBePayingResponse, error) {
	return &billingpb.IsOrderCanBePayingResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.Order{},
	}, nil
}

func (s *BillingServerOkMock) GetCountry(ctx context.Context, in *billingpb.GetCountryRequest, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) UpdateCountry(ctx context.Context, in *billingpb.Country, opts ...client.CallOption) (*billingpb.Country, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPriceGroup(ctx context.Context, in *billingpb.GetPriceGroupRequest, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	return &billingpb.PriceGroup{
		Id: "some_id",
	}, nil
}

func (s *BillingServerOkMock) UpdatePriceGroup(ctx context.Context, in *billingpb.PriceGroup, opts ...client.CallOption) (*billingpb.PriceGroup, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetUserNotifySales(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetUserNotifyNewRegion(ctx context.Context, in *billingpb.SetUserNotifyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCountriesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.CountriesList, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystemRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	return &billingpb.PaymentChannelCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) SetPaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentChannelCostSystem, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemResponse, error) {
	return &billingpb.PaymentChannelCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) DeletePaymentChannelCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	return &billingpb.PaymentChannelCostMerchantResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) SetPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchant, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantResponse, error) {
	return &billingpb.PaymentChannelCostMerchantResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) DeletePaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystemRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	return &billingpb.MoneyBackCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) SetMoneyBackCostSystem(ctx context.Context, in *billingpb.MoneyBackCostSystem, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemResponse, error) {
	return &billingpb.MoneyBackCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) DeleteMoneyBackCostSystem(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	return &billingpb.MoneyBackCostMerchantResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.MoneyBackCostMerchant{},
	}, nil
}

func (s *BillingServerOkMock) SetMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchant, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantResponse, error) {
	return &billingpb.MoneyBackCostMerchantResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.MoneyBackCostMerchant{},
	}, nil
}

func (s *BillingServerOkMock) DeleteMoneyBackCostMerchant(ctx context.Context, in *billingpb.PaymentCostDeleteRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) CreateOrUpdatePaymentMethodTestSettings(ctx context.Context, in *billingpb.ChangePaymentMethodParamsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeletePaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) FindByZipCode(
	ctx context.Context,
	in *billingpb.FindByZipCodeRequest,
	opts ...client.CallOption,
) (*billingpb.FindByZipCodeResponse, error) {
	return &billingpb.FindByZipCodeResponse{
		Count: 1,
		Items: []*billingpb.ZipCode{
			{
				Zip:     in.Zip,
				Country: in.Country,
			},
		},
	}, nil
}

func (s *BillingServerOkMock) GetAllPaymentChannelCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostSystemListResponse, error) {
	return &billingpb.PaymentChannelCostSystemListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetAllPaymentChannelCostMerchant(ctx context.Context, in *billingpb.PaymentChannelCostMerchantListRequest, opts ...client.CallOption) (*billingpb.PaymentChannelCostMerchantListResponse, error) {
	return &billingpb.PaymentChannelCostMerchantListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetAllMoneyBackCostSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostSystemListResponse, error) {
	return &billingpb.MoneyBackCostSystemListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetAllMoneyBackCostMerchant(ctx context.Context, in *billingpb.MoneyBackCostMerchantListRequest, opts ...client.CallOption) (*billingpb.MoneyBackCostMerchantListResponse, error) {
	return &billingpb.MoneyBackCostMerchantListResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) CreateOrUpdatePaymentMethod(
	ctx context.Context,
	in *billingpb.PaymentMethod,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodResponse, error) {
	return &billingpb.ChangePaymentMethodResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) CreateOrUpdatePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.ChangePaymentMethodParamsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return &billingpb.ChangePaymentMethodParamsResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) DeletePaymentMethodProductionSettings(
	ctx context.Context,
	in *billingpb.GetPaymentMethodSettingsRequest,
	opts ...client.CallOption,
) (*billingpb.ChangePaymentMethodParamsResponse, error) {
	return &billingpb.ChangePaymentMethodParamsResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) CreateAccountingEntry(ctx context.Context, in *billingpb.CreateAccountingEntryRequest, opts ...client.CallOption) (*billingpb.CreateAccountingEntryResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) CreateRoyaltyReport(ctx context.Context, in *billingpb.CreateRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.CreateRoyaltyReportRequest, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ListRoyaltyReports(ctx context.Context, in *billingpb.ListRoyaltyReportsRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	return &billingpb.ListRoyaltyReportsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.RoyaltyReportsPaginate{
			Count: 1,
			Items: []*billingpb.RoyaltyReport{},
		},
	}, nil
}

func (s *BillingServerOkMock) ListRoyaltyReportOrders(ctx context.Context, in *billingpb.ListRoyaltyReportOrdersRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	return &billingpb.TransactionsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.TransactionsPaginate{
			Count: 100,
			Items: []*billingpb.OrderViewPublic{},
		},
	}, nil
}

func (s *BillingServerOkMock) GetVatReportsDashboard(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	return &billingpb.VatReportsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.VatReportsPaginate{
			Count: 1,
			Items: []*billingpb.VatReport{},
		},
	}, nil
}

func (s *BillingServerOkMock) GetVatReportsForCountry(ctx context.Context, in *billingpb.VatReportsRequest, opts ...client.CallOption) (*billingpb.VatReportsResponse, error) {
	return &billingpb.VatReportsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.VatReportsPaginate{
			Count: 100,
			Items: []*billingpb.VatReport{},
		},
	}, nil
}

func (s *BillingServerOkMock) CalcAnnualTurnovers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}
func (s *BillingServerOkMock) ProcessVatReports(ctx context.Context, in *billingpb.ProcessVatReportsRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) UpdateVatReportStatus(ctx context.Context, in *billingpb.UpdateVatReportStatusRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
	}, nil
}

func (s *BillingServerOkMock) GetPriceGroupByCountry(
	ctx context.Context,
	in *billingpb.PriceGroupByCountryRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroup, error) {
	return &billingpb.PriceGroup{}, nil
}

func (s *BillingServerOkMock) UpdateProductPrices(
	ctx context.Context,
	in *billingpb.UpdateProductPricesRequest,
	opts ...client.CallOption,
) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{}, nil
}

func (s *BillingServerOkMock) GetProductPrices(
	ctx context.Context,
	in *billingpb.RequestProduct,
	opts ...client.CallOption,
) (*billingpb.ProductPricesResponse, error) {
	return &billingpb.ProductPricesResponse{}, nil
}

func (s *BillingServerOkMock) GetPriceGroupRecommendedPrice(
	ctx context.Context,
	in *billingpb.RecommendedPriceRequest,
	opts ...client.CallOption,
) (*billingpb.RecommendedPriceResponse, error) {
	return &billingpb.RecommendedPriceResponse{}, nil
}

func (s *BillingServerOkMock) GetPriceGroupCurrencyByRegion(
	ctx context.Context,
	in *billingpb.PriceGroupByRegionRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return &billingpb.PriceGroupCurrenciesResponse{
		Region: []*billingpb.PriceGroupRegions{
			{Currency: "USD"},
		},
	}, nil
}

func (s *BillingServerOkMock) GetPriceGroupCurrencies(
	ctx context.Context,
	in *billingpb.EmptyRequest,
	opts ...client.CallOption,
) (*billingpb.PriceGroupCurrenciesResponse, error) {
	return &billingpb.PriceGroupCurrenciesResponse{}, nil
}

func (s *BillingServerOkMock) GetPaymentMethodProductionSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	return &billingpb.GetPaymentMethodSettingsResponse{
		Params: []*billingpb.PaymentMethodParams{
			{Currency: "RUB"},
		},
	}, nil
}

func (s *BillingServerOkMock) GetPaymentMethodTestSettings(ctx context.Context, in *billingpb.GetPaymentMethodSettingsRequest, opts ...client.CallOption) (*billingpb.GetPaymentMethodSettingsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ChangeRoyaltyReport(ctx context.Context, in *billingpb.ChangeRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) AutoAcceptRoyaltyReports(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetUserProfile(
	ctx context.Context,
	in *billingpb.GetUserProfileRequest,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return &billingpb.GetUserProfileResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.UserProfile{},
	}, nil
}

func (s *BillingServerOkMock) CreateOrUpdateUserProfile(
	ctx context.Context,
	in *billingpb.UserProfile,
	opts ...client.CallOption,
) (*billingpb.GetUserProfileResponse, error) {
	return &billingpb.GetUserProfileResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   &billingpb.UserProfile{},
	}, nil
}

func (s *BillingServerOkMock) ConfirmUserEmail(
	ctx context.Context,
	in *billingpb.ConfirmUserEmailRequest,
	opts ...client.CallOption,
) (*billingpb.ConfirmUserEmailResponse, error) {
	return &billingpb.ConfirmUserEmailResponse{
		Status: billingpb.ResponseStatusOk,
		Profile: &billingpb.UserProfile{
			Id:     bson.NewObjectId().Hex(),
			UserId: bson.NewObjectId().Hex(),
			Email:  &billingpb.UserProfileEmail{Email: "test@test.com"},
		},
	}, nil
}

func (s *BillingServerOkMock) CreatePageReview(
	ctx context.Context,
	in *billingpb.CreatePageReviewRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) MerchantReviewRoyaltyReport(ctx context.Context, in *billingpb.MerchantReviewRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.ResponseError, error) {
	return &billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetMerchantOnboardingCompleteData(
	ctx context.Context,
	in *billingpb.SetMerchantS3AgreementRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantOnboardingCompleteDataResponse, error) {
	return &billingpb.GetMerchantOnboardingCompleteDataResponse{Status: billingpb.ResponseStatusOk}, nil
}

func (s *BillingServerOkMock) GetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.GetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.GetMerchantTariffRatesResponse, error) {
	return &billingpb.GetMerchantTariffRatesResponse{}, nil
}

func (s *BillingServerOkMock) SetMerchantTariffRates(
	ctx context.Context,
	in *billingpb.SetMerchantTariffRatesRequest,
	opts ...client.CallOption,
) (*billingpb.CheckProjectRequestSignatureResponse, error) {
	return &billingpb.CheckProjectRequestSignatureResponse{}, nil
}

func (s *BillingServerOkMock) CreateOrUpdateKeyProduct(ctx context.Context, in *billingpb.CreateOrUpdateKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusOk,
		Product: &billingpb.KeyProduct{},
	}, nil
}

func (s *BillingServerOkMock) GetKeyProducts(ctx context.Context, in *billingpb.ListKeyProductsRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return &billingpb.ListKeyProductsResponse{
		Status: billingpb.ResponseStatusOk,
		Count:  1,
		Products: []*billingpb.KeyProduct{
			{},
		},
	}, nil
}

func (s *BillingServerOkMock) GetKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusOk,
		Product: &billingpb.KeyProduct{},
	}, nil
}

func (s *BillingServerOkMock) DeleteKeyProduct(ctx context.Context, in *billingpb.RequestKeyProductMerchant, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) PublishKeyProduct(ctx context.Context, in *billingpb.PublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	return &billingpb.KeyProductResponse{
		Status:  billingpb.ResponseStatusOk,
		Product: &billingpb.KeyProduct{},
	}, nil
}

func (s *BillingServerOkMock) GetKeyProductsForOrder(ctx context.Context, in *billingpb.GetKeyProductsForOrderRequest, opts ...client.CallOption) (*billingpb.ListKeyProductsResponse, error) {
	return &billingpb.ListKeyProductsResponse{
		Status: billingpb.ResponseStatusOk,
		Count:  1,
		Products: []*billingpb.KeyProduct{
			{},
		},
	}, nil
}

func (s *BillingServerOkMock) GetPlatforms(ctx context.Context, in *billingpb.ListPlatformsRequest, opts ...client.CallOption) (*billingpb.ListPlatformsResponse, error) {
	return &billingpb.ListPlatformsResponse{
		Status: billingpb.ResponseStatusOk,
		Count:  1,
		Platforms: []*billingpb.Platform{
			{},
		},
	}, nil
}

func (s *BillingServerOkMock) DeletePlatformFromProduct(ctx context.Context, in *billingpb.RemovePlatformRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetAvailableKeysCount(ctx context.Context, in *billingpb.GetPlatformKeyCountRequest, opts ...client.CallOption) (*billingpb.GetPlatformKeyCountResponse, error) {
	return &billingpb.GetPlatformKeyCountResponse{
		Status: billingpb.ResponseStatusOk,
		Count:  1000,
	}, nil
}

func (s *BillingServerOkMock) UploadKeysFile(ctx context.Context, in *billingpb.PlatformKeysFileRequest, opts ...client.CallOption) (*billingpb.PlatformKeysFileResponse, error) {
	return &billingpb.PlatformKeysFileResponse{
		Status:        billingpb.ResponseStatusOk,
		KeysProcessed: 1000,
		TotalCount:    2000,
	}, nil
}

func (s *BillingServerOkMock) GetKeyByID(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return &billingpb.GetKeyForOrderRequestResponse{
		Status: billingpb.ResponseStatusOk,
		Key:    &billingpb.Key{},
	}, nil
}

func (s *BillingServerOkMock) ReserveKeyForOrder(ctx context.Context, in *billingpb.PlatformKeyReserveRequest, opts ...client.CallOption) (*billingpb.PlatformKeyReserveResponse, error) {
	return &billingpb.PlatformKeyReserveResponse{
		KeyId:  SomeMerchantId,
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) FinishRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.GetKeyForOrderRequestResponse, error) {
	return &billingpb.GetKeyForOrderRequestResponse{
		Status: billingpb.ResponseStatusOk,
		Key:    &billingpb.Key{},
	}, nil
}

func (s *BillingServerOkMock) CancelRedeemKeyForOrder(ctx context.Context, in *billingpb.KeyForOrderRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) GetKeyProductInfo(ctx context.Context, in *billingpb.GetKeyProductInfoRequest, opts ...client.CallOption) (*billingpb.GetKeyProductInfoResponse, error) {
	return &billingpb.GetKeyProductInfoResponse{
		Status: billingpb.ResponseStatusOk,
	}, nil
}

func (s *BillingServerOkMock) ChangeCodeInOrder(ctx context.Context, in *billingpb.ChangeCodeInOrderRequest, opts ...client.CallOption) (*billingpb.ChangeCodeInOrderResponse, error) {
	return &billingpb.ChangeCodeInOrderResponse{
		Status: billingpb.ResponseStatusOk,
		Order:  &billingpb.Order{},
	}, nil
}

func (s *BillingServerOkMock) GetOrderPublic(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPublicResponse, error) {
	item := new(billingpb.OrderViewPublic)
	_ = faker.FakeData(item)

	return &billingpb.GetOrderPublicResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   item,
	}, nil
}

func (s *BillingServerOkMock) GetOrderPrivate(
	ctx context.Context,
	in *billingpb.GetOrderRequest,
	opts ...client.CallOption,
) (*billingpb.GetOrderPrivateResponse, error) {
	item := new(billingpb.OrderViewPrivate)
	_ = faker.FakeData(item)

	return &billingpb.GetOrderPrivateResponse{
		Status: billingpb.ResponseStatusOk,
		Item:   item,
	}, nil
}

func (s *BillingServerOkMock) FindAllOrdersPublic(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPublicResponse, error) {
	count := 5
	items := make([]*billingpb.OrderViewPublic, 0, count)

	for i := 0; i < count; i++ {
		item := new(billingpb.OrderViewPublic)
		_ = faker.FakeData(item)
		items = append(items, item)
	}

	return &billingpb.ListOrdersPublicResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.ListOrdersPublicResponseItem{
			Count: int64(count),
			Items: items,
		},
	}, nil
}

func (s *BillingServerOkMock) FindAllOrdersPrivate(
	ctx context.Context,
	in *billingpb.ListOrdersRequest,
	opts ...client.CallOption,
) (*billingpb.ListOrdersPrivateResponse, error) {
	count := 5
	items := make([]*billingpb.OrderViewPrivate, 0, count)

	for i := 0; i < count; i++ {
		item := new(billingpb.OrderViewPrivate)
		_ = faker.FakeData(item)
		items = append(items, item)
	}

	return &billingpb.ListOrdersPrivateResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.ListOrdersPrivateResponseItem{
			Count: int64(count),
			Items: items,
		},
	}, nil
}

func (s *BillingServerOkMock) GetDashboardMainReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardMainResponse, error) {
	return &billingpb.GetDashboardMainResponse{}, nil
}
func (s *BillingServerOkMock) GetDashboardRevenueDynamicsReport(
	ctx context.Context,
	in *billingpb.GetDashboardMainRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardRevenueDynamicsReportResponse, error) {
	return &billingpb.GetDashboardRevenueDynamicsReportResponse{}, nil
}

func (s *BillingServerOkMock) GetDashboardBaseReport(
	ctx context.Context,
	in *billingpb.GetDashboardBaseReportRequest,
	opts ...client.CallOption,
) (*billingpb.GetDashboardBaseReportResponse, error) {
	return &billingpb.GetDashboardBaseReportResponse{}, nil
}

func (s *BillingServerOkMock) CreatePayoutDocument(ctx context.Context, in *billingpb.CreatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.CreatePayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) UpdatePayoutDocument(ctx context.Context, in *billingpb.UpdatePayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPayoutDocuments(ctx context.Context, in *billingpb.GetPayoutDocumentsRequest, opts ...client.CallOption) (*billingpb.GetPayoutDocumentsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) UpdatePayoutDocumentSignatures(ctx context.Context, in *billingpb.UpdatePayoutDocumentSignaturesRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetMerchantBalance(ctx context.Context, in *billingpb.GetMerchantBalanceRequest, opts ...client.CallOption) (*billingpb.GetMerchantBalanceResponse, error) {
	return &billingpb.GetMerchantBalanceResponse{
		Status: billingpb.ResponseStatusOk,
		Item: &billingpb.MerchantBalance{
			Id:             bson.NewObjectId().Hex(),
			MerchantId:     bson.NewObjectId().Hex(),
			Currency:       "RUB",
			Debit:          0,
			Credit:         0,
			RollingReserve: 0,
			Total:          0,
			CreatedAt:      ptypes.TimestampNow(),
		},
	}, nil
}

func (s *BillingServerOkMock) PayoutDocumentPdfUploaded(ctx context.Context, in *billingpb.PayoutDocumentPdfUploadedRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetRoyaltyReport(ctx context.Context, in *billingpb.GetRoyaltyReportRequest, opts ...client.CallOption) (*billingpb.GetRoyaltyReportResponse, error) {
	return &billingpb.GetRoyaltyReportResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Item:    &billingpb.RoyaltyReport{},
	}, nil
}

func (s *BillingServerOkMock) UnPublishKeyProduct(ctx context.Context, in *billingpb.UnPublishKeyProductRequest, opts ...client.CallOption) (*billingpb.KeyProductResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) PaymentFormPlatformChanged(
	ctx context.Context,
	in *billingpb.PaymentFormUserChangePlatformRequest,
	opts ...client.CallOption,
) (*billingpb.PaymentFormDataChangeResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) OrderReceipt(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) OrderReceiptRefund(ctx context.Context, in *billingpb.OrderReceiptRequest, opts ...client.CallOption) (*billingpb.OrderReceiptResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetRecommendedPriceByPriceGroup(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetRecommendedPriceByConversion(ctx context.Context, in *billingpb.RecommendedPriceRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) CheckSkuAndKeyProject(ctx context.Context, in *billingpb.CheckSkuAndKeyProjectRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPriceGroupByRegion(ctx context.Context, in *billingpb.GetPriceGroupByRegionRequest, opts ...client.CallOption) (*billingpb.GetPriceGroupByRegionResponse, error) {
	return &billingpb.GetPriceGroupByRegionResponse{
		Status: 200,
		Group: &billingpb.PriceGroup{
			Id: "some id",
		},
	}, nil
}

func (s *BillingServerOkMock) GetMerchantUsers(ctx context.Context, in *billingpb.GetMerchantUsersRequest, opts ...client.CallOption) (*billingpb.GetMerchantUsersResponse, error) {
	return &billingpb.GetMerchantUsersResponse{
		Status: 200,
		Users: []*billingpb.UserRole{
			{MerchantId: in.MerchantId, Id: SomeMerchantId},
		},
	}, nil
}
func (s *BillingServerOkMock) FindAllOrders(ctx context.Context, in *billingpb.ListOrdersRequest, opts ...client.CallOption) (*billingpb.ListOrdersResponse, error) {
	return &billingpb.ListOrdersResponse{Status: http.StatusOK}, nil
}

func (s *BillingServerOkMock) GetAdminUsers(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetAdminUsersResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetMerchantsForUser(ctx context.Context, in *billingpb.GetMerchantsForUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantsForUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) InviteUserMerchant(ctx context.Context, in *billingpb.InviteUserMerchantRequest, opts ...client.CallOption) (*billingpb.InviteUserMerchantResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) InviteUserAdmin(ctx context.Context, in *billingpb.InviteUserAdminRequest, opts ...client.CallOption) (*billingpb.InviteUserAdminResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ResendInviteMerchant(ctx context.Context, in *billingpb.ResendInviteMerchantRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ResendInviteAdmin(ctx context.Context, in *billingpb.ResendInviteAdminRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetMerchantUser(ctx context.Context, in *billingpb.GetMerchantUserRequest, opts ...client.CallOption) (*billingpb.GetMerchantUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetAdminUser(ctx context.Context, in *billingpb.GetAdminUserRequest, opts ...client.CallOption) (*billingpb.GetAdminUserResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) AcceptInvite(ctx context.Context, in *billingpb.AcceptInviteRequest, opts ...client.CallOption) (*billingpb.AcceptInviteResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) CheckInviteToken(ctx context.Context, in *billingpb.CheckInviteTokenRequest, opts ...client.CallOption) (*billingpb.CheckInviteTokenResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ChangeRoleForMerchantUser(ctx context.Context, in *billingpb.ChangeRoleForMerchantUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ChangeRoleForAdminUser(ctx context.Context, in *billingpb.ChangeRoleForAdminUserRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetRoleList(ctx context.Context, in *billingpb.GetRoleListRequest, opts ...client.CallOption) (*billingpb.GetRoleListResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) ChangeMerchantManualPayouts(ctx context.Context, in *billingpb.ChangeMerchantManualPayoutsRequest, opts ...client.CallOption) (*billingpb.ChangeMerchantManualPayoutsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeleteMerchantUser(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeleteAdminUser(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) OrderCreateByPaylink(ctx context.Context, in *billingpb.OrderCreateByPaylink, opts ...client.CallOption) (*billingpb.OrderCreateProcessResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPaylinks(ctx context.Context, in *billingpb.GetPaylinksRequest, opts ...client.CallOption) (*billingpb.GetPaylinksResponse, error) {
	return &billingpb.GetPaylinksResponse{Status: http.StatusOK, Data: &billingpb.PaylinksPaginate{Count: 0, Items: []*billingpb.Paylink{}}}, nil
}

func (s *BillingServerOkMock) GetPaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	return &billingpb.GetPaylinkResponse{Status: http.StatusOK, Item: &billingpb.Paylink{}}, nil
}

func (s *BillingServerOkMock) IncrPaylinkVisits(ctx context.Context, in *billingpb.PaylinkRequestById, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	return &billingpb.EmptyResponse{}, nil
}

func (s *BillingServerOkMock) GetPaylinkURL(ctx context.Context, in *billingpb.GetPaylinkURLRequest, opts ...client.CallOption) (*billingpb.GetPaylinkUrlResponse, error) {
	return &billingpb.GetPaylinkUrlResponse{Status: http.StatusOK, Url: "http://someurl"}, nil
}

func (s *BillingServerOkMock) CreateOrUpdatePaylink(ctx context.Context, in *billingpb.CreatePaylinkRequest, opts ...client.CallOption) (*billingpb.GetPaylinkResponse, error) {
	return &billingpb.GetPaylinkResponse{Status: http.StatusOK, Item: &billingpb.Paylink{}}, nil
}

func (s *BillingServerOkMock) DeletePaylink(ctx context.Context, in *billingpb.PaylinkRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	return &billingpb.EmptyResponseWithStatus{Status: http.StatusOK}, nil
}

func (s *BillingServerOkMock) GetPaylinkStatTotal(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonResponse, error) {
	return &billingpb.GetPaylinkStatCommonResponse{Status: http.StatusOK, Item: &billingpb.StatCommon{}}, nil
}

func (s *BillingServerOkMock) GetPaylinkStatByCountry(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	return &billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil
}

func (s *BillingServerOkMock) GetPaylinkStatByReferrer(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	return &billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil
}

func (s *BillingServerOkMock) GetPaylinkStatByDate(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	return &billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil
}

func (s *BillingServerOkMock) GetPaylinkStatByUtm(ctx context.Context, in *billingpb.GetPaylinkStatCommonRequest, opts ...client.CallOption) (*billingpb.GetPaylinkStatCommonGroupResponse, error) {
	return &billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil
}

func (s *BillingServerOkMock) GetRecommendedPriceTable(ctx context.Context, in *billingpb.RecommendedPriceTableRequest, opts ...client.CallOption) (*billingpb.RecommendedPriceTableResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) RoyaltyReportPdfUploaded(ctx context.Context, in *billingpb.RoyaltyReportPdfUploadedRequest, opts ...client.CallOption) (*billingpb.RoyaltyReportPdfUploadedResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPayoutDocument(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.PayoutDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPayoutDocumentRoyaltyReports(ctx context.Context, in *billingpb.GetPayoutDocumentRequest, opts ...client.CallOption) (*billingpb.ListRoyaltyReportsResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) AutoCreatePayoutDocuments(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.EmptyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetAdminUserRole(ctx context.Context, in *billingpb.AdminRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetMerchantUserRole(ctx context.Context, in *billingpb.MerchantRoleRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCommonUserProfile(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.CommonUserProfileResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeleteSavedCard(ctx context.Context, in *billingpb.DeleteSavedCardRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetMerchantOperatingCompany(ctx context.Context, in *billingpb.SetMerchantOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.SetMerchantOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetMerchantAcceptedStatus(ctx context.Context, in *billingpb.SetMerchantAcceptedStatusRequest, opts ...client.CallOption) (*billingpb.SetMerchantAcceptedStatusResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetOperatingCompaniesList(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompaniesListResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) AddOperatingCompany(ctx context.Context, in *billingpb.OperatingCompany, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPaymentMinLimitsSystem(ctx context.Context, in *billingpb.EmptyRequest, opts ...client.CallOption) (*billingpb.GetPaymentMinLimitsSystemResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetPaymentMinLimitSystem(ctx context.Context, in *billingpb.PaymentMinLimitSystem, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetOperatingCompany(ctx context.Context, in *billingpb.GetOperatingCompanyRequest, opts ...client.CallOption) (*billingpb.GetOperatingCompanyResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetCountriesListForOrder(ctx context.Context, in *billingpb.GetCountriesListForOrderRequest, opts ...client.CallOption) (*billingpb.GetCountriesListForOrderResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetPaylinkTransactions(ctx context.Context, in *billingpb.GetPaylinkTransactionsRequest, opts ...client.CallOption) (*billingpb.TransactionsResponse, error) {
	return &billingpb.TransactionsResponse{
		Status:  billingpb.ResponseStatusOk,
		Message: nil,
		Data: &billingpb.TransactionsPaginate{
			Count: 100,
			Items: []*billingpb.OrderViewPublic{},
		},
	}, nil
}

func (s *BillingServerOkMock) SendWebhookToMerchant(ctx context.Context, in *billingpb.OrderCreateRequest, opts ...client.CallOption) (*billingpb.SendWebhookToMerchantResponse, error) {
	return &billingpb.SendWebhookToMerchantResponse{
		Status:  200,
		OrderId: bson.NewObjectId().Hex(),
		Message: nil,
	}, nil
}

func (s *BillingServerOkMock) NotifyWebhookTestResults(ctx context.Context, in *billingpb.NotifyWebhookTestResultsRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetAdminByUserId(ctx context.Context, in *billingpb.CommonUserProfileRequest, opts ...client.CallOption) (*billingpb.UserRoleResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) RoyaltyReportFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) PayoutFinanceDone(ctx context.Context, in *billingpb.ReportFinanceDoneRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) GetActOfCompletion(ctx context.Context, in *billingpb.ActOfCompletionRequest, opts ...client.CallOption) (*billingpb.ActOfCompletionResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) SetCustomerPaymentActivity(ctx context.Context, in *billingpb.SetCustomerPaymentActivityRequest, opts ...client.CallOption) (*billingpb.EmptyResponseWithStatus, error) {
	panic("implement me")
}

func (s *BillingServerOkMock) DeserializeCookie(ctx context.Context, in *billingpb.DeserializeCookieRequest, opts ...client.CallOption) (*billingpb.DeserializeCookieResponse, error) {
	panic("implement me")
}
