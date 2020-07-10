package common

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"math"
	"sync"
)

const (
	MoneyDefaultPrecision = 2
)

type MoneyOptions struct {
	precision      int64
	logger         logger.Logger
	isPrecisionSet bool
}

type MoneyOption func(*MoneyOptions)

func MoneyPrecision(precision int64) MoneyOption {
	return func(opts *MoneyOptions) {
		opts.precision = precision
		opts.isPrecisionSet = true
	}
}

func MoneyLogger(logger logger.Logger) MoneyOption {
	return func(opts *MoneyOptions) {
		opts.logger = logger
	}
}

type Money struct {
	registry  map[string]*tools.Money
	precision int64
	logger    logger.Logger
	mx        sync.Mutex
}

func NewMoney(opts ...MoneyOption) *Money {
	options := &MoneyOptions{}

	for _, opt := range opts {
		opt(options)
	}

	money := &Money{
		registry:  make(map[string]*tools.Money),
		logger:    options.logger,
		precision: options.precision,
	}

	return money
}

func (m *Money) Round(key string, val float64) (float64, error) {
	inst, ok := m.registry[key]

	if !ok {
		m.mx.Lock()
		inst = tools.New()
		m.registry[key] = inst
		m.mx.Unlock()
	}

	rounded, err := inst.Round(val, m.precision)

	if err != nil {
		m.logger.Error(
			billingpb.ErrorUnableRound,
			logger.WithPrettyFields(logger.Fields{
				"err":                     err,
				billingpb.ErrorFieldValue: val,
			}),
		)
	}

	return math.Round(rounded*100) / 100, err
}
