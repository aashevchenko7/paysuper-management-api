// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import client "github.com/micro/go-micro/client"
import context "context"
import currencies "github.com/paysuper/paysuper-currencies/pkg/proto/currencies"
import mock "github.com/stretchr/testify/mock"

// CurrencyratesService is an autogenerated mock type for the CurrencyratesService type
type CurrencyratesService struct {
	mock.Mock
}

// AddCommonRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) AddCommonRateCorrectionRule(ctx context.Context, in *currencies.CommonCorrectionRule, opts ...client.CallOption) (*currencies.EmptyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.EmptyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.CommonCorrectionRule, ...client.CallOption) *currencies.EmptyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.EmptyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.CommonCorrectionRule, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddMerchantRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) AddMerchantRateCorrectionRule(ctx context.Context, in *currencies.CorrectionRule, opts ...client.CallOption) (*currencies.EmptyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.EmptyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.CorrectionRule, ...client.CallOption) *currencies.EmptyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.EmptyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.CorrectionRule, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyByDateCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) ExchangeCurrencyByDateCommon(ctx context.Context, in *currencies.ExchangeCurrencyByDateCommonRequest, opts ...client.CallOption) (*currencies.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.ExchangeCurrencyByDateCommonRequest, ...client.CallOption) *currencies.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.ExchangeCurrencyByDateCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyByDateForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) ExchangeCurrencyByDateForMerchant(ctx context.Context, in *currencies.ExchangeCurrencyByDateForMerchantRequest, opts ...client.CallOption) (*currencies.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.ExchangeCurrencyByDateForMerchantRequest, ...client.CallOption) *currencies.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.ExchangeCurrencyByDateForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyCurrentCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) ExchangeCurrencyCurrentCommon(ctx context.Context, in *currencies.ExchangeCurrencyCurrentCommonRequest, opts ...client.CallOption) (*currencies.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.ExchangeCurrencyCurrentCommonRequest, ...client.CallOption) *currencies.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.ExchangeCurrencyCurrentCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeCurrencyCurrentForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *currencies.ExchangeCurrencyCurrentForMerchantRequest, opts ...client.CallOption) (*currencies.ExchangeCurrencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.ExchangeCurrencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.ExchangeCurrencyCurrentForMerchantRequest, ...client.CallOption) *currencies.ExchangeCurrencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.ExchangeCurrencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.ExchangeCurrencyCurrentForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountingCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetAccountingCurrencies(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetCommonRateCorrectionRule(ctx context.Context, in *currencies.CommonCorrectionRuleRequest, opts ...client.CallOption) (*currencies.CorrectionRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CorrectionRule
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.CommonCorrectionRuleRequest, ...client.CallOption) *currencies.CorrectionRule); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CorrectionRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.CommonCorrectionRuleRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyByRegion provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetCurrencyByRegion(ctx context.Context, in *currencies.RegionRequest, opts ...client.CallOption) (*currencies.CurrencyListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrencyListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.RegionRequest, ...client.CallOption) *currencies.CurrencyListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrencyListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.RegionRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyList provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetCurrencyList(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrencyListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrencyListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrencyListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrencyListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyRegionByCountry provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetCurrencyRegionByCountry(ctx context.Context, in *currencies.CountryRequest, opts ...client.CallOption) (*currencies.CurrencyRegion, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrencyRegion
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.CountryRequest, ...client.CallOption) *currencies.CurrencyRegion); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrencyRegion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.CountryRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerchantRateCorrectionRule provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetMerchantRateCorrectionRule(ctx context.Context, in *currencies.MerchantCorrectionRuleRequest, opts ...client.CallOption) (*currencies.CorrectionRule, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CorrectionRule
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.MerchantCorrectionRuleRequest, ...client.CallOption) *currencies.CorrectionRule); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CorrectionRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.MerchantCorrectionRuleRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPriceCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetPriceCurrencies(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateByDateCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetRateByDateCommon(ctx context.Context, in *currencies.GetRateByDateCommonRequest, opts ...client.CallOption) (*currencies.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.GetRateByDateCommonRequest, ...client.CallOption) *currencies.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.GetRateByDateCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateByDateForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetRateByDateForMerchant(ctx context.Context, in *currencies.GetRateByDateForMerchantRequest, opts ...client.CallOption) (*currencies.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.GetRateByDateForMerchantRequest, ...client.CallOption) *currencies.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.GetRateByDateForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateCurrentCommon provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetRateCurrentCommon(ctx context.Context, in *currencies.GetRateCurrentCommonRequest, opts ...client.CallOption) (*currencies.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.GetRateCurrentCommonRequest, ...client.CallOption) *currencies.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.GetRateCurrentCommonRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateCurrentForMerchant provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetRateCurrentForMerchant(ctx context.Context, in *currencies.GetRateCurrentForMerchantRequest, opts ...client.CallOption) (*currencies.RateData, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.RateData
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.GetRateCurrentForMerchantRequest, ...client.CallOption) *currencies.RateData); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.RateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.GetRateCurrentForMerchantRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecommendedPrice provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetRecommendedPrice(ctx context.Context, in *currencies.RecommendedPriceRequest, opts ...client.CallOption) (*currencies.RecommendedPriceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.RecommendedPriceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.RecommendedPriceRequest, ...client.CallOption) *currencies.RecommendedPriceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.RecommendedPriceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.RecommendedPriceRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSettlementCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetSettlementCurrencies(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSupportedCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetSupportedCurrencies(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVatCurrencies provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) GetVatCurrencies(ctx context.Context, in *currencies.EmptyRequest, opts ...client.CallOption) (*currencies.CurrenciesList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.CurrenciesList
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) *currencies.CurrenciesList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.CurrenciesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.EmptyRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPaysuperCorrectionCorridor provides a mock function with given fields: ctx, in, opts
func (_m *CurrencyratesService) SetPaysuperCorrectionCorridor(ctx context.Context, in *currencies.CorrectionCorridor, opts ...client.CallOption) (*currencies.EmptyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *currencies.EmptyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *currencies.CorrectionCorridor, ...client.CallOption) *currencies.EmptyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*currencies.EmptyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *currencies.CorrectionCorridor, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
