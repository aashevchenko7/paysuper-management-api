package dispatcher

import (
	"context"
	jwtverifier "github.com/ProtocolONE/authone-jwt-verifier-golang"
	geoip "github.com/ProtocolONE/geoip-service/pkg"
	"github.com/ProtocolONE/geoip-service/pkg/proto"
	"github.com/ProtocolONE/go-core/v2/pkg/config"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/google/wire"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/validators"
	"github.com/paysuper/paysuper-management-api/pkg/micro"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"github.com/paysuper/paysuper-proto/go/recurringpb"
	"github.com/paysuper/paysuper-proto/go/reporterpb"
	"github.com/paysuper/paysuper-proto/go/taxpb"
	"gopkg.in/go-playground/validator.v9"
)

// ProviderCfg
func ProviderCfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{
		WorkDir: cfg.WorkDir(),
		invoker: invoker.NewInvoker(),
	}
	e := cfg.UnmarshalKeyOnReload(common.UnmarshalKey, c)
	return c, func() {}, e
}

// ProviderGlobalCfg
func ProviderGlobalCfg(cfg config.Configurator) (*common.Config, func(), error) {
	c := &common.Config{}
	e := cfg.UnmarshalKey(common.UnmarshalGlobalConfigKey, c)
	return c, func() {}, e
}

// ProviderJwtVerifier
func ProviderJwtVerifier(cfg *common.Config) *jwtverifier.JwtVerifier {
	return jwtverifier.NewJwtVerifier(jwtverifier.Config{
		ClientID:     cfg.Auth1.ClientId,
		ClientSecret: cfg.Auth1.ClientSecret,
		Scopes:       []string{"openid", "offline"},
		RedirectURL:  cfg.Auth1.RedirectUrl,
		Issuer:       cfg.Auth1.Issuer,
	})
}

// ProviderServices
func ProviderServices(srv *micro.Micro, cfg *micro.Config) common.Services {
	return common.Services{
		Repository: recurringpb.NewRepositoryService(recurringpb.PayOneRepositoryServiceName, srv.Client("", "")),
		Geo:        proto.NewGeoIpService(geoip.ServiceName, srv.Client("", "")),
		Billing:    billingpb.NewBillingService(billingpb.ServiceName, srv.Client(cfg.BillingVersion, cfg.BillingFallbackVersion)),
		Tax:        taxpb.NewTaxService(taxpb.ServiceName, srv.Client("", "")),
		Reporter:   reporterpb.NewReporterService(reporterpb.ServiceName, srv.Client("", "")),
	}
}

// ProviderValidators
func ProviderValidators(v *validators.ValidatorSet) (validate *validator.Validate, _ func(), err error) {
	validate = validator.New()
	if err = validate.RegisterValidation("currency_price", v.ProductPriceValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("region_price", v.PriceRegionValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("phone", v.PhoneValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("uuid", v.UuidValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("zip_usa", v.ZipUsaValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("name", v.NameValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("position", v.PositionValidator); err != nil {
		return
	}
	validate.RegisterStructValidation(v.CompanyValidator, billingpb.UserProfileCompany{})
	validate.RegisterStructValidation(v.MerchantCompanyValidator, billingpb.MerchantCompanyInfo{})
	if err = validate.RegisterValidation("company_name", v.CompanyNameValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("swift", v.SwiftValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("city", v.CityValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("world_region", v.WorldRegionValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("tariff_region", v.TariffRegionValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("iban", v.IBANValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("locale", v.UserLocaleValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("date", v.DateValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("datetime", v.DateTimeValidator); err != nil {
		return
	}
	if err = validate.RegisterValidation("datetime_rfc3339", v.DateTimeRFC3339Validator); err != nil {
		return
	}
	return validate, func() {}, nil
}

// ProviderDispatcher
func ProviderDispatcher(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg *Config, globalCfg *common.Config, ms *micro.Micro) (*Dispatcher, func(), error) {
	d := New(ctx, set, appSet, cfg, globalCfg, ms)
	return d, func() {}, nil
}

var (
	// Dependencies: go-shared/provider.AwareSet, internal/*validators.ValidatorSet, pkg/micro.Micro, go-shared/config.Configurator, ProviderHandlers
	WireSet = wire.NewSet(
		ProviderDispatcher,
		ProviderServices,
		ProviderJwtVerifier,
		ProviderValidators,
		ProviderCfg,
		ProviderGlobalCfg,
		wire.Struct(new(AppSet), "*"),
	)
	// Dependencies: go-shared/provider.AwareSet, internal/*validators.ValidatorSet, common.Services, common.Handlers, go-shared/config.Configurator
	WireTestSet = wire.NewSet(
		ProviderDispatcher,
		ProviderJwtVerifier,
		ProviderValidators,
		ProviderCfg,
		ProviderGlobalCfg,
		wire.Struct(new(AppSet), "*"),
	)
)
