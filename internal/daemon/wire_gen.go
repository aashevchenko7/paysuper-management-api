// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package daemon

import (
	"context"
	"github.com/ProtocolONE/go-core/config"
	"github.com/ProtocolONE/go-core/invoker"
	"github.com/ProtocolONE/go-core/logger"
	"github.com/ProtocolONE/go-core/metric"
	"github.com/ProtocolONE/go-core/provider"
	"github.com/ProtocolONE/go-core/tracing"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher"
	"github.com/paysuper/paysuper-management-api/internal/handlers"
	"github.com/paysuper/paysuper-management-api/internal/validators"
	"github.com/paysuper/paysuper-management-api/pkg/http"
	"github.com/paysuper/paysuper-management-api/pkg/micro"
)

// Injectors from injector.go:

func BuildHTTP(ctx context.Context, initial config.Initial, observer invoker.Observer) (*http.HTTP, func(), error) {
	configurator, cleanup, err := config.Provider(initial, observer)
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	awareSet := provider.AwareSet{
		Logger: zap,
		Metric: scope,
		Tracer: tracer,
	}
	microConfig, cleanup8, err := micro.Cfg(configurator)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	microMicro, cleanup9, err := micro.Provider(ctx, awareSet, microConfig)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	services := dispatcher.ProviderServices(microMicro)
	validatorSet, cleanup10, err := validators.Provider(services)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	validate, cleanup11, err := dispatcher.ProviderValidators(validatorSet)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commonConfig, cleanup12, err := dispatcher.ProviderGlobalCfg(configurator)
	if err != nil {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commonHandlers, cleanup13, err := handlers.ProviderHandlers(initial, services, validate, awareSet, commonConfig)
	if err != nil {
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	jwtVerifier := dispatcher.ProviderJwtVerifier(commonConfig)
	appSet := dispatcher.AppSet{
		Handlers:    commonHandlers,
		Services:    services,
		JwtVerifier: jwtVerifier,
	}
	dispatcherConfig, cleanup14, err := dispatcher.ProviderCfg(configurator)
	if err != nil {
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dispatcherDispatcher, cleanup15, err := dispatcher.ProviderDispatcher(ctx, awareSet, appSet, dispatcherConfig, commonConfig)
	if err != nil {
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpConfig, cleanup16, err := http.Cfg(configurator)
	if err != nil {
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpHTTP, cleanup17, err := http.Provider(ctx, awareSet, dispatcherDispatcher, httpConfig)
	if err != nil {
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return httpHTTP, func() {
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildMicro(ctx context.Context, initial config.Initial, observer invoker.Observer) (*micro.Micro, func(), error) {
	configurator, cleanup, err := config.Provider(initial, observer)
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	awareSet := provider.AwareSet{
		Logger: zap,
		Metric: scope,
		Tracer: tracer,
	}
	microConfig, cleanup8, err := micro.Cfg(configurator)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	microMicro, cleanup9, err := micro.Provider(ctx, awareSet, microConfig)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return microMicro, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
