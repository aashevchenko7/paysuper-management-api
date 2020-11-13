package mock

import (
	"github.com/globalsign/mgo/bson"

	"github.com/paysuper/paysuper-proto/go/billingpb"
)

const (
	SomeAgreementName  = "some_name.pdf"
	SomeAgreementName1 = "some_name1.pdf"
)

var (
	SomeError = &billingpb.ResponseErrorMessage{Message: "some error"}

	SomeMerchantId  = bson.NewObjectId().Hex()
	SomeMerchantId1 = bson.NewObjectId().Hex()

	OnboardingMerchantMock = &billingpb.Merchant{
		Id: bson.NewObjectId().Hex(),
		Company: &billingpb.MerchantCompanyInfo{
			Name:    "merchant1",
			Country: "RU",
			Zip:     "190000",
			City:    "St.Petersburg",
		},
		Contacts: &billingpb.MerchantContact{
			Authorized: &billingpb.MerchantContactAuthorized{
				Name:     "Unit Test",
				Email:    "test@unit.test",
				Phone:    "123456789",
				Position: "Unit Test",
			},
			Technical: &billingpb.MerchantContactTechnical{
				Name:  "Unit Test",
				Email: "test@unit.test",
				Phone: "123456789",
			},
		},
		Banking: &billingpb.MerchantBanking{
			Currency: "RUB",
			Name:     "Bank name",
		},
		IsVatEnabled:              true,
		IsCommissionToUserEnabled: true,
		Status:                    billingpb.MerchantStatusAgreementSigning,
		LastPayout:                &billingpb.MerchantLastPayout{},
		IsSigned:                  true,
		PaymentMethods: map[string]*billingpb.MerchantPaymentMethod{
			bson.NewObjectId().Hex(): {
				PaymentMethod: &billingpb.MerchantPaymentMethodIdentification{
					Id:   bson.NewObjectId().Hex(),
					Name: "Bank card",
				},
				Commission: &billingpb.MerchantPaymentMethodCommissions{
					Fee: 2.5,
					PerTransaction: &billingpb.MerchantPaymentMethodPerTransactionCommission{
						Fee:      30,
						Currency: "RUB",
					},
				},
				Integration: &billingpb.MerchantPaymentMethodIntegration{
					TerminalId:       "1234567890",
					TerminalPassword: "0987654321",
					Integrated:       true,
				},
				IsActive: true,
			},
		},
	}

	OnboardingMerchantShortInfoMock = &billingpb.MerchantShortInfo{
		Id:       OnboardingMerchantMock.Id,
		Company:  OnboardingMerchantMock.Company,
		Contacts: OnboardingMerchantMock.Contacts,
		Banking:  OnboardingMerchantMock.Banking,
		Status:   OnboardingMerchantMock.Status,
	}

	ProductPrice = &billingpb.ProductPrice{
		Currency: "USD",
		Amount:   1010.23,
	}

	Product = &billingpb.Product{
		Id:              "5c99391568add439ccf0ffaf",
		Object:          "product",
		Type:            "simple_product",
		Sku:             "ru_double_yeti_rel",
		Name:            map[string]string{"en": "Double Yeti"},
		DefaultCurrency: "USD",
		Enabled:         true,
		Description:     map[string]string{"en": "Yet another cool game"},
		LongDescription: map[string]string{"en": "Super game steam keys"},
		Url:             "http://mygame.ru/duoble_yeti",
		Images:          []string{"/home/image.jpg"},
		MerchantId:      "5bdc35de5d1e1100019fb7db",
		Metadata: map[string]string{
			"SomeKey": "SomeValue",
		},
		Prices: []*billingpb.ProductPrice{
			ProductPrice,
		},
	}

	GetProductResponse = &billingpb.GetProductResponse{
		Status: 200,
		Item: &billingpb.Product{
			Id:              "5c99391568add439ccf0ffaf",
			Object:          "product",
			Type:            "simple_product",
			Sku:             "ru_double_yeti_rel",
			Name:            map[string]string{"en": "Double Yeti"},
			DefaultCurrency: "USD",
			Enabled:         true,
			Description:     map[string]string{"en": "Yet another cool game"},
			LongDescription: map[string]string{"en": "Super game steam keys"},
			Url:             "http://mygame.ru/duoble_yeti",
			Images:          []string{"/home/image.jpg"},
			MerchantId:      "5bdc35de5d1e1100019fb7db",
			Metadata: map[string]string{
				"SomeKey": "SomeValue",
			},
			Prices: []*billingpb.ProductPrice{
				ProductPrice,
			},
		},
	}
)
