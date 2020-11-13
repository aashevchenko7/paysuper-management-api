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
