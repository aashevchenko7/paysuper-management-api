package micro

import (
	"context"
	"github.com/ProtocolONE/go-core/v2/pkg/invoker"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/ProtocolONE/go-micro-plugins/wrapper/select/version"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/client/selector/static"
)

// Micro
type Micro struct {
	ctx context.Context
	cfg Config
	provider.LMT
}

// Client
func (m *Micro) Client(serviceVersion, fallback string) client.Client {
	options := []micro.Option{
		micro.Name(m.cfg.Name),
		micro.Version(m.cfg.Version),
	}

	if len(serviceVersion) > 0 {
		wrapper := version.NewClientWrapper(serviceVersion, fallback)
		options = append(options, micro.WrapClient(wrapper))
	}

	if m.cfg.Selector == "static" {
		options = append(options, micro.Selector(static.NewSelector()))
	}

	service := micro.NewService(options...)
	service.Init()

	return service.Client()
}

// Config
type Config struct {
	Debug                  bool `fallback:"shared.debug"`
	Name                   string
	Version                string `default:"latest"`
	BillingVersion         string `default:"latest"`
	BillingFallbackVersion string `default:"latest"`
	Selector               string
	Bind                   string
	invoker                *invoker.Invoker
}

// OnReload
func (c *Config) OnReload(callback func(ctx context.Context)) {
	c.invoker.OnReload(callback)
}

// Reload
func (c *Config) Reload(ctx context.Context) {
	c.invoker.Reload(ctx)
}

// New
func New(ctx context.Context, set provider.AwareSet, cfg *Config) *Micro {
	set.Logger = set.Logger.WithFields(logger.Fields{"service": Prefix, "service_name": cfg.Name})

	return &Micro{
		ctx: ctx,
		cfg: *cfg,
		LMT: &set,
	}
}
