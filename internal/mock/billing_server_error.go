package mock

import (
	"github.com/paysuper/paysuper-proto/go/billingpb"
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func NewBillingServerErrorMock() billingpb.BillingService {
	bill := &billingMocks.BillingService{}

	bill.On("ProcessRefundCallback", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentNotifyResponse{
			Status: billingpb.ResponseStatusNotFound,
			Error:  SomeError.Message,
		}, nil)

	bill.On("CreateRefund", mock.Anything, mock.Anything).
		Return(&billingpb.CreateRefundResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("GetRefund", mock.Anything, mock.Anything).
		Return(&billingpb.CreateRefundResponse{
			Status:  billingpb.ResponseStatusNotFound,
			Message: SomeError,
		}, nil)

	bill.On("FindByZipCode", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("ConfirmUserEmail", mock.Anything, mock.Anything).
		Return(&billingpb.ConfirmUserEmailResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("ChangeMerchant", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeMerchantResponse{
			Status:  http.StatusBadRequest,
			Message: SomeError,
		}, nil)

	bill.On("CreatePageReview", mock.Anything, mock.Anything).
		Return(&billingpb.CheckProjectRequestSignatureResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("GetUserProfile", mock.Anything, mock.Anything).
		Return(&billingpb.GetUserProfileResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("CreateOrUpdateUserProfile", mock.Anything, mock.Anything).
		Return(&billingpb.GetUserProfileResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("CheckProjectRequestSignature", mock.Anything, mock.Anything).
		Return(&billingpb.CheckProjectRequestSignatureResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("CreateToken", mock.Anything, mock.Anything).
		Return(&billingpb.TokenResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("ChangeProject", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeProjectResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	bill.On("DeleteProject", mock.Anything, mock.Anything).
		Return(&billingpb.ChangeProjectResponse{
			Status:  billingpb.ResponseStatusBadData,
			Message: SomeError,
		}, nil)

	return bill
}

func (s *BillingServerErrorMock) AddMerchantDocument(ctx context.Context, in *billingpb.MerchantDocument, opts ...client.CallOption) (*billingpb.AddMerchantDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantDocument(ctx context.Context, in *billingpb.GetMerchantDocumentRequest, opts ...client.CallOption) (*billingpb.GetMerchantDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerErrorMock) GetMerchantDocuments(ctx context.Context, in *billingpb.GetMerchantDocumentsRequest, opts ...client.CallOption) (*billingpb.GetMerchantDocumentsResponse, error) {
	panic("implement me")
}
