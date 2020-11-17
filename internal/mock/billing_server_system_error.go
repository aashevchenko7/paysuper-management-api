package mock

import (
	billingMocks "github.com/paysuper/paysuper-proto/go/billingpb/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/paysuper/paysuper-proto/go/billingpb"
)

func NewBillingServerSystemErrorMock() billingpb.BillingService {
	bill := &billingMocks.BillingService{}

	bill.On("ProcessRefundCallback", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("ListRefunds", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("CreateRefund", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("GetRefund", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("ListProjects", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("GetProject", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("ChangeProject", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("DeleteProject", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("CheckProjectRequestSignature", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("CreateToken", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("ConfirmUserEmail", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("CreatePageReview", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("GetUserProfile", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	bill.On("CreateOrUpdateUserProfile", mock.Anything, mock.Anything).
		Return(nil, SomeError)

	return bill
}
