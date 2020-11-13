package mock

import (
	"github.com/bxcodec/faker"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"

	"github.com/paysuper/paysuper-proto/go/billingpb"

	"net/http"
)

func NewBillingServerOkMock() billingpb.BillingService {
	_ = faker.AddProvider("objectIdString", func(_ reflect.Value) (interface{}, error) {
		return "ffffffffffffffffffffffff", nil
	})

	bill := &billingMocks.BillingService{}

	bill.On("ProcessRefundCallback", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentNotifyResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("CreateOrUpdateKeyProduct", mock.Anything, mock.Anything).
		Return(&billingpb.KeyProductResponse{
			Status:  billingpb.ResponseStatusOk,
			Product: &billingpb.KeyProduct{},
		}, nil)

	bill.On("GetKeyProduct", mock.Anything, mock.Anything).
		Return(&billingpb.KeyProductResponse{
			Status:  billingpb.ResponseStatusOk,
			Product: &billingpb.KeyProduct{},
		}, nil)

	bill.On("GetKeyProducts", mock.Anything, mock.Anything).
		Return(&billingpb.ListKeyProductsResponse{
			Status: billingpb.ResponseStatusOk,
			Count:  1,
			Products: []*billingpb.KeyProduct{
				{},
			},
		}, nil)

	bill.On("GetPlatforms", mock.Anything, mock.Anything).
		Return(&billingpb.ListPlatformsResponse{
			Status: billingpb.ResponseStatusOk,
			Count:  1,
			Platforms: []*billingpb.Platform{
				{},
			},
		}, nil)

	bill.On("PublishKeyProduct", mock.Anything, mock.Anything).
		Return(&billingpb.KeyProductResponse{
			Status:  billingpb.ResponseStatusOk,
			Product: &billingpb.KeyProduct{},
		}, nil)

	bill.On("GetMerchantBalance", mock.Anything, mock.Anything).
		Return(&billingpb.GetMerchantBalanceResponse{
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
		}, nil)

	bill.On("ChangeMerchantData", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeMerchantDataResponse{
			Status: billingpb.ResponseStatusOk,
			Item:   OnboardingMerchantMock,
		}, nil)

	bill.On("ChangeMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeMerchantResponse{
			Status: billingpb.ResponseStatusOk,
			Item: &billingpb.Merchant{
				User: &billingpb.MerchantUser{
					Id:    bson.NewObjectId().Hex(),
					Email: "test@unit.test",
				},
				Status: billingpb.MerchantStatusDraft,
			},
		}, nil)

	bill.On("CreateNotification", mock.Anything, mock.Anything).
		Return(&billingpb.CreateNotificationResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("GetNotification", mock.Anything, mock.Anything).
		Return(&billingpb.Notification{}, nil)

	bill.On("ListNotifications", mock.Anything, mock.Anything).
		Return(&billingpb.Notifications{}, nil)

	bill.On("MarkNotificationAsRead", mock.Anything, mock.Anything).
		Return(&billingpb.Notification{}, nil)

	bill.On("GetMerchantOnboardingCompleteData", mock.Anything, mock.Anything).
		Return(&billingpb.GetMerchantOnboardingCompleteDataResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("ListMerchants", mock.Anything, mock.Anything).
		Return(&billingpb.MerchantListingResponse{
			Count: 3,
			Items: []*billingpb.MerchantShortInfo{OnboardingMerchantShortInfoMock, OnboardingMerchantShortInfoMock, OnboardingMerchantShortInfoMock},
		}, nil)

	OnboardingMerchantMock.S3AgreementName = SomeAgreementName1
	OnboardingMerchantMock.Status = billingpb.MerchantStatusAgreementSigning
	bill.On("GetMerchantBy", mock.Anything, mock.Anything).
		Return(&billingpb.GetMerchantResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: &billingpb.ResponseErrorMessage{},
			Item:    OnboardingMerchantMock,
		}, nil)

	bill.On("CreateRefund", mock.Anything, mock.Anything).
		Return(&billingpb.CreateRefundResponse{
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
		}, nil)

	bill.On("GetRefund", mock.Anything, mock.Anything).
		Return(&billingpb.CreateRefundResponse{
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
		}, nil)

	bill.On("ListRefunds", mock.Anything, mock.Anything).
		Return(&billingpb.ListRefundsResponse{
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
		}, nil)

	bill.On("FindAllOrders", mock.Anything, mock.Anything).
		Return(&billingpb.ListOrdersResponse{Status: http.StatusOK}, nil)

	bill.On("CreateOrUpdatePaylink", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkResponse{Status: http.StatusOK, Item: &billingpb.Paylink{}}, nil)

	bill.On("DeletePaylink", mock.Anything, mock.Anything).
		Return(&billingpb.EmptyResponseWithStatus{Status: http.StatusOK}, nil)

	bill.On("GetPaylinkStatByCountry", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil)

	bill.On("GetPaylinkStatByDate", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil)

	bill.On("GetPaylinkStatByReferrer", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil)

	bill.On("GetPaylinkStatByUtm", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkStatCommonGroupResponse{Status: http.StatusOK, Item: &billingpb.GroupStatCommon{}}, nil)

	bill.On("GetPaylinkStatTotal", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkStatCommonResponse{Status: http.StatusOK, Item: &billingpb.StatCommon{}}, nil)

	bill.On("GetPaylinkTransactions", mock.Anything, mock.Anything).
		Return(&billingpb.TransactionsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.TransactionsPaginate{
				Count: 100,
				Items: []*billingpb.OrderViewPublic{},
			},
		}, nil)

	bill.On("GetPaylinkURL", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkUrlResponse{Status: http.StatusOK, Url: "http://someurl"}, nil)

	bill.On("GetPaylink", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkResponse{Status: http.StatusOK, Item: &billingpb.Paylink{}}, nil)

	bill.On("GetPaylinks", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinksResponse{
			Status: http.StatusOK,
			Data: &billingpb.PaylinksPaginate{
				Count: 0,
				Items: []*billingpb.Paylink{},
			},
		}, nil)

	bill.On("CreateOrUpdatePaylink", mock.Anything, mock.Anything).
		Return(&billingpb.GetPaylinkResponse{Status: http.StatusOK, Item: &billingpb.Paylink{}}, nil)

	bill.On("SetMoneyBackCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostMerchantResponse{
			Status: billingpb.ResponseStatusOk,
			Item:   &billingpb.MoneyBackCostMerchant{},
		}, nil)

	bill.On("DeleteMoneyBackCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetMoneyBackCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostMerchantResponse{
			Status: billingpb.ResponseStatusOk,
			Item:   &billingpb.MoneyBackCostMerchant{},
		}, nil)

	bill.On("GetAllMoneyBackCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostMerchantListResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("SetMoneyBackCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("DeleteMoneyBackCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetMoneyBackCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetAllMoneyBackCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostSystemListResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("SetPaymentChannelCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentChannelCostMerchantResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("DeletePaymentChannelCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetPaymentChannelCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentChannelCostMerchantResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetAllMoneyBackCostMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.MoneyBackCostMerchantListResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("SetPaymentChannelCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentChannelCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("DeletePaymentChannelCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetPaymentChannelCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentChannelCostSystemResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetAllPaymentChannelCostSystem", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentChannelCostSystemListResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("GetPriceGroupByRegion", mock.Anything, mock.Anything).
		Return(&billingpb.GetPriceGroupByRegionResponse{
			Status: 200,
			Group: &billingpb.PriceGroup{
				Id: "some id",
			},
		}, nil)

	bill.On("DeleteProduct", mock.Anything, mock.Anything).
		Return(&billingpb.EmptyResponse{}, nil)

	bill.On("GetProduct", mock.Anything, mock.Anything).
		Return(GetProductResponse, nil)

	bill.On("ListProducts", mock.Anything, mock.Anything).
		Return(&billingpb.ListProductsResponse{
			Limit:  1,
			Offset: 0,
			Total:  200,
			Products: []*billingpb.Product{
				Product,
			},
		}, nil)

	bill.On("CreateOrUpdateProduct", mock.Anything, mock.Anything).
		Return(Product, nil)

	bill.On("GetPriceGroupByRegion", mock.Anything, mock.Anything).
		Return(&billingpb.GetPriceGroupByRegionResponse{
			Status: 200,
			Group: &billingpb.PriceGroup{
				Id: "some id",
			},
		}, nil)

	bill.On("ChangeProject", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeProjectResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("DeleteProject", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeProjectResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetProject", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeProjectResponse{
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
		}, nil)

	bill.On("ListProjects", mock.Anything, mock.Anything).
		Return(&billingpb.ListProjectsResponse{Count: 1, Items: []*billingpb.Project{{Id: "id"}}}, nil)

	bill.On("MerchantReviewRoyaltyReport", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("ChangeRoyaltyReport", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetRoyaltyReport", mock.Anything, mock.Anything).
		Return(&billingpb.GetRoyaltyReportResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Item:    &billingpb.RoyaltyReport{},
		}, nil)

	bill.On("ListRoyaltyReports", mock.Anything, mock.Anything).
		Return(&billingpb.ListRoyaltyReportsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.RoyaltyReportsPaginate{
				Count: 1,
				Items: []*billingpb.RoyaltyReport{},
			},
		}, nil)

	bill.On("ListRoyaltyReportOrders", mock.Anything, mock.Anything).
		Return(&billingpb.TransactionsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.TransactionsPaginate{
				Count: 100,
				Items: []*billingpb.OrderViewPublic{},
			},
		}, nil)

	bill.On("MerchantReviewRoyaltyReport", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("CheckProjectRequestSignature", mock.Anything, mock.Anything).
		Return(&billingpb.CheckProjectRequestSignatureResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	bill.On("ConfirmUserEmail", mock.Anything, mock.Anything).
		Return(&billingpb.ConfirmUserEmailResponse{
			Status: billingpb.ResponseStatusOk,
			Profile: &billingpb.UserProfile{
				Id:     bson.NewObjectId().Hex(),
				UserId: bson.NewObjectId().Hex(),
				Email:  &billingpb.UserProfileEmail{Email: "test@test.com"},
			},
		}, nil)

	bill.On("CreatePageReview", mock.Anything, mock.Anything).
		Return(&billingpb.CheckProjectRequestSignatureResponse{Status: billingpb.ResponseStatusOk}, nil)

	bill.On("GetUserProfile", mock.Anything, mock.Anything).
		Return(&billingpb.GetUserProfileResponse{
			Status: billingpb.ResponseStatusOk,
			Item:   &billingpb.UserProfile{},
		}, nil)

	bill.On("CreateOrUpdateUserProfile", mock.Anything, mock.Anything).
		Return(&billingpb.GetUserProfileResponse{
			Status: billingpb.ResponseStatusOk,
			Item:   &billingpb.UserProfile{},
		}, nil)

	bill.On("GetVatReportTransactions", mock.Anything, mock.Anything).
		Return(&billingpb.PrivateTransactionsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.PrivateTransactionsPaginate{
				Count: 100,
				Items: []*billingpb.OrderViewPrivate{},
			},
		}, nil)

	bill.On("GetVatReportsDashboard", mock.Anything, mock.Anything).
		Return(&billingpb.VatReportsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.VatReportsPaginate{
				Count: 1,
				Items: []*billingpb.VatReport{},
			},
		}, nil)

	bill.On("GetVatReportsForCountry", mock.Anything, mock.Anything).
		Return(&billingpb.VatReportsResponse{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
			Data: &billingpb.VatReportsPaginate{
				Count: 100,
				Items: []*billingpb.VatReport{},
			},
		}, nil)

	bill.On("UpdateVatReportStatus", mock.Anything, mock.Anything).
		Return(&billingpb.ResponseError{
			Status:  billingpb.ResponseStatusOk,
			Message: nil,
		}, nil)

	bill.On("SendWebhookToMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.SendWebhookToMerchantResponse{
			Status:  200,
			OrderId: bson.NewObjectId().Hex(),
			Message: nil,
		}, nil)

	bill.On("CreateToken", mock.Anything, mock.Anything).
		Return(&billingpb.TokenResponse{
			Status: billingpb.ResponseStatusOk,
		}, nil)

	return bill
}
