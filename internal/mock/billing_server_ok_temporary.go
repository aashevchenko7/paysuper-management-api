package mock

import (
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/paysuper/paysuper-proto/go/billingpb"
)

func NewBillingServerOkTemporaryMock() billingpb.BillingService {
	bill := &billingMocks.BillingService{}

	bill.On("ProcessRefundCallback", mock.Anything, mock.Anything).
		Return(&billingpb.PaymentNotifyResponse{
			Status: billingpb.ResponseStatusOk,
			Error:  SomeError.Message,
		}, nil)

	return bill
}

func (s *BillingServerOkTemporaryMock) AddMerchantDocument(ctx context.Context, in *billingpb.MerchantDocument, opts ...client.CallOption) (*billingpb.AddMerchantDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantDocument(ctx context.Context, in *billingpb.GetMerchantDocumentRequest, opts ...client.CallOption) (*billingpb.GetMerchantDocumentResponse, error) {
	panic("implement me")
}

func (s *BillingServerOkTemporaryMock) GetMerchantDocuments(ctx context.Context, in *billingpb.GetMerchantDocumentsRequest, opts ...client.CallOption) (*billingpb.GetMerchantDocumentsResponse, error) {
	panic("implement me")
}
